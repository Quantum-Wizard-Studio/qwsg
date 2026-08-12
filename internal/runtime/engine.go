// Package runtime coordinates one explicit, bounded QWSG execution cycle. It
// owns orchestration only; Scheduler, Pipeline, Alert, and Notification retain
// their canonical decisions and state semantics.
package runtime

import (
	"bytes"
	"context"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"sort"
	"time"

	"quantumwizard.hu/qwsg/internal/alert"
	"quantumwizard.hu/qwsg/internal/command"
	"quantumwizard.hu/qwsg/internal/configuration"
	"quantumwizard.hu/qwsg/internal/health"
	"quantumwizard.hu/qwsg/internal/notification"
	"quantumwizard.hu/qwsg/internal/pipeline"
	"quantumwizard.hu/qwsg/internal/policy"
	"quantumwizard.hu/qwsg/internal/report"
	"quantumwizard.hu/qwsg/internal/rule"
	"quantumwizard.hu/qwsg/internal/scheduler"
)

const (
	SchemaVersion         = "1.0"
	ModelVersion          = "1.0"
	StateSchema           = "qwsg.runtime-state"
	ContextSchema         = "qwsg.runtime-execution-context"
	ResultSchema          = "qwsg.runtime-result"
	EventSchema           = "qwsg.runtime-event"
	MaxEvents             = 64
	MaxEvidenceReferences = 8192
)

type Lifecycle string

const (
	Idle    Lifecycle = "idle"
	Running Lifecycle = "running"
)

type Outcome string

const (
	Completed Outcome = "completed"
	Partial   Outcome = "partial"
	Failed    Outcome = "failed"
	Cancelled Outcome = "cancelled"
	TimedOut  Outcome = "timed_out"
)

type Component string

const (
	SchedulerComponent            Component = "scheduler"
	AlertComponent                Component = "alert"
	NotificationPlanComponent     Component = "notification_plan"
	NotificationDeliveryComponent Component = "notification_delivery"
)

type ComponentStatus string

const (
	ComponentCompleted ComponentStatus = "completed"
	ComponentFailed    ComponentStatus = "failed"
	ComponentSkipped   ComponentStatus = "skipped"
)

type Clock interface{ Observe() time.Time }
type SchedulerRunner interface {
	Run(context.Context) (scheduler.CycleResult, error)
}

type ExecutionContext struct {
	SchemaName    string    `json:"schema_name"`
	SchemaVersion string    `json:"schema_version"`
	ID            string    `json:"id"`
	InitiatorRef  string    `json:"initiator_ref"`
	StartedAt     time.Time `json:"started_at"`
	Deadline      time.Time `json:"deadline"`
}

type State struct {
	SchemaName           string    `json:"schema_name"`
	SchemaVersion        string    `json:"schema_version"`
	ID                   string    `json:"id"`
	Lifecycle            Lifecycle `json:"lifecycle"`
	ActiveCycleID        string    `json:"active_cycle_id,omitempty"`
	LastCompletedCycleID string    `json:"last_completed_cycle_id,omitempty"`
}

type Input struct {
	Context                   ExecutionContext          `json:"context"`
	Configuration             configuration.Effective   `json:"configuration"`
	PreviousState             State                     `json:"previous_state"`
	PreviousAlertState        alert.State               `json:"previous_alert_state"`
	PreviousNotificationQueue notification.QueueState   `json:"previous_notification_queue"`
	AlertEvidenceTTLNS        int64                     `json:"alert_evidence_ttl_ns"`
	Acknowledgements          []alert.Acknowledgement   `json:"acknowledgements"`
	Suppressions              []alert.SuppressionWindow `json:"suppressions"`
	NotificationPolicy        notification.Policy       `json:"notification_policy"`
}

type Event struct {
	SchemaName    string    `json:"schema_name"`
	SchemaVersion string    `json:"schema_version"`
	ID            string    `json:"id"`
	Sequence      int       `json:"sequence"`
	At            time.Time `json:"at"`
	Component     Component `json:"component"`
	Kind          string    `json:"kind"`
	ReasonToken   string    `json:"reason_token"`
}

type Evidence struct {
	Component Component `json:"component"`
	Contract  string    `json:"contract"`
	RecordIDs []string  `json:"record_ids"`
}

type ComponentResult struct {
	Component    Component       `json:"component"`
	Status       ComponentStatus `json:"status"`
	FailureToken string          `json:"failure_token,omitempty"`
	ReferenceIDs []string        `json:"reference_ids"`
}

type Result struct {
	SchemaName             string                    `json:"schema_name"`
	SchemaVersion          string                    `json:"schema_version"`
	ModelVersion           string                    `json:"model_version"`
	ID                     string                    `json:"id"`
	Context                ExecutionContext          `json:"context"`
	Outcome                Outcome                   `json:"outcome"`
	Components             []ComponentResult         `json:"components"`
	Events                 []Event                   `json:"events"`
	Evidence               []Evidence                `json:"evidence"`
	SchedulerResultID      string                    `json:"scheduler_result_id,omitempty"`
	FinalSchedulerState    *scheduler.State          `json:"final_scheduler_state,omitempty"`
	AlertResults           []alert.Result            `json:"alert_results"`
	NotificationPlan       *notification.Plan        `json:"notification_plan,omitempty"`
	NotificationCycle      *notification.CycleResult `json:"notification_cycle,omitempty"`
	FinalAlertState        alert.State               `json:"final_alert_state"`
	FinalNotificationQueue notification.QueueState   `json:"final_notification_queue"`
	NextState              State                     `json:"next_state"`
}

type Coordinator struct {
	Scheduler SchedulerRunner
	Clock     Clock
	Providers notification.Registry
}

func NewState() State {
	v := State{SchemaName: StateSchema, SchemaVersion: SchemaVersion, Lifecycle: Idle}
	v.ID = stateID(v)
	return v
}

// Run performs at most one call to each Scheduler and Notification cycle. An
// operational component failure is returned as a validated terminal Result;
// error is reserved for invalid contracts or an inability to prove the result.
func (c Coordinator) Run(ctx context.Context, input Input) (Result, error) {
	if c.Scheduler == nil || c.Clock == nil {
		return Result{}, fmt.Errorf("runtime dependencies are incomplete")
	}
	if err := ValidateInput(input); err != nil {
		return Result{}, err
	}
	bounded, cancel := context.WithDeadline(ctx, input.Context.Deadline)
	defer cancel()
	b := builder{input: input, clock: c.Clock, result: Result{SchemaName: ResultSchema, SchemaVersion: SchemaVersion, ModelVersion: ModelVersion, Context: input.Context, Components: []ComponentResult{}, Events: []Event{}, Evidence: []Evidence{}, AlertResults: []alert.Result{}, FinalAlertState: input.PreviousAlertState, FinalNotificationQueue: input.PreviousNotificationQueue}}
	b.event(SchedulerComponent, "cycle_started", "runtime_cycle_started")
	if stop := b.contextStop(bounded); stop != "" {
		return b.finish(stop), nil
	}
	scheduled, err := c.Scheduler.Run(bounded)
	if err != nil {
		b.component(SchedulerComponent, ComponentFailed, "scheduler_cycle_failed", nil)
		return b.finish(Failed), nil
	}
	if err := scheduler.ValidateCycleResult(scheduled); err != nil {
		return Result{}, fmt.Errorf("invalid scheduler result: %w", err)
	}
	b.result.SchedulerResultID = scheduled.ID
	finalSchedulerState := scheduled.FinalState
	b.result.FinalSchedulerState = &finalSchedulerState
	b.component(SchedulerComponent, ComponentCompleted, "", []string{scheduled.ID})
	b.evidence(SchedulerComponent, scheduler.CycleResultSchema, []string{scheduled.ID})

	alertState := input.PreviousAlertState
	alertRecords := []alert.Record{}
	evaluate := func(projection projected, schedulerEvaluation *scheduler.Evaluation) error {
		if stop := b.contextStop(bounded); stop != "" {
			return contextError(stop)
		}
		at, observeErr := b.observe()
		if observeErr != nil {
			return observeErr
		}
		ai := alert.Input{SchemaName: alert.InputSchema, SchemaVersion: alert.SchemaVersion, EvaluatedAt: at, EvidenceTTLNS: input.AlertEvidenceTTLNS, Configuration: &input.Configuration, Health: projection.health, Rules: projection.rules, Policies: projection.policies, Scheduler: schedulerEvaluation, PolicyReport: projection.policyReport, PreviousState: alertState, Acknowledgements: input.Acknowledgements, Suppressions: input.Suppressions}
		ar, evaluateErr := alert.Evaluate(ai)
		if evaluateErr != nil {
			return evaluateErr
		}
		alertState = ar.NextState
		b.result.FinalAlertState = alertState
		b.result.AlertResults = append(b.result.AlertResults, ar)
		alertRecords = append(alertRecords, ar.Records...)
		b.evidence(AlertComponent, alert.ResultSchema, []string{ar.ID})
		return nil
	}
	if err := evaluate(projected{}, &scheduled.Evaluation); err != nil {
		return b.alertFailure(err), nil
	}
	for _, trace := range scheduled.Traces {
		if trace.FailureCode != "" {
			continue
		}
		projection, err := project(trace)
		if err != nil {
			return Result{}, err
		}
		if projection.empty() {
			continue
		}
		if err := evaluate(projection, nil); err != nil {
			return b.alertFailure(err), nil
		}
	}
	b.result.FinalAlertState = alertState
	b.component(AlertComponent, ComponentCompleted, "", resultIDs(b.result.AlertResults))
	sort.Slice(alertRecords, func(i, j int) bool { return alertRecords[i].ID < alertRecords[j].ID })
	if stop := b.contextStop(bounded); stop != "" {
		return b.finish(stop), nil
	}
	at, err := b.observe()
	if err != nil {
		return b.finish(classifyError(err)), nil
	}
	plan, err := notification.PlanDeliveries(notification.PlanningInput{SchemaName: notification.InputSchema, SchemaVersion: notification.SchemaVersion, EvaluatedAt: at, Alerts: alertRecords, Policy: input.NotificationPolicy, PreviousQueue: input.PreviousNotificationQueue})
	if err != nil {
		b.component(NotificationPlanComponent, ComponentFailed, "notification_planning_failed", nil)
		return b.finish(Partial), nil
	}
	b.result.NotificationPlan = &plan
	b.result.FinalNotificationQueue = plan.NextQueue
	b.component(NotificationPlanComponent, ComponentCompleted, "", []string{plan.ID})
	b.evidence(NotificationPlanComponent, notification.PlanSchema, []string{plan.ID})
	if len(plan.Requests) == 0 {
		b.component(NotificationDeliveryComponent, ComponentSkipped, "no_delivery_requests", []string{})
		return b.finish(Completed), nil
	}
	if stop := b.contextStop(bounded); stop != "" {
		return b.finish(stop), nil
	}
	at, err = b.observe()
	if err != nil {
		return b.finish(classifyError(err)), nil
	}
	cycle, err := notification.ExecuteCycle(bounded, plan, c.Providers, at)
	if err != nil {
		b.component(NotificationDeliveryComponent, ComponentFailed, "notification_delivery_failed", nil)
		return b.finish(Partial), nil
	}
	b.result.NotificationCycle = &cycle
	b.result.FinalNotificationQueue = cycle.NextQueue
	b.component(NotificationDeliveryComponent, ComponentCompleted, "", []string{cycle.ID})
	b.evidence(NotificationDeliveryComponent, notification.CycleResultSchema, []string{cycle.ID})
	return b.finish(Completed), nil
}

type projected struct {
	health       *health.Result
	rules        *rule.Result
	policies     *policy.Result
	policyReport *report.PolicyReport
}

func (p projected) empty() bool {
	return p.health == nil && p.rules == nil && p.policies == nil && p.policyReport == nil
}

func project(trace scheduler.ExecutionTrace) (projected, error) {
	plan, err := command.PlanDefinition(trace.Definition)
	if err != nil {
		return projected{}, fmt.Errorf("invalid traced command: %w", err)
	}
	if err := pipeline.ValidateExecution(trace.Execution, plan); err != nil {
		return projected{}, fmt.Errorf("invalid traced execution: %w", err)
	}
	var p projected
	for _, stage := range trace.Execution.Stages {
		switch stage.Stage {
		case command.Health:
			v, ok := stage.Value.(health.Result)
			if !ok {
				return projected{}, fmt.Errorf("health stage has non-canonical value")
			}
			p.health = &v
		case command.Rule:
			v, ok := stage.Value.(rule.Result)
			if !ok {
				return projected{}, fmt.Errorf("rule stage has non-canonical value")
			}
			p.rules = &v
		case command.Policy:
			v, ok := stage.Value.(policy.Result)
			if !ok {
				return projected{}, fmt.Errorf("policy stage has non-canonical value")
			}
			p.policies = &v
		case command.Report:
			v, ok := stage.Value.(report.PolicyReport)
			if !ok {
				return projected{}, fmt.Errorf("report stage has non-canonical value")
			}
			p.policyReport = &v
		}
	}
	return p, nil
}

type builder struct {
	input  Input
	clock  Clock
	result Result
}

func (b *builder) observe() (time.Time, error) {
	at := b.clock.Observe().UTC()
	if at.IsZero() || at.Before(b.input.Context.StartedAt) {
		return time.Time{}, fmt.Errorf("runtime_clock_invalid")
	}
	if !at.Before(b.input.Context.Deadline) {
		return time.Time{}, context.DeadlineExceeded
	}
	return at, nil
}
func (b *builder) contextStop(ctx context.Context) Outcome {
	select {
	case <-ctx.Done():
		return classifyError(ctx.Err())
	default:
		return ""
	}
}
func (b *builder) event(component Component, kind, reason string) {
	at, err := b.observe()
	if err != nil {
		at = b.input.Context.StartedAt
	}
	e := Event{SchemaName: EventSchema, SchemaVersion: SchemaVersion, Sequence: len(b.result.Events) + 1, At: at, Component: component, Kind: kind, ReasonToken: reason}
	e.ID = stableID("runtime-event", e)
	b.result.Events = append(b.result.Events, e)
}
func (b *builder) component(component Component, status ComponentStatus, failure string, refs []string) {
	if refs == nil {
		refs = []string{}
	}
	sort.Strings(refs)
	b.result.Components = append(b.result.Components, ComponentResult{Component: component, Status: status, FailureToken: failure, ReferenceIDs: refs})
}
func (b *builder) evidence(component Component, contract string, ids []string) {
	sort.Strings(ids)
	b.result.Evidence = append(b.result.Evidence, Evidence{Component: component, Contract: contract, RecordIDs: ids})
}
func (b *builder) alertFailure(err error) Result {
	token := "alert_evaluation_failed"
	outcome := Partial
	if v := classifyError(err); v == Cancelled || v == TimedOut {
		outcome = v
		token = string(v)
	}
	b.component(AlertComponent, ComponentFailed, token, nil)
	return b.finish(outcome)
}
func (b *builder) finish(outcome Outcome) Result {
	b.result.Outcome = outcome
	next := State{SchemaName: StateSchema, SchemaVersion: SchemaVersion, Lifecycle: Idle, LastCompletedCycleID: b.input.Context.ID}
	next.ID = stateID(next)
	b.result.NextState = next
	b.result.ID = resultID(b.result)
	return b.result
}

func ValidateInput(v Input) error {
	if v.Context.SchemaName != ContextSchema || v.Context.SchemaVersion != SchemaVersion || v.Context.ID == "" || v.Context.InitiatorRef == "" || v.Context.StartedAt.IsZero() || !v.Context.StartedAt.Before(v.Context.Deadline) {
		return fmt.Errorf("invalid runtime execution context")
	}
	if err := configuration.ValidateEffective(v.Configuration); err != nil {
		return err
	}
	if err := ValidateState(v.PreviousState); err != nil || v.PreviousState.Lifecycle != Idle {
		return fmt.Errorf("runtime must start idle")
	}
	if err := alert.ValidateState(v.PreviousAlertState); err != nil || v.PreviousAlertState.ConfigurationID != v.Configuration.ID {
		return fmt.Errorf("invalid runtime alert state")
	}
	if err := notification.ValidateQueue(v.PreviousNotificationQueue); err != nil {
		return err
	}
	if err := notification.ValidatePolicy(v.NotificationPolicy); err != nil {
		return err
	}
	if v.Acknowledgements == nil || v.Suppressions == nil || v.AlertEvidenceTTLNS <= 0 || time.Duration(v.AlertEvidenceTTLNS) > alert.MaxEvidenceTTL {
		return fmt.Errorf("invalid runtime alert controls")
	}
	return nil
}

func ValidateState(v State) error {
	if v.SchemaName != StateSchema || v.SchemaVersion != SchemaVersion || (v.Lifecycle != Idle && v.Lifecycle != Running) || (v.Lifecycle == Idle && v.ActiveCycleID != "") || (v.Lifecycle == Running && v.ActiveCycleID == "") || v.ID != stateID(v) {
		return fmt.Errorf("invalid runtime state")
	}
	return nil
}
func ValidateResult(v Result) error {
	if v.SchemaName != ResultSchema || v.SchemaVersion != SchemaVersion || v.ModelVersion != ModelVersion || v.ID != resultID(v) || v.Components == nil || v.Events == nil || v.Evidence == nil || v.AlertResults == nil || len(v.Events) > MaxEvents || len(v.Evidence) > MaxEvidenceReferences {
		return fmt.Errorf("invalid runtime result")
	}
	if v.Outcome != Completed && v.Outcome != Partial && v.Outcome != Failed && v.Outcome != Cancelled && v.Outcome != TimedOut {
		return fmt.Errorf("invalid runtime outcome")
	}
	if err := ValidateState(v.NextState); err != nil || v.NextState.Lifecycle != Idle || v.NextState.LastCompletedCycleID != v.Context.ID {
		return fmt.Errorf("runtime did not return idle")
	}
	if v.Context.SchemaName != ContextSchema || v.Context.SchemaVersion != SchemaVersion || v.Context.ID == "" || v.Context.InitiatorRef == "" || v.Context.StartedAt.IsZero() || !v.Context.StartedAt.Before(v.Context.Deadline) {
		return fmt.Errorf("invalid runtime result context")
	}
	componentOrder := map[Component]int{SchedulerComponent: 1, AlertComponent: 2, NotificationPlanComponent: 3, NotificationDeliveryComponent: 4}
	lastComponent := 0
	for _, item := range v.Components {
		order := componentOrder[item.Component]
		if order == 0 || order <= lastComponent || !validComponentResult(item) || item.ReferenceIDs == nil || !sortedUniqueStrings(item.ReferenceIDs) {
			return fmt.Errorf("invalid runtime component result")
		}
		lastComponent = order
	}
	for i, event := range v.Events {
		copy := event
		copy.ID = ""
		if event.SchemaName != EventSchema || event.SchemaVersion != SchemaVersion || event.Sequence != i+1 || event.ID != stableID("runtime-event", copy) || event.At.Before(v.Context.StartedAt) || !event.At.Before(v.Context.Deadline) {
			return fmt.Errorf("invalid runtime event")
		}
	}
	for _, evidence := range v.Evidence {
		if componentOrder[evidence.Component] == 0 || evidence.Contract == "" || evidence.RecordIDs == nil || !sortedUniqueStrings(evidence.RecordIDs) {
			return fmt.Errorf("invalid runtime evidence")
		}
	}
	if v.FinalSchedulerState != nil {
		if v.SchedulerResultID == "" {
			return fmt.Errorf("scheduler state lacks result reference")
		}
		if err := scheduler.ValidateState(*v.FinalSchedulerState); err != nil {
			return err
		}
	} else if v.SchedulerResultID != "" {
		return fmt.Errorf("scheduler result lacks final state")
	}
	for _, item := range v.AlertResults {
		if err := alert.ValidateResult(item); err != nil {
			return err
		}
	}
	if err := alert.ValidateState(v.FinalAlertState); err != nil {
		return err
	}
	if err := notification.ValidateQueue(v.FinalNotificationQueue); err != nil {
		return err
	}
	if v.NotificationPlan != nil {
		if err := notification.ValidatePlan(*v.NotificationPlan); err != nil {
			return err
		}
	}
	if v.NotificationCycle != nil {
		if err := notification.ValidateCycleResult(*v.NotificationCycle); err != nil {
			return err
		}
	}
	return nil
}

func validComponentResult(v ComponentResult) bool {
	if v.Status == ComponentCompleted {
		return v.FailureToken == ""
	}
	allowed := map[ComponentStatus]map[Component]map[string]bool{
		ComponentFailed: {
			SchedulerComponent:            {"scheduler_cycle_failed": true},
			AlertComponent:                {"alert_evaluation_failed": true, string(Cancelled): true, string(TimedOut): true},
			NotificationPlanComponent:     {"notification_planning_failed": true},
			NotificationDeliveryComponent: {"notification_delivery_failed": true, "notification_cycle_failed": true},
		},
		ComponentSkipped: {
			NotificationDeliveryComponent: {"no_delivery_requests": true},
		},
	}
	return allowed[v.Status][v.Component][v.FailureToken]
}
func MarshalCanonical(v Result) ([]byte, error) {
	if err := ValidateResult(v); err != nil {
		return nil, err
	}
	return json.Marshal(v)
}
func DecodeInput(data []byte) (Input, error) {
	var v Input
	if err := strictDecode(data, &v); err != nil {
		return Input{}, err
	}
	return v, ValidateInput(v)
}
func DecodeResult(data []byte) (Result, error) {
	var v Result
	if err := strictDecode(data, &v); err != nil {
		return Result{}, err
	}
	return v, ValidateResult(v)
}
func NewExecutionContext(id, initiator string, started, deadline time.Time) (ExecutionContext, error) {
	v := ExecutionContext{SchemaName: ContextSchema, SchemaVersion: SchemaVersion, ID: id, InitiatorRef: initiator, StartedAt: started.UTC(), Deadline: deadline.UTC()}
	if v.ID == "" || v.InitiatorRef == "" || v.StartedAt.IsZero() || !v.StartedAt.Before(v.Deadline) {
		return ExecutionContext{}, fmt.Errorf("invalid runtime execution context")
	}
	return v, nil
}

func resultIDs(values []alert.Result) []string {
	ids := make([]string, len(values))
	for i := range values {
		ids[i] = values[i].ID
	}
	sort.Strings(ids)
	return ids
}
func sortedUniqueStrings(values []string) bool {
	for i, v := range values {
		if v == "" || (i > 0 && values[i-1] >= v) {
			return false
		}
	}
	return true
}
func contextError(v Outcome) error {
	if v == TimedOut {
		return context.DeadlineExceeded
	}
	return context.Canceled
}
func classifyError(err error) Outcome {
	if err == context.DeadlineExceeded {
		return TimedOut
	}
	if err == context.Canceled {
		return Cancelled
	}
	return Failed
}
func stateID(v State) string   { v.ID = ""; return stableID("runtime-state", v) }
func resultID(v Result) string { v.ID = ""; return stableID("runtime-result", v) }
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
		return fmt.Errorf("trailing runtime data")
	}
	return nil
}
