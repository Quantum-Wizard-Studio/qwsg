package collector

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"sort"
	"sync"
	"time"

	"quantumwizard.hu/qwsg/internal/inventory"
)

const ContractVersion = "1.0"

var (
	ErrDuplicateCollector = errors.New("collector already registered")
	ErrInvalidDescriptor  = errors.New("invalid collector descriptor")
)

type Capability struct {
	ID          string
	Description string
}

type Descriptor struct {
	Name                   string
	Version                string
	ContractVersion        string
	InventoryCompatibility string
	Capability             Capability
	SupportedPlatforms     []string
	PrivilegeClass         string
	Timeout                time.Duration
	MaxTimeout             time.Duration
	OutputLimitBytes       int64
	Dependencies           []string
	SensitivityClasses     []string
}

type Request struct {
	RequestID  string
	Capability string
	Deadline   time.Time
}

type Availability struct {
	Available  bool
	Code       string
	MessageKey string
}

type Warning struct {
	Code       string `json:"code"`
	MessageKey string `json:"message_key"`
}

type Result struct {
	CollectorName      string
	Version            string
	Capability         string
	SupportedPlatforms []string
	ExecutionTimeMS    int64
	Timestamp          time.Time
	HealthStatus       inventory.CategoryStatus
	Warnings           []Warning
	Errors             []inventory.InventoryError
	CollectedData      inventory.Category
	Metadata           map[string]string
}

type Collector interface {
	Descriptor() Descriptor
	Availability(context.Context) Availability
	Collect(context.Context, Request) Result
}

type Registry struct {
	mu         sync.RWMutex
	collectors map[string]Collector
}

func NewRegistry() *Registry { return &Registry{collectors: make(map[string]Collector)} }

func (r *Registry) Register(c Collector) error {
	if c == nil {
		return fmt.Errorf("%w: nil collector", ErrInvalidDescriptor)
	}
	d := c.Descriptor()
	if err := validateDescriptor(d); err != nil {
		return err
	}
	r.mu.Lock()
	defer r.mu.Unlock()
	if _, exists := r.collectors[d.Name]; exists {
		return fmt.Errorf("%w: %s", ErrDuplicateCollector, d.Name)
	}
	r.collectors[d.Name] = c
	return nil
}

func (r *Registry) Names() []string {
	r.mu.RLock()
	defer r.mu.RUnlock()
	names := make([]string, 0, len(r.collectors))
	for name := range r.collectors {
		names = append(names, name)
	}
	sort.Strings(names)
	return names
}

func (r *Registry) Execute(ctx context.Context, requestID string) []Result {
	collectors := plan(r.snapshot())
	results := make([]Result, 0, len(collectors))
	completed := make(map[string]Result, len(collectors))
	for _, c := range collectors {
		d := c.Descriptor()
		if dependency := unavailableDependency(d.Dependencies, completed); dependency != "" {
			result := skippedResult(d, inventory.Unavailable, "dependency_unavailable", "collector.dependency_unavailable")
			result.Metadata["dependency"] = dependency
			results = append(results, result)
			completed[d.Name] = result
			continue
		}
		result := r.executeOne(ctx, c, requestID)
		results = append(results, result)
		completed[d.Name] = result
	}
	return results
}

func (r *Registry) snapshot() []Collector {
	r.mu.RLock()
	defer r.mu.RUnlock()
	collectors := make([]Collector, 0, len(r.collectors))
	for _, c := range r.collectors {
		collectors = append(collectors, c)
	}
	sort.Slice(collectors, func(i, j int) bool {
		return collectors[i].Descriptor().Name < collectors[j].Descriptor().Name
	})
	return collectors
}

func (r *Registry) executeOne(parent context.Context, c Collector, requestID string) Result {
	d := c.Descriptor()
	availability, panicked := safeAvailability(parent, c)
	if panicked {
		return skippedResult(d, inventory.Error, "availability_check_panic", "collector.availability_check_panic")
	}
	if !availability.Available {
		status := inventory.Unavailable
		if availability.Code == "unsupported" {
			status = inventory.Unsupported
		}
		return skippedResult(d, status, availability.Code, availability.MessageKey)
	}

	timeout := d.Timeout
	if timeout <= 0 || timeout > d.MaxTimeout {
		timeout = d.MaxTimeout
	}
	ctx, cancel := context.WithTimeout(parent, timeout)
	defer cancel()
	request := Request{RequestID: requestID, Capability: d.Capability.ID}
	if deadline, ok := ctx.Deadline(); ok {
		request.Deadline = deadline
	}
	resultChannel := make(chan Result, 1)
	go func() {
		defer func() {
			if recovered := recover(); recovered != nil {
				resultChannel <- skippedResult(d, inventory.Error, "collector_panic", "collector.panic")
			}
		}()
		resultChannel <- c.Collect(ctx, request)
	}()

	select {
	case result := <-resultChannel:
		result = normalizeResult(d, result)
		encoded, err := json.Marshal(result.CollectedData)
		if err != nil || int64(len(encoded)) > d.OutputLimitBytes {
			return skippedResult(d, inventory.Error, "resource_limit", "collector.output_limit")
		}
		return result
	case <-ctx.Done():
		status, code, key := inventory.Cancelled, "cancelled", "collector.cancelled"
		if errors.Is(ctx.Err(), context.DeadlineExceeded) {
			status, code, key = inventory.Timeout, "timeout", "collector.timeout"
		}
		return skippedResult(d, status, code, key)
	}
}

func plan(collectors []Collector) []Collector {
	remaining := make(map[string]Collector, len(collectors))
	for _, c := range collectors {
		remaining[c.Descriptor().Name] = c
	}
	planned := make([]Collector, 0, len(collectors))
	resolved := make(map[string]bool, len(collectors))
	for len(remaining) > 0 {
		progress := false
		for _, c := range collectors {
			d := c.Descriptor()
			if _, pending := remaining[d.Name]; !pending {
				continue
			}
			ready := true
			for _, dependency := range d.Dependencies {
				if _, registered := remaining[dependency]; registered && !resolved[dependency] {
					ready = false
					break
				}
			}
			if ready {
				planned = append(planned, c)
				resolved[d.Name] = true
				delete(remaining, d.Name)
				progress = true
			}
		}
		if progress {
			continue
		}
		// Cycles remain deterministic and fail dependency checks at execution.
		for _, c := range collectors {
			if _, pending := remaining[c.Descriptor().Name]; pending {
				planned = append(planned, c)
				delete(remaining, c.Descriptor().Name)
			}
		}
	}
	return planned
}

func safeAvailability(ctx context.Context, c Collector) (availability Availability, panicked bool) {
	defer func() {
		if recover() != nil {
			panicked = true
		}
	}()
	return c.Availability(ctx), false
}

func validateDescriptor(d Descriptor) error {
	if d.Name == "" || d.Version == "" || d.ContractVersion != ContractVersion || d.InventoryCompatibility == "" || d.Capability.ID == "" || len(d.SupportedPlatforms) == 0 || d.PrivilegeClass == "" || d.Timeout <= 0 || d.MaxTimeout <= 0 || d.Timeout > d.MaxTimeout || d.OutputLimitBytes <= 0 {
		return fmt.Errorf("%w: %s", ErrInvalidDescriptor, d.Name)
	}
	for _, dependency := range d.Dependencies {
		if dependency == "" || dependency == d.Name {
			return fmt.Errorf("%w: invalid dependency for %s", ErrInvalidDescriptor, d.Name)
		}
	}
	return nil
}

func unavailableDependency(dependencies []string, completed map[string]Result) string {
	for _, dependency := range dependencies {
		result, ok := completed[dependency]
		if !ok || result.HealthStatus != inventory.Available {
			return dependency
		}
	}
	return ""
}

func normalizeResult(d Descriptor, result Result) Result {
	result.CollectorName = d.Name
	result.Version = d.Version
	result.Capability = d.Capability.ID
	result.SupportedPlatforms = append([]string(nil), d.SupportedPlatforms...)
	if result.Timestamp.IsZero() {
		result.Timestamp = time.Now().UTC()
	}
	if result.Warnings == nil {
		result.Warnings = []Warning{}
	}
	if result.Errors == nil {
		result.Errors = []inventory.InventoryError{}
	}
	if result.Metadata == nil {
		result.Metadata = map[string]string{}
	}
	result.CollectedData.CategoryID = d.Capability.ID
	result.CollectedData.CollectorID = d.Name
	result.CollectedData.ContractVersion = d.ContractVersion
	result.CollectedData.Status = result.HealthStatus
	return result
}

func skippedResult(d Descriptor, status inventory.CategoryStatus, code, key string) Result {
	now := time.Now().UTC()
	errors := []inventory.InventoryError{}
	if code != "" {
		errors = append(errors, inventory.InventoryError{Code: code, CategoryID: d.Capability.ID, Class: string(status), MessageKey: key, OccurredAt: now})
	}
	category := inventory.Category{CategoryID: d.Capability.ID, ContractVersion: d.ContractVersion, Status: status, ObservedAt: now, CompletedAt: now, FreshUntil: now, CollectorID: d.Name, PrivilegeUsed: d.PrivilegeClass, SourceSummary: []string{}, Items: []inventory.Item{}, Errors: errors, Redactions: []string{}}
	return Result{CollectorName: d.Name, Version: d.Version, Capability: d.Capability.ID, SupportedPlatforms: append([]string(nil), d.SupportedPlatforms...), Timestamp: now, HealthStatus: status, Warnings: []Warning{}, Errors: errors, CollectedData: category, Metadata: map[string]string{}}
}
