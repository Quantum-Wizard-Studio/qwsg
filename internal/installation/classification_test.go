package installation

import (
	"context"
	"os"
	"path/filepath"
	"testing"
)

const (
	fixtureCommit = "1111111111111111111111111111111111111111"
	fixtureBuilt  = "2026-08-29T00:00:00Z"
)

func TestClassificationMatrix(t *testing.T) {
	tests := []struct {
		name      string
		prepare   func(*testing.T, string)
		candidate string
		want      State
		why       Reason
	}{
		{"clean host", nil, "", NoInstallation, ReasonNoPackageArtifacts},
		{"verified current", func(t *testing.T, root string) { packageFixture(t, root, "1.2.0") }, "", VerifiedSupported, ReasonPackageVerified},
		{"supported older", func(t *testing.T, root string) { packageFixture(t, root, "1.1.0") }, "", VerifiedSupported, ReasonPackageVerified},
		{"supported upgrade source", func(t *testing.T, root string) { packageFixture(t, root, "1.1.0") }, "1.2.0-rc.1", SupportedUpgradeSource, ReasonUpgradeRouteVerified},
		{"unsupported upgrade source", func(t *testing.T, root string) { packageFixture(t, root, "1.0.0") }, "1.2.0", UnknownInstallation, ReasonUnsupportedVersion},
		{"legacy prealpha binary", func(t *testing.T, root string) { binaryFixture(t, root, "0.0.1-prealpha", fixtureCommit, fixtureBuilt) }, "", LegacyInstallation, ReasonLegacyBinaryOnly},
		{"arbitrary executable", func(t *testing.T, root string) {
			write(t, filepath.Join(root, packagePaths[0]), "#!/bin/sh\nexit 1\n", 0700)
		}, "", UnknownInstallation, ReasonBinaryOnlyUnverified},
		{"version output without package", func(t *testing.T, root string) { binaryFixture(t, root, "1.2.0", fixtureCommit, fixtureBuilt) }, "", UnknownInstallation, ReasonBinaryOnlyUnverified},
		{"partial package", func(t *testing.T, root string) {
			binaryFixture(t, root, "1.2.0", fixtureCommit, fixtureBuilt)
			write(t, filepath.Join(root, packagePaths[1]), "unit", 0600)
		}, "", InconsistentInstallation, ReasonPackageLayoutIncomplete},
		{"provenance mismatch", func(t *testing.T, root string) {
			packageFixture(t, root, "1.2.0")
			binaryFixture(t, root, "1.2.0", "2222222222222222222222222222222222222222", fixtureBuilt)
		}, "", InconsistentInstallation, ReasonProvenanceMismatch},
		{"unsupported package version", func(t *testing.T, root string) { packageFixture(t, root, "2.0.0") }, "", UnknownInstallation, ReasonUnsupportedVersion},
		{"malformed metadata", func(t *testing.T, root string) {
			packageFixture(t, root, "1.2.0")
			write(t, filepath.Join(root, packagePaths[2]), "{", 0600)
		}, "", InconsistentInstallation, ReasonReleaseMetadataInvalid},
		{"unsafe artifact", func(t *testing.T, root string) {
			if err := os.MkdirAll(filepath.Dir(filepath.Join(root, packagePaths[0])), 0700); err != nil {
				t.Fatal(err)
			}
			if err := os.Symlink("elsewhere", filepath.Join(root, packagePaths[0])); err != nil {
				t.Fatal(err)
			}
		}, "", InconsistentInstallation, ReasonArtifactTypeUnsafe},
		{"unsafe ancestor symlink", func(t *testing.T, root string) {
			outside := t.TempDir()
			if err := os.MkdirAll(filepath.Join(root, "usr/local"), 0700); err != nil {
				t.Fatal(err)
			}
			if err := os.Symlink(outside, filepath.Join(root, "usr/local/bin")); err != nil {
				t.Fatal(err)
			}
			write(t, filepath.Join(outside, "qwsg"), "#!/bin/sh\nexit 1\n", 0700)
		}, "", InconsistentInstallation, ReasonArtifactTypeUnsafe},
		{"valid package with unrelated stale state", func(t *testing.T, root string) {
			packageFixture(t, root, "1.2.0")
			write(t, filepath.Join(root, "home/user/.config/qwsg/config.json"), "stale", 0600)
		}, "", VerifiedSupported, ReasonPackageVerified},
		{"stale state without package", func(t *testing.T, root string) {
			write(t, filepath.Join(root, "home/user/.local/state/qwsg/checkpoint.json"), "stale", 0600)
		}, "", NoInstallation, ReasonNoPackageArtifacts},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			root := t.TempDir()
			if tc.prepare != nil {
				tc.prepare(t, root)
			}
			got := Classify(Options{Root: root, CandidateVersion: tc.candidate})
			if got.State != tc.want || got.Reason != tc.why {
				t.Fatalf("got %+v want state=%s reason=%s", got, tc.want, tc.why)
			}
			if got.MigrationID != "" && got.State != SupportedUpgradeSource {
				t.Fatalf("migration leaked into non-upgrade result: %+v", got)
			}
		})
	}
}

func TestSpoofedVersionOutputCannotOverridePackageProvenance(t *testing.T) {
	root := t.TempDir()
	packageFixture(t, root, "1.2.0")
	got := Classify(Options{Root: root, RunVersion: func(context.Context, string) ([]byte, error) {
		return []byte("QWSG 1.2.0\ncommit: 9999999999999999999999999999999999999999\nbuilt: " + fixtureBuilt + "\n"), nil
	}})
	if got.State != InconsistentInstallation || got.Reason != ReasonProvenanceMismatch {
		t.Fatalf("spoof accepted: %+v", got)
	}
}

func packageFixture(t *testing.T, root, version string) {
	t.Helper()
	binaryFixture(t, root, version, fixtureCommit, fixtureBuilt)
	for _, relative := range packagePaths[1:] {
		body := "fixture"
		if relative == packagePaths[2] {
			body = `{"Schema":"qwsg.release/1","Version":"` + version + `","Commit":"` + fixtureCommit + `","Built":"` + fixtureBuilt + `","Platform":"linux-amd64"}`
		}
		write(t, filepath.Join(root, relative), body, 0600)
	}
}

func binaryFixture(t *testing.T, root, version, commit, built string) {
	t.Helper()
	body := "#!/bin/sh\nprintf 'QWSG " + version + "\\ncommit: " + commit + "\\nbuilt: " + built + "\\n'\n"
	write(t, filepath.Join(root, packagePaths[0]), body, 0700)
}

func write(t *testing.T, path, body string, mode os.FileMode) {
	t.Helper()
	if err := os.MkdirAll(filepath.Dir(path), 0700); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(path, []byte(body), mode); err != nil {
		t.Fatal(err)
	}
}
