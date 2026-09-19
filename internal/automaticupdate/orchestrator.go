// Package automaticupdate coordinates the common authenticated update engine.
// It has no scheduler, trigger, entitlement source, or background goroutine.
package automaticupdate

import (
	"context"
	"errors"
	"net/http"
	"os"
	"path/filepath"
	"sync/atomic"
	"time"

	"quantumwizard.hu/qwsg/internal/productcapability"
	"quantumwizard.hu/qwsg/internal/releasediscovery"
	"quantumwizard.hu/qwsg/internal/update"
	"quantumwizard.hu/qwsg/internal/updateauthority"
	"quantumwizard.hu/qwsg/internal/updatemutation"
	"quantumwizard.hu/qwsg/internal/updatepolicy"
)

type State string

const (
	Idle                 State = "idle"
	PolicyCheck          State = "policy_check"
	CandidateCheck       State = "candidate_check"
	EligibilityCheck     State = "eligibility_check"
	Staging              State = "staging"
	Preflight            State = "preflight"
	Backup               State = "backup"
	Apply                State = "apply"
	PostUpdateValidation State = "post_update_validation"
	Success              State = "success"
	Failure              State = "failure"
	Rollback             State = "rollback"
	RollbackSuccess      State = "rollback_success"
	RollbackFailure      State = "rollback_failure"
)

type Result struct {
	SourceVersion     string             `json:"source_version"`
	TargetVersion     string             `json:"target_version"`
	CapabilityAllowed bool               `json:"capability_allowed"`
	PolicyAllowed     bool               `json:"policy_allowed"`
	Policy            updatepolicy.State `json:"policy"`
	State             State              `json:"state"`
	Stages            []State            `json:"stages"`
	FailureStage      State              `json:"failure_stage,omitempty"`
	FailureCategory   string             `json:"failure_category,omitempty"`
	MutationStarted   bool               `json:"mutation_started"`
	MutationKnown     bool               `json:"mutation_known"`
	RollbackAttempted bool               `json:"rollback_attempted"`
	RollbackResult    string             `json:"rollback_result"`
}

// ApplyResult is a bounded helper receipt. Known=false means communication was
// lost: mutation cannot be ruled out and recovery must be attempted. A successful
// rollback is never inferred from an apply error alone.
type ApplyResult struct {
	Known             bool `json:"known"`
	MutationStarted   bool `json:"mutation_started"`
	RollbackAttempted bool `json:"rollback_attempted"`
	RollbackSucceeded bool `json:"rollback_succeeded"`
}

type PackageInput struct {
	Staged                       update.Staged
	AuthorityPath                string
	SourceVersion, TargetVersion string
}

// Host is a trusted local composition boundary, never metadata-supplied code.
// Apply MUST invoke the common independently authenticating privileged helper.
// Preflight cannot mutate the installation. Rollback and Validate must work even
// when the caller's context has been cancelled. Commit records local rollback
// metadata atomically and must not discard previous rollback material on error.
// Run owns the common mutation lease across every Host call. Host implementations
// must not call another coordinator or acquire the mutation lease recursively.
type Host interface {
	Preflight(context.Context, string, string) error
	Apply(context.Context, PackageInput) (ApplyResult, error)
	Validate(context.Context, string) error
	Rollback(context.Context) error
	Commit(context.Context, string, string) error
}

// Request is supplied by trusted code. Capabilities come from Task 083 authority,
// Policy from its canonical parser; neither is taken from release metadata.
// StageParent is the canonical private update directory for this installation.
// Verifier/Evaluator are fixed production dependencies or isolated test fixtures.
type Request struct {
	Capabilities                           productcapability.Set
	Policy                                 updatepolicy.Request
	Metadata                               []byte
	Verifier                               releasediscovery.Verifier
	Evaluator                              releasediscovery.Evaluator
	Now                                    time.Time
	NotBefore                              time.Time
	SourceVersion, TargetVersion, Platform string
	StageParent, LocalArchive              string
	Client                                 *http.Client
}

var running atomic.Bool

// Run is synchronous and explicitly invoked. The process guard also rejects
// recursive invocation across different Runner/Host instances. The nonblocking
// file lock covers independent processes sharing the canonical update directory.
func Run(ctx context.Context, req Request, host Host) (Result, error) {
	r := Result{SourceVersion: req.SourceVersion, TargetVersion: req.TargetVersion, RollbackResult: "not_required", MutationKnown: true}
	enter := func(s State) { r.State = s; r.Stages = append(r.Stages, s) }
	enter(Idle)
	fail := func(category string) (Result, error) {
		r.FailureStage = r.State
		r.FailureCategory = category
		enter(Failure)
		return r, errors.New(category)
	}
	enter(PolicyCheck)
	r.CapabilityAllowed = req.Capabilities.Has(productcapability.UpdateAutomatic)
	policy, err := updatepolicy.Evaluate(req.Policy, req.Capabilities)
	r.Policy = policy
	if !r.CapabilityAllowed {
		return fail("automatic_capability_unavailable")
	}
	if err != nil || policy.Effective != updatepolicy.Automatic {
		return fail("automatic_policy_refused")
	}
	r.PolicyAllowed = true
	if host == nil {
		return fail("host_unavailable")
	}
	if !running.CompareAndSwap(false, true) {
		return fail("transaction_conflict")
	}
	defer running.Store(false)
	lock, err := updatemutation.Acquire(req.StageParent)
	if errors.Is(err, updatemutation.ErrContended) {
		return fail("transaction_conflict")
	}
	if err != nil {
		return fail("transaction_lock_unavailable")
	}
	defer lock.Close() // Closing the descriptor releases flock; never unlink it.
	if ctx.Err() != nil {
		return fail("cancelled")
	}
	enter(CandidateCheck)
	candidate, err := updateauthority.Authorize(req.Metadata, req.Verifier, req.Evaluator, "stable", req.Platform, req.TargetVersion, req.Now)
	if err != nil {
		return fail("release_authority_refused")
	}
	e := candidate.Evaluation()
	r.TargetVersion = e.Release.Version
	enter(EligibilityCheck)
	if e.InstalledVersion != req.SourceVersion || e.MigrationID != update.PreservePackageV1 || candidate.CheckWatermark(req.NotBefore) != nil {
		return fail("candidate_eligibility_refused")
	}
	enter(Staging)
	var staged update.Staged
	if req.LocalArchive != "" {
		staged, err = update.StageLocal(req.LocalArchive, req.LocalArchive+".sha256", r.TargetVersion, req.StageParent)
	} else if req.Client != nil {
		a := e.Artifact
		staged, err = update.AcquireDigest(ctx, req.Client, update.Asset{Name: a.Name, URL: a.URL, Size: a.Size}, r.TargetVersion, a.SHA256, req.StageParent)
	} else {
		return fail("acquisition_unavailable")
	}
	if err != nil {
		return fail("package_integrity_refused")
	}
	defer os.RemoveAll(staged.Root)
	if _, err = candidate.VerifyStaged(staged); err != nil {
		return fail("package_authority_refused")
	}
	authority := filepath.Join(staged.Root, "release-index.json")
	if err = os.WriteFile(authority, candidate.Metadata(), 0600); err != nil {
		return fail("authority_staging_failed")
	}
	enter(Preflight)
	if err = host.Preflight(ctx, r.SourceVersion, r.TargetVersion); err != nil {
		return fail("preflight_failed")
	}
	if ctx.Err() != nil {
		return fail("cancelled")
	}
	// Backup and apply are one common helper transaction: each destination's
	// rollback material is captured before that destination is changed.
	enter(Backup)
	enter(Apply)
	receipt, applyErr := host.Apply(ctx, PackageInput{Staged: staged, AuthorityPath: authority, SourceVersion: r.SourceVersion, TargetVersion: r.TargetVersion})
	r.MutationKnown = receipt.Known
	r.MutationStarted = receipt.MutationStarted || !receipt.Known
	r.RollbackAttempted = receipt.RollbackAttempted
	if applyErr == nil && (!receipt.Known || !receipt.MutationStarted || receipt.RollbackAttempted) {
		r.MutationKnown = false
		r.MutationStarted = true // Inconsistent success cannot prove absence of mutation.
		applyErr = errors.New("invalid apply receipt")
	}
	category := "apply_failed"
	if applyErr == nil {
		enter(PostUpdateValidation)
		applyErr = host.Validate(ctx, r.TargetVersion)
		category = "post_update_validation_failed"
	}
	if applyErr == nil {
		applyErr = host.Commit(ctx, r.SourceVersion, r.TargetVersion)
		category = "commit_failed"
	}
	if applyErr == nil {
		enter(Success)
		return r, nil
	}
	r.FailureStage = r.State
	r.FailureCategory = category
	enter(Failure)
	if !r.MutationStarted {
		return r, errors.New(category)
	}
	enter(Rollback)
	r.RollbackAttempted = true
	// Cancellation must not prevent recovery. A separate bounded recovery context
	// prevents a failed/cancelled apply from silently abandoning rollback.
	recovery, cancel := context.WithTimeout(context.WithoutCancel(ctx), 2*time.Minute)
	defer cancel()
	var recoveryErr error
	if receipt.RollbackAttempted {
		if !receipt.RollbackSucceeded {
			recoveryErr = errors.New("transaction rollback failed")
		}
	} else {
		recoveryErr = host.Rollback(recovery)
	}
	if recoveryErr == nil {
		recoveryErr = host.Validate(recovery, r.SourceVersion)
	}
	if recoveryErr != nil {
		r.RollbackResult = "failed"
		enter(RollbackFailure)
		return r, errors.New("rollback_failed")
	}
	r.RollbackResult = "succeeded"
	enter(RollbackSuccess)
	return r, errors.New(category)
}
