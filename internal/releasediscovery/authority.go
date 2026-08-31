package releasediscovery

import (
	"crypto/ed25519"
	"crypto/sha256"
	_ "embed"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"io"
	"strings"
)

const (
	ProductionSourceID        = "community-release-index"
	ProductionEndpoint        = "https://releases.quantumwizard.hu/qwsg/v1/release-index.json"
	ProductionKeyID           = "qwsg-community-release-2026-01"
	ProductionPublicKeyBase64 = "r+iDDJJlGRzU/1bv7aSlVl63PcipILaGmdk7130drHQ="
	ProductionFingerprint     = "0d17ea178bb27820d5c7ca44c539dbf9d6ec1e399b29c536252e4658a5d1dcf6"
	TrustSchema               = "qwsg.release-trust/1"
)

//go:embed trust/production.json
var productionTrustJSON []byte

type TrustIdentity struct {
	Schema          string `json:"schema"`
	KeyID           string `json:"key_id"`
	Algorithm       string `json:"algorithm"`
	PublicKeyBase64 string `json:"public_key_base64"`
	Fingerprint     string `json:"fingerprint_sha256"`
	Epoch           uint64 `json:"epoch"`
	Status          string `json:"status"`
}

func ProductionTrustIdentity() (TrustIdentity, ed25519.PublicKey, error) {
	decoder := json.NewDecoder(strings.NewReader(string(productionTrustJSON)))
	decoder.DisallowUnknownFields()
	var identity TrustIdentity
	if err := decoder.Decode(&identity); err != nil {
		return TrustIdentity{}, nil, fail(SourceAuthority)
	}
	var trailing any
	if decoder.Decode(&trailing) != io.EOF {
		return TrustIdentity{}, nil, fail(SourceAuthority)
	}
	publicKey, err := base64.StdEncoding.Strict().DecodeString(identity.PublicKeyBase64)
	digest := sha256.Sum256(publicKey)
	if err != nil || len(publicKey) != ed25519.PublicKeySize || identity.Schema != TrustSchema || identity.KeyID != ProductionKeyID || identity.Algorithm != "ed25519" || identity.PublicKeyBase64 != ProductionPublicKeyBase64 || identity.Fingerprint != ProductionFingerprint || identity.Fingerprint != hex.EncodeToString(digest[:]) || identity.Epoch != 1 || identity.Status != "active" {
		return TrustIdentity{}, nil, fail(SourceAuthority)
	}
	return identity, ed25519.PublicKey(append([]byte(nil), publicKey...)), nil
}

func ProductionVerifier() (Verifier, error) {
	identity, publicKey, err := ProductionTrustIdentity()
	if err != nil {
		return Verifier{}, err
	}
	return NewVerifier(map[string]ed25519.PublicKey{identity.KeyID: publicKey})
}

func ProductionDiscoverer() (Discoverer, error) {
	source, err := NewStaticHTTPSource(ProductionEndpoint, ProductionSourceID, nil)
	if err != nil {
		return Discoverer{}, err
	}
	verifier, err := ProductionVerifier()
	if err != nil {
		return Discoverer{}, err
	}
	return NewDiscoverer(source, verifier, LocalEvaluator())
}
