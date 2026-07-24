# Quantum Wizard Server Guardian Product Definition

## Document purpose

This document is the single product-level definition of Quantum Wizard Server Guardian (QWSG). It defines what the product is, why it exists, whom it serves, which problems it addresses, which boundaries constrain it, and which principles must guide its evolution.

It intentionally does not define implementation, internal architecture, modules, APIs, databases, programming languages, frameworks, installers, deployment topology, or service configuration. Future architecture must conform to this Product Definition after its strategic proposals receive owner decisions.

## Status and authority

- Status: **owner-review draft**
- Created: `2026-07-18` UTC under Task 003
- Engineering language: English
- Product authority: the human project owner
- Current effect: statements marked **Established** restate already approved QWSG governance; statements marked **Proposed — owner approval required** are recommendations and are not approved business commitments.

This document becomes fully authoritative when the owner explicitly approves or resolves the proposed strategic decisions. Until then, it is authoritative only as a consolidated statement of the established constraints it cites.

Task 019 approved the edition philosophy, offline and privacy principles, and
long-term ecosystem boundaries in `docs/PRODUCT_ARCHITECTURE.md`. That canonical
architecture resolves the corresponding earlier proposals below; unrelated
open pricing, legal, distribution, and service commitments remain open.

## Decision labels

- **Established:** already mandated by approved project governance and may guide future work.
- **Proposed — owner approval required:** recommended product direction that must not be treated as an approved fact.
- **Open decision:** a choice intentionally left to the owner; engineering work must not silently resolve it.

## Product identity

**Established.** Quantum Wizard Server Guardian is an independent, modular Linux server protection product. QWSG exists to protect servers, not merely observe them. It must be installable on independent Linux servers, must not depend on QUWIP or another Quantum Wizard product, and must provide its own independent QWSG console.

QWSG behaves as a careful and experienced system administrator would: it observes, verifies, analyzes, explains, and warns. It performs automatic corrective action only when explicitly authorized. It never substitutes assumption for available verification and never hides relevant changes.

## Product purpose

**Established.** QWSG exists to help responsible server operators understand risk, detect meaningful problems, make safer decisions, and protect Linux systems without surrendering control or visibility.

The product's value is not the volume of collected data. Its value is trustworthy interpretation and controlled protection: turning verified server facts into understandable findings, warnings, and explicitly authorized corrective outcomes.

## Product philosophy

**Established.** QWSG follows these product-level principles:

1. Protection is the purpose; observation is a necessary means.
2. Stability has priority over convenience.
3. Security has priority over shortcuts.
4. Relevant changes are logged and understandable.
5. Destructive actions are reversible where technically possible.
6. Automatic correction requires explicit authorization.
7. Verification is preferred over assumption.
8. The product remains modular, independently installable, explainable, and maintainable over years.
9. QWSG remains independent of other Quantum Wizard products.
10. User-facing software is localization-ready from the beginning.

## Core values

### Trust through transparency

**Established.** The Guardian never hides changes. Findings, decisions, and relevant actions must be explainable to the operator.

### Operator authority

**Established.** The operator remains in control. Observation, analysis, explanation, and warning do not imply permission to change a server.

### Conservative safety

**Established.** Small, verified, reversible actions are preferred. QWSG must not trade server stability for apparent convenience.

### Security without shortcuts

**Established.** Secrets, credentials, privileges, and corrective authority must be treated as security boundaries rather than implementation conveniences.

### Independence and longevity

**Established.** A QWSG installation must not require another Quantum Wizard product, and the product must remain maintainable long after its initial creation.

### Respect for user language and data

**Established.** Engineering artifacts remain English, while user-facing content is designed for localization. Privacy-sensitive data is not exposed or collected merely because collection is possible.

## Target users

The following positioning is **Proposed — owner approval required**.

- Linux system administrators responsible for the safety and continuity of one or more servers.
- Hosting and infrastructure operators who need consistent, explainable server protection.
- Small organizations that operate Linux services without a large dedicated security or operations team.
- Experienced technical owners who want assistance without surrendering administrative authority.
- Managed-service and professional operations teams that need auditable, repeatable governance across servers.

QWSG is not positioned as a substitute for basic Linux competence. It is a guardian and decision-support product for people who remain accountable for their systems.

## User personas

The following personas are **Proposed — owner approval required** and describe needs, not implementation roles.

### Independent server administrator

Maintains a small number of Linux servers and needs trustworthy warnings, clear explanations, safe recommendations, and control over every corrective action.

### Small-team technical owner

Is accountable for business services but has limited operations capacity. Needs prioritization, understandable risk, and confidence that QWSG will not silently destabilize the server.

### Hosting or managed-service operator

Oversees multiple independently administered systems. Needs consistent policy, traceability, explainability, and clear separation between observation and authorized action.

### Security-conscious infrastructure lead

Needs evidence, accountability, privacy boundaries, and a product that integrates into responsible operational practice without concealing decisions.

## Problems QWSG solves

### Established problem domain

QWSG addresses the gap between raw server state and safe operator action. It observes and verifies Linux server conditions, analyzes their significance, explains findings, warns about relevant risks, and supports explicitly authorized correction.

### Proposed product outcomes — owner approval required

- Reduce the chance that important server risks remain unnoticed or misunderstood.
- Reduce cognitive load by prioritizing findings rather than presenting undifferentiated telemetry.
- Make corrective choices easier to understand before authorization.
- Preserve an auditable record of relevant observations, warnings, decisions, and changes.
- Improve consistency without removing human accountability.
- Help operators distinguish urgent risk from routine noise.

## Problems intentionally not solved

The following boundaries are **Proposed — owner approval required**, except where an established rule is noted.

- QWSG does not replace the human owner or administrator as the final authority. **Established.**
- QWSG does not silently remediate infrastructure. **Established.**
- QWSG does not promise that every destructive operation can be reversed when the underlying technology makes reversal impossible; it must disclose that limitation. **Established principle.**
- QWSG is not a dependency or extension of QUWIP or another Quantum Wizard product. **Established.**
- QWSG is not a general-purpose business management, billing, customer relationship, or content-management product.
- QWSG does not define or own the operator's entire security program, backup strategy, disaster-recovery plan, or organizational compliance program.
- QWSG does not guarantee absolute security, uninterrupted availability, or error-free third-party systems.
- QWSG does not treat high-volume monitoring data as useful protection without verified interpretation.

## Product goals

### Established goals

- Protect Linux servers through verification, analysis, explanation, warning, and controlled action.
- Preserve stability, security, transparency, modularity, independence, reversibility where possible, and long-term maintainability.
- Provide an independent QWSG console.
- Keep automatic correction behind explicit authorization.

### Proposed goals — owner approval required

- Be useful for a single independently operated server without requiring an external commercial service.
- Scale product value from individual administrators to professional multi-server operations without weakening local control.
- Make important findings understandable to both technical decision-makers and day-to-day operators.
- Establish a sustainable product whose commercial value comes from additional operational capability and support rather than hidden data use or unsafe pressure.

## Non-goals

- Selecting internal architecture, modules, protocols, storage, languages, frameworks, or deployment topology.
- Implementing or pre-authorizing corrective actions.
- Creating a cloud dependency for fundamental server protection.
- Using obscurity, silent modification, or unexplained scoring as a substitute for evidence.
- Defining pricing, license grants, service-level commitments, or a final edition feature matrix in this task.
- Beginning Product Architecture or implementation.

## Product boundaries

### Guardian responsibility

**Established.** The Guardian's product responsibility is to observe, verify, analyze, explain, warn, and—only when explicitly authorized—perform corrective action. Its responsibility ends where facts cannot be verified, authority has not been granted, or safe behavior cannot be explained.

### Operator responsibility

**Established.** The human owner retains final authority, approves strategic decisions, grants corrective authority, and remains accountable for server and organizational choices outside QWSG's verified scope.

### Environmental boundary

**Established.** QWSG must remain installable on independent Linux servers and must not require another Quantum Wizard product. Compatibility with particular distributions, control panels, hosting environments, or external services is not defined here.

### Product-definition boundary

This document constrains future architecture but does not design it. Terms such as Agent and Console describe user-visible product responsibilities, not process layout, network topology, API shape, storage design, or deployment decisions.

## Agent and Console relationship

The following product relationship is **Proposed — owner approval required**.

- The **Agent** represents QWSG's server-guardian responsibility: acquiring and evaluating verified server facts and carrying out only authorized product behavior.
- The **Console** represents QWSG's independent human control and explanation surface: presenting findings, evidence, warnings, policy choices, authorization boundaries, and relevant history.
- The Agent must not become a hidden authority independent of the operator.
- The Console must not conceal, reinterpret, or imply actions that the operator cannot understand.
- Their relationship must preserve traceability and explicit authorization.

This statement does not decide whether Agent and Console are separate processes, packages, services, repositories, or deployment units.

## Offline philosophy

**Established constraint.** QWSG must remain installable on independent Linux servers.

**Proposed — owner approval required.** Fundamental local protection should remain useful without a continuous connection to QWSG-operated cloud infrastructure. Loss of external connectivity should not silently disable essential local observation, explanation, warning, or already authorized local behavior. Any feature that cannot operate offline should be clearly identified before purchase or activation.

“Offline” does not mean that a protected server itself never uses a network. It means QWSG's fundamental value should not depend on a mandatory vendor-controlled connection.

## Cloud philosophy

**Open decision.** No QWSG cloud service is currently approved, specified, or implemented.

**Proposed — owner approval required.** Future cloud capability, if approved, should be optional, additive, transparent, and privacy-conscious. It may provide convenience, coordination, or professional services, but must not become a concealed prerequisite for fundamental independent-server protection. Operators must understand which data leaves a server, why, under whose control, and for how long.

## Privacy principles

The following principles are **Proposed — owner approval required**, consistent with established secret-handling and transparency rules:

1. Collect the minimum data necessary for an explicitly defined product purpose.
2. Prefer local processing when it can satisfy the purpose safely.
3. Never expose secrets or credentials in repositories, logs, reports, support data, or user interfaces.
4. Make external transmission, retention, and deletion behavior understandable.
5. Do not monetize private server data or telemetry as an undisclosed product model.
6. Separate operational evidence from unnecessary personal or customer content.
7. Give operators meaningful control over optional data sharing.

## Security principles

### Established

- Security has priority over shortcuts.
- Automatic corrective action requires explicit authorization.
- Relevant changes are logged and understandable.
- Destructive action is reversible where technically possible.
- Credentials and secrets are never written into the repository or exposed.
- Stability is a security property and must not be sacrificed for convenience.

### Proposed — owner approval required

- Default behavior should minimize privilege and corrective authority.
- Trust boundaries and limitations should be visible to operators.
- Security findings should distinguish evidence, interpretation, uncertainty, and recommended action.
- Commercial edition boundaries must never make the Community Edition unsafe,
  incomplete, or less correct.

## Product editions

**Approved product direction under Task 019.** QWSG uses Community,
Professional, and Provider edition identities. Every edition shares the same
deterministic engineering core. The canonical boundary and comparison are
defined in `docs/PRODUCT_ARCHITECTURE.md`. Pricing, distribution rights, public
license terms, and service commitments remain open. “Community” is a product
edition name and does not by itself grant redistribution or modification
permission.

### Community Edition

Community is a complete professional Linux engineering toolkit for unlimited
manual local operation. It provides the full deterministic engineering
capability, transparent history, reports and exports, local terminal
experiences, offline operation, privacy-first behavior, and human control
without requiring an account.

### Professional Edition

Professional is the automation and management platform around the same core. It
adds supported scheduling, background operation, notifications, managed
history, dashboards, multi-server coordination, updates, licensing, and support
as separately implemented. Its value is time saved, not better engineering
correctness.

### Provider Edition

Provider adds tenant-aware fleet, customer-portal, integration, white-label,
and provider-control-plane operations. Tenant isolation and audit are security
properties. Provider uses the same core and does not establish a private
correctness tier.

## Commercial philosophy

The following is **Proposed — owner approval required**.

- Commercial sustainability should fund long-term maintenance, security work, documentation, and support.
- Revenue should come from declared product capability and service value, not undisclosed exploitation of server data.
- Commercial messaging must not exaggerate certainty or promise absolute security.
- Operators should be able to understand edition limits before installation or purchase.
- Product independence should be preserved; bundling with another Quantum Wizard product must not become a dependency.
- The final licensing model remains an owner decision before public release.

## Community, Professional, and Provider positioning

**Approved product direction under Task 019.** Community earns trust through
complete local engineering. Professional saves time through supported
automation and management. Provider enables tenant-aware operations at scale.

The distinction is never “unsafe versus safe,” “incorrect versus correct,”
“opaque versus explainable,” or “operator-controlled versus vendor-controlled.”
Those principles define QWSG itself and are not premium-only.

## Long-term evolution

### Established constraints

QWSG must remain modular, independently installable, product-independent, transparent, and maintainable over years. Every relevant evolution must respect explicit authorization and documented change.

### Proposed direction — owner approval required

- Evolve from verified operator problems rather than feature volume.
- Preserve compatibility and migration paths where practical.
- Add product breadth only when responsibilities and boundaries remain understandable.
- Treat localization, privacy, security, and auditability as product foundations rather than later additions.
- Avoid architectural lock-in before evidence justifies it.
- Review this Product Definition when strategic scope changes, but preserve decision history.

## Guiding engineering principles

Future product and architecture work must:

1. Read and conform to the Project Philosophy and Constitution.
2. Trace decisions back to an established product requirement or an explicit owner decision.
3. Keep product requirements separate from implementation choices.
4. Verify environmental facts rather than embedding assumptions.
5. Prefer the smallest safe, reversible, understandable change.
6. Protect independent-server operation and QWSG product independence.
7. Preserve explicit authorization for corrective action.
8. Design user-visible content for localization and avoid hardcoded strings where technically feasible.
9. Document privacy, security, stability, and maintenance consequences.
10. Stop and request owner direction when a strategic choice is unresolved.

## Owner decision register

The following decisions remain open and must not be silently converted into engineering requirements:

| Decision | Current status | Engineering recommendation |
|---|---|---|
| Target-user and persona positioning | Owner approval required | Approve the four proposed personas as initial product audiences. |
| Offline product promise | Approved by Task 019 | Local deterministic engineering remains useful without Internet, account, cloud, telemetry, or remote licensing checks. |
| Cloud product role | Approved boundary; service details open | Future services are optional adapters and never prerequisites for the local core. |
| Edition names | Approved by Task 019 | Community, Professional, and Provider. |
| Edition feature boundary | Approved by Task 019 | Community is complete engineering; Professional adds automation; Provider adds operational scale. |
| Commercial model and pricing | Open | Decide only after product scope and cost assumptions are verified. |
| Final license and distribution rights | Open | Resolve before public release; the temporary proprietary notice remains controlling. |
| Privacy and telemetry commitments | Approved boundary; implementation details open | Privacy first, local by default, explicit minimized transfers, and opt-in telemetry. |
| Agent and Console product roles | Owner approval required | Approve the responsibility split without treating it as architecture. |

## Relationship to Product Architecture

The canonical Product Architecture is
`docs/PRODUCT_ARCHITECTURE.md`. It refines this product definition into approved
ecosystem, edition, experience, deployment, licensing, privacy, automation, AI,
and roadmap boundaries. It does not silently resolve the pricing, legal,
distribution, or service-level decisions that remain open here.

## Maintenance and change control

This document evolves through explicit product-governance tasks. Every change must record its reason, decision authority, snapshot, rollback, verification, and history. Revisions must preserve the distinction between established constraints, approved product decisions, proposals, and open questions.

Implementation details, live configuration, secrets, pricing tables, contractual terms, and transient task instructions do not belong here.
