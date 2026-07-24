package drift

import (
	"encoding/json"
	"reflect"
	"testing"
	"time"

	"quantumwizard.hu/qwsg/internal/comparison"
	"quantumwizard.hu/qwsg/internal/inventory"
)

func TestTaxonomyAndClassificationCoverage(t *testing.T) {
	tests := []struct {
		layer, object, kind, fact string
		want                      Category
	}{
		{"host", "host:one", "fact", "name", IdentityDrift},
		{"accounts", "account:one", "fact", "state", IdentityDrift},
		{"software", "package:one", "fact", "version", SoftwareDrift},
		{"hardware", "cpu:one", "fact", "cores", HardwareDrift},
		{"operating_system", "os:one", "fact", "version", PlatformDrift},
		{"storage", "filesystem:one", "fact", "capacity", FilesystemDrift},
		{"storage", "volume:one", "fact", "capacity", StorageDrift},
		{"network", "interface:one", "fact", "mtu", NetworkDrift},
		{"services", "service:one", "fact", "state", ServiceDrift},
		{"configuration", "config:one", "fact", "value", ConfigurationDrift},
		{"security", "control:one", "fact", "state", SecurityDrift},
		{"capabilities", "capability:one", "fact", "state", CapabilityDrift},
		{"environment", "environment:one", "fact", "value", EnvironmentDrift},
		{"future_layer", "future:one", "fact", "value", ExtensionDrift},
	}
	for i, tc := range tests {
		change := changeRecord(tc.layer, tc.object, comparison.Modified)
		change.Metadata = map[string]string{"object_kind": tc.kind, "fact_name": tc.fact}
		change.ID = string(rune('a' + i))
		result, err := Classify([]comparison.ChangeRecord{change})
		if err != nil {
			t.Fatalf("%s: %v", tc.layer, err)
		}
		if got := result.Records[0].Category; got != tc.want {
			t.Errorf("%s: got %q want %q", tc.layer, got, tc.want)
		}
	}
}

func TestOneRecordPerChangeStableOrderAndJSON(t *testing.T) {
	changes := []comparison.ChangeRecord{
		changeRecord("network", "interface:b", comparison.Added),
		changeRecord("hardware", "cpu:a", comparison.Removed),
		changeRecord("services", "service:c", comparison.Unchanged),
	}
	changes[0].ID, changes[1].ID, changes[2].ID = "change-b", "change-a", "change-c"
	before, _ := json.Marshal(changes)
	first, err := Classify(changes)
	if err != nil {
		t.Fatal(err)
	}
	second, err := Classify(changes)
	if err != nil {
		t.Fatal(err)
	}
	left, _ := json.Marshal(first)
	right, _ := json.Marshal(second)
	if string(left) != string(right) {
		t.Fatal("repeat classification is not byte-identical")
	}
	if len(first.Records) != len(changes) {
		t.Fatalf("got %d records for %d changes", len(first.Records), len(changes))
	}
	if first.Records[0].Scope.Layer != "hardware" ||
		first.Records[1].Scope.Layer != "network" ||
		first.Records[2].Scope.Layer != "services" {
		t.Fatalf("unstable ordering: %#v", first.Records)
	}
	if first.Records[0].Classification != PresenceRemoved ||
		first.Records[1].Classification != PresenceAdded ||
		first.Records[2].Classification != StateUnchanged {
		t.Fatal("change classifications differ")
	}
	after, _ := json.Marshal(changes)
	if !reflect.DeepEqual(before, after) {
		t.Fatal("classification mutated input")
	}
}

func TestRejectsInvalidInputAndOutputContracts(t *testing.T) {
	valid := changeRecord("hardware", "cpu:a", comparison.Modified)
	valid.ID = "change"
	tests := []comparison.ChangeRecord{
		func() comparison.ChangeRecord { c := valid; c.ID = ""; return c }(),
		func() comparison.ChangeRecord { c := valid; c.Path = "not-canonical"; return c }(),
		func() comparison.ChangeRecord { c := valid; c.Type = "future"; return c }(),
		func() comparison.ChangeRecord { c := valid; c.Current = c.Previous; return c }(),
	}
	for i, change := range tests {
		if _, err := Classify([]comparison.ChangeRecord{change}); err == nil {
			t.Errorf("invalid input %d accepted", i)
		}
	}
	if _, err := Classify([]comparison.ChangeRecord{valid, valid}); err == nil {
		t.Fatal("duplicate change accepted")
	}

	result, err := Classify([]comparison.ChangeRecord{valid})
	if err != nil {
		t.Fatal(err)
	}
	result.SchemaVersion = "2.0"
	if Validate(result) == nil {
		t.Fatal("unsupported output contract accepted")
	}
	result, _ = Classify([]comparison.ChangeRecord{valid})
	result.Records[0].Metadata["private"] = "not-bounded"
	if Validate(result) == nil {
		t.Fatal("unbounded output metadata accepted")
	}
}

func TestMetadataIsBoundedAndDoesNotCopyValues(t *testing.T) {
	change := changeRecord("configuration", "config:a", comparison.Modified)
	change.ID = "change"
	change.Metadata["private_note"] = "must-not-propagate"
	result, err := Classify([]comparison.ChangeRecord{change})
	if err != nil {
		t.Fatal(err)
	}
	encoded, _ := json.Marshal(result)
	if string(encoded) == "" || contains(string(encoded), "must-not-propagate") ||
		contains(string(encoded), "previous-secret") || contains(string(encoded), "current-secret") {
		t.Fatalf("source values or metadata leaked: %s", encoded)
	}
}

func TestCompareToDriftPipeline(t *testing.T) {
	from := snapshot("from", "old")
	to := snapshot("to", "new")
	compared, err := comparison.Compare(from, to, "from.json", "to.json")
	if err != nil {
		t.Fatal(err)
	}
	result, err := Classify(compared.Changes)
	if err != nil {
		t.Fatal(err)
	}
	if len(result.Records) != len(compared.Changes) {
		t.Fatalf("got %d drift records for %d comparison records", len(result.Records), len(compared.Changes))
	}
	for _, record := range result.Records {
		if record.Category != IdentityDrift {
			t.Fatalf("host change classified as %q", record.Category)
		}
	}
}

func changeRecord(layer, object string, changeType comparison.ChangeType) comparison.ChangeRecord {
	old := &comparison.TypedValue{Type: "string", Value: "previous-secret"}
	current := &comparison.TypedValue{Type: "string", Value: "current-secret"}
	switch changeType {
	case comparison.Added:
		old = nil
	case comparison.Removed:
		current = nil
	case comparison.Unchanged:
		current = &comparison.TypedValue{Type: "string", Value: "previous-secret"}
	}
	return comparison.ChangeRecord{
		ID: "change", Layer: layer, ObjectID: object,
		Path: "/layers/" + layer + "/resources/" + object + "/facts/value",
		Type: changeType, Previous: old, Current: current,
		ComparisonTimestamp: time.Date(2026, 7, 24, 12, 0, 0, 0, time.UTC),
		Metadata:            map[string]string{"object_kind": "fact", "fact_name": "value"},
	}
}

func contains(value, token string) bool {
	for i := 0; i+len(token) <= len(value); i++ {
		if value[i:i+len(token)] == token {
			return true
		}
	}
	return false
}

func snapshot(id, value string) inventory.Snapshot {
	now := time.Date(2026, 7, 24, 12, 0, 0, 0, time.UTC)
	producer := inventory.Producer{ToolVersion: "test", ContractVersion: "1.0"}
	fact := inventory.CanonicalFact{
		Value: value, ValueType: "string", Quality: "observed", Sensitivity: "public",
		ObservedAt: now, Provenance: inventory.Provenance{
			SourceType: "fixture", SourceLabel: "fixture", ObservedAt: now,
		},
	}
	resource := inventory.Resource{
		ResourceID: "host:fixture", Kind: "host", LayerID: "host", LifecycleState: "observed",
		Facts:         map[string]inventory.CanonicalFact{"name": fact},
		Relationships: []inventory.Relationship{}, Labels: map[string]string{},
		ObservedAt: now, CollectorID: "host", Metadata: map[string]string{},
	}
	layer := inventory.Layer{
		LayerID: "host", ContractVersion: inventory.ContractVersionForLayer,
		Status: inventory.Available, ObservedAt: now, CompletedAt: now,
		CollectorIDs: []string{"host"}, Resources: []inventory.Resource{resource},
		Issues: []inventory.InventoryError{}, Redactions: []string{}, Metadata: map[string]string{},
	}
	category := inventory.Category{
		CategoryID: "host", ContractVersion: "1.0", Status: inventory.Available,
		ObservedAt: now, CompletedAt: now, FreshUntil: now.Add(time.Minute),
		CollectorID: "host", SourceSummary: []string{"fixture"},
		Items: []inventory.Item{}, Errors: []inventory.InventoryError{}, Redactions: []string{},
	}
	canonical := inventory.SystemInventory{
		SchemaName: inventory.CanonicalSchemaName, SchemaVersion: inventory.SchemaVersion,
		Profile: "canonical-system-inventory-v1", SnapshotID: id, RequestID: id,
		SubjectID: "subject", ObservedAt: now, CompletedAt: now, FreshUntil: now.Add(time.Minute),
		Status: inventory.Complete, Producer: producer,
		CollectorResults: []inventory.CollectorExecution{{
			CollectorName: "host", Version: "1.0", Capability: "host",
			SupportedPlatforms: []string{"linux"}, Timestamp: now, Status: inventory.Available,
			Warnings: []inventory.InventoryWarning{}, Errors: []inventory.InventoryError{},
			Metadata: map[string]string{},
		}},
		Layers: []inventory.Layer{layer}, Issues: []inventory.InventoryError{},
		Redactions: []string{}, Metadata: map[string]string{},
	}
	return inventory.Snapshot{
		SchemaVersion: inventory.SchemaVersion, SnapshotID: id, RequestID: id,
		InstanceID: "subject", ObservedAt: now, CompletedAt: now, FreshUntil: now.Add(time.Minute),
		Status: inventory.Complete, Categories: []inventory.Category{category},
		Errors: []inventory.InventoryError{}, Redactions: []string{},
		Producer: producer, Canonical: canonical,
	}
}
