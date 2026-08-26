package update

import "testing"

func TestClassifyVersions(t *testing.T) {
	tests := []struct {
		installed, candidate string
		want                 Relation
	}{
		{"1.1.0", "1.2.0-rc.1", Newer}, {"1.2.0-rc.1", "1.2.0", Newer},
		{"1.2.0", "1.2.0", Equal}, {"1.2.0", "1.1.9", Older},
		{"1.2.0", "2.0.0", Unsupported}, {"1.2", "1.2.0", Invalid},
		{"1.2.0-rc.2", "1.2.0-rc.10", Newer},
	}
	for _, tt := range tests {
		if got := Classify(tt.installed, tt.candidate); got != tt.want {
			t.Fatalf("%s -> %s: got %s want %s", tt.installed, tt.candidate, got, tt.want)
		}
	}
}

func TestRejectsNonCanonicalVersions(t *testing.T) {
	for _, raw := range []string{"", "01.2.3", "1.02.3", "1.2.03", "1.2.3-", "1.2.3-rc.01", "1.2.3+x"} {
		if _, err := ParseVersion(raw); err == nil {
			t.Fatalf("accepted %q", raw)
		}
	}
}
