package comparison

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"sort"
	"strings"
	"time"

	"quantumwizard.hu/qwsg/internal/inventory"
)

const (
	SchemaName    = "qwsg.comparison"
	SchemaVersion = "1.0"
	EngineVersion = "1.0"
)

type ChangeType string

const (
	Added     ChangeType = "added"
	Removed   ChangeType = "removed"
	Modified  ChangeType = "modified"
	Unchanged ChangeType = "unchanged"
)

type SnapshotReference struct {
	Selector   string           `json:"selector"`
	SnapshotID string           `json:"snapshot_id"`
	Completed  time.Time        `json:"completed_at"`
	Status     inventory.Status `json:"status"`
}

type TypedValue struct {
	Type        string `json:"type"`
	Value       any    `json:"value"`
	Unit        string `json:"unit,omitempty"`
	Quality     string `json:"quality,omitempty"`
	Sensitivity string `json:"sensitivity,omitempty"`
	ReasonCode  string `json:"reason_code,omitempty"`
}

type ChangeRecord struct {
	ID                  string            `json:"id"`
	Layer               string            `json:"layer"`
	ObjectID            string            `json:"object_id"`
	Path                string            `json:"path"`
	Type                ChangeType        `json:"change_type"`
	Previous            *TypedValue       `json:"previous"`
	Current             *TypedValue       `json:"current"`
	ComparisonTimestamp time.Time         `json:"comparison_timestamp"`
	Metadata            map[string]string `json:"metadata"`
}

type Counts struct {
	Added     int `json:"added"`
	Removed   int `json:"removed"`
	Modified  int `json:"modified"`
	Unchanged int `json:"unchanged"`
}

type Result struct {
	SchemaName          string            `json:"schema_name"`
	SchemaVersion       string            `json:"schema_version"`
	EngineVersion       string            `json:"engine_version"`
	ComparisonID        string            `json:"comparison_id"`
	SubjectID           string            `json:"subject_id"`
	InventorySchema     string            `json:"inventory_schema"`
	InventoryProfile    string            `json:"inventory_profile"`
	ComparisonTimestamp time.Time         `json:"comparison_timestamp"`
	From                SnapshotReference `json:"from"`
	To                  SnapshotReference `json:"to"`
	Counts              Counts            `json:"counts"`
	Changes             []ChangeRecord    `json:"changes"`
	Metadata            map[string]string `json:"metadata"`
}

func Compare(from, to inventory.Snapshot, fromSelector, toSelector string) (Result, error) {
	if err := validateSource(from); err != nil {
		return Result{}, fmt.Errorf("invalid from snapshot: %w", err)
	}
	if err := validateSource(to); err != nil {
		return Result{}, fmt.Errorf("invalid to snapshot: %w", err)
	}
	if from.InstanceID != to.InstanceID {
		return Result{}, fmt.Errorf("snapshot subjects differ")
	}
	if from.Canonical.SchemaVersion != to.Canonical.SchemaVersion || from.Canonical.Profile != to.Canonical.Profile {
		return Result{}, fmt.Errorf("snapshot canonical contracts differ")
	}
	if fromSelector == "" || toSelector == "" {
		return Result{}, fmt.Errorf("snapshot selectors are required")
	}

	timestamp := to.CompletedAt.UTC()
	result := Result{
		SchemaName: SchemaName, SchemaVersion: SchemaVersion, EngineVersion: EngineVersion,
		ComparisonID: stableID("comparison", from.SnapshotID, to.SnapshotID, from.Canonical.SchemaVersion, from.Canonical.Profile),
		SubjectID:    from.InstanceID, InventorySchema: from.Canonical.SchemaVersion,
		InventoryProfile: from.Canonical.Profile, ComparisonTimestamp: timestamp,
		From:    SnapshotReference{Selector: fromSelector, SnapshotID: from.SnapshotID, Completed: from.CompletedAt.UTC(), Status: from.Status},
		To:      SnapshotReference{Selector: toSelector, SnapshotID: to.SnapshotID, Completed: timestamp, Status: to.Status},
		Changes: []ChangeRecord{},
		Metadata: map[string]string{
			"comparison_scope": "canonical-layer-resource-fact-v1",
			"from_status":      string(from.Status),
			"to_status":        string(to.Status),
		},
	}

	fromLayers, err := layerIndex(from.Canonical.Layers)
	if err != nil {
		return Result{}, fmt.Errorf("invalid from layers: %w", err)
	}
	toLayers, err := layerIndex(to.Canonical.Layers)
	if err != nil {
		return Result{}, fmt.Errorf("invalid to layers: %w", err)
	}
	for _, layerID := range sortedUnion(fromLayers, toLayers) {
		fromLayer, fromOK := fromLayers[layerID]
		toLayer, toOK := toLayers[layerID]
		result.Changes = append(result.Changes, compareValue(
			result.ComparisonID, layerID, layerID, layerPath(layerID, "status"),
			categoryValue(fromLayer.Status, fromOK), categoryValue(toLayer.Status, toOK),
			timestamp, map[string]string{"object_kind": "layer"},
		))
		fromResources, err := resourceIndex(fromLayer, fromOK)
		if err != nil {
			return Result{}, fmt.Errorf("invalid from layer %s: %w", layerID, err)
		}
		toResources, err := resourceIndex(toLayer, toOK)
		if err != nil {
			return Result{}, fmt.Errorf("invalid to layer %s: %w", layerID, err)
		}
		for _, resourceID := range sortedUnion(fromResources, toResources) {
			fromResource, resourceFromOK := fromResources[resourceID]
			toResource, resourceToOK := toResources[resourceID]
			result.Changes = append(result.Changes, compareValue(
				result.ComparisonID, layerID, resourceID, resourcePath(layerID, resourceID, "kind"),
				stringValue(fromResource.Kind, resourceFromOK), stringValue(toResource.Kind, resourceToOK),
				timestamp, map[string]string{"object_kind": "resource"},
			))
			for _, factName := range sortedFactUnion(fromResource, resourceFromOK, toResource, resourceToOK) {
				previous := factValue(fromResource.Facts[factName], resourceFromOK && factExists(fromResource.Facts, factName))
				current := factValue(toResource.Facts[factName], resourceToOK && factExists(toResource.Facts, factName))
				result.Changes = append(result.Changes, compareValue(
					result.ComparisonID, layerID, resourceID, factPath(layerID, resourceID, factName),
					previous, current, timestamp,
					map[string]string{"object_kind": "fact", "fact_name": factName},
				))
			}
		}
	}
	sort.Slice(result.Changes, func(i, j int) bool {
		a, b := result.Changes[i], result.Changes[j]
		if a.Layer != b.Layer {
			return a.Layer < b.Layer
		}
		if a.ObjectID != b.ObjectID {
			return a.ObjectID < b.ObjectID
		}
		if a.Path != b.Path {
			return a.Path < b.Path
		}
		return a.Type < b.Type
	})
	for _, record := range result.Changes {
		switch record.Type {
		case Added:
			result.Counts.Added++
		case Removed:
			result.Counts.Removed++
		case Modified:
			result.Counts.Modified++
		case Unchanged:
			result.Counts.Unchanged++
		}
	}
	if err := Validate(result); err != nil {
		return Result{}, err
	}
	return result, nil
}

func Validate(result Result) error {
	if result.SchemaName != SchemaName || result.SchemaVersion != SchemaVersion || result.EngineVersion != EngineVersion {
		return fmt.Errorf("unsupported comparison contract")
	}
	if result.ComparisonID == "" || result.SubjectID == "" || result.InventorySchema == "" || result.InventoryProfile == "" ||
		result.ComparisonTimestamp.IsZero() || result.From.Selector == "" || result.To.Selector == "" {
		return fmt.Errorf("incomplete comparison envelope")
	}
	ids := map[string]bool{}
	counts := Counts{}
	lastKey := ""
	for _, record := range result.Changes {
		if record.ID == "" || ids[record.ID] || record.Layer == "" || record.ObjectID == "" || record.Path == "" ||
			!record.ComparisonTimestamp.Equal(result.ComparisonTimestamp) {
			return fmt.Errorf("invalid change record")
		}
		ids[record.ID] = true
		if _, err := json.Marshal(record.Previous); err != nil {
			return fmt.Errorf("invalid previous typed value: %w", err)
		}
		if _, err := json.Marshal(record.Current); err != nil {
			return fmt.Errorf("invalid current typed value: %w", err)
		}
		key := record.Layer + "\x00" + record.ObjectID + "\x00" + record.Path + "\x00" + string(record.Type)
		if lastKey != "" && key <= lastKey {
			return fmt.Errorf("change records are not uniquely ordered")
		}
		lastKey = key
		switch record.Type {
		case Added:
			if record.Previous != nil || record.Current == nil {
				return fmt.Errorf("invalid added record")
			}
			counts.Added++
		case Removed:
			if record.Previous == nil || record.Current != nil {
				return fmt.Errorf("invalid removed record")
			}
			counts.Removed++
		case Modified:
			if record.Previous == nil || record.Current == nil || valuesEqual(record.Previous, record.Current) {
				return fmt.Errorf("invalid modified record")
			}
			counts.Modified++
		case Unchanged:
			if record.Previous == nil || record.Current == nil || !valuesEqual(record.Previous, record.Current) {
				return fmt.Errorf("invalid unchanged record")
			}
			counts.Unchanged++
		default:
			return fmt.Errorf("invalid change type")
		}
	}
	if counts != result.Counts {
		return fmt.Errorf("comparison counts differ")
	}
	return nil
}

func validateSource(snapshot inventory.Snapshot) error {
	if err := inventory.Validate(snapshot); err != nil {
		return err
	}
	if snapshot.Canonical.SchemaName != inventory.CanonicalSchemaName || snapshot.Canonical.SchemaVersion != inventory.SchemaVersion ||
		snapshot.Canonical.Profile == "" || snapshot.Canonical.SubjectID == "" {
		return fmt.Errorf("canonical inventory is required")
	}
	return nil
}

func compareValue(comparisonID, layer, objectID, path string, previous, current *TypedValue, timestamp time.Time, metadata map[string]string) ChangeRecord {
	changeType := Unchanged
	switch {
	case previous == nil && current != nil:
		changeType = Added
	case previous != nil && current == nil:
		changeType = Removed
	case !valuesEqual(previous, current):
		changeType = Modified
	}
	record := ChangeRecord{
		Layer: layer, ObjectID: objectID, Path: path, Type: changeType,
		Previous: previous, Current: current, ComparisonTimestamp: timestamp, Metadata: metadata,
	}
	previousJSON, _ := json.Marshal(previous)
	currentJSON, _ := json.Marshal(current)
	record.ID = stableID("record", comparisonID, layer, objectID, path, string(changeType), string(previousJSON), string(currentJSON))
	return record
}

func valuesEqual(a, b *TypedValue) bool {
	left, leftErr := json.Marshal(a)
	right, rightErr := json.Marshal(b)
	return leftErr == nil && rightErr == nil && string(left) == string(right)
}

func factValue(fact inventory.CanonicalFact, exists bool) *TypedValue {
	if !exists {
		return nil
	}
	return &TypedValue{
		Type: fact.ValueType, Value: fact.Value, Unit: fact.Unit, Quality: fact.Quality,
		Sensitivity: fact.Sensitivity, ReasonCode: fact.ReasonCode,
	}
}

func stringValue(value string, exists bool) *TypedValue {
	if !exists {
		return nil
	}
	return &TypedValue{Type: "string", Value: value}
}

func categoryValue(value inventory.CategoryStatus, exists bool) *TypedValue {
	if !exists {
		return nil
	}
	return &TypedValue{Type: "string", Value: string(value)}
}

func layerIndex(layers []inventory.Layer) (map[string]inventory.Layer, error) {
	result := make(map[string]inventory.Layer, len(layers))
	for _, layer := range layers {
		if layer.LayerID == "" || layer.ContractVersion != inventory.ContractVersionForLayer {
			return nil, fmt.Errorf("invalid layer identity")
		}
		if _, exists := result[layer.LayerID]; exists {
			return nil, fmt.Errorf("duplicate layer %s", layer.LayerID)
		}
		result[layer.LayerID] = layer
	}
	return result, nil
}

func resourceIndex(layer inventory.Layer, exists bool) (map[string]inventory.Resource, error) {
	result := map[string]inventory.Resource{}
	if !exists {
		return result, nil
	}
	for _, resource := range layer.Resources {
		if resource.ResourceID == "" || resource.LayerID != layer.LayerID {
			return nil, fmt.Errorf("invalid resource identity")
		}
		if _, duplicate := result[resource.ResourceID]; duplicate {
			return nil, fmt.Errorf("duplicate resource %s", resource.ResourceID)
		}
		result[resource.ResourceID] = resource
	}
	return result, nil
}

func sortedUnion[T any](a, b map[string]T) []string {
	seen := map[string]bool{}
	for key := range a {
		seen[key] = true
	}
	for key := range b {
		seen[key] = true
	}
	result := make([]string, 0, len(seen))
	for key := range seen {
		result = append(result, key)
	}
	sort.Strings(result)
	return result
}

func sortedFactUnion(a inventory.Resource, aExists bool, b inventory.Resource, bExists bool) []string {
	left, right := map[string]inventory.CanonicalFact{}, map[string]inventory.CanonicalFact{}
	if aExists {
		left = a.Facts
	}
	if bExists {
		right = b.Facts
	}
	return sortedUnion(left, right)
}

func factExists(facts map[string]inventory.CanonicalFact, name string) bool {
	_, exists := facts[name]
	return exists
}

func stableID(domain string, parts ...string) string {
	hash := sha256.New()
	hash.Write([]byte("qwsg.comparison/" + EngineVersion + "/" + domain))
	for _, part := range parts {
		hash.Write([]byte{0})
		hash.Write([]byte(part))
	}
	return hex.EncodeToString(hash.Sum(nil))
}

func layerPath(layer, field string) string {
	return "/layers/" + escape(layer) + "/" + escape(field)
}

func resourcePath(layer, resource, field string) string {
	return "/layers/" + escape(layer) + "/resources/" + escape(resource) + "/" + escape(field)
}

func factPath(layer, resource, fact string) string {
	return "/layers/" + escape(layer) + "/resources/" + escape(resource) + "/facts/" + escape(fact)
}

func escape(value string) string {
	return strings.ReplaceAll(strings.ReplaceAll(value, "~", "~0"), "/", "~1")
}
