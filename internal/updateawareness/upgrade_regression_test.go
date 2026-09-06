package updateawareness

import (
	"context"
	"reflect"
	"testing"
	"time"

	"quantumwizard.hu/qwsg/internal/installation"
	"quantumwizard.hu/qwsg/internal/releasediscovery"
	"quantumwizard.hu/qwsg/internal/update"
)

func TestUpgradeAndLegacy304CacheRequireFreshEvaluation(t *testing.T) {
	for _, legacy := range []bool{false, true} {
		t.Run(map[bool]string{false: "installed_changed", true: "legacy_stale_relation"}[legacy], func(t *testing.T) {
			checker := &fakeChecker{result: result("1.3.0", update.Equal, releasediscovery.CompatibilityNotApplicable)}
			checker.result.Evaluation.InstalledVersion = "1.3.0"
			m := managerFixture(t, checker)
			before := validState(t)
			if legacy {
				before.Installed = installed("1.3.0")
				before.LastSuccess.Installed = before.Installed
				var err error
				before, err = Normalize(before)
				if err != nil {
					t.Fatal(err)
				}
			}
			if err := m.Store.Publish(before); err != nil {
				t.Fatal(err)
			}
			m.Classify = func() installation.Result {
				return installation.Result{State: installation.VerifiedSupported, Version: "1.3.0"}
			}
			m.Now = func() time.Time { return testTime.Add(time.Hour) }
			got, err := m.Check(context.Background())
			if err != nil || got.Status != Current || got.LastSuccess.Relation != update.Equal {
				t.Fatalf("fresh evaluation=%+v err=%v", got, err)
			}
			if checker.requests[0].Validators != (releasediscovery.Validators{}) {
				t.Fatal("stale validators sent after upgrade")
			}
			// Once current/equal has been authenticated, ordinary 304 reuse is safe.
			checker.result = releasediscovery.CheckResult{NotModified: true, Source: releasediscovery.SourceEvidence{SourceID: m.SourceID, TransportAuthenticated: true}}
			got, err = m.Check(context.Background())
			if err != nil || got.Status != Current || got.LastSuccess.Relation != update.Equal || checker.requests[1].Validators.ETag == "" {
				t.Fatal("current 304 reuse failed")
			}
			// A server may not invent 304 for a cache unusable for the installed version.
			if _, err = NewSuccess(&before, checker.result, m.SourceID, m.Channel, installed("1.3.0"), testTime.Add(time.Hour), time.Hour); err == nil {
				t.Fatal("unsafe 304 accepted")
			}
		})
	}
}

func TestUpgradeFailurePreservesNotificationAndRollbackWatermark(t *testing.T) {
	checker := &fakeChecker{err: &releasediscovery.ContractError{Category: releasediscovery.SourceTimeout}}
	m := managerFixture(t, checker)
	before := validState(t)
	if err := m.Store.Publish(before); err != nil {
		t.Fatal(err)
	}
	o := before.LastSuccess
	id := NotificationIdentity(before.SourceID, before.Channel, o.ReleaseVersion, o.ArtifactSHA256, o.Authenticity.KeyID)
	if err := m.Store.RecordNotification(id, testTime.Add(time.Minute)); err != nil {
		t.Fatal(err)
	}
	before, err := m.Store.Load()
	if err != nil {
		t.Fatal(err)
	}
	m.Classify = func() installation.Result {
		return installation.Result{State: installation.VerifiedSupported, Version: "1.3.0"}
	}
	m.Now = func() time.Time { return testTime.Add(time.Hour) }
	failed, err := m.Check(context.Background())
	if err == nil || failed.Status != Unknown || failed.LastSuccess == nil || !reflect.DeepEqual(failed.LastNotification, before.LastNotification) {
		t.Fatal("failure discarded history or claimed a valid relation")
	}
	// Restart from disk: preserve the prior authenticated generated-at watermark.
	failed, err = m.Store.Load()
	if err != nil {
		t.Fatal(err)
	}
	checker.err = nil
	checker.result = result("1.3.0", update.Equal, releasediscovery.CompatibilityNotApplicable)
	checker.result.Evaluation.InstalledVersion = "1.3.0"
	checker.result.IndexGeneratedAt = "2026-08-30T10:00:00Z"
	if _, err = m.Check(context.Background()); releasediscovery.FailureOf(err) != releasediscovery.MetadataRollback {
		t.Fatal("upgrade failure lost anti-rollback watermark")
	}
	checker.result.IndexGeneratedAt = "2026-08-30T11:00:00Z"
	current, err := m.Check(context.Background())
	if err != nil || current.Status != Current || !reflect.DeepEqual(current.LastNotification, before.LastNotification) {
		t.Fatal("successful upgrade refresh lost deduplication")
	}
	if len(checker.requests) != 3 {
		t.Fatal("unexpected requests")
	}
	for _, request := range checker.requests {
		if request.Validators != (releasediscovery.Validators{}) {
			t.Fatal("historical cache reused")
		}
	}
}
