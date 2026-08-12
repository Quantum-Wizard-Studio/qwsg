package guardian

import (
	"errors"
	"os"
	"path/filepath"
	"testing"
	"time"

	"quantumwizard.hu/qwsg/internal/alert"
	"quantumwizard.hu/qwsg/internal/app"
	"quantumwizard.hu/qwsg/internal/notification"
	"quantumwizard.hu/qwsg/internal/operatorstate"
	"quantumwizard.hu/qwsg/internal/presentationmodel"
	"quantumwizard.hu/qwsg/internal/runtime"
	"quantumwizard.hu/qwsg/internal/runtimeservice"
)

func TestCheckpointRoundTripIntegrityAndInstanceLock(t *testing.T) {
	root := filepath.Join(t.TempDir(), "guardian")
	store, err := OpenStore(root)
	if err != nil {
		t.Fatal(err)
	}
	value := checkpoint("generation.test")
	if err = store.Save(value); err != nil {
		t.Fatal(err)
	}
	loaded, err := store.Load()
	if err != nil || loaded.Generation != value.Generation {
		t.Fatalf("checkpoint round trip: %#v %v", loaded, err)
	}
	info, _ := os.Stat(filepath.Join(root, "checkpoint.json"))
	if info.Mode().Perm() != 0600 {
		t.Fatalf("checkpoint mode %o", info.Mode().Perm())
	}
	lock, err := Acquire(root, "owner.one")
	if err != nil {
		t.Fatal(err)
	}
	if _, err = Acquire(root, "owner.two"); !errors.Is(err, ErrActive) {
		t.Fatalf("second instance: %v", err)
	}
	if err = lock.Release(); err != nil {
		t.Fatal(err)
	}
	document, _ := os.ReadFile(filepath.Join(root, "checkpoint.json"))
	document[len(document)/2] ^= 1
	if err = os.WriteFile(filepath.Join(root, "checkpoint.json"), document, 0600); err != nil {
		t.Fatal(err)
	}
	if _, err = store.Load(); !errors.Is(err, ErrCheckpoint) {
		t.Fatalf("corrupt checkpoint accepted: %v", err)
	}
}

func TestExitReportIsGenerationCorrelatedAndDemotionOnly(t *testing.T) {
	root := t.TempDir()
	if err := os.Chmod(root, 0700); err != nil {
		t.Fatal(err)
	}
	checkpoints, _ := OpenStore(filepath.Join(root, "guardian"))
	value := checkpoint("generation.active")
	if err := checkpoints.Save(value); err != nil {
		t.Fatal(err)
	}
	current, _ := operatorstate.Open(root)
	at := time.Date(2099, 1, 2, 3, 4, 5, 0, time.UTC)
	service := runtimeservice.NewState("qwsg.guardian.local")
	overview, err := app.ProjectGuardianLifecycle(service, at, time.Minute)
	if err != nil {
		t.Fatal(err)
	}
	stored, err := operatorstate.Normalize(operatorstate.State{ObservedAt: at, PublishedAt: at, FreshUntil: at.Add(time.Minute), Coverage: operatorstate.CoverageGuardianOperation, Provenance: operatorstate.Provenance{DefinitionID: "configuration.test", ExecutionID: service.ID, Profile: "guardian", Source: "live", Stages: []string{"runtime_service"}, Reason: operatorstate.PublicationGuardian, ApplicationVersion: "test"}, Overview: overview})
	if err != nil {
		t.Fatal(err)
	}
	if err = current.Publish(stored); err != nil {
		t.Fatal(err)
	}
	if err = ReportExit(checkpoints, current, "different.generation", "signal", at.Add(time.Second), time.Minute); err != nil {
		t.Fatal(err)
	}
	unchanged, _ := current.Load()
	if unchanged.Overview.Guardian != presentationmodel.GuardianStarting {
		t.Fatal("mismatched generation changed state")
	}
	if err = ReportExit(checkpoints, current, value.Generation, "signal", at.Add(2*time.Second), time.Minute); err != nil {
		t.Fatal(err)
	}
	demoted, _ := current.Load()
	if demoted.Overview.Guardian != presentationmodel.GuardianDegraded || demoted.Overview.Condition == presentationmodel.Healthy {
		t.Fatalf("unsafe exit projection: %#v", demoted.Overview)
	}
}

func TestExitEvidenceAcceptsSystemdTokens(t *testing.T) {
	if !token("7d43c574f6f94a17a242cc25ba097d28") {
		t.Fatal("systemd invocation id rejected")
	}
}

func checkpoint(generation string) Checkpoint {
	configurationID := "configuration.test"
	return Checkpoint{SchemaName: "qwsg.guardian-checkpoint", SchemaVersion: SchemaVersion, ModelVersion: ModelVersion, ServiceID: "qwsg.guardian.local", ConfigurationID: configurationID, Generation: generation, Active: true, RuntimeState: runtime.NewState(), AlertState: alert.NewState(configurationID), NotificationQueueState: notification.NewQueueState()}
}
