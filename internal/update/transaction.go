package update

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"time"
)

type InstalledFile struct {
	Destination, Backup, SHA256 string
	Mode                        uint32
	Existed                     bool
}
type Transaction struct {
	Schema, FromVersion, ToVersion, ToCommit, Created string
	Complete                                          bool
	Files                                             []InstalledFile
}

var packageDestinations = map[string]string{
	"bin/qwsg":                               "usr/local/bin/qwsg",
	"lib/systemd/user/qwsg-guardian.service": "usr/local/lib/systemd/user/qwsg-guardian.service",
}

func destination(rel string) (string, bool) {
	if d, ok := packageDestinations[rel]; ok {
		return d, true
	}
	switch rel {
	case "README.md", "INSTALL.md", "LICENSE", "CHANGELOG.md", "qwsg-config.json", "RELEASE.json":
		return "usr/local/share/doc/qwsg/" + filepath.Base(rel), true
	}
	if filepath.Dir(rel) == "docs" && filepath.Ext(rel) == ".md" {
		return "usr/local/share/doc/qwsg/" + filepath.Base(rel), true
	}
	return "", false
}

func Apply(packageRoot, destRoot, backupRoot, fromVersion string) (tx Transaction, err error) {
	if err = ensureNewPrivateRoot(backupRoot); err != nil {
		return tx, err
	}
	tx = Transaction{Schema: "qwsg.update-transaction/1", FromVersion: fromVersion, ToVersion: "", Created: time.Now().UTC().Format(time.RFC3339Nano)}
	data, e := os.ReadFile(filepath.Join(packageRoot, "RELEASE.json"))
	if e != nil {
		return tx, e
	}
	var p Provenance
	if e = json.Unmarshal(data, &p); e != nil {
		return tx, e
	}
	tx.ToVersion = p.Version
	tx.ToCommit = p.Commit
	var pairs [][2]string
	for rel := range manifestSet(packageRoot) {
		if d, ok := destination(rel); ok {
			pairs = append(pairs, [2]string{rel, d})
		}
	}
	sort.Slice(pairs, func(i, j int) bool { return pairs[i][1] < pairs[j][1] })
	if len(pairs) < 8 {
		return tx, fmt.Errorf("package destination set incomplete")
	}
	defer func() {
		if err != nil {
			_ = restore(tx, destRoot, backupRoot)
		}
	}()
	for _, pair := range pairs {
		src := filepath.Join(packageRoot, pair[0])
		dst := filepath.Join(destRoot, pair[1])
		info, e := os.Lstat(dst)
		backup, hash := filepath.Join("files", pair[1]), ""
		existed := e == nil
		if e != nil && !os.IsNotExist(e) {
			return tx, e
		}
		if existed {
			if !info.Mode().IsRegular() || info.Mode()&os.ModeSymlink != 0 {
				return tx, fmt.Errorf("unsafe installed artifact: %s", pair[1])
			}
			hash, e = fileSHA(dst)
			if e != nil {
				return tx, e
			}
			if e = copyExclusive(dst, filepath.Join(backupRoot, backup), info.Mode().Perm()); e != nil {
				return tx, e
			}
			tx.Files = append(tx.Files, InstalledFile{Destination: pair[1], Backup: backup, SHA256: hash, Mode: uint32(info.Mode().Perm()), Existed: true})
		} else {
			tx.Files = append(tx.Files, InstalledFile{Destination: pair[1], Backup: backup, Mode: 0, Existed: false})
			if e = os.MkdirAll(filepath.Dir(dst), 0755); e != nil {
				return tx, e
			}
			if e = os.WriteFile(dst, []byte{}, 0600); e != nil {
				return tx, e
			}
		}
		mode := os.FileMode(0644)
		if pair[0] == "bin/qwsg" {
			mode = 0755
		}
		if e = replaceFile(src, dst, mode); e != nil {
			return tx, e
		}
	}
	tx.Complete = true
	if err = writeTransaction(backupRoot, tx); err != nil {
		return tx, err
	}
	return tx, nil
}

func Rollback(destRoot, backupRoot string) error {
	tx, err := ReadTransaction(backupRoot)
	if err != nil {
		return err
	}
	if !tx.Complete {
		return fmt.Errorf("rollback transaction incomplete")
	}
	return restore(tx, destRoot, backupRoot)
}
func ReadTransaction(root string) (Transaction, error) {
	info, err := os.Lstat(root)
	if err != nil || !info.IsDir() || info.Mode()&os.ModeSymlink != 0 || info.Mode().Perm() != 0700 {
		return Transaction{}, fmt.Errorf("unsafe rollback root")
	}
	data, err := os.ReadFile(filepath.Join(root, "transaction.json"))
	if err != nil {
		return Transaction{}, err
	}
	var tx Transaction
	if err = json.Unmarshal(data, &tx); err != nil || tx.Schema != "qwsg.update-transaction/1" || len(tx.Files) == 0 {
		return Transaction{}, fmt.Errorf("invalid rollback metadata")
	}
	return tx, nil
}
func restore(tx Transaction, destRoot, backupRoot string) error {
	for _, f := range tx.Files {
		if filepath.IsAbs(f.Destination) || filepath.IsAbs(f.Backup) || filepath.Clean(f.Destination) != f.Destination || filepath.Clean(f.Backup) != f.Backup {
			return fmt.Errorf("unsafe rollback metadata")
		}
		if !f.Existed {
			if err := os.Remove(filepath.Join(destRoot, f.Destination)); err != nil && !os.IsNotExist(err) {
				return err
			}
			continue
		}
		src := filepath.Join(backupRoot, f.Backup)
		hash, err := fileSHA(src)
		if err != nil || hash != f.SHA256 {
			return fmt.Errorf("rollback integrity mismatch")
		}
		if err = replaceFile(src, filepath.Join(destRoot, f.Destination), os.FileMode(f.Mode)); err != nil {
			return err
		}
	}
	return nil
}
func ensureNewPrivateRoot(path string) error {
	if filepath.IsAbs(path) == false {
		return fmt.Errorf("rollback root must be absolute")
	}
	if _, err := os.Lstat(path); !os.IsNotExist(err) {
		return fmt.Errorf("rollback root already exists")
	}
	return os.MkdirAll(path, 0700)
}
func copyExclusive(src, dst string, mode os.FileMode) error {
	if err := os.MkdirAll(filepath.Dir(dst), 0700); err != nil {
		return err
	}
	in, err := os.Open(src)
	if err != nil {
		return err
	}
	defer in.Close()
	out, err := os.OpenFile(dst, os.O_CREATE|os.O_EXCL|os.O_WRONLY, mode)
	if err != nil {
		return err
	}
	_, copyErr := io.Copy(out, in)
	closeErr := out.Close()
	if copyErr != nil {
		return copyErr
	}
	return closeErr
}
func replaceFile(src, dst string, mode os.FileMode) error {
	info, err := os.Lstat(dst)
	if err != nil || !info.Mode().IsRegular() || info.Mode()&os.ModeSymlink != 0 {
		return fmt.Errorf("unsafe destination")
	}
	if err = os.MkdirAll(filepath.Dir(dst), 0755); err != nil {
		return err
	}
	tmp, err := os.CreateTemp(filepath.Dir(dst), ".qwsg-update-")
	if err != nil {
		return err
	}
	name := tmp.Name()
	defer os.Remove(name)
	in, err := os.Open(src)
	if err != nil {
		tmp.Close()
		return err
	}
	_, e := io.Copy(tmp, in)
	in.Close()
	if e == nil {
		e = tmp.Chmod(mode)
	}
	if e == nil {
		e = tmp.Sync()
	}
	if ce := tmp.Close(); e == nil {
		e = ce
	}
	if e != nil {
		return e
	}
	return os.Rename(name, dst)
}
func writeTransaction(root string, tx Transaction) error {
	data, err := json.MarshalIndent(tx, "", "  ")
	if err != nil {
		return err
	}
	data = append(data, '\n')
	path := filepath.Join(root, "transaction.json")
	return os.WriteFile(path, data, 0600)
}
func manifestSet(root string) map[string]bool {
	result := map[string]bool{}
	data, err := os.ReadFile(filepath.Join(root, "MANIFEST.sha256"))
	if err != nil {
		return result
	}
	for _, line := range strings.Split(strings.TrimSpace(string(data)), "\n") {
		parts := strings.SplitN(line, "  ", 2)
		if len(parts) == 2 {
			result[parts[1]] = true
		}
	}
	return result
}
