package updatenotification

import (
	"context"
	"errors"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"testing"
	"time"

	"quantumwizard.hu/qwsg/internal/installation"
	"quantumwizard.hu/qwsg/internal/releasediscovery"
	"quantumwizard.hu/qwsg/internal/update"
	"quantumwizard.hu/qwsg/internal/updateawareness"
)

var notificationTime = time.Date(2026, 9, 3, 12, 0, 0, 0, time.UTC)

type fakeSender struct {
	mu       sync.Mutex
	attempts int
	succeed  bool
	subject  string
	body     string
}

func (f *fakeSender) DeliverText(_ context.Context, subject, body string) bool {
	f.mu.Lock()
	defer f.mu.Unlock()
	f.attempts++
	f.subject, f.body = subject, body
	return f.succeed
}

func awareness(t *testing.T, version string, relation update.Relation, compatibility releasediscovery.Compatibility) updateawareness.State {
	t.Helper()
	auth := releasediscovery.AuthenticityEvidence{Scheme: "ed25519", KeyID: "community-test"}
	result := releasediscovery.CheckResult{
		Source:           releasediscovery.SourceEvidence{SourceID: "community-release-index", TransportAuthenticated: true},
		IndexGeneratedAt: "2026-09-03T11:00:00Z",
		Authenticity:     auth,
		Evaluation: releasediscovery.Evaluation{
			InstalledVersion: "1.2.0", Channel: "stable", Platform: "linux-amd64",
			Release:       releasediscovery.Release{Version: version, PublishedAt: "2026-09-03T10:00:00Z", Status: "active"},
			Artifact:      releasediscovery.Artifact{Name: "qwsg-" + version + "-linux-amd64.tar.gz", SHA256: strings.Repeat(string(version[0]), 64), Size: 1234},
			Relation:      relation,
			Compatibility: compatibility,
			MigrationID:   map[bool]string{true: "compat-1.2.0-to-" + version}[compatibility == releasediscovery.CompatibilitySupported],
			Authenticity:  auth,
		},
	}
	state, err := updateawareness.NewSuccess(nil, result, "community-release-index", "stable", updateawareness.Installed{Classification: installation.VerifiedSupported, Version: "1.2.0"}, notificationTime, 48*time.Hour)
	if err != nil {
		t.Fatal(err)
	}
	return state
}

func notificationStore(t *testing.T, state updateawareness.State) *updateawareness.Store {
	t.Helper()
	root := filepath.Join(t.TempDir(), "state")
	if err := os.Mkdir(root, 0700); err != nil {
		t.Fatal(err)
	}
	store, err := updateawareness.Open(root)
	if err != nil || store.Publish(state) != nil {
		t.Fatalf("store setup: %v", err)
	}
	return store
}

func TestEligibilityFailsClosed(t *testing.T) {
	newer := awareness(t, "1.3.0", update.Newer, releasediscovery.CompatibilitySupported)
	if identity, ok := Eligible(newer); !ok || len(identity) != 64 {
		t.Fatalf("authenticated newer release ineligible: %q %t", identity, ok)
	}
	for name, state := range map[string]updateawareness.State{
		"current": awareness(t, "1.2.0", update.Equal, releasediscovery.CompatibilityNotApplicable),
		"older":   awareness(t, "1.1.0", update.Older, releasediscovery.CompatibilityNotApplicable),
	} {
		t.Run(name, func(t *testing.T) {
			if _, ok := Eligible(state); ok {
				t.Fatal("non-newer release eligible")
			}
		})
	}
	untrusted := newer
	untrusted.LastSuccess.TransportAuthenticated = false
	if _, ok := Eligible(untrusted); ok {
		t.Fatal("tampered/untrusted release eligible")
	}
}

func TestSuccessfulDeliveryPersistsAndDeduplicatesAcrossRestart(t *testing.T) {
	state := awareness(t, "1.3.0", update.Newer, releasediscovery.CompatibilitySupported)
	store := notificationStore(t, state)
	sender := &fakeSender{succeed: true}
	service := Service{Enabled: true, Locale: "hu-HU", Store: store, Sender: sender, Now: func() time.Time { return notificationTime.Add(time.Minute) }}
	if err := service.Notify(context.Background(), state); err != nil {
		t.Fatal(err)
	}
	stored, err := store.Load()
	if err != nil || stored.LastNotification == nil || sender.attempts != 1 {
		t.Fatalf("stored=%+v attempts=%d err=%v", stored.LastNotification, sender.attempts, err)
	}
	if !strings.Contains(sender.subject, "Hitelesített") || !strings.Contains(sender.body, "1.2.0") || !strings.Contains(sender.body, "1.3.0") || !strings.Contains(sender.body, "stable") || !strings.Contains(sender.body, "nem automatikus") || strings.Contains(strings.ToLower(sender.body), "hostname") {
		t.Fatalf("unexpected localized message: %q %q", sender.subject, sender.body)
	}
	restarted := Service{Enabled: true, Locale: "en", Store: store, Sender: sender}
	if err = restarted.Notify(context.Background(), stored); err != nil || sender.attempts != 1 {
		t.Fatalf("restart duplicate attempts=%d err=%v", sender.attempts, err)
	}
}

func TestFailureDoesNotDeduplicateAndLaterScheduledAttemptMayRetry(t *testing.T) {
	state := awareness(t, "1.3.0", update.Newer, releasediscovery.CompatibilitySupported)
	store := notificationStore(t, state)
	sender := &fakeSender{}
	service := Service{Enabled: true, Locale: "en", Store: store, Sender: sender}
	if err := service.Notify(context.Background(), state); err == nil {
		t.Fatal("delivery failure accepted")
	}
	stored, err := store.Load()
	if err != nil || stored.LastNotification != nil || sender.attempts != 1 {
		t.Fatalf("failed delivery persisted: %+v attempts=%d err=%v", stored.LastNotification, sender.attempts, err)
	}
	sender.succeed = true
	if err = service.Notify(context.Background(), stored); err != nil {
		t.Fatal(err)
	}
	stored, err = store.Load()
	if err != nil || stored.LastNotification == nil || sender.attempts != 2 {
		t.Fatalf("retry result=%+v attempts=%d err=%v", stored.LastNotification, sender.attempts, err)
	}
}

func TestDifferentAuthenticatedReleaseAndDisabledChannel(t *testing.T) {
	first := awareness(t, "1.3.0", update.Newer, releasediscovery.CompatibilitySupported)
	store := notificationStore(t, first)
	sender := &fakeSender{succeed: true}
	service := Service{Enabled: false, Store: store, Sender: sender}
	if err := service.Notify(context.Background(), first); err != nil || sender.attempts != 0 {
		t.Fatalf("disabled delivery attempted: %d %v", sender.attempts, err)
	}
	service.Enabled = true
	if err := service.Notify(context.Background(), first); err != nil {
		t.Fatal(err)
	}
	previous, _ := store.Load()
	second := awareness(t, "1.4.0", update.Newer, releasediscovery.CompatibilitySupported)
	second.LastNotification = previous.LastNotification
	second, _ = updateawareness.Normalize(second)
	if err := store.Publish(second); err != nil {
		t.Fatal(err)
	}
	if err := service.Notify(context.Background(), second); err != nil || sender.attempts != 2 {
		t.Fatalf("different release not delivered: attempts=%d err=%v", sender.attempts, err)
	}
}

type failingStore struct{}

func (failingStore) RecordNotification(string, time.Time) error { return errors.New("persist") }

func TestPersistenceFailureIsReportedAfterAcceptedDelivery(t *testing.T) {
	state := awareness(t, "1.3.0", update.Newer, releasediscovery.CompatibilitySupported)
	sender := &fakeSender{succeed: true}
	service := &Service{Enabled: true, Store: failingStore{}, Sender: sender}
	if err := service.Notify(context.Background(), state); err == nil || sender.attempts != 1 {
		t.Fatalf("persistence uncertainty hidden: attempts=%d err=%v", sender.attempts, err)
	}
}

type blockingSender struct {
	entered chan<- struct{}
	release <-chan struct{}
}

func (s blockingSender) DeliverText(context.Context, string, string) bool {
	s.entered <- struct{}{}
	<-s.release
	return true
}

func TestConcurrentAttemptsDoNotOverlap(t *testing.T) {
	state := awareness(t, "1.3.0", update.Newer, releasediscovery.CompatibilitySupported)
	store := notificationStore(t, state)
	entered, release := make(chan struct{}, 1), make(chan struct{})
	service := &Service{Enabled: true, Store: store, Sender: blockingSender{entered: entered, release: release}}
	done := make(chan error, 2)
	go func() { done <- service.Notify(context.Background(), state) }()
	<-entered
	go func() { done <- service.Notify(context.Background(), state) }()
	select {
	case <-entered:
		t.Fatal("notification attempts overlapped")
	case <-time.After(20 * time.Millisecond):
	}
	close(release)
	if err := <-done; err != nil {
		t.Fatal(err)
	}
	if err := <-done; err != nil {
		t.Fatal(err)
	}
}
