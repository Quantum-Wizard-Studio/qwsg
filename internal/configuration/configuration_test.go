package configuration

import (
	"bytes"
	"encoding/json"
	"strings"
	"testing"

	"quantumwizard.hu/qwsg/internal/drift"
	"quantumwizard.hu/qwsg/internal/policy"
	"quantumwizard.hu/qwsg/internal/rule"
)

func TestResolveDeterministicPrecedenceAndProvenance(t *testing.T) {
	base := mustBuiltIn(t)
	locale, retention := "hu-HU", 42
	primary := mustSource(t, Source{SchemaName: SourceSchema, SchemaVersion: SchemaVersion, ModelVersion: ModelVersion, ID: "local.primary", SourceVersion: "1.0", Kind: PrimaryLocal, Patch: Patch{Locale: &locale}})
	temporary := mustSource(t, Source{SchemaName: SourceSchema, SchemaVersion: SchemaVersion, ModelVersion: ModelVersion, ID: "command.override", SourceVersion: "2.0", Kind: TemporaryOverride, Patch: Patch{SnapshotRetention: &retention}})

	first, err := Resolve([]Source{temporary, base, primary})
	if err != nil {
		t.Fatal(err)
	}
	second, err := Resolve([]Source{primary, temporary, base})
	if err != nil {
		t.Fatal(err)
	}
	one, _ := MarshalEffectiveCanonical(first)
	two, _ := MarshalEffectiveCanonical(second)
	if !bytes.Equal(one, two) || first.ID != second.ID {
		t.Fatal("source enumeration changed canonical configuration")
	}
	if first.Values.Locale != "hu-HU" || first.Values.SnapshotRetention != 42 {
		t.Fatalf("precedence failed: %#v", first.Values)
	}
	if first.Provenance["locale"].SourceIDs[0] != "local.primary" || first.Provenance["snapshot_retention"].Resolution != "overridden" {
		t.Fatalf("provenance failed: %#v", first.Provenance)
	}
}

func TestEqualPrecedenceConflictAndEqualMerge(t *testing.T) {
	base := mustBuiltIn(t)
	a, b := "hu", "de"
	left := mustSource(t, Source{SchemaName: SourceSchema, SchemaVersion: SchemaVersion, ModelVersion: ModelVersion, ID: "local.a", SourceVersion: "1.0", Kind: PrimaryLocal, Patch: Patch{Locale: &a}})
	right := mustSource(t, Source{SchemaName: SourceSchema, SchemaVersion: SchemaVersion, ModelVersion: ModelVersion, ID: "local.b", SourceVersion: "1.0", Kind: PrimaryLocal, Patch: Patch{Locale: &b}})
	if _, err := Resolve([]Source{base, left, right}); err == nil || !strings.Contains(err.Error(), "equal-precedence conflict") {
		t.Fatalf("expected explicit conflict, got %v", err)
	}
	right.Patch.Locale = &a
	right = mustSource(t, withoutClaim(right))
	result, err := Resolve([]Source{right, base, left})
	if err != nil {
		t.Fatal(err)
	}
	if result.Provenance["locale"].Resolution != "equal_values_merged" || len(result.Provenance["locale"].SourceIDs) != 2 {
		t.Fatalf("equal merge not explained: %#v", result.Provenance["locale"])
	}
}

func TestScheduleDefinitionAndReferences(t *testing.T) {
	base := mustBuiltIn(t)
	checks := []Check{{ID: "check.health", Enabled: true, TargetIDs: []string{"target.local"}}}
	targets := []Target{{ID: "target.local", Kind: "host", Metadata: map[string]string{}}}
	retries := []RetryPolicy{{ID: "retry.normal", MaxAttempts: 3, InitialDelayNS: 1_000_000_000, MaxDelayNS: 4_000_000_000}}
	schedules := []Schedule{{ID: "schedule.health", ContractVersion: ScheduleVersion, Enabled: true, TimeZone: "Europe/Budapest", Trigger: CalendarTrigger, Calendar: Calendar{Minutes: []int{0, 30}, Hours: []int{1, 12}}, DSTPolicy: DSTFirstOccurrence, Priority: 100, MisfirePolicy: MisfireRunOnce, OverlapPolicy: OverlapForbid, ExecutionTimeoutNS: 60_000_000_000, RetryPolicyID: "retry.normal", CheckIDs: []string{"check.health"}, CommandProfile: "report"}}
	source := mustSource(t, Source{SchemaName: SourceSchema, SchemaVersion: SchemaVersion, ModelVersion: ModelVersion, ID: "local.scheduler", SourceVersion: "1.0", Kind: PrimaryLocal, Patch: Patch{Checks: &checks, Targets: &targets, RetryPolicies: &retries, Schedules: &schedules}})
	result, err := Resolve([]Source{source, base})
	if err != nil {
		t.Fatal(err)
	}
	if len(result.Values.Schedules) != 1 || result.Values.Schedules[0].Trigger != CalendarTrigger {
		t.Fatalf("schedule missing: %#v", result.Values.Schedules)
	}

	bad := schedules
	bad[0].RetryPolicyID = "retry.missing"
	source = mustSource(t, Source{SchemaName: SourceSchema, SchemaVersion: SchemaVersion, ModelVersion: ModelVersion, ID: "local.invalid", SourceVersion: "1.0", Kind: PrimaryLocal, Patch: Patch{Checks: &checks, Targets: &targets, RetryPolicies: &retries, Schedules: &bad}})
	if _, err := Resolve([]Source{base, source}); err == nil || !strings.Contains(err.Error(), "unknown retry policy") {
		t.Fatalf("expected reference failure, got %v", err)
	}
}

func TestIntervalAndCalendarBoundaries(t *testing.T) {
	base := mustBuiltIn(t)
	schedules := []Schedule{{ID: "schedule.bad", ContractVersion: ScheduleVersion, Enabled: true, TimeZone: "UTC", Trigger: IntervalTrigger, IntervalNS: 0, Calendar: Calendar{}, DSTPolicy: DSTSkipNonexistent, MisfirePolicy: MisfireSkip, OverlapPolicy: OverlapForbid, ExecutionTimeoutNS: 1, RetryPolicyID: "canonical.default", CheckIDs: []string{}, CommandProfile: "check"}}
	_, err := NormalizeSource(Source{SchemaName: SourceSchema, SchemaVersion: SchemaVersion, ModelVersion: ModelVersion, ID: "local.bad", SourceVersion: "1.0", Kind: PrimaryLocal, Patch: Patch{Schedules: &schedules}})
	if err == nil || !strings.Contains(err.Error(), "interval") {
		t.Fatalf("expected interval failure, got %v", err)
	}

	schedules[0].Trigger, schedules[0].IntervalNS = CalendarTrigger, 0
	schedules[0].Calendar = Calendar{Minutes: []int{60}, Hours: []int{0}}
	_, err = NormalizeSource(Source{SchemaName: SourceSchema, SchemaVersion: SchemaVersion, ModelVersion: ModelVersion, ID: "local.bad", SourceVersion: "1.0", Kind: PrimaryLocal, Patch: Patch{Schedules: &schedules}})
	if err == nil || !strings.Contains(err.Error(), "calendar") {
		t.Fatalf("expected calendar failure, got %v", err)
	}
	if _, err := Resolve([]Source{base}); err != nil {
		t.Fatal(err)
	}
}

func TestStrictDecodeIdentityVersionBoundsAndExtensions(t *testing.T) {
	base := mustBuiltIn(t)
	data, _ := MarshalSourceCanonical(base)
	var object map[string]any
	_ = json.Unmarshal(data, &object)
	object["unknown"] = true
	unknown, _ := json.Marshal(object)
	if _, err := DecodeSource(unknown); err == nil {
		t.Fatal("unknown key accepted")
	}

	tampered := base
	tampered.SourceVersion = "2.0"
	if _, err := NormalizeSource(tampered); err == nil || !strings.Contains(err.Error(), "identity") {
		t.Fatalf("tampered identity accepted: %v", err)
	}
	unsupported := withoutClaim(base)
	unsupported.SchemaVersion = "2.0"
	if _, err := NormalizeSource(unsupported); err == nil {
		t.Fatal("unsupported version accepted")
	}
	tooMany := MaxConcurrency + 1
	if _, err := NormalizeSource(Source{SchemaName: SourceSchema, SchemaVersion: SchemaVersion, ModelVersion: ModelVersion, ID: "local.large", SourceVersion: "1.0", Kind: PrimaryLocal, Patch: Patch{Concurrency: &tooMany}}); err == nil {
		t.Fatal("unbounded concurrency accepted")
	}
	extensions := []Extension{{ID: "future.required", Version: "1.0", Required: true, Fields: map[string]string{}}}
	if _, err := NormalizeSource(Source{SchemaName: SourceSchema, SchemaVersion: SchemaVersion, ModelVersion: ModelVersion, ID: "local.extension", SourceVersion: "1.0", Kind: PrimaryLocal, Patch: Patch{Extensions: &extensions}}); err == nil {
		t.Fatal("required unsupported extension accepted")
	}
}

func TestSecretMaterialCannotEnterTypedContract(t *testing.T) {
	raw := `{"schema_name":"qwsg.configuration-source","schema_version":"1.0","model_version":"1.0","id":"local.secret","identity":"","source_version":"1.0","kind":"primary_local","patch":{"secret_references":[{"id":"mail.password","provider":"vault","reference":"mail-key","value":"seeded-secret"}]}}`
	if _, err := DecodeSource([]byte(raw)); err == nil {
		t.Fatal("secret value field was accepted")
	}
	references := []SecretReference{{ID: "mail.password", Provider: "vault", Reference: "mail-key"}}
	source := mustSource(t, Source{SchemaName: SourceSchema, SchemaVersion: SchemaVersion, ModelVersion: ModelVersion, ID: "local.secret", SourceVersion: "1.0", Kind: PrimaryLocal, Patch: Patch{SecretReferences: &references}})
	data, err := MarshalSourceCanonical(source)
	if err != nil || bytes.Contains(data, []byte("seeded-secret")) {
		t.Fatalf("unsafe serialization: %s %v", data, err)
	}
}

func TestEffectiveIdentityDetectsTampering(t *testing.T) {
	result, err := Resolve([]Source{mustBuiltIn(t)})
	if err != nil {
		t.Fatal(err)
	}
	result.Values.Locale = "de"
	if err := ValidateEffective(result); err == nil || !strings.Contains(err.Error(), "identity") {
		t.Fatalf("tampering accepted: %v", err)
	}
}

func TestNormalizationDoesNotMutateInputAndRejectsInvalidRules(t *testing.T) {
	checks := []Check{{ID: "check.z", Enabled: true, TargetIDs: []string{"target.z", "target.a"}}, {ID: "check.a", Enabled: true, TargetIDs: []string{}}}
	source := Source{SchemaName: SourceSchema, SchemaVersion: SchemaVersion, ModelVersion: ModelVersion, ID: "local.immutable", SourceVersion: "1.0", Kind: PrimaryLocal, Patch: Patch{Checks: &checks}}
	if _, err := NormalizeSource(source); err != nil {
		t.Fatal(err)
	}
	if checks[0].ID != "check.z" || checks[0].TargetIDs[0] != "target.z" {
		t.Fatal("normalization mutated caller-owned slices")
	}

	invalid := []rule.Definition{{ID: "invalid.rule", ContractVersion: rule.RuleVersion, Category: rule.StatusRule, Enabled: true, Description: "invalid condition", Scope: rule.Scope{HealthIDs: []string{}, Categories: []drift.Category{}}, InputRequirements: []rule.Field{}, Condition: rule.Condition{Operator: "unsupported", Values: []rule.Value{}, Children: []rule.Condition{}}, Metadata: map[string]string{}}}
	if _, err := NormalizeSource(Source{SchemaName: SourceSchema, SchemaVersion: SchemaVersion, ModelVersion: ModelVersion, ID: "local.rule", SourceVersion: "1.0", Kind: PrimaryLocal, Patch: Patch{RuleDefinitions: &invalid}}); err == nil {
		t.Fatal("invalid Rule definition was accepted")
	}
}

func mustBuiltIn(t *testing.T) Source {
	t.Helper()
	source, err := BuiltIn([]rule.Definition{}, []policy.Profile{})
	if err != nil {
		t.Fatal(err)
	}
	return source
}

func mustSource(t *testing.T, source Source) Source {
	t.Helper()
	result, err := NormalizeSource(source)
	if err != nil {
		t.Fatal(err)
	}
	return result
}

func withoutClaim(source Source) Source { source.Identity = ""; return source }
