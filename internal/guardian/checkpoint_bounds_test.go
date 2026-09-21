package guardian

import (
	"errors"
	"os"
	"path/filepath"
	"testing"
)

func TestCheckpointBoundAndIncompleteEvidence(t *testing.T) {
	s, err := OpenStore(filepath.Join(t.TempDir(), "guardian"))
	if err != nil {
		t.Fatal(err)
	}
	path := filepath.Join(s.directory, "checkpoint.json")
	if err := s.Save(checkpoint("valid")); err != nil {
		t.Fatal(err)
	}
	if _, err := s.Load(); err != nil {
		t.Fatal(err)
	}
	for _, data := range [][]byte{[]byte(`{"schema_name":`), nil} {
		if err := os.WriteFile(path, data, 0600); err != nil {
			t.Fatal(err)
		}
		if _, err := s.Load(); !errors.Is(err, ErrCheckpoint) {
			t.Fatalf("incomplete accepted: %v", err)
		}
	}
	f, err := os.OpenFile(path, os.O_WRONLY|os.O_TRUNC, 0600)
	if err != nil {
		t.Fatal(err)
	}
	if err := f.Truncate(64 * MaxSize); err != nil {
		t.Fatal(err)
	}
	f.Close()
	if _, err := s.Load(); !errors.Is(err, ErrCheckpoint) {
		t.Fatalf("oversize accepted: %v", err)
	}
	if err := s.Save(checkpoint("replacement")); err != nil {
		t.Fatal(err)
	}
	if err := os.Chmod(path, 0644); err != nil {
		t.Fatal(err)
	}
	if _, err := s.Load(); !errors.Is(err, ErrUnsafePath) {
		t.Fatalf("permissions: %v", err)
	}
}
