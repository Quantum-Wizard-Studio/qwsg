package userservice

import (
	"context"
	"errors"
	"quantumwizard.hu/qwsg/internal/runner"
	"quantumwizard.hu/qwsg/internal/userruntime"
	"testing"
)

type recording struct {
	calls   [][]string
	results []runner.Result
	errors  []error
}

func (r *recording) Run(_ context.Context, id string, args ...string) (runner.Result, error) {
	r.calls = append(r.calls, append([]string{id}, args...))
	index := len(r.calls) - 1
	var result runner.Result
	var err error
	if index < len(r.results) {
		result = r.results[index]
	}
	if index < len(r.errors) {
		err = r.errors[index]
	}
	return result, err
}

func TestControllerOwnsExactOperations(t *testing.T) {
	r := &recording{results: []runner.Result{{Stdout: []byte("running\n")}, {}, {}}}
	c := Controller{runner: r, runtimeOutcome: userruntime.Valid}
	if err := c.Activate(context.Background()); err != nil {
		t.Fatal(err)
	}
	if len(r.calls) != 3 || r.calls[0][0] != "systemd_user_manager" || r.calls[0][2] != "is-system-running" || r.calls[1][0] != "systemd_daemon_reload" || r.calls[1][2] != "daemon-reload" || r.calls[2][0] != "systemd_enable_start" || r.calls[2][2] != "enable" || r.calls[2][4] != Unit {
		t.Fatalf("calls=%q", r.calls)
	}
}

func TestNewUsesOnlyCurrentValidatedRuntimeContext(t *testing.T) {
	c := New()
	runtimeContext, outcome := userruntime.Current()
	if c.runtimeOutcome != outcome {
		t.Fatalf("controller outcome=%q current=%q", c.runtimeOutcome, outcome)
	}
	bounded, ok := c.runner.(runner.Bounded)
	if !ok {
		t.Fatalf("runner type=%T", c.runner)
	}
	for _, id := range []string{"systemd_user_manager", "systemd_daemon_reload", "systemd_enable_start"} {
		entries := bounded.TrustedEnvironment[id]
		if outcome == userruntime.Valid {
			if len(entries) != 1 || entries[0] != runtimeContext.Environment() {
				t.Fatalf("id=%q entries=%q", id, entries)
			}
		} else if len(entries) != 0 {
			t.Fatalf("invalid context received environment: id=%q entries=%q", id, entries)
		}
	}
}

func TestRuntimeContextFailureStopsBeforeCommands(t *testing.T) {
	for _, tc := range []struct {
		outcome userruntime.Outcome
		cause   Cause
	}{{userruntime.Missing, CauseContextMissing}, {userruntime.Unsafe, CauseContextUnsafe}} {
		r := &recording{}
		err := (Controller{runner: r, runtimeOutcome: tc.outcome}).Activate(context.Background())
		assertActivationError(t, err, StageRuntimeContext, tc.cause)
		if len(r.calls) != 0 {
			t.Fatalf("outcome=%q calls=%q", tc.outcome, r.calls)
		}
	}
}

func TestManagerReachabilityFailuresStopBeforeMutation(t *testing.T) {
	cases := []struct {
		result runner.Result
		err    error
		cause  Cause
	}{
		{runner.Result{}, errors.New("unavailable"), CauseManagerUnreachable},
		{runner.Result{Stdout: []byte("starting\n")}, nil, CauseManagerStarting},
		{runner.Result{Stdout: []byte("future\n")}, nil, CauseManagerStateUnrecognized},
		{runner.Result{Stdout: []byte("bounded\n")}, errors.New("exit"), CauseManagerProbeFailed},
		{runner.Result{}, context.DeadlineExceeded, CauseTimeout},
		{runner.Result{}, runner.ErrOutputLimit, CauseOutputLimit},
	}
	for _, tc := range cases {
		r := &recording{results: []runner.Result{tc.result}, errors: []error{tc.err}}
		err := (Controller{runner: r, runtimeOutcome: userruntime.Valid}).Activate(context.Background())
		assertActivationError(t, err, StageManager, tc.cause)
		if len(r.calls) != 1 {
			t.Fatalf("cause=%q calls=%q", tc.cause, r.calls)
		}
	}
}

func TestFixedOperationFailuresPreserveStageAndBounds(t *testing.T) {
	cases := []struct {
		name   string
		errors []error
		stage  Stage
		cause  Cause
		calls  int
	}{
		{"reload", []error{nil, errors.New("exit")}, StageDaemonReload, CauseOperationFailed, 2},
		{"reload timeout", []error{nil, context.DeadlineExceeded}, StageDaemonReload, CauseTimeout, 2},
		{"reload output", []error{nil, runner.ErrOutputLimit}, StageDaemonReload, CauseOutputLimit, 2},
		{"enable", []error{nil, nil, errors.New("exit")}, StageEnableStart, CauseOperationFailed, 3},
		{"enable timeout", []error{nil, nil, context.DeadlineExceeded}, StageEnableStart, CauseTimeout, 3},
		{"enable output", []error{nil, nil, runner.ErrOutputLimit}, StageEnableStart, CauseOutputLimit, 3},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			r := &recording{results: []runner.Result{{Stdout: []byte("degraded\n")}}, errors: tc.errors}
			err := (Controller{runner: r, runtimeOutcome: userruntime.Valid}).Activate(context.Background())
			assertActivationError(t, err, tc.stage, tc.cause)
			if len(r.calls) != tc.calls {
				t.Fatalf("calls=%q", r.calls)
			}
		})
	}
}

func TestActivationErrorIsDeterministicAndPrivate(t *testing.T) {
	err := &ActivationError{Stage: StageEnableStart, Cause: CauseOperationFailed}
	if got := err.Error(); got != "guardian activation failed: stage=enable_start cause=operation_failed" {
		t.Fatalf("error=%q", got)
	}
}

func assertActivationError(t *testing.T, err error, stage Stage, cause Cause) {
	t.Helper()
	var activationErr *ActivationError
	if !errors.As(err, &activationErr) || activationErr.Stage != stage || activationErr.Cause != cause {
		t.Fatalf("error=%v stage=%q cause=%q", err, stage, cause)
	}
}
