package main

import (
	"bytes"
	"encoding/json"
	"os"
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
	}{
		{[]string{"help"}, 0, "snapshot explorer"},
		{[]string{"help", "inventory"}, 0, "inventory list"},
		{[]string{"inventory", "list", "--help"}, 0, "inventory list"},
		{[]string{"version"}, 0, "commit:"},
		{[]string{"unknown"}, 1, "unknown command"},
	} {
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
		{[]string{"inventory", "other", "--store", absolute}, "unknown inventory subcommand"},
		{[]string{"inventory", "--format", "yaml"}, "--format must be json or human"},
		{[]string{"inventory", "list", "--store", absolute, "--snapshot", "x"}, "--snapshot is valid only"},
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

func TestSnapshotExplorerListInfoAndHumanLoad(t *testing.T) {
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
	tests := []struct {
		args     []string
		code     int
		contains []string
	}{
		{[]string{"inventory", "list", "--store", root, "--retention", "2"}, 0, []string{"Snapshots: 1", name, "status=complete"}},
		{[]string{"inventory", "list", "--store", root, "--retention", "2", "--format", "json"}, 0, []string{`"integrity": "verified"`, `"name":`}},
		{[]string{"inventory", "info", "--store", root, "--retention", "2", "--snapshot", name}, 0, []string{"Snapshot:", "Integrity: verified", "Status: complete"}},
		{[]string{"inventory", "load", "--store", root, "--retention", "2", "--format", "human"}, 0, []string{"Stored inventory", "Canonical layers:", "Status: complete"}},
	}
	for _, tc := range tests {
		var out, errout bytes.Buffer
		if code := run(tc.args, &out, &errout); code != tc.code {
			t.Fatalf("%v: code=%d error=%q", tc.args, code, errout.String())
		}
		for _, expected := range tc.contains {
			if !strings.Contains(out.String(), expected) {
				t.Fatalf("%v: output %q does not contain %q", tc.args, out.String(), expected)
			}
		}
	}
}

func TestDocumentedEnvironmentSupportsConciseExplorerCommands(t *testing.T) {
	root := filepath.Join(t.TempDir(), "store")
	store, err := inventorystore.Open(root, inventorystore.DefaultRetention)
	if err != nil {
		t.Fatal(err)
	}
	if _, err := store.Save(cliFixtureSnapshot()); err != nil {
		t.Fatal(err)
	}
	t.Setenv("QWSG_STORE", root)
	t.Setenv("QWSG_FORMAT", "human")
	for _, args := range [][]string{
		{"inventory", "list"},
		{"inventory", "info"},
		{"inventory", "load"},
	} {
		var out, errout bytes.Buffer
		if code := run(args, &out, &errout); code != 0 || out.Len() == 0 {
			t.Fatalf("%v: code=%d out=%q err=%q", args, code, out.String(), errout.String())
		}
	}
}

func TestSnapshotExplorerEmptyAndInvalidStore(t *testing.T) {
	root := filepath.Join(t.TempDir(), "store")
	store, err := inventorystore.Open(root, 2)
	if err != nil {
		t.Fatal(err)
	}
	name, err := store.Save(cliFixtureSnapshot())
	if err != nil {
		t.Fatal(err)
	}
	if err := os.Remove(filepath.Join(root, "snapshots", name)); err != nil {
		t.Fatal(err)
	}
	var out, errout bytes.Buffer
	if code := run([]string{"inventory", "list", "--store", root, "--retention", "2"}, &out, &errout); code != 0 || out.String() != "Snapshots: none\n" {
		t.Fatalf("empty list: code=%d out=%q err=%q", code, out.String(), errout.String())
	}
	out.Reset()
	errout.Reset()
	if code := run([]string{"inventory", "info", "--store", root, "--retention", "2"}, &out, &errout); code != 1 || !strings.Contains(errout.String(), "file does not exist") {
		t.Fatalf("empty info: code=%d out=%q err=%q", code, out.String(), errout.String())
	}
}

func TestHumanOutputEscapesTerminalControls(t *testing.T) {
	input := "name\x1b[31m\n\t\u007f"
	got := safeText(input)
	if strings.ContainsAny(got, "\x1b\n\t\u007f") || got != `name\u001b[31m\u000a\u0009\u007f` {
		t.Fatalf("unsafe rendering: %q", got)
	}
}

func TestJSONCompatibilityOutput(t *testing.T) {
	snapshot := cliFixtureSnapshot()
	var out, errout bytes.Buffer
	if code := writeInventory(&out, &errout, snapshot, formatJSON, "ignored"); code != 0 {
		t.Fatalf("code=%d error=%q", code, errout.String())
	}
	var got inventory.Snapshot
	if err := json.Unmarshal(out.Bytes(), &got); err != nil {
		t.Fatal(err)
	}
	if got.SnapshotID != snapshot.SnapshotID || got.Canonical.SchemaName != inventory.CanonicalSchemaName {
		t.Fatalf("compatibility envelope changed: %#v", got)
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
