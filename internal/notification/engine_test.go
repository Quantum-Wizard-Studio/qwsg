package notification

import (
	"bytes"
	"context"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"sort"
	"testing"
	"time"

	"quantumwizard.hu/qwsg/internal/alert"
)

var testTime = time.Date(2026, 8, 7, 12, 0, 0, 0, time.UTC)

type fakeProvider struct {
	descriptor ProviderDescriptor
	result     ProviderResult
	calls      int
}

func (f *fakeProvider) Descriptor() ProviderDescriptor { return f.descriptor }
func (f *fakeProvider) Deliver(_ context.Context, _ Request) ProviderResult {
	f.calls++
	return f.result
}

func TestDeterministicPlanningFanoutAndAlertBoundary(t *testing.T) {
	record := alertRecord(t, alert.DecisionAlert, alert.StateActive, alert.EventEntered)
	policy := deliveryPolicy(t, []string{"endpoint-email", "endpoint-webhook"})
	input := planningInput(testTime, []alert.Record{record}, policy, NewQueueState())
	first := plan(t, input)
	second := plan(t, input)
	if first.ID != second.ID || len(first.Requests) != 2 || !bytes.Equal(mustPlanJSON(t, first), mustPlanJSON(t, second)) {
		t.Fatalf("planning was not deterministic: %#v %#v", first, second)
	}
	if len(first.Eligibility) != 1 || !first.Eligibility[0].Eligible || first.Eligibility[0].ReasonToken != "delivery_planned" {
		t.Fatalf("unexpected eligibility: %#v", first.Eligibility)
	}
	for _, request := range first.Requests {
		if request.Envelope.AlertRecordID != record.ID || request.Envelope.Severity != record.Severity || request.IdempotencyKey != request.DeliveryID {
			t.Fatalf("alert record was not preserved: %#v", request)
		}
	}

	suppressed := alertRecord(t, alert.DecisionSuppressed, alert.StateSuppressed, alert.EventSuppressionStarted)
	suppressedPlan := plan(t, planningInput(testTime, []alert.Record{suppressed}, policy, NewQueueState()))
	if len(suppressedPlan.Requests) != 0 || suppressedPlan.Eligibility[0].Eligible || suppressedPlan.Eligibility[0].ReasonToken != "alert_record_suppressed" {
		t.Fatalf("suppressed alert was delivered: %#v", suppressedPlan)
	}
}

func TestRouteMatchingCanonicalPolicyAndReplayDeduplication(t *testing.T) {
	record := alertRecord(t, alert.DecisionLifecycle, alert.StateResolved, alert.EventRecovered)
	policyA := deliveryPolicy(t, []string{"endpoint-webhook", "endpoint-email"})
	policyB := deliveryPolicy(t, []string{"endpoint-email", "endpoint-webhook"})
	if policyA.ID != policyB.ID {
		t.Fatal("equivalent policy ordering changed identity")
	}
	first := plan(t, planningInput(testTime, []alert.Record{record}, policyA, NewQueueState()))
	replay := plan(t, planningInput(testTime.Add(time.Minute), []alert.Record{record}, policyA, first.NextQueue))
	if len(replay.Requests) != 0 || len(replay.NextQueue.Entries) != 2 {
		t.Fatalf("replay duplicated delivery: %#v", replay)
	}

	filtered := policyA
	filtered.Routes[0].Severities = []alert.Severity{alert.Emergency}
	filtered.ID = contentIdentity("notification-policy", filtered)
	noMatch := plan(t, planningInput(testTime, []alert.Record{record}, filtered, NewQueueState()))
	if len(noMatch.Requests) != 0 || noMatch.Eligibility[0].ReasonToken != "no_matching_route" {
		t.Fatalf("route filter ignored: %#v", noMatch)
	}
}

func TestQueuePriorityUsesSeverityThenEventTime(t *testing.T) {
	critical := alertRecord(t, alert.DecisionAlert, alert.StateActive, alert.EventEntered)
	critical.EventTime = testTime.Add(time.Minute)
	critical.EvaluationTime = critical.EventTime
	critical.ID = alertRecordIdentity(critical)
	emergency := alertRecord(t, alert.DecisionAlert, alert.StateActive, alert.EventEntered)
	emergency.Severity = alert.Emergency
	emergency.ConditionKey = "condition:emergency"
	emergency.LifecycleID = "lifecycle:emergency"
	emergency.ID = alertRecordIdentity(emergency)
	policy := deliveryPolicy(t, []string{"endpoint-email"})
	planned := plan(t, planningInput(testTime.Add(time.Minute), []alert.Record{critical, emergency}, policy, NewQueueState()))
	if len(planned.Requests) != 2 || planned.Requests[0].Envelope.Severity != alert.Emergency || planned.NextQueue.Entries[0].Request.Envelope.Severity != alert.Emergency {
		t.Fatalf("queue priority was not severity-first: %#v", planned.Requests)
	}
}

func TestOneCycleProviderAcknowledgementAndEvidence(t *testing.T) {
	record := alertRecord(t, alert.DecisionAlert, alert.StateActive, alert.EventEntered)
	policy := deliveryPolicy(t, []string{"endpoint-email"})
	planned := plan(t, planningInput(testTime, []alert.Record{record}, policy, NewQueueState()))
	provider := &fakeProvider{
		descriptor: descriptor("provider-email", []Channel{Email}),
		result:     ProviderResult{SchemaName: ProviderResultSchema, SchemaVersion: SchemaVersion, Status: StatusDelivered, Failure: FailureNone, CompletedAt: testTime.Add(time.Second), ProviderReference: "provider:message-1", EvidenceTokens: []string{"accepted", "delivered"}},
	}
	registry, err := NewRegistry(provider)
	if err != nil {
		t.Fatal(err)
	}
	cycle, err := ExecuteCycle(context.Background(), planned, registry, testTime)
	if err != nil {
		t.Fatal(err)
	}
	if provider.calls != 1 || len(cycle.Attempts) != 1 || cycle.Attempts[0].Status.Status != StatusDelivered {
		t.Fatalf("unexpected cycle result: %#v", cycle)
	}
	ack := cycle.Attempts[0].Acknowledgement
	if ack == nil || ack.Kind != AcknowledgedDelivered || ack.RequestID != planned.Requests[0].ID {
		t.Fatalf("provider acknowledgement was confused with delivery evidence: %#v", ack)
	}
	if cycle.Attempts[0].Evidence.AlertRecordID != record.ID || cycle.NextQueue.Entries[0].Status != StatusDelivered {
		t.Fatalf("delivery evidence was incomplete: %#v", cycle)
	}
	if _, err := MarshalCanonicalCycle(cycle); err != nil {
		t.Fatal(err)
	}
}

func TestRetrySchedulingDeadlineAndExhaustion(t *testing.T) {
	record := alertRecord(t, alert.DecisionAlert, alert.StateActive, alert.EventEntered)
	policy := deliveryPolicy(t, []string{"endpoint-email"})
	first := plan(t, planningInput(testTime, []alert.Record{record}, policy, NewQueueState()))
	provider := &fakeProvider{
		descriptor: descriptor("provider-email", []Channel{Email}),
		result:     ProviderResult{SchemaName: ProviderResultSchema, SchemaVersion: SchemaVersion, Status: StatusRetryableFailure, Failure: FailureRateLimited, CompletedAt: testTime.Add(time.Second), EvidenceTokens: []string{"rate_limited"}},
	}
	registry, _ := NewRegistry(provider)
	cycle, err := ExecuteCycle(context.Background(), first, registry, testTime)
	if err != nil {
		t.Fatal(err)
	}

	early := plan(t, planningInput(testTime.Add(30*time.Second), []alert.Record{record}, policy, cycle.NextQueue))
	if len(early.Requests) != 0 || early.NextQueue.Entries[0].Status != StatusRetryScheduled || early.NextQueue.Entries[0].NextAttemptAt != testTime.Add(time.Minute+time.Second) {
		t.Fatalf("retry was not scheduled deterministically: %#v", early.NextQueue)
	}
	due := plan(t, planningInput(testTime.Add(time.Minute+time.Second), []alert.Record{record}, policy, early.NextQueue))
	if len(due.Requests) != 1 || due.Requests[0].AttemptNumber != 2 || due.Requests[0].IdempotencyKey != first.Requests[0].IdempotencyKey {
		t.Fatalf("due retry was invalid: %#v", due.Requests)
	}

	exhaustedQueue := cycle.NextQueue
	exhaustedQueue.Entries[0].Request.Deadline = testTime.Add(30 * time.Second)
	exhaustedQueue.Entries[0].Request.ID = stableID("request", exhaustedQueue.Entries[0].Request.DeliveryID, "1")
	exhaustedQueue.ID = queueIdentity(exhaustedQueue)
	exhausted := plan(t, planningInput(testTime.Add(30*time.Second), []alert.Record{record}, policy, exhaustedQueue))
	if exhausted.NextQueue.Entries[0].Status != StatusExhausted || len(exhausted.Requests) != 0 {
		t.Fatalf("delivery window did not exhaust retries: %#v", exhausted)
	}
}

func TestUnknownProviderIsolationAndProviderValidation(t *testing.T) {
	record := alertRecord(t, alert.DecisionAlert, alert.StateActive, alert.EventEntered)
	policy := deliveryPolicy(t, []string{"endpoint-email", "endpoint-webhook"})
	planned := plan(t, planningInput(testTime, []alert.Record{record}, policy, NewQueueState()))
	empty, _ := NewRegistry()
	cycle, err := ExecuteCycle(context.Background(), planned, empty, testTime)
	if err != nil {
		t.Fatal(err)
	}
	if len(cycle.Attempts) != 2 {
		t.Fatalf("provider failures were not isolated: %#v", cycle)
	}
	for _, attempt := range cycle.Attempts {
		if attempt.Status.Status != StatusTerminalFailure || attempt.Status.Failure != FailureUnsupportedProvider || attempt.Acknowledgement != nil {
			t.Fatalf("unsupported provider classification failed: %#v", attempt)
		}
	}
	duplicate := &fakeProvider{descriptor: descriptor("same", []Channel{Email}), result: ProviderResult{}}
	if _, err := NewRegistry(duplicate, duplicate); err == nil {
		t.Fatal("duplicate provider accepted")
	}
}

func TestStrictValidationPrivacyBoundsAndTampering(t *testing.T) {
	record := alertRecord(t, alert.DecisionAlert, alert.StateActive, alert.EventEntered)
	policy := deliveryPolicy(t, []string{"endpoint-email"})
	input := planningInput(testTime, []alert.Record{record}, policy, NewQueueState())
	data, _ := json.Marshal(input)
	data = append(data[:len(data)-1], []byte(`,"credential":"secret"}`)...)
	if _, err := DecodePlanningInput(data); err == nil {
		t.Fatal("unknown credential field accepted")
	}

	planned := plan(t, input)
	encoded := mustPlanJSON(t, planned)
	decoded, err := DecodePlan(encoded)
	if err != nil || decoded.ID != planned.ID {
		t.Fatalf("canonical plan round trip failed: %v", err)
	}
	planned.Requests[0].Envelope.ReasonToken = "tampered"
	if err := ValidatePlan(planned); err == nil {
		t.Fatal("tampered request accepted")
	}

	badPolicy := policy
	badPolicy.Endpoints[0].DestinationRef = "https://contains/non-token"
	badPolicy.ID = contentIdentity("notification-policy", badPolicy)
	if err := ValidatePolicy(badPolicy); err == nil {
		t.Fatal("raw destination accepted")
	}

	over := make([]alert.Record, MaxAlerts+1)
	input.Alerts = over
	if err := ValidatePlanningInput(input); err == nil {
		t.Fatal("alert resource limit ignored")
	}

	endpoints := make([]EndpointReference, MaxEndpoints)
	ids := make([]string, MaxEndpoints)
	for i := range endpoints {
		ids[i] = stableID("endpoint", string(rune(i+1)))
		endpoints[i] = EndpointReference{ID: ids[i], Channel: Email, DestinationRef: stableID("destination", ids[i])}
	}
	routes := make([]Route, 9)
	for i := range routes {
		routes[i] = Route{ID: stableID("route", string(rune(i+1))), Enabled: true, Severities: []alert.Severity{}, Categories: []alert.Category{}, Events: []alert.EventKind{}, EndpointIDs: append([]string{}, ids...)}
	}
	if _, err := NewPolicy(policy.Retry, routes, endpoints, []ProviderBinding{{ID: "binding-email", Channel: Email, ProviderID: "provider-email"}}); err == nil {
		t.Fatal("aggregate route fanout limit ignored")
	}
}

func planningInput(at time.Time, records []alert.Record, policy Policy, queue QueueState) PlanningInput {
	sort.Slice(records, func(i, j int) bool { return records[i].ID < records[j].ID })
	return PlanningInput{SchemaName: InputSchema, SchemaVersion: SchemaVersion, EvaluatedAt: at, Alerts: records, Policy: policy, PreviousQueue: queue}
}

func deliveryPolicy(t *testing.T, endpointIDs []string) Policy {
	t.Helper()
	endpoints := []EndpointReference{
		{ID: "endpoint-email", Channel: Email, DestinationRef: "destination:operations", SecretRef: "secret:email-provider"},
		{ID: "endpoint-webhook", Channel: Webhook, DestinationRef: "destination:automation", SecretRef: "secret:webhook-provider"},
	}
	route := Route{ID: "route-primary", Enabled: true, Severities: []alert.Severity{}, Categories: []alert.Category{}, Events: []alert.EventKind{}, EndpointIDs: endpointIDs}
	policy, err := NewPolicy(RetryPolicy{MaxAttempts: 3, DeliveryWindowNS: int64(15 * time.Minute), BackoffNS: []int64{int64(time.Minute), int64(5 * time.Minute)}}, []Route{route}, endpoints, []ProviderBinding{{ID: "binding-email", Channel: Email, ProviderID: "provider-email"}, {ID: "binding-webhook", Channel: Webhook, ProviderID: "provider-webhook"}})
	if err != nil {
		t.Fatal(err)
	}
	return policy
}

func descriptor(id string, channels []Channel) ProviderDescriptor {
	sort.Slice(channels, func(i, j int) bool { return channels[i] < channels[j] })
	return ProviderDescriptor{SchemaName: ProviderSchema, SchemaVersion: SchemaVersion, ID: id, Channels: channels}
}

func alertRecord(t *testing.T, decision alert.Decision, state alert.LifecycleState, event alert.EventKind) alert.Record {
	t.Helper()
	record := alert.Record{
		SchemaName: alert.RecordSchema, SchemaVersion: alert.SchemaVersion,
		ConditionKey: "condition:test", LifecycleID: "lifecycle:test", Generation: 1,
		Event: event, Decision: decision, LifecycleState: state, Severity: alert.Critical,
		Category:  alert.EngineeringCondition,
		Source:    alert.SourceReference{Kind: alert.HealthSource, SchemaName: "qwsg.health-record", SchemaVersion: alert.SchemaVersion, RecordID: "health:test", Subject: "subject:test", EvidenceReferences: []string{"evidence:test"}},
		EventTime: testTime, ObservationTime: testTime, EvaluationTime: testTime,
		SuppressionIDs: []string{}, ReasonToken: "condition_detected",
		Versions: alert.VersionInfo{AlertSchema: alert.SchemaVersion, AlertEngine: alert.EngineVersion, AlertTaxonomy: alert.TaxonomyVersion, AlertModel: alert.ModelVersion},
	}
	if event == alert.EventRecovered {
		record.RecoveryTime = testTime
	}
	if event == alert.EventExpired {
		record.ExpirationTime = testTime
	}
	record.ID = alertRecordIdentity(record)
	if err := alert.ValidateRecord(record); err != nil {
		t.Fatalf("invalid fixture: %v", err)
	}
	return record
}

func alertRecordIdentity(v alert.Record) string {
	v.ID = ""
	data, _ := json.Marshal(v)
	sum := sha256.Sum256(data)
	return "alert:" + hex.EncodeToString(sum[:])
}

func plan(t *testing.T, input PlanningInput) Plan {
	t.Helper()
	v, err := PlanDeliveries(input)
	if err != nil {
		t.Fatal(err)
	}
	return v
}

func mustPlanJSON(t *testing.T, value Plan) []byte {
	t.Helper()
	data, err := MarshalCanonicalPlan(value)
	if err != nil {
		t.Fatal(err)
	}
	return data
}
