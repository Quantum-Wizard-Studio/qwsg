package scheduler

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
)

const (
	stateFileName = "scheduler-state.json"
	lockFileName  = "scheduler.lock"
)

type stateEnvelope struct {
	SchemaName    string `json:"schema_name"`
	SchemaVersion string `json:"schema_version"`
	PayloadSHA256 string `json:"payload_sha256"`
	State         State  `json:"state"`
}

type StateStore interface {
	Load() (State, error)
	Save(State) error
}

type FileStore struct{ directory string }

func OpenFileStore(directory string) (*FileStore, error) {
	if directory == "" || !filepath.IsAbs(directory) {
		return nil, fmt.Errorf("scheduler state directory must be an explicit absolute path")
	}
	if err := os.MkdirAll(directory, 0700); err != nil {
		return nil, fmt.Errorf("scheduler state directory unavailable")
	}
	if err := os.Chmod(directory, 0700); err != nil {
		return nil, fmt.Errorf("scheduler state directory permissions unavailable")
	}
	return &FileStore{directory: directory}, nil
}

func (store *FileStore) Load() (State, error) {
	data, err := os.ReadFile(filepath.Join(store.directory, stateFileName))
	if err != nil {
		return State{}, err
	}
	var envelope stateEnvelope
	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&envelope); err != nil {
		return State{}, fmt.Errorf("scheduler state is malformed")
	}
	var extra any
	if err := decoder.Decode(&extra); err != io.EOF {
		return State{}, fmt.Errorf("scheduler state contains trailing data")
	}
	if envelope.SchemaName != "qwsg.scheduler-state-envelope" || envelope.SchemaVersion != SchemaVersion {
		return State{}, fmt.Errorf("unsupported scheduler state envelope")
	}
	payload, err := json.Marshal(envelope.State)
	if err != nil {
		return State{}, err
	}
	sum := sha256.Sum256(payload)
	if envelope.PayloadSHA256 != hex.EncodeToString(sum[:]) {
		return State{}, fmt.Errorf("scheduler state integrity check failed")
	}
	if err := ValidateState(envelope.State); err != nil {
		return State{}, err
	}
	return envelope.State, nil
}

func (store *FileStore) Save(state State) error {
	if err := ValidateState(state); err != nil {
		return err
	}
	payload, err := json.Marshal(state)
	if err != nil {
		return err
	}
	sum := sha256.Sum256(payload)
	envelope := stateEnvelope{SchemaName: "qwsg.scheduler-state-envelope", SchemaVersion: SchemaVersion, PayloadSHA256: hex.EncodeToString(sum[:]), State: state}
	data, err := json.Marshal(envelope)
	if err != nil {
		return err
	}
	temporary, err := os.CreateTemp(store.directory, ".scheduler-state-*")
	if err != nil {
		return fmt.Errorf("scheduler state temporary file unavailable")
	}
	name := temporary.Name()
	installed := false
	defer func() {
		if !installed {
			_ = os.Remove(name)
		}
	}()
	if err := temporary.Chmod(0600); err != nil {
		_ = temporary.Close()
		return err
	}
	if _, err := temporary.Write(data); err != nil {
		_ = temporary.Close()
		return err
	}
	if err := temporary.Sync(); err != nil {
		_ = temporary.Close()
		return err
	}
	if err := temporary.Close(); err != nil {
		return err
	}
	if err := os.Rename(name, filepath.Join(store.directory, stateFileName)); err != nil {
		return fmt.Errorf("scheduler state atomic replacement failed")
	}
	installed = true
	directory, err := os.Open(store.directory)
	if err == nil {
		err = directory.Sync()
		_ = directory.Close()
	}
	return err
}

type Lock interface{ Release() error }
type Locker interface{ Acquire(string) (Lock, error) }

var ErrLockContended = errors.New("scheduler lock is already held")

type FileLocker struct{ directory string }

func NewFileLocker(directory string) (*FileLocker, error) {
	if directory == "" || !filepath.IsAbs(directory) {
		return nil, fmt.Errorf("scheduler lock directory must be explicit")
	}
	if err := os.MkdirAll(directory, 0700); err != nil {
		return nil, err
	}
	return &FileLocker{directory: directory}, nil
}

type fileLock struct {
	file     *os.File
	released bool
}

func (locker *FileLocker) Acquire(owner string) (Lock, error) {
	if !idPattern.MatchString(owner) {
		return nil, fmt.Errorf("invalid scheduler lock owner")
	}
	path := filepath.Join(locker.directory, lockFileName)
	file, err := os.OpenFile(path, os.O_CREATE|os.O_RDWR, 0600)
	if err != nil {
		return nil, fmt.Errorf("scheduler lock unavailable")
	}
	if err := syscall.Flock(int(file.Fd()), syscall.LOCK_EX|syscall.LOCK_NB); err != nil {
		_ = file.Close()
		if errors.Is(err, syscall.EWOULDBLOCK) || errors.Is(err, syscall.EAGAIN) {
			return nil, ErrLockContended
		}
		return nil, fmt.Errorf("scheduler lock acquisition failed")
	}
	if err := file.Truncate(0); err == nil {
		_, err = file.WriteString(owner + "\n")
	}
	if err != nil {
		_ = syscall.Flock(int(file.Fd()), syscall.LOCK_UN)
		_ = file.Close()
		return nil, fmt.Errorf("scheduler lock evidence failed")
	}
	_ = file.Sync()
	return &fileLock{file: file}, nil
}
func (lock *fileLock) Release() error {
	if lock == nil || lock.file == nil || lock.released {
		return fmt.Errorf("scheduler lock is not active")
	}
	lock.released = true
	unlock := syscall.Flock(int(lock.file.Fd()), syscall.LOCK_UN)
	closeErr := lock.file.Close()
	if unlock != nil {
		return unlock
	}
	return closeErr
}
