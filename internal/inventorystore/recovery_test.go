package inventorystore

import (
	"bytes"
	"errors"
	"os"
	"os/exec"
	"path/filepath"
	"reflect"
	"strings"
	"testing"
	"time"
)

var recoveryTime = time.Date(2026, 9, 21, 12, 0, 0, 0, time.UTC)

func saveRecovery(t *testing.T, s *Store, id string, offset int) string {
	t.Helper()
	name, err := s.Save(fixtureSnapshot(id, recoveryTime.Add(time.Duration(offset)*time.Second), false))
	if err != nil {
		t.Fatal(err)
	}
	return name
}
func snapshotPath(s *Store, name string) string { return filepath.Join(s.root, snapshotsDir, name) }
func retireFixture(t *testing.T, s *Store, name string) {
	t.Helper()
	if err := os.Rename(snapshotPath(s, name), snapshotPath(s, ".retire-"+name)); err != nil {
		t.Fatal(err)
	}
}
func assertLatest(t *testing.T, s *Store, id string) {
	t.Helper()
	got, _, err := s.LoadLatest()
	if err != nil || got.SnapshotID != id {
		t.Fatalf("latest=%s want=%s err=%v", got.SnapshotID, id, err)
	}
}
func assertNoTransition(t *testing.T, s *Store) {
	t.Helper()
	entries, err := os.ReadDir(filepath.Join(s.root, snapshotsDir))
	if err != nil {
		t.Fatal(err)
	}
	for _, e := range entries {
		if strings.HasPrefix(e.Name(), ".") {
			t.Fatalf("artifact remains: %s", e.Name())
		}
	}
}

// The child exits at the existing actual pre-install hook, bypassing all Go
// defers, while holding the lock and a synced temporary plus retired A.
func TestInterruptedWriterProcess(t *testing.T) {
	if root := os.Getenv("QWSG_TEST_INTERRUPTED_STORE"); root != "" {
		s, err := Open(root, 1)
		if err != nil {
			os.Exit(81)
		}
		s.beforeInstall = func() error { os.Exit(73); return nil }
		_, _ = s.Save(fixtureSnapshot("B", recoveryTime.Add(time.Second), false))
		os.Exit(82)
	}
	s := openFixtureStore(t, 1)
	a := saveRecovery(t, s, "A", 0)
	before, err := os.ReadFile(snapshotPath(s, a))
	if err != nil {
		t.Fatal(err)
	}
	cmd := exec.Command(os.Args[0], "-test.run=^TestInterruptedWriterProcess$")
	cmd.Env = append(os.Environ(), "QWSG_TEST_INTERRUPTED_STORE="+s.root)
	out, err := cmd.CombinedOutput()
	var exit *exec.ExitError
	if !errors.As(err, &exit) || exit.ExitCode() != 73 {
		t.Fatalf("child: %v %s", err, out)
	}
	if _, err := os.Stat(snapshotPath(s, ".retire-"+a)); err != nil {
		t.Fatalf("not interrupted at retention: %v", err)
	}
	for i := 0; i < 3; i++ {
		assertLatest(t, s, "A")
		assertNoTransition(t, s)
	}
	after, err := os.ReadFile(snapshotPath(s, a))
	if err != nil || !bytes.Equal(before, after) {
		t.Fatal("prior committed bytes changed")
	}
	saveRecovery(t, s, "B", 1)
	assertLatest(t, s, "B")
}

func TestLiveLockAndPermanentLegacyExclusion(t *testing.T) {
	s := openFixtureStore(t, 2)
	saveRecovery(t, s, "A", 0)
	unlock, err := s.lock()
	if err != nil {
		t.Fatal(err)
	}
	other, err := Open(s.root, 2)
	if err != nil {
		t.Fatal(err)
	}
	if _, err := other.Save(fixtureSnapshot("B", recoveryTime.Add(time.Second), false)); !errors.Is(err, ErrLocked) {
		t.Fatalf("live writer stolen: %v", err)
	}
	if _, err := other.List(); !errors.Is(err, ErrLocked) {
		t.Fatalf("reader raced writer: %v", err)
	}
	f, err := os.OpenFile(filepath.Join(s.root, lockName), os.O_CREATE|os.O_EXCL|os.O_WRONLY, 0600)
	if err == nil {
		f.Close()
		t.Fatal("older writer could start")
	}
	unlock()
	saveRecovery(t, other, "B", 1)
	assertLatest(t, other, "B")
}

func TestLegacyAndMalformedLockRequireExplicitRecovery(t *testing.T) {
	for _, content := range []string{"", "unknown owner"} {
		t.Run(content, func(t *testing.T) {
			s := openFixtureStore(t, 2)
			if err := s.ensureLayout(); err != nil {
				t.Fatal(err)
			}
			path := filepath.Join(s.root, lockName)
			if err := os.WriteFile(path, []byte(content), 0600); err != nil {
				t.Fatal(err)
			}
			for i := 0; i < 2; i++ {
				_, err := s.Save(fixtureSnapshot("A", recoveryTime, false))
				if !errors.Is(err, ErrCorrupt) || !strings.Contains(err.Error(), "stop all older QWSG writers") {
					t.Fatalf("ambiguous lock: %v", err)
				}
			}
			got, _ := os.ReadFile(path)
			if string(got) != content {
				t.Fatal("legacy lock modified")
			}
			// Test operator has established that no old process exists; preserve the
			// artifact outside the store as instructed, never age-based deletion.
			if err := os.Rename(path, filepath.Join(t.TempDir(), "preserved-lock")); err != nil {
				t.Fatal(err)
			}
			saveRecovery(t, s, "A", 0)
			assertLatest(t, s, "A")
		})
	}
}

func TestDurableInstallWithInterruptedRetirementCleanup(t *testing.T) {
	for _, retention := range []int{1, 2} {
		t.Run(string(rune('0'+retention)), func(t *testing.T) {
			s := openFixtureStore(t, retention)
			a := saveRecovery(t, s, "A", 0)
			if retention == 2 {
				saveRecovery(t, s, "middle", 1)
			}
			fixture := openFixtureStore(t, 2)
			b := saveRecovery(t, fixture, "B", 2)
			document, err := os.ReadFile(snapshotPath(fixture, b))
			if err != nil {
				t.Fatal(err)
			}
			unlock, err := s.lock()
			if err != nil {
				t.Fatal(err)
			}
			retireFixture(t, s, a)
			if err := atomicInstall(snapshotPath(s, b), document, nil, nil); err != nil {
				t.Fatal(err)
			}
			// A crash after install but before temporary unlink can leave both links.
			if err := os.Link(snapshotPath(s, b), snapshotPath(s, ".tmp-inventory-123")); err != nil {
				t.Fatal(err)
			}
			unlock()
			for i := 0; i < 3; i++ {
				assertLatest(t, s, "B")
				assertNoTransition(t, s)
			}
			got, err := os.ReadFile(snapshotPath(s, b))
			if err != nil || !bytes.Equal(got, document) {
				t.Fatal("committed B reverted/changed")
			}
			if _, err := os.Stat(snapshotPath(s, a)); !errors.Is(err, os.ErrNotExist) {
				t.Fatal("retired A became current")
			}
			names, err := s.List()
			if err != nil || len(names) != retention {
				t.Fatalf("retention: %v %v", names, err)
			}
		})
	}
}

func TestAmbiguousRecoveryPreservesEveryArtifact(t *testing.T) {
	for _, mode := range []string{"conflicting", "multiple", "wrong-count", "malformed-retired", "malformed-committed", "unknown", "unsafe-temp"} {
		t.Run(mode, func(t *testing.T) {
			s := openFixtureStore(t, 2)
			a := saveRecovery(t, s, "A", 0)
			b := saveRecovery(t, s, "B", 1)
			retireFixture(t, s, a)
			switch mode {
			case "conflicting":
				if err := os.Link(snapshotPath(s, ".retire-"+a), snapshotPath(s, a)); err != nil {
					t.Fatal(err)
				}
			case "multiple":
				retireFixture(t, s, b)
			case "wrong-count":
				if err := os.Remove(snapshotPath(s, b)); err != nil {
					t.Fatal(err)
				}
			case "malformed-retired":
				if err := os.WriteFile(snapshotPath(s, ".retire-"+a), []byte(`{"format_name":`), 0600); err != nil {
					t.Fatal(err)
				}
			case "malformed-committed":
				if err := os.WriteFile(snapshotPath(s, b), []byte(`{"format_name":`), 0600); err != nil {
					t.Fatal(err)
				}
			case "unknown":
				if err := os.WriteFile(snapshotPath(s, ".unrecognized"), []byte("evidence"), 0600); err != nil {
					t.Fatal(err)
				}
			case "unsafe-temp":
				if err := os.Symlink(snapshotPath(s, b), snapshotPath(s, ".tmp-inventory-123")); err != nil {
					t.Fatal(err)
				}
			}
			before := captureDirectory(t, filepath.Join(s.root, snapshotsDir))
			for i := 0; i < 2; i++ {
				if _, err := s.List(); err == nil {
					t.Fatal("ambiguous/incomplete recovery accepted")
				}
			}
			if !reflect.DeepEqual(before, captureDirectory(t, filepath.Join(s.root, snapshotsDir))) {
				t.Fatal("ambiguous evidence mutated")
			}
		})
	}
}
func captureDirectory(t *testing.T, dir string) map[string]string {
	t.Helper()
	entries, err := os.ReadDir(dir)
	if err != nil {
		t.Fatal(err)
	}
	out := map[string]string{}
	for _, e := range entries {
		path := filepath.Join(dir, e.Name())
		if e.Type()&os.ModeSymlink != 0 {
			v, err := os.Readlink(path)
			if err != nil {
				t.Fatal(err)
			}
			out[e.Name()] = "link:" + v
		} else {
			v, err := os.ReadFile(path)
			if err != nil {
				t.Fatal(err)
			}
			out[e.Name()] = string(v)
		}
	}
	return out
}

func TestAbandonedTemporaryNeverPromoted(t *testing.T) {
	s := openFixtureStore(t, 2)
	saveRecovery(t, s, "A", 0)
	for _, dir := range []string{s.root, filepath.Join(s.root, snapshotsDir)} {
		if err := os.WriteFile(filepath.Join(dir, ".tmp-inventory-123"), []byte(`{"incomplete":`), 0600); err != nil {
			t.Fatal(err)
		}
	}
	for i := 0; i < 2; i++ {
		assertLatest(t, s, "A")
		assertNoTransition(t, s)
	}
	if _, err := os.Stat(filepath.Join(s.root, ".tmp-inventory-123")); !errors.Is(err, os.ErrNotExist) {
		t.Fatal("metadata temp remains")
	}
}

func TestRecoveryDirectoryBound(t *testing.T) {
	s := openFixtureStore(t, 1)
	saveRecovery(t, s, "A", 0)
	for _, name := range []string{"1", "2", "3"} {
		if err := os.WriteFile(snapshotPath(s, ".tmp-inventory-"+name), []byte("partial"), 0600); err != nil {
			t.Fatal(err)
		}
	}
	before := captureDirectory(t, filepath.Join(s.root, snapshotsDir))
	if _, err := s.List(); !errors.Is(err, ErrCorrupt) {
		t.Fatalf("unbounded artifacts accepted: %v", err)
	}
	if !reflect.DeepEqual(before, captureDirectory(t, filepath.Join(s.root, snapshotsDir))) {
		t.Fatal("over-limit store mutated")
	}
}

func TestDirectorySyncFailureDoesNotRevertInstalledEvidence(t *testing.T) {
	s := openFixtureStore(t, 1)
	a := saveRecovery(t, s, "A", 0)
	injected := errors.New("directory sync fault")
	s.beforeDirectorySync = func() error { return injected }
	if _, err := s.Save(fixtureSnapshot("B", recoveryTime.Add(time.Second), false)); !errors.Is(err, injected) {
		t.Fatalf("false success: %v", err)
	}
	// Both objects survive uncertainty. Recovery syncs the observed directory
	// before finishing retention; it must not arbitrarily undo installed B.
	if _, err := os.Stat(snapshotPath(s, ".retire-"+a)); err != nil {
		t.Fatal("A destroyed before durable boundary")
	}
	b := snapshotName(fixtureSnapshot("B", recoveryTime.Add(time.Second), false))
	if _, err := os.Stat(snapshotPath(s, b)); err != nil {
		t.Fatal("B reverted after link")
	}
	s.beforeDirectorySync = nil
	assertLatest(t, s, "B")
	assertNoTransition(t, s)
}

func TestLegacyEvidenceReadableWithoutRewriting(t *testing.T) {
	s := openFixtureStore(t, 2)
	a := saveRecovery(t, s, "A", 0)
	before := captureDirectory(t, filepath.Join(s.root, snapshotsDir))
	// Legacy 1.0 has the same metadata/envelope and no permanent lock files.
	for _, name := range []string{kernelLockName, lockName} {
		if err := os.Remove(filepath.Join(s.root, name)); err != nil {
			t.Fatal(err)
		}
	}
	if _, err := s.Load(a); err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(before, captureDirectory(t, filepath.Join(s.root, snapshotsDir))) {
		t.Fatal("historical evidence rewritten")
	}
	saveRecovery(t, s, "B", 1)
}

func TestMissingMetadataPreservesExistingEvidence(t *testing.T) {
	s := openFixtureStore(t, 1)
	saveRecovery(t, s, "A", 0)
	if err := os.Remove(filepath.Join(s.root, metadataName)); err != nil {
		t.Fatal(err)
	}
	before := captureDirectory(t, filepath.Join(s.root, snapshotsDir))
	if _, err := s.Save(fixtureSnapshot("B", recoveryTime.Add(time.Second), false)); !errors.Is(err, ErrCorrupt) {
		t.Fatalf("metadata silently reset: %v", err)
	}
	if !reflect.DeepEqual(before, captureDirectory(t, filepath.Join(s.root, snapshotsDir))) {
		t.Fatal("evidence modified")
	}
}

func TestOversizedInventoryAndMetadataRejected(t *testing.T) {
	for _, metadata := range []bool{false, true} {
		t.Run(map[bool]string{false: "snapshot", true: "metadata"}[metadata], func(t *testing.T) {
			s := openFixtureStore(t, 1)
			a := saveRecovery(t, s, "A", 0)
			path := snapshotPath(s, a)
			size := int64(MaxDocumentSize + 1)
			if metadata {
				path = filepath.Join(s.root, metadataName)
				size = 4097
			}
			f, err := os.OpenFile(path, os.O_WRONLY|os.O_TRUNC, 0600)
			if err != nil {
				t.Fatal(err)
			}
			if err := f.Truncate(size); err != nil {
				t.Fatal(err)
			}
			f.Close()
			if _, err := s.Load(a); err == nil {
				t.Fatal("oversize accepted")
			}
		})
	}
}

func TestConcurrentSaveCannotEnterActiveTransaction(t *testing.T) {
	s := openFixtureStore(t, 1)
	saveRecovery(t, s, "A", 0)
	entered := make(chan struct{})
	release := make(chan struct{})
	done := make(chan error, 1)
	s.beforeInstall = func() error { close(entered); <-release; return nil }
	go func() { _, err := s.Save(fixtureSnapshot("B", recoveryTime.Add(time.Second), false)); done <- err }()
	<-entered
	other, err := Open(s.root, 1)
	if err != nil {
		close(release)
		t.Fatal(err)
	}
	_, collision := other.Save(fixtureSnapshot("C", recoveryTime.Add(2*time.Second), false))
	close(release)
	if err := <-done; err != nil {
		t.Fatal(err)
	}
	if !errors.Is(collision, ErrLocked) {
		t.Fatalf("active transaction not excluded: %v", collision)
	}
	assertLatest(t, other, "B")
}

func TestInterruptedInitializationRecoversOnlyKnownTemps(t *testing.T) {
	s := openFixtureStore(t, 1)
	if err := os.Mkdir(s.root, 0700); err != nil {
		t.Fatal(err)
	}
	if err := os.Mkdir(filepath.Join(s.root, snapshotsDir), 0700); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(s.root, ".tmp-inventory-123"), []byte(`{"format_name":`), 0600); err != nil {
		t.Fatal(err)
	}
	saveRecovery(t, s, "A", 0)
	assertLatest(t, s, "A")
	if _, err := os.Stat(filepath.Join(s.root, ".tmp-inventory-123")); !errors.Is(err, os.ErrNotExist) {
		t.Fatal("initialization temp remains")
	}
}

func TestUnsafeKernelLockNeverFollowedOrReset(t *testing.T) {
	for _, mode := range []string{"symlink", "content", "permissions"} {
		t.Run(mode, func(t *testing.T) {
			s := openFixtureStore(t, 1)
			saveRecovery(t, s, "A", 0)
			path := filepath.Join(s.root, kernelLockName)
			switch mode {
			case "symlink":
				if err := os.Remove(path); err != nil {
					t.Fatal(err)
				}
				if err := os.Symlink(filepath.Join(s.root, metadataName), path); err != nil {
					t.Fatal(err)
				}
			case "content":
				if err := os.WriteFile(path, []byte("unknown"), 0600); err != nil {
					t.Fatal(err)
				}
			case "permissions":
				if err := os.Chmod(path, 0644); err != nil {
					t.Fatal(err)
				}
			}
			before := captureDirectory(t, filepath.Join(s.root, snapshotsDir))
			if _, err := s.List(); err == nil {
				t.Fatal("unsafe lock accepted")
			}
			if !reflect.DeepEqual(before, captureDirectory(t, filepath.Join(s.root, snapshotsDir))) {
				t.Fatal("unsafe lock recovery mutated evidence")
			}
		})
	}
}
