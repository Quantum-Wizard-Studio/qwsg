// Package updateawareness owns QWSG's local, read-only release-awareness
// record. It does not schedule checks, notify, acquire artifacts, or install.
package updateawareness

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"strings"
	"time"

	"quantumwizard.hu/qwsg/internal/installation"
	"quantumwizard.hu/qwsg/internal/releasediscovery"
	"quantumwizard.hu/qwsg/internal/update"
)

const (
	Schema         = "qwsg.update-awareness/1"
	DigestScheme   = "sha256"
	MaxEncodedSize = 64 << 10
	DefaultFresh   = 48 * time.Hour
)

var (
	ErrMissing      = errors.New("update awareness state missing")
	ErrCorrupt      = errors.New("update awareness state corrupt")
	ErrIncompatible = errors.New("update awareness state incompatible")
	ErrUnsafePath   = errors.New("update awareness state path unsafe")
	ErrPermission   = errors.New("update awareness state permission denied")
	ErrContended    = errors.New("update awareness check already active")
)

type Status string

const (
	NeverChecked      Status = "never_checked"
	Current           Status = "current"
	UpdateAvailable   Status = "update_available"
	UpdateUnsupported Status = "update_available_unsupported_source"
	Withdrawn         Status = "withdrawn"
	Unknown           Status = "unknown"
)

type Outcome string

const (
	AttemptSuccess Outcome = "success"
	AttemptFailure Outcome = "failure"
)

type Installed struct {
	Classification installation.State `json:"classification"`
	Version        string             `json:"version"`
}

type Attempt struct {
	At      time.Time `json:"at"`
	Outcome Outcome   `json:"outcome"`
	Failure string    `json:"failure,omitempty"`
}

type Observation struct {
	ObservedAt             time.Time                             `json:"observed_at"`
	FreshUntil             time.Time                             `json:"fresh_until"`
	SourceID               string                                `json:"source_id"`
	Channel                string                                `json:"channel"`
	IndexGeneratedAt       string                                `json:"index_generated_at"`
	TransportAuthenticated bool                                  `json:"transport_authenticated"`
	Authenticity           releasediscovery.AuthenticityEvidence `json:"authenticity"`
	Validators             releasediscovery.Validators           `json:"validators,omitempty"`
	Installed              Installed                             `json:"installed"`
	Status                 Status                                `json:"status"`
	Relation               update.Relation                       `json:"relation"`
	Compatibility          releasediscovery.Compatibility        `json:"compatibility"`
	MigrationID            string                                `json:"migration_id,omitempty"`
	ReleaseVersion         string                                `json:"release_version,omitempty"`
	ReleasePublishedAt     string                                `json:"release_published_at,omitempty"`
	ReleaseStatus          string                                `json:"release_status,omitempty"`
	ArtifactName           string                                `json:"artifact_name,omitempty"`
	ArtifactSHA256         string                                `json:"artifact_sha256,omitempty"`
	ArtifactSize           int64                                 `json:"artifact_size,omitempty"`
}

type State struct {
	Schema              string       `json:"schema"`
	DigestScheme        string       `json:"digest_scheme"`
	Digest              string       `json:"digest"`
	SourceID            string       `json:"source_id"`
	Channel             string       `json:"channel"`
	Installed           Installed    `json:"installed"`
	Status              Status       `json:"status"`
	LastAttempt         Attempt      `json:"last_attempt"`
	LastSuccess         *Observation `json:"last_success,omitempty"`
	ConsecutiveFailures uint32       `json:"consecutive_failures"`
}

func NewFailure(previous *State, sourceID, channel string, installed Installed, at time.Time, failure string) (State, error) {
	value := State{Schema: Schema, DigestScheme: DigestScheme, SourceID: sourceID, Channel: channel, Installed: installed, Status: Unknown, LastAttempt: Attempt{At: at.UTC(), Outcome: AttemptFailure, Failure: failure}, ConsecutiveFailures: 1}
	if previous != nil && Validate(*previous) == nil && previous.SourceID == sourceID && previous.Channel == channel && previous.Installed == installed {
		value.LastSuccess = cloneObservation(previous.LastSuccess)
		value.Status = previous.Status
		value.ConsecutiveFailures = previous.ConsecutiveFailures + 1
	}
	return Normalize(value)
}

func NewSuccess(previous *State, result releasediscovery.CheckResult, sourceID, channel string, installed Installed, at time.Time, freshness time.Duration) (State, error) {
	if freshness <= 0 {
		freshness = DefaultFresh
	}
	when := at.UTC()
	var observation Observation
	if result.NotModified {
		if previous == nil || previous.LastSuccess == nil || previous.SourceID != sourceID || previous.Channel != channel || result.Source.SourceID != sourceID || !result.Source.TransportAuthenticated || previous.LastSuccess.Authenticity.Scheme != "ed25519" {
			return State{}, fmt.Errorf("%w: not-modified without authenticated cache", ErrCorrupt)
		}
		observation = *cloneObservation(previous.LastSuccess)
		observation.ObservedAt, observation.FreshUntil = when, when.Add(freshness)
		if result.Source.Validators.ETag != "" {
			observation.Validators.ETag = result.Source.Validators.ETag
		}
		if result.Source.Validators.LastModified != "" {
			observation.Validators.LastModified = result.Source.Validators.LastModified
		}
		observation.Installed = installed
	} else {
		e := result.Evaluation
		if result.Source.SourceID != sourceID || !result.Source.TransportAuthenticated || result.Authenticity.Scheme != "ed25519" || result.Authenticity != e.Authenticity {
			return State{}, ErrCorrupt
		}
		status := Current
		if e.Relation == update.Newer {
			status = UpdateUnsupported
			if e.Compatibility == releasediscovery.CompatibilitySupported {
				status = UpdateAvailable
			}
		}
		observation = Observation{ObservedAt: when, FreshUntil: when.Add(freshness), SourceID: result.Source.SourceID, Channel: channel, IndexGeneratedAt: result.IndexGeneratedAt, TransportAuthenticated: result.Source.TransportAuthenticated, Authenticity: e.Authenticity, Validators: result.Source.Validators, Installed: installed, Status: status, Relation: e.Relation, Compatibility: e.Compatibility, MigrationID: e.MigrationID, ReleaseVersion: e.Release.Version, ReleasePublishedAt: e.Release.PublishedAt, ReleaseStatus: e.Release.Status, ArtifactName: e.Artifact.Name, ArtifactSHA256: e.Artifact.SHA256, ArtifactSize: e.Artifact.Size}
	}
	value := State{Schema: Schema, DigestScheme: DigestScheme, SourceID: sourceID, Channel: channel, Installed: installed, Status: observation.Status, LastAttempt: Attempt{At: when, Outcome: AttemptSuccess}, LastSuccess: &observation}
	return Normalize(value)
}

func NewWithdrawn(previous State, result releasediscovery.CheckResult, installed Installed, at time.Time, freshness time.Duration) (State, error) {
	if previous.LastSuccess == nil || result.Source.SourceID != previous.SourceID || !result.Source.TransportAuthenticated || result.Authenticity.Scheme != "ed25519" || !contains(result.WithdrawnVersions, previous.LastSuccess.ReleaseVersion) {
		return State{}, ErrCorrupt
	}
	if freshness <= 0 {
		freshness = DefaultFresh
	}
	observation := *cloneObservation(previous.LastSuccess)
	observation.ObservedAt, observation.FreshUntil, observation.Installed = at.UTC(), at.UTC().Add(freshness), installed
	observation.Status, observation.ReleaseStatus, observation.TransportAuthenticated = Withdrawn, "withdrawn", result.Source.TransportAuthenticated
	observation.Authenticity, observation.IndexGeneratedAt = result.Authenticity, result.IndexGeneratedAt
	observation.Validators = result.Source.Validators
	value := State{Schema: Schema, DigestScheme: DigestScheme, SourceID: previous.SourceID, Channel: previous.Channel, Installed: installed, Status: Withdrawn, LastAttempt: Attempt{At: at.UTC(), Outcome: AttemptSuccess}, LastSuccess: &observation}
	return Normalize(value)
}

func Normalize(value State) (State, error) {
	value.Schema, value.DigestScheme, value.Digest = Schema, DigestScheme, ""
	if err := validateContent(value); err != nil {
		return State{}, err
	}
	document, _ := json.Marshal(value)
	sum := sha256.Sum256(document)
	value.Digest = hex.EncodeToString(sum[:])
	return value, Validate(value)
}

func Validate(value State) error {
	if value.Schema != Schema || value.DigestScheme != DigestScheme {
		return ErrIncompatible
	}
	claimed := value.Digest
	value.Digest = ""
	if err := validateContent(value); err != nil {
		return err
	}
	document, _ := json.Marshal(value)
	sum := sha256.Sum256(document)
	if claimed != hex.EncodeToString(sum[:]) {
		return ErrCorrupt
	}
	return nil
}

func Marshal(value State) ([]byte, error) {
	if err := Validate(value); err != nil {
		return nil, err
	}
	document, err := json.Marshal(value)
	if err != nil || len(document) > MaxEncodedSize {
		return nil, ErrCorrupt
	}
	return append(document, '\n'), nil
}

func Decode(document []byte) (State, error) {
	if len(document) == 0 || len(document) > MaxEncodedSize {
		return State{}, ErrCorrupt
	}
	decoder := json.NewDecoder(bytes.NewReader(document))
	decoder.DisallowUnknownFields()
	var value State
	if decoder.Decode(&value) != nil {
		return State{}, ErrCorrupt
	}
	var extra any
	if decoder.Decode(&extra) != io.EOF {
		return State{}, ErrCorrupt
	}
	if err := Validate(value); err != nil {
		return State{}, err
	}
	return value, nil
}

func IsStale(value State, now time.Time) bool {
	return value.LastSuccess == nil || !now.UTC().Before(value.LastSuccess.FreshUntil)
}

func Reconcile(value State, installed Installed) (State, bool) {
	return value, value.Installed != installed
}

func validateContent(value State) error {
	if value.Schema != Schema || value.DigestScheme != DigestScheme {
		return ErrIncompatible
	}
	if !safeToken(value.SourceID, 64) || (value.Channel != "stable" && value.Channel != "preview") || value.Installed.Classification != installation.VerifiedSupported || !validVersion(value.Installed.Version) || !utc(value.LastAttempt.At) {
		return ErrCorrupt
	}
	if value.LastAttempt.Outcome == AttemptFailure {
		if !safeToken(value.LastAttempt.Failure, 64) || value.ConsecutiveFailures == 0 {
			return ErrCorrupt
		}
	} else if value.LastAttempt.Outcome != AttemptSuccess || value.LastAttempt.Failure != "" || value.ConsecutiveFailures != 0 {
		return ErrCorrupt
	}
	if value.LastSuccess == nil {
		return nilIf(value.Status != Unknown || value.LastAttempt.Outcome != AttemptFailure)
	}
	o := value.LastSuccess
	if o.SourceID != value.SourceID || o.Channel != value.Channel || o.Installed != value.Installed || !utc(o.ObservedAt) || !utc(o.FreshUntil) || !o.FreshUntil.After(o.ObservedAt) || o.ObservedAt.After(value.LastAttempt.At) || !o.TransportAuthenticated || o.Authenticity.Scheme != "ed25519" || !safeToken(o.Authenticity.KeyID, 64) || !validVersion(o.Installed.Version) {
		return ErrCorrupt
	}
	if o.Status != value.Status || !validStatus(o.Status) || o.ReleaseVersion == "" || !validVersion(o.ReleaseVersion) || !canonicalTime(o.ReleasePublishedAt) || !canonicalTime(o.IndexGeneratedAt) || o.ArtifactName == "" || !lowerHex(o.ArtifactSHA256, 64) || o.ArtifactSize <= 0 || !safeValidator(o.Validators.ETag) || !safeValidator(o.Validators.LastModified) {
		return ErrCorrupt
	}
	if o.Status == Withdrawn {
		return nilIf(o.ReleaseStatus != "withdrawn")
	}
	if o.ReleaseStatus != "active" {
		return ErrCorrupt
	}
	if o.Status == UpdateAvailable && (o.Relation != update.Newer || o.Compatibility != releasediscovery.CompatibilitySupported || o.MigrationID == "") {
		return ErrCorrupt
	}
	if o.Status == UpdateUnsupported && (o.Relation != update.Newer || o.Compatibility != releasediscovery.CompatibilityUnsupported) {
		return ErrCorrupt
	}
	if o.Status == Current && o.Relation == update.Newer {
		return ErrCorrupt
	}
	return nil
}

func nilIf(condition bool) error {
	if condition {
		return ErrCorrupt
	}
	return nil
}
func validStatus(s Status) bool {
	return s == Current || s == UpdateAvailable || s == UpdateUnsupported || s == Withdrawn
}
func validVersion(v string) bool { _, err := update.ParseVersion(v); return err == nil }
func canonicalTime(v string) bool {
	parsed, err := time.Parse(time.RFC3339, v)
	return err == nil && parsed.Location() == time.UTC && parsed.Nanosecond() == 0 && parsed.Format(time.RFC3339) == v
}
func lowerHex(v string, n int) bool {
	if len(v) != n {
		return false
	}
	for _, r := range v {
		if !(r >= '0' && r <= '9') && !(r >= 'a' && r <= 'f') {
			return false
		}
	}
	return true
}
func safeValidator(v string) bool {
	if len(v) > 256 {
		return false
	}
	for _, r := range v {
		if r < 0x20 || r > 0x7e {
			return false
		}
	}
	return true
}
func safeToken(v string, n int) bool {
	return v != "" && len(v) <= n && !strings.ContainsAny(v, "\x00\r\n")
}
func utc(v time.Time) bool { return !v.IsZero() && v.Location() == time.UTC }
func cloneObservation(v *Observation) *Observation {
	if v == nil {
		return nil
	}
	copy := *v
	return &copy
}
func contains(values []string, want string) bool {
	for _, value := range values {
		if value == want {
			return true
		}
	}
	return false
}
