package alert

import (
	"bytes"
	"encoding/json"
	"fmt"
	"testing"
	"time"

	"quantumwizard.hu/qwsg/internal/comparison"
	"quantumwizard.hu/qwsg/internal/configuration"
	"quantumwizard.hu/qwsg/internal/drift"
	"quantumwizard.hu/qwsg/internal/health"
	"quantumwizard.hu/qwsg/internal/pipeline"
	"quantumwizard.hu/qwsg/internal/policy"
	"quantumwizard.hu/qwsg/internal/report"
	"quantumwizard.hu/qwsg/internal/rule"
	"quantumwizard.hu/qwsg/internal/scheduler"
)

var baseTime = time.Date(2026, 8, 7, 8, 0, 0, 0, time.UTC)

func TestValidateRecordStandaloneBoundary(t *testing.T) {
	info := healthFixture(t, "standalone", comparison.Added, baseTime)
	result := evaluate(t, inputAt(baseTime, &info, NewState("")))
	record := result.Records[0]
	if err := ValidateRecord(record); err != nil {
		t.Fatalf("valid canonical alert record rejected: %v", err)
	}
	record.ReasonToken = "tampered"
	if err := ValidateRecord(record); err == nil {
		t.Fatal("tampered canonical alert record accepted")
	}
}

func TestHealthLifecycleDedupEscalationRecoveryAndRecurrence(t *testing.T) {
	info := healthFixture(t, "info", comparison.Added, baseTime)
	first := evaluate(t, inputAt(baseTime, &info, NewState("")))
	assertEvent(t, first, EventEntered, Informational, 1)
	if !validCondition(first.NextState.Conditions[0]) {
		t.Fatalf("invalid generated condition: %#v", first.NextState.Conditions[0])
	}

	unchanged := evaluate(t, inputAt(baseTime.Add(time.Hour), &info, first.NextState))
	if len(unchanged.Records) != 0 {
		t.Fatalf("unchanged evidence emitted alerts: %#v", unchanged.Records)
	}

	critical := healthFixture(t, "critical", comparison.Removed, baseTime.Add(2*time.Hour))
	escalated := evaluate(t, inputAt(baseTime.Add(2*time.Hour), &critical, unchanged.NextState))
	assertEvent(t, escalated, EventEscalated, Critical, 1)

	warning := healthFixture(t, "warning", comparison.Modified, baseTime.Add(3*time.Hour))
	deescalated := evaluate(t, inputAt(baseTime.Add(3*time.Hour), &warning, escalated.NextState))
	assertEvent(t, deescalated, EventDeescalated, Warning, 1)

	healthy := healthFixture(t, "healthy", comparison.Unchanged, baseTime.Add(4*time.Hour))
	recovered := evaluate(t, inputAt(baseTime.Add(4*time.Hour), &healthy, deescalated.NextState))
	assertEvent(t, recovered, EventRecovered, Warning, 1)
	if recovered.Records[0].RecoveryTime != baseTime.Add(4*time.Hour) || !recovered.Records[0].ExpirationTime.IsZero() {
		t.Fatalf("recovery and expiration were conflated: %#v", recovered.Records[0])
	}

	recurrenceEvidence := healthFixture(t, "recurrence", comparison.Removed, baseTime.Add(5*time.Hour))
	recurrence := evaluate(t, inputAt(baseTime.Add(5*time.Hour), &recurrenceEvidence, recovered.NextState))
	assertEvent(t, recurrence, EventEntered, Critical, 2)
	if recurrence.Records[0].LifecycleID == first.Records[0].LifecycleID || recurrence.Records[0].ConditionKey != first.Records[0].ConditionKey {
		t.Fatal("recurrence did not preserve condition correlation with a new lifecycle")
	}
}

func TestAcknowledgementAndMaintenanceSuppression(t *testing.T) {
	evidence := healthFixture(t, "condition", comparison.Removed, baseTime)
	probe := inputAt(baseTime, &evidence, NewState(""))
	key := conditionKey(EngineeringCondition, healthSubject(evidence.Records[0]))
	window, err := NewSuppressionWindow(MaintenanceSuppression, []string{key}, []Category{}, []Severity{}, baseTime.Add(15*time.Minute), baseTime.Add(time.Hour), "operator", "local-admin", "planned-maintenance", false)
	if err != nil {
		t.Fatal(err)
	}
	probe.Suppressions = []SuppressionWindow{window}
	active := evaluate(t, probe)
	assertEvent(t, active, EventEntered, Critical, 1)

	suppressionInput := inputAt(baseTime.Add(30*time.Minute), &evidence, active.NextState)
	suppressionInput.Suppressions = []SuppressionWindow{window}
	suppressed := evaluate(t, suppressionInput)
	assertEvent(t, suppressed, EventSuppressionStarted, Critical, 1)
	if suppressed.Records[0].Decision != DecisionSuppressed || suppressed.Records[0].LifecycleState != StateSuppressed {
		t.Fatalf("condition was not retained as suppressed: %#v", suppressed.Records[0])
	}

	after := inputAt(baseTime.Add(2*time.Hour), &evidence, suppressed.NextState)
	after.Suppressions = []SuppressionWindow{window}
	maintenanceEnded := evaluate(t, after)
	assertEvent(t, maintenanceEnded, EventMaintenanceEnded, Critical, 1)

	lifecycle := maintenanceEnded.NextState.Conditions[0].LifecycleID
	ack, err := NewAcknowledgement(lifecycle, "operator", "local-admin", baseTime.Add(3*time.Hour), "investigating")
	if err != nil {
		t.Fatal(err)
	}
	ackInput := inputAt(baseTime.Add(3*time.Hour), &evidence, maintenanceEnded.NextState)
	ackInput.Acknowledgements = []Acknowledgement{ack}
	acknowledged := evaluate(t, ackInput)
	assertEvent(t, acknowledged, EventAcknowledged, Critical, 1)
	if acknowledged.NextState.Conditions[0].Severity != Critical || acknowledged.NextState.Conditions[0].LifecycleState != StateAcknowledged {
		t.Fatal("acknowledgement changed condition severity or failed to record awareness")
	}

	bad, _ := NewAcknowledgement("lifecycle:aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa", "operator", "local-admin", baseTime.Add(3*time.Hour), "investigating")
	ackInput.Acknowledgements = []Acknowledgement{bad}
	if _, err := Evaluate(ackInput); err == nil {
		t.Fatal("accepted acknowledgement for a nonexistent lifecycle")
	}
}

func TestExpirationDoesNotFabricateRecovery(t *testing.T) {
	evidence := healthFixture(t, "condition", comparison.Removed, baseTime)
	firstInput := inputAt(baseTime, &evidence, NewState(""))
	firstInput.EvidenceTTLNS = int64(time.Hour)
	first := evaluate(t, firstInput)

	expireInput := inputAt(baseTime.Add(2*time.Hour), nil, first.NextState)
	expireInput.EvidenceTTLNS = int64(time.Hour)
	expired := evaluate(t, expireInput)
	assertEvent(t, expired, EventExpired, Critical, 1)
	if !expired.Records[0].RecoveryTime.IsZero() || expired.Records[0].ExpirationTime != baseTime.Add(time.Hour) {
		t.Fatalf("expiration fabricated recovery: %#v", expired.Records[0])
	}

	staleEvidence := healthFixture(t, "stale", comparison.Removed, baseTime)
	staleInput := inputAt(baseTime.Add(2*time.Hour), &staleEvidence, NewState(""))
	staleInput.EvidenceTTLNS = int64(time.Hour)
	stale := evaluate(t, staleInput)
	assertEvent(t, stale, EventExpired, Critical, 1)
	if stale.Records[0].ReasonToken != "source_evidence_expired" || !stale.Records[0].RecoveryTime.IsZero() {
		t.Fatalf("stale evidence was not explicit expiration: %#v", stale.Records[0])
	}
}

func TestPolicySupersedesRuleAndHealthAndRemindsEmergency(t *testing.T) {
	healthResult := healthFixture(t, "policy", comparison.Removed, baseTime)
	ruleResult := matchedRule(t, healthResult)
	policyResult := escalatedPolicy(t, ruleResult)
	policyReport, err := report.GeneratePolicy(policyResult)
	if err != nil {
		t.Fatal(err)
	}
	input := inputAt(baseTime, &healthResult, NewState(""))
	input.Rules, input.Policies, input.PolicyReport = &ruleResult, &policyResult, &policyReport
	first := evaluate(t, input)
	if len(first.Records) != 1 || first.Records[0].Source.Kind != PolicySource || first.Records[0].Severity != Emergency {
		t.Fatalf("source precedence produced competing alerts: %#v", first.Records)
	}

	reminderInput := input
	reminderInput.EvaluatedAt = baseTime.Add(EmergencyReminder)
	reminderInput.PreviousState = first.NextState
	reminder := evaluate(t, reminderInput)
	assertEvent(t, reminder, EventReminder, Emergency, 1)
}

func TestCanonicalReportAndSchedulerSources(t *testing.T) {
	emptyHealth := emptyHealth(t)
	emptyRules, err := rule.Evaluate([]rule.Definition{}, emptyHealth)
	if err != nil {
		t.Fatal(err)
	}
	incomplete, err := report.Generate(emptyRules)
	if err != nil {
		t.Fatal(err)
	}
	input := inputAt(baseTime, nil, NewState(""))
	input.Report = &incomplete
	reportResult := evaluate(t, input)
	if len(reportResult.Records) != 1 || reportResult.Records[0].Category != ReportCompleteness || reportResult.Records[0].Source.Kind != ReportSource {
		t.Fatalf("incomplete report did not create one report-level alert: %#v", reportResult.Records)
	}

	config := effectiveConfiguration(t)
	firstSchedule, err := scheduler.Evaluate(config, scheduler.NewState(config.ID), scheduler.ClockObservation{WallTime: baseTime, SessionID: "session", MonotonicNS: 0}, scheduler.SystemTimeZones{})
	if err != nil {
		t.Fatal(err)
	}
	secondSchedule, err := scheduler.Evaluate(config, firstSchedule.NextState, scheduler.ClockObservation{WallTime: baseTime.Add(10 * time.Second), SessionID: "session", MonotonicNS: int64(time.Second)}, scheduler.SystemTimeZones{})
	if err != nil {
		t.Fatal(err)
	}
	schedulerInput := inputAt(baseTime.Add(10*time.Second), nil, NewState(config.ID))
	schedulerInput.Configuration, schedulerInput.Scheduler = &config, &secondSchedule
	schedulerResult := evaluate(t, schedulerInput)
	if len(schedulerResult.Records) != 1 || schedulerResult.Records[0].Category != SchedulerOperation || schedulerResult.Records[0].Severity != Critical {
		t.Fatalf("scheduler discontinuity was not mapped canonically: %#v", schedulerResult.Records)
	}
}

func TestPolicyReportUsesBoundedAggregateEvidenceAtLiveScale(t *testing.T) {
	incomplete := policyReportFixture(t, 0)
	enteredInput := inputAt(baseTime, nil, NewState(""))
	enteredInput.PolicyReport = &incomplete
	entered := evaluate(t, enteredInput)
	assertEvent(t, entered, EventIndeterminate, Warning, 1)

	for _, count := range []int{64, 65, 366, 1024} {
		t.Run(fmt.Sprintf("sources_%d", count), func(t *testing.T) {
			value := policyReportFixture(t, count)
			if len(value.Sources) != count {
				t.Fatalf("report sources=%d", len(value.Sources))
			}
			input := inputAt(baseTime.Add(time.Minute), nil, entered.NextState)
			input.PolicyReport = &value
			first := evaluate(t, input)
			second := evaluate(t, input)
			assertEvent(t, first, EventRecovered, Warning, 1)
			if first.ID != second.ID || len(first.Records[0].Source.EvidenceReferences) != 1 || first.Records[0].Source.EvidenceReferences[0] != value.ID || first.Records[0].Source.RecordID != value.ID {
				t.Fatalf("report envelope correlation is not stable and bounded: %#v", first.Records[0].Source)
			}
			if err := report.ValidatePolicyReport(value); err != nil || len(value.Sources) != count {
				t.Fatalf("full canonical traceability was not preserved: sources=%d err=%v", len(value.Sources), err)
			}
		})
	}
}

func TestDeterminismCanonicalJSONStrictDecodingAndTamperRejection(t *testing.T) {
	evidence := healthFixture(t, "deterministic", comparison.Removed, baseTime)
	input := inputAt(baseTime, &evidence, NewState(""))
	left := evaluate(t, input)
	right := evaluate(t, input)
	leftJSON, err := MarshalCanonical(left)
	if err != nil {
		t.Fatal(err)
	}
	rightJSON, err := MarshalCanonical(right)
	if err != nil {
		t.Fatal(err)
	}
	if !bytes.Equal(leftJSON, rightJSON) || left.ID != right.ID {
		t.Fatal("equivalent explicit input was not byte-stable")
	}
	decoded, err := DecodeResult(leftJSON)
	if err != nil || decoded.ID != left.ID {
		t.Fatalf("canonical result did not round-trip: %v", err)
	}

	tampered := left
	tampered.Records = append([]Record{}, left.Records...)
	tampered.Records[0].Severity = Warning
	if err := ValidateResult(tampered); err == nil {
		t.Fatal("accepted tampered canonical Alert Record")
	}

	inputJSON, err := json.Marshal(input)
	if err != nil {
		t.Fatal(err)
	}
	inputJSON = append(inputJSON[:len(inputJSON)-1], []byte(`,"unknown":true}`)...)
	if _, err := DecodeInput(inputJSON); err == nil {
		t.Fatal("strict decoder accepted an unknown input field")
	}
}

func TestControlAndResourceValidation(t *testing.T) {
	if _, err := NewSuppressionWindow(OperationalSuppression, []string{}, []Category{}, []Severity{Emergency}, baseTime, baseTime.Add(time.Hour), "operator", "local-admin", "quiet", false); err == nil {
		t.Fatal("emergency suppression lacked explicit authorization")
	}
	if _, err := NewSuppressionWindow(OperationalSuppression, []string{}, []Category{}, []Severity{}, baseTime, baseTime.Add(MaxSuppression+time.Second), "operator", "local-admin", "unbounded", false); err == nil {
		t.Fatal("accepted an unbounded suppression window")
	}
	input := inputAt(baseTime, nil, NewState(""))
	input.EvidenceTTLNS = int64(MaxEvidenceTTL + time.Second)
	if _, err := Evaluate(input); err == nil {
		t.Fatal("accepted an unbounded evidence lifetime")
	}
}

func inputAt(at time.Time, healthResult *health.Result, state State) Input {
	return Input{SchemaName: InputSchema, SchemaVersion: SchemaVersion, EvaluatedAt: at, EvidenceTTLNS: int64(48 * time.Hour), Health: healthResult, PreviousState: state, Acknowledgements: []Acknowledgement{}, Suppressions: []SuppressionWindow{}}
}

func evaluate(t *testing.T, input Input) Result {
	t.Helper()
	result, err := Evaluate(input)
	if err != nil {
		t.Fatal(err)
	}
	return result
}

func assertEvent(t *testing.T, result Result, event EventKind, severity Severity, generation int) {
	t.Helper()
	if len(result.Records) != 1 || result.Records[0].Event != event || result.Records[0].Severity != severity || result.Records[0].Generation != generation {
		t.Fatalf("unexpected Alert Record: %#v", result.Records)
	}
}

func healthFixture(t *testing.T, id string, kind comparison.ChangeType, at time.Time) health.Result {
	t.Helper()
	previous := &comparison.TypedValue{Type: "string", Value: "private-before"}
	current := &comparison.TypedValue{Type: "string", Value: "private-after"}
	switch kind {
	case comparison.Added:
		previous = nil
	case comparison.Removed:
		current = nil
	case comparison.Unchanged:
		current = &comparison.TypedValue{Type: "string", Value: "private-before"}
	}
	change := comparison.ChangeRecord{ID: id, Layer: "security", ObjectID: "canonical-subject", Path: "/layers/security/resources/canonical-subject/facts/value", Type: kind, Previous: previous, Current: current, ComparisonTimestamp: at, Metadata: map[string]string{"object_kind": "test", "fact_name": "value"}}
	driftResult, err := drift.Classify([]comparison.ChangeRecord{change})
	if err != nil {
		t.Fatal(err)
	}
	result, err := health.Evaluate(driftResult)
	if err != nil {
		t.Fatal(err)
	}
	return result
}

func emptyHealth(t *testing.T) health.Result {
	t.Helper()
	driftResult, err := drift.Classify([]comparison.ChangeRecord{})
	if err != nil {
		t.Fatal(err)
	}
	result, err := health.Evaluate(driftResult)
	if err != nil {
		t.Fatal(err)
	}
	return result
}

func matchedRule(t *testing.T, input health.Result) rule.Result {
	t.Helper()
	definition := rule.Definition{ID: "canonical.alert.test", ContractVersion: rule.RuleVersion, Category: rule.StatusRule, Scope: rule.Scope{HealthIDs: []string{}, Categories: []drift.Category{}}, Enabled: true, InputRequirements: []rule.Field{rule.FieldStatus}, Condition: rule.Condition{Operator: rule.Exists, Field: rule.FieldStatus, Values: []rule.Value{}, Children: []rule.Condition{}}, Description: "Canonical Alert test rule.", Metadata: map[string]string{}}
	result, err := rule.Evaluate([]rule.Definition{definition}, input)
	if err != nil {
		t.Fatal(err)
	}
	return result
}

func escalatedPolicy(t *testing.T, input rule.Result) policy.Result {
	t.Helper()
	profile, err := policy.NormalizeProfile(policy.Profile{ID: "canonical.alert.profile", ContractVersion: policy.ProfileVersion, Version: "1.0", Priority: 100, Extends: []string{}, Enabled: true, Scope: policy.Selector{RuleIDs: []string{}, Outcomes: []rule.Outcome{}}, Statements: []policy.Statement{{ID: "escalate-match", Priority: 100, Selector: policy.Selector{RuleIDs: []string{}, Outcomes: []rule.Outcome{rule.Matched}}, Outcome: policy.Escalated, Explanation: "canonical_alert_escalation", Metadata: map[string]string{}}}, DefaultOutcome: policy.Accepted, Metadata: map[string]string{}})
	if err != nil {
		t.Fatal(err)
	}
	result, err := policy.Evaluate([]policy.Profile{profile}, input)
	if err != nil {
		t.Fatal(err)
	}
	return result
}

func effectiveConfiguration(t *testing.T) configuration.Effective {
	t.Helper()
	source, err := configuration.BuiltIn([]rule.Definition{}, []policy.Profile{})
	if err != nil {
		t.Fatal(err)
	}
	result, err := configuration.Resolve([]configuration.Source{source})
	if err != nil {
		t.Fatal(err)
	}
	return result
}

func policyReportFixture(t *testing.T, count int) report.PolicyReport {
	t.Helper()
	changes := make([]comparison.ChangeRecord, 0, count)
	for index := 0; index < count; index++ {
		previous := &comparison.TypedValue{Type: "integer", Value: index}
		current := &comparison.TypedValue{Type: "integer", Value: index + 1}
		changes = append(changes, comparison.ChangeRecord{ID: fmt.Sprintf("report-change-%04d", index), Layer: "configuration", ObjectID: fmt.Sprintf("item-%04d", index), Path: fmt.Sprintf("/layers/configuration/resources/item-%04d/facts/value", index), Type: comparison.Modified, Previous: previous, Current: current, ComparisonTimestamp: baseTime, Metadata: map[string]string{"fact_name": "value", "object_kind": "configuration"}})
	}
	drifted, err := drift.Classify(changes)
	if err != nil {
		t.Fatal(err)
	}
	evaluated, err := health.Evaluate(drifted)
	if err != nil {
		t.Fatal(err)
	}
	rules := pipeline.CanonicalObservationRules()
	profiles := pipeline.CanonicalPolicyProfiles()
	if count == 0 {
		rules = []rule.Definition{}
	}
	ruled, err := rule.Evaluate(rules, evaluated)
	if err != nil {
		t.Fatal(err)
	}
	governed, err := policy.Evaluate(profiles, ruled)
	if err != nil {
		t.Fatal(err)
	}
	value, err := report.GeneratePolicy(governed)
	if err != nil {
		t.Fatal(err)
	}
	return value
}
