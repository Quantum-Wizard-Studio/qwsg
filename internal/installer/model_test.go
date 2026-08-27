package installer

import (
	"strings"
	"testing"
)

func TestProgressCountsOnlyCompletedWork(t *testing.T) {
	p := NewProgress()
	if p.Percent() != 0 {
		t.Fatal(p.Percent())
	}
	_ = p.Start(PhasePreflight)
	if p.Percent() != 0 {
		t.Fatal("active work counted complete")
	}
	p.Fail(PhasePreflight, false)
	if p.Percent() != 0 {
		t.Fatal("failed work counted complete")
	}
	_ = p.Start(PhasePreflight)
	_ = p.Complete(PhasePreflight)
	if p.Percent() != 12 {
		t.Fatal(p.Percent())
	}
	for _, phase := range Phases[1:] {
		_ = p.Start(phase.ID)
		_ = p.Complete(phase.ID)
	}
	if p.Percent() != 100 {
		t.Fatal(p.Percent())
	}
}

func TestCatalogsAndFallback(t *testing.T) {
	if err := ValidateCatalogs(); err != nil {
		t.Fatal(err)
	}
	if (Catalog{Language: Language("future")}).Text("preflight") != english["preflight"] {
		t.Fatal("fallback failed")
	}
}

func TestPlatformContract(t *testing.T) {
	p, err := Detect(strings.NewReader("ID=ubuntu\nVERSION_ID=\"24.04\"\n"), "x86_64")
	if err != nil || !p.Supported() {
		t.Fatalf("%+v %v", p, err)
	}
	p.Version = "24.10"
	if p.Supported() {
		t.Fatal("unsupported version accepted")
	}
}
