package runtimeservice

import (
	"context"
	"errors"
	"os"
	"reflect"
	"syscall"
	"testing"
	"time"

	"quantumwizard.hu/qwsg/internal/alert"
	"quantumwizard.hu/qwsg/internal/command"
	"quantumwizard.hu/qwsg/internal/configuration"
	"quantumwizard.hu/qwsg/internal/notification"
	"quantumwizard.hu/qwsg/internal/pipeline"
	"quantumwizard.hu/qwsg/internal/runtime"
	"quantumwizard.hu/qwsg/internal/scheduler"
)

func TestServiceRunsFixedRateAndHandsOffExactRuntimeState(t *testing.T) {
	input, scheduled := fixture(t)
	clock := &fakeClock{now: input.StartedAt}
	ctx, cancel := context.WithCancel(context.Background())
	runner := &validRunner{scheduled: scheduled, serviceClock: clock, duration: time.Second, cancelAfter: 2, cancel: cancel}
	sink := &memorySink{}
	result, err := (Service{Clock: clock, Waiter: fakeWaiter{clock}, Runner: runner, Sink: sink}).Run(ctx, input)
	if err != nil {
		t.Fatal(err)
	}
	if result.FinalState.Lifecycle != Stopped || runner.calls != 2 || result.FinalState.CyclesCompleted != 2 || result.FinalState.IntervalsMissed != 0 {
		t.Fatalf("result=%#v calls=%d", result, runner.calls)
	}
	if len(runner.inputs) != 2 || runner.inputs[1].PreviousState.LastCompletedCycleID != runner.inputs[0].Context.ID || runner.inputs[0].Configuration.ID != runner.inputs[1].Configuration.ID || !reflect.DeepEqual(runner.inputs[0].NotificationPolicy, runner.inputs[1].NotificationPolicy) {
		t.Fatal("runtime state handoff changed immutable seed or lost proposed state")
	}
	if runner.inputs[1].Context.StartedAt != input.StartedAt.Add(10*time.Second) {
		t.Fatalf("second nominal=%s", runner.inputs[1].Context.StartedAt)
	}
	if runner.inputs[0].Context.Deadline != input.StartedAt.Add(5*time.Second) {
		t.Fatalf("cycle deadline=%s", runner.inputs[0].Context.Deadline)
	}
	if err := ValidateResult(result); err != nil {
		t.Fatal(err)
	}
	data, err := MarshalCanonical(result)
	if err != nil {
		t.Fatal(err)
	}
	decoded, err := DecodeResult(data)
	if err != nil || !reflect.DeepEqual(result, decoded) {
		t.Fatalf("round trip=%v", err)
	}
	if _, err := DecodeResult(append(data, []byte("{}")...)); err == nil {
		t.Fatal("trailing data accepted")
	}
	assertEvidenceSequence(t, sink.events)
	if len(sink.events) != len(sink.evidence) {
		t.Fatal("event/evidence cardinality mismatch")
	}
	for i, evidence := range sink.evidence {
		if ValidateEvidence(evidence) != nil || evidence.EventID != sink.events[i].ID {
			t.Fatalf("evidence[%d]=%#v", i, evidence)
		}
	}
	definitionData, _ := MarshalDefinition(input.Definition)
	if decoded, err := DecodeDefinition(definitionData); err != nil || !reflect.DeepEqual(decoded, input.Definition) {
		t.Fatalf("definition round trip: %v", err)
	}
	stateData, _ := MarshalState(input.InitialState)
	if decoded, err := DecodeState(stateData); err != nil || !reflect.DeepEqual(decoded, input.InitialState) {
		t.Fatalf("state round trip: %v", err)
	}
	inputData, _ := MarshalInput(input)
	if decoded, err := DecodeInput(inputData); err != nil {
		t.Fatalf("input round trip: %v", err)
	} else if roundTrip, marshalErr := MarshalInput(decoded); marshalErr != nil || !reflect.DeepEqual(roundTrip, inputData) {
		t.Fatalf("input canonical round trip: %v", marshalErr)
	}
	eventData, _ := MarshalEvent(sink.events[0])
	if decoded, err := DecodeEvent(eventData); err != nil || !reflect.DeepEqual(decoded, sink.events[0]) {
		t.Fatalf("event round trip: %v", err)
	}
	evidenceData, _ := MarshalEvidence(sink.evidence[0])
	if decoded, err := DecodeEvidence(evidenceData); err != nil || !reflect.DeepEqual(decoded, sink.evidence[0]) {
		t.Fatalf("evidence round trip: %v", err)
	}
}

func TestServiceSkipsElapsedNominalBoundariesWithoutBurst(t *testing.T) {
	input, scheduled := fixture(t)
	clock := &fakeClock{now: input.StartedAt}
	ctx, cancel := context.WithCancel(context.Background())
	runner := &validRunner{scheduled: scheduled, serviceClock: clock, duration: 25 * time.Second, cancelAfter: 2, cancel: cancel}
	sink := &memorySink{}
	result, err := (Service{Clock: clock, Waiter: fakeWaiter{clock}, Runner: runner, Sink: sink}).Run(ctx, input)
	if err != nil {
		t.Fatal(err)
	}
	if runner.calls != 2 || runner.inputs[1].Context.StartedAt != input.StartedAt.Add(30*time.Second) || result.FinalState.IntervalsMissed != 2 {
		t.Fatalf("calls=%d second=%s missed=%d", runner.calls, runner.inputs[1].Context.StartedAt, result.FinalState.IntervalsMissed)
	}
	missedRecords := 0
	for _, value := range sink.events {
		if value.Kind == EvidenceIntervalsMissed {
			missedRecords++
			if value.MissedCount != 2 {
				t.Fatalf("missed=%d", value.MissedCount)
			}
		}
	}
	if missedRecords != 1 {
		t.Fatalf("missed records=%d", missedRecords)
	}
}

func TestServiceCancellationBeforeFirstCycleIsGraceful(t *testing.T) {
	input, scheduled := fixture(t)
	clock := &fakeClock{now: input.StartedAt}
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	runner := &validRunner{scheduled: scheduled, serviceClock: clock}
	result, err := (Service{Clock: clock, Waiter: fakeWaiter{clock}, Runner: runner, Sink: &memorySink{}}).Run(ctx, input)
	if err != nil || runner.calls != 0 || result.FinalState.Lifecycle != Stopped || result.TerminalReason != "context_cancelled" {
		t.Fatalf("result=%#v calls=%d err=%v", result, runner.calls, err)
	}
}

func TestServicePropagatesCancellationDuringActiveCycle(t *testing.T) {
	input, scheduled := fixture(t)
	clock := &fakeClock{now: input.StartedAt}
	ctx, cancel := context.WithCancel(context.Background())
	inner := &validRunner{scheduled: scheduled, serviceClock: clock}
	runner := cancellingRunner{cancel: cancel, inner: inner}
	result, err := (Service{Clock: clock, Waiter: fakeWaiter{clock}, Runner: runner, Sink: &memorySink{}}).Run(ctx, input)
	if err != nil || result.FinalState.Lifecycle != Stopped || result.LastRuntimeOutcome != runtime.Cancelled || inner.calls != 1 {
		t.Fatalf("result=%#v calls=%d err=%v", result, inner.calls, err)
	}
}

func TestServiceFailsClosedWithoutLeakingDependencyErrors(t *testing.T) {
	input, _ := fixture(t)
	clock := &fakeClock{now: input.StartedAt}
	runner := errorRunner{err: errors.New("secret destination /private/path")}
	result, err := (Service{Clock: clock, Waiter: fakeWaiter{clock}, Runner: runner, Sink: &memorySink{}}).Run(context.Background(), input)
	if err != nil || result.FinalState.Lifecycle != Failed || result.TerminalReason != "runtime_runner_failed" {
		t.Fatalf("result=%#v err=%v", result, err)
	}
	data, _ := MarshalCanonical(result)
	if bytesContain(data, []byte("secret")) || bytesContain(data, []byte("private")) {
		t.Fatal("raw dependency error leaked")
	}
}

func TestServiceFailsClosedOnEvidenceRefusal(t *testing.T) {
	input, scheduled := fixture(t)
	clock := &fakeClock{now: input.StartedAt}
	result, err := (Service{Clock: clock, Waiter: fakeWaiter{clock}, Runner: &validRunner{scheduled: scheduled, serviceClock: clock}, Sink: &memorySink{failAt: 1}}).Run(context.Background(), input)
	if err != nil || result.FinalState.Lifecycle != Failed || result.TerminalReason != "evidence_sink_failed" {
		t.Fatalf("result=%#v err=%v", result, err)
	}
}

func TestSignalAdapterRegistersOnlyInterruptAndTerminate(t *testing.T) {
	input, scheduled := fixture(t)
	clock := &fakeClock{now: input.StartedAt}
	stopped := false
	factory := func(parent context.Context, values ...os.Signal) (context.Context, context.CancelFunc) {
		if !reflect.DeepEqual(values, []os.Signal{os.Interrupt, syscall.SIGTERM}) {
			t.Fatalf("signals=%v", values)
		}
		ctx, cancel := context.WithCancel(parent)
		cancel()
		return ctx, func() { stopped = true }
	}
	result, err := RunWithSignals(context.Background(), Service{Clock: clock, Waiter: fakeWaiter{clock}, Runner: &validRunner{scheduled: scheduled, serviceClock: clock}, Sink: &memorySink{}}, input, factory)
	if err != nil || !stopped || result.FinalState.Lifecycle != Stopped {
		t.Fatalf("result=%#v stopped=%t err=%v", result, stopped, err)
	}
	if _, err := RunWithSignals(context.Background(), Service{}, input, nil); !errors.Is(err, os.ErrInvalid) {
		t.Fatalf("nil factory=%v", err)
	}
}

func TestDefinitionStateAndInputRejectTamperingAndLimits(t *testing.T) {
	input, _ := fixture(t)
	definition := input.Definition
	definition.IntervalNS = 0
	if ValidateDefinition(definition) == nil {
		t.Fatal("zero interval accepted")
	}
	state := input.InitialState
	state.ID = "tampered"
	if ValidateState(state) == nil {
		t.Fatal("tampered state accepted")
	}
	input.InitialState.ServiceID = "different"
	if ValidateInput(input) == nil {
		t.Fatal("mismatched service state accepted")
	}
}

type fakeClock struct{ now time.Time }

func (v *fakeClock) Observe() time.Time { return v.now }

type fakeWaiter struct{ clock *fakeClock }

func (v fakeWaiter) Wait(ctx context.Context, at time.Time) error {
	if err := ctx.Err(); err != nil {
		return err
	}
	v.clock.now = at
	return nil
}

type memorySink struct {
	events   []Event
	evidence []Evidence
	failAt   int
}

func (v *memorySink) Emit(state State, event Event, evidence Evidence) error {
	if v.failAt > 0 && len(v.events)+1 == v.failAt {
		return errors.New("sink private error")
	}
	if ValidateState(state) != nil || state.ServiceID != event.ServiceID {
		return errors.New("invalid observed service state")
	}
	v.events = append(v.events, event)
	v.evidence = append(v.evidence, evidence)
	return nil
}

type errorRunner struct{ err error }

func (v errorRunner) Run(context.Context, runtime.Input) (runtime.Result, error) {
	return runtime.Result{}, v.err
}

type cancellingRunner struct {
	cancel context.CancelFunc
	inner  *validRunner
}

func (v cancellingRunner) Run(ctx context.Context, input runtime.Input) (runtime.Result, error) {
	v.cancel()
	return v.inner.Run(ctx, input)
}

type validRunner struct {
	scheduled    scheduler.CycleResult
	serviceClock *fakeClock
	duration     time.Duration
	cancelAfter  int
	cancel       context.CancelFunc
	calls        int
	inputs       []runtime.Input
}

func (v *validRunner) Run(ctx context.Context, input runtime.Input) (runtime.Result, error) {
	v.calls++
	v.inputs = append(v.inputs, input)
	clock := &runtimeClock{values: []time.Time{input.Context.StartedAt, input.Context.StartedAt.Add(time.Millisecond), input.Context.StartedAt.Add(2 * time.Millisecond), input.Context.StartedAt.Add(3 * time.Millisecond)}}
	registry, _ := notification.NewRegistry()
	result, err := (runtime.Coordinator{Scheduler: fixedScheduler{v.scheduled}, Clock: clock, Providers: registry}).Run(ctx, input)
	v.serviceClock.now = v.serviceClock.now.Add(v.duration)
	if v.cancelAfter > 0 && v.calls == v.cancelAfter {
		v.cancel()
	}
	return result, err
}

type fixedScheduler struct{ result scheduler.CycleResult }

func (v fixedScheduler) Run(context.Context) (scheduler.CycleResult, error) { return v.result, nil }

type runtimeClock struct {
	values []time.Time
	index  int
}

func (v *runtimeClock) Observe() time.Time {
	if v.index >= len(v.values) {
		return v.values[len(v.values)-1]
	}
	result := v.values[v.index]
	v.index++
	return result
}

type memoryStore struct {
	value  scheduler.State
	exists bool
}

func (v *memoryStore) Load() (scheduler.State, error) {
	if !v.exists {
		return scheduler.State{}, os.ErrNotExist
	}
	return v.value, nil
}
func (v *memoryStore) Save(value scheduler.State) error { v.value = value; v.exists = true; return nil }

type memoryLocker struct{}
type memoryLock struct{}

func (memoryLocker) Acquire(string) (scheduler.Lock, error) { return memoryLock{}, nil }
func (memoryLock) Release() error                           { return nil }

type schedulerClock struct{ value scheduler.ClockObservation }

func (v schedulerClock) Observe() scheduler.ClockObservation { return v.value }

type unusedPipeline struct{}

func (unusedPipeline) Execute(context.Context, command.Definition) (command.Execution, error) {
	return command.Execution{}, errors.New("unexpected pipeline call")
}

func fixture(t *testing.T) (Input, scheduler.CycleResult) {
	t.Helper()
	base, err := configuration.BuiltIn(pipeline.CanonicalObservationRules(), pipeline.CanonicalPolicyProfiles())
	if err != nil {
		t.Fatal(err)
	}
	config, err := configuration.Resolve([]configuration.Source{base})
	if err != nil {
		t.Fatal(err)
	}
	start := time.Date(2099, 1, 2, 3, 4, 5, 0, time.UTC)
	cycle := scheduler.Cycle{Configuration: config, Selection: command.Selection{Source: "live"}, LockOwnerID: "service.test", Store: &memoryStore{}, Locker: memoryLocker{}, Clock: schedulerClock{scheduler.ClockObservation{WallTime: start, SessionID: "service.test"}}, TimeZones: scheduler.SystemTimeZones{}, ResolveCommand: scheduler.ResolveCanonicalCommand, Pipeline: unusedPipeline{}}
	scheduled, err := cycle.Run(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	policyValue, err := notification.NewPolicy(notification.RetryPolicy{MaxAttempts: 1, DeliveryWindowNS: int64(time.Hour), BackoffNS: []int64{}}, []notification.Route{}, []notification.EndpointReference{}, []notification.ProviderBinding{})
	if err != nil {
		t.Fatal(err)
	}
	seedContext, err := runtime.NewExecutionContext("seed.fixture", "service.test", start, start.Add(time.Hour))
	if err != nil {
		t.Fatal(err)
	}
	seed := runtime.Input{Context: seedContext, Configuration: config, PreviousState: runtime.NewState(), PreviousAlertState: alert.NewState(config.ID), PreviousNotificationQueue: notification.NewQueueState(), AlertEvidenceTTLNS: int64(time.Hour), Acknowledgements: []alert.Acknowledgement{}, Suppressions: []alert.SuppressionWindow{}, NotificationPolicy: policyValue}
	definition, err := NewDefinition("service.test", 10*time.Second, 5*time.Second)
	if err != nil {
		t.Fatal(err)
	}
	return Input{Definition: definition, StartedAt: start, InitialState: NewState(definition.ServiceID), Seed: seed}, scheduled
}

func assertEvidenceSequence(t *testing.T, values []Event) {
	t.Helper()
	for i, v := range values {
		if v.Sequence != uint64(i+1) || ValidateEvent(v) != nil {
			t.Fatalf("evidence[%d]=%#v", i, v)
		}
	}
}
func bytesContain(data, fragment []byte) bool {
	if len(fragment) == 0 {
		return true
	}
	for i := 0; i+len(fragment) <= len(data); i++ {
		if reflect.DeepEqual(data[i:i+len(fragment)], fragment) {
			return true
		}
	}
	return false
}
