package updatenotification

import (
	"context"
	"reflect"
	"testing"
	"time"

	"quantumwizard.hu/qwsg/internal/releasediscovery"
	"quantumwizard.hu/qwsg/internal/update"
	"quantumwizard.hu/qwsg/internal/updateawareness"
)

func TestDeliveredReleaseDedupSurvivesUpgradeFailureRestartAndRollback(t *testing.T) {
	first := awareness(t, "1.3.0", update.Newer, releasediscovery.CompatibilitySupported)
	store := notificationStore(t, first)
	sender := &fakeSender{succeed: true}
	notifier := &Service{Enabled: true, Store: store, Sender: sender, Now: func() time.Time { return notificationTime }}
	if err := notifier.Notify(context.Background(), first); err != nil {
		t.Fatal(err)
	}
	delivered, err := store.Load()
	if err != nil {
		t.Fatal(err)
	}
	upgraded := delivered.Installed
	upgraded.Version = "1.3.0"
	failed, err := updateawareness.NewFailure(&delivered, delivered.SourceID, delivered.Channel, upgraded, notificationTime.Add(time.Minute), "source_timeout")
	if err != nil {
		t.Fatal(err)
	}
	if err = store.Publish(failed); err != nil {
		t.Fatal(err)
	}
	failed, err = store.Load()
	if err != nil {
		t.Fatal(err)
	}
	o := delivered.LastSuccess
	result := releasediscovery.CheckResult{Source: releasediscovery.SourceEvidence{SourceID: delivered.SourceID, TransportAuthenticated: true}, IndexGeneratedAt: o.IndexGeneratedAt, Authenticity: o.Authenticity, Evaluation: releasediscovery.Evaluation{InstalledVersion: "1.3.0", Channel: "stable", Platform: "linux-amd64", Release: releasediscovery.Release{Version: o.ReleaseVersion, PublishedAt: o.ReleasePublishedAt, Status: "active"}, Artifact: releasediscovery.Artifact{Name: o.ArtifactName, SHA256: o.ArtifactSHA256, Size: o.ArtifactSize}, Relation: update.Equal, Compatibility: releasediscovery.CompatibilityNotApplicable, Authenticity: o.Authenticity}}
	current, err := updateawareness.NewSuccess(&failed, result, delivered.SourceID, delivered.Channel, upgraded, notificationTime.Add(2*time.Minute), time.Hour)
	if err != nil || !reflect.DeepEqual(current.LastNotification, delivered.LastNotification) {
		t.Fatal("upgrade lost delivery record")
	}
	if err = store.Publish(current); err != nil {
		t.Fatal(err)
	}
	current, err = store.Load()
	if err != nil {
		t.Fatal(err)
	}
	notifier = &Service{Enabled: true, Store: store, Sender: sender}
	if err = notifier.Notify(context.Background(), current); err != nil {
		t.Fatal(err)
	}
	// Returning to the old installed version must not notify for the same release again.
	result.Evaluation.InstalledVersion = "1.2.0"
	result.Evaluation.Relation = update.Newer
	result.Evaluation.Compatibility = releasediscovery.CompatibilitySupported
	result.Evaluation.MigrationID = "compat-1.2.0-to-1.3.0"
	rolled, err := updateawareness.NewSuccess(&current, result, delivered.SourceID, delivered.Channel, delivered.Installed, notificationTime.Add(3*time.Minute), time.Hour)
	if err != nil {
		t.Fatal(err)
	}
	if err = store.Publish(rolled); err != nil {
		t.Fatal(err)
	}
	rolled, err = store.Load()
	if err != nil {
		t.Fatal(err)
	}
	notifier = &Service{Enabled: true, Store: store, Sender: sender}
	if err = notifier.Notify(context.Background(), rolled); err != nil {
		t.Fatal(err)
	}
	if sender.attempts != 1 {
		t.Fatalf("duplicate delivery attempts=%d", sender.attempts)
	}
}

func TestLegacyStaleEqualStateCannotNotify(t *testing.T) {
	state := awareness(t, "1.3.0", update.Newer, releasediscovery.CompatibilitySupported)
	state.Installed.Version = "1.3.0"
	state.LastSuccess.Installed = state.Installed
	var err error
	state, err = updateawareness.Normalize(state)
	if err != nil {
		t.Fatal(err)
	}
	if _, ok := Eligible(state); ok {
		t.Fatal("legacy stale equal release eligible for notification")
	}
}
