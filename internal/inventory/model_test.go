package inventory

import (
	"bytes"
	"encoding/json"
	"reflect"
	"testing"
	"time"
)

func category(id string, s CategoryStatus) Category {
	n := time.Now().UTC()
	return Category{CategoryID: id, Status: s, ObservedAt: n, CompletedAt: n, FreshUntil: n.Add(time.Minute)}
}

func TestCanonicalAssemblyIsDeterministicAndValidated(t *testing.T) {
	now := time.Date(2026, 7, 21, 12, 0, 0, 0, time.UTC)
	fact := Fact{Value: int64(1024), Unit: "bytes", Quality: "observed", Sensitivity: "operational", Provenance: Provenance{SourceType: "fixture", SourceLabel: "test", ObservedAt: now}}
	categories := []Category{
		{CategoryID: "network", ContractVersion: "1.0", Status: Available, ObservedAt: now, CompletedAt: now, CollectorID: "network", Items: []Item{{ID: "b", Kind: "interface", Facts: map[string]Fact{"mtu": fact}}}, Errors: []InventoryError{}},
		{CategoryID: "host", ContractVersion: "1.0", Status: Available, ObservedAt: now, CompletedAt: now, CollectorID: "host", Items: []Item{{ID: "a", Kind: "host", Facts: map[string]Fact{"capacity": fact}}}, Errors: []InventoryError{}},
	}
	executions := []CollectorExecution{{CollectorName: "network", Version: "1", Capability: "network", SupportedPlatforms: []string{"linux"}, Status: Available}, {CollectorName: "host", Version: "1", Capability: "host", SupportedPlatforms: []string{"linux"}, Status: Available}}
	producer := Producer{ToolVersion: "test", ContractVersion: "1.0"}
	got := AssembleSystemInventory(categories, executions, "snapshot", "request", "subject", now, now.Add(time.Second), now.Add(time.Minute), 1000, producer)
	if err := ValidateSystemInventory(got); err != nil {
		t.Fatal(err)
	}
	if got.Layers[0].LayerID != "host" || got.Layers[1].LayerID != "network" || got.CollectorResults[0].CollectorName != "host" {
		t.Fatalf("not sorted: %#v", got)
	}
	reversed := []Category{categories[1], categories[0]}
	reverseExecutions := []CollectorExecution{executions[1], executions[0]}
	want := AssembleSystemInventory(reversed, reverseExecutions, "snapshot", "request", "subject", now, now.Add(time.Second), now.Add(time.Minute), 1000, producer)
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("assembly differs by input order:\n%#v\n%#v", got, want)
	}
	gotJSON, err := json.Marshal(got)
	if err != nil {
		t.Fatal(err)
	}
	wantJSON, err := json.Marshal(want)
	if err != nil {
		t.Fatal(err)
	}
	if !bytes.Equal(gotJSON, wantJSON) {
		t.Fatal("canonical JSON differs by input order")
	}
}

func TestCanonicalValidationRejectsSecretsAndBrokenRelationships(t *testing.T) {
	now := time.Now().UTC()
	system := SystemInventory{
		SchemaName: CanonicalSchemaName, SchemaVersion: SchemaVersion, SnapshotID: "snapshot", RequestID: "request", SubjectID: "subject",
		ObservedAt: now, CompletedAt: now, FreshUntil: now.Add(time.Minute), Status: Complete, Producer: Producer{ToolVersion: "test", ContractVersion: "1.0"},
		CollectorResults: []CollectorExecution{{CollectorName: "host", Capability: "host", SupportedPlatforms: []string{"linux"}}},
		Layers: []Layer{{
			LayerID: "host", ContractVersion: "1.0", Status: Available, ObservedAt: now, CompletedAt: now, CollectorIDs: []string{"host"},
			Resources: []Resource{{
				ResourceID: "host:one", Kind: "host", LayerID: "host", LifecycleState: "observed", ObservedAt: now, CollectorID: "host",
				Facts: map[string]CanonicalFact{"secret": {
					ValueType: "string", Quality: "observed", Sensitivity: "secret_prohibited",
					Provenance: Provenance{SourceType: "fixture"},
				}},
				Relationships: []Relationship{},
			}},
		}},
	}
	if err := ValidateSystemInventory(system); err == nil {
		t.Fatal("secret_prohibited fact accepted")
	}
	system.Layers[0].Resources[0].Facts = map[string]CanonicalFact{}
	system.Layers[0].Resources[0].Relationships = []Relationship{{RelationshipType: "contains", SourceResourceID: "host:one", TargetResourceID: "missing"}}
	if err := ValidateSystemInventory(system); err == nil {
		t.Fatal("broken relationship accepted")
	}
}
func TestAggregationAndExitCodes(t *testing.T) {
	cases := []struct {
		cs     []Category
		status Status
		code   int
	}{{[]Category{category("a", Available)}, Complete, 0}, {[]Category{category("a", Available), category("b", Timeout)}, Partial, 2}, {[]Category{category("a", Unavailable)}, Failed, 1}}
	for _, tc := range cases {
		if got := Aggregate(tc.cs); got != tc.status || ExitCode(got) != tc.code {
			t.Fatalf("got %s/%d", got, ExitCode(got))
		}
	}
}
func TestValidateRejectsFutureSchema(t *testing.T) {
	n := time.Now().UTC()
	s := Snapshot{SchemaVersion: "2.0", ObservedAt: n, CompletedAt: n, FreshUntil: n.Add(time.Minute), Status: Complete, Categories: []Category{category("a", Available)}}
	if Validate(s) == nil {
		t.Fatal("future schema accepted")
	}
}

func TestValidateRejectsCanonicalEnvelopeMismatch(t *testing.T) {
	now := time.Now().UTC()
	category := Category{CategoryID: "host", ContractVersion: "1.0", Status: Available, ObservedAt: now, CompletedAt: now, FreshUntil: now.Add(time.Minute), CollectorID: "host", Items: []Item{}, Errors: []InventoryError{}, Redactions: []string{}}
	producer := Producer{ToolVersion: "test", ContractVersion: "1.0"}
	execution := CollectorExecution{CollectorName: "host", Version: "1", Capability: "host", SupportedPlatforms: []string{"linux"}, Timestamp: now, Status: Available, Warnings: []InventoryWarning{}, Errors: []InventoryError{}, Metadata: map[string]string{}}
	snapshot := Snapshot{SchemaVersion: SchemaVersion, SnapshotID: "snapshot", RequestID: "request", InstanceID: "subject", ObservedAt: now, CompletedAt: now, FreshUntil: now.Add(time.Minute), Status: Complete, Categories: []Category{category}, Producer: producer}
	snapshot.Canonical = AssembleSystemInventory(snapshot.Categories, []CollectorExecution{execution}, snapshot.SnapshotID, snapshot.RequestID, snapshot.InstanceID, snapshot.ObservedAt, snapshot.CompletedAt, snapshot.FreshUntil, snapshot.DurationMS, producer)
	if err := Validate(snapshot); err != nil {
		t.Fatal(err)
	}
	snapshot.Canonical.SubjectID = "different"
	if err := Validate(snapshot); err == nil {
		t.Fatal("mismatched canonical envelope accepted")
	}
}
