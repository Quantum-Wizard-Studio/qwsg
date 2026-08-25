package scheduler

import (
	"context"
	"errors"
	"fmt"
	"os"
	"sort"
	"time"

	"quantumwizard.hu/qwsg/internal/command"
	"quantumwizard.hu/qwsg/internal/configuration"
)

type Clock interface{ Observe() ClockObservation }
type CommandResolver func(string, command.Selection) (command.Definition, error)
type PipelineExecutor interface {
	Execute(context.Context, command.Definition) (command.Execution, error)
}

type Cycle struct {
	Configuration  configuration.Effective
	Selection      command.Selection
	LockOwnerID    string
	Store          StateStore
	Locker         Locker
	Clock          Clock
	TimeZones      TimeZoneResolver
	ResolveCommand CommandResolver
	Pipeline       PipelineExecutor
}

type CycleResult struct {
	SchemaName    string            `json:"schema_name"`
	SchemaVersion string            `json:"schema_version"`
	ID            string            `json:"id"`
	Evaluation    Evaluation        `json:"evaluation"`
	Results       []ExecutionResult `json:"results"`
	Traces        []ExecutionTrace  `json:"traces"`
	Events        []Event           `json:"events"`
	FinalState    State             `json:"final_state"`
}

// Run performs exactly one explicitly requested local scheduling cycle. It
// owns no loop and never discovers runtime paths or services.
func (cycle Cycle) Run(ctx context.Context) (CycleResult, error) {
	if err := configuration.ValidateEffective(cycle.Configuration); err != nil {
		return CycleResult{}, err
	}
	if cycle.Store == nil || cycle.Locker == nil || cycle.Clock == nil || cycle.TimeZones == nil || cycle.ResolveCommand == nil || cycle.Pipeline == nil {
		return CycleResult{}, fmt.Errorf("scheduler cycle dependencies are incomplete")
	}
	lock, err := cycle.Locker.Acquire(cycle.LockOwnerID)
	if err != nil {
		return CycleResult{}, err
	}
	released := false
	defer func() {
		if !released {
			_ = lock.Release()
		}
	}()
	state, err := cycle.Store.Load()
	if errors.Is(err, os.ErrNotExist) {
		state = NewState(cycle.Configuration.ID)
	} else if err != nil {
		return CycleResult{}, err
	} else if state.ConfigurationID != cycle.Configuration.ID {
		state = NewState(cycle.Configuration.ID)
	}
	observation := cycle.Clock.Observe()
	evaluation, err := Evaluate(cycle.Configuration, state, observation, cycle.TimeZones)
	if err != nil {
		return CycleResult{}, err
	}
	if err := cycle.Store.Save(evaluation.NextState); err != nil {
		return CycleResult{}, err
	}
	completions := make([]Completion, 0, len(evaluation.Requests))
	traces := make([]ExecutionTrace, 0, len(evaluation.Requests))
	for _, request := range evaluation.Requests {
		started := cycle.Clock.Observe().WallTime
		definition, resolveErr := cycle.ResolveCommand(request.CommandProfile, cycle.Selection)
		execution := command.Execution{}
		failureCode := ""
		if resolveErr != nil {
			failureCode = "command_resolution_failed"
		} else {
			timeout := time.Duration(request.ExecutionTimeoutNS)
			runContext, cancel := context.WithTimeout(ctx, timeout)
			execution, err = cycle.Pipeline.Execute(runContext, definition)
			cancel()
			if err != nil {
				failureCode = "pipeline_execution_failed"
			}
		}
		completed := cycle.Clock.Observe().WallTime
		completions = append(completions, Completion{Request: request, StartedAt: started, CompletedAt: completed, Execution: execution, FailureCode: failureCode})
		trace := ExecutionTrace{SchemaName: TraceSchema, SchemaVersion: SchemaVersion, RequestID: request.ID, Definition: definition, Execution: execution, FailureCode: failureCode}
		trace.ID = traceID(trace)
		traces = append(traces, trace)
	}
	finalState, results, events, err := ApplyCompletions(cycle.Configuration, evaluation.NextState, completions)
	if err != nil {
		return CycleResult{}, err
	}
	if err := cycle.Store.Save(finalState); err != nil {
		return CycleResult{}, err
	}
	if err := lock.Release(); err != nil {
		return CycleResult{}, err
	}
	released = true
	allEvents := append(append([]Event(nil), evaluation.Events...), events...)
	if results == nil {
		results = []ExecutionResult{}
	}
	if traces == nil {
		traces = []ExecutionTrace{}
	}
	if allEvents == nil {
		allEvents = []Event{}
	}
	sort.Slice(allEvents, func(i, j int) bool {
		if allEvents[i].At.Equal(allEvents[j].At) {
			return allEvents[i].ID < allEvents[j].ID
		}
		return allEvents[i].At.Before(allEvents[j].At)
	})
	result := CycleResult{SchemaName: CycleResultSchema, SchemaVersion: SchemaVersion, Evaluation: evaluation, Results: results, Traces: traces, Events: allEvents, FinalState: finalState}
	result.ID = cycleResultID(result)
	if err := ValidateCycleResult(result); err != nil {
		return CycleResult{}, err
	}
	return result, nil
}

// ValidateCycleResult validates the Scheduler-owned one-cycle adapter result.
func ValidateCycleResult(value CycleResult) error {
	if value.SchemaName != CycleResultSchema || value.SchemaVersion != SchemaVersion || value.Results == nil || value.Traces == nil || value.Events == nil || len(value.Traces) > MaxOccurrences {
		return fmt.Errorf("invalid scheduler cycle result envelope")
	}
	if err := ValidateEvaluation(value.Evaluation); err != nil {
		return err
	}
	if err := ValidateState(value.FinalState); err != nil {
		return err
	}
	if len(value.Results) != len(value.Evaluation.Requests) || len(value.Traces) != len(value.Evaluation.Requests) {
		return fmt.Errorf("scheduler cycle cardinality mismatch")
	}
	for i := range value.Evaluation.Requests {
		request, result, trace := value.Evaluation.Requests[i], value.Results[i], value.Traces[i]
		if err := validateResult(result); err != nil {
			return err
		}
		if err := validateTrace(trace); err != nil {
			return err
		}
		if request.ID != result.RequestID || request.ID != trace.RequestID {
			return fmt.Errorf("scheduler cycle request correlation mismatch")
		}
		if trace.FailureCode == "" && result.CommandExecutionID != trace.Execution.ID {
			return fmt.Errorf("scheduler cycle execution correlation mismatch")
		}
	}
	for _, event := range value.Events {
		if event.ID != eventID(event) {
			return fmt.Errorf("invalid scheduler cycle event")
		}
	}
	if value.ID != cycleResultID(value) {
		return fmt.Errorf("invalid scheduler cycle result identity")
	}
	return nil
}

func cycleResultID(value CycleResult) string {
	value.ID = ""
	return stableID("scheduler-cycle", value)
}

func ResolveCanonicalCommand(profile string, selection command.Selection) (command.Definition, error) {
	return command.ResolveProfile(profile, selection)
}
