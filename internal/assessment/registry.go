package assessment

func Registry() []Requirement {
	registry := []Requirement{
		{ID: "configuration.present", PurposeToken: "configuration_present", Class: ConfigurationRequirement, Disposition: Required, Capability: "configuration", ProbeID: "configuration", Platforms: []string{"ubuntu-24.04-amd64"}, PrivacyClass: "configuration_metadata", Remediations: []Remediation{{PlatformID: "ubuntu-24.04-amd64", Commands: [][]string{{"qwsg", "setup"}}, DisplayCommands: []string{"qwsg setup"}, Revalidate: true}}},
		{ID: "filesystem.local_semantics", PurposeToken: "filesystem_semantics", Class: RecommendedCapability, Disposition: Recommended, Capability: "environment", ProbeID: "filesystem", Platforms: []string{"ubuntu-24.04-amd64"}, PrivacyClass: "operational", Remediations: []Remediation{}, GuidanceRules: filesystemGuidanceRules()},
		{ID: "guardian.canonical_evidence", PurposeToken: "guardian_evidence", Class: EnvironmentCapability, Disposition: Required, Capability: "guardian", ProbeID: "guardian_evidence", Platforms: []string{"ubuntu-24.04-amd64"}, PrivacyClass: "operational", Remediations: []Remediation{}},
		{ID: "guardian.service_active", PurposeToken: "guardian_service_active", Class: EnvironmentCapability, Disposition: Required, Capability: "guardian", ProbeID: "systemd_service_active", Platforms: []string{"ubuntu-24.04-amd64"}, PrivacyClass: "operational", Remediations: []Remediation{{PlatformID: "ubuntu-24.04-amd64", Commands: [][]string{{"systemctl", "--user", "enable", "--now", "qwsg-guardian.service"}}, DisplayCommands: []string{"systemctl --user enable --now qwsg-guardian.service"}, Revalidate: true}}},
		{ID: "guardian.service_enabled", PurposeToken: "guardian_service_enabled", Class: EnvironmentCapability, Disposition: Required, Capability: "guardian", ProbeID: "systemd_service_enabled", Platforms: []string{"ubuntu-24.04-amd64"}, PrivacyClass: "operational", Remediations: []Remediation{{PlatformID: "ubuntu-24.04-amd64", Commands: [][]string{{"systemctl", "--user", "enable", "qwsg-guardian.service"}}, DisplayCommands: []string{"systemctl --user enable qwsg-guardian.service"}, Revalidate: true}}},
		{ID: "guardian.unit_installed", PurposeToken: "guardian_unit_installed", Class: InstallTimeDependency, Disposition: Required, Capability: "guardian", ProbeID: "systemd_unit_installed", Platforms: []string{"ubuntu-24.04-amd64"}, PrivacyClass: "operational", Remediations: []Remediation{{PlatformID: "ubuntu-24.04-amd64", Commands: [][]string{{"sudo", "./install.sh"}}, DisplayCommands: []string{"sudo ./install.sh"}, ElevatedPrivileges: true, Revalidate: true}}},
		{ID: "notification.external", PurposeToken: "external_notification", Class: OptionalFeature, Disposition: Optional, Capability: "notification", ProbeID: "notification", Platforms: []string{"ubuntu-24.04-amd64"}, PrivacyClass: "configuration_metadata", Remediations: []Remediation{}},
		{ID: "platform.architecture", PurposeToken: "supported_architecture", Class: RuntimeDependency, Disposition: Required, Capability: "environment", ProbeID: "go_arch", Platforms: []string{"ubuntu-24.04-amd64"}, PrivacyClass: "operational", Remediations: []Remediation{}},
		{ID: "platform.operating_system", PurposeToken: "supported_operating_system", Class: RuntimeDependency, Disposition: Required, Capability: "environment", ProbeID: "os_release", Platforms: []string{"ubuntu-24.04-amd64"}, PrivacyClass: "operational", Remediations: []Remediation{}},
		{ID: "runtime.glibc", PurposeToken: "glibc_compatible_userspace", Class: RuntimeDependency, Disposition: Required, Capability: "environment", ProbeID: "glibc_version", Platforms: []string{"ubuntu-24.04-amd64"}, PrivacyClass: "operational", Remediations: []Remediation{}},
		{ID: "runtime.non_root", PurposeToken: "ordinary_runtime_user", Class: EnvironmentCapability, Disposition: Required, Capability: "environment", ProbeID: "effective_user", Platforms: []string{"ubuntu-24.04-amd64"}, PrivacyClass: "operational", Remediations: []Remediation{}},
		{ID: "systemd.lingering", PurposeToken: "boot_before_login", Class: RecommendedCapability, Disposition: Recommended, Capability: "guardian", ProbeID: "lingering", Platforms: []string{"ubuntu-24.04-amd64"}, PrivacyClass: "operational", Remediations: []Remediation{}},
		{ID: "systemd.user_manager", PurposeToken: "systemd_user_manager", Class: EnvironmentCapability, Disposition: Required, Capability: "environment", ProbeID: "systemd_user", Platforms: []string{"ubuntu-24.04-amd64"}, PrivacyClass: "operational", Remediations: []Remediation{}, GuidanceRules: userManagerGuidanceRules()},
		{ID: "systemd.version", PurposeToken: "systemd_minimum_version", Class: RuntimeDependency, Disposition: Required, Capability: "environment", ProbeID: "systemd_version", Platforms: []string{"ubuntu-24.04-amd64"}, MinimumVersion: "255", PrivacyClass: "operational", Remediations: []Remediation{}},
	}
	if err := ValidateRegistry(registry); err != nil {
		panic(err)
	}
	return registry
}

func guidance(explanation, effect string, privilege PrivilegeRequirement, manual bool, verify, act []string, notes ...string) Guidance {
	return Guidance{ExplanationToken: explanation, BlockingEffect: effect, VerificationActions: verify, OperatorActions: act, PrivilegeRequirement: privilege, ManualVerification: manual, SafetyNotes: notes, RevalidationAction: "rerun_qwsg_install_check"}
}

func userManagerGuidanceRules() []GuidanceRule {
	rules := []struct {
		classification        Classification
		evidence, explanation string
		privilege             PrivilegeRequirement
		manual                bool
		verify, act           []string
	}{
		{MissingRequired, "systemd_user_runtime_directory_missing", "user_runtime_directory_missing", PrivilegeManualVerification, true, []string{"verify_login_session_runtime"}, []string{"contact_host_administrator_for_user_session"}},
		{UnknownVerification, "systemd_user_runtime_directory_unsafe", "user_runtime_directory_unsafe", PrivilegeManualVerification, true, []string{"verify_runtime_directory_ownership_and_type"}, []string{"contact_host_administrator_for_runtime_directory"}},
		{UnknownVerification, "systemd_user_manager_starting", "user_manager_transient", PrivilegeNone, false, []string{"wait_for_user_manager_startup"}, []string{"retry_assessment_after_bounded_wait"}},
		{MissingRequired, "systemd_user_manager_unavailable", "user_manager_unavailable", PrivilegeManualVerification, true, []string{"verify_systemctl_user_manager"}, []string{"contact_host_administrator_for_user_manager"}},
		{UnknownVerification, "systemd_user_probe_timeout", "user_manager_probe_timeout", PrivilegeNone, false, []string{"retry_systemctl_user_manager"}, []string{"retry_assessment"}},
		{UnknownVerification, "systemd_user_probe_output_limit", "user_manager_probe_output_unbounded", PrivilegeManualVerification, true, []string{"verify_systemctl_user_manager"}, []string{"contact_host_administrator_if_repeated"}},
		{UnknownVerification, "systemd_user_probe_failed", "user_manager_probe_failed", PrivilegeManualVerification, true, []string{"verify_systemctl_user_manager"}, []string{"contact_host_administrator_if_repeated"}},
		{UnknownVerification, "systemd_user_state_unrecognized", "user_manager_state_unrecognized", PrivilegeManualVerification, true, []string{"verify_systemctl_user_manager"}, []string{"contact_host_administrator_for_user_manager"}},
	}
	out := make([]GuidanceRule, 0, len(rules))
	for _, rule := range rules {
		out = append(out, GuidanceRule{PlatformID: "ubuntu-24.04-amd64", Classification: rule.classification, EvidenceToken: rule.evidence, Guidance: guidance(rule.explanation, "blocks_installation", rule.privilege, rule.manual, rule.verify, rule.act, "assessment_does_not_modify_host", "do_not_guess_systemd_repair")})
	}
	return out
}

func filesystemGuidanceRules() []GuidanceRule {
	evidence := []string{"filesystem_type_unknown", "filesystem_remote_or_overlay", "filesystem_path_unavailable", "filesystem_path_unsafe"}
	out := make([]GuidanceRule, 0, len(evidence))
	for _, token := range evidence {
		out = append(out, GuidanceRule{PlatformID: "ubuntu-24.04-amd64", Classification: UnknownVerification, EvidenceToken: token, Guidance: guidance(token, "reduces_operational_confidence", PrivilegeManualVerification, true, []string{"verify_local_unix_filesystem_semantics"}, []string{"confirm_atomic_rename_flock_and_private_modes"}, "assessment_remains_read_only", "no_install_command_is_proven")})
	}
	return out
}

func RequirementByID(id string) (Requirement, bool) {
	for _, requirement := range Registry() {
		if requirement.ID == id {
			return requirement, true
		}
	}
	return Requirement{}, false
}

func Recommendation(id, platformID string) *Remediation {
	requirement, ok := RequirementByID(id)
	if !ok {
		return nil
	}
	for _, remediation := range requirement.Remediations {
		if remediation.PlatformID == platformID {
			copy := remediation
			return &copy
		}
	}
	return nil
}

func GuidanceFor(id, platformID string, classification Classification, evidence string) *Guidance {
	requirement, ok := RequirementByID(id)
	if !ok {
		return nil
	}
	for _, rule := range requirement.GuidanceRules {
		if rule.PlatformID == platformID && rule.Classification == classification && rule.EvidenceToken == evidence {
			copy := rule.Guidance
			return &copy
		}
	}
	return nil
}

// AttachGuidance selects only exact supported-platform registry mappings.
func AttachGuidance(report *Report) {
	if !report.Platform.Supported {
		return
	}
	for index := range report.Findings {
		finding := &report.Findings[index]
		if finding.Classification == Satisfied {
			continue
		}
		finding.Remediation = Recommendation(finding.RequirementID, report.Platform.ID)
		finding.Guidance = GuidanceFor(finding.RequirementID, report.Platform.ID, finding.Classification, finding.EvidenceToken)
	}
}
