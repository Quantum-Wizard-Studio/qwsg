package updateauthority

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"os"
	"os/exec"
	"path/filepath"
	"testing"
	"time"

	"quantumwizard.hu/qwsg/internal/installation"
	"quantumwizard.hu/qwsg/internal/releasediscovery"
	"quantumwizard.hu/qwsg/internal/releasepublication"
	"quantumwizard.hu/qwsg/internal/update"
)

// This bounded package test never changes a host installation or grants
// production signing authority. The actual frozen unsigned candidate is copied
// into a test-only signature domain; the production verifier must reject it.
func TestReleaseCandidateBootstrapPackage(t *testing.T) {
	candidatePath, archive, historical := os.Getenv("QWSG_CANDIDATE_INDEX"), os.Getenv("QWSG_CANDIDATE_ARCHIVE"), os.Getenv("QWSG_ACCEPTANCE_131_ARCHIVE")
	if candidatePath == "" || archive == "" || historical == "" {
		t.Skip("requires frozen candidate index/archive and official 1.3.1 archive")
	}
	raw, err := os.ReadFile(candidatePath)
	if err != nil {
		t.Fatal(err)
	}
	canonical, err := releasepublication.Generate(raw)
	if err != nil {
		t.Fatal(err)
	}
	index, err := releasediscovery.Parse(canonical)
	if err != nil {
		t.Fatal(err)
	}
	r := index.Channels[0].Releases[0]
	if r.Version != "1.4.0" || len(r.Compatibility) != 1 || r.Compatibility[0].SourceVersion != "1.3.1" {
		t.Fatal("unexpected candidate boundary")
	}
	if _, err := update.PlanMigration("1.3.1", r.Version); err == nil {
		t.Fatal("bootstrap must not invent a historical route")
	}
	data, err := os.ReadFile(historical)
	if err != nil {
		t.Fatal(err)
	}
	digest := sha256.Sum256(data)
	if hex.EncodeToString(digest[:]) != "0ab726bcde36182232ff89e3ea33e2d5ab77d8cad2b94ce56c9534a8147d9cd7" {
		t.Fatal("historical artifact changed")
	}
	old, err := update.VerifyPackage(update.Staged{Root: t.TempDir(), Archive: historical, Release: update.Release{Version: "1.3.1"}})
	if err != nil {
		t.Fatal(err)
	}
	if old.Provenance.Commit != "45009fe6169bff00842a4c4e9561bf339a5db81e" {
		t.Fatal("historical source changed")
	}
	root := t.TempDir()
	if out, err := exec.Command(filepath.Join(old.Root, "install.sh"), "--destdir", root).CombinedOutput(); err != nil {
		t.Fatalf("isolated install: %v %s", err, out)
	}
	evaluator, err := releasediscovery.NewEvaluator(func(target string) installation.Result {
		return installation.Classify(installation.Options{Root: root, CandidateVersion: target})
	})
	if err != nil {
		t.Fatal(err)
	}
	before := map[string][]byte{}
	if err := filepath.WalkDir(root, func(p string, d os.DirEntry, e error) error {
		if e != nil {
			return e
		}
		if d.IsDir() {
			return nil
		}
		b, e := os.ReadFile(p)
		before[p] = b
		return e
	}); err != nil {
		t.Fatal(err)
	}
	private := filepath.Join(root, "home/operator/private-state")
	if err := os.MkdirAll(filepath.Dir(private), 0700); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(private, []byte("preserved-private-fixture"), 0600); err != nil {
		t.Fatal(err)
	}
	before[private] = []byte("preserved-private-fixture")
	signed, verifier := signedCapability(t, index)
	if err := releasepublication.VerifyProduction(signed); err == nil {
		t.Fatal("test key became production authority")
	}
	if _, err := Authorize(raw, verifier, evaluator, "stable", "linux-amd64", r.Version, time.Now()); err == nil {
		t.Fatal("unsigned candidate authorized")
	}
	for _, scenario := range []string{"unknown-capability", "unsupported-schema", "wrong-source", "altered-signature"} {
		t.Run(scenario, func(t *testing.T) {
			changed, err := releasediscovery.Parse(canonical)
			if err != nil {
				t.Fatal(err)
			}
			d := &changed.Channels[0].Releases[0].Compatibility[0]
			switch scenario {
			case "unknown-capability":
				d.Capability = "unknown-v9"
			case "unsupported-schema":
				d.ConfigurationSchema = "9.0"
			case "wrong-source":
				d.SourceVersion = "1.3.2"
			}
			payload, v := signedCapability(t, changed)
			if scenario == "altered-signature" {
				payload = bytes.Replace(payload, []byte("preserve-package-v1"), []byte("preserve-package-v2"), 1)
			}
			if _, err := Authorize(payload, v, evaluator, "stable", "linux-amd64", r.Version, time.Now()); err == nil {
				t.Fatal("unsupported authority accepted")
			}
		})
	}
	authority, err := Authorize(signed, verifier, evaluator, "stable", "linux-amd64", r.Version, time.Now())
	if err != nil {
		t.Fatal(err)
	}
	stageParent := t.TempDir()
	if err := os.Chmod(stageParent, 0700); err != nil {
		t.Fatal(err)
	}
	staged, err := update.StageLocal(archive, archive+".sha256", r.Version, stageParent)
	if err != nil {
		t.Fatal(err)
	}
	pkg, err := authority.VerifyStaged(staged)
	if err != nil {
		t.Fatal(err)
	}
	out, err := exec.Command(filepath.Join(pkg.Root, "bin/qwsg"), "version").CombinedOutput()
	if err != nil || !bytes.Contains(out, []byte(r.Version)) || !bytes.Contains(out, []byte(r.SourceCommit)) {
		t.Fatalf("embedded provenance: %v %s", err, out)
	}
	backup := filepath.Join(t.TempDir(), "transaction")
	tx, err := update.Apply(pkg.Root, root, backup, "1.3.1")
	if err != nil || !tx.Complete {
		t.Fatalf("isolated transaction: %+v %v", tx, err)
	}
	identity := installation.Classify(installation.Options{Root: root})
	if identity.State != installation.VerifiedSupported || identity.Version != r.Version {
		t.Fatalf("identity: %+v", identity)
	}
	if data, err := os.ReadFile(private); err != nil || !bytes.Equal(data, before[private]) {
		t.Fatal("private state changed")
	}
	if err := update.Rollback(root, backup); err != nil {
		t.Fatal(err)
	}
	for p, want := range before {
		got, err := os.ReadFile(p)
		if err != nil || !bytes.Equal(got, want) {
			t.Fatalf("rollback mismatch %s", p)
		}
	}
	identity = installation.Classify(installation.Options{Root: root})
	if identity.State != installation.VerifiedSupported || identity.Version != "1.3.1" {
		t.Fatalf("rollback identity: %+v", identity)
	}
	t.Log("actual candidate metadata/package verified; official 1.3.1 isolated package bootstrap/rollback PASS; no real systemd/sudo or production signature acceptance")
}
