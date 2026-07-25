package main

import (
	"bytes"
	"encoding/json"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"

	canonicalcommand "quantumwizard.hu/qwsg/internal/command"
	"quantumwizard.hu/qwsg/internal/comparison"
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

func TestCanonicalProfilesAndAdvancedGrammarUseSamePipeline(t *testing.T) {
	root := filepath.Join(t.TempDir(), "store")
	store, err := inventorystore.Open(root, inventorystore.DefaultRetention)
	if err != nil {
		t.Fatal(err)
	}
	from := cliFixtureSnapshot()
	fromName, err := store.Save(from)
	if err != nil {
		t.Fatal(err)
	}
	to := laterCLIFixtureSnapshot(from)
	toName, err := store.Save(to)
	if err != nil {
		t.Fatal(err)
	}
	t.Setenv("QWSG_STORE", root)

	for _, profile := range []string{"changes", "health", "report"} {
		var out, errout bytes.Buffer
		if code := run([]string{profile, "--output", "json", "--presentation", "structured"}, &out, &errout); code != 0 {
			t.Fatalf("%s: code=%d error=%q", profile, code, errout.String())
		}
		var execution canonicalcommand.Execution
		if err := json.Unmarshal(out.Bytes(), &execution); err != nil {
			t.Fatal(err)
		}
		if execution.SchemaName != canonicalcommand.ExecutionSchema ||
			execution.CommandID == "" || execution.PlanID == "" || len(execution.Stages) == 0 {
			t.Fatalf("%s returned invalid execution: %#v", profile, execution)
		}
	}

	var simple, advanced, errout bytes.Buffer
	simpleArgs := []string{"health", "--output", "json", "--presentation", "structured"}
	advancedArgs := []string{
		"analyze", "--source", "store", "--store", root, "--pipeline", "health",
		"--output", "json", "--presentation", "structured",
	}
	if code := run(simpleArgs, &simple, &errout); code != 0 {
		t.Fatalf("simple: %d %s", code, errout.String())
	}
	errout.Reset()
	if code := run(advancedArgs, &advanced, &errout); code != 0 {
		t.Fatalf("advanced: %d %s", code, errout.String())
	}
	var simpleExecution, advancedExecution canonicalcommand.Execution
	if err := json.Unmarshal(simple.Bytes(), &simpleExecution); err != nil {
		t.Fatal(err)
	}
	if err := json.Unmarshal(advanced.Bytes(), &advancedExecution); err != nil {
		t.Fatal(err)
	}
	if len(simpleExecution.Stages) != len(advancedExecution.Stages) {
		t.Fatal("simple and advanced pipeline lengths differ")
	}
	for index := range simpleExecution.Stages {
		if simpleExecution.Stages[index].Stage != advancedExecution.Stages[index].Stage {
			t.Fatal("simple and advanced stage ordering differs")
		}
	}
	if !strings.Contains(simple.String(), fromName) || !strings.Contains(simple.String(), toName) {
		t.Fatal("canonical execution lost snapshot traceability")
	}
}

func TestCanonicalCLIRejectsInvalidAndUnsafeComposition(t *testing.T) {
	for _, tc := range []struct {
		args     []string
		contains string
	}{
		{[]string{"changes"}, "--store is required"},
		{[]string{"health", "--exclude", "compare"}, "not valid for a predefined profile"},
		{[]string{"analyze", "--source", "store", "--store", "/tmp/x", "--pipeline", "health", "--exclude", "compare"}, "required"},
		{[]string{"analyze", "--source", "remote", "--pipeline", "health"}, "live or store"},
		{[]string{"analyze", "--source", "store", "--store", "/tmp/x", "--pipeline", "health", "--presentation", "dashboard"}, "unsupported presentation"},
	} {
		var out, errout bytes.Buffer
		if code := run(tc.args, &out, &errout); code != 1 || !strings.Contains(errout.String(), tc.contains) {
			t.Fatalf("%v: code=%d output=%q", tc.args, code, out.String()+errout.String())
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

func TestCompareCLIDefaultExplicitAndHuman(t *testing.T) {
	root := filepath.Join(t.TempDir(), "store")
	store, err := inventorystore.Open(root, 3)
	if err != nil {
		t.Fatal(err)
	}
	from := cliFixtureSnapshot()
	fromName, err := store.Save(from)
	if err != nil {
		t.Fatal(err)
	}
	to := laterCLIFixtureSnapshot(from)
	toName, err := store.Save(to)
	if err != nil {
		t.Fatal(err)
	}

	var first, second, errout bytes.Buffer
	args := []string{"compare", "--store", root, "--retention", "3"}
	if code := run(args, &first, &errout); code != 0 {
		t.Fatalf("default compare: code=%d error=%q", code, errout.String())
	}
	errout.Reset()
	if code := run(args, &second, &errout); code != 0 || first.String() != second.String() {
		t.Fatalf("repeat compare: code=%d equal=%v error=%q", code, first.String() == second.String(), errout.String())
	}
	var result comparison.Result
	if err := json.Unmarshal(first.Bytes(), &result); err != nil {
		t.Fatal(err)
	}
	if result.From.Selector != fromName || result.To.Selector != toName || result.Counts.Added != 2 {
		t.Fatalf("unexpected comparison: %#v", result)
	}

	var human bytes.Buffer
	errout.Reset()
	explicit := []string{"compare", "--store", root, "--retention", "3", "--from", fromName, "--to", toName, "--format", "human"}
	if code := run(explicit, &human, &errout); code != 0 {
		t.Fatalf("human compare: code=%d error=%q", code, errout.String())
	}
	for _, expected := range []string{"Added (2)", "Removed (0)", "Modified (0)", "Unchanged (", ` = "changed"`} {
		if !strings.Contains(human.String(), expected) {
			t.Fatalf("human output %q lacks %q", human.String(), expected)
		}
	}
}

func TestCompareCLIRejectsIncompleteSelectionAndInsufficientHistory(t *testing.T) {
	root := filepath.Join(t.TempDir(), "store")
	store, err := inventorystore.Open(root, 2)
	if err != nil {
		t.Fatal(err)
	}
	name, err := store.Save(cliFixtureSnapshot())
	if err != nil {
		t.Fatal(err)
	}
	for _, tc := range []struct {
		args     []string
		contains string
	}{
		{[]string{"compare", "--store", root, "--retention", "2"}, "at least two"},
		{[]string{"compare", "--store", root, "--retention", "2", "--from", name}, "must be provided together"},
		{[]string{"compare", "--store", root, "--retention", "2", "--from", name, "--to", "missing.json"}, "to snapshot load failed"},
	} {
		var out, errout bytes.Buffer
		if code := run(tc.args, &out, &errout); code != 1 || !strings.Contains(errout.String(), tc.contains) {
			t.Fatalf("%v: code=%d output=%q", tc.args, code, out.String()+errout.String())
		}
	}
}

func laterCLIFixtureSnapshot(snapshot inventory.Snapshot) inventory.Snapshot {
	snapshot.SnapshotID = "cli-fixture-later"
	snapshot.RequestID = "cli-fixture-later"
	snapshot.CompletedAt = snapshot.CompletedAt.Add(time.Minute)
	snapshot.FreshUntil = snapshot.FreshUntil.Add(time.Minute)
	snapshot.Canonical.SnapshotID = snapshot.SnapshotID
	snapshot.Canonical.RequestID = snapshot.RequestID
	snapshot.Canonical.CompletedAt = snapshot.CompletedAt
	snapshot.Canonical.FreshUntil = snapshot.FreshUntil
	layer := &snapshot.Canonical.Layers[0]
	layer.Resources = append(layer.Resources, inventory.Resource{
		ResourceID: "task018:fixture", Kind: "fixture", LayerID: layer.LayerID,
		LifecycleState: "observed", CollectorID: "host", ObservedAt: snapshot.ObservedAt,
		Facts: map[string]inventory.CanonicalFact{
			"task_018_fixture": {
				Value: "changed", ValueType: "string", Quality: "observed", Sensitivity: "public",
				ObservedAt: snapshot.ObservedAt,
				Provenance: inventory.Provenance{
					SourceType: "fixture", SourceLabel: "fixture", ObservedAt: snapshot.ObservedAt,
				},
			},
		},
		Relationships: []inventory.Relationship{}, Labels: map[string]string{},
		Metadata: map[string]string{},
	})
	return snapshot
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
