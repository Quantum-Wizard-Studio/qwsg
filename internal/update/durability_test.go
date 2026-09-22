package update

import (
	"errors"
	"os"
	"path/filepath"
	"testing"
)

func TestPreparedJournalPrecedesMutation(t *testing.T) {
	old := beforeMutation
	defer func() { beforeMutation = old }()
	reached := false
	beforeMutation = func(tx Transaction) error {
		reached = true
		if !tx.Prepared || tx.MutationStarted || tx.Complete {
			t.Fatalf("unsafe prepare: %+v", tx)
		}
		return errors.New("interrupted after prepare")
	}
	// Existing transaction fixture exercises complete real backup preparation.
	t.Run("prepare", func(t *testing.T) {
		root := t.TempDir()
		pkg := filepath.Join(root, "pkg")
		dest := filepath.Join(root, "dest")
		backup := filepath.Join(root, "backup")
		files := map[string]string{"bin/qwsg": "new", "lib/systemd/user/qwsg-guardian.service": "unit", "README.md": "readme", "INSTALL.md": "install", "LICENSE": "license", "CHANGELOG.md": "changes", "qwsg-config.json": "config", "RELEASE.json": `{"Version":"1.9.0"}`}
		manifest := ""
		for rel, body := range files {
			writeTestFile(t, filepath.Join(pkg, rel), body)
			manifest += bytesSHA([]byte(body)) + "  " + rel + "\n"
			d, _ := destination(rel)
			writeTestFile(t, filepath.Join(dest, d), "old:"+rel)
		}
		writeTestFile(t, filepath.Join(pkg, "MANIFEST.sha256"), manifest)
		tx, err := Apply(pkg, dest, backup, "1.3.1")
		if err == nil || tx.MutationStarted {
			t.Fatalf("%+v %v", tx, err)
		}
		saved, err := ReadTransaction(backup)
		if err != nil || !saved.Prepared || len(saved.Files) != len(files) {
			t.Fatalf("%+v %v", saved, err)
		}
		for i := 0; i < 2; i++ {
			if err := Rollback(dest, backup); err != nil {
				t.Fatal(err)
			}
		}
		for rel := range files {
			d, _ := destination(rel)
			b, _ := os.ReadFile(filepath.Join(dest, d))
			if string(b) != "old:"+rel {
				t.Fatal("prepare changed destination")
			}
		}
	})
	if !reached {
		t.Fatal("prepare boundary not reached")
	}
}

func TestInvalidLaterRollbackSourceDoesNotMutate(t *testing.T) {
	root := t.TempDir()
	dest := filepath.Join(root, "dest")
	backup := filepath.Join(root, "backup")
	if err := os.Mkdir(backup, 0700); err != nil {
		t.Fatal(err)
	}
	tx := Transaction{Schema: "qwsg.update-transaction/1", Complete: true}
	for _, d := range []string{"usr/local/bin/qwsg", "usr/local/share/doc/qwsg/README.md"} {
		writeTestFile(t, filepath.Join(dest, d), "valid-current")
		writeTestFile(t, filepath.Join(backup, "files", d), "old")
		tx.Files = append(tx.Files, InstalledFile{Destination: d, Backup: "files/" + d, SHA256: bytesSHA([]byte("old")), Mode: 0644, Existed: true})
	}
	writeTestFile(t, filepath.Join(backup, tx.Files[1].Backup), "corrupt")
	if err := writeTransaction(backup, tx); err != nil {
		t.Fatal(err)
	}
	if err := Rollback(dest, backup); err == nil {
		t.Fatal("accepted corrupt rollback")
	}
	for _, f := range tx.Files {
		b, _ := os.ReadFile(filepath.Join(dest, f.Destination))
		if string(b) != "valid-current" {
			t.Fatal("partial rollback on invalid source")
		}
	}
}

func durabilityFixture(t *testing.T) (string, string, string) {
	t.Helper()
	root := t.TempDir()
	pkg := filepath.Join(root, "pkg")
	dest := filepath.Join(root, "dest")
	backup := filepath.Join(root, "backup")
	manifest := ""
	for _, rel := range []string{"bin/qwsg", "lib/systemd/user/qwsg-guardian.service", "README.md", "INSTALL.md", "LICENSE", "CHANGELOG.md", "qwsg-config.json", "RELEASE.json"} {
		body := "new:" + rel
		if rel == "RELEASE.json" {
			body = `{"Version":"1.9.0"}`
		}
		writeTestFile(t, filepath.Join(pkg, rel), body)
		manifest += bytesSHA([]byte(body)) + "  " + rel + "\n"
		d, _ := destination(rel)
		writeTestFile(t, filepath.Join(dest, d), "old:"+rel)
	}
	writeTestFile(t, filepath.Join(pkg, "MANIFEST.sha256"), manifest)
	return pkg, dest, backup
}

func TestInterruptedMutationUsesPreparedRollbackJournal(t *testing.T) {
	pkg, dest, backup := durabilityFixture(t)
	old := afterMutation
	defer func() { afterMutation = old }()
	afterMutation = func() { panic("process loss after first replacement") }
	func() {
		defer func() {
			if recover() == nil {
				t.Fatal("missing boundary")
			}
		}()
		_, _ = Apply(pkg, dest, backup, "1.3.1")
	}()
	tx, err := ReadTransaction(backup)
	if err != nil || !tx.Prepared || tx.Complete {
		t.Fatalf("%+v %v", tx, err)
	}
	b, _ := os.ReadFile(filepath.Join(dest, "usr/local/bin/qwsg"))
	if string(b) != "new:bin/qwsg" {
		t.Fatal("interruption did not reach actual mutation")
	}
	for i := 0; i < 2; i++ {
		if err := Rollback(dest, backup); err != nil {
			t.Fatal(err)
		}
	}
	b, _ = os.ReadFile(filepath.Join(dest, "usr/local/bin/qwsg"))
	if string(b) != "old:bin/qwsg" {
		t.Fatal("interrupted mutation not restored")
	}
}

func TestJournalSyncFailureNeverCrossesMutationBoundary(t *testing.T) {
	pkg, dest, backup := durabilityFixture(t)
	old := syncDirectory
	defer func() { syncDirectory = old }()
	syncDirectory = func(path string) error {
		if path == backup {
			return errors.New("directory sync failed")
		}
		return old(path)
	}
	tx, err := Apply(pkg, dest, backup, "1.3.1")
	if err == nil || tx.MutationStarted {
		t.Fatalf("sync failure mutated installation %+v %v", tx, err)
	}
	b, _ := os.ReadFile(filepath.Join(dest, "usr/local/bin/qwsg"))
	if string(b) != "old:bin/qwsg" {
		t.Fatal("sync failure changed binary")
	}
	if _, err := os.Stat(filepath.Join(backup, "files/usr/local/bin/qwsg")); err != nil {
		t.Fatal("backup removed after durability uncertainty")
	}
}

func TestDestinationSyncFailurePreservesRecoverableBackup(t *testing.T) {
	pkg, dest, backup := durabilityFixture(t)
	old := syncDirectory
	defer func() { syncDirectory = old }()
	syncDirectory = func(path string) error {
		if path == filepath.Join(dest, "usr/local/bin") {
			return errors.New("destination sync failed")
		}
		return old(path)
	}
	tx, err := Apply(pkg, dest, backup, "1.3.1")
	if err == nil || !tx.MutationStarted || !tx.RollbackAttempted || tx.RollbackSucceeded {
		t.Fatalf("dishonest durability result %+v %v", tx, err)
	}
	syncDirectory = old
	if err := Rollback(dest, backup); err != nil {
		t.Fatal(err)
	}
}
