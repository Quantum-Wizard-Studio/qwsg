package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"time"
	"unicode"

	"quantumwizard.hu/qwsg/internal/app"
	"quantumwizard.hu/qwsg/internal/collector"
	"quantumwizard.hu/qwsg/internal/inventory"
	"quantumwizard.hu/qwsg/internal/inventorystore"
	"quantumwizard.hu/qwsg/internal/runner"
)

var (
	version     = "0.0.1-prealpha"
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

func main() { os.Exit(run(os.Args[1:], os.Stdout, os.Stderr)) }

func run(args []string, out, errout io.Writer) int {
	if len(args) == 0 {
		writeRootHelp(out)
		return 0
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
	default:
		return usageError(errout, "unknown command: %s", safeText(args[0]))
	}
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
	if args[0] != "inventory" {
		return usageError(errout, "unknown help topic: %s", safeText(strings.Join(args, " ")))
	}
	return writeInventoryHelp(args[1:], out, errout)
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
	r := runner.Bounded{Allowed: map[string]string{"systemctl": "/usr/bin/systemctl", "go": "/usr/local/go/bin/go"}, Timeout: 2 * time.Second, MaxOutput: 1 << 20}
	registry, err := collector.DefaultRegistry(r)
	if err != nil {
		return inventory.Snapshot{}, fmt.Errorf("collector registry initialization failed: %w", err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
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
	fmt.Fprintln(out, `QWSG — read-only Linux inventory and snapshot explorer

Usage:
  qwsg help [command]
  qwsg version
  qwsg inventory [--format json|human]
  qwsg inventory <save|list|info|load> [options]

Commands:
  inventory  Collect Inventory 1.0 or browse an explicit private snapshot store
  version    Show version and build information
  help       Show root or contextual help

JSON remains the compatibility default for inventory, save, and load.
Run 'qwsg help inventory' for store options and exit semantics.`)
}

func writeVersionHelp(out io.Writer) {
	fmt.Fprintln(out, "Usage: qwsg version\n\nShows the QWSG version, build commit, and controlled build date.")
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
