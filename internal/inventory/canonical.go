package inventory

import (
	"fmt"
	"reflect"
	"sort"
	"time"
)

const CanonicalSchemaName = "qwsg.inventory"
const ContractVersionForLayer = "1.0"

type CollectorExecution struct {
	CollectorName      string             `json:"collector_name"`
	Version            string             `json:"version"`
	Capability         string             `json:"capability"`
	SupportedPlatforms []string           `json:"supported_platforms"`
	ExecutionTimeMS    int64              `json:"execution_time_ms"`
	Timestamp          time.Time          `json:"timestamp"`
	Status             CategoryStatus     `json:"health_status"`
	Warnings           []InventoryWarning `json:"warnings"`
	Errors             []InventoryError   `json:"errors"`
	Metadata           map[string]string  `json:"metadata"`
}

type InventoryWarning struct {
	Code       string `json:"code"`
	MessageKey string `json:"message_key"`
}

type SystemInventory struct {
	SchemaName       string               `json:"schema_name"`
	SchemaVersion    string               `json:"schema_version"`
	Profile          string               `json:"profile"`
	SnapshotID       string               `json:"snapshot_id"`
	RequestID        string               `json:"request_id"`
	SubjectID        string               `json:"subject_id"`
	ObservedAt       time.Time            `json:"observed_at"`
	CompletedAt      time.Time            `json:"completed_at"`
	FreshUntil       time.Time            `json:"fresh_until"`
	DurationMS       int64                `json:"duration_ms"`
	Status           Status               `json:"status"`
	Producer         Producer             `json:"producer"`
	CollectorResults []CollectorExecution `json:"collector_results"`
	Layers           []Layer              `json:"layers"`
	Issues           []InventoryError     `json:"issues"`
	Redactions       []string             `json:"redactions"`
	Metadata         map[string]string    `json:"metadata"`
}

type Layer struct {
	LayerID         string            `json:"layer_id"`
	ContractVersion string            `json:"contract_version"`
	Status          CategoryStatus    `json:"status"`
	ObservedAt      time.Time         `json:"observed_at"`
	CompletedAt     time.Time         `json:"completed_at"`
	CollectorIDs    []string          `json:"collector_ids"`
	Resources       []Resource        `json:"resources"`
	Issues          []InventoryError  `json:"issues"`
	Redactions      []string          `json:"redactions"`
	Metadata        map[string]string `json:"metadata"`
}

type Resource struct {
	ResourceID     string                   `json:"resource_id"`
	Kind           string                   `json:"kind"`
	LayerID        string                   `json:"layer_id"`
	LifecycleState string                   `json:"lifecycle_state"`
	Facts          map[string]CanonicalFact `json:"facts"`
	Relationships  []Relationship           `json:"relationships"`
	Labels         map[string]string        `json:"labels"`
	ObservedAt     time.Time                `json:"observed_at"`
	CollectorID    string                   `json:"collector_id"`
	Metadata       map[string]string        `json:"metadata"`
}

type CanonicalFact struct {
	Value       any        `json:"value,omitempty"`
	ValueType   string     `json:"value_type"`
	Unit        string     `json:"unit,omitempty"`
	Quality     string     `json:"quality"`
	Sensitivity string     `json:"sensitivity"`
	ObservedAt  time.Time  `json:"observed_at"`
	Provenance  Provenance `json:"provenance"`
	ReasonCode  string     `json:"reason_code,omitempty"`
}

type Relationship struct {
	RelationshipType string `json:"relationship_type"`
	SourceResourceID string `json:"source_resource_id"`
	TargetResourceID string `json:"target_resource_id"`
}

func AssembleSystemInventory(categories []Category, executions []CollectorExecution, snapshotID, requestID, subjectID string, observedAt, completedAt, freshUntil time.Time, durationMS int64, producer Producer) SystemInventory {
	layersByID := make(map[string][]Category)
	for _, category := range categories {
		layerID := layerForCategory(category.CategoryID)
		layersByID[layerID] = append(layersByID[layerID], category)
	}
	layerIDs := make([]string, 0, len(layersByID))
	for layerID := range layersByID {
		layerIDs = append(layerIDs, layerID)
	}
	sort.Strings(layerIDs)
	layers := make([]Layer, 0, len(layerIDs))
	for _, layerID := range layerIDs {
		layers = append(layers, assembleLayer(layerID, layersByID[layerID]))
	}
	sort.Slice(executions, func(i, j int) bool { return executions[i].CollectorName < executions[j].CollectorName })
	return SystemInventory{
		SchemaName: CanonicalSchemaName, SchemaVersion: SchemaVersion, Profile: "canonical-system-inventory-v1",
		SnapshotID: snapshotID, RequestID: requestID, SubjectID: subjectID, ObservedAt: observedAt, CompletedAt: completedAt, FreshUntil: freshUntil, DurationMS: durationMS, Status: Aggregate(categories), Producer: producer,
		CollectorResults: executions, Layers: layers, Issues: []InventoryError{},
		Redactions: []string{"hostnames", "network_addresses", "hardware_addresses", "mount_paths", "device_names", "service_identities"},
		Metadata:   map[string]string{"projection": "inventory-1.0-additive"},
	}
}

func assembleLayer(layerID string, categories []Category) Layer {
	status := Available
	available := 0
	collectorIDs := make([]string, 0, len(categories))
	resources := []Resource{}
	issues := []InventoryError{}
	redactions := []string{}
	observedAt := categories[0].ObservedAt
	completedAt := categories[0].CompletedAt
	for _, category := range categories {
		if category.ObservedAt.Before(observedAt) {
			observedAt = category.ObservedAt
		}
		if category.CompletedAt.After(completedAt) {
			completedAt = category.CompletedAt
		}
		collectorIDs = append(collectorIDs, category.CollectorID)
		if category.Status == Available {
			available++
		}
		issues = append(issues, category.Errors...)
		redactions = append(redactions, category.Redactions...)
		for _, legacy := range category.Items {
			facts := make(map[string]CanonicalFact, len(legacy.Facts))
			observedAt := category.ObservedAt
			for name, fact := range legacy.Facts {
				if !fact.Provenance.ObservedAt.IsZero() {
					observedAt = fact.Provenance.ObservedAt
				}
				facts[name] = CanonicalFact{Value: fact.Value, ValueType: valueType(fact.Value), Unit: fact.Unit, Quality: fact.Quality, Sensitivity: fact.Sensitivity, ObservedAt: observedAt, Provenance: fact.Provenance, ReasonCode: fact.Reason}
			}
			resources = append(resources, Resource{ResourceID: category.CategoryID + ":" + legacy.ID, Kind: legacy.Kind, LayerID: layerID, LifecycleState: "observed", Facts: facts, Relationships: []Relationship{}, Labels: map[string]string{}, ObservedAt: observedAt, CollectorID: category.CollectorID, Metadata: map[string]string{"legacy_category_id": category.CategoryID}})
		}
	}
	if available == 0 {
		status = categories[0].Status
	} else if available != len(categories) {
		status = CategoryPartial
	}
	sort.Strings(collectorIDs)
	sort.Slice(resources, func(i, j int) bool { return resources[i].ResourceID < resources[j].ResourceID })
	sort.Strings(redactions)
	return Layer{LayerID: layerID, ContractVersion: ContractVersionForLayer, Status: status, ObservedAt: observedAt, CompletedAt: completedAt, CollectorIDs: collectorIDs, Resources: resources, Issues: issues, Redactions: redactions, Metadata: map[string]string{}}
}

func layerForCategory(categoryID string) string {
	switch categoryID {
	case "host", "virtualization":
		return "host"
	case "cpu", "memory":
		return "hardware"
	case "os", "kernel":
		return "operating_system"
	case "storage", "filesystem":
		return "storage"
	case "network":
		return "network"
	case "services":
		return "services"
	case "components":
		return "applications"
	default:
		return "metadata"
	}
}

func valueType(value any) string {
	if value == nil {
		return "unknown"
	}
	switch reflect.TypeOf(value).Kind() {
	case reflect.Bool:
		return "boolean"
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64, reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return "integer"
	case reflect.Float32, reflect.Float64:
		return "number"
	case reflect.Array, reflect.Slice:
		return "array"
	case reflect.Map, reflect.Struct:
		return "object"
	default:
		return "string"
	}
}

func ValidateSystemInventory(system SystemInventory) error {
	if system.SchemaName != CanonicalSchemaName || system.SchemaVersion != SchemaVersion || system.SnapshotID == "" || system.RequestID == "" || system.SubjectID == "" || system.ObservedAt.After(system.CompletedAt) || system.CompletedAt.After(system.FreshUntil) || system.DurationMS < 0 || system.Producer.ContractVersion == "" || len(system.CollectorResults) == 0 || len(system.Layers) == 0 {
		return fmt.Errorf("invalid canonical envelope")
	}
	lastCollector := ""
	seenCollectors := map[string]bool{}
	for _, result := range system.CollectorResults {
		if result.CollectorName == "" || result.CollectorName < lastCollector || seenCollectors[result.CollectorName] || result.Capability == "" || result.ExecutionTimeMS < 0 {
			return fmt.Errorf("invalid collector result")
		}
		seenCollectors[result.CollectorName] = true
		lastCollector = result.CollectorName
	}
	seenLayers := map[string]bool{}
	seenResources := map[string]bool{}
	lastLayer := ""
	for _, layer := range system.Layers {
		if layer.LayerID == "" || layer.ContractVersion == "" || layer.ObservedAt.After(layer.CompletedAt) || seenLayers[layer.LayerID] || layer.LayerID < lastLayer {
			return fmt.Errorf("invalid layer identity or order: %s", layer.LayerID)
		}
		seenLayers[layer.LayerID] = true
		lastLayer = layer.LayerID
		lastResource := ""
		for _, resource := range layer.Resources {
			if resource.ResourceID == "" || resource.Kind == "" || resource.LayerID != layer.LayerID || resource.LifecycleState == "" || resource.CollectorID == "" || seenResources[resource.ResourceID] || resource.ResourceID < lastResource {
				return fmt.Errorf("invalid resource identity or order: %s", resource.ResourceID)
			}
			seenResources[resource.ResourceID] = true
			lastResource = resource.ResourceID
			for _, fact := range resource.Facts {
				if fact.Sensitivity == "secret_prohibited" || fact.ValueType == "" || fact.Quality == "" || fact.Provenance.SourceType == "" {
					return fmt.Errorf("invalid canonical fact")
				}
			}
		}
	}
	for _, layer := range system.Layers {
		for _, resource := range layer.Resources {
			for _, relationship := range resource.Relationships {
				if relationship.SourceResourceID != resource.ResourceID || !seenResources[relationship.TargetResourceID] {
					return fmt.Errorf("invalid relationship")
				}
			}
		}
	}
	return nil
}
