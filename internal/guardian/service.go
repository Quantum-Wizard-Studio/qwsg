package guardian

import (
	"context"
	"fmt"
	"sync"
	"time"

	"quantumwizard.hu/qwsg/internal/app"
	"quantumwizard.hu/qwsg/internal/command"
	"quantumwizard.hu/qwsg/internal/operatorstate"
	"quantumwizard.hu/qwsg/internal/presentationmodel"
	"quantumwizard.hu/qwsg/internal/runtime"
	"quantumwizard.hu/qwsg/internal/runtimeservice"
	"quantumwizard.hu/qwsg/internal/scheduler"
)

type CapturingScheduler struct {
	Cycle scheduler.Cycle
	mu    sync.Mutex
	last  scheduler.CycleResult
}

type SchedulerClock struct {
	SessionID string
	started   time.Time
}

func NewSchedulerClock(sessionID string) *SchedulerClock {
	return &SchedulerClock{SessionID: sessionID, started: time.Now()}
}

func (v *SchedulerClock) Observe() scheduler.ClockObservation {
	return scheduler.ClockObservation{WallTime: time.Now().UTC(), SessionID: v.SessionID, MonotonicNS: int64(time.Since(v.started))}
}

func (v *CapturingScheduler) Run(ctx context.Context) (scheduler.CycleResult, error) {
	result, err := v.Cycle.Run(ctx)
	if err == nil {
		v.mu.Lock()
		v.last = result
		v.mu.Unlock()
	}
	return result, err
}

func (v *CapturingScheduler) Last() scheduler.CycleResult {
	v.mu.Lock()
	defer v.mu.Unlock()
	return v.last
}

// Take returns the most recent cycle for end-of-cycle publication and releases
// the retained execution graph. A cycle trace can contain two inventories and
// every derived pipeline stage, so retaining it until the next collection
// needlessly overlaps two loaded-host working sets.
func (v *CapturingScheduler) Take() scheduler.CycleResult {
	v.mu.Lock()
	defer v.mu.Unlock()
	result := v.last
	v.last = scheduler.CycleResult{}
	return result
}

type RuntimeRunner struct {
	Coordinator runtime.Coordinator
	Store       *Store
	Checkpoint  *Checkpoint
	mu          sync.Mutex
	last        runtime.Result
}

func (v *RuntimeRunner) Run(ctx context.Context, input runtime.Input) (runtime.Result, error) {
	result, err := v.Coordinator.Run(ctx, input)
	if err != nil {
		return runtime.Result{}, err
	}
	if err := runtime.ValidateResult(result); err != nil {
		return runtime.Result{}, err
	}
	v.mu.Lock()
	checkpoint := *v.Checkpoint
	checkpoint.RuntimeState = result.NextState
	checkpoint.AlertState = result.FinalAlertState
	checkpoint.NotificationQueueState = result.FinalNotificationQueue
	checkpoint.LastCompletedCycleID = result.Context.ID
	checkpoint.LastCompletedAt = result.Context.StartedAt
	if len(result.Events) > 0 {
		checkpoint.LastCompletedAt = result.Events[len(result.Events)-1].At
	}
	if err = v.Store.Save(checkpoint); err == nil {
		*v.Checkpoint = checkpoint
		v.last = result
	}
	v.mu.Unlock()
	if err != nil {
		return runtime.Result{}, fmt.Errorf("guardian checkpoint publication failed")
	}
	return result, nil
}

func (v *RuntimeRunner) Last() runtime.Result {
	v.mu.Lock()
	defer v.mu.Unlock()
	return v.last
}

type Publisher struct {
	Store              *operatorstate.Store
	Scheduler          *CapturingScheduler
	Runner             *RuntimeRunner
	DefinitionID       string
	ApplicationVersion string
	FreshFor           time.Duration
	mu                 sync.Mutex
	hasEngineering     bool
}

func (v *Publisher) Publish(state runtimeservice.State, event runtimeservice.Event) error {
	v.mu.Lock()
	defer v.mu.Unlock()
	if v.hasEngineering && event.Kind != runtimeservice.EvidenceCycleCompleted && event.Kind != runtimeservice.EvidenceShutdownRequested && event.Kind != runtimeservice.EvidenceShutdownCompleted && event.Kind != runtimeservice.EvidenceTerminalFailure {
		return nil
	}
	var (
		overview  presentationmodel.Overview
		execution command.Execution
		runtimeV  runtime.Result
		complete  bool
		err       error
	)
	if event.Kind == runtimeservice.EvidenceCycleCompleted {
		runtimeV = v.Runner.Last()
		cycle := v.Scheduler.Take()
		for index := len(cycle.Traces) - 1; index >= 0; index-- {
			trace := cycle.Traces[index]
			if trace.FailureCode == "" && trace.Definition.Profile == "observe" && trace.Execution.Complete {
				execution, complete = trace.Execution, true
				break
			}
		}
	}
	stages := []string{"runtime_service"}
	if complete {
		overview, err = app.ProjectOperatorEvaluation(execution, event.At, v.FreshFor, &runtimeV, &state)
		stages = []string{"inventory", "snapshot", "compare", "drift", "health", "rule", "policy", "report", "runtime", "runtime_service"}
	} else {
		overview, err = app.ProjectGuardianLifecycle(state, event.At, v.FreshFor)
	}
	if err != nil {
		return fmt.Errorf("guardian projection failed: %w", err)
	}
	current, err := operatorstate.Normalize(operatorstate.State{ObservedAt: event.At, PublishedAt: time.Now().UTC(), FreshUntil: event.At.Add(v.FreshFor), Coverage: operatorstate.CoverageGuardianOperation, Provenance: operatorstate.Provenance{DefinitionID: v.DefinitionID, ExecutionID: state.ID, Profile: "guardian", Source: "live", Stages: stages, Reason: operatorstate.PublicationGuardian, ApplicationVersion: v.ApplicationVersion}, Overview: overview})
	if err != nil {
		return fmt.Errorf("guardian current state invalid: %w", err)
	}
	if err = v.Store.Publish(current); err != nil {
		return fmt.Errorf("guardian current state publication failed")
	}
	if complete {
		v.hasEngineering = true
	}
	return nil
}

type Sink struct {
	Publisher *Publisher
}

func (v Sink) Emit(state runtimeservice.State, event runtimeservice.Event, evidence runtimeservice.Evidence) error {
	if runtimeservice.ValidateState(state) != nil || runtimeservice.ValidateEvent(event) != nil || runtimeservice.ValidateEvidence(evidence) != nil {
		return fmt.Errorf("guardian service evidence invalid")
	}
	return v.Publisher.Publish(state, event)
}

func ReportExit(checkpoints *Store, current *operatorstate.Store, generation, result string, at time.Time, freshFor time.Duration) error {
	allowed := map[string]bool{"success": true, "protocol": true, "timeout": true, "exit-code": true, "signal": true, "core-dump": true, "watchdog": true, "resources": true, "oom-kill": true}
	if checkpoints == nil || current == nil || !token(generation) || !allowed[result] || at.IsZero() || freshFor <= 0 || freshFor > presentationmodel.MaxFreshFor {
		return ErrExitEvidence
	}
	checkpoint, err := checkpoints.Load()
	if err != nil {
		return fmt.Errorf("%w: %v", ErrExitCheckpoint, err)
	}
	if checkpoint.Generation != generation || !checkpoint.Active {
		return nil
	}
	stored, err := current.Load()
	if err != nil {
		return fmt.Errorf("%w: %v", ErrExitCurrent, err)
	}
	status := presentationmodel.GuardianDegraded
	if result == "success" {
		status = presentationmodel.GuardianStopped
	}
	overview, err := presentationmodel.TransitionGuardian(stored.Overview, status, at.UTC())
	if err != nil {
		return ErrExitState
	}
	stored.ObservedAt, stored.PublishedAt, stored.FreshUntil, stored.Overview = at.UTC(), at.UTC(), at.UTC().Add(freshFor), overview
	stored, err = operatorstate.Normalize(stored)
	if err != nil {
		return fmt.Errorf("%w: %v", ErrExitState, err)
	}
	if err = current.Publish(stored); err != nil {
		return fmt.Errorf("%w: %v", ErrExitCurrent, err)
	}
	checkpoint.Active = false
	return checkpoints.Save(checkpoint)
}
