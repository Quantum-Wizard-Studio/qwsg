package inventorystore

import (
	"bytes"
	"encoding/json"
	"errors"
	"os"
	"path/filepath"
	"reflect"
	"strings"
	"testing"
	"time"

	"quantumwizard.hu/qwsg/internal/inventory"
)

func fixtureSnapshot(id string, completed time.Time, partial bool) inventory.Snapshot {
	observed := completed.Add(-time.Second)
	fact := inventory.Fact{
		Value: "fixture", Quality: "observed", Sensitivity: "operational",
		Provenance: inventory.Provenance{SourceType: "fixture", SourceLabel: "synthetic", ObservedAt: observed},
	}
	categories := []inventory.Category{{
		CategoryID: "host", ContractVersion: "1.0", Status: inventory.Available,
		ObservedAt: observed, CompletedAt: completed, FreshUntil: completed.Add(time.Minute),
		CollectorID: "host", PrivilegeUsed: "ordinary-user", SourceSummary: []string{"fixture"},
		Items:  []inventory.Item{{ID: "subject", Kind: "host", Facts: map[string]inventory.Fact{"name": fact}}},
		Errors: []inventory.InventoryError{}, Redactions: []string{"hostnames"},
	}}
	executions := []inventory.CollectorExecution{{
		CollectorName: "host", Version: "1", Capability: "host", SupportedPlatforms: []string{"linux"},
		Timestamp: observed, Status: inventory.Available, Warnings: []inventory.InventoryWarning{},
		Errors: []inventory.InventoryError{}, Metadata: map[string]string{},
	}}
	if partial {
		categories = append(categories, inventory.Category{
			CategoryID: "network", ContractVersion: "1.0", Status: inventory.PermissionDenied,
			ObservedAt: observed, CompletedAt: completed, FreshUntil: completed.Add(time.Minute),
			CollectorID: "network", PrivilegeUsed: "ordinary-user", SourceSummary: []string{},
			Items: []inventory.Item{}, Errors: []inventory.InventoryError{{Code: "permission_denied", Class: "permission", MessageKey: "inventory.permission_denied", OccurredAt: completed}},
			Redactions: []string{},
		})
		executions = append(executions, inventory.CollectorExecution{
			CollectorName: "network", Version: "1", Capability: "network", SupportedPlatforms: []string{"linux"},
			Timestamp: observed, Status: inventory.PermissionDenied, Warnings: []inventory.InventoryWarning{},
			Errors: categories[1].Errors, Metadata: map[string]string{},
		})
	}
	producer := inventory.Producer{ToolVersion: "test", ContractVersion: "1.0"}
	snapshot := inventory.Snapshot{
		SchemaVersion: inventory.SchemaVersion, SnapshotID: id, RequestID: id, InstanceID: "subject",
		ObservedAt: observed, CompletedAt: completed, FreshUntil: completed.Add(time.Minute),
		DurationMS: 1000, Status: inventory.Aggregate(categories), Categories: categories,
		Errors: []inventory.InventoryError{}, Redactions: []string{"hostnames"}, Producer: producer,
	}
	snapshot.Canonical = inventory.AssembleSystemInventory(
		categories, executions, id, id, "subject", observed, completed,
		completed.Add(time.Minute), 1000, producer,
	)
	return snapshot
}

func openFixtureStore(t *testing.T, retention int) *Store {
	t.Helper()
	store, err := Open(filepath.Join(t.TempDir(), "inventory-store"), retention)
	if err != nil {
		t.Fatal(err)
	}
	return store
}

func TestRoundTripDeterministicBytesAndPermissions(t *testing.T) {
	now := time.Date(2026, 7, 24, 1, 2, 3, 4, time.UTC)
	snapshot := fixtureSnapshot("snapshot-one", now, false)
	first := openFixtureStore(t, 3)
	second := openFixtureStore(t, 3)
	firstName, err := first.Save(snapshot)
	if err != nil {
		t.Fatal(err)
	}
	secondName, err := second.Save(snapshot)
	if err != nil {
		t.Fatal(err)
	}
	if firstName != secondName {
		t.Fatalf("non-deterministic names: %s != %s", firstName, secondName)
	}
	firstBytes, err := os.ReadFile(filepath.Join(first.root, snapshotsDir, firstName))
	if err != nil {
		t.Fatal(err)
	}
	secondBytes, err := os.ReadFile(filepath.Join(second.root, snapshotsDir, secondName))
	if err != nil {
		t.Fatal(err)
	}
	if !bytes.Equal(firstBytes, secondBytes) {
		t.Fatal("persisted bytes differ for identical snapshots")
	}
	loaded, loadedName, err := first.LoadLatest()
	if err != nil {
		t.Fatal(err)
	}
	if loadedName != firstName || !reflect.DeepEqual(snapshot, loaded) {
		t.Fatalf("round trip mismatch: %s %#v", loadedName, loaded)
	}
	for _, path := range []string{first.root, filepath.Join(first.root, snapshotsDir)} {
		info, err := os.Stat(path)
		if err != nil || info.Mode().Perm() != 0o700 {
			t.Fatalf("directory mode for %s: %v %v", path, info, err)
		}
	}
	for _, path := range []string{filepath.Join(first.root, metadataName), filepath.Join(first.root, snapshotsDir, firstName)} {
		info, err := os.Stat(path)
		if err != nil || info.Mode().Perm() != 0o600 {
			t.Fatalf("file mode for %s: %v %v", path, info, err)
		}
	}
}

func TestPartialSnapshotAndRetention(t *testing.T) {
	store := openFixtureStore(t, 2)
	base := time.Date(2026, 7, 24, 2, 0, 0, 0, time.UTC)
	var names []string
	for i, id := range []string{"one", "two", "three"} {
		name, err := store.Save(fixtureSnapshot(id, base.Add(time.Duration(i)*time.Second), i == 2))
		if err != nil {
			t.Fatal(err)
		}
		names = append(names, name)
	}
	got, err := store.List()
	if err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(got, names[1:]) {
		t.Fatalf("retention mismatch: %v", got)
	}
	loaded, _, err := store.LoadLatest()
	if err != nil {
		t.Fatal(err)
	}
	if loaded.Status != inventory.Partial || len(loaded.Errors) != 0 || len(loaded.Redactions) == 0 {
		t.Fatalf("partial semantics lost: %#v", loaded)
	}
}

func TestInvalidAndSecretSnapshotsAreRejected(t *testing.T) {
	store := openFixtureStore(t, 2)
	invalid := fixtureSnapshot("invalid", time.Now().UTC(), false)
	invalid.SchemaVersion = "2.0"
	if _, err := store.Save(invalid); err == nil {
		t.Fatal("unsupported inventory schema persisted")
	}
	secret := fixtureSnapshot("secret", time.Now().UTC(), false)
	secret.Canonical.Layers[0].Resources[0].Facts["secret"] = inventory.CanonicalFact{
		Value: "not-a-real-secret", ValueType: "string", Quality: "observed",
		Sensitivity: "secret_prohibited", Provenance: inventory.Provenance{SourceType: "fixture"},
	}
	if _, err := store.Save(secret); err == nil {
		t.Fatal("secret_prohibited fact persisted")
	}
	legacySecret := fixtureSnapshot("legacy-secret", time.Now().UTC(), false)
	legacySecret.Categories[0].Items[0].Facts["secret"] = inventory.Fact{
		Value: "not-a-real-secret", Quality: "observed", Sensitivity: "secret_prohibited",
		Provenance: inventory.Provenance{SourceType: "fixture"},
	}
	if _, err := store.Save(legacySecret); err == nil {
		t.Fatal("legacy secret_prohibited fact persisted")
	}
	if _, err := store.List(); !errors.Is(err, os.ErrNotExist) {
		t.Fatalf("invalid save created a store: %v", err)
	}
}

func TestAtomicFailurePreservesPreviousSnapshot(t *testing.T) {
	store := openFixtureStore(t, 1)
	base := time.Date(2026, 7, 24, 3, 0, 0, 0, time.UTC)
	old := fixtureSnapshot("old", base, false)
	oldName, err := store.Save(old)
	if err != nil {
		t.Fatal(err)
	}
	store.beforeInstall = func() error { return errors.New("injected failure") }
	if _, err := store.Save(fixtureSnapshot("new", base.Add(time.Second), false)); err == nil {
		t.Fatal("injected failure did not fail")
	}
	store.beforeInstall = nil
	names, err := store.List()
	if err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(names, []string{oldName}) {
		t.Fatalf("old snapshot not preserved: %v", names)
	}
	loaded, _, err := store.LoadLatest()
	if err != nil || loaded.SnapshotID != "old" {
		t.Fatalf("old snapshot unreadable: %#v %v", loaded, err)
	}
}

func TestCorruptionCompatibilityAndDuplicateKeysFailClosed(t *testing.T) {
	cases := []struct {
		name   string
		mutate func(t *testing.T, path string)
		target error
	}{
		{"integrity", func(t *testing.T, path string) {
			document, err := os.ReadFile(path)
			if err != nil {
				t.Fatal(err)
			}
			var envelope Envelope
			if err := json.Unmarshal(document, &envelope); err != nil {
				t.Fatal(err)
			}
			envelope.PayloadSHA256 = strings.Repeat("0", 64)
			document, _ = json.MarshalIndent(envelope, "", "  ")
			if err := os.WriteFile(path, append(document, '\n'), 0o600); err != nil {
				t.Fatal(err)
			}
		}, ErrCorrupt},
		{"incompatible", func(t *testing.T, path string) {
			document, err := os.ReadFile(path)
			if err != nil {
				t.Fatal(err)
			}
			document = bytes.Replace(document, []byte(`"format_version": "1.0"`), []byte(`"format_version": "2.0"`), 1)
			if err := os.WriteFile(path, document, 0o600); err != nil {
				t.Fatal(err)
			}
		}, ErrIncompatible},
		{"duplicate", func(t *testing.T, path string) {
			if err := os.WriteFile(path, []byte(`{"format_name":"qwsg.digital-twin","format_name":"duplicate"}`), 0o600); err != nil {
				t.Fatal(err)
			}
		}, ErrCorrupt},
		{"truncated", func(t *testing.T, path string) {
			if err := os.WriteFile(path, []byte(`{"format_name":`), 0o600); err != nil {
				t.Fatal(err)
			}
		}, ErrCorrupt},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			store := openFixtureStore(t, 2)
			name, err := store.Save(fixtureSnapshot("snapshot-"+tc.name, time.Now().UTC(), false))
			if err != nil {
				t.Fatal(err)
			}
			tc.mutate(t, filepath.Join(store.root, snapshotsDir, name))
			if _, err := store.Load(name); !errors.Is(err, tc.target) {
				t.Fatalf("got %v, want %v", err, tc.target)
			}
		})
	}
}

func TestUnsafePathsFilesAndConflictsAreRejected(t *testing.T) {
	if _, err := Open("relative", 1); !errors.Is(err, ErrUnsafePath) {
		t.Fatalf("relative root accepted: %v", err)
	}
	if _, err := Open(t.TempDir()+"/../unclean", 1); !errors.Is(err, ErrUnsafePath) {
		t.Fatalf("unclean root accepted: %v", err)
	}
	target := filepath.Join(t.TempDir(), "real")
	if err := os.Mkdir(target, 0o700); err != nil {
		t.Fatal(err)
	}
	link := filepath.Join(t.TempDir(), "link")
	if err := os.Symlink(target, link); err != nil {
		t.Fatal(err)
	}
	store, err := Open(link, 1)
	if !errors.Is(err, ErrUnsafePath) {
		t.Fatalf("symlink root accepted: store=%v err=%v", store, err)
	}
	nestedLink := filepath.Join(t.TempDir(), "parent-link")
	if err := os.Symlink(t.TempDir(), nestedLink); err != nil {
		t.Fatal(err)
	}
	if _, err := Open(filepath.Join(nestedLink, "store"), 1); !errors.Is(err, ErrUnsafePath) {
		t.Fatalf("symlink parent accepted: %v", err)
	}

	store = openFixtureStore(t, 2)
	snapshot := fixtureSnapshot("same", time.Now().UTC(), false)
	name, err := store.Save(snapshot)
	if err != nil {
		t.Fatal(err)
	}
	if _, err := store.Save(snapshot); !errors.Is(err, ErrConflict) {
		t.Fatalf("duplicate snapshot accepted: %v", err)
	}
	if _, err := store.Load("../" + name); !errors.Is(err, ErrUnsafePath) {
		t.Fatalf("path traversal accepted: %v", err)
	}
	if err := os.Chmod(filepath.Join(store.root, snapshotsDir, name), 0o644); err != nil {
		t.Fatal(err)
	}
	if _, err := store.Load(name); !errors.Is(err, ErrUnsafePath) {
		t.Fatalf("permissive snapshot accepted: %v", err)
	}
}

func TestRenamedSnapshotIsRejected(t *testing.T) {
	store := openFixtureStore(t, 2)
	name, err := store.Save(fixtureSnapshot("named", time.Now().UTC(), false))
	if err != nil {
		t.Fatal(err)
	}
	renamed := "20260724T000000.000000000Z_0000000000000000.json"
	if err := os.Rename(
		filepath.Join(store.root, snapshotsDir, name),
		filepath.Join(store.root, snapshotsDir, renamed),
	); err != nil {
		t.Fatal(err)
	}
	if _, err := store.Load(renamed); !errors.Is(err, ErrCorrupt) {
		t.Fatalf("renamed snapshot accepted: %v", err)
	}
}

func TestWriteLockAndRetentionConfiguration(t *testing.T) {
	store := openFixtureStore(t, 2)
	if err := store.ensureLayout(); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(store.root, lockName), nil, 0o600); err != nil {
		t.Fatal(err)
	}
	if _, err := store.Save(fixtureSnapshot("locked", time.Now().UTC(), false)); err == nil {
		t.Fatal("concurrent write lock ignored")
	}
	if err := os.Remove(filepath.Join(store.root, lockName)); err != nil {
		t.Fatal(err)
	}
	other, err := Open(store.root, 3)
	if err != nil {
		t.Fatal(err)
	}
	if _, err := other.List(); err == nil {
		t.Fatal("retention mismatch accepted")
	}
}
