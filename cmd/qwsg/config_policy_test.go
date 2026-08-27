package main

import (
	"bytes"
	"path/filepath"
	"strings"
	"testing"
)

func TestUpdatePolicyConfigurationIsExplicitAndBounded(t *testing.T) {
	home := t.TempDir()
	t.Setenv("HOME", home)
	t.Setenv("XDG_CONFIG_HOME", filepath.Join(home, "config"))
	var out, errout bytes.Buffer
	if code := run([]string{"setup", "--accept-defaults"}, &out, &errout); code != 0 {
		t.Fatalf("setup: %d %s", code, errout.String())
	}
	out.Reset()
	if code := run([]string{"config", "set", "update.policy", "notify"}, &out, &errout); code != 0 {
		t.Fatalf("set policy: %d %s", code, errout.String())
	}
	out.Reset()
	if code := run([]string{"config", "get", "update.policy"}, &out, &errout); code != 0 || strings.TrimSpace(out.String()) != "notify" {
		t.Fatalf("get policy: %d %q %s", code, out.String(), errout.String())
	}
	if code := run([]string{"config", "set", "update.policy", "automatic"}, &out, &errout); code == 0 {
		t.Fatal("unsupported automatic policy accepted")
	}
}
