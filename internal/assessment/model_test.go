package assessment

import (
	"context"
	"encoding/json"
	"errors"
	"os"
	"path/filepath"
	"testing"
	"time"

	"quantumwizard.hu/qwsg/internal/runner"
)

type fixtureHost struct {
	os, arch string
	uid      int
	results  map[string]runner.Result
	errors   map[string]error
}

type recordingRunner struct {
	id   string
	args []string
}

func (r *recordingRunner) Run(_ context.Context, id string, args ...string) (runner.Result, error) {
	r.id, r.args = id, append([]string(nil), args...)
	return runner.Result{}, nil
}

func (f fixtureHost) ReadOSRelease() ([]byte, error) {
	if f.os == "error" {
		return nil, errors.New("unavailable")
	}
	return []byte(f.os), nil
}
func (f fixtureHost) Architecture() string { return f.arch }
func (f fixtureHost) EffectiveUID() int    { return f.uid }
func (f fixtureHost) Run(_ context.Context, id string) (runner.Result, error) {
	return f.results[id], f.errors[id]
}

func supportedFixture() fixtureHost {
	return fixtureHost{os: "ID=ubuntu\nVERSION_ID=24.04\n", arch: "amd64", uid: 1000, results: map[string]runner.Result{"glibc_version": {Stdout: []byte("glibc 2.39\n")}, "systemd_version": {Stdout: []byte("systemd 255 (255.4)\n")}, "systemd_user": {Stdout: []byte("running\n")}}, errors: map[string]error{}}
}

func TestRegistryAndExactRemediation(t *testing.T) {
	registry := Registry()
	if err := ValidateRegistry(registry); err != nil || len(registry) < 10 {
		t.Fatalf("registry=%d err=%v", len(registry), err)
	}
	remediation := Recommendation("configuration.present", "ubuntu-24.04-amd64")
	if remediation == nil || len(remediation.Commands) != 1 || remediation.DisplayCommands[0] != "qwsg setup" || !remediation.Revalidate {
		t.Fatalf("remediation=%+v", remediation)
	}
	if Recommendation("configuration.present", "debian-12-amd64") != nil {
		t.Fatal("guessed unsupported remediation")
	}
}

func TestAssessmentClassifiesSupportedAndDeterministic(t *testing.T) {
	now := time.Date(2026, 8, 20, 10, 0, 0, 0, time.UTC)
	first := AssessInstall(context.Background(), supportedFixture(), now)
	second := AssessInstall(context.Background(), supportedFixture(), now)
	one, _ := json.Marshal(first)
	two, _ := json.Marshal(second)
	if string(one) != string(two) || first.Platform.ID != "ubuntu-24.04-amd64" {
		t.Fatalf("nondeterministic or unsupported: %s", one)
	}
	if err := ValidateReport(first); err != nil {
		t.Fatal(err)
	}
	if first.Domains[0].State != Partial {
		t.Fatalf("filesystem uncertainty not represented: %+v", first.Domains)
	}
}

func TestAssessmentRejectsUnsupportedOldRootAndFailedProbe(t *testing.T) {
	host := supportedFixture()
	host.os = "ID=debian\nVERSION_ID=12\n"
	host.arch = "arm64"
	host.uid = 0
	host.results["systemd_version"] = runner.Result{Stdout: []byte("systemd 252\n")}
	host.errors["systemd_user"] = errors.New("no bus")
	report := AssessInstall(context.Background(), host, time.Now())
	if report.Domains[0].State != NotReady {
		t.Fatalf("domain=%s", report.Domains[0].State)
	}
	for _, finding := range report.Findings {
		if finding.Remediation != nil {
			t.Fatal("unsupported platform received remediation")
		}
	}
}

func TestActionableGuidanceIsDeterministicAndCommandFree(t *testing.T) {
	host := supportedFixture()
	host.results["systemd_user"] = runner.Result{}
	host.errors["systemd_user"] = errors.New("no bus")
	report := AssessInstall(context.Background(), host, time.Unix(1, 0))
	for _, finding := range report.Findings {
		if finding.RequirementID != "systemd.user_manager" {
			continue
		}
		if finding.Guidance == nil || finding.Guidance.PrivilegeRequirement != PrivilegeManualVerification || !finding.Guidance.ManualVerification || finding.Guidance.RevalidationAction != "rerun_qwsg_install_check" {
			t.Fatalf("guidance=%+v", finding.Guidance)
		}
		if finding.Remediation != nil {
			t.Fatal("ambiguous user-manager state received a command")
		}
		return
	}
	t.Fatal("user-manager finding missing")
}

func TestEveryUserManagerFailureHasExactSupportedGuidance(t *testing.T) {
	cases := []struct {
		classification Classification
		evidence       string
	}{
		{MissingRequired, "systemd_user_runtime_directory_missing"},
		{UnknownVerification, "systemd_user_runtime_directory_unsafe"},
		{UnknownVerification, "systemd_user_manager_starting"},
		{MissingRequired, "systemd_user_manager_unavailable"},
		{UnknownVerification, "systemd_user_probe_timeout"},
		{UnknownVerification, "systemd_user_probe_output_limit"},
		{UnknownVerification, "systemd_user_probe_failed"},
		{UnknownVerification, "systemd_user_state_unrecognized"},
	}
	for _, tc := range cases {
		guidance := GuidanceFor("systemd.user_manager", "ubuntu-24.04-amd64", tc.classification, tc.evidence)
		if guidance == nil || guidance.RevalidationAction != "rerun_qwsg_install_check" {
			t.Fatalf("evidence=%s guidance=%+v", tc.evidence, guidance)
		}
		if Recommendation("systemd.user_manager", "ubuntu-24.04-amd64") != nil {
			t.Fatal("ambiguous user-manager state received remediation")
		}
	}
	if GuidanceFor("systemd.user_manager", "unsupported", MissingRequired, "systemd_user_manager_unavailable") != nil {
		t.Fatal("unsupported platform received guidance mapping")
	}
}

func TestSystemdUserManagerEvidenceStates(t *testing.T) {
	cases := []struct {
		result         runner.Result
		err            error
		classification Classification
		evidence       string
	}{
		{runner.Result{Stdout: []byte("running\n")}, nil, Satisfied, "systemd_user_manager_available"},
		{runner.Result{Stdout: []byte("degraded\n"), ExitCode: 1}, errors.New("exit 1"), Satisfied, "systemd_user_manager_available"},
		{runner.Result{Stdout: []byte("starting\n")}, nil, UnknownVerification, "systemd_user_manager_starting"},
		{runner.Result{}, context.DeadlineExceeded, UnknownVerification, "systemd_user_probe_timeout"},
		{runner.Result{}, runner.ErrOutputLimit, UnknownVerification, "systemd_user_probe_output_limit"},
		{runner.Result{}, errors.New("unavailable"), MissingRequired, "systemd_user_manager_unavailable"},
		{runner.Result{Stdout: []byte("future-state\n")}, nil, UnknownVerification, "systemd_user_state_unrecognized"},
	}
	for _, tc := range cases {
		finding := classifyUserManager(tc.result, tc.err)
		if finding.Classification != tc.classification || finding.EvidenceToken != tc.evidence {
			t.Fatalf("result=%q got=%+v", tc.result.Stdout, finding)
		}
	}
}

func TestRuntimeDirectoryValidationIsReadOnlyAndFailClosed(t *testing.T) {
	root := t.TempDir()
	directory := filepath.Join(root, "runtime")
	if _, evidence := validateRuntimeDirectory(directory, os.Geteuid()); evidence != "systemd_user_runtime_directory_missing" {
		t.Fatal(evidence)
	}
	if err := os.Mkdir(directory, 0700); err != nil {
		t.Fatal(err)
	}
	if _, evidence := validateRuntimeDirectory(directory, os.Geteuid()); evidence != "systemd_user_runtime_directory_valid" {
		t.Fatal(evidence)
	}
	if err := os.Chmod(directory, 0755); err != nil {
		t.Fatal(err)
	}
	if _, evidence := validateRuntimeDirectory(directory, os.Geteuid()); evidence != "systemd_user_runtime_directory_unsafe" {
		t.Fatal(evidence)
	}
	link := filepath.Join(root, "link")
	if err := os.Symlink(directory, link); err != nil {
		t.Fatal(err)
	}
	if _, evidence := validateRuntimeDirectory(link, os.Geteuid()); evidence != "systemd_user_runtime_directory_unsafe" {
		t.Fatal(evidence)
	}
}

func TestFilesystemPathEvidenceRejectsSymlinkAndRelativeSelection(t *testing.T) {
	root := t.TempDir()
	owned := filepath.Join(root, "owned")
	if err := os.Mkdir(owned, 0700); err != nil {
		t.Fatal(err)
	}
	if ancestor, evidence := safeExistingAncestor(filepath.Join(owned, "future", "qwsg"), os.Geteuid()); evidence != "" || ancestor != owned {
		t.Fatalf("ancestor=%q evidence=%q", ancestor, evidence)
	}
	link := filepath.Join(root, "link")
	if err := os.Symlink(owned, link); err != nil {
		t.Fatal(err)
	}
	if _, evidence := safeExistingAncestor(filepath.Join(link, "qwsg"), os.Geteuid()); evidence != "filesystem_path_unsafe" {
		t.Fatal(evidence)
	}
	t.Setenv("XDG_CONFIG_HOME", "relative")
	if _, ok := effectiveAssessmentPaths(root); ok {
		t.Fatal("relative environment path accepted")
	}
}

func TestFilesystemTypeEvidenceIsAllowlisted(t *testing.T) {
	for _, value := range []uint64{0xef53, 0x58465342, 0x9123683e} {
		classification, evidence := classifyFilesystemType(value)
		if classification != Satisfied || evidence != "filesystem_local_semantics_supported" {
			t.Fatalf("type=%x class=%s evidence=%s", value, classification, evidence)
		}
	}
	for _, value := range []uint64{0x794c7630, 0x6969, 0x01021994, 0} { // overlay, NFS, tmpfs, unknown
		classification, evidence := classifyFilesystemType(value)
		if classification != UnknownVerification || evidence != "filesystem_remote_or_overlay" {
			t.Fatalf("type=%x class=%s evidence=%s", value, classification, evidence)
		}
	}
}

func TestAggregationPrecedence(t *testing.T) {
	if Summarize([]Finding{{Classification: MissingOptional}}) != Partial || Summarize([]Finding{{Classification: UnknownVerification}}) != Unknown || Summarize([]Finding{{Classification: MissingRequired}, {Classification: MissingOptional}}) != NotReady {
		t.Fatal("aggregation precedence")
	}
}

func TestRegistryRejectsInjectedCommand(t *testing.T) {
	registry := []Requirement{{ID: "bad.command", PurposeToken: "bad", Class: RuntimeDependency, Disposition: Required, Capability: "bad", ProbeID: "bad", Platforms: []string{"ubuntu"}, PrivacyClass: "operational", Remediations: []Remediation{{PlatformID: "ubuntu", Commands: [][]string{{"sh", "-c", "evil\ncommand"}}, DisplayCommands: []string{"evil\ncommand"}, Revalidate: true}}}}
	if ValidateRegistry(registry) == nil {
		t.Fatal("unsafe remediation accepted")
	}
}

func TestRegistryRejectsUnsafeGuidance(t *testing.T) {
	registry := []Requirement{{ID: "bad.guidance", PurposeToken: "bad", Class: RuntimeDependency, Disposition: Required, Capability: "bad", ProbeID: "bad", Platforms: []string{"ubuntu"}, PrivacyClass: "operational", GuidanceRules: []GuidanceRule{{PlatformID: "ubuntu", Classification: MissingRequired, EvidenceToken: "bad", Guidance: Guidance{ExplanationToken: "bad\ntext", BlockingEffect: "blocks", VerificationActions: []string{"verify"}, OperatorActions: []string{"act"}, PrivilegeRequirement: PrivilegeNone, RevalidationAction: "retry"}}}}}
	if ValidateRegistry(registry) == nil {
		t.Fatal("unsafe guidance accepted")
	}
}

func TestFixedRunnerRejectsUnknownAndOwnsArguments(t *testing.T) {
	recording := &recordingRunner{}
	fixed := fixedRunner{base: recording}
	if _, err := fixed.Run(context.Background(), "unknown", "sh", "-c", "evil"); err == nil {
		t.Fatal("unknown probe accepted")
	}
	if _, err := fixed.Run(context.Background(), "systemd_version", "sh", "-c", "evil"); err != nil {
		t.Fatal(err)
	}
	if recording.id != "systemd_version" || len(recording.args) != 1 || recording.args[0] != "--version" {
		t.Fatalf("id=%q args=%q", recording.id, recording.args)
	}
}
