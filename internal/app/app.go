package app

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"sort"
	"time"

	"quantumwizard.hu/qwsg/internal/collector"
	"quantumwizard.hu/qwsg/internal/inventory"
)

func Collect(ctx context.Context, version string, registry *collector.Registry) (inventory.Snapshot, error) {
	start := time.Now().UTC()
	id := randomID()
	results := registry.Execute(ctx, id)
	cats := make([]inventory.Category, 0, len(results))
	executions := make([]inventory.CollectorExecution, 0, len(results))
	for _, result := range results {
		cats = append(cats, result.CollectedData)
		warnings := make([]inventory.InventoryWarning, 0, len(result.Warnings))
		for _, warning := range result.Warnings {
			warnings = append(warnings, inventory.InventoryWarning{Code: warning.Code, MessageKey: warning.MessageKey})
		}
		executions = append(executions, inventory.CollectorExecution{CollectorName: result.CollectorName, Version: result.Version, Capability: result.Capability, SupportedPlatforms: append([]string(nil), result.SupportedPlatforms...), ExecutionTimeMS: result.ExecutionTimeMS, Timestamp: result.Timestamp, Status: result.HealthStatus, Warnings: warnings, Errors: result.Errors, Metadata: result.Metadata})
	}
	sort.Slice(cats, func(i, j int) bool { return cats[i].CategoryID < cats[j].CategoryID })
	end := time.Now().UTC()
	subjectID := instanceID(cats)
	durationMS := time.Since(start).Milliseconds()
	producer := inventory.Producer{ToolVersion: version, ContractVersion: "1.0"}
	s := inventory.Snapshot{SchemaVersion: inventory.SchemaVersion, SnapshotID: id, RequestID: id, InstanceID: subjectID, ObservedAt: start, CompletedAt: end, FreshUntil: end.Add(5 * time.Minute), DurationMS: durationMS, Status: inventory.Aggregate(cats), Categories: cats, Errors: []inventory.InventoryError{}, Redactions: []string{"hostnames", "network addresses", "hardware addresses", "mount paths", "service identities"}, Producer: producer}
	s.Canonical = inventory.AssembleSystemInventory(cats, executions, id, id, subjectID, start, end, s.FreshUntil, durationMS, producer)
	return s, inventory.Validate(s)
}
func randomID() string {
	b := make([]byte, 16)
	if _, e := rand.Read(b); e != nil {
		return "unavailable"
	}
	return hex.EncodeToString(b)
}
func instanceID(cs []inventory.Category) string {
	for _, c := range cs {
		if c.CategoryID == "host" && len(c.Items) > 0 {
			if f, ok := c.Items[0].Facts["instance_id"]; ok {
				if s, ok := f.Value.(string); ok {
					return s
				}
			}
		}
	}
	return "unavailable"
}
