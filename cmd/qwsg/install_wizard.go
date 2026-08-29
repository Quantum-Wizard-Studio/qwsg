package main

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
	"time"

	"quantumwizard.hu/qwsg/internal/assessment"
	"quantumwizard.hu/qwsg/internal/changenotification"
	"quantumwizard.hu/qwsg/internal/configurationstore"
	"quantumwizard.hu/qwsg/internal/installation"
	"quantumwizard.hu/qwsg/internal/installer"
	"quantumwizard.hu/qwsg/internal/userservice"
)

var wizardInstallPackage = installReleasePackage
var wizardClassifyInstallation = func(candidate string) installation.Result {
	return installation.Classify(installation.Options{Root: "/", CandidateVersion: candidate})
}
var wizardEvidenceProbe guardianEvidenceProbe = currentGuardianEvidence
var wizardEvidencePause guardianEvidencePause = guardianEvidenceSleep
var wizardOperationalReport = func() assessment.Report {
	return buildOperationalReport(context.Background(), assessment.LocalHost{Runner: assessment.DefaultRunner()}, time.Now().UTC())
}

func runInstall(args []string, in io.Reader, out, errout io.Writer, interactive bool) int {
	if len(args) > 0 && args[0] == "--check" {
		return runInstallAssessment(args, out, errout)
	}
	if len(args) > 0 && isHelp(args[0]) {
		fmt.Fprintln(out, "Usage: qwsg install [--guided|--check] [--language en|hu|de] [--line-mode]")
		return 0
	}
	language, lineMode := installer.English, false
	for i := 0; i < len(args); i++ {
		switch args[i] {
		case "--guided":
		case "--line-mode":
			lineMode = true
		case "--language":
			i++
			if i >= len(args) {
				return usageError(errout, "--language requires en, hu, or de")
			}
			var ok bool
			language, ok = installer.ParseLanguage(args[i])
			if !ok {
				return usageError(errout, "--language requires en, hu, or de")
			}
		default:
			return usageError(errout, "install supports only --guided, --check, --language en|hu|de, and --line-mode")
		}
	}
	if !interactive {
		return usageError(errout, "guided installation requires a terminal; use --check for assessment")
	}
	reader := bufio.NewReader(in)
	if !hasArg(args, "--language") {
		fmt.Fprintln(out, "Quantum Wizard Server Guardian — Installation Wizard")
		fmt.Fprintln(out, "Choose language / Válasszon nyelvet / Sprache wählen: 1 English, 2 Magyar, 3 Deutsch [1]")
		choice, _ := reader.ReadString('\n')
		switch strings.TrimSpace(choice) {
		case "2":
			language = installer.Hungarian
		case "3":
			language = installer.German
		}
	}
	catalog := installer.Catalog{Language: language}
	progress := installer.NewProgress()
	render := func(phase installer.PhaseID) {
		_ = progress.Start(phase)
		if !lineMode && os.Getenv("TERM") != "" && os.Getenv("TERM") != "dumb" {
			fmt.Fprint(out, "\033[2J\033[H")
		}
		fmt.Fprintf(out, "%s\n[%s] %d%%\n%s: %s\n\n%s\n\n", catalog.Text("title"), progressBar(progress.Percent()), progress.Percent(), catalog.Text("stage"), catalog.Text(installer.MessageID(phase)), catalog.Text(installer.MessageID(string(phase)+".help")))
	}
	complete := func(phase installer.PhaseID) { _ = progress.Complete(phase) }

	render(installer.PhasePreflight)
	f, err := os.Open("/etc/os-release")
	if err != nil {
		fmt.Fprintln(errout, catalog.Text("unsupported", "unknown"))
		return 3
	}
	platform, err := installer.Detect(f, runtime.GOARCH)
	_ = f.Close()
	if err != nil || !platform.Supported() {
		progress.Fail(installer.PhasePreflight, false)
		fmt.Fprintln(errout, catalog.Text("unsupported", platform.String()))
		return 3
	}
	report := assessment.AssessInstall(context.Background(), assessment.LocalHost{Runner: assessment.DefaultRunner()}, time.Now().UTC())
	blocked := false
	for _, domain := range report.Domains {
		if domain.Domain == "installation" && domain.State == assessment.NotReady {
			blocked = true
		}
	}
	if blocked {
		progress.Fail(installer.PhasePreflight, false)
		return writeAssessment(out, errout, report, formatHuman)
	}
	complete(installer.PhasePreflight)

	render(installer.PhasePlan)
	classification := wizardClassifyInstallation(version)
	fresh, modeID, refusalID := guidedPackageDecision(classification)
	mode := catalog.Text(modeID)
	if refusalID != "" {
		progress.Fail(installer.PhasePlan, false)
		fmt.Fprintln(out, catalog.Text("plan.mode", mode))
		fmt.Fprintln(errout, catalog.Text(refusalID))
		return 1
	}
	fmt.Fprintln(out, catalog.Text("plan.mode", mode))
	fmt.Fprintln(out, catalog.Text("plan.package"))
	fmt.Fprintln(out, catalog.Text("plan.data"))
	fmt.Fprintln(out, catalog.Text("plan.service"))
	fmt.Fprint(out, catalog.Text("plan.confirm")+" ")
	answer, _ := reader.ReadString('\n')
	if !yes(answer, language, false) {
		fmt.Fprintln(out, catalog.Text("cancelled"))
		return 0
	}
	complete(installer.PhasePlan)

	render(installer.PhaseInstall)
	if fresh {
		if err = wizardInstallPackage(); err != nil {
			progress.Fail(installer.PhaseInstall, false)
			fmt.Fprintln(errout, catalog.Text("install.failure"))
			return 1
		}
	} else {
		fmt.Fprintln(out, catalog.Text("install.existing"))
	}
	complete(installer.PhaseInstall)

	render(installer.PhaseConfiguration)
	var legacyOut, legacyErr strings.Builder
	if code := runSetup([]string{"--accept-defaults", "--set", "locale=" + string(language)}, reader, &legacyOut, &legacyErr, true); code != 0 {
		progress.Fail(installer.PhaseConfiguration, false)
		fmt.Fprintln(errout, catalog.Text("configuration.failure"))
		return code
	}
	fmt.Fprintln(out, catalog.Text("configuration.done"))
	complete(installer.PhaseConfiguration)
	installOutcome := changenotification.Failed
	installOperationID := "installation-" + time.Now().UTC().Format("20060102T150405.000000000Z")
	defer func() {
		reason := ""
		if installOutcome == changenotification.Failed {
			reason = "installation_failed"
		}
		managedChangeDelivery(managedEvent(changenotification.Install, installOutcome, installOperationID, "", version, reason), errout)
	}()

	render(installer.PhaseNotification)
	fmt.Fprint(out, catalog.Text("notification.prompt")+" ")
	answer, _ = reader.ReadString('\n')
	if yes(answer, language, false) {
		fmt.Fprintln(out, catalog.Text("notification.later"))
	} else {
		fmt.Fprintln(out, catalog.Text("notification.skipped"))
	}
	complete(installer.PhaseNotification)

	render(installer.PhaseUpdatePolicy)
	fmt.Fprint(out, catalog.Text("update_policy.prompt")+" ")
	answer, _ = reader.ReadString('\n')
	policy := "manual"
	if strings.TrimSpace(answer) == "2" {
		policy = "notify"
	}
	legacyOut.Reset()
	legacyErr.Reset()
	if code := runConfig([]string{"set", "update.policy", policy}, &legacyOut, &legacyErr); code != 0 {
		progress.Fail(installer.PhaseUpdatePolicy, false)
		return code
	}
	fmt.Fprintln(out, catalog.Text("update_policy.saved", policy))
	complete(installer.PhaseUpdatePolicy)

	render(installer.PhaseActivation)
	fmt.Fprint(out, catalog.Text("activation.prompt")+" ")
	answer, _ = reader.ReadString('\n')
	active := yes(answer, language, true)
	previousEvidence, _ := wizardEvidenceProbe()
	if active {
		if err = prepareGuardianState(); err == nil {
			err = userservice.New().Activate(context.Background())
		}
		if err != nil {
			progress.Fail(installer.PhaseActivation, false)
			fmt.Fprintln(errout, catalog.Text("activation.failure"))
			return 1
		}
		managedChangeDelivery(managedEvent(changenotification.GuardianChanged, changenotification.Success, installOperationID+"-guardian", "inactive", "active", ""), out)
	} else {
		fmt.Fprintln(out, catalog.Text("activation.skipped"))
	}
	complete(installer.PhaseActivation)
	render(installer.PhaseReadiness)
	if active {
		path, pathErr := configurationstore.SelectPath("", os.Getenv)
		if pathErr != nil {
			progress.Fail(installer.PhaseReadiness, false)
			fmt.Fprintln(errout, catalog.Text("readiness.timeout"))
			return 4
		}
		source, found, loadErr := configurationstore.Load(path)
		effective, configErr := resolveLocalConfiguration(source, found, nil)
		if loadErr != nil || configErr != nil || waitForFreshGuardianEvidence(context.Background(), previousEvidence, guardianTimeout(effective)+5*time.Second, wizardEvidenceProbe, wizardEvidencePause) != nil {
			progress.Fail(installer.PhaseReadiness, false)
			fmt.Fprintln(errout, catalog.Text("readiness.timeout"))
			return 4
		}
	}

	report = wizardOperationalReport()
	code := finishGuidedInstallation(catalog, &progress, active, policy, report, out, errout, render, complete)
	if code == 0 {
		installOutcome = changenotification.Success
	}
	return code
}

func guidedPackageDecision(result installation.Result) (bool, installer.MessageID, installer.MessageID) {
	switch result.State {
	case installation.NoInstallation:
		return true, "plan.mode.fresh", ""
	case installation.VerifiedSupported:
		return false, "plan.mode.existing", ""
	case installation.SupportedUpgradeSource:
		return false, "plan.mode.upgrade", "install.upgrade"
	case installation.LegacyInstallation:
		return false, "plan.mode.legacy", "install.legacy"
	case installation.InconsistentInstallation:
		return false, "plan.mode.inconsistent", "install.inconsistent"
	default:
		return false, "plan.mode.unknown", "install.unknown"
	}
}

func finishGuidedInstallation(catalog installer.Catalog, progress *installer.Progress, active bool, policy string, report assessment.Report, out, errout io.Writer, render func(installer.PhaseID), complete func(installer.PhaseID)) int {
	readinessState := assessment.Unknown
	for _, domain := range report.Domains {
		if domain.Domain == "overall" {
			readinessState = domain.State
		}
	}
	if active && readinessState == assessment.NotReady {
		progress.Fail(installer.PhaseReadiness, false)
		fmt.Fprintln(errout, catalog.Text("readiness.partial"))
		return 4
	}
	if readinessState == assessment.Ready {
		fmt.Fprintln(out, catalog.Text("readiness.ready"))
	} else {
		fmt.Fprintln(out, catalog.Text("readiness.partial"))
	}
	complete(installer.PhaseReadiness)
	render(installer.PhaseCompletion)
	complete(installer.PhaseCompletion)
	guardian := catalog.Text("summary.inactive")
	if active {
		guardian = catalog.Text("summary.active")
	}
	readiness := catalog.Text("summary.partial")
	if readinessState == assessment.Ready {
		readiness = catalog.Text("summary.ready")
	}
	fmt.Fprintf(out, "[%s] 100%%\n%s\n%s\n%s\n%s\n%s\n%s\n", progressBar(100), catalog.Text("summary.version", version), catalog.Text("summary.guardian", guardian), catalog.Text("summary.notification"), catalog.Text("summary.policy", policy), catalog.Text("summary.readiness", readiness), catalog.Text("next"))
	return 0
}

func installReleasePackage() error {
	executable, err := os.Executable()
	if err != nil {
		return err
	}
	root := filepath.Dir(filepath.Dir(executable))
	script := filepath.Join(root, "install.sh")
	info, err := os.Lstat(script)
	if err != nil || !info.Mode().IsRegular() || info.Mode()&os.ModeSymlink != 0 {
		return fmt.Errorf("verified release install.sh is unavailable; run the wizard from an extracted QWSG release archive")
	}
	cmd := exec.Command("sudo", script)
	cmd.Stdin, cmd.Stdout, cmd.Stderr = os.Stdin, os.Stdout, os.Stderr
	return cmd.Run()
}

func progressBar(percent int) string {
	filled := percent / 5
	return strings.Repeat("█", filled) + strings.Repeat("░", 20-filled)
}
func yes(value string, language installer.Language, defaultYes bool) bool {
	v := strings.ToLower(strings.TrimSpace(value))
	if v == "" {
		return defaultYes
	}
	return v == "y" || v == "yes" || v == "i" || v == "igen" || v == "j" || v == "ja"
}
func hasArg(args []string, target string) bool {
	for _, arg := range args {
		if arg == target {
			return true
		}
	}
	return false
}
