package update

import "testing"

func TestCompatibilityMigration(t *testing.T) {
	for _, path := range [][2]string{{"1.1.0", "1.2.0-rc.1"}, {"1.1.0", "1.2.0-rc.2"}, {"1.2.0-rc.1", "1.2.0-rc.2"}} {
		m, err := PlanMigration(path[0], path[1])
		if err != nil || m.Validate() != nil || m.Mutation {
			t.Fatalf("migration %v: %+v %v", path, m, err)
		}
	}
}
func TestUnknownMigrationFailsClosed(t *testing.T) {
	if _, err := PlanMigration("1.1.0", "1.3.0"); err == nil {
		t.Fatal("unknown migration accepted")
	}
}
