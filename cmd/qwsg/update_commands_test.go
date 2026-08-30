package main

import (
	"bytes"
	"context"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"quantumwizard.hu/qwsg/internal/releasediscovery"
	updatecore "quantumwizard.hu/qwsg/internal/update"
)

type commandAwarenessChecker struct {
	calls  int
	result releasediscovery.CheckResult
	err    error
}

func (c *commandAwarenessChecker) Check(context.Context, releasediscovery.FetchRequest, string, bool) (releasediscovery.CheckResult, error) {
	c.calls++
	return c.result, c.err
}

func TestUpdateHelp(t *testing.T) {
	var out, errout bytes.Buffer
	if code := run([]string{"help", "update"}, &out, &errout); code != 0 || !strings.Contains(out.String(), "qwsg update rollback") {
		t.Fatalf("code=%d out=%q err=%q", code, out.String(), errout.String())
	}
}

func TestBootstrapUsesInstalledIdentity(t *testing.T) {
	root := t.TempDir()
	installedPackageFixture(t, root, "1.1.0", false)
	previousRoot := installedQWSGRoot
	installedQWSGRoot = root
	defer func() { installedQWSGRoot = previousRoot }()
	got, err := installedVersion()
	if err != nil || got != "1.1.0" {
		t.Fatalf("got %q %v", got, err)
	}
	if updatecore.Classify(got, "1.2.0-rc.1") != updatecore.Newer {
		t.Fatal("bootstrap target not newer than installed identity")
	}
}

func TestRC2InstalledIdentityAndConfigurationPreflight(t *testing.T) {
	root := t.TempDir()
	binary := installedPackageFixture(t, root, "1.2.0-rc.2", true)
	previousBinary, previousRoot := installedQWSGBinary, installedQWSGRoot
	installedQWSGBinary, installedQWSGRoot = binary, root
	defer func() { installedQWSGBinary, installedQWSGRoot = previousBinary, previousRoot }()
	got, err := installedVersion()
	if err != nil || got != "1.2.0-rc.2" || validateInstalledConfiguration() != nil {
		t.Fatalf("RC.2 preflight failed: %q %v", got, err)
	}
}

func TestBinaryVersionOutputAloneIsNotInstalledIdentity(t *testing.T) {
	root := t.TempDir()
	binary := filepath.Join(root, "usr/local/bin/qwsg")
	if err := os.MkdirAll(filepath.Dir(binary), 0700); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(binary, []byte("#!/bin/sh\nprintf 'QWSG 1.2.0\\ncommit: 1111111111111111111111111111111111111111\\nbuilt: 2026-08-29T00:00:00Z\\n'\n"), 0700); err != nil {
		t.Fatal(err)
	}
	previousRoot := installedQWSGRoot
	installedQWSGRoot = root
	defer func() { installedQWSGRoot = previousRoot }()
	if _, err := installedVersion(); err == nil {
		t.Fatal("binary-only version output established installed identity")
	}
}

func installedPackageFixture(t *testing.T, root, releaseVersion string, configValid bool) string {
	t.Helper()
	commit, built := "1111111111111111111111111111111111111111", "2026-08-29T00:00:00Z"
	binary := filepath.Join(root, "usr/local/bin/qwsg")
	for _, path := range []string{
		binary,
		filepath.Join(root, "usr/local/lib/systemd/user/qwsg-guardian.service"),
		filepath.Join(root, "usr/local/share/doc/qwsg/RELEASE.json"),
		filepath.Join(root, "usr/local/share/doc/qwsg/README.md"),
		filepath.Join(root, "usr/local/share/doc/qwsg/INSTALL.md"),
		filepath.Join(root, "usr/local/share/doc/qwsg/LICENSE"),
		filepath.Join(root, "usr/local/share/doc/qwsg/CHANGELOG.md"),
		filepath.Join(root, "usr/local/share/doc/qwsg/qwsg-config.json"),
	} {
		if err := os.MkdirAll(filepath.Dir(path), 0700); err != nil {
			t.Fatal(err)
		}
		body, mode := []byte("fixture"), os.FileMode(0600)
		switch filepath.Base(path) {
		case "qwsg":
			configurationExit := "exit 1"
			if configValid {
				configurationExit = "exit 0"
			}
			body = []byte("#!/bin/sh\nif test \"$1\" = version; then printf 'QWSG " + releaseVersion + "\\ncommit: " + commit + "\\nbuilt: " + built + "\\n'; exit 0; fi\nif test \"$1 $2\" = 'config validate'; then " + configurationExit + "; fi\nexit 1\n")
			mode = 0700
		case "RELEASE.json":
			body = []byte(`{"Schema":"qwsg.release/1","Version":"` + releaseVersion + `","Commit":"` + commit + `","Built":"` + built + `","Platform":"linux-amd64"}`)
		}
		if err := os.WriteFile(path, body, mode); err != nil {
			t.Fatal(err)
		}
	}
	return binary
}
func TestUpdateLocalArguments(t *testing.T) {
	archive, target, err := parseUpdateArgs([]string{"--archive", "/tmp/qwsg-1.2.0-rc.1-linux-amd64.tar.gz", "--version", "1.2.0-rc.1"})
	if err != nil || archive == "" || target != "1.2.0-rc.1" {
		t.Fatalf("%q %q %v", archive, target, err)
	}
	if _, _, err = parseUpdateArgs([]string{"--archive", "x"}); err == nil {
		t.Fatal("incomplete local identity accepted")
	}
}

func TestUpdateCheckPublishesAuthenticatedAwarenessAndStatusIsNetworkFree(t *testing.T) {
	root := t.TempDir()
	installedPackageFixture(t, root, "1.2.0", true)
	t.Setenv("QWSG_STATE_DIR", filepath.Join(t.TempDir(), "state"))
	checker := &commandAwarenessChecker{result: releasediscovery.CheckResult{
		Source:           releasediscovery.SourceEvidence{SourceID: "community-release-index", TransportAuthenticated: true, Validators: releasediscovery.Validators{ETag: "\"one\""}},
		IndexGeneratedAt: "2026-08-30T11:00:00Z",
		Authenticity:     releasediscovery.AuthenticityEvidence{Scheme: "ed25519", KeyID: "test-key"},
		Evaluation: releasediscovery.Evaluation{
			InstalledVersion: "1.2.0", Channel: "stable", Platform: "linux-amd64",
			Release:  releasediscovery.Release{Version: "1.3.0", PublishedAt: "2026-08-30T10:00:00Z", Status: "active"},
			Artifact: releasediscovery.Artifact{Name: "qwsg-1.3.0-linux-amd64.tar.gz", SHA256: strings.Repeat("a", 64), Size: 1234},
			Relation: updatecore.Newer, Compatibility: releasediscovery.CompatibilitySupported, MigrationID: "compat-1.2.0-to-1.3.0",
			Authenticity: releasediscovery.AuthenticityEvidence{Scheme: "ed25519", KeyID: "test-key"},
		},
	}}
	previousRoot, previousChecker := installedQWSGRoot, updateAwarenessChecker
	installedQWSGRoot, updateAwarenessChecker = root, checker
	defer func() { installedQWSGRoot, updateAwarenessChecker = previousRoot, previousChecker }()
	var out, errout bytes.Buffer
	if code := runUpdateCheck(&out, &errout); code != 0 || !strings.Contains(out.String(), "update_available") {
		t.Fatalf("code=%d out=%q err=%q", code, out.String(), errout.String())
	}
	checker.calls = 0
	out.Reset()
	errout.Reset()
	if code := runUpdateStatus(&out, &errout); code != 0 || checker.calls != 0 || !strings.Contains(out.String(), "update_available") {
		t.Fatalf("code=%d calls=%d out=%q err=%q", code, checker.calls, out.String(), errout.String())
	}
}
func TestPrivilegedBackupPathBoundary(t *testing.T) {
	for _, path := range []string{"/tmp/x", "/var/lib/qwsg/rollback/../x", "relative"} {
		if validBackup(path) {
			t.Fatalf("accepted %q", path)
		}
	}
	if !validBackup("/var/lib/qwsg/rollback/1000/transaction") {
		t.Fatal("valid backup refused")
	}
}
