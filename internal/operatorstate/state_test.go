package operatorstate

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"os"
	"path/filepath"
	"syscall"
	"testing"
	"time"

	"quantumwizard.hu/qwsg/internal/command"
	"quantumwizard.hu/qwsg/internal/presentationmodel"
)

func fixtureState(t *testing.T) State {
	t.Helper()
	observed := time.Date(2026, 8, 8, 0, 0, 0, 0, time.UTC)
	execution := command.Execution{SchemaName: command.ExecutionSchema, SchemaVersion: command.SchemaVersion, ID: "execution", CommandID: "definition", PlanID: "plan", Stages: []command.StageResult{}, View: command.View{Rows: []command.ViewRow{}, Groups: []command.ViewGroup{}}, Diagnostics: []string{}, Complete: true}
	overview, err := presentationmodel.Project(presentationmodel.Input{SchemaName: presentationmodel.InputSchema, SchemaVersion: presentationmodel.SchemaVersion, ObservedAt: observed, FreshForNS: int64(time.Hour), Command: &presentationmodel.CommandObservation{ObservedAt: observed, Value: execution}})
	if err != nil {
		t.Fatal(err)
	}
	value, err := Normalize(State{ObservedAt: observed, PublishedAt: observed.Add(time.Second), FreshUntil: observed.Add(time.Hour), Coverage: CoverageInventorySnapshot, Provenance: Provenance{DefinitionID: "definition", ExecutionID: "execution", Profile: "check", Source: "live", Stages: []string{"inventory", "snapshot"}, Reason: PublicationCheck, ApplicationVersion: "test"}, Overview: overview})
	if err != nil {
		t.Fatal(err)
	}
	return value
}

func TestCanonicalIdentityIntegrityAndStrictDecode(t *testing.T) {
	value := fixtureState(t)
	one, _ := MarshalCanonical(value)
	two, _ := MarshalCanonical(value)
	if !bytes.Equal(one, two) {
		t.Fatal("serialization is not deterministic")
	}
	decoded, err := Decode(one)
	if err != nil || decoded.ID != value.ID {
		t.Fatalf("decode: %v", err)
	}
	tampered := append([]byte(nil), one...)
	tampered[len(tampered)/2] ^= 1
	if _, err = Decode(tampered); err == nil {
		t.Fatal("tampering accepted")
	}
	if _, err = Decode(append(one, []byte(`{}`)...)); err == nil {
		t.Fatal("trailing data accepted")
	}
	value.SchemaVersion = "2.0"
	document, _ := json.Marshal(value)
	if _, err = Decode(document); !errors.Is(err, ErrIncompatible) {
		t.Fatalf("version: %v", err)
	}
}

func TestLegacyStateRemainsReadableAfterOverviewVersionChange(t *testing.T) {
	legacy := fixtureState(t)
	legacy.SchemaVersion = LegacySchemaVersion
	legacy.ModelVersion = LegacyModelVersion
	legacy.Overview.SchemaVersion = presentationmodel.LegacySchemaVersion
	legacy.Overview.ModelVersion = presentationmodel.LegacyModelVersion
	legacy.Overview.AttentionSummary = nil
	legacy.Overview.ID = ""
	document, _ := json.Marshal(legacy.Overview)
	sum := sha256.Sum256(append([]byte(presentationmodel.SchemaName+"/"+presentationmodel.LegacySchemaVersion+"\x00"), document...))
	legacy.Overview.ID = hex.EncodeToString(sum[:])
	legacy.ID, legacy.PayloadDigest = "", ""
	legacy, err := NormalizeWithoutClaims(legacy)
	if err != nil {
		t.Fatal(err)
	}
	encoded, err := MarshalCanonical(legacy)
	if err != nil {
		t.Fatal(err)
	}
	decoded, err := Decode(encoded)
	if err != nil || decoded.SchemaVersion != LegacySchemaVersion || decoded.Overview.SchemaVersion != presentationmodel.LegacySchemaVersion {
		t.Fatalf("legacy state rejected: %v", err)
	}
}

func TestFullOperatorEvaluationCoverageIsAcceptedWithoutGuardianClaim(t *testing.T) {
	value := fixtureState(t)
	value.ID, value.PayloadDigest = "", ""
	value.Coverage = CoverageOperatorEvaluation
	value.Provenance.Profile = "observe"
	value.Provenance.Reason = PublicationObserve
	value.Provenance.Stages = []string{"inventory", "snapshot", "compare", "drift", "health", "rule", "policy", "report"}
	full, err := Normalize(value)
	if err != nil {
		t.Fatal(err)
	}
	if full.Overview.Guardian != presentationmodel.GuardianNotObserved {
		t.Fatalf("one-shot evaluation claimed Guardian state: %s", full.Overview.Guardian)
	}
	full.Provenance.Stages[7] = "runtime"
	if err = Validate(full); !errors.Is(err, ErrCorrupt) {
		t.Fatalf("invalid full coverage accepted: %v", err)
	}
}

func TestPrivateAtomicStoreAndLastValidPreservation(t *testing.T) {
	root := filepath.Join(t.TempDir(), "state")
	store, err := Open(root)
	if err != nil {
		t.Fatal(err)
	}
	if _, err = store.Load(); !errors.Is(err, ErrMissing) {
		t.Fatalf("missing: %v", err)
	}
	value := fixtureState(t)
	if err = store.Publish(value); err != nil {
		t.Fatal(err)
	}
	loaded, err := store.Load()
	if err != nil || loaded.ID != value.ID {
		t.Fatalf("load: %v", err)
	}
	if mode := mustMode(t, root); mode != 0o700 {
		t.Fatalf("dir mode %o", mode)
	}
	path := filepath.Join(root, FileName)
	if mode := mustMode(t, path); mode != 0o600 {
		t.Fatalf("file mode %o", mode)
	}
	if err = os.WriteFile(path, []byte("broken"), 0o600); err != nil {
		t.Fatal(err)
	}
	if _, err = store.Load(); !errors.Is(err, ErrCorrupt) {
		t.Fatalf("corrupt: %v", err)
	}
}

func TestPathAndPermissionRejection(t *testing.T) {
	if _, err := Open("relative"); !errors.Is(err, ErrUnsafePath) {
		t.Fatalf("relative: %v", err)
	}
	root := filepath.Join(t.TempDir(), "state")
	store, _ := Open(root)
	if err := store.Publish(fixtureState(t)); err != nil {
		t.Fatal(err)
	}
	if err := os.Chmod(filepath.Join(root, FileName), 0o644); err != nil {
		t.Fatal(err)
	}
	if _, err := store.Load(); !errors.Is(err, ErrPermission) {
		t.Fatalf("mode: %v", err)
	}
}

func TestPublishRejectsExistingNonPrivateDirectory(t *testing.T) {
	root := filepath.Join(t.TempDir(), "state")
	if err := os.Mkdir(root, 0o755); err != nil {
		t.Fatal(err)
	}
	store, _ := Open(root)
	if err := store.Publish(fixtureState(t)); !errors.Is(err, ErrPermission) {
		t.Fatalf("non-private directory accepted: %v", err)
	}
}

func TestEnsurePrivateRootCreatesMissingHierarchyWithoutChangingAncestor(t *testing.T) {
	home := filepath.Join(t.TempDir(), "home")
	if err := os.Mkdir(home, 0o755); err != nil {
		t.Fatal(err)
	}
	before, err := os.Stat(home)
	if err != nil {
		t.Fatal(err)
	}
	root := filepath.Join(home, ".local", "state", "qwsg")
	if err = EnsurePrivateRoot(root); err != nil {
		t.Fatal(err)
	}
	after, err := os.Stat(home)
	if err != nil {
		t.Fatal(err)
	}
	if before.Mode().Perm() != after.Mode().Perm() || before.Sys().(*syscall.Stat_t).Uid != after.Sys().(*syscall.Stat_t).Uid {
		t.Fatal("existing ancestor identity changed")
	}
	if mode := mustMode(t, root); mode != 0o700 {
		t.Fatalf("root mode %o", mode)
	}
	if err = EnsurePrivateRoot(root); err != nil {
		t.Fatalf("repeat ensure: %v", err)
	}
}

func TestEnsurePrivateRootRejectsUnsafeExistingPaths(t *testing.T) {
	base := t.TempDir()
	target := filepath.Join(base, "target")
	if err := os.Mkdir(target, 0o700); err != nil {
		t.Fatal(err)
	}
	link := filepath.Join(base, "link")
	if err := os.Symlink(target, link); err != nil {
		t.Fatal(err)
	}
	if err := EnsurePrivateRoot(filepath.Join(link, "qwsg")); !errors.Is(err, ErrUnsafePath) {
		t.Fatalf("symlink accepted: %v", err)
	}
	permissive := filepath.Join(base, "permissive")
	if err := os.Mkdir(permissive, 0o755); err != nil {
		t.Fatal(err)
	}
	if err := EnsurePrivateRoot(permissive); !errors.Is(err, ErrPermission) {
		t.Fatalf("permissive root accepted: %v", err)
	}
}

func TestPrivateDirRejectsWrongOwnershipWithoutMutation(t *testing.T) {
	root := filepath.Join(t.TempDir(), "state")
	if err := os.Mkdir(root, 0o700); err != nil {
		t.Fatal(err)
	}
	before, err := os.Lstat(root)
	if err != nil {
		t.Fatal(err)
	}
	if err = privateDirForUID(root, os.Geteuid()+1); !errors.Is(err, ErrPermission) {
		t.Fatalf("wrong owner accepted: %v", err)
	}
	after, err := os.Lstat(root)
	if err != nil {
		t.Fatal(err)
	}
	if before.Mode() != after.Mode() || before.Sys().(*syscall.Stat_t).Uid != after.Sys().(*syscall.Stat_t).Uid {
		t.Fatal("ownership rejection mutated the directory")
	}
}

func TestEnsurePrivateRootDoesNotRepairRestrictiveCreationMode(t *testing.T) {
	root := filepath.Join(t.TempDir(), "state")
	previous := syscall.Umask(0o777)
	defer syscall.Umask(previous)
	if err := EnsurePrivateRoot(root); !errors.Is(err, ErrPermission) {
		t.Fatalf("restrictive creation mode accepted or repaired: %v", err)
	}
	info, err := os.Lstat(root)
	if err != nil {
		t.Fatal(err)
	}
	if info.Mode().Perm() != 0 {
		t.Fatalf("restrictive mode was changed: %o", info.Mode().Perm())
	}
}

func TestPublicationFailureWindowsAreDeterministic(t *testing.T) {
	root := filepath.Join(t.TempDir(), "state")
	store, _ := Open(root)
	first := fixtureState(t)
	if err := store.Publish(first); err != nil {
		t.Fatal(err)
	}
	store.beforeRename = func() error { return errors.New("injected pre-rename") }
	if err := store.Publish(first); err == nil {
		t.Fatal("pre-rename failure ignored")
	}
	loaded, err := store.Load()
	if err != nil || loaded.ID != first.ID {
		t.Fatal("pre-rename failure lost last valid state")
	}
	store.beforeRename = nil
	store.beforeDirectorySync = func() error { return errors.New("injected directory-sync") }
	if err := store.Publish(first); err == nil {
		t.Fatal("directory-sync failure ignored")
	}
	loaded, err = store.Load()
	if err != nil || loaded.ID != first.ID {
		t.Fatal("post-rename state is not recoverable")
	}
	matches, err := filepath.Glob(filepath.Join(root, ".current-operator-state-*"))
	if err != nil || len(matches) != 0 {
		t.Fatalf("temporary files remain: %v %v", matches, err)
	}
}
func mustMode(t *testing.T, path string) os.FileMode {
	t.Helper()
	info, err := os.Stat(path)
	if err != nil {
		t.Fatal(err)
	}
	return info.Mode().Perm()
}
