package health

import (
	"encoding/json"
	"strings"
	"testing"
	"time"

	"quantumwizard.hu/qwsg/internal/comparison"
	"quantumwizard.hu/qwsg/internal/drift"
	"quantumwizard.hu/qwsg/internal/inventory"
)

func TestEvaluateTaxonomyAndAggregation(t *testing.T) {
	input := driftResult(t,
		change("healthy", "hardware", comparison.Unchanged),
		change("info", "software", comparison.Added),
		change("advisory", "configuration", comparison.Modified),
		change("warning", "services", comparison.Removed),
		change("critical", "security", comparison.Removed),
		change("unsupported", "future_layer", comparison.Modified),
	)
	result, err := Evaluate(input)
	if err != nil {
		t.Fatal(err)
	}
	want := map[string]struct {
		status Status
		state  EvidenceState
		reason string
	}{
		"healthy":     {Healthy, EvidenceSufficient, "canonical_state_unchanged"},
		"info":        {Informational, EvidenceSufficient, "canonical_presence_added"},
		"advisory":    {Advisory, EvidenceSufficient, "canonical_value_modified"},
		"warning":     {Warning, EvidenceSufficient, "canonical_presence_removed"},
		"critical":    {Critical, EvidenceSufficient, "canonical_security_presence_removed"},
		"unsupported": {Unsupported, EvidenceUnsupported, "unsupported_drift_category"},
	}
	if result.OverallStatus != Critical || result.EvidenceState != EvidenceUnsupported {
		t.Fatalf("unexpected summary: %s/%s", result.OverallStatus, result.EvidenceState)
	}
	for _, record := range result.Records {
		expected, ok := want[record.ChangeID]
		if !ok {
			t.Fatalf("unexpected record: %#v", record)
		}
		if record.Status != expected.status || record.EvidenceState != expected.state ||
			record.Reason != expected.reason {
			t.Fatalf("%s: got %s/%s/%s", record.ChangeID, record.Status, record.EvidenceState, record.Reason)
		}
	}
}

func TestEmptyEvidenceIsUnknownAndInsufficient(t *testing.T) {
	result, err := Evaluate(driftResult(t))
	if err != nil {
		t.Fatal(err)
	}
	if result.OverallStatus != Unknown || result.EvidenceState != EvidenceInsufficient ||
		len(result.Records) != 0 {
		t.Fatalf("unexpected empty evaluation: %#v", result)
	}
}

func TestEveryDriftTaxonomyCategoryIsPreserved(t *testing.T) {
	layers := []string{
		"host", "software", "hardware", "platform", "filesystem", "storage",
		"network", "services", "configuration", "security", "capability",
		"environment", "future_layer",
	}
	changes := make([]comparison.ChangeRecord, 0, len(layers))
	for _, layer := range layers {
		changes = append(changes, change(layer, layer, comparison.Unchanged))
	}
	input := driftResult(t, changes...)
	result, err := Evaluate(input)
	if err != nil {
		t.Fatal(err)
	}
	if len(result.Records) != len(input.Records) {
		t.Fatalf("got %d Health Records for %d Drift Records", len(result.Records), len(input.Records))
	}
	byDriftID := map[string]Record{}
	for _, record := range result.Records {
		byDriftID[record.DriftID] = record
	}
	for _, evidence := range input.Records {
		record, ok := byDriftID[evidence.ID]
		if !ok || record.Category != evidence.Category {
			t.Fatalf("category was not preserved for %s: %#v", evidence.ID, record)
		}
		if evidence.Category == drift.ExtensionDrift && record.Status != Unsupported {
			t.Fatalf("extension category is not unsupported: %#v", record)
		}
	}
}

func TestDeterministicOrderingIDsAndJSON(t *testing.T) {
	a := driftResult(t,
		change("z", "services", comparison.Removed),
		change("a", "hardware", comparison.Unchanged),
	)
	b := driftResult(t,
		change("a", "hardware", comparison.Unchanged),
		change("z", "services", comparison.Removed),
	)
	first, err := Evaluate(a)
	if err != nil {
		t.Fatal(err)
	}
	second, err := Evaluate(b)
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
		t.Fatalf("non-deterministic output:\n%s\n%s", left, right)
	}
	if first.Records[0].ChangeID != "a" || first.Records[1].ChangeID != "z" {
		t.Fatalf("records are not canonically ordered: %#v", first.Records)
	}
}

func TestDoesNotMutateOrDiscloseDriftInput(t *testing.T) {
	input := driftResult(t, change("private", "configuration", comparison.Modified))
	before, err := json.Marshal(input)
	if err != nil {
		t.Fatal(err)
	}
	result, err := Evaluate(input)
	if err != nil {
		t.Fatal(err)
	}
	after, err := json.Marshal(input)
	if err != nil {
		t.Fatal(err)
	}
	if string(before) != string(after) {
		t.Fatal("Health mutated canonical Drift evidence")
	}
	output, err := json.Marshal(result)
	if err != nil {
		t.Fatal(err)
	}
	if strings.Contains(string(output), "private-previous") ||
		strings.Contains(string(output), "private-current") {
		t.Fatalf("Health disclosed source values: %s", output)
	}
}

func TestRejectsInvalidAndUnsupportedContracts(t *testing.T) {
	valid := driftResult(t, change("one", "hardware", comparison.Unchanged))
	cases := []drift.Result{valid, valid, valid}
	cases[0].SchemaVersion = "2.0"
	cases[1].Records[0].ID = "tampered"
	cases[2].Records = append(cases[2].Records, cases[2].Records[0])
	for i, input := range cases {
		if _, err := Evaluate(input); err == nil {
			t.Fatalf("case %d accepted invalid Drift evidence", i)
		}
	}
}

func TestValidateRejectsTamperingAndUnknownRecord(t *testing.T) {
	result, err := Evaluate(driftResult(t, change("one", "hardware", comparison.Unchanged)))
	if err != nil {
		t.Fatal(err)
	}
	tampered := result
	tampered.Records = append([]Record(nil), result.Records...)
	tampered.Records[0].Status = Unknown
	if err := Validate(tampered); err == nil {
		t.Fatal("accepted a record-level unknown without insufficient evidence")
	}
	tampered = result
	tampered.OverallStatus = Critical
	if err := Validate(tampered); err == nil {
		t.Fatal("accepted invalid aggregation")
	}
	tampered = result
	tampered.SchemaVersion = "2.0"
	if _, err := MarshalCanonical(tampered); err == nil {
		t.Fatal("serialized unsupported Health contract")
	}
}

func TestRealComparisonToDriftToHealthPipeline(t *testing.T) {
	previous := comparisonSnapshot("previous", "old")
	current := comparisonSnapshot("current", "new")
	compared, err := comparison.Compare(previous, current, "previous.json", "current.json")
	if err != nil {
		t.Fatal(err)
	}
	drifted, err := drift.Classify(compared.Changes)
	if err != nil {
		t.Fatal(err)
	}
	evaluated, err := Evaluate(drifted)
	if err != nil {
		t.Fatal(err)
	}
	if len(evaluated.Records) != len(drifted.Records) || len(evaluated.Records) == 0 {
		t.Fatalf("pipeline cardinality mismatch: %d/%d", len(evaluated.Records), len(drifted.Records))
	}
	for i, record := range evaluated.Records {
		if record.DriftID != drifted.Records[i].ID || record.ChangeID != drifted.Records[i].ChangeID {
			t.Fatal("Health did not preserve canonical evidence references")
		}
	}
}

func driftResult(t *testing.T, changes ...comparison.ChangeRecord) drift.Result {
	t.Helper()
	result, err := drift.Classify(changes)
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

func comparisonSnapshot(id, value string) inventory.Snapshot {
	now := time.Date(2026, 7, 24, 12, 0, 0, 0, time.UTC)
	producer := inventory.Producer{ToolVersion: "test", ContractVersion: "1.0"}
	fact := inventory.CanonicalFact{
		Value: value, ValueType: "string", Quality: "observed", Sensitivity: "public",
		ObservedAt: now, Provenance: inventory.Provenance{
			SourceType: "fixture", SourceLabel: "fixture", ObservedAt: now,
		},
	}
	resource := inventory.Resource{
		ResourceID: "host:fixture", Kind: "host", LayerID: "host",
		LifecycleState: "observed", Facts: map[string]inventory.CanonicalFact{"name": fact},
		Relationships: []inventory.Relationship{}, Labels: map[string]string{},
		ObservedAt: now, CollectorID: "host", Metadata: map[string]string{},
	}
	layer := inventory.Layer{
		LayerID: "host", ContractVersion: inventory.ContractVersionForLayer,
		Status: inventory.Available, ObservedAt: now, CompletedAt: now,
		CollectorIDs: []string{"host"}, Resources: []inventory.Resource{resource},
		Issues: []inventory.InventoryError{}, Redactions: []string{},
		Metadata: map[string]string{},
	}
	category := inventory.Category{
		CategoryID: "host", ContractVersion: "1.0", Status: inventory.Available,
		ObservedAt: now, CompletedAt: now, FreshUntil: now.Add(time.Minute),
		CollectorID: "host", SourceSummary: []string{"fixture"},
		Items: []inventory.Item{}, Errors: []inventory.InventoryError{},
		Redactions: []string{},
	}
	canonical := inventory.SystemInventory{
		SchemaName: inventory.CanonicalSchemaName, SchemaVersion: inventory.SchemaVersion,
		Profile: "canonical-system-inventory-v1", SnapshotID: id, RequestID: id,
		SubjectID: "subject", ObservedAt: now, CompletedAt: now,
		FreshUntil: now.Add(time.Minute), Status: inventory.Complete, Producer: producer,
		CollectorResults: []inventory.CollectorExecution{{
			CollectorName: "host", Version: "1.0", Capability: "host",
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
