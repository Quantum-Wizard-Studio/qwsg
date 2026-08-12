package main

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"
	"time"

	canonicalcommand "quantumwizard.hu/qwsg/internal/command"
	"quantumwizard.hu/qwsg/internal/comparison"
	"quantumwizard.hu/qwsg/internal/configuration"
	"quantumwizard.hu/qwsg/internal/configurationstore"
	"quantumwizard.hu/qwsg/internal/inventory"
	"quantumwizard.hu/qwsg/internal/inventorystore"
	"quantumwizard.hu/qwsg/internal/operatorstate"
	"quantumwizard.hu/qwsg/internal/pipeline"
	"quantumwizard.hu/qwsg/internal/presentationmodel"
)

func TestCLIHelpVersionAndInvalid(t *testing.T) {
	for _, tc := range []struct {
		args     []string
		code     int
		contains string
	}{
		{[]string{"help"}, 0, "snapshot explorer"},
		{[]string{"help", "inventory"}, 0, "inventory list"},
		{[]string{"inventory", "list", "--help"}, 0, "inventory list"},
		{[]string{"version"}, 0, "commit:"},
		{[]string{"unknown"}, 1, "unknown command"},
	} {
		var out, err bytes.Buffer
		code := run(tc.args, &out, &err)
		if code != tc.code || !strings.Contains(out.String()+err.String(), tc.contains) {
			t.Fatalf("%v: %d %q", tc.args, code, out.String()+err.String())
		}
	}
}

func TestSetupAndConfigCommandLifecycle(t *testing.T) {
	home := filepath.Join(t.TempDir(), "home")
	if err := os.Mkdir(home, 0700); err != nil {
		t.Fatal(err)
	}
	t.Setenv("HOME", home)
	t.Setenv("XDG_CONFIG_HOME", "")
	path := filepath.Join(home, ".config", "qwsg", "config.json")
	var out, errout bytes.Buffer
	if code := run([]string{"config", "show"}, &out, &errout); code != 0 || !strings.Contains(out.String(), "Configured: false") {
		t.Fatalf("show defaults code=%d out=%q err=%q", code, out.String(), errout.String())
	}
	out.Reset()
	errout.Reset()
	if code := run([]string{"setup"}, &out, &errout); code != 1 || !strings.Contains(errout.String(), "requires a terminal") {
		t.Fatalf("noninteractive setup code=%d out=%q err=%q", code, out.String(), errout.String())
	}
	if _, err := os.Stat(path); !errors.Is(err, os.ErrNotExist) {
		t.Fatal("setup wrote before acceptance")
	}
	out.Reset()
	errout.Reset()
	if code := run([]string{"setup", "--accept-defaults", "--set", "locale=hu", "--set", "guardian.interval=10m"}, &out, &errout); code != 0 {
		t.Fatalf("setup code=%d out=%q err=%q", code, out.String(), errout.String())
	}
	if info, err := os.Stat(path); err != nil || info.Mode().Perm() != 0600 {
		t.Fatalf("config mode=%v err=%v", info, err)
	}
	out.Reset()
	errout.Reset()
	if code := run([]string{"config", "get", "locale"}, &out, &errout); code != 0 || strings.TrimSpace(out.String()) != "hu" {
		t.Fatalf("get code=%d out=%q err=%q", code, out.String(), errout.String())
	}
	out.Reset()
	errout.Reset()
	if code := run([]string{"config", "set", "guardian.cycle_timeout", "3m"}, &out, &errout); code != 0 {
		t.Fatalf("set code=%d out=%q err=%q", code, out.String(), errout.String())
	}
	out.Reset()
	errout.Reset()
	if code := run([]string{"config", "show", "--format", "json"}, &out, &errout); code != 0 || strings.Contains(out.String(), "QWSG setup plan") {
		t.Fatalf("json show code=%d out=%q err=%q", code, out.String(), errout.String())
	}
	var result configResult
	if err := json.Unmarshal(out.Bytes(), &result); err != nil || !result.Configured || guardianInterval(result.Effective) != 10*time.Minute || guardianTimeout(result.Effective) != 3*time.Minute {
		t.Fatalf("result=%+v err=%v", result, err)
	}
	before, _ := os.ReadFile(path)
	out.Reset()
	errout.Reset()
	if code := run([]string{"config", "set", "guardian.cycle_timeout", "20m"}, &out, &errout); code != 1 || !strings.Contains(errout.String(), "configuration_invalid") {
		t.Fatalf("inconsistent timing code=%d err=%q", code, errout.String())
	}
	unchanged, _ := os.ReadFile(path)
	if !bytes.Equal(before, unchanged) {
		t.Fatal("invalid timing replaced configuration")
	}
	out.Reset()
	errout.Reset()
	if code := run([]string{"setup", "--accept-defaults"}, &out, &errout); code != 0 {
		t.Fatalf("repeat setup: %d %q", code, errout.String())
	}
	after, _ := os.ReadFile(path)
	if !bytes.Equal(before, after) {
		t.Fatal("repeat setup changed valid configuration")
	}
}

func TestConfigurationFailuresAreBoundedAndGuardianPreflights(t *testing.T) {
	home := filepath.Join(t.TempDir(), "home")
	configDir := filepath.Join(home, ".config", "qwsg")
	if err := os.MkdirAll(configDir, 0700); err != nil {
		t.Fatal(err)
	}
	path := filepath.Join(configDir, "config.json")
	if err := os.WriteFile(path, []byte(`{"unknown":true}`), 0600); err != nil {
		t.Fatal(err)
	}
	t.Setenv("HOME", home)
	t.Setenv("XDG_CONFIG_HOME", "")
	t.Setenv("QWSG_STATE_DIR", filepath.Join(home, "state"))
	for _, args := range [][]string{{"config", "validate"}, {"console"}} {
		var out, errout bytes.Buffer
		if code := run(args, &out, &errout); code != 1 || !strings.Contains(errout.String(), "configuration_invalid") || strings.Contains(errout.String(), home) {
			t.Fatalf("%v code=%d out=%q err=%q", args, code, out.String(), errout.String())
		}
	}
	options, err := parseGuardianOptions([]string{"--interval", "10m", "--cycle-timeout", "2m"})
	if err != nil {
		t.Fatal(err)
	}
	if err = executeGuardian(options); !errors.Is(err, configurationstore.ErrInvalid) {
		t.Fatalf("guardian invalid config: %v", err)
	}
	if _, statErr := os.Stat(filepath.Join(home, "state")); !errors.Is(statErr, os.ErrNotExist) {
		t.Fatal("Guardian mutated state before configuration validation")
	}
}

func TestInteractiveSetupRequiresExplicitYes(t *testing.T) {
	home := filepath.Join(t.TempDir(), "home")
	if err := os.Mkdir(home, 0700); err != nil {
		t.Fatal(err)
	}
	t.Setenv("HOME", home)
	t.Setenv("XDG_CONFIG_HOME", "")
	path := filepath.Join(home, ".config", "qwsg", "config.json")
	for _, tc := range []struct {
		answer  string
		written bool
	}{{"n\n", false}, {"y\n", true}} {
		var out, errout bytes.Buffer
		if code := runWithConsole([]string{"setup"}, strings.NewReader(tc.answer), &out, &errout, true); code != 0 {
			t.Fatalf("answer=%q code=%d err=%q", tc.answer, code, errout.String())
		}
		_, err := os.Stat(path)
		if tc.written != (err == nil) {
			t.Fatalf("answer=%q written=%v err=%v", tc.answer, tc.written, err)
		}
	}
}

func TestConfigRejectsUnsafePathsPermissionsAndUnknownKeys(t *testing.T) {
	home := filepath.Join(t.TempDir(), "home")
	if err := os.Mkdir(home, 0700); err != nil {
		t.Fatal(err)
	}
	t.Setenv("HOME", home)
	t.Setenv("XDG_CONFIG_HOME", "relative")
	var out, errout bytes.Buffer
	if code := run([]string{"config", "show"}, &out, &errout); code != 1 || !strings.Contains(errout.String(), "configuration_path_unsafe") {
		t.Fatalf("unsafe XDG: %d %q", code, errout.String())
	}
	t.Setenv("XDG_CONFIG_HOME", "")
	if code := run([]string{"config", "set", "smtp.password", "secret"}, &out, &errout); code != 1 || !strings.Contains(errout.String(), "unsupported configuration key") {
		t.Fatalf("secret key accepted: %d %q", code, errout.String())
	}
	configDir := filepath.Join(home, ".config", "qwsg")
	if err := os.MkdirAll(configDir, 0700); err != nil {
		t.Fatal(err)
	}
	path := filepath.Join(configDir, "config.json")
	locale := "en"
	source, _ := configuration.NormalizeSource(configuration.Source{SchemaName: configuration.SourceSchema, SchemaVersion: configuration.SchemaVersion, ModelVersion: configuration.ModelVersion, ID: "local.operator", SourceVersion: "1.0", Kind: configuration.PrimaryLocal, Patch: configuration.Patch{Locale: &locale}})
	document, _ := configuration.MarshalSourceCanonical(source)
	if err := os.WriteFile(path, document, 0644); err != nil {
		t.Fatal(err)
	}
	out.Reset()
	errout.Reset()
	if code := run([]string{"config", "validate"}, &out, &errout); code != 1 || !strings.Contains(errout.String(), "configuration_permission_unsafe") {
		t.Fatalf("permission accepted: %d %q", code, errout.String())
	}
}

func TestNotificationRecipientExtensionRemainsInertAndRepresentable(t *testing.T) {
	for _, recipients := range []string{`["admin@example.invalid"]`, `["a@example.invalid","b@example.invalid","c@example.invalid"]`} {
		extensions := []configuration.Extension{{ID: "notification.recipients", Version: "1.0", Required: false, Fields: map[string]string{"addresses_json": recipients}}}
		source, err := configuration.NormalizeSource(configuration.Source{SchemaName: configuration.SourceSchema, SchemaVersion: configuration.SchemaVersion, ModelVersion: configuration.ModelVersion, ID: "local.operator", SourceVersion: "1.0", Kind: configuration.PrimaryLocal, Patch: configuration.Patch{Extensions: &extensions}})
		if err != nil {
			t.Fatal(err)
		}
		effective, err := resolveLocalConfiguration(source, true, nil)
		if err != nil || effective.Values.Extensions[0].Fields["addresses_json"] != recipients {
			t.Fatalf("recipients=%s effective=%+v err=%v", recipients, effective.Values.Extensions, err)
		}
	}
}

func TestBareAndExplicitConsoleAreNonBlockingWithoutTerminal(t *testing.T) {
	t.Setenv("QWSG_STATE_DIR", filepath.Join(t.TempDir(), "missing-state"))
	for _, args := range [][]string{{}, {"console"}} {
		var out, errout bytes.Buffer
		if code := run(args, &out, &errout); code != 0 {
			t.Fatalf("%v: code=%d err=%q", args, code, errout.String())
		}
		for _, expected := range []string{"QWSG Operator Console", "Server condition: unavailable", "Guardian: not observed", "Recommended action: run qwsg observe"} {
			if !strings.Contains(out.String(), expected) {
				t.Fatalf("%v missing %q in %q", args, expected, out.String())
			}
		}
	}
}

func TestCheckPublicationSurvivesProcessBoundaryForBareConsole(t *testing.T) {
	root := filepath.Join(t.TempDir(), "current")
	t.Setenv("QWSG_STATE_DIR", root)
	snapshot := cliFixtureSnapshot()
	definition, err := canonicalcommand.ResolveProfile("check", canonicalcommand.Selection{Source: "live"})
	if err != nil {
		t.Fatal(err)
	}
	orchestrator := pipeline.Orchestrator{Collect: func(context.Context) (inventory.Snapshot, error) { return snapshot, nil }, Retention: inventorystore.DefaultRetention, Rules: pipeline.CanonicalObservationRules()}
	execution, err := orchestrator.Execute(context.Background(), definition)
	if err != nil {
		t.Fatal(err)
	}
	if _, err = publishCheck(definition, execution, snapshot.CompletedAt.Add(time.Second)); err != nil {
		t.Fatal(err)
	}
	var out, errout bytes.Buffer
	if code := run([]string{}, &out, &errout); code != 0 {
		t.Fatalf("code=%d err=%q", code, errout.String())
	}
	if strings.Contains(out.String(), "Server condition: unavailable") || !strings.Contains(out.String(), "Evidence: stale / partial") {
		t.Fatalf("persisted state not consumed: %q", out.String())
	}
}

func TestCurrentStateAcrossSeparateProcesses(t *testing.T) {
	root := filepath.Join(t.TempDir(), "current")
	writer := exec.Command(os.Args[0], "-test.run=TestCurrentStateSubprocessHelper", "--", "write")
	writer.Env = append(os.Environ(), "QWSG_STATE_DIR="+root, "QWSG_SUBPROCESS_HELPER=1")
	if output, err := writer.CombinedOutput(); err != nil {
		t.Fatalf("writer: %v %s", err, output)
	}
	reader := exec.Command(os.Args[0], "-test.run=TestCurrentStateSubprocessHelper", "--", "read")
	reader.Env = append(os.Environ(), "QWSG_STATE_DIR="+root, "QWSG_SUBPROCESS_HELPER=1")
	output, err := reader.CombinedOutput()
	if err != nil {
		t.Fatalf("reader: %v %s", err, output)
	}
	text := string(output)
	if strings.Contains(text, "Server condition: unavailable") || !strings.Contains(text, "Evidence: stale / partial") {
		t.Fatalf("state did not cross process boundary: %q", text)
	}
}

func TestObserveBootstrapQualifiedEvaluationAndConsoleAcrossProcesses(t *testing.T) {
	root := filepath.Join(t.TempDir(), "state")
	base := time.Now().UTC().Add(-70 * time.Second).Truncate(time.Second).Format(time.RFC3339Nano)
	for _, mode := range []string{"observe-bootstrap", "observe-qualified"} {
		process := exec.Command(os.Args[0], "-test.run=TestCurrentStateSubprocessHelper", "--", mode)
		process.Env = append(os.Environ(), "QWSG_STATE_DIR="+root, "QWSG_STORE=", "QWSG_SUBPROCESS_HELPER=1", "QWSG_TEST_OBSERVED_AT="+base)
		if output, err := process.CombinedOutput(); err != nil {
			t.Fatalf("%s: %v %s", mode, err, output)
		}
	}
	stored, loadErr := func() (operatorstate.State, error) {
		store, openErr := operatorstate.Open(root)
		if openErr != nil {
			return operatorstate.State{}, openErr
		}
		return store.Load()
	}()
	if loadErr != nil || !time.Now().UTC().Before(stored.FreshUntil) {
		t.Fatalf("qualified subprocess state is not fresh: now=%s fresh_until=%s err=%v", time.Now().UTC(), stored.FreshUntil, loadErr)
	}
	reader := exec.Command(os.Args[0], "-test.run=TestCurrentStateSubprocessHelper", "--", "read-qualified")
	reader.Env = append(os.Environ(), "QWSG_STATE_DIR="+root, "QWSG_STORE=", "QWSG_SUBPROCESS_HELPER=1", "QWSG_TEST_OBSERVED_AT="+base)
	output, err := reader.CombinedOutput()
	if err != nil {
		t.Fatalf("reader: %v %s", err, output)
	}
	text := string(output)
	for _, forbidden := range []string{"Server condition: unavailable", "Server condition: unknown", "Guardian: running"} {
		if strings.Contains(text, forbidden) {
			t.Fatalf("qualified state was misrepresented: %q", text)
		}
	}
	for _, expected := range []string{"Changes:", "Alerts: 0", "Guardian: not observed", "Evidence: current / complete"} {
		if !strings.Contains(text, expected) {
			t.Fatalf("qualified console lacks %q: %q", expected, text)
		}
	}
}

func TestLargeObservePublishesAndConsoleConsumesAcrossProcesses(t *testing.T) {
	root := filepath.Join(t.TempDir(), "state")
	base := time.Now().UTC().Add(-90 * time.Second).Truncate(time.Second).Format(time.RFC3339Nano)
	for _, mode := range []string{"observe-large-bootstrap", "observe-large-qualified"} {
		process := exec.Command(os.Args[0], "-test.run=TestCurrentStateSubprocessHelper", "--", mode)
		process.Env = append(os.Environ(), "QWSG_STATE_DIR="+root, "QWSG_STORE=", "QWSG_SUBPROCESS_HELPER=1", "QWSG_TEST_OBSERVED_AT="+base)
		if output, err := process.CombinedOutput(); err != nil {
			t.Fatalf("%s: %v %s", mode, err, output)
		}
	}
	store, _ := operatorstate.Open(root)
	stored, err := store.Load()
	if err != nil || stored.Coverage != operatorstate.CoverageOperatorEvaluation || stored.Overview.AttentionSummary == nil || stored.Overview.AttentionSummary.TotalCandidates <= presentationmodel.MaxAttention || len(stored.Overview.AttentionItems) > presentationmodel.MaxAttention {
		t.Fatalf("large current state was not published safely: %+v err=%v", stored.Overview.AttentionSummary, err)
	}
	reader := exec.Command(os.Args[0], "-test.run=TestCurrentStateSubprocessHelper", "--", "read-qualified")
	reader.Env = append(os.Environ(), "QWSG_STATE_DIR="+root, "QWSG_STORE=", "QWSG_SUBPROCESS_HELPER=1")
	output, err := reader.CombinedOutput()
	if err != nil {
		t.Fatalf("reader: %v %s", err, output)
	}
	text := string(output)
	for _, expected := range []string{"Additional concerns summarized:", "Guardian: not observed", "Evidence: current / complete"} {
		if !strings.Contains(text, expected) {
			t.Fatalf("large Console output lacks %q: %q", expected, text)
		}
	}
	if strings.Contains(text, "Server condition: unknown") || strings.Contains(text, "Server condition: unavailable") || strings.Contains(text, "Guardian: running") {
		t.Fatalf("large qualified state was misrepresented: %q", text)
	}
}

func TestCurrentStateSubprocessHelper(t *testing.T) {
	if os.Getenv("QWSG_SUBPROCESS_HELPER") != "1" {
		return
	}
	mode := os.Args[len(os.Args)-1]
	if mode == "write" {
		snapshot := cliFixtureSnapshot()
		if value := os.Getenv("QWSG_TEST_OBSERVED_AT"); value != "" {
			at, parseErr := time.Parse(time.RFC3339Nano, value)
			if parseErr != nil {
				t.Fatal(parseErr)
			}
			snapshot = shiftCLIFixtureSnapshot(snapshot, at.UTC())
		}
		definition, _ := canonicalcommand.ResolveProfile("check", canonicalcommand.Selection{Source: "live"})
		orchestrator := pipeline.Orchestrator{Collect: func(context.Context) (inventory.Snapshot, error) { return snapshot, nil }, Retention: inventorystore.DefaultRetention, Rules: pipeline.CanonicalObservationRules()}
		execution, err := orchestrator.Execute(context.Background(), definition)
		if err != nil {
			t.Fatal(err)
		}
		if _, err = publishCheck(definition, execution, snapshot.CompletedAt.Add(time.Second)); err != nil {
			t.Fatal(err)
		}
		return
	}
	if mode == "read" {
		if code := run([]string{}, os.Stdout, os.Stderr); code != 0 {
			os.Exit(code)
		}
		return
	}
	if mode == "observe-bootstrap" || mode == "observe-qualified" || mode == "observe-large-bootstrap" || mode == "observe-large-qualified" {
		storeRoot, err := observationStoreRoot()
		if err != nil {
			t.Fatal(err)
		}
		definition, err := canonicalcommand.ResolveProfile("observe", canonicalcommand.Selection{Source: "live", Store: storeRoot})
		if err != nil {
			t.Fatal(err)
		}
		snapshot := cliFixtureSnapshot()
		if value := os.Getenv("QWSG_TEST_OBSERVED_AT"); value != "" {
			at, parseErr := time.Parse(time.RFC3339Nano, value)
			if parseErr != nil {
				t.Fatal(parseErr)
			}
			snapshot = shiftCLIFixtureSnapshot(snapshot, at.UTC())
		}
		if mode == "observe-large-bootstrap" || mode == "observe-large-qualified" {
			snapshot = largeCLIFixtureSnapshot(snapshot.CompletedAt, 0)
		}
		if mode == "observe-qualified" {
			snapshot = laterCLIFixtureSnapshot(snapshot)
		} else if mode == "observe-large-qualified" {
			snapshot = largeCLIFixtureSnapshot(snapshot.CompletedAt.Add(time.Minute), 1)
		}
		result, err := observeOnce(context.Background(), definition, func(context.Context) (inventory.Snapshot, error) { return snapshot, nil }, func() time.Time { return snapshot.CompletedAt.Add(time.Second) })
		if err != nil {
			t.Fatal(err)
		}
		bootstrapMode := mode == "observe-bootstrap" || mode == "observe-large-bootstrap"
		if bootstrapMode != result.Bootstrap {
			t.Fatalf("unexpected bootstrap state: %s %+v", mode, result)
		}
		return
	}
	if mode == "read-qualified" {
		if code := run([]string{}, os.Stdout, os.Stderr); code != 0 {
			os.Exit(code)
		}
		return
	}
	t.Fatalf("unknown helper mode %q", mode)
}

func TestCheckPublicationRejectsUntypedCoverage(t *testing.T) {
	t.Setenv("QWSG_STATE_DIR", filepath.Join(t.TempDir(), "current"))
	snapshot := cliFixtureSnapshot()
	definition, _ := canonicalcommand.ResolveProfile("check", canonicalcommand.Selection{Source: "live"})
	orchestrator := pipeline.Orchestrator{Collect: func(context.Context) (inventory.Snapshot, error) { return snapshot, nil }, Retention: inventorystore.DefaultRetention, Rules: pipeline.CanonicalObservationRules()}
	execution, err := orchestrator.Execute(context.Background(), definition)
	if err != nil {
		t.Fatal(err)
	}
	execution.Stages[0].Value = map[string]any{"snapshot_id": snapshot.SnapshotID}
	if _, err = publishCheck(definition, execution, snapshot.CompletedAt.Add(time.Second)); err == nil {
		t.Fatal("untyped stage published")
	}
}

func TestCheckPublishesValidPartialEvidence(t *testing.T) {
	root := filepath.Join(t.TempDir(), "current")
	t.Setenv("QWSG_STATE_DIR", root)
	snapshot := partialCLIFixtureSnapshot()
	definition, _ := canonicalcommand.ResolveProfile("check", canonicalcommand.Selection{Source: "live"})
	orchestrator := pipeline.Orchestrator{Collect: func(context.Context) (inventory.Snapshot, error) { return snapshot, nil }, Retention: inventorystore.DefaultRetention, Rules: pipeline.CanonicalObservationRules()}
	execution, err := orchestrator.Execute(context.Background(), definition)
	if err != nil {
		t.Fatal(err)
	}
	if execution.Stages[0].Complete || execution.Stages[1].Complete {
		t.Fatal("partial evidence was marked complete")
	}
	overview, err := publishCheck(definition, execution, snapshot.CompletedAt.Add(time.Second))
	if err != nil {
		t.Fatal(err)
	}
	if overview.Completeness != presentationmodel.CompletenessPartial || overview.Condition != presentationmodel.Degraded {
		t.Fatalf("partial evidence misrepresented: %+v", overview)
	}
	store, _ := operatorstate.Open(root)
	if current, loadErr := store.Load(); loadErr != nil || current.Overview.Completeness != presentationmodel.CompletenessPartial {
		t.Fatalf("partial current state: %+v err=%v", current, loadErr)
	}
}

func TestObserveBootstrapsThenPublishesFullEvaluation(t *testing.T) {
	stateRoot := filepath.Join(t.TempDir(), "state")
	t.Setenv("QWSG_STATE_DIR", stateRoot)
	storeRoot, err := observationStoreRoot()
	if err != nil {
		t.Fatal(err)
	}
	definition, err := canonicalcommand.ResolveProfile("observe", canonicalcommand.Selection{Source: "live", Store: storeRoot})
	if err != nil {
		t.Fatal(err)
	}
	first := cliFixtureSnapshot()
	result, err := observeOnce(context.Background(), definition, func(context.Context) (inventory.Snapshot, error) { return first, nil }, func() time.Time { return first.CompletedAt.Add(time.Second) })
	if err != nil {
		t.Fatal(err)
	}
	if !result.Bootstrap || len(result.Execution.Stages) != 2 || result.Overview.Condition != presentationmodel.Unknown || result.Overview.Guardian != presentationmodel.GuardianNotObserved {
		t.Fatalf("dishonest bootstrap: %+v", result)
	}
	store, _ := inventorystore.Open(storeRoot, inventorystore.DefaultRetention)
	if names, listErr := store.List(); listErr != nil || len(names) != 1 {
		t.Fatalf("bootstrap store: names=%v err=%v", names, listErr)
	}
	currentStore, _ := operatorstate.Open(stateRoot)
	current, err := currentStore.Load()
	if err != nil || current.Coverage != operatorstate.CoverageInventorySnapshot {
		t.Fatalf("bootstrap current state: %+v %v", current, err)
	}

	second := laterCLIFixtureSnapshot(first)
	result, err = observeOnce(context.Background(), definition, func(context.Context) (inventory.Snapshot, error) { return second, nil }, func() time.Time { return second.CompletedAt.Add(time.Second) })
	if err != nil {
		t.Fatal(err)
	}
	if result.Bootstrap || len(result.Execution.Stages) != 8 || result.Overview.Condition == presentationmodel.Unknown || result.Overview.Guardian != presentationmodel.GuardianNotObserved {
		t.Fatalf("unqualified full observation: %+v", result)
	}
	current, err = currentStore.Load()
	if err != nil || current.Coverage != operatorstate.CoverageOperatorEvaluation || current.Provenance.Profile != "observe" || len(current.Provenance.Stages) != 8 {
		t.Fatalf("full current state: %+v %v", current, err)
	}
	if names, listErr := store.List(); listErr != nil || len(names) != 2 {
		t.Fatalf("qualified store: names=%v err=%v", names, listErr)
	}
}

func TestObserveRejectsCorruptStoreAndTypedTamperingWithoutReplacingState(t *testing.T) {
	stateRoot := filepath.Join(t.TempDir(), "state")
	t.Setenv("QWSG_STATE_DIR", stateRoot)
	storeRoot, err := observationStoreRoot()
	if err != nil {
		t.Fatal(err)
	}
	definition, _ := canonicalcommand.ResolveProfile("observe", canonicalcommand.Selection{Source: "live", Store: storeRoot})
	first := cliFixtureSnapshot()
	if _, err = observeOnce(context.Background(), definition, func(context.Context) (inventory.Snapshot, error) { return first, nil }, func() time.Time { return first.CompletedAt.Add(time.Second) }); err != nil {
		t.Fatal(err)
	}
	second := laterCLIFixtureSnapshot(first)
	orchestrator := pipeline.Orchestrator{Collect: func(context.Context) (inventory.Snapshot, error) { return second, nil }, Retention: inventorystore.DefaultRetention, Rules: pipeline.CanonicalObservationRules()}
	execution, err := orchestrator.Execute(context.Background(), definition)
	if err != nil {
		t.Fatal(err)
	}
	currentStore, _ := operatorstate.Open(stateRoot)
	before, _ := currentStore.Load()
	execution.Stages[4].Value = map[string]any{"overall_status": "healthy"}
	if _, err = publishObserve(definition, execution, second.CompletedAt.Add(time.Second)); err == nil || observeDiagnostic(err) != string(observationProjectionFailure) {
		t.Fatal("untyped health stage published")
	}
	after, err := currentStore.Load()
	if err != nil || after.ID != before.ID {
		t.Fatalf("last valid state replaced: before=%s after=%s err=%v", before.ID, after.ID, err)
	}

	corruptRoot := filepath.Join(t.TempDir(), "corrupt")
	if err = os.Mkdir(corruptRoot, 0o700); err != nil {
		t.Fatal(err)
	}
	if err = os.Mkdir(filepath.Join(corruptRoot, "snapshots"), 0o700); err != nil {
		t.Fatal(err)
	}
	if err = os.WriteFile(filepath.Join(corruptRoot, "store.json"), []byte("broken"), 0o600); err != nil {
		t.Fatal(err)
	}
	corruptDefinition, _ := canonicalcommand.ResolveProfile("observe", canonicalcommand.Selection{Source: "live", Store: corruptRoot})
	if _, err = observeOnce(context.Background(), corruptDefinition, func(context.Context) (inventory.Snapshot, error) { return first, nil }, func() time.Time { return first.CompletedAt.Add(time.Second) }); !strings.Contains(observeDiagnostic(err), "corrupt") {
		t.Fatalf("corrupt store silently bootstrapped: %v", err)
	}
}

func TestObserveDiagnosticsAreTypedBoundedAndPrivacySafe(t *testing.T) {
	stateRoot := filepath.Join(t.TempDir(), "state")
	t.Setenv("QWSG_STATE_DIR", stateRoot)
	storeRoot, err := observationStoreRoot()
	if err != nil {
		t.Fatal(err)
	}
	definition, _ := canonicalcommand.ResolveProfile("observe", canonicalcommand.Selection{Source: "live", Store: storeRoot})
	private := errors.New("/private/host/path secret=value")
	_, err = observeOnce(context.Background(), definition, func(context.Context) (inventory.Snapshot, error) { return inventory.Snapshot{}, private }, func() time.Time { return time.Now().UTC() })
	if got := observeDiagnostic(err); got != string(observationPipelineFailure) || strings.Contains(got, "private") || strings.Contains(err.Error(), "private") {
		t.Fatalf("unsafe pipeline diagnostic: public=%q error=%q", got, err)
	}
	for kind, want := range map[observationFailureKind]string{
		observationBootstrapFailure:   "state_bootstrap_failed",
		observationProjectionFailure:  "operator_projection_failed",
		observationPublicationFailure: "current_state_publication_failed",
	} {
		failure := classifiedObservationFailure(kind, private)
		if got := observeDiagnostic(failure); got != want || strings.Contains(got, "private") || strings.Contains(failure.Error(), "private") {
			t.Fatalf("unsafe %s diagnostic: public=%q error=%q", kind, got, failure)
		}
	}
}

func TestPublicationDiagnosticDoesNotMislabelNonStateFailure(t *testing.T) {
	private := errors.New("/private/host/path secret=value")
	if got := publicationDiagnostic(private); got != "state_publication_failed" || strings.Contains(got, "private") {
		t.Fatalf("unsafe publication diagnostic: %q", got)
	}
	if got := publicationDiagnostic(operatorstate.ErrUnsafePath); got != "state_unsafe" {
		t.Fatalf("unsafe path diagnostic: %q", got)
	}
}

func TestObserveDefaultStoreIsPrivateStateRelative(t *testing.T) {
	root := filepath.Join(t.TempDir(), "state")
	t.Setenv("QWSG_STATE_DIR", root)
	t.Setenv("QWSG_STORE", "")
	got, err := observationStoreRoot()
	if err != nil || got != filepath.Join(root, "inventory") {
		t.Fatalf("default store=%q err=%v", got, err)
	}
	explicit := filepath.Join(t.TempDir(), "explicit")
	t.Setenv("QWSG_STORE", explicit)
	if got, err = observationStoreRoot(); err != nil || got != explicit {
		t.Fatalf("explicit store=%q err=%v", got, err)
	}
}

func TestObserveDefaultStoreBootstrapsEmptyHome(t *testing.T) {
	home := filepath.Join(t.TempDir(), "home")
	if err := os.Mkdir(home, 0o755); err != nil {
		t.Fatal(err)
	}
	t.Setenv("HOME", home)
	t.Setenv("XDG_STATE_HOME", "")
	t.Setenv("QWSG_STATE_DIR", "")
	t.Setenv("QWSG_STORE", "")
	storeRoot, err := observationStoreRoot()
	if err != nil {
		t.Fatal(err)
	}
	stateRoot := filepath.Join(home, ".local", "state", "qwsg")
	if storeRoot != filepath.Join(stateRoot, "inventory") {
		t.Fatalf("store root %q", storeRoot)
	}
	if mode := mustPathMode(t, stateRoot); mode != 0o700 {
		t.Fatalf("state root mode %o", mode)
	}
	definition, _ := canonicalcommand.ResolveProfile("observe", canonicalcommand.Selection{Source: "live", Store: storeRoot})
	snapshot := partialCLIFixtureSnapshot()
	result, err := observeOnce(context.Background(), definition, func(context.Context) (inventory.Snapshot, error) { return snapshot, nil }, func() time.Time { return snapshot.CompletedAt.Add(time.Second) })
	if err != nil {
		t.Fatal(err)
	}
	if !result.Bootstrap || result.Overview.Completeness != presentationmodel.CompletenessPartial {
		t.Fatalf("bootstrap result: %+v", result)
	}
	for _, path := range []string{filepath.Join(storeRoot, "store.json"), filepath.Join(stateRoot, operatorstate.FileName)} {
		if mode := mustPathMode(t, path); mode != 0o600 {
			t.Fatalf("%s mode %o", filepath.Base(path), mode)
		}
	}
	store, _ := inventorystore.Open(storeRoot, inventorystore.DefaultRetention)
	if names, listErr := store.List(); listErr != nil || len(names) != 1 {
		t.Fatalf("baseline names=%v err=%v", names, listErr)
	}
}

func mustPathMode(t *testing.T, path string) os.FileMode {
	t.Helper()
	info, err := os.Stat(path)
	if err != nil {
		t.Fatal(err)
	}
	return info.Mode().Perm()
}

func TestObservePartialBootstrapPublishesUnknownPartialEvidence(t *testing.T) {
	stateRoot := filepath.Join(t.TempDir(), "state")
	t.Setenv("QWSG_STATE_DIR", stateRoot)
	storeRoot, err := observationStoreRoot()
	if err != nil {
		t.Fatal(err)
	}
	definition, _ := canonicalcommand.ResolveProfile("observe", canonicalcommand.Selection{Source: "live", Store: storeRoot})
	snapshot := partialCLIFixtureSnapshot()
	result, err := observeOnce(context.Background(), definition, func(context.Context) (inventory.Snapshot, error) { return snapshot, nil }, func() time.Time { return snapshot.CompletedAt.Add(time.Second) })
	if err != nil {
		t.Fatal(err)
	}
	if !result.Bootstrap || result.Execution.Stages[0].Complete || result.Overview.Condition != presentationmodel.Degraded || result.Overview.Completeness != presentationmodel.CompletenessPartial || result.Overview.Guardian != presentationmodel.GuardianNotObserved {
		t.Fatalf("partial bootstrap was misrepresented: %+v", result)
	}
}

func TestLiveObserveAcceptance(t *testing.T) {
	if os.Getenv("QWSG_LIVE_ACCEPTANCE") != "1" {
		t.Skip("live read-only acceptance is explicitly invoked")
	}
	stateRoot := filepath.Join(t.TempDir(), "state")
	t.Setenv("QWSG_STATE_DIR", stateRoot)
	t.Setenv("QWSG_STORE", "")
	storeRoot, err := observationStoreRoot()
	if err != nil {
		t.Fatal(err)
	}
	definition, err := canonicalcommand.ResolveProfile("observe", canonicalcommand.Selection{Source: "live", Store: storeRoot})
	if err != nil {
		t.Fatal(err)
	}
	for cycle := 0; cycle < 2; cycle++ {
		collect := func(ctx context.Context) (inventory.Snapshot, error) {
			snapshot, collectErr := collectInventoryContext(ctx)
			t.Logf("cycle %d inventory status=%s errors=%d", cycle+1, snapshot.Status, len(snapshot.Errors))
			return snapshot, collectErr
		}
		result, observeErr := observeOnce(context.Background(), definition, collect, func() time.Time { return time.Now().UTC() })
		if observeErr != nil {
			if cycle == 1 && strings.Contains(observeErr.Error(), "invalid observe stage coverage") {
				store, _ := operatorstate.Open(stateRoot)
				current, loadErr := store.Load()
				if loadErr == nil && current.Coverage == operatorstate.CoverageInventorySnapshot && current.Overview.Completeness == presentationmodel.CompletenessPartial {
					t.Skipf("live host inventory is partial; full evaluation correctly remained unavailable: %v", observeErr)
				}
			}
			t.Fatalf("cycle %d: %v", cycle+1, observeErr)
		}
		if (cycle == 0) != result.Bootstrap {
			t.Fatalf("cycle %d bootstrap=%v", cycle+1, result.Bootstrap)
		}
	}
	store, _ := operatorstate.Open(stateRoot)
	current, err := store.Load()
	if err != nil || current.Coverage != operatorstate.CoverageOperatorEvaluation {
		t.Fatalf("current state: coverage=%s err=%v", current.Coverage, err)
	}
}

func TestInteractiveConsoleRefreshRequiresOperatorAction(t *testing.T) {
	var out, errout bytes.Buffer
	if code := runWithConsole([]string{"console"}, strings.NewReader("q\n"), &out, &errout, true); code != 0 {
		t.Fatalf("code=%d err=%q", code, errout.String())
	}
	if !strings.Contains(out.String(), "QWSG Operator Console") {
		t.Fatal("console did not render")
	}
}

func TestConsoleProviderPropagatesCancellation(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	if _, err := (localOverviewProvider{}).Refresh(ctx); err == nil {
		t.Fatal("cancelled refresh unexpectedly succeeded")
	}
}

func TestConsoleRefreshReadsCurrentStateWhileGuardianLockIsHeld(t *testing.T) {
	stateRoot := filepath.Join(t.TempDir(), "state")
	t.Setenv("QWSG_STATE_DIR", stateRoot)
	snapshot := cliFixtureSnapshot()
	definition, err := canonicalcommand.ResolveProfile("check", canonicalcommand.Selection{Source: "live"})
	if err != nil {
		t.Fatal(err)
	}
	orchestrator := pipeline.Orchestrator{Collect: func(context.Context) (inventory.Snapshot, error) { return snapshot, nil }, Retention: inventorystore.DefaultRetention, Rules: pipeline.CanonicalObservationRules()}
	execution, err := orchestrator.Execute(context.Background(), definition)
	if err != nil {
		t.Fatal(err)
	}
	if _, err = publishCheck(definition, execution, snapshot.CompletedAt.Add(time.Second)); err != nil {
		t.Fatal(err)
	}
	statePath := filepath.Join(stateRoot, operatorstate.FileName)
	before, err := os.ReadFile(statePath)
	if err != nil {
		t.Fatal(err)
	}
	lock, err := acquireOneShotLock()
	if err != nil {
		t.Fatal(err)
	}
	defer lock.Release()
	refreshed, err := (localOverviewProvider{}).Refresh(context.Background())
	if err != nil || refreshed.ID == "" {
		t.Fatalf("read-only refresh failed while Guardian lock was held: %v", err)
	}
	after, err := os.ReadFile(statePath)
	if err != nil || !bytes.Equal(before, after) {
		t.Fatalf("refresh changed Current Operator State: %v", err)
	}
	var out, errout bytes.Buffer
	if code := runWithConsole([]string{"console"}, strings.NewReader("r\nq\n"), &out, &errout, true); code != 0 || strings.Contains(out.String(), "Refresh failed") || strings.Contains(out.String(), "guardian_active") {
		t.Fatalf("interactive read-only refresh failed: code=%d out=%q err=%q", code, out.String(), errout.String())
	}
}

func TestCanonicalProfilesAndAdvancedGrammarUseSamePipeline(t *testing.T) {
	root := filepath.Join(t.TempDir(), "store")
	store, err := inventorystore.Open(root, inventorystore.DefaultRetention)
	if err != nil {
		t.Fatal(err)
	}
	from := cliFixtureSnapshot()
	fromName, err := store.Save(from)
	if err != nil {
		t.Fatal(err)
	}
	to := laterCLIFixtureSnapshot(from)
	toName, err := store.Save(to)
	if err != nil {
		t.Fatal(err)
	}
	t.Setenv("QWSG_STORE", root)

	for _, profile := range []string{"changes", "health", "report"} {
		var out, errout bytes.Buffer
		if code := run([]string{profile, "--output", "json", "--presentation", "structured"}, &out, &errout); code != 0 {
			t.Fatalf("%s: code=%d error=%q", profile, code, errout.String())
		}
		var execution canonicalcommand.Execution
		if err := json.Unmarshal(out.Bytes(), &execution); err != nil {
			t.Fatal(err)
		}
		if execution.SchemaName != canonicalcommand.ExecutionSchema ||
			execution.CommandID == "" || execution.PlanID == "" || len(execution.Stages) == 0 {
			t.Fatalf("%s returned invalid execution: %#v", profile, execution)
		}
	}

	var simple, advanced, errout bytes.Buffer
	simpleArgs := []string{"health", "--output", "json", "--presentation", "structured"}
	advancedArgs := []string{
		"analyze", "--source", "store", "--store", root, "--pipeline", "health",
		"--output", "json", "--presentation", "structured",
	}
	if code := run(simpleArgs, &simple, &errout); code != 0 {
		t.Fatalf("simple: %d %s", code, errout.String())
	}
	errout.Reset()
	if code := run(advancedArgs, &advanced, &errout); code != 0 {
		t.Fatalf("advanced: %d %s", code, errout.String())
	}
	var simpleExecution, advancedExecution canonicalcommand.Execution
	if err := json.Unmarshal(simple.Bytes(), &simpleExecution); err != nil {
		t.Fatal(err)
	}
	if err := json.Unmarshal(advanced.Bytes(), &advancedExecution); err != nil {
		t.Fatal(err)
	}
	if len(simpleExecution.Stages) != len(advancedExecution.Stages) {
		t.Fatal("simple and advanced pipeline lengths differ")
	}
	for index := range simpleExecution.Stages {
		if simpleExecution.Stages[index].Stage != advancedExecution.Stages[index].Stage {
			t.Fatal("simple and advanced stage ordering differs")
		}
	}
	if !strings.Contains(simple.String(), fromName) || !strings.Contains(simple.String(), toName) {
		t.Fatal("canonical execution lost snapshot traceability")
	}
}

func TestCanonicalCLIRejectsInvalidAndUnsafeComposition(t *testing.T) {
	for _, tc := range []struct {
		args     []string
		contains string
	}{
		{[]string{"changes"}, "--store is required"},
		{[]string{"health", "--exclude", "compare"}, "not valid for a predefined profile"},
		{[]string{"analyze", "--source", "store", "--store", "/tmp/x", "--pipeline", "health", "--exclude", "compare"}, "required"},
		{[]string{"analyze", "--source", "remote", "--pipeline", "health"}, "live or store"},
		{[]string{"analyze", "--source", "store", "--store", "/tmp/x", "--pipeline", "health", "--presentation", "dashboard"}, "unsupported presentation"},
	} {
		var out, errout bytes.Buffer
		if code := run(tc.args, &out, &errout); code != 1 || !strings.Contains(errout.String(), tc.contains) {
			t.Fatalf("%v: code=%d output=%q", tc.args, code, out.String()+errout.String())
		}
	}
}

func TestInventoryStoreArguments(t *testing.T) {
	absolute := filepath.Join(t.TempDir(), "store")
	for _, tc := range []struct {
		args     []string
		contains string
	}{
		{[]string{"inventory", "load"}, "--store is required"},
		{[]string{"inventory", "load", "--store", "relative"}, "clean absolute path"},
		{[]string{"inventory", "save", "--store", absolute, "--snapshot", "x"}, "--snapshot is valid only"},
		{[]string{"inventory", "load", "--store", absolute, "--retention", "zero"}, "--retention must be an integer"},
		{[]string{"inventory", "other", "--store", absolute}, "unknown inventory subcommand"},
		{[]string{"inventory", "--format", "yaml"}, "--format must be json or human"},
		{[]string{"inventory", "list", "--store", absolute, "--snapshot", "x"}, "--snapshot is valid only"},
	} {
		var out, err bytes.Buffer
		if code := run(tc.args, &out, &err); code != 1 || !strings.Contains(err.String(), tc.contains) {
			t.Fatalf("%v: code=%d output=%q", tc.args, code, out.String()+err.String())
		}
	}
}

func TestInventoryLoadCLIUsesPersistedSnapshot(t *testing.T) {
	root := filepath.Join(t.TempDir(), "store")
	store, err := inventorystore.Open(root, 2)
	if err != nil {
		t.Fatal(err)
	}
	snapshot := cliFixtureSnapshot()
	name, err := store.Save(snapshot)
	if err != nil {
		t.Fatal(err)
	}
	for _, args := range [][]string{
		{"inventory", "load", "--store", root, "--retention", "2"},
		{"inventory", "load", "--store", root, "--retention", "2", "--snapshot", name},
	} {
		var out, errout bytes.Buffer
		if code := run(args, &out, &errout); code != 0 {
			t.Fatalf("%v: code=%d error=%q", args, code, errout.String())
		}
		var got inventory.Snapshot
		if err := json.Unmarshal(out.Bytes(), &got); err != nil {
			t.Fatal(err)
		}
		if got.SnapshotID != snapshot.SnapshotID || got.Canonical.SchemaName != inventory.CanonicalSchemaName {
			t.Fatalf("unexpected loaded snapshot: %#v", got)
		}
	}
}

func TestSnapshotExplorerListInfoAndHumanLoad(t *testing.T) {
	root := filepath.Join(t.TempDir(), "store")
	store, err := inventorystore.Open(root, 2)
	if err != nil {
		t.Fatal(err)
	}
	snapshot := cliFixtureSnapshot()
	name, err := store.Save(snapshot)
	if err != nil {
		t.Fatal(err)
	}
	tests := []struct {
		args     []string
		code     int
		contains []string
	}{
		{[]string{"inventory", "list", "--store", root, "--retention", "2"}, 0, []string{"Snapshots: 1", name, "status=complete"}},
		{[]string{"inventory", "list", "--store", root, "--retention", "2", "--format", "json"}, 0, []string{`"integrity": "verified"`, `"name":`}},
		{[]string{"inventory", "info", "--store", root, "--retention", "2", "--snapshot", name}, 0, []string{"Snapshot:", "Integrity: verified", "Status: complete"}},
		{[]string{"inventory", "load", "--store", root, "--retention", "2", "--format", "human"}, 0, []string{"Stored inventory", "Canonical layers:", "Status: complete"}},
	}
	for _, tc := range tests {
		var out, errout bytes.Buffer
		if code := run(tc.args, &out, &errout); code != tc.code {
			t.Fatalf("%v: code=%d error=%q", tc.args, code, errout.String())
		}
		for _, expected := range tc.contains {
			if !strings.Contains(out.String(), expected) {
				t.Fatalf("%v: output %q does not contain %q", tc.args, out.String(), expected)
			}
		}
	}
}

func TestDocumentedEnvironmentSupportsConciseExplorerCommands(t *testing.T) {
	root := filepath.Join(t.TempDir(), "store")
	store, err := inventorystore.Open(root, inventorystore.DefaultRetention)
	if err != nil {
		t.Fatal(err)
	}
	if _, err := store.Save(cliFixtureSnapshot()); err != nil {
		t.Fatal(err)
	}
	t.Setenv("QWSG_STORE", root)
	t.Setenv("QWSG_FORMAT", "human")
	for _, args := range [][]string{
		{"inventory", "list"},
		{"inventory", "info"},
		{"inventory", "load"},
	} {
		var out, errout bytes.Buffer
		if code := run(args, &out, &errout); code != 0 || out.Len() == 0 {
			t.Fatalf("%v: code=%d out=%q err=%q", args, code, out.String(), errout.String())
		}
	}
}

func TestSnapshotExplorerEmptyAndInvalidStore(t *testing.T) {
	root := filepath.Join(t.TempDir(), "store")
	store, err := inventorystore.Open(root, 2)
	if err != nil {
		t.Fatal(err)
	}
	name, err := store.Save(cliFixtureSnapshot())
	if err != nil {
		t.Fatal(err)
	}
	if err := os.Remove(filepath.Join(root, "snapshots", name)); err != nil {
		t.Fatal(err)
	}
	var out, errout bytes.Buffer
	if code := run([]string{"inventory", "list", "--store", root, "--retention", "2"}, &out, &errout); code != 0 || out.String() != "Snapshots: none\n" {
		t.Fatalf("empty list: code=%d out=%q err=%q", code, out.String(), errout.String())
	}
	out.Reset()
	errout.Reset()
	if code := run([]string{"inventory", "info", "--store", root, "--retention", "2"}, &out, &errout); code != 1 || !strings.Contains(errout.String(), "file does not exist") {
		t.Fatalf("empty info: code=%d out=%q err=%q", code, out.String(), errout.String())
	}
}

func TestHumanOutputEscapesTerminalControls(t *testing.T) {
	input := "name\x1b[31m\n\t\u007f"
	got := safeText(input)
	if strings.ContainsAny(got, "\x1b\n\t\u007f") || got != `name\u001b[31m\u000a\u0009\u007f` {
		t.Fatalf("unsafe rendering: %q", got)
	}
}

func TestJSONCompatibilityOutput(t *testing.T) {
	snapshot := cliFixtureSnapshot()
	var out, errout bytes.Buffer
	if code := writeInventory(&out, &errout, snapshot, formatJSON, "ignored"); code != 0 {
		t.Fatalf("code=%d error=%q", code, errout.String())
	}
	var got inventory.Snapshot
	if err := json.Unmarshal(out.Bytes(), &got); err != nil {
		t.Fatal(err)
	}
	if got.SnapshotID != snapshot.SnapshotID || got.Canonical.SchemaName != inventory.CanonicalSchemaName {
		t.Fatalf("compatibility envelope changed: %#v", got)
	}
}

func TestCompareCLIDefaultExplicitAndHuman(t *testing.T) {
	root := filepath.Join(t.TempDir(), "store")
	store, err := inventorystore.Open(root, 3)
	if err != nil {
		t.Fatal(err)
	}
	from := cliFixtureSnapshot()
	fromName, err := store.Save(from)
	if err != nil {
		t.Fatal(err)
	}
	to := laterCLIFixtureSnapshot(from)
	toName, err := store.Save(to)
	if err != nil {
		t.Fatal(err)
	}

	var first, second, errout bytes.Buffer
	args := []string{"compare", "--store", root, "--retention", "3"}
	if code := run(args, &first, &errout); code != 0 {
		t.Fatalf("default compare: code=%d error=%q", code, errout.String())
	}
	errout.Reset()
	if code := run(args, &second, &errout); code != 0 || first.String() != second.String() {
		t.Fatalf("repeat compare: code=%d equal=%v error=%q", code, first.String() == second.String(), errout.String())
	}
	var result comparison.Result
	if err := json.Unmarshal(first.Bytes(), &result); err != nil {
		t.Fatal(err)
	}
	if result.From.Selector != fromName || result.To.Selector != toName || result.Counts.Added != 2 {
		t.Fatalf("unexpected comparison: %#v", result)
	}

	var human bytes.Buffer
	errout.Reset()
	explicit := []string{"compare", "--store", root, "--retention", "3", "--from", fromName, "--to", toName, "--format", "human"}
	if code := run(explicit, &human, &errout); code != 0 {
		t.Fatalf("human compare: code=%d error=%q", code, errout.String())
	}
	for _, expected := range []string{"Added (2)", "Removed (0)", "Modified (0)", "Unchanged (", ` = "changed"`} {
		if !strings.Contains(human.String(), expected) {
			t.Fatalf("human output %q lacks %q", human.String(), expected)
		}
	}
}

func TestCompareCLIRejectsIncompleteSelectionAndInsufficientHistory(t *testing.T) {
	root := filepath.Join(t.TempDir(), "store")
	store, err := inventorystore.Open(root, 2)
	if err != nil {
		t.Fatal(err)
	}
	name, err := store.Save(cliFixtureSnapshot())
	if err != nil {
		t.Fatal(err)
	}
	for _, tc := range []struct {
		args     []string
		contains string
	}{
		{[]string{"compare", "--store", root, "--retention", "2"}, "at least two"},
		{[]string{"compare", "--store", root, "--retention", "2", "--from", name}, "must be provided together"},
		{[]string{"compare", "--store", root, "--retention", "2", "--from", name, "--to", "missing.json"}, "to snapshot load failed"},
	} {
		var out, errout bytes.Buffer
		if code := run(tc.args, &out, &errout); code != 1 || !strings.Contains(errout.String(), tc.contains) {
			t.Fatalf("%v: code=%d output=%q", tc.args, code, out.String()+errout.String())
		}
	}
}

func laterCLIFixtureSnapshot(snapshot inventory.Snapshot) inventory.Snapshot {
	snapshot.SnapshotID = "cli-fixture-later"
	snapshot.RequestID = "cli-fixture-later"
	snapshot.CompletedAt = snapshot.CompletedAt.Add(time.Minute)
	snapshot.FreshUntil = snapshot.FreshUntil.Add(time.Minute)
	snapshot.Canonical.SnapshotID = snapshot.SnapshotID
	snapshot.Canonical.RequestID = snapshot.RequestID
	snapshot.Canonical.CompletedAt = snapshot.CompletedAt
	snapshot.Canonical.FreshUntil = snapshot.FreshUntil
	layer := &snapshot.Canonical.Layers[0]
	layer.Resources = append(layer.Resources, inventory.Resource{
		ResourceID: "task018:fixture", Kind: "fixture", LayerID: layer.LayerID,
		LifecycleState: "observed", CollectorID: "host", ObservedAt: snapshot.ObservedAt,
		Facts: map[string]inventory.CanonicalFact{
			"task_018_fixture": {
				Value: "changed", ValueType: "string", Quality: "observed", Sensitivity: "public",
				ObservedAt: snapshot.ObservedAt,
				Provenance: inventory.Provenance{
					SourceType: "fixture", SourceLabel: "fixture", ObservedAt: snapshot.ObservedAt,
				},
			},
		},
		Relationships: []inventory.Relationship{}, Labels: map[string]string{},
		Metadata: map[string]string{},
	})
	return snapshot
}

func shiftCLIFixtureSnapshot(snapshot inventory.Snapshot, completedAt time.Time) inventory.Snapshot {
	delta := completedAt.Sub(snapshot.CompletedAt)
	shift := func(value time.Time) time.Time { return value.Add(delta) }
	snapshot.ObservedAt, snapshot.CompletedAt, snapshot.FreshUntil = shift(snapshot.ObservedAt), shift(snapshot.CompletedAt), shift(snapshot.FreshUntil)
	snapshot.Canonical.ObservedAt, snapshot.Canonical.CompletedAt, snapshot.Canonical.FreshUntil = shift(snapshot.Canonical.ObservedAt), shift(snapshot.Canonical.CompletedAt), shift(snapshot.Canonical.FreshUntil)
	for categoryIndex := range snapshot.Categories {
		category := &snapshot.Categories[categoryIndex]
		category.ObservedAt, category.CompletedAt, category.FreshUntil = shift(category.ObservedAt), shift(category.CompletedAt), shift(category.FreshUntil)
		for itemIndex := range category.Items {
			for name, fact := range category.Items[itemIndex].Facts {
				fact.Provenance.ObservedAt = shift(fact.Provenance.ObservedAt)
				category.Items[itemIndex].Facts[name] = fact
			}
		}
	}
	for executionIndex := range snapshot.Canonical.CollectorResults {
		snapshot.Canonical.CollectorResults[executionIndex].Timestamp = shift(snapshot.Canonical.CollectorResults[executionIndex].Timestamp)
	}
	for layerIndex := range snapshot.Canonical.Layers {
		layer := &snapshot.Canonical.Layers[layerIndex]
		layer.ObservedAt, layer.CompletedAt = shift(layer.ObservedAt), shift(layer.CompletedAt)
		for resourceIndex := range layer.Resources {
			resource := &layer.Resources[resourceIndex]
			resource.ObservedAt = shift(resource.ObservedAt)
			for name, fact := range resource.Facts {
				fact.ObservedAt, fact.Provenance.ObservedAt = shift(fact.ObservedAt), shift(fact.Provenance.ObservedAt)
				resource.Facts[name] = fact
			}
		}
	}
	return snapshot
}

func cliFixtureSnapshot() inventory.Snapshot {
	now := time.Date(2026, 7, 24, 4, 0, 0, 0, time.UTC)
	category := inventory.Category{
		CategoryID: "host", ContractVersion: "1.0", Status: inventory.Available,
		ObservedAt: now, CompletedAt: now, FreshUntil: now.Add(time.Minute),
		CollectorID: "host", PrivilegeUsed: "ordinary-user", SourceSummary: []string{"fixture"},
		Items: []inventory.Item{}, Errors: []inventory.InventoryError{}, Redactions: []string{},
	}
	execution := inventory.CollectorExecution{
		CollectorName: "host", Version: "1", Capability: "host", SupportedPlatforms: []string{"linux"},
		Timestamp: now, Status: inventory.Available, Warnings: []inventory.InventoryWarning{},
		Errors: []inventory.InventoryError{}, Metadata: map[string]string{},
	}
	producer := inventory.Producer{ToolVersion: "test", ContractVersion: "1.0"}
	snapshot := inventory.Snapshot{
		SchemaVersion: "1.0", SnapshotID: "cli-fixture", RequestID: "cli-fixture", InstanceID: "subject",
		ObservedAt: now, CompletedAt: now, FreshUntil: now.Add(time.Minute), Status: inventory.Complete,
		Categories: []inventory.Category{category}, Errors: []inventory.InventoryError{},
		Redactions: []string{}, Producer: producer,
	}
	snapshot.Canonical = inventory.AssembleSystemInventory(
		snapshot.Categories, []inventory.CollectorExecution{execution},
		snapshot.SnapshotID, snapshot.RequestID, snapshot.InstanceID,
		now, now, now.Add(time.Minute), 0, producer,
	)
	return snapshot
}

func largeCLIFixtureSnapshot(now time.Time, generation int) inventory.Snapshot {
	items := make([]inventory.Item, 0, 367)
	for index := 0; index < 367; index++ {
		items = append(items, inventory.Item{ID: fmt.Sprintf("component-%04d", index), Kind: "runtime_component", Facts: map[string]inventory.Fact{"value": {Value: index + generation, Quality: "observed", Sensitivity: "operational", Provenance: inventory.Provenance{SourceType: "fixture", SourceLabel: "bounded", ObservedAt: now}}}})
	}
	category := inventory.Category{CategoryID: "components", ContractVersion: "1.0", Status: inventory.Available, ObservedAt: now, CompletedAt: now, FreshUntil: now.Add(5 * time.Minute), CollectorID: "components", PrivilegeUsed: "ordinary-user", SourceSummary: []string{"bounded fixture"}, Items: items, Errors: []inventory.InventoryError{}, Redactions: []string{}}
	execution := inventory.CollectorExecution{CollectorName: "components", Version: "1", Capability: "components", SupportedPlatforms: []string{"linux"}, Timestamp: now, Status: inventory.Available, Warnings: []inventory.InventoryWarning{}, Errors: []inventory.InventoryError{}, Metadata: map[string]string{}}
	producer := inventory.Producer{ToolVersion: "test", ContractVersion: "1.0"}
	id := fmt.Sprintf("large-fixture-%d", generation)
	snapshot := inventory.Snapshot{SchemaVersion: inventory.SchemaVersion, SnapshotID: id, RequestID: id, InstanceID: "subject", ObservedAt: now, CompletedAt: now, FreshUntil: now.Add(5 * time.Minute), Status: inventory.Complete, Categories: []inventory.Category{category}, Errors: []inventory.InventoryError{}, Redactions: []string{}, Producer: producer}
	snapshot.Canonical = inventory.AssembleSystemInventory(snapshot.Categories, []inventory.CollectorExecution{execution}, snapshot.SnapshotID, snapshot.RequestID, snapshot.InstanceID, now, now, snapshot.FreshUntil, 0, producer)
	return snapshot
}

func partialCLIFixtureSnapshot() inventory.Snapshot {
	snapshot := cliFixtureSnapshot()
	now := snapshot.CompletedAt
	category := inventory.Category{CategoryID: "components", ContractVersion: "1.0", Status: inventory.Unsupported, ObservedAt: now, CompletedAt: now, FreshUntil: now.Add(time.Minute), CollectorID: "components", PrivilegeUsed: "ordinary-user", SourceSummary: []string{}, Items: []inventory.Item{}, Errors: []inventory.InventoryError{{Code: "unsupported", CategoryID: "components", Class: string(inventory.Unsupported), MessageKey: "collector.unsupported", OccurredAt: now}}, Redactions: []string{}}
	snapshot.Categories = append(snapshot.Categories, category)
	snapshot.Status = inventory.Partial
	executions := []inventory.CollectorExecution{
		{CollectorName: "components", Version: "1", Capability: "components", SupportedPlatforms: []string{"linux"}, Timestamp: now, Status: inventory.Unsupported, Warnings: []inventory.InventoryWarning{}, Errors: category.Errors, Metadata: map[string]string{}},
		{CollectorName: "host", Version: "1", Capability: "host", SupportedPlatforms: []string{"linux"}, Timestamp: now, Status: inventory.Available, Warnings: []inventory.InventoryWarning{}, Errors: []inventory.InventoryError{}, Metadata: map[string]string{}},
	}
	snapshot.Canonical = inventory.AssembleSystemInventory(snapshot.Categories, executions, snapshot.SnapshotID, snapshot.RequestID, snapshot.InstanceID, snapshot.ObservedAt, snapshot.CompletedAt, snapshot.FreshUntil, 0, snapshot.Producer)
	snapshot.Errors = []inventory.InventoryError{category.Errors[0]}
	return snapshot
}
