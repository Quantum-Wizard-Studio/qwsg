package presentationmodel

import (
	"bytes"
	"encoding/json"
	"fmt"
	"math/rand"
	"reflect"
	"testing"
	"time"

	"quantumwizard.hu/qwsg/internal/alert"
	"quantumwizard.hu/qwsg/internal/command"
	"quantumwizard.hu/qwsg/internal/comparison"
	"quantumwizard.hu/qwsg/internal/drift"
	"quantumwizard.hu/qwsg/internal/health"
	"quantumwizard.hu/qwsg/internal/pipeline"
	"quantumwizard.hu/qwsg/internal/policy"
	"quantumwizard.hu/qwsg/internal/rule"
	"quantumwizard.hu/qwsg/internal/runtime"
	"quantumwizard.hu/qwsg/internal/runtimeservice"
)

var testNow = time.Date(2026, 8, 7, 18, 0, 0, 0, time.UTC)

func baseInput() Input {
	return Input{SchemaName: InputSchema, SchemaVersion: SchemaVersion, ObservedAt: testNow, FreshForNS: int64(time.Hour)}
}

func TestMissingEvidenceIsUnavailableAndNeverHealthy(t *testing.T) {
	overview, err := Project(baseInput())
	if err != nil {
		t.Fatal(err)
	}
	if overview.Condition != Unavailable || overview.Attention != AttentionUnknown || overview.Completeness != CompletenessMissing || overview.Freshness != FreshnessNotObserved || overview.Guardian != GuardianNotObserved {
		t.Fatalf("missing evidence was hidden: %+v", overview)
	}
	if len(overview.Recommendations) != 2 || overview.Recommendations[0].Token != RecommendRunFreshCheck || overview.Recommendations[1].Token != RecommendVerifyGuardian {
		t.Fatalf("unexpected recommendations: %+v", overview.Recommendations)
	}
}

func TestStoredOverviewFreshnessRequalificationIsOneWay(t *testing.T) {
	input := baseInput()
	input.Command = &CommandObservation{ObservedAt: testNow, Value: command.Execution{SchemaName: command.ExecutionSchema, SchemaVersion: command.SchemaVersion, ID: "execution", CommandID: "definition", PlanID: "plan", Stages: []command.StageResult{}, View: command.View{Rows: []command.ViewRow{}, Groups: []command.ViewGroup{}}, Diagnostics: []string{}, Complete: true}}
	input.ServiceState = &ServiceStateObservation{ObservedAt: testNow, Value: runtimeservice.NewState("guardian-local")}
	overview, err := Project(input)
	if err != nil {
		t.Fatal(err)
	}
	freshUntil := overview.ObservedAt.Add(time.Hour)
	fresh, err := RequalifyFreshness(overview, freshUntil.Add(-time.Nanosecond), freshUntil)
	if err != nil || fresh.ID != overview.ID {
		t.Fatalf("fresh changed: %v", err)
	}
	stale, err := RequalifyFreshness(overview, freshUntil, freshUntil)
	if err != nil {
		t.Fatal(err)
	}
	if stale.Freshness != FreshnessStale || stale.Completeness != CompletenessPartial || stale.Condition != Degraded || stale.Guardian != GuardianUnavailable || stale.ID == overview.ID {
		t.Fatalf("unexpected stale overview: %+v", stale)
	}
	again, err := RequalifyFreshness(stale, overview.ObservedAt, freshUntil)
	if err != nil || again.ID != stale.ID || again.Freshness != FreshnessStale {
		t.Fatal("stale overview was upgraded")
	}
}

func TestHealthyEvidenceAndDeterministicCanonicalJSON(t *testing.T) {
	comparisonResult, driftResult, healthResult := chain(t, comparison.Unchanged)
	input := baseInput()
	input.Comparison = &ComparisonObservation{ObservedAt: testNow, Value: comparisonResult}
	input.Drift = &DriftObservation{ObservedAt: testNow, Value: driftResult}
	input.Health = &HealthObservation{ObservedAt: testNow, Value: healthResult}
	first, err := Project(input)
	if err != nil {
		t.Fatal(err)
	}
	second, err := Project(input)
	if err != nil {
		t.Fatal(err)
	}
	if first.Condition != Healthy || first.Attention != AttentionNone || first.Freshness != FreshnessCurrent || first.Completeness != CompletenessComplete {
		t.Fatalf("unexpected healthy overview: %+v", first)
	}
	one, _ := MarshalCanonical(first)
	two, _ := MarshalCanonical(second)
	if !bytes.Equal(one, two) || first.ID != second.ID {
		t.Fatal("equivalent input was not deterministic")
	}
	decoded, err := Decode(one)
	if err != nil || decoded.ID != first.ID {
		t.Fatalf("decode failed: %v", err)
	}
}

func TestCriticalPrecedenceAndReadOnlyRecommendations(t *testing.T) {
	comparisonResult, driftResult, healthResult := chain(t, comparison.Removed)
	input := baseInput()
	input.Comparison = &ComparisonObservation{ObservedAt: testNow, Value: comparisonResult}
	input.Drift = &DriftObservation{ObservedAt: testNow, Value: driftResult}
	input.Health = &HealthObservation{ObservedAt: testNow, Value: healthResult}
	overview, err := Project(input)
	if err != nil {
		t.Fatal(err)
	}
	if overview.Condition != Critical || overview.Attention != AttentionUrgent || overview.Summary.CriticalHealth != 1 {
		t.Fatalf("critical evidence was masked: %+v", overview)
	}
	want := []RecommendationToken{RecommendInspectAttention, RecommendReviewChanges, RecommendVerifyGuardian}
	if len(overview.Recommendations) != len(want) {
		t.Fatalf("recommendations: %+v", overview.Recommendations)
	}
	for i := range want {
		if overview.Recommendations[i].Token != want[i] {
			t.Fatalf("recommendation order: %+v", overview.Recommendations)
		}
	}
}

func TestExactFreshnessBoundaryAndStaleEvidence(t *testing.T) {
	_, driftResult, healthResult := chain(t, comparison.Unchanged)
	input := baseInput()
	input.Drift = &DriftObservation{ObservedAt: testNow.Add(-time.Hour), Value: driftResult}
	input.Health = &HealthObservation{ObservedAt: testNow.Add(-time.Hour), Value: healthResult}
	overview, err := Project(input)
	if err != nil {
		t.Fatal(err)
	}
	if overview.Freshness != FreshnessCurrent {
		t.Fatalf("exact boundary is not current: %s", overview.Freshness)
	}
	input.Health.ObservedAt = testNow.Add(-time.Hour - time.Nanosecond)
	overview, err = Project(input)
	if err != nil {
		t.Fatal(err)
	}
	if overview.Freshness != FreshnessStale || overview.Completeness != CompletenessPartial || overview.Condition != Degraded {
		t.Fatalf("stale evidence hidden: %+v", overview)
	}
}

func TestInputStrictnessCorrelationAndMutualExclusion(t *testing.T) {
	input := baseInput()
	input.ObservedAt = testNow.In(time.FixedZone("bad", 0))
	overview, err := Project(input)
	if err != nil || overview.ObservedAt.Location() != time.UTC {
		t.Fatalf("observation was not UTC-normalized: %v", err)
	}
	input = baseInput()
	input.ServiceState = &ServiceStateObservation{ObservedAt: testNow}
	input.ServiceResult = &ServiceResultObservation{ObservedAt: testNow}
	if _, err := Project(input); err == nil {
		t.Fatal("competing service observations accepted")
	}
	input = baseInput()
	input.Drift = &DriftObservation{ObservedAt: testNow, Value: drift.Result{}}
	if _, err := Project(input); err == nil {
		t.Fatal("invalid source accepted")
	}
	input = baseInput()
	document, _ := json.Marshal(input)
	document = append(document[:len(document)-1], []byte(`,"unknown":true}`)...)
	if _, err := DecodeInput(document); err == nil {
		t.Fatal("unknown input field accepted")
	}
	valid, _ := json.Marshal(baseInput())
	if _, err := DecodeInput(append(valid, []byte("{}")...)); err == nil {
		t.Fatal("trailing input accepted")
	}
}

func TestClosedMappingTables(t *testing.T) {
	guardianCases := map[runtimeservice.Lifecycle]Guardian{runtimeservice.Created: GuardianStarting, runtimeservice.Starting: GuardianStarting, runtimeservice.Running: GuardianRunning, runtimeservice.Stopping: GuardianStopping, runtimeservice.Stopped: GuardianStopped, runtimeservice.Failed: GuardianDegraded}
	for input, want := range guardianCases {
		if got := guardianFromLifecycle(input); got != want {
			t.Fatalf("%s: %s", input, got)
		}
	}
	for _, severity := range []alert.Severity{alert.Informational, alert.Warning} {
		if alertAttention(severity) != AttentionReview {
			t.Fatalf("%s mapping", severity)
		}
	}
	for _, severity := range []alert.Severity{alert.Critical, alert.Emergency} {
		if alertAttention(severity) != AttentionUrgent {
			t.Fatalf("%s mapping", severity)
		}
	}
	runtimeCases := map[runtime.Outcome]Attention{runtime.Completed: AttentionNone, runtime.Partial: AttentionReview, runtime.Failed: AttentionUrgent, runtime.Cancelled: AttentionReview, runtime.TimedOut: AttentionUrgent}
	for input, want := range runtimeCases {
		if got := runtimeAttention(input); got != want {
			t.Fatalf("runtime %s: %s", input, got)
		}
	}
	ruleCases := map[rule.Outcome]Attention{rule.Matched: AttentionReview, rule.NotMatched: AttentionNone, rule.InsufficientEvidence: AttentionUnknown, rule.UnsupportedRule: AttentionUnknown, rule.InvalidRule: AttentionUnknown, rule.EvaluationError: AttentionUnknown, rule.DisabledRule: AttentionNone}
	for input, want := range ruleCases {
		if got := ruleAttention(input); got != want {
			t.Fatalf("rule %s: %s", input, got)
		}
	}
	policyCases := map[policy.Outcome]Attention{policy.Accepted: AttentionNone, policy.Observe: AttentionReview, policy.Suppressed: AttentionNone, policy.Escalated: AttentionUrgent, policy.Indeterminate: AttentionReview, policy.NotApplicable: AttentionNone, policy.Conflict: AttentionUrgent}
	for input, want := range policyCases {
		if got := policyAttention(input); got != want {
			t.Fatalf("policy %s: %s", input, got)
		}
	}
	for _, value := range []Condition{Healthy, Degraded, Critical, Unknown, Unavailable} {
		if !validCondition(value) {
			t.Fatalf("condition %s", value)
		}
	}
	for _, value := range []Freshness{FreshnessCurrent, FreshnessStale, FreshnessInvalid, FreshnessNotObserved} {
		if !validFreshness(value) {
			t.Fatalf("freshness %s", value)
		}
	}
	for _, value := range []Completeness{CompletenessComplete, CompletenessPartial, CompletenessUnsupported, CompletenessMissing, CompletenessInvalid, CompletenessNotObserved} {
		if !validCompleteness(value) {
			t.Fatalf("completeness %s", value)
		}
	}
}

func TestCommandValidationAndTamperRejection(t *testing.T) {
	input := baseInput()
	input.Command = &CommandObservation{ObservedAt: testNow, Value: command.Execution{SchemaName: command.ExecutionSchema, SchemaVersion: command.SchemaVersion, ID: "execution", CommandID: "command", PlanID: "plan", Stages: []command.StageResult{}, View: command.View{Rows: []command.ViewRow{}, Groups: []command.ViewGroup{}}, Diagnostics: []string{}, Complete: true}}
	overview, err := Project(input)
	if err != nil {
		t.Fatal(err)
	}
	overview.Condition = Healthy
	if Validate(overview) == nil {
		t.Fatal("tampered overview accepted")
	}
	document, _ := MarshalCanonical(func() Overview { v, _ := Project(input); return v }())
	document = append(document, []byte("{}")...)
	if _, err := Decode(document); err == nil {
		t.Fatal("trailing JSON accepted")
	}
}

func TestExplicitServiceObservationAndAlertLifecycleProjection(t *testing.T) {
	input := baseInput()
	state := runtimeservice.NewState("guardian-local")
	input.ServiceState = &ServiceStateObservation{ObservedAt: testNow, Value: state}
	overview, err := Project(input)
	if err != nil {
		t.Fatal(err)
	}
	if overview.Guardian != GuardianStarting || overview.Attention != AttentionReview || overview.Condition != Degraded {
		t.Fatalf("explicit service state lost: %+v", overview)
	}
	b := builder{overview: Overview{AttentionItems: []AttentionItem{}}}
	sourceTime := testNow
	cases := []struct {
		state                                 alert.LifecycleState
		active, ack, supp, recovered, expired int
		visible                               bool
	}{{alert.StateActive, 1, 0, 0, 0, 0, true}, {alert.StateAcknowledged, 0, 1, 0, 0, 0, true}, {alert.StateSuppressed, 0, 0, 1, 0, 0, false}, {alert.StateResolved, 0, 0, 0, 1, 0, false}, {alert.StateExpired, 0, 0, 0, 0, 1, false}}
	for index, c := range cases {
		before := len(b.attentionCandidates)
		b.consumeAlert(alert.Record{ID: fmt.Sprintf("alert-%d", index), LifecycleState: c.state, Severity: alert.Warning, ReasonToken: "canonical_reason"}, sourceTime)
		if len(b.attentionCandidates) > before != c.visible {
			t.Fatalf("visibility %s", c.state)
		}
	}
	if b.overview.Summary.ActiveAlerts != 1 || b.overview.Summary.AcknowledgedAlerts != 1 || b.overview.Summary.SuppressedAlerts != 1 || b.overview.Summary.RecoveredAlerts != 1 || b.overview.Summary.ExpiredAlerts != 1 {
		t.Fatalf("alert lifecycle counts: %+v", b.overview.Summary)
	}
}

func TestPrivacyAndResourceLimits(t *testing.T) {
	comparisonResult, driftResult, healthResult := chain(t, comparison.Removed)
	input := baseInput()
	input.Comparison = &ComparisonObservation{ObservedAt: testNow, Value: comparisonResult}
	input.Drift = &DriftObservation{ObservedAt: testNow, Value: driftResult}
	input.Health = &HealthObservation{ObservedAt: testNow, Value: healthResult}
	overview, err := Project(input)
	if err != nil {
		t.Fatal(err)
	}
	document, err := MarshalCanonical(overview)
	if err != nil {
		t.Fatal(err)
	}
	for _, forbidden := range [][]byte{[]byte("/layers/"), []byte("enabled"), []byte("before"), []byte("credential"), []byte("provider_payload")} {
		if bytes.Contains(document, forbidden) {
			t.Fatalf("private source value leaked: %s", forbidden)
		}
	}
	overview.AttentionItems = make([]AttentionItem, MaxAttention+1)
	overview.ID = overviewID(overview)
	if Validate(overview) == nil {
		t.Fatal("attention resource limit accepted")
	}
}

func TestLargeAttentionProjectionIsCorrelatedPrioritizedAndDeterministic(t *testing.T) {
	firstInput := largeAttentionInput(t, 367, 1)
	secondInput := largeAttentionInput(t, 367, 99)
	first, err := Project(firstInput)
	if err != nil {
		t.Fatal(err)
	}
	second, err := Project(secondInput)
	if err != nil {
		t.Fatal(err)
	}
	if first.AttentionSummary == nil {
		t.Fatal("large projection did not disclose reduction")
	}
	summary := first.AttentionSummary
	if summary.TotalCandidates <= 732 || summary.Represented >= 16 || summary.CorrelatedDuplicates <= 700 || summary.TotalCandidates != summary.Represented+summary.CorrelatedDuplicates+summary.Omitted {
		t.Fatalf("invalid reduction accounting: %+v", summary)
	}
	if len(first.AttentionItems) != summary.Represented || first.AttentionItems[0].Severity != AttentionUrgent || first.AttentionItems[0].Source.Kind != "health_record" {
		t.Fatalf("late urgent concern was not preserved: %+v", first.AttentionItems[0])
	}
	seenProjection := map[string]bool{}
	for _, item := range first.AttentionItems {
		if item.Source.Kind == "rule_evaluation" {
			t.Fatal("correlated intermediate Rule concern was retained")
		}
		if !validReference(item.Source) {
			t.Fatalf("untraceable represented concern: %+v", item)
		}
		key := string(item.Severity) + "\x00" + item.TitleToken + "\x00" + item.ReasonToken + "\x00" + item.Source.Kind + "\x00" + item.Source.Contract + "\x00" + item.Source.Version
		if seenProjection[key] {
			t.Fatalf("duplicate operator meaning retained: %s", key)
		}
		seenProjection[key] = true
	}
	one, _ := MarshalCanonical(first)
	two, _ := MarshalCanonical(second)
	if first.ID != second.ID || !bytes.Equal(one, two) {
		t.Fatal("equivalent shuffled input changed bounded projection")
	}
	tampered := first
	copy := *tampered.AttentionSummary
	copy.Omitted++
	tampered.AttentionSummary = &copy
	tampered.ID = overviewID(tampered)
	if Validate(tampered) == nil {
		t.Fatal("inconsistent overflow accounting accepted")
	}
}

func TestLegacyOverviewAndRecommendationCompatibility(t *testing.T) {
	legacy, err := Project(baseInput())
	if err != nil {
		t.Fatal(err)
	}
	legacy.SchemaVersion = LegacySchemaVersion
	legacy.ModelVersion = LegacyModelVersion
	legacy.AttentionSummary = nil
	legacy.ID = overviewID(legacy)
	document, err := MarshalCanonical(legacy)
	if err != nil {
		t.Fatal(err)
	}
	decoded, err := Decode(document)
	if err != nil || decoded.SchemaVersion != LegacySchemaVersion {
		t.Fatalf("legacy overview rejected: %v", err)
	}

	b := builder{input: Input{Runtime: &RuntimeObservation{Value: runtime.Result{Outcome: runtime.Partial}}}, overview: Overview{Freshness: FreshnessCurrent, Completeness: CompletenessPartial, Guardian: GuardianNotObserved, Recommendations: []Recommendation{}}}
	b.recommend()
	seen := map[RecommendationToken]bool{}
	for _, recommendation := range b.overview.Recommendations {
		seen[recommendation.Token] = true
	}
	if !seen[RecommendInspectEvidence] || !seen[RecommendInspectFailedOperation] || seen[RecommendRunFreshCheck] {
		t.Fatalf("non-self-healing partial result recommends blind retry: %+v", b.overview.Recommendations)
	}
}

func TestRuntimeFailureDiagnosticsRemainSpecificAndBounded(t *testing.T) {
	result := runtime.Result{Outcome: runtime.Partial, Components: []runtime.ComponentResult{
		{Component: runtime.AlertComponent, Status: runtime.ComponentFailed, FailureToken: "alert_evaluation_failed"},
		{Component: runtime.NotificationPlanComponent, Status: runtime.ComponentFailed, FailureToken: "notification_planning_failed"},
	}}
	got := runtimeFailureTokens(result)
	if !reflect.DeepEqual(got, []string{"alert_evaluation_failed", "notification_planning_failed"}) {
		t.Fatalf("specific Runtime failures were lost: %v", got)
	}
	for outcome, want := range map[runtime.Outcome]string{
		runtime.Cancelled: "runtime_cancelled",
		runtime.TimedOut:  "runtime_timeout",
		runtime.Partial:   "runtime_not_completed",
	} {
		if got := runtimeOutcomeFailureToken(outcome); got != want {
			t.Fatalf("%s mapped to %s", outcome, got)
		}
	}
}

func largeAttentionInput(t *testing.T, count int, seed int64) Input {
	t.Helper()
	changes := make([]comparison.ChangeRecord, 0, count)
	for index := 0; index < count-1; index++ {
		previous := &comparison.TypedValue{Type: "integer", Value: index}
		current := &comparison.TypedValue{Type: "integer", Value: index + 1}
		changes = append(changes, comparison.ChangeRecord{ID: fmt.Sprintf("change-%04d", index), Layer: "configuration", ObjectID: fmt.Sprintf("item-%04d", index), Path: fmt.Sprintf("/layers/configuration/resources/item-%04d/facts/value", index), Type: comparison.Modified, Previous: previous, Current: current, ComparisonTimestamp: testNow, Metadata: map[string]string{"fact_name": "value", "object_kind": "configuration"}})
	}
	changes = append(changes, comparison.ChangeRecord{ID: "change-critical", Layer: "security", ObjectID: "z-critical", Path: "/layers/security/resources/z-critical/facts/enabled", Type: comparison.Removed, Previous: &comparison.TypedValue{Type: "boolean", Value: true}, ComparisonTimestamp: testNow, Metadata: map[string]string{"fact_name": "enabled", "object_kind": "security_control"}})
	rand.New(rand.NewSource(seed)).Shuffle(len(changes), func(i, j int) { changes[i], changes[j] = changes[j], changes[i] })
	drifted, err := drift.Classify(changes)
	if err != nil {
		t.Fatal(err)
	}
	evaluated, err := health.Evaluate(drifted)
	if err != nil {
		t.Fatal(err)
	}
	ruled, err := rule.Evaluate(pipeline.CanonicalObservationRules(), evaluated)
	if err != nil {
		t.Fatal(err)
	}
	governed, err := policy.Evaluate(pipeline.CanonicalPolicyProfiles(), ruled)
	if err != nil {
		t.Fatal(err)
	}
	input := baseInput()
	input.Health = &HealthObservation{ObservedAt: testNow, Value: evaluated}
	input.Rule = &RuleObservation{ObservedAt: testNow, Value: ruled}
	input.Policy = &PolicyObservation{ObservedAt: testNow, Value: governed}
	return input
}

func chain(t *testing.T, changeType comparison.ChangeType) (comparison.Result, drift.Result, health.Result) {
	t.Helper()
	value := &comparison.TypedValue{Type: "string", Value: "enabled"}
	record := comparison.ChangeRecord{ID: "change-security", Layer: "security", ObjectID: "ssh", Path: "/layers/security/resources/ssh/facts/enabled", Type: changeType, ComparisonTimestamp: testNow, Metadata: map[string]string{}}
	switch changeType {
	case comparison.Unchanged:
		record.Previous = value
		record.Current = &comparison.TypedValue{Type: "string", Value: "enabled"}
	case comparison.Removed:
		record.Previous = value
	}
	counts := comparison.Counts{}
	if changeType == comparison.Unchanged {
		counts.Unchanged = 1
	} else {
		counts.Removed = 1
	}
	c := comparison.Result{SchemaName: comparison.SchemaName, SchemaVersion: comparison.SchemaVersion, EngineVersion: comparison.EngineVersion, ComparisonID: "comparison", SubjectID: "subject", InventorySchema: "1.0", InventoryProfile: "canonical", ComparisonTimestamp: testNow, From: comparison.SnapshotReference{Selector: "from"}, To: comparison.SnapshotReference{Selector: "to"}, Counts: counts, Changes: []comparison.ChangeRecord{record}, Metadata: map[string]string{}}
	if err := comparison.Validate(c); err != nil {
		t.Fatal(err)
	}
	d, err := drift.Classify(c.Changes)
	if err != nil {
		t.Fatal(err)
	}
	h, err := health.Evaluate(d)
	if err != nil {
		t.Fatal(err)
	}
	return c, d, h
}
