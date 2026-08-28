package changenotification

import (
	"errors"
	"strings"
	"testing"
	"time"
)

type fakeSender struct {
	messages []Message
	fail     bool
}

func (f *fakeSender) Send(subject, body string) error {
	f.messages = append(f.messages, Message{subject, body})
	if f.fail {
		return errors.New("transport")
	}
	return nil
}

func TestRequiredEventsLocalizationRedactionAndIdempotency(t *testing.T) {
	now := time.Date(2026, 8, 28, 9, 0, 0, 0, time.UTC)
	for _, typ := range []EventType{Install, Update, Rollback, VersionChanged, ConfigurationChanged, GuardianChanged} {
		for _, locale := range []string{"en", "hu", "de"} {
			e := Event{ID: string(typ) + locale, Host: "server.example", Type: typ, Result: Success, PreviousVersion: "1.2.0-rc.2", NewVersion: "1.2.0-rc.6", At: now}
			f := &fakeSender{}
			d := &Dispatcher{}
			if got := d.Deliver(true, locale, e, f); got != DeliveryAccepted {
				t.Fatalf("%s %s: %s", typ, locale, got)
			}
			if len(f.messages) != 1 || !strings.Contains(f.messages[0].Subject, "[QWSG]") || !strings.Contains(f.messages[0].Body, "1.2.0-rc.2") || !strings.Contains(f.messages[0].Body, "1.2.0-rc.6") {
				t.Fatalf("bad message: %#v", f.messages)
			}
			if got := d.Deliver(true, locale, e, f); got != DeliveryDuplicate || len(f.messages) != 1 {
				t.Fatalf("duplicate: %s %d", got, len(f.messages))
			}
		}
	}
	e := Event{ID: "failure", Host: "server.example", Type: Update, Result: Failed, Reason: "SMTP password=unsafe", At: now, ActionRequired: true}
	f := &fakeSender{}
	if (&Dispatcher{}).Deliver(true, "en", e, f) != DeliveryAccepted || !strings.Contains(f.messages[0].Body, "redacted") || strings.Contains(strings.ToLower(f.messages[0].Body), "password") {
		t.Fatal("secret not redacted")
	}
}

func TestDisabledAndTransportFailurePreserveOperationResult(t *testing.T) {
	e := Event{ID: "update-1", Host: "server.example", Type: Update, Result: Success, At: time.Now().UTC()}
	f := &fakeSender{}
	if (&Dispatcher{}).Deliver(false, "en", e, f) != DeliveryDisabled || len(f.messages) != 0 {
		t.Fatal("disabled delivery attempted")
	}
	f.fail = true
	if (&Dispatcher{}).Deliver(true, "en", e, f) != DeliveryFailed || e.Result != Success {
		t.Fatal("transport changed operation result")
	}
}

func TestUpdateAndRollbackFailureDirection(t *testing.T) {
	now := time.Date(2026, 8, 28, 9, 0, 0, 0, time.UTC)
	for _, e := range []Event{
		{ID: "update-failed", Host: "server.example", Type: Update, Result: Failed, PreviousVersion: "1.2.0-rc.2", NewVersion: "1.2.0-rc.2", Reason: "update_failed", At: now, ActionRequired: true},
		{ID: "rollback-failed", Host: "server.example", Type: Rollback, Result: Failed, PreviousVersion: "1.2.0-rc.6", NewVersion: "1.2.0-rc.2", Reason: "rollback_failed", At: now, ActionRequired: true},
	} {
		f := &fakeSender{}
		if (&Dispatcher{}).Deliver(true, "en", e, f) != DeliveryAccepted {
			t.Fatal("failure event not delivered")
		}
		if !strings.Contains(f.messages[0].Subject, "FAILED") || !strings.Contains(f.messages[0].Body, e.PreviousVersion) || !strings.Contains(f.messages[0].Body, e.NewVersion) {
			t.Fatalf("direction missing: %#v", f.messages[0])
		}
	}
}
