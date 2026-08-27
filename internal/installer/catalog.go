package installer

import "fmt"

type Language string
type MessageID string

const (
	English   Language = "en"
	Hungarian Language = "hu"
	German    Language = "de"
)

var languages = map[Language]string{English: "English", Hungarian: "Magyar", German: "Deutsch"}

var english = map[MessageID]string{
	"plan.mode": "Mode: %s", "plan.mode.fresh": "fresh installation", "plan.mode.existing": "existing installation / setup",
	"plan.package": "Package files: /usr/local/bin/qwsg, /usr/local/lib/qwsg, /usr/local/share/doc/qwsg", "plan.data": "Operator data: private configuration and state under your account", "plan.service": "Service: systemd user unit qwsg-guardian.service",
	"install.existing": "The verified QWSG package is already installed; package mutation is skipped.", "install.failure": "Package installation failed. No successful package mutation is claimed; review diagnostics before retrying.",
	"configuration.done": "Private operator configuration was initialized safely.", "configuration.failure": "Configuration could not be initialized. Correct the configuration path or value and retry.",
	"notification.later": "You will need the recipient address, SMTP host such as mail.example.com, port, sender address, encryption mode, and—when required—a provider username and password. Obtain these from your hosting or email provider. Store the password only with `qwsg notification credential set --from-file PRIVATE_FILE`, then run preflight and test. This optional step can be completed later.", "notification.skipped": "Email notification was skipped. Guardian can operate without it; rerun `qwsg setup` later.",
	"update_policy.saved": "Update policy saved: %s. Automatic privileged updates remain disabled.", "activation.skipped": "Guardian activation was skipped; rerun `qwsg setup` later.", "activation.failure": "Guardian activation failed. Configuration remains preserved. Run `qwsg readiness`, follow its service guidance, and resume with `qwsg setup`.",
	"readiness.ready": "Readiness: ready. Guardian produced the required operational evidence.", "readiness.partial": "Readiness: incomplete. Installation data is preserved; follow `qwsg readiness` and resume setup after correcting the reported requirement.",
	"summary.version": "Installed QWSG version: %s", "summary.guardian": "Guardian: %s", "summary.active": "active", "summary.inactive": "not activated", "summary.notification": "Notification: optional; configurable later", "summary.policy": "Update policy: %s (automatic privileged updates disabled)", "summary.readiness": "Readiness: %s", "summary.ready": "ready", "summary.partial": "partial (optional capabilities remain unverified or disabled)",
	"title":                "Quantum Wizard Server Guardian — Installation Wizard",
	"stage":                "Current stage",
	"language.prompt":      "Choose the language used for the rest of the installation",
	"preflight":            "Environment and compatibility check",
	"preflight.help":       "QWSG is inspecting the operating system, architecture, user session, and existing QWSG state. This read-only step prevents an unsafe or incompatible installation.",
	"plan":                 "Installation plan",
	"plan.help":            "Review what QWSG intends to change. No privileged installation change has occurred yet.",
	"plan.confirm":         "Continue with this plan? [y/N]",
	"install":              "QWSG package installation",
	"install.help":         "QWSG will install only the verified release-owned program, service unit, and documentation at the listed system paths.",
	"configuration":        "Guardian configuration",
	"configuration.help":   "QWSG will create your private operator configuration and state directories. Safe defaults run Guardian every five minutes.",
	"notification":         "Optional email notification",
	"notification.help":    "Email notification lets Guardian send alerts through an SMTP mail server. SMTP is the outgoing mail service normally provided by your hosting company, workplace, or email provider. You can skip this and configure it later with `qwsg setup`.",
	"notification.prompt":  "Review the values needed for optional email notification now? [y/N]",
	"update_policy":        "Update policy",
	"update_policy.help":   "Manual keeps every update operator-initiated. Notify records your preference to be informed when a future supported notifier is available; it never installs an update automatically.",
	"update_policy.prompt": "Select 1 for manual updates or 2 for update notifications [1]",
	"activation":           "Guardian service activation",
	"activation.help":      "QWSG will enable and start the per-user Guardian service. Starting before login may additionally require administrator-approved lingering.",
	"activation.prompt":    "Activate Guardian now? [Y/n]",
	"readiness":            "Readiness verification",
	"readiness.help":       "QWSG is checking configuration, service state, fresh Guardian evidence, notification status, and update integration.",
	"completion":           "Installation complete",
	"completion.help":      "All selected stages have finished. The summary below records actual Guardian, notification, update-policy, and readiness outcomes.",
	"unsupported":          "Installation cannot continue. Detected: %s. Supported: Ubuntu 24.04 LTS amd64. Use a supported server or consult a future QWSG release for expanded support.",
	"cancelled":            "Installation cancelled before privileged mutation.",
	"invalid":              "That value is not valid. Please use one of the displayed choices.",
	"next":                 "Next commands: `qwsg readiness`, `qwsg update check`, `qwsg setup`.",
}

var hungarian = map[MessageID]string{
	"plan.mode": "Mód: %s", "plan.mode.fresh": "új telepítés", "plan.mode.existing": "meglévő telepítés / beállítás", "plan.package": "Csomagfájlok: /usr/local/bin/qwsg, /usr/local/lib/qwsg, /usr/local/share/doc/qwsg", "plan.data": "Kezelői adatok: privát konfiguráció és állapot a felhasználói fiók alatt", "plan.service": "Szolgáltatás: qwsg-guardian.service systemd felhasználói egység",
	"install.existing": "Az ellenőrzött QWSG-csomag már telepítve van; a csomagmódosítás kimarad.", "install.failure": "A csomagtelepítés sikertelen. A QWSG nem állítja, hogy a csomagmódosítás sikerült; újrapróbálás előtt tekintse át a diagnosztikát.", "configuration.done": "A privát kezelői konfiguráció biztonságosan létrejött.", "configuration.failure": "A konfiguráció nem hozható létre. Javítsa a konfigurációs útvonalat vagy értéket, majd próbálja újra.",
	"notification.later": "Szükség lesz a címzett címére, egy SMTP-gépre (például mail.example.com), portra, feladói címre, titkosítási módra, és szükség esetén szolgáltatói felhasználónévre és jelszóra. Ezeket a tárhely- vagy e-mail-szolgáltató adja meg. A jelszót kizárólag a `qwsg notification credential set --from-file PRIVATE_FILE` paranccsal tárolja, majd futtassa az előellenőrzést és a tesztet. Ez később is elvégezhető.", "notification.skipped": "Az e-mail-értesítés kimaradt. A Guardian nélküle is működhet; később futtassa újra a `qwsg setup` parancsot.",
	"update_policy.saved": "Frissítési szabály mentve: %s. Az automatikus jogosultságemelt frissítés kikapcsolva marad.", "activation.skipped": "A Guardian aktiválása kimaradt; később futtassa újra a `qwsg setup` parancsot.", "activation.failure": "A Guardian aktiválása sikertelen. A konfiguráció megmaradt. Futtassa a `qwsg readiness` parancsot, kövesse a szolgáltatási útmutatást, majd folytassa a `qwsg setup` paranccsal.",
	"readiness.ready": "Üzemkészség: kész. A Guardian előállította a szükséges működési bizonyítékot.", "readiness.partial": "Üzemkészség: részleges. A kötelező Guardian-alap működik, de opcionális képességek nincsenek ellenőrizve vagy engedélyezve.", "summary.version": "Telepített QWSG-verzió: %s", "summary.guardian": "Guardian: %s", "summary.active": "aktív", "summary.inactive": "nincs aktiválva", "summary.notification": "Értesítés: opcionális; később beállítható", "summary.policy": "Frissítési szabály: %s (az automatikus jogosultságemelt frissítés kikapcsolva)", "summary.readiness": "Üzemkészség: %s", "summary.ready": "kész", "summary.partial": "részleges (opcionális képességek nincsenek ellenőrizve vagy engedélyezve)",
	"title": "Quantum Wizard Server Guardian — Telepítési varázsló", "stage": "Jelenlegi szakasz",
	"language.prompt": "Válassza ki a telepítés további részének nyelvét", "preflight": "Környezet- és kompatibilitás-ellenőrzés",
	"preflight.help": "A QWSG ellenőrzi az operációs rendszert, az architektúrát, a felhasználói munkamenetet és a meglévő QWSG-állapotot. Ez az írásvédett lépés megelőzi a nem biztonságos vagy nem kompatibilis telepítést.",
	"plan":           "Telepítési terv", "plan.help": "Tekintse át, mit kíván módosítani a QWSG. Jogosultságot igénylő telepítési változás még nem történt.", "plan.confirm": "Folytatja ezzel a tervvel? [i/N]",
	"install": "QWSG-csomag telepítése", "install.help": "A QWSG kizárólag az ellenőrzött kiadáshoz tartozó programot, szolgáltatásegységet és dokumentációt telepíti a felsorolt rendszerútvonalakra.",
	"configuration": "Guardian beállítása", "configuration.help": "A QWSG létrehozza a privát kezelői konfigurációt és állapotkönyvtárakat. A biztonságos alapérték ötpercenként futtatja a Guardiant.",
	"notification": "Opcionális e-mail-értesítés", "notification.help": "Az e-mail-értesítés lehetővé teszi, hogy a Guardian SMTP levelezőszerveren keresztül riasztásokat küldjön. Az SMTP kimenő levelezési szolgáltatást általában a tárhelyszolgáltató, a munkahely vagy az e-mail-szolgáltató biztosítja. Ez kihagyható és később a `qwsg setup` paranccsal beállítható.", "notification.prompt": "Áttekinti most az opcionális e-mail-értesítéshez szükséges adatokat? [i/N]",
	"update_policy": "Frissítési szabály", "update_policy.help": "A kézi mód minden frissítést kezelői indításhoz köt. Az értesítési mód eltárolja, hogy később értesítést kér; soha nem telepít automatikusan frissítést.", "update_policy.prompt": "Válasszon: 1 kézi frissítés, 2 frissítési értesítés [1]",
	"activation": "Guardian szolgáltatás aktiválása", "activation.help": "A QWSG engedélyezi és elindítja a felhasználói Guardian szolgáltatást. Bejelentkezés előtti induláshoz rendszergazda által engedélyezett lingering is szükséges lehet.", "activation.prompt": "Aktiválja most a Guardiant? [I/n]",
	"readiness": "Üzemkészség ellenőrzése", "readiness.help": "A QWSG ellenőrzi a konfigurációt, a szolgáltatás állapotát, a friss Guardian-bizonyítékot, az értesítést és a frissítési integrációt.",
	"completion.help": "Minden kiválasztott szakasz befejeződött. Az összefoglaló a Guardian, az értesítés, a frissítési szabály és az üzemkészség tényleges eredményét mutatja.",
	"completion":      "A telepítés befejeződött", "unsupported": "A telepítés nem folytatható. Észlelt rendszer: %s. Támogatott: Ubuntu 24.04 LTS amd64. Használjon támogatott szervert, vagy ellenőrizze egy későbbi QWSG-kiadás támogatását.", "cancelled": "A telepítés jogosultságot igénylő módosítás előtt megszakadt.", "invalid": "Az érték érvénytelen. Válasszon a megjelenített lehetőségek közül.", "next": "Következő parancsok: `qwsg readiness`, `qwsg update check`, `qwsg setup`.",
}

var german = map[MessageID]string{
	"plan.mode": "Modus: %s", "plan.mode.fresh": "Neuinstallation", "plan.mode.existing": "vorhandene Installation / Einrichtung", "plan.package": "Paketdateien: /usr/local/bin/qwsg, /usr/local/lib/qwsg, /usr/local/share/doc/qwsg", "plan.data": "Bedienerdaten: private Konfiguration und Zustände im Benutzerkonto", "plan.service": "Dienst: systemd-Benutzereinheit qwsg-guardian.service",
	"install.existing": "Das verifizierte QWSG-Paket ist bereits installiert; die Paketänderung wird übersprungen.", "install.failure": "Paketinstallation fehlgeschlagen. QWSG behauptet keine erfolgreiche Paketänderung; prüfen Sie vor einem neuen Versuch die Diagnose.", "configuration.done": "Die private Bedienerkonfiguration wurde sicher initialisiert.", "configuration.failure": "Die Konfiguration konnte nicht initialisiert werden. Korrigieren Sie Pfad oder Wert und versuchen Sie es erneut.",
	"notification.later": "Sie benötigen Empfängeradresse, SMTP-Host wie mail.example.com, Port, Absenderadresse, Verschlüsselungsmodus und bei Bedarf Benutzername und Passwort des Anbieters. Diese Werte erhalten Sie vom Hosting- oder E-Mail-Anbieter. Speichern Sie das Passwort nur mit `qwsg notification credential set --from-file PRIVATE_FILE`; führen Sie danach Vorprüfung und Test aus. Dieser optionale Schritt ist später möglich.", "notification.skipped": "E-Mail-Benachrichtigung wurde übersprungen. Guardian kann ohne sie arbeiten; führen Sie später `qwsg setup` aus.",
	"update_policy.saved": "Aktualisierungsrichtlinie gespeichert: %s. Automatische privilegierte Aktualisierungen bleiben deaktiviert.", "activation.skipped": "Guardian-Aktivierung wurde übersprungen; führen Sie später `qwsg setup` aus.", "activation.failure": "Guardian-Aktivierung fehlgeschlagen. Die Konfiguration bleibt erhalten. Führen Sie `qwsg readiness` aus, folgen Sie den Diensthinweisen und setzen Sie mit `qwsg setup` fort.",
	"readiness.ready": "Bereitschaft: bereit. Guardian hat den erforderlichen Betriebsnachweis erzeugt.", "readiness.partial": "Bereitschaft: teilweise. Die erforderliche Guardian-Basis funktioniert; optionale Fähigkeiten sind nicht verifiziert oder deaktiviert.", "summary.version": "Installierte QWSG-Version: %s", "summary.guardian": "Guardian: %s", "summary.active": "aktiv", "summary.inactive": "nicht aktiviert", "summary.notification": "Benachrichtigung: optional; später konfigurierbar", "summary.policy": "Aktualisierungsrichtlinie: %s (automatische privilegierte Aktualisierungen deaktiviert)", "summary.readiness": "Bereitschaft: %s", "summary.ready": "bereit", "summary.partial": "teilweise (optionale Fähigkeiten sind nicht verifiziert oder deaktiviert)",
	"title": "Quantum Wizard Server Guardian — Installationsassistent", "stage": "Aktueller Abschnitt", "language.prompt": "Wählen Sie die Sprache für die weitere Installation",
	"preflight": "Umgebungs- und Kompatibilitätsprüfung", "preflight.help": "QWSG prüft Betriebssystem, Architektur, Benutzersitzung und vorhandenen QWSG-Zustand. Dieser schreibgeschützte Schritt verhindert eine unsichere oder inkompatible Installation.",
	"plan": "Installationsplan", "plan.help": "Prüfen Sie, was QWSG ändern wird. Es wurde noch keine privilegierte Installationsänderung vorgenommen.", "plan.confirm": "Mit diesem Plan fortfahren? [j/N]",
	"install": "QWSG-Paketinstallation", "install.help": "QWSG installiert ausschließlich das verifizierte Programm, die Diensteinheit und Dokumentation der Version in die aufgeführten Systempfade.",
	"configuration": "Guardian-Konfiguration", "configuration.help": "QWSG erstellt Ihre private Bedienerkonfiguration und Zustandsverzeichnisse. Sichere Standardwerte führen Guardian alle fünf Minuten aus.",
	"notification": "Optionale E-Mail-Benachrichtigung", "notification.help": "E-Mail-Benachrichtigungen erlauben Guardian, Warnungen über einen SMTP-Mailserver zu senden. SMTP ist der ausgehende Maildienst Ihres Hosting-, Arbeits- oder E-Mail-Anbieters. Sie können dies überspringen und später mit `qwsg setup` konfigurieren.", "notification.prompt": "Benötigte Werte für optionale E-Mail-Benachrichtigung jetzt ansehen? [j/N]",
	"update_policy": "Aktualisierungsrichtlinie", "update_policy.help": "Manuell bedeutet, dass jede Aktualisierung vom Bediener gestartet wird. Benachrichtigen speichert den Wunsch für einen zukünftigen unterstützten Benachrichtigungsdienst; Aktualisierungen werden niemals automatisch installiert.", "update_policy.prompt": "Wählen Sie 1 für manuell oder 2 für Aktualisierungshinweise [1]",
	"activation": "Guardian-Dienstaktivierung", "activation.help": "QWSG aktiviert und startet den Guardian-Benutzerdienst. Start vor der Anmeldung kann zusätzlich administrativ genehmigtes Lingering erfordern.", "activation.prompt": "Guardian jetzt aktivieren? [J/n]",
	"readiness": "Bereitschaftsprüfung", "readiness.help": "QWSG prüft Konfiguration, Dienstzustand, aktuelle Guardian-Nachweise, Benachrichtigungsstatus und Aktualisierungsintegration.",
	"completion.help": "Alle ausgewählten Abschnitte sind abgeschlossen. Die Zusammenfassung zeigt die tatsächlichen Ergebnisse für Guardian, Benachrichtigung, Aktualisierungsrichtlinie und Bereitschaft.",
	"completion":      "Installation abgeschlossen", "unsupported": "Installation kann nicht fortgesetzt werden. Erkannt: %s. Unterstützt: Ubuntu 24.04 LTS amd64. Verwenden Sie einen unterstützten Server oder prüfen Sie eine spätere QWSG-Version.", "cancelled": "Installation wurde vor privilegierten Änderungen abgebrochen.", "invalid": "Dieser Wert ist ungültig. Verwenden Sie eine der angezeigten Optionen.", "next": "Nächste Befehle: `qwsg readiness`, `qwsg update check`, `qwsg setup`.",
}

type Catalog struct{ Language Language }

func ParseLanguage(value string) (Language, bool) {
	lang := Language(value)
	_, ok := languages[lang]
	return lang, ok
}

func (c Catalog) Text(id MessageID, args ...any) string {
	selected := map[Language]map[MessageID]string{Hungarian: hungarian, German: german}[c.Language]
	value := ""
	if selected != nil {
		value = selected[id]
	}
	if value == "" {
		value = english[id]
	}
	if value == "" {
		value = string(id)
	}
	if len(args) > 0 {
		return fmt.Sprintf(value, args...)
	}
	return value
}

func ValidateCatalogs() error {
	for _, catalog := range []map[MessageID]string{hungarian, german} {
		for id := range english {
			if catalog[id] == "" {
				return fmt.Errorf("missing translation %s", id)
			}
		}
	}
	return nil
}
