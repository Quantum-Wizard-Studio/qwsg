# Current Engineering Task 025: Canonical Policy Engine

## Task Metadata

- Task ID: `025`
- Task slug: `canonical-policy-engine`
- Status: `complete`
- Date opened: `2026-08-06` UTC
- Human authority: Project Owner
- Owner or lead-developer communication language: English

## Title

Canonical Policy Engine


## Objective

Implement the permanent Canonical Policy Engine of Quantum Wizard Server Guardian.

The Policy Engine establishes the missing governance layer between Rule Evaluation and operational interpretation.

It shall evaluate immutable Rule Evaluation results through deterministic, versioned Policy Profiles without modifying Rule logic or technical evidence.

The resulting Policy Evaluation shall become the single canonical governance source for all future Scheduler, Alert Engine, Dashboard, Terminal UI, REST API, remote management and automation capabilities.



## Scope

Task 025 shall design and implement the permanent Canonical Policy Engine of Quantum Wizard Server Guardian.

The task scope includes:

- defining the versioned Canonical Policy Model 1.0;
- defining Policy Profile 1.0;
- defining Policy Evaluation Record 1.0;
- defining canonical policy identity, provenance, scope, applicability and version semantics;
- consuming immutable canonical Rule Evaluation Records as policy input;
- producing deterministic, immutable and presentation-independent Policy Evaluation Records;
- defining canonical policy outcomes and their exact semantics;
- implementing deterministic policy-profile selection;
- implementing deterministic policy precedence;
- implementing explicit conflict detection and conflict resolution;
- defining profile inheritance or composition only where it can remain deterministic, bounded and explainable;
- preserving complete traceability from every policy result to its source Rule Evaluation Records and applicable Policy Profile;
- separating policy decisions from technical Rule, Health, Drift and Inventory evidence;
- integrating Policy into the existing canonical pipeline and orchestration layer without creating a second execution path;
- extending the existing command and report architecture only as required to consume canonical Policy Evaluation results;
- defining versioning, compatibility, unsupported-version and extension behavior;
- defining validation for malformed, incomplete, contradictory, unsupported or inapplicable policy definitions;
- defining safe behavior for missing evidence and indeterminate evaluation;
- preserving deterministic ordering, stable identifiers and byte-stable canonical serialization where applicable;
- preserving privacy, redaction, bounded-resource, offline and AI-independent guarantees;
- adding focused unit, contract, integration, determinism, conflict, compatibility and regression tests;
- creating permanent Canonical Policy Engine architecture documentation;
- updating only the directly affected architecture, system map, roadmap, engineering history, README and lifecycle records.

The Policy Engine shall remain a pure governance-evaluation component.

It may determine how canonical Rule Evaluation results are classified or treated under an applicable Policy Profile, but it must not:

- collect system data;
- perform Inventory, Compare, Drift, Health or Rule evaluation;
- modify source evidence;
- execute commands;
- schedule work;
- send notifications;
- trigger remediation;
- mutate the host;
- make network calls;
- introduce hidden side effects.

Task 025 must preserve all existing public contracts and CLI behavior unless a narrowly required, explicitly versioned and fully tested compatible extension is necessary for Policy integration.



## Out of Scope

Do NOT implement:

- Scheduler
- Alert Engine
- Monitoring
- Daemon mode
- Notifications
- Email
- REST API
- Dashboard
- Terminal UI
- AI
- Machine Learning
- Remediation
- Remote execution
- Configuration UI
- Host mutation

Policy evaluates only.

It performs no operational action.



## Required Reading

- `ai/core/00_PROJECT_PHILOSOPHY.md`
- `ai/core/01_CONSTITUTION.md`
- `ai/core/03_AGENTS.md`
- `ai/core/08_JOB_TEMPLATE.md`
- `ai/core/11_ENGINEERING_LIFECYCLE.md`
- `ai/core/14_PROMPT_WORKFLOW.md`
- `ai/core/16_GIT_POLICY.md`
- `ai/config/engineering-project.conf`

## Starting State Verification

Before making any repository modification:

- verify canonical lifecycle;
- verify Task 025 is the sole active approved task;
- verify Task 024 is the latest completed task and the canonical implementation baseline;
- verify repository validation passes;
- verify Git working tree is suitable for Builder installation;
- report every failed precondition and stop if validation fails.



## Snapshot Requirements

Before modifying any Task 025 implementation or documentation target, create and verify the mandatory rollback-capable snapshot:

- create the mandatory pre-task snapshot;
- verify snapshot integrity;
- generate checksum list;
- ensure rollback artifacts are complete.



## Risk Assessment

Primary risks:

- introducing governance responsibilities into Rule Engine;
- mixing Policy with Scheduler or Alert behaviour;
- creating non-deterministic Policy evaluation;
- changing existing canonical pipeline behaviour.

Risk mitigation:

- Policy must remain read-only;
- Rule Engine responsibilities must remain unchanged;
- Policy must not execute actions;
- deterministic output must be preserved.



## Planned Work

Task 025 shall:

- define Canonical Policy Model 1.0;
- define Policy Profile 1.0;
- define Policy Evaluation Record 1.0;
- implement deterministic Policy Engine;
- implement deterministic Policy precedence;
- implement deterministic conflict resolution;
- define canonical Policy Outcomes;
- support versioned Policy Profiles;
- define Policy serialization;
- integrate Policy into the existing canonical pipeline;
- create permanent architecture documentation;
- implement engineering tests;
- update canonical documentation where required.



## Rollback Plan

If any validation fails:

- restore the repository from the pre-task snapshot;
- restore lifecycle state;
- verify snapshot integrity;
- confirm repository consistency before exiting.



## Deliverables

- Canonical Policy Engine
- Canonical Policy Model 1.0
- Policy Profile 1.0
- Policy Evaluation Record 1.0
- deterministic Policy taxonomy
- Policy precedence model
- conflict resolution model
- Policy serialization
- architecture documentation
- engineering tests
- updated canonical documentation



## Verification

Successful completion requires:

- deterministic Policy Evaluation;
- deterministic Policy serialization;
- stable Policy identity;
- complete engineering tests;
- Go tests PASS;
- race tests PASS;
- go vet PASS;
- gofmt PASS;
- framework validation PASS;
- lifecycle validation PASS;
- Builder validation PASS;
- repository validation PASS.



## Documentation Updates

Create:

docs/architecture/CANONICAL_POLICY_ENGINE.md

Update only the canonical documentation strictly required by Task 025, including architecture references, engineering history, roadmap, system map and other directly affected permanent engineering documents.



## Completion Criteria

Task 025 is complete only when:

- Policy becomes the single governance layer between Rule and Report;
- Rule Engine responsibilities remain unchanged;
- Scheduler and Alert Engine remain unimplemented;
- future Scheduler, Alert Engine, Dashboard and Terminal UI can consume Policy Evaluation directly;
- all engineering validations pass;
- repository returns to canonical idle lifecycle state after task completion.



## Owner Approval Requirements

Approved by Project Owner through the Engineering Task Builder on 2026-08-06 UTC.

The structured task definition has been explicitly approved for implementation. Further scope changes require explicit Project Owner approval.
