package inventorystore

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"quantumwizard.hu/qwsg/internal/evidenceio"
	"strconv"
	"strings"
	"syscall"
)

const kernelLockName = ".write.flock"

const lockMarker = "qwsg.inventory-kernel-lock/1\n"

var ErrLocked = errors.New("inventory store has an active writer or reader")

// The inode is permanent: unlinking a flock file can create two lock domains.
// Legacy O_EXCL locks have no ownership evidence and cannot be stolen safely.
func (s *Store) lock() (func(), error) {
	if err := validatePrivateDir(s.root); err != nil {
		return nil, err
	}
	file, err := os.OpenFile(filepath.Join(s.root, kernelLockName), os.O_CREATE|os.O_RDWR|syscall.O_NOFOLLOW|syscall.O_NONBLOCK, 0600)
	if err != nil {
		return nil, fmt.Errorf("acquire inventory kernel lock: %w", err)
	}
	fail := func(err error) (func(), error) { file.Close(); return nil, err }
	info, err := file.Stat()
	if err != nil {
		return fail(err)
	}
	if !info.Mode().IsRegular() || info.Mode().Perm()&0077 != 0 || info.Size() != 0 {
		return fail(fmt.Errorf("%w: malformed kernel lock; preserve store for inspection", ErrUnsafePath))
	}
	if err := syscall.Flock(int(file.Fd()), syscall.LOCK_EX|syscall.LOCK_NB); err != nil {
		if errors.Is(err, syscall.EWOULDBLOCK) || errors.Is(err, syscall.EAGAIN) {
			return fail(ErrLocked)
		}
		return fail(fmt.Errorf("inventory kernel lock unavailable: %w", err))
	}
	// A permanent versioned guard also excludes older O_EXCL writers. Install
	// complete marker bytes atomically; an empty old lock is never overwritten.
	guard := filepath.Join(s.root, lockName)
	if _, err := os.Lstat(guard); errors.Is(err, os.ErrNotExist) {
		if err := atomicInstall(guard, []byte(lockMarker), nil, nil); err != nil {
			return fail(err)
		}
	} else if err != nil {
		return fail(err)
	}
	marker, err := evidenceio.ReadFile(guard, 128)
	if err != nil || !bytes.Equal(marker, []byte(lockMarker)) {
		return fail(fmt.Errorf("%w: legacy or malformed .write.lock ownership unknown; stop all older QWSG writers, preserve the store, then move the legacy lock outside the store before retrying", ErrCorrupt))
	}
	return func() { _ = file.Close() }, nil
}

// Directory input is bounded too. A limit violation leaves all evidence intact.
func boundedEntries(path string, limit int) ([]os.DirEntry, error) {
	dir, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer dir.Close()
	entries, err := dir.ReadDir(limit + 1)
	if err != nil && !errors.Is(err, io.EOF) {
		return nil, err
	}
	if len(entries) > limit {
		return nil, fmt.Errorf("%w: recovery entry limit exceeded; preserve store for inspection", ErrCorrupt)
	}
	return entries, nil
}

func temporaryName(name string) bool {
	const prefix = ".tmp-inventory-"
	if !strings.HasPrefix(name, prefix) {
		return false
	}
	suffix := strings.TrimPrefix(name, prefix)
	_, err := strconv.ParseUint(suffix, 10, 32)
	return err == nil && suffix != "" && !strings.HasPrefix(suffix, "+")
}

func privateArtifact(path string) error {
	info, err := os.Lstat(path)
	if err != nil {
		return err
	}
	if !info.Mode().IsRegular() || info.Mode().Perm()&0077 != 0 {
		return fmt.Errorf("%w: unsafe recovery artifact", ErrUnsafePath)
	}
	return nil
}

// Called only with the exclusive kernel lock and validated metadata. Classify
// and verify the entire transition before mutating any artifact. Counts follow
// the existing one-retirement/one-install transaction, not timestamps or age.
func (s *Store) recoverUnlocked() error {
	rootEntries, err := boundedEntries(s.root, 6)
	if err != nil {
		return err
	}
	var temps []string
	for _, entry := range rootEntries {
		name := entry.Name()
		switch name {
		case metadataName, snapshotsDir, kernelLockName, lockName:
		default:
			if !temporaryName(name) {
				return fmt.Errorf("%w: unknown root artifact; preserve store for inspection", ErrCorrupt)
			}
			path := filepath.Join(s.root, name)
			if err := privateArtifact(path); err != nil {
				return err
			}
			temps = append(temps, path)
		}
	}
	dir := filepath.Join(s.root, snapshotsDir)
	entries, err := boundedEntries(dir, s.retention+2)
	if err != nil {
		return err
	}
	var retired string
	var names []string
	for _, entry := range entries {
		name := entry.Name()
		path := filepath.Join(dir, name)
		if err := privateArtifact(path); err != nil {
			return err
		}
		switch {
		case temporaryName(name):
			temps = append(temps, path)
		case strings.HasPrefix(name, ".retire-"):
			if retired != "" {
				return fmt.Errorf("%w: multiple retirements; preserve store for inspection", ErrCorrupt)
			}
			retired = name
		case !strings.HasPrefix(name, ".") && strings.HasSuffix(name, ".json"):
			names = append(names, name)
		default:
			return fmt.Errorf("%w: unknown snapshot artifact; preserve store for inspection", ErrCorrupt)
		}
	}
	if len(names) > s.retention {
		return fmt.Errorf("%w: snapshot count exceeds retention", ErrCorrupt)
	}
	if retired != "" {
		original := strings.TrimPrefix(retired, ".retire-")
		if len(names) != s.retention-1 && len(names) != s.retention {
			return fmt.Errorf("%w: ambiguous retirement count; preserve store for inspection", ErrCorrupt)
		}
		for _, name := range names {
			if name == original {
				return fmt.Errorf("%w: conflicting retirement; preserve store for inspection", ErrCorrupt)
			}
			if _, err := s.loadUnlocked(name, name); err != nil {
				return err
			}
		}
		if _, err := s.loadUnlocked(retired, original); err != nil {
			return err
		}
		// Re-establish durability of observed committed links before deleting old
		// evidence, including after a previous directory-sync failure.
		if err := syncDir(dir); err != nil {
			return err
		}
		if len(names) == s.retention-1 {
			if err := os.Rename(filepath.Join(dir, retired), filepath.Join(dir, original)); err != nil {
				return err
			}
		} else {
			if err := os.Remove(filepath.Join(dir, retired)); err != nil {
				return err
			}
		}
		if err := syncDir(dir); err != nil {
			return err
		}
	}
	// Temps are never authoritative, even when they contain a complete envelope.
	// Sync directories first: a temp may be another hard link to the final object.
	if err := syncDir(s.root); err != nil {
		return err
	}
	if err := syncDir(dir); err != nil {
		return err
	}
	for _, path := range temps {
		if err := os.Remove(path); err != nil {
			return err
		}
	}
	if len(temps) != 0 {
		if err := syncDir(s.root); err != nil {
			return err
		}
		return syncDir(dir)
	}
	return nil
}
