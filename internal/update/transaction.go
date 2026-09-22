package update

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"quantumwizard.hu/qwsg/internal/evidenceio"
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
	Prepared                                          bool // Complete rollback set durably recorded before any destination change.
	// MutationStarted is conservative: a destination-changing operation was entered.
	MutationStarted, RollbackAttempted, RollbackSucceeded bool
	Files                                                 []InstalledFile
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
		if err != nil && tx.MutationStarted {
			tx.RollbackAttempted = true
			rollbackErr := restore(tx, destRoot, backupRoot)
			if rollbackErr == nil {
				rollbackErr = validateRestored(tx, destRoot)
			}
			tx.RollbackSucceeded = rollbackErr == nil
			if journalErr := writeTransaction(backupRoot, tx); journalErr != nil {
				err = errors.Join(err, journalErr)
			}
			if rollbackErr != nil {
				err = errors.Join(err, fmt.Errorf("rollback failed: %w", rollbackErr))
			}
		}
	}()
	for _, pair := range pairs {
		if err = safeArtifactPath(destRoot, pair[1]); err != nil {
			return tx, err
		}
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
		}
	}
	if err = validateRollbackSource(tx, destRoot, backupRoot); err != nil {
		return tx, err
	}
	tx.Prepared = true
	if err = writeTransaction(backupRoot, tx); err != nil {
		return tx, err
	}
	if err = syncTreeDirectories(backupRoot); err != nil {
		return tx, err
	}
	if err = beforeMutation(tx); err != nil {
		return tx, err
	}
	for _, pair := range pairs {
		src := filepath.Join(packageRoot, pair[0])
		dst := filepath.Join(destRoot, pair[1])
		if _, e := os.Lstat(dst); os.IsNotExist(e) {
			tx.MutationStarted = true
			if err = os.MkdirAll(filepath.Dir(dst), 0755); err != nil {
				return tx, err
			}
			if err = os.WriteFile(dst, nil, 0600); err != nil {
				return tx, err
			}
		}
		mode := os.FileMode(0644)
		if pair[0] == "bin/qwsg" {
			mode = 0755
		}
		tx.MutationStarted = true
		if err = replaceFile(src, dst, mode); err != nil {
			return tx, err
		}
		afterMutation()
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
	if !tx.Complete && !tx.Prepared {
		return fmt.Errorf("rollback transaction incomplete")
	}
	if err = restore(tx, destRoot, backupRoot); err != nil {
		return err
	}
	return validateRestored(tx, destRoot)
}
func ReadTransaction(root string) (Transaction, error) {
	info, err := os.Lstat(root)
	if err != nil || !info.IsDir() || info.Mode()&os.ModeSymlink != 0 || info.Mode().Perm() != 0700 {
		return Transaction{}, fmt.Errorf("unsafe rollback root")
	}
	data, err := evidenceio.ReadFile(filepath.Join(root, "transaction.json"), 1<<20)
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
	if err := validateRollbackSource(tx, destRoot, backupRoot); err != nil {
		return err
	}
	for _, f := range tx.Files {
		if filepath.IsAbs(f.Destination) || filepath.IsAbs(f.Backup) || filepath.Clean(f.Destination) != f.Destination || filepath.Clean(f.Backup) != f.Backup {
			return fmt.Errorf("unsafe rollback metadata")
		}
		if !f.Existed {
			if err := os.Remove(filepath.Join(destRoot, f.Destination)); err != nil && !os.IsNotExist(err) {
				return err
			}
			if err := syncDirectory(filepath.Dir(filepath.Join(destRoot, f.Destination))); err != nil {
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
	if copyErr == nil {
		copyErr = out.Sync()
	}
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
	if err = os.Rename(name, dst); err != nil {
		return err
	}
	return syncDirectory(filepath.Dir(dst))
}
func writeTransaction(root string, tx Transaction) error {
	data, err := json.MarshalIndent(tx, "", "  ")
	if err != nil {
		return err
	}
	data = append(data, '\n')
	path := filepath.Join(root, "transaction.json")
	tmp, err := os.CreateTemp(root, ".transaction-")
	if err != nil {
		return err
	}
	defer os.Remove(tmp.Name())
	if _, err = tmp.Write(data); err == nil {
		err = tmp.Sync()
	}
	if ce := tmp.Close(); err == nil {
		err = ce
	}
	if err != nil {
		return err
	}
	if err = os.Rename(tmp.Name(), path); err != nil {
		return err
	}
	return syncDirectory(root)
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

// validateRestored checks the complete recorded write set, not only the binary.
func validateRestored(tx Transaction, destRoot string) error {
	for _, f := range tx.Files {
		path := filepath.Join(destRoot, f.Destination)
		info, err := os.Lstat(path)
		if !f.Existed {
			if !os.IsNotExist(err) {
				return fmt.Errorf("rollback removal validation failed")
			}
			continue
		}
		if err != nil || !info.Mode().IsRegular() || uint32(info.Mode().Perm()) != f.Mode {
			return fmt.Errorf("rollback destination validation failed")
		}
		hash, err := fileSHA(path)
		if err != nil || hash != f.SHA256 {
			return fmt.Errorf("rollback restored integrity mismatch")
		}
	}
	return nil
}

// Test seam at the actual durable prepare boundary; no runtime fault controls.
var beforeMutation = func(Transaction) error { return nil }
var afterMutation = func() {}
var syncDirectory = func(path string) error {
	d, err := os.Open(path)
	if err != nil {
		return err
	}
	defer d.Close()
	return d.Sync()
}

func syncTreeDirectories(root string) error {
	var dirs []string
	err := filepath.WalkDir(root, func(path string, d os.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() {
			dirs = append(dirs, path)
		}
		return nil
	})
	if err != nil {
		return err
	}
	for i := len(dirs) - 1; i >= 0; i-- {
		if err = syncDirectory(dirs[i]); err != nil {
			return err
		}
	}
	// MkdirAll may have created rollback/user ancestors too.
	for p := filepath.Dir(root); ; p = filepath.Dir(p) {
		if err = syncDirectory(p); err != nil {
			return err
		}
		if p == filepath.Dir(p) {
			break
		}
	}
	return nil
}

func safeArtifactPath(root, rel string) error {
	if filepath.IsAbs(rel) || filepath.Clean(rel) != rel || rel == ".." || strings.HasPrefix(rel, "../") {
		return fmt.Errorf("unsafe rollback path")
	}
	p := root
	for _, part := range strings.Split(rel, string(filepath.Separator)) {
		p = filepath.Join(p, part)
		info, err := os.Lstat(p)
		if os.IsNotExist(err) {
			continue
		}
		if err != nil {
			return err
		}
		if info.Mode()&os.ModeSymlink != 0 {
			return fmt.Errorf("symlink in rollback path")
		}
	}
	return nil
}

// Check the WHOLE source before the first restore; a later corrupt member must
// never replace an earlier valid installed artifact.
func validateRollbackSource(tx Transaction, destRoot, backupRoot string) error {
	seen := map[string]bool{}
	for _, f := range tx.Files {
		allowed := false
		for _, d := range packageDestinations {
			if f.Destination == d {
				allowed = true
			}
		}
		if strings.HasPrefix(f.Destination, "usr/local/share/doc/qwsg/") {
			base := strings.TrimPrefix(f.Destination, "usr/local/share/doc/qwsg/")
			d, ok := destination(base)
			if !ok {
				d, ok = destination("docs/" + base)
			}
			allowed = ok && d == f.Destination
		}
		if !allowed || seen[f.Destination] || f.Backup != "files/"+f.Destination || f.Mode > 0777 {
			return fmt.Errorf("invalid rollback write set")
		}
		seen[f.Destination] = true
		if err := safeArtifactPath(destRoot, f.Destination); err != nil {
			return err
		}
		if err := safeArtifactPath(backupRoot, f.Backup); err != nil {
			return err
		}
		if f.Existed {
			info, err := os.Lstat(filepath.Join(backupRoot, f.Backup))
			if err != nil || !info.Mode().IsRegular() {
				return fmt.Errorf("invalid rollback source")
			}
			hash, err := fileSHA(filepath.Join(backupRoot, f.Backup))
			if err != nil || hash != f.SHA256 {
				return fmt.Errorf("rollback integrity mismatch")
			}
		}
	}
	return nil
}

// ValidateApplied verifies every installed package destination against the
// authenticated staged package before the coordinator starts Guardian.
func ValidateApplied(packageRoot, destRoot string) error {
	for rel := range manifestSet(packageRoot) {
		d, ok := destination(rel)
		if !ok {
			continue
		}
		if err := safeArtifactPath(destRoot, d); err != nil {
			return err
		}
		path := filepath.Join(destRoot, d)
		info, err := os.Lstat(path)
		mode := os.FileMode(0644)
		if rel == "bin/qwsg" {
			mode = 0755
		}
		if err != nil || !info.Mode().IsRegular() || info.Mode().Perm() != mode {
			return fmt.Errorf("installed package mode mismatch")
		}
		want, err := fileSHA(filepath.Join(packageRoot, rel))
		if err != nil {
			return err
		}
		got, err := fileSHA(path)
		if err != nil || got != want {
			return fmt.Errorf("installed package integrity mismatch")
		}
	}
	return nil
}
