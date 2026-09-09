package main

import (
	"archive/tar"
	"bytes"
	"compress/gzip"
	"context"
	"crypto/ed25519"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"sort"
	"strings"
	"testing"
	"time"

	"quantumwizard.hu/qwsg/internal/installation"
	"quantumwizard.hu/qwsg/internal/releasediscovery"
	"quantumwizard.hu/qwsg/internal/update"
	"quantumwizard.hu/qwsg/internal/updateauthority"
)

// TestMain is present only in test-linked binaries. The acceptance executable
// substitutes host boundaries, never migration planning, authentication,
// staging, package verification, preflight, apply or rollback.
func TestMain(m *testing.M) {
	if root := os.Getenv("QWSG_FORWARD_FIXTURE_ROOT"); root != "" {
		version = "1.3.1"
		buildCommit = strings.Repeat("a", 40)
		buildDate = "2026-09-09T00:00:00Z"
		installedQWSGRoot = root
		installedQWSGBinary = filepath.Join(root, "usr/local/bin/qwsg")
		updateSystemRoot = filepath.Join(root, "var/lib/qwsg")
		updateRollbackRoot = filepath.Join(updateSystemRoot, "rollback")
		updateEffectiveUID = func() int { return 1000 }
		updateMetadataFetch = func(context.Context) ([]byte, error) {
			return updateauthority.ReadMetadata(os.Getenv("QWSG_FORWARD_METADATA"))
		}
		updateVerifier = func() (releasediscovery.Verifier, error) {
			key := ed25519.NewKeyFromSeed(make([]byte, ed25519.SeedSize))
			return releasediscovery.NewVerifier(map[string]ed25519.PublicKey{"forward-test": key.Public().(ed25519.PublicKey)})
		}
		updateHTTPClient = func() *http.Client {
			return &http.Client{Transport: forwardTransport{directory: os.Getenv("QWSG_FORWARD_ARTIFACTS")}}
		}
		commandState = func(string) string { return "yes" }
		runSystemctl = func(action string) error {
			if action == "daemon-reload" && os.Getenv("QWSG_FORWARD_FAIL_POST") == "1" {
				return fmt.Errorf("injected post-apply failure")
			}
			return nil
		}
		runSudo = func(args ...string) error {
			// Invoke the old executable's actual privileged dispatch with its real gates.
			if os.Getenv("QWSG_FORWARD_HELPER_TAMPER") == "1" && len(args) > 0 && args[0] == "privileged-apply" {
				for i := 1; i+1 < len(args); i += 2 {
					if args[i] == "--authority" {
						data, err := os.ReadFile(args[i+1])
						if err != nil {
							return err
						}
						var index releasediscovery.Index
						if err = json.Unmarshal(data, &index); err != nil {
							return err
						}
						index.Signatures = nil
						data, err = json.Marshal(index)
						if err != nil {
							return err
						}
						if err = os.WriteFile(args[i+1], data, 0600); err != nil {
							return err
						}
					}
				}
			}
			old := updateEffectiveUID
			updateEffectiveUID = func() int { return 0 }
			defer func() { updateEffectiveUID = old }()
			if err := os.WriteFile(filepath.Join(root, "helper-old-client"), []byte(version), 0600); err != nil {
				return err
			}
			if code := runUpdate(args, os.Stdout, os.Stderr); code != 0 {
				return fmt.Errorf("helper exit %d", code)
			}
			return nil
		}
		verifier, err := updateVerifier()
		if err != nil {
			panic(err)
		}
		discoverer, err := releasediscovery.NewDiscoverer(forwardSource{}, verifier, installedUpdateEvaluator())
		if err != nil {
			panic(err)
		}
		updateAwarenessChecker = discoverer
		os.Exit(run(os.Args[1:], os.Stdout, os.Stderr))
	}
	os.Exit(m.Run())
}

type forwardSource struct{}

func (forwardSource) Fetch(ctx context.Context, _ releasediscovery.FetchRequest) (releasediscovery.FetchResult, error) {
	p, e := updateMetadataFetch(ctx)
	return releasediscovery.FetchResult{Manifest: p, Evidence: releasediscovery.SourceEvidence{SourceID: releasediscovery.ProductionSourceID, TransportAuthenticated: true}}, e
}

type forwardTransport struct{ directory string }

func (f forwardTransport) RoundTrip(r *http.Request) (*http.Response, error) {
	data, err := os.ReadFile(filepath.Join(f.directory, filepath.Base(r.URL.Path)))
	if err != nil {
		return nil, err
	}
	return &http.Response{StatusCode: 200, Body: io.NopCloser(bytes.NewReader(data)), Header: make(http.Header)}, nil
}

func TestOldBinaryForwardAuthenticatedUpdate(t *testing.T) {
	repo, err := filepath.Abs("../..")
	if err != nil {
		t.Fatal(err)
	}
	work := t.TempDir()
	old := filepath.Join(work, "old-client")
	build := exec.Command("go", "test", "-c", "-o", old, "./cmd/qwsg")
	build.Dir = repo
	if out, err := build.CombinedOutput(); err != nil {
		t.Fatalf("build old client: %v %s", err, out)
	}
	oldBytes, err := os.ReadFile(old)
	if err != nil {
		t.Fatal(err)
	}
	frozen := sha256.Sum256(oldBytes)
	// Select the target only after the old client artifact is frozen. Neither
	// this target nor a target-specific route is added to the migration registry.
	target := fmt.Sprintf("1.%d.%d", int(frozen[0])+20, int(frozen[1])+20)
	if _, err := update.PlanMigration("1.3.1", target); err == nil {
		t.Fatal("future target unexpectedly hard-coded")
	}
	newer := filepath.Join(work, "target-client")
	targetCommit := strings.Repeat("b", 40)
	build = exec.Command("go", "build", "-trimpath", "-buildvcs=false", "-ldflags", "-X main.version="+target+" -X main.buildCommit="+targetCommit+" -X main.buildDate=2026-09-09T00:00:00Z", "-o", newer, "./cmd/qwsg")
	build.Dir = repo
	if out, err := build.CombinedOutput(); err != nil {
		t.Fatalf("build future artifact: %v %s", err, out)
	}
	targetBytes, err := os.ReadFile(newer)
	if err != nil {
		t.Fatal(err)
	}
	archive, digest := forwardArchive(t, work, target, targetCommit, targetBytes)
	info, err := os.Stat(archive)
	if err != nil {
		t.Fatal(err)
	}
	now := time.Now().UTC().Truncate(time.Second).Format(time.RFC3339)
	declaration := update.CompatibilityDeclaration{Schema: update.MigrationSchema, SourceVersion: "1.3.1", TargetVersion: target, Platform: "linux-amd64", Capability: update.PreservePackageV1, ConfigurationSchema: "1.0", GuardianSchema: "1.0", SchedulerSchema: "1.0", OperatorState: "1.0-1.2"}
	index := releasediscovery.Index{Schema: releasediscovery.CapabilitySchema, Product: "qwsg", GeneratedAt: now, Channels: []releasediscovery.Channel{{Name: "stable", Releases: []releasediscovery.Release{{Version: target, PublishedAt: now, Status: "active", SourceCommit: targetCommit, ReleaseNotesURL: "https://example.invalid/notes", MinimumSourceVersion: "1.3.1", Compatibility: []update.CompatibilityDeclaration{declaration}, Artifacts: []releasediscovery.Artifact{{Platform: "linux-amd64", Name: filepath.Base(archive), URL: "https://example.invalid/" + filepath.Base(archive), Size: info.Size(), SHA256: digest}}}}}}}
	for _, scenario := range []string{"online", "archive", "rollback-after-failure", "unknown-capability", "unsigned", "source-mismatch", "provenance-mismatch", "digest-mismatch", "archive-missing-authority", "helper-reauthentication", "archive-forged-sidecar"} {
		t.Run(scenario, func(t *testing.T) {
			root := filepath.Join(t.TempDir(), "root")
			if err := os.Mkdir(root, 0700); err != nil {
				t.Fatal(err)
			}
			for rel, body := range forwardInstalledFiles("1.3.1", strings.Repeat("a", 40), oldBytes) {
				p := filepath.Join(root, rel)
				forwardWrite(t, p, body, 0644)
				if strings.HasSuffix(rel, "/bin/qwsg") {
					if err := os.Chmod(p, 0755); err != nil {
						t.Fatal(err)
					}
				}
			}
			state := filepath.Join(root, "state")
			if err := os.Mkdir(state, 0700); err != nil {
				t.Fatal(err)
			}
			preserved := filepath.Join(state, "operator-state-preserved")
			forwardWrite(t, preserved, []byte("preserve user state"), 0600)
			cfg := filepath.Join(root, "config")
			if err := os.Mkdir(cfg, 0700); err != nil {
				t.Fatal(err)
			}
			// Clone the signed input so subtests cannot contaminate authority.
			data, _ := json.Marshal(index)
			var candidate releasediscovery.Index
			if err := json.Unmarshal(data, &candidate); err != nil {
				t.Fatal(err)
			}
			release := &candidate.Channels[0].Releases[0]
			switch scenario {
			case "unknown-capability":
				release.Compatibility[0].Capability = "unimplemented-v9"
			case "source-mismatch":
				release.Compatibility[0].SourceVersion = "1.3.2"
			case "provenance-mismatch":
				release.SourceCommit = strings.Repeat("c", 40)
			case "digest-mismatch":
				release.Artifacts[0].SHA256 = strings.Repeat("d", 64)
			}
			signed := forwardSign(t, candidate, scenario != "unsigned")
			metadata := filepath.Join(root, "release-index.json")
			forwardWrite(t, metadata, signed, 0600)
			env := append(os.Environ(), "QWSG_FORWARD_FIXTURE_ROOT="+root, "QWSG_FORWARD_METADATA="+metadata, "QWSG_FORWARD_ARTIFACTS="+work, "QWSG_STATE_DIR="+state, "XDG_CONFIG_HOME="+cfg)
			runOld := func(args ...string) ([]byte, error) {
				cmd := exec.Command(old, args...)
				cmd.Env = env
				return cmd.CombinedOutput()
			}
			if scenario == "online" {
				out, err := runOld("update", "check")
				if err != nil || !bytes.Contains(out, []byte("Update awareness: update_available\n")) {
					t.Fatalf("discovery %v %s", err, out)
				}
				got, _ := os.ReadFile(filepath.Join(root, "usr/local/bin/qwsg"))
				if !bytes.Equal(got, oldBytes) {
					t.Fatal("discovery installed release")
				}
				if _, err := os.Stat(filepath.Join(root, "helper-old-client")); !os.IsNotExist(err) {
					t.Fatal("discovery invoked helper")
				}
			}
			args := []string{"update"}
			if scenario == "archive" || scenario == "archive-missing-authority" || scenario == "archive-forged-sidecar" {
				local := filepath.Join(root, filepath.Base(archive))
				b, err := os.ReadFile(archive)
				if err != nil {
					t.Fatal(err)
				}
				localDigest := digest
				if scenario == "archive-forged-sidecar" {
					b[len(b)/2] ^= 1
					sum := sha256.Sum256(b)
					localDigest = hex.EncodeToString(sum[:])
				}
				forwardWrite(t, local, b, 0600)
				forwardWrite(t, local+".sha256", []byte(localDigest+"  "+filepath.Base(archive)+"\n"), 0600)
				if scenario == "archive" || scenario == "archive-forged-sidecar" {
					forwardWrite(t, local+".release-index.json", signed, 0600)
				}
				args = append(args, "--archive", local, "--version", target)
			}
			if scenario == "helper-reauthentication" {
				env = append(env, "QWSG_FORWARD_HELPER_TAMPER=1")
			}
			if scenario == "rollback-after-failure" {
				env = append(env, "QWSG_FORWARD_FAIL_POST=1")
			}
			out, err := runOld(args...)
			success := scenario == "online" || scenario == "archive"
			if success != (err == nil) {
				t.Fatalf("update err=%v output=%s", err, out)
			}
			if !success {
				expected := map[string]string{
					"unknown-capability":        "authenticated compatible migration capability unavailable",
					"unsigned":                  "authenticated compatible migration capability unavailable",
					"source-mismatch":           "authenticated compatible migration capability unavailable",
					"provenance-mismatch":       "authenticated package verification failed",
					"digest-mismatch":           "candidate acquisition or integrity verification failed",
					"archive-missing-authority": "authenticated release metadata unavailable",
					"helper-reauthentication":   "privileged migration authority refused",
					"archive-forged-sidecar":    "authenticated package verification failed",
					"rollback-after-failure":    "automatic package rollback was attempted",
				}[scenario]
				if expected == "" || !bytes.Contains(out, []byte(expected)) {
					t.Fatalf("wrong refusal gate for %s: %s", scenario, out)
				}
			}

			got, readErr := os.ReadFile(filepath.Join(root, "usr/local/bin/qwsg"))
			if readErr != nil {
				t.Fatal(readErr)
			}
			if success {
				if !bytes.Equal(got, targetBytes) || !bytes.Contains(out, []byte("updated safely")) {
					t.Fatalf("not installed: %s", out)
				}
				helper, _ := os.ReadFile(filepath.Join(root, "helper-old-client"))
				if string(helper) != "1.3.1" {
					t.Fatal("old client did not execute helper")
				}
				out, err = runOld("update", "rollback")
				if err != nil {
					t.Fatalf("rollback %v %s", err, out)
				}
				got, _ = os.ReadFile(filepath.Join(root, "usr/local/bin/qwsg"))
			}
			if !bytes.Equal(got, oldBytes) {
				t.Fatal("refusal/rollback did not preserve old binary")
			}
			if scenario == "rollback-after-failure" && !bytes.Contains(out, []byte("automatic package rollback")) {
				t.Fatalf("failure never reached rollback: %s", out)
			}
			p, err := os.ReadFile(preserved)
			if err != nil || string(p) != "preserve user state" {
				t.Fatal("operator state changed")
			}
			classifier := installation.Classify(installation.Options{Root: root, RunVersion: func(_ context.Context, binary string) ([]byte, error) {
				cmd := exec.Command(binary, "version")
				cmd.Env = env
				return cmd.Output()
			}})
			if classifier.State != installation.VerifiedSupported || classifier.Version != "1.3.1" {
				t.Fatalf("final identity %+v", classifier)
			}
		})
	}
	after, err := os.ReadFile(old)
	if err != nil || sha256.Sum256(after) != frozen {
		t.Fatal("frozen old executable changed")
	}
	t.Logf("old executable frozen before target selection; future target %s; old migration table has no route", target)
}

func forwardWrite(t *testing.T, p string, b []byte, mode os.FileMode) {
	t.Helper()
	if err := os.MkdirAll(filepath.Dir(p), 0700); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(p, b, mode); err != nil {
		t.Fatal(err)
	}
}
func forwardInstalledFiles(v, commit string, binary []byte) map[string][]byte {
	files := map[string][]byte{"usr/local/bin/qwsg": binary, "usr/local/lib/systemd/user/qwsg-guardian.service": []byte("[Service]\nExecStart=/usr/local/bin/qwsg\n")}
	for _, name := range []string{"README.md", "INSTALL.md", "LICENSE", "CHANGELOG.md", "qwsg-config.json"} {
		files["usr/local/share/doc/qwsg/"+name] = []byte("fixture\n")
	}
	files["usr/local/share/doc/qwsg/RELEASE.json"], _ = json.Marshal(update.Provenance{Schema: "qwsg.release/1", Version: v, Commit: commit, Built: "2026-09-09T00:00:00Z", Platform: "linux-amd64"})
	return files
}
func forwardArchive(t *testing.T, dir, v, commit string, binary []byte) (string, string) {
	t.Helper()
	files := map[string][]byte{"bin/qwsg": binary, "install.sh": []byte("#!/bin/sh\nexit 97\n"), "uninstall.sh": []byte("#!/bin/sh\nexit 97\n")}
	for rel, b := range forwardInstalledFiles(v, commit, binary) {
		if rel == "usr/local/bin/qwsg" {
			continue
		}
		name := filepath.Base(rel)
		if strings.HasSuffix(rel, ".service") {
			name = "lib/systemd/user/qwsg-guardian.service"
		}
		files[name] = b
	}
	names := []string{}
	for name := range files {
		names = append(names, name)
	}
	sort.Strings(names)
	var manifest strings.Builder
	for _, name := range names {
		h := sha256.Sum256(files[name])
		fmt.Fprintf(&manifest, "%x  %s\n", h, name)
	}
	files["MANIFEST.sha256"] = []byte(manifest.String())
	names = append(names, "MANIFEST.sha256")
	sort.Strings(names)
	var b bytes.Buffer
	gz := gzip.NewWriter(&b)
	tr := tar.NewWriter(gz)
	root := "qwsg-" + v + "-linux-amd64"
	if err := tr.WriteHeader(&tar.Header{Name: root + "/", Typeflag: tar.TypeDir, Mode: 0755}); err != nil {
		t.Fatal(err)
	}
	for _, name := range names {
		if err := tr.WriteHeader(&tar.Header{Name: root + "/" + name, Mode: 0644, Size: int64(len(files[name]))}); err != nil {
			t.Fatal(err)
		}
		if _, err := tr.Write(files[name]); err != nil {
			t.Fatal(err)
		}
	}
	if err := tr.Close(); err != nil {
		t.Fatal(err)
	}
	if err := gz.Close(); err != nil {
		t.Fatal(err)
	}
	archive := filepath.Join(dir, root+".tar.gz")
	forwardWrite(t, archive, b.Bytes(), 0600)
	hash := sha256.Sum256(b.Bytes())
	return archive, hex.EncodeToString(hash[:])
}
func forwardSign(t *testing.T, i releasediscovery.Index, sign bool) []byte {
	t.Helper()
	message, err := releasediscovery.SigningBytes(i)
	if err != nil {
		t.Fatal(err)
	}
	if sign {
		key := ed25519.NewKeyFromSeed(make([]byte, ed25519.SeedSize))
		i.Signatures = []releasediscovery.Signature{{Algorithm: "ed25519", KeyID: "forward-test", Value: base64.StdEncoding.EncodeToString(ed25519.Sign(key, message))}}
	}
	b, err := json.Marshal(i)
	if err != nil {
		t.Fatal(err)
	}
	return b
}
