package main

import (
	"os"
	"path/filepath"
	"testing"

	"quantumwizard.hu/qwsg/internal/releasepublication"
)

func TestGenerateWritesExactCanonicalSigningBytes(t *testing.T) {
	candidate := filepath.Join("..", "..", "internal", "releasepublication", "testdata", "unsigned-candidate.json")
	wantInput, err := os.ReadFile(candidate)
	if err != nil {
		t.Fatal(err)
	}
	want, err := releasepublication.Generate(wantInput)
	if err != nil {
		t.Fatal(err)
	}
	output := filepath.Join(t.TempDir(), "signing-input.json")
	if err = run([]string{"generate", candidate, output}); err != nil {
		t.Fatal(err)
	}
	got, err := os.ReadFile(output)
	if err != nil || string(got) != string(want) {
		t.Fatalf("CLI altered signing bytes: err=%v got=%q want=%q", err, got, want)
	}
}
