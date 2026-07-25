// Package report deterministically transforms canonical Rule Evaluation
// Records into presentation-neutral Canonical Reports. It performs no
// collection, comparison, drift classification, health or rule evaluation,
// policy, monitoring, alerting, delivery, remediation, networking, or AI work.
package report

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"sort"
	"strconv"
	"strings"
	"unicode"

	"quantumwizard.hu/qwsg/internal/health"
	"quantumwizard.hu/qwsg/internal/rule"
)

const (
	SchemaName      = "qwsg.report"
	SchemaVersion   = "1.0"
	EngineVersion   = "1.0"
	TaxonomyVersion = "1.0"
	MaxItems        = 4096
)

type Type string

const EngineeringSummary Type = "engineering_summary"

type SourceType string

const RuleEvaluationSource SourceType = "rule_evaluation"

type Completeness string

const (
	Complete   Completeness = "complete"
	Incomplete Completeness = "incomplete"
)

type SourceReference struct {
	Type            SourceType `json:"type"`
	ID              string     `json:"id"`
	ContractName    string     `json:"contract_name"`
	ContractVersion string     `json:"contract_version"`
}

type Item struct {
	ID                    string                `json:"id"`
	RuleID                string                `json:"rule_id"`
	HealthRecordID        string                `json:"health_record_id,omitempty"`
	Outcome               rule.Outcome          `json:"outcome"`
	EvaluationStatus      rule.EvaluationStatus `json:"evaluation_status"`
	MatchResult           rule.MatchResult      `json:"match_result"`
	ConfidenceBasisPoints int                   `json:"confidence_basis_points"`
	ExplanationToken      string                `json:"explanation_token"`
	EvidenceReferences    []string              `json:"evidence_references"`
	Source                SourceReference       `json:"source"`
}

type Section struct {
	ID         string       `json:"id"`
	Outcome    rule.Outcome `json:"outcome"`
	TitleToken string       `json:"title_token"`
	Items      []Item       `json:"items"`
}

type Summary struct {
	Total                int `json:"total"`
	Matched              int `json:"matched"`
	NotMatched           int `json:"not_matched"`
	InsufficientEvidence int `json:"insufficient_evidence"`
	UnsupportedRule      int `json:"unsupported_rule"`
	InvalidRule          int `json:"invalid_rule"`
	EvaluationError      int `json:"evaluation_error"`
	DisabledRule         int `json:"disabled_rule"`
}

type VersionInfo struct {
	ReportSchema   string `json:"report_schema"`
	ReportEngine   string `json:"report_engine"`
	ReportTaxonomy string `json:"report_taxonomy"`
	RuleEvaluation string `json:"rule_evaluation"`
	RuleEngine     string `json:"rule_engine"`
	RuleTaxonomy   string `json:"rule_taxonomy"`
	RuleContract   string `json:"rule_contract"`
	HealthSchema   string `json:"health_schema"`
	HealthTaxonomy string `json:"health_taxonomy"`
}

// Report is the Canonical Report 1.0 public contract.
type Report struct {
	ID            string            `json:"id"`
	SchemaName    string            `json:"schema_name"`
	SchemaVersion string            `json:"schema_version"`
	Type          Type              `json:"type"`
	TitleToken    string            `json:"title_token"`
	Completeness  Completeness      `json:"completeness"`
	Summary       Summary           `json:"summary"`
	Sections      []Section         `json:"sections"`
	Sources       []SourceReference `json:"sources"`
	Metadata      map[string]string `json:"metadata"`
	Versions      VersionInfo       `json:"versions"`
}

var outcomeOrder = []rule.Outcome{
	rule.Matched,
	rule.NotMatched,
	rule.InsufficientEvidence,
	rule.UnsupportedRule,
	rule.InvalidRule,
	rule.EvaluationError,
	rule.DisabledRule,
}

// Generate transforms a validated Rule Result without re-evaluating it.
func Generate(input rule.Result) (Report, error) {
	if err := rule.Validate(input); err != nil {
		return Report{}, fmt.Errorf("invalid rule evaluation source: %w", err)
	}
	if len(input.Records) > MaxItems {
		return Report{}, fmt.Errorf("report item limit exceeded")
	}
	report := Report{
		SchemaName: SchemaName, SchemaVersion: SchemaVersion,
		Type: EngineeringSummary, TitleToken: "canonical_engineering_summary",
		Completeness: Complete,
		Sections:     []Section{}, Sources: []SourceReference{},
		Metadata: map[string]string{
			"pipeline":        "canonical-rule-evaluation-to-canonical-report",
			"rendering_model": "presentation-neutral",
		},
		Versions: expectedVersions(),
	}
	if len(input.Records) == 0 {
		report.Completeness = Incomplete
	}
	sections := make(map[rule.Outcome][]Item, len(outcomeOrder))
	for _, record := range input.Records {
		source := SourceReference{
			Type: RuleEvaluationSource, ID: record.ID,
			ContractName: rule.SchemaName, ContractVersion: rule.SchemaVersion,
		}
		item := Item{
			RuleID: record.RuleID, HealthRecordID: record.HealthRecordID,
			Outcome: record.Outcome, EvaluationStatus: record.EvaluationStatus,
			MatchResult:           record.MatchResult,
			ConfidenceBasisPoints: record.ConfidenceBasisPoints,
			ExplanationToken:      record.Explanation,
			EvidenceReferences:    append([]string{}, record.EvidenceReferences...),
			Source:                source,
		}
		item.ID = stableItemID(item)
		sections[record.Outcome] = append(sections[record.Outcome], item)
		report.Sources = append(report.Sources, source)
		increment(&report.Summary, record.Outcome)
	}
	report.Summary.Total = len(input.Records)
	for _, outcome := range outcomeOrder {
		items := sections[outcome]
		if len(items) == 0 {
			continue
		}
		sort.Slice(items, func(i, j int) bool {
			if items[i].RuleID != items[j].RuleID {
				return items[i].RuleID < items[j].RuleID
			}
			if items[i].HealthRecordID != items[j].HealthRecordID {
				return items[i].HealthRecordID < items[j].HealthRecordID
			}
			return items[i].ID < items[j].ID
		})
		section := Section{
			Outcome: outcome, TitleToken: "rule_outcome_" + string(outcome),
			Items: items,
		}
		section.ID = stableSectionID(section)
		report.Sections = append(report.Sections, section)
	}
	sort.Slice(report.Sources, func(i, j int) bool {
		return report.Sources[i].ID < report.Sources[j].ID
	})
	report.ID = stableReportID(report)
	if err := Validate(report); err != nil {
		return Report{}, err
	}
	return report, nil
}

// Validate rejects unsupported or inconsistent Canonical Report contracts.
func Validate(report Report) error {
	if report.SchemaName != SchemaName || report.SchemaVersion != SchemaVersion ||
		report.Type != EngineeringSummary ||
		report.TitleToken != "canonical_engineering_summary" ||
		(report.Completeness != Complete && report.Completeness != Incomplete) ||
		report.Sections == nil || report.Sources == nil ||
		report.Metadata == nil || len(report.Metadata) != 2 ||
		report.Metadata["pipeline"] != "canonical-rule-evaluation-to-canonical-report" ||
		report.Metadata["rendering_model"] != "presentation-neutral" ||
		report.Versions != expectedVersions() {
		return fmt.Errorf("invalid report envelope")
	}
	if report.Summary.Total > MaxItems || report.Summary.Total != len(report.Sources) ||
		report.Summary.Total != summaryTotal(report.Summary) ||
		(report.Summary.Total == 0) != (report.Completeness == Incomplete) {
		return fmt.Errorf("invalid report completeness or summary")
	}
	sourceByID := make(map[string]SourceReference, len(report.Sources))
	lastSource := ""
	for _, source := range report.Sources {
		if !validSource(source) || source.ID <= lastSource {
			return fmt.Errorf("invalid or unordered report source")
		}
		lastSource, sourceByID[source.ID] = source.ID, source
	}
	itemCount := 0
	lastRank := -1
	seenItems := map[string]bool{}
	for _, section := range report.Sections {
		rank := outcomeRank(section.Outcome)
		if rank < 0 || rank <= lastRank || section.TitleToken != "rule_outcome_"+string(section.Outcome) ||
			len(section.Items) == 0 || section.ID != stableSectionID(section) {
			return fmt.Errorf("invalid report section")
		}
		lastRank = rank
		lastKey := ""
		for _, item := range section.Items {
			key := item.RuleID + "\x00" + item.HealthRecordID + "\x00" + item.ID
			if item.Outcome != section.Outcome || key <= lastKey || seenItems[item.ID] ||
				item.ID != stableItemID(item) || !validItem(item) ||
				sourceByID[item.Source.ID] != item.Source {
				return fmt.Errorf("invalid report item")
			}
			lastKey, seenItems[item.ID] = key, true
			itemCount++
		}
	}
	if itemCount != report.Summary.Total || len(seenItems) != len(sourceByID) ||
		report.ID != stableReportID(report) {
		return fmt.Errorf("invalid report identity or traceability")
	}
	return nil
}

// MarshalCanonical validates and emits byte-stable canonical JSON.
func MarshalCanonical(report Report) ([]byte, error) {
	if err := Validate(report); err != nil {
		return nil, err
	}
	return json.Marshal(report)
}

// RenderText is a deterministic, control-character-safe view derived solely
// from a validated Canonical Report. It is not an export or delivery engine.
func RenderText(report Report) (string, error) {
	if err := Validate(report); err != nil {
		return "", err
	}
	var builder strings.Builder
	fmt.Fprintf(&builder, "%s %s %s\n", safe(report.TitleToken), report.ID, report.Completeness)
	for _, section := range report.Sections {
		fmt.Fprintf(&builder, "[%s]\n", safe(section.TitleToken))
		for _, item := range section.Items {
			fmt.Fprintf(&builder, "- %s %s %s %s\n",
				safe(item.RuleID), item.Outcome, item.EvaluationStatus,
				safe(item.ExplanationToken))
		}
	}
	return builder.String(), nil
}

func validItem(item Item) bool {
	if item.RuleID == "" || item.ExplanationToken == "" ||
		item.EvidenceReferences == nil || !validSource(item.Source) ||
		item.Source.ContractVersion != rule.SchemaVersion ||
		item.ConfidenceBasisPoints < 0 || item.ConfidenceBasisPoints > 10000 {
		return false
	}
	switch item.Outcome {
	case rule.Matched:
		return completedItem(item, rule.Match, "condition_matched")
	case rule.NotMatched:
		return completedItem(item, rule.NoMatch, "condition_not_matched")
	case rule.InsufficientEvidence:
		return terminalItem(item, rule.EvaluationSkipped,
			item.ExplanationToken == "no_health_record_in_scope")
	case rule.UnsupportedRule:
		return terminalItem(item, rule.EvaluationSkipped,
			item.ExplanationToken == "unsupported_rule_contract" ||
				item.ExplanationToken == "unsupported_rule_category" ||
				item.ExplanationToken == "unsupported_operator")
	case rule.InvalidRule:
		return terminalItem(item, rule.EvaluationFailed,
			strings.HasPrefix(item.ExplanationToken, "invalid_") ||
				item.ExplanationToken == "condition_bounds_exceeded")
	case rule.EvaluationError:
		return terminalItem(item, rule.EvaluationFailed,
			item.ExplanationToken == "invalid_health_contract")
	case rule.DisabledRule:
		return terminalItem(item, rule.EvaluationSkipped,
			item.ExplanationToken == "rule_disabled")
	default:
		return false
	}
}

func completedItem(item Item, match rule.MatchResult, explanation string) bool {
	return item.HealthRecordID != "" && len(item.EvidenceReferences) == 1 &&
		item.EvidenceReferences[0] == item.HealthRecordID &&
		item.EvaluationStatus == rule.EvaluationComplete &&
		item.MatchResult == match && item.ConfidenceBasisPoints == 10000 &&
		item.ExplanationToken == explanation
}

func terminalItem(item Item, status rule.EvaluationStatus, explanation bool) bool {
	return item.HealthRecordID == "" && len(item.EvidenceReferences) == 0 &&
		item.EvaluationStatus == status && item.MatchResult == rule.Indeterminate &&
		item.ConfidenceBasisPoints == 0 && explanation
}

func validSource(source SourceReference) bool {
	return source.Type == RuleEvaluationSource && source.ID != "" &&
		source.ContractName == rule.SchemaName &&
		source.ContractVersion == rule.SchemaVersion
}

func increment(summary *Summary, outcome rule.Outcome) {
	switch outcome {
	case rule.Matched:
		summary.Matched++
	case rule.NotMatched:
		summary.NotMatched++
	case rule.InsufficientEvidence:
		summary.InsufficientEvidence++
	case rule.UnsupportedRule:
		summary.UnsupportedRule++
	case rule.InvalidRule:
		summary.InvalidRule++
	case rule.EvaluationError:
		summary.EvaluationError++
	case rule.DisabledRule:
		summary.DisabledRule++
	}
}

func summaryTotal(summary Summary) int {
	return summary.Matched + summary.NotMatched + summary.InsufficientEvidence +
		summary.UnsupportedRule + summary.InvalidRule + summary.EvaluationError +
		summary.DisabledRule
}

func stableItemID(item Item) string {
	return hash("item", item.RuleID, item.HealthRecordID, string(item.Outcome),
		string(item.EvaluationStatus), string(item.MatchResult),
		strconv.Itoa(item.ConfidenceBasisPoints), item.ExplanationToken,
		strings.Join(item.EvidenceReferences, "\x1f"), item.Source.ID)
}

func stableSectionID(section Section) string {
	ids := make([]string, len(section.Items))
	for i := range section.Items {
		ids[i] = section.Items[i].ID
	}
	return hash("section", string(section.Outcome), section.TitleToken,
		strings.Join(ids, "\x1f"))
}

func stableReportID(report Report) string {
	sections := make([]string, len(report.Sections))
	for i := range report.Sections {
		sections[i] = report.Sections[i].ID
	}
	sources := make([]string, len(report.Sources))
	for i := range report.Sources {
		sources[i] = report.Sources[i].ID
	}
	return hash("report", string(report.Type), report.TitleToken,
		string(report.Completeness), strings.Join(sections, "\x1f"),
		strings.Join(sources, "\x1f"))
}

func hash(kind string, parts ...string) string {
	digest := sha256.New()
	digest.Write([]byte("qwsg.report/" + EngineVersion + "/" + kind))
	for _, part := range parts {
		digest.Write([]byte{0})
		digest.Write([]byte(part))
	}
	return hex.EncodeToString(digest.Sum(nil))
}

func expectedVersions() VersionInfo {
	return VersionInfo{
		ReportSchema: SchemaVersion, ReportEngine: EngineVersion,
		ReportTaxonomy: TaxonomyVersion, RuleEvaluation: rule.SchemaVersion,
		RuleEngine: rule.EngineVersion, RuleTaxonomy: rule.TaxonomyVersion,
		RuleContract: rule.RuleVersion, HealthSchema: health.SchemaVersion,
		HealthTaxonomy: health.TaxonomyVersion,
	}
}

func outcomeRank(outcome rule.Outcome) int {
	for i, candidate := range outcomeOrder {
		if outcome == candidate {
			return i
		}
	}
	return -1
}

func safe(value string) string {
	var builder strings.Builder
	for _, character := range value {
		if unicode.IsControl(character) {
			fmt.Fprintf(&builder, "\\u%04X", character)
			continue
		}
		builder.WriteRune(character)
	}
	return builder.String()
}
