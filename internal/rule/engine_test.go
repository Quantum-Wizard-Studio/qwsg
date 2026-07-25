package rule

import (
	"encoding/json"
	"strings"
	"testing"
	"time"

	"quantumwizard.hu/qwsg/internal/comparison"
	"quantumwizard.hu/qwsg/internal/drift"
	"quantumwizard.hu/qwsg/internal/health"
)

func TestCanonicalOutcomesRemainDistinct(t *testing.T) {
	input := healthInput(t, change("one", "security", comparison.Removed))
	rules := []Definition{
		definition("a.match", StatusRule, leaf(StatusMatches, FieldStatus, text("critical"))),
		definition("b.no-match", StatusRule, leaf(StatusMatches, FieldStatus, text("healthy"))),
		definition("c.insufficient", StatusRule, leaf(Exists, FieldStatus, Value{})),
		definition("d.unsupported", ExtensionRule, leaf(Equal, FieldStatus, text("critical"))),
		definition("e.invalid", StatusRule, Condition{Operator: All}),
	}
	rules[2].Scope.HealthIDs = []string{"absent-health-id"}
	result, err := Evaluate(rules, input)
	if err != nil {
		t.Fatal(err)
	}
	want := map[string]Outcome{
		"a.match": Matched, "b.no-match": NotMatched,
		"c.insufficient": InsufficientEvidence,
		"d.unsupported":  UnsupportedRule,
		"e.invalid":      InvalidRule,
	}
	for _, record := range result.Records {
		if record.Outcome != want[record.RuleID] {
			t.Fatalf("%s: got %s, want %s", record.RuleID, record.Outcome, want[record.RuleID])
		}
	}
	invalidHealth := input
	invalidHealth.SchemaVersion = "2.0"
	failed, err := Evaluate([]Definition{rules[0]}, invalidHealth)
	if err != nil {
		t.Fatal(err)
	}
	if failed.Records[0].Outcome != EvaluationError ||
		failed.Records[0].Outcome == NotMatched {
		t.Fatalf("technical failure was conflated: %#v", failed.Records[0])
	}
}

func TestAllSupportedOperators(t *testing.T) {
	input := healthInput(t, change("one", "security", comparison.Removed))
	n9999, n10000, n10001 := int64(9999), int64(10000), int64(10001)
	conditions := map[string]Condition{
		"eq":       leaf(Equal, FieldStatus, text("critical")),
		"ne":       leaf(NotEqual, FieldStatus, text("healthy")),
		"gt":       leaf(GreaterThan, FieldConfidence, Value{Number: &n9999}),
		"gte":      leaf(GreaterThanOrEqual, FieldConfidence, Value{Number: &n10000}),
		"lt":       leaf(LessThan, FieldConfidence, Value{Number: &n10001}),
		"lte":      leaf(LessThanOrEqual, FieldConfidence, Value{Number: &n10000}),
		"in":       {Operator: In, Field: FieldStatus, Values: []Value{text("critical"), text("warning")}},
		"exists":   leaf(Exists, FieldReason, Value{}),
		"status":   leaf(StatusMatches, FieldStatus, text("critical")),
		"category": leaf(CategoryMatches, FieldCategory, text("security")),
		"and":      {Operator: All, Children: []Condition{leaf(Exists, FieldStatus, Value{}), leaf(Equal, FieldStatus, text("critical"))}},
		"or":       {Operator: Any, Children: []Condition{leaf(Equal, FieldStatus, text("healthy")), leaf(Equal, FieldStatus, text("critical"))}},
		"not":      {Operator: Not, Children: []Condition{leaf(Equal, FieldStatus, text("healthy"))}},
	}
	rules := make([]Definition, 0, len(conditions))
	for name, condition := range conditions {
		rules = append(rules, definition("operator."+name, CompositeRule, condition))
	}
	result, err := Evaluate(rules, input)
	if err != nil {
		t.Fatal(err)
	}
	if len(result.Records) != len(conditions) {
		t.Fatalf("got %d evaluations", len(result.Records))
	}
	for _, record := range result.Records {
		if record.Outcome != Matched {
			t.Fatalf("%s did not match: %#v", record.RuleID, record)
		}
	}
}

func TestDeterministicOrderingIdentifiersAndJSON(t *testing.T) {
	input := healthInput(t,
		change("z", "services", comparison.Removed),
		change("a", "hardware", comparison.Unchanged),
	)
	firstRules := []Definition{
		definition("z.rule", StatusRule, leaf(Exists, FieldStatus, Value{})),
		definition("a.rule", StatusRule, leaf(Exists, FieldStatus, Value{})),
	}
	secondRules := []Definition{firstRules[1], firstRules[0]}
	first, err := Evaluate(firstRules, input)
	if err != nil {
		t.Fatal(err)
	}
	second, err := Evaluate(secondRules, input)
	if err != nil {
		t.Fatal(err)
	}
	left, err := MarshalCanonical(first)
	if err != nil {
		t.Fatal(err)
	}
	right, err := MarshalCanonical(second)
	if err != nil {
		t.Fatal(err)
	}
	if string(left) != string(right) {
		t.Fatalf("non-deterministic Rule output:\n%s\n%s", left, right)
	}
	if first.Records[0].RuleID != "a.rule" ||
		first.Records[len(first.Records)-1].RuleID != "z.rule" {
		t.Fatalf("unstable Rule order: %#v", first.Records)
	}
}

func TestRuleContractVersionUnsupportedOperatorInvalidAndDisabled(t *testing.T) {
	input := healthInput(t, change("one", "hardware", comparison.Unchanged))
	unsupportedVersion := definition("a.version", StatusRule, leaf(Exists, FieldStatus, Value{}))
	unsupportedVersion.ContractVersion = "2.0"
	unsupportedOperator := definition("b.operator", StatusRule, Condition{Operator: "future"})
	invalid := definition("c.invalid", StatusRule, leaf(GreaterThan, FieldStatus, text("healthy")))
	disabled := definition("d.disabled", StatusRule, leaf(Exists, FieldStatus, Value{}))
	disabled.Enabled = false
	result, err := Evaluate([]Definition{unsupportedVersion, unsupportedOperator, invalid, disabled}, input)
	if err != nil {
		t.Fatal(err)
	}
	want := []Outcome{UnsupportedRule, UnsupportedRule, InvalidRule, DisabledRule}
	for i, record := range result.Records {
		if record.Outcome != want[i] {
			t.Fatalf("%d: got %s, want %s", i, record.Outcome, want[i])
		}
	}
}

func TestCompositionAndInputBounds(t *testing.T) {
	input := healthInput(t, change("one", "hardware", comparison.Unchanged))
	tooDeep := leaf(Equal, FieldStatus, text("healthy"))
	for i := 0; i < MaxDepth; i++ {
		tooDeep = Condition{Operator: Not, Children: []Condition{tooDeep}}
	}
	rule := definition("deep.rule", CompositeRule, tooDeep)
	result, err := Evaluate([]Definition{rule}, input)
	if err != nil {
		t.Fatal(err)
	}
	if result.Records[0].Outcome != InvalidRule ||
		result.Records[0].Explanation != "condition_bounds_exceeded" {
		t.Fatalf("unbounded composition accepted: %#v", result.Records[0])
	}
	many := make([]Definition, MaxRules+1)
	for i := range many {
		many[i] = definition("rule."+string(rune('a'+i%26)), StatusRule, leaf(Exists, FieldStatus, Value{}))
	}
	if _, err := Evaluate(many, input); err == nil {
		t.Fatal("accepted rule set above the canonical limit")
	}
}

func TestScopeRequirementsAndPrivacy(t *testing.T) {
	input := healthInput(t, change("private", "configuration", comparison.Modified))
	before, _ := json.Marshal(input)
	rule := definition("privacy.rule", EvidenceRule, leaf(Exists, FieldReason, Value{}))
	rule.Scope.Categories = []drift.Category{drift.ConfigurationDrift}
	rule.InputRequirements = []Field{FieldReason, FieldStatus}
	result, err := Evaluate([]Definition{rule}, input)
	if err != nil {
		t.Fatal(err)
	}
	after, _ := json.Marshal(input)
	output, _ := json.Marshal(result)
	if string(before) != string(after) {
		t.Fatal("Rule mutated canonical Health input")
	}
	if strings.Contains(string(output), "private-previous") ||
		strings.Contains(string(output), "private-current") {
		t.Fatalf("Rule disclosed upstream values: %s", output)
	}
	if len(result.Records[0].EvidenceReferences) != 1 ||
		result.Records[0].EvidenceReferences[0] != input.Records[0].ID {
		t.Fatalf("canonical Health evidence was not preserved: %#v", result.Records[0])
	}
}

func TestInvalidDefinitionShapesAndDuplicates(t *testing.T) {
	input := healthInput(t, change("one", "hardware", comparison.Unchanged))
	badValue := text("healthy")
	number := int64(1)
	badValue.Number = &number
	cases := []Definition{
		definition("bad.value", StatusRule, leaf(Equal, FieldStatus, badValue)),
		definition("bad.membership", StatusRule, Condition{Operator: In, Field: FieldStatus, Values: []Value{}}),
		definition("bad.not", CompositeRule, Condition{Operator: Not, Children: []Condition{}}),
	}
	for _, rule := range cases {
		result, err := Evaluate([]Definition{rule}, input)
		if err != nil {
			t.Fatal(err)
		}
		if result.Records[0].Outcome != InvalidRule {
			t.Fatalf("accepted invalid definition: %#v", rule)
		}
	}
	duplicate := definition("duplicate.rule", StatusRule, leaf(Exists, FieldStatus, Value{}))
	if _, err := Evaluate([]Definition{duplicate, duplicate}, input); err == nil {
		t.Fatal("accepted duplicate Rule IDs")
	}
}

func TestHealthToRulePipeline(t *testing.T) {
	driftResult, err := drift.Classify([]comparison.ChangeRecord{
		change("pipeline", "security", comparison.Removed),
	})
	if err != nil {
		t.Fatal(err)
	}
	healthResult, err := health.Evaluate(driftResult)
	if err != nil {
		t.Fatal(err)
	}
	rule := definition(
		"security.removal.critical", StatusRule,
		Condition{Operator: All, Children: []Condition{
			leaf(CategoryMatches, FieldCategory, text("security")),
			leaf(StatusMatches, FieldStatus, text("critical")),
		}},
	)
	result, err := Evaluate([]Definition{rule}, healthResult)
	if err != nil {
		t.Fatal(err)
	}
	if len(result.Records) != 1 || result.Records[0].Outcome != Matched ||
		result.Records[0].HealthRecordID != healthResult.Records[0].ID {
		t.Fatalf("Health-to-Rule pipeline failed: %#v", result)
	}
}

func TestValidateRejectsTampering(t *testing.T) {
	input := healthInput(t, change("one", "hardware", comparison.Unchanged))
	result, err := Evaluate([]Definition{
		definition("valid.rule", StatusRule, leaf(Exists, FieldStatus, Value{})),
	}, input)
	if err != nil {
		t.Fatal(err)
	}
	tampered := result
	tampered.Records = append([]EvaluationRecord(nil), result.Records...)
	tampered.Records[0].Outcome = EvaluationError
	if err := Validate(tampered); err == nil {
		t.Fatal("accepted tampered evaluation")
	}
	tampered = result
	tampered.SchemaVersion = "2.0"
	if _, err := MarshalCanonical(tampered); err == nil {
		t.Fatal("serialized unsupported evaluation contract")
	}
}

func definition(id string, category Category, condition Condition) Definition {
	return Definition{
		ID: id, ContractVersion: RuleVersion, Category: category,
		Scope:   Scope{HealthIDs: []string{}, Categories: []drift.Category{}},
		Enabled: true, InputRequirements: []Field{},
		Condition: condition, Description: "Canonical test rule.",
		Metadata: map[string]string{},
	}
}

func leaf(operator Operator, field Field, value Value) Condition {
	return Condition{Operator: operator, Field: field, Value: value}
}

func text(value string) Value { return Value{String: &value} }

func healthInput(t *testing.T, changes ...comparison.ChangeRecord) health.Result {
	t.Helper()
	driftResult, err := drift.Classify(changes)
	if err != nil {
		t.Fatal(err)
	}
	result, err := health.Evaluate(driftResult)
	if err != nil {
		t.Fatal(err)
	}
	return result
}

func change(id, layer string, kind comparison.ChangeType) comparison.ChangeRecord {
	previous := &comparison.TypedValue{Type: "string", Value: "private-previous"}
	current := &comparison.TypedValue{Type: "string", Value: "private-current"}
	switch kind {
	case comparison.Added:
		previous = nil
	case comparison.Removed:
		current = nil
	case comparison.Unchanged:
		current = &comparison.TypedValue{Type: "string", Value: "private-previous"}
	}
	return comparison.ChangeRecord{
		ID: id, Type: kind, Layer: layer, ObjectID: "object-" + id,
		Path:     "/layers/" + layer + "/resources/object-" + id + "/facts/value",
		Previous: previous, Current: current,
		ComparisonTimestamp: time.Date(2026, 7, 24, 12, 0, 0, 0, time.UTC),
		Metadata:            map[string]string{"object_kind": "test", "fact_name": "value"},
	}
}
