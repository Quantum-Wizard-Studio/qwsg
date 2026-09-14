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
	"quantumwizard.hu/qwsg/internal/configurationstore"
	"quantumwizard.hu/qwsg/internal/productcapability"
	"quantumwizard.hu/qwsg/internal/updateawareness"
	"quantumwizard.hu/qwsg/internal/updatepolicy"
)

// executeAutomaticUpdate is an explicit code entry point, deliberately absent
// from command dispatch and Guardian startup. Production currently resolves
// Community authority and therefore refuses before fetching or staging.
func executeAutomaticUpdate(ctx context.Context) (automaticupdate.Result, error) {
	req := automaticupdate.Request{Platform: "linux-amd64", Now: time.Now().UTC()}
	capabilities, err := installationCapabilities()
	if err != nil {
		return automaticPreparationFailure(req, "capability_unavailable")
	}
	req.Capabilities = capabilities
	if !capabilities.Has(productcapability.UpdateAutomatic) {
		return automaticupdate.Run(ctx, req, nil)
	}
	path, err := configurationstore.DefaultPath(os.Getenv)
	if err != nil {
		return automaticPreparationFailure(req, "configuration_unavailable")
	}
	source, found, err := configurationstore.Load(path)
	if err != nil {
		return automaticPreparationFailure(req, "configuration_invalid")
	}
	effective, err := resolveLocalConfiguration(source, found, nil)
	if err != nil {
		return automaticPreparationFailure(req, "policy_invalid")
	}
	req.Policy, err = configuration.UpdatePolicy(effective.Values)
	if err != nil {
		return automaticPreparationFailure(req, "policy_invalid")
	}
	policy, err := updatepolicy.Evaluate(req.Policy, capabilities)
	if err != nil || policy.Effective != updatepolicy.Automatic {
		return automaticupdate.Run(ctx, req, nil)
	}
	if updateEffectiveUID() == 0 {
		return automaticPreparationFailure(req, "non_root_orchestrator_required")
	}
	req.SourceVersion, err = installedVersion()
	if err != nil {
		return automaticPreparationFailure(req, "installed_identity_unavailable")
	}
	state, err := localStateRoot()
	if err != nil {
		return automaticPreparationFailure(req, "state_unavailable")
	}
	req.StageParent = filepath.Join(state, "update")
	if err = ensureUpdateRoot(req.StageParent); err != nil {
		return automaticPreparationFailure(req, "state_unavailable")
	}
	store, err := updateawareness.Open(state)
	if err != nil {
		return automaticPreparationFailure(req, "watermark_unavailable")
	}
	awareness, err := store.Load()
	if err != nil && !errors.Is(err, updateawareness.ErrMissing) {
		return automaticPreparationFailure(req, "watermark_invalid")
	}
	if err == nil && awareness.LastSuccess != nil {
		req.NotBefore, err = time.Parse(time.RFC3339, awareness.LastSuccess.IndexGeneratedAt)
		if err != nil {
			return automaticPreparationFailure(req, "watermark_invalid")
		}
	}
	req.Verifier, err = updateVerifier()
	if err != nil {
		return automaticPreparationFailure(req, "release_authority_unavailable")
	}
	req.Metadata, err = updateMetadataFetch(ctx)
	if err != nil {
		return automaticPreparationFailure(req, "release_metadata_unavailable")
	}
	req.Evaluator = installedUpdateEvaluator()
	req.Client = updateHTTPClient()
	host := &automaticHost{root: req.StageParent, backup: filepath.Join(updateRollbackRoot, strconv.Itoa(os.Getuid()), time.Now().UTC().Format("20060102T150405.000000000Z"))}
	return automaticupdate.Run(ctx, req, host)
}

func automaticPreparationFailure(req automaticupdate.Request, category string) (automaticupdate.Result, error) {
	policy, err := updatepolicy.Evaluate(req.Policy, req.Capabilities)
	allowed := err == nil && policy.Effective == updatepolicy.Automatic
	stages := []automaticupdate.State{automaticupdate.Idle, automaticupdate.PolicyCheck}
	stage := automaticupdate.PolicyCheck
	if allowed {
		stage = automaticupdate.CandidateCheck
		stages = append(stages, stage)
	}
	stages = append(stages, automaticupdate.Failure)
	return automaticupdate.Result{SourceVersion: req.SourceVersion, TargetVersion: req.TargetVersion, CapabilityAllowed: req.Capabilities.Has(productcapability.UpdateAutomatic), PolicyAllowed: allowed, Policy: policy, State: automaticupdate.Failure, Stages: stages, FailureStage: stage, FailureCategory: category, RollbackResult: "not_required", MutationKnown: true}, errors.New(category)
}

type automaticHost struct{ root, backup string }

func (h *automaticHost) Preflight(ctx context.Context, from, to string) error {
	actual, err := installedVersion()
	if err != nil || actual != from {
		return fmt.Errorf("source identity changed")
	}
	if err = validateInstalledConfiguration(); err != nil {
		return err
	}
	return automaticGuardianInactive(ctx)
}

var automaticGuardianInactive = realAutomaticGuardianInactive

func realAutomaticGuardianInactive(ctx context.Context) error {
	// Task 084 deliberately does not manage Guardian service intent. Only an
	// explicitly quiescent installation is eligible; unknown/failed state refuses.
	cmd := exec.CommandContext(ctx, "/usr/bin/systemctl", "--user", "is-active", "qwsg-guardian.service")
	output, err := cmd.Output()
	var exit *exec.ExitError
	if strings.TrimSpace(string(output)) != "inactive" || !errors.As(err, &exit) || exit.ExitCode() != 3 {
		return fmt.Errorf("automatic update requires verified inactive Guardian")
	}
	return nil
}
func (h *automaticHost) Apply(ctx context.Context, p automaticupdate.PackageInput) (automaticupdate.ApplyResult, error) {
	args := []string{"privileged-apply-report", "--archive", p.Staged.Archive, "--sidecar", p.Staged.Sidecar, "--version", p.TargetVersion, "--sha256", p.Staged.SHA256, "--backup", h.backup, "--from", p.SourceVersion, "--authority", p.AuthorityPath}
	return runAutomaticApply(ctx, args...)
}

var runAutomaticApply = realRunAutomaticApply

func realRunAutomaticApply(ctx context.Context, args ...string) (automaticupdate.ApplyResult, error) {
	data, err := runAutomaticHelper(ctx, args...)
	return decodeApplyReceipt(data, err)
}

func runAutomaticHelper(ctx context.Context, args ...string) ([]byte, error) {
	executable, err := os.Executable()
	if err != nil {
		return nil, err
	}
	// Once launched, do not kill sudo and leave its privileged child applying in
	// parallel with rollback. The finite local helper must finish before recovery.
	if ctx.Err() != nil {
		return nil, ctx.Err()
	}
	cmd := exec.Command("/usr/bin/sudo", append([]string{"-n", "--", executable, "update"}, args...)...)
	var output boundedReceipt
	cmd.Stdout = &output
	err = cmd.Run()
	if output.overflow {
		return nil, errors.New("helper receipt exceeds bound")
	}
	return output.Bytes(), err
}

// Bounded capture continues draining rather than closing/killing the helper.
type boundedReceipt struct {
	bytes.Buffer
	overflow bool
}

func (b *boundedReceipt) Write(p []byte) (int, error) {
	n := len(p)
	if n > 4096-b.Len() {
		b.overflow = true
	}
	if b.Len() < 4096 {
		keep := 4096 - b.Len()
		if keep > n {
			keep = n
		}
		_, _ = b.Buffer.Write(p[:keep])
	}
	return n, nil
}
func decodeApplyReceipt(data []byte, commandErr error) (automaticupdate.ApplyResult, error) {
	var receipt automaticupdate.ApplyResult
	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&receipt); err != nil {
		return automaticupdate.ApplyResult{}, errors.New("helper receipt unavailable")
	}
	var extra any
	if decoder.Decode(&extra) != io.EOF || !receipt.Known || receipt.RollbackSucceeded && !receipt.RollbackAttempted || receipt.RollbackAttempted && !receipt.MutationStarted {
		return automaticupdate.ApplyResult{}, errors.New("helper receipt invalid")
	}
	return receipt, commandErr
}
func (h *automaticHost) Validate(ctx context.Context, version string) error {
	if ctx.Err() != nil {
		return ctx.Err()
	}
	if err := runSystemctl("daemon-reload"); err != nil {
		return err
	}
	return validateInstalledVersion(version)
}
func (h *automaticHost) Rollback(ctx context.Context) error {
	if ctx.Err() != nil {
		return ctx.Err()
	}
	return runAutomaticRollback(ctx, h.backup)
}
func (h *automaticHost) Commit(ctx context.Context, from, to string) error {
	if ctx.Err() != nil {
		return ctx.Err()
	}
	return saveUpdateRecord(h.root, localUpdateRecord{Schema: "qwsg.update-local/1", Installed: to, Previous: from, Backup: h.backup, UpdatedAt: time.Now().UTC().Format(time.RFC3339Nano)})
}

var runAutomaticRollback = func(ctx context.Context, backup string) error {
	_, err := runAutomaticHelper(ctx, "privileged-rollback", "--backup", backup)
	return err
}
