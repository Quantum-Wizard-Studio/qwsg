# Current Engineering Task 021: Canonical Health Engine

## Task Metadata

- Task ID: `021`
- Task slug: `canonical-health-engine`
- Status: `complete`
- Date opened: `2026-07-24` UTC
- Human authority: Attila — Project Owner
- Owner or lead-developer communication language: Hungarian

## Title

Canonical Health Engine


## Objective

Design and implement the permanent, deterministic Health evaluation layer of QWSG.

Health evaluates the current engineering condition of the system from canonical engineering evidence produced by previous layers. It must not duplicate Compare or Drift responsibilities, and it must remain completely deterministic, reproducible, versioned, explainable, and independent of AI.

The Engineering Goal is to establish the canonical deterministic Health model upon which all future Rule, Policy, Report, and Automation engines will rely.


## Scope

Design, implement, test, and document:

- the Canonical Health Engine;
- Canonical Health Record 1.0;
- a canonical Health taxonomy;
- a deterministic Health evaluation pipeline;
- versioned public contracts;
- architecture documentation;
- engineering tests;
- the lifecycle and chronological engineering-documentation updates required by the Engineering Framework.

The Health Engine shall:

- evaluate the current engineering condition represented by canonical input evidence;
- consume canonical engineering evidence produced by previous QWSG layers, including canonical Drift output where applicable;
- preserve a strict Drift → Health boundary;
- produce stable Health classifications and records from equivalent canonical evidence;
- expose explicit, versioned, testable public contracts;
- make evaluation inputs, outputs, taxonomy terms, reasons, and compatibility behavior auditable;
- remain deterministic and reproducible across repeated evaluations of identical canonical evidence;
- avoid wall-clock, network, monitoring, scheduler, daemon, alerting, remediation, policy, compliance, and AI dependencies in the evaluation path.

The task includes only the smallest safe changes required to establish this permanent Health foundation and integrate it consistently with the existing Inventory, Snapshot, Compare, and Drift architecture.


## Out of Scope

The following are explicitly prohibited:

- monitoring;
- scheduler implementation;
- daemon implementation;
- Alert Engine implementation;
- alerts or alert delivery;
- email notifications;
- automatic remediation or remediation execution;
- Policy Engine implementation;
- Compliance Engine implementation;
- AI integration or AI-dependent evaluation;
- remote communication or network-dependent Health behavior.

Health must never perform monitoring, alerting, notification, remediation, policy enforcement, compliance enforcement, scheduling, daemon orchestration, or remote communication.

Health must not:

- rediscover or recollect system inventory;
- create snapshots;
- compare snapshots;
- reimplement comparison semantics;
- detect or classify canonical drift already owned by the Drift layer;
- silently absorb future Rule, Policy, Report, Alert, Automation, or Compliance responsibilities;
- begin implementation of any future Rule Engine, Policy Engine, Report Engine, Alert Engine, Automation Engine, or Compliance Engine.

No release, deployment, infrastructure mutation, dependency installation, remote Git operation, or unrelated refactor is authorized.


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

Before implementation:

1. Run `bin/job --check` and confirm Task 021 is the single active approved task with a matching history record.
2. Read every item in Required Reading and record the completed reading in the Task 021 history.
3. Run `ai/scripts/framework-check.sh` and all configured baseline lifecycle validations.
4. Verify the UTC date, effective user, repository root, configured project identity, current branch, HEAD, remotes, ahead/behind relationship where safely available, complete Git status, and relevant tags.
5. Confirm Task 020 is complete and its Canonical Drift Engine deliverables and contracts exist as the immediate architectural dependency.
6. Inspect the current Inventory → Snapshot → Compare → Drift contracts, packages, tests, architecture documents, ownership, permissions, and ACLs before proposing Health interfaces.
7. Confirm there is no existing Health implementation or overlapping untracked work that would be overwritten.
8. Record all existing unrelated tracked and untracked changes and leave them untouched.
9. Stop and report any material difference, lifecycle inconsistency, unresolved authority field, incompatible contract, unexpected Git divergence, permission problem, or target collision.


## Snapshot Requirements

Before modifying any Task 021 target:

- create a timestamped rollback-capable snapshot under `/tmp` using a Task 021-specific directory name;
- capture every existing file that may be modified, together with relevant Git state, permissions, ownership, ACL evidence where supported, and checksums;
- include the active Task 021 prompt and matching history record in the governance snapshot;
- record the exact snapshot path, captured targets, creation command, integrity verification, and retention limitation in the Task 021 history;
- verify that the snapshot is readable and sufficient for the bounded rollback procedure before implementation.

Do not place credentials, secrets, generated build caches, or unrelated private data in the snapshot. Do not modify task targets until snapshot verification passes.


## Risk Assessment

- Architectural-boundary risk — high: Health could duplicate Compare or Drift semantics. Mitigate with explicit responsibility tables, contract-level boundaries, dependency direction, and tests proving Health consumes rather than recreates canonical evidence.
- Determinism risk — high: time, iteration order, environment, or hidden mutable state could make Health output unstable. Mitigate with canonical ordering, explicit inputs, controlled serialization, golden/repeatability tests, and prohibition of wall-clock, network, AI, and ambient-state dependencies.
- Taxonomy ambiguity risk — high: unclear Health terms could become incompatible foundations for later engines. Mitigate with normative definitions, stable identifiers, precedence rules, reason semantics, examples, and versioned compatibility rules.
- Contract-compatibility risk — high: early public types may constrain Rule, Policy, Report, and Automation integration. Mitigate with Canonical Health Record 1.0, additive evolution rules, explicit version fields, validation, and compatibility tests.
- Evidence-integrity risk — medium: incomplete or unsupported upstream evidence could be mistaken for healthy state. Mitigate with explicit unknown, unavailable, invalid, or insufficient-evidence semantics and fail-closed evaluation rules where appropriate.
- Serialization and ordering risk — medium: maps or unstable ordering could produce non-reproducible records. Mitigate with deterministic ordering and byte-stability tests for canonical representations.
- Security and privacy risk — medium: Health reasons could echo sensitive upstream content. Mitigate by preserving canonical redaction boundaries and testing that Health does not introduce raw secret or private-host disclosure.
- Performance and stability risk — medium: evaluation could become unbounded. Mitigate with pure bounded processing of supplied canonical records and representative tests.
- Localization risk — low: engineering contracts remain English; any future user-facing rendering belongs to Report or presentation layers and is not implemented here.
- Permission and rollback risk — low: preserve existing modes, ownership, and ACL behavior; use exact target lists and the verified snapshot for recovery.
- Scope-expansion risk — high: future engine concepts could be prematurely implemented. Mitigate with explicit stubs-free documentation boundaries and excluded-work verification.


## Planned Work

1. Validate the active Task 021 lifecycle, framework configuration, repository identity, Git state, required reading, and Task 020 completion evidence.
2. Inspect the canonical Inventory, Snapshot, Compare, and Drift contracts and document the exact evidence boundary entering Health.
3. Create and verify the Task 021 rollback snapshot before changing targets.
4. Define the Health evaluation model and responsibility boundary: Health evaluates current engineering condition from canonical evidence; Compare determines differences; Drift produces canonical drift evidence; Health neither repeats nor owns those operations.
5. Define the canonical Health taxonomy with stable machine identifiers, normative meanings, precedence or aggregation rules, insufficient-evidence behavior, and deterministic reason semantics.
6. Define Canonical Health Record 1.0, including explicit schema/version identity, evaluated evidence references, overall condition, component or finding results where justified, deterministic reasons, and canonical ordering rules.
7. Implement the Canonical Health Engine as a pure deterministic evaluation pipeline over explicit canonical inputs, with no monitoring, scheduler, daemon, alerting, remediation, policy, compliance, AI, network, or remote communication dependency.
8. Publish versioned public contracts and validation behavior with compatibility rules for supported and unsupported versions.
9. Add engineering tests covering taxonomy, input validation, evaluation precedence, Drift → Health handoff, repeatability, ordering, serialization stability where applicable, unsupported evidence, invalid data, privacy boundaries, and strict exclusion of upstream/downstream responsibilities.
10. Write architecture documentation describing the Health Engine architecture, Health evaluation model, Health taxonomy, Drift → Health pipeline, future Rule Engine integration, future Policy Engine integration, future Report Engine integration, and compatibility strategy.
11. Update only the necessary canonical architecture, system map, roadmap, engineering history, user/developer references, and Task 021 history so terminology and dependency direction remain consistent.
12. Run all task-specific and repository-wide verification required by the framework, review exact diffs and permissions, and produce the English delivery report and Hungarian owner-facing handoff.

Do not begin future engines. Future integration sections must define boundaries and contracts only.


## Rollback Plan

Rollback is bounded to the exact Task 021 target paths recorded before implementation.

1. Stop Health implementation and preserve failure evidence.
2. Confirm the repository root, active Task ID, current Git state, exact target list, and verified Task 021 snapshot path.
3. Review the proposed restore set and obtain explicit confirmation before any destructive replacement.
4. Restore only pre-existing modified targets from their exact snapshot counterparts, preserving recorded permissions and ownership where authorized.
5. Remove only Task 021-created files whose exact paths are recorded in the history and confirmed not to contain later owner work. Never use wildcard deletion, broad `git reset`, `git checkout`, `git restore`, or `git clean`.
6. Restore the Task 021 prompt/history only if governance files were damaged; do not erase truthful execution or failure evidence merely to make the task appear unstarted.
7. Re-run framework, lifecycle, relevant package, documentation, permission, Git diff, and status checks.
8. Record the rollback commands, restored paths, verification results, retained evidence, and unresolved issues in the Task 021 history.

If exact restoration cannot be proven safe, stop and request Project Owner direction.


## Deliverables

- Canonical Health Engine implementation.
- Canonical Health Record 1.0 with explicit version identity and validation rules.
- Canonical Health taxonomy with stable identifiers and normative semantics.
- Deterministic Health evaluation pipeline consuming canonical engineering evidence.
- Versioned public Health contracts and compatibility behavior.
- Engineering tests proving correctness, determinism, reproducibility, boundary separation, and excluded behavior.
- Architecture documentation covering:
  - Health Engine architecture;
  - Health evaluation model;
  - Health taxonomy;
  - Drift → Health pipeline;
  - future Rule Engine integration;
  - future Policy Engine integration;
  - future Report Engine integration;
  - compatibility strategy.
- Necessary canonical architecture, system map, roadmap, developer/user documentation, and lifecycle updates.
- A complete Task 021 chronological history and delivery report with snapshot, decisions, exact changes, validation evidence, rollback, Git state, limitations, and result.


## Verification

Verification must include:

- `bin/job --check`;
- `ai/scripts/framework-check.sh`;
- `ai/scripts/next-task.sh --check`;
- `bin/job --check-test-tasks`;
- all configured Engineering Framework, lifecycle, diversion, and Task Builder assertions;
- targeted Health package tests;
- all repository Go package tests;
- race-enabled Go tests where supported;
- `go vet` and formatting checks;
- deterministic repeat evaluation of identical inputs;
- deterministic output ordering and canonical serialization checks where serialization is exposed;
- golden or equivalent contract tests for Canonical Health Record 1.0;
- tests for every Health taxonomy class and precedence/aggregation rule;
- tests for empty, incomplete, unknown, unsupported-version, malformed, and contradictory canonical evidence;
- tests proving the Drift → Health boundary and proving Health does not perform Compare or Drift work;
- tests proving no monitoring, scheduler, daemon, alerting, notification, remediation, Policy Engine, Compliance Engine, AI, network, or remote communication behavior was introduced;
- privacy and secret-disclosure checks;
- architecture and terminology consistency checks;
- backward-compatibility and version-negotiation checks;
- exact target, permission, ownership, and ACL review where supported;
- snapshot integrity and bounded rollback review;
- `git diff --check`, exact unstaged/staged path review, and `git diff --cached --check` if anything is staged;
- confirmation that unrelated tracked and untracked files remain untouched;
- confirmation that no prohibited dependency, infrastructure change, commit, push, release, deployment, or future-engine implementation occurred without separate authority.

Record exact commands, relevant output summaries, assertion counts, limitations, and failures truthfully in the Task 021 history. A failing mandatory check blocks completion.


## Documentation Updates

Create or update only the documentation needed to make the Health model canonical and traceable:

- a dedicated Health Engine architecture document;
- canonical architecture and system-map references establishing Inventory → Snapshot → Compare → Drift → Health dependency direction;
- documentation of the Health evaluation model;
- documentation of the Health taxonomy;
- documentation of the Drift → Health pipeline;
- future Rule Engine integration boundaries;
- future Policy Engine integration boundaries;
- future Report Engine integration boundaries;
- Canonical Health Record 1.0 and public-contract compatibility strategy;
- developer and user-facing conceptual documentation where needed, while keeping runtime rendering and localization work out of scope;
- roadmap and chronological Engineering History milestone updates;
- Task 021 history with the complete implementation and verification record;
- any required changelog or index updates mandated by the existing repository conventions.

All engineering artifacts must be in English. Documentation must clearly distinguish implemented Health behavior from future integration contracts and must not imply that monitoring, alerting, remediation, Policy, Compliance, Report, Rule, or Automation engines were implemented.


## Completion Criteria

Task 021 is complete only when:

- the Canonical Health Engine exists and evaluates current engineering condition solely from explicit canonical engineering evidence;
- Canonical Health Record 1.0, the Health taxonomy, deterministic evaluation pipeline, and versioned public contracts are implemented and documented;
- equivalent canonical inputs produce reproducibly equivalent, deterministically ordered Health outputs;
- the responsibility boundary is explicit and tested: Compare compares, Drift produces canonical drift evidence, and Health evaluates engineering condition without duplicating either layer;
- Health performs no monitoring, scheduling, daemon operation, alerting, notification, remediation, policy enforcement, compliance evaluation, AI processing, network access, or remote communication;
- all required architecture topics and future Rule, Policy, and Report integration boundaries are documented without implementing those engines;
- engineering tests and all mandatory repository/framework validations pass;
- compatibility, privacy, permission, ownership, ACL, Git diff, and rollback reviews pass or any platform limitation is explicitly documented and accepted;
- required lifecycle and chronological documentation is complete and truthful;
- the Task 021 history records starting state, verified snapshot, decisions, exact changes, validation evidence, rollback, Git state, unresolved issues, and delivery result;
- no excluded work or unrelated owner content was changed.

The valid delivery result is `complete`, `complete with disclosed limitations`, or `blocked`, as defined by the Engineering Framework. Completion does not authorize Task 022 or implementation of any future engine.


## Owner Approval Requirements

Approved by Attila — Project Owner through the Engineering Task Builder on 2026-07-24 UTC.

The structured task definition has been explicitly approved for implementation. Further scope changes require explicit Project Owner approval.
