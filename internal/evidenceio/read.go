// Package evidenceio bounds reads of private local evidence without trusting a
// pre-read file size. It does not interpret or repair evidence.
package evidenceio

import (
	"errors"
	"io"
	"os"
	"syscall"
)

var (
	ErrSize   = errors.New("local evidence exceeds size limit")
	ErrUnsafe = errors.New("local evidence must be a private regular file")
)

// Read consumes at most limit+1 bytes, including the oversize detection byte.
func Read(r io.Reader, limit int64) ([]byte, error) {
	b, err := io.ReadAll(io.LimitReader(r, limit+1))
	if err != nil {
		return nil, err
	}
	if int64(len(b)) > limit {
		return nil, ErrSize
	}
	return b, nil
}

func ReadFile(path string, limit int64) ([]byte, error) {
	f, err := os.OpenFile(path, os.O_RDONLY|syscall.O_NOFOLLOW|syscall.O_NONBLOCK, 0)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	info, err := f.Stat()
	if err != nil {
		return nil, err
	}
	if !info.Mode().IsRegular() || info.Mode().Perm()&0077 != 0 {
		return nil, ErrUnsafe
	}
	return Read(f, limit)
}
