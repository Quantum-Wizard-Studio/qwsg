package updateauthority

import (
	"crypto/ed25519"
	"encoding/base64"
	"encoding/json"
	"strings"
	"testing"
	"time"

	"quantumwizard.hu/qwsg/internal/installation"
	"quantumwizard.hu/qwsg/internal/releasediscovery"
	"quantumwizard.hu/qwsg/internal/update"
)

func capabilityFixture() releasediscovery.Index {
	i := indexFixture()
	i.Schema = releasediscovery.CapabilitySchema
	r := &i.Channels[0].Releases[0]
	r.MigrationRoutes = nil
	r.Compatibility = []update.CompatibilityDeclaration{{Schema: update.MigrationSchema, SourceVersion: "1.1.0", TargetVersion: r.Version, Platform: "linux-amd64", Capability: update.PreservePackageV1, ConfigurationSchema: "1.0", GuardianSchema: "1.0", SchedulerSchema: "1.0", OperatorState: "1.0-1.2"}}
	return i
}

func signedCapability(t *testing.T, i releasediscovery.Index) ([]byte, releasediscovery.Verifier) {
	t.Helper()
	key := ed25519.NewKeyFromSeed(make([]byte, ed25519.SeedSize))
	b, err := releasediscovery.SigningBytes(i)
	if err != nil {
		t.Fatal(err)
	}
	i.Signatures = []releasediscovery.Signature{{Algorithm: "ed25519", KeyID: "test", Value: base64.StdEncoding.EncodeToString(ed25519.Sign(key, b))}}
	payload, err := json.Marshal(i)
	if err != nil {
		t.Fatal(err)
	}
	v, err := releasediscovery.NewVerifier(map[string]ed25519.PublicKey{"test": key.Public().(ed25519.PublicKey)})
	if err != nil {
		t.Fatal(err)
	}
	return payload, v
}

func TestCapabilityAuthority(t *testing.T) {
	now, _ := time.Parse(time.RFC3339, "2026-08-30T01:00:00Z")
	evaluator, _ := releasediscovery.NewEvaluator(func(candidate string) installation.Result {
		if candidate != "" {
			t.Fatal("future authority consulted target-specific classification")
		}
		return installation.Result{State: installation.VerifiedSupported, Version: "1.1.0"}
	})
	for _, tc := range []struct {
		name   string
		mutate func(*releasediscovery.Index)
		refuse bool
	}{
		{"future target", func(i *releasediscovery.Index) {
			r := &i.Channels[0].Releases[0]
			r.Version = "1.987.654"
			r.Compatibility[0].TargetVersion = r.Version
			r.Artifacts[0].Name = "qwsg-" + r.Version + "-linux-amd64.tar.gz"
			r.Artifacts[0].URL = "https://example.invalid/" + r.Artifacts[0].Name
		}, false},
		{"unknown capability", func(i *releasediscovery.Index) {
			i.Channels[0].Releases[0].Compatibility[0].Capability = "future-code-v9"
		}, true},
		{"source mismatch", func(i *releasediscovery.Index) { i.Channels[0].Releases[0].Compatibility[0].SourceVersion = "1.1.1" }, true},
		{"configuration constraint", func(i *releasediscovery.Index) {
			i.Channels[0].Releases[0].Compatibility[0].ConfigurationSchema = "2.0"
		}, true},
		{"guardian constraint", func(i *releasediscovery.Index) { i.Channels[0].Releases[0].Compatibility[0].GuardianSchema = "2.0" }, true},
		{"scheduler constraint", func(i *releasediscovery.Index) { i.Channels[0].Releases[0].Compatibility[0].SchedulerSchema = "2.0" }, true},
		{"state constraint", func(i *releasediscovery.Index) { i.Channels[0].Releases[0].Compatibility[0].OperatorState = "9.0" }, true},
	} {
		t.Run(tc.name, func(t *testing.T) {
			i := capabilityFixture()
			tc.mutate(&i)
			p, v := signedCapability(t, i)
			c, err := Authorize(p, v, evaluator, "stable", "linux-amd64", "", now)
			if (err != nil) != tc.refuse {
				t.Fatalf("candidate=%+v err=%v", c, err)
			}
			if !tc.refuse && c.Evaluation().MigrationID != update.PreservePackageV1 {
				t.Fatal("wrong capability")
			}
		})
	}
	p, v := signedCapability(t, capabilityFixture())
	for _, tc := range []struct {
		name     string
		payload  []byte
		target   string
		platform string
	}{
		{"tampered signature", []byte(strings.Replace(string(p), "preserve-package-v1", "preserve-package-v2", 1)), "", "linux-amd64"},
		{"target mismatch", p, "1.9.9", "linux-amd64"},
		{"platform mismatch", p, "", "linux-arm64"},
	} {
		t.Run(tc.name, func(t *testing.T) {
			if _, err := Authorize(tc.payload, v, evaluator, "stable", tc.platform, tc.target, now); err == nil {
				t.Fatal("unsafe authority accepted")
			}
		})
	}
	i := capabilityFixture()
	p, _ = json.Marshal(i)
	if _, err := Authorize(p, v, evaluator, "stable", "linux-amd64", "", now); releasediscovery.FailureOf(err) != releasediscovery.UnauthenticatedMetadata {
		t.Fatalf("unsigned err=%v", err)
	}
	if _, err := evaluator.Evaluate(releasediscovery.AuthenticatedIndex{}, "stable", "linux-amd64", false); err == nil {
		t.Fatal("zero proof accepted")
	}
}

func TestCapabilityContractRejectsAmbiguousMalformedAndConflictingAuthority(t *testing.T) {
	for _, tc := range []struct {
		name   string
		mutate func(*releasediscovery.Index)
	}{
		{"duplicate", func(i *releasediscovery.Index) {
			r := &i.Channels[0].Releases[0]
			r.Compatibility = append(r.Compatibility, r.Compatibility[0])
		}},
		{"conflict", func(i *releasediscovery.Index) {
			r := &i.Channels[0].Releases[0]
			d := r.Compatibility[0]
			d.Capability = "other-v1"
			r.Compatibility = append(r.Compatibility, d)
		}},
		{"legacy conflict", func(i *releasediscovery.Index) {
			i.Channels[0].Releases[0].MigrationRoutes = []string{"compat-1.1.0-to-1.2.0"}
		}},
		{"wrong target", func(i *releasediscovery.Index) { i.Channels[0].Releases[0].Compatibility[0].TargetVersion = "1.9.0" }},
		{"downgrade", func(i *releasediscovery.Index) { i.Channels[0].Releases[0].Compatibility[0].SourceVersion = "1.9.0" }},
		{"unknown schema", func(i *releasediscovery.Index) {
			i.Channels[0].Releases[0].Compatibility[0].Schema = "qwsg.migration/999"
		}},
		{"platform", func(i *releasediscovery.Index) { i.Channels[0].Releases[0].Compatibility[0].Platform = "linux-arm64" }},
		{"missing", func(i *releasediscovery.Index) { i.Channels[0].Releases[0].Compatibility = nil }},
		{"v1 mixed authority", func(i *releasediscovery.Index) { i.Schema = releasediscovery.Schema }},
	} {
		t.Run(tc.name, func(t *testing.T) {
			i := capabilityFixture()
			tc.mutate(&i)
			p, _ := json.Marshal(i)
			if _, err := releasediscovery.Parse(p); err == nil {
				t.Fatal("invalid authority accepted")
			}
		})
	}
}

func TestCapabilitySigningCloneIsolation(t *testing.T) {
	i := capabilityFixture()
	a := authenticatedFixture(t, i)
	i.Channels[0].Releases[0].Compatibility[0].Capability = "mutated"
	copy := a.Index()
	copy.Channels[0].Releases[0].Compatibility[0].Capability = "mutated"
	if a.Index().Channels[0].Releases[0].Compatibility[0].Capability != update.PreservePackageV1 {
		t.Fatal("authenticated authority mutated")
	}
}

func indexFixture() releasediscovery.Index {
	return releasediscovery.Index{Schema: releasediscovery.Schema, Product: "qwsg", GeneratedAt: "2026-08-30T00:00:00Z", Channels: []releasediscovery.Channel{{Name: "stable", Releases: []releasediscovery.Release{{Version: "1.2.0", PublishedAt: "2026-08-29T00:00:00Z", Status: "active", SourceCommit: strings.Repeat("a", 40), ReleaseNotesURL: "https://example.invalid/notes", MinimumSourceVersion: "1.1.0", Artifacts: []releasediscovery.Artifact{{Platform: "linux-amd64", Name: "qwsg-1.2.0-linux-amd64.tar.gz", URL: "https://example.invalid/qwsg-1.2.0-linux-amd64.tar.gz", Size: 1234, SHA256: strings.Repeat("a", 64)}}}}}}}
}
func authenticatedFixture(t *testing.T, i releasediscovery.Index) releasediscovery.AuthenticatedIndex {
	t.Helper()
	payload, v := signedCapability(t, i)
	parsed, err := releasediscovery.Parse(payload)
	if err != nil {
		t.Fatal(err)
	}
	a, err := v.Verify(parsed)
	if err != nil {
		t.Fatal(err)
	}
	return a
}

func TestCandidateCannotGrantAuthorityFromAwarenessOrNoUpdate(t *testing.T) {
	now, _ := time.Parse(time.RFC3339, "2026-08-30T01:00:00Z")
	evaluator, _ := releasediscovery.NewEvaluator(func(string) installation.Result {
		return installation.Result{State: installation.VerifiedSupported, Version: "1.1.0"}
	})
	payload, verifier := signedCapability(t, capabilityFixture())
	candidate, err := Authorize(payload, verifier, evaluator, "stable", "linux-amd64", "", now)
	if err != nil {
		t.Fatal(err)
	}
	if err = candidate.CheckWatermark(now); releasediscovery.FailureOf(err) != releasediscovery.MetadataRollback {
		t.Fatalf("older than known authority accepted: %v", err)
	}
	if err = candidate.CheckWatermark(now.Add(-time.Hour)); err != nil {
		t.Fatal(err)
	}
	if _, err := Authorize(payload, verifier, evaluator, "stable", "linux-amd64", "", now.Add(-2*time.Hour)); releasediscovery.FailureOf(err) != releasediscovery.MetadataFreshness {
		t.Fatalf("future metadata accepted: %v", err)
	}
	if _, err := (Candidate{}).VerifyStaged(update.Staged{}); err == nil {
		t.Fatal("zero candidate authorized package")
	}
	equal, _ := releasediscovery.NewEvaluator(func(string) installation.Result {
		return installation.Result{State: installation.VerifiedSupported, Version: "1.2.0"}
	})
	candidate, err = Authorize(payload, verifier, equal, "stable", "linux-amd64", "", now)
	if err != ErrNoUpdate {
		t.Fatalf("equal relation %v", err)
	}
	if _, err = candidate.VerifyStaged(update.Staged{}); err == nil {
		t.Fatal("non-upgrade candidate authorized package")
	}
}
