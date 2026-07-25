// Package health evaluates the current engineering condition represented by
// canonical Drift Records. It performs no comparison, drift classification,
// monitoring, policy, alerting, remediation, networking, or AI processing.
package health

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"sort"
	"time"

	"quantumwizard.hu/qwsg/internal/drift"
)

const (
	SchemaName      = "qwsg.health"
	SchemaVersion   = "1.0"
	EngineVersion   = "1.0"
	TaxonomyVersion = "1.0"
)

// Status is a deterministic engineering-condition classification, not a
// probability, policy decision, operational risk score, or alert severity.
type Status string

const (
	Healthy       Status = "healthy"
	Informational Status = "informational"
	Advisory      Status = "advisory"
	Warning       Status = "warning"
	Critical      Status = "critical"
	Unknown       Status = "unknown"
	Unsupported   Status = "unsupported"
)

// EvidenceState makes uncertainty explicit and keeps it separate from Status.
type EvidenceState string

const (
	EvidenceSufficient   EvidenceState = "sufficient"
	EvidenceInsufficient EvidenceState = "insufficient"
	EvidenceUnsupported  EvidenceState = "unsupported"
)

// Scope preserves the canonical privacy-safe Drift scope without examining the
// values represented by it.
type Scope struct {
	Layer    string `json:"layer"`
	ObjectID string `json:"object_id"`
	Path     string `json:"path"`
}

// VersionInfo pins every public semantic dependency of a Health Record.
type VersionInfo struct {
	HealthSchema   string `json:"health_schema"`
	HealthEngine   string `json:"health_engine"`
	HealthTaxonomy string `json:"health_taxonomy"`
	DriftSchema    string `json:"drift_schema"`
	DriftTaxonomy  string `json:"drift_taxonomy"`
}

// Record is the canonical public Health Record 1.0 contract.
type Record struct {
	ID                    string            `json:"id"`
	DriftID               string            `json:"drift_id"`
	ChangeID              string            `json:"change_id"`
	Category              drift.Category    `json:"category"`
	Status                Status            `json:"status"`
	EvidenceState         EvidenceState     `json:"evidence_state"`
	ConfidenceBasisPoints int               `json:"confidence_basis_points"`
	Reason                string            `json:"reason"`
	Scope                 Scope             `json:"scope"`
	EvidenceTimestamp     time.Time         `json:"evidence_timestamp"`
	Metadata              map[string]string `json:"metadata"`
	Versions              VersionInfo       `json:"versions"`
}

// Result is the deterministic Health evaluation envelope.
type Result struct {
	SchemaName      string            `json:"schema_name"`
	SchemaVersion   string            `json:"schema_version"`
	EngineVersion   string            `json:"engine_version"`
	TaxonomyVersion string            `json:"taxonomy_version"`
	OverallStatus   Status            `json:"overall_status"`
	EvidenceState   EvidenceState     `json:"evidence_state"`
	Records         []Record          `json:"records"`
	Metadata        map[string]string `json:"metadata"`
}

// Evaluate converts validated canonical Drift evidence to Health Records.
func Evaluate(input drift.Result) (Result, error) {
	if err := drift.Validate(input); err != nil {
		return Result{}, fmt.Errorf("invalid canonical drift evidence: %w", err)
	}
	result := Result{
		SchemaName: SchemaName, SchemaVersion: SchemaVersion,
		EngineVersion: EngineVersion, TaxonomyVersion: TaxonomyVersion,
		OverallStatus: Unknown, EvidenceState: EvidenceInsufficient,
		Records: []Record{},
		Metadata: map[string]string{
			"input_contract": drift.SchemaName + "/" + drift.SchemaVersion,
			"pipeline":       "canonical-drift-record-to-health-record",
		},
	}
	for _, evidence := range input.Records {
		result.Records = append(result.Records, evaluateOne(evidence))
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
		return a.DriftID < b.DriftID
	})
	if len(result.Records) > 0 {
		result.OverallStatus, result.EvidenceState = aggregate(result.Records)
	}
	if err := Validate(result); err != nil {
		return Result{}, err
	}
	return result, nil
}

func evaluateOne(evidence drift.Record) Record {
	status, state, confidence, reason := classify(evidence)
	record := Record{
		DriftID: evidence.ID, ChangeID: evidence.ChangeID,
		Category: evidence.Category, Status: status, EvidenceState: state,
		ConfidenceBasisPoints: confidence, Reason: reason,
		Scope: Scope{
			Layer: evidence.Scope.Layer, ObjectID: evidence.Scope.ObjectID,
			Path: evidence.Scope.Path,
		},
		EvidenceTimestamp: evidence.Timestamp.UTC(),
		Metadata: map[string]string{
			"source_classification": string(evidence.Classification),
		},
		Versions: VersionInfo{
			HealthSchema: SchemaVersion, HealthEngine: EngineVersion,
			HealthTaxonomy: TaxonomyVersion, DriftSchema: drift.SchemaVersion,
			DriftTaxonomy: drift.TaxonomyVersion,
		},
	}
	record.ID = stableID(record)
	return record
}

func classify(evidence drift.Record) (Status, EvidenceState, int, string) {
	if evidence.Category == drift.ExtensionDrift {
		return Unsupported, EvidenceUnsupported, 0, "unsupported_drift_category"
	}
	switch evidence.Classification {
	case drift.StateUnchanged:
		return Healthy, EvidenceSufficient, 10000, "canonical_state_unchanged"
	case drift.PresenceAdded:
		return Informational, EvidenceSufficient, 10000, "canonical_presence_added"
	case drift.ValueModified:
		return Advisory, EvidenceSufficient, 10000, "canonical_value_modified"
	case drift.PresenceRemoved:
		if evidence.Category == drift.SecurityDrift {
			return Critical, EvidenceSufficient, 10000, "canonical_security_presence_removed"
		}
		return Warning, EvidenceSufficient, 10000, "canonical_presence_removed"
	default:
		panic("validated drift classification is unreachable")
	}
}

func aggregate(records []Record) (Status, EvidenceState) {
	status, state := Healthy, EvidenceSufficient
	for _, record := range records {
		if statusPriority(record.Status) > statusPriority(status) {
			status = record.Status
		}
		if evidencePriority(record.EvidenceState) > evidencePriority(state) {
			state = record.EvidenceState
		}
	}
	return status, state
}

// Validate rejects unsupported or internally inconsistent Health contracts.
func Validate(result Result) error {
	if result.SchemaName != SchemaName || result.SchemaVersion != SchemaVersion ||
		result.EngineVersion != EngineVersion || result.TaxonomyVersion != TaxonomyVersion {
		return fmt.Errorf("unsupported health contract")
	}
	if result.Records == nil || result.Metadata == nil ||
		result.Metadata["input_contract"] != drift.SchemaName+"/"+drift.SchemaVersion ||
		len(result.Metadata) != 2 ||
		result.Metadata["pipeline"] != "canonical-drift-record-to-health-record" {
		return fmt.Errorf("invalid health envelope")
	}
	if !validStatus(result.OverallStatus) || !validEvidenceState(result.EvidenceState) {
		return fmt.Errorf("invalid health summary")
	}
	if len(result.Records) == 0 {
		if result.OverallStatus != Unknown || result.EvidenceState != EvidenceInsufficient {
			return fmt.Errorf("empty evidence must remain unknown and insufficient")
		}
		return nil
	}

	seenHealth, seenDrift, seenChange := map[string]bool{}, map[string]bool{}, map[string]bool{}
	lastKey := ""
	for _, record := range result.Records {
		if record.ID == "" || record.DriftID == "" || record.ChangeID == "" ||
			seenHealth[record.ID] || seenDrift[record.DriftID] || seenChange[record.ChangeID] ||
			record.Scope.Layer == "" || record.Scope.ObjectID == "" || record.Scope.Path == "" ||
			record.EvidenceTimestamp.IsZero() || record.Metadata == nil ||
			!validStatus(record.Status) || record.Status == Unknown ||
			!validEvidenceState(record.EvidenceState) ||
			record.ConfidenceBasisPoints < 0 || record.ConfidenceBasisPoints > 10000 ||
			record.Versions != expectedVersions() {
			return fmt.Errorf("invalid health record")
		}
		if len(record.Metadata) != 1 || record.Metadata["source_classification"] == "" ||
			!validDerivation(record) || record.ID != stableID(record) {
			return fmt.Errorf("invalid health derivation")
		}
		key := record.Scope.Layer + "\x00" + record.Scope.ObjectID + "\x00" +
			record.Scope.Path + "\x00" + record.DriftID
		if lastKey != "" && key <= lastKey {
			return fmt.Errorf("health records are not uniquely ordered")
		}
		lastKey = key
		seenHealth[record.ID], seenDrift[record.DriftID], seenChange[record.ChangeID] = true, true, true
	}
	status, state := aggregate(result.Records)
	if result.OverallStatus != status || result.EvidenceState != state {
		return fmt.Errorf("invalid health aggregation")
	}
	return nil
}

func validDerivation(record Record) bool {
	classification := drift.Classification(record.Metadata["source_classification"])
	switch {
	case record.Category == drift.ExtensionDrift:
		return record.Status == Unsupported && record.EvidenceState == EvidenceUnsupported &&
			record.ConfidenceBasisPoints == 0 && record.Reason == "unsupported_drift_category"
	case classification == drift.StateUnchanged:
		return record.Status == Healthy && sufficient(record, "canonical_state_unchanged")
	case classification == drift.PresenceAdded:
		return record.Status == Informational && sufficient(record, "canonical_presence_added")
	case classification == drift.ValueModified:
		return record.Status == Advisory && sufficient(record, "canonical_value_modified")
	case classification == drift.PresenceRemoved && record.Category == drift.SecurityDrift:
		return record.Status == Critical && sufficient(record, "canonical_security_presence_removed")
	case classification == drift.PresenceRemoved:
		return record.Status == Warning && sufficient(record, "canonical_presence_removed")
	default:
		return false
	}
}

func sufficient(record Record, reason string) bool {
	return record.EvidenceState == EvidenceSufficient &&
		record.ConfidenceBasisPoints == 10000 && record.Reason == reason
}

func expectedVersions() VersionInfo {
	return VersionInfo{
		HealthSchema: SchemaVersion, HealthEngine: EngineVersion,
		HealthTaxonomy: TaxonomyVersion, DriftSchema: drift.SchemaVersion,
		DriftTaxonomy: drift.TaxonomyVersion,
	}
}

func validStatus(status Status) bool {
	switch status {
	case Healthy, Informational, Advisory, Warning, Critical, Unknown, Unsupported:
		return true
	default:
		return false
	}
}

func validEvidenceState(state EvidenceState) bool {
	switch state {
	case EvidenceSufficient, EvidenceInsufficient, EvidenceUnsupported:
		return true
	default:
		return false
	}
}

func statusPriority(status Status) int {
	switch status {
	case Critical:
		return 7
	case Unknown:
		return 6
	case Unsupported:
		return 5
	case Warning:
		return 4
	case Advisory:
		return 3
	case Informational:
		return 2
	case Healthy:
		return 1
	default:
		return 0
	}
}

func evidencePriority(state EvidenceState) int {
	switch state {
	case EvidenceInsufficient:
		return 3
	case EvidenceUnsupported:
		return 2
	case EvidenceSufficient:
		return 1
	default:
		return 0
	}
}

func stableID(record Record) string {
	hash := sha256.New()
	hash.Write([]byte("qwsg.health/" + EngineVersion + "/record"))
	for _, part := range []string{
		record.DriftID, record.ChangeID, string(record.Category), string(record.Status),
		string(record.EvidenceState), fmt.Sprintf("%d", record.ConfidenceBasisPoints),
		record.Reason, record.Scope.Layer, record.Scope.ObjectID, record.Scope.Path,
		record.Metadata["source_classification"],
	} {
		hash.Write([]byte{0})
		hash.Write([]byte(part))
	}
	return hex.EncodeToString(hash.Sum(nil))
}

// MarshalCanonical validates and serializes a byte-stable public Health result.
func MarshalCanonical(result Result) ([]byte, error) {
	if err := Validate(result); err != nil {
		return nil, err
	}
	return json.Marshal(result)
}
