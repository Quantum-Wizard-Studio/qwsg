package main

import (
	"bytes"
	"os"
	"path/filepath"
	"strings"
	"testing"

	updatecore "quantumwizard.hu/qwsg/internal/update"
)

func TestUpdateHelp(t *testing.T) {
	var out, errout bytes.Buffer
	if code := run([]string{"help", "update"}, &out, &errout); code != 0 || !strings.Contains(out.String(), "qwsg update rollback") {
		t.Fatalf("code=%d out=%q err=%q", code, out.String(), errout.String())
	}
}

func TestBootstrapUsesInstalledIdentity(t *testing.T) {
	root := t.TempDir()
	binary := filepath.Join(root, "qwsg")
	if err := os.WriteFile(binary, []byte("#!/bin/sh\nprintf 'QWSG 1.1.0\\ncommit: old\\nbuilt: old\\n'\n"), 0700); err != nil {
		t.Fatal(err)
	}
	previous := installedQWSGBinary
	installedQWSGBinary = binary
	defer func() { installedQWSGBinary = previous }()
	got, err := installedVersion()
	if err != nil || got != "1.1.0" {
		t.Fatalf("got %q %v", got, err)
	}
	if updatecore.Classify(got, "1.2.0-rc.1") != updatecore.Newer {
		t.Fatal("bootstrap target not newer than installed identity")
	}
}
func TestUpdateLocalArguments(t *testing.T) {
	archive, target, err := parseUpdateArgs([]string{"--archive", "/tmp/qwsg-1.2.0-rc.1-linux-amd64.tar.gz", "--version", "1.2.0-rc.1"})
	if err != nil || archive == "" || target != "1.2.0-rc.1" {
		t.Fatalf("%q %q %v", archive, target, err)
	}
	if _, _, err = parseUpdateArgs([]string{"--archive", "x"}); err == nil {
		t.Fatal("incomplete local identity accepted")
	}
}
func TestPrivilegedBackupPathBoundary(t *testing.T) {
	for _, path := range []string{"/tmp/x", "/var/lib/qwsg/rollback/../x", "relative"} {
		if validBackup(path) {
			t.Fatalf("accepted %q", path)
		}
	}
	if !validBackup("/var/lib/qwsg/rollback/1000/transaction") {
		t.Fatal("valid backup refused")
	}
}
