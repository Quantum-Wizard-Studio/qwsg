package update

import "testing"

func TestCompatibilityMigration(t *testing.T) {
	for _, path := range [][2]string{{"1.1.0", "1.2.0-rc.1"}, {"1.1.0", "1.2.0-rc.2"}, {"1.2.0-rc.1", "1.2.0-rc.2"}, {"1.2.0-rc.2", "1.2.0-rc.7"}} {
		m, err := PlanMigration(path[0], path[1])
		if err != nil || m.Validate() != nil || m.Mutation || !m.PreserveConfiguration || !m.PreserveCredentials || !m.PreserveState {
			t.Fatalf("migration %v: %+v %v", path, m, err)
		}
	}
}

func TestRC2ToRC6SelectsCanonicalPath(t *testing.T) {
	m, err := PlanMigration("1.2.0-rc.2", "1.2.0-rc.7")
	if err != nil || m.ID != "compat-1.2.0-rc.2-to-1.2.0-rc.7" || !m.ReplaceGuardianUnit {
		t.Fatalf("unexpected path: %+v %v", m, err)
	}
}

func TestUnknownMigrationFailsClosed(t *testing.T) {
	for _, path := range [][2]string{{"1.1.0", "1.3.0"}, {"1.2.0-rc.2", "1.2.0-rc.5"}, {"unknown", "1.2.0-rc.7"}, {"1.2.0-rc.7", "1.2.0-rc.7"}} {
		if _, err := PlanMigration(path[0], path[1]); err == nil {
			t.Fatalf("unknown migration accepted: %v", path)
		}
	}
}
