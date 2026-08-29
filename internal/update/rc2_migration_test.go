package update

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"testing"
)

type rc2Fixture struct {
	Schema              string `json:"schema"`
	Version             string `json:"version"`
	ConfigurationSchema string `json:"configuration_schema"`
	GuardianSchema      string `json:"guardian_schema"`
	SchedulerSchema     string `json:"scheduler_schema"`
	OperatorState       string `json:"operator_state"`
	GuardianUnit        string `json:"guardian_unit"`
	Configuration       string `json:"configuration"`
	Credential          string `json:"credential"`
	State               string `json:"state"`
}

func TestRC2FixtureMigratesToRC7AndRollsBackExactly(t *testing.T) {
	data, err := os.ReadFile("testdata/rc2-installed-state.json")
	if err != nil {
		t.Fatal(err)
	}
	var fixture rc2Fixture
	if json.Unmarshal(data, &fixture) != nil || fixture.Schema != "qwsg.update-fixture/1" || fixture.Version != "1.2.0-rc.2" {
		t.Fatalf("invalid RC.2 fixture: %+v", fixture)
	}
	migration, err := PlanMigration(fixture.Version, "1.2.0-rc.7")
	if err != nil || migration.Validate() != nil || migration.ConfigurationSchema != fixture.ConfigurationSchema || migration.GuardianSchema != fixture.GuardianSchema || migration.SchedulerSchema != fixture.SchedulerSchema || migration.OperatorState != fixture.OperatorState {
		t.Fatalf("fixture compatibility mismatch: %+v %v", migration, err)
	}

	root := t.TempDir()
	pkg, dest, backup := filepath.Join(root, "pkg"), filepath.Join(root, "dest"), filepath.Join(root, "rollback")
	files := map[string]string{
		"bin/qwsg": "QWSG 1.2.0-rc.7\n", "lib/systemd/user/qwsg-guardian.service": "[Service]\nExecStart=/usr/local/bin/qwsg guardian run\nEnvironment=GOMEMLIMIT=64MiB\nMemoryMax=128M\nTasksMax=32\n",
		"README.md": "readme", "INSTALL.md": "install", "LICENSE": "license", "CHANGELOG.md": "changes", "qwsg-config.json": "example",
		"RELEASE.json":       `{"Schema":"qwsg.release/1","Version":"1.2.0-rc.7","Commit":"1111111111111111111111111111111111111111","Built":"2026-08-28T00:00:00Z","Platform":"linux-amd64"}`,
		"docs/OPERATIONS.md": "ops",
	}
	manifest := ""
	for rel, body := range files {
		writeTestFile(t, filepath.Join(pkg, rel), body)
		manifest += fmt.Sprintf("%s  %s\n", bytesSHA([]byte(body)), rel)
		if d, ok := destination(rel); ok && rel != "RELEASE.json" {
			old := "old:" + rel
			if rel == "bin/qwsg" {
				old = "QWSG 1.2.0-rc.2\n"
			} else if rel == "lib/systemd/user/qwsg-guardian.service" {
				old = fixture.GuardianUnit
			}
			writeTestFile(t, filepath.Join(dest, d), old)
		}
	}
	writeTestFile(t, filepath.Join(pkg, "MANIFEST.sha256"), manifest)
	preserved := map[string]string{
		"home/user/.config/qwsg/config.json":               fixture.Configuration,
		"home/user/.config/qwsg/credentials/smtp-password": fixture.Credential,
		"home/user/.local/state/qwsg/guardian.json":        fixture.State,
	}
	for rel, body := range preserved {
		writeTestFile(t, filepath.Join(dest, rel), body)
	}

	tx, err := Apply(pkg, dest, backup, fixture.Version)
	if err != nil || !tx.Complete || tx.FromVersion != fixture.Version || tx.ToVersion != "1.2.0-rc.7" {
		t.Fatalf("migration failed: %+v %v", tx, err)
	}
	for rel, want := range preserved {
		if got, _ := os.ReadFile(filepath.Join(dest, rel)); string(got) != want {
			t.Fatalf("preserved data changed: %s", rel)
		}
	}
	if got, _ := os.ReadFile(filepath.Join(dest, "usr/local/bin/qwsg")); string(got) != "QWSG 1.2.0-rc.7\n" {
		t.Fatalf("resulting version is not RC.7: %q", got)
	}
	if err = Rollback(dest, backup); err != nil {
		t.Fatal(err)
	}
	if got, _ := os.ReadFile(filepath.Join(dest, "usr/local/bin/qwsg")); string(got) != "QWSG 1.2.0-rc.2\n" {
		t.Fatalf("RC.2 binary not restored: %q", got)
	}
	if got, _ := os.ReadFile(filepath.Join(dest, "usr/local/lib/systemd/user/qwsg-guardian.service")); string(got) != fixture.GuardianUnit {
		t.Fatal("RC.2 Guardian unit not restored byte-for-byte")
	}
	for rel, want := range preserved {
		if got, _ := os.ReadFile(filepath.Join(dest, rel)); string(got) != want {
			t.Fatalf("rollback changed preserved data: %s", rel)
		}
	}
}
