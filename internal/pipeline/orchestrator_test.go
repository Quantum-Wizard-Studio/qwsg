package pipeline

import (
	"context"
	"path/filepath"
	"reflect"
	"testing"
	"time"

	"quantumwizard.hu/qwsg/internal/command"
	"quantumwizard.hu/qwsg/internal/drift"
	"quantumwizard.hu/qwsg/internal/health"
	"quantumwizard.hu/qwsg/internal/inventory"
	"quantumwizard.hu/qwsg/internal/inventorystore"
	"quantumwizard.hu/qwsg/internal/report"
	"quantumwizard.hu/qwsg/internal/rule"
)

func TestOnlySelectedStagesExecute(t *testing.T) {
	calls := 0
	orchestrator := Orchestrator{Collect: func(context.Context) (inventory.Snapshot, error) {
		calls++
		return fixture("current", time.Date(2026, 7, 24, 12, 0, 0, 0, time.UTC), "a"), nil
	}}
	definition, _ := command.ResolveProfile("status", command.Selection{})
	execution, err := orchestrator.Execute(context.Background(), definition)
	if err != nil {
		t.Fatal(err)
	}
	if calls != 1 || len(execution.Stages) != 1 || execution.Stages[0].Stage != command.Inventory {
		t.Fatalf("unexpected execution: calls=%d stages=%v", calls, execution.Stages)
	}
}

func TestCanonicalFullPipelineIsDeterministicAndTraceable(t *testing.T) {
	root := filepath.Join(t.TempDir(), "store")
	store, err := inventorystore.Open(root, 3)
	if err != nil {
		t.Fatal(err)
	}
	from := fixture("from", time.Date(2026, 7, 24, 12, 0, 0, 0, time.UTC), "a")
	to := fixture("to", time.Date(2026, 7, 24, 12, 1, 0, 0, time.UTC), "b")
	if _, err := store.Save(from); err != nil {
		t.Fatal(err)
	}
	if _, err := store.Save(to); err != nil {
		t.Fatal(err)
	}
	definition, err := command.ResolveProfile("report", command.Selection{Source: "store", Store: root})
	if err != nil {
		t.Fatal(err)
	}
	warning := string(health.Warning)
	rules := []rule.Definition{{
		ID: "canonical.warning", ContractVersion: rule.RuleVersion,
		Category: rule.StatusRule, Enabled: true,
		Scope:             rule.Scope{HealthIDs: []string{}, Categories: []drift.Category{}},
		InputRequirements: []rule.Field{rule.FieldStatus},
		Condition: rule.Condition{
			Operator: rule.Equal, Field: rule.FieldStatus,
			Value: rule.Value{String: &warning}, Values: []rule.Value{}, Children: []rule.Condition{},
		},
		Description: "Canonical warning observation", Metadata: map[string]string{},
	}}
	orchestrator := Orchestrator{Retention: 3, Rules: rules}
	first, err := orchestrator.Execute(context.Background(), definition)
	if err != nil {
		t.Fatal(err)
	}
	second, err := orchestrator.Execute(context.Background(), definition)
	if err != nil {
		t.Fatal(err)
	}
	expected := []command.Stage{command.Compare, command.Drift, command.Health, command.Rule, command.Report}
	got := make([]command.Stage, len(first.Stages))
	for index, result := range first.Stages {
		got[index] = result.Stage
	}
	if !reflect.DeepEqual(got, expected) || first.ID != second.ID {
		t.Fatalf("stages=%v deterministic=%v", got, first.ID == second.ID)
	}
	final, ok := first.Stages[len(first.Stages)-1].Value.(report.Report)
	if !ok || len(final.Sources) != final.Summary.Total || final.Summary.Total == 0 {
		t.Fatalf("untraceable report: %#v", final)
	}
}

func TestStageFailureStopsPipeline(t *testing.T) {
	definition, _ := command.ResolveProfile("health", command.Selection{Source: "store", Store: filepath.Join(t.TempDir(), "missing")})
	execution, err := (Orchestrator{}).Execute(context.Background(), definition)
	if err == nil || execution.Complete || len(execution.Stages) != 1 ||
		execution.Stages[0].Stage != command.Compare || len(execution.Diagnostics) != 1 {
		t.Fatalf("failure did not stop cleanly: execution=%#v error=%v", execution, err)
	}
}

func fixture(id string, at time.Time, value string) inventory.Snapshot {
	category := inventory.Category{
		CategoryID: "host", ContractVersion: "1.0", Status: inventory.Available,
		ObservedAt: at, CompletedAt: at, FreshUntil: at.Add(time.Minute),
		CollectorID: "host", PrivilegeUsed: "ordinary-user", SourceSummary: []string{"fixture"},
		Items: []inventory.Item{{ID: "host", Kind: "host", Facts: map[string]inventory.Fact{
			"value": {Value: value, Quality: "observed", Sensitivity: "public", Provenance: inventory.Provenance{SourceType: "fixture", SourceLabel: "fixture", ObservedAt: at}},
		}}},
		Errors: []inventory.InventoryError{}, Redactions: []string{},
	}
	execution := inventory.CollectorExecution{
		CollectorName: "host", Version: "1", Capability: "host", SupportedPlatforms: []string{"linux"},
		Timestamp: at, Status: inventory.Available, Warnings: []inventory.InventoryWarning{},
		Errors: []inventory.InventoryError{}, Metadata: map[string]string{},
	}
	producer := inventory.Producer{ToolVersion: "test", ContractVersion: "1.0"}
	snapshot := inventory.Snapshot{
		SchemaVersion: "1.0", SnapshotID: id, RequestID: id, InstanceID: "subject",
		ObservedAt: at, CompletedAt: at, FreshUntil: at.Add(time.Minute), Status: inventory.Complete,
		Categories: []inventory.Category{category}, Errors: []inventory.InventoryError{},
		Redactions: []string{}, Producer: producer,
	}
	snapshot.Canonical = inventory.AssembleSystemInventory(
		snapshot.Categories, []inventory.CollectorExecution{execution},
		id, id, "subject", at, at, at.Add(time.Minute), 0, producer,
	)
	return snapshot
}
