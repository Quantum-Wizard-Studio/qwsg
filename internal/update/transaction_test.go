package update

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"
)

func TestApplyAndRollbackPreserveUserOwnedData(t *testing.T) {
	root := t.TempDir()
	pkg := filepath.Join(root, "pkg")
	dest := filepath.Join(root, "dest")
	backup := filepath.Join(root, "rollback")
	files := map[string]string{"bin/qwsg": "new-binary", "lib/systemd/user/qwsg-guardian.service": "new-unit", "README.md": "new-readme", "INSTALL.md": "new-install", "LICENSE": "license", "CHANGELOG.md": "changes", "qwsg-config.json": "example", "RELEASE.json": `{"Schema":"qwsg.release/1","Version":"1.2.0-rc.1","Commit":"1111111111111111111111111111111111111111","Built":"2026-08-26T00:00:00Z","Platform":"linux-amd64"}`, "docs/OPERATIONS.md": "ops"}
	var manifest string
	for rel, body := range files {
		writeTestFile(t, filepath.Join(pkg, rel), body)
		manifest += fmt.Sprintf("%s  %s\n", bytesSHA([]byte(body)), rel)
	}
	writeTestFile(t, filepath.Join(pkg, "MANIFEST.sha256"), manifest)
	for rel := range files {
		d, ok := destination(rel)
		if !ok {
			continue
		}
		if rel == "RELEASE.json" {
			continue
		}
		writeTestFile(t, filepath.Join(dest, d), "old:"+rel)
	}
	userConfig := filepath.Join(dest, "home/user/.config/qwsg/config.json")
	writeTestFile(t, userConfig, "private-config")
	tx, err := Apply(pkg, dest, backup, "1.1.0")
	if err != nil {
		t.Fatal(err)
	}
	if !tx.Complete || tx.ToVersion != "1.2.0-rc.1" {
		t.Fatalf("bad transaction: %+v", tx)
	}
	if got, _ := os.ReadFile(userConfig); string(got) != "private-config" {
		t.Fatal("user data changed")
	}
	if err = Rollback(dest, backup); err != nil {
		t.Fatal(err)
	}
	if _, err = os.Stat(filepath.Join(dest, "usr/local/share/doc/qwsg/RELEASE.json")); !os.IsNotExist(err) {
		t.Fatal("new artifact survived rollback")
	}
	if got, _ := os.ReadFile(filepath.Join(dest, "usr/local/bin/qwsg")); string(got) != "old:bin/qwsg" {
		t.Fatalf("binary not restored: %q", got)
	}
	if got, _ := os.ReadFile(userConfig); string(got) != "private-config" {
		t.Fatal("user data changed by rollback")
	}
}

func TestRollbackRefusesTamperedBackup(t *testing.T) {
	root := t.TempDir()
	backup := filepath.Join(root, "rollback")
	if err := os.MkdirAll(filepath.Join(backup, "files/usr/local/bin"), 0700); err != nil {
		t.Fatal(err)
	}
	tx := Transaction{Schema: "qwsg.update-transaction/1", Complete: true, Files: []InstalledFile{{Destination: "usr/local/bin/qwsg", Backup: "files/usr/local/bin/qwsg", SHA256: bytesSHA([]byte("old")), Mode: 0755, Existed: true}}}
	if err := writeTransaction(backup, tx); err != nil {
		t.Fatal(err)
	}
	writeTestFile(t, filepath.Join(backup, "files/usr/local/bin/qwsg"), "tampered")
	if err := Rollback(filepath.Join(root, "dest"), backup); err == nil {
		t.Fatal("tampered backup accepted")
	}
}

func TestApplyAutomaticallyRestoresAfterMutationFailure(t *testing.T) {
	root := t.TempDir()
	pkg := filepath.Join(root, "pkg")
	dest := filepath.Join(root, "dest")
	backup := filepath.Join(root, "rollback")
	files := map[string]string{"bin/qwsg": "new", "lib/systemd/user/qwsg-guardian.service": "unit", "README.md": "readme", "INSTALL.md": "install", "LICENSE": "license", "CHANGELOG.md": "changes", "qwsg-config.json": "config", "RELEASE.json": `{"Schema":"qwsg.release/1","Version":"1.2.0-rc.1","Commit":"1111111111111111111111111111111111111111","Built":"2026-08-26T00:00:00Z","Platform":"linux-amd64"}`, "docs/OPERATIONS.md": "ops"}
	manifest := ""
	for rel, body := range files {
		writeTestFile(t, filepath.Join(pkg, rel), body)
		manifest += fmt.Sprintf("%s  %s\n", bytesSHA([]byte(body)), rel)
		if d, ok := destination(rel); ok && rel != "RELEASE.json" {
			writeTestFile(t, filepath.Join(dest, d), "old:"+rel)
		}
	}
	writeTestFile(t, filepath.Join(pkg, "MANIFEST.sha256"), manifest)
	if err := os.Remove(filepath.Join(pkg, "README.md")); err != nil {
		t.Fatal(err)
	}
	if _, err := Apply(pkg, dest, backup, "1.1.0"); err == nil {
		t.Fatal("broken package applied")
	}
	got, err := os.ReadFile(filepath.Join(dest, "usr/local/bin/qwsg"))
	if err != nil || string(got) != "old:bin/qwsg" {
		t.Fatalf("automatic restore failed: %q %v", got, err)
	}
}

func writeTestFile(t *testing.T, path, body string) {
	t.Helper()
	if err := os.MkdirAll(filepath.Dir(path), 0755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(path, []byte(body), 0644); err != nil {
		t.Fatal(err)
	}
}
