# QWSG Product & System Blueprint

## 1. Document Purpose and Authority

This document is the authoritative product-level blueprint for Quantum Wizard Server Guardian (QWSG). It translates the Project Philosophy, Product Definition, and original comprehensive planning material into durable boundaries and high-level system direction for later specification, architecture, implementation, validation, packaging, deployment, and roadmap work.

The Project Constitution remains supreme. The Product Definition remains the parent statement of product intent. Where early planning examples suggest implementation details, this blueprint classifies them as illustrative or deferred rather than final. The following labels apply throughout:

- **Principle:** settled, mandatory product direction.
- **Product requirement:** required behavior or outcome, without prescribing implementation.
- **Preferred direction:** current recommendation requiring validation in later architecture.
- **Illustrative example:** explanatory, not a commitment.
- **Deferred decision:** deliberately reserved for a later milestone.
- **Open question:** requires evidence or Project Owner direction.

## 2. Executive Summary

QWSG is an independent, modular Linux server guardian for small-server operators and a future path to professional multi-server operations. It detects its environment, evaluates actual operational outcomes, retains state, alerts on meaningful transitions, explains evidence, and performs corrective action only when explicitly authorized.

The product has three major parts. The **Agent** is the useful standalone guardian on each server. The **Installer** performs transparent, consent-based privileged lifecycle operations. The optional **Console** provides secure administration, history, and visualization without becoming a hidden root shell. Shared product responsibilities cover detection, execution coordination, checks, state, alerts, reporting, configuration, secrets, storage, audit, diagnostics, dependencies, updates, and removal.

The first usable release targets Ubuntu and Debian and focuses on system detection; disk, inode, memory, swap, load, systemd, HTTP/HTTPS, SSL, and existing-backup checks; e-mail alerts driven by state transitions; recovery and bounded emergency reminders; daily reporting; local state; lifecycle tooling; diagnostics; and a secure standalone Console direction. Broader checks, channels, platforms, automation, central fleet management, and convenience extensions follow after the core is trustworthy.

## 3. Product Origin

QWSG originated from a preventable production failure pattern: a Linux server resource such as disk space approaches exhaustion, but the operator learns about it only after services degrade. The initial idea was a small reusable health script. The planning work showed that durable protection also requires environment detection, consistent state, controlled notifications, lifecycle management, modular checks, transparent dependencies, and approachable administration. This evolved into an independent product rather than a host-specific script.

## 4. Problem Statement

Small Linux servers often combine web, database, mail, certificate, and backup responsibilities without dedicated operations staff. Existing signals are fragmented, repetitive, or limited to process existence. Operators need early, trustworthy warnings, evidence about real service outcomes, low notification noise, and safe guidance without granting an opaque tool unrestricted authority.

QWSG addresses that gap. It does not replace professional judgment; it makes relevant server state visible, understandable, historically traceable, and actionable before routine degradation becomes an outage.

## 5. Mission

QWSG protects compatible Linux servers by behaving like a careful and experienced administrator: observe, verify, analyze, explain, warn, and act only within explicit authorization.

## 6. Vision

QWSG should grow from a dependable standalone guardian for one server into a maintainable platform that can oversee many independent servers while preserving local usefulness, operator control, product independence, and clear trust boundaries.

## 7. Product Philosophy

- Stability takes priority over convenience; security takes priority over shortcuts.
- Verify available facts instead of guessing.
- Show evidence, planned changes, and outcomes.
- Prefer prevention and early detection to disruptive recovery.
- Alert on meaningful change, not every polling cycle.
- Keep optional capability optional; degraded capability must be explicit.
- Make installation, operation, update, diagnosis, and removal understandable.
- Keep the core serious and dependable while allowing safely bounded extensions.

## 8. Target Users

The initial audience is the technically capable owner or administrator of one or several Linux servers who needs dependable protection without a dedicated operations team. Secondary audiences are hosting operators, developers responsible for deployed services, managed-service providers, security reviewers, and contributors extending checks or integrations.

The interface should remain approachable for occasional operators without concealing the evidence and controls required by experts.

## 9. Supported Deployment Contexts

QWSG is designed for independent compatible Linux servers, including VPS instances, dedicated servers, and physical machines. It must not depend on QUWIP, the Quantum Wizard website, Hestia, a particular hosting provider, or a specific VPS.

The initial supported distributions are Ubuntu and Debian. A deployment may be a general-purpose server, web server, mail server, database host, or mixed-role server. Hardware-only capabilities are conditional because VPS environments frequently do not expose them.

## 10. Product Goals

- Detect environmental capabilities and relevant operational risks.
- Evaluate actual outcomes where feasible, not merely process presence.
- Provide low-noise state-transition, escalation, recovery, and reminder notifications.
- Remain useful as an Agent-only installation.
- Provide a safe privileged installation and lifecycle path.
- Offer an independent, secure, localization-ready Console.
- Preserve understandable configuration, evidence, history, and auditability.
- Support modular growth across server roles, integrations, and deployment scale.

## 11. Explicit Non-Goals

- QWSG is not a hidden autonomous root administrator.
- It is not tied to or embedded inside another Quantum Wizard product.
- It is not initially a backup creation product; it first monitors existing backups.
- It is not a replacement for a full security program, disaster-recovery plan, or compliance system.
- It does not promise every check on every environment.
- The MVP does not include automatic security repair or broad automatic remediation.
- This blueprint does not select languages, frameworks, protocols, schemas, package layouts, or final process topology.

## 12. Immutable Design Principles

1. Product independence and standalone Agent usefulness.
2. Explicit consent before optional dependency installation or system modification.
3. Preview of planned changes before they occur.
4. Opt-in, bounded, observable remediation; observation is the default.
5. State-transition-based alerting with recovery and controlled reminders.
6. Correct Linux semantics, including treating reclaimable memory cache appropriately.
7. Outcome-oriented service checks where practical.
8. Secure separation of privileged lifecycle work from routine administration.
9. No plain-text sensitive credentials in ordinary configuration.
10. Auditability for every system modification.
11. Capability-dependent optional hardware checks.
12. First-class installation, update, diagnostics, migration, and removal.
13. Localization-ready user experiences and English engineering artifacts.
14. Reversibility where technically possible and explicit disclosure where it is not.

## 13. Trust, Safety, and Privilege Principles

Routine monitoring should run with the least privileges sufficient for its enabled capabilities. Privileged bootstrap and system modification belong to the Installer or another narrowly defined privileged boundary, not to arbitrary Console requests. The Console must expose supported operations through validated product controls and must never provide a disguised general root shell.

Every proposed modification must identify its purpose, target, dependency, privilege requirement, expected effect, and rollback implications before approval. Every attempted modification must yield an auditable result. Automatic remediation requires explicit enablement per bounded capability, retry limits, outcome verification, and escalation when it fails.

## 14. Product Boundaries

QWSG owns detection, enabled checks, check coordination, normalization into product state, transition evaluation, notification orchestration, reporting, its own configuration and protected secrets references, its own local history, audit evidence, diagnostics, and its product lifecycle.

QWSG observes but does not automatically own external services, backups, certificates, firewalls, package policy, or application configuration. Integration with those systems must respect their authority. Future central management may coordinate QWSG instances but must not eliminate local operation or silently broaden privileges.

## 15. Functional Domain Map

| Domain | Product responsibility |
| --- | --- |
| Environment and inventory | Detect platform, roles, capabilities, installed services, and constraints. |
| Execution coordination | Schedule and coordinate enabled checks without uncontrolled overlap. |
| Checks and modules | Gather evidence and return normalized outcomes with capability status. |
| State | Retain current and previous states and evaluate transitions. |
| Alerts and notifications | Route escalation, recovery, reminders, and delivery outcomes. |
| Reporting | Produce understandable periodic and on-demand summaries. |
| Configuration and secrets | Validate operator intent and protect sensitive material. |
| Storage and history | Retain bounded operational evidence and trends. |
| Audit | Record product and system modifications and authorization context. |
| Dependencies | Inspect requirements and obtain consent before any installation. |
| Diagnostics | Explain product health and generate redacted support evidence. |
| Lifecycle | Install, update, migrate, verify, repair, and remove QWSG safely. |

## 16. High-Level Product Model

Information flows from environment inventory and enabled checks into normalized observations. The state responsibility compares observations with retained state, classifies transitions, and supplies alerting, history, reporting, and Console views. Operator policy flows through validated configuration. Modification requests flow through authorization and privilege boundaries to lifecycle or explicitly enabled remediation operations. Audit evidence accompanies every product-controlled change.

The model is conceptual: a later architecture decides process boundaries, communication mechanisms, storage technology, schedules, failure isolation, and deployment topology.

## 17. QWSG Agent

The Agent is the mandatory operational core on a monitored server. It inventories the environment, coordinates enabled checks, evaluates observations, persists enough state to detect transitions, produces alerts and reports, records operational evidence, and exposes diagnostics through supported interfaces.

The Agent must remain fully useful without the Console. Loss or absence of the Console must not disable local checks or core alerting. The Agent does not gain unrestricted repair authority merely because it is installed.

## 18. QWSG Installer

The Installer is the trusted lifecycle assistant for bootstrap, capability inspection, dependency consent, installation planning, configuration initialization, update, migration, repair, diagnostics, and removal. It handles privileged operations that a browser cannot safely perform.

The preferred initial experience is an interactive terminal installer with non-interactive automation considered later under equivalent safety guarantees. TUI technologies shown in the source plan are examples, not selections. The Installer must support preview, explicit consent, partial-failure reporting, verification, and bounded rollback or recovery instructions.

## 19. QWSG Console

The Console is an optional, independent QWSG administration and visualization surface. It presents current health, evidence, active incidents, history, reports, module capability, configuration, alert delivery, updates, and diagnostics. It may guide supported changes but cannot bypass Installer or Agent authorization and privilege boundaries.

The Console direction is secure standalone deployment with strong authentication, session controls, brute-force resistance, auditable changes, and the ability to restrict exposure, including localhost or reverse-proxy arrangements. Exact authentication, network, framework, and deployment designs are deferred.

## 20. Shared Core Responsibilities

The product requires conceptual responsibilities for detection and inventory, scheduling or execution coordination, module execution, state management, alert management, reporting, dependency inspection, configuration validation, secrets handling, local storage and history, audit logging, diagnostics, and lifecycle coordination.

These are responsibilities, not prescribed services or directories. Later architecture must assign ownership and interfaces while avoiding duplicated authority and single points that unnecessarily disable Agent-only operation.

## 21. Modules and Extension Model

Modules extend evidence collection, policy evaluation, reporting, notification, integration, or carefully bounded lifecycle recipes. A module must declare capabilities, prerequisites, privilege needs, configuration, data sensitivity, outputs, failure behavior, and compatibility expectations in a later-defined form.

Serious domains include disk, memory, CPU/load, services, network, HTTP, SSL, backups, databases, mail, logs, security reporting, and hardware health. Optional harmless convenience extensions—such as MOTD information, shell conveniences, `sl`, `cowsay`, or `fortune`—are preserved as post-MVP examples that test modularity. They must be explicitly selected, isolated from the protection core, and never silently installed.

## 22. State and Severity Model

The conceptual severity set is `OK`, `WARNING`, `CRITICAL`, `EMERGENCY`, and `UNKNOWN`. Modules may provide evidence and a proposed status; common state policy must preserve consistent transition semantics. `UNKNOWN` represents insufficient or failed observation and must not be silently treated as healthy.

Thresholds and persistence rules may vary by resource and environment. Example disk thresholds such as 80/90/95 percent are illustrative defaults pending functional specification and validation. Hysteresis, observation windows, flapping control, maintenance states, acknowledgement, and aggregation rules are deferred.

## 23. Alerting and Recovery Philosophy

Notifications are primarily triggered by meaningful state entry, escalation, recovery, and delivery failure. Repeated polling in an unchanged state should not create repeated noise. Controlled reminders are permitted for unresolved emergency conditions according to explicit policy. De-escalation and full recovery must be distinguishable so operators understand whether risk remains.

Alert content should identify the server, affected subject, old and new state, evidence, time, policy basis, suggested next action, and whether any remediation was attempted. Delivery channels must report their own success or failure. E-mail is required for MVP; additional channels are post-MVP.

## 24. Reporting

QWSG provides periodic and on-demand summaries of overall health, resource state, services, checks, active incidents, recoveries, backups, certificates, and check freshness. The MVP includes a configurable daily report and an alert-only mode. Weekly reports, richer trends, comparisons, and fleet summaries are later capabilities.

Reports must distinguish healthy, unhealthy, unknown, disabled, unsupported, and stale data rather than collapsing them into a misleading overall `OK`.

## 25. Configuration and Secrets

Configuration expresses desired checks, targets, thresholds, notification policies, reporting choices, operating profile, and bounded remediation consent. It must be validated before activation and support understandable defaults and explicit overrides.

Secrets are not ordinary configuration. Credentials and private tokens must use protected handling appropriate to the eventual architecture; they must not appear in normal configuration exports, logs, diagnostics, reports, or repository files. The final secrets backend, configuration format, precedence model, and reload behavior are deferred.

## 26. Data, History, Logs, and Audit

- **Current state** supports transition decisions and must survive routine restarts.
- **History** supports incidents, recoveries, trends, and reports under bounded retention.
- **Operational logs** explain check and product behavior without leaking secrets.
- **Audit records** capture authorization and results for configuration, dependency, lifecycle, and remediation changes.
- **Cache** may improve performance but must not be the sole authoritative state.

The MVP requires an implementation-neutral local state store; SQLite is the preferred early candidate, not a final schema decision. External databases, retention periods, rotation, export, integrity mechanisms, and fleet aggregation are deferred.

## 27. Diagnostics and Supportability

QWSG must expose its version, platform compatibility, configuration validity, enabled capabilities, check freshness, scheduler health, notification health, storage health, and recent errors. A diagnostic command or Console action should produce a redacted support bundle containing relevant system facts, product configuration without secrets, recent logs, service state, and error identifiers.

Diagnostic generation must preview included categories, exclude secrets by design, and disclose any information that may identify a host or deployment.

## 28. Installation, Update, and Removal Lifecycle

Installation begins with environment and privilege verification, followed by an explicit plan and dependency choices. Completion requires verification of installed components and enabled checks. Updates require compatibility checks, migration planning, rollback or recovery guidance, integrity verification, and clear release information. Removal must distinguish QWSG-owned artifacts from operator data and external dependencies, preserve or explicitly offer export of useful history, and avoid deleting shared dependencies or external configuration automatically.

Repair and reconfiguration are first-class lifecycle states. Interrupted operations must produce resumable or safely repeatable guidance rather than an ambiguous installation.

## 29. Supported Platform Strategy

Ubuntu and Debian are the initial supported distributions. Support must be expressed as tested distribution/version/capability combinations rather than broad assumptions. Detection should identify architecture, kernel, init system, package manager, virtualization or physical context, installed services, storage, network, and relevant monitoring tools.

AlmaLinux, Rocky Linux, Fedora Server, CentOS Stream, and wider Linux support are post-MVP directions. Platform adapters and capability declarations are preferred conceptually; exact mechanisms await architecture. Unsupported or partially supported environments must fail clearly and without mutation.

## 30. Operating Profiles

- **Agent-only:** mandatory core protection, local configuration, alerts, reports, history, and diagnostics without a Console.
- **Installer-assisted:** Agent-only or Console-enabled installation managed through the privileged lifecycle assistant.
- **Console-enabled:** Agent plus optional secure administration and visualization.

Suggested installation presets include Minimal, Web Server, Mail Server, Full, and Custom. Presets are starting selections, not distinct products or irreversible editions. Environment detection may recommend but must not silently enable modules.

## 31. MVP Definition

The first usable release must provide a coherent, operable slice rather than every planned feature:

- verified Ubuntu and Debian support boundaries;
- environment and capability detection;
- dependency inspection, change preview, and explicit consent;
- interactive TUI installer direction;
- disk capacity and inode checks;
- Linux-aware memory and swap checks;
- CPU load checks normalized to available capacity;
- selected systemd service checks;
- HTTP/HTTPS outcome and SSL-expiry checks;
- existing-backup age and size checks;
- e-mail alert delivery;
- retained state with transition, escalation, recovery, and controlled emergency reminders;
- configurable daily report and alert-only operation;
- implementation-neutral local state storage, with SQLite as a preferred candidate;
- secure standalone Console direction, while Agent-only remains complete;
- installation, update, removal, diagnostics, operational logging, and audit evidence.

MVP acceptance requires coherent end-to-end behavior across detection, observation, state, notification, recovery, diagnostics, and lifecycle—not merely the presence of individual checks. Automatic remediation, broad security repair, fleet management, and optional entertainment modules are excluded.

## 32. Post-MVP Capability Groups

- **Operational depth:** resource trends, I/O wait, OOM analysis, process outliers, log aggregation, port and network changes.
- **Service depth:** databases, mail queues and protocols, DNS, richer web content and certificate-chain checks.
- **Hardware:** SMART, NVMe, RAID, temperatures, UPS where capabilities permit.
- **Security reporting:** SSH posture, fail2ban, update exposure, account and privilege changes, firewall state, critical-file change detection.
- **Notifications:** Telegram, Discord, Slack, Teams, webhooks, SMS, and mobile push.
- **Experience:** richer history, graphs, weekly reports, installation profiles, localization, and accessibility refinement.
- **Lifecycle and integrations:** signed distribution, additional platforms, import/export, external tooling, and carefully governed remediation.
- **Optional extensions:** MOTD/status display, shell conveniences, and harmless entertainment packages as explicit modular recipes.

## 33. Long-Term Direction

Long term, QWSG may provide secure multi-server visibility and policy coordination while each Agent remains locally useful and resilient. Professional deployment may add fleet inventory, central incident views, delegated administration, integration APIs, stronger compliance evidence, and scalable history. Any hosted or cloud role must remain optional, privacy-conscious, and explicitly approved.

QWSG may be suitable for public or open-source distribution, but licensing, editions, commercial services, hosted offerings, and external branding are business decisions outside this blueprint.

## 34. Security and Privacy Expectations

QWSG must minimize privilege, attack surface, collected data, and network exposure. It must authenticate administrative access, authorize operations, protect secrets at rest and in transit as applicable, validate untrusted inputs, resist brute force, time out sessions appropriately, and log security-relevant changes. Supply-chain integrity and update authenticity require later security architecture.

Local-first operation is the preferred product direction. No telemetry or external transmission may be assumed. Any future remote service must disclose data categories, purpose, retention, and control, and requires explicit owner and operator decisions.

## 35. Reliability and Failure-Handling Expectations

One failed check, unavailable optional capability, notification outage, Console failure, or malformed module result must not silently disable unrelated protection. QWSG must surface stale observations, partial operation, delivery failures, storage pressure, clock concerns, and internal health failures. Restart behavior must preserve enough state to avoid false recoveries or alert floods.

Later specifications must define timeouts, retries, concurrency limits, backpressure, corruption handling, clock semantics, and recovery objectives.

## 36. Compatibility and Migration Expectations

Configuration, state, module compatibility, and lifecycle data need explicit versioning before implementation. Updates must detect incompatible changes and either migrate safely or stop with guidance. Downgrade expectations, schema evolution, module compatibility ranges, supported upgrade paths, and export formats are deferred to architecture and release policy work.

## 37. Documentation and User-Experience Expectations

The Installer, Console, messages, reports, diagnostics, and end-user documentation must be designed for localization. Hungarian-friendly usability may be part of product identity, while engineering artifacts remain English. User-visible states and actions must use consistent terminology and explain consequences before confirmation.

Documentation must cover deployment prerequisites, permissions, configuration, secrets, checks, state semantics, alerts, reports, lifecycle operations, troubleshooting, privacy, data retention, extension safety, and support boundaries. Examples must be labeled so they do not become accidental commitments.

## 38. Open Questions

1. Which exact Ubuntu and Debian versions and CPU architectures form the first support matrix?
2. Is the Console required in the first released package, or should MVP delivery sequence Agent and Installer before the Console while preserving the secure Console requirement?
3. Which local e-mail methods are mandatory, and is direct SMTP required for the first release?
4. What are acceptable default retention periods and storage budgets for state, history, logs, reports, and audit records?
5. What authentication and recovery model should the standalone Console use initially?
6. Which configuration and secrets-management approaches satisfy both small-server usability and privilege separation?
7. Which remediation, if any, is safe enough for an early opt-in capability after the reporting-only baseline?
8. What licensing, distribution, naming, telemetry, commercial, and hosted-service positions does the Project Owner approve?

## 39. Deferred Decisions

The following must not be inferred from examples in the source plan: implementation language; web or TUI framework; process and service topology; module interfaces; scheduling technology; IPC, CLI, REST, gRPC, WebSocket, or remote protocols; directory and package layout; database schema; configuration syntax and precedence; secrets backend; authentication implementation; cryptographic and update-signing design; exact thresholds; data retention; deployment ports and reverse-proxy setup; central-console topology; licensing; editions; pricing; hosted services; and implementation roadmap estimates.

## 40. Recommended Follow-Up Architecture Documents

Recommended sequence:

1. **Task 006 — Functional Specification**: define observable product behavior, actors, use cases, state semantics, MVP acceptance, and failure behavior without choosing internal implementation.
2. **Core System Architecture**: allocate responsibilities and trust boundaries.
3. **Security Architecture**: threat model, privilege separation, identity, secrets, updates, and audit integrity.
4. **Data and State Architecture**: observation, state, history, retention, migration, and recovery.
5. **Module Architecture**: capability contracts, lifecycle, isolation, and compatibility.
6. **Installer Architecture**: planning, consent, privilege, dependency, update, and removal workflows.
7. **Console Architecture**: administration boundary, deployment model, authentication, and localization.
8. **Implementation Roadmap**: sequence verified vertical slices after specifications and architectures are approved.

Proposed Task 006 slug: `functional-specification`. Proposed objective: create one authoritative QWSG Functional Specification that turns this blueprint's MVP and product behaviors into testable, implementation-neutral requirements, including actors, operating profiles, state transitions, alerts, lifecycle behavior, failure cases, and acceptance criteria. This is next because architecture should allocate responsibilities only after required behavior is precise.

## 41. Glossary

- **Agent:** standalone operational guardian installed on a monitored server.
- **Installer:** privileged lifecycle assistant for transparent bootstrap and maintenance.
- **Console:** optional independent administration and visualization surface.
- **Check:** one evidence-gathering evaluation of a target condition.
- **Module:** bounded extension contributing checks, policy, notification, integration, reporting, or lifecycle recipes.
- **Observation:** evidence produced by a check at a point in time.
- **State:** retained interpretation of observations for an identified subject.
- **Transition:** meaningful change between retained states.
- **Recovery:** transition from an unhealthy state toward or into healthy operation.
- **Capability:** functionality available only when environmental and dependency conditions are satisfied.
- **Remediation:** product-initiated corrective action under explicit bounded authorization.

## 42. Source Coverage and Traceability

| Source or planning domain | Blueprint coverage |
| --- | --- |
| Project Philosophy and Constitution | Sections 5, 7, 12–14, 34–37 |
| Product Definition | Sections 3–14, 33–39 |
| Original `QWSG_MASTER_PLAN.md` origin and alert examples | Sections 3–4, 22–24 |
| Independence and reusable server deployment | Sections 6, 9, 12, 14, 29–30 |
| Hybrid bootstrap, TUI, and web setup ideas | Sections 18–19, 28, 30–31; technologies deferred |
| Detection and dependency planning | Sections 15–16, 18, 20, 28–31 |
| Disk, memory, load, service, HTTP, SSL, and backup checks | Sections 15, 21, 31 |
| Network, database, mail, logs, hardware, and security checks | Sections 21 and 32 |
| State severities, escalation, recovery, and reminders | Sections 22–23 |
| Notification channels and periodic reports | Sections 23–24, 31–32 |
| Standalone web UI, history, and graphs | Sections 19, 24, 26, 32–33 |
| Diagnostics and support bundle | Section 27 |
| Installation profiles and operating modes | Sections 28 and 30 |
| Convenience and entertainment extensions, including `sl` | Sections 21 and 32 |
| Suggested directories, commands, systemd units, SQLite, and frameworks | Recorded only as illustrative source material; final choices deferred in Sections 26 and 39 |
| Three-component Agent/Installer/Console direction | Sections 16–20 |
| First usable release list | Section 31, separated from post-MVP in Section 32 |

No original planning category was discarded. Host-specific examples, short effort estimates, named tools, paths, ports, and structures are not treated as product commitments. They remain available in the preserved master plan for later evidence-based design work.
