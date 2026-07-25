# Current Engineering Task 023: Canonical Report Engine

## Task Metadata

- Task ID: `023`
- Task slug: `canonical-report-engine`
- Status: `complete — all authorized implementation and verification gates passed`
- Date opened: `2026-07-24` UTC
- Human authority: Attila — Project Owner
- Owner or lead-developer communication language: Hungarian

## Title

Canonical Report Engine


## Objective

Design and implement the permanent deterministic Report Engine of Quantum Wizard Server Guardian.

The Report Engine is the presentation-contract layer of the canonical engineering pipeline. It consumes canonical engineering records produced by previous layers and transforms deterministic engineering evidence into deterministic canonical engineering reports.

The Report Engine must never collect Inventory, perform Compare, classify Drift, evaluate Health, evaluate Rules, execute Policy, perform monitoring, generate alerts, execute remediation, or depend on AI.

The Engineering Goal is to establish QWSG's canonical deterministic reporting model. All future dashboards, exports, notifications, and management interfaces shall consume Canonical Reports instead of rebuilding engineering summaries independently.


## Scope

Design, implement, test, and document:

- the Canonical Report Engine;
- Canonical Report 1.0;
- a canonical Report taxonomy;
- a deterministic report-generation pipeline;
- a canonical report rendering model that remains structured and presentation-neutral;
- a source traceability and provenance model;
- stable identifiers and deterministic section, finding, evidence, and output ordering;
- versioned public contracts;
- compatibility and extension strategy;
- permanent architecture documentation;
- engineering tests;
- required Engineering History and lifecycle updates.

The permanent relationship is:

Inventory → Snapshot → Compare → Drift → Health → Rule → Report.

Report may consume validated canonical records from authorized previous layers through their versioned public contracts. Its primary decision-layer input shall preserve Rule Evaluation Record identity and the originating Health evidence chain. Where the final design supports Inventory, Snapshot, Compare, Drift, or Health material directly, it must use explicit typed source adapters and must never reinterpret or regenerate upstream semantics.

Canonical Report 1.0 shall provide a stable structured representation suitable for future renderers and interfaces. It shall include concepts equivalent to Report ID, report type, generation contract, deterministic title/summary data, ordered sections/items, canonical source references, source contract versions, traceability/provenance, metadata, completeness/unsupported evidence state, and report schema/version information.

The rendering model shall define deterministic structured or plain-text-safe presentation from Canonical Report data only. It must preserve localization readiness and must not implement HTML, PDF, Dashboard, email, notification delivery, or remote export.

The smallest safe implementation shall establish the complete permanent reporting contract without introducing future presentation channels or operational behavior.


## Out of Scope

The following are explicitly prohibited:

- Policy Engine implementation or policy execution;
- Compliance Engine or risk scoring;
- monitoring;
- scheduler or daemon implementation;
- Alert Engine implementation;
- notifications;
- email delivery;
- Dashboard UI;
- HTML rendering;
- PDF export;
- automatic remediation;
- script, command, process, or plugin execution;
- AI summarization or AI integration;
- machine learning;
- probabilistic summarization;
- remote communication;
- cloud dependency.

The Report Engine must not collect Inventory, create or persist snapshots, compare snapshots, classify Drift, reevaluate Health, reevaluate Rules, change Rule outcomes, execute Policy, authorize actions, produce alerts, deliver content, or mutate the host.

A Canonical Report is deterministic presentation evidence. It is not a policy decision, compliance result, risk score, alert, notification, authorization, remediation instruction, or action.

No existing CLI behavior change, Dashboard, HTML/PDF renderer, exporter, remote transport, dependency installation, infrastructure mutation, release, deployment, staging, commit, push, unrelated refactor, or future-engine implementation is authorized.


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

1. Run `bin/job --check` and confirm Task 023 is the single active approved task with one matching history record.
2. Read and record every Required Reading item.
3. Run `ai/scripts/framework-check.sh` and all configured baseline lifecycle validations.
4. Verify UTC date, effective user, repository root, project identity, branch, HEAD, remotes, ahead/behind relation where safely available, tags, and complete Git status.
5. Confirm Task 022 is complete and Canonical Rule Engine, Rule Definition 1.0, Rule Evaluation Record 1.0, tests, architecture, and lifecycle evidence exist.
6. Inspect Inventory → Snapshot → Compare → Drift → Health → Rule public contracts, identifiers, ordering, validation, privacy, traceability, tests, documentation, modes, ownership, and ACLs before designing Report contracts.
7. Confirm no Report Engine implementation, target collision, unresolved editing marker, or overlapping untracked implementation exists.
8. Record and leave untouched all unrelated tracked and untracked owner content and prior-task unstaged deliverables.
9. Stop on every material lifecycle, contract, compatibility, Git, permission, ACL, ownership, dependency, or environment difference.


## Snapshot Requirements

Before modifying Task 023 implementation or documentation targets:

- create a unique UTC-timestamped Task 023 rollback snapshot under `/tmp`;
- capture every pre-existing target, the active Task 023 prompt/history, relevant upstream untracked dependency files, Git state, modes, ownership, ACL evidence where supported, and checksums;
- record verified absence for each planned new path;
- create and verify a complete Git bundle where supported;
- document the exact snapshot path, commands, target list, SHA-256 evidence, archive inspection, absence evidence, retention limitation, and guarded restore procedure in Task 023 history;
- verify readability, privacy, integrity, completeness, and exact bounded rollback capability.

Do not capture credentials, secrets, build caches, or unrelated private payloads. No task target may change before snapshot verification passes.


## Risk Assessment

- Responsibility-boundary risk — high: Report could duplicate Inventory, Compare, Drift, Health, or Rule semantics. Mitigate with typed validated source contracts, preserved source IDs, no upstream derivation, and boundary tests.
- Traceability risk — high: summaries could lose or misattribute evidence. Mitigate with explicit source type, source ID, contract version, ordered evidence references, validation, and end-to-end pipeline tests.
- Determinism risk — high: maps, locale, time, rendering, or ambient state could change output. Mitigate with explicit inputs, canonical sorting, stable IDs, fixed rendering rules, canonical serialization, and repeatability tests.
- Presentation-truth risk — high: report text could introduce diagnosis, policy, or alert meaning absent from evidence. Mitigate with fixed reason/status vocabulary, source-derived structured content, and no free-form generated conclusions.
- Compatibility risk — high: future dashboards and exports will depend on Report 1.0. Mitigate with explicit schema/version fields, validation, additive evolution rules, unsupported-version behavior, and compatibility tests.
- Completeness risk — medium: missing or unsupported evidence could be presented as complete. Mitigate with explicit completeness, insufficient, invalid, and unsupported source semantics.
- Privacy risk — high: reports aggregate evidence and could disclose private values. Mitigate with redacted public contracts only, minimal source references, bounded metadata, safe text escaping, and disclosure tests.
- Rendering/security risk — medium: unsafe terminal or markup content could be interpreted. Mitigate with structured canonical data, control-character-safe plain text, no HTML/PDF/template/script runtime, and escaping tests.
- Localization risk — medium: hardcoded prose could block future localized renderers. Mitigate with machine-readable tokens, presentation-neutral report data, and separation of canonical engineering artifacts from future localized views.
- Resource risk — medium: large evidence sets could create unbounded reports. Mitigate with documented limits, bounded processing, stable truncation/error semantics, and tests.
- Side-effect risk — high: report generation could drift into export or delivery. Mitigate with pure offline APIs, no filesystem/process/network dependency, and exclusion audits.
- Permission and rollback risk — low: preserve inherited ownership, modes, and ACLs; use exact snapshots and bounded restore.


## Planned Work

1. Validate Task 023 lifecycle, required reading, framework configuration, repository identity, Git state, Task 022 completion, and baseline regressions.
2. Inspect every canonical upstream contract and establish the exclusive evidence-to-Report boundary.
3. Create and verify the exact rollback snapshot before modifying targets.
4. Define Report taxonomy, supported source types, report types, completeness/evidence semantics, and responsibility boundaries.
5. Define Canonical Report 1.0, including stable identity, ordered structured sections/items, source references, source-contract versions, traceability, bounded metadata, completeness, and version information.
6. Define a deterministic report-generation pipeline over explicit validated canonical source records. Preserve source identity and never regenerate upstream classifications or decisions.
7. Define a minimal presentation-neutral rendering model and, if justified, a safe deterministic plain-text representation produced solely from Canonical Report data. Do not add HTML, PDF, Dashboard, email, or remote rendering.
8. Implement strict contract validation, deterministic ordering, stable IDs, canonical serialization, bounded processing, privacy controls, and unsupported/invalid source handling.
9. Add unit tests for taxonomy, contracts, each supported source type, ordering, stable IDs, serialization, completeness, invalid/unsupported evidence, privacy, escaping, limits, and no input mutation.
10. Add integration tests covering Rule → Report and at least one real Inventory/Compare/Drift/Health/Rule → Report traceability pipeline.
11. Document Report architecture, Canonical Report model, taxonomy, generation pipeline, complete pipeline relationship, traceability, compatibility, future Dashboard/Export/Policy integration, privacy/security, and explicit exclusions.
12. Update only required canonical architecture, system map, roadmap, README, Engineering History, and Task 023 lifecycle records.
13. Run all required focused and repository-wide validation, inspect exact diffs, permissions, ACLs, privacy, excluded behavior, rollback evidence, and provide the English delivery record plus Hungarian owner handoff.

Future dashboards, exports, notifications, and management interfaces must consume Canonical Reports rather than rebuild engineering summaries.


## Rollback Plan

Rollback is bounded to the exact Task 023 targets recorded before implementation.

1. Stop and preserve truthful failure evidence.
2. Verify repository identity, Task ID, current Git state, exact target list, snapshot manifest, hashes, bundle, archive, and absence records.
3. Review the restore set and obtain explicit confirmation before destructive replacement.
4. Restore only pre-existing Task 023 targets from exact snapshot counterparts, preserving modes, ownership, and ACLs where authorized.
5. Remove only Task 023-created paths recorded as absent and confirmed not to contain later owner work.
6. Preserve truthful lifecycle evidence; restore prompt/history only if governance files are damaged.
7. Never use broad reset, checkout, restore, clean, wildcard deletion, or archive extraction over the live worktree.
8. Re-run framework, lifecycle, test-task audit, focused and repository tests, race, vet, format, documentation, privacy, permissions, ACL, Git diff, and status checks.
9. Record exact rollback operations, restored and retained paths, results, and unresolved issues in Task 023 history.

If safe exact restoration cannot be proven, stop and request Project Owner direction.


## Deliverables

- Canonical Report Engine implementation.
- Canonical Report 1.0 contract.
- Canonical Report taxonomy.
- Deterministic report-generation pipeline.
- Presentation-neutral deterministic report rendering model.
- Canonical source traceability and provenance model.
- Stable report, section, item, and evidence identity/order where applicable.
- Versioned public contracts with strict validation, compatibility, and extension strategy.
- Engineering unit and integration/pipeline tests.
- Permanent architecture documentation covering:
  - Report Engine architecture;
  - Canonical Report model;
  - Report taxonomy;
  - report-generation pipeline;
  - Inventory → Snapshot → Compare → Drift → Health → Rule → Report relationship;
  - traceability model;
  - compatibility strategy;
  - future Dashboard integration;
  - future Export integration;
  - future Policy integration;
  - privacy, security, localization, and no-side-effect boundaries.
- Necessary canonical architecture, system-map, roadmap, README, Engineering History, and lifecycle updates.
- Complete Task 023 history with starting state, snapshot, design decisions, changes, validation, rollback, Git state, limitations, and result.


## Verification

Verification must include:

- `bin/job --check`;
- `ai/scripts/framework-check.sh` and all configured validations;
- `ai/scripts/next-task.sh --check`;
- `bin/job --check-test-tasks`;
- all Engineering Framework, lifecycle, diversion, and Task Builder assertion suites;
- focused Report, Rule, and Health tests;
- all repository Go tests;
- race-enabled Go tests where supported;
- `go vet`, formatting, and Git diff checks;
- deterministic repeated generation and byte-stable canonical serialization;
- stable report, section, item, evidence identity, and ordering where implemented;
- golden or equivalent Canonical Report 1.0 contract tests;
- tests for every Report taxonomy type and supported source adapter;
- missing, incomplete, invalid, duplicate, contradictory, unsupported-source, and unsupported-version handling;
- source traceability from every report item to canonical source identity and contract version;
- privacy, redaction, bounded metadata, terminal/control-character safety, and localization-readiness checks;
- rendering-model tests proving output derives only from Canonical Report data;
- Rule → Report integration and at least one real canonical pipeline test through Report output;
- tests proving Report does not collect Inventory, perform Compare, classify Drift, evaluate Health/Rules, execute Policy, monitor, alert, notify, remediate, or introduce side effects;
- direct dependency audits proving no scheduler, daemon, process/script/plugin execution, HTML/PDF renderer, Dashboard, email, remote/cloud, AI, machine learning, probabilistic, host-mutation, or delivery dependency;
- preservation of existing CLI and public-contract behavior;
- architecture-topic and terminology consistency;
- compatibility and extension-strategy review;
- exact target mode, ownership, and ACL review where supported;
- snapshot checksum, archive, bundle, absence-record, and bounded rollback validation;
- `git diff --check`, exact staged/unstaged review, and `git diff --cached --check` if staged;
- confirmation that unrelated owner content remains untouched and no unauthorized dependency, infrastructure, release, deployment, staging, commit, push, or future capability occurred.

Record exact commands, meaningful output summaries, assertion counts, limitations, and failures truthfully. Any mandatory failure blocks completion.


## Documentation Updates

Create or update only documentation needed to make reporting canonical and traceable:

- a dedicated Canonical Report Engine architecture document;
- Canonical Report 1.0 model and examples;
- Report taxonomy;
- deterministic report-generation and rendering model;
- Inventory → Snapshot → Compare → Drift → Health → Rule → Report relationship;
- source traceability/provenance and evidence-reference model;
- completeness, invalid, and unsupported-source behavior;
- deterministic identity, ordering, limits, and canonical serialization;
- versioning, compatibility, and extension strategy;
- privacy, security, safe text, offline, AI-independent, localization-ready, and no-side-effect boundaries;
- future Dashboard integration;
- future Export integration;
- future Policy integration;
- canonical architecture, system map, roadmap, README, and concise Engineering History references;
- Task 023 history and required repository indexes.

All engineering artifacts must be English. Documentation must distinguish implemented Report contracts from future Policy, Dashboard, Export, Alert, Notification, HTML/PDF, Automation, and delivery channels. It must state that future interfaces consume Canonical Reports rather than independently rebuilding summaries.


## Completion Criteria

Task 023 is complete only when:

- the Canonical Report Engine transforms only explicit validated canonical evidence into Canonical Reports without duplicating upstream responsibilities;
- Canonical Report 1.0, Report taxonomy, deterministic generation pipeline, rendering model, traceability model, and versioned public contracts exist and are documented;
- stable identity/order, deterministic repeatability, canonical serialization, source provenance, completeness semantics, privacy, safe rendering, bounds, strict validation, and compatibility behavior are tested;
- every report item remains traceable to canonical source identity and source-contract version;
- missing, invalid, incomplete, and unsupported evidence is explicit and never silently presented as complete;
- Rule → Report and a real canonical pipeline test produce deterministic Canonical Report output;
- future Dashboard, Export, Notification, and management interfaces are documented as consumers of Canonical Reports rather than independent summary engines;
- all required architecture and future Dashboard, Export, and Policy boundaries are documented without implementing those components;
- no Policy, Compliance, risk, monitoring, scheduler, daemon, alerting, notification, email, Dashboard, HTML/PDF, remediation, execution, remote/cloud, AI, machine learning, probabilistic, host-mutation, delivery, or side-effect behavior exists;
- all focused, repository, race, vet, format, framework, Builder, lifecycle, diversion, regression, privacy, compatibility, permission, ACL, rollback, and Git checks pass;
- existing CLI and public contracts do not regress;
- Task 023 history is complete and truthful, unrelated owner content remains untouched, and no excluded work occurred.

The valid delivery result is `complete`, `complete with disclosed limitations`, or `blocked` under the Engineering Framework. Completion does not authorize Task 024 or any future Policy, Dashboard, Export, Alert, Notification, Automation, monitoring, or remediation work.


## Owner Approval Requirements

Approved by Attila — Project Owner through the Engineering Task Builder on 2026-07-24 UTC.

The structured task definition has been explicitly approved for implementation. Further scope changes require explicit Project Owner approval.
