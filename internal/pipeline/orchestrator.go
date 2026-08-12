// Package pipeline is the only canonical orchestration layer. It connects the
// existing deterministic engineering engines; it contains no CLI, terminal,
// dashboard, HTTP, scheduling, alerting, or remediation behavior.
package pipeline

import (
	"context"
	"fmt"

	"quantumwizard.hu/qwsg/internal/command"
	"quantumwizard.hu/qwsg/internal/comparison"
	"quantumwizard.hu/qwsg/internal/configuration"
	"quantumwizard.hu/qwsg/internal/drift"
	"quantumwizard.hu/qwsg/internal/health"
	"quantumwizard.hu/qwsg/internal/inventory"
	"quantumwizard.hu/qwsg/internal/inventorystore"
	"quantumwizard.hu/qwsg/internal/policy"
	"quantumwizard.hu/qwsg/internal/report"
	"quantumwizard.hu/qwsg/internal/rule"
)

const EngineVersion = "1.0"

type Collector func(context.Context) (inventory.Snapshot, error)

type Orchestrator struct {
	Collect       Collector
	Configuration *configuration.Effective
	// Retention, Rules, and Policies are compatibility inputs. They are
	// projected through the Canonical Configuration resolver before use.
	Retention int
	Rules     []rule.Definition
	Policies  []policy.Profile
}

// CanonicalPolicyProfiles returns the presentation-independent default policy
// used by Command 1.0. It interprets matched observations as observable,
// non-matches as accepted, and incomplete technical outcomes as indeterminate.
func CanonicalPolicyProfiles() []policy.Profile {
	profile, err := policy.NormalizeProfile(policy.Profile{
		ID: "qwsg.canonical.observation", ContractVersion: policy.ProfileVersion,
		Version: "1.0", Priority: 0, Extends: []string{}, Enabled: true,
		Scope: policy.Selector{RuleIDs: []string{}, Outcomes: []rule.Outcome{}},
		Statements: []policy.Statement{
			{ID: "matched.observe", Priority: 100, Selector: policy.Selector{RuleIDs: []string{}, Outcomes: []rule.Outcome{rule.Matched}}, Outcome: policy.Observe, Explanation: "matched_rule_requires_observation", Metadata: map[string]string{}},
			{ID: "not-matched.accept", Priority: 100, Selector: policy.Selector{RuleIDs: []string{}, Outcomes: []rule.Outcome{rule.NotMatched}}, Outcome: policy.Accepted, Explanation: "unmatched_rule_is_accepted", Metadata: map[string]string{}},
		},
		DefaultOutcome: policy.Indeterminate,
		Metadata:       map[string]string{"profile": "canonical-command-1.0"},
	})
	if err != nil {
		return nil
	}
	return []policy.Profile{profile}
}

// CanonicalObservationRules returns the fixed Rule Definition configuration
// used by the Command 1.0 report profile. Presentation adapters never define
// or evaluate these rules.
func CanonicalObservationRules() []rule.Definition {
	return []rule.Definition{{
		ID: "qwsg.canonical.health-observed", ContractVersion: rule.RuleVersion,
		Category: rule.StatusRule, Enabled: true,
		Scope:             rule.Scope{HealthIDs: []string{}, Categories: []drift.Category{}},
		InputRequirements: []rule.Field{rule.FieldStatus},
		Condition: rule.Condition{
			Operator: rule.Exists, Field: rule.FieldStatus,
			Values: []rule.Value{}, Children: []rule.Condition{},
		},
		Description: "Canonical health evidence is present",
		Metadata:    map[string]string{"profile": "canonical-command-1.0"},
	}}
}

func (orchestrator Orchestrator) Execute(ctx context.Context, definition command.Definition) (command.Execution, error) {
	plan, err := command.PlanDefinition(definition)
	if err != nil {
		return command.Execution{}, err
	}
	effective, err := orchestrator.effectiveConfiguration()
	if err != nil {
		return command.Execution{}, fmt.Errorf("configuration resolution failed: %w", err)
	}
	execution := command.Execution{
		SchemaName: command.ExecutionSchema, SchemaVersion: command.SchemaVersion,
		CommandID: definition.ID, PlanID: plan.ID,
		Stages: []command.StageResult{}, View: command.View{Rows: []command.ViewRow{}, Groups: []command.ViewGroup{}},
		Diagnostics: []string{}, Complete: true,
	}
	var current, from inventory.Snapshot
	var fromSelector, toSelector string
	var compared comparison.Result
	var drifted drift.Result
	var evaluated health.Result
	var ruled rule.Result
	var governed policy.Result

	if plan.Selection.Source == "live" && containsStage(plan.Stages, command.Compare) {
		from, fromSelector, err = orchestrator.loadBaseline(plan.Selection, effective.Values.SnapshotRetention)
		if err != nil {
			return command.Execution{}, fmt.Errorf("live baseline selection failed: %w", err)
		}
	}
	for _, stage := range plan.Stages {
		switch stage {
		case command.Inventory:
			if orchestrator.Collect == nil {
				return command.Execution{}, fmt.Errorf("inventory collector is unavailable")
			}
			current, err = orchestrator.Collect(ctx)
			if err == nil {
				err = inventory.Validate(current)
			}
			execution.Stages = append(execution.Stages, stageResult(stage, inventory.CanonicalSchemaName, inventory.SchemaVersion, len(current.Canonical.Layers), current.Status == inventory.Complete, current))
		case command.Snapshot:
			if current.SnapshotID == "" {
				return command.Execution{}, fmt.Errorf("snapshot stage requires inventory")
			}
			toSelector = current.SnapshotID
			if plan.Selection.Store != "" {
				var store *inventorystore.Store
				store, err = inventorystore.Open(plan.Selection.Store, effective.Values.SnapshotRetention)
				if err == nil {
					toSelector, err = store.Save(current)
				}
			}
			execution.Stages = append(execution.Stages, stageResult(stage, inventorystore.FormatName, inventorystore.FormatVersion, 1, current.Status == inventory.Complete, current))
		case command.Compare:
			if plan.Selection.Source == "store" {
				from, current, fromSelector, toSelector, err = orchestrator.loadPair(plan.Selection, effective.Values.SnapshotRetention)
			}
			if err == nil {
				compared, err = comparison.Compare(from, current, fromSelector, toSelector)
			}
			execution.Stages = append(execution.Stages, stageResult(stage, comparison.SchemaName, comparison.SchemaVersion, len(compared.Changes), err == nil, compared))
		case command.Drift:
			drifted, err = drift.Classify(compared.Changes)
			execution.Stages = append(execution.Stages, stageResult(stage, drift.SchemaName, drift.SchemaVersion, len(drifted.Records), err == nil, drifted))
		case command.Health:
			evaluated, err = health.Evaluate(drifted)
			execution.Stages = append(execution.Stages, stageResult(stage, health.SchemaName, health.SchemaVersion, len(evaluated.Records), err == nil && evaluated.EvidenceState == health.EvidenceSufficient, evaluated))
		case command.Rule:
			ruled, err = rule.Evaluate(effective.Values.RuleDefinitions, evaluated)
			execution.Stages = append(execution.Stages, stageResult(stage, rule.SchemaName, rule.SchemaVersion, len(ruled.Records), err == nil, ruled))
		case command.Policy:
			governed, err = policy.Evaluate(effective.Values.PolicyProfiles, ruled)
			execution.Stages = append(execution.Stages, stageResult(stage, policy.SchemaName, policy.SchemaVersion, len(governed.Records), err == nil, governed))
		case command.Report:
			var generated report.PolicyReport
			generated, err = report.GeneratePolicy(governed)
			execution.Stages = append(execution.Stages, stageResult(stage, report.PolicyReportSchemaName, report.PolicyReportSchemaVersion, generated.Summary.Total, err == nil && generated.Completeness == report.Complete, generated))
		default:
			err = fmt.Errorf("unsupported pipeline stage %q", stage)
		}
		if err != nil {
			execution.Complete = false
			execution.Diagnostics = append(execution.Diagnostics, string(stage)+": "+err.Error())
			return execution, fmt.Errorf("%s stage failed: %w", stage, err)
		}
	}
	execution.View, err = command.BuildView(execution.Stages, plan.Parameters)
	if err != nil {
		return command.Execution{}, fmt.Errorf("command result projection failed: %w", err)
	}
	execution.ID = executionID(execution)
	return execution, ValidateExecution(execution, plan)
}

func (orchestrator Orchestrator) effectiveConfiguration() (configuration.Effective, error) {
	if orchestrator.Configuration != nil {
		if orchestrator.Retention != 0 || orchestrator.Rules != nil || orchestrator.Policies != nil {
			return configuration.Effective{}, fmt.Errorf("effective configuration cannot be combined with compatibility configuration fields")
		}
		if err := configuration.ValidateEffective(*orchestrator.Configuration); err != nil {
			return configuration.Effective{}, err
		}
		return *orchestrator.Configuration, nil
	}
	rules := orchestrator.Rules
	if rules == nil {
		rules = CanonicalObservationRules()
	}
	policies := orchestrator.Policies
	if policies == nil {
		policies = CanonicalPolicyProfiles()
	}
	source, err := configuration.BuiltIn(rules, policies)
	if err != nil {
		return configuration.Effective{}, err
	}
	if orchestrator.Retention != 0 {
		retention := orchestrator.Retention
		override, normalizeErr := configuration.NormalizeSource(configuration.Source{
			SchemaName: configuration.SourceSchema, SchemaVersion: configuration.SchemaVersion,
			ModelVersion: configuration.ModelVersion, ID: "qwsg.compatibility.retention",
			SourceVersion: "1.0", Kind: configuration.TemporaryOverride,
			Patch: configuration.Patch{SnapshotRetention: &retention},
		})
		if normalizeErr != nil {
			return configuration.Effective{}, normalizeErr
		}
		return configuration.Resolve([]configuration.Source{source, override})
	}
	return configuration.Resolve([]configuration.Source{source})
}

func containsStage(stages []command.Stage, wanted command.Stage) bool {
	for _, stage := range stages {
		if stage == wanted {
			return true
		}
	}
	return false
}

func (orchestrator Orchestrator) loadPair(selection command.Selection, retention int) (inventory.Snapshot, inventory.Snapshot, string, string, error) {
	store, err := inventorystore.Open(selection.Store, retention)
	if err != nil {
		return inventory.Snapshot{}, inventory.Snapshot{}, "", "", err
	}
	fromName, toName := selection.FromSnapshot, selection.ToSnapshot
	if fromName == "" {
		names, listErr := store.List()
		if listErr != nil {
			return inventory.Snapshot{}, inventory.Snapshot{}, "", "", listErr
		}
		if len(names) < 2 {
			return inventory.Snapshot{}, inventory.Snapshot{}, "", "", fmt.Errorf("at least two snapshots are required")
		}
		fromName, toName = names[len(names)-2], names[len(names)-1]
	}
	from, err := store.Load(fromName)
	if err != nil {
		return inventory.Snapshot{}, inventory.Snapshot{}, "", "", fmt.Errorf("load from snapshot: %w", err)
	}
	to, err := store.Load(toName)
	return from, to, fromName, toName, err
}

func (orchestrator Orchestrator) loadBaseline(selection command.Selection, retention int) (inventory.Snapshot, string, error) {
	if selection.Store == "" {
		return inventory.Snapshot{}, "", fmt.Errorf("live comparison requires a snapshot store")
	}
	store, err := inventorystore.Open(selection.Store, retention)
	if err != nil {
		return inventory.Snapshot{}, "", err
	}
	name := selection.FromSnapshot
	if name == "" {
		names, listErr := store.List()
		if listErr != nil || len(names) == 0 {
			if listErr != nil {
				return inventory.Snapshot{}, "", listErr
			}
			return inventory.Snapshot{}, "", fmt.Errorf("a baseline snapshot is required")
		}
		name = names[len(names)-1]
	}
	snapshot, err := store.Load(name)
	return snapshot, name, err
}

func stageResult(stage command.Stage, contract, version string, count int, complete bool, value any) command.StageResult {
	return command.StageResult{Stage: stage, ContractName: contract, Version: version, RecordCount: count, Complete: complete, Value: value}
}
