// Package configurationstore activates the Canonical Configuration Contract
// at the per-user filesystem boundary. It owns path and persistence safety, not
// configuration semantics.
package configurationstore

import (
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"syscall"

	"quantumwizard.hu/qwsg/internal/configuration"
)

const MaxDocumentSize = 4 << 20

var (
	ErrUnavailable = errors.New("configuration location unavailable")
	ErrUnsafe      = errors.New("configuration path unsafe")
	ErrPermission  = errors.New("configuration permissions unsafe")
	ErrInvalid     = errors.New("configuration invalid")
)

// DefaultPath resolves the canonical per-user configuration location.
func DefaultPath(getenv func(string) string) (string, error) {
	if base := getenv("XDG_CONFIG_HOME"); base != "" {
		if cleanAbsolute(base) {
			return filepath.Join(base, "qwsg", "config.json"), nil
		}
		return "", ErrUnsafe
	}
	home := getenv("HOME")
	if !cleanAbsolute(home) {
		return "", ErrUnavailable
	}
	return filepath.Join(home, ".config", "qwsg", "config.json"), nil
}

// SelectPath uses an explicit path instead of discovery when supplied.
func SelectPath(explicit string, getenv func(string) string) (string, error) {
	if explicit != "" {
		if !cleanAbsolute(explicit) {
			return "", ErrUnsafe
		}
		return explicit, nil
	}
	return DefaultPath(getenv)
}

// Load returns found=false only when the selected file does not exist.
func Load(path string) (source configuration.Source, found bool, err error) {
	if !cleanAbsolute(path) {
		return source, false, ErrUnsafe
	}
	if err = inspectParents(filepath.Dir(path), true); errors.Is(err, os.ErrNotExist) {
		return source, false, nil
	} else if err != nil {
		return source, false, err
	}
	info, err := os.Lstat(path)
	if errors.Is(err, os.ErrNotExist) {
		return source, false, nil
	}
	if err != nil || info.Mode()&os.ModeSymlink != 0 || !info.Mode().IsRegular() {
		return source, false, ErrUnsafe
	}
	if info.Mode().Perm() != 0600 || !ownedByCurrentUser(info) {
		return source, false, ErrPermission
	}
	file, err := os.Open(path)
	if err != nil {
		return source, false, ErrPermission
	}
	defer file.Close()
	document, err := io.ReadAll(io.LimitReader(file, MaxDocumentSize+1))
	if err != nil || len(document) == 0 || len(document) > MaxDocumentSize {
		return source, false, ErrInvalid
	}
	after, err := file.Stat()
	if err != nil || !os.SameFile(info, after) || after.Mode().Perm() != 0600 || !ownedByCurrentUser(after) {
		return source, false, ErrUnsafe
	}
	source, err = configuration.DecodeSource(document)
	if err != nil {
		return source, false, fmt.Errorf("%w: %v", ErrInvalid, err)
	}
	if source.Kind != configuration.PrimaryLocal {
		return source, false, fmt.Errorf("%w: primary file has non-primary source kind", ErrInvalid)
	}
	return source, true, nil
}

// Save atomically publishes one normalized Source Record.
func Save(path string, source configuration.Source) error {
	return saveWithHook(path, source, func(string) error { return nil })
}

func saveWithHook(path string, source configuration.Source, hook func(string) error) error {
	if !cleanAbsolute(path) {
		return ErrUnsafe
	}
	if source.Kind != configuration.PrimaryLocal {
		return fmt.Errorf("%w: primary file requires primary source kind", ErrInvalid)
	}
	document, err := configuration.MarshalSourceCanonical(source)
	if err != nil || len(document) > MaxDocumentSize {
		return fmt.Errorf("%w: source record", ErrInvalid)
	}
	directory := filepath.Dir(path)
	if err = ensureDirectory(directory); err != nil {
		return err
	}
	if _, _, loadErr := Load(path); loadErr != nil {
		return loadErr
	}
	if info, statErr := os.Lstat(path); statErr == nil {
		if info.Mode()&os.ModeSymlink != 0 || !info.Mode().IsRegular() {
			return ErrUnsafe
		}
		if info.Mode().Perm() != 0600 || !ownedByCurrentUser(info) {
			return ErrPermission
		}
	} else if !errors.Is(statErr, os.ErrNotExist) {
		return ErrUnsafe
	}
	temporary, err := os.CreateTemp(directory, ".config-*")
	if err != nil {
		return err
	}
	name, installed := temporary.Name(), false
	defer func() {
		if !installed {
			_ = os.Remove(name)
		}
	}()
	if err = temporary.Chmod(0600); err == nil {
		err = hook("before_write")
	}
	if err == nil {
		_, err = temporary.Write(append(document, '\n'))
	}
	if err == nil {
		err = hook("before_file_sync")
	}
	if err == nil {
		err = temporary.Sync()
	}
	if closeErr := temporary.Close(); err == nil {
		err = closeErr
	}
	if err == nil {
		err = inspectParents(directory, true)
	}
	if err == nil {
		if info, statErr := os.Lstat(path); statErr == nil && (info.Mode()&os.ModeSymlink != 0 || !info.Mode().IsRegular() || info.Mode().Perm() != 0600 || !ownedByCurrentUser(info)) {
			err = ErrUnsafe
		} else if statErr != nil && !errors.Is(statErr, os.ErrNotExist) {
			err = ErrUnsafe
		}
	}
	if err == nil {
		err = hook("before_rename")
	}
	if err == nil {
		err = os.Rename(name, path)
	}
	if err != nil {
		return err
	}
	installed = true
	dir, err := os.Open(directory)
	if err == nil {
		err = hook("before_directory_sync")
		if err == nil {
			err = dir.Sync()
		}
		_ = dir.Close()
	}
	return err
}

func ensureDirectory(path string) error {
	if !cleanAbsolute(path) {
		return ErrUnsafe
	}
	if info, err := os.Lstat(path); err == nil {
		if info.Mode()&os.ModeSymlink != 0 || !info.IsDir() {
			return ErrUnsafe
		}
		if info.Mode().Perm() != 0700 || !ownedByCurrentUser(info) {
			return ErrPermission
		}
		return inspectParents(path, true)
	} else if !errors.Is(err, os.ErrNotExist) {
		return ErrUnsafe
	}
	if err := inspectParents(filepath.Dir(path), false); err != nil && !errors.Is(err, os.ErrNotExist) {
		return err
	}
	if err := os.MkdirAll(path, 0700); err != nil {
		return err
	}
	return inspectParents(path, true)
}

func inspectParents(path string, requireLeafPrivate bool) error {
	leafMissing := false
	for current := path; ; current = filepath.Dir(current) {
		info, err := os.Lstat(current)
		if errors.Is(err, os.ErrNotExist) {
			if current == path {
				leafMissing = true
			}
		} else if err != nil || info.Mode()&os.ModeSymlink != 0 || !info.IsDir() {
			return ErrUnsafe
		} else if current == path && requireLeafPrivate && (info.Mode().Perm() != 0700 || !ownedByCurrentUser(info)) {
			return ErrPermission
		}
		parent := filepath.Dir(current)
		if parent == current {
			break
		}
	}
	if leafMissing {
		return os.ErrNotExist
	}
	return nil
}

func ownedByCurrentUser(info os.FileInfo) bool {
	stat, ok := info.Sys().(*syscall.Stat_t)
	return ok && int(stat.Uid) == os.Geteuid()
}

func cleanAbsolute(path string) bool {
	return path != "" && filepath.IsAbs(path) && filepath.Clean(path) == path
}
