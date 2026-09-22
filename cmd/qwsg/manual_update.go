package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"time"

	"quantumwizard.hu/qwsg/internal/automaticupdate"
	"quantumwizard.hu/qwsg/internal/evidenceio"
	"quantumwizard.hu/qwsg/internal/guardian"
	updatecore "quantumwizard.hu/qwsg/internal/update"
)

// Manual coordination shares the authenticated helper and Guardian recovery
// contracts. The common mutation lease belongs to executeUpdate throughout.
type manualUpdateEvidence struct {
	Schema             string                      `json:"schema"`
	At                 string                      `json:"at"`
	Source             string                      `json:"source_version"`
	Target             string                      `json:"target_version"`
	Backup             string                      `json:"backup"`
	Phase              string                      `json:"phase"`
	Outcome            string                      `json:"outcome"`
	Failure            string                      `json:"failure,omitempty"`
	Apply              automaticupdate.ApplyResult `json:"apply"`
	Validation         string                      `json:"validation"`
	Rollback           string                      `json:"rollback"`
	RollbackValidation string                      `json:"rollback_validation"`
	Recovery           guardian.Recovery           `json:"recovery"`
	Intervention       bool                        `json:"operator_intervention_required"`
}

var manualApply = realManualApply
var manualValidate = validateInstalledVersion
var manualPersist = saveLocalEvidence

func realManualApply(args ...string) (automaticupdate.ApplyResult, error) {
	executable, err := os.Executable()
	if err != nil {
		return automaticupdate.ApplyResult{}, err
	}
	// Synchronous interactive sudo is intentional for explicit operator action.
	// Never kill a live privileged child and race rollback against it.
	cmd := exec.Command("/usr/bin/sudo", append([]string{"--", executable, "update"}, args...)...)
	var out boundedReceipt
	cmd.Stdout = &out
	cmd.Stderr = os.Stderr
	err = cmd.Run()
	if out.overflow {
		return automaticupdate.ApplyResult{}, errors.New("helper receipt exceeds bound")
	}
	return decodeApplyReceipt(out.Bytes(), err)
}

func executeManualTransaction(root, from string, pkg updatecore.Package, staged updatecore.Staged, authority, backup string, errout io.Writer) (success bool) {
	r := manualUpdateEvidence{Schema: "qwsg.manual-update/1", At: time.Now().UTC().Format(time.RFC3339Nano), Source: from, Target: pkg.Provenance.Version, Backup: backup, Phase: "prepare", Outcome: "incomplete", Validation: "not_attempted", Rollback: "not_attempted", RollbackValidation: "not_attempted", Recovery: guardian.Recovery{Intent: "unknown", State: "not_changed"}, Intervention: true}
	persist := func() error { return manualPersist(root, "update-result.json", r) }
	if err := persist(); err != nil {
		fmt.Fprintln(errout, "Update refused: transaction intent evidence unavailable.")
		return false
	}
	// An abnormal exit retains incomplete evidence, never an optimistic terminal
	// defer or a restart based on an absent helper receipt.
	defer func() {
		if !success {
			fmt.Fprintf(errout, "Update FAILED: %s; rollback=%s; rollback_validation=%s; Guardian=%s.\n", r.Outcome, r.Rollback, r.RollbackValidation, r.Recovery.State)
			if r.Intervention {
				fmt.Fprintln(errout, "Inspect qwsg update status; preserve recovery evidence, then run qwsg update rollback. Do not start Guardian before package recovery validates.")
			}
		}
	}()
	finish := func(outcome string) bool {
		r.Outcome = outcome
		r.Phase = "finished"
		if err := persist(); err != nil {
			r.Outcome = "evidence_failed"
			r.Intervention = true
			_ = persist()
			fmt.Fprintln(errout, "Update terminal evidence failed; recovery is not committed.")
			return false
		}
		return outcome == "success"
	}
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	active, err := guardianServiceState(ctx)
	if err != nil {
		r.Failure = "guardian_state_unknown"
		r.Intervention = false
		return finish("prepare_failed")
	}
	r.Recovery = guardian.Recovery{Intent: active, StopRequested: active == "active", State: "pending"}
	if err = persist(); err != nil {
		return false
	}
	stopped := false
	if active == "active" {
		stopped = true
		_, err = automaticSystemctl(ctx, "stop", "qwsg-guardian.service")
	}
	if err == nil {
		err = automaticGuardianInactive(ctx)
	}
	if err != nil {
		r.Failure = "guardian_stop_failed"
		if stopped {
			r.Recovery, err = recoverManualGuardian(r.Recovery)
		}
		r.Intervention = err != nil
		return finish("prepare_failed")
	}
	r.Phase = "mutate"
	if err = persist(); err != nil {
		return false
	}
	r.Apply, err = manualApply("privileged-apply-report", "--archive", staged.Archive, "--sidecar", staged.Sidecar, "--version", r.Target, "--sha256", staged.SHA256, "--backup", backup, "--from", from, "--authority", authority)
	if err == nil && (!r.Apply.Known || !r.Apply.MutationStarted || r.Apply.RollbackAttempted) {
		err = errors.New("invalid successful helper receipt")
	}
	if err != nil {
		r.Failure = "apply_failed"
	} else {
		r.Phase = "validate"
		r.Validation = "attempted"
		if err = persist(); err == nil {
			err = updatecore.ValidateApplied(pkg.Root, installedQWSGRoot)
		}
		if err == nil {
			err = manualValidate(r.Target)
		}
		if err == nil {
			err = runSystemctl("daemon-reload")
		}
		if err != nil {
			r.Validation = "failed"
			r.Failure = "validation_failed"
		} else {
			r.Validation = "passed"
			r.Phase = "recover_guardian"
			if err = persist(); err == nil {
				r.Recovery, err = recoverManualGuardian(r.Recovery)
			}
			if err == nil {
				err = saveUpdateRecord(root, localUpdateRecord{Schema: "qwsg.update-local/1", Installed: r.Target, Previous: from, Backup: backup, UpdatedAt: r.At})
			}
			if err == nil {
				r.Intervention = false
				return finish("success")
			}
			r.Failure = "guardian_recovery_failed"
		}
	}
	// All failures after entering the helper stay failed, even after recovery.
	// A possibly started new Guardian must be stopped before restoring files.
	recoveryCtx, cancelRecovery := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancelRecovery()
	current, stateErr := guardianServiceState(recoveryCtx)
	if stateErr == nil && current == "active" {
		_, stateErr = automaticSystemctl(recoveryCtx, "stop", "qwsg-guardian.service")
	}
	if stateErr == nil {
		stateErr = automaticGuardianInactive(recoveryCtx)
	}
	if stateErr != nil {
		r.Recovery.State = "blocked"
		return finish("recovery_incomplete")
	}
	r.Phase = "rollback"
	r.Recovery.State = "blocked"
	if r.Apply.Known && !r.Apply.MutationStarted {
		r.Rollback = "not_required"
	} else if r.Apply.Known && r.Apply.RollbackSucceeded {
		r.Rollback = "succeeded"
	} else {
		r.Rollback = "attempted"
		if err = persist(); err != nil {
			return false
		}
		if err = runSudo("privileged-rollback", "--backup", backup); err != nil {
			r.Rollback = "failed"
			return finish("rollback_failed")
		}
		r.Rollback = "succeeded"
	}
	r.Phase = "validate_rollback"
	r.RollbackValidation = "attempted"
	if err = persist(); err != nil {
		return false
	}
	err = manualValidate(from)
	if err == nil {
		err = runSystemctl("daemon-reload")
	}
	if err != nil {
		r.RollbackValidation = "failed"
		return finish("rollback_validation_failed")
	}
	r.RollbackValidation = "passed"
	r.Phase = "recover_guardian"
	if err = persist(); err != nil {
		return false
	}
	r.Recovery, err = recoverManualGuardian(r.Recovery)
	if err != nil {
		return finish("recovery_failed")
	}
	r.Intervention = false
	return finish("update_failed_recovered")
}

func recoverManualGuardian(r guardian.Recovery) (guardian.Recovery, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	return recoverGuardian(ctx, r, true)
}
func loadManualEvidence(root string) (manualUpdateEvidence, error) {
	var r manualUpdateEvidence
	data, err := evidenceio.ReadFile(filepath.Join(root, "update-result.json"), 16384)
	if err != nil {
		return r, err
	}
	if json.Unmarshal(data, &r) != nil || r.Schema != "qwsg.manual-update/1" || !validBackup(r.Backup) {
		return r, errors.New("invalid manual update evidence")
	}
	if _, err := updatecore.ParseVersion(r.Source); err != nil {
		return r, errors.New("invalid source evidence")
	}
	if _, err := updatecore.ParseVersion(r.Target); err != nil {
		return r, errors.New("invalid target evidence")
	}
	if _, err := time.Parse(time.RFC3339Nano, r.At); err != nil {
		return r, errors.New("invalid evidence timestamp")
	}
	if r.Outcome == "success" && (r.Validation != "passed" || !r.Apply.Known || !r.Apply.MutationStarted || r.Apply.RollbackAttempted || r.Intervention || (r.Recovery.State != "verified_running" && r.Recovery.State != "preserved_stopped")) {
		return r, errors.New("conflicting success evidence")
	}
	return r, nil
}
func manualTransactionReady(root string) error {
	r, err := loadManualEvidence(root)
	if os.IsNotExist(err) {
		return nil
	}
	if err != nil {
		return err
	}
	switch r.Outcome {
	case "success", "update_failed_recovered", "prepare_failed", "resolved_by_rollback":
		if !r.Intervention {
			return nil
		}
	}
	return errors.New("incomplete manual update recovery")
}
func writeManualEvidenceStatus(root string, out io.Writer) error {
	r, err := loadManualEvidence(root)
	if os.IsNotExist(err) {
		return nil
	}
	if err != nil {
		return err
	}
	fmt.Fprintf(out, "Manual update: %s\nPhase: %s\nFailure: %s\nPackage validation: %s\nRollback: %s (%s)\nGuardian recovery: %s (intent=%s)\nOperator intervention: %t\n", safeText(r.Outcome), safeText(r.Phase), safeText(r.Failure), safeText(r.Validation), safeText(r.Rollback), safeText(r.RollbackValidation), safeText(r.Recovery.State), safeText(r.Recovery.Intent), r.Intervention)
	if r.Intervention {
		fmt.Fprintln(out, "Next action: preserve evidence; qwsg update rollback. Do not start Guardian before package recovery validates.")
	}
	return nil
}

func writeRollbackEvidenceStatus(root string, out io.Writer) error {
	data, err := evidenceio.ReadFile(filepath.Join(root, "rollback-result.json"), 16384)
	if os.IsNotExist(err) {
		return nil
	}
	if err != nil {
		return err
	}
	var r rollbackEvidence
	if json.Unmarshal(data, &r) != nil || r.Schema != "qwsg.rollback-result/1" {
		return errors.New("invalid rollback evidence")
	}
	fmt.Fprintf(out, "Manual rollback: %s; package=%s; Guardian=%s\n", safeText(r.Outcome), safeText(r.Package), safeText(r.Recovery.State))
	if r.Outcome != "success" {
		fmt.Fprintln(out, "Next action: preserve evidence; correct the reported failure, then retry qwsg update rollback.")
	}
	return nil
}
