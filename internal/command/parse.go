package command

import (
	"fmt"
	"strings"
)

// ParseProfile applies the bounded common parameter grammar to a predefined
// profile. Environment and transport concerns remain outside this package.
func ParseProfile(name string, args []string, defaults Selection) (Definition, error) {
	profile, ok := profiles[name]
	if !ok {
		return Definition{}, fmt.Errorf("unsupported command profile %q", name)
	}
	advanced := []string{"--source", profile.Source, "--pipeline", string(profile.Target)}
	if defaults.Store != "" {
		advanced = append(advanced, "--store", defaults.Store)
	}
	if defaults.FromSnapshot != "" {
		advanced = append(advanced, "--from", defaults.FromSnapshot, "--to", defaults.ToSnapshot)
	}
	for _, argument := range args {
		if argument == "--source" || argument == "--pipeline" || argument == "--include" || argument == "--exclude" {
			return Definition{}, fmt.Errorf("%s is not valid for a predefined profile", argument)
		}
	}
	advanced = append(advanced, args...)
	definition, err := Parse(advanced)
	if err != nil {
		return Definition{}, err
	}
	definition.Name, definition.Profile = name, name
	if !containsOption(args, "--output") {
		definition.Parameters.Output = profile.Output
	}
	if !containsOption(args, "--presentation") {
		definition.Parameters.Presentation = profile.Presentation
	}
	definition.ID = ""
	return Normalize(definition)
}

// Parse converts the versioned advanced grammar into Command Definition 1.0.
// It has no knowledge of terminal streams or engine implementations.
func Parse(args []string) (Definition, error) {
	definition := Definition{
		SchemaName: SchemaName, SchemaVersion: SchemaVersion, Name: "analyze",
		Selection: Selection{Source: "store"},
		Pipeline:  []Stage{}, IncludedEngines: []Stage{}, ExcludedEngines: []Stage{},
		Parameters: Parameters{
			Filters: []string{}, GroupBy: []string{}, SortBy: []string{},
			Output: JSON, Presentation: Structured,
		},
	}
	seenSingle := map[string]bool{}
	for index := 0; index < len(args); index++ {
		option := args[index]
		value := func() (string, error) {
			index++
			if index >= len(args) || args[index] == "" {
				return "", fmt.Errorf("%s requires one value", option)
			}
			return args[index], nil
		}
		switch option {
		case "--source":
			if seenSingle[option] {
				return Definition{}, fmt.Errorf("duplicate option %s", option)
			}
			seenSingle[option] = true
			got, err := value()
			if err != nil {
				return Definition{}, err
			}
			if got != "live" && got != "store" {
				return Definition{}, fmt.Errorf("--source must be live or store")
			}
			definition.Selection.Source = got
		case "--store":
			got, err := value()
			if err != nil {
				return Definition{}, err
			}
			definition.Selection.Store = got
		case "--from", "--to":
			got, err := value()
			if err != nil {
				return Definition{}, err
			}
			if option == "--from" {
				definition.Selection.FromSnapshot = got
			} else {
				definition.Selection.ToSnapshot = got
			}
		case "--pipeline":
			got, err := value()
			if err != nil {
				return Definition{}, err
			}
			stages, err := parseStages(got)
			if err != nil {
				return Definition{}, err
			}
			definition.Pipeline = stages
		case "--include", "--exclude":
			got, err := value()
			if err != nil {
				return Definition{}, err
			}
			stages, err := parseStages(got)
			if err != nil {
				return Definition{}, err
			}
			if option == "--include" {
				definition.IncludedEngines = append(definition.IncludedEngines, stages...)
			} else {
				definition.ExcludedEngines = append(definition.ExcludedEngines, stages...)
			}
		case "--filter", "--group", "--sort":
			got, err := value()
			if err != nil {
				return Definition{}, err
			}
			switch option {
			case "--filter":
				definition.Parameters.Filters = append(definition.Parameters.Filters, got)
			case "--group":
				definition.Parameters.GroupBy = append(definition.Parameters.GroupBy, got)
			case "--sort":
				definition.Parameters.SortBy = append(definition.Parameters.SortBy, got)
			}
		case "--output":
			got, err := value()
			if err != nil {
				return Definition{}, err
			}
			definition.Parameters.Output = OutputFormat(got)
		case "--presentation":
			got, err := value()
			if err != nil {
				return Definition{}, err
			}
			definition.Parameters.Presentation = Presentation(got)
		default:
			return Definition{}, fmt.Errorf("unsupported command parameter %q", option)
		}
	}
	if definition.Selection.Store == "" && definition.Selection.Source == "store" {
		return Definition{}, fmt.Errorf("--store is required for store source")
	}
	if (definition.Selection.FromSnapshot == "") != (definition.Selection.ToSnapshot == "") {
		return Definition{}, fmt.Errorf("--from and --to must be provided together")
	}
	return Normalize(definition)
}

func containsOption(args []string, option string) bool {
	for _, argument := range args {
		if argument == option {
			return true
		}
	}
	return false
}

func parseStages(value string) ([]Stage, error) {
	parts := strings.Split(value, ",")
	result := make([]Stage, 0, len(parts))
	for _, part := range parts {
		stage := Stage(part)
		if !validStage(stage) {
			return nil, fmt.Errorf("unsupported pipeline stage %q", part)
		}
		result = append(result, stage)
	}
	return result, nil
}
