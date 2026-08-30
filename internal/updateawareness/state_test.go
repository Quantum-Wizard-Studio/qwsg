package updateawareness

import (
	"context"
	"errors"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"

	"quantumwizard.hu/qwsg/internal/installation"
	"quantumwizard.hu/qwsg/internal/releasediscovery"
	"quantumwizard.hu/qwsg/internal/update"
)

var testTime = time.Date(2026, 8, 30, 12, 0, 0, 0, time.UTC)

func installed(version string) Installed {
	return Installed{Classification: installation.VerifiedSupported, Version: version}
}

func result(version string, relation update.Relation, compatibility releasediscovery.Compatibility) releasediscovery.CheckResult {
	auth := releasediscovery.AuthenticityEvidence{Scheme: "ed25519", KeyID: "community-test"}
	return releasediscovery.CheckResult{Source: releasediscovery.SourceEvidence{SourceID: "community-release-index", TransportAuthenticated: true, Validators: releasediscovery.Validators{ETag: "\"one\""}}, IndexGeneratedAt: "2026-08-30T11:00:00Z", Authenticity: auth, Evaluation: releasediscovery.Evaluation{InstalledVersion: "1.2.0", Channel: "stable", Platform: "linux-amd64", Release: releasediscovery.Release{Version: version, PublishedAt: "2026-08-30T10:00:00Z", Status: "active"}, Artifact: releasediscovery.Artifact{Name: "qwsg-" + version + "-linux-amd64.tar.gz", SHA256: strings.Repeat("a", 64), Size: 1234}, Relation: relation, Compatibility: compatibility, MigrationID: map[bool]string{true: "compat-1.2.0-to-1.3.0"}[compatibility == releasediscovery.CompatibilitySupported], Authenticity: auth}}
}

func validState(t *testing.T) State {
	t.Helper()
	value, err := NewSuccess(nil, result("1.3.0", update.Newer, releasediscovery.CompatibilitySupported), "community-release-index", "stable", installed("1.2.0"), testTime, 48*time.Hour)
	if err != nil {
		t.Fatal(err)
	}
	return value
}

func TestStateRoundTripIntegrityAndStrictDecoding(t *testing.T) {
	value := validState(t)
	document, err := Marshal(value)
	if err != nil {
		t.Fatal(err)
	}
	decoded, err := Decode(document)
	if err != nil || decoded.Digest != value.Digest {
		t.Fatalf("decoded=%+v err=%v", decoded, err)
	}
	for name, mutated := range map[string][]byte{"tampered": []byte(strings.Replace(string(document), "1.3.0", "1.4.0", 1)), "unknown": []byte(strings.Replace(string(document), "{", "{\"unknown\":true,", 1)), "trailing": append(document, []byte("{}")...)} {
		t.Run(name, func(t *testing.T) {
			if _, err := Decode(mutated); err == nil {
				t.Fatal("invalid state accepted")
			}
		})
	}
}

func TestFailurePreservesAuthenticatedSuccessButIdentityChangeDoesNot(t *testing.T) {
	previous := validState(t)
	failed, err := NewFailure(&previous, "community-release-index", "stable", installed("1.2.0"), testTime.Add(time.Hour), "source_timeout")
	if err != nil || failed.LastSuccess == nil || failed.ConsecutiveFailures != 1 {
		t.Fatalf("failed=%+v err=%v", failed, err)
	}
	again, err := NewFailure(&failed, "community-release-index", "stable", installed("1.2.0"), testTime.Add(2*time.Hour), "source_timeout")
	if err != nil || again.ConsecutiveFailures != 2 {
		t.Fatalf("again=%+v err=%v", again, err)
	}
	changed, err := NewFailure(&again, "community-release-index", "stable", installed("1.3.0"), testTime.Add(3*time.Hour), "installed_identity_changed")
	if err != nil || changed.LastSuccess != nil || changed.Status != Unknown {
		t.Fatalf("changed=%+v err=%v", changed, err)
	}
}

func TestAuthenticatedNotModifiedRequiresMatchingCache(t *testing.T) {
	previous := validState(t)
	notModified := releasediscovery.CheckResult{NotModified: true, Source: releasediscovery.SourceEvidence{SourceID: "community-release-index", TransportAuthenticated: true, Validators: releasediscovery.Validators{ETag: "\"two\""}}}
	refreshed, err := NewSuccess(&previous, notModified, "community-release-index", "stable", installed("1.2.0"), testTime.Add(time.Hour), 48*time.Hour)
	if err != nil || refreshed.LastSuccess.Validators.ETag != "\"two\"" || !refreshed.LastSuccess.ObservedAt.Equal(testTime.Add(time.Hour)) {
		t.Fatalf("refreshed=%+v err=%v", refreshed, err)
	}
	if _, err = NewSuccess(nil, notModified, "community-release-index", "stable", installed("1.2.0"), testTime, 48*time.Hour); err == nil {
		t.Fatal("304 without authenticated cache accepted")
	}
}

func TestStoreAtomicPrivateAndIntegrityBoundaries(t *testing.T) {
	root := filepath.Join(t.TempDir(), "state")
	if err := os.Mkdir(root, 0700); err != nil {
		t.Fatal(err)
	}
	store, err := Open(root)
	if err != nil {
		t.Fatal(err)
	}
	value := validState(t)
	if err = store.Publish(value); err != nil {
		t.Fatal(err)
	}
	path := filepath.Join(root, "update", "awareness.json")
	info, err := os.Stat(path)
	if err != nil || info.Mode().Perm() != 0600 {
		t.Fatalf("info=%v err=%v", info, err)
	}
	loaded, err := store.Load()
	if err != nil || loaded.Digest != value.Digest {
		t.Fatalf("loaded=%+v err=%v", loaded, err)
	}
	old := loaded
	store.beforeRename = func() error { return errors.New("injected") }
	failed, _ := NewFailure(&old, "community-release-index", "stable", installed("1.2.0"), testTime.Add(time.Hour), "source_timeout")
	if store.Publish(failed) == nil {
		t.Fatal("pre-rename failure ignored")
	}
	loaded, err = store.Load()
	if err != nil || loaded.Digest != old.Digest {
		t.Fatal("pre-rename failure lost prior state")
	}
	store.beforeRename = nil
	store.beforeDirectorySync = func() error { return errors.New("injected post rename") }
	if store.Publish(failed) == nil {
		t.Fatal("post-rename uncertainty ignored")
	}
	loaded, err = store.Load()
	if err != nil || loaded.Digest != failed.Digest {
		t.Fatal("post-rename state not recoverable")
	}
}

func TestStoreRejectsUnsafeTargetsAndContendedLock(t *testing.T) {
	root := filepath.Join(t.TempDir(), "state")
	if err := os.Mkdir(root, 0700); err != nil {
		t.Fatal(err)
	}
	store, _ := Open(root)
	if err := store.Publish(validState(t)); err != nil {
		t.Fatal(err)
	}
	lock, err := store.Lock()
	if err != nil {
		t.Fatal(err)
	}
	if _, err = store.Lock(); !errors.Is(err, ErrContended) {
		t.Fatalf("contended err=%v", err)
	}
	if err = lock.Release(); err != nil {
		t.Fatal(err)
	}
	path := filepath.Join(root, "update", "awareness.json")
	if err = os.Remove(path); err != nil {
		t.Fatal(err)
	}
	if err = os.Symlink("missing", path); err != nil {
		t.Fatal(err)
	}
	if _, err = store.Load(); !errors.Is(err, ErrUnsafePath) {
		t.Fatalf("symlink err=%v", err)
	}
	root2 := filepath.Join(t.TempDir(), "state")
	if err = os.Mkdir(root2, 0700); err != nil {
		t.Fatal(err)
	}
	store2, _ := Open(root2)
	if err = store2.Publish(validState(t)); err != nil {
		t.Fatal(err)
	}
	target := filepath.Join(root2, "update", "awareness.json")
	if err = os.Link(target, target+".alias"); err != nil {
		t.Fatal(err)
	}
	if _, err = store2.Load(); !errors.Is(err, ErrPermission) {
		t.Fatalf("hardlink err=%v", err)
	}
}

type fakeChecker struct {
	result   releasediscovery.CheckResult
	err      error
	requests []releasediscovery.FetchRequest
}

func (f *fakeChecker) Check(_ context.Context, r releasediscovery.FetchRequest, _ string, _ bool) (releasediscovery.CheckResult, error) {
	f.requests = append(f.requests, r)
	return f.result, f.err
}

func managerFixture(t *testing.T, checker Checker) Manager {
	t.Helper()
	root := filepath.Join(t.TempDir(), "state")
	if err := os.Mkdir(root, 0700); err != nil {
		t.Fatal(err)
	}
	store, err := Open(root)
	if err != nil {
		t.Fatal(err)
	}
	return Manager{Store: store, Checker: checker, Classify: func() installation.Result {
		return installation.Result{State: installation.VerifiedSupported, Version: "1.2.0"}
	}, SourceID: "community-release-index", Channel: "stable", Platform: "linux-amd64", Freshness: 48 * time.Hour, Now: func() time.Time { return testTime }}
}

func TestManagerSuccessFailure304AndWithdrawal(t *testing.T) {
	checker := &fakeChecker{result: result("1.3.0", update.Newer, releasediscovery.CompatibilitySupported)}
	manager := managerFixture(t, checker)
	state, err := manager.Check(context.Background())
	if err != nil || state.Status != UpdateAvailable {
		t.Fatalf("state=%+v err=%v", state, err)
	}
	checker.err = &releasediscovery.ContractError{Category: releasediscovery.SourceTimeout}
	manager.Now = func() time.Time { return testTime.Add(time.Hour) }
	failed, err := manager.Check(context.Background())
	if releasediscovery.FailureOf(err) != releasediscovery.SourceTimeout || failed.LastSuccess == nil {
		t.Fatalf("failed=%+v err=%v", failed, err)
	}
	checker.err = nil
	checker.result = releasediscovery.CheckResult{NotModified: true, Source: releasediscovery.SourceEvidence{SourceID: "community-release-index", TransportAuthenticated: true, Validators: releasediscovery.Validators{ETag: "\"two\""}}}
	manager.Now = func() time.Time { return testTime.Add(2 * time.Hour) }
	refreshed, err := manager.Check(context.Background())
	if err != nil || refreshed.LastSuccess.Validators.ETag != "\"two\"" {
		t.Fatalf("refreshed=%+v err=%v", refreshed, err)
	}
	checker.result = releasediscovery.CheckResult{Source: releasediscovery.SourceEvidence{SourceID: "community-release-index", TransportAuthenticated: true}, IndexGeneratedAt: "2026-08-30T14:00:00Z", Authenticity: releasediscovery.AuthenticityEvidence{Scheme: "ed25519", KeyID: "community-test"}, WithdrawnVersions: []string{"1.3.0"}}
	checker.err = &releasediscovery.ContractError{Category: releasediscovery.NoEligibleRelease}
	manager.Now = func() time.Time { return testTime.Add(3 * time.Hour) }
	withdrawn, err := manager.Check(context.Background())
	if err != nil || withdrawn.Status != Withdrawn {
		t.Fatalf("withdrawn=%+v err=%v", withdrawn, err)
	}
}

func TestManagerFailsClosedForUnverifiedInstalledIdentityAndMissingAuthority(t *testing.T) {
	manager := managerFixture(t, nil)
	manager.Classify = func() installation.Result {
		return installation.Result{State: installation.LegacyInstallation, Version: "0.0.1-prealpha"}
	}
	if _, err := manager.Check(context.Background()); !errors.Is(err, ErrInstalledIdentity) {
		t.Fatalf("err=%v", err)
	}
	manager = managerFixture(t, nil)
	state, err := manager.Check(context.Background())
	if err == nil || state.LastAttempt.Failure != "source_authority_refused" || state.LastSuccess != nil {
		t.Fatalf("state=%+v err=%v", state, err)
	}
}
