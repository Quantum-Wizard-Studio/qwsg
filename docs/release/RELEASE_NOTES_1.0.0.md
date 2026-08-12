# QWSG 1.0.0 Release Notes

QWSG 1.0.0 is the accepted local Linux Server Guardian for Ubuntu 24.04 LTS,
linux-amd64 and systemd 255 or later. It promotes the behavior proven through
RC.1–RC.3 and the real Task 043 clean-host acceptance without adding a new
product feature or changing canonical engine decisions.

The Community Edition includes the complete accepted local QWSG 1.0 product:

- local Inventory and canonical Snapshot/Digital Twin evidence;
- Comparison, Drift, Health, Rule, Policy and Report evaluation;
- bounded Alert and provider-neutral Notification contracts;
- the recurring non-root local Guardian under a systemd user service;
- private local state/history, diagnostics and read-only Operator Console;
- safe installation, replacement, rollback and uninstall tooling.

The Task 039 bounded aggregate Alert reference supports large valid Policy
Reports while preserving full traceability in the canonical Report. Console
refresh remains read-only, Runtime diagnostics remain privacy-safe and
Guardian liveness remains freshness-bounded. The Task 041 secure first-use
bootstrap permits a clean account to establish a truthful partial baseline
without Go, Git, a repository checkout or manual state-directory preparation.

Task 043 accepted the exact RC.3 behavior on a freshly installed disposable
Ubuntu 24.04 host, including first use, later full evaluation, Guardian start,
boot-before-login lingering, physical reboot recovery, recurring post-reboot
cycles, controlled restart and uninstall with private state preserved.

QWSG 1.0 is distributed under the QWS Community / Free License Version 1.0, a
proprietary source-available Community license rather than an OSI open-source
license. No paid license or Quantum Wizard API key is required for the complete
local Community functionality shipped in this release. Future central APIs,
managed notification/alerting, remote history, Dashboard, fleet, team/role,
compliance, remote-management and managed backup services are not implemented
or promised by this release. Their future absence or failure must not disable
the local Community Guardian or corrupt local evidence.

The release archive remains a deterministic linux-amd64 package containing the
static binary, systemd user unit, installer, uninstaller, configuration example,
license, changelog, internal manifest and operator documentation. SHA-256
provides integrity but does not authenticate an untrusted distribution source.
Tagging, push and external publication require separate Project Owner authority.
