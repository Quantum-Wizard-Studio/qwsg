package main

import (
	"bytes"
	"context"
	"crypto/ed25519"
	"crypto/sha256"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"reflect"
	"strings"
	"testing"
	"time"

	"quantumwizard.hu/qwsg/internal/automaticupdate"
	"quantumwizard.hu/qwsg/internal/installation"
	"quantumwizard.hu/qwsg/internal/productcapability"
	"quantumwizard.hu/qwsg/internal/releasediscovery"
	"quantumwizard.hu/qwsg/internal/update"
	"quantumwizard.hu/qwsg/internal/updateauthority"
	"quantumwizard.hu/qwsg/internal/updatepolicy"
)

// This host substitutes privilege/service boundaries only. Authentication,
// extraction, backup, application and rollback use the actual common engine.
type automaticFixtureHost struct {
	t                                                     *testing.T
	req                                                   automaticupdate.Request
	dest, backup                                          string
	preflightError, postError, rollbackError, commitError bool
	helperTamper, applyFailure                            bool
	applyCalls, rollbackCalls                             int
	onPreflight                                           func()
	cancel                                                context.CancelFunc
}

func (h *automaticFixtureHost) Preflight(context.Context, string, string) error {
	if h.onPreflight != nil {
		h.onPreflight()
	}
	if h.preflightError {
		return errors.New("injected preflight")
	}
	return nil
}
func (h *automaticFixtureHost) Apply(_ context.Context, p automaticupdate.PackageInput) (automaticupdate.ApplyResult, error) {
	h.applyCalls++
	receipt := automaticupdate.ApplyResult{Known: true}
	// Repeat the privileged helper's re-staging and reauthentication, including
	// a deterministic attack between ordinary and privileged verification.
	metadata, err := updateauthority.ReadMetadata(p.AuthorityPath)
	if h.helperTamper {
		metadata = []byte(`{}`)
	}
	c, err := updateauthority.Authorize(metadata, h.req.Verifier, h.req.Evaluator, "stable", h.req.Platform, p.TargetVersion, h.req.Now)
	if err != nil {
		return receipt, err
	}
	parent := h.t.TempDir()
	_ = os.Chmod(parent, 0700)
	staged, err := update.StageLocal(p.Staged.Archive, p.Staged.Sidecar, p.TargetVersion, parent)
	if err != nil {
		return receipt, err
	}
	pkg, err := c.VerifyStaged(staged)
	if err != nil {
		return receipt, err
	}
	if h.applyFailure {
		// Sorted destination order replaces binary first, then fails a later source.
		if err = os.Remove(filepath.Join(pkg.Root, "README.md")); err != nil {
			h.t.Fatal(err)
		}
	}

	tx, err := update.Apply(pkg.Root, h.dest, h.backup, p.SourceVersion)
	receipt.MutationStarted = tx.MutationStarted
	receipt.RollbackAttempted = tx.RollbackAttempted
	receipt.RollbackSucceeded = tx.RollbackSucceeded
	if h.cancel != nil {
		h.cancel()
	}
	return receipt, err
}
func (h *automaticFixtureHost) Validate(ctx context.Context, version string) error {
	if ctx.Err() != nil {
		return ctx.Err()
	}
	if h.postError && version == h.req.TargetVersion {
		return errors.New("injected health failure")
	}
	data, err := os.ReadFile(filepath.Join(h.dest, "usr/local/share/doc/qwsg/RELEASE.json"))
	if err != nil {
		return err
	}
	var p update.Provenance
	if json.Unmarshal(data, &p) != nil || p.Version != version {
		return errors.New("installed provenance mismatch")
	}
	data, err = os.ReadFile(filepath.Join(h.dest, "usr/local/bin/qwsg"))
	if err != nil || string(data) != version {
		return errors.New("installed binary mismatch")
	}
	return nil
}
func (h *automaticFixtureHost) Rollback(ctx context.Context) error {
	h.rollbackCalls++
	if ctx.Err() != nil {
		return ctx.Err()
	}
	if h.rollbackError {
		return errors.New("injected rollback failure")
	}
	return update.Rollback(h.dest, h.backup)
}
func (h *automaticFixtureHost) Commit(context.Context, string, string) error {
	if h.commitError {
		return errors.New("injected record failure")
	}
	return nil
}
func automaticFixture(t *testing.T) (automaticupdate.Request, *automaticFixtureHost, releasediscovery.Index) {
	t.Helper()
	dir := t.TempDir()
	_ = os.Chmod(dir, 0700)
	target := "1.84.0"
	commit := strings.Repeat("b", 40)
	archive, digest := forwardArchive(t, dir, target, commit, []byte(target))
	forwardWrite(t, archive+".sha256", []byte(digest+"  "+filepath.Base(archive)+"\n"), 0600)
	info, err := os.Stat(archive)
	if err != nil {
		t.Fatal(err)
	}
	now := time.Date(2026, 9, 14, 12, 0, 0, 0, time.UTC)
	stamp := now.Format(time.RFC3339)
	declaration := update.CompatibilityDeclaration{Schema: update.MigrationSchema, SourceVersion: "1.3.1", TargetVersion: target, Platform: "linux-amd64", Capability: update.PreservePackageV1, ConfigurationSchema: "1.0", GuardianSchema: "1.0", SchedulerSchema: "1.0", OperatorState: "1.0-1.2"}
	index := releasediscovery.Index{Schema: releasediscovery.CapabilitySchema, Product: "qwsg", GeneratedAt: stamp, Channels: []releasediscovery.Channel{{Name: "stable", Releases: []releasediscovery.Release{{Version: target, PublishedAt: stamp, Status: "active", SourceCommit: commit, ReleaseNotesURL: "https://example.invalid/notes", MinimumSourceVersion: "1.3.1", Compatibility: []update.CompatibilityDeclaration{declaration}, Artifacts: []releasediscovery.Artifact{{Platform: "linux-amd64", Name: filepath.Base(archive), URL: "https://example.invalid/" + filepath.Base(archive), Size: info.Size(), SHA256: digest}}}}}}}
	caps, err := proTestAuthority()
	if err != nil {
		t.Fatal(err)
	}
	key := ed25519.NewKeyFromSeed(make([]byte, ed25519.SeedSize))
	verifier, err := releasediscovery.NewVerifier(map[string]ed25519.PublicKey{"forward-test": key.Public().(ed25519.PublicKey)})
	if err != nil {
		t.Fatal(err)
	}
	evaluator, _ := releasediscovery.NewEvaluator(func(string) installation.Result {
		return installation.Result{State: installation.VerifiedSupported, Version: "1.3.1"}
	})
	req := automaticupdate.Request{Capabilities: caps, Policy: updatepolicy.Request{Mode: updatepolicy.Automatic}, Metadata: forwardSign(t, index, true), Verifier: verifier, Evaluator: evaluator, Now: now, SourceVersion: "1.3.1", TargetVersion: target, Platform: "linux-amd64", StageParent: dir, LocalArchive: archive}
	host := &automaticFixtureHost{t: t, req: req, dest: filepath.Join(dir, "installed"), backup: filepath.Join(dir, "backup")}
	for path, body := range forwardInstalledFiles("1.3.1", strings.Repeat("a", 40), []byte("1.3.1")) {
		forwardWrite(t, filepath.Join(host.dest, path), body, 0644)
	}
	forwardWrite(t, filepath.Join(host.dest, "private/config"), []byte("preserved"), 0600)
	return req, host, index
}
func installationDigest(t *testing.T, root string) map[string]string {
	t.Helper()
	result := map[string]string{}
	err := filepath.WalkDir(root, func(path string, e os.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if e.IsDir() {
			return nil
		}
		b, err := os.ReadFile(path)
		if err != nil {
			return err
		}
		info, err := e.Info()
		if err != nil {
			return err
		}
		rel, _ := filepath.Rel(root, path)
		result[rel] = fmt.Sprintf("%x:%o", sha256.Sum256(b), info.Mode().Perm())
		return nil
	})
	if err != nil {
		t.Fatal(err)
	}
	return result
}

func TestAutomaticUpdateRefusesBeforeMutation(t *testing.T) {
	for _, scenario := range []string{"community", "zero capability", "manual", "unknown policy", "conflicting policy", "unsigned", "tampered", "unknown migration", "source mismatch", "target mismatch", "platform mismatch", "provenance mismatch", "artifact mismatch", "invalid package", "preflight", "helper rejection", "watermark"} {
		t.Run(scenario, func(t *testing.T) {
			req, host, index := automaticFixture(t)
			switch scenario {
			case "community":
				req.Capabilities, _ = productcapability.Resolve(nil)
			case "zero capability":
				req.Capabilities = productcapability.Set{}
			case "manual":
				req.Policy.Mode = updatepolicy.Manual
			case "unknown policy":
				req.Policy.Mode = "unknown"
			case "conflicting policy":
				req.Policy.Notify = true
			case "unsigned":
				req.Metadata = forwardSign(t, index, false)
			case "tampered":
				req.Metadata = bytes.Replace(req.Metadata, []byte("preserve-package-v1"), []byte("preserve-package-v2"), 1)
			case "unknown migration":
				index.Channels[0].Releases[0].Compatibility[0].Capability = "unknown-v1"
				req.Metadata = forwardSign(t, index, true)
			case "source mismatch":
				req.SourceVersion = "1.3.0"
			case "target mismatch":
				req.TargetVersion = "1.84.1"
			case "platform mismatch":
				req.Platform = "linux-arm64"
			case "provenance mismatch":
				index.Channels[0].Releases[0].SourceCommit = strings.Repeat("c", 40)
				req.Metadata = forwardSign(t, index, true)
			case "artifact mismatch":
				index.Channels[0].Releases[0].Artifacts[0].SHA256 = strings.Repeat("0", 64)
				req.Metadata = forwardSign(t, index, true)
			case "invalid package":
				b := []byte("invalid archive")
				hash := fmt.Sprintf("%x", sha256.Sum256(b))
				forwardWrite(t, req.LocalArchive, b, 0600)
				forwardWrite(t, req.LocalArchive+".sha256", []byte(hash+"  "+filepath.Base(req.LocalArchive)+"\n"), 0600)
				index.Channels[0].Releases[0].Artifacts[0].SHA256 = hash
				index.Channels[0].Releases[0].Artifacts[0].Size = int64(len(b))
				req.Metadata = forwardSign(t, index, true)
			case "preflight":
				host.preflightError = true
			case "helper rejection":
				host.helperTamper = true
			case "watermark":
				req.NotBefore = req.Now.Add(time.Second)
			}
			before := installationDigest(t, host.dest)
			r, err := automaticupdate.Run(context.Background(), req, host)
			if err == nil || r.State != automaticupdate.Failure || r.MutationStarted || r.RollbackAttempted || host.rollbackCalls != 0 || r.FailureCategory == "" {
				t.Fatalf("unsafe refusal: %+v err=%v", r, err)
			}
			if !reflect.DeepEqual(before, installationDigest(t, host.dest)) {
				t.Fatal("refusal changed installed files or modes")
			}
			if scenario != "helper rejection" && host.applyCalls != 0 {
				t.Fatal("refusal reached privileged apply")
			}
		})
	}
}

func TestAutomaticUpdateTransactionOutcomes(t *testing.T) {
	for _, scenario := range []string{"success", "during apply", "post validation", "commit failure", "rollback failure", "cancel after apply"} {
		t.Run(scenario, func(t *testing.T) {
			req, host, _ := automaticFixture(t)
			before := installationDigest(t, host.dest)
			ctx, cancel := context.WithCancel(context.Background())
			defer cancel()
			switch scenario {
			case "during apply":
				host.applyFailure = true
			case "post validation":
				host.postError = true
			case "commit failure":
				host.commitError = true
			case "rollback failure":
				host.postError = true
				host.rollbackError = true
			case "cancel after apply":
				host.cancel = cancel
			}
			r, err := automaticupdate.Run(ctx, req, host)
			if !r.CapabilityAllowed || !r.PolicyAllowed || !r.MutationStarted || r.SourceVersion != "1.3.1" || r.TargetVersion != "1.84.0" {
				t.Fatalf("missing evidence: %+v", r)
			}
			if scenario == "success" {
				want := []automaticupdate.State{automaticupdate.Idle, automaticupdate.PolicyCheck, automaticupdate.CandidateCheck, automaticupdate.EligibilityCheck, automaticupdate.Staging, automaticupdate.Preflight, automaticupdate.Backup, automaticupdate.Apply, automaticupdate.PostUpdateValidation, automaticupdate.Success}
				if err != nil || r.RollbackAttempted || !reflect.DeepEqual(r.Stages, want) {
					t.Fatalf("bad success: %+v %v", r, err)
				}
				if err = host.Validate(ctx, req.TargetVersion); err != nil {
					t.Fatal(err)
				}
			} else {
				if err == nil || !r.RollbackAttempted {
					t.Fatalf("missing rollback: %+v %v", r, err)
				}
				if scenario == "rollback failure" {
					if r.State != automaticupdate.RollbackFailure || r.RollbackResult != "failed" {
						t.Fatalf("hidden degraded outcome: %+v", r)
					}
				} else {
					if r.State != automaticupdate.RollbackSuccess || r.RollbackResult != "succeeded" || !reflect.DeepEqual(before, installationDigest(t, host.dest)) {
						t.Fatalf("rollback did not restore exact files/modes: %+v", r)
					}
				}
				if scenario == "during apply" && host.rollbackCalls != 0 {
					t.Fatal("duplicate rollback after common engine already restored")
				}
			}
			data, _ := os.ReadFile(filepath.Join(host.dest, "private/config"))
			if string(data) != "preserved" {
				t.Fatal("private configuration changed")
			}
		})
	}
}

func TestAutomaticUpdateReentryAndConcurrentConflict(t *testing.T) {
	req, host, _ := automaticFixture(t)
	entered, release := make(chan struct{}), make(chan struct{})
	host.onPreflight = func() { close(entered); <-release }
	done := make(chan error, 1)
	go func() { _, err := automaticupdate.Run(context.Background(), req, host); done <- err }()
	<-entered
	r, err := automaticupdate.Run(context.Background(), req, host)
	if err == nil || r.FailureCategory != "transaction_conflict" || r.MutationStarted {
		t.Fatalf("conflicting transaction entered: %+v %v", r, err)
	}
	close(release)
	if err = <-done; err != nil {
		t.Fatal(err)
	}
	// A separate invocation after completion is admitted and then fails its real
	// source precondition in production; the lock itself must not remain held.
	req2, host2, _ := automaticFixture(t)
	host2.onPreflight = func() {
		r, err := automaticupdate.Run(context.Background(), req2, host2)
		if err == nil || r.FailureCategory != "transaction_conflict" {
			t.Errorf("reentry accepted: %+v", r)
		}
	}
	if _, err = automaticupdate.Run(context.Background(), req2, host2); err != nil {
		t.Fatal(err)
	}
}

func TestAutomaticCommunityCompositionDoesNotFetch(t *testing.T) {
	oldCaps, oldFetch := installationCapabilities, updateMetadataFetch
	defer func() { installationCapabilities = oldCaps; updateMetadataFetch = oldFetch }()
	installationCapabilities = func() (productcapability.Set, error) { return productcapability.Resolve(nil) }
	updateMetadataFetch = func(context.Context) ([]byte, error) {
		t.Fatal("Community automatic entry fetched metadata")
		return nil, nil
	}
	r, err := executeAutomaticUpdate(context.Background())
	if err == nil || r.CapabilityAllowed || r.MutationStarted {
		t.Fatalf("Community authorized: %+v %v", r, err)
	}
}

func TestAutomaticHelperReceiptFailClosed(t *testing.T) {
	for _, data := range []string{"", `{}`, `{"known":true,"unexpected":true}`, `{"known":true} {}`, `{"known":true,"rollback_succeeded":true}`} {
		receipt, err := decodeApplyReceipt([]byte(data), nil)
		if err == nil || receipt.Known {
			t.Fatalf("invalid receipt accepted: %s", data)
		}
	}
	var out, diagnostic bytes.Buffer
	code := runPrivilegedApplyReport(nil, &out, &diagnostic)
	receipt, err := decodeApplyReceipt(out.Bytes(), errors.New("helper refusal"))
	if code == 0 || err == nil || !receipt.Known || receipt.MutationStarted {
		t.Fatalf("helper pre-mutation refusal lost: %+v %v", receipt, err)
	}
}

// Receipt-only faults model helper transport loss and an already failed internal
// rollback. Neither condition may be flattened into ordinary update failure.
type automaticReceiptHost struct {
	*automaticFixtureHost
	receipt automaticupdate.ApplyResult
}

func (h automaticReceiptHost) Apply(context.Context, automaticupdate.PackageInput) (automaticupdate.ApplyResult, error) {
	return h.receipt, errors.New("injected helper result")
}
func TestAutomaticDegradedHelperResults(t *testing.T) {
	for _, receipt := range []automaticupdate.ApplyResult{
		{},
		{Known: true, MutationStarted: true, RollbackAttempted: true},
	} {
		req, fixture, _ := automaticFixture(t)
		h := automaticReceiptHost{automaticFixtureHost: fixture, receipt: receipt}
		r, err := automaticupdate.Run(context.Background(), req, h)
		if err == nil || r.State != automaticupdate.RollbackFailure || !r.MutationStarted || !r.RollbackAttempted || r.RollbackResult != "failed" {
			t.Fatalf("degraded helper failure hidden: %+v %v", r, err)
		}
	}
}
