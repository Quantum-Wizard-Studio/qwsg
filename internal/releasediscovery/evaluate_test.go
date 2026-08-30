package releasediscovery

import (
	"context"
	"crypto/ed25519"
	"encoding/base64"
	"encoding/json"
	"testing"

	"quantumwizard.hu/qwsg/internal/installation"
	"quantumwizard.hu/qwsg/internal/update"
)

type fixtureSource struct {
	result FetchResult
	err    error
}

func (s fixtureSource) Fetch(context.Context, FetchRequest) (FetchResult, error) {
	return s.result, s.err
}

func TestDiscovererEnforcesFetchParseAuthenticateEvaluateOrder(t *testing.T) {
	index, publicKey := signedFixture(t)
	payload, _ := json.Marshal(index)
	verifier, _ := NewVerifier(map[string]ed25519.PublicKey{"community-test-2026": publicKey})
	evaluator, _ := NewEvaluator(func(candidate string) installation.Result {
		if candidate == "" {
			return installation.Result{State: installation.VerifiedSupported, Version: "1.2.0-rc.2"}
		}
		return installation.Result{State: installation.SupportedUpgradeSource, Version: "1.2.0-rc.2", MigrationID: "compat-1.2.0-rc.2-to-1.2.0"}
	})
	discoverer, err := NewDiscoverer(fixtureSource{result: FetchResult{Manifest: payload, Evidence: SourceEvidence{SourceID: "fixture", TransportAuthenticated: true}}}, verifier, evaluator)
	if err != nil {
		t.Fatal(err)
	}
	result, err := discoverer.Check(context.Background(), FetchRequest{Channel: "stable"}, "linux-amd64", false)
	if err != nil || result.Evaluation.Compatibility != CompatibilitySupported || result.Source.SourceID != "fixture" {
		t.Fatalf("result=%+v err=%v", result, err)
	}

	unsigned, _ := json.Marshal(indexFixture())
	discoverer.source = fixtureSource{result: FetchResult{Manifest: unsigned}}
	if _, err = discoverer.Check(context.Background(), FetchRequest{Channel: "stable"}, "linux-amd64", false); FailureOf(err) != UnauthenticatedMetadata {
		t.Fatalf("unsigned failure=%s", FailureOf(err))
	}
	discoverer.source = fixtureSource{result: FetchResult{Manifest: []byte(`{"schema":`)}}
	if _, err = discoverer.Check(context.Background(), FetchRequest{Channel: "stable"}, "linux-amd64", false); FailureOf(err) != MalformedMetadata {
		t.Fatalf("malformed failure=%s", FailureOf(err))
	}
	discoverer.source = fixtureSource{result: FetchResult{NotModified: true, Evidence: SourceEvidence{SourceID: "fixture", TransportAuthenticated: true}}}
	if result, err = discoverer.Check(context.Background(), FetchRequest{Channel: "stable"}, "linux-amd64", false); err != nil || !result.NotModified {
		t.Fatalf("not-modified result=%+v err=%v", result, err)
	}
}

func TestEvaluationConsumesVerifiedInstalledClassificationAndMigration(t *testing.T) {
	authenticated := authenticatedFixture(t, indexFixture())
	calls := []string{}
	evaluator, err := NewEvaluator(func(candidate string) installation.Result {
		calls = append(calls, candidate)
		if candidate == "" {
			return installation.Result{State: installation.VerifiedSupported, Reason: installation.ReasonPackageVerified, Version: "1.2.0-rc.2"}
		}
		return installation.Result{State: installation.SupportedUpgradeSource, Reason: installation.ReasonUpgradeRouteVerified, Version: "1.2.0-rc.2", MigrationID: "compat-1.2.0-rc.2-to-1.2.0"}
	})
	if err != nil {
		t.Fatal(err)
	}
	result, err := evaluator.Evaluate(authenticated, "stable", "linux-amd64", false)
	if err != nil || result.Relation != update.Newer || result.Compatibility != CompatibilitySupported || result.MigrationID != "compat-1.2.0-rc.2-to-1.2.0" || len(calls) != 2 || calls[0] != "" || calls[1] != "1.2.0" {
		t.Fatalf("result=%+v calls=%v err=%v", result, calls, err)
	}
}

func TestEvaluationRejectsUnsafeInstalledStates(t *testing.T) {
	authenticated := authenticatedFixture(t, indexFixture())
	for _, state := range []installation.State{installation.NoInstallation, installation.LegacyInstallation, installation.UnknownInstallation, installation.InconsistentInstallation, installation.SupportedUpgradeSource} {
		t.Run(string(state), func(t *testing.T) {
			evaluator, _ := NewEvaluator(func(string) installation.Result { return installation.Result{State: state, Version: "1.2.0-rc.2"} })
			if _, err := evaluator.Evaluate(authenticated, "stable", "linux-amd64", false); FailureOf(err) != InstalledUnverified {
				t.Fatalf("state=%s failure=%s err=%v", state, FailureOf(err), err)
			}
		})
	}
}

func TestManifestAdvisoryCannotCreateOrOverrideLocalCompatibility(t *testing.T) {
	tests := []struct {
		name      string
		mutate    func(*Index)
		upgrade   installation.Result
		installed string
	}{
		{"missing advertised route", func(index *Index) { index.Channels[0].Releases[0].MigrationRoutes = nil }, installation.Result{State: installation.SupportedUpgradeSource, Version: "1.2.0-rc.2", MigrationID: "compat-1.2.0-rc.2-to-1.2.0"}, "1.2.0-rc.2"},
		{"local route missing", func(*Index) {}, installation.Result{State: installation.UnknownInstallation, Version: "1.2.0-rc.2"}, "1.2.0-rc.2"},
		{"minimum source unmet", func(index *Index) { index.Channels[0].Releases[0].MinimumSourceVersion = "1.2.0-rc.2" }, installation.Result{State: installation.SupportedUpgradeSource, Version: "1.1.0", MigrationID: "compat-1.2.0-rc.2-to-1.2.0"}, "1.1.0"},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			index := indexFixture()
			tc.mutate(&index)
			authenticated := authenticatedFixture(t, index)
			evaluator, _ := NewEvaluator(func(candidate string) installation.Result {
				if candidate == "" {
					return installation.Result{State: installation.VerifiedSupported, Version: tc.installed}
				}
				return tc.upgrade
			})
			result, err := evaluator.Evaluate(authenticated, "stable", "linux-amd64", false)
			if err != nil || result.Compatibility != CompatibilityUnsupported || result.MigrationID != "" {
				t.Fatalf("result=%+v err=%v", result, err)
			}
		})
	}
}

func TestWithdrawalAndPrereleaseEligibility(t *testing.T) {
	withdrawn := indexFixture()
	withdrawn.Channels[0].Releases[0].Status = "withdrawn"
	evaluator, _ := NewEvaluator(func(string) installation.Result {
		return installation.Result{State: installation.VerifiedSupported, Version: "1.2.0-rc.2"}
	})
	if _, err := evaluator.Evaluate(authenticatedFixture(t, withdrawn), "stable", "linux-amd64", false); FailureOf(err) != NoEligibleRelease {
		t.Fatalf("withdrawn failure=%s", FailureOf(err))
	}

	preview := indexFixture()
	release := &preview.Channels[0].Releases[0]
	preview.Channels[0].Name = "preview"
	release.Version = "1.2.0-rc.1"
	release.Artifacts[0].Name = "qwsg-1.2.0-rc.1-linux-amd64.tar.gz"
	release.Artifacts[0].URL = "https://git.quantumwizard.hu/Quantum_Wizard_Studio/qwsg/releases/download/v1.2.0-rc.1/qwsg-1.2.0-rc.1-linux-amd64.tar.gz"
	authenticatedPreview := authenticatedFixture(t, preview)
	if _, err := evaluator.Evaluate(authenticatedPreview, "preview", "linux-amd64", false); FailureOf(err) != NoEligibleRelease {
		t.Fatalf("prerelease opt-out failure=%s", FailureOf(err))
	}
	if _, err := evaluator.Evaluate(authenticatedPreview, "preview", "linux-amd64", true); err != nil {
		t.Fatalf("prerelease opt-in failed: %v", err)
	}
}

func authenticatedFixture(t *testing.T, index Index) AuthenticatedIndex {
	t.Helper()
	seed := [ed25519.SeedSize]byte{1, 2, 3, 4, 5}
	privateKey := ed25519.NewKeyFromSeed(seed[:])
	message, err := SigningBytes(index)
	if err != nil {
		t.Fatal(err)
	}
	index.Signatures = []Signature{{Algorithm: "ed25519", KeyID: "community-test-2026", Value: base64.StdEncoding.EncodeToString(ed25519.Sign(privateKey, message))}}
	verifier, _ := NewVerifier(map[string]ed25519.PublicKey{"community-test-2026": privateKey.Public().(ed25519.PublicKey)})
	authenticated, err := verifier.Verify(index)
	if err != nil {
		t.Fatal(err)
	}
	return authenticated
}
