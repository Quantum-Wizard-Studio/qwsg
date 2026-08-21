package runner

import (
	"context"
	"errors"
	"strings"
	"testing"
	"time"
)

func TestRejectsUnlistedCommand(t *testing.T) {
	b := Bounded{Allowed: map[string]string{}, Timeout: time.Second, MaxOutput: 10}
	if _, e := b.Run(context.Background(), "sh", "-c", "id"); e == nil {
		t.Fatal("unlisted command accepted")
	}
}

func TestTrustedEnvironmentIsNarrow(t *testing.T) {
	t.Setenv("HOME", "/tmp/ambient-home")
	t.Setenv("DBUS_SESSION_BUS_ADDRESS", "unix:path=/tmp/ambient-bus")
	b := Bounded{Allowed: map[string]string{"env": "/usr/bin/env"}, TrustedEnvironment: map[string][]string{"env": {"XDG_RUNTIME_DIR=/run/user/1000"}}, Timeout: time.Second, MaxOutput: 4096}
	result, err := b.Run(context.Background(), "env")
	if err != nil || !strings.Contains(string(result.Stdout), "XDG_RUNTIME_DIR=/run/user/1000") {
		t.Fatalf("result=%q err=%v", result.Stdout, err)
	}
	if strings.Contains(string(result.Stdout), "ambient-home") || strings.Contains(string(result.Stdout), "ambient-bus") {
		t.Fatalf("ambient environment inherited: %q", result.Stdout)
	}
	b.TrustedEnvironment["env"] = []string{"HOME=/tmp/hostile"}
	if _, err := b.Run(context.Background(), "env"); err == nil {
		t.Fatal("arbitrary environment accepted")
	}
	b.TrustedEnvironment["env"] = []string{"XDG_RUNTIME_DIR=/run/user/01000"}
	if _, err := b.Run(context.Background(), "env"); err == nil {
		t.Fatal("non-canonical uid environment accepted")
	}
	for _, entries := range [][]string{
		{"XDG_RUNTIME_DIR=/run/user/1000", "XDG_RUNTIME_DIR=/run/user/1000"},
		{"DBUS_SESSION_BUS_ADDRESS=unix:path=/run/user/1000/bus"},
		{"HOME=/tmp/hostile"},
	} {
		b.TrustedEnvironment["env"] = entries
		if _, err := b.Run(context.Background(), "env"); err == nil {
			t.Fatalf("unsafe environment accepted: %q", entries)
		}
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
