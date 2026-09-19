package main

import (
	"bytes"
	"context"
	"errors"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"reflect"
	"strings"
	"testing"
	"time"

	"quantumwizard.hu/qwsg/internal/automaticupdate"
	"quantumwizard.hu/qwsg/internal/releasediscovery"
	"quantumwizard.hu/qwsg/internal/updatemutation"
)

func mutationCommandFixture(t *testing.T) string {
	t.Helper()
	state := t.TempDir()
	t.Setenv("QWSG_STATE_DIR", state)
	root := filepath.Join(state, "update")
	if err := ensureUpdateRoot(root); err != nil {
		t.Fatal(err)
	}
	oldUID, oldRoot, oldSudo, oldState, oldCtl, oldFetch := updateEffectiveUID, installedQWSGRoot, runSudo, commandState, runSystemctl, updateMetadataFetch
	t.Cleanup(func() {
		updateEffectiveUID, installedQWSGRoot, runSudo, commandState, runSystemctl, updateMetadataFetch = oldUID, oldRoot, oldSudo, oldState, oldCtl, oldFetch
	})
	updateEffectiveUID = func() int { return 1000 }
	installedQWSGRoot = t.TempDir()
	installedPackageFixture(t, installedQWSGRoot, "1.3.1", true)
	runSudo = func(...string) error { t.Fatal("contender reached privileged helper"); return nil }
	commandState = func(string) string { t.Fatal("contender reached service control"); return "no" }
	runSystemctl = func(string) error { t.Fatal("contender changed service"); return nil }
	return root
}

func assertManualMutationConflict(t *testing.T) {
	t.Helper()
	for _, args := range [][]string{nil, {"rollback"}} {
		var out, diagnostic bytes.Buffer
		if code := runUpdate(args, &out, &diagnostic); code != 3 || !strings.Contains(diagnostic.String(), "transaction_conflict") {
			t.Fatalf("%v: code=%d diagnostic=%s", args, code, diagnostic.String())
		}
	}
}

func TestAutomaticOwnsCommonMutationBoundary(t *testing.T) {
	root := mutationCommandFixture(t)
	req, host, _ := automaticFixture(t)
	req.StageParent = root
	before := installationDigest(t, installedQWSGRoot)
	host.onPreflight = func() {
		assertManualMutationConflict(t)
		assertMutationProcessConflict(t, root)
	}
	host.postError = true // Recovery must also stay inside the same lease.
	wrapped := &mutationRecoveryHost{automaticFixtureHost: host, t: t}
	result, err := automaticupdate.Run(context.Background(), req, wrapped)
	if err == nil || result.State != automaticupdate.RollbackSuccess {
		t.Fatalf("%+v %v", result, err)
	}
	if !reflect.DeepEqual(before, installationDigest(t, installedQWSGRoot)) {
		t.Fatal("manual contender mutated installation")
	}
	lease, err := updatemutation.Acquire(root)
	if err != nil {
		t.Fatal(err)
	}
	lease.Close()
}

type mutationRecoveryHost struct {
	*automaticFixtureHost
	t *testing.T
}

func (h *mutationRecoveryHost) Rollback(ctx context.Context) error {
	assertManualMutationConflict(h.t)
	return h.automaticFixtureHost.Rollback(ctx)
}
func (h *mutationRecoveryHost) Validate(ctx context.Context, version string) error {
	assertManualMutationConflict(h.t)
	return h.automaticFixtureHost.Validate(ctx, version)
}

func TestManualOwnersExcludeAllMutationPathsAndReleaseOnFailure(t *testing.T) {
	for _, owner := range []string{"update", "rollback"} {
		t.Run(owner, func(t *testing.T) {
			root := mutationCommandFixture(t)
			req, host, _ := automaticFixture(t)
			req.StageParent = root
			before := installationDigest(t, installedQWSGRoot)
			called := false
			compete := func() {
				called = true
				assertMutationProcessConflict(t, root)
				assertManualMutationConflict(t)
				result, err := automaticupdate.Run(context.Background(), req, host)
				if err == nil || result.FailureCategory != "transaction_conflict" || result.MutationStarted || host.applyCalls != 0 || host.rollbackCalls != 0 {
					t.Fatalf("automatic contender: %+v %v", result, err)
				}
			}
			var diagnostic bytes.Buffer
			if owner == "update" {
				updateMetadataFetch = func(context.Context) ([]byte, error) { compete(); return nil, errors.New("injected metadata failure") }
				if executeUpdate("", "", io.Discard, &diagnostic) != 1 {
					t.Fatal("failure not reported")
				}
			} else {
				record := localUpdateRecord{Schema: "qwsg.update-local/1", Installed: "1.3.1", Previous: "1.3.0", Backup: filepath.Join(updateRollbackRoot, "1000", "test"), UpdatedAt: "2026-09-19T00:00:00Z"}
				if err := saveUpdateRecord(root, record); err != nil {
					t.Fatal(err)
				}
				commandState = func(string) string { return "yes" }
				runSystemctl = func(string) error { compete(); return errors.New("injected stop failure") }
				if runUpdateRollback(io.Discard, &diagnostic) != 1 {
					t.Fatal("failure not reported")
				}
				after, err := loadUpdateRecord(root)
				if err != nil || after != record {
					t.Fatal("contender changed rollback record")
				}
			}
			if !called {
				t.Fatalf("owner did not reach exclusion test: %s", diagnostic.String())
			}
			if !reflect.DeepEqual(before, installationDigest(t, installedQWSGRoot)) {
				t.Fatal("contender mutated installed package")
			}
			// A legitimate transaction can acquire and complete after the failed owner.
			result, err := automaticupdate.Run(context.Background(), req, host)
			if err != nil || result.State != automaticupdate.Success {
				t.Fatalf("lock leaked: %+v %v", result, err)
			}
		})
	}
}

func TestManualUnsafeMutationLockRefused(t *testing.T) {
	root := mutationCommandFixture(t)
	if err := os.Symlink(filepath.Join(root, "missing"), filepath.Join(root, "automatic.lock")); err != nil {
		t.Fatal(err)
	}
	for _, args := range [][]string{nil, {"rollback"}} {
		var diagnostic bytes.Buffer
		if code := runUpdate(args, io.Discard, &diagnostic); code != 1 || !strings.Contains(diagnostic.String(), "lock unavailable") {
			t.Fatalf("code=%d %s", code, diagnostic.String())
		}
	}
}

func TestManualMutationOwnershipThroughHelperAndRecovery(t *testing.T) {
	for _, owner := range []string{"update", "rollback"} {
		t.Run(owner, func(t *testing.T) {
			root := mutationCommandFixture(t)
			req, host, index := automaticFixture(t)
			req.StageParent = root
			oldVerifier, oldBinary := updateVerifier, installedQWSGBinary
			t.Cleanup(func() { updateVerifier, installedQWSGBinary = oldVerifier, oldBinary })
			installedQWSGBinary = filepath.Join(installedQWSGRoot, "usr/local/bin/qwsg")
			updateVerifier = func() (releasediscovery.Verifier, error) { return req.Verifier, nil }
			index.GeneratedAt = time.Now().UTC().Format(time.RFC3339)
			updateMetadataFetch = func(context.Context) ([]byte, error) { return forwardSign(t, index, true), nil }
			// Local authenticated staging uses the same signed fixture without transport.
			forwardWrite(t, req.LocalArchive+".release-index.json", forwardSign(t, index, true), 0600)
			commandState = func(string) string { return "yes" }
			runSystemctl = func(string) error { return nil }
			calls := []string{}
			runSudo = func(args ...string) error {
				calls = append(calls, args[0])
				assertManualMutationConflict(t)
				result, err := automaticupdate.Run(context.Background(), req, host)
				if err == nil || result.FailureCategory != "transaction_conflict" || result.MutationStarted {
					t.Fatalf("helper scope lost lease: %+v %v", result, err)
				}
				return errors.New("injected helper failure")
			}
			var diagnostic bytes.Buffer
			if owner == "update" {
				if code := executeUpdate(req.LocalArchive, req.TargetVersion, io.Discard, &diagnostic); code != 1 {
					t.Fatalf("code %d", code)
				}
				if !reflect.DeepEqual(calls, []string{"privileged-apply", "privileged-rollback"}) {
					t.Fatalf("helper chain %v: %s", calls, diagnostic.String())
				}
			} else {
				record := localUpdateRecord{Schema: "qwsg.update-local/1", Installed: "1.3.1", Previous: "1.3.0", Backup: filepath.Join(updateRollbackRoot, "1000", "test")}
				if err := saveUpdateRecord(root, record); err != nil {
					t.Fatal(err)
				}
				if code := runUpdateRollback(io.Discard, &diagnostic); code != 1 {
					t.Fatalf("code %d", code)
				}
				if !reflect.DeepEqual(calls, []string{"privileged-rollback"}) {
					t.Fatalf("helper chain %v: %s", calls, diagnostic.String())
				}
			}
			lease, err := updatemutation.Acquire(root)
			if err != nil {
				t.Fatal(err)
			}
			lease.Close()
		})
	}
}

// The child exercises real entry points, not a mocked lock, while each parent
// coordinator is paused deterministically inside its owned transaction.
func TestMutationProcessProbe(t *testing.T) {
	root := os.Getenv("QWSG_MUTATION_COMMAND_PROBE")
	if root == "" {
		return
	}
	t.Setenv("QWSG_STATE_DIR", filepath.Dir(root))
	oldUID := updateEffectiveUID
	updateEffectiveUID = func() int { return 1000 }
	defer func() { updateEffectiveUID = oldUID }()
	assertManualMutationConflict(t)
	req, host, _ := automaticFixture(t)
	req.StageParent = root
	result, err := automaticupdate.Run(context.Background(), req, host)
	if err == nil || result.FailureCategory != "transaction_conflict" || result.MutationStarted || host.applyCalls != 0 || host.rollbackCalls != 0 {
		t.Fatalf("process contender: %+v %v", result, err)
	}
}

func assertMutationProcessConflict(t *testing.T, root string) {
	t.Helper()
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	cmd := exec.CommandContext(ctx, os.Args[0], "-test.run=^TestMutationProcessProbe$")
	cmd.Env = append(os.Environ(), "QWSG_MUTATION_COMMAND_PROBE="+root)
	if out, err := cmd.CombinedOutput(); err != nil {
		t.Fatalf("mutation process: %v %s", err, out)
	}
}
