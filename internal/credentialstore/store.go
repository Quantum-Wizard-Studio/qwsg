// Package credentialstore provides the narrow private per-user secret boundary.
package credentialstore

import (
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"regexp"
	"syscall"
)

const MaxSecretSize = 64 << 10

var (
	ErrUnsafe      = errors.New("credential path unsafe")
	ErrPermission  = errors.New("credential permissions unsafe")
	ErrUnavailable = errors.New("credential unavailable")
	namePattern    = regexp.MustCompile(`^[a-z0-9]+([._-][a-z0-9]+)*$`)
)

func Directory(configPath string) (string, error) {
	if !cleanAbsolute(configPath) {
		return "", ErrUnsafe
	}
	return filepath.Join(filepath.Dir(configPath), "credentials"), nil
}

func Path(configPath, reference string) (string, error) {
	dir, err := Directory(configPath)
	if err != nil || !namePattern.MatchString(reference) || len(reference) > 128 {
		return "", ErrUnsafe
	}
	return filepath.Join(dir, reference), nil
}

func Load(configPath, reference string) ([]byte, error) {
	path, err := Path(configPath, reference)
	if err != nil {
		return nil, err
	}
	if err = inspectDir(filepath.Dir(path), false); err != nil {
		return nil, err
	}
	info, err := os.Lstat(path)
	if errors.Is(err, os.ErrNotExist) {
		return nil, ErrUnavailable
	}
	if err != nil || !info.Mode().IsRegular() || info.Mode()&os.ModeSymlink != 0 || info.Sys().(*syscall.Stat_t).Nlink != 1 {
		return nil, ErrUnsafe
	}
	if info.Mode().Perm() != 0600 || !owned(info) {
		return nil, ErrPermission
	}
	f, err := os.Open(path)
	if err != nil {
		return nil, ErrPermission
	}
	defer f.Close()
	b, err := io.ReadAll(io.LimitReader(f, MaxSecretSize+1))
	if err != nil || len(b) == 0 || len(b) > MaxSecretSize {
		return nil, ErrUnavailable
	}
	after, err := f.Stat()
	if err != nil || !os.SameFile(info, after) || after.Mode().Perm() != 0600 || !owned(after) {
		return nil, ErrUnsafe
	}
	for _, c := range b {
		if c == 0 || c == '\n' || c == '\r' {
			return nil, ErrUnsafe
		}
	}
	return b, nil
}

func Save(configPath, reference string, secret []byte) error {
	path, err := Path(configPath, reference)
	if err != nil {
		return err
	}
	if len(secret) == 0 || len(secret) > MaxSecretSize {
		return ErrUnsafe
	}
	for _, c := range secret {
		if c == 0 || c == '\n' || c == '\r' {
			return ErrUnsafe
		}
	}
	dir := filepath.Dir(path)
	if err = ensureDir(dir); err != nil {
		return err
	}
	if _, err = os.Lstat(path); err == nil {
		if _, loadErr := Load(configPath, reference); loadErr != nil {
			return loadErr
		}
	} else if !errors.Is(err, os.ErrNotExist) {
		return ErrUnsafe
	}
	f, err := os.CreateTemp(dir, ".credential-*")
	if err != nil {
		return err
	}
	name, installed := f.Name(), false
	defer func() {
		if !installed {
			_ = os.Remove(name)
		}
	}()
	if err = f.Chmod(0600); err == nil {
		_, err = f.Write(secret)
	}
	if err == nil {
		err = f.Sync()
	}
	if closeErr := f.Close(); err == nil {
		err = closeErr
	}
	if err == nil {
		err = os.Rename(name, path)
	}
	if err != nil {
		return fmt.Errorf("credential persistence failed")
	}
	installed = true
	d, err := os.Open(dir)
	if err == nil {
		err = d.Sync()
		_ = d.Close()
	}
	return err
}

func ensureDir(path string) error {
	if !cleanAbsolute(path) {
		return ErrUnsafe
	}
	if err := os.MkdirAll(path, 0700); err != nil {
		return err
	}
	if err := os.Chmod(path, 0700); err != nil {
		return err
	}
	return inspectDir(path, false)
}
func inspectDir(path string, allowMissing bool) error {
	for current := path; ; current = filepath.Dir(current) {
		info, err := os.Lstat(current)
		if errors.Is(err, os.ErrNotExist) && allowMissing {
			return nil
		}
		if err != nil || !info.IsDir() || info.Mode()&os.ModeSymlink != 0 {
			return ErrUnsafe
		}
		if current == path && (info.Mode().Perm() != 0700 || !owned(info)) {
			return ErrPermission
		}
		if filepath.Dir(current) == current {
			break
		}
	}
	return nil
}
func owned(info os.FileInfo) bool {
	s, ok := info.Sys().(*syscall.Stat_t)
	return ok && int(s.Uid) == os.Geteuid()
}
func cleanAbsolute(path string) bool {
	return path != "" && filepath.IsAbs(path) && filepath.Clean(path) == path
}
