package assessment

import (
	"context"
	"encoding/json"
	"errors"
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
