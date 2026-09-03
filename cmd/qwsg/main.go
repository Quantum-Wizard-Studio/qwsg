package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"reflect"
	"strconv"
	"strings"
	"sync"
	"syscall"
	"time"
	"unicode"

	"quantumwizard.hu/qwsg/internal/alert"
	"quantumwizard.hu/qwsg/internal/app"
	"quantumwizard.hu/qwsg/internal/collector"
	canonicalcommand "quantumwizard.hu/qwsg/internal/command"
	"quantumwizard.hu/qwsg/internal/comparison"
	"quantumwizard.hu/qwsg/internal/configuration"
	"quantumwizard.hu/qwsg/internal/configurationstore"
	"quantumwizard.hu/qwsg/internal/credentialstore"
	"quantumwizard.hu/qwsg/internal/guardian"
	"quantumwizard.hu/qwsg/internal/inventory"
	"quantumwizard.hu/qwsg/internal/inventorystore"
	"quantumwizard.hu/qwsg/internal/notification"
	"quantumwizard.hu/qwsg/internal/operatorconsole"
	"quantumwizard.hu/qwsg/internal/operatorstate"
	"quantumwizard.hu/qwsg/internal/pipeline"
	"quantumwizard.hu/qwsg/internal/presentation"
	"quantumwizard.hu/qwsg/internal/presentationmodel"
	"quantumwizard.hu/qwsg/internal/runner"
	"quantumwizard.hu/qwsg/internal/runtime"
	"quantumwizard.hu/qwsg/internal/runtimeservice"
	"quantumwizard.hu/qwsg/internal/scheduler"
	"quantumwizard.hu/qwsg/internal/smtpnotification"
	"quantumwizard.hu/qwsg/internal/updateawareness"
)

var (
	version     = "1.0.0"
	buildCommit = "unknown"
	buildDate   = "unknown"
)

const (
	formatJSON  = "json"
	formatHuman = "human"
)

type storeOptions struct {
	storePath    string
	snapshotName string
	format       string
	retention    int
}

type compareOptions struct {
	storePath string
	from      string
	to        string
	format    string
	retention int
}

type snapshotSummary struct {
	Name          string           `json:"name"`
	Status        inventory.Status `json:"status"`
	ObservedAt    time.Time        `json:"observed_at"`
	CompletedAt   time.Time        `json:"completed_at"`
	SchemaVersion string           `json:"schema_version"`
	SchemaName    string           `json:"canonical_schema"`
	CategoryCount int              `json:"category_count"`
	IssueCount    int              `json:"issue_count"`
	Integrity     string           `json:"integrity"`
}

func main() {
	interactive := terminalFile(os.Stdin) && terminalFile(os.Stdout)
	os.Exit(runWithConsole(os.Args[1:], os.Stdin, os.Stdout, os.Stderr, interactive))
}

func run(args []string, out, errout io.Writer) int {
	return runWithConsole(args, strings.NewReader(""), out, errout, false)
}

func runWithConsole(args []string, in io.Reader, out, errout io.Writer, interactive bool) int {
	if len(args) == 0 {
		return runConsole(in, out, errout, interactive)
	}
	switch args[0] {
	case "help", "--help", "-h":
		return runHelp(args[1:], out, errout)
	case "version":
		if len(args) == 2 && isHelp(args[1]) {
			writeVersionHelp(out)
			return 0
		}
		if len(args) != 1 {
			return usageError(errout, "version does not accept options")
		}
		fmt.Fprintf(out, "QWSG %s\ncommit: %s\nbuilt: %s\n", safeText(version), safeText(buildCommit), safeText(buildDate))
		return 0
	case "inventory":
		return runInventory(args[1:], out, errout)
	case "compare":
		return runCompare(args[1:], out, errout)
	case "status", "check", "observe", "changes", "health", "report":
		return runCanonicalProfile(args[0], args[1:], out, errout)
	case "analyze":
		return runCanonicalAdvanced(args[1:], out, errout)
	case "console":
		if len(args) == 2 && isHelp(args[1]) {
			fmt.Fprintln(out, "Usage: qwsg console\n\nOpens the read-only local Operator Console. Refresh is explicit (r).")
			return 0
		}
		if len(args) != 1 {
			return usageError(errout, "console does not accept options")
		}
		return runConsole(in, out, errout, interactive)
	case "guardian":
		return runGuardian(args[1:], out, errout)
	case "setup":
		return runSetup(args[1:], in, out, errout, interactive)
	case "config":
		return runConfig(args[1:], out, errout)
	case "notification":
		return runNotification(args[1:], out, errout)
	case "install":
		return runInstall(args[1:], in, out, errout, interactive)
	case "readiness":
		return runReadiness(args[1:], out, errout)
	case "update":
		return runUpdate(args[1:], out, errout)
	default:
		return usageError(errout, "unknown command: %s", safeText(args[0]))
	}
}

const (
	guardianDefaultInterval = 5 * time.Minute
	guardianDefaultTimeout  = 2 * time.Minute
)

type guardianOptions struct {
	stateRoot    string
	storeRoot    string
	configSource string
	interval     *time.Duration
	timeout      *time.Duration
	generation   string
}

type guardianPipeline struct {
	orchestrator *pipeline.Orchestrator
	storeRoot    string
}

func (v guardianPipeline) Execute(ctx context.Context, definition canonicalcommand.Definition) (canonicalcommand.Execution, error) {
	empty, err := observationStoreEmpty(v.storeRoot)
	if err != nil {
		return canonicalcommand.Execution{}, err
	}
	if empty && definition.Profile == "observe" {
		check, resolveErr := canonicalcommand.ResolveProfile("check", canonicalcommand.Selection{Source: "live", Store: v.storeRoot})
		if resolveErr != nil {
			return canonicalcommand.Execution{}, resolveErr
		}
		if _, executeErr := v.orchestrator.Execute(ctx, check); executeErr != nil {
			return canonicalcommand.Execution{}, executeErr
		}
		return canonicalcommand.Execution{}, fmt.Errorf("observation baseline initialized")
	}
	return v.orchestrator.Execute(ctx, definition)
}

func runGuardian(args []string, out, errout io.Writer) int {
	if len(args) == 0 || isHelp(args[0]) {
		fmt.Fprintln(out, "Usage: qwsg guardian run [--state-dir DIR] [--store DIR] [--config FILE] [--interval DURATION] [--cycle-timeout DURATION]")
		return 0
	}
	if args[0] == "report-exit" {
		return runGuardianExit(args[1:], errout)
	}
	if args[0] != "run" {
		return usageError(errout, "unknown guardian operation: %s", safeText(args[0]))
	}
	options, err := parseGuardianOptions(args[1:])
	if err != nil {
		return usageError(errout, "%v", err)
	}
	if err = executeGuardian(options); err != nil {
		diagnostic := "guardian_failed"
		if errors.Is(err, guardian.ErrActive) {
			diagnostic = "guardian_active"
		} else if errors.Is(err, guardian.ErrCheckpoint) || errors.Is(err, guardian.ErrIncompatible) {
			diagnostic = "guardian_checkpoint_invalid"
		} else if errors.Is(err, guardian.ErrUnsafePath) {
			diagnostic = "guardian_state_unsafe"
		} else if errors.Is(err, configurationstore.ErrUnsafe) || errors.Is(err, configurationstore.ErrPermission) || errors.Is(err, configurationstore.ErrInvalid) || errors.Is(err, configurationstore.ErrUnavailable) {
			diagnostic = "guardian_configuration_invalid"
		}
		fmt.Fprintf(errout, "guardian operation failed: %s\n", diagnostic)
		return 1
	}
	return 0
}

func runGuardianExit(args []string, errout io.Writer) int {
	root, err := localStateRoot()
	generation, result := "", ""
	for index := 0; err == nil && index < len(args); index++ {
		if index+1 >= len(args) {
			err = fmt.Errorf("exit report option requires a value")
			break
		}
		name, value := args[index], args[index+1]
		index++
		switch name {
		case "--state-dir":
			root = value
		case "--generation":
			generation = value
		case "--result":
			result = value
		default:
			err = fmt.Errorf("unknown exit report option")
		}
	}
	if err == nil {
		var checkpoints *guardian.Store
		checkpoints, err = guardian.OpenStore(filepath.Join(root, "guardian"))
		if err != nil {
			err = fmt.Errorf("%w", guardian.ErrExitCheckpoint)
		}
		if err == nil {
			var current *operatorstate.Store
			current, err = operatorstate.Open(root)
			if err != nil {
				err = fmt.Errorf("%w", guardian.ErrExitCurrent)
			}
			if err == nil {
				err = guardian.ReportExit(checkpoints, current, generation, result, time.Now().UTC(), 2*guardianDefaultInterval)
			}
		}
	}
	if err != nil {
		diagnostic := "exit_operation_failed"
		if errors.Is(err, guardian.ErrCheckpoint) || errors.Is(err, guardian.ErrIncompatible) {
			diagnostic = "exit_checkpoint_invalid"
		} else if errors.Is(err, guardian.ErrExitEvidence) {
			diagnostic = "exit_evidence_invalid"
		} else if errors.Is(err, guardian.ErrExitState) {
			diagnostic = "exit_projection_invalid"
		} else if errors.Is(err, guardian.ErrExitCheckpoint) {
			diagnostic = "exit_checkpoint_unavailable"
		} else if errors.Is(err, guardian.ErrExitCurrent) {
			diagnostic = "exit_current_state_unavailable"
		} else if errors.Is(err, operatorstate.ErrCorrupt) || errors.Is(err, operatorstate.ErrIncompatible) || errors.Is(err, operatorstate.ErrPermission) || errors.Is(err, operatorstate.ErrUnsafePath) {
			diagnostic = "exit_current_state_invalid"
		}
		fmt.Fprintf(errout, "guardian exit report failed: %s\n", diagnostic)
		return 1
	}
	return 0
}

func parseGuardianOptions(args []string) (guardianOptions, error) {
	stateRoot, err := localStateRoot()
	if err != nil {
		return guardianOptions{}, err
	}
	value := guardianOptions{stateRoot: stateRoot, generation: os.Getenv("INVOCATION_ID")}
	if value.generation == "" {
		value.generation = fmt.Sprintf("manual-%d-%d", os.Getpid(), time.Now().UTC().UnixNano())
	}
	for index := 0; index < len(args); index++ {
		if index+1 >= len(args) {
			return guardianOptions{}, fmt.Errorf("guardian option requires a value")
		}
		name, item := args[index], args[index+1]
		index++
		switch name {
		case "--state-dir":
			value.stateRoot = item
		case "--store":
			value.storeRoot = item
		case "--config", "--config-source":
			if value.configSource != "" {
				return guardianOptions{}, fmt.Errorf("configuration path may be selected once")
			}
			value.configSource = item
		case "--interval":
			parsed, parseErr := time.ParseDuration(item)
			value.interval, err = &parsed, parseErr
		case "--cycle-timeout":
			parsed, parseErr := time.ParseDuration(item)
			value.timeout, err = &parsed, parseErr
		default:
			return guardianOptions{}, fmt.Errorf("unknown guardian option: %s", safeText(name))
		}
		if err != nil {
			return guardianOptions{}, fmt.Errorf("invalid guardian duration")
		}
	}
	if value.storeRoot == "" {
		value.storeRoot = filepath.Join(value.stateRoot, "inventory")
	}
	if !filepath.IsAbs(value.stateRoot) || !filepath.IsAbs(value.storeRoot) || (value.interval != nil && (*value.interval <= 0 || *value.interval > runtimeservice.MaxInterval)) || (value.timeout != nil && (*value.timeout <= 0 || *value.timeout > runtimeservice.MaxCycleTimeout)) {
		return guardianOptions{}, fmt.Errorf("invalid guardian operating bounds")
	}
	return value, nil
}

func executeGuardian(options guardianOptions) error {
	effective, interval, timeout, err := guardianConfiguration(options)
	if err != nil {
		return err
	}
	if interval <= 0 || interval > runtimeservice.MaxInterval || timeout <= 0 || timeout >= interval || timeout > runtimeservice.MaxCycleTimeout {
		return fmt.Errorf("%w: guardian operating bounds", configurationstore.ErrInvalid)
	}
	configPath, err := configurationstore.SelectPath(options.configSource, os.Getenv)
	if err != nil {
		return err
	}
	emailConfig, err := smtpnotification.FromEffective(effective)
	if err != nil {
		return fmt.Errorf("%w: notification configuration", configurationstore.ErrInvalid)
	}
	password := []byte(nil)
	credentialAvailable := emailConfig.Auth != "password"
	if emailConfig.Enabled && emailConfig.Auth == "password" {
		password, err = credentialstore.Load(configPath, emailConfig.CredentialRef)
		credentialAvailable = err == nil
	}
	if !smtpnotification.Ready(smtpnotification.Preflight(emailConfig, credentialAvailable)) {
		return fmt.Errorf("%w: notification preflight", configurationstore.ErrInvalid)
	}
	guardianRoot := filepath.Join(options.stateRoot, "guardian")
	lock, err := guardian.Acquire(guardianRoot, options.generation)
	if err != nil {
		return err
	}
	defer lock.Release()
	checkpointStore, err := guardian.OpenStore(guardianRoot)
	if err != nil {
		return err
	}
	previous, loadErr := checkpointStore.Load()
	if loadErr != nil && !errors.Is(loadErr, os.ErrNotExist) {
		return loadErr
	}
	checkpoint := guardian.Checkpoint{SchemaName: "qwsg.guardian-checkpoint", SchemaVersion: guardian.SchemaVersion, ModelVersion: guardian.ModelVersion, ServiceID: "qwsg.guardian.local", ConfigurationID: effective.ID, Generation: options.generation, Active: true, RuntimeState: runtime.NewState(), AlertState: alert.NewState(effective.ID), NotificationQueueState: notification.NewQueueState()}
	if loadErr == nil {
		if previous.ServiceID != checkpoint.ServiceID {
			return guardian.ErrCheckpoint
		}
		if previous.ConfigurationID == effective.ID {
			checkpoint.RuntimeState, checkpoint.AlertState, checkpoint.NotificationQueueState = previous.RuntimeState, previous.AlertState, previous.NotificationQueueState
		}
	}
	if err = checkpointStore.Save(checkpoint); err != nil {
		return err
	}
	schedulerStore, err := scheduler.OpenFileStore(filepath.Join(guardianRoot, "scheduler"))
	if err != nil {
		return err
	}
	schedulerLocker, err := scheduler.NewFileLocker(filepath.Join(guardianRoot, "scheduler"))
	if err != nil {
		return err
	}
	orchestrator := &pipeline.Orchestrator{Collect: collectInventoryContext, Configuration: &effective}
	cycle := scheduler.Cycle{Configuration: effective, Selection: canonicalcommand.Selection{Source: "live", Store: options.storeRoot}, LockOwnerID: "qwsg.guardian.cycle", Store: schedulerStore, Locker: schedulerLocker, Clock: guardian.NewSchedulerClock(options.generation), TimeZones: scheduler.SystemTimeZones{}, ResolveCommand: scheduler.ResolveCanonicalCommand, Pipeline: guardianPipeline{orchestrator: orchestrator, storeRoot: options.storeRoot}}
	captured := &guardian.CapturingScheduler{Cycle: cycle}
	providers := []notification.Provider{}
	if emailConfig.Enabled {
		providers = append(providers, smtpnotification.Provider{Config: emailConfig, Password: password})
	}
	registry, err := notification.NewRegistry(providers...)
	if err != nil {
		return err
	}
	policy, err := smtpnotification.Policy(emailConfig)
	if err != nil {
		return err
	}
	runner := &guardian.RuntimeRunner{Coordinator: runtime.Coordinator{Scheduler: captured, Clock: runtimeservice.SystemClock{}, Providers: registry}, Store: checkpointStore, Checkpoint: &checkpoint}
	stateStore, err := operatorstate.Open(options.stateRoot)
	if err != nil {
		return err
	}
	publisher := &guardian.Publisher{Store: stateStore, Scheduler: captured, Runner: runner, DefinitionID: effective.ID, ApplicationVersion: version, FreshFor: 2 * interval}
	definition, err := runtimeservice.NewDefinition(checkpoint.ServiceID, interval, timeout)
	if err != nil {
		return err
	}
	started := time.Now().UTC()
	seedContext, err := runtime.NewExecutionContext("guardian-seed", checkpoint.ServiceID, started, started.Add(timeout))
	if err != nil {
		return err
	}
	seed := runtime.Input{Context: seedContext, Configuration: effective, PreviousState: checkpoint.RuntimeState, PreviousAlertState: checkpoint.AlertState, PreviousNotificationQueue: checkpoint.NotificationQueueState, AlertEvidenceTTLNS: int64(24 * time.Hour), Acknowledgements: []alert.Acknowledgement{}, Suppressions: []alert.SuppressionWindow{}, NotificationPolicy: policy}
	ready := make(chan struct{})
	startupSink := &guardian.StartupSink{Next: guardian.Sink{Publisher: publisher}, Ready: ready}
	service := runtimeservice.Service{Clock: runtimeservice.SystemClock{}, Waiter: runtimeservice.TimerWaiter{}, Runner: runner, Sink: startupSink}
	awarenessStore, awarenessErr := updateawareness.Open(options.stateRoot)
	if awarenessErr != nil {
		return awarenessErr
	}
	awarenessManager := updateAwarenessManager(awarenessStore)
	signalContext, stopSignals := runtimeservice.OSSignalContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	var releaseChecks sync.WaitGroup
	releaseChecks.Add(1)
	go func() {
		defer releaseChecks.Done()
		(guardian.ReleaseCheckService{Store: awarenessStore, Ready: ready, Interval: guardian.DefaultReleaseCheckInterval, Timeout: guardian.DefaultReleaseCheckTimeout, Check: func(ctx context.Context) error {
			_, checkErr := awarenessManager.Check(ctx)
			return checkErr
		}}).Run(signalContext)
	}()
	result, err := service.Run(signalContext, runtimeservice.Input{Definition: definition, StartedAt: started, InitialState: runtimeservice.NewState(checkpoint.ServiceID), Seed: seed})
	stopSignals()
	releaseChecks.Wait()
	checkpoint.Active = false
	_ = checkpointStore.Save(checkpoint)
	if err != nil || result.FinalState.Lifecycle == runtimeservice.Failed {
		return fmt.Errorf("guardian service terminated")
	}
	return nil
}

func guardianConfiguration(options guardianOptions) (configuration.Effective, time.Duration, time.Duration, error) {
	path, err := configurationstore.SelectPath(options.configSource, os.Getenv)
	if err != nil {
		return configuration.Effective{}, 0, 0, err
	}
	user, found, err := configurationstore.Load(path)
	if err != nil {
		return configuration.Effective{}, 0, 0, err
	}
	if options.configSource != "" && !found {
		return configuration.Effective{}, 0, 0, configurationstore.ErrUnavailable
	}
	base, err := resolveLocalConfiguration(user, found, nil)
	if err != nil {
		return configuration.Effective{}, 0, 0, fmt.Errorf("%w: effective configuration", configurationstore.ErrInvalid)
	}
	temporary, err := temporaryGuardianSource(base, options.interval, options.timeout)
	if err != nil {
		return configuration.Effective{}, 0, 0, err
	}
	effective, err := resolveLocalConfiguration(user, found, temporary)
	if err != nil {
		return configuration.Effective{}, 0, 0, fmt.Errorf("%w: effective configuration", configurationstore.ErrInvalid)
	}
	return effective, guardianInterval(effective), guardianTimeout(effective), nil
}

type localOverviewProvider struct{}

func (localOverviewProvider) Refresh(ctx context.Context) (presentationmodel.Overview, error) {
	if err := ctx.Err(); err != nil {
		return presentationmodel.Overview{}, err
	}
	return loadCurrentOverview(time.Now().UTC())
}

func loadCurrentOverview(now time.Time) (presentationmodel.Overview, error) {
	store, err := currentOperatorStore()
	if err != nil {
		return presentationmodel.Overview{}, err
	}
	current, err := store.Load()
	if err != nil {
		return presentationmodel.Overview{}, err
	}
	overview, err := presentationmodel.RequalifyFreshness(current.Overview, now, current.FreshUntil)
	if err != nil {
		return presentationmodel.Overview{}, operatorstate.ErrCorrupt
	}
	return overview, nil
}

func runConsole(in io.Reader, out, errout io.Writer, interactive bool) int {
	path, configErr := configurationstore.DefaultPath(os.Getenv)
	if configErr == nil {
		_, _, configErr = configurationstore.Load(path)
	}
	if configErr != nil {
		return configFailure(errout, configErr)
	}
	now := time.Unix(0, 0).UTC()
	overview, err := presentationmodel.Project(presentationmodel.Input{SchemaName: presentationmodel.InputSchema, SchemaVersion: presentationmodel.SchemaVersion, ObservedAt: now, FreshForNS: int64(time.Hour)})
	if err != nil {
		fmt.Fprintf(errout, "console initialization failed: %s\n", safeText(err.Error()))
		return 1
	}
	locale := os.Getenv("QWSG_LOCALE")
	if locale != "hu" {
		locale = "en"
	}
	diagnostic := ""
	if qualified, loadErr := loadCurrentOverview(time.Now().UTC()); loadErr == nil {
		overview = qualified
	} else if !errors.Is(loadErr, operatorstate.ErrMissing) {
		diagnostic = stateDiagnostic(loadErr)
	}
	capabilities := operatorconsole.Capabilities{Interactive: interactive, Width: 80, Height: 30}
	state, err := operatorconsole.NewState(overview, locale, capabilities)
	if err != nil {
		fmt.Fprintf(errout, "console initialization failed: %s\n", safeText(err.Error()))
		return 1
	}
	state.Diagnostic = diagnostic
	if !interactive {
		text := operatorconsole.Render(state)
		if _, err = io.WriteString(out, text); err != nil {
			fmt.Fprintf(errout, "console output failed: %s\n", safeText(err.Error()))
			return 1
		}
		return 0
	}
	if err = operatorconsole.Run(context.Background(), in, out, localOverviewProvider{}, state); err != nil {
		fmt.Fprintf(errout, "console failed: %s\n", safeText(err.Error()))
		return 1
	}
	return 0
}

func terminalFile(file *os.File) bool {
	info, err := file.Stat()
	return err == nil && info.Mode()&os.ModeCharDevice != 0
}

func runCanonicalProfile(name string, args []string, out, errout io.Writer) int {
	if len(args) > 0 && isHelp(args[0]) {
		writeCanonicalHelp(out, name)
		return 0
	}
	selection := canonicalcommand.Selection{Store: os.Getenv("QWSG_STORE")}
	if name == "observe" && selection.Store == "" {
		var err error
		selection.Store, err = observationStoreRoot()
		if err != nil {
			fmt.Fprintf(errout, "operator observation failed: %s\n", observeDiagnostic(err))
			return 1
		}
	}
	definition, err := canonicalcommand.ParseProfile(name, args, selection)
	if err != nil {
		return usageError(errout, "%v", err)
	}
	if name == "observe" {
		return executeObserve(definition, out, errout)
	}
	return executeCanonical(definition, out, errout)
}

type observationResult struct {
	Definition canonicalcommand.Definition
	Execution  canonicalcommand.Execution
	Overview   presentationmodel.Overview
	Bootstrap  bool
}

type observationFailureKind string

const (
	observationBootstrapFailure   observationFailureKind = "state_bootstrap_failed"
	observationPipelineFailure    observationFailureKind = "evaluation_pipeline_failed"
	observationProjectionFailure  observationFailureKind = "operator_projection_failed"
	observationPublicationFailure observationFailureKind = "current_state_publication_failed"
)

type observationFailure struct {
	kind observationFailureKind
	err  error
}

func (failure *observationFailure) Error() string { return string(failure.kind) }
func (failure *observationFailure) Unwrap() error { return failure.err }

func classifiedObservationFailure(kind observationFailureKind, err error) error {
	if err == nil {
		return nil
	}
	return &observationFailure{kind: kind, err: err}
}

func executeObserve(definition canonicalcommand.Definition, out, errout io.Writer) int {
	lock, lockErr := acquireOneShotLock()
	if lockErr != nil {
		diagnostic := "evaluation_failed"
		if errors.Is(lockErr, guardian.ErrActive) {
			diagnostic = "guardian_active"
		}
		fmt.Fprintf(errout, "operator observation failed: %s\n", diagnostic)
		return 1
	}
	defer lock.Release()
	result, err := observeOnce(context.Background(), definition, collectInventoryContext, func() time.Time { return time.Now().UTC() })
	if err != nil {
		fmt.Fprintf(errout, "operator observation failed: %s\n", observeDiagnostic(err))
		return 1
	}
	if err = renderCanonical(result.Definition, result.Execution, out); err != nil {
		fmt.Fprintf(errout, "command presentation failed: %s\n", safeText(err.Error()))
		return 1
	}
	if result.Bootstrap {
		writeBaselineGuidance(definition.Parameters.Output, out, errout)
	}
	return 0
}

func acquireOneShotLock() (*guardian.Lock, error) {
	root, err := localStateRoot()
	if err != nil {
		return nil, err
	}
	return guardian.Acquire(filepath.Join(root, "guardian"), "qwsg.observe.once")
}

func observeOnce(ctx context.Context, definition canonicalcommand.Definition, collect pipeline.Collector, clock func() time.Time) (observationResult, error) {
	if err := ctx.Err(); err != nil {
		return observationResult{}, err
	}
	if clock == nil {
		return observationResult{}, fmt.Errorf("observation clock is unavailable")
	}
	plan, err := canonicalcommand.PlanDefinition(definition)
	if err != nil || definition.Profile != "observe" || definition.Selection.Source != "live" || definition.Selection.Store == "" || !reflect.DeepEqual(plan.Stages, canonicalcommand.CanonicalStages) {
		return observationResult{}, fmt.Errorf("invalid observe definition")
	}
	bootstrap, err := observationStoreEmpty(definition.Selection.Store)
	if err != nil {
		return observationResult{}, err
	}
	orchestrator := pipeline.Orchestrator{Collect: collect, Retention: inventorystore.DefaultRetention, Rules: pipeline.CanonicalObservationRules()}
	if bootstrap {
		check, resolveErr := canonicalcommand.ResolveProfile("check", canonicalcommand.Selection{Source: "live", Store: definition.Selection.Store})
		if resolveErr != nil {
			return observationResult{}, resolveErr
		}
		check.Parameters = definition.Parameters
		check.ID = ""
		check, resolveErr = canonicalcommand.Normalize(check)
		if resolveErr != nil {
			return observationResult{}, resolveErr
		}
		execution, executeErr := orchestrator.Execute(ctx, check)
		if executeErr != nil {
			return observationResult{}, classifiedObservationFailure(observationPipelineFailure, executeErr)
		}
		overview, publishErr := publishObserveBootstrap(check, execution, clock().UTC())
		return observationResult{Definition: check, Execution: execution, Overview: overview, Bootstrap: true}, publishErr
	}
	execution, err := orchestrator.Execute(ctx, definition)
	if err != nil {
		return observationResult{}, classifiedObservationFailure(observationPipelineFailure, err)
	}
	overview, err := publishObserve(definition, execution, clock().UTC())
	return observationResult{Definition: definition, Execution: execution, Overview: overview}, err
}

func publishObserveBootstrap(definition canonicalcommand.Definition, execution canonicalcommand.Execution, publishedAt time.Time) (presentationmodel.Overview, error) {
	plan, err := canonicalcommand.PlanDefinition(definition)
	if err != nil || definition.Profile != "check" || definition.Selection.Source != "live" || !execution.Complete || execution.CommandID != definition.ID || execution.PlanID != plan.ID || len(execution.Stages) != 2 {
		return presentationmodel.Overview{}, classifiedObservationFailure(observationProjectionFailure, fmt.Errorf("ineligible observe bootstrap"))
	}
	inventoryStage, snapshotStage := execution.Stages[0], execution.Stages[1]
	if inventoryStage.Stage != canonicalcommand.Inventory || snapshotStage.Stage != canonicalcommand.Snapshot || inventoryStage.ContractName != inventory.CanonicalSchemaName || inventoryStage.Version != inventory.SchemaVersion || snapshotStage.ContractName != inventorystore.FormatName || snapshotStage.Version != inventorystore.FormatVersion {
		return presentationmodel.Overview{}, classifiedObservationFailure(observationProjectionFailure, fmt.Errorf("invalid observe bootstrap coverage"))
	}
	observed, ok := inventoryStage.Value.(inventory.Snapshot)
	if !ok || inventory.Validate(observed) != nil {
		return presentationmodel.Overview{}, classifiedObservationFailure(observationProjectionFailure, fmt.Errorf("invalid typed bootstrap inventory"))
	}
	snapshotted, ok := snapshotStage.Value.(inventory.Snapshot)
	if !ok || inventory.Validate(snapshotted) != nil || !reflect.DeepEqual(observed, snapshotted) {
		return presentationmodel.Overview{}, classifiedObservationFailure(observationProjectionFailure, fmt.Errorf("invalid typed bootstrap snapshot"))
	}
	observationTime, freshUntil := observed.CompletedAt.UTC(), observed.FreshUntil.UTC()
	if !freshUntil.After(observationTime) || publishedAt.Before(observationTime) {
		return presentationmodel.Overview{}, classifiedObservationFailure(observationProjectionFailure, fmt.Errorf("invalid bootstrap freshness"))
	}
	overview, err := presentationmodel.Project(presentationmodel.Input{SchemaName: presentationmodel.InputSchema, SchemaVersion: presentationmodel.SchemaVersion, ObservedAt: observationTime, FreshForNS: int64(freshUntil.Sub(observationTime)), Command: &presentationmodel.CommandObservation{ObservedAt: observationTime, Value: execution}, Inventory: &presentationmodel.InventoryObservation{ObservedAt: observationTime, Value: observed}})
	if err != nil {
		return presentationmodel.Overview{}, classifiedObservationFailure(observationProjectionFailure, err)
	}
	current, err := operatorstate.Normalize(operatorstate.State{ObservedAt: observationTime, PublishedAt: publishedAt.UTC(), FreshUntil: freshUntil, Coverage: operatorstate.CoverageInventorySnapshot, Provenance: operatorstate.Provenance{DefinitionID: definition.ID, ExecutionID: execution.ID, Profile: "check", Source: "live", Stages: []string{"inventory", "snapshot"}, Reason: operatorstate.PublicationCheck, ApplicationVersion: version}, Overview: overview})
	if err != nil {
		return presentationmodel.Overview{}, classifiedObservationFailure(observationPublicationFailure, err)
	}
	store, err := currentOperatorStore()
	if err != nil {
		return presentationmodel.Overview{}, classifiedObservationFailure(observationPublicationFailure, err)
	}
	if err = store.Publish(current); err != nil {
		return presentationmodel.Overview{}, classifiedObservationFailure(observationPublicationFailure, err)
	}
	return overview, nil
}

func observationStoreEmpty(root string) (bool, error) {
	if root == "" {
		return false, fmt.Errorf("observation store is unavailable")
	}
	if _, err := os.Lstat(root); errors.Is(err, os.ErrNotExist) {
		return true, nil
	} else if err != nil {
		return false, err
	}
	store, err := inventorystore.Open(root, inventorystore.DefaultRetention)
	if err != nil {
		return false, err
	}
	names, err := store.List()
	return len(names) == 0, err
}

func renderCanonical(definition canonicalcommand.Definition, execution canonicalcommand.Execution, out io.Writer) error {
	var document []byte
	var err error
	if definition.Parameters.Output == canonicalcommand.JSON {
		document, err = presentation.JSON(execution)
	} else {
		var value string
		value, err = presentation.Terminal(execution)
		document = []byte(value)
	}
	if err != nil {
		return err
	}
	_, err = out.Write(document)
	return err
}

func writeBaselineGuidance(format canonicalcommand.OutputFormat, out, errout io.Writer) {
	message := "Observation baseline established; condition remains unknown until a later 'qwsg observe'."
	if os.Getenv("QWSG_LOCALE") == "hu" {
		message = "A megfigyelési alapállapot elkészült; az állapot egy későbbi 'qwsg observe' futásig ismeretlen marad."
	}
	if format == canonicalcommand.JSON {
		fmt.Fprintln(errout, message)
	} else {
		fmt.Fprintln(out, message)
	}
}

func observeDiagnostic(err error) string {
	switch {
	case errors.Is(err, inventorystore.ErrCorrupt):
		return "inventory_store_corrupt"
	case errors.Is(err, inventorystore.ErrIncompatible):
		return "inventory_store_incompatible"
	case errors.Is(err, inventorystore.ErrUnsafePath):
		return "inventory_store_unsafe"
	case errors.Is(err, os.ErrPermission):
		return "inventory_store_permission_denied"
	case errors.Is(err, operatorstate.ErrCorrupt), errors.Is(err, operatorstate.ErrIncompatible), errors.Is(err, operatorstate.ErrPermission), errors.Is(err, operatorstate.ErrUnsafePath):
		return stateDiagnostic(err)
	}
	var failure *observationFailure
	if errors.As(err, &failure) {
		return string(failure.kind)
	}
	return "evaluation_failed"
}

func runCanonicalAdvanced(args []string, out, errout io.Writer) int {
	if len(args) > 0 && isHelp(args[0]) {
		writeCanonicalHelp(out, "analyze")
		return 0
	}
	definition, err := canonicalcommand.Parse(args)
	if err != nil {
		return usageError(errout, "%v", err)
	}
	return executeCanonical(definition, out, errout)
}

func executeCanonical(definition canonicalcommand.Definition, out, errout io.Writer) int {
	orchestrator := pipeline.Orchestrator{
		Collect: func(ctx context.Context) (inventory.Snapshot, error) {
			return collectInventory()
		},
		Retention: inventorystore.DefaultRetention,
		Rules:     pipeline.CanonicalObservationRules(),
	}
	execution, err := orchestrator.Execute(context.Background(), definition)
	if err != nil {
		fmt.Fprintf(errout, "canonical command execution failed: %s\n", safeText(err.Error()))
		return 1
	}
	if definition.Profile == "check" {
		if _, err = publishCheck(definition, execution, time.Now().UTC()); err != nil {
			fmt.Fprintf(errout, "current operator state publication failed: %s\n", safeText(publicationDiagnostic(err)))
			return 1
		}
	}
	var document []byte
	if definition.Parameters.Output == canonicalcommand.JSON {
		document, err = presentation.JSON(execution)
	} else {
		var text string
		text, err = presentation.Terminal(execution)
		document = []byte(text)
	}
	if err != nil {
		fmt.Fprintf(errout, "command presentation failed: %s\n", safeText(err.Error()))
		return 1
	}
	if _, err := out.Write(document); err != nil {
		fmt.Fprintf(errout, "command output failed: %s\n", safeText(err.Error()))
		return 1
	}
	return 0
}

func publishCheck(definition canonicalcommand.Definition, execution canonicalcommand.Execution, publishedAt time.Time) (presentationmodel.Overview, error) {
	plan, err := canonicalcommand.PlanDefinition(definition)
	if err != nil || definition.Profile != "check" || definition.Selection.Source != "live" || !execution.Complete || execution.CommandID != definition.ID || execution.PlanID != plan.ID || len(execution.Stages) != 2 {
		return presentationmodel.Overview{}, fmt.Errorf("ineligible check execution")
	}
	inventoryStage, snapshotStage := execution.Stages[0], execution.Stages[1]
	if inventoryStage.Stage != canonicalcommand.Inventory || snapshotStage.Stage != canonicalcommand.Snapshot || inventoryStage.ContractName != inventory.CanonicalSchemaName || inventoryStage.Version != inventory.SchemaVersion || snapshotStage.ContractName != inventorystore.FormatName || snapshotStage.Version != inventorystore.FormatVersion {
		return presentationmodel.Overview{}, fmt.Errorf("invalid check stage coverage")
	}
	observed, ok := inventoryStage.Value.(inventory.Snapshot)
	if !ok {
		return presentationmodel.Overview{}, fmt.Errorf("inventory stage is not typed")
	}
	snapshotted, ok := snapshotStage.Value.(inventory.Snapshot)
	if !ok || snapshotted.SnapshotID != observed.SnapshotID {
		return presentationmodel.Overview{}, fmt.Errorf("snapshot stage is not correlated")
	}
	if err = inventory.Validate(observed); err != nil {
		return presentationmodel.Overview{}, err
	}
	if err = inventory.Validate(snapshotted); err != nil || !reflect.DeepEqual(observed, snapshotted) {
		return presentationmodel.Overview{}, fmt.Errorf("snapshot stage payload mismatch")
	}
	observationTime := observed.CompletedAt.UTC()
	freshUntil := observed.FreshUntil.UTC()
	if !freshUntil.After(observationTime) || publishedAt.Before(observationTime) {
		return presentationmodel.Overview{}, fmt.Errorf("invalid observation freshness")
	}
	overview, err := presentationmodel.Project(presentationmodel.Input{SchemaName: presentationmodel.InputSchema, SchemaVersion: presentationmodel.SchemaVersion, ObservedAt: observationTime, FreshForNS: int64(freshUntil.Sub(observationTime)), Command: &presentationmodel.CommandObservation{ObservedAt: observationTime, Value: execution}, Inventory: &presentationmodel.InventoryObservation{ObservedAt: observationTime, Value: observed}})
	if err != nil {
		return presentationmodel.Overview{}, err
	}
	current, err := operatorstate.Normalize(operatorstate.State{ObservedAt: observationTime, PublishedAt: publishedAt.UTC(), FreshUntil: freshUntil, Coverage: operatorstate.CoverageInventorySnapshot, Provenance: operatorstate.Provenance{DefinitionID: definition.ID, ExecutionID: execution.ID, Profile: "check", Source: "live", Stages: []string{"inventory", "snapshot"}, Reason: operatorstate.PublicationCheck, ApplicationVersion: version}, Overview: overview})
	if err != nil {
		return presentationmodel.Overview{}, err
	}
	store, err := currentOperatorStore()
	if err != nil {
		return presentationmodel.Overview{}, err
	}
	if err = store.Publish(current); err != nil {
		return presentationmodel.Overview{}, err
	}
	return overview, nil
}

func publishObserve(definition canonicalcommand.Definition, execution canonicalcommand.Execution, publishedAt time.Time) (presentationmodel.Overview, error) {
	plan, err := canonicalcommand.PlanDefinition(definition)
	if err != nil || definition.Profile != "observe" || definition.Selection.Source != "live" || !execution.Complete || execution.CommandID != definition.ID || execution.PlanID != plan.ID || len(execution.Stages) != 8 {
		return presentationmodel.Overview{}, classifiedObservationFailure(observationProjectionFailure, fmt.Errorf("ineligible observe execution"))
	}
	observed, ok := execution.Stages[0].Value.(inventory.Snapshot)
	if !ok || inventory.Validate(observed) != nil {
		return presentationmodel.Overview{}, classifiedObservationFailure(observationProjectionFailure, fmt.Errorf("invalid typed inventory stage"))
	}
	observationTime, freshUntil := observed.CompletedAt.UTC(), observed.FreshUntil.UTC()
	if !freshUntil.After(observationTime) || publishedAt.Before(observationTime) {
		return presentationmodel.Overview{}, classifiedObservationFailure(observationProjectionFailure, fmt.Errorf("invalid observation freshness"))
	}
	overview, err := app.ProjectOperatorEvaluation(execution, observationTime, freshUntil.Sub(observationTime), nil, nil)
	if err != nil {
		return presentationmodel.Overview{}, classifiedObservationFailure(observationProjectionFailure, err)
	}
	stages := []string{"inventory", "snapshot", "compare", "drift", "health", "rule", "policy", "report"}
	current, err := operatorstate.Normalize(operatorstate.State{ObservedAt: observationTime, PublishedAt: publishedAt.UTC(), FreshUntil: freshUntil, Coverage: operatorstate.CoverageOperatorEvaluation, Provenance: operatorstate.Provenance{DefinitionID: definition.ID, ExecutionID: execution.ID, Profile: "observe", Source: "live", Stages: stages, Reason: operatorstate.PublicationObserve, ApplicationVersion: version}, Overview: overview})
	if err != nil {
		return presentationmodel.Overview{}, classifiedObservationFailure(observationPublicationFailure, err)
	}
	store, err := currentOperatorStore()
	if err != nil {
		return presentationmodel.Overview{}, classifiedObservationFailure(observationPublicationFailure, err)
	}
	if err = store.Publish(current); err != nil {
		return presentationmodel.Overview{}, classifiedObservationFailure(observationPublicationFailure, err)
	}
	return overview, nil
}

func currentOperatorStore() (*operatorstate.Store, error) {
	root, err := localStateRoot()
	if err != nil {
		return nil, err
	}
	return operatorstate.Open(root)
}

func observationStoreRoot() (string, error) {
	if root := os.Getenv("QWSG_STORE"); root != "" {
		return root, nil
	}
	root, err := localStateRoot()
	if err != nil {
		return "", err
	}
	if err = ensureLocalStateRoot(root); err != nil {
		return "", err
	}
	return filepath.Join(root, "inventory"), nil
}

func ensureLocalStateRoot(root string) error {
	if err := operatorstate.EnsurePrivateRoot(root); err != nil {
		return classifiedObservationFailure(observationBootstrapFailure, err)
	}
	return nil
}

func localStateRoot() (string, error) {
	if root := os.Getenv("QWSG_STATE_DIR"); root != "" {
		return root, nil
	}
	if base := os.Getenv("XDG_STATE_HOME"); base != "" {
		return filepath.Join(base, "qwsg"), nil
	}
	if home := os.Getenv("HOME"); home != "" {
		return filepath.Join(home, ".local", "state", "qwsg"), nil
	}
	return "", fmt.Errorf("local state root is unavailable")
}

func stateDiagnostic(err error) string {
	switch {
	case errors.Is(err, operatorstate.ErrCorrupt):
		return "state_corrupt"
	case errors.Is(err, operatorstate.ErrIncompatible):
		return "state_incompatible"
	case errors.Is(err, operatorstate.ErrPermission):
		return "state_permission"
	case errors.Is(err, operatorstate.ErrUnsafePath):
		return "state_unsafe"
	default:
		return "state_unreadable"
	}
}

func publicationDiagnostic(err error) string {
	switch {
	case errors.Is(err, operatorstate.ErrCorrupt), errors.Is(err, operatorstate.ErrIncompatible), errors.Is(err, operatorstate.ErrPermission), errors.Is(err, operatorstate.ErrUnsafePath):
		return stateDiagnostic(err)
	default:
		return "state_publication_failed"
	}
}

func writeCanonicalHelp(out io.Writer, name string) {
	fmt.Fprintf(out, "QWSG canonical command profile: %s\n\n", safeText(name))
	fmt.Fprintln(out, "Common parameters:")
	fmt.Fprintln(out, "  --store DIR")
	fmt.Fprintln(out, "  --from SNAPSHOT --to SNAPSHOT")
	fmt.Fprintln(out, "  --filter FIELD=VALUE  --group FIELD  --sort FIELD")
	fmt.Fprintln(out, "  --output json|human  --presentation structured|terminal")
	if name == "analyze" {
		fmt.Fprintln(out, "  --source live|store --pipeline STAGE[,STAGE]")
		fmt.Fprintln(out, "  --include STAGE[,STAGE] --exclude STAGE[,STAGE]")
	}
	fmt.Fprintln(out, "\nAll profiles resolve to Command Definition 1.0 and the same canonical pipeline.")
}

func runHelp(args []string, out, errout io.Writer) int {
	if len(args) == 0 {
		writeRootHelp(out)
		return 0
	}
	if args[0] == "version" && len(args) == 1 {
		writeVersionHelp(out)
		return 0
	}
	if args[0] == "update" && len(args) == 1 {
		writeUpdateHelp(out)
		return 0
	}
	if args[0] == "compare" {
		if len(args) != 1 {
			return usageError(errout, "compare help does not accept a subcommand")
		}
		writeCompareHelp(out)
		return 0
	}
	for _, topic := range []string{"status", "check", "observe", "changes", "health", "report", "analyze"} {
		if args[0] == topic {
			if len(args) != 1 {
				return usageError(errout, "%s help does not accept a subcommand", topic)
			}
			writeCanonicalHelp(out, topic)
			return 0
		}
	}
	if args[0] != "inventory" {
		return usageError(errout, "unknown help topic: %s", safeText(strings.Join(args, " ")))
	}
	return writeInventoryHelp(args[1:], out, errout)
}

func runCompare(args []string, out, errout io.Writer) int {
	if len(args) > 0 && isHelp(args[0]) {
		if len(args) != 1 {
			return usageError(errout, "compare help does not accept options")
		}
		writeCompareHelp(out)
		return 0
	}
	options, err := parseCompareArgs(args)
	if err != nil {
		return usageError(errout, "%v", err)
	}
	store, err := inventorystore.Open(options.storePath, options.retention)
	if err != nil {
		fmt.Fprintf(errout, "comparison store configuration failed: %s\n", safeText(err.Error()))
		return 1
	}
	fromName, toName := options.from, options.to
	if fromName == "" {
		names, err := store.List()
		if err != nil {
			fmt.Fprintf(errout, "comparison snapshot selection failed: %s\n", safeText(err.Error()))
			return 1
		}
		if len(names) < 2 {
			fmt.Fprintln(errout, "comparison requires at least two stored snapshots")
			return 1
		}
		fromName, toName = names[len(names)-2], names[len(names)-1]
	}
	from, err := store.Load(fromName)
	if err != nil {
		fmt.Fprintf(errout, "comparison from snapshot load failed: %s\n", safeText(err.Error()))
		return 1
	}
	to, err := store.Load(toName)
	if err != nil {
		fmt.Fprintf(errout, "comparison to snapshot load failed: %s\n", safeText(err.Error()))
		return 1
	}
	result, err := comparison.Compare(from, to, fromName, toName)
	if err != nil {
		fmt.Fprintf(errout, "snapshot comparison failed: %s\n", safeText(err.Error()))
		return 1
	}
	if options.format == formatJSON {
		return writeJSON(out, errout, result)
	}
	return writeHumanComparison(out, result)
}

func parseCompareArgs(args []string) (compareOptions, error) {
	options := compareOptions{format: formatJSON, retention: inventorystore.DefaultRetention}
	if configured := os.Getenv("QWSG_FORMAT"); configured != "" {
		if !validFormat(configured) {
			return compareOptions{}, fmt.Errorf("QWSG_FORMAT must be json or human")
		}
		options.format = configured
	}
	for i := 0; i < len(args); i++ {
		switch args[i] {
		case "--store":
			i++
			if i >= len(args) || options.storePath != "" {
				return compareOptions{}, fmt.Errorf("--store requires one value")
			}
			options.storePath = args[i]
		case "--from":
			i++
			if i >= len(args) || options.from != "" {
				return compareOptions{}, fmt.Errorf("--from requires one snapshot name")
			}
			options.from = args[i]
		case "--to":
			i++
			if i >= len(args) || options.to != "" {
				return compareOptions{}, fmt.Errorf("--to requires one snapshot name")
			}
			options.to = args[i]
		case "--retention":
			i++
			if i >= len(args) {
				return compareOptions{}, fmt.Errorf("--retention requires one value")
			}
			value, err := strconv.Atoi(args[i])
			if err != nil {
				return compareOptions{}, fmt.Errorf("--retention must be an integer")
			}
			options.retention = value
		case "--format":
			i++
			if i >= len(args) || !validFormat(args[i]) {
				return compareOptions{}, fmt.Errorf("--format must be json or human")
			}
			options.format = args[i]
		default:
			return compareOptions{}, fmt.Errorf("unknown compare option: %s", safeText(args[i]))
		}
	}
	if options.storePath == "" {
		options.storePath = os.Getenv("QWSG_STORE")
	}
	if options.storePath == "" {
		return compareOptions{}, fmt.Errorf("--store is required (or set QWSG_STORE)")
	}
	if (options.from == "") != (options.to == "") {
		return compareOptions{}, fmt.Errorf("--from and --to must be provided together")
	}
	return options, nil
}

func writeHumanComparison(out io.Writer, result comparison.Result) int {
	fmt.Fprintf(out, "Snapshot comparison\nFrom: %s\nTo: %s\nCompared: %s\n",
		safeText(result.From.Selector), safeText(result.To.Selector),
		result.ComparisonTimestamp.UTC().Format(time.RFC3339Nano))
	groups := []struct {
		name       string
		changeType comparison.ChangeType
		count      int
	}{
		{"Added", comparison.Added, result.Counts.Added},
		{"Removed", comparison.Removed, result.Counts.Removed},
		{"Modified", comparison.Modified, result.Counts.Modified},
		{"Unchanged", comparison.Unchanged, result.Counts.Unchanged},
	}
	for _, group := range groups {
		fmt.Fprintf(out, "\n%s (%d)\n", group.name, group.count)
		if group.count == 0 {
			fmt.Fprintln(out, "- none")
			continue
		}
		for _, record := range result.Changes {
			if record.Type != group.changeType {
				continue
			}
			fmt.Fprintf(out, "- [%s] %s", safeText(record.Layer), safeText(record.Path))
			switch record.Type {
			case comparison.Added:
				fmt.Fprintf(out, " = %s", humanValue(record.Current))
			case comparison.Removed:
				fmt.Fprintf(out, " (was %s)", humanValue(record.Previous))
			case comparison.Modified:
				fmt.Fprintf(out, ": %s -> %s", humanValue(record.Previous), humanValue(record.Current))
			}
			fmt.Fprintln(out)
		}
	}
	return 0
}

func humanValue(value *comparison.TypedValue) string {
	if value == nil {
		return "null"
	}
	document, err := json.Marshal(value.Value)
	if err != nil {
		return "<unavailable>"
	}
	return safeText(string(document))
}

func runInventory(args []string, out, errout io.Writer) int {
	if len(args) > 0 && isHelp(args[0]) {
		return writeInventoryHelp(args[1:], out, errout)
	}
	action := "collect"
	if len(args) > 0 && !strings.HasPrefix(args[0], "-") {
		action = args[0]
		args = args[1:]
	}
	if len(args) > 0 && isHelp(args[0]) {
		return writeInventoryHelp([]string{action}, out, errout)
	}
	switch action {
	case "collect":
		format, err := parseCollectArgs(args)
		if err != nil {
			return usageError(errout, "%v", err)
		}
		snapshot, err := collectInventory()
		if err != nil {
			fmt.Fprintln(errout, safeText(err.Error()))
			return 1
		}
		return writeInventory(out, errout, snapshot, format, "Live inventory")
	case "save", "load", "list", "info":
		return runStoreAction(action, args, out, errout)
	default:
		return usageError(errout, "unknown inventory subcommand: %s", safeText(action))
	}
}

func runStoreAction(action string, args []string, out, errout io.Writer) int {
	options, err := parseStoreArgs(action, args)
	if err != nil {
		return usageError(errout, "%v", err)
	}
	store, err := inventorystore.Open(options.storePath, options.retention)
	if err != nil {
		fmt.Fprintf(errout, "inventory store configuration failed: %s\n", safeText(err.Error()))
		return 1
	}
	switch action {
	case "save":
		snapshot, err := collectInventory()
		if err == nil {
			_, err = store.Save(snapshot)
		}
		if err != nil {
			fmt.Fprintf(errout, "inventory persistence failed: %s\n", safeText(err.Error()))
			return 1
		}
		return writeInventory(out, errout, snapshot, options.format, "Saved inventory")
	case "load":
		snapshot, name, err := loadSnapshot(store, options.snapshotName)
		if err != nil {
			fmt.Fprintf(errout, "inventory load failed: %s\n", safeText(err.Error()))
			return 1
		}
		return writeInventory(out, errout, snapshot, options.format, "Stored inventory "+name)
	case "list":
		return listSnapshots(store, options.format, out, errout)
	case "info":
		snapshot, name, err := loadSnapshot(store, options.snapshotName)
		if err != nil {
			fmt.Fprintf(errout, "inventory info failed: %s\n", safeText(err.Error()))
			return 1
		}
		return writeSummary(out, errout, summarize(name, snapshot), options.format)
	}
	return 1
}

func parseCollectArgs(args []string) (string, error) {
	format := formatJSON
	if configured := os.Getenv("QWSG_FORMAT"); configured != "" {
		if !validFormat(configured) {
			return "", fmt.Errorf("QWSG_FORMAT must be json or human")
		}
		format = configured
	}
	for i := 0; i < len(args); i++ {
		if args[i] != "--format" || i+1 >= len(args) {
			return "", fmt.Errorf("inventory accepts only --format <json|human>")
		}
		i++
		if !validFormat(args[i]) {
			return "", fmt.Errorf("--format must be json or human")
		}
		format = args[i]
	}
	return format, nil
}

func parseStoreArgs(action string, args []string) (storeOptions, error) {
	options := storeOptions{retention: inventorystore.DefaultRetention, format: formatJSON}
	if action == "list" || action == "info" {
		options.format = formatHuman
	}
	if configured := os.Getenv("QWSG_FORMAT"); configured != "" {
		if !validFormat(configured) {
			return storeOptions{}, fmt.Errorf("QWSG_FORMAT must be json or human")
		}
		options.format = configured
	}
	for i := 0; i < len(args); i++ {
		switch args[i] {
		case "--store":
			i++
			if i >= len(args) || options.storePath != "" {
				return storeOptions{}, fmt.Errorf("--store requires one value")
			}
			options.storePath = args[i]
		case "--snapshot":
			i++
			if i >= len(args) || options.snapshotName != "" {
				return storeOptions{}, fmt.Errorf("--snapshot requires one value")
			}
			options.snapshotName = args[i]
		case "--retention":
			i++
			if i >= len(args) {
				return storeOptions{}, fmt.Errorf("--retention requires one value")
			}
			value, err := strconv.Atoi(args[i])
			if err != nil {
				return storeOptions{}, fmt.Errorf("--retention must be an integer")
			}
			options.retention = value
		case "--format":
			i++
			if i >= len(args) || !validFormat(args[i]) {
				return storeOptions{}, fmt.Errorf("--format must be json or human")
			}
			options.format = args[i]
		default:
			return storeOptions{}, fmt.Errorf("unknown inventory store option: %s", safeText(args[i]))
		}
	}
	if options.storePath == "" {
		options.storePath = os.Getenv("QWSG_STORE")
	}
	if options.storePath == "" {
		return storeOptions{}, fmt.Errorf("--store is required (or set QWSG_STORE)")
	}
	if options.snapshotName != "" && action != "load" && action != "info" {
		return storeOptions{}, fmt.Errorf("--snapshot is valid only for inventory load or info")
	}
	return options, nil
}

func loadSnapshot(store *inventorystore.Store, name string) (inventory.Snapshot, string, error) {
	if name == "" {
		return store.LoadLatest()
	}
	snapshot, err := store.Load(name)
	return snapshot, name, err
}

func listSnapshots(store *inventorystore.Store, format string, out, errout io.Writer) int {
	names, err := store.List()
	if err != nil {
		fmt.Fprintf(errout, "inventory list failed: %s\n", safeText(err.Error()))
		return 1
	}
	summaries := make([]snapshotSummary, 0, len(names))
	for _, name := range names {
		snapshot, err := store.Load(name)
		if err != nil {
			fmt.Fprintf(errout, "inventory list failed while validating %s: %s\n", safeText(name), safeText(err.Error()))
			return 1
		}
		summaries = append(summaries, summarize(name, snapshot))
	}
	if format == formatJSON {
		return writeJSON(out, errout, summaries)
	}
	if len(summaries) == 0 {
		fmt.Fprintln(out, "Snapshots: none")
		return 0
	}
	fmt.Fprintf(out, "Snapshots: %d (oldest to newest)\n", len(summaries))
	for _, summary := range summaries {
		fmt.Fprintf(out, "- %s  %s  %s  status=%s\n",
			safeText(summary.Name), summary.CompletedAt.UTC().Format(time.RFC3339Nano),
			safeText(summary.SchemaVersion), safeText(string(summary.Status)))
	}
	return 0
}

func summarize(name string, snapshot inventory.Snapshot) snapshotSummary {
	return snapshotSummary{
		Name: name, Status: snapshot.Status, ObservedAt: snapshot.ObservedAt,
		CompletedAt: snapshot.CompletedAt, SchemaVersion: snapshot.SchemaVersion,
		SchemaName: snapshot.Canonical.SchemaName, CategoryCount: len(snapshot.Categories),
		IssueCount: len(snapshot.Errors) + len(snapshot.Canonical.Issues), Integrity: "verified",
	}
}

func writeSummary(out, errout io.Writer, summary snapshotSummary, format string) int {
	if format == formatJSON {
		return writeJSON(out, errout, summary)
	}
	fmt.Fprintf(out, "Snapshot: %s\nStatus: %s\nObserved: %s\nCompleted: %s\nInventory schema: %s\nCanonical schema: %s\nCategories: %d\nIssues: %d\nIntegrity: %s\n",
		safeText(summary.Name), safeText(string(summary.Status)),
		summary.ObservedAt.UTC().Format(time.RFC3339Nano), summary.CompletedAt.UTC().Format(time.RFC3339Nano),
		safeText(summary.SchemaVersion), safeText(summary.SchemaName),
		summary.CategoryCount, summary.IssueCount, safeText(summary.Integrity))
	return inventory.ExitCode(summary.Status)
}

func collectInventory() (inventory.Snapshot, error) {
	return collectInventoryContext(context.Background())
}

func collectInventoryContext(parent context.Context) (inventory.Snapshot, error) {
	r := runner.Bounded{Allowed: map[string]string{"systemctl": "/usr/bin/systemctl", "go": "/usr/local/go/bin/go"}, Timeout: 2 * time.Second, MaxOutput: 1 << 20}
	registry, err := collector.DefaultRegistry(r)
	if err != nil {
		return inventory.Snapshot{}, fmt.Errorf("collector registry initialization failed: %w", err)
	}
	ctx, cancel := context.WithTimeout(parent, 15*time.Second)
	defer cancel()
	snapshot, err := app.Collect(ctx, version, registry)
	if err != nil {
		return inventory.Snapshot{}, fmt.Errorf("inventory validation failed: %w", err)
	}
	return snapshot, nil
}

func writeInventory(out, errout io.Writer, snapshot inventory.Snapshot, format, label string) int {
	if format == formatJSON {
		if code := writeJSON(out, errout, snapshot); code != 0 {
			return code
		}
		return inventory.ExitCode(snapshot.Status)
	}
	fmt.Fprintf(out, "%s\nStatus: %s\nObserved: %s\nCompleted: %s\nCategories: %d\nCanonical layers: %d\nIssues: %d\nRedactions: %d\n",
		safeText(label), safeText(string(snapshot.Status)),
		snapshot.ObservedAt.UTC().Format(time.RFC3339Nano), snapshot.CompletedAt.UTC().Format(time.RFC3339Nano),
		len(snapshot.Categories), len(snapshot.Canonical.Layers),
		len(snapshot.Errors)+len(snapshot.Canonical.Issues), len(snapshot.Redactions)+len(snapshot.Canonical.Redactions))
	if snapshot.Status != inventory.Complete {
		fmt.Fprintln(out, "Result is not a health verdict; inspect JSON for structured collector issues.")
	}
	return inventory.ExitCode(snapshot.Status)
}

func writeJSON(out, errout io.Writer, value any) int {
	encoder := json.NewEncoder(out)
	encoder.SetIndent("", "  ")
	if err := encoder.Encode(value); err != nil {
		fmt.Fprintf(errout, "output encoding failed: %s\n", safeText(err.Error()))
		return 1
	}
	return 0
}

func safeText(value string) string {
	var result strings.Builder
	for _, r := range value {
		if unicode.IsControl(r) || r == '\u007f' {
			fmt.Fprintf(&result, "\\u%04x", r)
			continue
		}
		result.WriteRune(r)
	}
	return result.String()
}

func validFormat(value string) bool { return value == formatJSON || value == formatHuman }
func isHelp(value string) bool      { return value == "help" || value == "--help" || value == "-h" }

func usageError(errout io.Writer, format string, values ...any) int {
	fmt.Fprintf(errout, format+"\n", values...)
	fmt.Fprintln(errout, "Run 'qwsg help' for usage.")
	return 1
}

func writeRootHelp(out io.Writer) {
	fmt.Fprintln(out, `QWSG — deterministic Linux engineering analysis and snapshot explorer

Usage:
  qwsg
	qwsg console
  qwsg setup [--accept-defaults] [--set KEY=VALUE]
  qwsg config <show|validate|get|set> ...
  qwsg notification <preflight|test|credential> ...
  qwsg install --check [--format human|json]
  qwsg readiness [--format human|json]
  qwsg update <check|status|rollback> | qwsg update
  qwsg help [command]
  qwsg version
  qwsg <status|check|observe|changes|health|report> [options]
  qwsg analyze --source live|store --pipeline STAGE[,STAGE] [options]
  qwsg inventory [--format json|human]
  qwsg inventory <save|list|info|load> [options]
  qwsg compare [--store DIR] [--from SNAPSHOT --to SNAPSHOT] [--format json|human]

Commands:
	console    Open the read-only local Operator Console
  setup      Safely create or review the per-user QWSG configuration
  config     Show, validate, read, or update canonical configuration
  notification Assess readiness, store a private SMTP credential, or send a test
  install     Assess installation and host requirements without mutation
  readiness   Show composite operational readiness
  update      Discover, verify, install, inspect, or roll back QWSG updates
  status     Execute the canonical live Inventory profile
  check      Execute the canonical live Inventory and Snapshot profile
  observe    Establish a baseline or run the full canonical operator evaluation
  changes    Execute the canonical Compare profile over stored snapshots
  health     Execute Compare → Drift → Health over stored snapshots
  report     Execute Compare → Drift → Health → Rule → Policy → Report
  analyze    Compose Command Definition 1.0 with structured parameters
  inventory  Collect Inventory 1.0 or browse an explicit private snapshot store
  compare    Compare canonical Inventory snapshots without making a health judgement
  version    Show version and build information
  help       Show root or contextual help

Every command resolves through the presentation-independent Canonical Command
Architecture. JSON remains the compatibility default for legacy inventory,
save, load, and compare commands.`)
}

func writeVersionHelp(out io.Writer) {
	fmt.Fprintln(out, "Usage: qwsg version\n\nShows the QWSG version, build commit, and controlled build date.")
}

func writeCompareHelp(out io.Writer) {
	fmt.Fprintln(out, `Usage:
  qwsg compare --store DIR [--retention N] [--format json|human]
  qwsg compare --store DIR --from SNAPSHOT --to SNAPSHOT [--retention N] [--format json|human]

Without selectors, compare uses the previous and latest snapshots in the store.
--from and --to select exact names returned by 'qwsg inventory list'. QWSG_STORE
and QWSG_FORMAT may provide the store and format; command-line options take
precedence. JSON is the canonical default.

Exit 0 means comparison succeeded. Exit 1 means usage, selection, store,
integrity, compatibility, or output failure. Change Records report facts only:
they are not drift, health, alert, scoring, or recommendation results.`)
}

func writeInventoryHelp(args []string, out, errout io.Writer) int {
	if len(args) > 1 {
		return usageError(errout, "inventory help accepts at most one subcommand")
	}
	if len(args) == 0 || args[0] == "collect" {
		fmt.Fprintln(out, `Usage:
  qwsg inventory [--format json|human]
  qwsg inventory save --store DIR [--retention N] [--format json|human]
  qwsg inventory list --store DIR [--retention N] [--format human|json]
  qwsg inventory info --store DIR [--retention N] [--snapshot NAME] [--format human|json]
  qwsg inventory load --store DIR [--retention N] [--snapshot NAME] [--format json|human]

The store path must be clean, absolute, and private. Retention defaults to 10
and is fixed when the store is created. Info and load use the latest snapshot
unless --snapshot selects an exact name returned by list.
QWSG_STORE may provide the same explicit store path, and QWSG_FORMAT may be
json or human. Command-line options take precedence.

Exit 0 means success/complete Inventory, 2 means partial but usable Inventory,
and 1 means usage, store, validation, permission, corruption, or runtime failure.
Stored observations are evidence, not current state or a health verdict.`)
		return 0
	}
	switch args[0] {
	case "save", "list", "info", "load":
		fmt.Fprintf(out, "Usage: qwsg inventory %s --store DIR [--retention N]", args[0])
		if args[0] == "info" || args[0] == "load" {
			fmt.Fprint(out, " [--snapshot NAME]")
		}
		fmt.Fprintln(out, " [--format json|human]")
		return 0
	default:
		return usageError(errout, "unknown inventory help topic: %s", safeText(args[0]))
	}
}
