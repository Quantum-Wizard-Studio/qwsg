package main

import (
	"context"
	"encoding/json"
	"errors"
	"io"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"quantumwizard.hu/qwsg/internal/automaticupdate"
	"quantumwizard.hu/qwsg/internal/guardian"
	"quantumwizard.hu/qwsg/internal/updateawareness"
	"quantumwizard.hu/qwsg/internal/updatemutation"
)

type recoveryServiceFixture struct {
	active                                              bool
	startFailure, falseStart, queryFailure, stopFailure bool
	starts, stops, verified                             int
	onCommand                                           func(string)
}

func recoveryService(t *testing.T, active bool) *recoveryServiceFixture {
	t.Helper()
	oldCtl, oldState, oldInactive := automaticSystemctl, guardianServiceState, automaticGuardianInactive
	t.Cleanup(func() {
		automaticSystemctl, guardianServiceState, automaticGuardianInactive = oldCtl, oldState, oldInactive
	})
	guardianServiceState = realGuardianServiceState
	f := &recoveryServiceFixture{active: active}
	automaticSystemctl = func(_ context.Context, args ...string) ([]byte, error) {
		if f.onCommand != nil {
			f.onCommand(args[0])
		}
		switch args[0] {
		case "stop":
			f.stops++
			f.active = false
			if f.stopFailure {
				return nil, errors.New("partial stop failure")
			}
			return nil, nil
		case "start":
			f.starts++
			if f.startFailure {
				return nil, errors.New("start refused")
			}
			f.active = !f.falseStart
			return nil, nil
		case "is-active":
			if f.active {
				return []byte("active"), nil
			}
			return []byte("inactive"), errors.New("inactive")
		case "show":
			if f.queryFailure {
				return nil, errors.New("query failed")
			}
			switch args[2] {
			case "--property=ActiveState":
				if f.active {
					return []byte("active"), nil
				}
				return []byte("inactive"), nil
			case "--property=InvocationID":
				return []byte(strings.Repeat("a", 32)), nil
			case "--property=SubState":
				return []byte("running"), nil
			case "--property=Result":
				return []byte("success"), nil
			case "--property=MainPID":
				f.verified++
				return []byte("42"), nil
			}
		}
		t.Fatalf("unexpected command %v", args)
		return nil, nil
	}
	automaticGuardianInactive = func(context.Context) error {
		if f.active {
			return errors.New("still active")
		}
		return nil
	}
	return f
}

func TestRecoveryHandoffTerminalEvidence(t *testing.T) {
	for _, scenario := range []string{"running", "stopped", "preflight failure", "staging failure", "rollback", "stopped rollback", "rollback failure", "restart failure", "false start", "rollback restart failure", "unknown service"} {
		t.Run(scenario, func(t *testing.T) {
			req, host, _ := automaticFixture(t)
			service := recoveryService(t, !strings.HasPrefix(scenario, "stopped"))
			switch scenario {
			case "preflight failure":
				host.preflightError = true
			case "staging failure":
				req.LocalArchive = "/missing/fixture.tar.gz"
			case "rollback", "stopped rollback":
				host.postError = true
			case "rollback failure":
				host.postError = true
				host.rollbackError = true
			case "restart failure":
				service.startFailure = true
			case "false start":
				service.falseStart = true
			case "rollback restart failure":
				host.postError = true
				service.startFailure = true
			case "unknown service":
				service.queryFailure = true
			}
			state := filepath.Dir(req.StageParent)
			store, err := updateawareness.Open(state)
			if err != nil {
				t.Fatal(err)
			}
			control := &automaticServiceControl{root: state, generation: strings.Repeat("a", 32)}
			control.beforeStop = func(recovery guardian.Recovery) error {
				return recordAutomatic(store, guardian.AutomaticResult{Outcome: "handoff_pending", Recovery: recovery}, true)
			}
			service.onCommand = func(string) {
				lease, err := updatemutation.Acquire(req.StageParent)
				if !errors.Is(err, updatemutation.ErrContended) {
					if lease != nil {
						lease.Close()
					}
					t.Fatalf("service control outside exclusion: %v", err)
				}
			}
			r := guardian.AutomaticHandoff(context.Background(), req, host, control, func(r guardian.AutomaticResult) error {
				lease, err := updatemutation.Acquire(req.StageParent)
				if !errors.Is(err, updatemutation.ErrContended) {
					if lease != nil {
						lease.Close()
					}
					t.Fatalf("evidence outside exclusion: %v", err)
				}
				return recordAutomatic(store, r, true)
			})
			data, err := store.LoadAutomation()
			if err != nil {
				t.Fatal(err)
			}
			var saved guardian.AutomaticResult
			if json.Unmarshal(data, &saved) != nil || saved.Outcome != r.Outcome || saved.Recovery != r.Recovery {
				t.Fatalf("lost evidence %s %+v", data, r)
			}
			want := "success"
			switch scenario {
			case "preflight failure", "staging failure":
				want = "orchestration_failed"
			case "unknown service":
				want = "handoff_failed"
			case "rollback", "stopped rollback":
				want = "orchestration_rolled_back"
			case "rollback failure":
				want = "rollback_failed"
			case "restart failure", "false start", "rollback restart failure":
				want = "recovery_failed"
			}
			if r.Outcome != want {
				t.Fatalf("%+v", r)
			}
			if strings.HasPrefix(scenario, "stopped") || scenario == "staging failure" || scenario == "unknown service" {
				if service.starts != 0 || service.stops != 0 {
					t.Fatalf("unexpected service mutation %+v", service)
				}
			} else {
				if service.stops != 1 {
					t.Fatal("missing stop")
				}
				if scenario == "rollback failure" {
					if service.starts != 0 || r.Recovery.State != "blocked" {
						t.Fatalf("unsafe rollback recovery %+v", r)
					}
				} else if service.starts != 1 {
					t.Fatal("restart count")
				}
			}
			if scenario == "running" && (service.verified != 1 || r.Recovery.State != "verified_running") {
				t.Fatal("recovery unverified")
			}
			if scenario == "preflight failure" && (host.rollbackCalls != 0 || r.Transaction.MutationStarted || !service.active) {
				t.Fatalf("pre-mutation failure %+v", r)
			}
			if strings.Contains(scenario, "rollback") && scenario != "rollback failure" && r.Transaction.RollbackResult != "succeeded" {
				t.Fatalf("package rollback evidence %+v", r)
			}
			if r.RestartFailed && automaticInhibition(store) != "handoff_incomplete" {
				t.Fatal("recovery failure permits retry")
			}
		})
	}
}

func TestWorkerLossHasNoUnconditionalRecovery(t *testing.T) {
	cmd := automaticHandoffCommand(context.Background(), strings.Repeat("a", 32), t.TempDir())
	for _, arg := range cmd.Args {
		if strings.Contains(arg, "ExecStopPost") || strings.Contains(arg, "Restart=") {
			t.Fatalf("unconditional fallback %v", cmd.Args)
		}
	}
	for _, running := range []bool{true, false} {
		t.Run(map[bool]string{true: "running", false: "stopped"}[running], func(t *testing.T) {
			f := recoveryService(t, running)
			root := t.TempDir()
			if err := os.Chmod(root, 0700); err != nil {
				t.Fatal(err)
			}
			store, err := updateawareness.Open(root)
			if err != nil {
				t.Fatal(err)
			}
			control := &automaticServiceControl{root: root, generation: strings.Repeat("a", 32)}
			control.beforeStop = func(r guardian.Recovery) error {
				return recordAutomatic(store, guardian.AutomaticResult{Outcome: "handoff_pending", Recovery: r}, true)
			}
			if err := control.Stop(context.Background()); err != nil {
				t.Fatal(err)
			}
			// Model process loss: close the instance lock, never invoke recovery.
			if control.lock != nil {
				if err := control.lock.Release(); err != nil {
					t.Fatal(err)
				}
			}
			if f.starts != 0 || automaticInhibition(store) != "handoff_incomplete" {
				t.Fatal("worker exit granted restart/retry")
			}
			data, _ := store.LoadAutomation()
			var r guardian.AutomaticResult
			_ = json.Unmarshal(data, &r)
			if r.Recovery.StopRequested != running || r.Recovery.State != "pending" {
				t.Fatalf("intent lost %s", data)
			}
		})
	}
}

func TestManualRollbackRecoveryEvidence(t *testing.T) {
	for _, scenario := range []string{"running", "stopped", "helper failure", "validation failure", "restart failure", "false start", "unknown service"} {
		t.Run(scenario, func(t *testing.T) {
			root := mutationCommandFixture(t)
			_, host, _ := automaticFixture(t)
			oldRoot, oldBinary := installedQWSGRoot, installedQWSGBinary
			t.Cleanup(func() { installedQWSGRoot, installedQWSGBinary = oldRoot, oldBinary })
			installedQWSGRoot = host.dest
			installedQWSGBinary = filepath.Join(host.dest, "usr/local/bin/qwsg")
			for path, body := range forwardInstalledFiles("1.3.1", strings.Repeat("a", 40), []byte("#!/bin/sh\nprintf 'QWSG 1.3.1\\ncommit: "+strings.Repeat("a", 40)+"\\nbuilt: 2026-09-09T00:00:00Z\\n'\n")) {
				forwardWrite(t, filepath.Join(host.dest, path), body, 0644)
			}
			if err := os.Chmod(installedQWSGBinary, 0755); err != nil {
				t.Fatal(err)
			}
			f := recoveryService(t, scenario != "stopped")
			f.startFailure = scenario == "restart failure"
			f.falseStart = scenario == "false start"
			f.queryFailure = scenario == "unknown service"
			calls := 0
			runSudo = func(...string) error {
				calls++
				if scenario == "helper failure" {
					return errors.New("helper failed")
				}
				return nil
			}
			runSystemctl = func(action string) error {
				if action == "daemon-reload" {
					return nil
				}
				_, err := automaticSystemctl(context.Background(), action, "qwsg-guardian.service")
				return err
			}
			previous := "1.3.1"
			if scenario == "validation failure" {
				previous = "9.9.9"
			}
			record := localUpdateRecord{Schema: "qwsg.update-local/1", Installed: "1.3.1", Previous: previous, Backup: filepath.Join(updateRollbackRoot, "1000", "fixture")}
			if err := saveUpdateRecord(root, record); err != nil {
				t.Fatal(err)
			}
			code := runUpdateRollback(io.Discard, io.Discard)
			data, err := os.ReadFile(filepath.Join(root, "rollback-result.json"))
			if err != nil {
				t.Fatal(err)
			}
			var r rollbackEvidence
			if json.Unmarshal(data, &r) != nil {
				t.Fatal("invalid evidence")
			}
			success := scenario == "running" || scenario == "stopped"
			if (code == 0) != success || (r.Outcome == "success") != success {
				t.Fatalf("code=%d %s", code, data)
			}
			if success && r.Package != "validated" {
				t.Fatalf("unvalidated success %s", data)
			}
			if scenario == "running" && (f.starts != 1 || f.verified != 1) {
				t.Fatal("recovery not verified")
			}
			if scenario == "stopped" && (f.starts != 0 || f.stops != 0 || r.Recovery.State != "preserved_stopped") {
				t.Fatal("stopped intent lost")
			}
			if scenario == "helper failure" || scenario == "validation failure" {
				if f.starts != 0 || r.Recovery.State != "blocked" {
					t.Fatalf("unsafe recovery %s", data)
				}
			}
			if scenario == "restart failure" || scenario == "false start" {
				if r.Package != "validated" || r.Outcome != "recovery_failed" {
					t.Fatalf("lost recovery failure %s", data)
				}
			}
			if scenario == "unknown service" && calls != 0 {
				t.Fatal("unknown state allowed rollback")
			}
			_, err = os.Stat(filepath.Join(root, "current.json"))
			if success != os.IsNotExist(err) {
				t.Fatal("rollback metadata discarded on failure")
			}
		})
	}
}

// A panic is an abnormal worker exit too. Even a deferred finalizer must not
// mistake the missing apply receipt for proof that the installation is safe.
type panicRecoveryHost struct{ automaticupdate.Host }

func (h panicRecoveryHost) Apply(context.Context, automaticupdate.PackageInput) (automaticupdate.ApplyResult, error) {
	panic("worker interrupted during apply")
}

func TestInterruptedWorkerAndEvidenceFailure(t *testing.T) {
	for _, scenario := range []string{"panic", "terminal evidence", "intent evidence"} {
		t.Run(scenario, func(t *testing.T) {
			req, host, _ := automaticFixture(t)
			service := recoveryService(t, true)
			store, err := updateawareness.Open(filepath.Dir(req.StageParent))
			if err != nil {
				t.Fatal(err)
			}
			control := &automaticServiceControl{root: filepath.Dir(req.StageParent), generation: strings.Repeat("a", 32)}
			control.beforeStop = func(r guardian.Recovery) error {
				if scenario == "intent evidence" {
					return errors.New("evidence refused")
				}
				return recordAutomatic(store, guardian.AutomaticResult{Outcome: "handoff_pending", Recovery: r}, true)
			}
			var target automaticupdate.Host = host
			if scenario == "panic" {
				target = panicRecoveryHost{host}
			}
			var r guardian.AutomaticResult
			panicked := false
			func() {
				defer func() { panicked = recover() != nil }()
				r = guardian.AutomaticHandoff(context.Background(), req, target, control, func(r guardian.AutomaticResult) error {
					if scenario == "terminal evidence" {
						return errors.New("evidence refused")
					}
					return recordAutomatic(store, r, true)
				})
			}()
			switch scenario {
			case "panic":
				if !panicked || service.starts != 0 || automaticInhibition(store) != "handoff_incomplete" {
					t.Fatal("interrupted apply caused recovery")
				}
				data, err := store.LoadAutomation()
				if err != nil {
					t.Fatal(err)
				}
				var saved guardian.AutomaticResult
				if json.Unmarshal(data, &saved) != nil || saved.Transaction == nil || saved.Transaction.MutationKnown || !saved.Transaction.MutationStarted {
					t.Fatalf("interrupted mutation falsely known: %s", data)
				}
			case "terminal evidence":
				if r.Outcome != "evidence_failed" || automaticInhibition(store) != "handoff_incomplete" {
					t.Fatalf("evidence failure lost %+v", r)
				}
			case "intent evidence":
				if service.starts != 0 || service.stops != 0 || host.applyCalls != 0 {
					t.Fatal("intent persistence failure mutated installation")
				}
			}
			lease, err := updatemutation.Acquire(req.StageParent)
			if err != nil {
				t.Fatal(err)
			}
			lease.Close()
		})
	}
}

func TestRecoveryRequiresCompleteServiceEvidence(t *testing.T) {
	for _, property := range []string{"ActiveState", "SubState", "Result", "MainPID"} {
		t.Run(property, func(t *testing.T) {
			recoveryService(t, true)
			base := automaticSystemctl
			automaticSystemctl = func(ctx context.Context, args ...string) ([]byte, error) {
				if args[0] == "show" && args[2] == "--property="+property {
					return []byte("0"), nil
				}
				return base(ctx, args...)
			}
			r, err := recoverGuardian(context.Background(), guardian.Recovery{Intent: "active", StopRequested: true}, true)
			if err == nil || r.State != "failed" {
				t.Fatalf("accepted incomplete %s evidence: %+v", property, r)
			}
		})
	}
}

type invalidRollbackRecoveryHost struct{ *automaticFixtureHost }

func (h invalidRollbackRecoveryHost) Validate(ctx context.Context, version string) error {
	if version == h.req.SourceVersion {
		return errors.New("rollback configuration invalid")
	}
	return h.automaticFixtureHost.Validate(ctx, version)
}
func TestRollbackValidationFailureBlocksGuardian(t *testing.T) {
	req, host, _ := automaticFixture(t)
	host.postError = true
	f := recoveryService(t, true)
	control := &automaticServiceControl{root: filepath.Dir(req.StageParent), generation: strings.Repeat("a", 32)}
	r := guardian.AutomaticHandoff(context.Background(), req, invalidRollbackRecoveryHost{host}, control)
	if host.rollbackCalls != 1 || r.Outcome != "rollback_failed" || r.Transaction.RollbackResult != "failed" || r.Recovery.State != "blocked" || f.starts != 0 {
		t.Fatalf("unvalidated rollback recovered: %+v", r)
	}
}
