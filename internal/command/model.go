// Package command defines the presentation-independent Canonical Command
// Architecture. It contains no terminal, HTTP, dashboard, collection, analysis,
// persistence, or host-operation code.
package command

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"sort"
	"strings"
)

const (
	SchemaName      = "qwsg.command"
	SchemaVersion   = "1.0"
	ExecutionSchema = "qwsg.command-execution"
	MaxValues       = 256
	MaxValueLength  = 4096
)

type Stage string

const (
	Inventory Stage = "inventory"
	Snapshot  Stage = "snapshot"
	Compare   Stage = "compare"
	Drift     Stage = "drift"
	Health    Stage = "health"
	Rule      Stage = "rule"
	Report    Stage = "report"
)

var CanonicalStages = []Stage{Inventory, Snapshot, Compare, Drift, Health, Rule, Report}

type OutputFormat string

const (
	JSON  OutputFormat = "json"
	Human OutputFormat = "human"
)

type Presentation string

const (
	Structured Presentation = "structured"
	Terminal   Presentation = "terminal"
)

type Selection struct {
	Source       string `json:"source"`
	Store        string `json:"store,omitempty"`
	FromSnapshot string `json:"from_snapshot,omitempty"`
	ToSnapshot   string `json:"to_snapshot,omitempty"`
}

type Parameters struct {
	Filters      []string     `json:"filters"`
	GroupBy      []string     `json:"group_by"`
	SortBy       []string     `json:"sort_by"`
	Output       OutputFormat `json:"output"`
	Presentation Presentation `json:"presentation"`
}

// Definition is Command Definition 1.0. Interfaces create or parse this value;
// they do not decide how its pipeline behaves.
type Definition struct {
	SchemaName      string     `json:"schema_name"`
	SchemaVersion   string     `json:"schema_version"`
	ID              string     `json:"id"`
	Name            string     `json:"name"`
	Profile         string     `json:"profile,omitempty"`
	Selection       Selection  `json:"selection"`
	Pipeline        []Stage    `json:"pipeline"`
	IncludedEngines []Stage    `json:"included_engines"`
	ExcludedEngines []Stage    `json:"excluded_engines"`
	Parameters      Parameters `json:"parameters"`
}

type Profile struct {
	Name         string       `json:"name"`
	Target       Stage        `json:"target"`
	Source       string       `json:"source"`
	Output       OutputFormat `json:"output"`
	Presentation Presentation `json:"presentation"`
}

type Plan struct {
	SchemaName    string     `json:"schema_name"`
	SchemaVersion string     `json:"schema_version"`
	ID            string     `json:"id"`
	CommandID     string     `json:"command_id"`
	Stages        []Stage    `json:"stages"`
	Selection     Selection  `json:"selection"`
	Parameters    Parameters `json:"parameters"`
}

type StageResult struct {
	Stage        Stage  `json:"stage"`
	ContractName string `json:"contract_name"`
	Version      string `json:"version"`
	RecordCount  int    `json:"record_count"`
	Complete     bool   `json:"complete"`
	Value        any    `json:"value,omitempty"`
}

type ViewRow struct {
	Stage       Stage  `json:"stage"`
	Contract    string `json:"contract"`
	Version     string `json:"version"`
	RecordCount int    `json:"record_count"`
	Complete    bool   `json:"complete"`
}

type ViewGroup struct {
	Key   string    `json:"key"`
	Value string    `json:"value"`
	Rows  []ViewRow `json:"rows"`
}

type View struct {
	Rows   []ViewRow   `json:"rows"`
	Groups []ViewGroup `json:"groups"`
}

// Execution is the canonical presentation-neutral command result. Presentation
// adapters consume this contract and may not execute stages or reinterpret it.
type Execution struct {
	SchemaName    string        `json:"schema_name"`
	SchemaVersion string        `json:"schema_version"`
	ID            string        `json:"id"`
	CommandID     string        `json:"command_id"`
	PlanID        string        `json:"plan_id"`
	Stages        []StageResult `json:"stages"`
	View          View          `json:"view"`
	Diagnostics   []string      `json:"diagnostics"`
	Complete      bool          `json:"complete"`
}

var profiles = map[string]Profile{
	"status":  {Name: "status", Target: Inventory, Source: "live", Output: Human, Presentation: Terminal},
	"check":   {Name: "check", Target: Snapshot, Source: "live", Output: Human, Presentation: Terminal},
	"changes": {Name: "changes", Target: Compare, Source: "store", Output: Human, Presentation: Terminal},
	"health":  {Name: "health", Target: Health, Source: "store", Output: Human, Presentation: Terminal},
	"report":  {Name: "report", Target: Report, Source: "store", Output: Human, Presentation: Terminal},
}

func Profiles() []Profile {
	result := make([]Profile, 0, len(profiles))
	for _, profile := range profiles {
		result = append(result, profile)
	}
	sort.Slice(result, func(i, j int) bool { return result[i].Name < result[j].Name })
	return result
}

func ResolveProfile(name string, selection Selection) (Definition, error) {
	profile, ok := profiles[name]
	if !ok {
		return Definition{}, fmt.Errorf("unsupported command profile %q", name)
	}
	if selection.Source == "" {
		selection.Source = profile.Source
	}
	definition := Definition{
		SchemaName: SchemaName, SchemaVersion: SchemaVersion,
		Name: name, Profile: name, Selection: selection,
		Pipeline:        []Stage{profile.Target},
		IncludedEngines: []Stage{}, ExcludedEngines: []Stage{},
		Parameters: Parameters{
			Filters: []string{}, GroupBy: []string{}, SortBy: []string{},
			Output: profile.Output, Presentation: profile.Presentation,
		},
	}
	return Normalize(definition)
}

func Normalize(definition Definition) (Definition, error) {
	claimedID := definition.ID
	if definition.SchemaName == "" {
		definition.SchemaName = SchemaName
	}
	if definition.SchemaVersion == "" {
		definition.SchemaVersion = SchemaVersion
	}
	if definition.Parameters.Filters == nil {
		definition.Parameters.Filters = []string{}
	}
	if definition.Parameters.GroupBy == nil {
		definition.Parameters.GroupBy = []string{}
	}
	if definition.Parameters.SortBy == nil {
		definition.Parameters.SortBy = []string{}
	}
	if definition.IncludedEngines == nil {
		definition.IncludedEngines = []Stage{}
	}
	if definition.ExcludedEngines == nil {
		definition.ExcludedEngines = []Stage{}
	}
	definition.Pipeline = uniqueStages(definition.Pipeline)
	definition.IncludedEngines = uniqueStages(definition.IncludedEngines)
	definition.ExcludedEngines = uniqueStages(definition.ExcludedEngines)
	sort.Strings(definition.Parameters.Filters)
	sort.Strings(definition.Parameters.GroupBy)
	sort.Strings(definition.Parameters.SortBy)
	definition.ID = ""
	if err := ValidateDefinition(definition); err != nil {
		return Definition{}, err
	}
	definition.ID = stableID("definition", canonical(definition))
	if claimedID != "" && claimedID != definition.ID {
		return Definition{}, fmt.Errorf("invalid command identity")
	}
	return definition, ValidateDefinition(definition)
}

func ValidateDefinition(definition Definition) error {
	if definition.SchemaName != SchemaName || definition.SchemaVersion != SchemaVersion {
		return fmt.Errorf("unsupported command contract")
	}
	if definition.Name == "" || len(definition.Name) > 128 || definition.Selection.Source == "" {
		return fmt.Errorf("incomplete command definition")
	}
	if len(definition.Pipeline) == 0 {
		return fmt.Errorf("pipeline selection is required")
	}
	if definition.Parameters.Output != JSON && definition.Parameters.Output != Human {
		return fmt.Errorf("unsupported output format")
	}
	if definition.Parameters.Presentation != Structured && definition.Parameters.Presentation != Terminal {
		return fmt.Errorf("unsupported presentation")
	}
	if definition.Parameters.Presentation == Structured && definition.Parameters.Output != JSON {
		return fmt.Errorf("structured presentation requires json output")
	}
	if err := validateValues(definition.Parameters.Filters, definition.Parameters.GroupBy, definition.Parameters.SortBy); err != nil {
		return err
	}
	if _, err := BuildView([]StageResult{}, definition.Parameters); err != nil {
		return err
	}
	excluded := map[Stage]bool{}
	for _, stage := range definition.ExcludedEngines {
		if !validStage(stage) {
			return fmt.Errorf("unsupported excluded engine %q", stage)
		}
		excluded[stage] = true
	}
	for _, group := range [][]Stage{definition.Pipeline, definition.IncludedEngines} {
		for _, stage := range group {
			if !validStage(stage) || excluded[stage] {
				return fmt.Errorf("invalid or contradictory engine selection")
			}
			for _, dependency := range dependencies(stage, definition.Selection.Source) {
				if excluded[dependency] {
					return fmt.Errorf("excluded engine %q is required by the selected pipeline", dependency)
				}
			}
		}
	}
	expected := definition
	expected.ID = ""
	if definition.ID != "" && definition.ID != stableID("definition", canonical(expected)) {
		return fmt.Errorf("invalid command identity")
	}
	return nil
}

func PlanDefinition(definition Definition) (Plan, error) {
	if err := ValidateDefinition(definition); err != nil {
		return Plan{}, err
	}
	if definition.ID == "" {
		return Plan{}, fmt.Errorf("normalized command identity is required")
	}
	selected := map[Stage]bool{}
	for _, stage := range append(append([]Stage{}, definition.Pipeline...), definition.IncludedEngines...) {
		for _, dependency := range dependencies(stage, definition.Selection.Source) {
			selected[dependency] = true
		}
	}
	for _, stage := range definition.ExcludedEngines {
		if selected[stage] {
			return Plan{}, fmt.Errorf("excluded engine %q is required by the selected pipeline", stage)
		}
	}
	stages := make([]Stage, 0, len(selected))
	for _, stage := range CanonicalStages {
		if selected[stage] {
			stages = append(stages, stage)
		}
	}
	plan := Plan{
		SchemaName: SchemaName + ".plan", SchemaVersion: SchemaVersion,
		CommandID: definition.ID, Stages: stages,
		Selection: definition.Selection, Parameters: definition.Parameters,
	}
	plan.ID = stableID("plan", canonical(plan))
	return plan, nil
}

func dependencies(target Stage, source string) []Stage {
	rank := stageRank(target)
	start := 0
	if source == "store" && rank >= stageRank(Compare) {
		start = stageRank(Compare)
	}
	return append([]Stage(nil), CanonicalStages[start:rank+1]...)
}

func uniqueStages(values []Stage) []Stage {
	seen := map[Stage]bool{}
	result := make([]Stage, 0, len(values))
	for _, stage := range CanonicalStages {
		for _, value := range values {
			if value == stage && !seen[value] {
				seen[value], result = true, append(result, value)
			}
		}
	}
	return result
}

func validStage(stage Stage) bool { return stageRank(stage) >= 0 }

func stageRank(stage Stage) int {
	for index, candidate := range CanonicalStages {
		if candidate == stage {
			return index
		}
	}
	return -1
}

func validateValues(groups ...[]string) error {
	total := 0
	for _, group := range groups {
		total += len(group)
		for _, value := range group {
			if value == "" || len(value) > MaxValueLength || strings.ContainsAny(value, "\x00\r\n") {
				return fmt.Errorf("invalid command parameter")
			}
		}
	}
	if total > MaxValues {
		return fmt.Errorf("command parameter limit exceeded")
	}
	return nil
}

func canonical(value any) string {
	document, _ := json.Marshal(value)
	return string(document)
}

func stableID(domain, value string) string {
	sum := sha256.Sum256([]byte(SchemaName + "/" + SchemaVersion + "/" + domain + "\x00" + value))
	return hex.EncodeToString(sum[:])
}
