// Package assessment owns the presentation-independent Smart Install and
// operational readiness contracts. It detects and plans; it never mutates a
// host or executes a remediation recommendation.
package assessment

import (
	"fmt"
	"sort"
	"strings"
	"time"
)

const (
	SchemaName      = "qwsg.readiness-assessment"
	SchemaVersion   = "1.1"
	ModelVersion    = "1.1"
	RegistryVersion = "1.1"
)

type Classification string

const (
	Satisfied           Classification = "satisfied"
	MissingRequired     Classification = "missing_required"
	MissingOptional     Classification = "missing_optional"
	UnknownVerification Classification = "unknown_requires_verification"
	Incompatible        Classification = "incompatible"
)

type RequirementClass string

const (
	RuntimeDependency        RequirementClass = "runtime_dependency"
	InstallTimeDependency    RequirementClass = "install_time_dependency"
	OptionalFeature          RequirementClass = "optional_feature_dependency"
	EnvironmentCapability    RequirementClass = "environment_capability"
	ConfigurationRequirement RequirementClass = "configuration_requirement"
	RecommendedCapability    RequirementClass = "recommended_non_blocking"
)

type Disposition string

const (
	Required    Disposition = "required"
	Optional    Disposition = "optional"
	Recommended Disposition = "recommended"
)

type SummaryState string

const (
	Ready    SummaryState = "ready"
	Partial  SummaryState = "partial"
	NotReady SummaryState = "not_ready"
	Unknown  SummaryState = "unknown"
)

type Remediation struct {
	PlatformID         string     `json:"platform_id"`
	Commands           [][]string `json:"commands"`
	DisplayCommands    []string   `json:"display_commands"`
	ElevatedPrivileges bool       `json:"elevated_privileges"`
	Revalidate         bool       `json:"revalidate"`
	CompatibilityGuard string     `json:"compatibility_guard,omitempty"`
}

type PrivilegeRequirement string

const (
	PrivilegeNone               PrivilegeRequirement = "none"
	PrivilegeAdministrator      PrivilegeRequirement = "administrator"
	PrivilegeManualVerification PrivilegeRequirement = "manual_verification"
)

// Guidance is registry-owned, presentation-independent operator guidance.
// Tokens are localization-ready identifiers; no field is executable.
type Guidance struct {
	ExplanationToken     string               `json:"explanation_token"`
	BlockingEffect       string               `json:"blocking_effect"`
	VerificationActions  []string             `json:"verification_actions"`
	OperatorActions      []string             `json:"operator_actions"`
	PrivilegeRequirement PrivilegeRequirement `json:"privilege_requirement"`
	ManualVerification   bool                 `json:"manual_verification"`
	SafetyNotes          []string             `json:"safety_notes"`
	RevalidationAction   string               `json:"revalidation_action"`
}

type GuidanceRule struct {
	PlatformID     string         `json:"platform_id"`
	Classification Classification `json:"classification"`
	EvidenceToken  string         `json:"evidence_token"`
	Guidance       Guidance       `json:"guidance"`
}

type Requirement struct {
	ID             string           `json:"id"`
	PurposeToken   string           `json:"purpose_token"`
	Class          RequirementClass `json:"requirement_class"`
	Disposition    Disposition      `json:"disposition"`
	Capability     string           `json:"capability"`
	ProbeID        string           `json:"probe_id"`
	Platforms      []string         `json:"supported_platforms"`
	MinimumVersion string           `json:"minimum_version,omitempty"`
	PrivacyClass   string           `json:"privacy_class"`
	Remediations   []Remediation    `json:"remediations"`
	GuidanceRules  []GuidanceRule   `json:"guidance_rules"`
}

type Finding struct {
	RequirementID  string         `json:"requirement_id"`
	Classification Classification `json:"classification"`
	EvidenceToken  string         `json:"evidence_token"`
	ObservedValue  string         `json:"observed_value,omitempty"`
	Remediation    *Remediation   `json:"remediation,omitempty"`
	Guidance       *Guidance      `json:"guidance,omitempty"`
}

type Platform struct {
	ID           string `json:"id"`
	Distribution string `json:"distribution"`
	Version      string `json:"version"`
	Architecture string `json:"architecture"`
	Supported    bool   `json:"supported"`
}

type DomainSummary struct {
	Domain string       `json:"domain"`
	State  SummaryState `json:"state"`
}

type Report struct {
	SchemaName      string          `json:"schema_name"`
	SchemaVersion   string          `json:"schema_version"`
	ModelVersion    string          `json:"model_version"`
	RegistryVersion string          `json:"registry_version"`
	AssessedAt      time.Time       `json:"assessed_at"`
	Phase           string          `json:"phase"`
	Platform        Platform        `json:"platform"`
	Findings        []Finding       `json:"findings"`
	Domains         []DomainSummary `json:"domains"`
	NextActions     []string        `json:"next_actions"`
}

func ValidateRegistry(registry []Requirement) error {
	seen := map[string]bool{}
	last := ""
	for _, requirement := range registry {
		if !token(requirement.ID) || requirement.ID <= last || seen[requirement.ID] ||
			!token(requirement.PurposeToken) || !validClass(requirement.Class) ||
			!validDisposition(requirement.Disposition) || !token(requirement.Capability) ||
			!token(requirement.ProbeID) || len(requirement.Platforms) == 0 ||
			!token(requirement.PrivacyClass) {
			return fmt.Errorf("invalid assessment registry")
		}
		seen[requirement.ID] = true
		last = requirement.ID
		remediationPlatforms := map[string]bool{}
		for _, remediation := range requirement.Remediations {
			if !token(remediation.PlatformID) || len(remediation.Commands) == 0 ||
				len(remediation.Commands) > 8 || len(remediation.Commands) != len(remediation.DisplayCommands) ||
				!remediation.Revalidate || remediationPlatforms[remediation.PlatformID] {
				return fmt.Errorf("invalid assessment remediation")
			}
			remediationPlatforms[remediation.PlatformID] = true
			for index, command := range remediation.Commands {
				if len(command) == 0 || len(command) > 32 || command[0] == "" ||
					unsafeCommandName(command[0]) || !safeBoundedText(remediation.DisplayCommands[index], 512) {
					return fmt.Errorf("unsafe assessment remediation")
				}
				for _, argument := range command {
					if argument == "" || !safeBoundedText(argument, 512) {
						return fmt.Errorf("unsafe assessment remediation")
					}
				}
			}
		}
		guidanceSelectors := map[string]bool{}
		for _, rule := range requirement.GuidanceRules {
			selector := rule.PlatformID + "\x00" + string(rule.Classification) + "\x00" + rule.EvidenceToken
			if !token(rule.PlatformID) || !validClassification(rule.Classification) ||
				!token(rule.EvidenceToken) || validateGuidance(rule.Guidance) != nil || guidanceSelectors[selector] {
				return fmt.Errorf("invalid assessment guidance")
			}
			guidanceSelectors[selector] = true
		}
	}
	return nil
}

func unsafeCommandName(value string) bool {
	base := value
	if index := strings.LastIndexByte(value, '/'); index >= 0 {
		base = value[index+1:]
	}
	return base == "sh" || base == "bash" || base == "dash" || base == "zsh" || base == "ksh"
}

func safeBoundedText(value string, limit int) bool {
	if value == "" || len(value) > limit {
		return false
	}
	for _, r := range value {
		if r < 0x20 || r == 0x7f {
			return false
		}
	}
	return true
}

func ValidateReport(report Report) error {
	if report.SchemaName != SchemaName || report.SchemaVersion != SchemaVersion ||
		report.ModelVersion != ModelVersion || report.RegistryVersion != RegistryVersion ||
		report.AssessedAt.IsZero() || (report.Phase != "install" && report.Phase != "operational") {
		return fmt.Errorf("invalid assessment report")
	}
	last := ""
	for _, finding := range report.Findings {
		if !token(finding.RequirementID) || finding.RequirementID <= last ||
			!validClassification(finding.Classification) || !token(finding.EvidenceToken) {
			return fmt.Errorf("invalid assessment finding")
		}
		last = finding.RequirementID
		if finding.Guidance != nil && validateGuidance(*finding.Guidance) != nil {
			return fmt.Errorf("invalid assessment guidance")
		}
	}
	return nil
}

func validateGuidance(value Guidance) error {
	if !token(value.ExplanationToken) || !token(value.BlockingEffect) ||
		!validPrivilege(value.PrivilegeRequirement) || !token(value.RevalidationAction) ||
		len(value.VerificationActions) == 0 || len(value.VerificationActions) > 8 ||
		len(value.OperatorActions) == 0 || len(value.OperatorActions) > 8 || len(value.SafetyNotes) > 8 {
		return fmt.Errorf("invalid guidance")
	}
	for _, values := range [][]string{value.VerificationActions, value.OperatorActions, value.SafetyNotes} {
		for _, item := range values {
			if !token(item) {
				return fmt.Errorf("invalid guidance token")
			}
		}
	}
	return nil
}

func validPrivilege(value PrivilegeRequirement) bool {
	return value == PrivilegeNone || value == PrivilegeAdministrator || value == PrivilegeManualVerification
}

func Summarize(findings []Finding) SummaryState {
	state := Ready
	for _, finding := range findings {
		switch finding.Classification {
		case MissingRequired, Incompatible:
			return NotReady
		case UnknownVerification:
			requirement, known := RequirementByID(finding.RequirementID)
			if !known || requirement.Disposition == Required {
				state = Unknown
			} else if state == Ready {
				state = Partial
			}
		case MissingOptional:
			if state == Ready {
				state = Partial
			}
		}
	}
	return state
}

func SortFindings(findings []Finding) {
	sort.Slice(findings, func(i, j int) bool { return findings[i].RequirementID < findings[j].RequirementID })
}

func token(value string) bool {
	if value == "" || len(value) > 128 {
		return false
	}
	for _, r := range value {
		if (r < 'a' || r > 'z') && (r < '0' || r > '9') && r != '_' && r != '.' && r != '-' {
			return false
		}
	}
	return true
}

func validClassification(value Classification) bool {
	return value == Satisfied || value == MissingRequired || value == MissingOptional || value == UnknownVerification || value == Incompatible
}

func validClass(value RequirementClass) bool {
	return value == RuntimeDependency || value == InstallTimeDependency || value == OptionalFeature || value == EnvironmentCapability || value == ConfigurationRequirement || value == RecommendedCapability
}

func validDisposition(value Disposition) bool {
	return value == Required || value == Optional || value == Recommended
}
