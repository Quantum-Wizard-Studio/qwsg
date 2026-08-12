# Quantum Wizard Server Guardian Product Architecture

## Version 1.0 frozen product

The 1.0 release is the implemented local Linux Server Guardian: canonical
evaluation, bounded operator state, local Terminal Console, and continuously
supervised non-root Runtime Service. Dashboard, API, provider/fleet,
remediation, AI and commercial sections are direction, not 1.0 release gates.

Canonical Current Operator State is the single-record process-boundary handoff between typed projection and replaceable interfaces. It provides current observation continuity, not general history or monitoring persistence.

Canonical Operator Evaluation is the minimal one-shot product composition over the existing engineering core. `qwsg observe` establishes a baseline on first use, then runs the complete canonical live pipeline and publishes its typed Overview. Its condition is limited to implemented evidence; it is not proof of every server, application, backup, certificate, or hardware concern, and it does not claim continuous Guardian operation.

Operator projection is a bounded product boundary: globally ranked attention
retains the most severe and important evidence, correlated Rule/Policy views
are reduced through canonical identities, and overflow is disclosed rather
than hidden. Pipeline, projection, and publication failures have separate
privacy-safe operator diagnostics.

The local interface follows `Canonical Engineering and Operational Data -> Canonical Operator Presentation Model -> Interactive Operator Console`. Bare interactive `qwsg` opens this read-only Console; non-interactive invocation prints a concise nonblocking view.

> The Community Edition exists to earn trust. The Professional Edition exists to save time. The Provider Edition exists to operate at scale. Every edition shares the same deterministic engineering core.

## 1. Purpose and authority

This document is the canonical product architecture for the Quantum Wizard
Server Guardian (QWSG) ecosystem. It defines durable product identity, edition
boundaries, experience principles, deployment and trust models, ecosystem
seams, and the structure by which future capability is planned.

Every future product implementation task must conform to this document instead
of redefining product philosophy. More specific canonical contracts govern
their own domains: the Product Definition governs purpose and operator
authority; the Product & System Blueprint governs established product
boundaries; the Functional Specification governs approved observable Core
Alpha behavior; and documents under `docs/architecture/` govern technical
contracts. If a conflict is discovered, engineering stops for an explicit
architecture decision rather than silently choosing an interpretation.

This is a target product architecture, not a statement that every described
capability exists. Implemented status remains defined by the README, release
notes, tests, and capability-specific records. Task 019 changes no runtime
behavior and makes no release, price, service-level, platform-support, or
delivery-date commitment.

## 2. Product vision

QWSG is a **Professional Linux Server Engineering Toolkit**: a coherent system
for observing Linux hosts, preserving trustworthy evidence, explaining change,
evaluating future health and policy, and supporting controlled operations.
Monitoring is one future use of this system; it is not the product identity.

QWSG serves an individual administrator at one terminal, an operations team
managing several servers, and a provider operating separated customer fleets.
Those contexts differ in orchestration and operating scale, not in the quality
of underlying engineering facts.

The ecosystem has four conceptual layers:

1. **Deterministic Engineering Core** — shared collection, canonical inventory,
   persistence, comparison, and future factual evaluation contracts.
2. **Community Edition** — the complete local professional engineering toolkit.
3. **Professional Edition** — automation and management around the same core.
4. **Provider Edition** — tenant-aware fleet operations and integration around
   the same core.

Optional future ecosystem services may distribute releases, licenses, support
content, or explicitly selected data. They are adapters around the product,
never prerequisites for the deterministic core or Community operation.

## 3. Permanent engineering philosophy

### 3.1 Deterministic core

The same valid input, configuration, contract version, and execution context
must produce reproducible structured results. Time, host identity, and other
contextual values must be explicit data rather than hidden sources of
non-determinism. Human views derive from canonical records; they do not create
independent truths.

The canonical flow is:

```text
bounded evidence acquisition
  -> canonical inventory
  -> validated snapshots
  -> canonical change records
  -> future deterministic health and policy records
  -> reports, alerts, interfaces, automation, and optional AI
```

Each stage consumes a documented upstream contract. Downstream products may
coordinate, retain, visualize, route, or explain records; they may not bypass
the core and invent competing collection, comparison, or correctness logic.

### 3.2 Offline capable

Local inventory, snapshots, comparison, future deterministic analysis, reports,
exports, help, and manual workflows must remain useful without Internet access,
a vendor account, telemetry, remote licensing checks, or another Quantum Wizard
product. Network-dependent capabilities must degrade explicitly and must not
invalidate local evidence.

### 3.3 Privacy first

Collection follows data minimization. Sensitive values are classified,
redacted, pseudonymized, or excluded before persistence and presentation.
Local data stays local by default. Every transfer across a host, organization,
tenant, or vendor boundary is explicit, purpose-limited, inspectable, and
revocable where technically possible.

### 3.4 AI independent

No collector, canonical model, validator, snapshot, comparison, health
decision, policy result, license decision, or required operational workflow may
depend on an AI model. QWSG remains correct and complete when no AI component is
installed or reachable.

### 3.5 Engineering before automation

QWSG first establishes evidence, contracts, failure behavior, reproducibility,
and operator understanding. Automation may execute or schedule a proven
engineering workflow; it may never conceal uncertainty, make incomplete
evidence appear complete, or substitute convenience for correctness.

### 3.6 One core, additive editions

Edition boundaries are drawn above engineering correctness. Professional does
not provide a “better” Inventory, more truthful comparison, or more accurate
deterministic analysis. Provider does not own a private correctness tier.
Edition-specific components consume the same versioned public core contracts.

## 4. Product identity

The canonical category is **Professional Linux Server Engineering Toolkit**.
The Guardian identity promises:

- deterministic engineering rather than opaque scoring;
- evidence and explanation rather than unexplained conclusions;
- privacy and local authority rather than mandatory data extraction;
- conservative, reversible operations rather than silent intervention;
- simple, composable workflows rather than unnecessary platform dependence;
- professional quality in every edition;
- stable contracts and migrations suitable for long-lived servers.

QWSG is not “another monitoring application,” a hosted dashboard that happens
to install an agent, an AI administrator, or a replacement for accountable
Linux administration. Its interfaces should feel calm, precise, and
trustworthy. Urgency is represented by evidence and severity, never by visual
noise or marketing pressure.

## 5. Shared deterministic engineering core

Every edition uses the same core capability families:

- bounded evidence acquisition through the Collector Framework;
- Canonical System Inventory and compatibility projections;
- validated Digital Twin snapshot persistence;
- the exclusive Snapshot Comparison Engine and canonical Change Records;
- future Drift, Health, Rule, Policy, and Report engines built on those
  contracts;
- stable machine-readable schemas and terminal-safe human renderers;
- explicit capability, completeness, provenance, error, and version metadata;
- import, export, validation, migration, and diagnostic boundaries.

The core owns facts and deterministic judgments. Edition shells own workflow:
when work runs, where approved records travel, how many systems are coordinated,
who may see them, and how operators are notified.

A public interface means a versioned CLI, file/schema, library boundary, or
network API approved for that purpose. Direct access to collector internals,
private storage structures, or undocumented database tables is not an edition
integration contract.

## 6. Community Edition

### 6.1 Promise

Community Edition is not a limited edition, trial, teaser, or reduced-quality
product. It is a complete professional Linux engineering toolkit that Linux
administrators should genuinely enjoy using every day.

Community includes unlimited local use of the shared deterministic core and the
accepted QWSG 1.0 supervised Guardian, subject only to host resources and
supported-platform contracts. No account, license key, telemetry or network
connection is required. It remains offline capable and privacy first.

### 6.2 Capability boundary

Community is the natural home for:

- full Inventory, Snapshot, and Comparison Engines;
- full local Drift, Health, Rule, Policy and Report evaluation;
- local alerts, reports, exports, diagnostics and canonical private history;
- the local Guardian service and read-only Operator Console;
- a private local workspace and inspectable configuration;
- manual observation, comparison, analysis, validation, and reporting;
- supported installation, upgrade, rollback and uninstall;
- stable machine-readable output for operator-authored scripts;
- local import/export and documented interoperability;
- localized user-facing guidance.

The accepted local Scheduler, Runtime and Guardian lifecycle are Community
capabilities in QWSG 1.0. QWSG must not sabotage pipes, scripts, cron, systemd,
or other operator-owned tools. Future central policy, coordinated managed
notification, remote history, Dashboard/API and fleet management may belong to
commercial services because those add cross-host operational responsibility.
Loss or rejection of a future API key or service must not disable the local
Community Guardian or corrupt local canonical evidence.

### 6.3 Trust contract

Community results must be as correct, complete, inspectable, and reproducible as
the same results in paid editions. Security fixes, schema validation, privacy
controls, integrity verification, and correctness fixes are never commercial
differentiators.

## 7. Professional Edition

### 7.1 Promise

Professional Edition is the QWSG automation and management platform. It saves
operator time by repeatedly running, coordinating, retaining, presenting, and
routing Community-grade engineering workflows.

Professional adds convenience, automation, and multi-system management. It
never replaces or improves engineering correctness.

### 7.2 Capability families

Subject to separate specifications and implementation tasks, Professional may
provide:

- central scheduling and coordinated automation beyond the shipped local
  Guardian, with bounded retries, managed history and self-observability;
- alert policies, suppression and maintenance windows, recovery notifications,
  email and multiple recipients, and future pluggable delivery channels;
- longer or policy-managed history, graphs, trends, historical dashboards, and
  scheduled reports;
- centrally managed configuration profiles and policy distribution;
- multiple-server views, remote agents, enrollment, identity, health, and
  lifecycle coordination;
- a secured Web Dashboard and management API;
- signed automatic-update orchestration with explicit policy and rollback;
- offline-capable license activation and entitlement inspection;
- premium support and privacy-reviewed diagnostic workflows.

Every automation action exposes its trigger, input record, policy, outcome, and
failure state. Destructive or corrective action remains separately authorized;
buying Professional is not authorization to modify a server.

Task 028 implements only the pure Alert decision subset of this architecture:
canonical Alert Records, lifecycle state, acknowledgement, bounded suppression
and maintenance evaluation, deduplication, expiration, correlation, and
recovery. Alert persistence, background monitoring, notification delivery,
recipients, channels, delivery audit, and interfaces remain future components.

Task 029 implements the provider-neutral delivery subset downstream of those
Alert Records: deterministic routing and fan-out, bounded queue/retry proposals,
idempotent delivery requests, canonical attempt/status/acknowledgement/evidence
records, and an explicitly invoked one-cycle injected-provider adapter. It does
not implement a concrete transport, durable queue, daemon, monitoring, Alert
creation, configuration/secret integration, or user interface.

### 7.3 Failure boundary

Failure of the scheduler, dashboard, notification transport, management plane,
subscription service, or network must not corrupt core records or prevent local
manual engineering. Professional must distinguish “not evaluated,” “evaluated
with incomplete evidence,” “notification not delivered,” and “management plane
unreachable.”

## 8. Provider Edition

### 8.1 Promise

Provider Edition adds operational capabilities for hosting providers, managed
service organizations, and large enterprise operators. It applies the same core
and Professional automation patterns across separated tenants and fleets.

### 8.2 Capability families

Subject to future architecture and threat-model approval, Provider may provide:

- multi-tenant organizations, projects, roles, delegated administration, and
  immutable audit boundaries;
- fleet enrollment, segmentation, policy rollout, staged updates, capacity
  views, and operational health of agents;
- provider-wide dashboards with tenant-scoped customer portals;
- white-label presentation without concealing QWSG evidence provenance or
  weakening security disclosures;
- integration APIs, webhooks, service-management connectors, and billing or
  entitlement integration;
- managed-hosting workflows, templates, compliance evidence, and support
  operations;
- high-availability control-plane, backup, disaster-recovery, and regional
  deployment patterns.

### 8.3 Tenant and control-plane boundary

Tenant isolation is a security property, not a UI filter. Identity, storage,
encryption, authorization, audit, export, retention, deletion, backup, and
support access must be tenant-aware by construction and independently tested.

A Provider control plane coordinates agents; it does not become the source of
truth for observations made on a host. An enrolled host retains validated local
records and explicit degraded behavior when disconnected. Cross-tenant
aggregation requires an approved, minimized derived-data contract.

## 9. Edition comparison

| Architectural concern | Community | Professional | Provider |
| --- | --- | --- | --- |
| Primary promise | Complete local engineering | Supported automation and management | Tenant-aware fleet operations |
| Deterministic core | Full shared core | Same core | Same core |
| Engineering correctness | Full | Identical | Identical |
| Manual local operation | Unlimited | Included | Included where role permits |
| Offline local core | Required | Required | Required for enrolled hosts |
| Account requirement | None | License or management identity only where needed | Organization and tenant identity for control-plane use |
| QWSG-managed scheduling | Accepted local 1.0 scheduling | Central coordination | Fleet-scale coordination |
| Background service | Accepted local Guardian | Centrally managed | Supported and orchestrated |
| Web management | Not required | Single-organization/multi-server | Multi-tenant/provider |
| Notifications | Manual outputs and operator composition | Managed policies and channels | Tenant/fleet routing and integrations |
| Fleet management | No | Multi-server | Multi-tenant and large-fleet |
| White label/customer portal | No | No | Optional |
| Commercial value | Trust and adoption | Time saved | Operations at scale |

This table defines product obligations, not technical sabotage. Community
artifacts remain open to documented, operator-controlled composition.

## 10. Workspace architecture

### 10.1 Canonical user workspace

The target default per-user workspace is:

```text
~/.qwsg/                         private directory, target mode 0700
├── config/                      operator-owned declarative configuration
├── snapshots/                   validated local Digital Twin observations
├── reports/                     generated human-oriented reports
├── exports/                     explicit interoperability outputs
├── logs/                        bounded local operational logs
├── cache/                       disposable, reconstructable data
├── state/                       durable workflow and engine state
└── run/                         ephemeral locks and local runtime metadata
```

This is a target product-level layout. The implemented Inventory Store
currently uses an operator-supplied absolute root with its own canonical
`store.json` and `snapshots/` contract. Adoption or migration to `~/.qwsg/`
requires a separate specification and must preserve compatibility; this
document does not change the current CLI default or storage behavior.

System-wide deployments use an explicitly configured service workspace under a
platform-appropriate state directory, owned by the dedicated QWSG service
identity. Provider control-plane storage is not an extension of a user's home
workspace and requires a separate data architecture.

### 10.2 Ownership and permissions

- Private directories default to `0700`; sensitive regular files default to
  `0600`, subject to a stricter platform contract.
- QWSG refuses unsafe ownership, symlink, traversal, or group/world-accessible
  sensitive paths rather than silently weakening them.
- Privileged and unprivileged workspaces are never silently mixed.
- Secrets do not belong in ordinary configuration, reports, logs, or exports.
- Export is an explicit disclosure action and preserves classification and
  provenance where the format supports it.
- Cache may be removed without losing canonical truth; snapshots, state,
  configuration, and audit records require defined retention and migration.

### 10.3 First run

First run must be inspectable and idempotent. It explains the intended paths and
permissions, creates only the minimum approved local structure, does not require
an account or network, does not scan or persist host data without the invoked
command's authority, and never imports legacy state implicitly. Existing
unsafe or incompatible state causes a clear stop with recovery guidance.

## 11. Terminal experience

### 11.1 Command-line contract

The CLI remains a stable, scriptable public interface. Explicit commands such as
`qwsg inventory` continue to support contextual help, predictable exit status,
machine-readable output, terminal-safe human output, and non-interactive use.
Interactive convenience must not make automation ambiguous.

The command vocabulary should follow operator intent:

```text
qwsg inventory
qwsg snapshot ...
qwsg compare ...
qwsg drift ...
qwsg health ...
qwsg report ...
qwsg config ...
qwsg workspace ...
```

Only implemented commands may appear as available product behavior. Future
names are architectural examples and require individual contract approval.

### 11.2 Bare `qwsg`

The long-term default for `qwsg` in an interactive terminal may be a
context-aware home experience. In a non-interactive context it must remain
predictable and must not unexpectedly start a full-screen interface. Flags or
configuration must permit an explicit choice between help, concise status, and
the future Terminal UI.

### 11.3 Interaction principles

- keyboard-first, fast, and usable over SSH and constrained terminals;
- clear current scope, selected host, time, profile, and data freshness;
- consistent navigation, terminology, colors, severity, and exit behavior;
- no color-only meaning; readable monochrome and reduced-color modes;
- responsive to terminal size, with plain-text fallback;
- screen-reader-conscious labels and a complete non-interactive equivalent;
- confirmation proportional to risk, with previews for state-changing actions;
- copyable evidence identifiers and direct routes to structured records;
- localized user text while commands, keys, and schemas remain stable.

## 12. Future Terminal UI philosophy

The Terminal UI (TUI) is a view and workflow coordinator over public
application contracts. It must not embed alternative collection, comparison,
health, policy, or licensing logic.

A durable information architecture is:

```text
Home
├── Health
├── Inventory
├── Change and Drift
├── History
├── Reports
├── Alerts
├── Configuration
└── Workspace and Diagnostics
```

Community receives the complete local Operator Console and Guardian views.
Professional may add central scheduling, multi-server, managed notification,
and licensing views. Provider may add tenant and fleet context. Edition changes
should add navigation domains, not replace familiar local workflows.

Every important TUI operation must have an equivalent explicit command or
public API operation. The TUI displays loading, stale, partial, unsupported,
empty, and failure states distinctly. It never turns unknown evidence into a
healthy state.

## 13. Web Dashboard vision

The Professional Web Dashboard is a secured remote and historical management
surface, not a second engineering engine. Its conceptual domains are:

- Overview and attention queue;
- Health and evidence;
- Inventory and relationships;
- History, Change, and Drift;
- Alerts, delivery, suppression, and maintenance;
- Storage, Network, Processes, and other canonical resource views;
- Reports and exports;
- servers, agents, profiles, and policy;
- users, roles, audit, workspace, and diagnostics;
- licensing, updates, and support.

Provider extends this information architecture with tenant selection, fleet
segments, customer portals, delegated roles, integrations, and provider
operations.

The browser never receives collector or shell authority. A versioned
application/API boundary enforces identity, authorization, tenant scope,
redaction, rate and resource limits, audit, and schema compatibility. Dashboard
unavailability must not stop local collection or invalidate local truth.
Accessibility, localization, responsive layout, and low-bandwidth operation are
architecture requirements from the first implementation.

## 14. Licensing philosophy

### 14.1 Boundary

Licensing may control access to future central, remote and managed service
capabilities. It does not control the accepted local Community Guardian,
correctness, security fixes, access to an operator's local canonical records,
or the ability to validate and export those records.

Community requires no account, activation, telemetry, or Internet connection.
Professional and Provider may use licenses or subscriptions, but entitlement
checking is outside the deterministic engineering core.

### 14.2 Offline activation

Professional must support a documented offline activation path suitable for
isolated servers. A future design may use a signed, bounded entitlement file
with edition, features, validity, and non-secret installation binding. Private
signing material never ships with the product. Online activation is optional
convenience, not the only recovery path.

License evaluation must be deterministic and inspectable locally. Clock
rollback, renewal, transfer, replacement, grace periods, and control-plane
outage require explicit behavior before implementation.

### 14.3 Expiry and upgrade

License expiry must fail safely: stop initiating paid automation according to a
documented grace policy, preserve local data, preserve Community workflows, and
provide export and recovery access. It must not corrupt snapshots, hide
evidence, or make the server unmanageable.

Upgrades are additive and reversible at the product boundary. Moving from
Community to Professional or Provider reuses the same core records. Downgrading
removes edition orchestration after an explicit transition while preserving
operator-owned local data in documented formats.

Final pricing, license text, feature packaging, support terms, and legal
obligations remain owner and legal decisions outside this architecture.

## 15. Privacy model

### 15.1 Data authority

The operator or tenant controls its data. QWSG collects only evidence required
for an enabled, documented purpose. Absence of telemetry is the default.
Consent to licensing, updates, support, notification delivery, or optional AI
is separate; none implies consent to unrelated data transfer.

### 15.2 Data classes and movement

Future data contracts must classify at least:

- public product metadata;
- operational metadata;
- host-identifying data;
- sensitive configuration and topology;
- credentials and secrets;
- derived aggregate data;
- audit and support evidence.

Each boundary defines collection purpose, fields, redaction, encryption,
retention, deletion, export, access roles, and failure behavior. Raw secrets,
environment values, command payloads, and unrelated file contents remain
prohibited unless a future dedicated contract explicitly requires and protects
them.

### 15.3 Telemetry and support

Telemetry is opt-in, documented, minimal, and independently disableable.
Support bundles are locally generated, previewable, redacted, and transferred
only through an explicit operator action. Provider support access is
time-bounded, role-controlled, tenant-visible, and audited.

## 16. Deployment models

### 16.1 Local standalone

Community's reference model is one installation on one Linux host, invoked by
the local operator, with a private local workspace and no network dependency.
This is the foundational model and remains supported as higher editions grow.

### 16.2 Local managed automation

Professional may add a dedicated least-privilege service identity, scheduler,
daemon, local state, and notification egress. Interactive local operation
continues against the same core contracts. Privileged helpers, if ever needed,
are narrowly scoped and independently threat-modeled.

### 16.3 Multi-server management

Professional may connect enrolled agents to a customer-controlled management
plane. Enrollment establishes explicit identity and trust; transport is
authenticated and encrypted; queued and disconnected operation is bounded; the
agent retains local truth and exposes only approved redacted contracts.

### 16.4 Provider control plane

Provider may be self-hosted, provider-managed, or offered as a future managed
service only after tenancy, regionality, backup, recovery, observability,
upgrade, and exit/export contracts exist. Control-plane failure cannot grant
broader agent authority or merge tenant data.

### 16.5 Packaging neutrality

Native packages, a self-contained binary, images, appliances, or orchestrated
deployments are delivery mechanisms. None may redefine product semantics.
Supported status requires an explicit platform matrix and verified lifecycle;
“platform independent” means contract and architecture portability, not an
unsupported claim that every operating system is implemented.

## 17. Automation philosophy

Automation is a transparent state machine around deterministic operations:

```text
declared trigger
  -> authenticated scope
  -> versioned profile and policy
  -> precondition validation
  -> deterministic core operation
  -> immutable result and audit
  -> notification or approved follow-up
```

Every run has a stable identity, initiator, target, configuration version,
timestamps, outcome, and links to evidence. Retries are bounded and idempotent
where possible. Concurrency, locking, missed schedules, backlog, clock changes,
partial fleets, cancellation, and recovery are explicit design inputs.

Observation never grants remediation authority. Recommendations, previews,
approval, execution, verification, and rollback are separate records. Future
automatic remediation requires a dedicated architecture, risk class, least
privilege, bounded target, opt-in policy, dry-run where meaningful, and
post-action verification.

## 18. AI separation

AI is an optional consumer/advisor layer:

```text
canonical public records -> minimized AI adapter -> explanation or proposal
                                              \-> never canonical truth
```

An AI component may summarize records, explain terminology, help query history,
draft reports, or propose investigation and remediation plans. Its output is
labeled advisory, includes provenance to deterministic evidence, and cannot
silently alter canonical records, severity, policy, licenses, configuration, or
hosts.

Local and remote AI providers are replaceable adapters. Data selection and
redaction occur before the adapter boundary. Remote transmission requires
explicit configuration and operator action appropriate to the workflow. Core
tests, builds, operation, and support do not require an AI service.

No edition may market probabilistic output as improved engineering correctness.
Paid AI convenience, if ever offered, remains separate from the shared core.

## 19. Future ecosystem

The ecosystem may grow through:

- a signed package and update channel;
- collector, rule, report, notification, and integration extension kits;
- documented schemas, CLI contracts, SDKs, and APIs;
- a privacy-reviewed catalog or registry;
- policy/profile libraries and report templates;
- training, certification, professional services, and support;
- optional managed control-plane and notification services;
- optional AI adapters.

Extensions declare identity, version, compatibility, capabilities, authority,
resource bounds, data classes, and provenance. Trust or signature verifies
publisher and integrity; it does not grant privilege. Extensions cannot bypass
Registry execution, canonical validation, tenant isolation, licensing
boundaries, or audit.

QWSG remains independently useful if every ecosystem service is unavailable.
Open formats and documented export prevent operational lock-in. Another Quantum
Wizard product may integrate only through the same public boundary available to
other approved consumers.

## 20. Long-term roadmap structure

The roadmap is organized by capability streams rather than speculative dates or
edition-led code forks:

1. **Engineering Core** — evidence, inventory, snapshots, comparison, drift,
   health, policy, reporting, contracts, safety, performance, and portability.
2. **Community Experience** — workspace, CLI, TUI, configuration, local reports,
   diagnostics, localization, packaging, and release readiness.
3. **Automation and Professional** — scheduler, daemon, history, alerts,
   dashboard, multi-server management, updates, licensing, and support.
4. **Provider Operations** — tenancy, fleet control, portals, integrations,
   white label, high availability, disaster recovery, and compliance evidence.
5. **Platform and Ecosystem** — public APIs, extension contracts, distribution,
   optional services, and AI adapters.

Movement through a stream uses architecture gates:

```text
product obligation
  -> threat and privacy model
  -> canonical contract
  -> bounded implementation slice
  -> verification and migration
  -> documented support decision
```

Dependencies flow from core contracts outward. Commercial priority may choose
which already-safe outer capability to implement next; it may not skip the
correctness, security, privacy, rollback, or support gates.

## 21. Engineering principles

All future decisions preserve these rules:

1. The deterministic core is the single source of engineering truth.
2. Community is complete professional engineering, not a crippled edition.
3. Professional extends automation; it never improves correctness.
4. Provider extends operational scale; it never weakens tenant isolation.
5. Offline local operation is foundational, including after network failure.
6. Privacy defaults to local, minimal, explicit, and inspectable.
7. AI is optional, advisory, replaceable, and outside canonical decisions.
8. Reproducible structured records precede every human view and automation.
9. Engineering contracts and failure behavior precede convenience features.
10. Operator authority is explicit; observation is not remediation permission.
11. Public versioned interfaces connect modules and editions.
12. Unknown, unsupported, stale, partial, and failed are distinct from healthy.
13. Security and correctness fixes are available across editions.
14. User data remains exportable in documented formats.
15. Localization and accessibility are architectural concerns, not polish.
16. Long-lived migrations and downgrade behavior are designed before release.
17. Claims of support require verified evidence, not architectural aspiration.
18. Every new capability declares its authority, data, failure, rollback, and
    audit boundaries.

## 22. Architecture governance

This document owns durable ecosystem and edition decisions. Capability
documents may refine but not weaken them. A future task that introduces a new
edition boundary, mandatory cloud dependency, data transfer, AI authority,
remediation capability, tenant model, licensing restriction, or incompatible
workspace/schema change requires explicit Project Owner approval and an
architecture decision with migration and rollback.

Examples in this document do not authorize implementation. Each capability
still requires a bounded engineering task, current-state verification,
snapshot, threat/privacy analysis where relevant, canonical contracts, tests,
documentation, and a truthful support decision.

The architecture is successful when QWSG can expand for decades without
forking its engineering truth, extracting operator trust, or making automation
more authoritative than evidence.

## Runtime coordination

The Canonical Runtime Engine performs one explicit deterministic cycle across
existing Scheduler, Pipeline evidence, Alert, and Notification contracts. It
adds orchestration and partial-result evidence only. Continuous operation,
service lifecycle, persistence, and monitoring remain outside Runtime 1.0.

The Canonical Runtime Service provides the continuous local process foundation
by repeatedly invoking that one-cycle boundary at fixed nominal intervals. The
Operational Guardian Service supplies the narrow production composition:
foreground non-root execution, systemd user supervision, canonical
configuration, private bounded restart handoff, generation-correlated exit
demotion, and Current Operator State publication. Production delivery and
release hardening remain separate concerns.

## Canonical operator presentation

Canonical engineering and operational data now feeds one presentation-
independent Operator Presentation Model before any replaceable interface. The
model preserves owned Health, change, Alert, Runtime, and Service meanings,
including missing, stale, partial, unsupported, and not-observed states. Future
CLI, Terminal Console, REST API, and Web Dashboard adapters must consume this
shared model rather than create interface-specific status semantics. The model
is not monitoring, persistence, a renderer, or a current-state database.
