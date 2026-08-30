package releasediscovery

import (
	"crypto/ed25519"
	"encoding/base64"
	"encoding/json"
	"strings"
	"testing"
)

func TestParseAndVerifySignedStableIndex(t *testing.T) {
	index, publicKey := signedFixture(t)
	payload, err := json.Marshal(index)
	if err != nil {
		t.Fatal(err)
	}
	parsed, err := Parse(payload)
	if err != nil {
		t.Fatal(err)
	}
	verifier, err := NewVerifier(map[string]ed25519.PublicKey{"community-test-2026": publicKey})
	if err != nil {
		t.Fatal(err)
	}
	authenticated, err := verifier.Verify(parsed)
	if err != nil || authenticated.Authenticity().Scheme != "ed25519" || authenticated.Authenticity().KeyID != "community-test-2026" {
		t.Fatalf("authenticated=%+v err=%v", authenticated.Authenticity(), err)
	}
}

func TestSigningBytesAreDeterministicAndExcludeSignatures(t *testing.T) {
	index := indexFixture()
	first, err := SigningBytes(index)
	if err != nil {
		t.Fatal(err)
	}
	want := `{"schema":"qwsg.release-index/1","product":"qwsg","generated_at":"2026-08-30T00:00:00Z","channels":[{"name":"stable","releases":[{"version":"1.2.0","published_at":"2026-08-29T16:53:34Z","status":"active","source_commit":"348d927ffcf4c8cd4c9a50fc3eacad71d8bfe5c2","release_notes_url":"https://git.quantumwizard.hu/Quantum_Wizard_Studio/qwsg/releases/tag/v1.2.0","minimum_source_version":"1.1.0","migration_routes":["compat-1.2.0-rc.2-to-1.2.0"],"artifacts":[{"platform":"linux-amd64","name":"qwsg-1.2.0-linux-amd64.tar.gz","url":"https://git.quantumwizard.hu/Quantum_Wizard_Studio/qwsg/releases/download/v1.2.0/qwsg-1.2.0-linux-amd64.tar.gz","size":3524214,"sha256":"44768af20c8456cde09f940590b8c4446f605b2af02866e1553705a01d1a4c11"}]}]}]}`
	if string(first) != want {
		t.Fatalf("canonical bytes changed:\n%s", first)
	}
	index.Signatures = []Signature{{Algorithm: "ed25519", KeyID: "unknown-key", Value: base64.StdEncoding.EncodeToString(make([]byte, ed25519.SignatureSize))}}
	second, err := SigningBytes(index)
	if err != nil || string(second) != want {
		t.Fatalf("signature affected signed bytes: %v %s", err, second)
	}
}

func TestAuthenticityFailsClosed(t *testing.T) {
	index, publicKey := signedFixture(t)
	otherSeed := [ed25519.SeedSize]byte{9, 8, 7, 6}
	otherPrivate := ed25519.NewKeyFromSeed(otherSeed[:])
	otherPublic := otherPrivate.Public().(ed25519.PublicKey)
	tests := []struct {
		name    string
		index   Index
		anchors map[string]ed25519.PublicKey
	}{
		{"unsigned", indexFixture(), map[string]ed25519.PublicKey{"community-test-2026": publicKey}},
		{"unknown key", index, map[string]ed25519.PublicKey{"different-approved-key": otherPublic}},
		{"wrong approved key", index, map[string]ed25519.PublicKey{"community-test-2026": otherPublic}},
	}
	altered := index
	altered.Channels[0].Releases[0].Artifacts[0].Size++
	tests = append(tests, struct {
		name    string
		index   Index
		anchors map[string]ed25519.PublicKey
	}{"altered payload", altered, map[string]ed25519.PublicKey{"community-test-2026": publicKey}})
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			verifier, err := NewVerifier(tc.anchors)
			if err != nil {
				t.Fatal(err)
			}
			if _, err = verifier.Verify(tc.index); FailureOf(err) != UnauthenticatedMetadata {
				t.Fatalf("failure=%s err=%v", FailureOf(err), err)
			}
		})
	}
}

func TestStrictJSONAndSemanticBoundaries(t *testing.T) {
	valid, _ := json.Marshal(indexFixture())
	tests := []struct {
		name    string
		payload []byte
		failure Failure
	}{
		{"duplicate member", []byte(`{"schema":"qwsg.release-index/1","schema":"qwsg.release-index/1"}`), MalformedMetadata},
		{"unknown member", append(valid[:len(valid)-1], []byte(`,"unexpected":true}`)...), MalformedMetadata},
		{"trailing data", append(valid, []byte(` {}`)...), MalformedMetadata},
		{"unsupported schema", mutateJSON(t, indexFixture(), func(index *Index) { index.Schema = "qwsg.release-index/2" }), UnsupportedContract},
		{"unsupported product", mutateJSON(t, indexFixture(), func(index *Index) { index.Product = "other" }), UnsupportedContract},
		{"bad generated time", mutateJSON(t, indexFixture(), func(index *Index) { index.GeneratedAt = "2026-08-30T00:00:00+00:00" }), MalformedMetadata},
		{"minimum newer than release", mutateJSON(t, indexFixture(), func(index *Index) { index.Channels[0].Releases[0].MinimumSourceVersion = "1.3.0" }), MalformedMetadata},
		{"unknown channel", mutateJSON(t, indexFixture(), func(index *Index) { index.Channels[0].Name = "nightly" }), MalformedMetadata},
		{"prerelease stable", mutateJSON(t, indexFixture(), func(index *Index) { index.Channels[0].Releases[0].Version = "1.2.0-rc.1" }), MalformedMetadata},
		{"bad status", mutateJSON(t, indexFixture(), func(index *Index) { index.Channels[0].Releases[0].Status = "deleted" }), MalformedMetadata},
		{"bad commit", mutateJSON(t, indexFixture(), func(index *Index) { index.Channels[0].Releases[0].SourceCommit = strings.Repeat("A", 40) }), MalformedMetadata},
		{"unsafe notes URL", mutateJSON(t, indexFixture(), func(index *Index) { index.Channels[0].Releases[0].ReleaseNotesURL = "http://example.invalid/notes" }), MalformedMetadata},
		{"duplicate route", mutateJSON(t, indexFixture(), func(index *Index) {
			route := index.Channels[0].Releases[0].MigrationRoutes[0]
			index.Channels[0].Releases[0].MigrationRoutes = append(index.Channels[0].Releases[0].MigrationRoutes, route)
		}), MalformedMetadata},
		{"wrong platform", mutateJSON(t, indexFixture(), func(index *Index) { index.Channels[0].Releases[0].Artifacts[0].Platform = "linux-arm64" }), MalformedMetadata},
		{"wrong artifact name", mutateJSON(t, indexFixture(), func(index *Index) { index.Channels[0].Releases[0].Artifacts[0].Name = "other.tar.gz" }), MalformedMetadata},
		{"unsafe artifact URL", mutateJSON(t, indexFixture(), func(index *Index) {
			index.Channels[0].Releases[0].Artifacts[0].URL = "https://user:secret@example.invalid/file"
		}), MalformedMetadata},
		{"bad artifact size", mutateJSON(t, indexFixture(), func(index *Index) { index.Channels[0].Releases[0].Artifacts[0].Size = 0 }), MalformedMetadata},
		{"uppercase digest", mutateJSON(t, indexFixture(), func(index *Index) { index.Channels[0].Releases[0].Artifacts[0].SHA256 = strings.Repeat("A", 64) }), MalformedMetadata},
		{"duplicate release identity", mutateJSON(t, indexFixture(), func(index *Index) {
			index.Channels[0].Releases = append(index.Channels[0].Releases, index.Channels[0].Releases[0])
		}), MalformedMetadata},
		{"unknown signature algorithm", mutateJSON(t, indexFixture(), func(index *Index) {
			index.Signatures = []Signature{{Algorithm: "rsa", KeyID: "key", Value: base64.StdEncoding.EncodeToString(make([]byte, ed25519.SignatureSize))}}
		}), MalformedMetadata},
		{"malformed signature", mutateJSON(t, indexFixture(), func(index *Index) {
			index.Signatures = []Signature{{Algorithm: "ed25519", KeyID: "key", Value: "not-base64"}}
		}), MalformedMetadata},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			_, err := Parse(tc.payload)
			if FailureOf(err) != tc.failure {
				t.Fatalf("failure=%s err=%v", FailureOf(err), err)
			}
		})
	}
	tooLarge := make([]byte, MaxIndexBytes+1)
	if _, err := Parse(tooLarge); FailureOf(err) != MalformedMetadata {
		t.Fatalf("oversize failure=%s", FailureOf(err))
	}
}

func TestAuthenticatedIndexIsImmutableFromCallerCopies(t *testing.T) {
	index, publicKey := signedFixture(t)
	verifier, _ := NewVerifier(map[string]ed25519.PublicKey{"community-test-2026": publicKey})
	authenticated, err := verifier.Verify(index)
	if err != nil {
		t.Fatal(err)
	}
	index.Channels[0].Releases[0].Version = "1.9.9"
	copyIndex := authenticated.Index()
	copyIndex.Channels[0].Releases[0].Version = "1.8.8"
	if got := authenticated.Index().Channels[0].Releases[0].Version; got != "1.2.0" {
		t.Fatalf("authenticated metadata mutated through caller alias: %s", got)
	}
}

func TestCollectionAndNestingBounds(t *testing.T) {
	tests := []struct {
		name   string
		mutate func(*Index)
	}{
		{"channels", func(index *Index) {
			for len(index.Channels) <= MaxChannels {
				index.Channels = append(index.Channels, Channel{Name: "stable"})
			}
		}},
		{"releases", func(index *Index) {
			release := index.Channels[0].Releases[0]
			for len(index.Channels[0].Releases) <= MaxReleases {
				copyRelease := release
				copyRelease.Version = "1.2." + string(rune('0'+len(index.Channels[0].Releases)%10))
				index.Channels[0].Releases = append(index.Channels[0].Releases, copyRelease)
			}
		}},
		{"artifacts", func(index *Index) {
			artifact := index.Channels[0].Releases[0].Artifacts[0]
			for len(index.Channels[0].Releases[0].Artifacts) <= MaxArtifacts {
				index.Channels[0].Releases[0].Artifacts = append(index.Channels[0].Releases[0].Artifacts, artifact)
			}
		}},
		{"routes", func(index *Index) {
			for len(index.Channels[0].Releases[0].MigrationRoutes) <= MaxMigrationRoutes {
				index.Channels[0].Releases[0].MigrationRoutes = append(index.Channels[0].Releases[0].MigrationRoutes, "route-"+string(rune('a'+len(index.Channels[0].Releases[0].MigrationRoutes)%26)))
			}
		}},
		{"signatures", func(index *Index) {
			value := base64.StdEncoding.EncodeToString(make([]byte, ed25519.SignatureSize))
			for len(index.Signatures) <= MaxSignatures {
				index.Signatures = append(index.Signatures, Signature{Algorithm: "ed25519", KeyID: "key-" + string(rune('a'+len(index.Signatures))), Value: value})
			}
		}},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			index := indexFixture()
			tc.mutate(&index)
			payload, _ := json.Marshal(index)
			if _, err := Parse(payload); FailureOf(err) != MalformedMetadata {
				t.Fatalf("failure=%s err=%v", FailureOf(err), err)
			}
		})
	}
	nested := []byte(`[[[[[[[[[[null]]]]]]]]]]`)
	if _, err := Parse(nested); FailureOf(err) != MalformedMetadata {
		t.Fatalf("nesting failure=%s", FailureOf(err))
	}
}

func signedFixture(t *testing.T) (Index, ed25519.PublicKey) {
	t.Helper()
	seed := [ed25519.SeedSize]byte{1, 2, 3, 4, 5}
	privateKey := ed25519.NewKeyFromSeed(seed[:])
	index := indexFixture()
	message, err := SigningBytes(index)
	if err != nil {
		t.Fatal(err)
	}
	index.Signatures = []Signature{{Algorithm: "ed25519", KeyID: "community-test-2026", Value: base64.StdEncoding.EncodeToString(ed25519.Sign(privateKey, message))}}
	return index, privateKey.Public().(ed25519.PublicKey)
}

func indexFixture() Index {
	return Index{
		Schema: Schema, Product: Product, GeneratedAt: "2026-08-30T00:00:00Z",
		Channels: []Channel{{Name: "stable", Releases: []Release{{
			Version: "1.2.0", PublishedAt: "2026-08-29T16:53:34Z", Status: "active", SourceCommit: "348d927ffcf4c8cd4c9a50fc3eacad71d8bfe5c2",
			ReleaseNotesURL: "https://git.quantumwizard.hu/Quantum_Wizard_Studio/qwsg/releases/tag/v1.2.0", MinimumSourceVersion: "1.1.0", MigrationRoutes: []string{"compat-1.2.0-rc.2-to-1.2.0"},
			Artifacts: []Artifact{{Platform: "linux-amd64", Name: "qwsg-1.2.0-linux-amd64.tar.gz", URL: "https://git.quantumwizard.hu/Quantum_Wizard_Studio/qwsg/releases/download/v1.2.0/qwsg-1.2.0-linux-amd64.tar.gz", Size: 3524214, SHA256: "44768af20c8456cde09f940590b8c4446f605b2af02866e1553705a01d1a4c11"}},
		}}}},
	}
}

func mutateJSON(t *testing.T, index Index, mutate func(*Index)) []byte {
	t.Helper()
	mutate(&index)
	payload, err := json.Marshal(index)
	if err != nil {
		t.Fatal(err)
	}
	return payload
}
