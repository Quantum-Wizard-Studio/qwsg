package credentialstore

import (
	"errors"
	"os"
	"path/filepath"
	"testing"
)

func TestSaveLoadAndPermissions(t *testing.T) {
	root := t.TempDir()
	if err := os.Chmod(root, 0700); err != nil {
		t.Fatal(err)
	}
	config := filepath.Join(root, "qwsg", "config.json")
	if err := Save(config, "smtp.admin", []byte("test-secret")); err != nil {
		t.Fatal(err)
	}
	b, err := Load(config, "smtp.admin")
	if err != nil || string(b) != "test-secret" {
		t.Fatalf("load: %q %v", b, err)
	}
	path, _ := Path(config, "smtp.admin")
	info, _ := os.Stat(path)
	if info.Mode().Perm() != 0600 {
		t.Fatalf("mode %o", info.Mode().Perm())
	}
}
func TestRejectsUnsafeSecretAndSymlink(t *testing.T) {
	root := t.TempDir()
	_ = os.Chmod(root, 0700)
	config := filepath.Join(root, "qwsg", "config.json")
	if !errors.Is(Save(config, "smtp.admin", []byte("bad\nsecret")), ErrUnsafe) {
		t.Fatal("newline accepted")
	}
	dir, _ := Directory(config)
	if err := os.MkdirAll(dir, 0700); err != nil {
		t.Fatal(err)
	}
	target := filepath.Join(root, "target")
	if err := os.WriteFile(target, []byte("secret"), 0600); err != nil {
		t.Fatal(err)
	}
	path, _ := Path(config, "smtp.admin")
	if err := os.Symlink(target, path); err != nil {
		t.Fatal(err)
	}
	if !errors.Is(func() error { _, e := Load(config, "smtp.admin"); return e }(), ErrUnsafe) {
		t.Fatal("symlink accepted")
	}
}
func TestAtomicUpdatePreservesReadableValue(t *testing.T) {
	root := t.TempDir()
	_ = os.Chmod(root, 0700)
	config := filepath.Join(root, "qwsg", "config.json")
	if err := Save(config, "smtp.admin", []byte("one")); err != nil {
		t.Fatal(err)
	}
	if err := Save(config, "smtp.admin", []byte("two")); err != nil {
		t.Fatal(err)
	}
	b, err := Load(config, "smtp.admin")
	if err != nil || string(b) != "two" {
		t.Fatalf("update %q %v", b, err)
	}
}
