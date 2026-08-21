// Package userruntime derives and validates the trusted local runtime context
// used for fixed communication with the effective user's systemd manager.
package userruntime

import (
	"os"
	"path/filepath"
	"strconv"
	"syscall"
)

type Outcome string

const (
	Valid   Outcome = "valid"
	Missing Outcome = "missing"
	Unsafe  Outcome = "unsafe"
)

// Context is created only after validating the canonical effective-UID path.
// Its value cannot be constructed or changed by callers.
type Context struct{ directory string }

func Current() (Context, Outcome) { return Resolve(os.Geteuid()) }

// Resolve deterministically derives /run/user/<uid> and validates local state.
// Callers cannot supply a path or environment value.
func Resolve(uid int) (Context, Outcome) {
	if uid < 0 {
		return Context{}, Unsafe
	}
	directory := filepath.Join("/run/user", strconv.Itoa(uid))
	return validate(directory, uid)
}

// Environment returns the sole trusted environment entry accepted by the
// bounded runner. It is empty for an invalid zero Context.
func (context Context) Environment() string {
	if context.directory == "" {
		return ""
	}
	return "XDG_RUNTIME_DIR=" + context.directory
}

func validate(directory string, uid int) (Context, Outcome) {
	info, err := os.Lstat(directory)
	if os.IsNotExist(err) {
		return Context{}, Missing
	}
	if err != nil || info.Mode()&os.ModeSymlink != 0 || !info.IsDir() || info.Mode().Perm()&0077 != 0 {
		return Context{}, Unsafe
	}
	stat, ok := info.Sys().(*syscall.Stat_t)
	if !ok || int(stat.Uid) != uid {
		return Context{}, Unsafe
	}
	return Context{directory: directory}, Valid
}
