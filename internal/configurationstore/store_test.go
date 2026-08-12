package configurationstore

import (
	"errors"
	"os"
	"path/filepath"
	"testing"

	"quantumwizard.hu/qwsg/internal/configuration"
)

func TestDefaultPathAndAtomicRoundTrip(t *testing.T) {
	home := filepath.Join(t.TempDir(), "home")
	if err := os.Mkdir(home, 0700); err != nil {
		t.Fatal(err)
	}
	path, err := DefaultPath(func(key string) string {
		if key == "HOME" {
			return home
		}
		return ""
	})
	if err != nil || path != filepath.Join(home, ".config", "qwsg", "config.json") {
		t.Fatalf("path=%q err=%v", path, err)
	}
	locale := "en"
	source, err := configuration.NormalizeSource(configuration.Source{SchemaName: configuration.SourceSchema, SchemaVersion: configuration.SchemaVersion, ModelVersion: configuration.ModelVersion, ID: "local.operator", SourceVersion: "1.0", Kind: configuration.PrimaryLocal, Patch: configuration.Patch{Locale: &locale}})
	if err != nil {
		t.Fatal(err)
	}
	if err = Save(path, source); err != nil {
		t.Fatal(err)
	}
	loaded, found, err := Load(path)
	if err != nil || !found || loaded.Identity != source.Identity {
		t.Fatalf("loaded=%#v found=%v err=%v", loaded, found, err)
	}
	for target, mode := range map[string]os.FileMode{filepath.Dir(path): 0700, path: 0600} {
		info, statErr := os.Stat(target)
		if statErr != nil || info.Mode().Perm() != mode {
			t.Fatalf("%s mode=%o err=%v", target, info.Mode().Perm(), statErr)
		}
	}
	before, _ := os.ReadFile(path)
	if err = Save(path, source); err != nil {
		t.Fatal(err)
	}
	after, _ := os.ReadFile(path)
	if string(before) != string(after) {
		t.Fatal("identical save was not deterministic")
	}
}

func TestUnsafeAndInvalidFilesFailClosed(t *testing.T) {
	root := t.TempDir()
	if _, err := SelectPath("relative", os.Getenv); !errors.Is(err, ErrUnsafe) {
		t.Fatalf("relative: %v", err)
	}
	real := filepath.Join(root, "real")
	if err := os.Mkdir(real, 0700); err != nil {
		t.Fatal(err)
	}
	link := filepath.Join(root, "link")
	if err := os.Symlink(real, link); err != nil {
		t.Fatal(err)
	}
	if _, _, err := Load(filepath.Join(link, "config.json")); !errors.Is(err, ErrUnsafe) {
		t.Fatalf("symlink parent: %v", err)
	}
	localeForSave := "en"
	safeSource, _ := configuration.NormalizeSource(configuration.Source{SchemaName: configuration.SourceSchema, SchemaVersion: configuration.SchemaVersion, ModelVersion: configuration.ModelVersion, ID: "local.operator", SourceVersion: "1.0", Kind: configuration.PrimaryLocal, Patch: configuration.Patch{Locale: &localeForSave}})
	if err := Save(filepath.Join(link, "missing", "config.json"), safeSource); !errors.Is(err, ErrUnsafe) {
		t.Fatalf("symlink ancestor behind missing leaf: %v", err)
	}
	path := filepath.Join(real, "config.json")
	if err := os.WriteFile(path, []byte("{}"), 0644); err != nil {
		t.Fatal(err)
	}
	if _, _, err := Load(path); !errors.Is(err, ErrPermission) {
		t.Fatalf("permissive: %v", err)
	}
	if err := os.Chmod(path, 0600); err != nil {
		t.Fatal(err)
	}
	if _, _, err := Load(path); !errors.Is(err, ErrInvalid) {
		t.Fatalf("invalid: %v", err)
	}
}

func TestSaveRefusesPermissiveExistingDirectory(t *testing.T) {
	root := filepath.Join(t.TempDir(), "config")
	if err := os.Mkdir(root, 0755); err != nil {
		t.Fatal(err)
	}
	locale := "en"
	source, _ := configuration.NormalizeSource(configuration.Source{SchemaName: configuration.SourceSchema, SchemaVersion: configuration.SchemaVersion, ModelVersion: configuration.ModelVersion, ID: "local.operator", SourceVersion: "1.0", Kind: configuration.PrimaryLocal, Patch: configuration.Patch{Locale: &locale}})
	if err := Save(filepath.Join(root, "config.json"), source); !errors.Is(err, ErrPermission) {
		t.Fatalf("permissive directory accepted: %v", err)
	}
	info, _ := os.Stat(root)
	if info.Mode().Perm() != 0755 {
		t.Fatal("existing permissions were silently changed")
	}
}

func TestSaveRefusesUnsafeExistingTargetAndPreservesIt(t *testing.T) {
	root := filepath.Join(t.TempDir(), "config")
	if err := os.Mkdir(root, 0700); err != nil {
		t.Fatal(err)
	}
	target := filepath.Join(root, "target")
	if err := os.WriteFile(target, []byte("owner"), 0600); err != nil {
		t.Fatal(err)
	}
	path := filepath.Join(root, "config.json")
	if err := os.Symlink(target, path); err != nil {
		t.Fatal(err)
	}
	locale := "en"
	source, _ := configuration.NormalizeSource(configuration.Source{SchemaName: configuration.SourceSchema, SchemaVersion: configuration.SchemaVersion, ModelVersion: configuration.ModelVersion, ID: "local.operator", SourceVersion: "1.0", Kind: configuration.PrimaryLocal, Patch: configuration.Patch{Locale: &locale}})
	if err := Save(path, source); !errors.Is(err, ErrUnsafe) {
		t.Fatalf("symlink save: %v", err)
	}
	data, _ := os.ReadFile(target)
	if string(data) != "owner" {
		t.Fatal("symlink target changed")
	}
}

func TestAtomicFailuresBeforeRenamePreservePreviousFile(t *testing.T) {
	root := filepath.Join(t.TempDir(), "config")
	if err := os.Mkdir(root, 0700); err != nil {
		t.Fatal(err)
	}
	path := filepath.Join(root, "config.json")
	locale := "en"
	old, _ := configuration.NormalizeSource(configuration.Source{SchemaName: configuration.SourceSchema, SchemaVersion: configuration.SchemaVersion, ModelVersion: configuration.ModelVersion, ID: "local.operator", SourceVersion: "1.0", Kind: configuration.PrimaryLocal, Patch: configuration.Patch{Locale: &locale}})
	if err := Save(path, old); err != nil {
		t.Fatal(err)
	}
	before, _ := os.ReadFile(path)
	locale = "hu"
	updated, _ := configuration.NormalizeSource(configuration.Source{SchemaName: configuration.SourceSchema, SchemaVersion: configuration.SchemaVersion, ModelVersion: configuration.ModelVersion, ID: "local.operator", SourceVersion: "1.0", Kind: configuration.PrimaryLocal, Patch: configuration.Patch{Locale: &locale}})
	for _, stage := range []string{"before_write", "before_file_sync", "before_rename"} {
		err := saveWithHook(path, updated, func(got string) error {
			if got == stage {
				return errors.New("injected")
			}
			return nil
		})
		if err == nil {
			t.Fatalf("%s succeeded", stage)
		}
		after, _ := os.ReadFile(path)
		if string(after) != string(before) {
			t.Fatalf("%s replaced prior file", stage)
		}
		matches, _ := filepath.Glob(filepath.Join(root, ".config-*"))
		if len(matches) != 0 {
			t.Fatalf("%s left temporary files: %v", stage, matches)
		}
	}
}

func TestDirectorySyncFailureLeavesOnlyCompleteValidReplacement(t *testing.T) {
	root := filepath.Join(t.TempDir(), "config")
	if err := os.Mkdir(root, 0700); err != nil {
		t.Fatal(err)
	}
	path := filepath.Join(root, "config.json")
	locale := "en"
	source, _ := configuration.NormalizeSource(configuration.Source{SchemaName: configuration.SourceSchema, SchemaVersion: configuration.SchemaVersion, ModelVersion: configuration.ModelVersion, ID: "local.operator", SourceVersion: "1.0", Kind: configuration.PrimaryLocal, Patch: configuration.Patch{Locale: &locale}})
	err := saveWithHook(path, source, func(stage string) error {
		if stage == "before_directory_sync" {
			return errors.New("injected")
		}
		return nil
	})
	if err == nil {
		t.Fatal("directory sync failure was hidden")
	}
	loaded, found, loadErr := Load(path)
	if loadErr != nil || !found || loaded.Identity != source.Identity {
		t.Fatalf("partial replacement: found=%v err=%v", found, loadErr)
	}
}
