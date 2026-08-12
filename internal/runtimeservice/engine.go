// Package runtimeservice owns recurrence and local process lifecycle for the
// Canonical Runtime Engine. It does not implement or reinterpret Runtime work.
package runtimeservice

import (
	"bytes"
	"context"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"math"
	"time"

	"quantumwizard.hu/qwsg/internal/runtime"
)

const (
	SchemaVersion     = "1.0"
	ModelVersion      = "1.0"
	DefinitionSchema  = "qwsg.runtime-service-definition"
	StateSchema       = "qwsg.runtime-service-state"
	EventSchema       = "qwsg.runtime-service-event"
	EvidenceSchema    = "qwsg.runtime-service-evidence"
	ResultSchema      = "qwsg.runtime-service-result"
	StartupImmediate  = "immediate"
	MaxInterval       = 24 * time.Hour
	MaxCycleTimeout   = 24 * time.Hour
	MaxCycleSequence  = uint64(1<<53 - 1)
	MaxEvidenceString = 256
)

type Lifecycle string

const (
	Created  Lifecycle = "created"
	Starting Lifecycle = "starting"
	Running  Lifecycle = "running"
	Stopping Lifecycle = "stopping"
	Stopped  Lifecycle = "stopped"
	Failed   Lifecycle = "failed"
)

type EvidenceKind string

const (
	EvidenceStartup           EvidenceKind = "startup"
	EvidenceCycleScheduled    EvidenceKind = "cycle_scheduled"
	EvidenceCycleStarted      EvidenceKind = "cycle_started"
	EvidenceCycleCompleted    EvidenceKind = "cycle_completed"
	EvidenceIntervalsMissed   EvidenceKind = "intervals_missed"
	EvidenceShutdownRequested EvidenceKind = "shutdown_requested"
	EvidenceShutdownCompleted EvidenceKind = "shutdown_completed"
	EvidenceTerminalFailure   EvidenceKind = "terminal_failure"
)

type Definition struct {
	SchemaName     string `json:"schema_name"`
	SchemaVersion  string `json:"schema_version"`
	ModelVersion   string `json:"model_version"`
	ID             string `json:"id"`
	ServiceID      string `json:"service_id"`
	IntervalNS     int64  `json:"interval_ns"`
	CycleTimeoutNS int64  `json:"cycle_timeout_ns"`
	StartupMode    string `json:"startup_mode"`
}

type State struct {
	SchemaName          string          `json:"schema_name"`
	SchemaVersion       string          `json:"schema_version"`
	ID                  string          `json:"id"`
	ServiceID           string          `json:"service_id"`
	Lifecycle           Lifecycle       `json:"lifecycle"`
	Sequence            uint64          `json:"sequence"`
	CyclesStarted       uint64          `json:"cycles_started"`
	CyclesCompleted     uint64          `json:"cycles_completed"`
	IntervalsMissed     uint64          `json:"intervals_missed"`
	ActiveCycleID       string          `json:"active_cycle_id,omitempty"`
	NextNominalAt       time.Time       `json:"next_nominal_at,omitempty"`
	LastRuntimeResultID string          `json:"last_runtime_result_id,omitempty"`
	LastRuntimeOutcome  runtime.Outcome `json:"last_runtime_outcome,omitempty"`
}

type Input struct {
	Definition   Definition    `json:"definition"`
	StartedAt    time.Time     `json:"started_at"`
	InitialState State         `json:"initial_state"`
	Seed         runtime.Input `json:"seed"`
}

type Event struct {
	SchemaName      string          `json:"schema_name"`
	SchemaVersion   string          `json:"schema_version"`
	ID              string          `json:"id"`
	Sequence        uint64          `json:"sequence"`
	Kind            EvidenceKind    `json:"kind"`
	At              time.Time       `json:"at"`
	ServiceID       string          `json:"service_id"`
	CycleID         string          `json:"cycle_id,omitempty"`
	RuntimeResultID string          `json:"runtime_result_id,omitempty"`
	RuntimeOutcome  runtime.Outcome `json:"runtime_outcome,omitempty"`
	MissedCount     uint64          `json:"missed_count,omitempty"`
	ReasonToken     string          `json:"reason_token"`
}

type Evidence struct {
	SchemaName    string   `json:"schema_name"`
	SchemaVersion string   `json:"schema_version"`
	ID            string   `json:"id"`
	EventID       string   `json:"event_id"`
	ServiceID     string   `json:"service_id"`
	ReferenceIDs  []string `json:"reference_ids"`
}

type Result struct {
	SchemaName          string          `json:"schema_name"`
	SchemaVersion       string          `json:"schema_version"`
	ModelVersion        string          `json:"model_version"`
	ID                  string          `json:"id"`
	ServiceID           string          `json:"service_id"`
	StartedAt           time.Time       `json:"started_at"`
	CompletedAt         time.Time       `json:"completed_at"`
	FinalState          State           `json:"final_state"`
	LastRuntimeResultID string          `json:"last_runtime_result_id,omitempty"`
	LastRuntimeOutcome  runtime.Outcome `json:"last_runtime_outcome,omitempty"`
	TerminalReason      string          `json:"terminal_reason"`
}

type Clock interface{ Observe() time.Time }
type Waiter interface {
	Wait(context.Context, time.Time) error
}
type RuntimeRunner interface {
	Run(context.Context, runtime.Input) (runtime.Result, error)
}

// EvidenceSink receives the exact validated proposed Service State associated
// with each Event/Evidence pair. It may persist or project the observation but
// cannot influence Service lifecycle decisions.
type EvidenceSink interface {
	Emit(State, Event, Evidence) error
}

// SystemClock and TimerWaiter are replaceable local standard-library adapters.
// Tests use deterministic fakes; neither adapter installs or supervises a service.
type SystemClock struct{}

func (SystemClock) Observe() time.Time { return time.Now().UTC() }

type TimerWaiter struct{}

func (TimerWaiter) Wait(ctx context.Context, at time.Time) error {
	delay := time.Until(at)
	if delay <= 0 {
		return ctx.Err()
	}
	timer := time.NewTimer(delay)
	defer timer.Stop()
	select {
	case <-ctx.Done():
		return ctx.Err()
	case <-timer.C:
		return nil
	}
}

type Service struct {
	Clock  Clock
	Waiter Waiter
	Runner RuntimeRunner
	Sink   EvidenceSink
}

func NewDefinition(serviceID string, interval, cycleTimeout time.Duration) (Definition, error) {
	v := Definition{SchemaName: DefinitionSchema, SchemaVersion: SchemaVersion, ModelVersion: ModelVersion, ServiceID: serviceID, IntervalNS: int64(interval), CycleTimeoutNS: int64(cycleTimeout), StartupMode: StartupImmediate}
	v.ID = definitionID(v)
	return v, ValidateDefinition(v)
}

func NewState(serviceID string) State {
	v := State{SchemaName: StateSchema, SchemaVersion: SchemaVersion, ServiceID: serviceID, Lifecycle: Created}
	v.ID = stateID(v)
	return v
}

func (service Service) Run(ctx context.Context, input Input) (Result, error) {
	if service.Clock == nil || service.Waiter == nil || service.Runner == nil || service.Sink == nil {
		return Result{}, fmt.Errorf("runtime service dependencies are incomplete")
	}
	if err := ValidateInput(input); err != nil {
		return Result{}, err
	}
	m := machine{service: service, input: input, state: input.InitialState, runtimeInput: input.Seed}
	now, err := m.observe()
	if err != nil || now.Before(input.StartedAt) {
		return Result{}, fmt.Errorf("invalid runtime service start observation")
	}
	m.state.Lifecycle = Starting
	m.state.NextNominalAt = input.StartedAt
	m.reidentifyState()
	if !m.emit(EvidenceStartup, now, "", "", "", 0, "service_starting") {
		return m.fail(now, "evidence_sink_failed"), nil
	}
	if ctx.Err() != nil {
		return m.stop(now, "context_cancelled"), nil
	}
	m.state.Lifecycle = Running
	m.reidentifyState()
	nominal := input.StartedAt
	interval := time.Duration(input.Definition.IntervalNS)
	timeout := time.Duration(input.Definition.CycleTimeoutNS)

	for {
		if ctx.Err() != nil {
			return m.stop(m.lastObserved, "context_cancelled"), nil
		}
		if nominal.After(m.lastObserved) {
			if err := service.Waiter.Wait(ctx, nominal); err != nil {
				if ctx.Err() != nil || err == context.Canceled || err == context.DeadlineExceeded {
					return m.stop(m.lastObserved, "context_cancelled"), nil
				}
				return m.fail(m.lastObserved, "waiter_failed"), nil
			}
			now, err = m.observe()
			if err != nil || now.Before(nominal) {
				return m.fail(m.lastObserved, "clock_contract_failed"), nil
			}
		}
		if m.state.Sequence >= MaxCycleSequence {
			return m.fail(m.lastObserved, "cycle_limit_exceeded"), nil
		}
		sequence := m.state.Sequence + 1
		cycleID := stableID("runtime-service-cycle", struct {
			ServiceID string    `json:"service_id"`
			Sequence  uint64    `json:"sequence"`
			Nominal   time.Time `json:"nominal"`
		}{input.Definition.ServiceID, sequence, nominal})
		if !m.emit(EvidenceCycleScheduled, m.lastObserved, cycleID, "", "", 0, "cycle_scheduled") {
			return m.fail(m.lastObserved, "evidence_sink_failed"), nil
		}
		m.state.Sequence = sequence
		m.state.CyclesStarted++
		m.state.ActiveCycleID = cycleID
		m.state.NextNominalAt = nominal
		m.reidentifyState()
		if !m.emit(EvidenceCycleStarted, m.lastObserved, cycleID, "", "", 0, "cycle_started") {
			return m.fail(m.lastObserved, "evidence_sink_failed"), nil
		}
		deadline, ok := addDuration(nominal, timeout)
		if !ok {
			return m.fail(m.lastObserved, "time_overflow"), nil
		}
		executionContext, err := runtime.NewExecutionContext(cycleID, input.Definition.ServiceID, nominal, deadline)
		if err != nil {
			return m.fail(m.lastObserved, "runtime_context_failed"), nil
		}
		m.runtimeInput.Context = executionContext
		cycleContext, cancel := context.WithDeadline(ctx, deadline)
		runtimeResult, runErr := service.Runner.Run(cycleContext, m.runtimeInput)
		cancel()
		if runErr != nil {
			return m.fail(m.lastObserved, "runtime_runner_failed"), nil
		}
		if err := runtime.ValidateResult(runtimeResult); err != nil || runtimeResult.Context.ID != cycleID {
			return m.fail(m.lastObserved, "runtime_result_invalid"), nil
		}
		m.runtimeInput.PreviousState = runtimeResult.NextState
		m.runtimeInput.PreviousAlertState = runtimeResult.FinalAlertState
		m.runtimeInput.PreviousNotificationQueue = runtimeResult.FinalNotificationQueue
		m.state.CyclesCompleted++
		m.state.ActiveCycleID = ""
		m.state.LastRuntimeResultID = runtimeResult.ID
		m.state.LastRuntimeOutcome = runtimeResult.Outcome
		m.reidentifyState()
		now, err = m.observe()
		if err != nil {
			return m.fail(m.lastObserved, "clock_contract_failed"), nil
		}
		if !m.emit(EvidenceCycleCompleted, now, cycleID, runtimeResult.ID, runtimeResult.Outcome, 0, "cycle_completed") {
			return m.fail(now, "evidence_sink_failed"), nil
		}
		if ctx.Err() != nil {
			return m.stop(now, "context_cancelled"), nil
		}
		next, ok := addDuration(nominal, interval)
		if !ok {
			return m.fail(now, "time_overflow"), nil
		}
		nominal = next
		if nominal.Before(now) {
			delta := now.Sub(nominal)
			missed := uint64((delta-1)/interval) + 1
			if missed > MaxCycleSequence-m.state.IntervalsMissed {
				return m.fail(now, "cycle_limit_exceeded"), nil
			}
			if missed > uint64(math.MaxInt64/int64(interval)) {
				return m.fail(now, "time_overflow"), nil
			}
			advance := time.Duration(missed) * interval
			nominal, ok = addDuration(nominal, advance)
			if !ok {
				return m.fail(now, "time_overflow"), nil
			}
			m.state.IntervalsMissed += missed
			m.state.NextNominalAt = nominal
			m.reidentifyState()
			if !m.emit(EvidenceIntervalsMissed, now, "", "", "", missed, "nominal_intervals_elapsed") {
				return m.fail(now, "evidence_sink_failed"), nil
			}
		}
		m.state.NextNominalAt = nominal
		m.reidentifyState()
	}
}

type machine struct {
	service      Service
	input        Input
	state        State
	runtimeInput runtime.Input
	evidenceSeq  uint64
	lastObserved time.Time
}

func (m *machine) observe() (time.Time, error) {
	now := m.service.Clock.Observe().UTC().Truncate(0)
	if now.IsZero() || (!m.lastObserved.IsZero() && now.Before(m.lastObserved)) {
		return time.Time{}, fmt.Errorf("non-monotonic service clock")
	}
	m.lastObserved = now
	return now, nil
}

func (m *machine) emit(kind EvidenceKind, at time.Time, cycleID, resultID string, outcome runtime.Outcome, missed uint64, reason string) bool {
	if m.evidenceSeq >= MaxCycleSequence {
		return false
	}
	m.evidenceSeq++
	event := Event{SchemaName: EventSchema, SchemaVersion: SchemaVersion, Sequence: m.evidenceSeq, Kind: kind, At: at.UTC(), ServiceID: m.input.Definition.ServiceID, CycleID: cycleID, RuntimeResultID: resultID, RuntimeOutcome: outcome, MissedCount: missed, ReasonToken: reason}
	event.ID = eventID(event)
	references := []string{}
	if resultID != "" {
		references = append(references, resultID)
	}
	if cycleID != "" {
		references = append(references, cycleID)
	}
	evidence := Evidence{SchemaName: EvidenceSchema, SchemaVersion: SchemaVersion, EventID: event.ID, ServiceID: m.input.Definition.ServiceID, ReferenceIDs: references}
	evidence.ID = evidenceID(evidence)
	return ValidateState(m.state) == nil && ValidateEvent(event) == nil && ValidateEvidence(evidence) == nil && m.service.Sink.Emit(m.state, event, evidence) == nil
}

func (m *machine) stop(at time.Time, reason string) Result {
	m.state.Lifecycle = Stopping
	m.state.ActiveCycleID = ""
	m.reidentifyState()
	if !m.emit(EvidenceShutdownRequested, at, "", m.state.LastRuntimeResultID, m.state.LastRuntimeOutcome, 0, reason) {
		return m.fail(at, "evidence_sink_failed")
	}
	m.state.Lifecycle = Stopped
	m.state.NextNominalAt = time.Time{}
	m.reidentifyState()
	if !m.emit(EvidenceShutdownCompleted, at, "", m.state.LastRuntimeResultID, m.state.LastRuntimeOutcome, 0, "service_stopped") {
		return m.fail(at, "evidence_sink_failed")
	}
	return m.result(at, reason)
}

func (m *machine) fail(at time.Time, reason string) Result {
	m.state.Lifecycle = Failed
	m.state.ActiveCycleID = ""
	m.state.NextNominalAt = time.Time{}
	m.reidentifyState()
	_ = m.emit(EvidenceTerminalFailure, at, "", m.state.LastRuntimeResultID, m.state.LastRuntimeOutcome, 0, reason)
	return m.result(at, reason)
}

func (m *machine) result(at time.Time, reason string) Result {
	v := Result{SchemaName: ResultSchema, SchemaVersion: SchemaVersion, ModelVersion: ModelVersion, ServiceID: m.input.Definition.ServiceID, StartedAt: m.input.StartedAt, CompletedAt: at.UTC(), FinalState: m.state, LastRuntimeResultID: m.state.LastRuntimeResultID, LastRuntimeOutcome: m.state.LastRuntimeOutcome, TerminalReason: reason}
	v.ID = resultID(v)
	return v
}

func (m *machine) reidentifyState() { m.state.ID = stateID(m.state) }

func ValidateDefinition(v Definition) error {
	if v.SchemaName != DefinitionSchema || v.SchemaVersion != SchemaVersion || v.ModelVersion != ModelVersion || !token(v.ServiceID) || v.StartupMode != StartupImmediate || v.IntervalNS <= 0 || time.Duration(v.IntervalNS) > MaxInterval || v.CycleTimeoutNS <= 0 || time.Duration(v.CycleTimeoutNS) > MaxCycleTimeout || v.ID != definitionID(v) {
		return fmt.Errorf("invalid runtime service definition")
	}
	return nil
}

func ValidateState(v State) error {
	if v.SchemaName != StateSchema || v.SchemaVersion != SchemaVersion || !token(v.ServiceID) || !validLifecycle(v.Lifecycle) || v.Sequence > MaxCycleSequence || v.CyclesStarted > v.Sequence || v.CyclesCompleted > v.CyclesStarted || v.IntervalsMissed > MaxCycleSequence || len(v.ActiveCycleID) > MaxEvidenceString || len(v.LastRuntimeResultID) > MaxEvidenceString {
		return fmt.Errorf("invalid runtime service state")
	}
	hasNominal := v.Lifecycle == Starting || v.Lifecycle == Running || v.Lifecycle == Stopping
	if hasNominal != !v.NextNominalAt.IsZero() || (v.ActiveCycleID != "" && v.Lifecycle != Running) || (v.Lifecycle == Created && (v.Sequence != 0 || v.CyclesStarted != 0 || v.CyclesCompleted != 0)) || v.ID != stateID(v) {
		return fmt.Errorf("invalid runtime service state lifecycle")
	}
	return nil
}

func ValidateInput(v Input) error {
	if err := ValidateDefinition(v.Definition); err != nil {
		return err
	}
	if v.StartedAt.IsZero() || v.StartedAt.Location() != time.UTC {
		return fmt.Errorf("invalid runtime service start time")
	}
	if err := ValidateState(v.InitialState); err != nil || v.InitialState.Lifecycle != Created || v.InitialState.ServiceID != v.Definition.ServiceID {
		return fmt.Errorf("invalid initial runtime service state")
	}
	if err := runtime.ValidateInput(v.Seed); err != nil {
		return fmt.Errorf("invalid runtime seed: %w", err)
	}
	return nil
}

func ValidateEvent(v Event) error {
	if v.SchemaName != EventSchema || v.SchemaVersion != SchemaVersion || v.Sequence == 0 || v.Sequence > MaxCycleSequence || !validEvidenceKind(v.Kind) || v.At.IsZero() || v.At.Location() != time.UTC || !token(v.ServiceID) || !token(v.ReasonToken) || len(v.CycleID) > MaxEvidenceString || len(v.RuntimeResultID) > MaxEvidenceString || v.ID != eventID(v) {
		return fmt.Errorf("invalid runtime service event")
	}
	if (v.Kind == EvidenceIntervalsMissed) != (v.MissedCount > 0) || (v.RuntimeOutcome != "" && !validRuntimeOutcome(v.RuntimeOutcome)) {
		return fmt.Errorf("invalid runtime service evidence content")
	}
	return nil
}

func ValidateEvidence(v Evidence) error {
	if v.SchemaName != EvidenceSchema || v.SchemaVersion != SchemaVersion || len(v.EventID) == 0 || len(v.EventID) > MaxEvidenceString || !token(v.ServiceID) || v.ReferenceIDs == nil || len(v.ReferenceIDs) > 2 || !sortedUnique(v.ReferenceIDs) || v.ID != evidenceID(v) {
		return fmt.Errorf("invalid runtime service evidence")
	}
	return nil
}

func ValidateResult(v Result) error {
	if v.SchemaName != ResultSchema || v.SchemaVersion != SchemaVersion || v.ModelVersion != ModelVersion || !token(v.ServiceID) || v.StartedAt.IsZero() || v.CompletedAt.Before(v.StartedAt) || !token(v.TerminalReason) || len(v.LastRuntimeResultID) > MaxEvidenceString || (v.FinalState.Lifecycle != Stopped && v.FinalState.Lifecycle != Failed) || v.FinalState.ServiceID != v.ServiceID || v.LastRuntimeResultID != v.FinalState.LastRuntimeResultID || v.LastRuntimeOutcome != v.FinalState.LastRuntimeOutcome || v.ID != resultID(v) {
		return fmt.Errorf("invalid runtime service result")
	}
	if err := ValidateState(v.FinalState); err != nil {
		return err
	}
	return nil
}

func MarshalCanonical(v Result) ([]byte, error) {
	if err := ValidateResult(v); err != nil {
		return nil, err
	}
	return json.Marshal(v)
}
func MarshalDefinition(v Definition) ([]byte, error) {
	if err := ValidateDefinition(v); err != nil {
		return nil, err
	}
	return json.Marshal(v)
}
func MarshalState(v State) ([]byte, error) {
	if err := ValidateState(v); err != nil {
		return nil, err
	}
	return json.Marshal(v)
}
func MarshalInput(v Input) ([]byte, error) {
	if err := ValidateInput(v); err != nil {
		return nil, err
	}
	return json.Marshal(v)
}
func MarshalEvent(v Event) ([]byte, error) {
	if err := ValidateEvent(v); err != nil {
		return nil, err
	}
	return json.Marshal(v)
}
func MarshalEvidence(v Evidence) ([]byte, error) {
	if err := ValidateEvidence(v); err != nil {
		return nil, err
	}
	return json.Marshal(v)
}
func DecodeDefinition(data []byte) (Definition, error) {
	var v Definition
	if err := strictDecode(data, &v); err != nil {
		return Definition{}, err
	}
	return v, ValidateDefinition(v)
}
func DecodeState(data []byte) (State, error) {
	var v State
	if err := strictDecode(data, &v); err != nil {
		return State{}, err
	}
	return v, ValidateState(v)
}
func DecodeInput(data []byte) (Input, error) {
	var v Input
	if err := strictDecode(data, &v); err != nil {
		return Input{}, err
	}
	return v, ValidateInput(v)
}
func DecodeEvent(data []byte) (Event, error) {
	var v Event
	if err := strictDecode(data, &v); err != nil {
		return Event{}, err
	}
	return v, ValidateEvent(v)
}
func DecodeEvidence(data []byte) (Evidence, error) {
	var v Evidence
	if err := strictDecode(data, &v); err != nil {
		return Evidence{}, err
	}
	return v, ValidateEvidence(v)
}
func DecodeResult(data []byte) (Result, error) {
	var v Result
	if err := strictDecode(data, &v); err != nil {
		return Result{}, err
	}
	return v, ValidateResult(v)
}

func definitionID(v Definition) string { v.ID = ""; return stableID("runtime-service-definition", v) }
func stateID(v State) string           { v.ID = ""; return stableID("runtime-service-state", v) }
func eventID(v Event) string           { v.ID = ""; return stableID("runtime-service-event", v) }
func evidenceID(v Evidence) string     { v.ID = ""; return stableID("runtime-service-evidence", v) }
func resultID(v Result) string         { v.ID = ""; return stableID("runtime-service-result", v) }
func stableID(prefix string, v any) string {
	data, _ := json.Marshal(v)
	sum := sha256.Sum256(data)
	return prefix + ":" + hex.EncodeToString(sum[:])
}
func strictDecode(data []byte, target any) error {
	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(target); err != nil {
		return err
	}
	var extra any
	if err := decoder.Decode(&extra); err != io.EOF {
		return fmt.Errorf("trailing runtime service data")
	}
	return nil
}
func addDuration(at time.Time, value time.Duration) (time.Time, bool) {
	result := at.Add(value)
	return result, (value > 0 && result.After(at))
}
func token(v string) bool {
	if v == "" || len(v) > MaxEvidenceString {
		return false
	}
	for i, r := range v {
		if (r >= 'a' && r <= 'z') || (r >= '0' && r <= '9') || (i > 0 && (r == '.' || r == '_' || r == ':' || r == '-')) {
			continue
		}
		return false
	}
	return true
}
func validLifecycle(v Lifecycle) bool {
	return v == Created || v == Starting || v == Running || v == Stopping || v == Stopped || v == Failed
}
func validEvidenceKind(v EvidenceKind) bool {
	return v == EvidenceStartup || v == EvidenceCycleScheduled || v == EvidenceCycleStarted || v == EvidenceCycleCompleted || v == EvidenceIntervalsMissed || v == EvidenceShutdownRequested || v == EvidenceShutdownCompleted || v == EvidenceTerminalFailure
}
func validRuntimeOutcome(v runtime.Outcome) bool {
	return v == runtime.Completed || v == runtime.Partial || v == runtime.Failed || v == runtime.Cancelled || v == runtime.TimedOut
}
func sortedUnique(values []string) bool {
	for i, value := range values {
		if value == "" || (i > 0 && values[i-1] >= value) {
			return false
		}
	}
	return true
}
