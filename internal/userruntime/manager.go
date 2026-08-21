package userruntime

import (
	"context"
	"errors"
	"strings"

	"quantumwizard.hu/qwsg/internal/runner"
)

type ManagerOutcome string

const (
	ManagerReachable    ManagerOutcome = "reachable"
	ManagerTransient    ManagerOutcome = "transient"
	ManagerUnavailable  ManagerOutcome = "unavailable"
	ManagerTimeout      ManagerOutcome = "timeout"
	ManagerOutputLimit  ManagerOutcome = "output_limit"
	ManagerProbeFailed  ManagerOutcome = "probe_failed"
	ManagerUnrecognized ManagerOutcome = "unrecognized"
)

// ClassifyManager is the shared interpretation boundary for the fixed
// systemctl --user is-system-running probe. It never inspects stderr or locale
// text and returns only a stable outcome plus a recognized state.
func ClassifyManager(result runner.Result, err error) (ManagerOutcome, string) {
	state := strings.TrimSpace(string(result.Stdout))
	if state == "running" || state == "degraded" {
		return ManagerReachable, state
	}
	for _, transient := range []string{"initializing", "starting", "maintenance", "stopping"} {
		if state == transient {
			return ManagerTransient, ""
		}
	}
	if errors.Is(err, context.DeadlineExceeded) {
		return ManagerTimeout, ""
	}
	if errors.Is(err, runner.ErrOutputLimit) {
		return ManagerOutputLimit, ""
	}
	if err != nil && state == "" {
		return ManagerUnavailable, ""
	}
	if err != nil {
		return ManagerProbeFailed, ""
	}
	return ManagerUnrecognized, ""
}
