package update

import (
	"os"
	"path/filepath"
	"testing"
)

func TestVerifyExternalReleasePackage(t *testing.T) {
	archive := os.Getenv("QWSG_TEST_ARCHIVE")
	if archive == "" {
		t.Skip("external package not selected")
	}
	version := os.Getenv("QWSG_TEST_VERSION")
	parent := t.TempDir()
	if err := os.Chmod(parent, 0700); err != nil {
		t.Fatal(err)
	}
	staged, err := StageLocal(archive, archive+".sha256", version, parent)
	if err != nil {
		t.Fatal(err)
	}
	pkg, err := VerifyPackage(staged)
	if err != nil {
		t.Fatal(err)
	}
	if pkg.Provenance.Version != version || pkg.Root == "" || len(pkg.Files) < 10 {
		t.Fatalf("invalid verified package: %+v", pkg)
	}
	if _, err = os.Stat(filepath.Join(pkg.Root, "bin/qwsg")); err != nil {
		t.Fatal(err)
	}
}
