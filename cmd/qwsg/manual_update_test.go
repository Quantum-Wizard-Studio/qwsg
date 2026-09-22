package main

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"io"
	"os"
	"path/filepath"
	"reflect"
	"strings"
	"testing"
	"time"

	"quantumwizard.hu/qwsg/internal/automaticupdate"
	"quantumwizard.hu/qwsg/internal/releasediscovery"
)

func manualFixture(t *testing.T) (string, string, *recoveryServiceFixture, string) {
	t.Helper()
	root := mutationCommandFixture(t)
	req, _, index := automaticFixture(t)
	oldBinary, oldSystem, oldRollback, oldVerifier, oldValidate, oldPersist := installedQWSGBinary, updateSystemRoot, updateRollbackRoot, updateVerifier, manualValidate, manualPersist
	t.Cleanup(func() {
		installedQWSGBinary, updateSystemRoot, updateRollbackRoot, updateVerifier, manualValidate, manualPersist = oldBinary, oldSystem, oldRollback, oldVerifier, oldValidate, oldPersist
	})
	installedQWSGBinary = filepath.Join(installedQWSGRoot, "usr/local/bin/qwsg")
	updateSystemRoot = filepath.Join(t.TempDir(), "system")
	updateRollbackRoot = filepath.Join(updateSystemRoot, "rollback")
	updateVerifier = func() (releasediscovery.Verifier, error) { return req.Verifier, nil }
	commit := strings.Repeat("b", 40)
	binary := []byte("#!/bin/sh\nprintf 'QWSG " + req.TargetVersion + "\\ncommit: " + commit + "\\nbuilt: 2026-09-09T00:00:00Z\\n'\n")
	archive, digest := forwardArchive(t, t.TempDir(), req.TargetVersion, commit, binary)
	forwardWrite(t, archive+".sha256", []byte(digest+"  "+filepath.Base(archive)+"\n"), 0600)
	info, _ := os.Stat(archive)
	a := &index.Channels[0].Releases[0].Artifacts[0]
	a.SHA256 = digest
	a.Size = info.Size()
	a.Name = filepath.Base(archive)
	index.GeneratedAt = time.Now().UTC().Format(time.RFC3339)
	forwardWrite(t, archive+".release-index.json", forwardSign(t, index, true), 0600)
	updateMetadataFetch = func(context.Context) ([]byte, error) { return forwardSign(t, index, true), nil }
	service := recoveryService(t, true)
	runSystemctl = func(string) error { return nil }
	manualApply = func(args ...string) (automaticupdate.ApplyResult, error) {
		old := updateEffectiveUID
		updateEffectiveUID = func() int { return 0 }
		defer func() { updateEffectiveUID = old }()
		var receipt bytes.Buffer
		code := runUpdate(args, &receipt, io.Discard)
		var err error
		if code != 0 {
			err = errors.New("helper failed")
		}
		return decodeApplyReceipt(receipt.Bytes(), err)
	}
	runSudo = func(args ...string) error {
		old := updateEffectiveUID
		updateEffectiveUID = func() int { return 0 }
		defer func() { updateEffectiveUID = old }()
		if runUpdate(args, io.Discard, io.Discard) != 0 {
			return errors.New("helper failed")
		}
		return nil
	}
	return archive, req.TargetVersion, service, root
}

func TestManualC7TransactionMatrix(t *testing.T) {
	cases := []struct{ name, want string }{
		{"success", "success"}, {"mutation failure", "update_failed_recovered"}, {"validation failure", "update_failed_recovered"},
		{"rollback failure", "rollback_failed"}, {"rollback validation failure", "rollback_validation_failed"},
		{"recovery failure", "recovery_failed"}, {"inactive", "success"}, {"ambiguous receipt", "update_failed_recovered"},
		{"terminal evidence failure", "evidence_failed"},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			archive, target, service, root := manualFixture(t)
			if tc.name == "inactive" {
				service.active = false
			}
			baseApply := manualApply
			baseSudo := runSudo
			applyCalls, rollbackCalls := 0, 0
			manualApply = func(args ...string) (automaticupdate.ApplyResult, error) {
				applyCalls++
				if service.active {
					t.Fatal("mutation before inactive Guardian")
				}
				r, err := baseApply(args...)
				if err != nil {
					t.Fatal(err)
				}
				if tc.name == "ambiguous receipt" {
					return automaticupdate.ApplyResult{}, errors.New("lost receipt")
				}
				if tc.name == "mutation failure" || tc.name == "rollback failure" || tc.name == "rollback validation failure" || tc.name == "recovery failure" {
					return r, errors.New("apply returned failure")
				}
				if tc.name == "validation failure" {
					forwardWrite(t, filepath.Join(installedQWSGRoot, "usr/local/lib/systemd/user/qwsg-guardian.service"), []byte("corrupt"), 0644)
				}
				return r, nil
			}
			runSudo = func(args ...string) error {
				if args[0] == "privileged-rollback" {
					rollbackCalls++
					if tc.name == "rollback failure" {
						return errors.New("rollback failed")
					}
				}
				return baseSudo(args...)
			}
			if tc.name == "rollback validation failure" {
				manualValidate = func(v string) error {
					if v == "1.3.1" {
						return errors.New("restored state invalid")
					}
					return validateInstalledVersion(v)
				}
			}
			if tc.name == "recovery failure" {
				service.startFailure = true
			}
			if tc.name == "terminal evidence failure" {
				manualPersist = func(root, name string, v any) error {
					r := v.(manualUpdateEvidence)
					if r.Outcome == "success" {
						return errors.New("sync failure")
					}
					return saveLocalEvidence(root, name, v)
				}
			}
			service.onCommand = func(action string) {
				if action == "start" {
					r, err := loadManualEvidence(root)
					if err != nil {
						t.Fatal(err)
					}
					if r.Validation != "passed" && r.RollbackValidation != "passed" {
						t.Fatal("Guardian started before package validation")
					}
				}
			}
			var out, diagnostic bytes.Buffer
			code := executeUpdate(archive, target, &out, &diagnostic)
			r, err := loadManualEvidence(root)
			if err != nil {
				t.Fatal(err)
			}
			success := tc.want == "success"
			if (code == 0) != success || r.Outcome != tc.want || applyCalls != 1 {
				t.Fatalf("code=%d evidence=%+v output=%s", code, r, diagnostic.String())
			}
			if !success && strings.Contains(out.String(), "updated safely") {
				t.Fatal("false success")
			}
			if tc.name == "inactive" && (service.starts != 0 || service.stops != 0 || r.Recovery.State != "preserved_stopped") {
				t.Fatal("inactive intent changed")
			}
			if r.Outcome == "update_failed_recovered" && (r.Rollback != "succeeded" || r.RollbackValidation != "passed" || r.Recovery.State != "verified_running" || rollbackCalls != 1) {
				t.Fatalf("recovery lost %+v", r)
			}
			if tc.name == "rollback failure" || tc.name == "rollback validation failure" {
				if service.starts != 0 || !r.Intervention || r.Recovery.State != "blocked" {
					t.Fatalf("unsafe recovery %+v", r)
				}
			}
			if tc.name == "recovery failure" && (r.RollbackValidation != "passed" || r.Recovery.State != "failed") {
				t.Fatalf("outcome conflated %+v", r)
			}
			if success && tc.name != "inactive" && service.verified != 1 {
				t.Fatal("runtime not verified")
			}
			if _, err = os.Stat(filepath.Join(r.Backup, "transaction.json")); err != nil {
				t.Fatal("rollback source discarded", err)
			}
			if r.Intervention && manualTransactionReady(root) == nil {
				t.Fatal("incomplete recovery allows update")
			}
			if tc.name == "recovery failure" {
				service.startFailure = false
				if code := runUpdateRollback(io.Discard, &diagnostic); code != 0 {
					t.Fatalf("retry=%d %s", code, diagnostic.String())
				}
				if !service.active {
					t.Fatal("retry lost original active intent")
				}
				if manualTransactionReady(root) != nil {
					t.Fatal("successful explicit recovery not resolved")
				}
			}
		})
	}
}

func TestManualInterruptedMutationPreservesRecoveryIntent(t *testing.T) {
	archive, target, service, root := manualFixture(t)
	base := manualApply
	manualApply = func(args ...string) (automaticupdate.ApplyResult, error) {
		_, err := base(args...)
		if err != nil {
			t.Fatal(err)
		}
		panic("interruption")
	}
	func() {
		defer func() {
			if recover() == nil {
				t.Fatal("missing interruption")
			}
		}()
		executeUpdate(archive, target, io.Discard, io.Discard)
	}()
	r, err := loadManualEvidence(root)
	if err != nil || r.Outcome != "incomplete" || r.Phase != "mutate" || service.starts != 0 || manualTransactionReady(root) == nil {
		t.Fatalf("interruption evidence %+v %v", r, err)
	}
	if runUpdateRollback(io.Discard, io.Discard) != 0 || !service.active {
		t.Fatal("interrupted update not recoverable")
	}
}

func TestManualC7RefusalAndCurrent(t *testing.T) {
	for _, scenario := range []string{"missing authority", "invalid authority", "conflicting authority", "current"} {
		t.Run(scenario, func(t *testing.T) {
			archive, target, service, root := manualFixture(t)
			before := installationDigest(t, installedQWSGRoot)
			manualApply = func(...string) (automaticupdate.ApplyResult, error) {
				t.Fatal("refusal reached mutation")
				return automaticupdate.ApplyResult{}, nil
			}
			switch scenario {
			case "missing authority":
				if err := os.Remove(archive + ".release-index.json"); err != nil {
					t.Fatal(err)
				}
			case "invalid authority":
				forwardWrite(t, archive+".release-index.json", []byte(`{}`), 0600)
			case "conflicting authority":
				data, err := os.ReadFile(archive + ".release-index.json")
				if err != nil {
					t.Fatal(err)
				}
				var index releasediscovery.Index
				if json.Unmarshal(data, &index) != nil {
					t.Fatal("fixture")
				}
				d := index.Channels[0].Releases[0].Compatibility[0]
				d.Capability = "unknown-v1"
				index.Channels[0].Releases[0].Compatibility = append(index.Channels[0].Releases[0].Compatibility, d)
				data, err = json.Marshal(index)
				if err != nil {
					t.Fatal(err)
				}
				forwardWrite(t, archive+".release-index.json", data, 0600)
			case "current":
				installedPackageFixture(t, installedQWSGRoot, target, true)
				before = installationDigest(t, installedQWSGRoot)
				archive, target = "", ""
			}
			var diagnostic bytes.Buffer
			code := executeUpdate(archive, target, io.Discard, &diagnostic)
			if (code == 0) != (scenario == "current") || service.stops != 0 || service.starts != 0 {
				t.Fatalf("code %d %s", code, diagnostic.String())
			}
			if !reflect.DeepEqual(before, installationDigest(t, installedQWSGRoot)) {
				t.Fatal("refusal changed installation")
			}
			data, err := os.ReadFile(filepath.Join(root, "manual-attempt.json"))
			if err != nil {
				t.Fatal(err)
			}
			phase := "authenticate"
			if scenario == "current" {
				phase = "no_update"
			}
			if !bytes.Contains(data, []byte(phase)) {
				t.Fatalf("lost refusal evidence %s", data)
			}
		})
	}
}

func TestManualC7InvalidRollbackAndRepeat(t *testing.T) {
	for _, scenario := range []string{"missing", "invalid", "repeat"} {
		t.Run(scenario, func(t *testing.T) {
			archive, target, _, root := manualFixture(t)
			if executeUpdate(archive, target, io.Discard, io.Discard) != 0 {
				t.Fatal("update fixture failed")
			}
			r, err := loadUpdateRecord(root)
			if err != nil {
				t.Fatal(err)
			}
			before := installationDigest(t, installedQWSGRoot)
			if scenario == "missing" {
				if err = os.Rename(r.Backup, r.Backup+"-preserved"); err != nil {
					t.Fatal(err)
				}
			}
			if scenario == "invalid" {
				forwardWrite(t, filepath.Join(r.Backup, "files/usr/local/share/doc/qwsg/README.md"), []byte("corrupt"), 0600)
			}
			code := runUpdateRollback(io.Discard, io.Discard)
			if scenario == "repeat" {
				if code != 0 {
					t.Fatal("valid rollback failed")
				}
				before = installationDigest(t, installedQWSGRoot)
				if runUpdateRollback(io.Discard, io.Discard) == 0 {
					t.Fatal("missing source reported as success")
				}
			} else if code == 0 {
				t.Fatal("invalid rollback succeeded")
			}
			if !reflect.DeepEqual(before, installationDigest(t, installedQWSGRoot)) {
				t.Fatal("invalid/repeated rollback changed installation")
			}
		})
	}
}

func TestManualRollbackRetryPreservesOriginalRuntimeIntent(t *testing.T) {
	archive, target, service, root := manualFixture(t)
	if executeUpdate(archive, target, io.Discard, io.Discard) != 0 {
		t.Fatal("fixture update failed")
	}
	service.startFailure = true
	if runUpdateRollback(io.Discard, io.Discard) == 0 {
		t.Fatal("restart failure hidden")
	}
	record, err := loadUpdateRecord(root)
	if err != nil {
		t.Fatal("retry metadata lost")
	}
	if _, err = os.Stat(filepath.Join(record.Backup, "transaction.json")); err != nil {
		t.Fatal("retry backup lost")
	}
	service.startFailure = false
	if runUpdateRollback(io.Discard, io.Discard) != 0 || !service.active {
		t.Fatal("retry lost original running intent")
	}
}

func TestManualInterruptionBeforeMutationRecoversWithoutBackup(t *testing.T) {
	archive, target, service, root := manualFixture(t)
	service.onCommand = func(action string) {
		if action == "stop" {
			service.active = false
			panic("interruption during stop")
		}
	}
	func() {
		defer func() {
			if recover() == nil {
				t.Fatal("no interruption")
			}
		}()
		executeUpdate(archive, target, io.Discard, io.Discard)
	}()
	service.onCommand = nil
	runSudo = func(...string) error { t.Fatal("pre-mutation recovery requested package rollback"); return nil }
	if runUpdateRollback(io.Discard, io.Discard) != 0 || !service.active || manualTransactionReady(root) != nil {
		t.Fatal("pre-mutation intent not recovered")
	}
}
