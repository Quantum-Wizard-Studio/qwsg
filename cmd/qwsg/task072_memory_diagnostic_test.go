package main

import (
	"context"
	"fmt"
	"os"
	"runtime"
	"testing"
	"time"

	canonicalcommand "quantumwizard.hu/qwsg/internal/command"
	"quantumwizard.hu/qwsg/internal/inventory"
	"quantumwizard.hu/qwsg/internal/inventorystore"
	"quantumwizard.hu/qwsg/internal/pipeline"
)

func task072Snapshot(now time.Time, generation, items int) inventory.Snapshot {
	values := make([]inventory.Item, 0, items)
	for index := 0; index < items; index++ {
		values = append(values, inventory.Item{ID: fmt.Sprintf("service-%06d", index), Kind: "service", Facts: map[string]inventory.Fact{
			"active":   {Value: true, Quality: "observed", Sensitivity: "operational", Provenance: inventory.Provenance{SourceType: "fixture", SourceLabel: "loaded-host", ObservedAt: now}},
			"identity": {Quality: "redacted", Sensitivity: "operational", Reason: "service_identity_hidden", Provenance: inventory.Provenance{SourceType: "fixture", SourceLabel: "loaded-host", ObservedAt: now}},
		}})
	}
	category := inventory.Category{CategoryID: "services", ContractVersion: "1.0", Status: inventory.Available, ObservedAt: now, CompletedAt: now, FreshUntil: now.Add(5 * time.Minute), CollectorID: "services", PrivilegeUsed: "ordinary-user", SourceSummary: []string{"bounded loaded-host fixture"}, Items: values, Errors: []inventory.InventoryError{}, Redactions: []string{}}
	execution := inventory.CollectorExecution{CollectorName: "services", Version: "1", Capability: "services", SupportedPlatforms: []string{"linux"}, Timestamp: now, Status: inventory.Available, Warnings: []inventory.InventoryWarning{}, Errors: []inventory.InventoryError{}, Metadata: map[string]string{}}
	producer := inventory.Producer{ToolVersion: "test", ContractVersion: "1.0"}
	id := fmt.Sprintf("task072-%d", generation)
	snapshot := inventory.Snapshot{SchemaVersion: inventory.SchemaVersion, SnapshotID: id, RequestID: id, InstanceID: "subject", ObservedAt: now, CompletedAt: now, FreshUntil: now.Add(5 * time.Minute), Status: inventory.Complete, Categories: []inventory.Category{category}, Errors: []inventory.InventoryError{}, Redactions: []string{}, Producer: producer}
	snapshot.Canonical = inventory.AssembleSystemInventory(snapshot.Categories, []inventory.CollectorExecution{execution}, id, id, "subject", now, now, snapshot.FreshUntil, 0, producer)
	return snapshot
}

func BenchmarkTask072QualifiedObservation(b *testing.B) {
	root := b.TempDir()
	if err := os.Chmod(root, 0o700); err != nil {
		b.Fatal(err)
	}
	definition, err := canonicalcommand.ResolveProfile("observe", canonicalcommand.Selection{Source: "live", Store: root})
	if err != nil {
		b.Fatal(err)
	}
	now := time.Date(2026, 8, 29, 0, 0, 0, 0, time.UTC)
	store, err := inventorystore.Open(root, inventorystore.DefaultRetention)
	if err != nil {
		b.Fatal(err)
	}
	if _, err = store.Save(task072Snapshot(now, 0, 367)); err != nil {
		b.Fatal(err)
	}
	b.ReportAllocs()
	b.ResetTimer()
	for index := 0; index < b.N; index++ {
		snapshot := task072Snapshot(now.Add(time.Duration(index+1)*time.Minute), index+1, 367)
		orchestrator := pipeline.Orchestrator{Collect: func(context.Context) (inventory.Snapshot, error) { return snapshot, nil }, Retention: inventorystore.DefaultRetention, Rules: pipeline.CanonicalObservationRules()}
		if _, err = orchestrator.Execute(context.Background(), definition); err != nil {
			b.Fatal(err)
		}
		runtime.GC()
	}
}
