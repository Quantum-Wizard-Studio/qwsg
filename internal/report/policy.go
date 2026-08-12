package report

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"sort"
	"strings"

	"quantumwizard.hu/qwsg/internal/policy"
)

const (
	PolicyReportSchemaName               = "qwsg.policy-report"
	PolicyReportSchemaVersion            = "1.0"
	PolicyEvaluationSource    SourceType = "policy_evaluation"
)

type PolicyItem struct {
	ID                  string                  `json:"id"`
	PolicyEvaluationID  string                  `json:"policy_evaluation_id"`
	RuleEvaluationID    string                  `json:"rule_evaluation_id"`
	RuleID              string                  `json:"rule_id"`
	Outcome             policy.Outcome          `json:"outcome"`
	EvaluationStatus    policy.EvaluationStatus `json:"evaluation_status"`
	AppliedProfileIDs   []string                `json:"applied_profile_ids"`
	AppliedStatementIDs []string                `json:"applied_statement_ids"`
	ExplanationToken    string                  `json:"explanation_token"`
	EvidenceReferences  []string                `json:"evidence_references"`
	Source              SourceReference         `json:"source"`
}

type PolicySection struct {
	ID         string         `json:"id"`
	Outcome    policy.Outcome `json:"outcome"`
	TitleToken string         `json:"title_token"`
	Items      []PolicyItem   `json:"items"`
}

type PolicySummary struct {
	Total         int `json:"total"`
	Accepted      int `json:"accepted"`
	Observe       int `json:"observe"`
	Suppressed    int `json:"suppressed"`
	Escalated     int `json:"escalated"`
	Indeterminate int `json:"indeterminate"`
	NotApplicable int `json:"not_applicable"`
	Conflict      int `json:"conflict"`
}

type PolicyVersionInfo struct {
	ReportSchema     string `json:"report_schema"`
	ReportEngine     string `json:"report_engine"`
	PolicyEvaluation string `json:"policy_evaluation"`
	PolicyEngine     string `json:"policy_engine"`
	PolicyTaxonomy   string `json:"policy_taxonomy"`
	PolicyProfile    string `json:"policy_profile"`
}

// PolicyReport is the additive Policy-backed Canonical Report contract. The
// legacy Rule-backed Report 1.0 remains byte-compatible for direct consumers.
type PolicyReport struct {
	ID            string            `json:"id"`
	SchemaName    string            `json:"schema_name"`
	SchemaVersion string            `json:"schema_version"`
	Type          Type              `json:"type"`
	TitleToken    string            `json:"title_token"`
	Completeness  Completeness      `json:"completeness"`
	Summary       PolicySummary     `json:"summary"`
	Sections      []PolicySection   `json:"sections"`
	Sources       []SourceReference `json:"sources"`
	Metadata      map[string]string `json:"metadata"`
	Versions      PolicyVersionInfo `json:"versions"`
}

var policyOutcomeOrder = []policy.Outcome{policy.Escalated, policy.Conflict, policy.Observe, policy.Indeterminate, policy.Suppressed, policy.Accepted, policy.NotApplicable}

// GeneratePolicy transforms validated Policy Evaluation Records without
// re-evaluating policy or upstream engineering evidence.
func GeneratePolicy(input policy.Result) (PolicyReport, error) {
	if err := policy.Validate(input); err != nil {
		return PolicyReport{}, fmt.Errorf("invalid policy evaluation source: %w", err)
	}
	if len(input.Records) > MaxItems {
		return PolicyReport{}, fmt.Errorf("policy report item limit exceeded")
	}
	report := PolicyReport{SchemaName: PolicyReportSchemaName, SchemaVersion: PolicyReportSchemaVersion, Type: EngineeringSummary,
		TitleToken: "canonical_policy_engineering_summary", Completeness: Complete, Sections: []PolicySection{}, Sources: []SourceReference{},
		Metadata: map[string]string{"pipeline": "canonical-policy-evaluation-to-canonical-report", "rendering_model": "presentation-neutral"},
		Versions: PolicyVersionInfo{PolicyReportSchemaVersion, EngineVersion, policy.SchemaVersion, policy.EngineVersion, policy.TaxonomyVersion, policy.ProfileVersion}}
	if len(input.Records) == 0 {
		report.Completeness = Incomplete
	}
	sections := map[policy.Outcome][]PolicyItem{}
	for _, record := range input.Records {
		source := SourceReference{Type: PolicyEvaluationSource, ID: record.ID, ContractName: policy.SchemaName, ContractVersion: policy.SchemaVersion}
		item := PolicyItem{PolicyEvaluationID: record.ID, RuleEvaluationID: record.RuleEvaluationID, RuleID: record.RuleID,
			Outcome: record.Outcome, EvaluationStatus: record.EvaluationStatus,
			AppliedProfileIDs: append([]string{}, record.AppliedProfileIDs...), AppliedStatementIDs: append([]string{}, record.AppliedStatementIDs...),
			ExplanationToken: record.Explanation, EvidenceReferences: append([]string{}, record.EvidenceReferences...), Source: source}
		item.ID = policyItemID(item)
		sections[record.Outcome] = append(sections[record.Outcome], item)
		report.Sources = append(report.Sources, source)
		incrementPolicy(&report.Summary, record.Outcome)
	}
	report.Summary.Total = len(input.Records)
	for _, outcome := range policyOutcomeOrder {
		items := sections[outcome]
		if len(items) == 0 {
			continue
		}
		sort.Slice(items, func(i, j int) bool {
			if items[i].RuleID != items[j].RuleID {
				return items[i].RuleID < items[j].RuleID
			}
			return items[i].PolicyEvaluationID < items[j].PolicyEvaluationID
		})
		section := PolicySection{Outcome: outcome, TitleToken: "policy_outcome_" + string(outcome), Items: items}
		section.ID = policySectionID(section)
		report.Sections = append(report.Sections, section)
	}
	sort.Slice(report.Sources, func(i, j int) bool { return report.Sources[i].ID < report.Sources[j].ID })
	report.ID = policyReportID(report)
	if err := ValidatePolicyReport(report); err != nil {
		return PolicyReport{}, err
	}
	return report, nil
}

func ValidatePolicyReport(report PolicyReport) error {
	if report.SchemaName != PolicyReportSchemaName || report.SchemaVersion != PolicyReportSchemaVersion || report.Type != EngineeringSummary ||
		report.TitleToken != "canonical_policy_engineering_summary" || (report.Completeness != Complete && report.Completeness != Incomplete) ||
		report.Sections == nil || report.Sources == nil || report.Metadata == nil || len(report.Metadata) != 2 ||
		report.Metadata["pipeline"] != "canonical-policy-evaluation-to-canonical-report" || report.Metadata["rendering_model"] != "presentation-neutral" ||
		report.Versions != (PolicyVersionInfo{PolicyReportSchemaVersion, EngineVersion, policy.SchemaVersion, policy.EngineVersion, policy.TaxonomyVersion, policy.ProfileVersion}) {
		return fmt.Errorf("invalid policy report envelope")
	}
	if report.Summary.Total != len(report.Sources) || report.Summary.Total != policySummaryTotal(report.Summary) || report.Summary.Total > MaxItems ||
		(report.Summary.Total == 0) != (report.Completeness == Incomplete) {
		return fmt.Errorf("invalid policy report summary")
	}
	sources := map[string]SourceReference{}
	lastSource := ""
	for _, source := range report.Sources {
		if source.Type != PolicyEvaluationSource || source.ID <= lastSource || source.ContractName != policy.SchemaName || source.ContractVersion != policy.SchemaVersion {
			return fmt.Errorf("invalid policy report source")
		}
		lastSource, sources[source.ID] = source.ID, source
	}
	count, lastRank := 0, -1
	seen := map[string]bool{}
	for _, section := range report.Sections {
		rank := policyOutcomeRank(section.Outcome)
		if rank < 0 || rank <= lastRank || len(section.Items) == 0 || section.TitleToken != "policy_outcome_"+string(section.Outcome) || section.ID != policySectionID(section) {
			return fmt.Errorf("invalid policy report section")
		}
		lastRank = rank
		lastKey := ""
		for _, item := range section.Items {
			key := item.RuleID + "\x00" + item.PolicyEvaluationID
			if item.Outcome != section.Outcome || key <= lastKey || seen[item.ID] || item.ID != policyItemID(item) || item.Source != sources[item.PolicyEvaluationID] || !validPolicyItem(item) {
				return fmt.Errorf("invalid policy report item")
			}
			lastKey, seen[item.ID] = key, true
			count++
		}
	}
	if count != report.Summary.Total || len(seen) != len(sources) || report.ID != policyReportID(report) {
		return fmt.Errorf("invalid policy report identity or traceability")
	}
	return nil
}

func MarshalPolicyCanonical(report PolicyReport) ([]byte, error) {
	if err := ValidatePolicyReport(report); err != nil {
		return nil, err
	}
	return json.Marshal(report)
}

func RenderPolicyText(report PolicyReport) (string, error) {
	if err := ValidatePolicyReport(report); err != nil {
		return "", err
	}
	var builder strings.Builder
	fmt.Fprintf(&builder, "%s %s %s\n", safe(report.TitleToken), report.ID, report.Completeness)
	for _, section := range report.Sections {
		fmt.Fprintf(&builder, "[%s]\n", safe(section.TitleToken))
		for _, item := range section.Items {
			fmt.Fprintf(&builder, "- %s %s %s %s\n", safe(item.RuleID), item.Outcome, item.EvaluationStatus, safe(item.ExplanationToken))
		}
	}
	return builder.String(), nil
}

func validPolicyItem(item PolicyItem) bool {
	if item.PolicyEvaluationID == "" || item.RuleEvaluationID == "" || item.RuleID == "" || item.ExplanationToken == "" || item.EvidenceReferences == nil || len(item.EvidenceReferences) != 1 || item.EvidenceReferences[0] != item.RuleEvaluationID || item.AppliedProfileIDs == nil || item.AppliedStatementIDs == nil {
		return false
	}
	if item.Outcome == policy.NotApplicable {
		return item.EvaluationStatus == policy.EvaluationSkipped && len(item.AppliedProfileIDs) == 0
	}
	return item.EvaluationStatus == policy.EvaluationComplete && len(item.AppliedProfileIDs) > 0
}

func incrementPolicy(summary *PolicySummary, outcome policy.Outcome) {
	switch outcome {
	case policy.Accepted:
		summary.Accepted++
	case policy.Observe:
		summary.Observe++
	case policy.Suppressed:
		summary.Suppressed++
	case policy.Escalated:
		summary.Escalated++
	case policy.Indeterminate:
		summary.Indeterminate++
	case policy.NotApplicable:
		summary.NotApplicable++
	case policy.Conflict:
		summary.Conflict++
	}
}
func policySummaryTotal(value PolicySummary) int {
	return value.Accepted + value.Observe + value.Suppressed + value.Escalated + value.Indeterminate + value.NotApplicable + value.Conflict
}
func policyOutcomeRank(value policy.Outcome) int {
	for index, candidate := range policyOutcomeOrder {
		if value == candidate {
			return index
		}
	}
	return -1
}
func policyItemID(value PolicyItem) string {
	copy := value
	copy.ID = ""
	document, _ := json.Marshal(copy)
	return reportStableID("policy-item", document)
}
func policySectionID(value PolicySection) string {
	copy := value
	copy.ID = ""
	document, _ := json.Marshal(copy)
	return reportStableID("policy-section", document)
}
func policyReportID(value PolicyReport) string {
	copy := value
	copy.ID = ""
	document, _ := json.Marshal(copy)
	return reportStableID("policy-report", document)
}
func reportStableID(domain string, document []byte) string {
	sum := sha256.Sum256(append([]byte(PolicyReportSchemaName+"/"+PolicyReportSchemaVersion+"/"+domain+"\x00"), document...))
	return hex.EncodeToString(sum[:])
}
