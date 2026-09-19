// Package updatemutation coordinates package mutation; it grants no authorization.
package updatemutation

import (
	"errors"
	"os"
	"path/filepath"
	"syscall"
)

// ErrContended means another transaction owns this installation's mutation lock.
var ErrContended = errors.New("transaction_conflict")

// Lease is owned by one top-level coordinator until all helper calls, validation,
// recovery and rollback-record changes finish. It must not be copied or shared.
// Helpers execute inside that ownership scope and must not acquire recursively.
// Read-only discovery and awareness do not need a Lease.
type Lease struct{ file *os.File }

// Close releases ownership, including on failure. Never unlink the lock file:
// other processes may already have the same inode open.
func (l *Lease) Close() error { return l.file.Close() }

// Acquire tries once, without waiting. root must be the canonical private update
// directory of the installed instance, shared by manual and automatic callers.
// Keep the historical automatic.lock name to coordinate with existing clients.
func Acquire(root string) (*Lease, error) {
	info, err := os.Lstat(root)
	if err != nil || !info.IsDir() || info.Mode().Perm() != 0700 || int(info.Sys().(*syscall.Stat_t).Uid) != os.Getuid() {
		return nil, errors.New("unsafe update directory")
	}
	fd, err := syscall.Open(filepath.Join(root, "automatic.lock"), syscall.O_CREAT|syscall.O_RDWR|syscall.O_CLOEXEC|syscall.O_NOFOLLOW, 0600)
	if err != nil {
		return nil, err
	}
	f := os.NewFile(uintptr(fd), "automatic.lock")
	info, err = f.Stat()
	if err != nil || !info.Mode().IsRegular() || info.Mode().Perm() != 0600 || int(info.Sys().(*syscall.Stat_t).Uid) != os.Getuid() {
		f.Close()
		return nil, errors.New("unsafe update lock")
	}
	if err = syscall.Flock(fd, syscall.LOCK_EX|syscall.LOCK_NB); err != nil {
		f.Close()
		if errors.Is(err, syscall.EWOULDBLOCK) || errors.Is(err, syscall.EAGAIN) {
			return nil, ErrContended
		}
		return nil, err
	}
	return &Lease{file: f}, nil
}
