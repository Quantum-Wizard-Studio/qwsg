package main

import (
	"bytes"
	"encoding/json"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"

	"quantumwizard.hu/qwsg/internal/assessment"
)

func TestInstallCheckIsStructuredAndReadOnly(t *testing.T) {
	root := t.TempDir()
	t.Setenv("HOME", filepath.Join(root, "absent-home"))
	t.Setenv("XDG_CONFIG_HOME", filepath.Join(root, "absent-config"))
	t.Setenv("QWSG_STATE_DIR", filepath.Join(root, "absent-state"))
	var out, errout bytes.Buffer
	code := run([]string{"install", "--check", "--format", "json"}, &out, &errout)
	if code != 0 && code != 4 {
		t.Fatalf("code=%d err=%q", code, errout.String())
	}
	var report assessment.Report
	if err := json.Unmarshal(out.Bytes(), &report); err != nil || report.Phase != "install" || len(report.Findings) == 0 {
		t.Fatalf("report=%+v err=%v", report, err)
	}
	for _, path := range []string{filepath.Join(root, "absent-home"), filepath.Join(root, "absent-config"), filepath.Join(root, "absent-state")} {
		if _, err := os.Lstat(path); !os.IsNotExist(err) {
			t.Fatalf("assessment mutated %s: %v", path, err)
		}
	}
}

func TestTask049FindingsProduceActionableGuidanceWithoutRepairCommand(t *testing.T) {
	report := assessment.Report{
		SchemaName: assessment.SchemaName, SchemaVersion: assessment.SchemaVersion,
		ModelVersion: assessment.ModelVersion, RegistryVersion: assessment.RegistryVersion,
		AssessedAt: time.Unix(1, 0).UTC(), Phase: "install",
		Platform: assessment.Platform{ID: "ubuntu-24.04-amd64", Distribution: "ubuntu", Version: "24.04", Architecture: "amd64", Supported: true},
		Findings: []assessment.Finding{
			{RequirementID: "filesystem.local_semantics", Classification: assessment.UnknownVerification, EvidenceToken: "filesystem_remote_or_overlay"},
			{RequirementID: "systemd.user_manager", Classification: assessment.MissingRequired, EvidenceToken: "systemd_user_manager_unavailable"},
		},
		Domains:     []assessment.DomainSummary{{Domain: "installation", State: assessment.NotReady}},
		NextActions: []string{"review_systemd_user_manager_guidance", "rerun_qwsg_install_check"},
	}
	attachRecommendations(&report)
	var out, errout bytes.Buffer
	if code := writeAssessment(&out, &errout, report, formatHuman); code != 4 {
		t.Fatalf("code=%d err=%q", code, errout.String())
	}
	for _, expected := range []string{"filesystem.local_semantics", "systemd.user_manager", "Explanation:", "Verify:", "Operator action:", "Privileges:", "Manual verification: required", "Safety:", "Revalidate: qwsg install --check"} {
		if !strings.Contains(out.String(), expected) {
			t.Fatalf("missing %q in %q", expected, out.String())
		}
	}
	if strings.Contains(out.String(), "Recommended:") {
		t.Fatalf("ambiguous findings received remediation: %q", out.String())
	}
	encoded, err := json.Marshal(report)
	if err != nil || !bytes.Contains(encoded, []byte(`"guidance"`)) || bytes.Contains(encoded, []byte(`"remediation"`)) {
		t.Fatalf("json=%s err=%v", encoded, err)
	}
}

func TestReadinessReportsMissingConfigurationWithoutWriting(t *testing.T) {
	root := t.TempDir()
	config := filepath.Join(root, "config")
	state := filepath.Join(root, "state")
	t.Setenv("XDG_CONFIG_HOME", config)
	t.Setenv("QWSG_STATE_DIR", state)
	var out, errout bytes.Buffer
	code := run([]string{"readiness", "--format", "json"}, &out, &errout)
	if code != 4 {
		t.Fatalf("code=%d out=%q err=%q", code, out.String(), errout.String())
	}
	var report assessment.Report
	if err := json.Unmarshal(out.Bytes(), &report); err != nil {
		t.Fatal(err)
	}
	foundConfiguration, foundOverall := false, false
	for _, finding := range report.Findings {
		if finding.RequirementID == "configuration.present" && finding.Classification == assessment.MissingRequired && finding.Remediation != nil && finding.Remediation.DisplayCommands[0] == "qwsg setup" {
			foundConfiguration = true
		}
	}
	for _, domain := range report.Domains {
		if domain.Domain == "overall" && domain.State == assessment.NotReady {
			foundOverall = true
		}
	}
	if !foundConfiguration || !foundOverall {
		t.Fatalf("report=%+v", report)
	}
	if _, err := os.Lstat(config); !os.IsNotExist(err) {
		t.Fatal("readiness wrote configuration")
	}
	if _, err := os.Lstat(state); !os.IsNotExist(err) {
		t.Fatal("readiness wrote state")
	}
}

func TestAssessmentUsageAndHumanOutput(t *testing.T) {
	for _, args := range [][]string{{"install", "--bad"}, {"readiness", "--format", "yaml"}} {
		var out, errout bytes.Buffer
		if code := run(args, &out, &errout); code != 1 || (!strings.Contains(errout.String(), "assessment accepts") && !strings.Contains(errout.String(), "supports only")) {
			t.Fatalf("args=%v code=%d err=%q", args, code, errout.String())
		}
	}
	var out, errout bytes.Buffer
	code := run([]string{"install", "--check"}, &out, &errout)
	if code != 0 && code != 4 {
		t.Fatalf("code=%d", code)
	}
	if !strings.Contains(out.String(), "QWSG install readiness") || !strings.Contains(out.String(), "systemd.version") {
		t.Fatalf("out=%q", out.String())
	}
	if strings.Contains(out.String(), "systemd.user_manager               missing_required") {
		for _, expected := range []string{"Explanation:", "Verify:", "Operator action:", "Privileges:", "Revalidate: qwsg install --check", "review_systemd_user_manager_guidance"} {
			if !strings.Contains(out.String(), expected) {
				t.Fatalf("missing %q in %q", expected, out.String())
			}
		}
	}
}

func TestJSONCarriesEquivalentActionableGuidance(t *testing.T) {
	var out, errout bytes.Buffer
	code := run([]string{"install", "--check", "--format", "json"}, &out, &errout)
	if code != 0 && code != 4 {
		t.Fatalf("code=%d err=%q", code, errout.String())
	}
	var report assessment.Report
	if err := json.Unmarshal(out.Bytes(), &report); err != nil {
		t.Fatal(err)
	}
	if report.SchemaVersion != "1.1" || report.ModelVersion != "1.1" || report.RegistryVersion != "1.1" {
		t.Fatalf("versions=%s/%s/%s", report.SchemaVersion, report.ModelVersion, report.RegistryVersion)
	}
	for _, finding := range report.Findings {
		if finding.RequirementID == "systemd.user_manager" && finding.Classification != assessment.Satisfied {
			if finding.Guidance == nil || finding.Guidance.RevalidationAction != "rerun_qwsg_install_check" {
				t.Fatalf("finding=%+v", finding)
			}
			return
		}
	}
}

func TestOperationalAggregationKeepsNotificationOptional(t *testing.T) {
	findings := []assessment.Finding{
		{RequirementID: "platform.operating_system", Classification: assessment.Satisfied},
		{RequirementID: "configuration.present", Classification: assessment.Satisfied},
		{RequirementID: "notification.external", Classification: assessment.MissingOptional},
		{RequirementID: "guardian.unit_installed", Classification: assessment.Satisfied},
		{RequirementID: "guardian.service_active", Classification: assessment.Satisfied},
		{RequirementID: "guardian.canonical_evidence", Classification: assessment.Satisfied},
	}
	domains := operationalDomains(findings)
	if domains[len(domains)-1].State != assessment.Partial {
		t.Fatalf("domains=%+v", domains)
	}
}

func TestSetupPlanJSONIsReadOnlyAndStructured(t *testing.T) {
	root := t.TempDir()
	t.Setenv("XDG_CONFIG_HOME", filepath.Join(root, "config"))
	t.Setenv("QWSG_STATE_DIR", filepath.Join(root, "state"))
	var out, errout bytes.Buffer
	if code := run([]string{"setup", "--plan", "--format", "json"}, &out, &errout); code != 0 {
		t.Fatalf("code=%d err=%q", code, errout.String())
	}
	if !strings.Contains(out.String(), `"schema_name": "qwsg.setup-flow"`) || !strings.Contains(out.String(), `"next_action"`) {
		t.Fatalf("out=%q", out.String())
	}
	for _, p := range []string{filepath.Join(root, "config"), filepath.Join(root, "state")} {
		if _, err := os.Lstat(p); !os.IsNotExist(err) {
			t.Fatalf("plan mutated %s", p)
		}
	}
}
