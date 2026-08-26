package update

import (
	"archive/tar"
	"compress/gzip"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

type Provenance struct{ Schema, Version, Commit, Built, Platform string }
type Package struct {
	Root       string
	Provenance Provenance
	Files      []string
}

func VerifyPackage(staged Staged) (Package, error) {
	extract := filepath.Join(staged.Root, "package")
	if err := os.Mkdir(extract, 0700); err != nil {
		return Package{}, err
	}
	f, err := os.Open(staged.Archive)
	if err != nil {
		return Package{}, err
	}
	defer f.Close()
	gz, err := gzip.NewReader(f)
	if err != nil {
		return Package{}, fmt.Errorf("archive gzip: %w", err)
	}
	defer gz.Close()
	tr := tar.NewReader(gz)
	rootName := "qwsg-" + staged.Release.Version + "-linux-amd64"
	seen := map[string]bool{}
	count := 0
	for {
		h, e := tr.Next()
		if e == io.EOF {
			break
		}
		if e != nil {
			return Package{}, fmt.Errorf("archive: %w", e)
		}
		count++
		if count > 128 {
			return Package{}, fmt.Errorf("archive member bound exceeded")
		}
		name := strings.TrimSuffix(h.Name, "/")
		if name == "" || filepath.IsAbs(name) || strings.Contains(name, "\\") || strings.Contains("/"+name+"/", "/../") || !(name == rootName || strings.HasPrefix(name, rootName+"/")) {
			return Package{}, fmt.Errorf("unsafe archive path")
		}
		if seen[name] {
			return Package{}, fmt.Errorf("duplicate archive member")
		}
		seen[name] = true
		target := filepath.Join(extract, filepath.FromSlash(name))
		if !strings.HasPrefix(target, extract+string(os.PathSeparator)) {
			return Package{}, fmt.Errorf("unsafe archive target")
		}
		switch h.Typeflag {
		case tar.TypeDir:
			if err = os.MkdirAll(target, 0700); err != nil {
				return Package{}, err
			}
		case tar.TypeReg, tar.TypeRegA:
			if h.Size < 0 || h.Size > 128<<20 {
				return Package{}, fmt.Errorf("unsafe archive member size")
			}
			if err = os.MkdirAll(filepath.Dir(target), 0700); err != nil {
				return Package{}, err
			}
			mode := os.FileMode(0600)
			if strings.HasSuffix(name, "/bin/qwsg") || strings.HasSuffix(name, "/install.sh") || strings.HasSuffix(name, "/uninstall.sh") {
				mode = 0700
			}
			out, e := os.OpenFile(target, os.O_CREATE|os.O_EXCL|os.O_WRONLY, mode)
			if e != nil {
				return Package{}, e
			}
			n, e := io.Copy(out, io.LimitReader(tr, h.Size+1))
			ce := out.Close()
			if e != nil {
				return Package{}, e
			}
			if ce != nil {
				return Package{}, ce
			}
			if n != h.Size {
				return Package{}, fmt.Errorf("archive member size mismatch")
			}
		default:
			return Package{}, fmt.Errorf("unsafe archive member type")
		}
	}
	root := filepath.Join(extract, rootName)
	if !seen[rootName] {
		return Package{}, fmt.Errorf("archive root missing")
	}
	files, err := verifyManifest(root)
	if err != nil {
		return Package{}, err
	}
	for _, required := range []string{"bin/qwsg", "install.sh", "uninstall.sh", "MANIFEST.sha256", "RELEASE.json", "LICENSE", "README.md", "INSTALL.md"} {
		if _, err = os.Stat(filepath.Join(root, required)); err != nil {
			return Package{}, fmt.Errorf("required package file missing: %s", required)
		}
	}
	data, err := os.ReadFile(filepath.Join(root, "RELEASE.json"))
	if err != nil {
		return Package{}, err
	}
	var p Provenance
	if err = json.Unmarshal(data, &p); err != nil {
		return Package{}, fmt.Errorf("release provenance: %w", err)
	}
	if p.Schema != "qwsg.release/1" || p.Version != staged.Release.Version || p.Platform != "linux-amd64" || len(p.Commit) != 40 || !lowerHex(p.Commit) || p.Built == "" {
		return Package{}, fmt.Errorf("release provenance mismatch")
	}
	return Package{Root: root, Provenance: p, Files: files}, nil
}

func verifyManifest(root string) ([]string, error) {
	data, err := os.ReadFile(filepath.Join(root, "MANIFEST.sha256"))
	if err != nil {
		return nil, err
	}
	lines := strings.Split(strings.TrimSuffix(string(data), "\n"), "\n")
	if len(lines) == 0 || len(lines) > 128 {
		return nil, fmt.Errorf("invalid manifest")
	}
	seen := map[string]bool{}
	files := make([]string, 0, len(lines))
	for _, line := range lines {
		parts := strings.SplitN(line, "  ", 2)
		if len(parts) != 2 || len(parts[0]) != 64 || !lowerHex(parts[0]) || parts[1] == "" || filepath.IsAbs(parts[1]) || strings.Contains("/"+parts[1]+"/", "/../") || seen[parts[1]] {
			return nil, fmt.Errorf("invalid manifest")
		}
		seen[parts[1]] = true
		path := filepath.Join(root, filepath.FromSlash(parts[1]))
		info, e := os.Lstat(path)
		if e != nil || !info.Mode().IsRegular() {
			return nil, fmt.Errorf("manifest file invalid: %s", parts[1])
		}
		actual, e := fileSHA(path)
		if e != nil || actual != parts[0] {
			return nil, fmt.Errorf("manifest checksum mismatch: %s", parts[1])
		}
		files = append(files, parts[1])
	}
	sort.Strings(files)
	return files, nil
}
func lowerHex(s string) bool {
	_, err := hex.DecodeString(s)
	if err != nil {
		return false
	}
	return s == strings.ToLower(s)
}
func bytesSHA(b []byte) string { h := sha256.Sum256(b); return hex.EncodeToString(h[:]) }
