package app

import (
	"context"
	"errors"
	"quantumwizard.hu/qwsg/internal/collector"
	"quantumwizard.hu/qwsg/internal/inventory"
	"testing"
)

func TestPartialSnapshotRemainsValid(t *testing.T) {
	cs := []collector.Collector{collector.Func{Name: "ok", Run: func(context.Context) ([]inventory.Item, []string, error) {
		return []inventory.Item{{ID: "one", Kind: "test", Facts: map[string]inventory.Fact{}}}, nil, nil
	}}, collector.Func{Name: "missing", Run: func(context.Context) ([]inventory.Item, []string, error) { return nil, nil, errors.New("bad evidence") }}}
	s, e := Collect(context.Background(), "test", registry(t, cs))
	if e != nil || s.Status != inventory.Partial || inventory.ExitCode(s.Status) != 2 {
		t.Fatalf("%s %v", s.Status, e)
	}
}
func TestFatalSnapshot(t *testing.T) {
	cs := []collector.Collector{collector.Func{Name: "bad", Run: func(context.Context) ([]inventory.Item, []string, error) { return nil, nil, errors.New("bad") }}}
	s, e := Collect(context.Background(), "test", registry(t, cs))
	if e != nil || inventory.ExitCode(s.Status) != 1 {
		t.Fatalf("%s %v", s.Status, e)
	}
}

func TestCollectBuildsCanonicalInventoryAndLegacyProjection(t *testing.T) {
	cs := []collector.Collector{
		collector.Func{Name: "network", Run: func(context.Context) ([]inventory.Item, []string, error) {
			return []inventory.Item{{ID: "two", Kind: "interface", Facts: map[string]inventory.Fact{}}}, []string{"fixture"}, nil
		}},
		collector.Func{Name: "host", Run: func(context.Context) ([]inventory.Item, []string, error) {
			return []inventory.Item{{ID: "one", Kind: "host", Facts: map[string]inventory.Fact{"instance_id": {Value: "subject", Quality: "derived", Sensitivity: "host_identifying", Provenance: inventory.Provenance{SourceType: "fixture", SourceLabel: "test"}}}}}, []string{"fixture"}, nil
		}},
	}
	snapshot, err := Collect(context.Background(), "test", registry(t, cs))
	if err != nil {
		t.Fatal(err)
	}
	if snapshot.SchemaVersion != "1.0" || len(snapshot.Categories) != 2 || snapshot.Canonical.SchemaName != inventory.CanonicalSchemaName || snapshot.Canonical.SubjectID != "subject" || len(snapshot.Canonical.Layers) != 2 {
		t.Fatalf("unexpected integrated inventory: %#v", snapshot)
	}
	if err := inventory.Validate(snapshot); err != nil {
		t.Fatal(err)
	}
}

func registry(t *testing.T, collectors []collector.Collector) *collector.Registry {
	t.Helper()
	r := collector.NewRegistry()
	for _, c := range collectors {
		if err := r.Register(c); err != nil {
			t.Fatal(err)
		}
	}
	return r
}
