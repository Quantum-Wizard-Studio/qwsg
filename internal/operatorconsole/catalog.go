package operatorconsole

import "fmt"

type Catalog map[string]string

var catalogs = map[string]Catalog{
	"en": {
		"title": "QWSG Operator Console", "condition": "Server condition", "attention": "Attention", "attention_summary": "Additional concerns summarized: %d correlated, %d omitted", "changes": "Changes", "alerts": "Alerts", "guardian": "Guardian", "evidence": "Evidence", "recommendation": "Recommended action", "navigation": "Views", "details": "Evidence details", "help": "Help", "none": "none", "empty": "No items", "keys": "j/k move, Enter select, b back, r refresh, h help, q quit", "refresh_failed": "Refresh failed; the last valid view is retained.", "refresh_invalid": "Refresh returned invalid data; the last valid view is retained.", "state_corrupt": "Saved current state is corrupt and was rejected.", "state_incompatible": "Saved current state uses an unsupported version.", "state_permission": "Saved current state is not private and was rejected.", "state_unsafe": "Saved current state path is unsafe and was rejected.", "state_unreadable": "Saved current state cannot be read safely.", "source": "Source", "observed": "Observed", "summary": "Summary",
	},
	"hu": {
		"title": "QWSG Operátori Konzol", "condition": "Szerver állapota", "attention": "Figyelmet igényel", "attention_summary": "További összegzett ügyek: %d korrelált, %d elhagyott", "changes": "Változások", "alerts": "Riasztások", "guardian": "Guardian", "evidence": "Bizonyíték", "recommendation": "Javasolt lépés", "navigation": "Nézetek", "details": "Bizonyíték részletei", "help": "Súgó", "none": "nincs", "empty": "Nincs megjeleníthető elem", "keys": "j/k mozgás, Enter kiválasztás, b vissza, r frissítés, h súgó, q kilépés", "refresh_failed": "A frissítés sikertelen; az utolsó érvényes nézet megmaradt.", "refresh_invalid": "A frissítés érvénytelen adatot adott; az utolsó érvényes nézet megmaradt.", "state_corrupt": "A mentett aktuális állapot sérült, ezért elutasítottuk.", "state_incompatible": "A mentett aktuális állapot verziója nem támogatott.", "state_permission": "A mentett aktuális állapot nem privát, ezért elutasítottuk.", "state_unsafe": "A mentett aktuális állapot útvonala nem biztonságos.", "state_unreadable": "A mentett aktuális állapot nem olvasható biztonságosan.", "source": "Forrás", "observed": "Megfigyelve", "summary": "Összegzés",
	},
}

var values = map[string]map[string]string{
	"en": {"healthy": "healthy", "degraded": "degraded", "critical": "critical", "unknown": "unknown", "unavailable": "unavailable", "none": "none", "review": "review needed", "urgent": "urgent", "running": "running", "starting": "starting", "stopping": "stopping", "stopped": "stopped", "failed": "failed", "not_observed": "not observed", "current": "current", "stale": "stale", "invalid": "invalid", "complete": "complete", "partial": "partial", "unsupported": "unsupported", "missing": "missing", "inspect_attention": "inspect attention", "review_changes": "review changes", "run_fresh_check": "run qwsg observe", "inspect_evidence": "inspect incomplete evidence", "inspect_failed_operation": "inspect failed operation", "verify_guardian_operation": "verify Guardian operation", "no_action": "no action", "alert_evaluation_failed": "Alert evaluation failed", "notification_planning_failed": "Notification planning failed", "notification_cycle_failed": "Notification delivery failed", "notification_delivery_failed": "Notification delivery failed", "scheduler_cycle_failed": "Scheduler cycle failed", "runtime_timeout": "Runtime timed out", "runtime_cancelled": "Runtime was cancelled", "runtime_not_completed": "Runtime did not complete"},
	"hu": {"healthy": "egészséges", "degraded": "korlátozott", "critical": "kritikus", "unknown": "ismeretlen", "unavailable": "nem elérhető", "none": "nincs", "review": "ellenőrzés szükséges", "urgent": "sürgős", "running": "fut", "starting": "indul", "stopping": "leáll", "stopped": "leállítva", "failed": "hibás", "not_observed": "nincs megfigyelve", "current": "friss", "stale": "elavult", "invalid": "érvénytelen", "complete": "teljes", "partial": "részleges", "unsupported": "nem támogatott", "missing": "hiányzik", "inspect_attention": "vizsgáld meg a figyelmet igénylő elemeket", "review_changes": "tekintsd át a változásokat", "run_fresh_check": "futtasd a qwsg observe parancsot", "inspect_evidence": "vizsgáld meg a hiányos bizonyítékot", "inspect_failed_operation": "vizsgáld meg a sikertelen műveletet", "verify_guardian_operation": "ellenőrizd a Guardian működését", "no_action": "nincs teendő", "alert_evaluation_failed": "A riasztáskiértékelés sikertelen", "notification_planning_failed": "Az értesítéstervezés sikertelen", "notification_cycle_failed": "Az értesítéskézbesítés sikertelen", "notification_delivery_failed": "Az értesítéskézbesítés sikertelen", "scheduler_cycle_failed": "Az ütemezőciklus sikertelen", "runtime_timeout": "A Runtime időtúllépés miatt leállt", "runtime_cancelled": "A Runtime futása megszakadt", "runtime_not_completed": "A Runtime nem fejeződött be"},
}

func text(locale, token string) string {
	if value, ok := catalogs[locale][token]; ok {
		return value
	}
	if value, ok := values[locale][token]; ok {
		return value
	}
	return fmt.Sprintf("[%s]", safe(token))
}
