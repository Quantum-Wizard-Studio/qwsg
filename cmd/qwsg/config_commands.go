package main

import (
	"errors"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"time"

	"quantumwizard.hu/qwsg/internal/configuration"
	"quantumwizard.hu/qwsg/internal/configurationstore"
	"quantumwizard.hu/qwsg/internal/pipeline"
	"quantumwizard.hu/qwsg/internal/runtimeservice"
)

type configOptions struct {
	path   string
	format string
}

type configResult struct {
	Status     string                  `json:"status"`
	Path       string                  `json:"path"`
	Configured bool                    `json:"configured"`
	Effective  configuration.Effective `json:"effective_configuration"`
}

type configValidationResult struct {
	Status      string `json:"status"`
	Path        string `json:"path"`
	Configured  bool   `json:"configured"`
	EffectiveID string `json:"effective_id"`
}

func runConfig(args []string, out, errout io.Writer) int {
	if len(args) == 0 || isHelp(args[0]) {
		writeConfigHelp(out)
		return 0
	}
	action, rest := args[0], args[1:]
	if action != "show" && action != "validate" && action != "get" && action != "set" {
		return usageError(errout, "unknown config operation: %s", safeText(action))
	}
	key, value := "", ""
	if action == "get" || action == "set" {
		if len(rest) == 0 || strings.HasPrefix(rest[0], "-") {
			return usageError(errout, "config %s requires a key", action)
		}
		key, rest = rest[0], rest[1:]
	}
	if action == "set" {
		if len(rest) == 0 || strings.HasPrefix(rest[0], "-") {
			return usageError(errout, "config set requires a value")
		}
		value, rest = rest[0], rest[1:]
	}
	options, err := parseConfigOptions(rest)
	if err != nil {
		return usageError(errout, "%v", err)
	}
	path, err := configurationstore.SelectPath(options.path, os.Getenv)
	if err != nil {
		return configFailure(errout, err)
	}
	source, found, err := configurationstore.Load(path)
	if err != nil {
		return configFailure(errout, err)
	}
	if options.path != "" && !found && action != "set" {
		return configFailure(errout, configurationstore.ErrUnavailable)
	}
	effective, err := resolveLocalConfiguration(source, found, nil)
	if err != nil {
		return configFailure(errout, err)
	}
	if err = validateGuardianTiming(effective); err != nil {
		return configFailure(errout, err)
	}
	switch action {
	case "show":
		result := configResult{Status: "valid", Path: path, Configured: found, Effective: effective}
		if options.format == formatJSON {
			return writeJSON(out, errout, result)
		}
		fmt.Fprintf(out, "Configuration: valid\nPath: %s\nConfigured: %t\nEffective ID: %s\nLocale: %s\nTime zone: %s\nGuardian interval: %s\nGuardian cycle timeout: %s\n", safeText(path), found, safeText(effective.ID), safeText(effective.Values.Locale), safeText(effective.Values.TimeZone), guardianInterval(effective), guardianTimeout(effective))
		return 0
	case "validate":
		if options.format == formatJSON {
			return writeJSON(out, errout, configValidationResult{Status: "valid", Path: path, Configured: found, EffectiveID: effective.ID})
		}
		fmt.Fprintf(out, "Configuration: valid\nPath: %s\nConfigured: %t\nEffective ID: %s\n", safeText(path), found, safeText(effective.ID))
		return 0
	case "get":
		got, getErr := getConfigValue(effective, key)
		if getErr != nil {
			return usageError(errout, "%v", getErr)
		}
		if options.format == formatJSON {
			return writeJSON(out, errout, map[string]string{"key": key, "value": got})
		}
		fmt.Fprintln(out, safeText(got))
		return 0
	case "set":
		updated, setErr := setConfigValue(source, found, effective, key, value)
		if setErr != nil {
			return usageError(errout, "%v", setErr)
		}
		resolved, setErr := resolveLocalConfiguration(updated, true, nil)
		if setErr != nil {
			return configFailure(errout, setErr)
		}
		if setErr = validateGuardianTiming(resolved); setErr != nil {
			return configFailure(errout, setErr)
		}
		if setErr = configurationstore.Save(path, updated); setErr != nil {
			return configFailure(errout, setErr)
		}
		if options.format == formatJSON {
			return writeJSON(out, errout, configResult{Status: "updated", Path: path, Configured: true, Effective: resolved})
		}
		fmt.Fprintf(out, "Configuration updated.\nPath: %s\n%s=%s\n", safeText(path), safeText(key), safeText(value))
		return 0
	}
	return 1
}

func runSetup(args []string, in io.Reader, out, errout io.Writer, interactive bool) int {
	accept, sets, remaining := false, []string{}, []string{}
	for i := 0; i < len(args); i++ {
		switch args[i] {
		case "--accept-defaults":
			accept = true
		case "--set":
			i++
			if i >= len(args) {
				return usageError(errout, "--set requires key=value")
			}
			sets = append(sets, args[i])
		default:
			remaining = append(remaining, args[i])
		}
	}
	options, err := parseConfigOptions(remaining)
	if err != nil {
		return usageError(errout, "%v", err)
	}
	if options.format == formatJSON && !accept {
		return usageError(errout, "JSON setup requires --accept-defaults")
	}
	path, err := configurationstore.SelectPath(options.path, os.Getenv)
	if err != nil {
		return configFailure(errout, err)
	}
	source, found, err := configurationstore.Load(path)
	if err != nil {
		return configFailure(errout, err)
	}
	effective, err := resolveLocalConfiguration(source, found, nil)
	if err != nil {
		return configFailure(errout, err)
	}
	updated := source
	if !found {
		updated = newLocalSource(effective)
	}
	for _, item := range sets {
		parts := strings.SplitN(item, "=", 2)
		if len(parts) != 2 {
			return usageError(errout, "--set requires key=value")
		}
		updated, err = setConfigValue(updated, true, effective, parts[0], parts[1])
		if err != nil {
			return usageError(errout, "%v", err)
		}
		effective, err = resolveLocalConfiguration(updated, true, nil)
		if err != nil {
			return configFailure(errout, err)
		}
	}
	if err = validateGuardianTiming(effective); err != nil {
		return configFailure(errout, err)
	}
	if options.format == formatHuman {
		fmt.Fprintf(out, "QWSG setup plan\nPath: %s\nExisting configuration: %t\nLocale: %s\nTime zone: %s\nGuardian interval: %s\nGuardian cycle timeout: %s\n", safeText(path), found, safeText(effective.Values.Locale), safeText(effective.Values.TimeZone), guardianInterval(effective), guardianTimeout(effective))
	}
	if !accept {
		if !interactive {
			return usageError(errout, "setup requires a terminal or --accept-defaults")
		}
		fmt.Fprint(out, "Write this configuration? [y/N]: ")
		var answer string
		if _, err = fmt.Fscanln(in, &answer); err != nil || strings.ToLower(answer) != "y" {
			fmt.Fprintln(out, "Setup cancelled; no changes written.")
			return 0
		}
	}
	if err = configurationstore.Save(path, updated); err != nil {
		return configFailure(errout, err)
	}
	if options.format == formatJSON {
		return writeJSON(out, errout, configResult{Status: "configured", Path: path, Configured: true, Effective: effective})
	}
	fmt.Fprintln(out, "Configuration written safely.")
	fmt.Fprintln(out, "Next: qwsg config validate; install and enable the user service explicitly; then run qwsg console.")
	return 0
}

func parseConfigOptions(args []string) (configOptions, error) {
	result := configOptions{format: formatHuman}
	for i := 0; i < len(args); i++ {
		if i+1 >= len(args) {
			return result, fmt.Errorf("configuration option requires a value")
		}
		name, value := args[i], args[i+1]
		i++
		switch name {
		case "--config":
			if result.path != "" {
				return result, fmt.Errorf("--config may be specified once")
			}
			result.path = value
		case "--format":
			if !validFormat(value) {
				return result, fmt.Errorf("--format must be json or human")
			}
			result.format = value
		default:
			return result, fmt.Errorf("unknown configuration option: %s", safeText(name))
		}
	}
	return result, nil
}

func resolveLocalConfiguration(source configuration.Source, found bool, temporary *configuration.Source) (configuration.Effective, error) {
	builtIn, err := configuration.BuiltIn(pipeline.CanonicalObservationRules(), pipeline.CanonicalPolicyProfiles())
	if err != nil {
		return configuration.Effective{}, err
	}
	schedules := []configuration.Schedule{defaultGuardianSchedule()}
	builtIn.Identity = ""
	builtIn.Patch.Schedules = &schedules
	builtIn, err = configuration.NormalizeSource(builtIn)
	if err != nil {
		return configuration.Effective{}, err
	}
	sources := []configuration.Source{builtIn}
	if found {
		sources = append(sources, source)
	}
	if temporary != nil {
		sources = append(sources, *temporary)
	}
	return configuration.Resolve(sources)
}

func newLocalSource(effective configuration.Effective) configuration.Source {
	locale, timezone := effective.Values.Locale, effective.Values.TimeZone
	result, _ := configuration.NormalizeSource(configuration.Source{SchemaName: configuration.SourceSchema, SchemaVersion: configuration.SchemaVersion, ModelVersion: configuration.ModelVersion, ID: "local.operator", SourceVersion: "1.0", Kind: configuration.PrimaryLocal, Patch: configuration.Patch{Locale: &locale, TimeZone: &timezone}})
	return result
}

func setConfigValue(source configuration.Source, found bool, effective configuration.Effective, key, value string) (configuration.Source, error) {
	if !found {
		source = newLocalSource(effective)
	}
	source.Identity = ""
	switch key {
	case "locale":
		source.Patch.Locale = &value
	case "time_zone":
		source.Patch.TimeZone = &value
	case "snapshot_retention":
		n, err := strconv.Atoi(value)
		if err != nil {
			return source, fmt.Errorf("snapshot_retention must be an integer")
		}
		source.Patch.SnapshotRetention = &n
	case "guardian.interval", "guardian.cycle_timeout":
		duration, err := time.ParseDuration(value)
		if err != nil {
			return source, fmt.Errorf("%s must be a duration", key)
		}
		schedule := defaultGuardianSchedule()
		schedule.IntervalNS, schedule.ExecutionTimeoutNS = int64(guardianInterval(effective)), int64(guardianTimeout(effective))
		if key == "guardian.interval" {
			schedule.IntervalNS = int64(duration)
		} else {
			schedule.ExecutionTimeoutNS = int64(duration)
		}
		schedules := []configuration.Schedule{schedule}
		source.Patch.Schedules = &schedules
	default:
		return source, fmt.Errorf("unsupported configuration key %q", key)
	}
	return configuration.NormalizeSource(source)
}

func getConfigValue(effective configuration.Effective, key string) (string, error) {
	switch key {
	case "locale":
		return effective.Values.Locale, nil
	case "time_zone":
		return effective.Values.TimeZone, nil
	case "snapshot_retention":
		return strconv.Itoa(effective.Values.SnapshotRetention), nil
	case "guardian.interval":
		return guardianInterval(effective).String(), nil
	case "guardian.cycle_timeout":
		return guardianTimeout(effective).String(), nil
	default:
		return "", fmt.Errorf("unsupported configuration key %q", key)
	}
}

func defaultGuardianSchedule() configuration.Schedule {
	return configuration.Schedule{ID: "guardian.observe", ContractVersion: configuration.ScheduleVersion, Enabled: true, TimeZone: "UTC", Trigger: configuration.IntervalTrigger, IntervalNS: int64(guardianDefaultInterval), Calendar: configuration.Calendar{Minutes: []int{}, Hours: []int{}, MonthDays: []int{}, Months: []int{}, Weekdays: []int{}}, DSTPolicy: configuration.DSTFirstOccurrence, MisfirePolicy: configuration.MisfireRunOnce, OverlapPolicy: configuration.OverlapForbid, ExecutionTimeoutNS: int64(guardianDefaultTimeout), RetryPolicyID: "canonical.default", CheckIDs: []string{}, CommandProfile: "observe"}
}

func guardianInterval(effective configuration.Effective) time.Duration {
	for _, schedule := range effective.Values.Schedules {
		if schedule.ID == "guardian.observe" {
			return time.Duration(schedule.IntervalNS)
		}
	}
	return guardianDefaultInterval
}
func guardianTimeout(effective configuration.Effective) time.Duration {
	for _, schedule := range effective.Values.Schedules {
		if schedule.ID == "guardian.observe" {
			return time.Duration(schedule.ExecutionTimeoutNS)
		}
	}
	return guardianDefaultTimeout
}

func validateGuardianTiming(effective configuration.Effective) error {
	interval, timeout := guardianInterval(effective), guardianTimeout(effective)
	if interval <= 0 || interval > runtimeservice.MaxInterval || timeout <= 0 || timeout >= interval || timeout > runtimeservice.MaxCycleTimeout {
		return fmt.Errorf("%w: guardian interval and timeout are inconsistent", configurationstore.ErrInvalid)
	}
	return nil
}

func temporaryGuardianSource(base configuration.Effective, interval, timeout *time.Duration) (*configuration.Source, error) {
	if interval == nil && timeout == nil {
		return nil, nil
	}
	schedule := defaultGuardianSchedule()
	schedule.IntervalNS, schedule.ExecutionTimeoutNS = int64(guardianInterval(base)), int64(guardianTimeout(base))
	if interval != nil {
		schedule.IntervalNS = int64(*interval)
	}
	if timeout != nil {
		schedule.ExecutionTimeoutNS = int64(*timeout)
	}
	schedules := []configuration.Schedule{schedule}
	source, err := configuration.NormalizeSource(configuration.Source{SchemaName: configuration.SourceSchema, SchemaVersion: configuration.SchemaVersion, ModelVersion: configuration.ModelVersion, ID: "command.guardian", SourceVersion: "1.0", Kind: configuration.TemporaryOverride, Patch: configuration.Patch{Schedules: &schedules}})
	return &source, err
}

func configFailure(errout io.Writer, err error) int {
	diagnostic := "configuration_failed"
	switch {
	case errors.Is(err, configurationstore.ErrUnavailable):
		diagnostic = "configuration_unavailable"
	case errors.Is(err, configurationstore.ErrUnsafe):
		diagnostic = "configuration_path_unsafe"
	case errors.Is(err, configurationstore.ErrPermission):
		diagnostic = "configuration_permission_unsafe"
	case errors.Is(err, configurationstore.ErrInvalid):
		diagnostic = "configuration_invalid"
	}
	fmt.Fprintf(errout, "configuration operation failed: %s\n", diagnostic)
	return 1
}

func writeConfigHelp(out io.Writer) {
	fmt.Fprintln(out, "Usage: qwsg config <show|validate|get|set> [KEY [VALUE]] [--config FILE] [--format human|json]")
}
