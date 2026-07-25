package pipeline

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"

	"quantumwizard.hu/qwsg/internal/command"
)

func ValidateExecution(execution command.Execution, plan command.Plan) error {
	if execution.SchemaName != command.ExecutionSchema ||
		execution.SchemaVersion != command.SchemaVersion ||
		execution.CommandID != plan.CommandID || execution.PlanID != plan.ID ||
		execution.Stages == nil || execution.Diagnostics == nil ||
		execution.View.Rows == nil || execution.View.Groups == nil ||
		len(execution.Stages) != len(plan.Stages) {
		return fmt.Errorf("invalid command execution envelope")
	}
	for index, result := range execution.Stages {
		if result.Stage != plan.Stages[index] || result.ContractName == "" ||
			result.Version == "" || result.RecordCount < 0 {
			return fmt.Errorf("invalid or unordered command stage result")
		}
	}
	if len(execution.Diagnostics) != 0 || !execution.Complete {
		return fmt.Errorf("completed execution contains diagnostics")
	}
	expectedView, err := command.BuildView(execution.Stages, plan.Parameters)
	if err != nil || !equalView(execution.View, expectedView) {
		return fmt.Errorf("invalid command execution view")
	}
	if execution.ID != executionID(execution) {
		return fmt.Errorf("invalid command execution identity")
	}
	return nil
}

func equalView(left, right command.View) bool {
	a, _ := json.Marshal(left)
	b, _ := json.Marshal(right)
	return string(a) == string(b)
}

func executionID(execution command.Execution) string {
	copy := execution
	copy.ID = ""
	document, _ := json.Marshal(copy)
	sum := sha256.Sum256(append([]byte(command.ExecutionSchema+"/"+EngineVersion+"\x00"), document...))
	return hex.EncodeToString(sum[:])
}
