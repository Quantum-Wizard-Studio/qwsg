// Package pipeline is the only canonical orchestration layer. It connects the
// existing deterministic engineering engines; it contains no CLI, terminal,
// dashboard, HTTP, scheduling, policy, alerting, or remediation behavior.
package pipeline

import (
	"context"
	"fmt"

	"quantumwizard.hu/qwsg/internal/command"
	"quantumwizard.hu/qwsg/internal/comparison"
	"quantumwizard.hu/qwsg/internal/drift"
	"quantumwizard.hu/qwsg/internal/health"
	"quantumwizard.hu/qwsg/internal/inventory"
	"quantumwizard.hu/qwsg/internal/inventorystore"
	"quantumwizard.hu/qwsg/internal/report"
	"quantumwizard.hu/qwsg/internal/rule"
)

const EngineVersion = "1.0"

type Collector func(context.Context) (inventory.Snapshot, error)

type Orchestrator struct {
	Collect   Collector
	Retention int
	Rules     []rule.Definition
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

	if plan.Selection.Source == "live" && containsStage(plan.Stages, command.Compare) {
		from, fromSelector, err = orchestrator.loadBaseline(plan.Selection)
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
				store, err = inventorystore.Open(plan.Selection.Store, retention(orchestrator.Retention))
				if err == nil {
					toSelector, err = store.Save(current)
				}
			}
			execution.Stages = append(execution.Stages, stageResult(stage, inventorystore.FormatName, inventorystore.FormatVersion, 1, current.Status == inventory.Complete, current))
		case command.Compare:
			if plan.Selection.Source == "store" {
				from, current, fromSelector, toSelector, err = orchestrator.loadPair(plan.Selection)
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
			ruled, err = rule.Evaluate(orchestrator.Rules, evaluated)
			execution.Stages = append(execution.Stages, stageResult(stage, rule.SchemaName, rule.SchemaVersion, len(ruled.Records), err == nil, ruled))
		case command.Report:
			var generated report.Report
			generated, err = report.Generate(ruled)
			execution.Stages = append(execution.Stages, stageResult(stage, report.SchemaName, report.SchemaVersion, generated.Summary.Total, err == nil && generated.Completeness == report.Complete, generated))
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

func containsStage(stages []command.Stage, wanted command.Stage) bool {
	for _, stage := range stages {
		if stage == wanted {
			return true
		}
	}
	return false
}

func (orchestrator Orchestrator) loadPair(selection command.Selection) (inventory.Snapshot, inventory.Snapshot, string, string, error) {
	store, err := inventorystore.Open(selection.Store, retention(orchestrator.Retention))
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

func (orchestrator Orchestrator) loadBaseline(selection command.Selection) (inventory.Snapshot, string, error) {
	if selection.Store == "" {
		return inventory.Snapshot{}, "", fmt.Errorf("live comparison requires a snapshot store")
	}
	store, err := inventorystore.Open(selection.Store, retention(orchestrator.Retention))
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

func retention(value int) int {
	if value == 0 {
		return inventorystore.DefaultRetention
	}
	return value
}
