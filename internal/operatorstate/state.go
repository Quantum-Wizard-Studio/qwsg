// Package operatorstate persists one validated current operator overview.
// It owns storage identity and integrity, never presentation semantics.
package operatorstate

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

	"quantumwizard.hu/qwsg/internal/presentationmodel"
)

const (
	SchemaName                 = "qwsg.current-operator-state"
	SchemaVersion              = "1.2"
	LegacySchemaVersion        = "1.0"
	LegacySchemaV11            = "1.1"
	ModelVersion               = "1.2"
	LegacyModelVersion         = "1.0"
	LegacyModelV11             = "1.1"
	DigestAlgorithm            = "sha256"
	CoverageInventorySnapshot  = "inventory_snapshot"
	CoverageOperatorEvaluation = "operator_evaluation"
	CoverageGuardianOperation  = "guardian_operation"
	PublicationCheck           = "check_completed"
	PublicationObserve         = "observe_completed"
	PublicationGuardian        = "guardian_observed"
	MaxEncodedSize             = 1 << 20
	MaxTokenLength             = 256
)

var (
	ErrMissing      = errors.New("current operator state missing")
	ErrCorrupt      = errors.New("current operator state corrupt")
	ErrIncompatible = errors.New("current operator state incompatible")
	ErrUnsafePath   = errors.New("current operator state path unsafe")
	ErrPermission   = errors.New("current operator state permission denied")
)

type Provenance struct {
	DefinitionID       string   `json:"definition_id"`
	ExecutionID        string   `json:"execution_id"`
	Profile            string   `json:"profile"`
	Source             string   `json:"source"`
	Stages             []string `json:"stages"`
	Reason             string   `json:"reason"`
	ApplicationVersion string   `json:"application_version"`
}

type State struct {
	SchemaName      string                     `json:"schema_name"`
	SchemaVersion   string                     `json:"schema_version"`
	ModelVersion    string                     `json:"model_version"`
	ID              string                     `json:"id"`
	DigestAlgorithm string                     `json:"digest_algorithm"`
	PayloadDigest   string                     `json:"payload_digest"`
	ObservedAt      time.Time                  `json:"observed_at"`
	PublishedAt     time.Time                  `json:"published_at"`
	FreshUntil      time.Time                  `json:"fresh_until"`
	Coverage        string                     `json:"coverage"`
	Provenance      Provenance                 `json:"provenance"`
	Overview        presentationmodel.Overview `json:"overview"`
}

func Normalize(value State) (State, error) {
	value.SchemaName, value.SchemaVersion, value.ModelVersion = SchemaName, SchemaVersion, ModelVersion
	value.DigestAlgorithm = DigestAlgorithm
	value.ID, value.PayloadDigest = "", ""
	if err := validateContent(value); err != nil {
		return State{}, err
	}
	payload, err := presentationmodel.MarshalCanonical(value.Overview)
	if err != nil {
		return State{}, err
	}
	sum := sha256.Sum256(payload)
	value.PayloadDigest = hex.EncodeToString(sum[:])
	identity := value
	identity.ID = ""
	document, _ := json.Marshal(identity)
	idSum := sha256.Sum256(document)
	value.ID = hex.EncodeToString(idSum[:])
	return value, Validate(value)
}

func Validate(value State) error {
	currentVersion := value.SchemaVersion == SchemaVersion && value.ModelVersion == ModelVersion
	legacyVersion := (value.SchemaVersion == LegacySchemaVersion && value.ModelVersion == LegacyModelVersion) || (value.SchemaVersion == LegacySchemaV11 && value.ModelVersion == LegacyModelV11)
	if value.SchemaName != SchemaName || (!currentVersion && !legacyVersion) || value.DigestAlgorithm != DigestAlgorithm {
		return ErrIncompatible
	}
	claimedID, claimedDigest := value.ID, value.PayloadDigest
	normalized, err := NormalizeWithoutClaims(value)
	if err != nil {
		return err
	}
	if claimedID != normalized.ID || claimedDigest != normalized.PayloadDigest {
		return ErrCorrupt
	}
	return nil
}

func NormalizeWithoutClaims(value State) (State, error) {
	value.ID, value.PayloadDigest = "", ""
	if err := validateContent(value); err != nil {
		return State{}, err
	}
	payload, err := presentationmodel.MarshalCanonical(value.Overview)
	if err != nil {
		return State{}, ErrCorrupt
	}
	sum := sha256.Sum256(payload)
	value.PayloadDigest = hex.EncodeToString(sum[:])
	identity := value
	identity.ID = ""
	document, _ := json.Marshal(identity)
	idSum := sha256.Sum256(document)
	value.ID = hex.EncodeToString(idSum[:])
	return value, nil
}

func validateContent(value State) error {
	currentVersion := value.SchemaVersion == SchemaVersion && value.ModelVersion == ModelVersion && value.Overview.SchemaVersion == presentationmodel.SchemaVersion
	legacyVersion := (value.SchemaVersion == LegacySchemaVersion && value.ModelVersion == LegacyModelVersion && value.Overview.SchemaVersion == presentationmodel.LegacySchemaVersion) || (value.SchemaVersion == LegacySchemaV11 && value.ModelVersion == LegacyModelV11 && value.Overview.SchemaVersion == presentationmodel.LegacySchemaV11)
	if value.SchemaName != SchemaName || (!currentVersion && !legacyVersion) || value.DigestAlgorithm != DigestAlgorithm {
		return ErrIncompatible
	}
	if err := presentationmodel.Validate(value.Overview); err != nil {
		return fmt.Errorf("%w: invalid overview", ErrCorrupt)
	}
	if !utc(value.ObservedAt) || !utc(value.PublishedAt) || !utc(value.FreshUntil) || value.PublishedAt.Before(value.ObservedAt) || !value.FreshUntil.After(value.ObservedAt) || value.Overview.ObservedAt != value.ObservedAt {
		return fmt.Errorf("%w: invalid times", ErrCorrupt)
	}
	if !validProvenance(value) {
		return fmt.Errorf("%w: invalid provenance", ErrCorrupt)
	}
	for _, token := range []string{value.Provenance.DefinitionID, value.Provenance.ExecutionID, value.Provenance.ApplicationVersion} {
		if !boundedToken(token) {
			return fmt.Errorf("%w: invalid provenance token", ErrCorrupt)
		}
	}
	return nil
}

func validProvenance(value State) bool {
	if value.Provenance.Source != "live" {
		return false
	}
	switch value.Coverage {
	case CoverageInventorySnapshot:
		return value.Provenance.Profile == "check" && value.Provenance.Reason == PublicationCheck && equalStages(value.Provenance.Stages, []string{"inventory", "snapshot"})
	case CoverageOperatorEvaluation:
		return value.Provenance.Profile == "observe" && value.Provenance.Reason == PublicationObserve && equalStages(value.Provenance.Stages, []string{"inventory", "snapshot", "compare", "drift", "health", "rule", "policy", "report"})
	case CoverageGuardianOperation:
		if value.Provenance.Profile != "guardian" || value.Provenance.Reason != PublicationGuardian {
			return false
		}
		return equalStages(value.Provenance.Stages, []string{"runtime_service"}) || equalStages(value.Provenance.Stages, []string{"inventory", "snapshot", "compare", "drift", "health", "rule", "policy", "report", "runtime", "runtime_service"})
	default:
		return false
	}
}

func equalStages(got, want []string) bool {
	if len(got) != len(want) {
		return false
	}
	for index := range want {
		if got[index] != want[index] {
			return false
		}
	}
	return true
}

func MarshalCanonical(value State) ([]byte, error) {
	if err := Validate(value); err != nil {
		return nil, err
	}
	document, err := json.Marshal(value)
	if err != nil {
		return nil, err
	}
	if len(document) > MaxEncodedSize {
		return nil, fmt.Errorf("%w: state too large", ErrCorrupt)
	}
	return append(document, '\n'), nil
}

func Decode(document []byte) (State, error) {
	if len(document) == 0 || len(document) > MaxEncodedSize {
		return State{}, fmt.Errorf("%w: invalid size", ErrCorrupt)
	}
	decoder := json.NewDecoder(bytes.NewReader(document))
	decoder.DisallowUnknownFields()
	var value State
	if err := decoder.Decode(&value); err != nil {
		return State{}, fmt.Errorf("%w: invalid document", ErrCorrupt)
	}
	var extra any
	if err := decoder.Decode(&extra); err != io.EOF {
		return State{}, fmt.Errorf("%w: trailing data", ErrCorrupt)
	}
	if err := Validate(value); err != nil {
		return State{}, err
	}
	return value, nil
}

func boundedToken(value string) bool {
	return value != "" && len(value) <= MaxTokenLength && !strings.ContainsAny(value, "\x00\r\n")
}
func utc(value time.Time) bool { return !value.IsZero() && value.Location() == time.UTC }
