package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"time"

	"quantumwizard.hu/qwsg/internal/app"
	"quantumwizard.hu/qwsg/internal/collector"
	"quantumwizard.hu/qwsg/internal/inventory"
	"quantumwizard.hu/qwsg/internal/runner"
)

const version = "0.0.1-prealpha"

func main() { os.Exit(run(os.Args[1:], os.Stdout, os.Stderr)) }
func run(args []string, out, errout io.Writer) int {
	if len(args) == 0 || args[0] == "help" || args[0] == "--help" || args[0] == "-h" {
		fmt.Fprintln(out, "Usage: qwsg <inventory|version|help>\n\ninventory  emit a read-only JSON inventory\nversion    print the QWSG version")
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
		if len(args) != 1 {
			fmt.Fprintln(errout, "inventory does not accept options")
			return 1
		}
		r := runner.Bounded{Allowed: map[string]string{"systemctl": "/usr/bin/systemctl", "go": "/usr/local/go/bin/go"}, Timeout: 2 * time.Second, MaxOutput: 1 << 20}
		registry, registryErr := collector.DefaultRegistry(r)
		if registryErr != nil {
			fmt.Fprintf(errout, "collector registry initialization failed: %v\n", registryErr)
			return 1
		}
		ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
		defer cancel()
		s, e := app.Collect(ctx, version, registry)
		if e != nil {
			fmt.Fprintf(errout, "inventory validation failed: %v\n", e)
			return 1
		}
		enc := json.NewEncoder(out)
		enc.SetIndent("", "  ")
		if e = enc.Encode(s); e != nil {
			fmt.Fprintf(errout, "inventory encoding failed: %v\n", e)
			return 1
		}
		return inventory.ExitCode(s.Status)
	default:
		fmt.Fprintf(errout, "unknown command: %s\n", args[0])
		return 1
	}
}
