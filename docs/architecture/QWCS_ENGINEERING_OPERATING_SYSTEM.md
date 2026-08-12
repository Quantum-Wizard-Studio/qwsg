# QWCS Engineering Operating System Architecture

## 1. Status, purpose, and authority

This document is the permanent architectural reference for the future Quantum
Wizard Creator System (QWCS). QWCS is a reusable AI Engineering Operating
System capable of supporting multiple independent software projects through a
shared, deterministic governance architecture.

This architecture preserves the accepted conclusions of the Engineering
Framework Review and defines the intended successor to Reusable Engineering
Framework 1.x. It is an architecture specification, not an implementation,
migration authorization, release claim, or modification of current engineering
rules. Framework 1.x remains authoritative for every adopting project until a
separately approved and verified migration activates a later QWCS version.

Normative architectural requirements use **MUST**, **MUST NOT**, **SHOULD**, and
**MAY**. Named implementation mechanisms are proposals unless this document
explicitly identifies their observable behavior as mandatory.

## 2. Long-term mission

QWCS exists to enable successful, safe, continuous, and accountable product
delivery by humans and AI engineering agents. It coordinates authority, intent,
risk, execution, evidence, validation, exceptions, lifecycle, and compatibility
without becoming a competing source of product intent or an unnecessary
obstacle to engineering progress.

QWCS MUST be reusable across independent repositories, products, technology
stacks, ownership models, languages, and release practices. Project policy MAY
specialize QWCS, but a project MUST NOT silently weaken constitutional truth,
authority, security, or evidence guarantees.

QWCS is not an autonomous project owner. It does not invent authority, approve
its own scope, conceal uncertainty, reinterpret failed work as success, or make
product decisions that belong to accountable humans.

## 3. Architectural hierarchy

QWCS separates durable purpose from replaceable implementation:

```text
Framework Philosophy
        ↓
Engineering Principles
        ↓
Constitution
        ↓
Framework Architecture
        ↓
Framework Components and Contracts
        ↓
Project Policy Profiles
        ↓
Task Manifests and Runtime Evidence
        ↓
Implementation Roadmap
```

The hierarchy has the following meanings:

- **Framework Philosophy** defines why QWCS exists.
- **Engineering Principles** define how competing goals are resolved.
- **Constitution** defines non-bypassable legitimacy and truth invariants.
- **Framework Architecture** defines responsibilities, information flow, and
  permanent separation of concerns.
- **Components and Contracts** define stable behavior independent of language,
  process topology, storage technology, or user interface.
- **Project Policy Profiles** select project-specific rules and authority.
- **Task Manifests and Evidence** represent individual governed work.
- **Implementation Roadmap** sequences delivery without changing architecture.

## 4. Framework philosophy

The primary engineering objective is successful product delivery. Governance
exists to make delivery safer, more truthful, more repeatable, and easier to
continue. A rule that cannot explain how it protects authority, safety,
correctness, compatibility, or maintainability does not qualify as a mandatory
engineering constraint.

QWCS MUST distinguish an unsafe action from an inconvenient procedure. It MUST
stop unauthorized, deceptive, destructive, or unacceptably risky work. When
authorized work can continue safely, it MUST provide a deterministic path
forward.

Governance strength is measured by reliable outcomes, not by the number of
interruptions, documents, approvals, or validations it produces.

## 5. Mandatory engineering principles

### 5.1 Human authority

Accountable human authority defines project intent, scope, acceptable risk,
delegation, and constitutional change. AI agents and framework components MUST
operate only within that authority.

### 5.2 Truthful evidence

QWCS MUST preserve the distinction between planned, approved, active, failed,
limited, complete, accepted, released, and abandoned work. Missing evidence
MUST remain missing. An override MUST NOT rewrite facts or manufacture success.

### 5.3 Engineering progress

When a rule fails, QWCS MUST identify the affected phase, protected interest,
smallest remediation, permissible remaining work, and override availability.
It MUST NOT block unrelated safe work merely because a later-phase gate is not
yet satisfied.

### 5.4 Safe forward progress

QWCS MUST prefer a safe continuation, isolation, compensating control,
disclosed limitation, or bounded override over an indefinite procedural stop.
Authority, constitutional truth, secrets, and immediate destructive safety are
not relaxed by this principle.

### 5.5 Minimal intervention

Every enforcement action MUST affect the smallest lifecycle phase and scope
needed to control the identified risk. A publication failure blocks
publication, not investigation. Missing completion evidence blocks completion,
not already authorized reversible implementation.

### 5.6 Proportionality

Validation, snapshot, review, and approval cost MUST be proportional to the
task's actual risk, reversibility, affected systems, and evidence needs.

### 5.7 Determinism

Identical canonical intent, policy, rule versions, project state, explicit
context, and evidence MUST produce identical governance decisions. Hidden time,
environment, identity, or network assumptions MUST NOT decide outcomes.

### 5.8 Reversibility and containment

Material changes MUST have recovery or containment evidence proportionate to
risk. Irreversible actions require stronger authority and validation than
reversible actions.

### 5.9 Separation of concerns

Project intent, governance policy, task compilation, execution, validation,
evidence, overrides, lifecycle transitions, and presentation MUST remain
separate responsibilities.

### 5.10 Compatibility

Historical evidence remains immutable. Migration MUST preserve the meaning and
auditability of Framework 1.x tasks without rewriting them into QWCS 2.x
claims.

## 6. Constitution boundary

The future QWCS Constitution contains only non-bypassable invariants:

- accountable human authority controls scope and risk acceptance;
- work MUST remain within granted authority;
- evidence and completion claims MUST be truthful;
- secrets and protected data MUST NOT be exposed;
- destructive or irreversible external actions require explicit authority;
- history MUST NOT be rewritten to invent legitimacy or success;
- task content and configuration MUST be treated as data, not executable
  authority;
- QWCS MUST NOT approve its own scope or constitutional exceptions.

These rules are not task-overridable. They may change only through a
prospective, versioned constitutional amendment by the designated human
constitutional authority. An amendment MUST NOT retroactively legitimize an
earlier violation.

Safety, workflow, quality, convenience, and recommendation rules do not belong
in the Constitution merely because they are important.

## 7. Rule architecture

### 7.1 Rule hierarchy

The Rule Registry MUST classify rules as:

| Category | Purpose | Default override model |
|---|---|---|
| Constitutional | Legitimacy, authority, and truth | No task-level override |
| Safety | Security, data protection, recoverability, destructive boundaries | Explicit authority plus compensating controls where allowed |
| Workflow | Predictable coordination and lifecycle shape | Project Owner or delegated workflow authority |
| Quality | Correctness, compatibility, maintainability, and support evidence | Technical authority; disclose limitations and follow-up |
| Convenience | Defaults and compatibility preferences | Project configuration or maintainer decision |
| Recommendation | Non-binding engineering guidance | Engineer judgment |

Rules MUST also carry priority. The canonical priority model is:

- `P0`: constitutional integrity;
- `P1`: immediate safety or security;
- `P2`: lifecycle and evidence integrity;
- `P3`: required quality;
- `P4`: workflow consistency;
- `P5`: convenience or project default;
- `P6`: recommendation.

Higher categories and priorities prevail in conflicts. Priority does not permit
a lower rule to weaken a higher rule.

### 7.2 Rule metadata

Every enforceable rule MUST have versioned machine-readable metadata:

- stable rule ID and title;
- category and priority;
- rationale and protected interest;
- applicable projects, task types, risks, targets, and lifecycle phases;
- deterministic condition and expected evidence;
- enforcement result and affected transition;
- override eligibility and required authority;
- required compensating controls;
- expiry behavior;
- dependencies, conflicts, compatibility, and remediation guidance;
- rule and framework versions.

The registry is canonical governance data. Human documentation explains it but
MUST NOT define conflicting hidden behavior.

## 8. Framework components

### 8.1 Engineering Principles

This component owns the mandatory decision principles in Section 5. It is the
highest interpretive layer beneath human authority and precedes the
Constitution. It does not contain project-specific procedures.

### 8.2 Constitution

This component owns only P0 invariants. It defines legitimacy boundaries and
the constitutional amendment process. It does not define filenames, command
syntax, branch names, validation tools, task templates, or convenience policy.

### 8.3 Rule Registry

The Rule Registry owns versioned rules, hierarchy, metadata, compatibility,
override contracts, and project-selectable profiles. It MUST support inspection
of why a rule applies and what it protects.

### 8.4 Builder / Task Compiler

The Builder compiles owner intent into a canonical Task Manifest. It MUST:

1. normalize structure without inventing substance;
2. validate authority, objective, scope, exclusions, and acceptance intent;
3. resolve applicable rule versions and project policy;
4. classify task risk;
5. derive phase-specific validation, evidence, snapshot, and rollback plans;
6. detect contradictions across task sections and installed lifecycle state;
7. render the complete review representation before approval;
8. bind approval to the exact manifest hash;
9. install lifecycle state transactionally;
10. validate that the installed task is executable in its post-install state;
11. emit a Validation Lock.

An `EXECUTABLE` manifest creates a presumption that implementation may start
without another procedural approval. Runtime interruption is justified only by
changed reality, failed authority or safety rules, scope expansion, destructive
external impact, or an override requirement.

The specific manifest encoding, Builder language, storage layout, and approval
interface are implementation decisions. Markdown MUST remain available as a
human-readable projection during Framework 1.x migration.

### 8.5 Governance Engine

The Governance Engine resolves rules against a Task Manifest, project policy,
explicit execution context, active overrides, and evidence. It owns governance
decisions but never performs product engineering actions.

Its canonical results are:

- `PASS`;
- `PASS_WITH_ADVISORY`;
- `REQUIRES_REMEDIATION`;
- `REQUIRES_OVERRIDE`;
- `BLOCKED_BY_SAFETY`;
- `BLOCKED_BY_AUTHORITY`.

Every result MUST identify rule IDs, phase, evidence, rationale, affected
transition, remediation, and override availability.

### 8.6 Validation Engine

The Validation Engine executes deterministic validation plans and produces
canonical Validation Records. Validation is phase-aware:

- intent validation;
- approval validation;
- executable-state validation;
- pre-mutation validation;
- implementation validation;
- completion validation;
- publication validation.

A validation result MAY be reused while its inputs, tool identity, policy,
environment dependencies, and evidence hashes remain unchanged. QWCS SHOULD
avoid repeated validation that cannot change the governance decision.

Validation implementation MUST isolate task content from command execution,
preserve exact argv, bound resources, record tool versions, distinguish
environment failure from product failure, and explain remediation.

### 8.7 Lifecycle Engine

The Lifecycle Engine owns transactional state transitions:

- create;
- review;
- approve;
- queue;
- start;
- amend;
- record evidence;
- complete;
- accept where separately required;
- archive;
- supersede;
- defer;
- divert or contain failed work.

It MUST preserve identity, history, collision safety, and rollback. Manual
status editing and manual file movement are compatibility behavior, not the
target architecture.

QWCS MAY support multiple queued or read-only tasks and disjoint concurrent
work. It MUST enforce one active writer for overlapping governed targets unless
a stronger isolation model proves safe composition.

### 8.8 Override Ledger

The Override Ledger is an append-only record of authorized exceptions. Each
record MUST identify:

- override ID;
- task and Validation Lock;
- rule ID and version;
- requested deviation and reason;
- authority and delegation basis;
- exact scope and lifecycle phases;
- compensating controls and evidence;
- accepted residual risk;
- effective time, expiry, and terminal event;
- outcome and supersession links.

Overrides MUST be specific, attributable, expiring where applicable, and
invalidated by material task or rule changes. They MUST NOT alter source
evidence, bypass constitutional rules, or conceal the original failure.

### 8.9 Evidence Store

The Evidence Store retains immutable or content-addressed Task Manifests,
Validation Locks, Validation Records, Override Records, lifecycle events,
snapshots, rollback evidence, decisions, attempts, results, and delivery
records.

It MUST preserve provenance, integrity, privacy classification, retention,
exportability, and project isolation. Evidence payloads and publication-safe
metadata MAY use different storage backends. QWCS MUST NOT require a cloud
service for local project governance.

### 8.10 Compatibility Layer

The Compatibility Layer reads and validates Framework 1.x configuration,
prompts, histories, Builder inputs, exact approval tokens, task numbering, and
lifecycle states. It projects them into QWCS contracts without rewriting
historical source records.

Compatibility adapters MUST distinguish an original Framework 1.x fact from a
QWCS-derived projection. Unsupported or ambiguous legacy states fail with
guidance rather than guessed semantics.

## 9. Validation Lock

The Validation Lock binds approval and executability to:

- canonical Task Manifest hash;
- authority and approval evidence;
- scope and exclusions;
- applicable rules and versions;
- project policy version;
- risk classification;
- required validation and evidence plan;
- active overrides;
- Builder and schema versions;
- explicit execution assumptions.

It does not freeze implementation choices or normal working-tree evolution
within scope. A substantive amendment invalidates the lock and requires
authorized recompilation. A mechanical or narrowly authorized correction
creates a linked superseding lock; the original remains evidence.

## 10. Risk and intervention model

The architecture requires proportional governance. A recommended baseline is:

| Risk | Typical work | Minimum recovery expectation |
|---|---|---|
| `R0` | Read-only review | No snapshot |
| `R1` | Reversible tracked documentation or configuration | Verified Git or exact-file recovery |
| `R2` | Normal source and test change | Bounded target snapshot or equivalent recoverability |
| `R3` | Migration, security, release, shared infrastructure | Verified full affected-scope snapshot and stronger approval |
| `R4` | Destructive or production-critical operation | Dedicated authority, rehearsal where possible, complete recovery and stop controls |

Projects MAY refine this model. They MUST preserve proportionality, authority,
truth, and minimum constitutional or safety constraints.

## 11. Deterministic engineering requirements

QWCS contracts MUST define:

- exact schemas and versions;
- canonical identity rules;
- explicit temporal and environmental context;
- deterministic ordering and conflict resolution;
- unsupported-version behavior;
- bounded inputs and outputs;
- stable serialization where interchange requires it;
- complete provenance;
- explicit unknown, unavailable, invalid, overridden, and failed states;
- separation of validation outcome from execution outcome;
- separation of technical completion from human or business acceptance.

Equal-priority governance conflicts MUST be represented explicitly. Input order,
filesystem enumeration, locale, presentation, or agent preference MUST NOT
silently resolve them.

## 12. Multi-project operating model

Each project owns an isolated Project Policy Profile containing identity,
repositories, authority roles, lifecycle storage, risk defaults, required rule
sets, validation adapters, languages, release policy, and compatibility mode.

The reusable QWCS core MUST NOT embed QWSG names, paths, remote transport,
branch names, task-width conventions, language policy, technology-stack tests,
or release assumptions. Project configuration cannot weaken P0 rules and may
override P1-P6 rules only through their declared contracts.

Cross-project operation MUST preserve data, authority, credential, evidence,
and lifecycle isolation. Failure in one project MUST NOT mutate or block an
unrelated project unless an explicit shared dependency exists.

## 13. Compatibility and migration strategy

Migration from Framework 1.x MUST be incremental and reversible:

1. inventory existing rules and assign stable QWCS IDs without changing
   behavior;
2. classify each rule and record its current enforcement point;
3. establish golden Framework 1.x fixtures and historical projections;
4. add structured validation results behind existing commands;
5. separate validation by lifecycle phase;
6. introduce Task Manifests and Validation Locks alongside Markdown records;
7. introduce Override and amendment transactions;
8. add proportional risk profiles and evidence reuse;
9. automate lifecycle transitions;
10. run Framework 1.x and QWCS validators in parallel;
11. require equivalent or stronger P0/P1 protection;
12. migrate projects individually through explicit owner acceptance;
13. retain legacy readers for the supported retention horizon.

Migration MUST NOT rewrite historical tasks, reinterpret incomplete work as
complete, require a mandatory hosted service, or force simultaneous migration
of independent projects.

## 14. Architecture versus implementation

The following are mandatory architectural outcomes:

- the hierarchy and component responsibilities in this document;
- constitutional non-bypassability;
- classified, versioned, explainable rules;
- deterministic phase-aware governance;
- approval bound to exact canonical task intent;
- safe forward progress and minimal intervention;
- explicit overrides with compensating controls;
- immutable evidence and lifecycle truth;
- Framework 1.x compatibility and reversible migration;
- multi-project isolation and reusable core boundaries.

The following remain implementation proposals requiring separate engineering
decisions:

- programming language and runtime;
- process, daemon, service, or library topology;
- JSON, YAML, database, or other canonical storage encoding;
- local or distributed Evidence Store backend;
- CLI, Terminal UI, Web, API, or agent protocol;
- authentication and signature mechanism;
- exact directory layout and executable names;
- cache technology and validation scheduler;
- packaging, installation, update, and deployment model;
- concurrency implementation;
- hosted or managed QWCS services.

No implementation may collapse the architectural boundaries merely for
convenience.

## 15. Implementation roadmap

### Phase 0 — Preserve and inventory

Assign stable IDs and metadata to existing Framework 1.x rules, document actual
enforcement, and establish compatibility fixtures without behavioral change.

### Phase 1 — Principles and Rule Registry

Create the Engineering Principles and reduced Constitution, introduce the
versioned Rule Registry, and represent existing behavior faithfully.

### Phase 2 — Structured governance results

Add canonical Governance and Validation Records while preserving existing
commands and exit behavior through adapters.

### Phase 3 — Phase-aware validation

Separate intent, approval, executable, mutation, completion, and publication
gates. Add deterministic evidence reuse and remediation reporting.

### Phase 4 — Builder / Task Compiler 2.0

Compile semantic Task Manifests, render before approval, bind approval to the
manifest hash, validate installed executability, and emit Validation Locks.

### Phase 5 — Overrides and amendments

Implement the Override Ledger, compensating-control validation, expiry, and
linked task amendments without rewriting prior evidence.

### Phase 6 — Proportional risk and recovery

Introduce risk-class-derived validation, snapshot, rollback, review, and
authority requirements.

### Phase 7 — Lifecycle Engine

Implement transactional lifecycle commands and remove manual status/file
mutation from the target workflow while preserving compatibility adapters.

### Phase 8 — Controlled concurrency and multi-project operation

Support queued, read-only, and safely isolated concurrent tasks; validate
project isolation and shared dependency behavior.

### Phase 9 — Framework 2.x adoption

Run dual validation, compare decisions and interruption rates, migrate projects
individually, retain legacy readers, and retire Framework 1.x enforcement only
after explicit acceptance and verified rollback readiness.

## 16. Success criteria for the architecture

Future QWCS work conforms to this architecture when:

- every mandatory rule has a rationale, category, priority, phase, evidence,
  and override contract;
- constitutional and immediate safety failures remain impossible to disguise;
- validated authorized tasks have a deterministic path to execution;
- irrelevant later-phase failures do not stop safe earlier-phase work;
- overrides are explicit, bounded, attributable, and auditable;
- governance decisions are reproducible and explainable;
- independent projects remain isolated;
- Framework 1.x projects can migrate incrementally without historical rewrite;
- QWCS demonstrably improves safe delivery rather than maximizing procedure.
