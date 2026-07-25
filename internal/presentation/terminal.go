// Package presentation contains replaceable consumers of canonical command
// executions. It may render data but may not plan or execute engineering work.
package presentation

import (
	"encoding/json"
	"fmt"
	"strings"
	"unicode"

	"quantumwizard.hu/qwsg/internal/command"
)

func JSON(execution command.Execution) ([]byte, error) {
	document, err := json.MarshalIndent(execution, "", "  ")
	if err != nil {
		return nil, err
	}
	return append(document, '\n'), nil
}

func Terminal(execution command.Execution) (string, error) {
	var output strings.Builder
	fmt.Fprintf(&output, "QWSG canonical command\nExecution: %s\nComplete: %t\n", safe(execution.ID), execution.Complete)
	for _, row := range execution.View.Rows {
		fmt.Fprintf(&output, "\n%s\nContract: %s/%s\nRecords: %d\nComplete: %t\n",
			strings.ToUpper(string(row.Stage)), safe(row.Contract), safe(row.Version), row.RecordCount, row.Complete)
	}
	return output.String(), nil
}

func safe(value string) string {
	var output strings.Builder
	for _, character := range value {
		if unicode.IsControl(character) {
			fmt.Fprintf(&output, "\\u%04x", character)
		} else {
			output.WriteRune(character)
		}
	}
	return output.String()
}
