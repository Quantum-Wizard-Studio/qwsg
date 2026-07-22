package runner

import (
	"context"
	"errors"
	"testing"
	"time"
)

func TestRejectsUnlistedCommand(t *testing.T) {
	b := Bounded{Allowed: map[string]string{}, Timeout: time.Second, MaxOutput: 10}
	if _, e := b.Run(context.Background(), "sh", "-c", "id"); e == nil {
		t.Fatal("unlisted command accepted")
	}
}
func TestOutputLimit(t *testing.T) {
	b := Bounded{Allowed: map[string]string{"printf": "/usr/bin/printf"}, Timeout: time.Second, MaxOutput: 2}
	_, e := b.Run(context.Background(), "printf", "abcdef")
	if !errors.Is(e, ErrOutputLimit) {
		t.Fatalf("got %v", e)
	}
}
func TestTimeout(t *testing.T) {
	b := Bounded{Allowed: map[string]string{"sleep": "/usr/bin/sleep"}, Timeout: time.Millisecond, MaxOutput: 10}
	_, e := b.Run(context.Background(), "sleep", "1")
	if !errors.Is(e, context.DeadlineExceeded) {
		t.Fatalf("got %v", e)
	}
}
