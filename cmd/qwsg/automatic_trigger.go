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
	// Fixed unit name excludes duplicate handoffs. ExecStopPost resumes Guardian
	// even if the coordinator unexpectedly exits after requesting the stop.
	cmd := exec.CommandContext(ctx, "/usr/bin/systemd-run", "--user", "--quiet", "--no-block", "--collect", "--unit="+automaticUnit,
		"--property=Type=exec", "--property=UMask=0077", "--setenv=QWSG_STATE_DIR="+root,
		"--property=ExecStopPost=/usr/bin/systemctl --user start qwsg-guardian.service",
		"/usr/local/bin/qwsg", "guardian", "automatic-handoff", generation)
	return cmd.Run()
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
	lock             *guardian.Lock
}

func (s *automaticServiceControl) Stop(ctx context.Context) error {
	if err := verifyAutomaticGeneration(ctx, s.generation); err != nil {
		return err
	}
	s.stopRequested = true
	if _, err := automaticSystemctl(ctx, "stop", "qwsg-guardian.service"); err != nil {
		return err
	}
	if err := automaticGuardianInactive(ctx); err != nil {
		return err
	}
	lock, err := guardian.Acquire(filepath.Join(s.root, "guardian"), s.generation)
	s.lock = lock
	return err
}
func (s *automaticServiceControl) Resume(ctx context.Context) error {
	if s.lock != nil {
		if err := s.lock.Release(); err != nil {
			return err
		}
		s.lock = nil
	}
	if !s.stopRequested {
		return nil
	}
	_, err := automaticSystemctl(ctx, "start", "qwsg-guardian.service")
	return err
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
	transaction, prepareErr := prepareAutomaticUpdate(ctx, func(ctx context.Context, req automaticupdate.Request, host automaticupdate.Host) (automaticupdate.Result, error) {
		prepared = true
		if outcome := automaticInhibition(store); outcome != "" {
			result = guardian.AutomaticResult{At: time.Now().UTC(), Outcome: outcome, CapabilityAllowed: true}
			return automaticupdate.Result{}, nil
		}
		pending := guardian.AutomaticResult{At: time.Now().UTC(), Outcome: "handoff_pending", CapabilityAllowed: true}
		if err := recordAutomatic(store, pending, true); err != nil {
			return automaticupdate.Result{}, err
		}
		result = guardian.AutomaticHandoff(ctx, req, host, &automaticServiceControl{generation: args[0], root: root})
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
	terminal := result.Outcome != "rollback_blocked" && result.Outcome != "handoff_incomplete" && result.Outcome != "evidence_invalid"
	if err = recordAutomatic(store, result, terminal); err != nil {
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
	if r.Outcome == "rollback_failed" || r.Transaction != nil && r.Transaction.State == automaticupdate.RollbackFailure {
		return "rollback_blocked"
	}
	if r.Outcome == "handoff_pending" {
		return "handoff_incomplete"
	}
	switch r.Outcome {
	case "success", "no_update", "not_authorized", "candidate_rejected", "handoff_failed", "orchestration_failed", "orchestration_rolled_back":
		return ""
	}
	return "evidence_invalid"
}
