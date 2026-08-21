// Package userservice owns the narrow, explicitly authorized QWSG user-unit
// activation boundary. It cannot select another executable, unit, or argv.
package userservice

import (
	"context"
	"errors"
	"fmt"
	"time"

	"quantumwizard.hu/qwsg/internal/runner"
	"quantumwizard.hu/qwsg/internal/userruntime"
)

const Unit = "qwsg-guardian.service"

type Stage string

const (
	StageRuntimeContext Stage = "runtime_context"
	StageManager        Stage = "manager_reachability"
	StageDaemonReload   Stage = "daemon_reload"
	StageEnableStart    Stage = "enable_start"
)

type Cause string

const (
	CauseContextMissing           Cause = "context_missing"
	CauseContextUnsafe            Cause = "context_unsafe"
	CauseManagerUnreachable       Cause = "manager_unreachable"
	CauseManagerStarting          Cause = "manager_starting"
	CauseManagerProbeFailed       Cause = "manager_probe_failed"
	CauseManagerStateUnrecognized Cause = "manager_state_unrecognized"
	CauseTimeout                  Cause = "timeout"
	CauseOutputLimit              Cause = "output_limit"
	CauseOperationFailed          Cause = "operation_failed"
)

// ActivationError exposes only deterministic fixed-operation identity. Raw
// command output and host environment never cross this boundary.
type ActivationError struct {
	Stage Stage
	Cause Cause
}

func (e *ActivationError) Error() string {
	return fmt.Sprintf("guardian activation failed: stage=%s cause=%s", e.Stage, e.Cause)
}

type Controller struct {
	runner         runner.Runner
	runtimeOutcome userruntime.Outcome
}

func New() Controller {
	runtimeContext, outcome := userruntime.Current()
	environment := map[string][]string{}
	if outcome == userruntime.Valid {
		for _, id := range []string{"systemd_user_manager", "systemd_daemon_reload", "systemd_enable_start"} {
			environment[id] = []string{runtimeContext.Environment()}
		}
	}
	return Controller{
		runner: runner.Bounded{
			Allowed: map[string]string{
				"systemd_user_manager":  "/usr/bin/systemctl",
				"systemd_daemon_reload": "/usr/bin/systemctl",
				"systemd_enable_start":  "/usr/bin/systemctl",
			},
			TrustedEnvironment: environment,
			Timeout:            10 * time.Second,
			MaxOutput:          64 << 10,
		},
		runtimeOutcome: outcome,
	}
}

func (c Controller) Activate(ctx context.Context) error {
	switch c.runtimeOutcome {
	case userruntime.Missing:
		return &ActivationError{Stage: StageRuntimeContext, Cause: CauseContextMissing}
	case userruntime.Unsafe:
		return &ActivationError{Stage: StageRuntimeContext, Cause: CauseContextUnsafe}
	case userruntime.Valid:
	default:
		return &ActivationError{Stage: StageRuntimeContext, Cause: CauseContextUnsafe}
	}
	result, err := c.runner.Run(ctx, "systemd_user_manager", "--user", "is-system-running")
	if activationErr := managerError(result, err); activationErr != nil {
		return activationErr
	}
	if _, err = c.runner.Run(ctx, "systemd_daemon_reload", "--user", "daemon-reload"); err != nil {
		return operationError(StageDaemonReload, err)
	}
	if _, err = c.runner.Run(ctx, "systemd_enable_start", "--user", "enable", "--now", Unit); err != nil {
		return operationError(StageEnableStart, err)
	}
	return nil
}

func managerError(result runner.Result, err error) error {
	outcome, _ := userruntime.ClassifyManager(result, err)
	switch outcome {
	case userruntime.ManagerReachable:
		return nil
	case userruntime.ManagerTimeout:
		return &ActivationError{Stage: StageManager, Cause: CauseTimeout}
	case userruntime.ManagerOutputLimit:
		return &ActivationError{Stage: StageManager, Cause: CauseOutputLimit}
	case userruntime.ManagerTransient:
		return &ActivationError{Stage: StageManager, Cause: CauseManagerStarting}
	case userruntime.ManagerUnavailable:
		return &ActivationError{Stage: StageManager, Cause: CauseManagerUnreachable}
	case userruntime.ManagerProbeFailed:
		return &ActivationError{Stage: StageManager, Cause: CauseManagerProbeFailed}
	default:
		return &ActivationError{Stage: StageManager, Cause: CauseManagerStateUnrecognized}
	}
}

func operationError(stage Stage, err error) error {
	switch {
	case errors.Is(err, context.DeadlineExceeded):
		return &ActivationError{Stage: stage, Cause: CauseTimeout}
	case errors.Is(err, runner.ErrOutputLimit):
		return &ActivationError{Stage: stage, Cause: CauseOutputLimit}
	default:
		return &ActivationError{Stage: stage, Cause: CauseOperationFailed}
	}
}
