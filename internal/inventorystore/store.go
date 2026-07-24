package inventorystore

import (
	"bufio"
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"time"

	"quantumwizard.hu/qwsg/internal/inventory"
)

const (
	FormatName       = "qwsg.digital-twin"
	FormatVersion    = "1.0"
	DefaultRetention = 10
	MaxRetention     = 1000
	MaxDocumentSize  = 16 << 20

	metadataName = "store.json"
	snapshotsDir = "snapshots"
	lockName     = ".write.lock"
)

var (
	ErrCorrupt      = errors.New("inventory store data is corrupt")
	ErrIncompatible = errors.New("inventory store format is incompatible")
	ErrUnsafePath   = errors.New("inventory store path is unsafe")
	ErrConflict     = errors.New("inventory snapshot already exists")
)

type Metadata struct {
	FormatName    string `json:"format_name"`
	FormatVersion string `json:"format_version"`
	Retention     int    `json:"retention"`
}

type Envelope struct {
	FormatName             string             `json:"format_name"`
	FormatVersion          string             `json:"format_version"`
	CreatedAt              time.Time          `json:"created_at"`
	SnapshotID             string             `json:"snapshot_id"`
	SubjectID              string             `json:"subject_id"`
	InventorySchemaVersion string             `json:"inventory_schema_version"`
	PayloadSHA256          string             `json:"payload_sha256"`
	Inventory              inventory.Snapshot `json:"inventory"`
}

type Store struct {
	root          string
	retention     int
	beforeInstall func() error
}

func Open(root string, retention int) (*Store, error) {
	if !filepath.IsAbs(root) || filepath.Clean(root) != root {
		return nil, fmt.Errorf("%w: store root must be a clean absolute path", ErrUnsafePath)
	}
	if err := rejectSymlinkComponents(root); err != nil {
		return nil, err
	}
	if retention < 1 || retention > MaxRetention {
		return nil, fmt.Errorf("retention must be between 1 and %d", MaxRetention)
	}
	return &Store{root: root, retention: retention}, nil
}

func (s *Store) Root() string { return s.root }

func (s *Store) Save(snapshot inventory.Snapshot) (string, error) {
	if err := validateSnapshot(snapshot); err != nil {
		return "", fmt.Errorf("refusing to persist invalid inventory: %w", err)
	}
	if err := s.ensureLayout(); err != nil {
		return "", err
	}
	unlock, err := s.lock()
	if err != nil {
		return "", err
	}
	defer unlock()

	names, err := s.listUnlocked()
	if err != nil {
		return "", err
	}
	if len(names) > s.retention {
		return "", fmt.Errorf("%w: snapshot count exceeds configured retention", ErrCorrupt)
	}

	payload, err := json.Marshal(snapshot)
	if err != nil {
		return "", fmt.Errorf("encode inventory payload: %w", err)
	}
	sum := sha256.Sum256(payload)
	envelope := Envelope{
		FormatName: FormatName, FormatVersion: FormatVersion,
		CreatedAt: snapshot.CompletedAt.UTC(), SnapshotID: snapshot.SnapshotID,
		SubjectID: snapshot.InstanceID, InventorySchemaVersion: snapshot.SchemaVersion,
		PayloadSHA256: hex.EncodeToString(sum[:]), Inventory: snapshot,
	}
	document, err := json.MarshalIndent(envelope, "", "  ")
	if err != nil {
		return "", fmt.Errorf("encode persisted inventory: %w", err)
	}
	document = append(document, '\n')
	if len(document) > MaxDocumentSize {
		return "", fmt.Errorf("persisted inventory exceeds %d bytes", MaxDocumentSize)
	}

	name := snapshotName(snapshot)
	target := filepath.Join(s.root, snapshotsDir, name)
	if _, err := os.Lstat(target); err == nil {
		return "", ErrConflict
	} else if !errors.Is(err, os.ErrNotExist) {
		return "", fmt.Errorf("inspect snapshot target: %w", err)
	}

	var retired, retiredOriginal string
	if len(names) == s.retention {
		retiredOriginal = filepath.Join(s.root, snapshotsDir, names[0])
		retired = filepath.Join(s.root, snapshotsDir, ".retire-"+names[0])
		if _, err := os.Lstat(retired); !errors.Is(err, os.ErrNotExist) {
			if err == nil {
				return "", fmt.Errorf("%w: stale retention artifact", ErrCorrupt)
			}
			return "", fmt.Errorf("inspect retention target: %w", err)
		}
		if err := os.Rename(retiredOriginal, retired); err != nil {
			return "", fmt.Errorf("prepare retention transaction: %w", err)
		}
		if err := syncDir(filepath.Dir(retired)); err != nil {
			_ = os.Rename(retired, retiredOriginal)
			return "", fmt.Errorf("sync retention transaction: %w", err)
		}
	}
	restoreRetired := func() {
		if retired != "" {
			_ = os.Rename(retired, retiredOriginal)
			_ = syncDir(filepath.Dir(retired))
		}
	}

	if err := atomicInstall(target, document, s.beforeInstall); err != nil {
		restoreRetired()
		return "", err
	}
	if retired != "" {
		if err := os.Remove(retired); err != nil {
			removeErr := os.Remove(target)
			restoreRetired()
			if removeErr != nil {
				return "", fmt.Errorf("finalize retention: %v; rollback new snapshot: %v", err, removeErr)
			}
			return "", fmt.Errorf("finalize retention: %w", err)
		}
		if err := syncDir(filepath.Dir(retired)); err != nil {
			return "", fmt.Errorf("sync completed retention: %w", err)
		}
	}
	return name, nil
}

func (s *Store) List() ([]string, error) {
	if err := s.validateLayout(); err != nil {
		return nil, err
	}
	return s.listUnlocked()
}

func (s *Store) LoadLatest() (inventory.Snapshot, string, error) {
	names, err := s.List()
	if err != nil {
		return inventory.Snapshot{}, "", err
	}
	if len(names) == 0 {
		return inventory.Snapshot{}, "", os.ErrNotExist
	}
	name := names[len(names)-1]
	snapshot, err := s.Load(name)
	return snapshot, name, err
}

func (s *Store) Load(name string) (inventory.Snapshot, error) {
	if filepath.Base(name) != name || !strings.HasSuffix(name, ".json") || strings.HasPrefix(name, ".") {
		return inventory.Snapshot{}, fmt.Errorf("%w: invalid snapshot name", ErrUnsafePath)
	}
	if err := s.validateLayout(); err != nil {
		return inventory.Snapshot{}, err
	}
	path := filepath.Join(s.root, snapshotsDir, name)
	info, err := os.Lstat(path)
	if err != nil {
		return inventory.Snapshot{}, err
	}
	if !info.Mode().IsRegular() || info.Mode().Perm()&0o077 != 0 || info.Size() > MaxDocumentSize {
		return inventory.Snapshot{}, fmt.Errorf("%w: invalid snapshot file", ErrUnsafePath)
	}
	document, err := os.ReadFile(path)
	if err != nil {
		return inventory.Snapshot{}, fmt.Errorf("read snapshot: %w", err)
	}
	if err := rejectDuplicateKeys(document); err != nil {
		return inventory.Snapshot{}, fmt.Errorf("%w: %v", ErrCorrupt, err)
	}
	var envelope Envelope
	decoder := json.NewDecoder(bytes.NewReader(document))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&envelope); err != nil {
		return inventory.Snapshot{}, fmt.Errorf("%w: decode envelope: %v", ErrCorrupt, err)
	}
	if err := requireJSONEOF(decoder); err != nil {
		return inventory.Snapshot{}, fmt.Errorf("%w: %v", ErrCorrupt, err)
	}
	if envelope.FormatName != FormatName || envelope.FormatVersion != FormatVersion {
		return inventory.Snapshot{}, ErrIncompatible
	}
	if envelope.SnapshotID != envelope.Inventory.SnapshotID ||
		envelope.SubjectID != envelope.Inventory.InstanceID ||
		envelope.InventorySchemaVersion != envelope.Inventory.SchemaVersion ||
		!envelope.CreatedAt.Equal(envelope.Inventory.CompletedAt) {
		return inventory.Snapshot{}, fmt.Errorf("%w: envelope and inventory identity differ", ErrCorrupt)
	}
	if name != snapshotName(envelope.Inventory) {
		return inventory.Snapshot{}, fmt.Errorf("%w: snapshot filename does not match its identity", ErrCorrupt)
	}
	payload, err := json.Marshal(envelope.Inventory)
	if err != nil {
		return inventory.Snapshot{}, fmt.Errorf("%w: encode payload: %v", ErrCorrupt, err)
	}
	sum := sha256.Sum256(payload)
	if envelope.PayloadSHA256 != hex.EncodeToString(sum[:]) {
		return inventory.Snapshot{}, fmt.Errorf("%w: payload integrity mismatch", ErrCorrupt)
	}
	if err := validateSnapshot(envelope.Inventory); err != nil {
		return inventory.Snapshot{}, fmt.Errorf("%w: invalid inventory: %v", ErrCorrupt, err)
	}
	return envelope.Inventory, nil
}

func validateSnapshot(snapshot inventory.Snapshot) error {
	if snapshot.Canonical.SchemaName == "" {
		return errors.New("canonical inventory is required")
	}
	for _, category := range snapshot.Categories {
		for _, item := range category.Items {
			for _, fact := range item.Facts {
				if fact.Sensitivity == "secret_prohibited" {
					return errors.New("legacy inventory contains a secret_prohibited fact")
				}
			}
		}
	}
	return inventory.Validate(snapshot)
}

func snapshotName(snapshot inventory.Snapshot) string {
	sum := sha256.Sum256([]byte(snapshot.SnapshotID))
	return snapshot.CompletedAt.UTC().Format("20060102T150405.000000000Z") + "_" + hex.EncodeToString(sum[:8]) + ".json"
}

func (s *Store) ensureLayout() error {
	if err := ensurePrivateDir(s.root); err != nil {
		return err
	}
	if err := ensurePrivateDir(filepath.Join(s.root, snapshotsDir)); err != nil {
		return err
	}
	metadataPath := filepath.Join(s.root, metadataName)
	metadata := Metadata{FormatName: FormatName, FormatVersion: FormatVersion, Retention: s.retention}
	document, err := json.MarshalIndent(metadata, "", "  ")
	if err != nil {
		return err
	}
	document = append(document, '\n')
	if _, err := os.Lstat(metadataPath); errors.Is(err, os.ErrNotExist) {
		if err := atomicInstall(metadataPath, document, nil); err != nil {
			return fmt.Errorf("install store metadata: %w", err)
		}
	} else if err != nil {
		return fmt.Errorf("inspect store metadata: %w", err)
	}
	return s.validateLayout()
}

func (s *Store) validateLayout() error {
	if err := validatePrivateDir(s.root); err != nil {
		return err
	}
	if err := validatePrivateDir(filepath.Join(s.root, snapshotsDir)); err != nil {
		return err
	}
	path := filepath.Join(s.root, metadataName)
	info, err := os.Lstat(path)
	if err != nil {
		return err
	}
	if !info.Mode().IsRegular() || info.Mode().Perm()&0o077 != 0 || info.Size() > 4096 {
		return fmt.Errorf("%w: invalid store metadata file", ErrUnsafePath)
	}
	document, err := os.ReadFile(path)
	if err != nil {
		return err
	}
	if err := rejectDuplicateKeys(document); err != nil {
		return fmt.Errorf("%w: metadata: %v", ErrCorrupt, err)
	}
	var metadata Metadata
	decoder := json.NewDecoder(bytes.NewReader(document))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&metadata); err != nil {
		return fmt.Errorf("%w: metadata: %v", ErrCorrupt, err)
	}
	if err := requireJSONEOF(decoder); err != nil {
		return fmt.Errorf("%w: metadata: %v", ErrCorrupt, err)
	}
	if metadata.FormatName != FormatName || metadata.FormatVersion != FormatVersion {
		return ErrIncompatible
	}
	if metadata.Retention != s.retention {
		return fmt.Errorf("configured retention %d differs from store retention %d", s.retention, metadata.Retention)
	}
	return nil
}

func (s *Store) listUnlocked() ([]string, error) {
	entries, err := os.ReadDir(filepath.Join(s.root, snapshotsDir))
	if err != nil {
		return nil, err
	}
	names := make([]string, 0, len(entries))
	for _, entry := range entries {
		name := entry.Name()
		if strings.HasPrefix(name, ".") {
			return nil, fmt.Errorf("%w: unexpected store artifact %s", ErrCorrupt, name)
		}
		if !strings.HasSuffix(name, ".json") {
			return nil, fmt.Errorf("%w: unexpected snapshot file %s", ErrCorrupt, name)
		}
		info, err := os.Lstat(filepath.Join(s.root, snapshotsDir, name))
		if err != nil {
			return nil, err
		}
		if !info.Mode().IsRegular() || info.Mode().Perm()&0o077 != 0 {
			return nil, fmt.Errorf("%w: unsafe snapshot file %s", ErrUnsafePath, name)
		}
		names = append(names, name)
	}
	sort.Strings(names)
	return names, nil
}

func (s *Store) lock() (func(), error) {
	path := filepath.Join(s.root, lockName)
	file, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_EXCL, 0o600)
	if err != nil {
		return nil, fmt.Errorf("acquire inventory store lock: %w", err)
	}
	if err := file.Close(); err != nil {
		_ = os.Remove(path)
		return nil, err
	}
	return func() {
		_ = os.Remove(path)
		_ = syncDir(s.root)
	}, nil
}

func ensurePrivateDir(path string) error {
	info, err := os.Lstat(path)
	if errors.Is(err, os.ErrNotExist) {
		if err := os.Mkdir(path, 0o700); err != nil {
			return fmt.Errorf("create inventory store directory: %w", err)
		}
		return syncDir(filepath.Dir(path))
	}
	if err != nil {
		return err
	}
	return validateDirInfo(path, info)
}

func validatePrivateDir(path string) error {
	info, err := os.Lstat(path)
	if err != nil {
		return err
	}
	return validateDirInfo(path, info)
}

func validateDirInfo(path string, info os.FileInfo) error {
	if !info.IsDir() || info.Mode()&os.ModeSymlink != 0 || info.Mode().Perm()&0o077 != 0 {
		return fmt.Errorf("%w: directory %s must be a private real directory", ErrUnsafePath, path)
	}
	return nil
}

func rejectSymlinkComponents(path string) error {
	for current := path; ; current = filepath.Dir(current) {
		info, err := os.Lstat(current)
		if err == nil && info.Mode()&os.ModeSymlink != 0 {
			return fmt.Errorf("%w: symlink path component %s", ErrUnsafePath, current)
		}
		if err != nil && !errors.Is(err, os.ErrNotExist) {
			return fmt.Errorf("inspect inventory store path: %w", err)
		}
		parent := filepath.Dir(current)
		if parent == current {
			return nil
		}
	}
}

func atomicInstall(target string, document []byte, beforeInstall func() error) error {
	dir := filepath.Dir(target)
	file, err := os.CreateTemp(dir, ".tmp-inventory-*")
	if err != nil {
		return fmt.Errorf("create temporary inventory: %w", err)
	}
	temp := file.Name()
	defer os.Remove(temp)
	if err := file.Chmod(0o600); err != nil {
		file.Close()
		return err
	}
	if _, err := file.Write(document); err != nil {
		file.Close()
		return fmt.Errorf("write temporary inventory: %w", err)
	}
	if err := file.Sync(); err != nil {
		file.Close()
		return fmt.Errorf("sync temporary inventory: %w", err)
	}
	if err := file.Close(); err != nil {
		return fmt.Errorf("close temporary inventory: %w", err)
	}
	if beforeInstall != nil {
		if err := beforeInstall(); err != nil {
			return fmt.Errorf("before inventory install: %w", err)
		}
	}
	if err := os.Link(temp, target); err != nil {
		if errors.Is(err, os.ErrExist) {
			return ErrConflict
		}
		return fmt.Errorf("install inventory: %w", err)
	}
	if err := syncDir(dir); err != nil {
		_ = os.Remove(target)
		return fmt.Errorf("sync installed inventory: %w", err)
	}
	return nil
}

func syncDir(path string) error {
	dir, err := os.Open(path)
	if err != nil {
		return err
	}
	defer dir.Close()
	return dir.Sync()
}

func requireJSONEOF(decoder *json.Decoder) error {
	var extra any
	if err := decoder.Decode(&extra); !errors.Is(err, io.EOF) {
		if err == nil {
			return errors.New("multiple JSON values")
		}
		return err
	}
	return nil
}

func rejectDuplicateKeys(document []byte) error {
	decoder := json.NewDecoder(bufio.NewReader(bytes.NewReader(document)))
	var walk func() error
	walk = func() error {
		token, err := decoder.Token()
		if err != nil {
			return err
		}
		delim, ok := token.(json.Delim)
		if !ok {
			return nil
		}
		switch delim {
		case '{':
			seen := map[string]struct{}{}
			for decoder.More() {
				keyToken, err := decoder.Token()
				if err != nil {
					return err
				}
				key, ok := keyToken.(string)
				if !ok {
					return errors.New("object key is not a string")
				}
				if _, exists := seen[key]; exists {
					return fmt.Errorf("duplicate object key %q", key)
				}
				seen[key] = struct{}{}
				if err := walk(); err != nil {
					return err
				}
			}
			end, err := decoder.Token()
			if err != nil || end != json.Delim('}') {
				return errors.New("unterminated object")
			}
		case '[':
			for decoder.More() {
				if err := walk(); err != nil {
					return err
				}
			}
			end, err := decoder.Token()
			if err != nil || end != json.Delim(']') {
				return errors.New("unterminated array")
			}
		default:
			return errors.New("unexpected closing delimiter")
		}
		return nil
	}
	if err := walk(); err != nil {
		return err
	}
	if _, err := decoder.Token(); !errors.Is(err, io.EOF) {
		if err == nil {
			return errors.New("multiple JSON values")
		}
		return err
	}
	return nil
}
