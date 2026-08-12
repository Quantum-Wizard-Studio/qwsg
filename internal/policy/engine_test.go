package policy

import (
	"encoding/json"
	"strings"
	"testing"
	"time"

	"quantumwizard.hu/qwsg/internal/comparison"
	"quantumwizard.hu/qwsg/internal/drift"
	"quantumwizard.hu/qwsg/internal/health"
	"quantumwizard.hu/qwsg/internal/rule"
)

func TestDeterministicPolicyEvaluationAndTraceability(t *testing.T) {
	input := ruleInput(t)
	first := normalized(t, profile("ops.primary", 10, Observe, statement("matched.observe", 20, Observe, rule.Matched)))
	second := normalized(t, profile("ops.secondary", 5, Accepted, statement("matched.accept", 20, Accepted, rule.Matched)))
	left, err := Evaluate([]Profile{first, second}, input)
	if err != nil {
		t.Fatal(err)
	}
	right, err := Evaluate([]Profile{second, first}, input)
	if err != nil {
		t.Fatal(err)
	}
	a, _ := MarshalCanonical(left)
	b, _ := MarshalCanonical(right)
	if string(a) != string(b) {
		t.Fatalf("non-deterministic policy output:\n%s\n%s", a, b)
	}
	if len(left.Records) != 1 || left.Records[0].Outcome != Observe || left.Records[0].AppliedProfileIDs[0] != "ops.primary" ||
		left.Records[0].EvidenceReferences[0] != input.Records[0].ID {
		t.Fatalf("unexpected policy result: %#v", left)
	}
	before, _ := json.Marshal(input)
	_, _ = Evaluate([]Profile{first}, input)
	after, _ := json.Marshal(input)
	if string(before) != string(after) {
		t.Fatal("policy mutated Rule evidence")
	}
}

func TestEqualPrecedenceConflictIsExplicit(t *testing.T) {
	input := ruleInput(t)
	a := normalized(t, profile("conflict.a", 10, Accepted, statement("decision.a", 10, Accepted, rule.Matched)))
	b := normalized(t, profile("conflict.b", 10, Observe, statement("decision.b", 10, Escalated, rule.Matched)))
	result, err := Evaluate([]Profile{b, a}, input)
	if err != nil {
		t.Fatal(err)
	}
	record := result.Records[0]
	if record.Outcome != Conflict || record.EvaluationStatus != EvaluationComplete || len(record.AppliedProfileIDs) != 2 ||
		record.Explanation != "equal_precedence_policy_conflict" {
		t.Fatalf("conflict was hidden: %#v", record)
	}
}

func TestInheritanceCompositionAndDefault(t *testing.T) {
	input := ruleInput(t)
	base := normalized(t, profile("base.profile", 1, Indeterminate, statement("base.matched", 7, Escalated, rule.Matched)))
	child := profile("child.profile", 20, Accepted)
	child.Extends = []string{"base.profile"}
	child = normalized(t, child)
	result, err := Evaluate([]Profile{base, child}, input)
	if err != nil {
		t.Fatal(err)
	}
	if result.Records[0].Outcome != Escalated || result.Records[0].AppliedProfileIDs[0] != "child.profile" ||
		result.Records[0].AppliedStatementIDs[0] != "base.matched" {
		t.Fatalf("inheritance was not applied: %#v", result.Records[0])
	}

	nonmatching := profile("default.profile", 30, Suppressed, statement("only.errors", 10, Escalated, rule.EvaluationError))
	nonmatching = normalized(t, nonmatching)
	result, err = Evaluate([]Profile{nonmatching}, input)
	if err != nil {
		t.Fatal(err)
	}
	if result.Records[0].Outcome != Suppressed || result.Records[0].Explanation != "profile_default_outcome" {
		t.Fatalf("default was not explicit: %#v", result.Records[0])
	}
}

func TestScopeNotApplicableAndIndeterminate(t *testing.T) {
	input := ruleInput(t)
	scoped := profile("scoped.profile", 1, Indeterminate)
	scoped.Scope.RuleIDs = []string{"other.rule"}
	scoped = normalized(t, scoped)
	result, err := Evaluate([]Profile{scoped}, input)
	if err != nil {
		t.Fatal(err)
	}
	if result.Records[0].Outcome != NotApplicable || result.Records[0].EvaluationStatus != EvaluationSkipped {
		t.Fatalf("inapplicable profile was treated as policy: %#v", result.Records[0])
	}

	unknown := normalized(t, profile("unknown.profile", 1, Indeterminate))
	result, err = Evaluate([]Profile{unknown}, input)
	if err != nil {
		t.Fatal(err)
	}
	if result.Records[0].Outcome != Indeterminate {
		t.Fatalf("missing policy evidence became acceptable: %#v", result.Records[0])
	}
}

func TestInvalidUnsupportedProfilesAndBoundsFailClosed(t *testing.T) {
	input := ruleInput(t)
	valid := normalized(t, profile("valid.profile", 1, Accepted))
	cases := []Profile{valid, valid, valid, valid}
	cases[0].ContractVersion = "2.0"
	cases[1].Identity = "tampered"
	cases[2].Extends = []string{"missing.profile"}
	cases[2] = resign(t, cases[2])
	cases[3].Statements = []Statement{{ID: "bad", Priority: 0, Selector: Selector{RuleIDs: []string{}, Outcomes: []rule.Outcome{}}, Outcome: Conflict, Explanation: "bad", Metadata: map[string]string{}}}
	cases[3] = resignWithoutValidation(cases[3])
	for _, candidate := range cases {
		if _, err := Evaluate([]Profile{candidate}, input); err == nil {
			t.Fatalf("accepted invalid profile: %#v", candidate)
		}
	}
	a := profile("cycle.a", 1, Accepted)
	a.Extends = []string{"cycle.b"}
	a = normalized(t, a)
	b := profile("cycle.b", 1, Accepted)
	b.Extends = []string{"cycle.a"}
	b = normalized(t, b)
	if _, err := Evaluate([]Profile{a, b}, input); err == nil || !strings.Contains(err.Error(), "cycle") {
		t.Fatalf("inheritance cycle accepted: %v", err)
	}
}

func TestCanonicalProfileAndResultSerializationRejectTampering(t *testing.T) {
	input := ruleInput(t)
	p := normalized(t, profile("stable.profile", 1, Observe, statement("stable.statement", 1, Observe, rule.Matched)))
	if _, err := MarshalProfilesCanonical([]Profile{p}); err != nil {
		t.Fatal(err)
	}
	result, err := Evaluate([]Profile{p}, input)
	if err != nil {
		t.Fatal(err)
	}
	tampered := result
	tampered.Records = append([]EvaluationRecord(nil), result.Records...)
	tampered.Records[0].Outcome = Escalated
	if err := Validate(tampered); err == nil {
		t.Fatal("accepted tampered Policy Evaluation Record")
	}
	result.SchemaVersion = "2.0"
	if _, err := MarshalCanonical(result); err == nil {
		t.Fatal("serialized unsupported Policy result")
	}
}

func normalized(t *testing.T, value Profile) Profile {
	t.Helper()
	result, err := NormalizeProfile(value)
	if err != nil {
		t.Fatal(err)
	}
	return result
}
func resign(t *testing.T, value Profile) Profile {
	t.Helper()
	value.Identity = ""
	value.Identity = profileIdentity(value)
	return value
}
func resignWithoutValidation(value Profile) Profile {
	value.Identity = ""
	value.Identity = profileIdentity(value)
	return value
}

func profile(id string, priority int, fallback Outcome, statements ...Statement) Profile {
	return Profile{ID: id, ContractVersion: ProfileVersion, Version: "1.0", Priority: priority, Extends: []string{}, Enabled: true,
		Scope: Selector{RuleIDs: []string{}, Outcomes: []rule.Outcome{}}, Statements: statements, DefaultOutcome: fallback, Metadata: map[string]string{}}
}

func statement(id string, priority int, outcome Outcome, sourceOutcomes ...rule.Outcome) Statement {
	return Statement{ID: id, Priority: priority, Selector: Selector{RuleIDs: []string{}, Outcomes: sourceOutcomes}, Outcome: outcome, Explanation: "canonical_test_policy", Metadata: map[string]string{}}
}

func ruleInput(t *testing.T) rule.Result {
	t.Helper()
	change := comparison.ChangeRecord{ID: "change", Type: comparison.Removed, Layer: "security", ObjectID: "object",
		Path: "/layers/security/resources/object/facts/value", Previous: &comparison.TypedValue{Type: "string", Value: "private"},
		ComparisonTimestamp: time.Date(2026, 8, 6, 12, 0, 0, 0, time.UTC), Metadata: map[string]string{"object_kind": "test", "fact_name": "value"}}
	drifted, err := drift.Classify([]comparison.ChangeRecord{change})
	if err != nil {
		t.Fatal(err)
	}
	healthResult, err := health.Evaluate(drifted)
	if err != nil {
		t.Fatal(err)
	}
	definition := rule.Definition{ID: "canonical.test", ContractVersion: rule.RuleVersion, Category: rule.StatusRule,
		Scope: rule.Scope{HealthIDs: []string{}, Categories: []drift.Category{}}, Enabled: true, InputRequirements: []rule.Field{rule.FieldStatus},
		Condition:   rule.Condition{Operator: rule.Exists, Field: rule.FieldStatus, Values: []rule.Value{}, Children: []rule.Condition{}},
		Description: "Canonical test rule", Metadata: map[string]string{}}
	result, err := rule.Evaluate([]rule.Definition{definition}, healthResult)
	if err != nil {
		t.Fatal(err)
	}
	return result
}
