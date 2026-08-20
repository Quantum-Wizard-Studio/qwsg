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
	assessment.AttachGuidance(report)
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
			if finding.Guidance != nil {
				fmt.Fprintf(out, "  Explanation: %s\n", guidanceText(finding.Guidance.ExplanationToken))
				fmt.Fprintf(out, "  Impact: %s\n", guidanceText(finding.Guidance.BlockingEffect))
				for _, action := range finding.Guidance.VerificationActions {
					fmt.Fprintf(out, "  Verify: %s\n", guidanceText(action))
				}
				for _, action := range finding.Guidance.OperatorActions {
					fmt.Fprintf(out, "  Operator action: %s\n", guidanceText(action))
				}
				fmt.Fprintf(out, "  Privileges: %s\n", guidanceText(string(finding.Guidance.PrivilegeRequirement)))
				if finding.Guidance.ManualVerification {
					fmt.Fprintln(out, "  Manual verification: required")
				}
				for _, note := range finding.Guidance.SafetyNotes {
					fmt.Fprintf(out, "  Safety: %s\n", guidanceText(note))
				}
				fmt.Fprintf(out, "  Revalidate: %s\n", guidanceText(finding.Guidance.RevalidationAction))
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

func guidanceText(token string) string {
	messages := map[string]string{
		"administrator":                                    "administrator privileges are required for any host change",
		"manual_verification":                              "manual verification; administrator help may be required",
		"none":                                             "ordinary-user action",
		"blocks_installation":                              "installation and activation cannot continue",
		"reduces_operational_confidence":                   "non-blocking, but readiness remains partial until verified",
		"user_runtime_directory_missing":                   "the ordinary user's systemd runtime directory is unavailable",
		"user_runtime_directory_unsafe":                    "the ordinary user's systemd runtime directory failed ownership, type, or permission checks",
		"user_manager_transient":                           "the user systemd manager has not reached a stable running state",
		"user_manager_unavailable":                         "the user systemd manager could not be reached with validated session context",
		"user_manager_probe_timeout":                       "the bounded user-manager verification timed out",
		"user_manager_probe_output_unbounded":              "the user-manager response exceeded the safe output limit",
		"user_manager_probe_failed":                        "the user-manager verification failed with an ambiguous state",
		"user_manager_state_unrecognized":                  "the user manager returned a state QWSG cannot safely classify",
		"verify_login_session_runtime":                     "log in as the intended non-root QWSG user and confirm that a normal user session is available",
		"verify_runtime_directory_ownership_and_type":      "have the host administrator verify the user's runtime directory is a private, user-owned directory",
		"wait_for_user_manager_startup":                    "wait briefly for the user manager to finish starting",
		"verify_systemctl_user_manager":                    "as the intended QWSG user, run: systemctl --user is-system-running",
		"retry_systemctl_user_manager":                     "retry: systemctl --user is-system-running",
		"contact_host_administrator_for_user_session":      "ask the host administrator to restore a normal systemd user session; QWSG cannot safely select a repair command from this evidence",
		"contact_host_administrator_for_runtime_directory": "ask the host administrator to inspect the user-session runtime directory; do not replace or chmod it based only on this result",
		"retry_assessment_after_bounded_wait":              "rerun the assessment after the user manager reaches a stable state",
		"contact_host_administrator_for_user_manager":      "ask the host administrator to verify the systemd user-session infrastructure; no exact repair command is proven",
		"retry_assessment":                                 "rerun the assessment once; seek administrator verification if the result repeats",
		"contact_host_administrator_if_repeated":           "if the bounded probe fails again, ask the host administrator to verify the user manager",
		"assessment_does_not_modify_host":                  "QWSG has not changed the user session or systemd configuration",
		"do_not_guess_systemd_repair":                      "do not apply package, PAM, lingering, or service changes without identifying the cause",
		"filesystem_type_unknown":                          "QWSG cannot prove the required local filesystem semantics from the available read-only evidence",
		"filesystem_remote_or_overlay":                     "the QWSG data path appears to use a remote, overlay, or unproven filesystem type",
		"filesystem_path_unavailable":                      "the QWSG data-path ancestor could not be inspected safely",
		"filesystem_path_unsafe":                           "the QWSG data-path ancestor is not a safe user-owned directory",
		"verify_local_unix_filesystem_semantics":           "verify that QWSG configuration and state reside on a local Unix filesystem supporting atomic rename, advisory flock, ownership, and private modes",
		"confirm_atomic_rename_flock_and_private_modes":    "confirm those semantics with the host or storage administrator before unattended operation",
		"assessment_remains_read_only":                     "the default assessment performs no filesystem write test",
		"no_install_command_is_proven":                     "no package or host-change command is justified by this evidence",
		"rerun_qwsg_install_check":                         "qwsg install --check",
	}
	if value, ok := messages[token]; ok {
		return value
	}
	return safeText(token)
}
