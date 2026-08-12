package runtime

import (
	"context"
	"errors"
	"fmt"
	"os"
	"reflect"
	"testing"
	"time"

	"quantumwizard.hu/qwsg/internal/alert"
	"quantumwizard.hu/qwsg/internal/command"
	"quantumwizard.hu/qwsg/internal/configuration"
	"quantumwizard.hu/qwsg/internal/inventory"
	"quantumwizard.hu/qwsg/internal/inventorystore"
	"quantumwizard.hu/qwsg/internal/notification"
	"quantumwizard.hu/qwsg/internal/pipeline"
	"quantumwizard.hu/qwsg/internal/report"
	"quantumwizard.hu/qwsg/internal/scheduler"
)

func TestCoordinatorRunsOneCycleAndReturnsIdle(t *testing.T) {
	input, scheduled, times := fixture(t)
	runner := &fakeScheduler{result: scheduled}
	coordinator := Coordinator{Scheduler: runner, Clock: &timeSequence{values: times}}
	result, err := coordinator.Run(context.Background(), input)
	if err != nil {
		t.Fatal(err)
	}
	if runner.calls != 1 || result.Outcome != Completed || result.NextState.Lifecycle != Idle {
		t.Fatalf("calls=%d outcome=%s state=%s", runner.calls, result.Outcome, result.NextState.Lifecycle)
	}
	if result.NotificationPlan == nil || result.NotificationCycle != nil || len(result.AlertResults) != 1 || result.SchedulerResultID != scheduled.ID {
		t.Fatalf("unexpected result: %#v", result)
	}
	if err := ValidateResult(result); err != nil {
		t.Fatal(err)
	}
	data, err := MarshalCanonical(result)
	if err != nil {
		t.Fatal(err)
	}
	if decoded, err := DecodeResult(data); err != nil || !reflect.DeepEqual(result, decoded) {
		t.Fatalf("canonical round trip: %v", err)
	}
	if _, err := DecodeResult(append(data, []byte("{}")...)); err == nil {
		t.Fatal("trailing data accepted")
	}

	second, _, times2 := fixture(t)
	result2, err := (Coordinator{Scheduler: &fakeScheduler{result: scheduled}, Clock: &timeSequence{values: times2}}).Run(context.Background(), second)
	if err != nil || !reflect.DeepEqual(result, result2) {
		t.Fatalf("non-deterministic result: %v", err)
	}
}

func TestCoordinatorHonorsCancellationBeforeScheduler(t *testing.T) {
	input, scheduled, times := fixture(t)
	runner := &fakeScheduler{result: scheduled}
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	result, err := (Coordinator{Scheduler: runner, Clock: &timeSequence{values: times}}).Run(ctx, input)
	if err != nil {
		t.Fatal(err)
	}
	if runner.calls != 0 || result.Outcome != Cancelled || result.NextState.Lifecycle != Idle {
		t.Fatalf("calls=%d result=%#v", runner.calls, result)
	}
}

func TestCoordinatorBoundsOperationalSchedulerFailure(t *testing.T) {
	input, _, times := fixture(t)
	result, err := (Coordinator{Scheduler: &fakeScheduler{err: errors.New("private detail")}, Clock: &timeSequence{values: times}}).Run(context.Background(), input)
	if err != nil {
		t.Fatal(err)
	}
	if result.Outcome != Failed || len(result.Components) != 1 || result.Components[0].FailureToken != "scheduler_cycle_failed" {
		t.Fatalf("unexpected failure result: %#v", result)
	}
	data, _ := MarshalCanonical(result)
	if contains(string(data), "private detail") {
		t.Fatal("raw component error leaked")
	}
}

func TestCoordinatorPreservesSchedulerTruthOnDeadline(t *testing.T) {
	input, scheduled, times := fixture(t)
	times[1] = input.Context.Deadline
	result, err := (Coordinator{Scheduler: &fakeScheduler{result: scheduled}, Clock: &timeSequence{values: times}}).Run(context.Background(), input)
	if err != nil {
		t.Fatal(err)
	}
	if result.Outcome != TimedOut || result.SchedulerResultID != scheduled.ID || result.FinalSchedulerState == nil || result.NextState.Lifecycle != Idle {
		t.Fatalf("deadline result lost partial truth: %#v", result)
	}
}

func TestRuntimeRejectsNonIdleAndTamperedResults(t *testing.T) {
	input, _, _ := fixture(t)
	input.PreviousState.Lifecycle = Running
	if _, err := (Coordinator{}).Run(context.Background(), input); err == nil {
		t.Fatal("non-idle state accepted")
	}
	state := NewState()
	state.ID = "tampered"
	if ValidateState(state) == nil {
		t.Fatal("tampered state accepted")
	}
}

func TestComponentFailureTokensAreClosedAndPrivacySafe(t *testing.T) {
	valid := []ComponentResult{
		{Component: SchedulerComponent, Status: ComponentFailed, FailureToken: "scheduler_cycle_failed"},
		{Component: AlertComponent, Status: ComponentFailed, FailureToken: "alert_evaluation_failed"},
		{Component: AlertComponent, Status: ComponentFailed, FailureToken: string(Cancelled)},
		{Component: AlertComponent, Status: ComponentFailed, FailureToken: string(TimedOut)},
		{Component: NotificationPlanComponent, Status: ComponentFailed, FailureToken: "notification_planning_failed"},
		{Component: NotificationDeliveryComponent, Status: ComponentFailed, FailureToken: "notification_delivery_failed"},
		{Component: NotificationDeliveryComponent, Status: ComponentSkipped, FailureToken: "no_delivery_requests"},
	}
	for _, item := range valid {
		if !validComponentResult(item) {
			t.Fatalf("canonical component token rejected: %+v", item)
		}
	}
	for _, item := range []ComponentResult{
		{Component: AlertComponent, Status: ComponentFailed, FailureToken: "private/path"},
		{Component: NotificationPlanComponent, Status: ComponentFailed, FailureToken: "alert_evaluation_failed"},
		{Component: SchedulerComponent, Status: ComponentCompleted, FailureToken: "scheduler_cycle_failed"},
		{Component: AlertComponent, Status: ComponentSkipped, FailureToken: "no_delivery_requests"},
	} {
		if validComponentResult(item) {
			t.Fatalf("non-canonical component token accepted: %+v", item)
		}
	}
}

func TestLargePolicyReportReachesNotificationPlanning(t *testing.T) {
	start := time.Date(2099, 1, 2, 3, 4, 5, 0, time.UTC)
	base, err := configuration.BuiltIn(pipeline.CanonicalObservationRules(), pipeline.CanonicalPolicyProfiles())
	if err != nil {
		t.Fatal(err)
	}
	schedules := []configuration.Schedule{{ID: "runtime.large", ContractVersion: configuration.ScheduleVersion, Enabled: true, TimeZone: "UTC", Trigger: configuration.IntervalTrigger, IntervalNS: int64(time.Minute), Calendar: configuration.Calendar{Minutes: []int{}, Hours: []int{}, MonthDays: []int{}, Months: []int{}, Weekdays: []int{}}, DSTPolicy: configuration.DSTFirstOccurrence, Priority: 0, MisfirePolicy: configuration.MisfireRunOnce, OverlapPolicy: configuration.OverlapForbid, ExecutionTimeoutNS: int64(30 * time.Second), RetryPolicyID: "canonical.default", CheckIDs: []string{}, CommandProfile: "observe"}}
	base.Identity = ""
	base.Patch.Schedules = &schedules
	base, err = configuration.NormalizeSource(base)
	if err != nil {
		t.Fatal(err)
	}
	config, err := configuration.Resolve([]configuration.Source{base})
	if err != nil {
		t.Fatal(err)
	}
	storeRoot := t.TempDir()
	if err = os.Chmod(storeRoot, 0o700); err != nil {
		t.Fatal(err)
	}
	store, err := inventorystore.Open(storeRoot, configuration.DefaultRetention)
	if err != nil {
		t.Fatal(err)
	}
	if _, err = store.Save(runtimeLargeSnapshot(start, 0, 367)); err != nil {
		t.Fatal(err)
	}
	current := runtimeLargeSnapshot(start.Add(time.Minute), 1, 367)
	scheduleClock := &schedulerClock{value: scheduler.ClockObservation{WallTime: start, SessionID: "runtime.large"}}
	cycle := scheduler.Cycle{Configuration: config, Selection: command.Selection{Source: "live", Store: storeRoot}, LockOwnerID: "runtime.large", Store: &memoryStore{}, Locker: memoryLocker{}, Clock: scheduleClock, TimeZones: scheduler.SystemTimeZones{}, ResolveCommand: scheduler.ResolveCanonicalCommand, Pipeline: pipeline.Orchestrator{Collect: func(context.Context) (inventory.Snapshot, error) { return current, nil }, Configuration: &config}}
	policyValue, err := notification.NewPolicy(notification.RetryPolicy{MaxAttempts: 1, DeliveryWindowNS: int64(time.Hour), BackoffNS: []int64{}}, []notification.Route{}, []notification.EndpointReference{}, []notification.ProviderBinding{})
	if err != nil {
		t.Fatal(err)
	}
	executionContext, err := NewExecutionContext("cycle.large", "owner.explicit", start.Add(time.Minute), start.Add(2*time.Minute))
	if err != nil {
		t.Fatal(err)
	}
	input := Input{Context: executionContext, Configuration: config, PreviousState: NewState(), PreviousAlertState: alert.NewState(config.ID), PreviousNotificationQueue: notification.NewQueueState(), AlertEvidenceTTLNS: int64(time.Hour), Acknowledgements: []alert.Acknowledgement{}, Suppressions: []alert.SuppressionWindow{}, NotificationPolicy: policyValue}
	if _, err = cycle.Run(context.Background()); err != nil {
		t.Fatal(err)
	}
	scheduleClock.value = scheduler.ClockObservation{WallTime: start.Add(time.Minute), SessionID: "runtime.large", MonotonicNS: int64(time.Minute)}
	scheduled, err := cycle.Run(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	if got := cycleResultPolicyReportSources(t, scheduled); got < 366 {
		t.Fatalf("Policy Report sources=%d, want at least 366", got)
	}
	times := []time.Time{start.Add(time.Minute), start.Add(time.Minute + time.Second), start.Add(time.Minute + 2*time.Second), start.Add(time.Minute + 3*time.Second)}
	result, err := (Coordinator{Scheduler: &fakeScheduler{result: scheduled}, Clock: &timeSequence{values: times}}).Run(context.Background(), input)
	if err != nil {
		t.Fatal(err)
	}
	if result.Outcome != Completed || result.NotificationPlan == nil || len(result.Components) != 4 || result.Components[1].Component != AlertComponent || result.Components[1].Status != ComponentCompleted || result.Components[2].Component != NotificationPlanComponent || result.Components[2].Status != ComponentCompleted {
		t.Fatalf("large report did not reach Notification planning: %+v", result.Components)
	}
}

func cycleResultPolicyReportSources(t *testing.T, result scheduler.CycleResult) int {
	t.Helper()
	if len(result.Traces) != 1 {
		t.Fatalf("Scheduler traces=%d", len(result.Traces))
	}
	stages := result.Traces[0].Execution.Stages
	if len(stages) != 8 {
		t.Fatalf("Pipeline stages=%d", len(stages))
	}
	reportValue, ok := stages[len(stages)-1].Value.(report.PolicyReport)
	if !ok || report.ValidatePolicyReport(reportValue) != nil {
		t.Fatal("Pipeline trace lacks a valid Policy Report")
	}
	return len(reportValue.Sources)
}

func runtimeLargeSnapshot(now time.Time, generation, count int) inventory.Snapshot {
	items := make([]inventory.Item, 0, count)
	for index := 0; index < count; index++ {
		items = append(items, inventory.Item{ID: fmt.Sprintf("component-%04d", index), Kind: "runtime_component", Facts: map[string]inventory.Fact{"value": {Value: index + generation, Quality: "observed", Sensitivity: "operational", Provenance: inventory.Provenance{SourceType: "fixture", SourceLabel: "bounded", ObservedAt: now}}}})
	}
	category := inventory.Category{CategoryID: "components", ContractVersion: "1.0", Status: inventory.Available, ObservedAt: now, CompletedAt: now, FreshUntil: now.Add(5 * time.Minute), CollectorID: "components", PrivilegeUsed: "ordinary-user", SourceSummary: []string{"bounded fixture"}, Items: items, Errors: []inventory.InventoryError{}, Redactions: []string{}}
	execution := inventory.CollectorExecution{CollectorName: "components", Version: "1", Capability: "components", SupportedPlatforms: []string{"linux"}, Timestamp: now, Status: inventory.Available, Warnings: []inventory.InventoryWarning{}, Errors: []inventory.InventoryError{}, Metadata: map[string]string{}}
	producer := inventory.Producer{ToolVersion: "test", ContractVersion: "1.0"}
	id := fmt.Sprintf("runtime-large-%d", generation)
	snapshot := inventory.Snapshot{SchemaVersion: inventory.SchemaVersion, SnapshotID: id, RequestID: id, InstanceID: "subject", ObservedAt: now, CompletedAt: now, FreshUntil: now.Add(5 * time.Minute), Status: inventory.Complete, Categories: []inventory.Category{category}, Errors: []inventory.InventoryError{}, Redactions: []string{}, Producer: producer}
	snapshot.Canonical = inventory.AssembleSystemInventory(snapshot.Categories, []inventory.CollectorExecution{execution}, snapshot.SnapshotID, snapshot.RequestID, snapshot.InstanceID, now, now, snapshot.FreshUntil, 0, producer)
	return snapshot
}

type fakeScheduler struct {
	result scheduler.CycleResult
	err    error
	calls  int
}

func (f *fakeScheduler) Run(context.Context) (scheduler.CycleResult, error) {
	f.calls++
	return f.result, f.err
}

type timeSequence struct {
	values []time.Time
	index  int
}

func (s *timeSequence) Observe() time.Time {
	if s.index >= len(s.values) {
		return s.values[len(s.values)-1]
	}
	v := s.values[s.index]
	s.index++
	return v
}

type memoryStore struct {
	value  scheduler.State
	exists bool
}

func (s *memoryStore) Load() (scheduler.State, error) {
	if !s.exists {
		return scheduler.State{}, os.ErrNotExist
	}
	return s.value, nil
}
func (s *memoryStore) Save(v scheduler.State) error { s.value = v; s.exists = true; return nil }

type memoryLocker struct{}
type memoryLock struct{}

func (memoryLocker) Acquire(string) (scheduler.Lock, error) { return memoryLock{}, nil }
func (memoryLock) Release() error                           { return nil }

type schedulerClock struct{ value scheduler.ClockObservation }

func (c schedulerClock) Observe() scheduler.ClockObservation { return c.value }

type unusedPipeline struct{}

func (unusedPipeline) Execute(context.Context, command.Definition) (command.Execution, error) {
	return command.Execution{}, errors.New("unexpected pipeline call")
}

func fixture(t *testing.T) (Input, scheduler.CycleResult, []time.Time) {
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
	cycle := scheduler.Cycle{Configuration: config, Selection: command.Selection{Source: "live"}, LockOwnerID: "runtime.test", Store: &memoryStore{}, Locker: memoryLocker{}, Clock: schedulerClock{value: scheduler.ClockObservation{WallTime: start, SessionID: "runtime.test"}}, TimeZones: scheduler.SystemTimeZones{}, ResolveCommand: scheduler.ResolveCanonicalCommand, Pipeline: unusedPipeline{}}
	scheduled, err := cycle.Run(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	policyValue, err := notification.NewPolicy(notification.RetryPolicy{MaxAttempts: 1, DeliveryWindowNS: int64(time.Hour), BackoffNS: []int64{}}, []notification.Route{}, []notification.EndpointReference{}, []notification.ProviderBinding{})
	if err != nil {
		t.Fatal(err)
	}
	executionContext, err := NewExecutionContext("cycle.fixture", "owner.explicit", start, start.Add(time.Hour))
	if err != nil {
		t.Fatal(err)
	}
	input := Input{Context: executionContext, Configuration: config, PreviousState: NewState(), PreviousAlertState: alert.NewState(config.ID), PreviousNotificationQueue: notification.NewQueueState(), AlertEvidenceTTLNS: int64(time.Hour), Acknowledgements: []alert.Acknowledgement{}, Suppressions: []alert.SuppressionWindow{}, NotificationPolicy: policyValue}
	return input, scheduled, []time.Time{start, start.Add(time.Second), start.Add(2 * time.Second), start.Add(3 * time.Second)}
}

func contains(value, fragment string) bool {
	for i := 0; i+len(fragment) <= len(value); i++ {
		if value[i:i+len(fragment)] == fragment {
			return true
		}
	}
	return false
}
