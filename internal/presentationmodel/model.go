// Package presentationmodel projects validated canonical QWSG records into a
// bounded, localization-ready operator overview. It owns no engineering or
// operational decision and performs no collection, execution, or persistence.
package presentationmodel

import (
	"quantumwizard.hu/qwsg/internal/alert"
	"quantumwizard.hu/qwsg/internal/command"
	"quantumwizard.hu/qwsg/internal/comparison"
	"quantumwizard.hu/qwsg/internal/drift"
	"quantumwizard.hu/qwsg/internal/health"
	"quantumwizard.hu/qwsg/internal/inventory"
	"quantumwizard.hu/qwsg/internal/policy"
	"quantumwizard.hu/qwsg/internal/report"
	"quantumwizard.hu/qwsg/internal/rule"
	"quantumwizard.hu/qwsg/internal/runtime"
	"quantumwizard.hu/qwsg/internal/runtimeservice"
	"time"
)

const (
	SchemaName          = "qwsg.operator-overview"
	SchemaVersion       = "1.2"
	LegacySchemaVersion = "1.0"
	LegacySchemaV11     = "1.1"
	ModelVersion        = "1.2"
	LegacyModelVersion  = "1.0"
	LegacyModelV11      = "1.1"
	InputSchema         = "qwsg.operator-projection-input"
	MaxSources          = 32
	MaxAttention        = 256
	MaxRecommendations  = 8
	MaxTokenLength      = 128
	MaxFreshFor         = 30 * 24 * time.Hour
)

type Condition string

const (
	Healthy     Condition = "healthy"
	Degraded    Condition = "degraded"
	Critical    Condition = "critical"
	Unknown     Condition = "unknown"
	Unavailable Condition = "unavailable"
)

type Attention string

const (
	AttentionNone    Attention = "none"
	AttentionReview  Attention = "review"
	AttentionUrgent  Attention = "urgent"
	AttentionUnknown Attention = "unknown"
)

type Guardian string

const (
	GuardianRunning     Guardian = "running"
	GuardianStarting    Guardian = "starting"
	GuardianStopping    Guardian = "stopping"
	GuardianStopped     Guardian = "stopped"
	GuardianFailed      Guardian = "failed"
	GuardianDegraded    Guardian = "degraded"
	GuardianUnavailable Guardian = "unavailable"
	GuardianNotObserved Guardian = "not_observed"
)

type Freshness string

const (
	FreshnessCurrent     Freshness = "current"
	FreshnessStale       Freshness = "stale"
	FreshnessInvalid     Freshness = "invalid"
	FreshnessNotObserved Freshness = "not_observed"
)

type Completeness string

const (
	CompletenessComplete    Completeness = "complete"
	CompletenessPartial     Completeness = "partial"
	CompletenessUnsupported Completeness = "unsupported"
	CompletenessMissing     Completeness = "missing"
	CompletenessInvalid     Completeness = "invalid"
	CompletenessNotObserved Completeness = "not_observed"
)

type RecommendationToken string

const (
	RecommendInspectAttention       RecommendationToken = "inspect_attention"
	RecommendReviewChanges          RecommendationToken = "review_changes"
	RecommendRunFreshCheck          RecommendationToken = "run_fresh_check"
	RecommendInspectEvidence        RecommendationToken = "inspect_evidence"
	RecommendInspectFailedOperation RecommendationToken = "inspect_failed_operation"
	RecommendVerifyGuardian         RecommendationToken = "verify_guardian_operation"
	RecommendNoAction               RecommendationToken = "no_action"
)

type CommandObservation struct {
	ObservedAt time.Time         `json:"observed_at"`
	Value      command.Execution `json:"value"`
}
type InventoryObservation struct {
	ObservedAt time.Time          `json:"observed_at"`
	Value      inventory.Snapshot `json:"value"`
}
type ComparisonObservation struct {
	ObservedAt time.Time         `json:"observed_at"`
	Value      comparison.Result `json:"value"`
}
type DriftObservation struct {
	ObservedAt time.Time    `json:"observed_at"`
	Value      drift.Result `json:"value"`
}
type HealthObservation struct {
	ObservedAt time.Time     `json:"observed_at"`
	Value      health.Result `json:"value"`
}
type RuleObservation struct {
	ObservedAt time.Time   `json:"observed_at"`
	Value      rule.Result `json:"value"`
}
type PolicyObservation struct {
	ObservedAt time.Time     `json:"observed_at"`
	Value      policy.Result `json:"value"`
}
type ReportObservation struct {
	ObservedAt time.Time     `json:"observed_at"`
	Value      report.Report `json:"value"`
}
type PolicyReportObservation struct {
	ObservedAt time.Time           `json:"observed_at"`
	Value      report.PolicyReport `json:"value"`
}
type RuntimeObservation struct {
	ObservedAt time.Time      `json:"observed_at"`
	Value      runtime.Result `json:"value"`
}
type ServiceStateObservation struct {
	ObservedAt time.Time            `json:"observed_at"`
	Value      runtimeservice.State `json:"value"`
}
type ServiceResultObservation struct {
	ObservedAt time.Time             `json:"observed_at"`
	Value      runtimeservice.Result `json:"value"`
}

type Input struct {
	SchemaName    string                    `json:"schema_name"`
	SchemaVersion string                    `json:"schema_version"`
	ObservedAt    time.Time                 `json:"observed_at"`
	FreshForNS    int64                     `json:"fresh_for_ns"`
	Command       *CommandObservation       `json:"command,omitempty"`
	Inventory     *InventoryObservation     `json:"inventory,omitempty"`
	Comparison    *ComparisonObservation    `json:"comparison,omitempty"`
	Drift         *DriftObservation         `json:"drift,omitempty"`
	Health        *HealthObservation        `json:"health,omitempty"`
	Rule          *RuleObservation          `json:"rule,omitempty"`
	Policy        *PolicyObservation        `json:"policy,omitempty"`
	Report        *ReportObservation        `json:"report,omitempty"`
	PolicyReport  *PolicyReportObservation  `json:"policy_report,omitempty"`
	Runtime       *RuntimeObservation       `json:"runtime,omitempty"`
	ServiceState  *ServiceStateObservation  `json:"service_state,omitempty"`
	ServiceResult *ServiceResultObservation `json:"service_result,omitempty"`
}

type SourceReference struct {
	Kind       string    `json:"kind"`
	Contract   string    `json:"contract"`
	Version    string    `json:"version"`
	RecordID   string    `json:"record_id"`
	ObservedAt time.Time `json:"observed_at"`
}

type Summary struct {
	Changes            int `json:"changes"`
	CriticalHealth     int `json:"critical_health"`
	WarningHealth      int `json:"warning_health"`
	ActiveAlerts       int `json:"active_alerts"`
	AcknowledgedAlerts int `json:"acknowledged_alerts"`
	SuppressedAlerts   int `json:"suppressed_alerts"`
	RecoveredAlerts    int `json:"recovered_alerts"`
	ExpiredAlerts      int `json:"expired_alerts"`
}

type ChangeSummary struct {
	Category string `json:"category"`
	Added    int    `json:"added"`
	Removed  int    `json:"removed"`
	Modified int    `json:"modified"`
}

type AttentionItem struct {
	Severity    Attention       `json:"severity"`
	TitleToken  string          `json:"title_token"`
	ReasonToken string          `json:"reason_token"`
	Source      SourceReference `json:"source"`
}

type Recommendation struct {
	Token RecommendationToken `json:"token"`
}

// AttentionSummary makes bounded reduction explicit without exposing source
// payloads. TotalCandidates counts attention-producing views before canonical
// Rule/Policy and identical operator-meaning correlation; the remaining fields
// partition that total. Full source traceability remains in Overview.Sources.
type AttentionSummary struct {
	TotalCandidates      int `json:"total_candidates"`
	Represented          int `json:"represented"`
	CorrelatedDuplicates int `json:"correlated_duplicates"`
	Omitted              int `json:"omitted"`
}

type Overview struct {
	SchemaName       string            `json:"schema_name"`
	SchemaVersion    string            `json:"schema_version"`
	ModelVersion     string            `json:"model_version"`
	ID               string            `json:"id"`
	ObservedAt       time.Time         `json:"observed_at"`
	Condition        Condition         `json:"condition"`
	Attention        Attention         `json:"attention"`
	Guardian         Guardian          `json:"guardian"`
	Freshness        Freshness         `json:"freshness"`
	Completeness     Completeness      `json:"completeness"`
	Summary          Summary           `json:"summary"`
	Changes          []ChangeSummary   `json:"changes"`
	AttentionItems   []AttentionItem   `json:"attention_items"`
	Recommendations  []Recommendation  `json:"recommendations"`
	Sources          []SourceReference `json:"sources"`
	AttentionSummary *AttentionSummary `json:"attention_summary,omitempty"`
}

func alertAttention(severity alert.Severity) Attention {
	if severity == alert.Critical || severity == alert.Emergency {
		return AttentionUrgent
	}
	return AttentionReview
}
