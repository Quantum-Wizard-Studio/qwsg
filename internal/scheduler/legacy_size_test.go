package scheduler

import (
	"context"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"

	"quantumwizard.hu/qwsg/internal/command"
	"quantumwizard.hu/qwsg/internal/configuration"
)

func TestLegacyOversizedStateFailsClosedAcrossRepeatedCycles(t *testing.T) {
	directory := t.TempDir()
	path := filepath.Join(directory, stateFileName)
	file, err := os.OpenFile(path, os.O_CREATE|os.O_EXCL|os.O_WRONLY, 0600)
	if err != nil {
		t.Fatal(err)
	}
	// A sparse 32 MiB legacy envelope must be rejected by stat, before decoding.
	if err = file.Truncate(32 << 20); err != nil {
		t.Fatal(err)
	}
	if err = file.Close(); err != nil {
		t.Fatal(err)
	}
	store, _ := OpenFileStore(directory)
	locker, _ := NewFileLocker(directory)
	executor := &fakePipeline{}
	cycle := Cycle{Configuration: effective(t, 1, intervalSchedule("schedule.legacy", 1, configuration.MisfireRunOnce)), Selection: command.Selection{Source: "live"}, LockOwnerID: "legacy.owner", Store: store, Locker: locker, Clock: &sequenceClock{values: []ClockObservation{{WallTime: time.Date(2026, 9, 6, 0, 0, 0, 0, time.UTC), SessionID: "session.legacy"}}}, TimeZones: SystemTimeZones{}, ResolveCommand: ResolveCanonicalCommand, Pipeline: executor}
	for i := 0; i < 50; i++ {
		if _, err := cycle.Run(context.Background()); err == nil || !strings.Contains(err.Error(), "size limit") {
			t.Fatalf("cycle %d: %v", i, err)
		}
	}
	if executor.calls != 0 {
		t.Fatal("oversized state executed work")
	}
	info, err := os.Stat(path)
	if err != nil || info.Size() != 32<<20 {
		t.Fatal("legacy state changed")
	}
	lock, err := locker.Acquire("legacy.verify")
	if err != nil {
		t.Fatal("failed cycle retained lock")
	}
	if err = lock.Release(); err != nil {
		t.Fatal(err)
	}
}
