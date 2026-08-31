package releasepublication

import (
	"crypto/ed25519"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"quantumwizard.hu/qwsg/internal/releasediscovery"
)

func TestFrozenProductionIndexAndCheckpoint(t *testing.T) {
	root := filepath.Join("..", "..", "release", "production")
	signed, err := os.ReadFile(filepath.Join(root, "qwsg-release-index-first-signed.json"))
	if err != nil {
		t.Fatal(err)
	}
	digest := sha256.Sum256(signed)
	if len(signed) != 918 || hex.EncodeToString(digest[:]) != "f9f95bf28d463a8403841d9cc56d817c248f1e0a01e3e65a5a9e1afc16d39704" {
		t.Fatalf("frozen signed index identity changed: size=%d sha256=%s", len(signed), hex.EncodeToString(digest[:]))
	}
	checkpoint, err := BuildCheckpoint(signed)
	if err != nil {
		t.Fatal(err)
	}
	want, err := os.ReadFile(filepath.Join(root, "qwsg-release-index-first-checkpoint.json"))
	if err != nil || string(checkpoint) != string(want) {
		t.Fatalf("checkpoint changed: err=%v got=%s want=%s", err, checkpoint, want)
	}
}

func TestGenerateAndAssembleAreDeterministic(t *testing.T) {
	seed := make([]byte, ed25519.SeedSize)
	for index := range seed {
		seed[index] = byte(index + 1)
	}
	privateKey := ed25519.NewKeyFromSeed(seed)
	candidate := fixture()
	first, err := Generate(candidate)
	if err != nil {
		t.Fatal(err)
	}
	second, err := Generate(append([]byte(nil), candidate...))
	if err != nil || string(first) != string(second) || strings.Contains(string(first), "signatures") {
		t.Fatalf("generation is not deterministic: %v", err)
	}
	signature := []byte(base64.StdEncoding.EncodeToString(ed25519.Sign(privateKey, first)) + "\n")
	signed, err := Assemble(first, signature, releasediscovery.ProductionKeyID)
	if err != nil {
		t.Fatal(err)
	}
	index, err := releasediscovery.Parse(signed)
	if err != nil || len(index.Signatures) != 1 || index.Signatures[0].KeyID != releasediscovery.ProductionKeyID {
		t.Fatalf("assembled index=%+v err=%v", index, err)
	}
	verifier, _ := releasediscovery.NewVerifier(map[string]ed25519.PublicKey{releasediscovery.ProductionKeyID: privateKey.Public().(ed25519.PublicKey)})
	if _, err = verifier.Verify(index); err != nil {
		t.Fatal(err)
	}
}

func TestPublicationInputsFailClosed(t *testing.T) {
	valid, err := Generate(fixture())
	if err != nil {
		t.Fatal(err)
	}
	tests := [][]byte{
		[]byte(`{"schema":"qwsg.release-index/1"}`),
		[]byte(strings.Replace(string(fixture()), "git.quantumwizard.hu", "github.com", 1)),
		append(fixture(), []byte(` trailing`)...),
	}
	for _, input := range tests {
		if _, err = Generate(input); err == nil {
			t.Fatalf("unsafe candidate accepted: %s", input)
		}
	}
	if _, err = Assemble(valid, []byte("not-base64\n"), releasediscovery.ProductionKeyID); err == nil {
		t.Fatal("malformed signature accepted")
	}
	if _, err = Assemble(valid, []byte(base64.StdEncoding.EncodeToString(make([]byte, ed25519.SignatureSize))+"\n"), "wrong-key"); err == nil {
		t.Fatal("wrong key identity accepted")
	}
	if err = VerifyProduction(valid); err == nil {
		t.Fatal("unsigned production candidate accepted")
	}
}

func fixture() []byte {
	value := releasediscovery.Index{
		Schema: releasediscovery.Schema, Product: releasediscovery.Product, GeneratedAt: "2026-08-30T00:00:00Z",
		Channels: []releasediscovery.Channel{{Name: "stable", Releases: []releasediscovery.Release{{
			Version: "1.2.0", PublishedAt: "2026-08-29T16:53:34Z", Status: "active", SourceCommit: "348d927ffcf4c8cd4c9a50fc3eacad71d8bfe5c2",
			ReleaseNotesURL: "https://git.quantumwizard.hu/Quantum_Wizard_Studio/qwsg/releases/tag/v1.2.0", MinimumSourceVersion: "1.1.0", MigrationRoutes: []string{"compat-1.2.0-rc.2-to-1.2.0"},
			Artifacts: []releasediscovery.Artifact{{Platform: "linux-amd64", Name: "qwsg-1.2.0-linux-amd64.tar.gz", URL: "https://git.quantumwizard.hu/Quantum_Wizard_Studio/qwsg/releases/download/v1.2.0/qwsg-1.2.0-linux-amd64.tar.gz", Size: 3524214, SHA256: "44768af20c8456cde09f940590b8c4446f605b2af02866e1553705a01d1a4c11"}},
		}}}},
	}
	payload, _ := json.Marshal(value)
	return payload
}
