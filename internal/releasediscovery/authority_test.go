package releasediscovery

import (
	"crypto/ed25519"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"testing"
)

func TestProductionAuthorityIdentityIsExact(t *testing.T) {
	identity, publicKey, err := ProductionTrustIdentity()
	if err != nil {
		t.Fatal(err)
	}
	digest := sha256.Sum256(publicKey)
	if identity.KeyID != "qwsg-community-release-2026-01" || identity.PublicKeyBase64 != "r+iDDJJlGRzU/1bv7aSlVl63PcipILaGmdk7130drHQ=" || identity.Fingerprint != "0d17ea178bb27820d5c7ca44c539dbf9d6ec1e399b29c536252e4658a5d1dcf6" || len(publicKey) != ed25519.PublicKeySize || hex.EncodeToString(digest[:]) != identity.Fingerprint {
		t.Fatalf("unexpected production identity: %+v key_bytes=%d digest=%s", identity, len(publicKey), hex.EncodeToString(digest[:]))
	}
	decoded, err := base64.StdEncoding.Strict().DecodeString(identity.PublicKeyBase64)
	if err != nil || string(decoded) != string(publicKey) {
		t.Fatal("embedded public key encoding mismatch")
	}
}

func TestProductionAuthoritySourceAndWrongSignaturesFailClosed(t *testing.T) {
	if ProductionEndpoint != "https://releases.quantumwizard.hu/qwsg/v1/release-index.json" || ProductionSourceID != "community-release-index" {
		t.Fatalf("unexpected production source: %s %s", ProductionSourceID, ProductionEndpoint)
	}
	verifier, err := ProductionVerifier()
	if err != nil {
		t.Fatal(err)
	}
	index := indexFixture()
	index.Signatures = []Signature{{Algorithm: "ed25519", KeyID: ProductionKeyID, Value: base64.StdEncoding.EncodeToString(make([]byte, ed25519.SignatureSize))}}
	if _, err = verifier.Verify(index); FailureOf(err) != UnauthenticatedMetadata {
		t.Fatalf("wrong production signature failure=%s", FailureOf(err))
	}
	index.Signatures[0].KeyID = "unknown-production-key"
	if _, err = verifier.Verify(index); FailureOf(err) != UnauthenticatedMetadata {
		t.Fatalf("unknown production key failure=%s", FailureOf(err))
	}
}
