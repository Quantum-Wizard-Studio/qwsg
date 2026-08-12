package operatorstate

import (
	"errors"
	"fmt"
	"io"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
	"syscall"
)

const FileName = "current-operator-state.json"

type Store struct {
	root                string
	beforeRename        func() error
	beforeDirectorySync func() error
}

func Open(root string) (*Store, error) {
	if root == "" || !filepath.IsAbs(root) || filepath.Clean(root) != root {
		return nil, ErrUnsafePath
	}
	if err := rejectSymlinks(root); err != nil && !errors.Is(err, fs.ErrNotExist) {
		return nil, err
	}
	return &Store{root: root}, nil
}

// EnsurePrivateRoot safely creates and validates an absent private state root.
// Existing ancestors are inspected for symlinks but are never chmodded or
// otherwise modified.
func EnsurePrivateRoot(root string) error {
	store, err := Open(root)
	if err != nil {
		return err
	}
	return ensurePrivateDir(store.root)
}

func (store *Store) Publish(value State) error {
	document, err := MarshalCanonical(value)
	if err != nil {
		return err
	}
	if err = ensurePrivateDir(store.root); err != nil {
		return err
	}
	target := filepath.Join(store.root, FileName)
	if info, targetErr := os.Lstat(target); targetErr == nil {
		if info.Mode()&os.ModeSymlink != 0 || !info.Mode().IsRegular() {
			return ErrUnsafePath
		}
		if info.Mode().Perm() != 0o600 {
			return ErrPermission
		}
		if stat, ok := info.Sys().(*syscall.Stat_t); !ok || int(stat.Uid) != os.Geteuid() {
			return ErrPermission
		}
	} else if !errors.Is(targetErr, fs.ErrNotExist) {
		return classify(targetErr)
	}
	temporary, err := os.CreateTemp(store.root, ".current-operator-state-*")
	if err != nil {
		return classify(err)
	}
	temporaryName := temporary.Name()
	keep := false
	defer func() {
		if !keep {
			_ = os.Remove(temporaryName)
		}
	}()
	if err = temporary.Chmod(0o600); err == nil {
		_, err = temporary.Write(document)
	}
	if err == nil {
		err = temporary.Sync()
	}
	closeErr := temporary.Close()
	if err == nil {
		err = closeErr
	}
	if err != nil {
		return classify(err)
	}
	if store.beforeRename != nil {
		if err = store.beforeRename(); err != nil {
			return err
		}
	}
	if err = os.Rename(temporaryName, target); err != nil {
		return classify(err)
	}
	keep = true
	if store.beforeDirectorySync != nil {
		if err = store.beforeDirectorySync(); err != nil {
			return err
		}
	}
	directory, err := os.Open(store.root)
	if err != nil {
		return classify(err)
	}
	syncErr := directory.Sync()
	closeErr = directory.Close()
	if syncErr != nil {
		return classify(syncErr)
	}
	return classify(closeErr)
}

func (store *Store) Load() (State, error) {
	if err := privateDir(store.root); err != nil {
		return State{}, err
	}
	path := filepath.Join(store.root, FileName)
	info, err := os.Lstat(path)
	if errors.Is(err, fs.ErrNotExist) {
		return State{}, ErrMissing
	}
	if err != nil {
		return State{}, classify(err)
	}
	if info.Mode()&os.ModeSymlink != 0 || !info.Mode().IsRegular() {
		return State{}, ErrUnsafePath
	}
	if info.Mode().Perm() != 0o600 {
		return State{}, ErrPermission
	}
	if stat, ok := info.Sys().(*syscall.Stat_t); !ok || int(stat.Uid) != os.Geteuid() {
		return State{}, ErrPermission
	}
	if info.Size() <= 0 || info.Size() > MaxEncodedSize {
		return State{}, ErrCorrupt
	}
	file, err := os.Open(path)
	if err != nil {
		return State{}, classify(err)
	}
	document, readErr := io.ReadAll(io.LimitReader(file, MaxEncodedSize+1))
	closeErr := file.Close()
	if readErr != nil {
		return State{}, classify(readErr)
	}
	if closeErr != nil {
		return State{}, classify(closeErr)
	}
	return Decode(document)
}

func ensurePrivateDir(root string) error {
	if err := rejectSymlinks(root); err != nil && !errors.Is(err, fs.ErrNotExist) {
		return err
	}
	if _, err := os.Lstat(root); errors.Is(err, fs.ErrNotExist) {
		if err = os.MkdirAll(root, 0o700); err != nil {
			return classify(err)
		}
		if err = os.Chmod(root, 0o700); err != nil {
			return classify(err)
		}
	} else if err != nil {
		return classify(err)
	}
	return privateDir(root)
}
func privateDir(root string) error {
	info, err := os.Lstat(root)
	if errors.Is(err, fs.ErrNotExist) {
		return ErrMissing
	}
	if err != nil {
		return classify(err)
	}
	if info.Mode()&os.ModeSymlink != 0 || !info.IsDir() {
		return ErrUnsafePath
	}
	if info.Mode().Perm() != 0o700 {
		return ErrPermission
	}
	if stat, ok := info.Sys().(*syscall.Stat_t); !ok || int(stat.Uid) != os.Geteuid() {
		return ErrPermission
	}
	return nil
}
func rejectSymlinks(path string) error {
	current := string(filepath.Separator)
	for _, part := range split(path) {
		current = filepath.Join(current, part)
		info, err := os.Lstat(current)
		if errors.Is(err, fs.ErrNotExist) {
			return err
		}
		if err != nil {
			return classify(err)
		}
		if info.Mode()&os.ModeSymlink != 0 {
			return ErrUnsafePath
		}
	}
	return nil
}
func split(path string) []string {
	relative := strings.TrimPrefix(path, string(filepath.Separator))
	if relative == "" {
		return nil
	}
	return strings.Split(relative, string(filepath.Separator))
}
func classify(err error) error {
	if err == nil {
		return nil
	}
	if errors.Is(err, fs.ErrPermission) {
		return fmt.Errorf("%w", ErrPermission)
	}
	return err
}
