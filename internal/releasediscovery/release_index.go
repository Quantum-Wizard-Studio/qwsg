// Package releasediscovery owns QWSG's read-only public release metadata
// contract. It does not acquire artifacts, persist awareness, notify, schedule,
// install, migrate, or cross a privileged boundary.
package releasediscovery

import (
	"bytes"
	"crypto/ed25519"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/url"
	"path"
	"strings"
	"time"
	"unicode/utf8"

	"quantumwizard.hu/qwsg/internal/update"
)

const (
	Schema             = "qwsg.release-index/1"
	Product            = "qwsg"
	MediaType          = "application/vnd.quantumwizard.qwsg-releases+json"
	MaxIndexBytes      = 1 << 20
	MaxChannels        = 4
	MaxReleases        = 50
	MaxArtifacts       = 4
	MaxMigrationRoutes = 32
	MaxSignatures      = 8
	maxJSONDepth       = 8
	maxArtifactBytes   = 128 << 20
)

type Index struct {
	Schema      string      `json:"schema"`
	Product     string      `json:"product"`
	GeneratedAt string      `json:"generated_at"`
	Channels    []Channel   `json:"channels"`
	Signatures  []Signature `json:"signatures"`
}

type Channel struct {
	Name     string    `json:"name"`
	Releases []Release `json:"releases"`
}

type Release struct {
	Version              string     `json:"version"`
	PublishedAt          string     `json:"published_at"`
	Status               string     `json:"status"`
	SourceCommit         string     `json:"source_commit"`
	ReleaseNotesURL      string     `json:"release_notes_url"`
	MinimumSourceVersion string     `json:"minimum_source_version"`
	MigrationRoutes      []string   `json:"migration_routes"`
	Artifacts            []Artifact `json:"artifacts"`
}

type Artifact struct {
	Platform string `json:"platform"`
	Name     string `json:"name"`
	URL      string `json:"url"`
	Size     int64  `json:"size"`
	SHA256   string `json:"sha256"`
}

type Signature struct {
	Algorithm string `json:"algorithm"`
	KeyID     string `json:"key_id"`
	Value     string `json:"value"`
}

type Failure string

const (
	MalformedMetadata       Failure = "malformed_metadata"
	UnsupportedContract     Failure = "unsupported_contract"
	UnauthenticatedMetadata Failure = "unauthenticated_metadata"
	SourceAuthority         Failure = "source_authority_refused"
	SourceCanceled          Failure = "source_canceled"
	SourceTimeout           Failure = "source_timeout"
	SourceTransport         Failure = "source_transport_failed"
	SourceHTTPStatus        Failure = "source_http_status"
	SourceTooLarge          Failure = "source_too_large"
	SourceMediaType         Failure = "source_media_type"
	InstalledUnverified     Failure = "installed_identity_unverified"
	NoEligibleRelease       Failure = "no_eligible_release"
)

type ContractError struct{ Category Failure }

func (e *ContractError) Error() string { return string(e.Category) }

func fail(category Failure) error { return &ContractError{Category: category} }

func FailureOf(err error) Failure {
	if err == nil {
		return ""
	}
	var contract *ContractError
	if errors.As(err, &contract) {
		return contract.Category
	}
	return SourceTransport
}

func Parse(payload []byte) (Index, error) {
	if len(payload) == 0 || len(payload) > MaxIndexBytes || !utf8.Valid(payload) {
		return Index{}, fail(MalformedMetadata)
	}
	if err := rejectDuplicateMembers(payload); err != nil {
		return Index{}, fail(MalformedMetadata)
	}
	decoder := json.NewDecoder(bytes.NewReader(payload))
	decoder.DisallowUnknownFields()
	var index Index
	if err := decoder.Decode(&index); err != nil {
		return Index{}, fail(MalformedMetadata)
	}
	var extra any
	if decoder.Decode(&extra) != io.EOF {
		return Index{}, fail(MalformedMetadata)
	}
	if err := validate(index); err != nil {
		return Index{}, err
	}
	return index, nil
}

func rejectDuplicateMembers(payload []byte) error {
	decoder := json.NewDecoder(bytes.NewReader(payload))
	decoder.UseNumber()
	if err := walkJSON(decoder, 0); err != nil {
		return err
	}
	if _, err := decoder.Token(); err != io.EOF {
		return fmt.Errorf("trailing JSON")
	}
	return nil
}

func walkJSON(decoder *json.Decoder, depth int) error {
	if depth > maxJSONDepth {
		return fmt.Errorf("JSON nesting exceeds bound")
	}
	token, err := decoder.Token()
	if err != nil {
		return err
	}
	delimiter, compound := token.(json.Delim)
	if !compound {
		return nil
	}
	switch delimiter {
	case '{':
		seen := make(map[string]struct{})
		for decoder.More() {
			nameToken, nameErr := decoder.Token()
			if nameErr != nil {
				return nameErr
			}
			name, ok := nameToken.(string)
			if !ok {
				return fmt.Errorf("invalid JSON member")
			}
			if _, exists := seen[name]; exists {
				return fmt.Errorf("duplicate JSON member")
			}
			seen[name] = struct{}{}
			if err = walkJSON(decoder, depth+1); err != nil {
				return err
			}
		}
	case '[':
		for decoder.More() {
			if err = walkJSON(decoder, depth+1); err != nil {
				return err
			}
		}
	default:
		return fmt.Errorf("invalid JSON delimiter")
	}
	closing, err := decoder.Token()
	if err != nil {
		return err
	}
	want := json.Delim('}')
	if delimiter == '[' {
		want = ']'
	}
	if closing != want {
		return fmt.Errorf("invalid JSON close")
	}
	return nil
}

func validate(index Index) error {
	if index.Schema != Schema || index.Product != Product {
		return fail(UnsupportedContract)
	}
	generated, ok := canonicalTime(index.GeneratedAt)
	if !ok || len(index.Channels) == 0 || len(index.Channels) > MaxChannels || len(index.Signatures) > MaxSignatures {
		return fail(MalformedMetadata)
	}
	channels, versions, signatureIDs := map[string]bool{}, map[string]bool{}, map[string]bool{}
	for _, channel := range index.Channels {
		if !validChannel(channel.Name) || channels[channel.Name] || len(channel.Releases) == 0 || len(channel.Releases) > MaxReleases {
			return fail(MalformedMetadata)
		}
		channels[channel.Name] = true
		for _, release := range channel.Releases {
			parsed, err := update.ParseVersion(release.Version)
			published, publishedOK := canonicalTime(release.PublishedAt)
			minimum, minimumErr := update.ParseVersion(release.MinimumSourceVersion)
			if err != nil || parsed.Major != 1 || !publishedOK || published.After(generated) || minimumErr != nil || minimum.Major != 1 || update.Compare(minimum, parsed) > 0 || versions[release.Version] {
				return fail(MalformedMetadata)
			}
			versions[release.Version] = true
			if (channel.Name == "stable" && len(parsed.Prerelease) != 0) || (channel.Name == "preview" && len(parsed.Prerelease) == 0) {
				return fail(MalformedMetadata)
			}
			if release.Status != "active" && release.Status != "withdrawn" {
				return fail(MalformedMetadata)
			}
			if len(release.SourceCommit) != 40 || !lowerHex(release.SourceCommit) || !safeHTTPSURL(release.ReleaseNotesURL, "") {
				return fail(MalformedMetadata)
			}
			if len(release.MigrationRoutes) > MaxMigrationRoutes || !uniqueTokens(release.MigrationRoutes) {
				return fail(MalformedMetadata)
			}
			if len(release.Artifacts) == 0 || len(release.Artifacts) > MaxArtifacts {
				return fail(MalformedMetadata)
			}
			platforms := map[string]bool{}
			for _, artifact := range release.Artifacts {
				want := "qwsg-" + release.Version + "-" + artifact.Platform + ".tar.gz"
				if artifact.Platform != "linux-amd64" || platforms[artifact.Platform] || artifact.Name != want || artifact.Size <= 0 || artifact.Size > maxArtifactBytes || len(artifact.SHA256) != 64 || !lowerHex(artifact.SHA256) || !safeHTTPSURL(artifact.URL, artifact.Name) {
					return fail(MalformedMetadata)
				}
				platforms[artifact.Platform] = true
			}
		}
	}
	for _, signature := range index.Signatures {
		decoded, err := base64.StdEncoding.Strict().DecodeString(signature.Value)
		if signature.Algorithm != "ed25519" || !safeToken(signature.KeyID, 64) || signatureIDs[signature.KeyID] || err != nil || len(decoded) != ed25519.SignatureSize {
			return fail(MalformedMetadata)
		}
		signatureIDs[signature.KeyID] = true
	}
	return nil
}

func canonicalTime(raw string) (time.Time, bool) {
	if len(raw) != len("2006-01-02T15:04:05Z") || !strings.HasSuffix(raw, "Z") {
		return time.Time{}, false
	}
	value, err := time.Parse(time.RFC3339, raw)
	return value, err == nil && value.Format(time.RFC3339) == raw
}

func validChannel(value string) bool { return value == "stable" || value == "preview" }

func uniqueTokens(values []string) bool {
	seen := map[string]bool{}
	for _, value := range values {
		if !safeToken(value, 128) || seen[value] {
			return false
		}
		seen[value] = true
	}
	return true
}

func safeToken(value string, bound int) bool {
	if value == "" || len(value) > bound {
		return false
	}
	for _, r := range value {
		if !(r >= 'a' && r <= 'z') && !(r >= '0' && r <= '9') && r != '-' && r != '.' {
			return false
		}
	}
	return true
}

func safeHTTPSURL(raw, basename string) bool {
	if len(raw) == 0 || len(raw) > 2048 {
		return false
	}
	parsed, err := url.Parse(raw)
	if err != nil || parsed.Scheme != "https" || parsed.Host == "" || parsed.User != nil || parsed.RawQuery != "" || parsed.Fragment != "" || parsed.Opaque != "" || parsed.EscapedPath() != parsed.Path || path.Clean(parsed.Path) != parsed.Path || parsed.Path == "/" {
		return false
	}
	return basename == "" || path.Base(parsed.EscapedPath()) == basename
}

func lowerHex(value string) bool {
	if value == "" || strings.ToLower(value) != value {
		return false
	}
	for _, r := range value {
		if !(r >= '0' && r <= '9') && !(r >= 'a' && r <= 'f') {
			return false
		}
	}
	return true
}

type signedIndex struct {
	Schema      string    `json:"schema"`
	Product     string    `json:"product"`
	GeneratedAt string    `json:"generated_at"`
	Channels    []Channel `json:"channels"`
}

func SigningBytes(index Index) ([]byte, error) {
	if err := validate(index); err != nil {
		return nil, err
	}
	return json.Marshal(signedIndex{Schema: index.Schema, Product: index.Product, GeneratedAt: index.GeneratedAt, Channels: index.Channels})
}

type AuthenticityEvidence struct {
	Scheme string `json:"scheme"`
	KeyID  string `json:"key_id"`
}

type AuthenticatedIndex struct {
	index    Index
	evidence AuthenticityEvidence
}

func (a AuthenticatedIndex) Index() Index                       { return cloneIndex(a.index) }
func (a AuthenticatedIndex) Authenticity() AuthenticityEvidence { return a.evidence }

type Verifier struct{ anchors map[string]ed25519.PublicKey }

func NewVerifier(anchors map[string]ed25519.PublicKey) (Verifier, error) {
	copySet := make(map[string]ed25519.PublicKey, len(anchors))
	for keyID, publicKey := range anchors {
		if !safeToken(keyID, 64) || len(publicKey) != ed25519.PublicKeySize {
			return Verifier{}, fail(UnauthenticatedMetadata)
		}
		copySet[keyID] = append(ed25519.PublicKey(nil), publicKey...)
	}
	return Verifier{anchors: copySet}, nil
}

func (v Verifier) Verify(index Index) (AuthenticatedIndex, error) {
	message, err := SigningBytes(index)
	if err != nil {
		return AuthenticatedIndex{}, err
	}
	for _, signature := range index.Signatures {
		publicKey, approved := v.anchors[signature.KeyID]
		if !approved {
			continue
		}
		value, decodeErr := base64.StdEncoding.Strict().DecodeString(signature.Value)
		if decodeErr == nil && ed25519.Verify(publicKey, message, value) {
			return AuthenticatedIndex{index: cloneIndex(index), evidence: AuthenticityEvidence{Scheme: "ed25519", KeyID: signature.KeyID}}, nil
		}
	}
	return AuthenticatedIndex{}, fail(UnauthenticatedMetadata)
}

func cloneIndex(index Index) Index {
	result := index
	result.Signatures = append([]Signature(nil), index.Signatures...)
	result.Channels = make([]Channel, len(index.Channels))
	for channelIndex, channel := range index.Channels {
		result.Channels[channelIndex] = channel
		result.Channels[channelIndex].Releases = make([]Release, len(channel.Releases))
		for releaseIndex, release := range channel.Releases {
			result.Channels[channelIndex].Releases[releaseIndex] = release
			result.Channels[channelIndex].Releases[releaseIndex].MigrationRoutes = append([]string(nil), release.MigrationRoutes...)
			result.Channels[channelIndex].Releases[releaseIndex].Artifacts = append([]Artifact(nil), release.Artifacts...)
		}
	}
	return result
}
