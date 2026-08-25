package scheduler

import (
	"context"
	"encoding/json"
	"errors"
	"os"
	"path/filepath"
	"reflect"
	"sort"
	"strings"
	"testing"
	"time"

	"quantumwizard.hu/qwsg/internal/command"
	"quantumwizard.hu/qwsg/internal/configuration"
	"quantumwizard.hu/qwsg/internal/pipeline"
	"quantumwizard.hu/qwsg/internal/policy"
)

func TestIntervalAnchorMissedRunAndDeterminism(t *testing.T) {
	config := effective(t, 1, intervalSchedule("schedule.hourly", 10, configuration.MisfireRunOnce))
	start := time.Date(2026, 8, 6, 10, 0, 0, 0, time.UTC)
	first, err := Evaluate(config, NewState(config.ID), ClockObservation{WallTime: start, SessionID: "session.a", MonotonicNS: 0}, SystemTimeZones{})
	if err != nil {
		t.Fatal(err)
	}
	if len(first.Requests) != 0 || !first.NextState.Schedules[0].AnchorAt.Equal(start) {
		t.Fatalf("first evaluation executed or failed to anchor: %#v", first)
	}
	observation := ClockObservation{WallTime: start.Add(35 * time.Minute), SessionID: "session.a", MonotonicNS: int64(35 * time.Minute)}
	second, err := Evaluate(config, first.NextState, observation, SystemTimeZones{})
	if err != nil {
		t.Fatal(err)
	}
	again, err := Evaluate(config, first.NextState, observation, SystemTimeZones{})
	if err != nil {
		t.Fatal(err)
	}
	one, _ := MarshalEvaluationCanonical(second)
	two, _ := MarshalEvaluationCanonical(again)
	if !reflect.DeepEqual(one, two) || len(second.Requests) != 1 || !second.Requests[0].ScheduledAt.Equal(start.Add(30*time.Minute)) {
		t.Fatalf("interval evaluation is not deterministic or coalesced: %#v", second)
	}
}

func TestDisabledInapplicableAndClockDiscontinuity(t *testing.T) {
	disabled := intervalSchedule("schedule.disabled", 1, configuration.MisfireSkip)
	disabled.Enabled = false
	scoped := intervalSchedule("schedule.scoped", 1, configuration.MisfireSkip)
	scoped.CheckIDs = []string{"check.health"}
	config := effective(t, 2, disabled, scoped)
	start := time.Date(2026, 8, 6, 10, 0, 0, 0, time.UTC)
	first, _ := Evaluate(config, NewState(config.ID), ClockObservation{WallTime: start, SessionID: "session.a"}, SystemTimeZones{})
	if first.Records[0].Decision != DecisionDisabled || first.Records[1].Decision != DecisionInapplicable {
		t.Fatalf("unexpected decisions: %#v", first.Records)
	}
	config = effective(t, 1, intervalSchedule("schedule.clock", 1, configuration.MisfireRunOnce))
	first, _ = Evaluate(config, NewState(config.ID), ClockObservation{WallTime: start, SessionID: "session.a"}, SystemTimeZones{})
	second, err := Evaluate(config, first.NextState, ClockObservation{WallTime: start.Add(20 * time.Minute), SessionID: "session.a", MonotonicNS: int64(time.Minute)}, SystemTimeZones{})
	if err != nil {
		t.Fatal(err)
	}
	if len(second.Requests) != 0 || second.Records[0].Decision != DecisionIndeterminate || second.Events[len(second.Events)-1].Kind != EventClockDiscontinuity {
		t.Fatalf("clock discontinuity fabricated work: %#v", second)
	}
}

func TestPriorityAndConcurrencyAreDeterministic(t *testing.T) {
	low := intervalSchedule("schedule.low", 1, configuration.MisfireRunOnce)
	low.Priority = 1
	high := intervalSchedule("schedule.high", 1, configuration.MisfireRunOnce)
	high.Priority = 100
	config := effective(t, 1, low, high)
	start := time.Date(2026, 8, 6, 10, 0, 0, 0, time.UTC)
	first, _ := Evaluate(config, NewState(config.ID), ClockObservation{WallTime: start, SessionID: "session.a"}, SystemTimeZones{})
	second, err := Evaluate(config, first.NextState, ClockObservation{WallTime: start.Add(time.Minute), SessionID: "session.a", MonotonicNS: int64(time.Minute)}, SystemTimeZones{})
	if err != nil {
		t.Fatal(err)
	}
	if len(second.Requests) != 1 || second.Requests[0].ScheduleID != "schedule.high" {
		t.Fatalf("priority failed: %#v", second.Requests)
	}
	if findRecord(second.Records, "schedule.low").Decision != DecisionDelayed {
		t.Fatalf("capacity was not explicit: %#v", second.Records)
	}
}

func TestDSTFirstAndSecondOccurrence(t *testing.T) {
	if _, err := time.LoadLocation("Europe/Budapest"); err != nil {
		t.Skip("zone database unavailable")
	}
	base := calendarSchedule("schedule.dst", configuration.DSTFirstOccurrence)
	start := time.Date(2026, 10, 24, 23, 0, 0, 0, time.UTC)
	end := time.Date(2026, 10, 25, 3, 0, 0, 0, time.UTC)
	firstConfig := effective(t, 1, base)
	first, _ := Evaluate(firstConfig, NewState(firstConfig.ID), ClockObservation{WallTime: start, SessionID: "session.a"}, SystemTimeZones{})
	evaluated, err := Evaluate(firstConfig, first.NextState, ClockObservation{WallTime: end, SessionID: "session.a", MonotonicNS: int64(end.Sub(start))}, SystemTimeZones{})
	if err != nil {
		t.Fatal(err)
	}
	if len(evaluated.Requests) != 1 || !evaluated.Requests[0].ScheduledAt.Equal(time.Date(2026, 10, 25, 0, 30, 0, 0, time.UTC)) {
		t.Fatalf("first DST occurrence wrong: %#v", evaluated.Requests)
	}
	base.DSTPolicy = configuration.DSTSecondOccurrence
	secondConfig := effective(t, 1, base)
	initial, _ := Evaluate(secondConfig, NewState(secondConfig.ID), ClockObservation{WallTime: start, SessionID: "session.a"}, SystemTimeZones{})
	evaluated, err = Evaluate(secondConfig, initial.NextState, ClockObservation{WallTime: end, SessionID: "session.a", MonotonicNS: int64(end.Sub(start))}, SystemTimeZones{})
	if err != nil {
		t.Fatal(err)
	}
	if len(evaluated.Requests) != 1 || !evaluated.Requests[0].ScheduledAt.Equal(time.Date(2026, 10, 25, 1, 30, 0, 0, time.UTC)) {
		t.Fatalf("second DST occurrence wrong: %#v", evaluated.Requests)
	}
}

func TestRestartMarksActiveInterruptedAndPlansRetry(t *testing.T) {
	schedule := intervalSchedule("schedule.restart", 1, configuration.MisfireRunOnce)
	schedule.RetryPolicyID = "retry.restart"
	config := effective(t, 1, schedule)
	start := time.Date(2026, 8, 6, 10, 0, 0, 0, time.UTC)
	first, err := Evaluate(config, NewState(config.ID), ClockObservation{WallTime: start, SessionID: "session.a"}, SystemTimeZones{})
	if err != nil {
		t.Fatal(err)
	}
	due, err := Evaluate(config, first.NextState, ClockObservation{WallTime: start.Add(time.Minute), SessionID: "session.a", MonotonicNS: int64(time.Minute)}, SystemTimeZones{})
	if err != nil {
		t.Fatal(err)
	}
	restarted, err := Evaluate(config, due.NextState, ClockObservation{WallTime: start.Add(2 * time.Minute), SessionID: "session.b"}, SystemTimeZones{})
	if err != nil {
		t.Fatal(err)
	}
	if len(restarted.NextState.Results) == 0 || restarted.NextState.Results[0].Outcome != ExecutionInterrupted {
		t.Fatalf("restart did not preserve interruption: %#v", restarted.NextState)
	}
	found := false
	for _, request := range restarted.Requests {
		if request.Attempt == 2 {
			found = true
		}
	}
	if !found {
		t.Fatalf("retry was not planned: %#v", restarted.Requests)
	}
}

func TestMisfireOverlapPendingAndUnavailableTimeZone(t *testing.T) {
	start := time.Date(2026, 8, 6, 10, 0, 0, 0, time.UTC)
	for _, test := range []struct {
		policy configuration.MisfirePolicy
		want   Decision
	}{{configuration.MisfireSkip, DecisionSkipped}, {configuration.MisfireIndeterminate, DecisionIndeterminate}} {
		config := effective(t, 1, intervalSchedule("schedule.misfire", 1, test.policy))
		first, _ := Evaluate(config, NewState(config.ID), ClockObservation{WallTime: start, SessionID: "session.a"}, SystemTimeZones{})
		result, err := Evaluate(config, first.NextState, ClockObservation{WallTime: start.Add(3 * time.Minute), SessionID: "session.a", MonotonicNS: int64(3 * time.Minute)}, SystemTimeZones{})
		if err != nil || len(result.Requests) != 0 || result.Records[0].Decision != test.want {
			t.Fatalf("misfire %q: result=%#v err=%v", test.policy, result, err)
		}
	}

	calendar := calendarSchedule("schedule.zone", configuration.DSTFirstOccurrence)
	calendar.TimeZone = "Unavailable/Zone"
	config := effective(t, 1, calendar)
	first, _ := Evaluate(config, NewState(config.ID), ClockObservation{WallTime: start, SessionID: "session.a"}, rejectingZones{})
	result, err := Evaluate(config, first.NextState, ClockObservation{WallTime: start.Add(time.Hour), SessionID: "session.a", MonotonicNS: int64(time.Hour)}, rejectingZones{})
	if err != nil || result.Records[0].Decision != DecisionIndeterminate || len(result.Requests) != 0 {
		t.Fatalf("unavailable zone was not isolated: %#v err=%v", result, err)
	}

	config = effective(t, 1, intervalSchedule("schedule.overlap", 1, configuration.MisfireRunOnce))
	first, _ = Evaluate(config, NewState(config.ID), ClockObservation{WallTime: start, SessionID: "session.a"}, SystemTimeZones{})
	due, _ := Evaluate(config, first.NextState, ClockObservation{WallTime: start.Add(time.Minute), SessionID: "session.a", MonotonicNS: int64(time.Minute)}, SystemTimeZones{})
	queued, err := Evaluate(config, due.NextState, ClockObservation{WallTime: start.Add(2 * time.Minute), SessionID: "session.a", MonotonicNS: int64(2 * time.Minute)}, SystemTimeZones{})
	if err != nil || queued.Records[0].Decision != DecisionQueued || queued.NextState.Schedules[0].Pending == nil || len(queued.Requests) != 0 {
		t.Fatalf("overlap was not bounded and queued: %#v err=%v", queued, err)
	}
}

type rejectingZones struct{}

func (rejectingZones) Resolve(string) (*time.Location, error) {
	return nil, errors.New("zone unavailable")
}

func TestFileStoreIntegrityPermissionsAndLockContention(t *testing.T) {
	directory := filepath.Join(t.TempDir(), "scheduler")
	store, err := OpenFileStore(directory)
	if err != nil {
		t.Fatal(err)
	}
	state := NewState("config:fixture")
	if err := store.Save(state); err != nil {
		t.Fatal(err)
	}
	loaded, err := store.Load()
	if err != nil || loaded.ID != state.ID {
		t.Fatalf("load failed: %v %#v", err, loaded)
	}
	info, _ := os.Stat(filepath.Join(directory, stateFileName))
	if info.Mode().Perm() != 0600 {
		t.Fatalf("state mode=%o", info.Mode().Perm())
	}
	data, _ := os.ReadFile(filepath.Join(directory, stateFileName))
	data[len(data)/2] ^= 1
	_ = os.WriteFile(filepath.Join(directory, stateFileName), data, 0600)
	if _, err := store.Load(); err == nil {
		t.Fatal("corrupt state accepted")
	}
	locker, _ := NewFileLocker(directory)
	if _, err := locker.Acquire("invalid owner"); err == nil {
		t.Fatal("invalid lock owner accepted")
	}
	first, err := locker.Acquire("owner.first")
	if err != nil {
		t.Fatal(err)
	}
	if _, err := locker.Acquire("owner.second"); !errors.Is(err, ErrLockContended) {
		t.Fatalf("contention=%v", err)
	}
	if err := first.Release(); err != nil {
		t.Fatal(err)
	}
	second, err := locker.Acquire("owner.second")
	if err != nil {
		t.Fatal(err)
	}
	_ = second.Release()
}

func TestCycleExecutesCanonicalProfileAndCapturesPolicy(t *testing.T) {
	config := effective(t, 1, intervalSchedule("schedule.cycle", 1, configuration.MisfireRunOnce))
	directory := filepath.Join(t.TempDir(), "cycle")
	store, _ := OpenFileStore(directory)
	locker, _ := NewFileLocker(directory)
	start := time.Date(2026, 8, 6, 10, 0, 0, 0, time.UTC)
	clock := &sequenceClock{values: []ClockObservation{
		{WallTime: start, SessionID: "session.a"},
		{WallTime: start.Add(time.Minute), SessionID: "session.a", MonotonicNS: int64(time.Minute)},
		{WallTime: start.Add(time.Minute), SessionID: "session.a", MonotonicNS: int64(time.Minute)},
		{WallTime: start.Add(time.Minute), SessionID: "session.a", MonotonicNS: int64(time.Minute)},
	}}
	executor := &fakePipeline{}
	cycle := Cycle{Configuration: config, Selection: command.Selection{Source: "live"}, LockOwnerID: "cycle.owner", Store: store, Locker: locker, Clock: clock, TimeZones: SystemTimeZones{}, ResolveCommand: ResolveCanonicalCommand, Pipeline: executor}
	if _, err := cycle.Run(context.Background()); err != nil {
		t.Fatal(err)
	}
	result, err := cycle.Run(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	if executor.calls != 1 || len(result.Results) != 1 || result.Results[0].CommandExecutionID != "execution.fixture" || !reflect.DeepEqual(result.Results[0].PolicyOutcomes, []string{"observe"}) {
		t.Fatalf("cycle result=%#v calls=%d", result, executor.calls)
	}
	if len(result.Traces) != 1 || result.Traces[0].RequestID != result.Results[0].RequestID || ValidateCycleResult(result) != nil {
		t.Fatalf("invalid execution trace: %#v", result.Traces)
	}
	tampered := result
	tampered.Traces = append([]ExecutionTrace(nil), result.Traces...)
	tampered.Traces[0].FailureCode = "private-error"
	if ValidateCycleResult(tampered) == nil {
		t.Fatal("tampered cycle result accepted")
	}
	if _, err := store.Load(); err != nil {
		t.Fatal(err)
	}
}

func TestCycleReinitializesPersistedStateForChangedConfiguration(t *testing.T) {
	oldConfig := effective(t, 1, intervalSchedule("schedule.old", 1, configuration.MisfireRunOnce))
	newConfig := effective(t, 1, intervalSchedule("schedule.new", 1, configuration.MisfireRunOnce))
	directory := filepath.Join(t.TempDir(), "cycle")
	store, _ := OpenFileStore(directory)
	locker, _ := NewFileLocker(directory)
	if err := store.Save(NewState(oldConfig.ID)); err != nil {
		t.Fatal(err)
	}
	start := time.Date(2026, 8, 25, 12, 0, 0, 0, time.UTC)
	executor := &fakePipeline{}
	cycle := Cycle{
		Configuration: newConfig,
		Selection:     command.Selection{Source: "live"},
		LockOwnerID:   "cycle.recovered",
		Store:         store,
		Locker:        locker,
		Clock: &sequenceClock{values: []ClockObservation{
			{WallTime: start, SessionID: "generation.recovered"},
			{WallTime: start.Add(time.Minute), SessionID: "generation.recovered", MonotonicNS: int64(time.Minute)},
			{WallTime: start.Add(time.Minute), SessionID: "generation.recovered", MonotonicNS: int64(time.Minute)},
			{WallTime: start.Add(time.Minute), SessionID: "generation.recovered", MonotonicNS: int64(time.Minute)},
		}},
		TimeZones:      SystemTimeZones{},
		ResolveCommand: ResolveCanonicalCommand,
		Pipeline:       executor,
	}
	result, err := cycle.Run(context.Background())
	if err != nil {
		t.Fatalf("recovered cycle did not converge after configuration change: %v", err)
	}
	if result.FinalState.ConfigurationID != newConfig.ID || result.FinalState.SessionID != "generation.recovered" {
		t.Fatalf("recovered state retained superseded ownership: %#v", result.FinalState)
	}
	loaded, err := store.Load()
	if err != nil || loaded.ID != result.FinalState.ID {
		t.Fatalf("recovered scheduler state was not published atomically: %#v %v", loaded, err)
	}
	result, err = cycle.Run(context.Background())
	if err != nil || executor.calls != 1 || len(result.Traces) != 1 || !result.Traces[0].Execution.Complete {
		t.Fatalf("recovered generation did not complete fresh canonical work: %#v calls=%d err=%v", result, executor.calls, err)
	}
}

func TestStrictStateValidationAndCanonicalSerialization(t *testing.T) {
	state := NewState("config:fixture")
	data, err := MarshalStateCanonical(state)
	if err != nil {
		t.Fatal(err)
	}
	var decoded State
	if err := json.Unmarshal(data, &decoded); err != nil || ValidateState(decoded) != nil {
		t.Fatalf("canonical state failed: %v", err)
	}
	decoded.ConfigurationID = "tampered"
	if ValidateState(decoded) == nil {
		t.Fatal("tampered state accepted")
	}
}

var _ PipelineExecutor = pipeline.Orchestrator{}

type sequenceClock struct {
	values []ClockObservation
	index  int
}

func (clock *sequenceClock) Observe() ClockObservation {
	if clock.index >= len(clock.values) {
		return clock.values[len(clock.values)-1]
	}
	value := clock.values[clock.index]
	clock.index++
	return value
}

type fakePipeline struct{ calls int }

func (pipeline *fakePipeline) Execute(context.Context, command.Definition) (command.Execution, error) {
	pipeline.calls++
	governed := policy.Result{Records: []policy.EvaluationRecord{{ID: "policy.record", Outcome: policy.Observe}}}
	return command.Execution{ID: "execution.fixture", Complete: true, Stages: []command.StageResult{{Stage: command.Policy, ContractName: policy.SchemaName, Version: policy.SchemaVersion, Complete: true, Value: governed}}}, nil
}

func effective(t *testing.T, concurrency int, schedules ...configuration.Schedule) configuration.Effective {
	t.Helper()
	rules := pipeline.CanonicalObservationRules()
	policies := pipeline.CanonicalPolicyProfiles()
	base, err := configuration.BuiltIn(rules, policies)
	if err != nil {
		t.Fatal(err)
	}
	sort.Slice(schedules, func(i, j int) bool { return schedules[i].ID < schedules[j].ID })
	checks := []configuration.Check{}
	retries := []configuration.RetryPolicy{{ID: "canonical.default", MaxAttempts: 1}}
	seen := map[string]bool{}
	for _, schedule := range schedules {
		if schedule.RetryPolicyID != "canonical.default" {
			retries = append(retries, configuration.RetryPolicy{ID: schedule.RetryPolicyID, MaxAttempts: 3, InitialDelayNS: 0, MaxDelayNS: 0})
		}
		for _, id := range schedule.CheckIDs {
			if !seen[id] {
				checks = append(checks, configuration.Check{ID: id, Enabled: true, TargetIDs: []string{}})
				seen[id] = true
			}
		}
	}
	sort.Slice(checks, func(i, j int) bool { return checks[i].ID < checks[j].ID })
	sort.Slice(retries, func(i, j int) bool { return retries[i].ID < retries[j].ID })
	source, err := configuration.NormalizeSource(configuration.Source{SchemaName: configuration.SourceSchema, SchemaVersion: configuration.SchemaVersion, ModelVersion: configuration.ModelVersion, ID: "scheduler.test", SourceVersion: "1.0", Kind: configuration.PrimaryLocal, Patch: configuration.Patch{Schedules: &schedules, Concurrency: &concurrency, Checks: &checks, RetryPolicies: &retries}})
	if err != nil {
		t.Fatal(err)
	}
	result, err := configuration.Resolve([]configuration.Source{base, source})
	if err != nil {
		t.Fatal(err)
	}
	return result
}
func intervalSchedule(id string, minutes int, misfire configuration.MisfirePolicy) configuration.Schedule {
	return configuration.Schedule{ID: id, ContractVersion: configuration.ScheduleVersion, Enabled: true, TimeZone: "UTC", Trigger: configuration.IntervalTrigger, IntervalNS: int64(time.Duration(minutes) * time.Minute), Calendar: configuration.Calendar{Minutes: []int{}, Hours: []int{}, MonthDays: []int{}, Months: []int{}, Weekdays: []int{}}, DSTPolicy: configuration.DSTSkipNonexistent, Priority: 1, MisfirePolicy: misfire, OverlapPolicy: configuration.OverlapForbid, ExecutionTimeoutNS: int64(time.Minute), RetryPolicyID: "canonical.default", CheckIDs: []string{}, CommandProfile: "status"}
}
func calendarSchedule(id string, dst configuration.DSTPolicy) configuration.Schedule {
	return configuration.Schedule{ID: id, ContractVersion: configuration.ScheduleVersion, Enabled: true, TimeZone: "Europe/Budapest", Trigger: configuration.CalendarTrigger, Calendar: configuration.Calendar{Minutes: []int{30}, Hours: []int{2}, MonthDays: []int{}, Months: []int{}, Weekdays: []int{}}, DSTPolicy: dst, Priority: 1, MisfirePolicy: configuration.MisfireRunOnce, OverlapPolicy: configuration.OverlapForbid, ExecutionTimeoutNS: int64(time.Minute), RetryPolicyID: "canonical.default", CheckIDs: []string{}, CommandProfile: "status"}
}

func TestNoSensitiveFailurePayload(t *testing.T) {
	config := effective(t, 1, intervalSchedule("schedule.safe", 1, configuration.MisfireRunOnce))
	text := string(mustJSON(config))
	if strings.Contains(text, "password=seeded-secret") {
		t.Fatal("secret leaked")
	}
}
func mustJSON(value any) []byte { data, _ := json.Marshal(value); return data }
