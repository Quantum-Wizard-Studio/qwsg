package main

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"quantumwizard.hu/qwsg/internal/automaticupdate"
	"quantumwizard.hu/qwsg/internal/configuration"
	"quantumwizard.hu/qwsg/internal/guardian"
	"quantumwizard.hu/qwsg/internal/updateawareness"
)

const automaticUnit = "qwsg-automatic-update.service"

func recordAutomatic(store *updateawareness.Store, r guardian.AutomaticResult, terminal bool) error {
	data, err := json.Marshal(r)
	if err != nil {
		return err
	}
	return store.RecordAutomation(data, terminal)
}

func guardianAutomaticCheck(ctx context.Context, options guardianOptions, effective configuration.Effective, store *updateawareness.Store, state updateawareness.State, checkErr error) error {
	caps, capErr := installationCapabilities()
	request, policyErr := configuration.UpdatePolicy(effective.Values)
	if capErr != nil || policyErr != nil {
		checkErr = errors.New("policy authority unavailable")
	}
	blocked := ""
	r := guardian.AutomaticDecision(ctx, caps, request, state, checkErr, func(ctx context.Context) error {
		blocked = automaticInhibition(store)
		if blocked != "" {
			return errors.New("automatic recovery requires review")
		}
		// Custom/manual Guardian instances are never allowed to stop the canonical
		// service or silently switch state/configuration roots during handoff.
		root, err := localStateRoot()
		if err != nil || root != options.stateRoot || options.configSource != "" {
			return errors.New("noncanonical guardian")
		}
		if err = verifyAutomaticGeneration(ctx, options.generation); err != nil {
			return err
		}
		pending := guardian.AutomaticResult{At: time.Now().UTC(), Outcome: "handoff_pending", CapabilityAllowed: true, CandidateVersion: state.LastSuccess.ReleaseVersion}
		pending.Policy, _ = effectiveUpdatePolicy(effective)
		if err = recordAutomatic(store, pending, false); err != nil {
			return err
		}
		return launchAutomaticHandoff(ctx, options.generation, root)
	})
	if blocked != "" {
		r.Outcome = blocked
	}
	return recordAutomatic(store, r, false)
}

var launchAutomaticHandoff = func(ctx context.Context, generation, root string) error {
	// The user manager creates a finite helper outside Guardian's cgroup and
	// NoNewPrivileges sandbox. A fork/setsid child would inherit both and would
	// be killed on stop. No timer, persistent updater unit or new scheduler exists.
	// Fixed unit name excludes duplicate handoffs. Worker loss leaves durable
	// incomplete evidence; only the Go recovery owner may restore Guardian.
	cmd := automaticHandoffCommand(ctx, generation, root)
	return cmd.Run()
}

func automaticHandoffCommand(ctx context.Context, generation, root string) *exec.Cmd {
	return exec.CommandContext(ctx, "/usr/bin/systemd-run", "--user", "--quiet", "--no-block", "--collect", "--unit="+automaticUnit,
		"--property=Type=exec", "--property=UMask=0077", "--setenv=QWSG_STATE_DIR="+root,
		"/usr/local/bin/qwsg", "guardian", "automatic-handoff", generation)
}

var automaticSystemctl = func(ctx context.Context, args ...string) ([]byte, error) {
	return exec.CommandContext(ctx, "/usr/bin/systemctl", append([]string{"--user"}, args...)...).Output()
}

func verifyAutomaticGeneration(ctx context.Context, generation string) error {
	if len(generation) != 32 || strings.Trim(generation, "0123456789abcdef") != "" {
		return errors.New("invalid generation")
	}
	data, err := automaticSystemctl(ctx, "show", "qwsg-guardian.service", "--property=InvocationID", "--value")
	if err != nil || strings.TrimSpace(string(data)) != generation {
		return errors.New("guardian generation changed")
	}
	data, err = automaticSystemctl(ctx, "is-active", "qwsg-guardian.service")
	if err != nil || strings.TrimSpace(string(data)) != "active" {
		return errors.New("guardian not active")
	}
	return nil
}

type automaticServiceControl struct {
	generation, root string
	stopRequested    bool
	intent           string
	beforeStop       func(guardian.Recovery) error
	lock             *guardian.Lock
}

func (s *automaticServiceControl) Stop(ctx context.Context) error {
	state, err := guardianServiceState(ctx)
	if err != nil {
		return err
	}
	s.intent = state
	if state == "active" {
		if err := verifyAutomaticGeneration(ctx, s.generation); err != nil {
			return err
		}
	} else if state != "inactive" {
		return errors.New("guardian state ambiguous")
	}
	pending := guardian.Recovery{Intent: state, State: "pending", StopRequested: state == "active"}
	if s.beforeStop != nil {
		if err := s.beforeStop(pending); err != nil {
			return err
		}
	}
	if state == "active" {
		s.stopRequested = true
		if _, err := automaticSystemctl(ctx, "stop", "qwsg-guardian.service"); err != nil {
			return err
		}
	}
	if err := automaticGuardianInactive(ctx); err != nil {
		return err
	}
	lock, err := guardian.Acquire(filepath.Join(s.root, "guardian"), s.generation)
	s.lock = lock
	return err
}
func (s *automaticServiceControl) Recover(ctx context.Context, safe bool) (guardian.Recovery, error) {
	r := guardian.Recovery{Intent: s.intent, State: "not_changed", StopRequested: s.stopRequested}
	if r.Intent == "" {
		r.Intent = "unchanged"
	}
	if s.lock != nil {
		if err := s.lock.Release(); err != nil {
			r.State = "failed"
			return r, err
		}
		s.lock = nil
	}
	return recoverGuardian(ctx, r, safe)
}

func recoverGuardian(ctx context.Context, r guardian.Recovery, safe bool) (guardian.Recovery, error) {
	if !safe {
		r.State = "blocked"
		return r, nil
	}
	if !r.StopRequested {
		if r.Intent == "inactive" {
			if err := automaticGuardianInactive(ctx); err != nil {
				r.State = "failed"
				return r, err
			}
			r.State = "preserved_stopped"
		}
		return r, nil
	}
	r.State = "failed"
	if _, err := automaticSystemctl(ctx, "start", "qwsg-guardian.service"); err != nil {
		return r, err
	}
	if err := verifyGuardianRecovery(ctx); err != nil {
		return r, err
	}
	r.State = "verified_running"
	return r, nil
}

// Read explicit service state: a failed query is never interpreted as stopped.
var guardianServiceState = realGuardianServiceState

func realGuardianServiceState(ctx context.Context) (string, error) {
	data, err := automaticSystemctl(ctx, "show", "qwsg-guardian.service", "--property=ActiveState", "--value")
	if err != nil {
		return "", err
	}
	state := strings.TrimSpace(string(data))
	if state != "active" && state != "inactive" {
		return "", errors.New("guardian state ambiguous")
	}
	return state, nil
}

// A successful start command is insufficient: require active/running with a
// live main process. All queries share the bounded recovery context.
func verifyGuardianRecovery(ctx context.Context) error {
	for _, check := range []struct{ property, want string }{{"ActiveState", "active"}, {"SubState", "running"}, {"Result", "success"}, {"MainPID", ""}} {
		data, err := automaticSystemctl(ctx, "show", "qwsg-guardian.service", "--property="+check.property, "--value")
		value := strings.TrimSpace(string(data))
		if err != nil {
			return err
		}
		if check.property == "MainPID" {
			pid, err := strconv.Atoi(value)
			if err != nil || pid <= 0 {
				return errors.New("guardian recovery unverified")
			}
		} else if value != check.want {
			return errors.New("guardian recovery unverified")
		}
	}
	return nil
}

func runAutomaticHandoff(args []string, errout io.Writer) int {
	if len(args) != 1 {
		return 1
	}
	// Require the manager-owned transient unit, not a general operator update CLI.
	// Capability and signed candidate authority are re-evaluated below regardless.
	data, err := os.ReadFile("/proc/self/cgroup")
	if err != nil || !strings.Contains(string(data), "/"+automaticUnit+"\n") {
		return 1
	}
	root, err := localStateRoot()
	if err != nil {
		return 1
	}
	store, err := updateawareness.Open(root)
	if err != nil {
		return 1
	}
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Minute)
	defer cancel()
	var result guardian.AutomaticResult
	prepared := false
	persisted := false
	claimed := false
	blocked := ""
	transaction, prepareErr := prepareAutomaticUpdate(ctx, func(ctx context.Context, req automaticupdate.Request, host automaticupdate.Host) (automaticupdate.Result, error) {
		prepared = true
		if outcome := automaticInhibition(store); outcome != "" {
			result = guardian.AutomaticResult{At: time.Now().UTC(), Outcome: outcome, CapabilityAllowed: true}
			return automaticupdate.Result{}, nil
		}
		control := &automaticServiceControl{generation: args[0], root: root}
		control.beforeStop = func(recovery guardian.Recovery) error {
			if blocked = automaticInhibition(store); blocked != "" {
				return errors.New("automatic recovery requires review")
			}
			err := recordAutomatic(store, guardian.AutomaticResult{At: time.Now().UTC(), Outcome: "handoff_pending", CapabilityAllowed: true, Recovery: recovery}, true)
			claimed = err == nil
			return err
		}
		result = guardian.AutomaticHandoff(ctx, req, host, control, func(r guardian.AutomaticResult) error {
			if !claimed {
				if blocked = automaticInhibition(store); blocked != "" {
					return nil
				}
			}
			persisted = true
			return recordAutomatic(store, r, true)
		})
		if result.Transaction != nil {
			return *result.Transaction, nil
		}
		return automaticupdate.Result{}, nil
	})
	if !prepared || prepareErr != nil {
		result = guardian.AutomaticResult{At: time.Now().UTC(), Outcome: "not_authorized", CapabilityAllowed: transaction.CapabilityAllowed, Policy: transaction.Policy, Transaction: &transaction}
		if transaction.PolicyAllowed {
			result.Outcome = "candidate_rejected"
		}
	}
	if blocked != "" {
		result.Outcome = blocked
	}
	// Without the mutation lease, update only decision evidence. A contender
	// must never overwrite another worker's durable pending/terminal receipt.
	if !persisted {
		err = recordAutomatic(store, result, false)
	}
	if err != nil || result.Outcome == "evidence_failed" {
		fmt.Fprintln(errout, "automatic_evidence_failed")
		return 1
	}
	if prepareErr != nil || result.Outcome != "success" && result.Outcome != "no_update" {
		return 1
	}
	return 0
}

// A severe or incomplete prior handoff is not an ordinary retry. Preserve its
// receipt and require operator recovery; corrupt evidence also fails closed.
func automaticInhibition(store *updateawareness.Store) string {
	data, err := store.LoadAutomation()
	if errors.Is(err, os.ErrNotExist) || errors.Is(err, updateawareness.ErrMissing) {
		return ""
	}
	if err != nil {
		return "evidence_invalid"
	}
	var r guardian.AutomaticResult
	d := json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if d.Decode(&r) != nil {
		return "evidence_invalid"
	}
	var extra any
	if d.Decode(&extra) != io.EOF {
		return "evidence_invalid"
	}
	// Legacy receipts have no Recovery member. New receipts must not turn an
	// unknown or pending recovery state into a retry-safe terminal result.
	switch r.Recovery.State {
	case "", "not_changed", "verified_running", "preserved_stopped", "failed", "blocked":
	case "pending":
		return "handoff_incomplete"
	default:
		return "evidence_invalid"
	}
	if r.Outcome == "rollback_failed" || r.Transaction != nil && r.Transaction.State == automaticupdate.RollbackFailure {
		return "rollback_blocked"
	}
	if r.RestartFailed || r.Outcome == "recovery_failed" || r.Outcome == "evidence_failed" || r.Recovery.State == "failed" || r.Recovery.State == "blocked" {
		return "handoff_incomplete"
	}
	if r.Outcome == "handoff_pending" || r.Outcome == "handoff_incomplete" {
		return "handoff_incomplete"
	}
	switch r.Outcome {
	case "success", "no_update", "not_authorized", "candidate_rejected", "handoff_failed", "orchestration_failed", "orchestration_rolled_back":
		return ""
	}
	return "evidence_invalid"
}
