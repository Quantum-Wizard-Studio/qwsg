package updateawareness

import (
	"context"
	"errors"
	"time"

	"quantumwizard.hu/qwsg/internal/installation"
	"quantumwizard.hu/qwsg/internal/releasediscovery"
)

var ErrInstalledIdentity = errors.New("installed identity is not verified")

type Checker interface {
	Check(context.Context, releasediscovery.FetchRequest, string, bool) (releasediscovery.CheckResult, error)
}

type Classifier func() installation.Result

type Manager struct {
	Store           *Store
	Checker         Checker
	Classify        Classifier
	SourceID        string
	Channel         string
	Platform        string
	AllowPrerelease bool
	Freshness       time.Duration
	Now             func() time.Time
}

func (m Manager) Check(ctx context.Context) (State, error) {
	if m.Store == nil || m.Classify == nil || m.SourceID == "" || m.Channel == "" {
		return State{}, ErrCorrupt
	}
	lock, err := m.Store.Lock()
	if err != nil {
		return State{}, err
	}
	defer lock.Release()
	installedResult := m.Classify()
	if installedResult.State != installation.VerifiedSupported {
		return State{}, ErrInstalledIdentity
	}
	installed := Installed{Classification: installedResult.State, Version: installedResult.Version}
	previous, loadErr := m.Store.Load()
	if loadErr != nil && !errors.Is(loadErr, ErrMissing) {
		return State{}, loadErr
	}
	var prior *State
	validators := releasediscovery.Validators{}
	if loadErr == nil {
		prior = &previous
		if previous.SourceID == m.SourceID && previous.Channel == m.Channel && previous.LastSuccess != nil {
			validators = previous.LastSuccess.Validators
		}
	}
	now := time.Now().UTC()
	if m.Now != nil {
		now = m.Now().UTC()
	}
	if m.Checker == nil {
		state, stateErr := NewFailure(prior, m.SourceID, m.Channel, installed, now, string(releasediscovery.SourceAuthority))
		if stateErr != nil {
			return State{}, stateErr
		}
		if stateErr = m.Store.Publish(state); stateErr != nil {
			return State{}, stateErr
		}
		return state, &releasediscovery.ContractError{Category: releasediscovery.SourceAuthority}
	}
	result, checkErr := m.Checker.Check(ctx, releasediscovery.FetchRequest{Channel: m.Channel, Validators: validators}, m.Platform, m.AllowPrerelease)
	if checkErr != nil {
		if releasediscovery.FailureOf(checkErr) == releasediscovery.NoEligibleRelease && prior != nil && prior.LastSuccess != nil && contains(result.WithdrawnVersions, prior.LastSuccess.ReleaseVersion) {
			state, stateErr := NewWithdrawn(*prior, result, installed, now, m.Freshness)
			if stateErr != nil {
				return State{}, stateErr
			}
			if stateErr = m.Store.Publish(state); stateErr != nil {
				return State{}, stateErr
			}
			return state, nil
		}
		failure := string(releasediscovery.FailureOf(checkErr))
		state, stateErr := NewFailure(prior, m.SourceID, m.Channel, installed, now, failure)
		if stateErr != nil {
			return State{}, stateErr
		}
		if stateErr = m.Store.Publish(state); stateErr != nil {
			return State{}, stateErr
		}
		return state, checkErr
	}
	state, err := NewSuccess(prior, result, m.SourceID, m.Channel, installed, now, m.Freshness)
	if err != nil {
		return State{}, err
	}
	if err = m.Store.Publish(state); err != nil {
		return State{}, err
	}
	return state, nil
}
