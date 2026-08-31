// Package releasepublication prepares deterministic release-index bytes for an
// offline signing and no-clobber publication workflow. It never performs a
// network request or publishes an object.
package releasepublication

import (
	"bytes"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"net/url"
	"path"
	"strings"

	"quantumwizard.hu/qwsg/internal/releasediscovery"
)

const forgejoPrefix = "/Quantum_Wizard_Studio/qwsg/"

func Generate(candidate []byte) ([]byte, error) {
	index, err := releasediscovery.Parse(candidate)
	if err != nil || len(index.Signatures) != 0 || !canonicalForgejoProvenance(index) {
		return nil, fmt.Errorf("invalid unsigned publication candidate")
	}
	return releasediscovery.SigningBytes(index)
}

func Assemble(unsigned, signatureText []byte, keyID string) ([]byte, error) {
	canonical, err := Generate(unsigned)
	if err != nil || !bytes.Equal(canonical, unsigned) || keyID != releasediscovery.ProductionKeyID {
		return nil, fmt.Errorf("invalid canonical signing input")
	}
	signature, err := base64.StdEncoding.Strict().DecodeString(strings.TrimSuffix(string(signatureText), "\n"))
	if err != nil || len(signature) != 64 || bytes.Contains(signatureText, []byte("\r")) || bytes.Count(signatureText, []byte("\n")) > 1 {
		return nil, fmt.Errorf("invalid detached signature")
	}
	index, err := releasediscovery.Parse(unsigned)
	if err != nil {
		return nil, err
	}
	index.Signatures = []releasediscovery.Signature{{Algorithm: "ed25519", KeyID: keyID, Value: base64.StdEncoding.EncodeToString(signature)}}
	return json.Marshal(index)
}

func VerifyProduction(signed []byte) error {
	index, err := releasediscovery.Parse(signed)
	if err != nil || !canonicalForgejoProvenance(index) {
		return fmt.Errorf("invalid signed publication candidate")
	}
	canonical, err := json.Marshal(index)
	if err != nil || !bytes.Equal(canonical, signed) {
		return fmt.Errorf("signed candidate is not canonical")
	}
	verifier, err := releasediscovery.ProductionVerifier()
	if err != nil {
		return err
	}
	_, err = verifier.Verify(index)
	return err
}

type Checkpoint struct {
	Schema                string `json:"schema"`
	SourceID              string `json:"source_id"`
	Endpoint              string `json:"endpoint"`
	KeyID                 string `json:"key_id"`
	PublicKeyFingerprint  string `json:"public_key_fingerprint_sha256"`
	IndexSHA256           string `json:"index_sha256"`
	IndexSize             int    `json:"index_size"`
	PublicationAuthorized bool   `json:"publication_authorized"`
}

func BuildCheckpoint(signed []byte) ([]byte, error) {
	if err := VerifyProduction(signed); err != nil {
		return nil, err
	}
	digest := sha256.Sum256(signed)
	record := Checkpoint{
		Schema: "qwsg.release-publication-checkpoint/1", SourceID: releasediscovery.ProductionSourceID,
		Endpoint: releasediscovery.ProductionEndpoint, KeyID: releasediscovery.ProductionKeyID,
		PublicKeyFingerprint: releasediscovery.ProductionFingerprint, IndexSHA256: hex.EncodeToString(digest[:]),
		IndexSize: len(signed), PublicationAuthorized: false,
	}
	return json.Marshal(record)
}

func canonicalForgejoProvenance(index releasediscovery.Index) bool {
	for _, channel := range index.Channels {
		for _, release := range channel.Releases {
			if !forgejoURL(release.ReleaseNotesURL, "/releases/tag/v"+release.Version) {
				return false
			}
			for _, artifact := range release.Artifacts {
				if !forgejoURL(artifact.URL, "/releases/download/v"+release.Version+"/"+artifact.Name) {
					return false
				}
			}
		}
	}
	return true
}

func forgejoURL(raw, suffix string) bool {
	parsed, err := url.Parse(raw)
	return err == nil && parsed.Scheme == "https" && parsed.Host == "git.quantumwizard.hu" && strings.HasPrefix(parsed.Path, forgejoPrefix) && strings.TrimPrefix(parsed.Path, forgejoPrefix) == strings.TrimPrefix(path.Clean(forgejoPrefix+suffix), forgejoPrefix)
}
