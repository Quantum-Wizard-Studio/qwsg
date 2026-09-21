package evidenceio

import (
	"bytes"
	"errors"
	"io"
	"os"
	"path/filepath"
	"testing"
)

type counted struct{ n int64 }

func (r *counted) Read(p []byte) (int, error) {
	for i := range p {
		p[i] = 'x'
	}
	r.n += int64(len(p))
	return len(p), nil
}

func TestReadStopsAtLimitPlusOne(t *testing.T) {
	for _, limit := range []int64{4096, 4 << 20, 16 << 20} {
		r := &counted{}
		if _, err := Read(r, limit); !errors.Is(err, ErrSize) {
			t.Fatalf("oversize accepted: %v", err)
		}
		if r.n != limit+1 {
			t.Fatalf("read %d bytes, limit %d", r.n, limit)
		}
	}
	for _, data := range [][]byte{nil, []byte("small"), bytes.Repeat([]byte("x"), 64)} {
		got, err := Read(bytes.NewReader(data), 64)
		if err != nil || !bytes.Equal(got, data) {
			t.Fatalf("within bound: %v", err)
		}
	}
}

func TestPrivateRegularFilesOnly(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "evidence")
	if err := os.WriteFile(path, []byte("valid"), 0600); err != nil {
		t.Fatal(err)
	}
	if _, err := ReadFile(path, 4); !errors.Is(err, ErrSize) {
		t.Fatal(err)
	}
	link := filepath.Join(dir, "link")
	if err := os.Symlink(path, link); err != nil {
		t.Fatal(err)
	}
	if _, err := ReadFile(link, 64); err == nil {
		t.Fatal("symlink accepted")
	}
	if err := os.Chmod(path, 0644); err != nil {
		t.Fatal(err)
	}
	if _, err := ReadFile(path, 64); !errors.Is(err, ErrUnsafe) {
		t.Fatal(err)
	}
	if _, err := ReadFile(dir, 64); !errors.Is(err, ErrUnsafe) {
		t.Fatal(err)
	}
}

func TestReadPropagatesFailure(t *testing.T) {
	_, err := Read(failingReader{}, 64)
	if !errors.Is(err, io.ErrUnexpectedEOF) {
		t.Fatal(err)
	}
}

type failingReader struct{}

func (failingReader) Read([]byte) (int, error) { return 0, io.ErrUnexpectedEOF }
