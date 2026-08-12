// Package configuration defines the immutable Canonical Configuration
// Contract. It normalizes, validates, and resolves supplied data; it performs
// no file discovery, activation, scheduling, clock access, secret resolution,
// pipeline execution, networking, process execution, or host mutation.
package configuration

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"regexp"
	"sort"
	"strings"

	"quantumwizard.hu/qwsg/internal/drift"
	"quantumwizard.hu/qwsg/internal/policy"
	"quantumwizard.hu/qwsg/internal/rule"
)

const (
	SourceSchema       = "qwsg.configuration-source"
	EffectiveSchema    = "qwsg.effective-configuration"
	SchemaVersion      = "1.0"
	ModelVersion       = "1.0"
	ScheduleVersion    = "1.0"
	MaxSources         = 32
	MaxItems           = 1024
	MaxStringLength    = 4096
	MaxConcurrency     = 256
	MaxSnapshotRetain  = 10000
	MaxDurationNS      = int64(365 * 24 * 60 * 60 * 1_000_000_000)
	DefaultTimeoutNS   = int64(5 * 60 * 1_000_000_000)
	DefaultRetention   = 10
	DefaultConcurrency = 1
)

type SourceKind string

const (
	BuiltInDefault    SourceKind = "built_in_default"
	PrimaryLocal      SourceKind = "primary_local"
	ActivatedOverride SourceKind = "activated_local_override"
	TemporaryOverride SourceKind = "command_temporary_override"
)

type TriggerKind string

const (
	IntervalTrigger TriggerKind = "interval"
	CalendarTrigger TriggerKind = "calendar"
)

type MisfirePolicy string

const (
	MisfireSkip          MisfirePolicy = "skip"
	MisfireRunOnce       MisfirePolicy = "run_once"
	MisfireIndeterminate MisfirePolicy = "indeterminate"
)

type OverlapPolicy string

const (
	OverlapForbid OverlapPolicy = "forbid"
	OverlapAllow  OverlapPolicy = "allow"
)

type DSTPolicy string

const (
	DSTFirstOccurrence  DSTPolicy = "first_occurrence"
	DSTSecondOccurrence DSTPolicy = "second_occurrence"
	DSTSkipNonexistent  DSTPolicy = "skip_nonexistent"
)

type Check struct {
	ID        string   `json:"id"`
	Enabled   bool     `json:"enabled"`
	TargetIDs []string `json:"target_ids"`
}

type Target struct {
	ID       string            `json:"id"`
	Kind     string            `json:"kind"`
	Metadata map[string]string `json:"metadata"`
}

type RetryPolicy struct {
	ID             string `json:"id"`
	MaxAttempts    int    `json:"max_attempts"`
	InitialDelayNS int64  `json:"initial_delay_ns"`
	MaxDelayNS     int64  `json:"max_delay_ns"`
}

type Calendar struct {
	Minutes   []int `json:"minutes"`
	Hours     []int `json:"hours"`
	MonthDays []int `json:"month_days"`
	Months    []int `json:"months"`
	Weekdays  []int `json:"weekdays"`
}

// Schedule is Schedule Definition 1.0. It describes intent but never computes
// due work or reads time.
type Schedule struct {
	ID                 string        `json:"id"`
	ContractVersion    string        `json:"contract_version"`
	Enabled            bool          `json:"enabled"`
	TimeZone           string        `json:"time_zone"`
	Trigger            TriggerKind   `json:"trigger"`
	IntervalNS         int64         `json:"interval_ns,omitempty"`
	Calendar           Calendar      `json:"calendar"`
	DSTPolicy          DSTPolicy     `json:"dst_policy"`
	Priority           int           `json:"priority"`
	MisfirePolicy      MisfirePolicy `json:"misfire_policy"`
	OverlapPolicy      OverlapPolicy `json:"overlap_policy"`
	ExecutionTimeoutNS int64         `json:"execution_timeout_ns"`
	RetryPolicyID      string        `json:"retry_policy_id"`
	CheckIDs           []string      `json:"check_ids"`
	CommandProfile     string        `json:"command_profile"`
}

type ReportPolicy struct {
	Enabled        bool   `json:"enabled"`
	ScheduleID     string `json:"schedule_id,omitempty"`
	RetentionCount int    `json:"retention_count"`
}

// SecretReference deliberately has no field capable of holding secret
// material. Provider identifies a future backend contract; Reference is a
// non-secret opaque lookup key.
type SecretReference struct {
	ID        string `json:"id"`
	Provider  string `json:"provider"`
	Reference string `json:"reference"`
}

type Extension struct {
	ID       string            `json:"id"`
	Version  string            `json:"version"`
	Required bool              `json:"required"`
	Fields   map[string]string `json:"fields"`
}

// Model is Canonical Configuration Model 1.0.
type Model struct {
	InstanceID         string            `json:"instance_id"`
	Locale             string            `json:"locale"`
	TimeZone           string            `json:"time_zone"`
	Checks             []Check           `json:"checks"`
	Targets            []Target          `json:"targets"`
	RuleDefinitions    []rule.Definition `json:"rule_definitions"`
	PolicyProfiles     []policy.Profile  `json:"policy_profiles"`
	SnapshotRetention  int               `json:"snapshot_retention"`
	Schedules          []Schedule        `json:"schedules"`
	ExecutionTimeoutNS int64             `json:"execution_timeout_ns"`
	Concurrency        int               `json:"concurrency"`
	RetryPolicies      []RetryPolicy     `json:"retry_policies"`
	Report             ReportPolicy      `json:"report"`
	SecretReferences   []SecretReference `json:"secret_references"`
	Extensions         []Extension       `json:"extensions"`
}

// Patch uses pointers to distinguish an omitted field from an explicit zero
// or empty replacement. Collections replace the lower-precedence field as one
// deterministic value; implicit list merging is prohibited.
type Patch struct {
	InstanceID         *string            `json:"instance_id,omitempty"`
	Locale             *string            `json:"locale,omitempty"`
	TimeZone           *string            `json:"time_zone,omitempty"`
	Checks             *[]Check           `json:"checks,omitempty"`
	Targets            *[]Target          `json:"targets,omitempty"`
	RuleDefinitions    *[]rule.Definition `json:"rule_definitions,omitempty"`
	PolicyProfiles     *[]policy.Profile  `json:"policy_profiles,omitempty"`
	SnapshotRetention  *int               `json:"snapshot_retention,omitempty"`
	Schedules          *[]Schedule        `json:"schedules,omitempty"`
	ExecutionTimeoutNS *int64             `json:"execution_timeout_ns,omitempty"`
	Concurrency        *int               `json:"concurrency,omitempty"`
	RetryPolicies      *[]RetryPolicy     `json:"retry_policies,omitempty"`
	Report             *ReportPolicy      `json:"report,omitempty"`
	SecretReferences   *[]SecretReference `json:"secret_references,omitempty"`
	Extensions         *[]Extension       `json:"extensions,omitempty"`
}

// Source is Configuration Source Record 1.0.
type Source struct {
	SchemaName    string     `json:"schema_name"`
	SchemaVersion string     `json:"schema_version"`
	ModelVersion  string     `json:"model_version"`
	ID            string     `json:"id"`
	Identity      string     `json:"identity"`
	SourceVersion string     `json:"source_version"`
	Kind          SourceKind `json:"kind"`
	Patch         Patch      `json:"patch"`
}

type FieldProvenance struct {
	Field          string     `json:"field"`
	SourceIDs      []string   `json:"source_ids"`
	SourceVersions []string   `json:"source_versions"`
	Kind           SourceKind `json:"kind"`
	Precedence     int        `json:"precedence"`
	Resolution     string     `json:"resolution"`
}

// Effective is Effective Configuration 1.0.
type Effective struct {
	SchemaName    string                     `json:"schema_name"`
	SchemaVersion string                     `json:"schema_version"`
	ModelVersion  string                     `json:"model_version"`
	ID            string                     `json:"id"`
	SourceIDs     []string                   `json:"source_ids"`
	Values        Model                      `json:"values"`
	Provenance    map[string]FieldProvenance `json:"provenance"`
}

var idPattern = regexp.MustCompile(`^[a-z0-9]+([._:-][a-z0-9]+)*$`)
var versionPattern = regexp.MustCompile(`^[0-9]+\.[0-9]+([.-][a-z0-9]+)*$`)
var localePattern = regexp.MustCompile(`^[a-z]{2,3}(-[A-Z]{2})?$`)
var timezonePattern = regexp.MustCompile(`^(UTC|[A-Za-z]+([_-][A-Za-z]+)*(\/[A-Za-z]+([_+-][A-Za-z]+)*)+)$`)

func BuiltIn(rules []rule.Definition, policies []policy.Profile) (Source, error) {
	instance, locale, timezone := "local", "en", "UTC"
	checks, targets, schedules := []Check{}, []Target{}, []Schedule{}
	retries := []RetryPolicy{{ID: "canonical.default", MaxAttempts: 1, InitialDelayNS: 0, MaxDelayNS: 0}}
	secrets, extensions := []SecretReference{}, []Extension{}
	retention, timeout, concurrency := DefaultRetention, DefaultTimeoutNS, DefaultConcurrency
	report := ReportPolicy{Enabled: true, RetentionCount: DefaultRetention}
	source := Source{
		SchemaName: SourceSchema, SchemaVersion: SchemaVersion, ModelVersion: ModelVersion,
		ID: "qwsg.builtin.default", SourceVersion: "1.0", Kind: BuiltInDefault,
		Patch: Patch{InstanceID: &instance, Locale: &locale, TimeZone: &timezone,
			Checks: &checks, Targets: &targets, RuleDefinitions: &rules, PolicyProfiles: &policies,
			SnapshotRetention: &retention, Schedules: &schedules, ExecutionTimeoutNS: &timeout,
			Concurrency: &concurrency, RetryPolicies: &retries, Report: &report,
			SecretReferences: &secrets, Extensions: &extensions},
	}
	return NormalizeSource(source)
}

func NormalizeSource(source Source) (Source, error) {
	source = cloneSource(source)
	claimed := source.Identity
	normalizePatch(&source.Patch)
	source.Identity = ""
	if err := validateSource(source, false); err != nil {
		return Source{}, err
	}
	source.Identity = stableID(source)
	if claimed != "" && claimed != source.Identity {
		return Source{}, fmt.Errorf("invalid configuration source identity")
	}
	return source, validateSource(source, true)
}

func Resolve(sources []Source) (Effective, error) {
	if len(sources) == 0 || len(sources) > MaxSources {
		return Effective{}, fmt.Errorf("configuration source count must be between 1 and %d", MaxSources)
	}
	normalized := make([]Source, len(sources))
	for i, source := range sources {
		var err error
		normalized[i], err = NormalizeSource(source)
		if err != nil {
			return Effective{}, fmt.Errorf("configuration source %q: %w", source.ID, err)
		}
	}
	sort.Slice(normalized, func(i, j int) bool {
		pi, pj := precedence(normalized[i].Kind), precedence(normalized[j].Kind)
		if pi != pj {
			return pi < pj
		}
		return normalized[i].ID < normalized[j].ID
	})
	seen := map[string]bool{}
	for _, source := range normalized {
		if seen[source.ID] {
			return Effective{}, fmt.Errorf("duplicate configuration source id %q", source.ID)
		}
		seen[source.ID] = true
	}
	values := Model{}
	provenance := map[string]FieldProvenance{}
	for _, source := range normalized {
		for _, field := range patchFields(source.Patch) {
			value := fieldValue(source.Patch, field)
			previous, exists := provenance[field]
			if exists && previous.Precedence == precedence(source.Kind) {
				if !bytes.Equal(canonical(fieldValueFromModel(values, field)), canonical(value)) {
					return Effective{}, fmt.Errorf("equal-precedence conflict for field %q between %q and %q", field, previous.SourceIDs[0], source.ID)
				}
				previous.SourceIDs = append(previous.SourceIDs, source.ID)
				previous.SourceVersions = append(previous.SourceVersions, source.SourceVersion)
				previous.Resolution = "equal_values_merged"
				provenance[field] = previous
				continue
			}
			setModelField(&values, field, value)
			provenance[field] = FieldProvenance{Field: field, SourceIDs: []string{source.ID}, SourceVersions: []string{source.SourceVersion}, Kind: source.Kind, Precedence: precedence(source.Kind), Resolution: resolution(exists)}
		}
	}
	if err := validateModel(values); err != nil {
		return Effective{}, err
	}
	required := allFields()
	for _, field := range required {
		if _, ok := provenance[field]; !ok {
			return Effective{}, fmt.Errorf("missing required configuration field %q", field)
		}
	}
	result := Effective{SchemaName: EffectiveSchema, SchemaVersion: SchemaVersion, ModelVersion: ModelVersion, SourceIDs: make([]string, len(normalized)), Values: values, Provenance: provenance}
	for i, source := range normalized {
		result.SourceIDs[i] = source.ID
	}
	sort.Strings(result.SourceIDs)
	result.ID = effectiveID(result)
	return result, ValidateEffective(result)
}

func ValidateEffective(result Effective) error {
	if result.SchemaName != EffectiveSchema || result.SchemaVersion != SchemaVersion || result.ModelVersion != ModelVersion || result.ID == "" || !sortedUniqueStrings(result.SourceIDs) || len(result.Provenance) != len(allFields()) {
		return fmt.Errorf("invalid effective configuration envelope")
	}
	if err := validateModel(result.Values); err != nil {
		return err
	}
	for _, field := range allFields() {
		item, ok := result.Provenance[field]
		if !ok || item.Field != field || len(item.SourceIDs) == 0 || len(item.SourceIDs) != len(item.SourceVersions) || !sortedUniqueStrings(item.SourceIDs) || precedence(item.Kind) != item.Precedence || (item.Resolution != "selected" && item.Resolution != "overridden" && item.Resolution != "equal_values_merged") {
			return fmt.Errorf("invalid provenance for field %q", field)
		}
	}
	if result.ID != effectiveID(result) {
		return fmt.Errorf("invalid effective configuration identity")
	}
	return nil
}

func MarshalSourceCanonical(source Source) ([]byte, error) {
	normalized, err := NormalizeSource(source)
	if err != nil {
		return nil, err
	}
	return json.Marshal(normalized)
}

func MarshalEffectiveCanonical(result Effective) ([]byte, error) {
	if err := ValidateEffective(result); err != nil {
		return nil, err
	}
	return json.Marshal(result)
}

func DecodeSource(data []byte) (Source, error) {
	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	var source Source
	if err := decoder.Decode(&source); err != nil {
		return Source{}, fmt.Errorf("invalid configuration source JSON: %w", err)
	}
	if err := ensureEOF(decoder); err != nil {
		return Source{}, err
	}
	return NormalizeSource(source)
}

func validateSource(source Source, requireIdentity bool) error {
	if source.SchemaName != SourceSchema || source.SchemaVersion != SchemaVersion || source.ModelVersion != ModelVersion || !idPattern.MatchString(source.ID) || !versionPattern.MatchString(source.SourceVersion) || precedence(source.Kind) < 0 || (requireIdentity && source.Identity != stableID(withoutIdentity(source))) || len(patchFields(source.Patch)) == 0 {
		return fmt.Errorf("invalid or unsupported configuration source contract")
	}
	return validatePatch(source.Patch)
}

func validatePatch(patch Patch) error {
	test := Model{}
	for _, field := range patchFields(patch) {
		setModelField(&test, field, fieldValue(patch, field))
	}
	return validatePresentFields(test, patchFields(patch))
}

func validateModel(model Model) error { return validatePresentFields(model, allFields()) }

func validatePresentFields(model Model, fields []string) error {
	present := map[string]bool{}
	for _, field := range fields {
		present[field] = true
	}
	if present["instance_id"] && !idPattern.MatchString(model.InstanceID) {
		return fmt.Errorf("invalid instance identity")
	}
	if present["locale"] && !localePattern.MatchString(model.Locale) {
		return fmt.Errorf("invalid locale")
	}
	if present["time_zone"] && !validTimeZone(model.TimeZone) {
		return fmt.Errorf("invalid time zone")
	}
	if present["snapshot_retention"] && (model.SnapshotRetention < 1 || model.SnapshotRetention > MaxSnapshotRetain) {
		return fmt.Errorf("snapshot retention out of bounds")
	}
	if present["execution_timeout_ns"] && (model.ExecutionTimeoutNS < 1 || model.ExecutionTimeoutNS > MaxDurationNS) {
		return fmt.Errorf("execution timeout out of bounds")
	}
	if present["concurrency"] && (model.Concurrency < 1 || model.Concurrency > MaxConcurrency) {
		return fmt.Errorf("concurrency out of bounds")
	}
	if present["checks"] {
		if err := validateChecks(model.Checks); err != nil {
			return err
		}
	}
	if present["targets"] {
		if err := validateTargets(model.Targets); err != nil {
			return err
		}
	}
	if present["rule_definitions"] {
		if err := validateRules(model.RuleDefinitions); err != nil {
			return err
		}
	}
	if present["policy_profiles"] {
		if _, err := policy.MarshalProfilesCanonical(model.PolicyProfiles); err != nil {
			return fmt.Errorf("invalid policy profiles: %w", err)
		}
	}
	if present["retry_policies"] {
		if err := validateRetries(model.RetryPolicies); err != nil {
			return err
		}
	}
	if present["schedules"] {
		if err := validateSchedules(model.Schedules, model); err != nil {
			return err
		}
	}
	if present["report"] && (model.Report.RetentionCount < 0 || model.Report.RetentionCount > MaxSnapshotRetain) {
		return fmt.Errorf("report retention out of bounds")
	}
	if present["secret_references"] {
		if err := validateSecrets(model.SecretReferences); err != nil {
			return err
		}
	}
	if present["extensions"] {
		if err := validateExtensions(model.Extensions); err != nil {
			return err
		}
	}
	if len(model.Checks)+len(model.Targets)+len(model.RuleDefinitions)+len(model.PolicyProfiles)+len(model.Schedules)+len(model.RetryPolicies)+len(model.SecretReferences)+len(model.Extensions) > MaxItems {
		return fmt.Errorf("configuration item limit exceeded")
	}
	if len(fields) == len(allFields()) {
		return validateReferences(model)
	}
	return nil
}

func validateChecks(values []Check) error {
	for i, value := range values {
		if !idPattern.MatchString(value.ID) || (i > 0 && values[i-1].ID >= value.ID) || !sortedUniqueStrings(value.TargetIDs) {
			return fmt.Errorf("invalid or unordered check")
		}
	}
	return nil
}
func validateTargets(values []Target) error {
	for i, value := range values {
		if !idPattern.MatchString(value.ID) || value.Kind == "" || len(value.Kind) > 128 || value.Metadata == nil || len(value.Metadata) > 16 || (i > 0 && values[i-1].ID >= value.ID) || !validMap(value.Metadata) {
			return fmt.Errorf("invalid or unordered target")
		}
	}
	return nil
}
func validateRules(values []rule.Definition) error {
	return rule.ValidateDefinitions(values)
}
func validateRetries(values []RetryPolicy) error {
	for i, value := range values {
		if !idPattern.MatchString(value.ID) || value.MaxAttempts < 1 || value.MaxAttempts > 16 || value.InitialDelayNS < 0 || value.MaxDelayNS < value.InitialDelayNS || value.MaxDelayNS > MaxDurationNS || (i > 0 && values[i-1].ID >= value.ID) {
			return fmt.Errorf("invalid or unordered retry policy")
		}
	}
	return nil
}
func validateSecrets(values []SecretReference) error {
	for i, value := range values {
		if !idPattern.MatchString(value.ID) || !idPattern.MatchString(value.Provider) || value.Reference == "" || len(value.Reference) > 256 || strings.ContainsAny(strings.ToLower(value.Reference), "\n\r\t") || (i > 0 && values[i-1].ID >= value.ID) {
			return fmt.Errorf("invalid or unordered secret reference")
		}
	}
	return nil
}
func validateExtensions(values []Extension) error {
	for i, value := range values {
		if !idPattern.MatchString(value.ID) || !versionPattern.MatchString(value.Version) || value.Required || value.Fields == nil || len(value.Fields) > 16 || !validMap(value.Fields) || (i > 0 && values[i-1].ID >= value.ID) {
			return fmt.Errorf("unsupported, invalid, or unordered extension")
		}
	}
	return nil
}

func validateSchedules(values []Schedule, model Model) error {
	for i, value := range values {
		if !idPattern.MatchString(value.ID) || value.ContractVersion != ScheduleVersion || !validTimeZone(value.TimeZone) || value.Priority < -1000 || value.Priority > 1000 || value.ExecutionTimeoutNS < 1 || value.ExecutionTimeoutNS > MaxDurationNS || !idPattern.MatchString(value.RetryPolicyID) || !sortedUniqueStrings(value.CheckIDs) || value.CommandProfile == "" || len(value.CommandProfile) > 128 || (i > 0 && values[i-1].ID >= value.ID) {
			return fmt.Errorf("invalid or unordered schedule")
		}
		if value.MisfirePolicy != MisfireSkip && value.MisfirePolicy != MisfireRunOnce && value.MisfirePolicy != MisfireIndeterminate {
			return fmt.Errorf("unsupported misfire policy")
		}
		if value.OverlapPolicy != OverlapForbid && value.OverlapPolicy != OverlapAllow {
			return fmt.Errorf("unsupported overlap policy")
		}
		if value.DSTPolicy != DSTFirstOccurrence && value.DSTPolicy != DSTSecondOccurrence && value.DSTPolicy != DSTSkipNonexistent {
			return fmt.Errorf("unsupported daylight-saving policy")
		}
		switch value.Trigger {
		case IntervalTrigger:
			if value.IntervalNS < 1 || value.IntervalNS > MaxDurationNS || !emptyCalendar(value.Calendar) {
				return fmt.Errorf("invalid interval schedule")
			}
		case CalendarTrigger:
			if value.IntervalNS != 0 || !validCalendar(value.Calendar) {
				return fmt.Errorf("invalid calendar schedule")
			}
		default:
			return fmt.Errorf("unsupported schedule trigger")
		}
	}
	return nil
}

func validateReferences(model Model) error {
	targets, checks, retries, schedules := map[string]bool{}, map[string]bool{}, map[string]bool{}, map[string]bool{}
	for _, item := range model.Targets {
		targets[item.ID] = true
	}
	for _, item := range model.Checks {
		checks[item.ID] = true
		for _, id := range item.TargetIDs {
			if !targets[id] {
				return fmt.Errorf("check references unknown target %q", id)
			}
		}
	}
	for _, item := range model.RetryPolicies {
		retries[item.ID] = true
	}
	for _, item := range model.Schedules {
		schedules[item.ID] = true
		if !retries[item.RetryPolicyID] {
			return fmt.Errorf("schedule references unknown retry policy %q", item.RetryPolicyID)
		}
		for _, id := range item.CheckIDs {
			if !checks[id] {
				return fmt.Errorf("schedule references unknown check %q", id)
			}
		}
	}
	if model.Report.ScheduleID != "" && !schedules[model.Report.ScheduleID] {
		return fmt.Errorf("report references unknown schedule %q", model.Report.ScheduleID)
	}
	return nil
}

func validCalendar(value Calendar) bool {
	return validInts(value.Minutes, 0, 59, true) && validInts(value.Hours, 0, 23, true) && validInts(value.MonthDays, 1, 31, false) && validInts(value.Months, 1, 12, false) && validInts(value.Weekdays, 0, 6, false)
}
func emptyCalendar(value Calendar) bool {
	return len(value.Minutes) == 0 && len(value.Hours) == 0 && len(value.MonthDays) == 0 && len(value.Months) == 0 && len(value.Weekdays) == 0
}
func validInts(values []int, min, max int, required bool) bool {
	if required && len(values) == 0 {
		return false
	}
	for i, v := range values {
		if v < min || v > max || (i > 0 && values[i-1] >= v) {
			return false
		}
	}
	return true
}
func validTimeZone(value string) bool { return len(value) <= 128 && timezonePattern.MatchString(value) }
func validMap(value map[string]string) bool {
	for key, item := range value {
		if !idPattern.MatchString(key) || len(item) > MaxStringLength {
			return false
		}
	}
	return true
}

func normalizePatch(patch *Patch) {
	if patch.Checks != nil {
		normalizeChecks(*patch.Checks)
	}
	if patch.Targets != nil {
		normalizeTargets(*patch.Targets)
	}
	if patch.RuleDefinitions != nil {
		normalizeRules(*patch.RuleDefinitions)
		sort.Slice(*patch.RuleDefinitions, func(i, j int) bool { return (*patch.RuleDefinitions)[i].ID < (*patch.RuleDefinitions)[j].ID })
	}
	if patch.PolicyProfiles != nil {
		sort.Slice(*patch.PolicyProfiles, func(i, j int) bool { return (*patch.PolicyProfiles)[i].ID < (*patch.PolicyProfiles)[j].ID })
	}
	if patch.Schedules != nil {
		normalizeSchedules(*patch.Schedules)
	}
	if patch.RetryPolicies != nil {
		sort.Slice(*patch.RetryPolicies, func(i, j int) bool { return (*patch.RetryPolicies)[i].ID < (*patch.RetryPolicies)[j].ID })
	}
	if patch.SecretReferences != nil {
		sort.Slice(*patch.SecretReferences, func(i, j int) bool { return (*patch.SecretReferences)[i].ID < (*patch.SecretReferences)[j].ID })
	}
	if patch.Extensions != nil {
		sort.Slice(*patch.Extensions, func(i, j int) bool { return (*patch.Extensions)[i].ID < (*patch.Extensions)[j].ID })
	}
}
func normalizeRules(values []rule.Definition) {
	for i := range values {
		if values[i].Scope.HealthIDs == nil {
			values[i].Scope.HealthIDs = []string{}
		}
		if values[i].Scope.Categories == nil {
			values[i].Scope.Categories = []drift.Category{}
		}
		if values[i].InputRequirements == nil {
			values[i].InputRequirements = []rule.Field{}
		}
		if values[i].Metadata == nil {
			values[i].Metadata = map[string]string{}
		}
		sort.Strings(values[i].Scope.HealthIDs)
		sort.Slice(values[i].Scope.Categories, func(a, b int) bool { return values[i].Scope.Categories[a] < values[i].Scope.Categories[b] })
		sort.Slice(values[i].InputRequirements, func(a, b int) bool { return values[i].InputRequirements[a] < values[i].InputRequirements[b] })
		normalizeCondition(&values[i].Condition)
	}
}
func normalizeCondition(value *rule.Condition) {
	if value.Values == nil {
		value.Values = []rule.Value{}
	}
	if value.Children == nil {
		value.Children = []rule.Condition{}
	}
	for i := range value.Children {
		normalizeCondition(&value.Children[i])
	}
}
func normalizeChecks(values []Check) {
	for i := range values {
		sort.Strings(values[i].TargetIDs)
		if values[i].TargetIDs == nil {
			values[i].TargetIDs = []string{}
		}
	}
	sort.Slice(values, func(i, j int) bool { return values[i].ID < values[j].ID })
}
func normalizeTargets(values []Target) {
	for i := range values {
		if values[i].Metadata == nil {
			values[i].Metadata = map[string]string{}
		}
	}
	sort.Slice(values, func(i, j int) bool { return values[i].ID < values[j].ID })
}
func normalizeSchedules(values []Schedule) {
	for i := range values {
		sort.Ints(values[i].Calendar.Minutes)
		sort.Ints(values[i].Calendar.Hours)
		sort.Ints(values[i].Calendar.MonthDays)
		sort.Ints(values[i].Calendar.Months)
		sort.Ints(values[i].Calendar.Weekdays)
		sort.Strings(values[i].CheckIDs)
		if values[i].CheckIDs == nil {
			values[i].CheckIDs = []string{}
		}
		normalizeCalendar(&values[i].Calendar)
	}
	sort.Slice(values, func(i, j int) bool { return values[i].ID < values[j].ID })
}
func normalizeCalendar(value *Calendar) {
	if value.Minutes == nil {
		value.Minutes = []int{}
	}
	if value.Hours == nil {
		value.Hours = []int{}
	}
	if value.MonthDays == nil {
		value.MonthDays = []int{}
	}
	if value.Months == nil {
		value.Months = []int{}
	}
	if value.Weekdays == nil {
		value.Weekdays = []int{}
	}
}

func allFields() []string {
	return []string{"checks", "concurrency", "execution_timeout_ns", "extensions", "instance_id", "locale", "policy_profiles", "report", "retry_policies", "rule_definitions", "schedules", "secret_references", "snapshot_retention", "targets", "time_zone"}
}
func patchFields(p Patch) []string {
	result := []string{}
	if p.Checks != nil {
		result = append(result, "checks")
	}
	if p.Concurrency != nil {
		result = append(result, "concurrency")
	}
	if p.ExecutionTimeoutNS != nil {
		result = append(result, "execution_timeout_ns")
	}
	if p.Extensions != nil {
		result = append(result, "extensions")
	}
	if p.InstanceID != nil {
		result = append(result, "instance_id")
	}
	if p.Locale != nil {
		result = append(result, "locale")
	}
	if p.PolicyProfiles != nil {
		result = append(result, "policy_profiles")
	}
	if p.Report != nil {
		result = append(result, "report")
	}
	if p.RetryPolicies != nil {
		result = append(result, "retry_policies")
	}
	if p.RuleDefinitions != nil {
		result = append(result, "rule_definitions")
	}
	if p.Schedules != nil {
		result = append(result, "schedules")
	}
	if p.SecretReferences != nil {
		result = append(result, "secret_references")
	}
	if p.SnapshotRetention != nil {
		result = append(result, "snapshot_retention")
	}
	if p.Targets != nil {
		result = append(result, "targets")
	}
	if p.TimeZone != nil {
		result = append(result, "time_zone")
	}
	return result
}
func fieldValue(p Patch, f string) any {
	switch f {
	case "checks":
		return *p.Checks
	case "concurrency":
		return *p.Concurrency
	case "execution_timeout_ns":
		return *p.ExecutionTimeoutNS
	case "extensions":
		return *p.Extensions
	case "instance_id":
		return *p.InstanceID
	case "locale":
		return *p.Locale
	case "policy_profiles":
		return *p.PolicyProfiles
	case "report":
		return *p.Report
	case "retry_policies":
		return *p.RetryPolicies
	case "rule_definitions":
		return *p.RuleDefinitions
	case "schedules":
		return *p.Schedules
	case "secret_references":
		return *p.SecretReferences
	case "snapshot_retention":
		return *p.SnapshotRetention
	case "targets":
		return *p.Targets
	case "time_zone":
		return *p.TimeZone
	}
	panic("unknown field")
}
func fieldValueFromModel(m Model, f string) any {
	switch f {
	case "checks":
		return m.Checks
	case "concurrency":
		return m.Concurrency
	case "execution_timeout_ns":
		return m.ExecutionTimeoutNS
	case "extensions":
		return m.Extensions
	case "instance_id":
		return m.InstanceID
	case "locale":
		return m.Locale
	case "policy_profiles":
		return m.PolicyProfiles
	case "report":
		return m.Report
	case "retry_policies":
		return m.RetryPolicies
	case "rule_definitions":
		return m.RuleDefinitions
	case "schedules":
		return m.Schedules
	case "secret_references":
		return m.SecretReferences
	case "snapshot_retention":
		return m.SnapshotRetention
	case "targets":
		return m.Targets
	case "time_zone":
		return m.TimeZone
	}
	panic("unknown field")
}
func setModelField(m *Model, f string, v any) {
	switch f {
	case "checks":
		m.Checks = v.([]Check)
	case "concurrency":
		m.Concurrency = v.(int)
	case "execution_timeout_ns":
		m.ExecutionTimeoutNS = v.(int64)
	case "extensions":
		m.Extensions = v.([]Extension)
	case "instance_id":
		m.InstanceID = v.(string)
	case "locale":
		m.Locale = v.(string)
	case "policy_profiles":
		m.PolicyProfiles = v.([]policy.Profile)
	case "report":
		m.Report = v.(ReportPolicy)
	case "retry_policies":
		m.RetryPolicies = v.([]RetryPolicy)
	case "rule_definitions":
		m.RuleDefinitions = v.([]rule.Definition)
	case "schedules":
		m.Schedules = v.([]Schedule)
	case "secret_references":
		m.SecretReferences = v.([]SecretReference)
	case "snapshot_retention":
		m.SnapshotRetention = v.(int)
	case "targets":
		m.Targets = v.([]Target)
	case "time_zone":
		m.TimeZone = v.(string)
	default:
		panic("unknown field")
	}
}

func precedence(kind SourceKind) int {
	switch kind {
	case BuiltInDefault:
		return 0
	case PrimaryLocal:
		return 100
	case ActivatedOverride:
		return 200
	case TemporaryOverride:
		return 300
	}
	return -1
}
func resolution(replaced bool) string {
	if replaced {
		return "overridden"
	}
	return "selected"
}
func withoutIdentity(source Source) Source { source.Identity = ""; return source }
func stableID(value any) string {
	sum := sha256.Sum256(canonical(value))
	return "sha256:" + hex.EncodeToString(sum[:])
}
func effectiveID(value Effective) string { value.ID = ""; return stableID(value) }
func canonical(value any) []byte {
	data, err := json.Marshal(value)
	if err != nil {
		panic(err)
	}
	return data
}
func cloneSource(source Source) Source {
	data, err := json.Marshal(source)
	if err != nil {
		panic(err)
	}
	var clone Source
	if err := json.Unmarshal(data, &clone); err != nil {
		panic(err)
	}
	return clone
}
func sortedUniqueStrings(values []string) bool {
	for i, v := range values {
		if v == "" || len(v) > MaxStringLength || (i > 0 && values[i-1] >= v) {
			return false
		}
	}
	return values != nil
}
func ensureEOF(decoder *json.Decoder) error {
	var extra any
	if err := decoder.Decode(&extra); err != io.EOF {
		if err == nil {
			return fmt.Errorf("multiple JSON values are not allowed")
		}
		return err
	}
	return nil
}
