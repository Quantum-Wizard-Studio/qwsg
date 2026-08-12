package app

import (
	"fmt"
	"reflect"
	"time"

	"quantumwizard.hu/qwsg/internal/command"
	"quantumwizard.hu/qwsg/internal/comparison"
	"quantumwizard.hu/qwsg/internal/drift"
	"quantumwizard.hu/qwsg/internal/health"
	"quantumwizard.hu/qwsg/internal/inventory"
	"quantumwizard.hu/qwsg/internal/inventorystore"
	"quantumwizard.hu/qwsg/internal/policy"
	"quantumwizard.hu/qwsg/internal/presentationmodel"
	"quantumwizard.hu/qwsg/internal/report"
	"quantumwizard.hu/qwsg/internal/rule"
	"quantumwizard.hu/qwsg/internal/runtime"
	"quantumwizard.hu/qwsg/internal/runtimeservice"
)

// ProjectOperatorEvaluation validates and projects the exact typed values from
// one already completed canonical observe execution. It performs no execution
// and owns no engineering decision.
func ProjectOperatorEvaluation(execution command.Execution, observedAt time.Time, freshFor time.Duration, runtimeResult *runtime.Result, serviceState *runtimeservice.State) (presentationmodel.Overview, error) {
	if !execution.Complete || len(execution.Stages) != 8 || observedAt.IsZero() || observedAt.Location() != time.UTC || freshFor <= 0 || freshFor > presentationmodel.MaxFreshFor {
		return presentationmodel.Overview{}, fmt.Errorf("ineligible operator evaluation")
	}
	expected := []struct {
		stage    command.Stage
		contract string
		version  string
	}{{command.Inventory, inventory.CanonicalSchemaName, inventory.SchemaVersion}, {command.Snapshot, inventorystore.FormatName, inventorystore.FormatVersion}, {command.Compare, comparison.SchemaName, comparison.SchemaVersion}, {command.Drift, drift.SchemaName, drift.SchemaVersion}, {command.Health, health.SchemaName, health.SchemaVersion}, {command.Rule, rule.SchemaName, rule.SchemaVersion}, {command.Policy, policy.SchemaName, policy.SchemaVersion}, {command.Report, report.PolicyReportSchemaName, report.PolicyReportSchemaVersion}}
	for index, want := range expected {
		got := execution.Stages[index]
		if got.Stage != want.stage || got.ContractName != want.contract || got.Version != want.version {
			return presentationmodel.Overview{}, fmt.Errorf("invalid operator stage coverage")
		}
	}
	observed, ok := execution.Stages[0].Value.(inventory.Snapshot)
	if !ok || inventory.Validate(observed) != nil {
		return presentationmodel.Overview{}, fmt.Errorf("invalid typed inventory")
	}
	snapshotted, ok := execution.Stages[1].Value.(inventory.Snapshot)
	if !ok || inventory.Validate(snapshotted) != nil || !reflect.DeepEqual(observed, snapshotted) {
		return presentationmodel.Overview{}, fmt.Errorf("invalid typed snapshot")
	}
	compared, ok := execution.Stages[2].Value.(comparison.Result)
	if !ok || comparison.Validate(compared) != nil || compared.SubjectID != observed.InstanceID || compared.To.SnapshotID != observed.SnapshotID || !compared.ComparisonTimestamp.Equal(observed.CompletedAt) {
		return presentationmodel.Overview{}, fmt.Errorf("invalid typed comparison")
	}
	drifted, ok := execution.Stages[3].Value.(drift.Result)
	if !ok || drift.Validate(drifted) != nil {
		return presentationmodel.Overview{}, fmt.Errorf("invalid typed drift")
	}
	evaluated, ok := execution.Stages[4].Value.(health.Result)
	if !ok || health.Validate(evaluated) != nil {
		return presentationmodel.Overview{}, fmt.Errorf("invalid typed health")
	}
	ruled, ok := execution.Stages[5].Value.(rule.Result)
	if !ok || rule.Validate(ruled) != nil {
		return presentationmodel.Overview{}, fmt.Errorf("invalid typed rule")
	}
	governed, ok := execution.Stages[6].Value.(policy.Result)
	if !ok || policy.Validate(governed) != nil {
		return presentationmodel.Overview{}, fmt.Errorf("invalid typed policy")
	}
	reported, ok := execution.Stages[7].Value.(report.PolicyReport)
	if !ok || report.ValidatePolicyReport(reported) != nil || !operatorCorrelated(evaluated, ruled, governed, reported) {
		return presentationmodel.Overview{}, fmt.Errorf("invalid typed report correlation")
	}
	input := presentationmodel.Input{SchemaName: presentationmodel.InputSchema, SchemaVersion: presentationmodel.SchemaVersion, ObservedAt: observedAt, FreshForNS: int64(freshFor), Command: &presentationmodel.CommandObservation{ObservedAt: observedAt, Value: execution}, Inventory: &presentationmodel.InventoryObservation{ObservedAt: observedAt, Value: observed}, Comparison: &presentationmodel.ComparisonObservation{ObservedAt: observedAt, Value: compared}, Drift: &presentationmodel.DriftObservation{ObservedAt: observedAt, Value: drifted}, Health: &presentationmodel.HealthObservation{ObservedAt: observedAt, Value: evaluated}, Rule: &presentationmodel.RuleObservation{ObservedAt: observedAt, Value: ruled}, Policy: &presentationmodel.PolicyObservation{ObservedAt: observedAt, Value: governed}, PolicyReport: &presentationmodel.PolicyReportObservation{ObservedAt: observedAt, Value: reported}}
	if runtimeResult != nil {
		input.Runtime = &presentationmodel.RuntimeObservation{ObservedAt: observedAt, Value: *runtimeResult}
	}
	if serviceState != nil {
		input.ServiceState = &presentationmodel.ServiceStateObservation{ObservedAt: observedAt, Value: *serviceState}
	}
	return presentationmodel.Project(input)
}

func operatorCorrelated(evaluated health.Result, ruled rule.Result, governed policy.Result, reported report.PolicyReport) bool {
	healthIDs := make(map[string]bool, len(evaluated.Records))
	for _, record := range evaluated.Records {
		healthIDs[record.ID] = true
	}
	ruleIDs := make(map[string]bool, len(ruled.Records))
	for _, record := range ruled.Records {
		if record.HealthRecordID != "" && !healthIDs[record.HealthRecordID] {
			return false
		}
		ruleIDs[record.ID] = true
	}
	policyIDs := make(map[string]bool, len(governed.Records))
	for _, record := range governed.Records {
		if !ruleIDs[record.RuleEvaluationID] {
			return false
		}
		policyIDs[record.ID] = true
	}
	seen := make(map[string]bool, len(reported.Sources))
	for _, source := range reported.Sources {
		if source.Type != report.PolicyEvaluationSource || !policyIDs[source.ID] {
			return false
		}
		seen[source.ID] = true
	}
	if len(seen) != len(policyIDs) {
		return false
	}
	for _, section := range reported.Sections {
		for _, item := range section.Items {
			if !policyIDs[item.PolicyEvaluationID] || !ruleIDs[item.RuleEvaluationID] {
				return false
			}
		}
	}
	return true
}

// ProjectGuardianLifecycle creates a bounded lifecycle-only Overview before a
// qualified engineering cycle exists.
func ProjectGuardianLifecycle(state runtimeservice.State, observedAt time.Time, freshFor time.Duration) (presentationmodel.Overview, error) {
	return presentationmodel.Project(presentationmodel.Input{SchemaName: presentationmodel.InputSchema, SchemaVersion: presentationmodel.SchemaVersion, ObservedAt: observedAt, FreshForNS: int64(freshFor), ServiceState: &presentationmodel.ServiceStateObservation{ObservedAt: observedAt, Value: state}})
}
