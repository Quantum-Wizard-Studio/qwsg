package main

import (
	"context"
	"errors"
	"fmt"
	"io"
	"os"
	"os/user"
	"path/filepath"
	"time"

	"quantumwizard.hu/qwsg/internal/assessment"
	"quantumwizard.hu/qwsg/internal/configurationstore"
	"quantumwizard.hu/qwsg/internal/credentialstore"
	"quantumwizard.hu/qwsg/internal/guardian"
	"quantumwizard.hu/qwsg/internal/notification"
	"quantumwizard.hu/qwsg/internal/operatorstate"
	"quantumwizard.hu/qwsg/internal/presentationmodel"
	"quantumwizard.hu/qwsg/internal/smtpnotification"
)

func runInstallAssessment(args []string, out, errout io.Writer) int {
	if len(args) == 0 || args[0] != "--check" {
		if len(args) == 0 || isHelp(args[0]) {
			fmt.Fprintln(out, "Usage: qwsg install --check [--format human|json]")
			return 0
		}
		return usageError(errout, "install supports only --check")
	}
	format, err := parseAssessmentFormat(args[1:])
	if err != nil {
		return usageError(errout, "%v", err)
	}
	host := assessment.LocalHost{Runner: assessment.DefaultRunner()}
	report := assessment.AssessInstall(context.Background(), host, time.Now().UTC())
	attachRecommendations(&report)
	return writeAssessment(out, errout, report, format)
}

func runReadiness(args []string, out, errout io.Writer) int {
	if len(args) > 0 && isHelp(args[0]) {
		fmt.Fprintln(out, "Usage: qwsg readiness [--format human|json]")
		return 0
	}
	format, err := parseAssessmentFormat(args)
	if err != nil {
		return usageError(errout, "%v", err)
	}
	host := assessment.LocalHost{Runner: assessment.DefaultRunner()}
	report := buildOperationalReport(context.Background(), host, time.Now().UTC())
	return writeAssessment(out, errout, report, format)
}

func buildOperationalReport(ctx context.Context, host assessment.Host, now time.Time) assessment.Report {
	report := assessment.AssessInstall(ctx, host, now)
	report.Phase = "operational"
	report.Findings = append(report.Findings, configurationFinding(), notificationFinding(), serviceFinding(ctx, host, "guardian.unit_installed", "systemd_unit_installed"), serviceFinding(ctx, host, "guardian.service_enabled", "systemd_service_enabled"), serviceFinding(ctx, host, "guardian.service_active", "systemd_service_active"), lingeringFinding(), guardianEvidenceFinding())
	assessment.SortFindings(report.Findings)
	attachRecommendations(&report)
	report.Domains = operationalDomains(report.Findings)
	report.NextActions = operationalNextActions(report.Domains)
	return report
}

func parseAssessmentFormat(args []string) (string, error) {
	format := formatHuman
	for i := 0; i < len(args); i++ {
		if args[i] != "--format" || i+1 >= len(args) || !validFormat(args[i+1]) {
			return "", fmt.Errorf("assessment accepts only --format human|json")
		}
		format = args[i+1]
		i++
	}
	return format, nil
}

func configurationFinding() assessment.Finding {
	path, err := configurationstore.SelectPath("", os.Getenv)
	if err != nil {
		return assessment.Finding{RequirementID: "configuration.present", Classification: assessment.Incompatible, EvidenceToken: "configuration_path_unsafe"}
	}
	source, found, err := configurationstore.Load(path)
	if err != nil {
		return assessment.Finding{RequirementID: "configuration.present", Classification: assessment.Incompatible, EvidenceToken: "configuration_invalid"}
	}
	if !found {
		return assessment.Finding{RequirementID: "configuration.present", Classification: assessment.MissingRequired, EvidenceToken: "configuration_missing"}
	}
	effective, err := resolveLocalConfiguration(source, true, nil)
	if err != nil || validateGuardianTiming(effective) != nil || validateNotificationReadiness(effective, path) != nil {
		return assessment.Finding{RequirementID: "configuration.present", Classification: assessment.Incompatible, EvidenceToken: "configuration_invalid"}
	}
	return assessment.Finding{RequirementID: "configuration.present", Classification: assessment.Satisfied, EvidenceToken: "configuration_valid"}
}

func notificationFinding() assessment.Finding {
	path, err := configurationstore.SelectPath("", os.Getenv)
	if err != nil {
		return assessment.Finding{RequirementID: "notification.external", Classification: assessment.MissingOptional, EvidenceToken: "notification_configuration_unavailable"}
	}
	source, found, err := configurationstore.Load(path)
	if err != nil || !found {
		return assessment.Finding{RequirementID: "notification.external", Classification: assessment.MissingOptional, EvidenceToken: "notification_not_configured"}
	}
	effective, err := resolveLocalConfiguration(source, true, nil)
	if err != nil {
		return assessment.Finding{RequirementID: "notification.external", Classification: assessment.MissingOptional, EvidenceToken: "notification_configuration_invalid"}
	}
	cfg, err := smtpnotification.FromEffective(effective)
	if err != nil {
		return assessment.Finding{RequirementID: "notification.external", Classification: assessment.MissingOptional, EvidenceToken: "notification_configuration_invalid"}
	}
	if !cfg.Enabled {
		return assessment.Finding{RequirementID: "notification.external", Classification: assessment.MissingOptional, EvidenceToken: "notification_disabled"}
	}
	available := cfg.Auth != "password"
	if cfg.Auth == "password" {
		_, err = credentialstore.Load(path, cfg.CredentialRef)
		available = err == nil
	}
	findings := smtpnotification.Preflight(cfg, available)
	for _, finding := range findings {
		if finding.Classification == assessment.MissingRequired || finding.Classification == assessment.Incompatible {
			return assessment.Finding{RequirementID: "notification.external", Classification: assessment.MissingOptional, EvidenceToken: "notification_not_ready"}
		}
		if finding.Classification == assessment.UnknownVerification {
			return notificationDeliveryEvidence()
		}
	}
	return notificationDeliveryEvidence()
}

func notificationDeliveryEvidence() assessment.Finding {
	root, err := localStateRoot()
	if err != nil {
		return assessment.Finding{RequirementID: "notification.external", Classification: assessment.UnknownVerification, EvidenceToken: "notification_configured_unverified"}
	}
	store, err := guardian.OpenStore(filepath.Join(root, "guardian"))
	if err != nil {
		return assessment.Finding{RequirementID: "notification.external", Classification: assessment.UnknownVerification, EvidenceToken: "notification_configured_unverified"}
	}
	checkpoint, err := store.Load()
	if err != nil {
		return assessment.Finding{RequirementID: "notification.external", Classification: assessment.UnknownVerification, EvidenceToken: "notification_configured_unverified"}
	}
	failure := false
	for _, entry := range checkpoint.NotificationQueueState.Entries {
		if entry.Status == notification.StatusAccepted || entry.Status == notification.StatusDelivered {
			return assessment.Finding{RequirementID: "notification.external", Classification: assessment.Satisfied, EvidenceToken: "notification_delivery_verified"}
		}
		if entry.Status == notification.StatusRetryableFailure || entry.Status == notification.StatusTerminalFailure || entry.Status == notification.StatusExhausted {
			failure = true
		}
	}
	if failure {
		return assessment.Finding{RequirementID: "notification.external", Classification: assessment.MissingOptional, EvidenceToken: "notification_delivery_failed"}
	}
	return assessment.Finding{RequirementID: "notification.external", Classification: assessment.UnknownVerification, EvidenceToken: "notification_configured_unverified"}
}

func serviceFinding(ctx context.Context, host assessment.Host, requirementID, probeID string) assessment.Finding {
	result, err := host.Run(ctx, probeID)
	if err == nil {
		return assessment.Finding{RequirementID: requirementID, Classification: assessment.Satisfied, EvidenceToken: probeID + "_confirmed"}
	}
	_ = result
	return assessment.Finding{RequirementID: requirementID, Classification: assessment.MissingRequired, EvidenceToken: probeID + "_not_confirmed"}
}

func guardianEvidenceFinding() assessment.Finding {
	root, err := localStateRoot()
	if err != nil {
		return assessment.Finding{RequirementID: "guardian.canonical_evidence", Classification: assessment.UnknownVerification, EvidenceToken: "guardian_state_root_unavailable"}
	}
	store, err := operatorstate.Open(root)
	if err != nil {
		return assessment.Finding{RequirementID: "guardian.canonical_evidence", Classification: assessment.UnknownVerification, EvidenceToken: "guardian_state_unsafe"}
	}
	state, err := store.Load()
	if errors.Is(err, operatorstate.ErrMissing) {
		return assessment.Finding{RequirementID: "guardian.canonical_evidence", Classification: assessment.MissingRequired, EvidenceToken: "guardian_evidence_missing"}
	}
	if err != nil {
		return assessment.Finding{RequirementID: "guardian.canonical_evidence", Classification: assessment.UnknownVerification, EvidenceToken: "guardian_evidence_invalid"}
	}
	if state.Overview.Guardian == presentationmodel.GuardianRunning && state.Overview.Freshness == presentationmodel.FreshnessCurrent {
		return assessment.Finding{RequirementID: "guardian.canonical_evidence", Classification: assessment.Satisfied, EvidenceToken: "guardian_running_fresh"}
	}
	return assessment.Finding{RequirementID: "guardian.canonical_evidence", Classification: assessment.MissingRequired, EvidenceToken: "guardian_not_running_fresh"}
}

func lingeringFinding() assessment.Finding {
	current, err := user.Current()
	if err != nil || current.Username == "" || filepath.Base(current.Username) != current.Username {
		return assessment.Finding{RequirementID: "systemd.lingering", Classification: assessment.UnknownVerification, EvidenceToken: "lingering_state_unavailable"}
	}
	info, err := os.Lstat(filepath.Join("/var/lib/systemd/linger", current.Username))
	if os.IsNotExist(err) {
		return assessment.Finding{RequirementID: "systemd.lingering", Classification: assessment.MissingOptional, EvidenceToken: "lingering_disabled"}
	}
	if err != nil || info.Mode()&os.ModeSymlink != 0 || !info.Mode().IsRegular() {
		return assessment.Finding{RequirementID: "systemd.lingering", Classification: assessment.UnknownVerification, EvidenceToken: "lingering_state_unavailable"}
	}
	return assessment.Finding{RequirementID: "systemd.lingering", Classification: assessment.Satisfied, EvidenceToken: "lingering_enabled"}
}

func attachRecommendations(report *assessment.Report) {
	if !report.Platform.Supported {
		return
	}
	for index := range report.Findings {
		if report.Findings[index].Classification == assessment.Satisfied {
			continue
		}
		report.Findings[index].Remediation = assessment.Recommendation(report.Findings[index].RequirementID, report.Platform.ID)
	}
}

func operationalDomains(findings []assessment.Finding) []assessment.DomainSummary {
	groups := map[string][]assessment.Finding{}
	for _, finding := range findings {
		domain := "environment_dependencies"
		switch finding.RequirementID {
		case "configuration.present":
			domain = "configuration"
		case "notification.external":
			domain = "notification"
		case "guardian.canonical_evidence", "guardian.service_active", "guardian.service_enabled", "guardian.unit_installed", "systemd.lingering":
			domain = "guardian_service"
		}
		groups[domain] = append(groups[domain], finding)
	}
	domains := []assessment.DomainSummary{{Domain: "installation", State: assessment.Summarize(groups["environment_dependencies"])}, {Domain: "environment_dependencies", State: assessment.Summarize(groups["environment_dependencies"])}, {Domain: "configuration", State: assessment.Summarize(groups["configuration"])}, {Domain: "notification", State: assessment.Summarize(groups["notification"])}, {Domain: "guardian_service", State: assessment.Summarize(groups["guardian_service"])}}
	overall := assessment.Ready
	for _, domain := range domains {
		if domain.Domain == "notification" && domain.State != assessment.Ready {
			if overall == assessment.Ready {
				overall = assessment.Partial
			}
			continue
		}
		if domain.State == assessment.NotReady {
			overall = assessment.NotReady
			break
		}
		if domain.State == assessment.Unknown && overall != assessment.NotReady {
			overall = assessment.Unknown
		}
		if domain.State == assessment.Partial && overall == assessment.Ready {
			overall = assessment.Partial
		}
	}
	return append(domains, assessment.DomainSummary{Domain: "overall", State: overall})
}

func operationalNextActions(domains []assessment.DomainSummary) []string {
	for _, domain := range domains {
		if domain.Domain == "overall" && domain.State == assessment.Ready {
			return []string{"no_action_required"}
		}
	}
	return []string{"resolve_required_findings", "configure_optional_notification", "rerun_qwsg_readiness"}
}

func writeAssessment(out, errout io.Writer, report assessment.Report, format string) int {
	if err := assessment.ValidateReport(report); err != nil {
		fmt.Fprintln(errout, "assessment failed: assessment_invalid")
		return 5
	}
	if format == formatJSON {
		if writeJSON(out, errout, report) != 0 {
			return 5
		}
	} else {
		fmt.Fprintf(out, "QWSG %s readiness\nPlatform: %s\n", report.Phase, safeText(report.Platform.ID))
		for _, finding := range report.Findings {
			fmt.Fprintf(out, "%-34s %s", safeText(finding.RequirementID), finding.Classification)
			if finding.ObservedValue != "" {
				fmt.Fprintf(out, " (%s)", safeText(finding.ObservedValue))
			}
			fmt.Fprintln(out)
			if finding.Remediation != nil {
				for _, command := range finding.Remediation.DisplayCommands {
					fmt.Fprintf(out, "  Recommended: %s\n", safeText(command))
				}
				fmt.Fprintln(out, "  Revalidate after operator action.")
			}
		}
		for _, domain := range report.Domains {
			fmt.Fprintf(out, "%s: %s\n", safeText(domain.Domain), domain.State)
		}
		if len(report.NextActions) > 0 {
			fmt.Fprintln(out, "Next actions:")
			for _, action := range report.NextActions {
				fmt.Fprintf(out, "- %s\n", safeText(action))
			}
		}
	}
	for _, domain := range report.Domains {
		if domain.Domain == "overall" || report.Phase == "install" && domain.Domain == "installation" {
			switch domain.State {
			case assessment.Ready, assessment.Partial:
				return 0
			case assessment.NotReady:
				return 4
			default:
				return 4
			}
		}
	}
	return 5
}
