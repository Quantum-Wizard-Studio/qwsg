// Package guardian provides the narrow local operational boundary around the
// Canonical Runtime Service. It owns instance exclusion and restart handoff,
// but no recurrence or engineering decisions.
package guardian

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"syscall"
	"time"

	"quantumwizard.hu/qwsg/internal/alert"
	"quantumwizard.hu/qwsg/internal/notification"
	"quantumwizard.hu/qwsg/internal/runtime"
)

const (
	SchemaVersion = "1.0"
	ModelVersion  = "1.0"
	MaxSize       = 4 << 20
)

var (
	ErrActive         = errors.New("guardian is already active")
	ErrCheckpoint     = errors.New("guardian checkpoint invalid")
	ErrUnsafePath     = errors.New("guardian state path unsafe")
	ErrIncompatible   = errors.New("guardian checkpoint incompatible")
	ErrExitEvidence   = errors.New("guardian exit evidence invalid")
	ErrExitState      = errors.New("guardian exit state invalid")
	ErrExitCheckpoint = errors.New("guardian exit checkpoint unavailable")
	ErrExitCurrent    = errors.New("guardian exit current state unavailable")
)

type Checkpoint struct {
	SchemaName             string                  `json:"schema_name"`
	SchemaVersion          string                  `json:"schema_version"`
	ModelVersion           string                  `json:"model_version"`
	ServiceID              string                  `json:"service_id"`
	ConfigurationID        string                  `json:"configuration_id"`
	Generation             string                  `json:"generation"`
	Active                 bool                    `json:"active"`
	LastCompletedCycleID   string                  `json:"last_completed_cycle_id,omitempty"`
	LastCompletedAt        time.Time               `json:"last_completed_at,omitempty"`
	RuntimeState           runtime.State           `json:"runtime_state"`
	AlertState             alert.State             `json:"alert_state"`
	NotificationQueueState notification.QueueState `json:"notification_queue_state"`
}

type envelope struct {
	SchemaName    string     `json:"schema_name"`
	SchemaVersion string     `json:"schema_version"`
	PayloadSHA256 string     `json:"payload_sha256"`
	Checkpoint    Checkpoint `json:"checkpoint"`
}

type Store struct{ directory string }

func OpenStore(directory string) (*Store, error) {
	if !cleanAbsolute(directory) {
		return nil, ErrUnsafePath
	}
	if err := privateDirectory(directory); err != nil {
		return nil, err
	}
	return &Store{directory: directory}, nil
}

func (s *Store) Load() (Checkpoint, error) {
	document, err := os.ReadFile(filepath.Join(s.directory, "checkpoint.json"))
	if err != nil {
		return Checkpoint{}, err
	}
	if len(document) == 0 || len(document) > MaxSize {
		return Checkpoint{}, ErrCheckpoint
	}
	var value envelope
	decoder := json.NewDecoder(bytes.NewReader(document))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&value); err != nil {
		return Checkpoint{}, ErrCheckpoint
	}
	var extra any
	if err := decoder.Decode(&extra); err != io.EOF {
		return Checkpoint{}, ErrCheckpoint
	}
	if value.SchemaName != "qwsg.guardian-checkpoint-envelope" || value.SchemaVersion != SchemaVersion {
		return Checkpoint{}, ErrIncompatible
	}
	payload, _ := json.Marshal(value.Checkpoint)
	sum := sha256.Sum256(payload)
	if value.PayloadSHA256 != hex.EncodeToString(sum[:]) || ValidateCheckpoint(value.Checkpoint) != nil {
		return Checkpoint{}, ErrCheckpoint
	}
	return value.Checkpoint, nil
}

func (s *Store) Save(value Checkpoint) error {
	if err := ValidateCheckpoint(value); err != nil {
		return err
	}
	payload, _ := json.Marshal(value)
	sum := sha256.Sum256(payload)
	document, err := json.Marshal(envelope{SchemaName: "qwsg.guardian-checkpoint-envelope", SchemaVersion: SchemaVersion, PayloadSHA256: hex.EncodeToString(sum[:]), Checkpoint: value})
	if err != nil || len(document) > MaxSize {
		return ErrCheckpoint
	}
	temporary, err := os.CreateTemp(s.directory, ".checkpoint-*")
	if err != nil {
		return err
	}
	name := temporary.Name()
	installed := false
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
	if err == nil {
		err = os.Rename(name, filepath.Join(s.directory, "checkpoint.json"))
	}
	if err != nil {
		return err
	}
	installed = true
	directory, err := os.Open(s.directory)
	if err == nil {
		err = directory.Sync()
		_ = directory.Close()
	}
	return err
}

func ValidateCheckpoint(v Checkpoint) error {
	if v.SchemaName != "qwsg.guardian-checkpoint" || v.SchemaVersion != SchemaVersion || v.ModelVersion != ModelVersion || !token(v.ServiceID) || !token(v.ConfigurationID) || !token(v.Generation) {
		return ErrCheckpoint
	}
	if err := runtime.ValidateState(v.RuntimeState); err != nil {
		return ErrCheckpoint
	}
	if err := alert.ValidateState(v.AlertState); err != nil || v.AlertState.ConfigurationID != v.ConfigurationID {
		return ErrCheckpoint
	}
	if err := notification.ValidateQueue(v.NotificationQueueState); err != nil {
		return ErrCheckpoint
	}
	if (v.LastCompletedCycleID == "") != v.LastCompletedAt.IsZero() || (!v.LastCompletedAt.IsZero() && v.LastCompletedAt.Location() != time.UTC) {
		return ErrCheckpoint
	}
	return nil
}

type Lock struct {
	file     *os.File
	released bool
}

func Acquire(directory, owner string) (*Lock, error) {
	if !cleanAbsolute(directory) || !token(owner) {
		return nil, ErrUnsafePath
	}
	if err := privateDirectory(directory); err != nil {
		return nil, err
	}
	file, err := os.OpenFile(filepath.Join(directory, "operation.lock"), os.O_CREATE|os.O_RDWR, 0600)
	if err != nil {
		return nil, err
	}
	if err = syscall.Flock(int(file.Fd()), syscall.LOCK_EX|syscall.LOCK_NB); err != nil {
		_ = file.Close()
		if errors.Is(err, syscall.EWOULDBLOCK) || errors.Is(err, syscall.EAGAIN) {
			return nil, ErrActive
		}
		return nil, fmt.Errorf("guardian lock unavailable")
	}
	return &Lock{file: file}, nil
}

func (l *Lock) Release() error {
	if l == nil || l.file == nil || l.released {
		return fmt.Errorf("guardian lock inactive")
	}
	l.released = true
	err := syscall.Flock(int(l.file.Fd()), syscall.LOCK_UN)
	closeErr := l.file.Close()
	if err != nil {
		return err
	}
	return closeErr
}

func privateDirectory(path string) error {
	for current := path; current != filepath.Dir(current); current = filepath.Dir(current) {
		info, err := os.Lstat(current)
		if errors.Is(err, os.ErrNotExist) {
			continue
		}
		if err != nil || info.Mode()&os.ModeSymlink != 0 || !info.IsDir() {
			return ErrUnsafePath
		}
	}
	if err := os.MkdirAll(path, 0700); err != nil {
		return err
	}
	return os.Chmod(path, 0700)
}

func cleanAbsolute(path string) bool {
	return path != "" && filepath.IsAbs(path) && filepath.Clean(path) == path
}
func token(value string) bool {
	if value == "" || len(value) > 256 {
		return false
	}
	for _, r := range value {
		if !(r >= 'a' && r <= 'z') && !(r >= '0' && r <= '9') && r != '.' && r != '_' && r != ':' && r != '-' {
			return false
		}
	}
	return true
}
