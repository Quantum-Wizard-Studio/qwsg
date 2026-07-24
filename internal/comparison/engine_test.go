package comparison

import (
	"encoding/json"
	"reflect"
	"testing"
	"time"

	"quantumwizard.hu/qwsg/internal/inventory"
)

func TestIdenticalSnapshotsAreDeterministicAndHaveNoChanges(t *testing.T) {
	from := fixture("from", "subject", "old")
	to := fixture("to", "subject", "old")
	first, err := Compare(from, to, "from.json", "to.json")
	if err != nil {
		t.Fatal(err)
	}
	second, err := Compare(from, to, "from.json", "to.json")
	if err != nil {
		t.Fatal(err)
	}
	firstJSON, _ := json.Marshal(first)
	secondJSON, _ := json.Marshal(second)
	if string(firstJSON) != string(secondJSON) {
		t.Fatal("repeat comparison is not byte-identical")
	}
	if first.Counts.Added != 0 || first.Counts.Removed != 0 || first.Counts.Modified != 0 || first.Counts.Unchanged == 0 {
		t.Fatalf("unexpected counts: %#v", first.Counts)
	}
	if !first.ComparisonTimestamp.Equal(to.CompletedAt) {
		t.Fatal("comparison timestamp is not derived from to snapshot")
	}
}

func TestModifiedAddedRemovedAndReverse(t *testing.T) {
	from := fixture("from", "subject", "old")
	to := fixture("to", "subject", "new")
	to.Canonical.Layers[0].Resources[0].Facts["added"] = canonicalFact(true)
	delete(to.Canonical.Layers[0].Resources[0].Facts, "removed")
	forward, err := Compare(from, to, "from", "to")
	if err != nil {
		t.Fatal(err)
	}
	if forward.Counts.Added != 1 || forward.Counts.Removed != 1 || forward.Counts.Modified != 1 {
		t.Fatalf("unexpected counts: %#v", forward.Counts)
	}
	reverse, err := Compare(to, from, "to", "from")
	if err != nil {
		t.Fatal(err)
	}
	if reverse.Counts.Added != forward.Counts.Removed || reverse.Counts.Removed != forward.Counts.Added ||
		reverse.Counts.Modified != forward.Counts.Modified {
		t.Fatalf("reverse counts differ: %#v %#v", forward.Counts, reverse.Counts)
	}
}

func TestOrderingAndVolatileMetadataDoNotChangeSemantics(t *testing.T) {
	from := fixture("from", "subject", "same")
	to := fixture("to", "subject", "same")
	to.Canonical.Layers[0].Resources[0].ObservedAt = to.Canonical.Layers[0].Resources[0].ObservedAt.Add(time.Hour)
	fact := to.Canonical.Layers[0].Resources[0].Facts["stable"]
	fact.ObservedAt = time.Now().UTC()
	fact.Provenance.ObservedAt = time.Now().UTC()
	to.Canonical.Layers[0].Resources[0].Facts["stable"] = fact
	to.Canonical.Layers[0].Resources[0].Metadata["volatile"] = "different"
	result, err := Compare(from, to, "from", "to")
	if err != nil {
		t.Fatal(err)
	}
	if result.Counts.Added+result.Counts.Removed+result.Counts.Modified != 0 {
		t.Fatalf("volatile metadata produced changes: %#v", result)
	}
}

func TestRejectsInvalidSourcesAndSubjects(t *testing.T) {
	valid := fixture("valid", "subject", "value")
	tests := []struct {
		name string
		from inventory.Snapshot
		to   inventory.Snapshot
	}{
		{"subject", valid, fixture("other", "different", "value")},
		{"missing canonical", withoutCanonical(valid), valid},
		{"invalid inventory", withSchema(valid, "2.0"), valid},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if _, err := Compare(tc.from, tc.to, "from", "to"); err == nil {
				t.Fatal("invalid comparison accepted")
			}
		})
	}
}

func TestEscapedPathsAndInputImmutability(t *testing.T) {
	from := fixture("from", "subject", "old")
	to := fixture("to", "subject", "new")
	resource := to.Canonical.Layers[0].Resources[0]
	delete(to.Canonical.Layers[0].Resources[0].Facts, "stable")
	resource.Facts = map[string]inventory.CanonicalFact{"a/b~c": canonicalFact("new")}
	to.Canonical.Layers[0].Resources[0] = resource
	before, _ := json.Marshal(to)
	result, err := Compare(from, to, "from", "to")
	if err != nil {
		t.Fatal(err)
	}
	after, _ := json.Marshal(to)
	if !reflect.DeepEqual(before, after) {
		t.Fatal("comparison mutated input")
	}
	found := false
	for _, change := range result.Changes {
		if change.Path == "/layers/host/resources/host:fixture/facts/a~1b~0c" {
			found = true
		}
	}
	if !found {
		t.Fatalf("escaped path missing: %#v", result.Changes)
	}
}

func TestCanonicalJSONValueSemanticsAndTypes(t *testing.T) {
	from := fixture("from", "subject", "same")
	to := fixture("to", "subject", "same")
	fromFacts := from.Canonical.Layers[0].Resources[0].Facts
	toFacts := to.Canonical.Layers[0].Resources[0].Facts
	values := map[string]any{
		"boolean": true,
		"empty":   "",
		"null":    nil,
		"number":  1.5,
		"array":   []any{"α", false, float64(2)},
		"object":  map[string]any{"z": "終", "a": float64(1)},
	}
	for name, value := range values {
		fromFacts[name] = canonicalFact(value)
		toFacts[name] = canonicalFact(value)
	}
	fromFacts["integer"] = canonicalFact(int64(4))
	toFacts["integer"] = canonicalFact(float64(4))
	fromFacts["integer"] = withValueType(fromFacts["integer"], "integer")
	toFacts["integer"] = withValueType(toFacts["integer"], "integer")
	result, err := Compare(from, to, "from", "to")
	if err != nil {
		t.Fatal(err)
	}
	if result.Counts.Added+result.Counts.Removed+result.Counts.Modified != 0 {
		t.Fatalf("canonical JSON-equivalent values differ: %#v", result.Counts)
	}
}

func fixture(snapshotID, subject, value string) inventory.Snapshot {
	now := time.Date(2026, 7, 24, 8, 0, 0, 0, time.UTC)
	producer := inventory.Producer{ToolVersion: "test", ContractVersion: "1.0"}
	resource := inventory.Resource{
		ResourceID: "host:fixture", Kind: "host", LayerID: "host", LifecycleState: "observed",
		Facts: map[string]inventory.CanonicalFact{
			"stable":  canonicalFact(value),
			"removed": canonicalFact(int64(4)),
		},
		Relationships: []inventory.Relationship{}, Labels: map[string]string{},
		ObservedAt: now, CollectorID: "fixture", Metadata: map[string]string{},
	}
	layer := inventory.Layer{
		LayerID: "host", ContractVersion: inventory.ContractVersionForLayer, Status: inventory.Available,
		ObservedAt: now, CompletedAt: now, CollectorIDs: []string{"fixture"},
		Resources: []inventory.Resource{resource}, Issues: []inventory.InventoryError{},
		Redactions: []string{}, Metadata: map[string]string{},
	}
	category := inventory.Category{
		CategoryID: "host", ContractVersion: "1.0", Status: inventory.Available,
		ObservedAt: now, CompletedAt: now, FreshUntil: now.Add(time.Minute),
		CollectorID: "fixture", PrivilegeUsed: "ordinary-user", SourceSummary: []string{"fixture"},
		Items: []inventory.Item{}, Errors: []inventory.InventoryError{}, Redactions: []string{},
	}
	canonical := inventory.SystemInventory{
		SchemaName: inventory.CanonicalSchemaName, SchemaVersion: inventory.SchemaVersion,
		Profile: "canonical-system-inventory-v1", SnapshotID: snapshotID, RequestID: snapshotID,
		SubjectID: subject, ObservedAt: now, CompletedAt: now, FreshUntil: now.Add(time.Minute),
		Status: inventory.Complete, Producer: producer,
		CollectorResults: []inventory.CollectorExecution{{
			CollectorName: "fixture", Version: "1.0", Capability: "host",
			SupportedPlatforms: []string{"linux"}, Timestamp: now, Status: inventory.Available,
			Warnings: []inventory.InventoryWarning{}, Errors: []inventory.InventoryError{},
			Metadata: map[string]string{},
		}},
		Layers: []inventory.Layer{layer}, Issues: []inventory.InventoryError{},
		Redactions: []string{}, Metadata: map[string]string{},
	}
	return inventory.Snapshot{
		SchemaVersion: inventory.SchemaVersion, SnapshotID: snapshotID, RequestID: snapshotID,
		InstanceID: subject, ObservedAt: now, CompletedAt: now, FreshUntil: now.Add(time.Minute),
		Status: inventory.Complete, Categories: []inventory.Category{category},
		Errors: []inventory.InventoryError{}, Redactions: []string{}, Producer: producer, Canonical: canonical,
	}
}

func canonicalFact(value any) inventory.CanonicalFact {
	return inventory.CanonicalFact{
		Value: value, ValueType: typeName(value), Quality: "observed", Sensitivity: "public",
		ObservedAt: time.Date(2026, 7, 24, 8, 0, 0, 0, time.UTC),
		Provenance: inventory.Provenance{SourceType: "fixture", SourceLabel: "fixture", ObservedAt: time.Date(2026, 7, 24, 8, 0, 0, 0, time.UTC)},
	}
}

func typeName(value any) string {
	switch value.(type) {
	case string:
		return "string"
	case bool:
		return "boolean"
	default:
		return "integer"
	}
}

func withValueType(fact inventory.CanonicalFact, valueType string) inventory.CanonicalFact {
	fact.ValueType = valueType
	return fact
}

func withoutCanonical(snapshot inventory.Snapshot) inventory.Snapshot {
	snapshot.Canonical = inventory.SystemInventory{}
	return snapshot
}

func withSchema(snapshot inventory.Snapshot, schema string) inventory.Snapshot {
	snapshot.SchemaVersion = schema
	return snapshot
}
