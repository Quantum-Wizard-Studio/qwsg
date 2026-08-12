// Package policy deterministically interprets immutable canonical Rule
// Evaluation Records through versioned Policy Profiles. It performs no rule
// evaluation, scheduling, alerting, notification, remediation, host mutation,
// networking, presentation, or AI work.
package policy

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"regexp"
	"sort"

	"quantumwizard.hu/qwsg/internal/rule"
)

const (
	SchemaName        = "qwsg.policy-evaluation"
	SchemaVersion     = "1.0"
	EngineVersion     = "1.0"
	TaxonomyVersion   = "1.0"
	ProfileVersion    = "1.0"
	MaxProfiles       = 64
	MaxStatements     = 256
	MaxRecords        = 4096
	MaxInheritance    = 8
	MaxMetadataFields = 16
)

type Outcome string

const (
	Accepted      Outcome = "accepted"
	Observe       Outcome = "observe"
	Suppressed    Outcome = "suppressed"
	Escalated     Outcome = "escalated"
	Indeterminate Outcome = "indeterminate"
	NotApplicable Outcome = "not_applicable"
	Conflict      Outcome = "conflict"
)

type EvaluationStatus string

const (
	EvaluationComplete EvaluationStatus = "complete"
	EvaluationSkipped  EvaluationStatus = "skipped"
)

// Selector restricts a profile or statement to canonical Rule Evaluation
// fields. Empty lists are wildcards. Lists must be sorted and unique.
type Selector struct {
	RuleIDs  []string       `json:"rule_ids"`
	Outcomes []rule.Outcome `json:"outcomes"`
}

type Statement struct {
	ID          string            `json:"id"`
	Priority    int               `json:"priority"`
	Selector    Selector          `json:"selector"`
	Outcome     Outcome           `json:"outcome"`
	Explanation string            `json:"explanation"`
	Metadata    map[string]string `json:"metadata"`
}

// Profile is the Policy Profile 1.0 public contract. Identity is a digest of
// every other field and makes profile provenance independently verifiable.
type Profile struct {
	ID              string            `json:"id"`
	Identity        string            `json:"identity"`
	ContractVersion string            `json:"contract_version"`
	Version         string            `json:"version"`
	Priority        int               `json:"priority"`
	Extends         []string          `json:"extends"`
	Enabled         bool              `json:"enabled"`
	Scope           Selector          `json:"scope"`
	Statements      []Statement       `json:"statements"`
	DefaultOutcome  Outcome           `json:"default_outcome"`
	Metadata        map[string]string `json:"metadata"`
}

type VersionInfo struct {
	EvaluationSchema   string `json:"evaluation_schema"`
	EvaluationEngine   string `json:"evaluation_engine"`
	EvaluationTaxonomy string `json:"evaluation_taxonomy"`
	PolicyProfile      string `json:"policy_profile"`
	RuleEvaluation     string `json:"rule_evaluation"`
	RuleEngine         string `json:"rule_engine"`
	RuleTaxonomy       string `json:"rule_taxonomy"`
	RuleContract       string `json:"rule_contract"`
}

// EvaluationRecord is Policy Evaluation Record 1.0.
type EvaluationRecord struct {
	ID                  string            `json:"id"`
	RuleEvaluationID    string            `json:"rule_evaluation_id"`
	RuleID              string            `json:"rule_id"`
	Outcome             Outcome           `json:"outcome"`
	EvaluationStatus    EvaluationStatus  `json:"evaluation_status"`
	AppliedProfileIDs   []string          `json:"applied_profile_ids"`
	AppliedStatementIDs []string          `json:"applied_statement_ids"`
	Explanation         string            `json:"explanation"`
	EvidenceReferences  []string          `json:"evidence_references"`
	Metadata            map[string]string `json:"metadata"`
	Versions            VersionInfo       `json:"versions"`
}

type Result struct {
	SchemaName      string             `json:"schema_name"`
	SchemaVersion   string             `json:"schema_version"`
	EngineVersion   string             `json:"engine_version"`
	TaxonomyVersion string             `json:"taxonomy_version"`
	ProfileIDs      []string           `json:"profile_ids"`
	Records         []EvaluationRecord `json:"records"`
	Metadata        map[string]string  `json:"metadata"`
}

type candidate struct {
	profileID     string
	statementID   string
	profileRank   int
	statementRank int
	outcome       Outcome
}

var canonicalID = regexp.MustCompile(`^[a-z0-9]+([._:-][a-z0-9]+)*$`)
var versionID = regexp.MustCompile(`^[0-9]+\.[0-9]+([.-][a-z0-9]+)*$`)

// NormalizeProfile supplies deterministic empty collections, sorts statements,
// and computes the content identity. It does not repair invalid semantics.
func NormalizeProfile(profile Profile) (Profile, error) {
	claimed := profile.Identity
	if profile.Extends == nil {
		profile.Extends = []string{}
	}
	if profile.Scope.RuleIDs == nil {
		profile.Scope.RuleIDs = []string{}
	}
	if profile.Scope.Outcomes == nil {
		profile.Scope.Outcomes = []rule.Outcome{}
	}
	if profile.Statements == nil {
		profile.Statements = []Statement{}
	}
	if profile.Metadata == nil {
		profile.Metadata = map[string]string{}
	}
	for i := range profile.Statements {
		if profile.Statements[i].Selector.RuleIDs == nil {
			profile.Statements[i].Selector.RuleIDs = []string{}
		}
		if profile.Statements[i].Selector.Outcomes == nil {
			profile.Statements[i].Selector.Outcomes = []rule.Outcome{}
		}
		if profile.Statements[i].Metadata == nil {
			profile.Statements[i].Metadata = map[string]string{}
		}
	}
	sort.Slice(profile.Statements, func(i, j int) bool { return profile.Statements[i].ID < profile.Statements[j].ID })
	profile.Identity = ""
	profile.Identity = profileIdentity(profile)
	if claimed != "" && claimed != profile.Identity {
		return Profile{}, fmt.Errorf("invalid policy profile identity")
	}
	if err := validateProfileShape(profile); err != nil {
		return Profile{}, err
	}
	return profile, nil
}

// Evaluate applies every enabled applicable profile. Higher profile priority
// wins; within it, higher statement priority wins. Equal-precedence different
// outcomes produce an explicit conflict rather than relying on input order.
func Evaluate(profiles []Profile, input rule.Result) (Result, error) {
	if err := rule.Validate(input); err != nil {
		return Result{}, fmt.Errorf("invalid rule evaluation source: %w", err)
	}
	if len(input.Records) > MaxRecords {
		return Result{}, fmt.Errorf("policy evaluation record limit exceeded")
	}
	if len(profiles) == 0 || len(profiles) > MaxProfiles {
		return Result{}, fmt.Errorf("policy profile count must be between 1 and %d", MaxProfiles)
	}
	profiles = append([]Profile(nil), profiles...)
	sort.Slice(profiles, func(i, j int) bool { return profiles[i].ID < profiles[j].ID })
	byID := make(map[string]Profile, len(profiles))
	profileIDs := make([]string, 0, len(profiles))
	for i, profile := range profiles {
		if err := validateProfileShape(profile); err != nil {
			return Result{}, fmt.Errorf("policy profile %q: %w", profile.ID, err)
		}
		if i > 0 && profile.ID == profiles[i-1].ID {
			return Result{}, fmt.Errorf("policy profile ids must be unique")
		}
		byID[profile.ID] = profile
		profileIDs = append(profileIDs, profile.ID)
	}
	if err := validateInheritance(profiles, byID); err != nil {
		return Result{}, err
	}
	result := Result{
		SchemaName: SchemaName, SchemaVersion: SchemaVersion, EngineVersion: EngineVersion,
		TaxonomyVersion: TaxonomyVersion, ProfileIDs: profileIDs, Records: []EvaluationRecord{},
		Metadata: map[string]string{"input_contract": rule.SchemaName + "/" + rule.SchemaVersion, "pipeline": "canonical-rule-evaluation-to-policy-evaluation"},
	}
	for _, source := range input.Records {
		candidates := []candidate{}
		for _, profile := range profiles {
			if !profile.Enabled || !matches(profile.Scope, source) {
				continue
			}
			statements := inheritedStatements(profile, byID, map[string]bool{})
			matched := false
			for _, statement := range statements {
				if matches(statement.Selector, source) {
					matched = true
					candidates = append(candidates, candidate{profile.ID, statement.ID, profile.Priority, statement.Priority, statement.Outcome})
				}
			}
			if !matched {
				candidates = append(candidates, candidate{profile.ID, "", profile.Priority, -1001, profile.DefaultOutcome})
			}
		}
		result.Records = append(result.Records, resolve(source, candidates))
	}
	sort.Slice(result.Records, func(i, j int) bool {
		if result.Records[i].RuleID != result.Records[j].RuleID {
			return result.Records[i].RuleID < result.Records[j].RuleID
		}
		return result.Records[i].RuleEvaluationID < result.Records[j].RuleEvaluationID
	})
	if err := Validate(result); err != nil {
		return Result{}, err
	}
	return result, nil
}

func resolve(source rule.EvaluationRecord, candidates []candidate) EvaluationRecord {
	record := EvaluationRecord{
		RuleEvaluationID: source.ID, RuleID: source.RuleID, Outcome: NotApplicable,
		EvaluationStatus: EvaluationSkipped, AppliedProfileIDs: []string{}, AppliedStatementIDs: []string{},
		Explanation: "no_applicable_policy", EvidenceReferences: []string{source.ID},
		Metadata: map[string]string{"source_outcome": string(source.Outcome)}, Versions: expectedVersions(),
	}
	if len(candidates) == 0 {
		record.ID = recordIdentity(record)
		return record
	}
	sort.Slice(candidates, func(i, j int) bool {
		if candidates[i].profileRank != candidates[j].profileRank {
			return candidates[i].profileRank > candidates[j].profileRank
		}
		if candidates[i].statementRank != candidates[j].statementRank {
			return candidates[i].statementRank > candidates[j].statementRank
		}
		if candidates[i].profileID != candidates[j].profileID {
			return candidates[i].profileID < candidates[j].profileID
		}
		return candidates[i].statementID < candidates[j].statementID
	})
	topProfile, topStatement := candidates[0].profileRank, candidates[0].statementRank
	winners := []candidate{}
	for _, value := range candidates {
		if value.profileRank != topProfile || value.statementRank != topStatement {
			break
		}
		winners = append(winners, value)
	}
	outcomes := map[Outcome]bool{}
	for _, winner := range winners {
		outcomes[winner.outcome] = true
		record.AppliedProfileIDs = append(record.AppliedProfileIDs, winner.profileID)
		if winner.statementID != "" {
			record.AppliedStatementIDs = append(record.AppliedStatementIDs, winner.statementID)
		}
	}
	record.AppliedProfileIDs = uniqueSorted(record.AppliedProfileIDs)
	record.AppliedStatementIDs = uniqueSorted(record.AppliedStatementIDs)
	record.EvaluationStatus = EvaluationComplete
	if len(outcomes) > 1 {
		record.Outcome, record.Explanation = Conflict, "equal_precedence_policy_conflict"
	} else {
		record.Outcome = winners[0].outcome
		if winners[0].statementID == "" {
			record.Explanation = "profile_default_outcome"
		} else {
			record.Explanation = "policy_statement_applied"
		}
	}
	record.ID = recordIdentity(record)
	return record
}

// Validate rejects unsupported, noncanonical, or untraceable Policy results.
func Validate(result Result) error {
	if result.SchemaName != SchemaName || result.SchemaVersion != SchemaVersion || result.EngineVersion != EngineVersion ||
		result.TaxonomyVersion != TaxonomyVersion || result.ProfileIDs == nil || !sortedUniqueStrings(result.ProfileIDs) ||
		result.Records == nil || len(result.Records) > MaxRecords || result.Metadata == nil || len(result.Metadata) != 2 ||
		result.Metadata["input_contract"] != rule.SchemaName+"/"+rule.SchemaVersion ||
		result.Metadata["pipeline"] != "canonical-rule-evaluation-to-policy-evaluation" {
		return fmt.Errorf("invalid policy result envelope")
	}
	last := ""
	seen := map[string]bool{}
	knownProfiles := map[string]bool{}
	for _, id := range result.ProfileIDs {
		knownProfiles[id] = true
	}
	for _, record := range result.Records {
		key := record.RuleID + "\x00" + record.RuleEvaluationID
		if key <= last || seen[record.ID] || !validRecord(record) || record.ID != recordIdentity(record) {
			return fmt.Errorf("invalid or unordered policy evaluation record")
		}
		for _, id := range record.AppliedProfileIDs {
			if !knownProfiles[id] {
				return fmt.Errorf("policy evaluation references an unknown profile")
			}
		}
		last, seen[record.ID] = key, true
	}
	return nil
}

func MarshalCanonical(result Result) ([]byte, error) {
	if err := Validate(result); err != nil {
		return nil, err
	}
	return json.Marshal(result)
}

func MarshalProfilesCanonical(profiles []Profile) ([]byte, error) {
	copy := append([]Profile(nil), profiles...)
	sort.Slice(copy, func(i, j int) bool { return copy[i].ID < copy[j].ID })
	seen := map[string]bool{}
	byID := map[string]Profile{}
	for _, profile := range copy {
		if seen[profile.ID] {
			return nil, fmt.Errorf("policy profile ids must be unique")
		}
		seen[profile.ID] = true
		if err := validateProfileShape(profile); err != nil {
			return nil, err
		}
		byID[profile.ID] = profile
	}
	if err := validateInheritance(copy, byID); err != nil {
		return nil, err
	}
	return json.Marshal(copy)
}

func validateProfileShape(profile Profile) error {
	if !canonicalID.MatchString(profile.ID) || profile.ContractVersion != ProfileVersion || !versionID.MatchString(profile.Version) ||
		profile.Priority < -1000 || profile.Priority > 1000 || profile.Identity != profileIdentity(profile) ||
		!sortedUniqueStrings(profile.Extends) || !validSelector(profile.Scope) || !configurableOutcome(profile.DefaultOutcome) ||
		profile.Metadata == nil || len(profile.Metadata) > MaxMetadataFields || len(profile.Statements) > MaxStatements {
		return fmt.Errorf("invalid or unsupported policy profile contract")
	}
	last := ""
	for _, statement := range profile.Statements {
		if !canonicalID.MatchString(statement.ID) || statement.ID <= last || statement.Priority < -1000 || statement.Priority > 1000 ||
			!validSelector(statement.Selector) || !configurableOutcome(statement.Outcome) || statement.Explanation == "" ||
			statement.Metadata == nil || len(statement.Metadata) > MaxMetadataFields {
			return fmt.Errorf("invalid policy statement")
		}
		last = statement.ID
	}
	return nil
}

func validateInheritance(profiles []Profile, byID map[string]Profile) error {
	for _, profile := range profiles {
		for _, parent := range profile.Extends {
			if _, ok := byID[parent]; !ok {
				return fmt.Errorf("unknown parent policy profile %q", parent)
			}
		}
		if inheritanceDepth(profile.ID, byID, map[string]bool{}) > MaxInheritance {
			return fmt.Errorf("policy inheritance cycle or depth exceeded")
		}
		seenStatements := map[string]bool{}
		for _, statement := range inheritedStatements(profile, byID, map[string]bool{}) {
			if seenStatements[statement.ID] {
				return fmt.Errorf("duplicate inherited policy statement %q", statement.ID)
			}
			seenStatements[statement.ID] = true
		}
	}
	return nil
}

func inheritanceDepth(id string, byID map[string]Profile, path map[string]bool) int {
	if path[id] {
		return MaxInheritance + 1
	}
	path[id] = true
	max := 1
	for _, parent := range byID[id].Extends {
		copyPath := make(map[string]bool, len(path))
		for key, value := range path {
			copyPath[key] = value
		}
		if depth := 1 + inheritanceDepth(parent, byID, copyPath); depth > max {
			max = depth
		}
	}
	return max
}

func inheritedStatements(profile Profile, byID map[string]Profile, seen map[string]bool) []Statement {
	if seen[profile.ID] {
		return nil
	}
	seen[profile.ID] = true
	result := []Statement{}
	for _, parent := range profile.Extends {
		result = append(result, inheritedStatements(byID[parent], byID, seen)...)
	}
	result = append(result, profile.Statements...)
	sort.Slice(result, func(i, j int) bool { return result[i].ID < result[j].ID })
	return result
}

func matches(selector Selector, record rule.EvaluationRecord) bool {
	return (len(selector.RuleIDs) == 0 || containsString(selector.RuleIDs, record.RuleID)) &&
		(len(selector.Outcomes) == 0 || containsRuleOutcome(selector.Outcomes, record.Outcome))
}

func validSelector(selector Selector) bool {
	return selector.RuleIDs != nil && selector.Outcomes != nil && sortedUniqueStrings(selector.RuleIDs) && sortedUniqueRuleOutcomes(selector.Outcomes)
}
func configurableOutcome(value Outcome) bool {
	switch value {
	case Accepted, Observe, Suppressed, Escalated, Indeterminate:
		return true
	}
	return false
}
func containsString(values []string, value string) bool {
	index := sort.SearchStrings(values, value)
	return index < len(values) && values[index] == value
}
func containsRuleOutcome(values []rule.Outcome, value rule.Outcome) bool {
	index := sort.Search(len(values), func(i int) bool { return values[i] >= value })
	return index < len(values) && values[index] == value
}

func validRecord(record EvaluationRecord) bool {
	if record.RuleEvaluationID == "" || record.RuleID == "" || record.EvidenceReferences == nil || len(record.EvidenceReferences) != 1 ||
		record.EvidenceReferences[0] != record.RuleEvaluationID || record.AppliedProfileIDs == nil || record.AppliedStatementIDs == nil ||
		!sortedUniqueStrings(record.AppliedProfileIDs) || !sortedUniqueStrings(record.AppliedStatementIDs) || record.Explanation == "" ||
		record.Metadata == nil || len(record.Metadata) != 1 || record.Metadata["source_outcome"] == "" || record.Versions != expectedVersions() {
		return false
	}
	if record.Outcome == NotApplicable {
		return record.EvaluationStatus == EvaluationSkipped && len(record.AppliedProfileIDs) == 0 && len(record.AppliedStatementIDs) == 0 && record.Explanation == "no_applicable_policy"
	}
	if record.EvaluationStatus != EvaluationComplete || len(record.AppliedProfileIDs) == 0 {
		return false
	}
	if record.Outcome == Conflict {
		return record.Explanation == "equal_precedence_policy_conflict" &&
			(len(record.AppliedProfileIDs) >= 2 || len(record.AppliedStatementIDs) >= 2)
	}
	return configurableOutcome(record.Outcome) && (record.Explanation == "policy_statement_applied" || record.Explanation == "profile_default_outcome")
}

func expectedVersions() VersionInfo {
	return VersionInfo{SchemaVersion, EngineVersion, TaxonomyVersion, ProfileVersion, rule.SchemaVersion, rule.EngineVersion, rule.TaxonomyVersion, rule.RuleVersion}
}
func profileIdentity(profile Profile) string {
	copy := profile
	copy.Identity = ""
	document, _ := json.Marshal(copy)
	return stableID("profile", document)
}
func recordIdentity(record EvaluationRecord) string {
	copy := record
	copy.ID = ""
	document, _ := json.Marshal(copy)
	return stableID("record", document)
}
func stableID(domain string, document []byte) string {
	sum := sha256.Sum256(append([]byte(SchemaName+"/"+SchemaVersion+"/"+domain+"\x00"), document...))
	return hex.EncodeToString(sum[:])
}

func uniqueSorted(values []string) []string {
	sort.Strings(values)
	result := values[:0]
	for _, value := range values {
		if len(result) == 0 || result[len(result)-1] != value {
			result = append(result, value)
		}
	}
	return result
}
func sortedUniqueStrings(values []string) bool {
	for i, value := range values {
		if value == "" || (i > 0 && value <= values[i-1]) {
			return false
		}
	}
	return true
}
func sortedUniqueRuleOutcomes(values []rule.Outcome) bool {
	for i, value := range values {
		if !validRuleOutcome(value) || (i > 0 && value <= values[i-1]) {
			return false
		}
	}
	return true
}
func validRuleOutcome(value rule.Outcome) bool {
	switch value {
	case rule.Matched, rule.NotMatched, rule.InsufficientEvidence, rule.UnsupportedRule, rule.InvalidRule, rule.EvaluationError, rule.DisabledRule:
		return true
	}
	return false
}
