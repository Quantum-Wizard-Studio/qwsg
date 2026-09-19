package main

import (
	"context"
	"encoding/json"
	"errors"
	"os"
	"os/exec"
	"path/filepath"
	"reflect"
	"strings"
	"syscall"
	"testing"
	"time"

	"quantumwizard.hu/qwsg/internal/automaticupdate"
	"quantumwizard.hu/qwsg/internal/guardian"
	"quantumwizard.hu/qwsg/internal/installation"
	"quantumwizard.hu/qwsg/internal/productcapability"
	"quantumwizard.hu/qwsg/internal/releasediscovery"
	"quantumwizard.hu/qwsg/internal/updateawareness"
	"quantumwizard.hu/qwsg/internal/updatepolicy"
)

func triggerAwareness(t *testing.T, req automaticupdate.Request) (updateawareness.State, error) {
	t.Helper()
	index, err := releasediscovery.Parse(req.Metadata)
	if err != nil {
		return updateawareness.State{}, err
	}
	auth, err := req.Verifier.Verify(index)
	if err != nil {
		return updateawareness.State{}, err
	}
	evaluation, err := req.Evaluator.Evaluate(auth, "stable", req.Platform, false)
	if err != nil {
		return updateawareness.State{}, err
	}
	return updateawareness.NewSuccess(nil, releasediscovery.CheckResult{Source: releasediscovery.SourceEvidence{SourceID: "fixture", TransportAuthenticated: true}, IndexGeneratedAt: index.GeneratedAt, Authenticity: auth.Authenticity(), Evaluation: evaluation}, "fixture", "stable", updateawareness.Installed{Classification: installation.VerifiedSupported, Version: req.SourceVersion}, time.Now().UTC(), time.Hour)
}

type triggerControl struct {
	stop, resume           int
	stopError, resumeError bool
	beforeStop             func()
	afterResume            func()
}

func (c *triggerControl) Stop(context.Context) error {
	c.stop++
	if c.beforeStop != nil {
		c.beforeStop()
	}
	if c.stopError {
		return errors.New("stop failed")
	}
	return nil
}
func (c *triggerControl) Recover(_ context.Context, safe bool) (guardian.Recovery, error) {
	if !safe {
		return guardian.Recovery{Intent: "active", State: "blocked", StopRequested: true}, nil
	}
	c.resume++
	if c.afterResume != nil {
		c.afterResume()
	}
	if c.resumeError {
		return guardian.Recovery{Intent: "active", State: "failed", StopRequested: true}, errors.New("restart failed")
	}
	return guardian.Recovery{Intent: "active", State: "verified_running", StopRequested: true}, nil
}

func TestAutomaticTriggerGatesAndNoOp(t *testing.T) {
	for _, scenario := range []string{"community", "pro manual", "current", "supported", "unsigned", "unsupported migration", "handoff failure", "failed check with old success"} {
		t.Run(scenario, func(t *testing.T) {
			req, host, index := automaticFixture(t)
			want := "not_authorized"
			switch scenario {
			case "community":
				req.Capabilities, _ = productcapability.Resolve(nil)
				req.Policy.Mode = updatepolicy.Manual
			case "pro manual":
				req.Policy.Mode = updatepolicy.Manual
			case "current":
				req.SourceVersion = req.TargetVersion
				req.Evaluator, _ = releasediscovery.NewEvaluator(func(string) installation.Result {
					return installation.Result{State: installation.VerifiedSupported, Version: req.SourceVersion}
				})
				want = "no_update"
			case "supported":
				want = "handoff_requested"
			case "unsigned":
				req.Metadata = forwardSign(t, index, false)
				want = "candidate_rejected"
			case "unsupported migration":
				index.Channels[0].Releases[0].Compatibility[0].Capability = "unknown-v9"
				req.Metadata = forwardSign(t, index, true)
				want = "candidate_rejected"
			case "handoff failure":
				want = "handoff_failed"
			case "failed check with old success":
				want = "candidate_rejected"
			}
			state, err := triggerAwareness(t, req)
			if scenario == "failed check with old success" {
				err = errors.New("authority failed")
			}
			calls := 0
			r := guardian.AutomaticDecision(context.Background(), req.Capabilities, req.Policy, state, err, func(context.Context) error {
				calls++
				if scenario == "handoff failure" {
					return errors.New("launch failed")
				}
				return nil
			})
			if r.Outcome != want || host.applyCalls != 0 || r.Invoked {
				t.Fatalf("%+v calls=%d", r, calls)
			}
			expectedCalls := 0
			if scenario == "supported" || scenario == "handoff failure" {
				expectedCalls = 1
			}
			if calls != expectedCalls {
				t.Fatalf("handoffs %d", calls)
			}
		})
	}
}

func TestAutomaticHandoffSecurityResultsAndRecovery(t *testing.T) {
	for _, scenario := range []string{"community", "manual", "current", "unsigned", "unsupported", "success", "stop failure", "concurrent", "rollback", "rollback failure", "restart failure"} {
		t.Run(scenario, func(t *testing.T) {
			req, host, index := automaticFixture(t)
			before := installationDigest(t, host.dest)
			control := &triggerControl{}
			want := "success"
			wantApply := 1
			wantStop := 1
			switch scenario {
			case "community":
				req.Capabilities, _ = productcapability.Resolve(nil)
				want = "not_authorized"
				wantApply = 0
				wantStop = 0
			case "manual":
				req.Policy.Mode = updatepolicy.Manual
				want = "not_authorized"
				wantApply = 0
				wantStop = 0
			case "current":
				req.SourceVersion = req.TargetVersion
				req.Evaluator, _ = releasediscovery.NewEvaluator(func(string) installation.Result {
					return installation.Result{State: installation.VerifiedSupported, Version: req.SourceVersion}
				})
				want = "no_update"
				wantApply = 0
				wantStop = 0
			case "unsigned":
				req.Metadata = forwardSign(t, index, false)
				want = "candidate_rejected"
				wantApply = 0
				wantStop = 0
			case "unsupported":
				index.Channels[0].Releases[0].Compatibility[0].Capability = "unknown-v9"
				req.Metadata = forwardSign(t, index, true)
				want = "candidate_rejected"
				wantApply = 0
				wantStop = 0
			case "stop failure":
				control.stopError = true
				want = "handoff_failed"
				wantApply = 0
			case "concurrent":
				f, err := os.OpenFile(filepath.Join(req.StageParent, "automatic.lock"), os.O_CREATE|os.O_RDWR, 0600)
				if err != nil {
					t.Fatal(err)
				}
				defer f.Close()
				if err = syscall.Flock(int(f.Fd()), syscall.LOCK_EX|syscall.LOCK_NB); err != nil {
					t.Fatal(err)
				}
				want = "orchestration_failed"
				wantApply = 0
				wantStop = 0
			case "rollback":
				host.postError = true
				want = "orchestration_rolled_back"
			case "rollback failure":
				host.postError = true
				host.rollbackError = true
				want = "rollback_failed"
			case "restart failure":
				control.resumeError = true
				want = "recovery_failed"
			}
			host.onPreflight = func() {
				if control.stop != 1 || control.resume != 0 {
					t.Fatal("apply outside stopped interval")
				}
			}
			r := guardian.AutomaticHandoff(context.Background(), req, host, control)
			if r.Outcome != want || host.applyCalls != wantApply || control.stop != wantStop || control.resume != func() int {
				if scenario == "rollback failure" {
					return 0
				}
				return wantStop
			}() {
				t.Fatalf("%+v apply=%d stop/resume=%d/%d", r, host.applyCalls, control.stop, control.resume)
			}
			if wantApply == 0 || scenario == "rollback" {
				if !reflect.DeepEqual(before, installationDigest(t, host.dest)) {
					t.Fatal("unexpected mutation")
				}
			}
			if scenario == "concurrent" && (r.Transaction == nil || r.Transaction.FailureCategory != "transaction_conflict") {
				t.Fatalf("%+v", r)
			}
			if scenario == "rollback" && (r.Transaction.RollbackResult != "succeeded" || !r.Transaction.MutationStarted) {
				t.Fatalf("%+v", r)
			}
			if scenario == "rollback failure" && r.Transaction.State != automaticupdate.RollbackFailure {
				t.Fatalf("%+v", r)
			}
			if scenario == "success" {
				req.SourceVersion = req.TargetVersion
				req.Evaluator, _ = releasediscovery.NewEvaluator(func(string) installation.Result {
					return installation.Result{State: installation.VerifiedSupported, Version: req.SourceVersion}
				})
				next := guardian.AutomaticHandoff(context.Background(), req, host, control)
				if next.Outcome != "no_update" || host.applyCalls != 1 || control.stop != 1 {
					t.Fatalf("duplicate transaction: %+v", next)
				}
			}
		})
	}
}

// Separate process acceptance runs the existing Guardian release scheduler,
// signed authority, Task 084 and real package mutation against isolated roots.
// Only service control and privilege are fixtures; nothing touches systemd/live
// installation. Async handoff must wait for the running scheduler to exit.
func TestAutomaticGuardianBlackBox(t *testing.T) {
	if product := os.Getenv("QWSG_TRIGGER_TEST_PRODUCT"); product != "" {
		req, host, _ := automaticFixture(t)
		if product == "community" {
			req.Capabilities, _ = productcapability.Resolve(nil)
			req.Policy.Mode = updatepolicy.Manual
		}
		stateRoot := filepath.Join(t.TempDir(), "state")
		if err := os.Mkdir(stateRoot, 0700); err != nil {
			t.Fatal(err)
		}
		store, err := updateawareness.Open(stateRoot)
		if err != nil {
			t.Fatal(err)
		}
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()
		schedulerDone := make(chan struct{})
		request := make(chan struct{}, 1)
		ready := make(chan struct{})
		close(ready)
		decisionDone := make(chan guardian.AutomaticResult, 1)
		service := guardian.ReleaseCheckService{Store: store, Ready: ready, Interval: time.Hour, Timeout: time.Minute, Check: func(ctx context.Context) error {
			state, checkErr := triggerAwareness(t, req)
			if checkErr != nil {
				return checkErr
			}
			if err := store.Publish(state); err != nil {
				return err
			}
			r := guardian.AutomaticDecision(ctx, req.Capabilities, req.Policy, state, nil, func(context.Context) error { request <- struct{}{}; return nil })
			if err := recordAutomatic(store, r, false); err != nil {
				return err
			}
			decisionDone <- r
			return nil
		}, Wait: func(ctx context.Context, _ time.Time) error { <-ctx.Done(); return ctx.Err() }}
		go func() { defer close(schedulerDone); service.Run(ctx) }()
		var decision guardian.AutomaticResult
		select {
		case decision = <-decisionDone:
		case <-time.After(5 * time.Second):
			t.Fatal("scheduler did not produce decision")
		}
		if product == "community" {
			if decision.Outcome != "not_authorized" || len(request) != 0 || host.applyCalls != 0 {
				t.Fatal("Community applied")
			}
			cancel()
			<-schedulerDone
			t.Log("community: running Guardian scheduler, same authenticated future candidate, no handoff or mutation")
			return
		}
		<-request
		active := true
		control := &triggerControl{beforeStop: func() { cancel(); <-schedulerDone; active = false }, afterResume: func() { active = true }}
		host.onPreflight = func() {
			if active {
				t.Fatal("Guardian still active during Task 084")
			}
		}
		result := guardian.AutomaticHandoff(context.Background(), req, host, control)
		if err := recordAutomatic(store, result, true); err != nil {
			t.Fatal(err)
		}
		if result.Outcome != "success" || !active || host.applyCalls != 1 {
			t.Fatalf("%+v", result)
		}
		data, err := os.ReadFile(filepath.Join(stateRoot, "update", "automatic-result.json"))
		if err != nil {
			t.Fatal(err)
		}
		var saved guardian.AutomaticResult
		if json.Unmarshal(data, &saved) != nil || saved.Transaction == nil || saved.Transaction.State != automaticupdate.Success {
			t.Fatalf("terminal receipt %s", data)
		}
		// A resumed scheduler reclassifies the installed target and does not apply it again.
		req.SourceVersion = req.TargetVersion
		req.Evaluator, _ = releasediscovery.NewEvaluator(func(string) installation.Result {
			return installation.Result{State: installation.VerifiedSupported, Version: req.SourceVersion}
		})
		state, err := triggerAwareness(t, req)
		if err != nil {
			t.Fatal(err)
		}
		next := guardian.AutomaticDecision(context.Background(), req.Capabilities, req.Policy, state, nil, func(context.Context) error { t.Fatal("duplicate handoff"); return nil })
		if next.Outcome != "no_update" {
			t.Fatalf("%+v", next)
		}
		if err = recordAutomatic(store, next, false); err != nil {
			t.Fatal(err)
		}
		retained, _ := os.ReadFile(filepath.Join(stateRoot, "update", "automatic-result.json"))
		if string(retained) != string(data) {
			t.Fatal("no-op lost transaction receipt")
		}
		t.Log("pro: running Guardian scheduler -> signed supported candidate -> async stop/join -> Task 084 real apply exactly once -> persisted success -> resumed current no-op")
		return
	}
	for _, product := range []string{"community", "pro"} {
		t.Run(product, func(t *testing.T) {
			cmd := exec.Command(os.Args[0], "-test.run=^TestAutomaticGuardianBlackBox$", "-test.v")
			cmd.Env = append(os.Environ(), "QWSG_TRIGGER_TEST_PRODUCT="+product)
			out, err := cmd.CombinedOutput()
			if err != nil {
				t.Fatalf("black-box %v\n%s", err, out)
			}
			if !strings.Contains(string(out), product+":") {
				t.Fatalf("missing acceptance evidence: %s", out)
			}
			t.Log(string(out))
		})
	}
}

func TestAutomaticServiceHandoffGenerationAndLock(t *testing.T) {
	originalCommand, originalInactive, originalState := automaticSystemctl, automaticGuardianInactive, guardianServiceState
	defer func() {
		automaticSystemctl = originalCommand
		automaticGuardianInactive = originalInactive
		guardianServiceState = originalState
	}()
	guardianServiceState = func(context.Context) (string, error) { return "active", nil }
	generation := strings.Repeat("a", 32)
	for _, scenario := range []string{"normal", "changed generation", "stop failure", "still active"} {
		t.Run(scenario, func(t *testing.T) {
			root := t.TempDir()
			var operations []string
			inactive := false
			automaticSystemctl = func(_ context.Context, args ...string) ([]byte, error) {
				operations = append(operations, args[0])
				switch args[0] {
				case "show":
					switch args[2] {
					case "--property=ActiveState":
						return []byte("active"), nil
					case "--property=SubState":
						return []byte("running"), nil
					case "--property=Result":
						return []byte("success"), nil
					case "--property=MainPID":
						return []byte("123"), nil
					}
					if scenario == "changed generation" {
						return []byte(strings.Repeat("b", 32)), nil
					}
					return []byte(generation), nil
				case "is-active":
					return []byte("active"), nil
				case "stop":
					if scenario == "stop failure" {
						return nil, errors.New("stop failed")
					}
					inactive = scenario != "still active"
					return nil, nil
				case "start":
					inactive = false
					return nil, nil
				}
				t.Fatal(args)
				return nil, nil
			}
			automaticGuardianInactive = func(context.Context) error {
				if !inactive {
					return errors.New("active")
				}
				return nil
			}
			control := &automaticServiceControl{generation: generation, root: root}
			err := control.Stop(context.Background())
			if (err == nil) != (scenario == "normal") {
				t.Fatalf("stop %v", err)
			}
			if scenario == "normal" {
				if _, err := guardian.Acquire(filepath.Join(root, "guardian"), generation); !errors.Is(err, guardian.ErrActive) {
					t.Fatalf("Guardian could enter during transaction: %v", err)
				}
			}
			if _, err := control.Recover(context.Background(), true); err != nil {
				t.Fatal(err)
			}
			if inactive {
				t.Fatal("Guardian not resumed")
			}
			if scenario == "normal" {
				lock, err := guardian.Acquire(filepath.Join(root, "guardian"), generation)
				if err != nil {
					t.Fatal(err)
				}
				_ = lock.Release()
			}
			if scenario == "changed generation" && len(operations) != 1 {
				t.Fatal("changed generation caused lifecycle mutation")
			}
		})
	}
}

func TestAutomaticTriggerFailureDoesNotDisableScheduler(t *testing.T) {
	req, _, _ := automaticFixture(t)
	state, err := triggerAwareness(t, req)
	if err != nil {
		t.Fatal(err)
	}
	root := filepath.Join(t.TempDir(), "state")
	if err = os.Mkdir(root, 0700); err != nil {
		t.Fatal(err)
	}
	store, _ := updateawareness.Open(root)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	checks := 0
	(guardian.ReleaseCheckService{Store: store, Interval: time.Hour, Timeout: time.Minute, Check: func(ctx context.Context) error {
		checks++
		r := guardian.AutomaticDecision(ctx, req.Capabilities, req.Policy, state, nil, func(context.Context) error { return errors.New("handoff unavailable") })
		if r.Outcome != "handoff_failed" {
			t.Fatalf("%+v", r)
		}
		if err := recordAutomatic(store, r, false); err != nil {
			t.Fatal(err)
		}
		return errors.New("handoff unavailable")
	}, Wait: func(context.Context, time.Time) error {
		if checks == 2 {
			cancel()
			return context.Canceled
		}
		return nil
	}}).Run(ctx)
	if checks != 2 {
		t.Fatalf("scheduler disabled after failure: %d", checks)
	}
}

func TestAutomaticSevereReceiptBlocksUnattendedRetry(t *testing.T) {
	root := filepath.Join(t.TempDir(), "state")
	if err := os.Mkdir(root, 0700); err != nil {
		t.Fatal(err)
	}
	store, _ := updateawareness.Open(root)
	if got := automaticInhibition(store); got != "" {
		t.Fatal(got)
	}
	for _, item := range []struct{ outcome, want string }{{"rollback_failed", "rollback_blocked"}, {"handoff_pending", "handoff_incomplete"}, {"orchestration_rolled_back", ""}, {"success", ""}, {"unknown", "evidence_invalid"}} {
		if err := recordAutomatic(store, guardian.AutomaticResult{At: time.Now().UTC(), Outcome: item.outcome}, true); err != nil {
			t.Fatal(err)
		}
		if got := automaticInhibition(store); got != item.want {
			t.Fatalf("%s: %s", item.outcome, got)
		}
	}
}
