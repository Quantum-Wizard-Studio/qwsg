// Package drift classifies canonical comparison Change Records without making
// health, risk, policy, or remediation judgements.
package drift

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"sort"
	"strings"
	"time"

	"quantumwizard.hu/qwsg/internal/comparison"
)

const (
	SchemaName       = "qwsg.drift"
	SchemaVersion    = "1.0"
	EngineVersion    = "1.0"
	TaxonomyVersion  = "1.0"
	ChangeSchemaName = comparison.SchemaName
)

type Category string

const (
	IdentityDrift      Category = "identity"
	SoftwareDrift      Category = "software"
	HardwareDrift      Category = "hardware"
	PlatformDrift      Category = "platform"
	FilesystemDrift    Category = "filesystem"
	StorageDrift       Category = "storage"
	NetworkDrift       Category = "network"
	ServiceDrift       Category = "service"
	ConfigurationDrift Category = "configuration"
	SecurityDrift      Category = "security"
	CapabilityDrift    Category = "capability"
	EnvironmentDrift   Category = "environment"
	ExtensionDrift     Category = "extension"
)

type Classification string

const (
	PresenceAdded   Classification = "presence_added"
	PresenceRemoved Classification = "presence_removed"
	ValueModified   Classification = "value_modified"
	StateUnchanged  Classification = "state_unchanged"
)

// Scope is a privacy-preserving reference to the canonical compared value.
type Scope struct {
	Layer    string `json:"layer"`
	ObjectID string `json:"object_id"`
	Path     string `json:"path"`
}

// VersionInfo pins every semantic dependency used to produce a record.
type VersionInfo struct {
	DriftSchema  string `json:"drift_schema"`
	DriftEngine  string `json:"drift_engine"`
	Taxonomy     string `json:"taxonomy"`
	ChangeSchema string `json:"change_schema"`
}

// Record is the canonical public Drift Record 1.0 contract.
type Record struct {
	ID             string            `json:"id"`
	ChangeID       string            `json:"change_id"`
	Category       Category          `json:"category"`
	Scope          Scope             `json:"scope"`
	Classification Classification    `json:"classification"`
	Confidence     int               `json:"confidence_basis_points"`
	Timestamp      time.Time         `json:"timestamp"`
	Metadata       map[string]string `json:"metadata"`
	Versions       VersionInfo       `json:"versions"`
}

// Result is a deterministic envelope containing only canonical Drift Records.
type Result struct {
	SchemaName      string            `json:"schema_name"`
	SchemaVersion   string            `json:"schema_version"`
	EngineVersion   string            `json:"engine_version"`
	TaxonomyVersion string            `json:"taxonomy_version"`
	Records         []Record          `json:"records"`
	Metadata        map[string]string `json:"metadata"`
}

// Classify converts canonical Change Records to exactly one Drift Record each.
func Classify(changes []comparison.ChangeRecord) (Result, error) {
	result := Result{
		SchemaName: SchemaName, SchemaVersion: SchemaVersion,
		EngineVersion: EngineVersion, TaxonomyVersion: TaxonomyVersion,
		Records: []Record{},
		Metadata: map[string]string{
			"input_contract": ChangeSchemaName + "/" + comparison.SchemaVersion,
			"pipeline":       "canonical-change-record-to-drift-record",
		},
	}
	seenChanges := map[string]bool{}
	for _, change := range changes {
		if err := validateChange(change); err != nil {
			return Result{}, fmt.Errorf("invalid canonical change record: %w", err)
		}
		if seenChanges[change.ID] {
			return Result{}, fmt.Errorf("duplicate change id")
		}
		seenChanges[change.ID] = true
		record := classifyOne(change)
		result.Records = append(result.Records, record)
	}
	sort.Slice(result.Records, func(i, j int) bool {
		a, b := result.Records[i], result.Records[j]
		if a.Scope.Layer != b.Scope.Layer {
			return a.Scope.Layer < b.Scope.Layer
		}
		if a.Scope.ObjectID != b.Scope.ObjectID {
			return a.Scope.ObjectID < b.Scope.ObjectID
		}
		if a.Scope.Path != b.Scope.Path {
			return a.Scope.Path < b.Scope.Path
		}
		return a.ChangeID < b.ChangeID
	})
	if err := Validate(result); err != nil {
		return Result{}, err
	}
	return result, nil
}

func classifyOne(change comparison.ChangeRecord) Record {
	category, confidence, rule := classifyCategory(change)
	classification := classificationFor(change.Type)
	scope := Scope{Layer: change.Layer, ObjectID: change.ObjectID, Path: change.Path}
	record := Record{
		ChangeID: change.ID, Category: category, Scope: scope,
		Classification: classification, Confidence: confidence,
		Timestamp: change.ComparisonTimestamp.UTC(),
		Metadata: map[string]string{
			"classification_rule": rule,
			"source_change_type":  string(change.Type),
		},
		Versions: VersionInfo{
			DriftSchema: SchemaVersion, DriftEngine: EngineVersion,
			Taxonomy: TaxonomyVersion, ChangeSchema: comparison.SchemaVersion,
		},
	}
	record.ID = stableID(record)
	return record
}

func classifyCategory(change comparison.ChangeRecord) (Category, int, string) {
	layer := strings.ToLower(change.Layer)
	objectKind := strings.ToLower(change.Metadata["object_kind"])
	factName := strings.ToLower(change.Metadata["fact_name"])
	objectID := strings.ToLower(change.ObjectID)

	switch layer {
	case "host", "identity", "accounts":
		return IdentityDrift, 10000, "layer"
	case "hardware":
		return HardwareDrift, 10000, "layer"
	case "operating_system", "platform":
		return PlatformDrift, 10000, "layer"
	case "software", "packages", "applications":
		return SoftwareDrift, 10000, "layer"
	case "storage":
		if containsAny(objectKind+" "+factName+" "+objectID, "filesystem", "mount") {
			return FilesystemDrift, 9500, "storage-subtype"
		}
		return StorageDrift, 10000, "layer"
	case "filesystem":
		return FilesystemDrift, 10000, "layer"
	case "network":
		return NetworkDrift, 10000, "layer"
	case "services", "service":
		return ServiceDrift, 10000, "layer"
	case "configuration", "config":
		return ConfigurationDrift, 10000, "layer"
	case "security":
		return SecurityDrift, 10000, "layer"
	case "capability", "capabilities":
		return CapabilityDrift, 10000, "layer"
	case "environment", "metadata":
		return EnvironmentDrift, 10000, "layer"
	default:
		return ExtensionDrift, 5000, "unregistered-layer"
	}
}

func classificationFor(changeType comparison.ChangeType) Classification {
	switch changeType {
	case comparison.Added:
		return PresenceAdded
	case comparison.Removed:
		return PresenceRemoved
	case comparison.Modified:
		return ValueModified
	default:
		return StateUnchanged
	}
}

func validateChange(change comparison.ChangeRecord) error {
	if change.ID == "" || change.Layer == "" || change.ObjectID == "" || change.Path == "" ||
		change.ComparisonTimestamp.IsZero() || change.Metadata == nil {
		return fmt.Errorf("incomplete record")
	}
	if !strings.HasPrefix(change.Path, "/layers/") {
		return fmt.Errorf("non-canonical path")
	}
	if _, err := json.Marshal(change.Previous); err != nil {
		return fmt.Errorf("invalid previous value: %w", err)
	}
	if _, err := json.Marshal(change.Current); err != nil {
		return fmt.Errorf("invalid current value: %w", err)
	}
	switch change.Type {
	case comparison.Added:
		if change.Previous != nil || change.Current == nil {
			return fmt.Errorf("invalid added record")
		}
	case comparison.Removed:
		if change.Previous == nil || change.Current != nil {
			return fmt.Errorf("invalid removed record")
		}
	case comparison.Modified:
		if change.Previous == nil || change.Current == nil || equalValues(change.Previous, change.Current) {
			return fmt.Errorf("invalid modified record")
		}
	case comparison.Unchanged:
		if change.Previous == nil || change.Current == nil || !equalValues(change.Previous, change.Current) {
			return fmt.Errorf("invalid unchanged record")
		}
	default:
		return fmt.Errorf("unsupported change type")
	}
	return nil
}

// Validate rejects unsupported or internally inconsistent Drift contracts.
func Validate(result Result) error {
	if result.SchemaName != SchemaName || result.SchemaVersion != SchemaVersion ||
		result.EngineVersion != EngineVersion || result.TaxonomyVersion != TaxonomyVersion {
		return fmt.Errorf("unsupported drift contract")
	}
	if result.Records == nil || result.Metadata["input_contract"] != ChangeSchemaName+"/"+comparison.SchemaVersion {
		return fmt.Errorf("incomplete drift envelope")
	}
	if len(result.Metadata) != 2 ||
		result.Metadata["pipeline"] != "canonical-change-record-to-drift-record" {
		return fmt.Errorf("invalid drift envelope metadata")
	}
	seenDrift, seenChange := map[string]bool{}, map[string]bool{}
	lastKey := ""
	for _, record := range result.Records {
		if record.ID == "" || record.ChangeID == "" || seenDrift[record.ID] || seenChange[record.ChangeID] ||
			record.Scope.Layer == "" || record.Scope.ObjectID == "" || record.Scope.Path == "" ||
			record.Timestamp.IsZero() || record.Metadata == nil ||
			record.Confidence < 0 || record.Confidence > 10000 ||
			!validCategory(record.Category) || !validClassification(record.Classification) ||
			record.Versions != (VersionInfo{SchemaVersion, EngineVersion, TaxonomyVersion, comparison.SchemaVersion}) {
			return fmt.Errorf("invalid drift record")
		}
		if len(record.Metadata) != 2 || record.Metadata["classification_rule"] == "" ||
			record.Metadata["source_change_type"] == "" {
			return fmt.Errorf("invalid drift metadata")
		}
		if !validDerivation(record) {
			return fmt.Errorf("invalid drift derivation")
		}
		if record.ID != stableID(record) {
			return fmt.Errorf("invalid drift id")
		}
		seenDrift[record.ID], seenChange[record.ChangeID] = true, true
		key := record.Scope.Layer + "\x00" + record.Scope.ObjectID + "\x00" + record.Scope.Path + "\x00" + record.ChangeID
		if lastKey != "" && key <= lastKey {
			return fmt.Errorf("drift records are not uniquely ordered")
		}
		lastKey = key
	}
	return nil
}

func validCategory(category Category) bool {
	switch category {
	case IdentityDrift, SoftwareDrift, HardwareDrift, PlatformDrift, FilesystemDrift,
		StorageDrift, NetworkDrift, ServiceDrift, ConfigurationDrift, SecurityDrift,
		CapabilityDrift, EnvironmentDrift, ExtensionDrift:
		return true
	default:
		return false
	}
}

func validClassification(classification Classification) bool {
	switch classification {
	case PresenceAdded, PresenceRemoved, ValueModified, StateUnchanged:
		return true
	default:
		return false
	}
}

func validDerivation(record Record) bool {
	expectedClassification := map[string]Classification{
		string(comparison.Added):     PresenceAdded,
		string(comparison.Removed):   PresenceRemoved,
		string(comparison.Modified):  ValueModified,
		string(comparison.Unchanged): StateUnchanged,
	}
	if expectedClassification[record.Metadata["source_change_type"]] != record.Classification {
		return false
	}
	switch record.Metadata["classification_rule"] {
	case "layer":
		return record.Confidence == 10000 && record.Category != ExtensionDrift
	case "storage-subtype":
		return record.Confidence == 9500 && record.Category == FilesystemDrift
	case "unregistered-layer":
		return record.Confidence == 5000 && record.Category == ExtensionDrift
	default:
		return false
	}
}

func stableID(record Record) string {
	hash := sha256.New()
	hash.Write([]byte("qwsg.drift/" + EngineVersion + "/record"))
	for _, part := range []string{
		record.ChangeID, string(record.Category), record.Scope.Layer, record.Scope.ObjectID,
		record.Scope.Path, string(record.Classification), fmt.Sprintf("%d", record.Confidence),
		record.Metadata["classification_rule"],
	} {
		hash.Write([]byte{0})
		hash.Write([]byte(part))
	}
	return hex.EncodeToString(hash.Sum(nil))
}

func equalValues(a, b *comparison.TypedValue) bool {
	left, leftErr := json.Marshal(a)
	right, rightErr := json.Marshal(b)
	return leftErr == nil && rightErr == nil && string(left) == string(right)
}

func containsAny(value string, tokens ...string) bool {
	for _, token := range tokens {
		if strings.Contains(value, token) {
			return true
		}
	}
	return false
}
