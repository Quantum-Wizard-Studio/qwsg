package update_test

import (
	"crypto/sha256"
	"encoding/hex"
	"os"
	"os/exec"
	"path/filepath"
	"testing"

	"quantumwizard.hu/qwsg/internal/installation"
	"quantumwizard.hu/qwsg/internal/update"
)

// Run against immutable pipeline outputs, not synthetic package fixtures.
func TestRealRelease130CleanMigrationRollback(t *testing.T) {
	realReleaseAcceptance(t, "QWSG_ACCEPTANCE_130_ARCHIVE", "QWSG_ACCEPTANCE_120_ARCHIVE", "1.3.0", "1.2.0", "44768af20c8456cde09f940590b8c4446f605b2af02866e1553705a01d1a4c11")
}

func TestRealRelease131CleanMigrationRollback(t *testing.T) {
	realReleaseAcceptance(t, "QWSG_ACCEPTANCE_131_ARCHIVE", "QWSG_ACCEPTANCE_130_ARCHIVE", "1.3.1", "1.3.0", "8e19f624ccd1a49f32f6127889390e3389aec6eda958c7fbaaaa657d7b7cf9a0")
}

func realReleaseAcceptance(t *testing.T, newEnv, oldEnv, target, source, oldHash string) {
	t.Helper()
	newer, older := os.Getenv(newEnv), os.Getenv(oldEnv)
	if newer == "" || older == "" {
		t.Skip("set both release archive paths for real-package acceptance")
	}
	verify := func(archive, version string) update.Package {
		t.Helper()
		root := t.TempDir()
		pkg, err := update.VerifyPackage(update.Staged{Root: root, Archive: archive, Release: update.Release{Version: version}})
		if err != nil {
			t.Fatal(err)
		}
		return pkg
	}
	oldData, err := os.ReadFile(older)
	if err != nil {
		t.Fatal(err)
	}
	sum := sha256.Sum256(oldData)
	if hex.EncodeToString(sum[:]) != oldHash {
		t.Fatal("protected source archive changed")
	}
	oldPkg, newPkg := verify(older, source), verify(newer, target)
	install := func(pkg update.Package, root string) {
		t.Helper()
		cmd := exec.Command(filepath.Join(pkg.Root, "install.sh"), "--destdir", root)
		if output, err := cmd.CombinedOutput(); err != nil {
			t.Fatalf("staged install: %v %s", err, output)
		}
	}
	clean, migrated := t.TempDir(), t.TempDir()
	install(newPkg, clean)
	install(oldPkg, migrated)
	before := map[string][]byte{}
	err = filepath.WalkDir(migrated, func(p string, d os.DirEntry, e error) error {
		if e != nil {
			return e
		}
		if d.IsDir() {
			return nil
		}
		rel, e := filepath.Rel(migrated, p)
		if e != nil {
			return e
		}
		before[rel], e = os.ReadFile(p)
		return e
	})
	if err != nil {
		t.Fatal(err)
	}
	preserved := map[string][]byte{
		"home/operator/.config/qwsg/config.json":                         []byte(`{"fixture":"configuration-and-smtp-reference"}`),
		"home/operator/.config/qwsg/credentials/smtp-password":           []byte("non-secret-test-credential"),
		"home/operator/.local/state/qwsg/update-awareness/state.json":    []byte("awareness-and-deduplication-fixture"),
		"home/operator/.local/state/qwsg/operator-state.json":            []byte("inventory-fixture"),
		"home/operator/.local/state/qwsg/scheduler/scheduler-state.json": make([]byte, 9<<20),
	}
	for rel, body := range preserved {
		p := filepath.Join(migrated, rel)
		if err = os.MkdirAll(filepath.Dir(p), 0700); err != nil {
			t.Fatal(err)
		}
		if err = os.WriteFile(p, body, 0600); err != nil {
			t.Fatal(err)
		}
	}
	checkPreserved := func() {
		t.Helper()
		for rel, want := range preserved {
			got, err := os.ReadFile(filepath.Join(migrated, rel))
			if err != nil || string(got) != string(want) {
				t.Fatalf("private state changed: %s", rel)
			}
		}
	}
	classification := installation.Classify(installation.Options{Root: migrated, CandidateVersion: target})
	if classification.State != installation.SupportedUpgradeSource || classification.MigrationID != "compat-"+source+"-to-"+target {
		t.Fatalf("source=%+v", classification)
	}
	backup := filepath.Join(t.TempDir(), "transaction")
	tx, err := update.Apply(newPkg.Root, migrated, backup, source)
	if err != nil || !tx.Complete || tx.ToCommit != newPkg.Provenance.Commit {
		t.Fatalf("apply=%+v err=%v", tx, err)
	}
	checkPreserved()
	for _, root := range []string{clean, migrated} {
		identity := installation.Classify(installation.Options{Root: root})
		if identity.State != installation.VerifiedSupported || identity.Version != target {
			t.Fatalf("identity=%+v", identity)
		}
	}
	for _, rel := range []string{"usr/local/bin/qwsg", "usr/local/share/doc/qwsg/RELEASE.json", "usr/local/lib/systemd/user/qwsg-guardian.service"} {
		a, _ := os.ReadFile(filepath.Join(clean, rel))
		b, _ := os.ReadFile(filepath.Join(migrated, rel))
		if string(a) != string(b) {
			t.Fatalf("clean/migration mismatch: %s", rel)
		}
	}
	if err = update.Rollback(migrated, backup); err != nil {
		t.Fatal(err)
	}
	checkPreserved()
	for rel, want := range before {
		got, err := os.ReadFile(filepath.Join(migrated, rel))
		if err != nil || string(got) != string(want) {
			t.Fatalf("rollback mismatch: %s", rel)
		}
	}
	identity := installation.Classify(installation.Options{Root: migrated})
	if identity.State != installation.VerifiedSupported || identity.Version != source {
		t.Fatalf("rollback identity=%+v", identity)
	}
}
