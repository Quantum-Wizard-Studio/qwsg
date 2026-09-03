package updateawareness

import (
	"errors"
	"io"
	"io/fs"
	"os"
	"path/filepath"
	"syscall"
	"time"
)

// RecordNotification atomically records a successful delivery only while the
// current authenticated actionable release still matches expectedIdentity.
func (s *Store) RecordNotification(expectedIdentity string, deliveredAt time.Time) error {
	lock, err := s.Lock()
	if err != nil {
		return err
	}
	defer lock.Release()
	value, err := s.Load()
	if err != nil {
		return err
	}
	if value.LastSuccess == nil || value.Status != UpdateAvailable {
		return ErrCorrupt
	}
	o := value.LastSuccess
	identity := NotificationIdentity(value.SourceID, value.Channel, o.ReleaseVersion, o.ArtifactSHA256, o.Authenticity.KeyID)
	if identity != expectedIdentity {
		return ErrCorrupt
	}
	value.LastNotification = &NotificationDelivery{Identity: identity, ReleaseVersion: o.ReleaseVersion, ArtifactSHA256: o.ArtifactSHA256, SigningKeyID: o.Authenticity.KeyID, DeliveredAt: deliveredAt.UTC()}
	value, err = Normalize(value)
	if err != nil {
		return err
	}
	return s.Publish(value)
}

const fileName = "awareness.json"

type Store struct {
	root                string
	beforeRename        func() error
	beforeDirectorySync func() error
}

func Open(stateRoot string) (*Store, error) {
	if stateRoot == "" || !filepath.IsAbs(stateRoot) || filepath.Clean(stateRoot) != stateRoot {
		return nil, ErrUnsafePath
	}
	if err := rejectSymlinks(stateRoot); err != nil && !errors.Is(err, fs.ErrNotExist) {
		return nil, err
	}
	return &Store{root: filepath.Join(stateRoot, "update")}, nil
}

func (s *Store) Publish(value State) error {
	document, err := Marshal(value)
	if err != nil {
		return err
	}
	if err = ensureDirectory(s.root); err != nil {
		return err
	}
	target := filepath.Join(s.root, fileName)
	if err = validateTarget(target, true); err != nil {
		return err
	}
	temporary, err := os.CreateTemp(s.root, ".awareness-*")
	if err != nil {
		return classify(err)
	}
	name, installed := temporary.Name(), false
	defer func() {
		if !installed {
			_ = os.Remove(name)
		}
	}()
	if err = temporary.Chmod(0600); err == nil {
		_, err = temporary.Write(document)
	}
	if err == nil {
		err = temporary.Sync()
	}
	if closeErr := temporary.Close(); err == nil {
		err = closeErr
	}
	if err != nil {
		return classify(err)
	}
	if s.beforeRename != nil {
		if err = s.beforeRename(); err != nil {
			return err
		}
	}
	if err = os.Rename(name, target); err != nil {
		return classify(err)
	}
	installed = true
	if s.beforeDirectorySync != nil {
		if err = s.beforeDirectorySync(); err != nil {
			return err
		}
	}
	directory, err := os.Open(s.root)
	if err != nil {
		return classify(err)
	}
	syncErr, closeErr := directory.Sync(), directory.Close()
	if syncErr != nil {
		return classify(syncErr)
	}
	return classify(closeErr)
}

func (s *Store) Load() (State, error) {
	if err := privateDirectory(s.root); err != nil {
		if errors.Is(err, fs.ErrNotExist) {
			return State{}, ErrMissing
		}
		return State{}, err
	}
	path := filepath.Join(s.root, fileName)
	if err := validateTarget(path, false); err != nil {
		return State{}, err
	}
	file, err := os.Open(path)
	if err != nil {
		if errors.Is(err, fs.ErrNotExist) {
			return State{}, ErrMissing
		}
		return State{}, classify(err)
	}
	info, err := file.Stat()
	if err != nil {
		_ = file.Close()
		return State{}, classify(err)
	}
	if info.Size() <= 0 || info.Size() > MaxEncodedSize {
		_ = file.Close()
		return State{}, ErrCorrupt
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

type Lock struct{ file *os.File }

func (s *Store) Lock() (*Lock, error) {
	if err := ensureDirectory(s.root); err != nil {
		return nil, err
	}
	path := filepath.Join(s.root, "awareness.lock")
	file, err := os.OpenFile(path, os.O_CREATE|os.O_RDWR, 0600)
	if err != nil {
		return nil, classify(err)
	}
	if err = validateOpenFile(file); err != nil {
		_ = file.Close()
		return nil, err
	}
	if err = syscall.Flock(int(file.Fd()), syscall.LOCK_EX|syscall.LOCK_NB); err != nil {
		_ = file.Close()
		if errors.Is(err, syscall.EWOULDBLOCK) || errors.Is(err, syscall.EAGAIN) {
			return nil, ErrContended
		}
		return nil, classify(err)
	}
	return &Lock{file: file}, nil
}
func (l *Lock) Release() error {
	if l == nil || l.file == nil {
		return ErrCorrupt
	}
	file := l.file
	l.file = nil
	unlock := syscall.Flock(int(file.Fd()), syscall.LOCK_UN)
	closeErr := file.Close()
	if unlock != nil {
		return classify(unlock)
	}
	return classify(closeErr)
}

func ensureDirectory(path string) error {
	if err := rejectSymlinks(path); err != nil && !errors.Is(err, fs.ErrNotExist) {
		return err
	}
	parent := filepath.Dir(path)
	if err := privateDirectory(parent); err != nil {
		return err
	}
	if err := os.Mkdir(path, 0700); err != nil && !errors.Is(err, fs.ErrExist) {
		return classify(err)
	}
	return privateDirectory(path)
}
func privateDirectory(path string) error {
	info, err := os.Lstat(path)
	if err != nil {
		return classify(err)
	}
	if info.Mode()&os.ModeSymlink != 0 || !info.IsDir() {
		return ErrUnsafePath
	}
	if info.Mode().Perm() != 0700 {
		return ErrPermission
	}
	stat, ok := info.Sys().(*syscall.Stat_t)
	if !ok || int(stat.Uid) != os.Geteuid() {
		return ErrPermission
	}
	return nil
}
func validateTarget(path string, absentOK bool) error {
	info, err := os.Lstat(path)
	if errors.Is(err, fs.ErrNotExist) && absentOK {
		return nil
	}
	if errors.Is(err, fs.ErrNotExist) {
		return ErrMissing
	}
	if err != nil {
		return classify(err)
	}
	if info.Mode()&os.ModeSymlink != 0 || !info.Mode().IsRegular() || info.Mode().Perm() != 0600 {
		return ErrUnsafePath
	}
	stat, ok := info.Sys().(*syscall.Stat_t)
	if !ok || int(stat.Uid) != os.Geteuid() || stat.Nlink != 1 {
		return ErrPermission
	}
	return nil
}
func validateOpenFile(file *os.File) error {
	info, err := file.Stat()
	if err != nil {
		return classify(err)
	}
	if !info.Mode().IsRegular() || info.Mode().Perm() != 0600 {
		return ErrPermission
	}
	stat, ok := info.Sys().(*syscall.Stat_t)
	if !ok || int(stat.Uid) != os.Geteuid() || stat.Nlink != 1 {
		return ErrPermission
	}
	return nil
}
func rejectSymlinks(path string) error {
	clean := filepath.Clean(path)
	for {
		info, err := os.Lstat(clean)
		if err == nil && info.Mode()&os.ModeSymlink != 0 {
			return ErrUnsafePath
		}
		if err != nil && !errors.Is(err, fs.ErrNotExist) {
			return classify(err)
		}
		parent := filepath.Dir(clean)
		if parent == clean {
			return nil
		}
		clean = parent
	}
}
func classify(err error) error {
	if err == nil {
		return nil
	}
	if errors.Is(err, fs.ErrNotExist) {
		return fs.ErrNotExist
	}
	if errors.Is(err, fs.ErrPermission) {
		return ErrPermission
	}
	return err
}
