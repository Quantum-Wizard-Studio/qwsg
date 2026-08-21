package userruntime

import (
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"testing"
)

func TestResolveRejectsInvalidUID(t *testing.T) {
	if context, outcome := Resolve(-1); outcome != Unsafe || context.Environment() != "" {
		t.Fatalf("context=%q outcome=%q", context.Environment(), outcome)
	}
}

func TestValidateRuntimeDirectory(t *testing.T) {
	root := t.TempDir()
	uid := os.Geteuid()
	missing := filepath.Join(root, "missing")
	if _, outcome := validate(missing, uid); outcome != Missing {
		t.Fatalf("missing outcome=%q", outcome)
	}
	directory := filepath.Join(root, "runtime")
	if err := os.Mkdir(directory, 0700); err != nil {
		t.Fatal(err)
	}
	context, outcome := validate(directory, uid)
	if outcome != Valid || context.Environment() != "XDG_RUNTIME_DIR="+directory {
		t.Fatalf("environment=%q outcome=%q", context.Environment(), outcome)
	}
	if err := os.Chmod(directory, 0750); err != nil {
		t.Fatal(err)
	}
	if _, outcome = validate(directory, uid); outcome != Unsafe {
		t.Fatalf("unsafe mode outcome=%q", outcome)
	}
	if err := os.Chmod(directory, 0700); err != nil {
		t.Fatal(err)
	}
	if _, outcome = validate(directory, uid+1); outcome != Unsafe {
		t.Fatalf("wrong owner outcome=%q", outcome)
	}
	link := filepath.Join(root, "link")
	if err := os.Symlink(directory, link); err != nil {
		t.Fatal(err)
	}
	if _, outcome = validate(link, uid); outcome != Unsafe {
		t.Fatalf("symlink outcome=%q", outcome)
	}
	file := filepath.Join(root, "file")
	if err := os.WriteFile(file, []byte("not a directory"), 0600); err != nil {
		t.Fatal(err)
	}
	if _, outcome = validate(file, uid); outcome != Unsafe {
		t.Fatalf("special type outcome=%q", outcome)
	}
}

func TestResolveUsesCanonicalDecimalEffectiveUID(t *testing.T) {
	context, outcome := Resolve(os.Geteuid())
	if outcome == Valid {
		want := "XDG_RUNTIME_DIR=/run/user/" + strconv.Itoa(os.Geteuid())
		if context.Environment() != want || strings.Contains(context.Environment(), "//") {
			t.Fatalf("environment=%q want=%q", context.Environment(), want)
		}
	}
}
