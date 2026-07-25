package presentation

import (
	"strings"
	"testing"

	"quantumwizard.hu/qwsg/internal/command"
)

func TestPresentationConsumesExecutionWithoutOrchestration(t *testing.T) {
	execution := command.Execution{
		SchemaName: command.ExecutionSchema, SchemaVersion: command.SchemaVersion,
		ID: "execution\x1b", CommandID: "command", PlanID: "plan",
		Stages: []command.StageResult{{
			Stage: command.Health, ContractName: "qwsg.health", Version: "1.0",
			RecordCount: 2, Complete: true,
		}},
		View: command.View{Rows: []command.ViewRow{{
			Stage: command.Health, Contract: "qwsg.health", Version: "1.0",
			RecordCount: 2, Complete: true,
		}}, Groups: []command.ViewGroup{}},
		Diagnostics: []string{}, Complete: true,
	}
	first, err := Terminal(execution)
	if err != nil {
		t.Fatal(err)
	}
	second, err := Terminal(execution)
	if err != nil || first != second {
		t.Fatalf("rendering is not deterministic: %v", err)
	}
	if strings.Contains(first, "\x1b") || !strings.Contains(first, `\u001b`) ||
		!strings.Contains(first, "HEALTH") || !strings.Contains(first, "Records: 2") {
		t.Fatalf("unsafe or incomplete terminal output: %q", first)
	}
	document, err := JSON(execution)
	if err != nil || !strings.Contains(string(document), `"schema_name": "qwsg.command-execution"`) {
		t.Fatalf("invalid JSON presentation: %s %v", document, err)
	}
}
