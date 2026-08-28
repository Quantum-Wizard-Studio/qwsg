// Package changenotification defines privacy-safe, localized notifications for
// QWSG-managed lifecycle and configuration changes. Transport remains owned by
// the existing notification provider adapters.
package changenotification

import (
	"fmt"
	"sort"
	"strings"
	"sync"
	"time"
)

type EventType string
type OperationResult string

const (
	Install              EventType       = "installation"
	Update               EventType       = "update"
	Rollback             EventType       = "rollback"
	VersionChanged       EventType       = "version_changed"
	ConfigurationChanged EventType       = "configuration_changed"
	GuardianChanged      EventType       = "guardian_changed"
	Success              OperationResult = "SUCCESS"
	Failed               OperationResult = "FAILED"
)

type Event struct {
	ID, Host, PreviousVersion, NewVersion, Reason string
	Type                                          EventType
	Result                                        OperationResult
	At                                            time.Time
	ActionRequired                                bool
}

type Message struct{ Subject, Body string }

type DeliveryResult string

const (
	DeliveryAccepted  DeliveryResult = "ACCEPTED"
	DeliveryFailed    DeliveryResult = "FAILED"
	DeliveryDisabled  DeliveryResult = "DISABLED"
	DeliveryDuplicate DeliveryResult = "DUPLICATE"
)

type Sender interface {
	Send(subject, body string) error
}

type Dispatcher struct {
	mu   sync.Mutex
	seen map[string]bool
}

func (d *Dispatcher) Deliver(enabled bool, locale string, event Event, sender Sender) DeliveryResult {
	if !enabled {
		return DeliveryDisabled
	}
	if err := Validate(event); err != nil {
		return DeliveryFailed
	}
	d.mu.Lock()
	if d.seen == nil {
		d.seen = map[string]bool{}
	}
	if d.seen[event.ID] {
		d.mu.Unlock()
		return DeliveryDuplicate
	}
	d.seen[event.ID] = true
	d.mu.Unlock()
	message := Render(locale, event)
	if err := sender.Send(message.Subject, message.Body); err != nil {
		return DeliveryFailed
	}
	return DeliveryAccepted
}

func Validate(e Event) error {
	if e.ID == "" || e.Host == "" || e.At.IsZero() || (e.Result != Success && e.Result != Failed) {
		return fmt.Errorf("invalid change event")
	}
	valid := map[EventType]bool{Install: true, Update: true, Rollback: true, VersionChanged: true, ConfigurationChanged: true, GuardianChanged: true}
	if !valid[e.Type] || strings.ContainsAny(e.ID+e.Host+e.PreviousVersion+e.NewVersion+e.Reason, "\r\n") {
		return fmt.Errorf("invalid change event")
	}
	return nil
}

func Render(locale string, e Event) Message {
	lang := strings.ToLower(strings.SplitN(locale, "-", 2)[0])
	if lang != "hu" && lang != "de" {
		lang = "en"
	}
	labels := catalogs[lang]
	name := labels[string(e.Type)]
	status := labels[strings.ToLower(string(e.Result))]
	subject := fmt.Sprintf("[QWSG] %s %s — %s", name, status, e.Host)
	rows := map[string]string{
		labels["identity"]: "QWSG", labels["host"]: e.Host, labels["event"]: name,
		labels["result"]: string(e.Result), labels["time"]: e.At.UTC().Format(time.RFC3339), labels["operation_id"]: e.ID,
		labels["action_required"]: labels[map[bool]string{true: "yes", false: "no"}[e.ActionRequired]],
	}
	if e.PreviousVersion != "" {
		rows[labels["previous_version"]] = e.PreviousVersion
	}
	if e.NewVersion != "" {
		rows[labels["new_version"]] = e.NewVersion
	}
	if e.Reason != "" {
		rows[labels["reason"]] = sanitizeReason(e.Reason)
	}
	keys := make([]string, 0, len(rows))
	for k := range rows {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	var body strings.Builder
	for _, k := range keys {
		fmt.Fprintf(&body, "%s: %s\n", k, rows[k])
	}
	return Message{Subject: subject, Body: body.String()}
}

func sanitizeReason(v string) string {
	lower := strings.ToLower(v)
	for _, forbidden := range []string{"password", "token", "secret", "credential", "private key", "api key"} {
		if strings.Contains(lower, forbidden) {
			return "redacted"
		}
	}
	if len(v) > 160 {
		return v[:160]
	}
	return v
}

var catalogs = map[string]map[string]string{
	"en": {"installation": "Installation", "update": "Update", "rollback": "Rollback", "version_changed": "Version changed", "configuration_changed": "Configuration changed", "guardian_changed": "Guardian state changed", "success": "completed", "failed": "FAILED", "identity": "Identity", "host": "Host", "event": "Event", "result": "Result", "time": "Timestamp", "operation_id": "Operation ID", "action_required": "Administrator action required", "previous_version": "Previous version", "new_version": "New version", "reason": "Reason", "yes": "yes", "no": "no"},
	"hu": {"installation": "Telepítés", "update": "Frissítés", "rollback": "Visszaállítás", "version_changed": "Verzióváltozás", "configuration_changed": "Konfiguráció módosult", "guardian_changed": "Guardian állapotváltozás", "success": "befejeződött", "failed": "SIKERTELEN", "identity": "Azonosító", "host": "Gép", "event": "Esemény", "result": "Eredmény", "time": "Időbélyeg", "operation_id": "Műveletazonosító", "action_required": "Rendszergazdai beavatkozás szükséges", "previous_version": "Előző verzió", "new_version": "Új verzió", "reason": "Ok", "yes": "igen", "no": "nem"},
	"de": {"installation": "Installation", "update": "Aktualisierung", "rollback": "Rollback", "version_changed": "Versionsänderung", "configuration_changed": "Konfiguration geändert", "guardian_changed": "Guardian-Status geändert", "success": "abgeschlossen", "failed": "FEHLGESCHLAGEN", "identity": "Identität", "host": "Host", "event": "Ereignis", "result": "Ergebnis", "time": "Zeitstempel", "operation_id": "Vorgangs-ID", "action_required": "Administratoraktion erforderlich", "previous_version": "Vorherige Version", "new_version": "Neue Version", "reason": "Grund", "yes": "ja", "no": "nein"},
}
