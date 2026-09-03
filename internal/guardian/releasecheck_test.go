package guardian

import (
	"context"
	"errors"
	"sync"
	"sync/atomic"
	"testing"
	"time"

	"quantumwizard.hu/qwsg/internal/updateawareness"
)

type awarenessStateStore struct {
	state updateawareness.State
	err   error
}

func (s awarenessStateStore) Load() (updateawareness.State, error) { return s.state, s.err }

func TestReleaseCheckDueAndRestartNotDue(t *testing.T) {
	interval := 24 * time.Hour
	now := time.Date(2026, 9, 3, 12, 0, 0, 0, time.UTC)
	cases := []struct {
		name      string
		store     AwarenessStateStore
		wantCheck int32
		wantWait  time.Time
	}{
		{name: "never checked is due", store: awarenessStateStore{err: updateawareness.ErrMissing}, wantCheck: 1, wantWait: now.Add(interval)},
		{name: "recent restart is not due", store: awarenessStateStore{state: updateawareness.State{LastAttempt: updateawareness.Attempt{At: now.Add(-time.Hour)}}}, wantWait: now.Add(23 * time.Hour)},
		{name: "interval elapsed is due", store: awarenessStateStore{state: updateawareness.State{LastAttempt: updateawareness.Attempt{At: now.Add(-interval)}}}, wantCheck: 1, wantWait: now.Add(interval)},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			ctx, cancel := context.WithCancel(context.Background())
			defer cancel()
			var checks int32
			var waited time.Time
			service := ReleaseCheckService{Store: tc.store, Interval: interval, Timeout: time.Minute, Now: func() time.Time { return now }, Check: func(context.Context) error {
				atomic.AddInt32(&checks, 1)
				return nil
			}, Wait: func(_ context.Context, at time.Time) error {
				waited = at
				cancel()
				return context.Canceled
			}}
			service.Run(ctx)
			if checks != tc.wantCheck || !waited.Equal(tc.wantWait) {
				t.Fatalf("checks=%d wait=%s; want %d %s", checks, waited, tc.wantCheck, tc.wantWait)
			}
		})
	}
}

func TestReleaseCheckWaitsForGuardianStartup(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	ready := make(chan struct{})
	called := make(chan struct{}, 1)
	done := make(chan struct{})
	go func() {
		defer close(done)
		(ReleaseCheckService{Store: awarenessStateStore{err: updateawareness.ErrMissing}, Ready: ready, Interval: time.Hour, Timeout: time.Minute, Check: func(context.Context) error {
			called <- struct{}{}
			return nil
		}, Wait: func(context.Context, time.Time) error { return context.Canceled }}).Run(ctx)
	}()
	select {
	case <-called:
		t.Fatal("release check ran before the first local cycle")
	default:
	}
	close(ready)
	select {
	case <-called:
	case <-time.After(time.Second):
		t.Fatal("release check did not run after startup gate")
	}
	<-done
}

func TestReleaseCheckFailureIsolatedSequentialAndCancellationAware(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	var active, maximum, checks int32
	var nowMu sync.Mutex
	now := time.Date(2026, 9, 3, 12, 0, 0, 0, time.UTC)
	service := ReleaseCheckService{Store: awarenessStateStore{err: updateawareness.ErrMissing}, Interval: time.Hour, Timeout: time.Minute, Now: func() time.Time {
		nowMu.Lock()
		defer nowMu.Unlock()
		return now
	}, Check: func(context.Context) error {
		current := atomic.AddInt32(&active, 1)
		if current > atomic.LoadInt32(&maximum) {
			atomic.StoreInt32(&maximum, current)
		}
		atomic.AddInt32(&checks, 1)
		atomic.AddInt32(&active, -1)
		return errors.New("bounded source failure")
	}, Wait: func(_ context.Context, at time.Time) error {
		nowMu.Lock()
		now = at
		nowMu.Unlock()
		if atomic.LoadInt32(&checks) == 2 {
			cancel()
			return context.Canceled
		}
		return nil
	}}
	service.Run(ctx)
	if checks != 2 || maximum != 1 || active != 0 {
		t.Fatalf("checks=%d maximum=%d active=%d", checks, maximum, active)
	}
}

func TestReleaseCheckCancelsInFlightCheckAndCleansUp(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	done := make(chan struct{})
	go func() {
		defer close(done)
		(ReleaseCheckService{Store: awarenessStateStore{err: updateawareness.ErrMissing}, Interval: time.Hour, Timeout: time.Minute, Check: func(checkCtx context.Context) error {
			cancel()
			<-checkCtx.Done()
			return checkCtx.Err()
		}, Wait: func(ctx context.Context, _ time.Time) error { return ctx.Err() }}).Run(ctx)
	}()
	select {
	case <-done:
	case <-time.After(time.Second):
		t.Fatal("canceled release check leaked its lifecycle")
	}
}

func TestUnsafeAwarenessStateFailsClosedWithoutNetwork(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	checks := 0
	service := ReleaseCheckService{Store: awarenessStateStore{err: updateawareness.ErrCorrupt}, Interval: time.Hour, Timeout: time.Minute, Check: func(context.Context) error {
		checks++
		return nil
	}, Wait: func(_ context.Context, _ time.Time) error {
		cancel()
		return context.Canceled
	}}
	service.Run(ctx)
	if checks != 0 {
		t.Fatal("unsafe awareness state triggered network retrieval")
	}
}
