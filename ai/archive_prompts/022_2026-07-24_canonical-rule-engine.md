# Current Engineering Task 022: Canonical Rule Engine

## Task Metadata

- Task ID: `022`
- Task slug: `canonical-rule-engine`
- Status: `complete`
- Date opened: `2026-07-24` UTC
- Human authority: Attila — Project Owner
- Owner or lead-developer communication language: Hungarian

## Title

Canonical Rule Engine


## Objective

Design and implement the permanent deterministic Rule Engine of Quantum Wizard Server Guardian.

The Rule Engine is the decision-matching layer above the Canonical Health Engine. It evaluates predefined canonical engineering rules against canonical Health Records and determines whether those rules match.

The Rule Engine must not reinterpret raw Inventory, Snapshot, Compare, or Drift data, must not duplicate Health evaluation, and must remain a pure, deterministic, reproducible, explainable, offline, AI-independent evaluation component.

The Engineering Goal is to establish QWSG's canonical deterministic rule-matching authority. Future Policy, Reporting, and Automation capabilities must consume versioned canonical Rule Evaluation Records instead of reimplementing rule logic.


## Scope

Design, implement, test, and document:

- the Canonical Rule Engine;
- a canonical versioned Rule definition contract;
- Canonical Rule Evaluation Record 1.0;
- a deterministic rule-matching model;
- a canonical Rule taxonomy;
- the Health → Rule evaluation pipeline;
- stable and reproducible identifiers;
- deterministic rule and result ordering;
- versioned public contracts;
- explainable evaluation evidence;
- compatibility and extension strategy;
- permanent architecture documentation;
- unit tests;
- integration and pipeline tests;
- Engineering History and lifecycle updates required by the Engineering Framework.

The permanent pipeline is:

Inventory → Snapshot → Compare → Drift → Health → Rule → future Policy → future Report → future Automation.

Responsibility boundaries:

- Compare determines what changed.
- Drift determines what kind of change occurred.
- Health determines the engineering health represented by canonical evidence.
- Rule determines whether a predefined deterministic condition matches.
- Future Policy decides how matched rules are interpreted or governed.
- Future Automation may act only after separate policy and authorization layers.

The Rule Engine shall consume canonical Health Records through public versioned contracts and produce canonical Rule Evaluation Records. It shall preserve Health IDs and evidence references, validate inputs and rule definitions, make every result explainable, reject or explicitly classify invalid and unsupported rules, preserve privacy and redaction boundaries, avoid hidden scoring and opaque logic, and never modify the host system.

The initial implementation must be the smallest complete canonical model. It may support deliberately limited deterministic equality, inequality, ordering, membership, existence, status/category matching, and bounded logical composition where justified by the final design, but it must remain extensible through versioned contracts.


## Out of Scope

The following are explicitly prohibited:

- Policy Engine implementation;
- Compliance Engine implementation;
- risk scoring;
- monitoring;
- scheduler or daemon implementation;
- Alert Engine implementation;
- notifications or email delivery;
- Report Engine implementation;
- dashboard implementation;
- automatic remediation;
- script execution;
- command or process execution;
- plugin execution;
- remote communication;
- cloud dependency;
- AI integration;
- machine learning;
- probabilistic evaluation;
- arbitrary code execution;
- a general-purpose programming language, scripting language, or arbitrary expression runtime.

The Rule Engine must not collect Inventory, create or persist snapshots, compare snapshots, classify Drift, reevaluate Health, reinterpret raw upstream values, produce policy decisions, authorize actions, generate alerts, render reports, execute remediation, or introduce side effects.

A matched rule is only a deterministic engineering result. It is not an alert, risk score, compliance result, policy decision, authorization, remediation instruction, or command.

No release, deployment, infrastructure mutation, dependency installation, remote Git operation, existing CLI behavior change, unrelated refactor, or implementation of future Policy, Report, Automation, Alert, Compliance, or monitoring capabilities is authorized.


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

1. Run `bin/job --check` and confirm Task 022 is the single active approved task with exactly one matching history record.
2. Read every Required Reading item and record completed reading in the Task 022 history.
3. Run `ai/scripts/framework-check.sh` and all configured baseline lifecycle validations.
4. Verify UTC date, effective user, repository root, configured project identity, branch, HEAD, remotes, ahead/behind relationship where safely available, relevant tags, and complete Git status.
5. Confirm Task 021 is complete and the Canonical Health Engine, Health Record 1.0, taxonomy, tests, architecture, and lifecycle evidence exist as the immediate dependency.
6. Inspect Inventory → Snapshot → Compare → Drift → Health public contracts, validation rules, ordering, identifiers, privacy boundaries, tests, architecture, permissions, ownership, and ACLs before designing Rule interfaces.
7. Confirm no Rule Engine implementation, target collision, unresolved editing marker, or overlapping untracked implementation exists.
8. Record and leave untouched all unrelated tracked and untracked owner content, including unfinished prior-task staging state.
9. Stop and report every material lifecycle, contract, Git, permission, ACL, ownership, dependency, or environment difference.


## Snapshot Requirements

Before modifying any Task 022 implementation or documentation target:

- create a unique UTC-timestamped Task 022 rollback snapshot under `/tmp`;
- capture every pre-existing target, active Task 022 prompt and history, relevant Git state, modes, ownership, ACL evidence where supported, and checksums;
- record verified absence for every planned new path;
- create and verify a complete Git bundle where supported;
- record the exact snapshot path, commands, target list, SHA-256 evidence, archive inspection, retention limitation, and guarded restore procedure in Task 022 history;
- verify the snapshot is readable, complete, private, and sufficient for exact bounded rollback.

Do not capture credentials, secrets, build caches, or unrelated private payloads. Do not modify task targets until snapshot verification passes.


## Risk Assessment

- Architectural-boundary risk — high: Rule could duplicate Health or reinterpret upstream evidence. Mitigate with a strict Health-only input contract, dependency-direction documentation, one-way provenance, and boundary tests.
- Language/runtime risk — high: an oversized expression system could enable arbitrary execution or unstable semantics. Mitigate with a deliberately small typed operator set, bounded composition, validation, resource limits, and no scripting or plugin runtime.
- Determinism risk — high: map iteration, ambiguous precedence, locale, time, randomness, or hidden state could alter results. Mitigate with canonical ordering, explicit inputs, fixed operator semantics, stable IDs, canonical serialization, and repeatability tests.
- Outcome-conflation risk — high: errors could be reported as non-matches. Mitigate with separate matched, not-matched, insufficient-evidence, unsupported-rule, invalid-rule, and evaluation-error outcomes and tests for each.
- Contract-compatibility risk — high: early Rule contracts will constrain Policy and Report consumers. Mitigate with explicit versions, validation, additive evolution rules, unsupported-version handling, and compatibility tests.
- Privacy risk — medium: explanations could disclose Health or upstream private values. Mitigate with bounded evidence references, safe reason tokens, redaction preservation, and disclosure tests.
- Logic-composition risk — medium: nested evaluation can become ambiguous or unbounded. Mitigate with explicit deterministic composition, depth/size bounds, stable traversal order, and rejection of unsupported structures.
- Numeric/type risk — medium: ordering operators can become coercive or platform-dependent. Mitigate with explicit supported types, no implicit coercion, canonical comparison rules, and invalid/unsupported classification.
- Stability and performance risk — medium: rule sets can create unbounded work. Mitigate with bounded validated inputs and predictable evaluation complexity.
- Security and side-effect risk — high: rule evaluation could be mistaken for authorization or execution. Mitigate with pure functions, no host/process/network imports, architecture exclusions, and dependency audits.
- Localization risk — low: engineering contracts remain English; descriptions are data and future presentation/localization belongs to Report/UI layers.
- Permission and rollback risk — low: preserve inherited ownership, modes, and ACLs; use exact target snapshots and bounded restoration.


## Planned Work

1. Validate Task 022 lifecycle, required reading, framework configuration, repository identity, Git state, Task 021 completion, and existing regression baseline.
2. Inspect Canonical Health Engine public contracts and establish the exclusive Health → Rule boundary.
3. Create and verify the exact Task 022 rollback snapshot before target modification.
4. Define the canonical Rule definition contract following QWSG conventions. It shall include concepts equivalent to Rule ID, contract version, category, scope, enabled state, input requirements, match conditions, expected Health category/status, deterministic operators, evaluation metadata, and human-readable description.
5. Define a minimal canonical operator model with explicit types, precedence, validation, composition bounds, and no arbitrary runtime. Select only operators needed for a complete safe 1.0 foundation.
6. Define Rule taxonomy and separate evaluation outcomes for matched, not matched, insufficient evidence, unsupported rule, invalid rule, and evaluation error.
7. Define Canonical Rule Evaluation Record 1.0 with stable evaluation identity, Rule ID, referenced Health Record ID, match result, evaluation status, deterministic non-probabilistic confidence or evidence sufficiency, explanation, bounded evidence references, metadata, and version information.
8. Implement the pure deterministic Rule Engine over explicit canonical Rule and Health inputs with stable evaluation order, result ordering, identifiers, validation, and canonical serialization.
9. Add unit tests for contracts, taxonomy, every outcome, every supported operator, type handling, bounded logical composition, determinism, stable IDs/order, invalid and unsupported rules, insufficient evidence, privacy, compatibility, and no input mutation.
10. Add integration tests proving real canonical Health input flows through Rule evaluation to Rule Evaluation Records without bypassing Health or reinterpreting Inventory, Compare, or Drift data.
11. Document Rule Engine architecture, Rule definition model, Rule Evaluation Record 1.0, taxonomy, operators, lifecycle, Health → Rule pipeline, evidence/explainability, invalid/unsupported handling, compatibility, privacy/security, and future Policy, Report, and Automation integration.
12. Update only required canonical architecture, system map, roadmap, README, Engineering History, and Task 022 lifecycle records.
13. Run every task-specific and repository-wide validation, inspect exact diffs, permissions, ACLs, privacy, excluded behavior, rollback evidence, and produce the English delivery record plus Hungarian owner handoff.

Future components must consume canonical Rule Evaluation Records. They must not create independent competing rule-matching logic.


## Rollback Plan

Rollback is bounded to the exact Task 022 target paths recorded before implementation.

1. Stop work and preserve truthful failure evidence.
2. Verify repository identity, Task ID, current Git state, exact target list, snapshot manifest, checksums, Git bundle, and archive listing.
3. Review the proposed restore set and obtain explicit confirmation before destructive replacement.
4. Restore only pre-existing Task 022 targets from exact snapshot counterparts, preserving recorded modes, ownership, and ACLs where authorized.
5. Remove only Task 022-created files whose exact paths were recorded as absent and are confirmed not to contain later owner work.
6. Preserve truthful lifecycle and failure evidence; restore prompt/history only if governance files are damaged.
7. Never use broad reset, checkout, restore, clean, wildcard deletion, or archive extraction over the live worktree.
8. Re-run framework, lifecycle, test-task audit, focused and repository Go tests, race, vet, format, documentation, privacy, permissions, ACL, Git diff, and status validation.
9. Record exact rollback actions, restored and retained paths, results, and unresolved issues in Task 022 history.

If exact safe restoration cannot be proven, stop and request Project Owner direction.


## Deliverables

- Canonical Rule Engine implementation.
- Canonical versioned Rule definition contract.
- Canonical Rule Evaluation Record 1.0.
- Deterministic rule-matching model and evaluation lifecycle.
- Canonical Rule taxonomy.
- Minimal safe deterministic operator model with bounded logical composition.
- Health → Rule pipeline preserving canonical Health evidence references.
- Stable identifiers, deterministic ordering, explainable outcomes, and canonical serialization.
- Versioned public contracts with validation, compatibility, and extension strategy.
- Unit tests and integration/pipeline tests.
- Permanent architecture documentation covering:
  - Rule Engine architecture;
  - Rule definition model;
  - Rule Evaluation Record 1.0;
  - Rule taxonomy;
  - deterministic operator model;
  - evaluation lifecycle;
  - Health → Rule pipeline;
  - explainability and evidence model;
  - invalid and unsupported rule handling;
  - versioning and compatibility strategy;
  - privacy and security boundaries;
  - future Policy Engine integration;
  - future Report Engine integration;
  - future Automation integration.
- Necessary canonical architecture, system-map, roadmap, README, Engineering History, and lifecycle updates.
- Complete Task 022 delivery history with starting state, snapshot, decisions, work, verification, rollback, Git state, limitations, and result.


## Verification

Verification must include:

- `bin/job --check`;
- `ai/scripts/framework-check.sh` and configured validations;
- `ai/scripts/next-task.sh --check`;
- `bin/job --check-test-tasks`;
- all Engineering Framework, lifecycle, diversion, and Task Builder assertion suites;
- focused Rule and Health package tests;
- all repository Go tests;
- race-enabled Go tests where supported;
- `go vet` and formatting checks;
- deterministic repeated evaluation and byte-stable canonical output;
- stable Rule and Evaluation identifiers and stable evaluation/result ordering;
- golden or equivalent public-contract tests for Rule definitions and Rule Evaluation Record 1.0;
- tests for matched, not matched, insufficient evidence, unsupported rule, invalid rule, and evaluation error, proving technical failure is never a normal non-match;
- tests for every supported equality, inequality, ordering, membership, existence, status/category, and logical-composition operator actually included in 1.0;
- explicit type, composition-depth/size, invalid-input, contradictory-input, duplicate, and unsupported-version tests;
- privacy and redaction-preservation tests;
- Health → Rule integration and at least one real pipeline test producing canonical Rule Evaluation output;
- tests proving Rule consumes Health contracts without reinterpreting Inventory, Compare, or Drift evidence or duplicating Health evaluation;
- direct dependency and behavior audits proving no Policy, Compliance, risk scoring, monitoring, scheduler, daemon, alerting, notification, reporting, dashboard, remediation, process/script/plugin execution, network, remote, cloud, AI, machine learning, probabilistic logic, host mutation, or side effect was added;
- existing CLI and public-contract regression tests;
- architecture-topic and terminology consistency checks;
- compatibility and extension-strategy review;
- exact target, mode, ownership, and ACL review where supported;
- snapshot checksum, archive, Git bundle, absence record, and bounded rollback review;
- `git diff --check`, exact unstaged/staged paths, and `git diff --cached --check` if staged;
- confirmation that unrelated owner-owned tracked and untracked content is untouched;
- confirmation that no dependency, infrastructure, release, deployment, staging, commit, push, or future-engine implementation occurred without authority.

Record exact commands, meaningful output summaries, assertion counts, limitations, and failures truthfully in Task 022 history. Any failing mandatory check blocks completion.


## Documentation Updates

Create or update only documentation necessary to make the Rule model permanent, canonical, and traceable:

- a dedicated Canonical Rule Engine architecture document;
- Rule definition contract and examples;
- Canonical Rule Evaluation Record 1.0;
- Rule taxonomy and deterministic operator semantics;
- evaluation lifecycle and deterministic ordering/identity rules;
- Health → Rule pipeline and responsibility boundary;
- explainability and bounded evidence model;
- invalid, unsupported, insufficient-evidence, and evaluation-error handling;
- versioning, backward compatibility, and extension strategy;
- privacy, security, offline, AI-independent, and no-side-effect boundaries;
- future Policy Engine integration;
- future Report Engine integration;
- future Automation integration;
- canonical architecture, system map, roadmap, README, and concise chronological Engineering History references;
- Task 022 history and any repository-mandated index or changelog updates.

All engineering artifacts must be English. Documentation must distinguish implemented Rule behavior from future Policy, Compliance, Alert, Report, Automation, UI, and operational behavior. It must state that future components consume canonical Rule Evaluation Records rather than implementing independent rule matching.


## Completion Criteria

Task 022 is complete only when:

- the Canonical Rule Engine evaluates only explicit canonical Rule definitions against validated canonical Health Records;
- the canonical Rule definition contract, Rule Evaluation Record 1.0, Rule taxonomy, minimal deterministic operator model, Health → Rule pipeline, and versioned public contracts exist and are documented;
- stable identifiers, canonical ordering, deterministic repeatability, explainable evidence, privacy preservation, strict validation, and compatibility behavior are implemented and tested;
- matched, not matched, insufficient evidence, unsupported rule, invalid rule, and evaluation error remain distinct and technical failure is never reported as normal non-match;
- every implemented operator and logical-composition boundary is documented and tested without creating an arbitrary language or execution runtime;
- at least one real Health → Rule pipeline test produces canonical Rule Evaluation output and proves the Rule layer does not duplicate upstream responsibilities;
- all required permanent architecture topics and future Policy, Report, and Automation boundaries are documented without implementing those engines;
- no Policy, Compliance, risk, monitoring, scheduler, daemon, alerting, notification, reporting, dashboard, remediation, command/script/plugin execution, remote/cloud, AI, machine learning, probabilistic, host-mutation, or side-effect behavior exists;
- all focused, repository, race, vet, format, framework, Builder, lifecycle, diversion, regression, privacy, compatibility, permission, ACL, rollback, and Git checks pass;
- existing CLI and public contracts do not regress;
- Task 022 history truthfully records starting state, verified snapshot, design decisions, exact work, validation evidence, rollback, Git state, unresolved issues, and delivery result;
- unrelated owner content remains untouched and no excluded work occurred.

The valid delivery result is `complete`, `complete with disclosed limitations`, or `blocked` under the Engineering Framework. Completion does not authorize Task 023 or implementation of any future Policy, Report, Alert, Automation, Compliance, monitoring, or remediation capability.


## Owner Approval Requirements

Approved by Attila — Project Owner through the Engineering Task Builder on 2026-07-24 UTC.

The structured task definition has been explicitly approved for implementation. Further scope changes require explicit Project Owner approval.
