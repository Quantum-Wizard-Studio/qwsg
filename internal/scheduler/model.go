// Package scheduler implements deterministic scheduling over the Canonical
// Configuration, Command, and Pipeline contracts. The Engine is pure; local
// persistence, locking, and execution are isolated in the Cycle adapter.
package scheduler

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"regexp"
	"sort"
	"time"

	"quantumwizard.hu/qwsg/internal/command"
	"quantumwizard.hu/qwsg/internal/configuration"
)

const (
	SchemaVersion     = "1.0"
	EngineVersion     = "1.0"
	TaxonomyVersion   = "1.0"
	StateSchema       = "qwsg.scheduler-state"
	EvaluationSchema  = "qwsg.scheduler-evaluation"
	RequestSchema     = "qwsg.scheduler-execution-request"
	ResultSchema      = "qwsg.scheduler-execution-result"
	TraceSchema       = "qwsg.scheduler-execution-trace"
	CycleResultSchema = "qwsg.scheduler-cycle-result"
	EventSchema       = "qwsg.scheduler-event"
	MaxOccurrences    = 1024
	// State retains enough recent outcomes for bounded retry and operational
	// continuity without allowing large policy result sets to exhaust Guardian's
	// cgroup during repeated JSON validation/publication.
	MaxStateResults     = 64
	MaxPolicyReferences = 4096
	MaxLookahead        = 366 * 24 * time.Hour
	ClockTolerance      = 5 * time.Second
)

type Decision string

const (
	DecisionDisabled      Decision = "disabled"
	DecisionNotDue        Decision = "not_due"
	DecisionDue           Decision = "due"
	DecisionSkipped       Decision = "skipped"
	DecisionQueued        Decision = "queued"
	DecisionDelayed       Decision = "delayed"
	DecisionInapplicable  Decision = "inapplicable"
	DecisionIndeterminate Decision = "indeterminate"
)

type EventKind string

const (
	EventInitialized        EventKind = "initialized"
	EventClockDiscontinuity EventKind = "clock_discontinuity"
	EventRestartRecovered   EventKind = "restart_recovered"
	EventRequestReserved    EventKind = "request_reserved"
	EventExecutionCompleted EventKind = "execution_completed"
	EventLockContended      EventKind = "lock_contended"
	EventStateFailure       EventKind = "state_failure"
)

type ExecutionOutcome string

const (
	ExecutionSucceeded   ExecutionOutcome = "succeeded"
	ExecutionFailed      ExecutionOutcome = "failed"
	ExecutionIncomplete  ExecutionOutcome = "incomplete"
	ExecutionInterrupted ExecutionOutcome = "interrupted"
)

type ClockObservation struct {
	WallTime    time.Time `json:"wall_time"`
	SessionID   string    `json:"session_id"`
	MonotonicNS int64     `json:"monotonic_ns"`
}

type Occurrence struct {
	ID          string    `json:"id"`
	ScheduleID  string    `json:"schedule_id"`
	ScheduledAt time.Time `json:"scheduled_at"`
	Trigger     string    `json:"trigger"`
}

type ActiveRun struct {
	RequestID    string    `json:"request_id"`
	OccurrenceID string    `json:"occurrence_id"`
	ScheduledAt  time.Time `json:"scheduled_at"`
	Attempt      int       `json:"attempt"`
	ReservedAt   time.Time `json:"reserved_at"`
}

type ScheduleState struct {
	ScheduleID      string      `json:"schedule_id"`
	AnchorAt        time.Time   `json:"anchor_at"`
	LastEvaluatedAt time.Time   `json:"last_evaluated_at"`
	LastScheduledAt time.Time   `json:"last_scheduled_at"`
	NextRunAt       time.Time   `json:"next_run_at"`
	Active          []ActiveRun `json:"active"`
	Pending         *Occurrence `json:"pending,omitempty"`
}

// State is Scheduler State Record 1.0.
type State struct {
	SchemaName      string            `json:"schema_name"`
	SchemaVersion   string            `json:"schema_version"`
	ID              string            `json:"id"`
	ConfigurationID string            `json:"configuration_id"`
	SessionID       string            `json:"session_id"`
	LastWallTime    time.Time         `json:"last_wall_time"`
	LastMonotonicNS int64             `json:"last_monotonic_ns"`
	Schedules       []ScheduleState   `json:"schedules"`
	Results         []ExecutionResult `json:"results"`
}

// Record is Scheduler Evaluation Record 1.0.
type Record struct {
	ID              string    `json:"id"`
	ScheduleID      string    `json:"schedule_id"`
	Decision        Decision  `json:"decision"`
	Reason          string    `json:"reason"`
	OccurrenceIDs   []string  `json:"occurrence_ids"`
	RequestIDs      []string  `json:"request_ids"`
	NextRunAt       time.Time `json:"next_run_at"`
	ConfigurationID string    `json:"configuration_id"`
	EvaluatedAt     time.Time `json:"evaluated_at"`
}

// Request is Scheduler Execution Request 1.0.
type Request struct {
	SchemaName         string    `json:"schema_name"`
	SchemaVersion      string    `json:"schema_version"`
	ID                 string    `json:"id"`
	EvaluationID       string    `json:"evaluation_id"`
	ConfigurationID    string    `json:"configuration_id"`
	ScheduleID         string    `json:"schedule_id"`
	OccurrenceID       string    `json:"occurrence_id"`
	ScheduledAt        time.Time `json:"scheduled_at"`
	Attempt            int       `json:"attempt"`
	Priority           int       `json:"priority"`
	CommandProfile     string    `json:"command_profile"`
	CheckIDs           []string  `json:"check_ids"`
	ExecutionTimeoutNS int64     `json:"execution_timeout_ns"`
	RetryPolicyID      string    `json:"retry_policy_id"`
}

// Event is Scheduler Event Record 1.0.
type Event struct {
	SchemaName    string            `json:"schema_name"`
	SchemaVersion string            `json:"schema_version"`
	ID            string            `json:"id"`
	Kind          EventKind         `json:"kind"`
	At            time.Time         `json:"at"`
	ScheduleID    string            `json:"schedule_id,omitempty"`
	RequestID     string            `json:"request_id,omitempty"`
	Reason        string            `json:"reason"`
	Metadata      map[string]string `json:"metadata"`
}

// ExecutionResult is Scheduler Execution Result Record 1.0.
type ExecutionResult struct {
	SchemaName          string           `json:"schema_name"`
	SchemaVersion       string           `json:"schema_version"`
	ID                  string           `json:"id"`
	RequestID           string           `json:"request_id"`
	ScheduleID          string           `json:"schedule_id"`
	OccurrenceID        string           `json:"occurrence_id"`
	ScheduledAt         time.Time        `json:"scheduled_at"`
	Attempt             int              `json:"attempt"`
	Outcome             ExecutionOutcome `json:"outcome"`
	StartedAt           time.Time        `json:"started_at"`
	CompletedAt         time.Time        `json:"completed_at"`
	CommandExecutionID  string           `json:"command_execution_id,omitempty"`
	CommandComplete     bool             `json:"command_complete"`
	StageContracts      []string         `json:"stage_contracts"`
	PolicyEvaluationIDs []string         `json:"policy_evaluation_ids"`
	PolicyOutcomes      []string         `json:"policy_outcomes"`
	FailureCode         string           `json:"failure_code,omitempty"`
	NextRetryAt         time.Time        `json:"next_retry_at"`
}

// Evaluation is the immutable Scheduler evaluation envelope.
type Evaluation struct {
	SchemaName      string    `json:"schema_name"`
	SchemaVersion   string    `json:"schema_version"`
	EngineVersion   string    `json:"engine_version"`
	TaxonomyVersion string    `json:"taxonomy_version"`
	ID              string    `json:"id"`
	ConfigurationID string    `json:"configuration_id"`
	EvaluatedAt     time.Time `json:"evaluated_at"`
	Records         []Record  `json:"records"`
	Requests        []Request `json:"requests"`
	Events          []Event   `json:"events"`
	NextState       State     `json:"next_state"`
}

type Completion struct {
	Request     Request
	StartedAt   time.Time
	CompletedAt time.Time
	Execution   command.Execution
	FailureCode string
}

// ExecutionTrace is the bounded immutable seam through which a Runtime
// coordinator may inspect outputs already produced by the Scheduler-owned
// Pipeline execution. It does not authorize execution or reinterpretation.
type ExecutionTrace struct {
	SchemaName    string             `json:"schema_name"`
	SchemaVersion string             `json:"schema_version"`
	ID            string             `json:"id"`
	RequestID     string             `json:"request_id"`
	Definition    command.Definition `json:"definition"`
	Execution     command.Execution  `json:"execution"`
	FailureCode   string             `json:"failure_code,omitempty"`
}

var idPattern = regexp.MustCompile(`^[a-z0-9]+([._:-][a-z0-9]+)*$`)

func NewState(configurationID string) State {
	state := State{SchemaName: StateSchema, SchemaVersion: SchemaVersion, ConfigurationID: configurationID, Schedules: []ScheduleState{}, Results: []ExecutionResult{}}
	state.ID = stateID(state)
	return state
}

func ValidateState(state State) error {
	if state.SchemaName != StateSchema || state.SchemaVersion != SchemaVersion || state.ConfigurationID == "" || state.Schedules == nil || state.Results == nil || len(state.Results) > MaxStateResults {
		return fmt.Errorf("invalid scheduler state envelope")
	}
	last := ""
	for _, item := range state.Schedules {
		if !idPattern.MatchString(item.ScheduleID) || item.ScheduleID <= last || item.Active == nil {
			return fmt.Errorf("invalid or unordered schedule state")
		}
		for i, active := range item.Active {
			if active.RequestID == "" || active.OccurrenceID == "" || active.Attempt < 1 {
				return fmt.Errorf("invalid active run contract for %q: request=%q occurrence=%q attempt=%d", item.ScheduleID, active.RequestID, active.OccurrenceID, active.Attempt)
			}
			if i > 0 && item.Active[i-1].RequestID >= active.RequestID {
				return fmt.Errorf("invalid active run ordering for %q", item.ScheduleID)
			}
		}
		last = item.ScheduleID
	}
	for _, result := range state.Results {
		if err := validateResult(result); err != nil {
			return err
		}
	}
	if state.ID != stateID(state) {
		return fmt.Errorf("invalid scheduler state identity")
	}
	return nil
}

func ValidateEvaluation(value Evaluation) error {
	if value.SchemaName != EvaluationSchema || value.SchemaVersion != SchemaVersion || value.EngineVersion != EngineVersion || value.TaxonomyVersion != TaxonomyVersion || value.ID == "" || value.ConfigurationID == "" || value.Records == nil || value.Requests == nil || value.Events == nil {
		return fmt.Errorf("invalid scheduler evaluation envelope")
	}
	for i, record := range value.Records {
		if record.ID != recordID(record) {
			return fmt.Errorf("invalid scheduler record identity for %q", record.ScheduleID)
		}
		if i > 0 && value.Records[i-1].ScheduleID >= record.ScheduleID {
			return fmt.Errorf("unordered scheduler record")
		}
	}
	for _, request := range value.Requests {
		if err := validateRequest(request); err != nil {
			return err
		}
	}
	for _, event := range value.Events {
		if event.ID != eventID(event) || event.Metadata == nil {
			return fmt.Errorf("invalid scheduler event")
		}
	}
	if err := ValidateState(value.NextState); err != nil {
		return err
	}
	if value.ID != evaluationID(value) {
		return fmt.Errorf("invalid scheduler evaluation identity")
	}
	return nil
}

func MarshalStateCanonical(value State) ([]byte, error) {
	if err := ValidateState(value); err != nil {
		return nil, err
	}
	return json.Marshal(value)
}
func MarshalEvaluationCanonical(value Evaluation) ([]byte, error) {
	if err := ValidateEvaluation(value); err != nil {
		return nil, err
	}
	return json.Marshal(value)
}

func validateRequest(value Request) error {
	if value.SchemaName != RequestSchema || value.SchemaVersion != SchemaVersion || value.EvaluationID == "" || value.ConfigurationID == "" || !idPattern.MatchString(value.ScheduleID) || value.OccurrenceID == "" || value.Attempt < 1 || value.CommandProfile == "" || value.CheckIDs == nil || !sortedUnique(value.CheckIDs) || value.ExecutionTimeoutNS < 1 || !idPattern.MatchString(value.RetryPolicyID) {
		return fmt.Errorf("invalid scheduler execution request contract: schema=%t evaluation=%t configuration=%t schedule=%t occurrence=%t attempt=%t profile=%t checks=%t timeout=%t retry=%t", value.SchemaName == RequestSchema && value.SchemaVersion == SchemaVersion, value.EvaluationID != "", value.ConfigurationID != "", idPattern.MatchString(value.ScheduleID), value.OccurrenceID != "", value.Attempt >= 1, value.CommandProfile != "", value.CheckIDs != nil && sortedUnique(value.CheckIDs), value.ExecutionTimeoutNS >= 1, idPattern.MatchString(value.RetryPolicyID))
	}
	if value.ID != requestID(value) {
		return fmt.Errorf("invalid scheduler execution request identity")
	}
	return nil
}

func validateResult(value ExecutionResult) error {
	if value.SchemaName != ResultSchema || value.SchemaVersion != SchemaVersion || value.ID != resultID(value) || value.RequestID == "" || !idPattern.MatchString(value.ScheduleID) || value.OccurrenceID == "" || value.ScheduledAt.IsZero() || value.Attempt < 1 || !validOutcome(value.Outcome) || value.StageContracts == nil || value.PolicyEvaluationIDs == nil || value.PolicyOutcomes == nil || !sortedUnique(value.StageContracts) || !sortedUnique(value.PolicyEvaluationIDs) || !sortedUnique(value.PolicyOutcomes) || len(value.PolicyEvaluationIDs) > MaxPolicyReferences {
		return fmt.Errorf("invalid scheduler execution result")
	}
	return nil
}

func validateTrace(value ExecutionTrace) error {
	if value.SchemaName != TraceSchema || value.SchemaVersion != SchemaVersion || value.RequestID == "" ||
		(value.FailureCode != "" && value.FailureCode != "command_resolution_failed" && value.FailureCode != "pipeline_execution_failed") {
		return fmt.Errorf("invalid scheduler execution trace")
	}
	if value.FailureCode == "command_resolution_failed" {
		if value.Definition.ID != "" || value.Execution.ID != "" {
			return fmt.Errorf("resolution failure contains command output")
		}
	} else {
		if err := command.ValidateDefinition(value.Definition); err != nil {
			return fmt.Errorf("invalid traced command definition: %w", err)
		}
		if value.Execution.CommandID != "" && value.Execution.CommandID != value.Definition.ID {
			return fmt.Errorf("traced command identity mismatch")
		}
	}
	if value.FailureCode == "" && (value.Execution.ID == "" || !value.Execution.Complete) {
		return fmt.Errorf("successful trace contains incomplete execution")
	}
	if value.ID != traceID(value) {
		return fmt.Errorf("invalid scheduler execution trace identity")
	}
	return nil
}

func validOutcome(value ExecutionOutcome) bool {
	return value == ExecutionSucceeded || value == ExecutionFailed || value == ExecutionIncomplete || value == ExecutionInterrupted
}
func sortedUnique(values []string) bool {
	for i, v := range values {
		if v == "" || (i > 0 && values[i-1] >= v) {
			return false
		}
	}
	return values != nil
}
func normalizeTime(value time.Time) time.Time {
	if value.IsZero() {
		return time.Time{}
	}
	return value.UTC().Truncate(0)
}
func stableID(prefix string, value any) string {
	data, err := json.Marshal(value)
	if err != nil {
		panic(err)
	}
	sum := sha256.Sum256(data)
	return prefix + ":" + hex.EncodeToString(sum[:])
}
func stateID(value State) string          { value.ID = ""; return stableID("scheduler-state", value) }
func traceID(value ExecutionTrace) string { value.ID = ""; return stableID("scheduler-trace", value) }
func evaluationID(value Evaluation) string {
	// Evaluation contains slices nested below records and next-state entries.
	// Clone before removing derived cross-references so identity calculation
	// cannot mutate the caller through shared slice backing arrays.
	data, err := json.Marshal(value)
	if err != nil {
		panic(err)
	}
	var clone Evaluation
	if err := json.Unmarshal(data, &clone); err != nil {
		panic(err)
	}
	value = clone
	value.ID = ""
	// Derived cross-references are excluded to avoid circular identity:
	// requests bind to the evaluation, while records and active reservations
	// bind to requests. Their own identities still cover those references.
	for i := range value.Requests {
		value.Requests[i].ID = ""
		value.Requests[i].EvaluationID = ""
	}
	for i := range value.Records {
		value.Records[i].ID = ""
		value.Records[i].RequestIDs = []string{}
	}
	for i := range value.Events {
		value.Events[i].ID = ""
		value.Events[i].RequestID = ""
	}
	for i := range value.NextState.Schedules {
		for j := range value.NextState.Schedules[i].Active {
			value.NextState.Schedules[i].Active[j].RequestID = ""
		}
	}
	value.NextState.ID = ""
	return stableID("scheduler-evaluation", value)
}
func recordID(value Record) string   { value.ID = ""; return stableID("scheduler-record", value) }
func requestID(value Request) string { value.ID = ""; return stableID("scheduler-request", value) }
func eventID(value Event) string     { value.ID = ""; return stableID("scheduler-event", value) }
func resultID(value ExecutionResult) string {
	value.ID = ""
	return stableID("scheduler-result", value)
}
func occurrenceID(scheduleID string, at time.Time, trigger string) string {
	return stableID("scheduler-occurrence", struct {
		ScheduleID string    `json:"schedule_id"`
		At         time.Time `json:"at"`
		Trigger    string    `json:"trigger"`
	}{scheduleID, normalizeTime(at), trigger})
}

func sortState(state *State) {
	sort.Slice(state.Schedules, func(i, j int) bool { return state.Schedules[i].ScheduleID < state.Schedules[j].ScheduleID })
	for i := range state.Schedules {
		sort.Slice(state.Schedules[i].Active, func(a, b int) bool {
			return state.Schedules[i].Active[a].RequestID < state.Schedules[i].Active[b].RequestID
		})
	}
	sort.Slice(state.Results, func(i, j int) bool {
		if state.Results[i].CompletedAt.Equal(state.Results[j].CompletedAt) {
			return state.Results[i].ID < state.Results[j].ID
		}
		return state.Results[i].CompletedAt.Before(state.Results[j].CompletedAt)
	})
	if len(state.Results) > MaxStateResults {
		state.Results = append([]ExecutionResult(nil), state.Results[len(state.Results)-MaxStateResults:]...)
	}
	state.ID = stateID(*state)
}

func cloneState(state State) State {
	data, _ := json.Marshal(state)
	var result State
	_ = json.Unmarshal(data, &result)
	return result
}

func findScheduleState(state *State, id string) *ScheduleState {
	index := sort.Search(len(state.Schedules), func(i int) bool { return state.Schedules[i].ScheduleID >= id })
	if index < len(state.Schedules) && state.Schedules[index].ScheduleID == id {
		return &state.Schedules[index]
	}
	state.Schedules = append(state.Schedules, ScheduleState{})
	copy(state.Schedules[index+1:], state.Schedules[index:])
	state.Schedules[index] = ScheduleState{ScheduleID: id, Active: []ActiveRun{}}
	return &state.Schedules[index]
}

func retryPolicy(config configuration.Effective, id string) (configuration.RetryPolicy, bool) {
	for _, item := range config.Values.RetryPolicies {
		if item.ID == id {
			return item, true
		}
	}
	return configuration.RetryPolicy{}, false
}
