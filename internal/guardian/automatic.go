package guardian

import (
	"context"
	"errors"
	"time"

	"quantumwizard.hu/qwsg/internal/automaticupdate"
	"quantumwizard.hu/qwsg/internal/productcapability"
	"quantumwizard.hu/qwsg/internal/updateauthority"
	"quantumwizard.hu/qwsg/internal/updateawareness"
	"quantumwizard.hu/qwsg/internal/updatepolicy"
)

// AutomaticResult composes the transaction receipt rather than copying its
// mutation/rollback fields. Decision and terminal evidence are separate records.
type AutomaticResult struct {
	At                time.Time               `json:"at"`
	Outcome           string                  `json:"outcome"`
	CapabilityAllowed bool                    `json:"capability_allowed"`
	Policy            updatepolicy.State      `json:"policy"`
	CandidateVersion  string                  `json:"candidate_version,omitempty"`
	Invoked           bool                    `json:"orchestration_invoked"`
	Recovery          Recovery                `json:"recovery"`
	RestartFailed     bool                    `json:"restart_failed,omitempty"`
	Transaction       *automaticupdate.Result `json:"transaction,omitempty"`
}

// AutomaticDecision runs only from the Guardian's existing release-check cycle.
// Cached awareness is a decision hint, never authority to mutate: the handoff
// authenticates again and Run independently authorizes the package transaction.
func AutomaticDecision(ctx context.Context, caps productcapability.Set, request updatepolicy.Request, state updateawareness.State, checkErr error, handoff func(context.Context) error) AutomaticResult {
	r := AutomaticResult{At: time.Now().UTC(), Outcome: "not_authorized", CapabilityAllowed: caps.Has(productcapability.UpdateAutomatic)}
	var err error
	r.Policy, err = updatepolicy.Evaluate(request, caps)
	if err != nil || !r.CapabilityAllowed || r.Policy.Effective != updatepolicy.Automatic {
		return r
	}
	r.Outcome = "candidate_rejected"
	if checkErr != nil || updateawareness.Validate(state) != nil || state.LastAttempt.Outcome != updateawareness.AttemptSuccess || state.LastSuccess == nil || !state.LastSuccess.FreshUntil.After(r.At) {
		return r
	}
	r.CandidateVersion = state.LastSuccess.ReleaseVersion
	if state.Status == updateawareness.Current {
		r.Outcome = "no_update"
		return r
	}
	if state.Status != updateawareness.UpdateAvailable {
		return r
	}
	r.Outcome = "handoff_failed"
	if handoff != nil && handoff(ctx) == nil {
		r.Outcome = "handoff_requested"
	}
	return r
}

// Recovery distinguishes preserved intent from package transaction success.
// Unknown/incomplete evidence never authorizes a service start.
type Recovery struct {
	Intent        string `json:"intent"`
	State         string `json:"state"`
	StopRequested bool   `json:"stop_requested"`
}

// HandoffControl belongs to the trusted local service adapter. Stop must verify
// the intended running generation, wait for full inactivity and hold Guardian's
// instance lock until Recover. Recover checks preserved intent and package safety;
// it also runs when stopping failed midway, and verifies any resulting restart.
type HandoffControl interface {
	Stop(context.Context) error
	Recover(context.Context, bool) (Recovery, error)
}

func AutomaticHandoff(ctx context.Context, req automaticupdate.Request, host automaticupdate.Host, control HandoffControl, persist ...func(AutomaticResult) error) AutomaticResult {
	r := AutomaticResult{At: time.Now().UTC(), Outcome: "not_authorized", CapabilityAllowed: req.Capabilities.Has(productcapability.UpdateAutomatic), Recovery: Recovery{Intent: "unchanged", State: "not_changed"}}
	var err error
	r.Policy, err = updatepolicy.Evaluate(req.Policy, req.Capabilities)
	if err != nil || !r.CapabilityAllowed || r.Policy.Effective != updatepolicy.Automatic {
		return r
	}
	r.Outcome = "candidate_rejected"
	candidate, err := updateauthority.Authorize(req.Metadata, req.Verifier, req.Evaluator, "stable", req.Platform, req.TargetVersion, req.Now)
	if err != nil && !errors.Is(err, updateauthority.ErrNoUpdate) {
		return r
	}
	if candidate.CheckWatermark(req.NotBefore) != nil || candidate.Evaluation().InstalledVersion != req.SourceVersion {
		return r
	}
	r.CandidateVersion = candidate.Evaluation().Release.Version
	if errors.Is(err, updateauthority.ErrNoUpdate) {
		r.Outcome = "no_update"
		return r
	}
	r.Outcome = "handoff_failed"
	if control == nil || host == nil {
		return r
	}
	h := &handoffHost{Host: host, control: control}
	h.finish = func(ctx context.Context, transaction automaticupdate.Result) error {
		r.Transaction = &transaction
		r.Outcome = "orchestration_failed"
		if h.stopFailed {
			r.Outcome = "handoff_failed"
		}
		switch transaction.State {
		case automaticupdate.RollbackFailure:
			r.Outcome = "rollback_failed"
		case automaticupdate.RollbackSuccess:
			r.Outcome = "orchestration_rolled_back"
		case automaticupdate.Success:
			r.Outcome = "success"
		}
		recovery, cancel := context.WithTimeout(ctx, 30*time.Second)
		defer cancel()
		var recoveryErr error
		safe := transaction.State == automaticupdate.Success || transaction.State == automaticupdate.RollbackSuccess ||
			transaction.State == automaticupdate.Failure && transaction.MutationKnown && !transaction.MutationStarted
		if !safe && transaction.State != automaticupdate.RollbackFailure {
			r.Outcome = "handoff_incomplete"
		}
		r.Recovery, recoveryErr = control.Recover(recovery, safe)
		r.RestartFailed = recoveryErr != nil
		if recoveryErr != nil && r.Outcome != "rollback_failed" {
			r.Outcome = "recovery_failed"
		}
		for _, save := range persist {
			if err := save(r); err != nil {
				r.Outcome = "evidence_failed"
				return err
			}
		}
		return recoveryErr
	}
	r.Invoked = true
	transaction, _ := automaticupdate.Run(ctx, req, h)
	if !h.finalized {
		r.Outcome = "orchestration_failed"
		r.Transaction = &transaction
		r.Recovery = Recovery{Intent: "unchanged", State: "not_changed"}
	}
	return r
}

// Stop, transaction, recovery and evidence share Run's mutation lease. Staging
// and authentication failures occur before Stop and cannot change service intent.
type handoffHost struct {
	automaticupdate.Host
	control               HandoffControl
	stopFailed, finalized bool
	finish                func(context.Context, automaticupdate.Result) error
}

func (h *handoffHost) Preflight(ctx context.Context, from, to string) error {
	if err := h.control.Stop(ctx); err != nil {
		h.stopFailed = true
		return err
	}
	return h.Host.Preflight(ctx, from, to)
}
func (h *handoffHost) Finalize(ctx context.Context, r automaticupdate.Result) error {
	h.finalized = true
	return h.finish(ctx, r)
}
