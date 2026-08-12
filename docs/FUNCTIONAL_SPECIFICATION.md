# QWSG Core Alpha Functional Specification

## Current operator state

A successful eligible `check` atomically publishes limited Inventory/Snapshot coverage. A separate bare process validates and displays it, ages it at the exclusive freshness deadline, and fails closed for missing, corrupt, incompatible, unsafe, or unreadable state. Inventory success is not Health.

`qwsg observe` provides the additive full operator workflow. The first invocation persists a baseline and remains unknown. A later invocation executes the existing live Inventory, Snapshot, Compare, Drift, Health, Rule, Policy, and Report stages, validates their typed correlation, and publishes full operator-evaluation coverage. It fails rather than self-comparing or silently replacing corrupt store evidence. This qualifies only the engineering condition represented by current canonical checks.

Valid evaluations with more than 256 candidate attention facts shall still
project and publish. Projection shall deterministically retain the
highest-severity and highest-importance facts, correlate Rule/Policy views only
through validated identities, disclose correlated and omitted counts, and
preserve source references. Operator diagnostics shall distinguish Pipeline,
projection, and Current State publication failures without exposing private
host data. Missing or stale evidence may recommend a fresh observation;
intrinsically partial current evidence shall recommend inspection instead.

## Local Operator Console

Bare `qwsg` presents model-owned condition, attention, changes, Alert summary, Guardian, evidence freshness/completeness, and recommendation. Refresh is explicit and one-shot. The interface performs no remediation, persistence, monitoring, or service control.

## 1. Purpose and authority

This document is the authoritative functional specification for Quantum Wizard Server Guardian (QWSG) Core Alpha. It defines externally observable behavior, required inputs and outputs, workflows, operational states, failure behavior, and acceptance criteria. It is implementation-neutral: it does not select languages, frameworks, process topology, protocols, package layout, database technology, or a secrets backend.

The Project Constitution is supreme. The Product Definition defines parent product intent, and the Product & System Blueprint defines product-level boundaries. If a future design cannot satisfy this specification, the conflict requires an explicit governed specification change; architecture must not reinterpret the requirement silently.

## 2. Status and requirement language

- Status: **authoritative Core Alpha specification**
- Created: `2026-07-20` UTC under Task 006
- Scope: the first coherent Agent, Installer, optional Console, monitoring, notification, reporting, diagnostics, and lifecycle behavior
- Engineering language: English; every user-visible interface is localization-ready

The terms **MUST**, **MUST NOT**, **SHOULD**, and **MAY** express mandatory, forbidden, recommended, and optional behavior. A requirement identified as `FR-*` is mandatory unless explicitly marked as a release gate or post-Alpha item.

## 3. Product boundary

Core Alpha protects one independently operated compatible Linux server. It observes and verifies evidence, interprets that evidence, explains findings, warns operators, retains relevant history, and supports lifecycle administration. It performs no automatic corrective action. The Agent remains useful without a Console and without a vendor-operated cloud service.

Core Alpha does not create backups, repair services, modify monitored application configuration, manage a fleet, provide broad security remediation, or guarantee absolute security or availability. It does not install optional dependencies without preview and explicit consent.

## 4. Actors and authority

| Actor | Functional authority |
| --- | --- |
| Operator | Views status, evidence, incidents, reports, history, and diagnostics; acknowledges incidents; requests on-demand checks and reports. |
| Administrator | Performs Operator actions and manages validated configuration, notification tests, maintenance windows, and supported lifecycle requests. |
| Installer operator | Reviews and explicitly authorizes privileged installation, update, repair, reconfiguration, export, and removal plans. |
| Agent | Observes, evaluates, persists state, creates incidents, sends notifications, reports, audits product behavior, and exposes supported controls. |
| Installer | Detects the environment and performs only the reviewed lifecycle plan within its privilege boundary. |
| Console | Optional localized administration surface; it never grants authority beyond the authenticated role or bypasses Agent/Installer controls. |
| External system | A monitored service, filesystem, endpoint, certificate, backup artifact, mail transport, clock, or operating-system capability. |

`FR-AUTH-001`: Human authority MUST remain final. Observation, recommendation, acknowledgement, or Console access MUST NOT imply remediation authority.

`FR-AUTH-002`: Every state-changing product operation MUST identify the authenticated or local initiating actor, requested operation, target, time, result, and failure reason in audit history.

`FR-AUTH-003`: Read-only checks and reports MAY run automatically according to validated policy. System modification requires a separate explicit lifecycle confirmation.

## 5. Operating profiles

`FR-PROFILE-001`: **Agent-only** MUST provide local configuration, scheduled and on-demand checks, state retention, incidents, e-mail alerts, reports, history, audit evidence, and diagnostics without a Console.

`FR-PROFILE-002`: **Installer-assisted** MUST add transparent installation, update, repair, reconfiguration, and removal workflows to Agent-only or Console-enabled operation.

`FR-PROFILE-003`: **Console-enabled** MUST add secure visualization and supported administration while preserving Agent-only operation during Console outage or removal.

`FR-PROFILE-004`: Installation presets named Minimal, Web Server, Mail Server, Full, and Custom MAY propose initial module selections. A preset MUST display its exact selections and MUST NOT silently enable a capability.

## 6. Core Alpha support and capability model

`FR-CAP-001`: Core Alpha targets Ubuntu and Debian. Exact distribution versions and CPU architectures MUST be held in a versioned, tested support matrix approved before a release is called supported. Absence from that matrix means unsupported, not implicitly compatible.

`FR-CAP-002`: Before mutation, the Installer MUST detect at least distribution and version, CPU architecture, kernel, init system, package manager, virtualization or physical context when detectable, storage mounts, relevant network capability, and required command or service capabilities.

`FR-CAP-003`: Every check target MUST have one capability condition: `available`, `unavailable`, `unsupported`, `disabled`, or `degraded`. The reason and any missing prerequisite MUST be visible.

`FR-CAP-004`: Detection MAY recommend checks from observed server roles but MUST NOT activate them without confirmation.

`FR-CAP-005`: An unavailable optional capability MUST NOT disable unrelated checks. A mandatory unavailable capability MUST stop the affected workflow safely and identify the blocking prerequisite without mutation.

## 7. Configuration model

### 7.1 Configuration content

`FR-CFG-001`: Configuration MUST express a schema version, instance identity, locale and time zone, enabled checks and targets, schedules, timeouts, thresholds, persistence and hysteresis rules, maintenance windows, notification policy, report policy, retention limits, and bounded interface settings.

`FR-CFG-002`: Each value MUST have one effective source and an explainable precedence. Core Alpha precedence is: command-specific temporary override, explicitly activated local override, primary local configuration, then documented built-in default. Temporary overrides MUST NOT persist unless separately confirmed.

`FR-CFG-003`: Configuration changes MUST follow `inspect -> edit or submit -> validate -> preview effective difference -> activate -> verify -> audit`. Invalid configuration MUST NOT replace the last valid active configuration.

`FR-CFG-004`: Activation MUST be atomic from the operator's perspective. On failure, QWSG MUST retain or restore the last valid configuration and report whether any runtime accepted the new version.

`FR-CFG-005`: Configuration MUST support stable identifiers for instances, checks, and targets. Renaming or removing a target MUST not silently merge or erase its history.

`FR-CFG-006`: Duration, size, percentage, and schedule values MUST use unambiguous units. The CLI and Console MUST show the normalized effective value.

`FR-CFG-007`: Unknown configuration keys MUST be errors by default. A forward-compatibility mode MAY preserve unknown keys but MUST NOT activate behavior based on them.

### 7.2 Secrets

`FR-SEC-001`: Secret material MUST be stored through a protected secret reference mechanism separate from ordinary exported configuration.

`FR-SEC-002`: CLI output, Console pages, logs, reports, diagnostics, audit records, exports, and validation errors MUST redact secrets by design. A secret value MUST never be returned after acceptance.

`FR-SEC-003`: Missing, unreadable, or invalid secret references MUST disable only dependent delivery or integration behavior, surface `UNKNOWN` internal health where relevant, and generate a safe diagnostic without exposing the value.

### 7.3 Default monitoring policy

Defaults are starting policy, not claims that every server has identical risk. Operators may override them after validation.

| Subject | WARNING | CRITICAL | EMERGENCY | Recovery boundary |
| --- | ---: | ---: | ---: | ---: |
| Filesystem capacity used | 80% | 90% | 95% | Below 78%, 88%, or 93% for the corresponding transition |
| Filesystem inodes used | 80% | 90% | 95% | Below 78%, 88%, or 93% |
| Available memory | Below 15% | Below 10% | Below 5% | Above 17%, 12%, or 7% |
| Swap used | 50% | 75% | 90% | Below 45%, 70%, or 85% |
| Normalized 15-minute load | Above 0.80 | Above 1.00 | Above 1.50 | Below 0.75, 0.95, or 1.40 |
| TLS certificate validity remaining | Below 30 days | Below 14 days | Below 7 days | Above the entered boundary plus 1 day |

`FR-CFG-008`: Percentage thresholds MUST be ordered so worsening evidence cannot map to a lower severity. Validation MUST reject inverted or overlapping policies.

`FR-CFG-009`: Non-emergency threshold entry requires two consecutive qualifying observations by default. Emergency entry is immediate. Recovery requires two consecutive qualifying observations by default. Operators MAY configure counts of one or more.

`FR-CFG-010`: Backup age and size have no universal safe threshold. Each backup target MUST define a maximum age and MAY define a minimum expected size or minimum change; otherwise it remains disabled with an explanation.

## 8. Observation and check contract

`FR-CHECK-001`: Every observation MUST contain a stable check ID, target ID, observation time, completion time, capability status, outcome state, summarized evidence, policy version or basis, freshness deadline, and error identifier when applicable.

`FR-CHECK-002`: Evidence MUST distinguish measured facts from interpretation. Sensitive source content MUST be minimized and redacted.

`FR-CHECK-003`: Every check MUST have a finite timeout. Timeout produces `UNKNOWN` for that target and MUST NOT block unrelated checks.

`FR-CHECK-004`: Schedules MUST prevent uncontrolled overlapping execution of the same check and target. If an earlier run remains active, policy MUST skip or queue one replacement run and record the decision; it MUST NOT create an unbounded backlog.

`FR-CHECK-005`: On-demand execution MUST use the same validation, timeout, normalization, persistence, and state-transition rules as scheduled execution. It MUST identify itself as operator-requested.

`FR-CHECK-006`: Malformed module output MUST be rejected, recorded as an internal error, and mapped to `UNKNOWN` only for the affected subject.

`FR-CHECK-007`: QWSG MUST use a monotonic elapsed-time source for timeouts and scheduling intervals where available, and wall-clock time for human timestamps. A significant wall-clock discontinuity MUST be reported and MUST NOT fabricate a recovery or duplicate incident.

## 9. Required Core Alpha checks

### 9.1 Disk capacity and inodes

`FR-DISK-001`: QWSG MUST evaluate configured mounted filesystems independently for byte capacity and inode consumption, excluding ephemeral or unsupported filesystem types only through visible policy.

`FR-DISK-002`: Evidence MUST include mount identity, filesystem type, total, used, available, percent used, threshold basis, and capability or read error. Byte and inode incidents remain distinct.

`FR-DISK-003`: A disappeared configured mount MUST become `UNKNOWN` or `CRITICAL` according to explicit target policy; it MUST NOT be reported as healthy or silently removed.

### 9.2 Memory and swap

`FR-MEM-001`: Memory evaluation MUST use Linux available-memory semantics and MUST NOT classify reclaimable cache as unavailable memory.

`FR-MEM-002`: Evidence MUST include total and available memory, calculated available percentage, swap availability and use, and policy basis.

`FR-MEM-003`: A host without configured swap MUST report swap as `unsupported` or `disabled` according to policy, not as failed usage.

### 9.3 CPU load

`FR-LOAD-001`: Load MUST be normalized by the CPU capacity visible to the monitored environment. Evidence MUST include raw 1-, 5-, and 15-minute load, visible capacity, normalized values, and the evaluated interval.

`FR-LOAD-002`: Core Alpha default severity uses normalized 15-minute load. Missing or invalid capacity evidence produces `UNKNOWN` rather than an unnormalized health claim.

### 9.4 systemd services

`FR-SVC-001`: QWSG MUST monitor explicitly selected systemd units and distinguish active, inactive, failed, activating, deactivating, not found, inaccessible, and unsupported states.

`FR-SVC-002`: The default mapping is `OK` for active, `WARNING` for prolonged activating or deactivating, `CRITICAL` for inactive or failed, and `UNKNOWN` for inaccessible or unsupported evidence. Per-target policy MAY allow a unit to be intentionally inactive.

`FR-SVC-003`: Process existence alone MUST NOT override a failed service-manager state.

### 9.5 HTTP and HTTPS outcomes

`FR-HTTP-001`: Each endpoint MUST define URL, accepted status outcome, timeout, redirect policy, and optional TLS validation policy. Request bodies and response bodies MUST NOT be retained by default.

`FR-HTTP-002`: Evidence MUST distinguish name-resolution, connection, timeout, TLS, redirect, protocol, and unacceptable-status failures and record latency without exposing credentials.

`FR-HTTP-003`: Redirect following MUST have a finite maximum and MUST not forward sensitive authentication across an unapproved origin.

### 9.6 TLS certificate expiry

`FR-TLS-001`: For configured TLS endpoints, QWSG MUST evaluate certificate validity dates, hostname validity, and chain-validation outcome available to the client environment.

`FR-TLS-002`: An expired, not-yet-valid, hostname-invalid, or untrusted certificate is at least `CRITICAL`; inability to complete observation is `UNKNOWN`. Expiry thresholds apply only to an otherwise valid observation.

### 9.7 Existing backups

`FR-BACKUP-001`: QWSG MUST monitor existing configured backup artifacts or sets; it MUST NOT create or modify backups.

`FR-BACKUP-002`: Evidence MUST include target identity, newest qualifying artifact time, age, size where available, matching basis, and observation error. Filenames or paths containing sensitive customer data SHOULD be minimized in user output.

`FR-BACKUP-003`: No qualifying artifact, an artifact older than the configured maximum, or an artifact below a configured minimum size maps to configured unhealthy severity. Inaccessible evidence maps to `UNKNOWN`.

## 10. State model

### 10.1 Subject state

The severity order is `OK < WARNING < CRITICAL < EMERGENCY`. `UNKNOWN` is not part of that health order; it represents missing, failed, invalid, or stale evidence and MUST never be treated as `OK`.

`FR-STATE-001`: Each enabled subject MUST retain current state, previous state, state-entry time, last observation time, last successful observation time, policy basis, evidence summary, and active incident reference if any.

`FR-STATE-002`: The first valid observation establishes baseline state. If unhealthy, it opens an incident and alerts after configured persistence. A first `OK` observation MUST NOT generate a recovery notification.

`FR-STATE-003`: Worsening severity is escalation; improving but still unhealthy severity is de-escalation; transition to `OK` is full recovery; transition to `UNKNOWN` is observation loss.

`FR-STATE-004`: Hysteresis and consecutive-observation policy MUST be evaluated before a transition is committed. Pending candidates MUST be visible diagnostically but MUST NOT overwrite current state.

`FR-STATE-005`: An observation is stale after its freshness deadline. Stale state becomes `UNKNOWN` with the last known health retained as context, not as current truth.

### 10.2 Incidents

`FR-INC-001`: A confirmed transition from `OK` or no baseline into `WARNING`, `CRITICAL`, `EMERGENCY`, or `UNKNOWN` MUST open an incident with a stable identifier.

`FR-INC-002`: Escalation, de-escalation, evidence update, acknowledgement, maintenance suppression, reminder, delivery attempt, and recovery MUST append to the same incident timeline until full recovery or explicit administrative closure permitted by policy.

`FR-INC-003`: Acknowledgement records actor, time, and optional localized note. It MUST NOT change measured health or suppress escalation and recovery notices. Policy MAY suppress unchanged-state reminders while acknowledged.

`FR-INC-004`: Full recovery closes the active incident automatically. A later recurrence opens a new incident rather than reopening the old one.

### 10.3 Aggregate health

`FR-STATE-006`: Overall health MUST be the highest confirmed unhealthy severity among enabled supported subjects, except that any `UNKNOWN` or stale subject MUST be separately visible and MUST prevent an unqualified overall `OK` claim.

`FR-STATE-007`: Disabled, unsupported, and maintenance-suppressed subjects MUST be counted and displayed separately from healthy subjects.

## 11. Maintenance and suppression

`FR-MAINT-001`: A maintenance window MUST identify scope, start, end, actor, reason, and creation time. Open-ended maintenance is forbidden in Core Alpha.

`FR-MAINT-002`: Checks and state evaluation continue during maintenance by default. Matching outbound incident notifications are suppressed, but transitions and suppression reasons remain recorded.

`FR-MAINT-003`: On maintenance end, QWSG MUST immediately evaluate current retained freshness. If an unhealthy incident remains active, it MUST send one maintenance-end status notification; it MUST not replay every suppressed event.

`FR-MAINT-004`: Emergency notification suppression MUST require an explicit policy choice visible in the maintenance preview.

## 12. Alert lifecycle and delivery

Implementation boundary: Task 028 implements the deterministic Alert decision
and lifecycle contract. Task 029 consumes its immutable Alert Records only and
implements provider-neutral delivery planning, bounded queue/retry proposals,
idempotent requests, delivery attempt/status/acknowledgement/evidence records,
and an explicitly invoked one-cycle injected-provider adapter. Neither task
implements incident persistence, monitoring, a durable queue, a daemon,
concrete production transport, channel-health Alert generation,
configuration/secret integration, or CLI/Console workflows.

`FR-ALERT-001`: Notification events are created for confirmed unhealthy entry, escalation, de-escalation, full recovery, observation loss, configured emergency reminder, maintenance-end active status, and notification-channel failure or recovery.

`FR-ALERT-002`: Unchanged polling MUST NOT repeatedly notify. The default reminder policy is one reminder every 24 hours for an unresolved `EMERGENCY`; reminders for other severities are disabled by default.

`FR-ALERT-003`: Every notification MUST identify instance, subject, old and new state, incident ID, evidence summary, event and observation times, policy basis, recommended next step, acknowledgement status, maintenance context, and whether remediation was attempted. Core Alpha always states that remediation was not attempted.

`FR-ALERT-004`: E-mail is a post-1.0 delivery capability. Version 1.0 ships
provider-neutral Notification contracts and locally visible Alert evidence;
it MUST NOT claim or require a concrete transport. A later transport release
must select, specify, and verify its supported local-submission or authenticated
SMTP boundary.

`FR-ALERT-005`: Delivery attempts MUST have finite timeouts and bounded retries with backoff. Default policy is three total attempts within 15 minutes. Failure after retries becomes a visible notification-health incident.

`FR-ALERT-006`: A delivery failure MUST NOT recursively use the same failed channel without limit. It MUST be visible in CLI, Console if present, daily reports, diagnostics, and audit/operational history.

`FR-ALERT-007`: Delivery recovery MUST be recorded and MAY generate one recovery notice through the recovered channel. A test notification MUST be clearly labeled and MUST not create a monitoring incident.

## 13. Reporting

`FR-REPORT-001`: Core Alpha MUST provide on-demand reports and a configurable daily report. Alert-only mode disables scheduled reports but not incident notifications.

`FR-REPORT-002`: A report MUST include generation time and coverage window, instance identity, aggregate health qualification, counts and details for unhealthy, unknown, stale, disabled, unsupported, and maintenance subjects, active incidents, recoveries, check freshness, backup and certificate status, notification health, scheduler health, storage health, and relevant recent errors.

`FR-REPORT-003`: Reports MUST distinguish “no unhealthy evidence” from “all enabled checks verified healthy.” Missing or stale evidence MUST be prominent.

`FR-REPORT-004`: Re-running a report for the same window MUST not mutate monitoring state. Report generation failure MUST not disable checks or alerts.

## 14. CLI behavior

The normative command name is `qwsg`. Packaging MAY additionally expose a privileged `qwsg-installer` entry point, but equivalent lifecycle actions invoked through `qwsg lifecycle` MUST preserve the same confirmation boundary.

`FR-CLI-001`: `qwsg --help` and every subcommand help MUST work without active configuration or elevated privilege and return success.

`FR-CLI-002`: Core Alpha MUST provide these logical commands:

| Command | Required behavior |
| --- | --- |
| `qwsg status` | Show qualified overall state, active incidents, unknown/stale counts, last successful cycle, and product-health summary. |
| `qwsg check [TARGET]` | Run all enabled checks or one authorized target on demand and show normalized results. |
| `qwsg incidents` | List and inspect incident timelines; support explicit acknowledgement. |
| `qwsg report` | Generate an on-demand report without changing health state. |
| `qwsg config validate` | Validate syntax, schema, references, ranges, and capability assumptions without activation. |
| `qwsg config show` | Show effective redacted configuration and value sources. |
| `qwsg config apply` | Validate and preview; activate only after confirmation or an explicit non-interactive confirmation flag. |
| `qwsg maintenance` | List, create, and cancel bounded maintenance windows with preview. |
| `qwsg notify test` | Send a clearly labeled test message and return delivery result. |
| `qwsg diagnostics` | Show product health or create a previewed, redacted support bundle. |
| `qwsg version` | Show product, configuration schema, state schema, and platform-support information. |
| `qwsg lifecycle` | Delegate supported install, update, repair, reconfigure, export, and remove workflows to the privileged lifecycle boundary. |

`FR-CLI-003`: Read-only commands MUST NOT require elevated privilege when their data can be safely exposed to the invoking role. A permission denial MUST identify the required role without suggesting unsafe broad privilege.

`FR-CLI-004`: Human output MUST be localizable. `--output json` MUST provide a versioned, locale-independent machine representation for read-only status, check, incident, report, config-validation, diagnostic-status, and version operations.

`FR-CLI-005`: Success returns exit code `0`; confirmed unhealthy monitoring result returns `1`; invalid input or configuration returns `2`; permission or authorization denial returns `3`; unavailable dependency or unsupported capability returns `4`; internal or persistence failure returns `5`. Commands MUST document exceptions and MUST not use a healthy exit code for an incomplete requested operation.

`FR-CLI-006`: Non-interactive state-changing commands MUST require an explicit confirmation flag, show or emit the exact plan before execution, and refuse if the plan changed after confirmation.

## 15. Installer and lifecycle behavior

### 15.1 Common lifecycle workflow

`FR-LIFE-001`: Install, update, repair, reconfigure, and remove MUST follow `inspect -> validate -> plan -> preview -> explicit consent -> execute -> verify -> report`. Inspection, validation, and preview MUST not mutate the server.

`FR-LIFE-002`: A plan MUST identify exact QWSG components, external dependencies, privileges, files or services affected by category, configuration/state migration, estimated service interruption, rollback capability, irreversible limitations, and retained operator data.

`FR-LIFE-003`: Consent MUST bind to the displayed plan. Environment drift that changes the plan MUST stop execution and require a new preview and consent.

`FR-LIFE-004`: Partial failure MUST identify completed, failed, and unattempted steps; resulting operability; safe retry conditions; and exact rollback or recovery guidance. Re-running after interruption MUST be safe or explicitly refused with guidance.

### 15.2 Installation

`FR-LIFE-005`: Installation MUST reject unsupported platforms without mutation. For partially supported platforms, it MUST show unavailable capabilities before consent.

`FR-LIFE-006`: Optional dependencies MUST be individually attributable to selected capabilities and require consent. Declining one MUST disable only dependent capability where a coherent installation remains possible.

`FR-LIFE-007`: Installation verification MUST confirm installed version, active valid configuration, required state access, enabled check execution, scheduler health, and configured notification transport test result. A failed verification means installation is incomplete.

### 15.3 Update and migration

`FR-LIFE-008`: Update MUST verify source integrity according to the later security architecture, compatibility, available space, configuration and state schema paths, and rollback or recovery prerequisites before mutation.

`FR-LIFE-009`: A migration MUST preserve configuration, current state, incident continuity, audit records, and retained history or stop before activation. Destructive or lossy migration requires explicit disclosure and separate confirmation.

`FR-LIFE-010`: Downgrade is unsupported unless the target release declares and verifies a compatible path; otherwise QWSG MUST refuse with export and recovery guidance.

### 15.4 Repair and reconfiguration

`FR-LIFE-011`: Repair MUST inspect expected QWSG-owned artifacts and propose only bounded restoration. It MUST NOT overwrite operator-managed external configuration silently.

`FR-LIFE-012`: Reconfiguration MUST preserve the last valid active configuration until the replacement validates and activates successfully.

### 15.5 Removal

`FR-LIFE-013`: Removal preview MUST distinguish QWSG executables, services, configuration, secrets, state, history, reports, logs, audit records, exports, and shared external dependencies.

`FR-LIFE-014`: Default removal MUST stop and remove QWSG runtime components while preserving operator data and an export path. Deleting retained data requires a separate explicit confirmation naming the data categories.

`FR-LIFE-015`: QWSG MUST NOT automatically delete shared dependencies or external service configuration. Completion MUST report preserved artifacts and manual actions.

## 16. Console behavior

`FR-CONSOLE-001`: The Console is optional. It MUST display current qualified health, evidence, active incidents, history, reports, capability state, redacted effective configuration, notification health, lifecycle status, and diagnostics.

`FR-CONSOLE-002`: Console authentication, session expiry, brute-force resistance, authorization roles, and recovery behavior MUST conform to a separately approved Security Architecture before network-exposed release. Until then, Console exposure beyond an explicitly protected local or administrative boundary is prohibited.

`FR-CONSOLE-003`: Console changes MUST use the same validation, preview, authorization, activation, verification, and audit behavior as CLI changes. It MUST NOT expose an arbitrary shell or arbitrary privileged file/command execution.

`FR-CONSOLE-004`: Loss, corruption, or unavailability of the Console MUST not stop Agent scheduling, checks, alerts, reports, or local CLI diagnostics.

`FR-CONSOLE-005`: Every page, label, message, date, number, duration, and accessibility name MUST support localization and locale-specific presentation without changing stored machine states or identifiers.

## 17. Persistence, history, logs, and audit

`FR-DATA-001`: Current state and active incident continuity MUST survive routine Agent restart. Restart MUST not generate false recovery, duplicate entry alerts, or reset reminder timing.

`FR-DATA-002`: State, configuration, and history formats MUST carry explicit schema versions. Unsupported schema versions MUST stop affected activation with diagnostics; they MUST not be overwritten.

`FR-DATA-003`: Core Alpha MUST retain observations sufficient for current state, incident timelines, delivery outcomes, daily reports, and diagnostics under a bounded policy. Exact default retention and storage budgets are release-gate decisions and MUST be documented before production use.

`FR-DATA-004`: Storage pressure or write failure MUST surface product health as degraded or `UNKNOWN`, continue safe in-memory observation where possible, avoid claiming durable recording, and prioritize evidence needed to explain the failure.

`FR-DATA-005`: Operational logs explain runtime behavior; history records monitoring events; audit records authorization and product-controlled changes. These categories MUST remain distinguishable even if one storage technology serves them.

`FR-DATA-006`: Retention cleanup MUST be bounded to QWSG-owned data, auditable, and unable to delete active incident state or the only record required to explain an ongoing failure.

## 18. Diagnostics and support bundles

`FR-DIAG-001`: Diagnostics MUST expose product version, platform-support result, configuration validity, enabled and unavailable capabilities, last and next check timing, scheduler health, notification health, persistence health, clock concerns, Console health if installed, and recent categorized errors.

`FR-DIAG-002`: Support-bundle generation MUST first preview included categories and host-identifying categories. It MUST exclude secrets and authentication material by design and redact ordinary configuration.

`FR-DIAG-003`: Bundle generation MUST be read-only with respect to monitored systems and MUST not upload or transmit the bundle. Any future transmission is a separate explicit operation and policy.

`FR-DIAG-004`: A failure to collect one diagnostic category MUST not discard successfully collected categories; the bundle manifest MUST disclose omissions and errors.

## 19. Internal product health and failure isolation

`FR-REL-001`: QWSG MUST monitor scheduler progress, check freshness, persistence writes, notification delivery, configuration validity, and clock discontinuity as product-health subjects.

`FR-REL-002`: One failed check, malformed result, unavailable optional capability, notification outage, report failure, or Console failure MUST not silently disable unrelated protection.

`FR-REL-003`: After restart, interrupted executions MUST be recorded as incomplete and reevaluated. They MUST not be treated as successful observations.

`FR-REL-004`: Retry policies MUST be finite and observable. Concurrency MUST be bounded by validated configuration. When capacity is exhausted, QWSG MUST preserve critical checks preferentially according to explicit priority and report delayed or skipped work.

`FR-REL-005`: Corrupt retained data MUST not be silently reset. QWSG MUST isolate the affected data where possible, preserve evidence, enter degraded operation, and provide repair or restore guidance.

## 20. Security, privacy, and localization requirements

`FR-NFR-001`: Routine monitoring MUST use least privilege sufficient for enabled checks. Privileged lifecycle actions and future remediation remain separate authority boundaries.

`FR-NFR-002`: Inputs from configuration, modules, monitored services, files, network endpoints, and user interfaces MUST be treated as untrusted and validated before interpretation or display.

`FR-NFR-003`: Logs and interfaces MUST prevent untrusted evidence from forging structure, terminal control, markup, or audit attribution.

`FR-NFR-004`: Core Alpha MUST perform no telemetry or vendor-controlled external transmission by default. E-mail destinations and monitored endpoints configured by the operator are explicit operational transmissions.

`FR-NFR-005`: User-visible strings MUST be externalizable. Machine identifiers, state tokens, JSON keys, and audit semantics remain stable and locale-independent.

`FR-NFR-006`: Time shown to users MUST include an unambiguous offset or named configured zone. Stored event ordering MUST remain consistent across locale changes and daylight-saving transitions.

`FR-NFR-007`: Accessibility semantics, keyboard operation, text alternatives, focus visibility, and non-color state indicators are mandatory for Console acceptance.

## 21. Explicit Core Alpha exclusions

The following are not Core Alpha requirements: automatic remediation; backup creation; broad security repair; databases, mail queues, DNS, hardware, SMART, RAID, UPS, log aggregation, port-change, and firewall checks; non-email notification channels; fleet management; hosted services; public APIs; weekly or comparative trend reports; entertainment or shell-convenience modules; edition enforcement; billing; and public licensing behavior.

A later feature MUST preserve this specification's authority, safety, state, audit, localization, and failure-isolation principles unless an explicitly approved change supersedes them.

## 22. Release gates and unresolved owner decisions

Task 038 resolves these historical Core Alpha gates for the narrow Version 1.0
local product as follows. Future families remain explicit rather than silently
inferred:

1. Supported: Ubuntu 24.04 LTS, systemd 255+, linux-amd64 only.
2. The local non-network Terminal Console ships.
3. Concrete notification transports are post-1.0.
4. Existing bounded file retention and engine caps ship; general history does not.
5. The Console is local and read-only, with no listener or authentication surface.
6. Configuration Source 1.0, private bounded files, and a systemd user service ship.
7. Deterministic archives and SHA-256 integrity ship; publisher authentication/signing is a disclosed later/publication decision.
8. Licensing and public/commercial publication remain separate Owner decisions; no telemetry or hosted dependency ships.

## 23. End-to-end workflows

### 23.1 Normal monitoring cycle

1. Scheduler selects due enabled checks under bounded concurrency.
2. Capability and configuration preconditions are confirmed.
3. Each check returns normalized evidence or a bounded failure.
4. State policy evaluates persistence, hysteresis, freshness, and transition.
5. Current state and history are durably recorded where storage is healthy.
6. Incident events and notification events are created only for meaningful transitions or configured reminders.
7. Delivery outcomes are recorded; unrelated checks continue.
8. Status, reports, CLI, and optional Console expose the same retained truth.

### 23.2 Configuration change

1. Administrator submits a candidate.
2. QWSG validates schema, values, references, capabilities, and secret references.
3. QWSG shows effective differences, affected checks, and restart or interruption implications.
4. Administrator confirms the unchanged plan.
5. QWSG atomically activates and verifies the candidate.
6. Success or rollback to the previous valid configuration is audited and reported.

### 23.3 Incident lifecycle

1. Qualifying observations confirm unhealthy or unknown entry.
2. QWSG opens one incident and sends one entry alert unless maintenance policy suppresses it.
3. Evidence updates without state change do not alert repeatedly.
4. Escalation alerts immediately; emergency reminders follow bounded policy.
5. Acknowledgement records operator awareness without changing health.
6. De-escalation communicates remaining risk.
7. Confirmed `OK` closes the incident and sends one recovery alert.

### 23.4 Lifecycle change

1. Installer inspects without mutation.
2. It validates compatibility and prerequisites.
3. It creates a complete bounded plan and rollback/recovery statement.
4. The installer operator explicitly consents to that exact plan.
5. Execution records every step and stops safely on material drift or failure.
6. Verification determines complete, incomplete, or recovered result.
7. The final report identifies changes, retained data, limitations, and next action.

## 24. Acceptance criteria

### 24.1 Agent and checks

- `AC-AGENT-001`: On a declared supported test platform, Agent-only installation completes, survives restart, schedules checks, retains state, sends e-mail, generates a daily report, and exposes diagnostics without Console availability.
- `AC-AGENT-002`: Deterministic fixtures drive every required check through `OK`, each applicable unhealthy severity, `UNKNOWN`, timeout, malformed evidence, and recovery with the required evidence fields.
- `AC-AGENT-003`: Disk/inode, memory/swap, normalized load, systemd, HTTP/HTTPS, TLS expiry, and existing-backup behaviors match Sections 7 and 9.
- `AC-AGENT-004`: A failed check or Console outage does not interrupt unrelated scheduled checks or alerts.

### 24.2 State and alerts

- `AC-STATE-001`: Tests prove baseline, consecutive-observation persistence, emergency immediate entry, hysteresis, escalation, de-escalation, recovery, stale-to-`UNKNOWN`, recurrence, and restart continuity.
- `AC-STATE-002`: Unchanged observations produce no repeated alert; one default 24-hour emergency reminder is produced per interval; restart does not duplicate it.
- `AC-STATE-003`: Acknowledgement and maintenance do not falsify health. Maintenance end produces at most one current-status notification per active incident.
- `AC-STATE-004`: Notification failure exhausts bounded retries, produces visible delivery-health failure without recursion, and records later recovery.

### 24.3 Configuration, CLI, and Console

- `AC-UX-001`: Invalid or inverted configuration never replaces the active version; effective-source and redaction tests pass.
- `AC-UX-002`: Every required CLI command, JSON output, exit-code class, permission denial, preview, and non-interactive confirmation behavior passes contract tests.
- `AC-UX-003`: If Console is included, role authorization, session controls, brute-force controls, supported-change parity, localization, accessibility, and Agent independence pass before network exposure.
- `AC-UX-004`: At least English and one non-English test catalog render without hardcoded user-visible strings changing machine identifiers or stored state.

### 24.4 Lifecycle and data safety

- `AC-LIFE-001`: Install, update, repair, reconfigure, and removal fixtures prove no mutation before consent, plan binding, drift refusal, partial-failure reporting, safe retry or recovery, and bounded audit evidence.
- `AC-LIFE-002`: Removal preserves operator data by default and never removes a shared dependency or external configuration automatically.
- `AC-DATA-001`: Restart, storage-full, write-failure, corrupt-state, incompatible-schema, and wall-clock-discontinuity tests produce explicit degraded behavior without false health or silent reset.
- `AC-DATA-002`: Support-bundle tests find no seeded secrets, preview identifying categories, disclose collection failures, and perform no upload.

### 24.5 Documentation and release readiness

- `AC-DOC-001`: Every implemented Core Alpha behavior traces to an `FR-*` requirement and every requirement traces to one or more tests, a documented release gate, or an explicit deferred exclusion.
- `AC-DOC-002`: Operator documentation covers prerequisites, support matrix, permissions, configuration, secrets, checks, states, incidents, alerts, reports, lifecycle, troubleshooting, privacy, retention, and removal.
- `AC-REL-001`: No release is described as supported until all eight Section 22 release gates are resolved, applicable acceptance tests pass, and limitations are published.

## 25. Traceability to parent documents

| Functional area | Product Definition | Product & System Blueprint | Specification |
| --- | --- | --- | --- |
| Protection, transparency, operator authority | Product identity; Product philosophy; Core values | Sections 5, 7, 12–14 | Sections 3–4, 20 |
| Agent, Installer, Console | Agent and Console relationship | Sections 16–20, 28, 30 | Sections 5, 14–16 |
| Detection and required checks | Problems and goals | Sections 15, 21, 29, 31 | Sections 6–9 |
| State, incidents, alerts, recovery | Guardian responsibility | Sections 22–24, 31, 35 | Sections 10–13 |
| Configuration, secrets, data, audit | Security and privacy principles | Sections 25–27, 34, 36 | Sections 7, 17–20 |
| Lifecycle and reversibility | Product philosophy | Sections 18, 28, 31 | Section 15 |
| Offline independence | Environmental boundary; Offline philosophy | Sections 6, 9, 17, 33–34 | Sections 3, 5, 20 |
| Localization | Product philosophy; Guiding engineering principles | Sections 12, 19, 37 | Sections 14, 16, 20 |
| MVP and deferred capabilities | Goals and non-goals | Sections 31–32, 38–40 | Sections 21–22, 24 |

## 26. Glossary

- **Capability:** functionality available only when verified environment and dependency conditions are satisfied.
- **Check:** one bounded evidence-gathering evaluation of a target.
- **Observation:** timestamped normalized evidence from a check.
- **Subject:** the stable check-and-target entity whose state is retained.
- **State:** current retained interpretation: `OK`, `WARNING`, `CRITICAL`, `EMERGENCY`, or `UNKNOWN`.
- **Transition:** confirmed change of retained state after policy evaluation.
- **Incident:** timeline from confirmed unhealthy or unknown entry through full recovery.
- **Acknowledgement:** operator record of awareness; it does not alter health.
- **Maintenance:** bounded policy that suppresses matching notifications while retaining observations and transitions.
- **Stale:** an observation whose freshness deadline has passed.
- **Recovery:** transition to `OK`; improvement that remains unhealthy is de-escalation, not full recovery.
- **Remediation:** product-initiated corrective action. It is excluded from Core Alpha.

## 27. Change control

Changes to mandatory behavior require a governed engineering task, explicit authority, snapshot and rollback, parent-document consistency review, requirement and acceptance-test updates, and chronological history. Implementation convenience, framework limitation, or undocumented existing behavior is not authority to weaken this specification.

## 28. Canonical Runtime cycle

The core shall accept an explicit bounded execution context, invoke one
Scheduler cycle, validate its execution traces, evaluate Alert inputs in
canonical order, plan Notification delivery from Alert Records only, and run at
most one provider cycle when requests exist. It shall preserve completed
evidence on cancellation or failure, return a proposed idle Runtime state, and
shall not run continuously or duplicate an owning engine's decisions.

## 29. Canonical Runtime Service

The local Runtime Service shall invoke the Canonical Runtime Engine sequentially
at fixed-rate nominal boundaries, start with one immediate cycle, skip elapsed
boundaries without overlap or catch-up bursts, propagate cancellation and
per-cycle deadlines, and stop gracefully on caller cancellation, SIGINT, or
SIGTERM. It shall forward Runtime-proposed states exactly in memory and shall
not reproduce Runtime or downstream decisions.

The Operational Guardian adapter shall run this service as one unprivileged
foreground process under the systemd user manager. It shall consume canonical
configuration, retain only a bounded validated Runtime/Alert/Notification
restart checkpoint, publish exact lifecycle evidence to Current Operator
State, reject competing writers, stop gracefully, and fail closed on invalid
state or configuration. A stale former-running claim shall become unavailable.
Neither systemd nor the adapter may reproduce an engineering decision.

## 30. Canonical operator overview

QWSG shall derive beginner-facing condition, attention, change, Alert,
Guardian, freshness/completeness, and recommended-next-step facts through one
presentation-independent model over validated canonical records. Missing,
stale, partial, unsupported, invalid, and not-observed evidence must remain
distinct and must never be displayed as healthy. Recommendations are bounded
read-only action tokens and grant no remediation authority. This contract does
not itself add bare-`qwsg` behavior, a Console, API, Dashboard, monitoring,
persistence, or process discovery.
