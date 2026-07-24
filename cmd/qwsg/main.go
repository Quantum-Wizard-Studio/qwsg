package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strconv"
	"time"

	"quantumwizard.hu/qwsg/internal/app"
	"quantumwizard.hu/qwsg/internal/collector"
	"quantumwizard.hu/qwsg/internal/inventory"
	"quantumwizard.hu/qwsg/internal/inventorystore"
	"quantumwizard.hu/qwsg/internal/runner"
)

const version = "0.0.1-prealpha"

func main() { os.Exit(run(os.Args[1:], os.Stdout, os.Stderr)) }
func run(args []string, out, errout io.Writer) int {
	if len(args) == 0 || args[0] == "help" || args[0] == "--help" || args[0] == "-h" {
		fmt.Fprintln(out, "Usage: qwsg <inventory|version|help>\n\ninventory                 emit a read-only JSON inventory\ninventory save --store DIR  collect and persist one inventory\ninventory load --store DIR  load the latest persisted inventory\nversion                   print the QWSG version")
		return 0
	}
	switch args[0] {
	case "version":
		if len(args) != 1 {
			fmt.Fprintln(errout, "version does not accept options")
			return 1
		}
		fmt.Fprintln(out, version)
		return 0
	case "inventory":
		return runInventory(args[1:], out, errout)
	default:
		fmt.Fprintf(errout, "unknown command: %s\n", args[0])
		return 1
	}
}

func runInventory(args []string, out, errout io.Writer) int {
	if len(args) == 0 {
		snapshot, err := collectInventory()
		if err != nil {
			fmt.Fprintf(errout, "%v\n", err)
			return 1
		}
		return writeInventory(out, errout, snapshot)
	}
	action := args[0]
	if action != "save" && action != "load" {
		fmt.Fprintln(errout, "inventory accepts only save or load subcommands")
		return 1
	}
	storePath, snapshotName, retention, err := parseStoreArgs(args[1:])
	if err != nil {
		fmt.Fprintln(errout, err)
		return 1
	}
	store, err := inventorystore.Open(storePath, retention)
	if err != nil {
		fmt.Fprintf(errout, "inventory store configuration failed: %v\n", err)
		return 1
	}
	var snapshot inventory.Snapshot
	switch action {
	case "save":
		if snapshotName != "" {
			fmt.Fprintln(errout, "--snapshot is valid only for inventory load")
			return 1
		}
		snapshot, err = collectInventory()
		if err == nil {
			_, err = store.Save(snapshot)
		}
		if err != nil {
			fmt.Fprintf(errout, "inventory persistence failed: %v\n", err)
			return 1
		}
	case "load":
		if snapshotName == "" {
			snapshot, _, err = store.LoadLatest()
		} else {
			snapshot, err = store.Load(snapshotName)
		}
		if err != nil {
			fmt.Fprintf(errout, "inventory load failed: %v\n", err)
			return 1
		}
	}
	return writeInventory(out, errout, snapshot)
}

func parseStoreArgs(args []string) (string, string, int, error) {
	storePath := ""
	snapshotName := ""
	retention := inventorystore.DefaultRetention
	for i := 0; i < len(args); i++ {
		switch args[i] {
		case "--store":
			i++
			if i >= len(args) || storePath != "" {
				return "", "", 0, fmt.Errorf("--store requires one value")
			}
			storePath = args[i]
		case "--snapshot":
			i++
			if i >= len(args) || snapshotName != "" {
				return "", "", 0, fmt.Errorf("--snapshot requires one value")
			}
			snapshotName = args[i]
		case "--retention":
			i++
			if i >= len(args) {
				return "", "", 0, fmt.Errorf("--retention requires one value")
			}
			value, err := strconv.Atoi(args[i])
			if err != nil {
				return "", "", 0, fmt.Errorf("--retention must be an integer")
			}
			retention = value
		default:
			return "", "", 0, fmt.Errorf("unknown inventory store option: %s", args[i])
		}
	}
	if storePath == "" {
		return "", "", 0, fmt.Errorf("--store is required")
	}
	return storePath, snapshotName, retention, nil
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

func writeInventory(out, errout io.Writer, snapshot inventory.Snapshot) int {
	encoder := json.NewEncoder(out)
	encoder.SetIndent("", "  ")
	if err := encoder.Encode(snapshot); err != nil {
		fmt.Fprintf(errout, "inventory encoding failed: %v\n", err)
		return 1
	}
	return inventory.ExitCode(snapshot.Status)
}
