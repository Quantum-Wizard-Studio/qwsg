// Package alert deterministically decides when canonical alerts exist.
// It is a pure decision engine: it performs no persistence, delivery,
// monitoring, remediation, presentation, networking, process, or AI work.
package alert

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"regexp"
	"sort"
	"strings"
	"time"

	"quantumwizard.hu/qwsg/internal/configuration"
	"quantumwizard.hu/qwsg/internal/health"
	"quantumwizard.hu/qwsg/internal/policy"
	"quantumwizard.hu/qwsg/internal/report"
	"quantumwizard.hu/qwsg/internal/rule"
	"quantumwizard.hu/qwsg/internal/scheduler"
)

const (
	SchemaVersion         = "1.0"
	EngineVersion         = "1.0"
	TaxonomyVersion       = "1.0"
	ModelVersion          = "1.0"
	InputSchema           = "qwsg.alert-evaluation-input"
	ResultSchema          = "qwsg.alert-evaluation-result"
	RecordSchema          = "qwsg.alert-record"
	StateSchema           = "qwsg.alert-state"
	AcknowledgementSchema = "qwsg.alert-acknowledgement"
	SuppressionSchema     = "qwsg.alert-suppression-window"
	MaxCandidates         = 4096
	MaxStateConditions    = 4096
	MaxControls           = 1024
	MaxReferences         = 64
	MaxStringLength       = 512
	MaxEvidenceTTL        = 30 * 24 * time.Hour
	MaxSuppression        = 366 * 24 * time.Hour
	EmergencyReminder     = 24 * time.Hour
)

type Severity string

const (
	Informational Severity = "informational"
	Warning       Severity = "warning"
	Critical      Severity = "critical"
	Emergency     Severity = "emergency"
)

type Category string

const (
	EngineeringCondition Category = "engineering_condition"
	RuleMatch            Category = "rule_match"
	PolicyGovernance     Category = "policy_governance"
	SchedulerOperation   Category = "scheduler_operation"
	ReportCompleteness   Category = "report_completeness"
	EvidenceLoss         Category = "evidence_loss"
)

type SourceKind string

const (
	HealthSource    SourceKind = "health"
	RuleSource      SourceKind = "rule"
	PolicySource    SourceKind = "policy"
	SchedulerSource SourceKind = "scheduler_event"
	ReportSource    SourceKind = "canonical_report"
)

type LifecycleState string

const (
	StateCandidate     LifecycleState = "candidate"
	StateActive        LifecycleState = "active"
	StateAcknowledged  LifecycleState = "acknowledged"
	StateSuppressed    LifecycleState = "suppressed"
	StateExpired       LifecycleState = "expired"
	StateResolved      LifecycleState = "resolved"
	StateIndeterminate LifecycleState = "indeterminate"
)

type EventKind string

const (
	EventEntered            EventKind = "entered"
	EventEscalated          EventKind = "escalated"
	EventDeescalated        EventKind = "deescalated"
	EventAcknowledged       EventKind = "acknowledged"
	EventSuppressionStarted EventKind = "suppression_started"
	EventSuppressionEnded   EventKind = "suppression_ended"
	EventMaintenanceEnded   EventKind = "maintenance_ended"
	EventReminder           EventKind = "reminder"
	EventExpired            EventKind = "expired"
	EventRecovered          EventKind = "recovered"
	EventIndeterminate      EventKind = "indeterminate"
)

type Decision string

const (
	DecisionAlert      Decision = "alert"
	DecisionSuppressed Decision = "suppressed"
	DecisionLifecycle  Decision = "lifecycle"
)

type SuppressionKind string

const (
	OperationalSuppression SuppressionKind = "operational"
	MaintenanceSuppression SuppressionKind = "maintenance"
)

type SourceReference struct {
	Kind               SourceKind `json:"kind"`
	SchemaName         string     `json:"schema_name"`
	SchemaVersion      string     `json:"schema_version"`
	RecordID           string     `json:"record_id"`
	Subject            string     `json:"subject"`
	EvidenceReferences []string   `json:"evidence_references"`
}

// Acknowledgement is immutable operator-awareness evidence. It changes no
// engineering condition and grants no remediation authority.
type Acknowledgement struct {
	SchemaName    string    `json:"schema_name"`
	SchemaVersion string    `json:"schema_version"`
	ID            string    `json:"id"`
	LifecycleID   string    `json:"lifecycle_id"`
	ActorID       string    `json:"actor_id"`
	AuthorityRef  string    `json:"authority_ref"`
	At            time.Time `json:"at"`
	NoteToken     string    `json:"note_token,omitempty"`
}

// SuppressionWindow is a bounded explicit decision input. Empty scope lists
// are wildcards. Maintenance is represented only as a suppression kind.
type SuppressionWindow struct {
	SchemaName        string          `json:"schema_name"`
	SchemaVersion     string          `json:"schema_version"`
	ID                string          `json:"id"`
	Kind              SuppressionKind `json:"kind"`
	ConditionKeys     []string        `json:"condition_keys"`
	Categories        []Category      `json:"categories"`
	Severities        []Severity      `json:"severities"`
	StartsAt          time.Time       `json:"starts_at"`
	EndsAt            time.Time       `json:"ends_at"`
	ActorID           string          `json:"actor_id"`
	AuthorityRef      string          `json:"authority_ref"`
	ReasonToken       string          `json:"reason_token"`
	SuppressEmergency bool            `json:"suppress_emergency"`
}

// Condition is retained lifecycle state. Resolved and expired generations are
// retained so recurrence receives a new lifecycle identity.
type Condition struct {
	ConditionKey      string          `json:"condition_key"`
	LifecycleID       string          `json:"lifecycle_id"`
	Generation        int             `json:"generation"`
	Category          Category        `json:"category"`
	Severity          Severity        `json:"severity"`
	LifecycleState    LifecycleState  `json:"lifecycle_state"`
	Source            SourceReference `json:"source"`
	ReasonToken       string          `json:"reason_token"`
	FirstObservedAt   time.Time       `json:"first_observed_at"`
	LastObservedAt    time.Time       `json:"last_observed_at"`
	LastAlertedAt     time.Time       `json:"last_alerted_at"`
	AcknowledgementID string          `json:"acknowledgement_id,omitempty"`
	SuppressionIDs    []string        `json:"suppression_ids"`
	ResolvedAt        time.Time       `json:"resolved_at,omitempty"`
	ExpiredAt         time.Time       `json:"expired_at,omitempty"`
}

type State struct {
	SchemaName      string      `json:"schema_name"`
	SchemaVersion   string      `json:"schema_version"`
	ID              string      `json:"id"`
	ConfigurationID string      `json:"configuration_id,omitempty"`
	Conditions      []Condition `json:"conditions"`
}

type Input struct {
	SchemaName       string                   `json:"schema_name"`
	SchemaVersion    string                   `json:"schema_version"`
	EvaluatedAt      time.Time                `json:"evaluated_at"`
	EvidenceTTLNS    int64                    `json:"evidence_ttl_ns"`
	Configuration    *configuration.Effective `json:"configuration,omitempty"`
	Health           *health.Result           `json:"health,omitempty"`
	Rules            *rule.Result             `json:"rules,omitempty"`
	Policies         *policy.Result           `json:"policies,omitempty"`
	Scheduler        *scheduler.Evaluation    `json:"scheduler,omitempty"`
	Report           *report.Report           `json:"report,omitempty"`
	PolicyReport     *report.PolicyReport     `json:"policy_report,omitempty"`
	PreviousState    State                    `json:"previous_state"`
	Acknowledgements []Acknowledgement        `json:"acknowledgements"`
	Suppressions     []SuppressionWindow      `json:"suppressions"`
}

type VersionInfo struct {
	AlertSchema   string `json:"alert_schema"`
	AlertEngine   string `json:"alert_engine"`
	AlertTaxonomy string `json:"alert_taxonomy"`
	AlertModel    string `json:"alert_model"`
}

// Record is an immutable Canonical Alert Record 1.0. Records represent only
// meaningful lifecycle decisions; unchanged evidence updates state silently.
type Record struct {
	SchemaName        string          `json:"schema_name"`
	SchemaVersion     string          `json:"schema_version"`
	ID                string          `json:"id"`
	ConditionKey      string          `json:"condition_key"`
	LifecycleID       string          `json:"lifecycle_id"`
	Generation        int             `json:"generation"`
	Event             EventKind       `json:"event"`
	Decision          Decision        `json:"decision"`
	LifecycleState    LifecycleState  `json:"lifecycle_state"`
	Severity          Severity        `json:"severity"`
	PreviousSeverity  Severity        `json:"previous_severity,omitempty"`
	Category          Category        `json:"category"`
	Source            SourceReference `json:"source"`
	EventTime         time.Time       `json:"event_time"`
	ObservationTime   time.Time       `json:"observation_time"`
	EvaluationTime    time.Time       `json:"evaluation_time"`
	AcknowledgementID string          `json:"acknowledgement_id,omitempty"`
	SuppressionIDs    []string        `json:"suppression_ids"`
	ExpirationTime    time.Time       `json:"expiration_time,omitempty"`
	RecoveryTime      time.Time       `json:"recovery_time,omitempty"`
	ReasonToken       string          `json:"reason_token"`
	Versions          VersionInfo     `json:"versions"`
}

type Result struct {
	SchemaName      string    `json:"schema_name"`
	SchemaVersion   string    `json:"schema_version"`
	EngineVersion   string    `json:"engine_version"`
	TaxonomyVersion string    `json:"taxonomy_version"`
	ModelVersion    string    `json:"model_version"`
	ID              string    `json:"id"`
	EvaluatedAt     time.Time `json:"evaluated_at"`
	Records         []Record  `json:"records"`
	NextState       State     `json:"next_state"`
}

type candidate struct {
	key           string
	category      Category
	severity      Severity
	reason        string
	observed      time.Time
	source        SourceReference
	indeterminate bool
}

var tokenPattern = regexp.MustCompile(`^[a-z0-9]+([._:-][a-z0-9]+)*$`)

func NewState(configurationID string) State {
	state := State{SchemaName: StateSchema, SchemaVersion: SchemaVersion, ConfigurationID: configurationID, Conditions: []Condition{}}
	state.ID = stateIdentity(state)
	return state
}

func NewAcknowledgement(lifecycleID, actorID, authorityRef string, at time.Time, noteToken string) (Acknowledgement, error) {
	value := Acknowledgement{SchemaName: AcknowledgementSchema, SchemaVersion: SchemaVersion, LifecycleID: lifecycleID, ActorID: actorID, AuthorityRef: authorityRef, At: at.UTC(), NoteToken: noteToken}
	value.ID = acknowledgementIdentity(value)
	return value, validateAcknowledgement(value)
}

func NewSuppressionWindow(kind SuppressionKind, conditionKeys []string, categories []Category, severities []Severity, startsAt, endsAt time.Time, actorID, authorityRef, reasonToken string, suppressEmergency bool) (SuppressionWindow, error) {
	value := SuppressionWindow{SchemaName: SuppressionSchema, SchemaVersion: SchemaVersion, Kind: kind,
		ConditionKeys: append([]string{}, conditionKeys...), Categories: append([]Category{}, categories...), Severities: append([]Severity{}, severities...),
		StartsAt: startsAt.UTC(), EndsAt: endsAt.UTC(), ActorID: actorID, AuthorityRef: authorityRef, ReasonToken: reasonToken, SuppressEmergency: suppressEmergency}
	sort.Strings(value.ConditionKeys)
	sort.Slice(value.Categories, func(i, j int) bool { return value.Categories[i] < value.Categories[j] })
	sort.Slice(value.Severities, func(i, j int) bool { return value.Severities[i] < value.Severities[j] })
	value.ID = suppressionIdentity(value)
	return value, validateSuppression(value)
}

func Evaluate(input Input) (Result, error) {
	if err := ValidateInput(input); err != nil {
		return Result{}, err
	}
	candidates, resolutions, err := buildCandidates(input)
	if err != nil {
		return Result{}, err
	}
	previous := make(map[string]Condition, len(input.PreviousState.Conditions))
	for _, item := range input.PreviousState.Conditions {
		if old, ok := previous[item.ConditionKey]; !ok || item.Generation > old.Generation {
			previous[item.ConditionKey] = cloneCondition(item)
		}
	}
	next := cloneState(input.PreviousState)
	byLifecycle := map[string]Acknowledgement{}
	for _, ack := range input.Acknowledgements {
		byLifecycle[ack.LifecycleID] = ack
	}
	records := []Record{}
	seen := map[string]bool{}
	for _, c := range candidates {
		seen[c.key] = true
		prior, exists := previous[c.key]
		expires := c.observed.Add(time.Duration(input.EvidenceTTLNS))
		if !input.EvaluatedAt.Before(expires) {
			if !exists || !terminalState(prior.LifecycleState) {
				generation := 1
				condition := newCondition(c, generation, input.EvaluatedAt)
				if exists {
					condition = cloneCondition(prior)
					condition.Source = c.source
					condition.LastObservedAt = c.observed
					condition.ReasonToken = "source_evidence_expired"
				} else {
					condition.ReasonToken = "source_evidence_expired"
				}
				condition.LifecycleState = StateExpired
				condition.ExpiredAt = expires
				condition.SuppressionIDs = []string{}
				records = append(records, makeRecord(condition, condition.Severity, EventExpired, input.EvaluatedAt, time.Time{}))
				condition.LastAlertedAt = input.EvaluatedAt
				upsertCondition(&next, condition)
			}
			continue
		}
		if !exists || prior.LifecycleState == StateResolved || prior.LifecycleState == StateExpired {
			generation := 1
			if exists {
				generation = prior.Generation + 1
			}
			condition := newCondition(c, generation, input.EvaluatedAt)
			condition.SuppressionIDs = matchingSuppressionIDs(input.Suppressions, condition, input.EvaluatedAt)
			if c.indeterminate {
				condition.LifecycleState = StateIndeterminate
			} else if len(condition.SuppressionIDs) > 0 {
				condition.LifecycleState = StateSuppressed
			} else {
				condition.LifecycleState = StateActive
			}
			event := EventEntered
			if c.indeterminate {
				event = EventIndeterminate
			}
			records = append(records, makeRecord(condition, Severity(""), event, input.EvaluatedAt, time.Time{}))
			condition.LastAlertedAt = input.EvaluatedAt
			upsertCondition(&next, condition)
			continue
		}
		condition := cloneCondition(prior)
		condition.Source, condition.ReasonToken, condition.LastObservedAt = c.source, c.reason, c.observed
		previousSeverity := condition.Severity
		condition.Category, condition.Severity = c.category, c.severity
		newSuppressions := matchingSuppressionIDs(input.Suppressions, condition, input.EvaluatedAt)
		event := EventKind("")
		if severityRank(condition.Severity) > severityRank(previousSeverity) {
			event = EventEscalated
		} else if severityRank(condition.Severity) < severityRank(previousSeverity) {
			event = EventDeescalated
		} else if len(prior.SuppressionIDs) == 0 && len(newSuppressions) > 0 {
			event = EventSuppressionStarted
		} else if len(prior.SuppressionIDs) > 0 && len(newSuppressions) == 0 {
			event = EventSuppressionEnded
			if endedMaintenance(prior.SuppressionIDs, input.Suppressions, input.EvaluatedAt) {
				event = EventMaintenanceEnded
			}
		}
		condition.SuppressionIDs = newSuppressions
		if len(newSuppressions) > 0 {
			condition.LifecycleState = StateSuppressed
		} else if ack, ok := byLifecycle[condition.LifecycleID]; ok {
			if condition.AcknowledgementID == "" {
				condition.AcknowledgementID = ack.ID
				if event == "" {
					event = EventAcknowledged
				}
			}
			condition.LifecycleState = StateAcknowledged
		} else if c.indeterminate {
			condition.LifecycleState = StateIndeterminate
			if event == "" && prior.LifecycleState != StateIndeterminate {
				event = EventIndeterminate
			}
		} else {
			condition.LifecycleState = StateActive
		}
		if event == "" && condition.Severity == Emergency && condition.AcknowledgementID == "" && len(condition.SuppressionIDs) == 0 && input.EvaluatedAt.Sub(condition.LastAlertedAt) >= EmergencyReminder {
			event = EventReminder
		}
		if event != "" {
			records = append(records, makeRecord(condition, previousSeverity, event, input.EvaluatedAt, time.Time{}))
			condition.LastAlertedAt = input.EvaluatedAt
		}
		upsertCondition(&next, condition)
	}
	for _, prior := range input.PreviousState.Conditions {
		if seen[prior.ConditionKey] || terminalState(prior.LifecycleState) {
			continue
		}
		condition := cloneCondition(prior)
		if source, ok := resolutions[prior.ConditionKey]; ok {
			condition.Source = source
			condition.LifecycleState = StateResolved
			condition.ResolvedAt = input.EvaluatedAt
			condition.SuppressionIDs = []string{}
			records = append(records, makeRecord(condition, prior.Severity, EventRecovered, input.EvaluatedAt, input.EvaluatedAt))
			condition.LastAlertedAt = input.EvaluatedAt
			upsertCondition(&next, condition)
			continue
		}
		expires := prior.LastObservedAt.Add(time.Duration(input.EvidenceTTLNS))
		if !input.EvaluatedAt.Before(expires) {
			condition.LifecycleState = StateExpired
			condition.ExpiredAt = expires
			condition.SuppressionIDs = []string{}
			records = append(records, makeRecord(condition, prior.Severity, EventExpired, input.EvaluatedAt, time.Time{}))
			condition.LastAlertedAt = input.EvaluatedAt
			upsertCondition(&next, condition)
		}
	}
	sortConditions(next.Conditions)
	if len(next.Conditions) > MaxStateConditions {
		return Result{}, fmt.Errorf("alert state condition limit exceeded")
	}
	next.ID = stateIdentity(next)
	sort.Slice(records, func(i, j int) bool {
		if records[i].ConditionKey != records[j].ConditionKey {
			return records[i].ConditionKey < records[j].ConditionKey
		}
		if records[i].Generation != records[j].Generation {
			return records[i].Generation < records[j].Generation
		}
		return records[i].ID < records[j].ID
	})
	result := Result{SchemaName: ResultSchema, SchemaVersion: SchemaVersion, EngineVersion: EngineVersion, TaxonomyVersion: TaxonomyVersion, ModelVersion: ModelVersion, EvaluatedAt: input.EvaluatedAt, Records: records, NextState: next}
	result.ID = resultIdentity(result)
	if err := ValidateResult(result); err != nil {
		return Result{}, err
	}
	return result, nil
}

func ValidateInput(input Input) error {
	if input.SchemaName != InputSchema || input.SchemaVersion != SchemaVersion || !validTime(input.EvaluatedAt) || input.EvidenceTTLNS <= 0 || time.Duration(input.EvidenceTTLNS) > MaxEvidenceTTL || input.Acknowledgements == nil || input.Suppressions == nil {
		return fmt.Errorf("invalid alert evaluation input envelope")
	}
	configurationID := ""
	if input.Configuration != nil {
		if err := configuration.ValidateEffective(*input.Configuration); err != nil {
			return fmt.Errorf("invalid effective configuration: %w", err)
		}
		configurationID = input.Configuration.ID
	}
	if err := ValidateState(input.PreviousState); err != nil {
		return fmt.Errorf("invalid previous alert state: %w", err)
	}
	if input.PreviousState.ConfigurationID != configurationID {
		return fmt.Errorf("alert state configuration identity mismatch")
	}
	if input.Health != nil {
		if err := health.Validate(*input.Health); err != nil {
			return fmt.Errorf("invalid health source: %w", err)
		}
	}
	if input.Rules != nil {
		if err := rule.Validate(*input.Rules); err != nil {
			return fmt.Errorf("invalid rule source: %w", err)
		}
	}
	if input.Policies != nil {
		if err := policy.Validate(*input.Policies); err != nil {
			return fmt.Errorf("invalid policy source: %w", err)
		}
	}
	if input.Scheduler != nil {
		if err := scheduler.ValidateEvaluation(*input.Scheduler); err != nil {
			return fmt.Errorf("invalid scheduler source: %w", err)
		}
	}
	if input.Report != nil {
		if err := report.Validate(*input.Report); err != nil {
			return fmt.Errorf("invalid report source: %w", err)
		}
	}
	if input.PolicyReport != nil {
		if err := report.ValidatePolicyReport(*input.PolicyReport); err != nil {
			return fmt.Errorf("invalid policy report source: %w", err)
		}
	}
	if input.Report != nil && input.PolicyReport != nil {
		return fmt.Errorf("ambiguous canonical report sources")
	}
	if len(input.Acknowledgements) > MaxControls || len(input.Suppressions) > MaxControls {
		return fmt.Errorf("alert control limit exceeded")
	}
	conditionsByLifecycle := make(map[string]Condition, len(input.PreviousState.Conditions))
	for _, condition := range input.PreviousState.Conditions {
		conditionsByLifecycle[condition.LifecycleID] = condition
	}
	acknowledgedLifecycle := map[string]bool{}
	last := ""
	for _, ack := range input.Acknowledgements {
		condition, exists := conditionsByLifecycle[ack.LifecycleID]
		if err := validateAcknowledgement(ack); err != nil || ack.ID <= last || ack.At.After(input.EvaluatedAt) ||
			!exists || terminalState(condition.LifecycleState) || ack.At.Before(condition.FirstObservedAt) || acknowledgedLifecycle[ack.LifecycleID] {
			return fmt.Errorf("invalid or unordered alert acknowledgement")
		}
		acknowledgedLifecycle[ack.LifecycleID] = true
		last = ack.ID
	}
	last = ""
	for _, window := range input.Suppressions {
		if err := validateSuppression(window); err != nil || window.ID <= last {
			return fmt.Errorf("invalid or unordered alert suppression")
		}
		last = window.ID
	}
	return nil
}

func ValidateState(state State) error {
	if state.SchemaName != StateSchema || state.SchemaVersion != SchemaVersion || state.Conditions == nil || len(state.Conditions) > MaxStateConditions {
		return fmt.Errorf("invalid alert state envelope")
	}
	last := ""
	seenLifecycle := map[string]bool{}
	for _, condition := range state.Conditions {
		key := condition.ConditionKey + "\x00" + fmt.Sprintf("%09d", condition.Generation)
		if key <= last || seenLifecycle[condition.LifecycleID] || !validCondition(condition) {
			return fmt.Errorf("invalid or unordered alert condition")
		}
		last, seenLifecycle[condition.LifecycleID] = key, true
	}
	if state.ID != stateIdentity(state) {
		return fmt.Errorf("invalid alert state identity")
	}
	return nil
}

func ValidateResult(result Result) error {
	if result.SchemaName != ResultSchema || result.SchemaVersion != SchemaVersion || result.EngineVersion != EngineVersion || result.TaxonomyVersion != TaxonomyVersion || result.ModelVersion != ModelVersion || !validTime(result.EvaluatedAt) || result.Records == nil {
		return fmt.Errorf("invalid alert result envelope")
	}
	if err := ValidateState(result.NextState); err != nil {
		return err
	}
	last := ""
	seen := map[string]bool{}
	for _, record := range result.Records {
		key := record.ConditionKey + "\x00" + fmt.Sprintf("%09d", record.Generation) + "\x00" + record.ID
		if key <= last || seen[record.ID] || !validRecord(record) || record.ID != recordIdentity(record) || record.EvaluationTime != result.EvaluatedAt {
			return fmt.Errorf("invalid or unordered alert record")
		}
		last, seen[record.ID] = key, true
	}
	if result.ID != resultIdentity(result) {
		return fmt.Errorf("invalid alert result identity")
	}
	return nil
}

// ValidateRecord validates one immutable Canonical Alert Record without
// evaluating alert existence or requiring its containing Result envelope.
func ValidateRecord(record Record) error {
	if !validRecord(record) || record.ID != recordIdentity(record) {
		return fmt.Errorf("invalid canonical alert record")
	}
	return nil
}

func MarshalCanonical(result Result) ([]byte, error) {
	if err := ValidateResult(result); err != nil {
		return nil, err
	}
	return json.Marshal(result)
}

func DecodeInput(data []byte) (Input, error) {
	var value Input
	if err := strictDecode(data, &value); err != nil {
		return Input{}, err
	}
	return value, ValidateInput(value)
}

func DecodeResult(data []byte) (Result, error) {
	var value Result
	if err := strictDecode(data, &value); err != nil {
		return Result{}, err
	}
	return value, ValidateResult(value)
}

func buildCandidates(input Input) ([]candidate, map[string]SourceReference, error) {
	healthByID := map[string]health.Record{}
	claimedHealth := map[string]bool{}
	claimedRules := map[string]bool{}
	if input.Health != nil {
		for _, r := range input.Health.Records {
			healthByID[r.ID] = r
		}
	}
	if input.Rules != nil {
		for _, r := range input.Rules.Records {
			for _, id := range r.EvidenceReferences {
				claimedHealth[id] = true
			}
		}
	}
	if input.Policies != nil {
		for _, r := range input.Policies.Records {
			claimedRules[r.RuleEvaluationID] = true
			for _, id := range r.EvidenceReferences {
				claimedHealth[id] = true
			}
		}
	}
	candidates := map[string]candidate{}
	resolutions := map[string]SourceReference{}
	add := func(c candidate) error {
		if c.observed.After(input.EvaluatedAt) {
			return fmt.Errorf("future-dated alert source")
		}
		if old, exists := candidates[c.key]; exists && old.source.RecordID != c.source.RecordID {
			return fmt.Errorf("ambiguous alert condition %q", c.key)
		}
		candidates[c.key] = c
		if len(candidates) > MaxCandidates {
			return fmt.Errorf("alert candidate limit exceeded")
		}
		return nil
	}
	if input.Health != nil {
		if len(input.Health.Records) == 0 {
			source := SourceReference{Kind: HealthSource, SchemaName: health.SchemaName, SchemaVersion: health.SchemaVersion, RecordID: "health-envelope-empty", Subject: "health-envelope", EvidenceReferences: []string{}}
			key := conditionKey(EvidenceLoss, source.Subject)
			if err := add(candidate{key: key, category: EvidenceLoss, severity: Warning, reason: "health_evidence_absent", observed: input.EvaluatedAt, source: source, indeterminate: true}); err != nil {
				return nil, nil, err
			}
		}
		for _, r := range input.Health.Records {
			if claimedHealth[r.ID] {
				continue
			}
			subject := healthSubject(r)
			source := SourceReference{Kind: HealthSource, SchemaName: health.SchemaName, SchemaVersion: health.SchemaVersion, RecordID: r.ID, Subject: subject, EvidenceReferences: sortedStrings([]string{r.ID, r.DriftID, r.ChangeID})}
			key := conditionKey(EngineeringCondition, subject)
			if r.Status == health.Healthy {
				resolutions[key] = source
				continue
			}
			sev, ind := healthSeverity(r.Status)
			if err := add(candidate{key: key, category: EngineeringCondition, severity: sev, reason: "health_status_" + string(r.Status), observed: r.EvidenceTimestamp.UTC(), source: source, indeterminate: ind}); err != nil {
				return nil, nil, err
			}
		}
	}
	if input.Rules != nil {
		for _, r := range input.Rules.Records {
			if claimedRules[r.ID] {
				continue
			}
			subject := ruleSubject(r, healthByID)
			source := SourceReference{Kind: RuleSource, SchemaName: rule.SchemaName, SchemaVersion: rule.SchemaVersion, RecordID: r.ID, Subject: subject, EvidenceReferences: sortedStrings(r.EvidenceReferences)}
			category := RuleMatch
			key := conditionKey(category, subject)
			switch r.Outcome {
			case rule.NotMatched, rule.DisabledRule:
				resolutions[key] = source
			case rule.Matched:
				if err := add(candidate{key: key, category: category, severity: Warning, reason: r.Explanation, observed: sourceTime(source, healthByID, input.EvaluatedAt), source: source}); err != nil {
					return nil, nil, err
				}
			case rule.InsufficientEvidence, rule.UnsupportedRule:
				if err := add(candidate{key: conditionKey(EvidenceLoss, subject), category: EvidenceLoss, severity: Warning, reason: r.Explanation, observed: input.EvaluatedAt, source: source, indeterminate: true}); err != nil {
					return nil, nil, err
				}
			case rule.InvalidRule, rule.EvaluationError:
				if err := add(candidate{key: conditionKey(EvidenceLoss, subject), category: EvidenceLoss, severity: Critical, reason: r.Explanation, observed: input.EvaluatedAt, source: source, indeterminate: true}); err != nil {
					return nil, nil, err
				}
			}
		}
	}
	if input.Policies != nil {
		for _, r := range input.Policies.Records {
			subject := policySubject(r, healthByID)
			source := SourceReference{Kind: PolicySource, SchemaName: policy.SchemaName, SchemaVersion: policy.SchemaVersion, RecordID: r.ID, Subject: subject, EvidenceReferences: sortedStrings(append([]string{r.RuleEvaluationID}, r.EvidenceReferences...))}
			key := conditionKey(PolicyGovernance, subject)
			switch r.Outcome {
			case policy.Escalated:
				if err := add(candidate{key: key, category: PolicyGovernance, severity: Emergency, reason: r.Explanation, observed: sourceTime(source, healthByID, input.EvaluatedAt), source: source}); err != nil {
					return nil, nil, err
				}
			case policy.Conflict:
				if err := add(candidate{key: key, category: PolicyGovernance, severity: Critical, reason: r.Explanation, observed: input.EvaluatedAt, source: source, indeterminate: true}); err != nil {
					return nil, nil, err
				}
			case policy.Indeterminate:
				if err := add(candidate{key: key, category: PolicyGovernance, severity: Warning, reason: r.Explanation, observed: input.EvaluatedAt, source: source, indeterminate: true}); err != nil {
					return nil, nil, err
				}
			case policy.Accepted, policy.Observe, policy.Suppressed, policy.NotApplicable:
				resolutions[key] = source
			}
		}
	}
	if input.Scheduler != nil {
		for _, e := range input.Scheduler.Events {
			subject := e.ScheduleID
			if subject == "" {
				subject = "scheduler"
			}
			source := SourceReference{Kind: SchedulerSource, SchemaName: scheduler.EventSchema, SchemaVersion: scheduler.SchemaVersion, RecordID: e.ID, Subject: subject, EvidenceReferences: sortedStrings(nonempty(e.ScheduleID, e.RequestID, input.Scheduler.ID))}
			key := conditionKey(SchedulerOperation, subject)
			switch e.Kind {
			case scheduler.EventStateFailure, scheduler.EventClockDiscontinuity:
				if err := add(candidate{key: key, category: SchedulerOperation, severity: Critical, reason: e.Reason, observed: e.At.UTC(), source: source}); err != nil {
					return nil, nil, err
				}
			case scheduler.EventLockContended:
				if err := add(candidate{key: key, category: SchedulerOperation, severity: Warning, reason: e.Reason, observed: e.At.UTC(), source: source}); err != nil {
					return nil, nil, err
				}
			case scheduler.EventExecutionCompleted, scheduler.EventRestartRecovered, scheduler.EventInitialized:
				resolutions[key] = source
			}
		}
	}
	if input.Report != nil {
		source := SourceReference{Kind: ReportSource, SchemaName: report.SchemaName, SchemaVersion: report.SchemaVersion, RecordID: input.Report.ID, Subject: "canonical-report", EvidenceReferences: []string{input.Report.ID}}
		key := conditionKey(ReportCompleteness, source.Subject)
		if input.Report.Completeness == report.Incomplete {
			if err := add(candidate{key: key, category: ReportCompleteness, severity: Warning, reason: "canonical_report_incomplete", observed: input.EvaluatedAt, source: source, indeterminate: true}); err != nil {
				return nil, nil, err
			}
		} else {
			resolutions[key] = source
		}
	}
	if input.PolicyReport != nil {
		source := SourceReference{Kind: ReportSource, SchemaName: report.PolicyReportSchemaName, SchemaVersion: report.PolicyReportSchemaVersion, RecordID: input.PolicyReport.ID, Subject: "canonical-policy-report", EvidenceReferences: []string{input.PolicyReport.ID}}
		key := conditionKey(ReportCompleteness, source.Subject)
		if input.PolicyReport.Completeness == report.Incomplete {
			if err := add(candidate{key: key, category: ReportCompleteness, severity: Warning, reason: "canonical_policy_report_incomplete", observed: input.EvaluatedAt, source: source, indeterminate: true}); err != nil {
				return nil, nil, err
			}
		} else {
			resolutions[key] = source
		}
	}
	list := make([]candidate, 0, len(candidates))
	for _, c := range candidates {
		list = append(list, c)
	}
	sort.Slice(list, func(i, j int) bool { return list[i].key < list[j].key })
	return list, resolutions, nil
}

func newCondition(c candidate, generation int, evaluatedAt time.Time) Condition {
	condition := Condition{ConditionKey: c.key, Generation: generation, Category: c.category, Severity: c.severity, LifecycleState: StateCandidate, Source: c.source, ReasonToken: c.reason, FirstObservedAt: c.observed, LastObservedAt: c.observed, LastAlertedAt: evaluatedAt, SuppressionIDs: []string{}}
	condition.LifecycleID = lifecycleIdentity(condition.ConditionKey, generation, condition.FirstObservedAt)
	return condition
}

func makeRecord(condition Condition, previous Severity, event EventKind, evaluatedAt, recoveryAt time.Time) Record {
	decision := DecisionAlert
	if condition.LifecycleState == StateSuppressed {
		decision = DecisionSuppressed
	}
	if event == EventRecovered || event == EventExpired {
		decision = DecisionLifecycle
	}
	record := Record{SchemaName: RecordSchema, SchemaVersion: SchemaVersion, ConditionKey: condition.ConditionKey, LifecycleID: condition.LifecycleID, Generation: condition.Generation,
		Event: event, Decision: decision, LifecycleState: condition.LifecycleState, Severity: condition.Severity, PreviousSeverity: previous, Category: condition.Category, Source: condition.Source,
		EventTime: evaluatedAt, ObservationTime: condition.LastObservedAt, EvaluationTime: evaluatedAt, AcknowledgementID: condition.AcknowledgementID,
		SuppressionIDs: append([]string{}, condition.SuppressionIDs...), ExpirationTime: condition.ExpiredAt, RecoveryTime: recoveryAt, ReasonToken: condition.ReasonToken, Versions: versions()}
	record.ID = recordIdentity(record)
	return record
}

func matchingSuppressionIDs(windows []SuppressionWindow, condition Condition, at time.Time) []string {
	ids := []string{}
	for _, w := range windows {
		if at.Before(w.StartsAt) || !at.Before(w.EndsAt) || !matchesString(w.ConditionKeys, condition.ConditionKey) || !matchesCategory(w.Categories, condition.Category) || !matchesSeverity(w.Severities, condition.Severity) || (condition.Severity == Emergency && !w.SuppressEmergency) {
			continue
		}
		ids = append(ids, w.ID)
	}
	sort.Strings(ids)
	return ids
}

func endedMaintenance(ids []string, windows []SuppressionWindow, at time.Time) bool {
	set := map[string]bool{}
	for _, id := range ids {
		set[id] = true
	}
	for _, w := range windows {
		if set[w.ID] && w.Kind == MaintenanceSuppression && !at.Before(w.EndsAt) {
			return true
		}
	}
	return false
}

func validateAcknowledgement(v Acknowledgement) error {
	if v.SchemaName != AcknowledgementSchema || v.SchemaVersion != SchemaVersion || !token(v.LifecycleID) || !token(v.ActorID) || !token(v.AuthorityRef) || !validTime(v.At) || (v.NoteToken != "" && !token(v.NoteToken)) || v.ID != acknowledgementIdentity(v) {
		return fmt.Errorf("invalid alert acknowledgement")
	}
	return nil
}

func validateSuppression(v SuppressionWindow) error {
	if v.SchemaName != SuppressionSchema || v.SchemaVersion != SchemaVersion || (v.Kind != OperationalSuppression && v.Kind != MaintenanceSuppression) || !validTime(v.StartsAt) || !validTime(v.EndsAt) || !v.StartsAt.Before(v.EndsAt) || v.EndsAt.Sub(v.StartsAt) > MaxSuppression || !token(v.ActorID) || !token(v.AuthorityRef) || !token(v.ReasonToken) || !sortedUnique(v.ConditionKeys) || !sortedUniqueCategories(v.Categories) || !sortedUniqueSeverities(v.Severities) || v.ID != suppressionIdentity(v) {
		return fmt.Errorf("invalid alert suppression window")
	}
	if containsSeverity(v.Severities, Emergency) && !v.SuppressEmergency {
		return fmt.Errorf("emergency suppression requires explicit authorization")
	}
	return nil
}

func validCondition(v Condition) bool {
	if !token(v.ConditionKey) || !token(v.LifecycleID) || v.Generation < 1 || !validCategory(v.Category) || !validSeverity(v.Severity) || !validLifecycle(v.LifecycleState) || !validSource(v.Source) || !token(v.ReasonToken) || !validTime(v.FirstObservedAt) || !validTime(v.LastObservedAt) || v.LastObservedAt.Before(v.FirstObservedAt) || !validTime(v.LastAlertedAt) || !sortedUnique(v.SuppressionIDs) || (v.AcknowledgementID != "" && !token(v.AcknowledgementID)) || v.LifecycleID != lifecycleIdentity(v.ConditionKey, v.Generation, v.FirstObservedAt) {
		return false
	}
	if v.LifecycleState == StateResolved {
		return validTime(v.ResolvedAt) && v.ExpiredAt.IsZero()
	}
	if v.LifecycleState == StateExpired {
		return validTime(v.ExpiredAt) && v.ResolvedAt.IsZero()
	}
	return v.ResolvedAt.IsZero() && v.ExpiredAt.IsZero()
}

func validRecord(v Record) bool {
	if v.SchemaName != RecordSchema || v.SchemaVersion != SchemaVersion || !token(v.ConditionKey) || !token(v.LifecycleID) || v.Generation < 1 || !validEvent(v.Event) || !validDecision(v.Decision) || !validLifecycle(v.LifecycleState) || !validSeverity(v.Severity) || (v.PreviousSeverity != "" && !validSeverity(v.PreviousSeverity)) || !validCategory(v.Category) || !validSource(v.Source) || !validTime(v.EventTime) || !validTime(v.ObservationTime) || !validTime(v.EvaluationTime) || v.EventTime != v.EvaluationTime || !sortedUnique(v.SuppressionIDs) || !token(v.ReasonToken) || v.Versions != versions() {
		return false
	}
	if v.Event == EventRecovered {
		return v.LifecycleState == StateResolved && validTime(v.RecoveryTime) && v.ExpirationTime.IsZero()
	}
	if v.Event == EventExpired {
		return v.LifecycleState == StateExpired && validTime(v.ExpirationTime) && v.RecoveryTime.IsZero()
	}
	return v.ExpirationTime.IsZero() && v.RecoveryTime.IsZero()
}

func validSource(v SourceReference) bool {
	if !validSourceKind(v.Kind) || v.SchemaName == "" || v.SchemaVersion != SchemaVersion || !token(v.RecordID) || !token(v.Subject) || len(v.EvidenceReferences) > MaxReferences || !sortedUnique(v.EvidenceReferences) {
		return false
	}
	return true
}

func healthSeverity(status health.Status) (Severity, bool) {
	switch status {
	case health.Informational:
		return Informational, false
	case health.Advisory, health.Warning:
		return Warning, false
	case health.Critical:
		return Critical, false
	case health.Unsupported, health.Unknown:
		return Warning, true
	default:
		return Informational, false
	}
}
func severityRank(s Severity) int {
	switch s {
	case Informational:
		return 1
	case Warning:
		return 2
	case Critical:
		return 3
	case Emergency:
		return 4
	}
	return 0
}
func terminalState(s LifecycleState) bool { return s == StateResolved || s == StateExpired }
func conditionKey(category Category, subject string) string {
	return "condition:" + digest(string(category), subject)
}
func lifecycleIdentity(key string, generation int, at time.Time) string {
	return "lifecycle:" + digest(key, fmt.Sprint(generation), at.UTC().Format(time.RFC3339Nano))
}
func acknowledgementIdentity(v Acknowledgement) string {
	copy := v
	copy.ID = ""
	return "ack:" + hashJSON(copy)
}
func suppressionIdentity(v SuppressionWindow) string {
	copy := v
	copy.ID = ""
	return "suppression:" + hashJSON(copy)
}
func stateIdentity(v State) string {
	copy := cloneState(v)
	copy.ID = ""
	return "alert-state:" + hashJSON(copy)
}
func recordIdentity(v Record) string { copy := v; copy.ID = ""; return "alert:" + hashJSON(copy) }
func resultIdentity(v Result) string {
	copy := v
	copy.ID = ""
	return "alert-result:" + hashJSON(copy)
}
func versions() VersionInfo {
	return VersionInfo{SchemaVersion, EngineVersion, TaxonomyVersion, ModelVersion}
}
func digest(parts ...string) string {
	h := sha256.New()
	for _, p := range parts {
		h.Write([]byte{0})
		h.Write([]byte(p))
	}
	return hex.EncodeToString(h.Sum(nil))
}
func hashJSON(v any) string {
	data, err := json.Marshal(v)
	if err != nil {
		panic(err)
	}
	sum := sha256.Sum256(data)
	return hex.EncodeToString(sum[:])
}

func upsertCondition(state *State, condition Condition) {
	for i := range state.Conditions {
		if state.Conditions[i].ConditionKey == condition.ConditionKey && state.Conditions[i].Generation == condition.Generation {
			state.Conditions[i] = cloneCondition(condition)
			return
		}
	}
	state.Conditions = append(state.Conditions, cloneCondition(condition))
}
func sortConditions(v []Condition) {
	sort.Slice(v, func(i, j int) bool {
		if v[i].ConditionKey != v[j].ConditionKey {
			return v[i].ConditionKey < v[j].ConditionKey
		}
		return v[i].Generation < v[j].Generation
	})
}
func cloneState(v State) State {
	out := v
	out.Conditions = make([]Condition, len(v.Conditions))
	for i := range v.Conditions {
		out.Conditions[i] = cloneCondition(v.Conditions[i])
	}
	return out
}
func cloneCondition(v Condition) Condition {
	out := v
	out.Source.EvidenceReferences = append([]string{}, v.Source.EvidenceReferences...)
	out.SuppressionIDs = append([]string{}, v.SuppressionIDs...)
	return out
}

func healthSubject(v health.Record) string {
	return canonicalSubject(v.Scope.Layer, v.Scope.ObjectID, v.Scope.Path)
}
func ruleSubject(v rule.EvaluationRecord, healthByID map[string]health.Record) string {
	if h, ok := healthByID[v.HealthRecordID]; ok {
		return canonicalSubject(v.RuleID, h.Scope.Layer, h.Scope.ObjectID, h.Scope.Path)
	}
	if v.HealthRecordID != "" {
		return canonicalSubject(v.RuleID, v.HealthRecordID)
	}
	return canonicalSubject(v.RuleID)
}
func policySubject(v policy.EvaluationRecord, healthByID map[string]health.Record) string {
	for _, id := range v.EvidenceReferences {
		if h, ok := healthByID[id]; ok {
			return canonicalSubject(v.RuleID, h.Scope.Layer, h.Scope.ObjectID, h.Scope.Path)
		}
	}
	return canonicalSubject(v.RuleID, strings.Join(v.EvidenceReferences, ":"))
}
func canonicalSubject(parts ...string) string { return "subject:" + digest(parts...) }
func sourceTime(source SourceReference, healthByID map[string]health.Record, fallback time.Time) time.Time {
	for _, id := range source.EvidenceReferences {
		if h, ok := healthByID[id]; ok {
			return h.EvidenceTimestamp.UTC()
		}
	}
	return fallback
}
func nonempty(v ...string) []string {
	out := []string{}
	for _, x := range v {
		if x != "" {
			out = append(out, x)
		}
	}
	return sortedStrings(out)
}
func sortedStrings(v []string) []string {
	out := append([]string{}, v...)
	sort.Strings(out)
	j := 0
	for _, x := range out {
		if j == 0 || out[j-1] != x {
			out[j] = x
			j++
		}
	}
	return out[:j]
}

func validTime(v time.Time) bool { return !v.IsZero() && v.Location() == time.UTC }
func token(v string) bool {
	return len(v) > 0 && len(v) <= MaxStringLength && tokenPattern.MatchString(v)
}
func sortedUnique(v []string) bool {
	if v == nil {
		return false
	}
	for i, x := range v {
		if !token(x) || (i > 0 && v[i-1] >= x) {
			return false
		}
	}
	return true
}
func sortedUniqueCategories(v []Category) bool {
	if v == nil {
		return false
	}
	for i, x := range v {
		if !validCategory(x) || (i > 0 && v[i-1] >= x) {
			return false
		}
	}
	return true
}
func sortedUniqueSeverities(v []Severity) bool {
	if v == nil {
		return false
	}
	for i, x := range v {
		if !validSeverity(x) || (i > 0 && v[i-1] >= x) {
			return false
		}
	}
	return true
}
func matchesString(scope []string, v string) bool {
	return len(scope) == 0 || sort.SearchStrings(scope, v) < len(scope) && scope[sort.SearchStrings(scope, v)] == v
}
func matchesCategory(scope []Category, v Category) bool {
	if len(scope) == 0 {
		return true
	}
	i := sort.Search(len(scope), func(i int) bool { return scope[i] >= v })
	return i < len(scope) && scope[i] == v
}
func matchesSeverity(scope []Severity, v Severity) bool {
	if len(scope) == 0 {
		return true
	}
	i := sort.Search(len(scope), func(i int) bool { return scope[i] >= v })
	return i < len(scope) && scope[i] == v
}
func containsSeverity(scope []Severity, v Severity) bool {
	return matchesSeverity(scope, v) && len(scope) > 0
}
func validSeverity(v Severity) bool {
	return v == Informational || v == Warning || v == Critical || v == Emergency
}
func validCategory(v Category) bool {
	return v == EngineeringCondition || v == RuleMatch || v == PolicyGovernance || v == SchedulerOperation || v == ReportCompleteness || v == EvidenceLoss
}
func validSourceKind(v SourceKind) bool {
	return v == HealthSource || v == RuleSource || v == PolicySource || v == SchedulerSource || v == ReportSource
}
func validLifecycle(v LifecycleState) bool {
	switch v {
	case StateCandidate, StateActive, StateAcknowledged, StateSuppressed, StateExpired, StateResolved, StateIndeterminate:
		return true
	}
	return false
}
func validEvent(v EventKind) bool {
	switch v {
	case EventEntered, EventEscalated, EventDeescalated, EventAcknowledged, EventSuppressionStarted, EventSuppressionEnded, EventMaintenanceEnded, EventReminder, EventExpired, EventRecovered, EventIndeterminate:
		return true
	}
	return false
}
func validDecision(v Decision) bool {
	return v == DecisionAlert || v == DecisionSuppressed || v == DecisionLifecycle
}

func strictDecode(data []byte, target any) error {
	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(target); err != nil {
		return fmt.Errorf("invalid alert JSON: %w", err)
	}
	if err := ensureEOF(decoder); err != nil {
		return err
	}
	return nil
}
func ensureEOF(decoder *json.Decoder) error {
	var extra any
	if err := decoder.Decode(&extra); err != io.EOF {
		if err == nil {
			return fmt.Errorf("unexpected trailing alert JSON")
		}
		return err
	}
	return nil
}
