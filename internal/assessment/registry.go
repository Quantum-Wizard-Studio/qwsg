package assessment

func Registry() []Requirement {
	registry := []Requirement{
		{ID: "configuration.present", PurposeToken: "configuration_present", Class: ConfigurationRequirement, Disposition: Required, Capability: "configuration", ProbeID: "configuration", Platforms: []string{"ubuntu-24.04-amd64"}, PrivacyClass: "configuration_metadata", Remediations: []Remediation{{PlatformID: "ubuntu-24.04-amd64", Commands: [][]string{{"qwsg", "setup"}}, DisplayCommands: []string{"qwsg setup"}, Revalidate: true}}},
		{ID: "filesystem.local_semantics", PurposeToken: "filesystem_semantics", Class: RecommendedCapability, Disposition: Recommended, Capability: "environment", ProbeID: "filesystem", Platforms: []string{"ubuntu-24.04-amd64"}, PrivacyClass: "operational", Remediations: []Remediation{}},
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
		{ID: "systemd.user_manager", PurposeToken: "systemd_user_manager", Class: EnvironmentCapability, Disposition: Required, Capability: "environment", ProbeID: "systemd_user", Platforms: []string{"ubuntu-24.04-amd64"}, PrivacyClass: "operational", Remediations: []Remediation{}},
		{ID: "systemd.version", PurposeToken: "systemd_minimum_version", Class: RuntimeDependency, Disposition: Required, Capability: "environment", ProbeID: "systemd_version", Platforms: []string{"ubuntu-24.04-amd64"}, MinimumVersion: "255", PrivacyClass: "operational", Remediations: []Remediation{}},
	}
	if err := ValidateRegistry(registry); err != nil {
		panic(err)
	}
	return registry
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
