# QWSG Community 1.0 and Pro 1.0 Scope Freeze & Release Gates

## Authority and change control

Status: **Owner-approved**, recorded by Task 088 on 2026-09-20 UTC following
the post-Task-087 planning audit. This is the canonical product-completion
boundary for Community 1.0 and Pro 1.0, subordinate to the Project Constitution.

For these completion decisions, this scope supersedes older, broader release
requirements in the Product Definition, Product & System Blueprint, Functional
Specification, Product Architecture and roadmap. Historical requirements and
engineering records remain valuable and are not deleted or retroactively
rewritten. Deferred requirements retain their design meaning for later work;
their implementation is not a 1.0 release gate. Existing security, privacy,
truthfulness and operator-authority guarantees are not weakened.

Progress MUST primarily be reported by the gates below, not speculative future
task numbers. New ideas default to POST-1.0 or FUTURE. A new capability MUST NOT
become a 1.0 release blocker without an explicit future Owner scope decision,
recorded through the governed lifecycle with its effect on this document.
An engine name, roadmap entry or useful enhancement alone creates no blocker.
Corrections needed to satisfy an existing gate remain within that gate.

Gate status changes require bounded implementation/acceptance evidence and a
reference to the corresponding task or acceptance record. PASS means the
defined gate is satisfied, not universal product correctness. PARTIAL means
existing capability has remaining bounded work; MISSING means the required
production path or acceptance is absent. Community completion requires all
C gates PASS; Pro completion requires all C and P gates PASS. Owner support
and release decisions remain separate from task completion.

## Product maturity and release versions

“Community 1.0” and “Pro 1.0” name product maturity/completion boundaries here,
not instructions to renumber software. Historical QWSG 1.2.0, 1.3.0 and 1.3.1
versions, artifacts and records remain immutable. The post-1.3.1 source needs
a distinct release identity under the release process. Task 088 neither selects
that version nor builds, signs, publishes or deploys a release.

## COMMUNITY 1.0 MUST HAVE

Community 1.0 is a trustworthy, privacy-preserving, operator-controlled local
Linux Server Guardian that observes the server, builds deterministic local
evidence, detects meaningful change/state, reports locally, can notify an
administrator, understands authenticated releases, and provides safe explicit
operator-controlled update and rollback behavior.

The supported boundary remains [Ubuntu 24.04 LTS, amd64, systemd 255+ and
documented local filesystem/user-manager conditions](release/SUPPORT.md).
It includes the accepted inventory, snapshots, comparison, Drift, Health,
Rule, Policy, local Report/Console, Scheduler and Guardian core. It does not
promise checks outside implemented coverage. In particular, “healthy” MUST NOT
be described as proof of universal server health: current Health evaluates
canonical engineering evidence, not every resource threshold, endpoint,
certificate, backup or application concern.

### Community release gates

| Gate | Current state | Required completion boundary / remaining work |
| --- | --- | --- |
| C1 — Supported installation, non-root operation and privacy boundary | PASS | Preserve the supported local Linux boundary, guided installation/setup/removal, ordinary-user Guardian, private local state, privacy/redaction and explicit operator authority. |
| C2 — Reliable inventory and meaningful change identity | PASS | Task 089 establishes deterministic protected systemd unit identity, canonical ordering and exact running-set appearance/disappearance/replacement evidence. Ordinal historical snapshots remain readable; exact comparison involving their nonempty service sets is explicitly refused. See the Task 089 evidence below. |
| C3 — Local evidence interruption integrity and recovery | PASS | Task 090 proves kernel-owned writer exclusion, explicit legacy-lock recovery, deterministic pre/post-commit retention recovery, preservation on durability errors and bounded evidence/checkpoint input. Unknown or conflicting states preserve evidence and fail closed. See Task 090 and the supported filesystem contract below. |
| C4 — Continuous Guardian and bounded resource behavior | PASS | Preserve accepted Guardian, resource, Scheduler and restart behavior. Existing accepted evidence remains authoritative; do not recreate completed work merely for task-number progression. |
| C5 — Honest Health / Rule / Policy and local reporting | PASS | Preserve deterministic evaluation and local reporting limited to actual evidence/check coverage. No universal-health claim is permitted. |
| C6 — Administrator SMTP notification and authenticated release awareness | PASS | Preserve security, TLS, privacy, bounded retry/deduplication and authenticated release boundaries. SMTP is operator-configurable and need not be configured on every Community installation. |
| C7 — Safe explicit update / rollback / recovery | PARTIAL | Bring manual update failure recovery to the established fail-closed standard. Do not discard rollback results. Required Guardian restart depends on verified recovery/package state. Rollback/backup/journal evidence remains sufficiently durable and actionable after failure/interruption. No entire-updater redesign is required. |
| C8 — Canonical scope, readiness and recovery documentation | PARTIAL | Authoritative documentation must reflect this freeze and disposition obsolete broader MUSTs. Supported coverage, optional SMTP, update bootstrap and recovery boundaries must be accurate and usable. See the Task 088 documentation disposition below. |
| C9 — Distinct signed release and real upgrade acceptance | MISSING | Give post-1.3.1 source a distinct identity, produce a reproducible artifact and authenticated release metadata/index, prove supported upgrade/bootstrap from released 1.3.1 and config/state preservation, and accept rollback plus Guardian recovery on a real supported system. |

### Evidence and documentation disposition

The accepted baseline is supported by [1.2.0 real-host acceptance](release/ACCEPTANCE_1.2.0.md),
[1.3.1 production acceptance](release/ACCEPTANCE_1.3.1.md), the
[authenticated migration contract](architecture/AUTHENTICATED_MIGRATION_AUTHORITY.md),
and [common mutation exclusion and recovery contracts](architecture/AUTOMATIC_UPDATE_ORCHESTRATION.md).
These records prove their stated scope, not completion of remaining gates.

Task 088 closes the scope-authority and obsolete-requirement disposition portion
of C8. C8 remains PARTIAL: operator recovery/readiness guidance must still be
checked against the eventual C2/C3/C7 fixes and C9 supported upgrade procedure.
No runtime gate is advanced by this documentation-only task.

[Task 089](../ai/history/089_2026-09-21_stable-privacy-preserving-service-identity.md)
closes C2 with collector-to-canonical-to-comparison/Drift/Report acceptance:
repeat and reorder preserve identity; disappearance, appearance and same-count
replacement preserve protected change references; raw names remain redacted;
retained ordinal evidence remains readable and unmodified with an explicit
comparison boundary. See the [identity/privacy contract](architecture/CANONICAL_SYSTEM_INVENTORY_V1.md#stable-protected-service-identity-task-089)
and [historical compatibility behavior](architecture/SNAPSHOT_COMPARISON_ENGINE.md#service-identity-compatibility).
No other gate status changes; C8's later readiness/recovery review remains PARTIAL.

[Task 090](../ai/history/090_2026-09-21_local-evidence-interruption-integrity-recovery.md)
closes C3 with nine interruption/integrity/recovery cases, including abrupt
pre-commit process exit, durable post-commit cleanup recovery, active/legacy
lock cases, limit+1 input measurement and repeated recovery. The
[persistence contract](architecture/INVENTORY_PERSISTENCE_AND_DIGITAL_TWIN.md#interruption-recovery-task-090-c3)
defines the commit boundary, artifact classification and operator steps for
ambiguous legacy locks. Existing evidence bytes remain compatible. C7 updater
recovery, C8 general documentation closure and C9 release acceptance are unchanged.

Current operational boundaries that documentation must preserve:

- Optional SMTP being unconfigured may produce overall readiness PARTIAL while
  mandatory Guardian requirements pass; that is not itself a core failure.
- Released 1.3.1 cannot retroactively gain Task 082's new metadata parser.
  C9 must prove an authenticated, explicitly authorized bootstrap to a capable
  client; the frozen-old-client fixture is not proof of that real-host upgrade.
- Task 087 verifies automatic handoff and separate manual rollback recovery
  within its scope. It does not close the manual update failure-path or journal
  durability work in C7. Package rollback and Guardian recovery are separate
  outcomes; uncertain recovery cannot be reported as success.
- Interrupted automatic work may retain incomplete evidence and leave Guardian
  stopped for operator review. Safe evidence, refusal and actionable recovery
  are required; autonomous replay and long health-observation windows are not.

## PRO 1.0 MUST HAVE

Pro 1.0 is the same trustworthy QWSG core with production-authorized Pro
capabilities and bounded unattended automation that Community intentionally
does not provide. It inherits **all nine Community gates** without weakening
their correctness, privacy, authority or recovery requirements.

### Additional Pro release gates

| Gate | Current state | Required completion boundary / remaining work |
| --- | --- | --- |
| P1 — Production Pro authority / entitlement | MISSING | Provide a minimal authenticated production authority path defining valid grants, missing/invalid authority, lost/revoked authority behavior and continued safe Community behavior. Billing, fleet management, hosted accounts and a commercial backend are not inherently required. |
| P2 — Explicit automatic policy and common security engine | PASS | Preserve existing capability/policy and authenticated update security boundaries. Configuration alone must not grant Pro authority; Pro uses the common authenticated engine. |
| P3 — Unattended privilege and handoff readiness | PARTIAL | Provide a narrow operator-authorized privilege path and deterministic readiness/preflight without an interactive password dependency. Detect insufficient privilege before unsafe Guardian handoff or mutation. |
| P4 — Bounded automation lifecycle and operator-visible outcome | PARTIAL | Bound automatic-update backup lifetime: protect active recovery material and prevent unbounded superseded backups. Distinguish success, rollback, recovery failure and incomplete state. Operators can inspect the last meaningful automation result and next action. |
| P5 — Real Pro end-to-end acceptance | MISSING | On a real supported system prove production Pro authority, automatic policy, release detection, privilege readiness, actual systemd handoff/worker, authenticated update, validation/recovery, repeated/current no-op, refusal paths, rollback and recovery-failure/incomplete behavior. Fixture-only Pro injection is insufficient. |

The [capability/policy foundation](architecture/PRODUCT_CAPABILITIES_AND_UPDATE_POLICY.md),
[automatic orchestration](architecture/AUTOMATIC_UPDATE_ORCHESTRATION.md) and
[Guardian trigger/recovery](architecture/AUTOMATIC_UPDATE_TRIGGER.md) are existing
engineering foundations. Production still resolves Community authority;
test-injected Pro does not satisfy P1 or P5. The PASS for P2 describes its
existing security/policy boundary, not production Pro availability.

## Current gate totals after Task 089

| Boundary | Total | PASS | PARTIAL | MISSING |
| --- | --- | --- | --- | --- |
| Community: C1–C9 | 9 | 6 | 2 | 1 |
| Additional Pro: P1–P5 | 5 | 1 | 2 | 2 |
| Pro including inherited Community gates | 14 | 7 | 4 | 3 |

Task 089 advances only C2 and Task 090 only C3 from PARTIAL to PASS. The Task 088 baseline was
Community 4/4/1 and inherited Pro 5/6/3 (PASS/PARTIAL/MISSING). Subsequent
governed work updates this register with evidence; completed task counts and
estimated future task numbers are not product-completion measures.

## POST-1.0

The following remain legitimate roadmap items, but are **not Community 1.0 or
Pro 1.0 blockers** without an explicit future Owner scope decision:

- **Monitoring/history:** long-term historical evidence database;
  CPU/memory/storage time-series; trend analytics; top resource consumers;
  broader threshold monitoring families; HTTP/TLS endpoint monitoring;
  backup-age monitoring.
- **Reporting:** daily, weekly and monthly reports; graphical/HTML/PDF periodic
  reporting; advanced recommendations. Existing local canonical reporting
  remains required by C5.
- **Security observation:** advanced security-log analytics; broader
  process/OOM/firewall/attack analytics.
- **Automation/notifications:** multiple notification recipients; additional
  notification transports; advanced maintenance windows/scheduling; autonomous
  crash replay.
- **Distribution:** native `.deb` packaging; RPM packaging; APT repository;
  broader distribution/platform support.
- **Management/UI:** web dashboard; central API; remote management;
  multi-server/fleet management.

This explicitly dispositions older daily-report and broad-monitoring MUSTs
for 1.0 completion without deleting their longer-term functional requirements.

## FUTURE

Longer-term product evolution outside the 1.0 boundary: provider/fleet platform;
tenancy; team/role management; enterprise compliance product; hosted control
plane; white-label product; billing platform integration beyond minimum P1
authority; AI adapters; autonomous remediation. Multi-server management above
is a POST-1.0 capability; a provider/tenant fleet platform is FUTURE evolution.

Neither deferred list authorizes implementation. Every future task retains
the established Owner authority and lifecycle requirements.
