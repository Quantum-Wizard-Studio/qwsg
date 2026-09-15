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

// HandoffControl belongs to the trusted local service adapter. Stop must verify
// the intended running generation, wait for full inactivity and hold Guardian's
// instance lock until Resume. Resume also runs when stopping failed midway.
type HandoffControl interface {
	Stop(context.Context) error
	Resume(context.Context) error
}

func AutomaticHandoff(ctx context.Context, req automaticupdate.Request, host automaticupdate.Host, control HandoffControl) AutomaticResult {
	r := AutomaticResult{At: time.Now().UTC(), Outcome: "not_authorized", CapabilityAllowed: req.Capabilities.Has(productcapability.UpdateAutomatic)}
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
	if control == nil {
		return r
	}
	if err = control.Stop(ctx); err == nil {
		r.Invoked = true
		result, runErr := automaticupdate.Run(ctx, req, host)
		r.Transaction = &result
		r.Outcome = "orchestration_failed"
		if result.State == automaticupdate.RollbackFailure {
			r.Outcome = "rollback_failed"
		} else if result.State == automaticupdate.RollbackSuccess {
			r.Outcome = "orchestration_rolled_back"
		} else if runErr == nil && result.State == automaticupdate.Success {
			r.Outcome = "success"
		}
	}
	recovery, cancel := context.WithTimeout(context.WithoutCancel(ctx), 3*time.Minute)
	defer cancel()
	r.RestartFailed = control.Resume(recovery) != nil
	if r.RestartFailed && r.Outcome != "rollback_failed" {
		r.Outcome = "handoff_failed"
	}
	return r
}
