package collector

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"reflect"
	"sort"
	"strings"
	"testing"
	"time"

	"quantumwizard.hu/qwsg/internal/command"
	"quantumwizard.hu/qwsg/internal/comparison"
	"quantumwizard.hu/qwsg/internal/drift"
	"quantumwizard.hu/qwsg/internal/health"
	"quantumwizard.hu/qwsg/internal/inventory"
	"quantumwizard.hu/qwsg/internal/inventorystore"
	"quantumwizard.hu/qwsg/internal/pipeline"
	"quantumwizard.hu/qwsg/internal/report"
	"quantumwizard.hu/qwsg/internal/rule"
	"quantumwizard.hu/qwsg/internal/runner"
)

type serviceRunner struct{ output string }

func (r serviceRunner) Run(_ context.Context, id string, args ...string) (runner.Result, error) {
	if id != "systemctl" || !reflect.DeepEqual(args, []string{"list-units", "--type=service", "--state=running", "--no-legend", "--plain", "--full"}) {
		return runner.Result{}, errors.New("service collection scope changed")
	}
	return runner.Result{Stdout: []byte(r.output)}, nil
}

func serviceItems(t *testing.T, names ...string) []inventory.Item {
	t.Helper()
	var output strings.Builder
	for _, name := range names {
		fmt.Fprintf(&output, "%s loaded active running private description\n", name)
	}
	items, _, err := collectServices(serviceRunner{output.String()})(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	return items
}

// Pin observation times only: independent collections must retain identical
// semantic IDs and ordering while real acquisition timestamps may differ.
func serviceSnapshot(t *testing.T, id string, items []inventory.Item) inventory.Snapshot {
	t.Helper()
	now := time.Date(2026, 9, 21, 0, 0, 0, 0, time.UTC)
	for i := range items {
		for key, fact := range items[i].Facts {
			fact.Provenance.ObservedAt = now
			items[i].Facts[key] = fact
		}
	}
	category := inventory.Category{CategoryID: "services", ContractVersion: "1.0", Status: inventory.Available, ObservedAt: now, CompletedAt: now, FreshUntil: now.Add(time.Minute), CollectorID: "services", PrivilegeUsed: "ordinary_user", SourceSummary: []string{"systemd unit metadata"}, Items: items, Errors: []inventory.InventoryError{}, Redactions: []string{}}
	producer := inventory.Producer{ToolVersion: "test", ContractVersion: "1.0"}
	executions := []inventory.CollectorExecution{{CollectorName: "services", Version: "1.0.0", Capability: "services", Timestamp: now, Status: inventory.Available}}
	s := inventory.Snapshot{SchemaVersion: inventory.SchemaVersion, SnapshotID: id, RequestID: id, InstanceID: "fixture-subject", ObservedAt: now, CompletedAt: now, FreshUntil: now.Add(time.Minute), Status: inventory.Complete, Categories: []inventory.Category{category}, Producer: producer, Errors: []inventory.InventoryError{}, Redactions: []string{}}
	s.Canonical = inventory.AssembleSystemInventory(s.Categories, executions, id, id, s.InstanceID, now, now, s.FreshUntil, 0, producer)
	if err := inventory.Validate(s); err != nil {
		t.Fatal(err)
	}
	return s
}

func assertServicePrivacy(t *testing.T, values ...any) {
	t.Helper()
	for _, value := range values {
		data, err := json.Marshal(value)
		if err != nil {
			t.Fatal(err)
		}
		for _, raw := range []string{"alpha.service", "beta.service", "gamma.service", "delta.service", "private description"} {
			if strings.Contains(string(data), raw) {
				t.Fatalf("raw service evidence leaked: %s", raw)
			}
		}
	}
}

func TestServiceIdentityAcceptance(t *testing.T) {
	abc := []string{"alpha.service", "beta.service", "gamma.service"}
	cases := []struct {
		name           string
		from, to       []string
		added, removed []string
	}{
		{"repeated observation", abc, abc, nil, nil},
		{"enumeration reorder", abc, []string{"gamma.service", "alpha.service", "beta.service"}, nil, nil},
		{"disappearance", abc, []string{"alpha.service", "gamma.service"}, nil, []string{"beta.service"}},
		{"appearance", []string{"alpha.service", "gamma.service"}, abc, []string{"beta.service"}, nil},
		{"same count replacement", abc, []string{"alpha.service", "delta.service", "gamma.service"}, []string{"delta.service"}, []string{"beta.service"}},
		{"empty to present", nil, abc, abc, nil},
		{"present to empty", abc, nil, nil, abc},
	}
	// Independent singleton observations identify each service without depending
	// on its position in any of the multi-service collections.
	ids := map[string]string{}
	for _, name := range []string{"alpha.service", "beta.service", "gamma.service", "delta.service"} {
		ids[name] = "services:" + serviceItems(t, name)[0].ID
	}
	unique := map[string]bool{}
	for _, id := range ids {
		if unique[id] {
			t.Fatal("distinct services collapsed")
		}
		unique[id] = true
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			from := serviceSnapshot(t, "from", serviceItems(t, tc.from...))
			to := serviceSnapshot(t, "to", serviceItems(t, tc.to...))
			for _, s := range []inventory.Snapshot{from, to} {
				items := s.Categories[0].Items
				if !sort.SliceIsSorted(items, func(i, j int) bool { return items[i].ID < items[j].ID }) {
					t.Fatal("collector order is not canonical")
				}
			}
			if len(tc.added)+len(tc.removed) == 0 {
				a, _ := json.Marshal(from.Canonical.Layers)
				b, _ := json.Marshal(to.Canonical.Layers)
				if string(a) != string(b) {
					t.Fatal("repeated/reordered canonical layers differ")
				}
			}
			result, err := comparison.Compare(from, to, "from", "to")
			if err != nil {
				t.Fatal(err)
			}
			again, err := comparison.Compare(from, to, "from", "to")
			if err != nil {
				t.Fatal(err)
			}
			a, _ := json.Marshal(result)
			b, _ := json.Marshal(again)
			if string(a) != string(b) {
				t.Fatal("comparison is nondeterministic")
			}
			expected := map[string]comparison.ChangeType{}
			for _, name := range tc.added {
				expected[ids[name]] = comparison.Added
			}
			for _, name := range tc.removed {
				expected[ids[name]] = comparison.Removed
			}
			found := map[string]comparison.ChangeType{}
			for _, c := range result.Changes {
				if c.Type == comparison.Unchanged {
					continue
				}
				if c.Type != expected[c.ObjectID] {
					t.Fatalf("unexpected change: %#v", c)
				}
				if c.Metadata["object_kind"] == "resource" {
					found[c.ObjectID] = c.Type
				}
			}
			if !reflect.DeepEqual(found, expected) {
				t.Fatalf("lost identity evidence: %v != %v", found, expected)
			}
			classified, err := drift.Classify(result.Changes)
			if err != nil {
				t.Fatal(err)
			}
			for _, d := range classified.Records {
				if d.Category != drift.ServiceDrift {
					t.Fatal("service change lost in Drift")
				}
				want := drift.StateUnchanged
				switch expected[d.Scope.ObjectID] {
				case comparison.Added:
					want = drift.PresenceAdded
				case comparison.Removed:
					want = drift.PresenceRemoved
				}
				if d.Classification != want {
					t.Fatal("incorrect Drift classification")
				}
			}
			evaluated, err := health.Evaluate(classified)
			if err != nil {
				t.Fatal(err)
			}
			status := "warning"
			ruled, err := rule.Evaluate([]rule.Definition{{ID: "fixture.service-evidence", ContractVersion: rule.RuleVersion, Category: rule.StatusRule, Scope: rule.Scope{HealthIDs: []string{}, Categories: []drift.Category{}}, Enabled: true, InputRequirements: []rule.Field{rule.FieldStatus}, Condition: rule.Condition{Operator: rule.StatusMatches, Field: rule.FieldStatus, Value: rule.Value{String: &status}}, Description: "Fixture only", Metadata: map[string]string{}}}, evaluated)
			if err != nil {
				t.Fatal(err)
			}
			generated, err := report.Generate(ruled)
			if err != nil {
				t.Fatal(err)
			}
			rendered, err := report.RenderText(generated)
			if err != nil {
				t.Fatal(err)
			}
			if generated.Summary.Total == 0 {
				t.Fatal("empty report is not privacy evidence")
			}
			assertServicePrivacy(t, from, to, result, classified, evaluated, ruled, generated, rendered)
		})
	}
}

func TestServiceIdentityInputContract(t *testing.T) {
	if got := serviceItems(t, "alpha.service")[0].ID; got != "systemd-unit-v1:816cca90ec6a90996e4bb15fb01abcfa" {
		t.Fatalf("versioned service identity changed: %s", got)
	}
	for _, output := range []string{"invalid", "alpha.service loaded inactive dead", "alpha.socket loaded active running", "alpha.service loaded active running\nalpha.service loaded active running"} {
		f := Func{Name: "services", Run: collectServices(serviceRunner{output})}
		got := f.Collect(context.Background(), Request{})
		if got.HealthStatus != inventory.Error || len(got.CollectedData.Items) != 0 {
			t.Fatal("invalid evidence accepted")
		}
		assertServicePrivacy(t, got)
	}
	items := serviceItems(t, "Alpha.service", "alpha.service", `worker@a\x2db.service`, `worker@a-b.service`, strings.Repeat("a", 180)+"x.service", strings.Repeat("a", 180)+"y.service")
	seen := map[string]bool{}
	for _, item := range items {
		if seen[item.ID] {
			t.Fatal("case or escaped unit identity collapsed")
		}
		seen[item.ID] = true
	}
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	if _, _, err := collectServices(serviceRunner{"alpha.service loaded active running"})(ctx); !errors.Is(err, context.Canceled) {
		t.Fatal("cancellation lost")
	}
}

func TestServiceHistoricalIdentityBoundary(t *testing.T) {
	current := serviceSnapshot(t, "current", serviceItems(t, "alpha.service", "beta.service", "gamma.service"))
	legacyItems := serviceItems(t, "alpha.service", "beta.service", "gamma.service")
	for i := range legacyItems {
		legacyItems[i].ID = fmt.Sprint(i + 1)
	}
	legacy := serviceSnapshot(t, "legacy", legacyItems)
	root := filepath.Join(t.TempDir(), "store")
	store, err := inventorystore.Open(root, 10)
	if err != nil {
		t.Fatal(err)
	}
	name, err := store.Save(legacy)
	if err != nil {
		t.Fatal(err)
	}
	path := filepath.Join(root, "snapshots", name)
	before, err := os.ReadFile(path)
	if err != nil {
		t.Fatal(err)
	}
	loaded, err := store.Load(name)
	if err != nil {
		t.Fatal(err)
	}
	if loaded.Canonical.Layers[0].Resources[0].ResourceID != "services:1" {
		t.Fatal("legacy identity reinterpreted")
	}
	currentName, err := store.Save(current)
	if err != nil {
		t.Fatal(err)
	}
	definition, err := command.ResolveProfile("report", command.Selection{Source: "store", Store: root, FromSnapshot: name, ToSnapshot: currentName})
	if err != nil {
		t.Fatal(err)
	}
	execution, err := (pipeline.Orchestrator{}).Execute(context.Background(), definition)
	if !errors.Is(err, comparison.ErrServiceIdentityUnavailable) || execution.Complete || len(execution.Stages) != 1 || execution.Stages[0].Stage != command.Compare || len(execution.Diagnostics) != 1 || !strings.Contains(execution.Diagnostics[0], "service identity comparison unavailable") {
		t.Fatalf("historical boundary did not stop report pipeline explicitly: %v, %#v", err, execution)
	}
	// A second corrected snapshot resumes the existing Community report pipeline
	// without migrating, deleting or rewriting the retained ordinal snapshot.
	next := serviceSnapshot(t, "next", serviceItems(t, "alpha.service", "delta.service", "gamma.service"))
	nextName, err := store.Save(next)
	if err != nil {
		t.Fatal(err)
	}
	definition, err = command.ResolveProfile("report", command.Selection{Source: "store", Store: root, FromSnapshot: currentName, ToSnapshot: nextName})
	if err != nil {
		t.Fatal(err)
	}
	execution, err = (pipeline.Orchestrator{}).Execute(context.Background(), definition)
	if err != nil || len(execution.Stages) != 6 || execution.Stages[5].Stage != command.Report {
		t.Fatalf("corrected snapshots did not resume report pipeline: %v", err)
	}
	assertServicePrivacy(t, execution)
	for _, pair := range [][2]inventory.Snapshot{{loaded, current}, {current, loaded}, {loaded, loaded}} {
		result, err := comparison.Compare(pair[0], pair[1], "from", "to")
		if !errors.Is(err, comparison.ErrServiceIdentityUnavailable) || len(result.Changes) != 0 {
			t.Fatalf("historical identity was treated as exact: %v", err)
		}
		assertServicePrivacy(t, err.Error())
	}
	after, err := os.ReadFile(path)
	if err != nil {
		t.Fatal(err)
	}
	if string(before) != string(after) {
		t.Fatal("historical evidence rewritten")
	}
	// Empty retained sets have no ordinal identity to misinterpret.
	empty := serviceSnapshot(t, "empty", nil)
	if _, err := comparison.Compare(empty, current, "empty", "current"); err != nil {
		t.Fatal(err)
	}
	for _, id := range []string{"services:1", "services:systemd-unit-v2:0123456789abcdef0123456789abcdef", "services:systemd-unit-v1:bad", "services:systemd-unit-v1:ABCDEF0123456789ABCDEF0123456789"} {
		s := serviceSnapshot(t, "unsupported", serviceItems(t, "alpha.service"))
		s.Canonical.Layers[0].Resources[0].ResourceID = id
		if _, err := comparison.Compare(s, current, "from", "to"); !errors.Is(err, comparison.ErrServiceIdentityUnavailable) {
			t.Fatalf("unsupported identity accepted: %v", err)
		}
	}
}
