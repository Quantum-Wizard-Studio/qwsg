package command

import (
	"encoding/json"
	"reflect"
	"strings"
	"testing"
)

func TestProfilesResolveDeterministically(t *testing.T) {
	expected := map[string][]Stage{
		"status":  {Inventory},
		"check":   {Inventory, Snapshot},
		"observe": {Inventory, Snapshot, Compare, Drift, Health, Rule, Policy, Report},
		"changes": {Compare},
		"health":  {Compare, Drift, Health},
		"report":  {Compare, Drift, Health, Rule, Policy, Report},
	}
	for _, profile := range Profiles() {
		first, err := ResolveProfile(profile.Name, Selection{Store: "/tmp/store"})
		if err != nil {
			t.Fatal(err)
		}
		second, err := ResolveProfile(profile.Name, Selection{Store: "/tmp/store"})
		if err != nil || first.ID != second.ID {
			t.Fatalf("%s is not deterministic: %v", profile.Name, err)
		}
		plan, err := PlanDefinition(first)
		if err != nil || !reflect.DeepEqual(plan.Stages, expected[profile.Name]) {
			t.Fatalf("%s stages=%v error=%v", profile.Name, plan.Stages, err)
		}
	}
}

func TestAdvancedGrammarAndSimpleProfileAreEquivalent(t *testing.T) {
	simple, err := ResolveProfile("health", Selection{Source: "store", Store: "/tmp/store"})
	if err != nil {
		t.Fatal(err)
	}
	advanced, err := Parse([]string{
		"--source", "store", "--store", "/tmp/store",
		"--pipeline", "health", "--output", "human", "--presentation", "terminal",
	})
	if err != nil {
		t.Fatal(err)
	}
	simplePlan, _ := PlanDefinition(simple)
	advancedPlan, _ := PlanDefinition(advanced)
	if !reflect.DeepEqual(simplePlan.Stages, advancedPlan.Stages) ||
		simplePlan.Selection != advancedPlan.Selection ||
		!reflect.DeepEqual(simplePlan.Parameters, advancedPlan.Parameters) {
		t.Fatalf("simple and advanced plans differ:\n%#v\n%#v", simplePlan, advancedPlan)
	}
}

func TestAdvancedParameterNormalizationAndCanonicalJSON(t *testing.T) {
	args := []string{
		"--source", "store", "--store", "/tmp/store", "--pipeline", "report",
		"--filter", "stage=report", "--filter", "complete=true",
		"--group", "contract", "--sort", "record_count",
		"--output", "json", "--presentation", "structured",
	}
	first, err := Parse(args)
	if err != nil {
		t.Fatal(err)
	}
	second, err := Parse(args)
	if err != nil || first.ID != second.ID {
		t.Fatalf("unstable parse: %v", err)
	}
	a, _ := json.Marshal(first)
	b, _ := json.Marshal(second)
	if string(a) != string(b) || !strings.Contains(string(a), `"schema_version":"1.0"`) {
		t.Fatalf("non-canonical definition: %s", a)
	}
}

func TestInvalidContradictoryAndBoundedDefinitionsFail(t *testing.T) {
	tests := [][]string{
		{"--source", "store", "--pipeline", "health"},
		{"--source", "store", "--store", "/tmp/store", "--pipeline", "unknown"},
		{"--source", "store", "--store", "/tmp/store", "--pipeline", "health", "--exclude", "compare"},
		{"--source", "store", "--store", "/tmp/store", "--pipeline", "health", "--from", "one"},
		{"--source", "store", "--store", "/tmp/store", "--pipeline", "health", "--output", "yaml"},
		{"--source", "store", "--store", "/tmp/store", "--pipeline", "health", "--presentation", "dashboard"},
		{"--source", "store", "--store", "/tmp/store", "--pipeline", "health", "--unknown", "x"},
	}
	for _, args := range tests {
		if _, err := Parse(args); err == nil {
			t.Fatalf("accepted invalid arguments: %v", args)
		}
	}
	tooLong := strings.Repeat("x", MaxValueLength+1)
	if _, err := Parse([]string{
		"--source", "store", "--store", "/tmp/store", "--pipeline", "health",
		"--filter", tooLong,
	}); err == nil {
		t.Fatal("accepted oversized parameter")
	}
}

func TestDefinitionRejectsTamperingAndUnsupportedVersion(t *testing.T) {
	definition, _ := ResolveProfile("status", Selection{})
	definition.ID = "tampered"
	if err := ValidateDefinition(definition); err == nil {
		t.Fatal("accepted tampered identity")
	}
	definition.ID = ""
	definition.SchemaVersion = "2.0"
	if err := ValidateDefinition(definition); err == nil {
		t.Fatal("accepted unsupported version")
	}
}

func TestViewFilteringGroupingAndSortingAreDeterministic(t *testing.T) {
	stages := []StageResult{
		{Stage: Inventory, ContractName: "qwsg.inventory", Version: "1.0", RecordCount: 9, Complete: true},
		{Stage: Compare, ContractName: "qwsg.comparison", Version: "1.0", RecordCount: 2, Complete: true},
		{Stage: Health, ContractName: "qwsg.health", Version: "1.0", RecordCount: 2, Complete: false},
	}
	parameters := Parameters{
		Filters: []string{"complete=true"}, GroupBy: []string{"record_count"},
		SortBy: []string{"record_count"}, Output: JSON, Presentation: Structured,
	}
	first, err := BuildView(stages, parameters)
	if err != nil {
		t.Fatal(err)
	}
	second, err := BuildView(stages, parameters)
	if err != nil || !reflect.DeepEqual(first, second) {
		t.Fatalf("view is not deterministic: %v", err)
	}
	if len(first.Rows) != 2 || first.Rows[0].Stage != Compare ||
		len(first.Groups) != 2 || first.Groups[0].Key != "record_count" {
		t.Fatalf("unexpected projected view: %#v", first)
	}
}
