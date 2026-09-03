package guardian

import (
	"context"
	"errors"
	"time"

	"quantumwizard.hu/qwsg/internal/updateawareness"
)

const (
	DefaultReleaseCheckInterval = 24 * time.Hour
	DefaultReleaseCheckTimeout  = 35 * time.Second
)

type AwarenessStateStore interface {
	Load() (updateawareness.State, error)
}

// ReleaseCheckService owns the Guardian's isolated low-frequency awareness
// loop. Check failures are deliberately not returned: they are persisted by
// the awareness manager and cannot change Guardian health or lifecycle.
type ReleaseCheckService struct {
	Store    AwarenessStateStore
	Check    func(context.Context) error
	Ready    <-chan struct{}
	Interval time.Duration
	Timeout  time.Duration
	Now      func() time.Time
	Wait     func(context.Context, time.Time) error
}

func (s ReleaseCheckService) Run(ctx context.Context) {
	if s.Store == nil || s.Check == nil || s.Interval <= 0 || s.Timeout <= 0 || s.Timeout >= s.Interval {
		return
	}
	if s.Ready != nil {
		select {
		case <-ctx.Done():
			return
		case <-s.Ready:
		}
	}
	for ctx.Err() == nil {
		now := s.now()
		due := now
		state, err := s.Store.Load()
		if err == nil {
			due = state.LastAttempt.At.Add(s.Interval)
			// A corrupt clock observation must not suppress checks indefinitely.
			if due.After(now.Add(s.Interval)) {
				due = now.Add(s.Interval)
			}
		} else if !errors.Is(err, updateawareness.ErrMissing) {
			// Unsafe or corrupt local state fails closed without a network request.
			due = now.Add(s.Interval)
		}
		if due.After(now) {
			if s.wait(ctx, due) != nil {
				return
			}
			continue
		}
		checkCtx, cancel := context.WithTimeout(ctx, s.Timeout)
		_ = s.Check(checkCtx)
		cancel()
		// Manager.Check normally persists LastAttempt. This bounded wait also
		// prevents tight retries when classification or local storage fails first.
		if s.wait(ctx, s.now().Add(s.Interval)) != nil {
			return
		}
	}
}

func (s ReleaseCheckService) now() time.Time {
	if s.Now != nil {
		return s.Now().UTC()
	}
	return time.Now().UTC()
}

func (s ReleaseCheckService) wait(ctx context.Context, at time.Time) error {
	if s.Wait != nil {
		return s.Wait(ctx, at)
	}
	delay := time.Until(at)
	if delay <= 0 {
		return nil
	}
	timer := time.NewTimer(delay)
	defer timer.Stop()
	select {
	case <-ctx.Done():
		return ctx.Err()
	case <-timer.C:
		return nil
	}
}
