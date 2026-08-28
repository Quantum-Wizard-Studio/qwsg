package main

import (
	"context"
	"errors"
	"time"

	"quantumwizard.hu/qwsg/internal/operatorstate"
	"quantumwizard.hu/qwsg/internal/presentationmodel"
)

var errGuardianEvidenceTimeout = errors.New("fresh Guardian evidence timeout")

type guardianEvidenceProbe func() (string, bool)
type guardianEvidencePause func(context.Context, time.Duration) error

func currentGuardianEvidence() (string, bool) {
	root, err := localStateRoot()
	if err != nil {
		return "", false
	}
	store, err := operatorstate.Open(root)
	if err != nil {
		return "", false
	}
	state, err := store.Load()
	if err != nil || state.Overview.Guardian != presentationmodel.GuardianRunning || state.Overview.Freshness != presentationmodel.FreshnessCurrent {
		return "", false
	}
	return state.ID, true
}

func waitForFreshGuardianEvidence(ctx context.Context, previousID string, limit time.Duration, probe guardianEvidenceProbe, pause guardianEvidencePause) error {
	if limit <= 0 || probe == nil || pause == nil {
		return errGuardianEvidenceTimeout
	}
	waitCtx, cancel := context.WithTimeout(ctx, limit)
	defer cancel()
	for {
		if id, ready := probe(); ready && id != "" && id != previousID {
			return nil
		}
		if err := pause(waitCtx, time.Second); err != nil {
			if errors.Is(waitCtx.Err(), context.DeadlineExceeded) {
				return errGuardianEvidenceTimeout
			}
			if waitCtx.Err() != nil {
				return waitCtx.Err()
			}
			return err
		}
	}
}

func guardianEvidenceSleep(ctx context.Context, duration time.Duration) error {
	timer := time.NewTimer(duration)
	defer timer.Stop()
	select {
	case <-ctx.Done():
		return ctx.Err()
	case <-timer.C:
		return nil
	}
}
