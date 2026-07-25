package command

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

var viewFields = map[string]bool{
	"stage": true, "contract": true, "version": true,
	"record_count": true, "complete": true,
}

// BuildView applies deterministic filtering, grouping, and sorting to stage
// metadata. It never changes or reinterprets canonical engine values.
func BuildView(stages []StageResult, parameters Parameters) (View, error) {
	rows := make([]ViewRow, 0, len(stages))
	for _, result := range stages {
		row := ViewRow{
			Stage: result.Stage, Contract: result.ContractName, Version: result.Version,
			RecordCount: result.RecordCount, Complete: result.Complete,
		}
		matches := true
		for _, filter := range parameters.Filters {
			field, expected, ok := strings.Cut(filter, "=")
			if !ok || !viewFields[field] {
				return View{}, fmt.Errorf("unsupported filter %q", filter)
			}
			if rowValue(row, field) != expected {
				matches = false
			}
		}
		if matches {
			rows = append(rows, row)
		}
	}
	sortFields := parameters.SortBy
	if len(sortFields) == 0 {
		sortFields = []string{"stage"}
	}
	for _, field := range sortFields {
		if !viewFields[field] {
			return View{}, fmt.Errorf("unsupported sort field %q", field)
		}
	}
	sort.SliceStable(rows, func(i, j int) bool {
		for _, field := range sortFields {
			if field == "stage" && rows[i].Stage != rows[j].Stage {
				return stageRank(rows[i].Stage) < stageRank(rows[j].Stage)
			}
			left, right := rowValue(rows[i], field), rowValue(rows[j], field)
			if left != right {
				return left < right
			}
		}
		return stageRank(rows[i].Stage) < stageRank(rows[j].Stage)
	})
	groups := []ViewGroup{}
	if len(parameters.GroupBy) > 1 {
		return View{}, fmt.Errorf("Command Definition 1.0 supports one grouping field")
	}
	if len(parameters.GroupBy) == 1 {
		field := parameters.GroupBy[0]
		if !viewFields[field] {
			return View{}, fmt.Errorf("unsupported grouping field %q", field)
		}
		byValue := map[string][]ViewRow{}
		for _, row := range rows {
			value := rowValue(row, field)
			byValue[value] = append(byValue[value], row)
		}
		values := make([]string, 0, len(byValue))
		for value := range byValue {
			values = append(values, value)
		}
		sort.Strings(values)
		for _, value := range values {
			groups = append(groups, ViewGroup{Key: field, Value: value, Rows: byValue[value]})
		}
	}
	return View{Rows: rows, Groups: groups}, nil
}

func rowValue(row ViewRow, field string) string {
	switch field {
	case "stage":
		return string(row.Stage)
	case "contract":
		return row.Contract
	case "version":
		return row.Version
	case "record_count":
		return fmt.Sprintf("%020d", row.RecordCount)
	case "complete":
		return strconv.FormatBool(row.Complete)
	default:
		return ""
	}
}
