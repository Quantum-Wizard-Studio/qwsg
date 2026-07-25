package report

import (
	"strings"
	"testing"
	"time"

	"quantumwizard.hu/qwsg/internal/comparison"
	"quantumwizard.hu/qwsg/internal/drift"
	"quantumwizard.hu/qwsg/internal/health"
	"quantumwizard.hu/qwsg/internal/inventory"
	"quantumwizard.hu/qwsg/internal/rule"
)

func TestDeterministicReportIdentityOrderingAndSerialization(t *testing.T) {
	input := ruleInput(t)
	first, err := Generate(input)
	if err != nil {
		t.Fatal(err)
	}
	second, err := Generate(input)
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
	if string(left) != string(right) || first.ID != second.ID {
		t.Fatal("report generation is not deterministic")
	}
	if first.Summary.Total != 2 || first.Sections[0].Outcome != rule.Matched ||
		first.Sections[1].Outcome != rule.NotMatched {
		t.Fatalf("unexpected taxonomy order or summary: %#v", first)
	}
}

func TestTraceabilityAndPrivacyPreservation(t *testing.T) {
	input := ruleInput(t)
	report, err := Generate(input)
	if err != nil {
		t.Fatal(err)
	}
	for _, section := range report.Sections {
		for _, item := range section.Items {
			found := false
			for _, source := range report.Sources {
				if source.ID == item.Source.ID {
					found = true
				}
			}
			if !found || item.Source.ContractName != rule.SchemaName ||
				item.Source.ContractVersion != rule.SchemaVersion {
				t.Fatalf("lost canonical source traceability: %#v", item)
			}
		}
	}
	data, err := MarshalCanonical(report)
	if err != nil {
		t.Fatal(err)
	}
	if strings.Contains(string(data), "private-previous") ||
		strings.Contains(string(data), "private-current") {
		t.Fatalf("upstream private values leaked: %s", data)
	}
}

func TestAllRuleOutcomesRemainDistinct(t *testing.T) {
	input := allOutcomeRuleInput(t)
	report, err := Generate(input)
	if err != nil {
		t.Fatal(err)
	}
	if report.Summary != (Summary{
		Total: 7, Matched: 1, NotMatched: 1, InsufficientEvidence: 1,
		UnsupportedRule: 1, InvalidRule: 1, EvaluationError: 1,
		DisabledRule: 1,
	}) {
		t.Fatalf("outcomes were conflated: %#v", report.Summary)
	}
	if len(report.Sections) != 7 {
		t.Fatalf("got %d outcome sections", len(report.Sections))
	}
}

func TestEmptyInputIsExplicitlyIncomplete(t *testing.T) {
	input, err := rule.Evaluate([]rule.Definition{}, healthInput(t))
	if err != nil {
		t.Fatal(err)
	}
	report, err := Generate(input)
	if err != nil {
		t.Fatal(err)
	}
	if report.Completeness != Incomplete || report.Summary.Total != 0 ||
		len(report.Sections) != 0 {
		t.Fatalf("empty evidence presented as complete: %#v", report)
	}
}

func TestInvalidAndUnsupportedContractsAreRejected(t *testing.T) {
	input := ruleInput(t)
	input.SchemaVersion = "2.0"
	if _, err := Generate(input); err == nil {
		t.Fatal("accepted unsupported source contract")
	}
	report, err := Generate(ruleInput(t))
	if err != nil {
		t.Fatal(err)
	}
	report.SchemaVersion = "2.0"
	if err := Validate(report); err == nil {
		t.Fatal("accepted unsupported report contract")
	}
	if _, err := MarshalCanonical(report); err == nil {
		t.Fatal("serialized unsupported report contract")
	}
}

func TestTamperingDuplicateTraceabilityAndBoundsAreRejected(t *testing.T) {
	report, err := Generate(ruleInput(t))
	if err != nil {
		t.Fatal(err)
	}
	tampered := report
	tampered.Sources = append([]SourceReference(nil), report.Sources...)
	tampered.Sources[0].ContractVersion = "2.0"
	if err := Validate(tampered); err == nil {
		t.Fatal("accepted tampered source")
	}
	duplicate := report
	duplicate.Sources = append(append([]SourceReference(nil), report.Sources...),
		report.Sources[len(report.Sources)-1])
	if err := Validate(duplicate); err == nil {
		t.Fatal("accepted duplicate source")
	}
	oversized := ruleInput(t)
	oversized.Records = make([]rule.EvaluationRecord, MaxItems+1)
	if _, err := Generate(oversized); err == nil {
		t.Fatal("accepted invalid oversized source")
	}
}

func TestSafeTextRenderingUsesOnlyCanonicalReport(t *testing.T) {
	report, err := Generate(ruleInput(t))
	if err != nil {
		t.Fatal(err)
	}
	rendered, err := RenderText(report)
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(rendered, report.ID) ||
		safe("safe\nforged") != `safe\u000Aforged` {
		t.Fatalf("unsafe control character rendering: %q", rendered)
	}
}

func TestRealInventoryToReportPipeline(t *testing.T) {
	from := snapshot("from", "old")
	to := snapshot("to", "new")
	compared, err := comparison.Compare(from, to, "from", "to")
	if err != nil {
		t.Fatal(err)
	}
	classified, err := drift.Classify(compared.Changes)
	if err != nil {
		t.Fatal(err)
	}
	conditionValue := "warning"
	healthResult, err := health.Evaluate(classified)
	if err != nil {
		t.Fatal(err)
	}
	evaluated, err := rule.Evaluate([]rule.Definition{{
		ID: "pipeline.warning", ContractVersion: rule.RuleVersion,
		Category: rule.StatusRule,
		Scope:    rule.Scope{HealthIDs: []string{}, Categories: []drift.Category{}},
		Enabled:  true, InputRequirements: []rule.Field{rule.FieldStatus},
		Condition: rule.Condition{
			Operator: rule.StatusMatches, Field: rule.FieldStatus,
			Value: rule.Value{String: &conditionValue},
		},
		Description: "Presentation pipeline fixture.", Metadata: map[string]string{},
	}}, healthResult)
	if err != nil {
		t.Fatal(err)
	}
	report, err := Generate(evaluated)
	if err != nil {
		t.Fatal(err)
	}
	found := false
	for _, source := range report.Sources {
		if source.ID == evaluated.Records[0].ID {
			found = true
		}
	}
	if report.Summary.Total == 0 || !found {
		t.Fatalf("pipeline source identity was lost: %#v", report)
	}
}

func ruleInput(t *testing.T) rule.Result {
	t.Helper()
	critical, healthy := "critical", "healthy"
	definitions := []rule.Definition{
		definition("a.match", critical),
		definition("b.no-match", healthy),
	}
	result, err := rule.Evaluate(definitions, healthInput(t))
	if err != nil {
		t.Fatal(err)
	}
	return result
}

func allOutcomeRuleInput(t *testing.T) rule.Result {
	t.Helper()
	critical, healthy := "critical", "healthy"
	definitions := []rule.Definition{
		definition("a.match", critical),
		definition("b.no-match", healthy),
		definition("c.insufficient", critical),
		definition("d.unsupported", critical),
		definition("e.invalid", critical),
		definition("f.error", critical),
		definition("g.disabled", critical),
	}
	definitions[2].Scope.HealthIDs = []string{"absent"}
	definitions[3].Category = rule.ExtensionRule
	definitions[4].Condition = rule.Condition{Operator: rule.All}
	definitions[6].Enabled = false
	valid, err := rule.Evaluate(definitions[:5], healthInput(t))
	if err != nil {
		t.Fatal(err)
	}
	invalidHealth := healthInput(t)
	invalidHealth.SchemaVersion = "2.0"
	failed, err := rule.Evaluate(definitions[5:6], invalidHealth)
	if err != nil {
		t.Fatal(err)
	}
	disabled, err := rule.Evaluate(definitions[6:], healthInput(t))
	if err != nil {
		t.Fatal(err)
	}
	valid.Records = append(valid.Records, failed.Records...)
	valid.Records = append(valid.Records, disabled.Records...)
	// Rule validation requires canonical order.
	sortRuleRecords(valid.Records)
	if err := rule.Validate(valid); err != nil {
		t.Fatal(err)
	}
	return valid
}

func sortRuleRecords(records []rule.EvaluationRecord) {
	for i := 1; i < len(records); i++ {
		for j := i; j > 0; j-- {
			left := records[j-1].RuleID + "\x00" + records[j-1].HealthRecordID
			right := records[j].RuleID + "\x00" + records[j].HealthRecordID
			if left <= right {
				break
			}
			records[j-1], records[j] = records[j], records[j-1]
		}
	}
}

func definition(id, status string) rule.Definition {
	return rule.Definition{
		ID: id, ContractVersion: rule.RuleVersion, Category: rule.StatusRule,
		Scope:   rule.Scope{HealthIDs: []string{}, Categories: []drift.Category{}},
		Enabled: true, InputRequirements: []rule.Field{rule.FieldStatus},
		Condition: rule.Condition{
			Operator: rule.StatusMatches, Field: rule.FieldStatus,
			Value: rule.Value{String: &status},
		},
		Description: "Canonical report fixture.", Metadata: map[string]string{},
	}
}

func healthInput(t *testing.T) health.Result {
	t.Helper()
	previous := &comparison.TypedValue{Type: "string", Value: "private-previous"}
	change := comparison.ChangeRecord{
		ID: "change", Type: comparison.Removed, Layer: "security",
		ObjectID: "object", Path: "/layers/security/resources/object/facts/value",
		Previous: previous, ComparisonTimestamp: time.Date(2026, 7, 24, 12, 0, 0, 0, time.UTC),
		Metadata: map[string]string{"object_kind": "test", "fact_name": "value"},
	}
	classified, err := drift.Classify([]comparison.ChangeRecord{change})
	if err != nil {
		t.Fatal(err)
	}
	result, err := health.Evaluate(classified)
	if err != nil {
		t.Fatal(err)
	}
	return result
}

func snapshot(id, value string) inventory.Snapshot {
	now := time.Date(2026, 7, 24, 8, 0, 0, 0, time.UTC)
	producer := inventory.Producer{ToolVersion: "test", ContractVersion: "1.0"}
	fact := inventory.CanonicalFact{
		Value: value, ValueType: "string", Quality: "observed", Sensitivity: "public",
		ObservedAt: now,
		Provenance: inventory.Provenance{
			SourceType: "fixture", SourceLabel: "fixture", ObservedAt: now,
		},
	}
	resource := inventory.Resource{
		ResourceID: "security:fixture", Kind: "fixture", LayerID: "security",
		LifecycleState: "observed", Facts: map[string]inventory.CanonicalFact{"state": fact},
		Relationships: []inventory.Relationship{}, Labels: map[string]string{},
		ObservedAt: now, CollectorID: "fixture", Metadata: map[string]string{},
	}
	layer := inventory.Layer{
		LayerID: "security", ContractVersion: inventory.ContractVersionForLayer,
		Status: inventory.Available, ObservedAt: now, CompletedAt: now,
		CollectorIDs: []string{"fixture"}, Resources: []inventory.Resource{resource},
		Issues: []inventory.InventoryError{}, Redactions: []string{},
		Metadata: map[string]string{},
	}
	category := inventory.Category{
		CategoryID: "security", ContractVersion: "1.0", Status: inventory.Available,
		ObservedAt: now, CompletedAt: now, FreshUntil: now.Add(time.Minute),
		CollectorID: "fixture", PrivilegeUsed: "ordinary-user",
		SourceSummary: []string{"fixture"}, Items: []inventory.Item{},
		Errors: []inventory.InventoryError{}, Redactions: []string{},
	}
	canonical := inventory.SystemInventory{
		SchemaName: inventory.CanonicalSchemaName, SchemaVersion: inventory.SchemaVersion,
		Profile: "canonical-system-inventory-v1", SnapshotID: id, RequestID: id,
		SubjectID: "subject", ObservedAt: now, CompletedAt: now,
		FreshUntil: now.Add(time.Minute), Status: inventory.Complete, Producer: producer,
		CollectorResults: []inventory.CollectorExecution{{
			CollectorName: "fixture", Version: "1.0", Capability: "security",
			SupportedPlatforms: []string{"linux"}, Timestamp: now,
			Status: inventory.Available, Warnings: []inventory.InventoryWarning{},
			Errors: []inventory.InventoryError{}, Metadata: map[string]string{},
		}},
		Layers: []inventory.Layer{layer}, Issues: []inventory.InventoryError{},
		Redactions: []string{}, Metadata: map[string]string{},
	}
	return inventory.Snapshot{
		SchemaVersion: inventory.SchemaVersion, SnapshotID: id, RequestID: id,
		InstanceID: "subject", ObservedAt: now, CompletedAt: now,
		FreshUntil: now.Add(time.Minute), Status: inventory.Complete,
		Categories: []inventory.Category{category}, Errors: []inventory.InventoryError{},
		Redactions: []string{}, Producer: producer, Canonical: canonical,
	}
}
