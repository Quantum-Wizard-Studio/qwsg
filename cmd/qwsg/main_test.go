package main

import (
	"bytes"
	"encoding/json"
	"path/filepath"
	"strings"
	"testing"
	"time"

	"quantumwizard.hu/qwsg/internal/inventory"
	"quantumwizard.hu/qwsg/internal/inventorystore"
)

func TestCLIHelpVersionAndInvalid(t *testing.T) {
	for _, tc := range []struct {
		args     []string
		code     int
		contains string
	}{{[]string{"help"}, 0, "Usage"}, {[]string{"version"}, 0, "0.0.1"}, {[]string{"unknown"}, 1, "unknown command"}} {
		var out, err bytes.Buffer
		code := run(tc.args, &out, &err)
		if code != tc.code || !strings.Contains(out.String()+err.String(), tc.contains) {
			t.Fatalf("%v: %d %q", tc.args, code, out.String()+err.String())
		}
	}
}

func TestInventoryStoreArguments(t *testing.T) {
	absolute := filepath.Join(t.TempDir(), "store")
	for _, tc := range []struct {
		args     []string
		contains string
	}{
		{[]string{"inventory", "load"}, "--store is required"},
		{[]string{"inventory", "load", "--store", "relative"}, "clean absolute path"},
		{[]string{"inventory", "save", "--store", absolute, "--snapshot", "x"}, "--snapshot is valid only"},
		{[]string{"inventory", "load", "--store", absolute, "--retention", "zero"}, "--retention must be an integer"},
		{[]string{"inventory", "other", "--store", absolute}, "only save or load"},
	} {
		var out, err bytes.Buffer
		if code := run(tc.args, &out, &err); code != 1 || !strings.Contains(err.String(), tc.contains) {
			t.Fatalf("%v: code=%d output=%q", tc.args, code, out.String()+err.String())
		}
	}
}

func TestInventoryLoadCLIUsesPersistedSnapshot(t *testing.T) {
	root := filepath.Join(t.TempDir(), "store")
	store, err := inventorystore.Open(root, 2)
	if err != nil {
		t.Fatal(err)
	}
	snapshot := cliFixtureSnapshot()
	name, err := store.Save(snapshot)
	if err != nil {
		t.Fatal(err)
	}
	for _, args := range [][]string{
		{"inventory", "load", "--store", root, "--retention", "2"},
		{"inventory", "load", "--store", root, "--retention", "2", "--snapshot", name},
	} {
		var out, errout bytes.Buffer
		if code := run(args, &out, &errout); code != 0 {
			t.Fatalf("%v: code=%d error=%q", args, code, errout.String())
		}
		var got inventory.Snapshot
		if err := json.Unmarshal(out.Bytes(), &got); err != nil {
			t.Fatal(err)
		}
		if got.SnapshotID != snapshot.SnapshotID || got.Canonical.SchemaName != inventory.CanonicalSchemaName {
			t.Fatalf("unexpected loaded snapshot: %#v", got)
		}
	}
}

func cliFixtureSnapshot() inventory.Snapshot {
	now := time.Date(2026, 7, 24, 4, 0, 0, 0, time.UTC)
	category := inventory.Category{
		CategoryID: "host", ContractVersion: "1.0", Status: inventory.Available,
		ObservedAt: now, CompletedAt: now, FreshUntil: now.Add(time.Minute),
		CollectorID: "host", PrivilegeUsed: "ordinary-user", SourceSummary: []string{"fixture"},
		Items: []inventory.Item{}, Errors: []inventory.InventoryError{}, Redactions: []string{},
	}
	execution := inventory.CollectorExecution{
		CollectorName: "host", Version: "1", Capability: "host", SupportedPlatforms: []string{"linux"},
		Timestamp: now, Status: inventory.Available, Warnings: []inventory.InventoryWarning{},
		Errors: []inventory.InventoryError{}, Metadata: map[string]string{},
	}
	producer := inventory.Producer{ToolVersion: "test", ContractVersion: "1.0"}
	snapshot := inventory.Snapshot{
		SchemaVersion: "1.0", SnapshotID: "cli-fixture", RequestID: "cli-fixture", InstanceID: "subject",
		ObservedAt: now, CompletedAt: now, FreshUntil: now.Add(time.Minute), Status: inventory.Complete,
		Categories: []inventory.Category{category}, Errors: []inventory.InventoryError{},
		Redactions: []string{}, Producer: producer,
	}
	snapshot.Canonical = inventory.AssembleSystemInventory(
		snapshot.Categories, []inventory.CollectorExecution{execution},
		snapshot.SnapshotID, snapshot.RequestID, snapshot.InstanceID,
		now, now, now.Add(time.Minute), 0, producer,
	)
	return snapshot
}
