package updatenotification

import (
	"context"
	"crypto/ed25519"
	"encoding/base64"
	"encoding/json"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"

	"quantumwizard.hu/qwsg/internal/installation"
	"quantumwizard.hu/qwsg/internal/releasediscovery"
	"quantumwizard.hu/qwsg/internal/update"
	"quantumwizard.hu/qwsg/internal/updateawareness"
)

type compatibilitySource struct {
	payload []byte
	calls   int
}

func (s *compatibilitySource) Fetch(context.Context, releasediscovery.FetchRequest) (releasediscovery.FetchResult, error) {
	s.calls++
	return releasediscovery.FetchResult{Manifest: s.payload, Evidence: releasediscovery.SourceEvidence{SourceID: "community-release-index", TransportAuthenticated: true}}, nil
}

// Exercise the actual package classifier, compiled migration table, signature
// verifier and notification eligibility together. The source supplies metadata
// only; no production signing material or installation operation is involved.
func TestCompatibilityBackportAuthenticatedLifecycle(t *testing.T) {
	for _, tc := range []struct {
		name, version, route         string
		unsigned, wrongKey, mismatch bool
		want                         updateawareness.Status
		eligible                     bool
	}{
		{name: "supported newer", version: "1.3.0", route: "compat-1.2.0-to-1.3.0", want: updateawareness.UpdateAvailable, eligible: true},
		{name: "current", version: "1.2.0", route: "compat-1.2.0-to-1.3.0", want: updateawareness.Current},
		{name: "unknown local route", version: "1.4.0", route: "compat-1.2.0-to-1.4.0", want: updateawareness.UpdateUnsupported},
		{name: "wrong advertised route", version: "1.3.0", route: "compat-1.1.0-to-1.3.0", want: updateawareness.UpdateUnsupported},
		{name: "unsigned", version: "1.3.0", route: "compat-1.2.0-to-1.3.0", unsigned: true},
		{name: "untrusted key", version: "1.3.0", route: "compat-1.2.0-to-1.3.0", wrongKey: true},
		{name: "provenance mismatch", version: "1.3.0", route: "compat-1.2.0-to-1.3.0", mismatch: true},
	} {
		t.Run(tc.name, func(t *testing.T) {
			root := t.TempDir()
			commit := strings.Repeat("1", 40)
			built := "2026-09-03T00:00:00Z"
			files := map[string]string{
				"usr/local/bin/qwsg": "#!/bin/sh\nprintf 'QWSG 1.2.0\\ncommit: " + commit + "\\nbuilt: " + built + "\\n'\n",
				"usr/local/lib/systemd/user/qwsg-guardian.service": "unit",
			}
			for _, name := range []string{"README.md", "INSTALL.md", "LICENSE", "CHANGELOG.md", "qwsg-config.json"} {
				files["usr/local/share/doc/qwsg/"+name] = "fixture"
			}
			metadataCommit := commit
			if tc.mismatch {
				metadataCommit = strings.Repeat("2", 40)
			}
			files["usr/local/share/doc/qwsg/RELEASE.json"] = `{"Schema":"qwsg.release/1","Version":"1.2.0","Commit":"` + metadataCommit + `","Built":"` + built + `","Platform":"linux-amd64"}`
			for name, body := range files {
				p := filepath.Join(root, name)
				if err := os.MkdirAll(filepath.Dir(p), 0700); err != nil {
					t.Fatal(err)
				}
				if err := os.WriteFile(p, []byte(body), 0700); err != nil {
					t.Fatal(err)
				}
			}
			index := releasediscovery.Index{Schema: releasediscovery.Schema, Product: releasediscovery.Product, GeneratedAt: "2026-09-03T11:00:00Z", Channels: []releasediscovery.Channel{{Name: "stable", Releases: []releasediscovery.Release{{Version: tc.version, PublishedAt: "2026-09-03T10:00:00Z", Status: "active", SourceCommit: commit, ReleaseNotesURL: "https://git.quantumwizard.hu/Quantum_Wizard_Studio/qwsg/releases/tag/v" + tc.version, MinimumSourceVersion: "1.2.0", MigrationRoutes: []string{tc.route}, Artifacts: []releasediscovery.Artifact{{Platform: "linux-amd64", Name: "qwsg-" + tc.version + "-linux-amd64.tar.gz", URL: "https://git.quantumwizard.hu/Quantum_Wizard_Studio/qwsg/releases/download/v" + tc.version + "/qwsg-" + tc.version + "-linux-amd64.tar.gz", Size: 1234, SHA256: strings.Repeat("a", 64)}}}}}}}
			key := ed25519.NewKeyFromSeed(make([]byte, ed25519.SeedSize))
			signing, err := releasediscovery.SigningBytes(index)
			if err != nil {
				t.Fatal(err)
			}
			if !tc.unsigned {
				index.Signatures = []releasediscovery.Signature{{Algorithm: "ed25519", KeyID: "fixture", Value: base64.StdEncoding.EncodeToString(ed25519.Sign(key, signing))}}
			}
			public := key.Public().(ed25519.PublicKey)
			if tc.wrongKey {
				public = ed25519.NewKeyFromSeed([]byte(strings.Repeat("x", 32))).Public().(ed25519.PublicKey)
			}
			verifier, err := releasediscovery.NewVerifier(map[string]ed25519.PublicKey{"fixture": public})
			if err != nil {
				t.Fatal(err)
			}
			evaluator, err := releasediscovery.NewEvaluator(func(candidate string) installation.Result {
				return installation.Classify(installation.Options{Root: root, CandidateVersion: candidate})
			})
			if err != nil {
				t.Fatal(err)
			}
			payload, err := json.Marshal(index)
			if err != nil {
				t.Fatal(err)
			}
			source := &compatibilitySource{payload: payload}
			discoverer, err := releasediscovery.NewDiscoverer(source, verifier, evaluator)
			if err != nil {
				t.Fatal(err)
			}
			result, err := discoverer.Check(context.Background(), releasediscovery.FetchRequest{Channel: "stable"}, "linux-amd64", false)
			if source.calls != 1 {
				t.Fatal("unexpected acquisition")
			}
			if tc.unsigned || tc.wrongKey || tc.mismatch {
				if err == nil {
					t.Fatal("unsafe input accepted")
				}
			} else {
				if err != nil {
					t.Fatal(err)
				}
				state, err := updateawareness.NewSuccess(nil, result, "community-release-index", "stable", updateawareness.Installed{Classification: installation.VerifiedSupported, Version: "1.2.0"}, notificationTime, 48*time.Hour)
				if err != nil {
					t.Fatal(err)
				}
				_, eligible := Eligible(state)
				if state.Status != tc.want || eligible != tc.eligible {
					t.Fatalf("status=%s eligible=%t", state.Status, eligible)
				}
				if tc.eligible && (result.Evaluation.Relation != update.Newer || result.Evaluation.MigrationID != "compat-1.2.0-to-1.3.0") {
					t.Fatal("wrong migration")
				}
			}
			// Checking may neither replace installed artifacts nor create download/state files.
			count := 0
			err = filepath.WalkDir(root, func(p string, d os.DirEntry, e error) error {
				if e != nil {
					return e
				}
				if d.IsDir() {
					return nil
				}
				count++
				relative, e := filepath.Rel(root, p)
				if e != nil {
					return e
				}
				body, e := os.ReadFile(p)
				if e != nil {
					return e
				}
				if string(body) != files[relative] {
					t.Errorf("package changed: %s", relative)
				}
				return nil
			})
			if err != nil || count != len(files) {
				t.Fatalf("unexpected filesystem mutation: count=%d err=%v", count, err)
			}
		})
	}
}
