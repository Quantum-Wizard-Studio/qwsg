package inventory

import (
	"fmt"
	"time"
)

const SchemaVersion = "1.0"

type Status string

const (
	Complete Status = "complete"
	Partial  Status = "partial"
	Failed   Status = "failed"
)

type CategoryStatus string

const (
	Available        CategoryStatus = "available"
	CategoryPartial  CategoryStatus = "partial"
	Unavailable      CategoryStatus = "unavailable"
	Unsupported      CategoryStatus = "unsupported"
	PermissionDenied CategoryStatus = "permission_denied"
	Timeout          CategoryStatus = "timeout"
	Error            CategoryStatus = "error"
	Cancelled        CategoryStatus = "cancelled"
)

type Provenance struct {
	SourceType     string    `json:"source_type"`
	SourceLabel    string    `json:"source_label"`
	ObservedAt     time.Time `json:"observed_at"`
	Transformation string    `json:"transformation,omitempty"`
}
type Fact struct {
	Value       any        `json:"value,omitempty"`
	Unit        string     `json:"unit,omitempty"`
	Quality     string     `json:"quality"`
	Sensitivity string     `json:"sensitivity"`
	Provenance  Provenance `json:"provenance"`
	Reason      string     `json:"reason,omitempty"`
}
type Item struct {
	ID    string          `json:"item_id"`
	Kind  string          `json:"kind"`
	Facts map[string]Fact `json:"facts"`
}
type InventoryError struct {
	Code        string    `json:"code"`
	CategoryID  string    `json:"category_id,omitempty"`
	Class       string    `json:"class"`
	MessageKey  string    `json:"safe_message_key"`
	Retryable   bool      `json:"retryable"`
	OccurredAt  time.Time `json:"occurred_at"`
	SafeDetails string    `json:"details,omitempty"`
}
type Category struct {
	CategoryID      string           `json:"category_id"`
	ContractVersion string           `json:"contract_version"`
	Status          CategoryStatus   `json:"status"`
	ObservedAt      time.Time        `json:"observed_at"`
	CompletedAt     time.Time        `json:"completed_at"`
	FreshUntil      time.Time        `json:"fresh_until"`
	DurationMS      int64            `json:"duration_ms"`
	CollectorID     string           `json:"collector_id"`
	PrivilegeUsed   string           `json:"privilege_used"`
	SourceSummary   []string         `json:"source_summary"`
	Items           []Item           `json:"items"`
	Errors          []InventoryError `json:"errors"`
	Redactions      []string         `json:"redactions"`
}
type Producer struct {
	ToolVersion     string `json:"tool_version"`
	ContractVersion string `json:"contract_version"`
}
type Snapshot struct {
	SchemaVersion string           `json:"schema_version"`
	SnapshotID    string           `json:"snapshot_id"`
	RequestID     string           `json:"request_id"`
	InstanceID    string           `json:"instance_id"`
	ObservedAt    time.Time        `json:"observed_at"`
	CompletedAt   time.Time        `json:"completed_at"`
	FreshUntil    time.Time        `json:"fresh_until"`
	DurationMS    int64            `json:"duration_ms"`
	Status        Status           `json:"status"`
	Categories    []Category       `json:"categories"`
	Errors        []InventoryError `json:"errors"`
	Redactions    []string         `json:"redactions"`
	Producer      Producer         `json:"producer"`
	Canonical     SystemInventory  `json:"canonical_inventory"`
}

func Aggregate(cs []Category) Status {
	n := 0
	for _, c := range cs {
		if c.Status == Available {
			n++
		}
	}
	if n == len(cs) && n > 0 {
		return Complete
	}
	if n > 0 {
		return Partial
	}
	return Failed
}
func ExitCode(s Status) int {
	if s == Complete {
		return 0
	}
	if s == Partial {
		return 2
	}
	return 1
}
func Validate(s Snapshot) error {
	if s.SchemaVersion != SchemaVersion {
		return fmt.Errorf("unsupported schema version")
	}
	if len(s.Categories) == 0 || s.ObservedAt.After(s.CompletedAt) || s.CompletedAt.After(s.FreshUntil) {
		return fmt.Errorf("invalid snapshot envelope")
	}
	seen := map[string]bool{}
	for _, c := range s.Categories {
		if c.CategoryID == "" || seen[c.CategoryID] {
			return fmt.Errorf("invalid category identity")
		}
		seen[c.CategoryID] = true
	}
	if Aggregate(s.Categories) != s.Status {
		return fmt.Errorf("invalid aggregate status")
	}
	if s.Canonical.SchemaName != "" {
		if err := ValidateSystemInventory(s.Canonical); err != nil {
			return fmt.Errorf("invalid canonical inventory: %w", err)
		}
		if s.Canonical.SnapshotID != s.SnapshotID || s.Canonical.RequestID != s.RequestID || s.Canonical.SubjectID != s.InstanceID || !s.Canonical.ObservedAt.Equal(s.ObservedAt) || !s.Canonical.CompletedAt.Equal(s.CompletedAt) || !s.Canonical.FreshUntil.Equal(s.FreshUntil) || s.Canonical.DurationMS != s.DurationMS || s.Canonical.Status != s.Status || s.Canonical.Producer != s.Producer {
			return fmt.Errorf("canonical and compatibility envelopes differ")
		}
	}
	return nil
}
