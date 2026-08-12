// Package rule deterministically matches canonical rule definitions against
// canonical Health Records. It performs no health evaluation, policy,
// alerting, reporting, remediation, process execution, networking, or AI work.
package rule

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"regexp"
	"sort"
	"strconv"
	"strings"

	"quantumwizard.hu/qwsg/internal/drift"
	"quantumwizard.hu/qwsg/internal/health"
)

const (
	SchemaName      = "qwsg.rule-evaluation"
	SchemaVersion   = "1.0"
	EngineVersion   = "1.0"
	TaxonomyVersion = "1.0"
	RuleVersion     = "1.0"
	MaxRules        = 1024
	MaxDepth        = 8
	MaxNodes        = 64
)

type Category string

const (
	StatusRule    Category = "health_status"
	CategoryRule  Category = "health_category"
	EvidenceRule  Category = "health_evidence"
	CompositeRule Category = "composite"
	ExtensionRule Category = "extension"
)

type Field string

const (
	FieldHealthID   Field = "health_id"
	FieldCategory   Field = "category"
	FieldStatus     Field = "status"
	FieldEvidence   Field = "evidence_state"
	FieldConfidence Field = "confidence_basis_points"
	FieldReason     Field = "reason"
	FieldLayer      Field = "scope.layer"
	FieldObjectID   Field = "scope.object_id"
	FieldPath       Field = "scope.path"
)

type Operator string

const (
	Equal              Operator = "eq"
	NotEqual           Operator = "ne"
	GreaterThan        Operator = "gt"
	GreaterThanOrEqual Operator = "gte"
	LessThan           Operator = "lt"
	LessThanOrEqual    Operator = "lte"
	In                 Operator = "in"
	Exists             Operator = "exists"
	StatusMatches      Operator = "status_matches"
	CategoryMatches    Operator = "category_matches"
	All                Operator = "and"
	Any                Operator = "or"
	Not                Operator = "not"
)

type Outcome string

const (
	Matched              Outcome = "matched"
	NotMatched           Outcome = "not_matched"
	InsufficientEvidence Outcome = "insufficient_evidence"
	UnsupportedRule      Outcome = "unsupported_rule"
	InvalidRule          Outcome = "invalid_rule"
	EvaluationError      Outcome = "evaluation_error"
	DisabledRule         Outcome = "disabled_rule"
)

type EvaluationStatus string

const (
	EvaluationComplete EvaluationStatus = "complete"
	EvaluationSkipped  EvaluationStatus = "skipped"
	EvaluationFailed   EvaluationStatus = "failed"
)

type MatchResult string

const (
	Match         MatchResult = "match"
	NoMatch       MatchResult = "no_match"
	Indeterminate MatchResult = "indeterminate"
)

// Value is an explicitly typed scalar. Exactly one field must be set.
type Value struct {
	String  *string `json:"string,omitempty"`
	Number  *int64  `json:"number,omitempty"`
	Boolean *bool   `json:"boolean,omitempty"`
}

// Condition is a bounded tree over the fixed canonical Health field registry.
type Condition struct {
	Operator Operator    `json:"operator"`
	Field    Field       `json:"field,omitempty"`
	Value    Value       `json:"value,omitempty"`
	Values   []Value     `json:"values,omitempty"`
	Children []Condition `json:"children,omitempty"`
}

// Scope restricts which canonical Health Records a Rule evaluates.
type Scope struct {
	HealthIDs  []string         `json:"health_ids"`
	Categories []drift.Category `json:"categories"`
}

// Definition is the canonical public Rule Definition 1.0 contract.
type Definition struct {
	ID                string            `json:"id"`
	ContractVersion   string            `json:"contract_version"`
	Category          Category          `json:"category"`
	Scope             Scope             `json:"scope"`
	Enabled           bool              `json:"enabled"`
	InputRequirements []Field           `json:"input_requirements"`
	Condition         Condition         `json:"condition"`
	Description       string            `json:"description"`
	Metadata          map[string]string `json:"metadata"`
}

type VersionInfo struct {
	EvaluationSchema   string `json:"evaluation_schema"`
	EvaluationEngine   string `json:"evaluation_engine"`
	EvaluationTaxonomy string `json:"evaluation_taxonomy"`
	RuleContract       string `json:"rule_contract"`
	HealthSchema       string `json:"health_schema"`
	HealthTaxonomy     string `json:"health_taxonomy"`
}

// EvaluationRecord is Canonical Rule Evaluation Record 1.0.
type EvaluationRecord struct {
	ID                    string            `json:"id"`
	RuleID                string            `json:"rule_id"`
	HealthRecordID        string            `json:"health_record_id,omitempty"`
	Outcome               Outcome           `json:"outcome"`
	EvaluationStatus      EvaluationStatus  `json:"evaluation_status"`
	MatchResult           MatchResult       `json:"match_result"`
	ConfidenceBasisPoints int               `json:"confidence_basis_points"`
	Explanation           string            `json:"explanation"`
	EvidenceReferences    []string          `json:"evidence_references"`
	Metadata              map[string]string `json:"metadata"`
	Versions              VersionInfo       `json:"versions"`
}

type Result struct {
	SchemaName      string             `json:"schema_name"`
	SchemaVersion   string             `json:"schema_version"`
	EngineVersion   string             `json:"engine_version"`
	TaxonomyVersion string             `json:"taxonomy_version"`
	Records         []EvaluationRecord `json:"records"`
	Metadata        map[string]string  `json:"metadata"`
}

// ValidateDefinitions validates configuration-time Rule Definition input
// without evaluating Health evidence. Rule semantics remain owned here while
// configuration consumers can reject malformed definitions before execution.
func ValidateDefinitions(definitions []Definition) error {
	if definitions == nil || len(definitions) > MaxRules {
		return fmt.Errorf("invalid rule definition collection")
	}
	last := ""
	for _, definition := range definitions {
		if definition.ID <= last {
			return fmt.Errorf("rule definitions must be ordered with unique ids")
		}
		if outcome, reason := classifyDefinition(definition); outcome != "" {
			return fmt.Errorf("rule definition %q is %s: %s", definition.ID, outcome, reason)
		}
		last = definition.ID
	}
	return nil
}

var ruleIDPattern = regexp.MustCompile(`^[a-z0-9]+([._:-][a-z0-9]+)*$`)

// Evaluate applies rules in stable Rule ID order and emits stable evaluation
// records. Invalid Health input becomes evaluation_error, never not_matched.
func Evaluate(definitions []Definition, input health.Result) (Result, error) {
	if len(definitions) > MaxRules {
		return Result{}, fmt.Errorf("rule limit exceeded")
	}
	definitions = append([]Definition(nil), definitions...)
	sort.Slice(definitions, func(i, j int) bool { return definitions[i].ID < definitions[j].ID })
	for i, definition := range definitions {
		if definition.ID == "" || (i > 0 && definition.ID == definitions[i-1].ID) {
			return Result{}, fmt.Errorf("rule ids must be nonempty and unique")
		}
	}
	result := newResult()
	healthErr := health.Validate(input)
	for _, definition := range definitions {
		if healthErr != nil {
			result.Records = append(result.Records, terminalRecord(
				definition, EvaluationError, EvaluationFailed,
				"invalid_health_contract",
			))
			continue
		}
		class, reason := classifyDefinition(definition)
		switch class {
		case UnsupportedRule:
			result.Records = append(result.Records, terminalRecord(
				definition, UnsupportedRule, EvaluationSkipped, reason,
			))
			continue
		case InvalidRule:
			result.Records = append(result.Records, terminalRecord(
				definition, InvalidRule, EvaluationFailed, reason,
			))
			continue
		}
		if !definition.Enabled {
			result.Records = append(result.Records, terminalRecord(
				definition, DisabledRule, EvaluationSkipped, "rule_disabled",
			))
			continue
		}
		selected := selectRecords(definition.Scope, input.Records)
		if len(selected) == 0 {
			result.Records = append(result.Records, terminalRecord(
				definition, InsufficientEvidence, EvaluationSkipped,
				"no_health_record_in_scope",
			))
			continue
		}
		for _, record := range selected {
			matched := evaluateCondition(definition.Condition, record)
			outcome, match, explanation := NotMatched, NoMatch, "condition_not_matched"
			if matched {
				outcome, match, explanation = Matched, Match, "condition_matched"
			}
			evaluation := EvaluationRecord{
				RuleID: definition.ID, HealthRecordID: record.ID,
				Outcome: outcome, EvaluationStatus: EvaluationComplete,
				MatchResult: match, ConfidenceBasisPoints: 10000,
				Explanation:        explanation,
				EvidenceReferences: []string{record.ID},
				Metadata: map[string]string{
					"rule_category": string(definition.Category),
					"root_operator": string(definition.Condition.Operator),
				},
				Versions: expectedVersions(),
			}
			evaluation.ID = stableID(evaluation)
			result.Records = append(result.Records, evaluation)
		}
	}
	sortRecords(result.Records)
	if err := Validate(result); err != nil {
		return Result{}, err
	}
	return result, nil
}

func newResult() Result {
	return Result{
		SchemaName: SchemaName, SchemaVersion: SchemaVersion,
		EngineVersion: EngineVersion, TaxonomyVersion: TaxonomyVersion,
		Records: []EvaluationRecord{},
		Metadata: map[string]string{
			"input_contract": health.SchemaName + "/" + health.SchemaVersion,
			"pipeline":       "canonical-health-record-to-rule-evaluation-record",
		},
	}
}

func classifyDefinition(definition Definition) (Outcome, string) {
	if definition.ContractVersion != RuleVersion {
		return UnsupportedRule, "unsupported_rule_contract"
	}
	if definition.Category == ExtensionRule {
		return UnsupportedRule, "unsupported_rule_category"
	}
	if !validCategory(definition.Category) {
		return InvalidRule, "invalid_rule_category"
	}
	if !ruleIDPattern.MatchString(definition.ID) || definition.Description == "" ||
		definition.Metadata == nil || len(definition.Metadata) > 16 ||
		!sortedUniqueStrings(definition.Scope.HealthIDs) ||
		!sortedUniqueCategories(definition.Scope.Categories) ||
		!sortedUniqueFields(definition.InputRequirements) {
		return InvalidRule, "invalid_rule_contract"
	}
	nodes := 0
	if outcome, reason := validateCondition(definition.Condition, 1, &nodes); outcome != "" {
		return outcome, reason
	}
	return "", ""
}

func validateCondition(condition Condition, depth int, nodes *int) (Outcome, string) {
	*nodes++
	if depth > MaxDepth || *nodes > MaxNodes {
		return InvalidRule, "condition_bounds_exceeded"
	}
	switch condition.Operator {
	case All, Any:
		if condition.Field != "" || valueCount(condition.Value) != 0 ||
			len(condition.Values) != 0 || len(condition.Children) < 2 {
			return InvalidRule, "invalid_logical_condition"
		}
		for _, child := range condition.Children {
			if outcome, reason := validateCondition(child, depth+1, nodes); outcome != "" {
				return outcome, reason
			}
		}
		return "", ""
	case Not:
		if condition.Field != "" || valueCount(condition.Value) != 0 ||
			len(condition.Values) != 0 || len(condition.Children) != 1 {
			return InvalidRule, "invalid_not_condition"
		}
		return validateCondition(condition.Children[0], depth+1, nodes)
	case Equal, NotEqual:
		if !validField(condition.Field) || valueCount(condition.Value) != 1 ||
			len(condition.Values) != 0 || len(condition.Children) != 0 ||
			!valueCompatible(condition.Field, condition.Value) {
			return InvalidRule, "invalid_comparison_condition"
		}
	case GreaterThan, GreaterThanOrEqual, LessThan, LessThanOrEqual:
		if condition.Field != FieldConfidence || condition.Value.Number == nil ||
			valueCount(condition.Value) != 1 || len(condition.Values) != 0 ||
			len(condition.Children) != 0 {
			return InvalidRule, "invalid_ordering_condition"
		}
	case In:
		if !validField(condition.Field) || valueCount(condition.Value) != 0 ||
			len(condition.Values) == 0 || len(condition.Children) != 0 {
			return InvalidRule, "invalid_membership_condition"
		}
		for _, value := range condition.Values {
			if valueCount(value) != 1 || !valueCompatible(condition.Field, value) {
				return InvalidRule, "invalid_membership_value"
			}
		}
	case Exists:
		if !validField(condition.Field) || valueCount(condition.Value) != 0 ||
			len(condition.Values) != 0 || len(condition.Children) != 0 {
			return InvalidRule, "invalid_exists_condition"
		}
	case StatusMatches:
		if condition.Field != FieldStatus || condition.Value.String == nil ||
			valueCount(condition.Value) != 1 || !validHealthStatus(*condition.Value.String) ||
			len(condition.Values) != 0 || len(condition.Children) != 0 {
			return InvalidRule, "invalid_status_condition"
		}
	case CategoryMatches:
		if condition.Field != FieldCategory || condition.Value.String == nil ||
			valueCount(condition.Value) != 1 ||
			!validHealthCategory(*condition.Value.String) ||
			len(condition.Values) != 0 ||
			len(condition.Children) != 0 {
			return InvalidRule, "invalid_category_condition"
		}
	default:
		return UnsupportedRule, "unsupported_operator"
	}
	return "", ""
}

func selectRecords(scope Scope, records []health.Record) []health.Record {
	selected := make([]health.Record, 0, len(records))
	for _, record := range records {
		if len(scope.HealthIDs) > 0 && !containsString(scope.HealthIDs, record.ID) {
			continue
		}
		if len(scope.Categories) > 0 && !containsCategory(scope.Categories, record.Category) {
			continue
		}
		selected = append(selected, record)
	}
	sort.Slice(selected, func(i, j int) bool { return selected[i].ID < selected[j].ID })
	return selected
}

func evaluateCondition(condition Condition, record health.Record) bool {
	switch condition.Operator {
	case All:
		for _, child := range condition.Children {
			if !evaluateCondition(child, record) {
				return false
			}
		}
		return true
	case Any:
		for _, child := range condition.Children {
			if evaluateCondition(child, record) {
				return true
			}
		}
		return false
	case Not:
		return !evaluateCondition(condition.Children[0], record)
	case Exists:
		return fieldValue(condition.Field, record) != (Value{})
	case In:
		actual := fieldValue(condition.Field, record)
		for _, expected := range condition.Values {
			if equalValue(actual, expected) {
				return true
			}
		}
		return false
	case StatusMatches, CategoryMatches, Equal:
		return equalValue(fieldValue(condition.Field, record), condition.Value)
	case NotEqual:
		return !equalValue(fieldValue(condition.Field, record), condition.Value)
	case GreaterThan, GreaterThanOrEqual, LessThan, LessThanOrEqual:
		actual, expected := fieldValue(condition.Field, record).Number, condition.Value.Number
		switch condition.Operator {
		case GreaterThan:
			return *actual > *expected
		case GreaterThanOrEqual:
			return *actual >= *expected
		case LessThan:
			return *actual < *expected
		default:
			return *actual <= *expected
		}
	default:
		return false
	}
}

func fieldValue(field Field, record health.Record) Value {
	switch field {
	case FieldHealthID:
		return stringValue(record.ID)
	case FieldCategory:
		return stringValue(string(record.Category))
	case FieldStatus:
		return stringValue(string(record.Status))
	case FieldEvidence:
		return stringValue(string(record.EvidenceState))
	case FieldConfidence:
		value := int64(record.ConfidenceBasisPoints)
		return Value{Number: &value}
	case FieldReason:
		return stringValue(record.Reason)
	case FieldLayer:
		return stringValue(record.Scope.Layer)
	case FieldObjectID:
		return stringValue(record.Scope.ObjectID)
	case FieldPath:
		return stringValue(record.Scope.Path)
	default:
		return Value{}
	}
}

func terminalRecord(definition Definition, outcome Outcome, status EvaluationStatus, explanation string) EvaluationRecord {
	record := EvaluationRecord{
		RuleID: definition.ID, Outcome: outcome, EvaluationStatus: status,
		MatchResult: Indeterminate, ConfidenceBasisPoints: 0,
		Explanation: explanation, EvidenceReferences: []string{},
		Metadata: map[string]string{
			"rule_category": string(definition.Category),
			"root_operator": string(definition.Condition.Operator),
		},
		Versions: expectedVersions(),
	}
	record.ID = stableID(record)
	return record
}

// Validate rejects unsupported or internally inconsistent result contracts.
func Validate(result Result) error {
	if result.SchemaName != SchemaName || result.SchemaVersion != SchemaVersion ||
		result.EngineVersion != EngineVersion || result.TaxonomyVersion != TaxonomyVersion ||
		result.Records == nil || result.Metadata == nil || len(result.Metadata) != 2 ||
		result.Metadata["input_contract"] != health.SchemaName+"/"+health.SchemaVersion ||
		result.Metadata["pipeline"] != "canonical-health-record-to-rule-evaluation-record" {
		return fmt.Errorf("invalid rule evaluation envelope")
	}
	seen := map[string]bool{}
	last := ""
	for _, record := range result.Records {
		if record.ID == "" || record.RuleID == "" || seen[record.ID] ||
			!validOutcome(record.Outcome) || !validEvaluationStatus(record.EvaluationStatus) ||
			!validMatchResult(record.MatchResult) ||
			record.ConfidenceBasisPoints < 0 || record.ConfidenceBasisPoints > 10000 ||
			record.Explanation == "" || record.EvidenceReferences == nil ||
			record.Metadata == nil || len(record.Metadata) != 2 ||
			record.Versions != expectedVersions() || record.ID != stableID(record) {
			return fmt.Errorf("invalid rule evaluation record")
		}
		if record.Outcome == Matched || record.Outcome == NotMatched {
			if record.HealthRecordID == "" || len(record.EvidenceReferences) != 1 ||
				record.EvidenceReferences[0] != record.HealthRecordID ||
				record.EvaluationStatus != EvaluationComplete ||
				record.ConfidenceBasisPoints != 10000 {
				return fmt.Errorf("invalid completed evaluation")
			}
			if (record.Outcome == Matched &&
				(record.MatchResult != Match || record.Explanation != "condition_matched")) ||
				(record.Outcome == NotMatched &&
					(record.MatchResult != NoMatch || record.Explanation != "condition_not_matched")) {
				return fmt.Errorf("invalid match result")
			}
		} else if record.HealthRecordID != "" || len(record.EvidenceReferences) != 0 ||
			record.MatchResult != Indeterminate || record.ConfidenceBasisPoints != 0 {
			return fmt.Errorf("invalid terminal evaluation")
		} else if !validTerminal(record) {
			return fmt.Errorf("invalid terminal outcome")
		}
		key := record.RuleID + "\x00" + record.HealthRecordID + "\x00" + record.ID
		if last != "" && key <= last {
			return fmt.Errorf("rule evaluations are not uniquely ordered")
		}
		last, seen[record.ID] = key, true
	}
	return nil
}

// MarshalCanonical validates and emits byte-stable canonical JSON.
func MarshalCanonical(result Result) ([]byte, error) {
	if err := Validate(result); err != nil {
		return nil, err
	}
	return json.Marshal(result)
}

func stableID(record EvaluationRecord) string {
	hash := sha256.New()
	hash.Write([]byte("qwsg.rule/" + EngineVersion + "/evaluation"))
	for _, part := range []string{
		record.RuleID, record.HealthRecordID, string(record.Outcome),
		string(record.EvaluationStatus), string(record.MatchResult),
		strconv.Itoa(record.ConfidenceBasisPoints), record.Explanation,
		strings.Join(record.EvidenceReferences, "\x1f"),
		record.Metadata["rule_category"], record.Metadata["root_operator"],
	} {
		hash.Write([]byte{0})
		hash.Write([]byte(part))
	}
	return hex.EncodeToString(hash.Sum(nil))
}

func expectedVersions() VersionInfo {
	return VersionInfo{
		EvaluationSchema: SchemaVersion, EvaluationEngine: EngineVersion,
		EvaluationTaxonomy: TaxonomyVersion, RuleContract: RuleVersion,
		HealthSchema: health.SchemaVersion, HealthTaxonomy: health.TaxonomyVersion,
	}
}

func sortRecords(records []EvaluationRecord) {
	sort.Slice(records, func(i, j int) bool {
		if records[i].RuleID != records[j].RuleID {
			return records[i].RuleID < records[j].RuleID
		}
		if records[i].HealthRecordID != records[j].HealthRecordID {
			return records[i].HealthRecordID < records[j].HealthRecordID
		}
		return records[i].ID < records[j].ID
	})
}

func valueCount(value Value) int {
	count := 0
	if value.String != nil {
		count++
	}
	if value.Number != nil {
		count++
	}
	if value.Boolean != nil {
		count++
	}
	return count
}

func valueCompatible(field Field, value Value) bool {
	if field == FieldConfidence {
		return value.Number != nil
	}
	if value.String == nil {
		return false
	}
	switch field {
	case FieldStatus:
		return validHealthStatus(*value.String)
	case FieldCategory:
		return validHealthCategory(*value.String)
	case FieldEvidence:
		return validHealthEvidence(*value.String)
	default:
		return true
	}
}

func equalValue(left, right Value) bool {
	switch {
	case left.String != nil && right.String != nil:
		return *left.String == *right.String
	case left.Number != nil && right.Number != nil:
		return *left.Number == *right.Number
	case left.Boolean != nil && right.Boolean != nil:
		return *left.Boolean == *right.Boolean
	default:
		return false
	}
}

func stringValue(value string) Value { return Value{String: &value} }

func validCategory(category Category) bool {
	switch category {
	case StatusRule, CategoryRule, EvidenceRule, CompositeRule:
		return true
	default:
		return false
	}
}

func validField(field Field) bool {
	switch field {
	case FieldHealthID, FieldCategory, FieldStatus, FieldEvidence, FieldConfidence,
		FieldReason, FieldLayer, FieldObjectID, FieldPath:
		return true
	default:
		return false
	}
}

func validHealthStatus(value string) bool {
	switch health.Status(value) {
	case health.Healthy, health.Informational, health.Advisory, health.Warning,
		health.Critical, health.Unknown, health.Unsupported:
		return true
	default:
		return false
	}
}

func validHealthCategory(value string) bool {
	switch drift.Category(value) {
	case drift.IdentityDrift, drift.SoftwareDrift, drift.HardwareDrift,
		drift.PlatformDrift, drift.FilesystemDrift, drift.StorageDrift,
		drift.NetworkDrift, drift.ServiceDrift, drift.ConfigurationDrift,
		drift.SecurityDrift, drift.CapabilityDrift, drift.EnvironmentDrift,
		drift.ExtensionDrift:
		return true
	default:
		return false
	}
}

func validHealthEvidence(value string) bool {
	switch health.EvidenceState(value) {
	case health.EvidenceSufficient, health.EvidenceInsufficient,
		health.EvidenceUnsupported:
		return true
	default:
		return false
	}
}

func validTerminal(record EvaluationRecord) bool {
	switch record.Outcome {
	case InsufficientEvidence:
		return record.EvaluationStatus == EvaluationSkipped &&
			record.Explanation == "no_health_record_in_scope"
	case UnsupportedRule:
		return record.EvaluationStatus == EvaluationSkipped &&
			(record.Explanation == "unsupported_rule_contract" ||
				record.Explanation == "unsupported_rule_category" ||
				record.Explanation == "unsupported_operator")
	case InvalidRule:
		return record.EvaluationStatus == EvaluationFailed &&
			(strings.HasPrefix(record.Explanation, "invalid_") ||
				record.Explanation == "condition_bounds_exceeded")
	case EvaluationError:
		return record.EvaluationStatus == EvaluationFailed &&
			record.Explanation == "invalid_health_contract"
	case DisabledRule:
		return record.EvaluationStatus == EvaluationSkipped &&
			record.Explanation == "rule_disabled"
	default:
		return false
	}
}

func validOutcome(outcome Outcome) bool {
	switch outcome {
	case Matched, NotMatched, InsufficientEvidence, UnsupportedRule, InvalidRule,
		EvaluationError, DisabledRule:
		return true
	default:
		return false
	}
}

func validEvaluationStatus(status EvaluationStatus) bool {
	return status == EvaluationComplete || status == EvaluationSkipped || status == EvaluationFailed
}

func validMatchResult(result MatchResult) bool {
	return result == Match || result == NoMatch || result == Indeterminate
}

func sortedUniqueStrings(values []string) bool {
	for i, value := range values {
		if value == "" || (i > 0 && value <= values[i-1]) {
			return false
		}
	}
	return values != nil
}

func sortedUniqueCategories(values []drift.Category) bool {
	for i, value := range values {
		if value == "" || (i > 0 && value <= values[i-1]) {
			return false
		}
	}
	return values != nil
}

func sortedUniqueFields(values []Field) bool {
	for i, value := range values {
		if !validField(value) || (i > 0 && value <= values[i-1]) {
			return false
		}
	}
	return values != nil
}

func containsString(values []string, value string) bool {
	index := sort.SearchStrings(values, value)
	return index < len(values) && values[index] == value
}

func containsCategory(values []drift.Category, value drift.Category) bool {
	index := sort.Search(len(values), func(i int) bool { return values[i] >= value })
	return index < len(values) && values[index] == value
}
