package update

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

type Staged struct {
	Root, Archive, Sidecar, SHA256 string
	Release                        Release
}

func StageLocal(archivePath, sidecarPath, version, parent string) (Staged, error) {
	if err := requirePrivateDir(parent); err != nil {
		return Staged{}, err
	}
	want := "qwsg-" + version + "-linux-amd64.tar.gz"
	if filepath.Base(archivePath) != want || filepath.Base(sidecarPath) != want+".sha256" {
		return Staged{}, fmt.Errorf("local candidate identity mismatch")
	}
	root, err := os.MkdirTemp(parent, "transaction-")
	if err != nil {
		return Staged{}, err
	}
	if err = os.Chmod(root, 0700); err != nil {
		os.RemoveAll(root)
		return Staged{}, err
	}
	fail := func(e error) (Staged, error) { os.RemoveAll(root); return Staged{}, e }
	archive := filepath.Join(root, want)
	sidecar := archive + ".sha256"
	for src, dst := range map[string]string{archivePath: archive, sidecarPath: sidecar} {
		info, e := os.Lstat(src)
		if e != nil || !info.Mode().IsRegular() || info.Mode()&os.ModeSymlink != 0 {
			return fail(fmt.Errorf("unsafe local candidate"))
		}
		if e = copyExclusive(src, dst, 0600); e != nil {
			return fail(e)
		}
	}
	expected, e := readSidecar(sidecar, want)
	if e != nil {
		return fail(e)
	}
	actual, e := fileSHA(archive)
	if e != nil || actual != expected {
		return fail(fmt.Errorf("archive checksum mismatch"))
	}
	rel := Release{Version: version, Tag: "private-acceptance", Archive: Asset{Name: want, Size: fileSize(archive)}, Sidecar: Asset{Name: want + ".sha256", Size: fileSize(sidecar)}}
	return Staged{Root: root, Archive: archive, Sidecar: sidecar, SHA256: actual, Release: rel}, nil
}
func fileSize(path string) int64 {
	info, err := os.Stat(path)
	if err != nil {
		return -1
	}
	return info.Size()
}

func Acquire(ctx context.Context, client *http.Client, release Release, parent string) (Staged, error) {
	if err := requirePrivateDir(parent); err != nil {
		return Staged{}, err
	}
	root, err := os.MkdirTemp(parent, "transaction-")
	if err != nil {
		return Staged{}, err
	}
	if err = os.Chmod(root, 0700); err != nil {
		os.RemoveAll(root)
		return Staged{}, err
	}
	fail := func(e error) (Staged, error) { os.RemoveAll(root); return Staged{}, e }
	archive := filepath.Join(root, release.Archive.Name)
	sidecar := filepath.Join(root, release.Sidecar.Name)
	if err = download(ctx, client, release.Sidecar, sidecar); err != nil {
		return fail(err)
	}
	if err = download(ctx, client, release.Archive, archive); err != nil {
		return fail(err)
	}
	expected, err := readSidecar(sidecar, release.Archive.Name)
	if err != nil {
		return fail(err)
	}
	actual, err := fileSHA(archive)
	if err != nil {
		return fail(err)
	}
	if actual != expected {
		return fail(fmt.Errorf("archive checksum mismatch"))
	}
	return Staged{Root: root, Archive: archive, Sidecar: sidecar, SHA256: actual, Release: release}, nil
}

func requirePrivateDir(path string) error {
	info, err := os.Lstat(path)
	if err != nil {
		return err
	}
	if !info.IsDir() || info.Mode()&os.ModeSymlink != 0 || info.Mode().Perm() != 0700 {
		return fmt.Errorf("staging parent is unsafe")
	}
	return nil
}
func download(ctx context.Context, client *http.Client, asset Asset, target string) error {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, asset.URL, nil)
	if err != nil {
		return err
	}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return fmt.Errorf("asset HTTP %d", resp.StatusCode)
	}
	f, err := os.OpenFile(target, os.O_CREATE|os.O_EXCL|os.O_WRONLY, 0600)
	if err != nil {
		return err
	}
	n, copyErr := io.Copy(f, io.LimitReader(resp.Body, asset.Size+1))
	closeErr := f.Close()
	if copyErr != nil {
		return copyErr
	}
	if closeErr != nil {
		return closeErr
	}
	if n != asset.Size {
		return fmt.Errorf("asset size mismatch")
	}
	return nil
}
func readSidecar(path, name string) (string, error) {
	b, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}
	fields := strings.Fields(string(b))
	if len(fields) != 2 || fields[1] != name || len(fields[0]) != 64 {
		return "", fmt.Errorf("invalid checksum sidecar")
	}
	if _, err = hex.DecodeString(fields[0]); err != nil {
		return "", fmt.Errorf("invalid checksum sidecar")
	}
	return strings.ToLower(fields[0]), nil
}
func fileSHA(path string) (string, error) {
	f, err := os.Open(path)
	if err != nil {
		return "", err
	}
	defer f.Close()
	h := sha256.New()
	if _, err = io.Copy(h, f); err != nil {
		return "", err
	}
	return hex.EncodeToString(h.Sum(nil)), nil
}
