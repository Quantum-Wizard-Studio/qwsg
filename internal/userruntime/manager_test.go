package userruntime

import (
	"context"
	"errors"
	"testing"

	"quantumwizard.hu/qwsg/internal/runner"
)

func TestClassifyManagerStableStates(t *testing.T) {
	cases := []struct {
		result  runner.Result
		err     error
		outcome ManagerOutcome
		state   string
	}{
		{runner.Result{Stdout: []byte("running\n")}, nil, ManagerReachable, "running"},
		{runner.Result{Stdout: []byte("degraded\n"), ExitCode: 1}, errors.New("exit"), ManagerReachable, "degraded"},
		{runner.Result{Stdout: []byte("starting\n")}, nil, ManagerTransient, ""},
		{runner.Result{}, context.DeadlineExceeded, ManagerTimeout, ""},
		{runner.Result{}, runner.ErrOutputLimit, ManagerOutputLimit, ""},
		{runner.Result{}, errors.New("unavailable"), ManagerUnavailable, ""},
		{runner.Result{Stdout: []byte("noise\n")}, errors.New("exit"), ManagerProbeFailed, ""},
		{runner.Result{Stdout: []byte("future\n")}, nil, ManagerUnrecognized, ""},
	}
	for _, tc := range cases {
		outcome, state := ClassifyManager(tc.result, tc.err)
		if outcome != tc.outcome || state != tc.state {
			t.Fatalf("result=%q outcome=%q state=%q", tc.result.Stdout, outcome, state)
		}
	}
}
