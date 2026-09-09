// Package updateauthority binds read-only release discovery to the existing
// update transaction without granting metadata executable authority.
package updateauthority

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"time"

	"quantumwizard.hu/qwsg/internal/releasediscovery"
	"quantumwizard.hu/qwsg/internal/update"
)

// ErrNoUpdate reports a truthful authenticated equal/older relation.
var ErrNoUpdate = errors.New("authenticated release is not newer")

// Candidate carries authenticated authority in private fields. Awareness and
// caller-created evaluation values cannot authorize a package transaction.
type Candidate struct {
	evaluation  releasediscovery.Evaluation
	payload     []byte
	generatedAt time.Time
}

func (c Candidate) Evaluation() releasediscovery.Evaluation { return c.evaluation }
func (c Candidate) Metadata() []byte                        { return append([]byte(nil), c.payload...) }

func Authorize(payload []byte, verifier releasediscovery.Verifier, evaluator releasediscovery.Evaluator, channel, platform, target string, now time.Time) (Candidate, error) {
	index, err := releasediscovery.Parse(payload)
	if err != nil {
		return Candidate{}, err
	}
	auth, err := verifier.Verify(index)
	if err != nil {
		return Candidate{}, err
	}
	generated, _ := time.Parse(time.RFC3339, index.GeneratedAt)
	if generated.After(now.Add(15 * time.Minute)) {
		return Candidate{}, &releasediscovery.ContractError{Category: releasediscovery.MetadataFreshness}
	}
	evaluation, err := evaluator.Evaluate(auth, channel, platform, false)
	if err != nil {
		return Candidate{}, err
	}
	if target != "" && target != evaluation.Release.Version {
		return Candidate{}, fmt.Errorf("target authority mismatch")
	}
	if evaluation.Relation != update.Newer {
		return Candidate{evaluation: evaluation, payload: append([]byte(nil), payload...), generatedAt: generated}, ErrNoUpdate
	}
	if evaluation.Compatibility != releasediscovery.CompatibilitySupported {
		return Candidate{}, fmt.Errorf("authenticated migration capability unavailable")
	}
	return Candidate{evaluation: evaluation, payload: append([]byte(nil), payload...), generatedAt: generated}, nil
}

// CheckWatermark allows an awareness observation to restrict authority, never
// create it. Its timestamp must come from the validated local awareness store.
func (c Candidate) CheckWatermark(notBefore time.Time) error {
	if len(c.payload) == 0 || c.generatedAt.Before(notBefore) {
		return &releasediscovery.ContractError{Category: releasediscovery.MetadataRollback}
	}
	return nil
}

// VerifyStaged binds the signed archive identity before extracting any package.
func (c Candidate) VerifyStaged(staged update.Staged) (update.Package, error) {
	e := c.evaluation
	info, err := os.Stat(staged.Archive)
	if len(c.payload) == 0 || e.Relation != update.Newer || e.Compatibility != releasediscovery.CompatibilitySupported || err != nil || staged.Release.Version != e.Release.Version || filepath.Base(staged.Archive) != e.Artifact.Name || staged.SHA256 != e.Artifact.SHA256 || info.Size() != e.Artifact.Size {
		return update.Package{}, fmt.Errorf("authenticated artifact identity mismatch")
	}
	f, err := os.Open(staged.Archive)
	if err != nil {
		return update.Package{}, err
	}
	h := sha256.New()
	_, err = io.Copy(h, io.LimitReader(f, e.Artifact.Size+1))
	closeErr := f.Close()
	if err != nil || closeErr != nil || hex.EncodeToString(h.Sum(nil)) != e.Artifact.SHA256 {
		return update.Package{}, fmt.Errorf("authenticated artifact digest mismatch")
	}
	pkg, err := update.VerifyPackage(staged)
	if err != nil {
		return update.Package{}, err
	}
	if pkg.Provenance.Version != e.Release.Version || pkg.Provenance.Commit != e.Release.SourceCommit || pkg.Provenance.Platform != e.Platform {
		return update.Package{}, fmt.Errorf("authenticated release provenance mismatch")
	}
	return pkg, nil
}

func ReadMetadata(name string) ([]byte, error) {
	info, err := os.Lstat(name)
	if err != nil || !info.Mode().IsRegular() || info.Size() > releasediscovery.MaxIndexBytes {
		return nil, &releasediscovery.ContractError{Category: releasediscovery.MalformedMetadata}
	}
	f, err := os.Open(name)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	data, err := io.ReadAll(io.LimitReader(f, releasediscovery.MaxIndexBytes+1))
	if err != nil || len(data) > releasediscovery.MaxIndexBytes {
		return nil, &releasediscovery.ContractError{Category: releasediscovery.MalformedMetadata}
	}
	return data, nil
}

func FetchProductionMetadata(ctx context.Context) ([]byte, error) {
	source, err := releasediscovery.NewStaticHTTPSource(releasediscovery.ProductionEndpoint, releasediscovery.ProductionSourceID, nil)
	if err != nil {
		return nil, err
	}
	fetched, err := source.Fetch(ctx, releasediscovery.FetchRequest{Channel: "stable"})
	if err != nil {
		return nil, err
	}
	return fetched.Manifest, nil
}
