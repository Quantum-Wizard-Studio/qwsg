# Current Engineering Task 062: Engineering Framework v2 — Development Freedom, Proportional Safety and Fast Execution

## Task Metadata

- Task ID: `062`
- Task slug: `engineering-framework-v2-development-freedom-proportional-safety-fast-execution`
- Status: `complete`
- Date opened: `2026-08-26` UTC
- Human authority: Project Owner
- Owner or lead-developer communication language: English

## Title

Engineering Framework v2 — Development Freedom, Proportional Safety and Fast Execution


## Objective

Design, implement, validate, document and migrate the reusable engineering framework to Framework 2.0.0 so one Owner-approved task grants broad normal in-scope engineering autonomy, proportional evidence and fast iterative execution while preserving the Builder, single active canonical task, snapshots, rollback, protected credentials, deterministic release construction, candidate immutability, repository cleanliness, validated integration and Project Owner control over architecture, external infrastructure, production and release/publication decisions. Use Tasks 057–061 as the concrete retrospective baseline and prove that v2 removes micro-gates, task proliferation and unnecessary Owner shell work without weakening genuine safety boundaries.


## Scope

- Audit Framework 1.1.0 policy, lifecycle, Builder, prompt workflow, Git policy, project-local job skill, validators and focused tests against Tasks 057–061.
- Define and implement the Framework 2.0.0 execution model: approved-scope default execution authority, iterative diagnose/fix/retest loops, proportional evidence tiers, explicit development/release separation, product-contract-first failure classification, reusable evidence, environment comparison and simplified release flow.
- Replace broad failure-as-STOP wording with precise boundary-based STOP semantics. A failed test, diagnostic, implementation attempt or incomplete evidence item must enter bounded classification/correction unless a real STOP boundary is reached.
- Preserve a concise Authority Envelope that inherits a standard reversible local execution baseline and records only task targets, task-specific external permissions, Owner-reserved decisions and genuine STOP conditions.
- Specify and implement a bounded diagnostic-runner pattern that consolidates multiple Owner-side checks into one reviewable command/artifact with declared mutation class, exact targets, privacy-safe machine-readable output, integrity identity and cleanup behavior.
- Update the exact reusable framework/core/scripts/tests and project-local integration files required for v2; add focused migration/compatibility tests and documentation.
- Preserve archived task/history bytes except where an index or prospective compatibility note is explicitly required; do not rewrite Tasks 057–061.
- Perform complete local validation, task-scoped Git integration, clean fast-forward push, direct Forgejo verification and canonical Task 062 closure.


## Out of Scope

- Do not modify QWSG product behavior, VERSION, candidate bytes, release notes or RC.6 acceptance conclusions.
- Do not access or mutate the test VPS, credentials, notification providers, private infrastructure or production systems.
- Do not create another QWSG release candidate, tag, Forgejo Release, upload, publication, deployment or announcement.
- Do not weaken the Project Constitution, credential/privacy protections, destructive-operation controls, candidate immutability after freeze, deterministic release requirements, one-active-task rule, repository cleanliness, rollback requirements or Owner authority over material architecture/product-scope/release decisions.
- Do not make all tasks release-grade, require a new task for ordinary in-scope correction, or add a second competing lifecycle system.
- Do not create or prepare Task 063.


## Authority Envelope

1. **Authorized paths/components/systems:** all reusable engineering-framework sources under `ai/framework`, `ai/core`, `ai/scripts`, `ai/tests`, relevant `ai/config` schema/validation declarations, `bin/job`, `.agents/skills/qwsg-job`, narrowly required repository guidance and Task 062 prompt/history/documentation. QWSG product source is read-only except test fixtures that directly validate framework integration and do not alter product behavior.
2. **Routine operations:** inspect repository/history, create private snapshots, edit/refactor framework code and documentation, add/remove/update focused tests and templates, run local checks/builds, classify failures, iterate implementation approaches, revert unsuccessful Task 062 changes through bounded rollback, update Task 062 evidence, explicitly stage reviewed paths, commit, push dry-run, clean fast-forward push, verify refs and close lifecycle without intermediate Owner gates.
3. **Correction/retest authority:** within the approved objective, independently repeat snapshot -> reproduce -> diagnose -> classify -> fix -> test -> retest -> integrate loops until Framework v2 satisfies its contract. Failed tests, diagnostics, implementation attempts, formatting, validation or evidence collection are recoverable work, not STOP conditions, unless they expose a genuine reserved boundary. Controlled refactoring and correction of the task's own framework implementation are authorized.
4. **Repository integration:** explicitly stage only reviewed Task 062 paths, inspect staged diff, create task-scoped implementation/evidence and lifecycle commits, push only clean fast-forwards to canonical `origin/main`, and verify direct Forgejo refs. No broad staging, history rewrite, force push or unrelated path inclusion.
5. **Lifecycle completion:** after truthful implementation and validation, finalize and archive Task 062, validate exactly zero active prompts and canonical idle, and report without another routine Owner gate. Do not prepare Task 063.
6. **Permitted external actions:** read-only canonical Git/Forgejo ref verification and the existing authorized clean fast-forward repository pushes only. No VPS, credential, infrastructure, publication, deployment, tag, Release or public asset action.
7. **Evidence and rollback:** create one verified private pre-change snapshot before tracked changes and a bounded closure snapshot; capture exact targets, hashes, Git/ref state, exclusions and literal rollback. Use lightweight evidence for low-risk documentation, focused evidence for framework behavior, and strong evidence for Builder/lifecycle/security/Git boundaries. Preserve the unrelated Owner blueprint untouched.
8. **Owner-reserved operations:** material changes to Project Constitution guarantees; QWSG product/architecture scope; credentials/secrets; destructive or irreversible external action; infrastructure mutation; production deployment; release candidate construction; tags; Forgejo Releases; upload/publication/announcement; force push/history rewrite; and Task 063.
9. **Mandatory STOP conditions:** credential/secret exposure risk; security/privacy regression or unresolved uncertainty; destructive external action outside this envelope; irreversible action without verified rollback; material QWSG product or architecture expansion; production/release/publication action; unauthorized external infrastructure mutation; Constitution conflict; or genuine uncertainty that bounded diagnosis/environment comparison cannot safely resolve. Ordinary failures, incomplete attempts and correctable test defects are explicitly not STOP conditions.


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

- Verify canonical idle after Task 061; exact `HEAD == origin/main ==` direct Forgejo `main` at the Task 061 closure commit; branch `main`; empty index; clean tracked worktree; Framework 1.1.0 VALID; zero active prompts; Task 061 complete/archived; and no Task 062 prompt/history.
- Verify the only unrelated state is the excluded Owner blueprint by pathname metadata only; do not read, hash, copy, modify or stage it.
- Record identities and permissions for Framework version/config, Constitution, agent rules, task standard, lifecycle, prompt workflow, Git policy, Builder, next-task/diversion/framework validators, project-local job skill, `bin/job`, and focused tests.
- Verify Tasks 057–061 histories and RC.6 acceptance evidence are immutable retrospective inputs. Record the Task 061 verdict: QWSG-059-F001 externally corrected, no RC.6 source defect, expected notification/overall partial state, and READY FOR RELEASE DECISION without release action.
- Run baseline Framework, Builder, lifecycle, diversion, active-job/test-task and Git checks before modification.


## Snapshot Requirements

Create one unique private mode-0700 Task 062 pre-change snapshot under `/tmp` before modifying tracked targets. Include a readable exact tracked-HEAD archive; literal before-images and hashes of every authorized framework/config/script/test/skill/guidance target; Task 062 lifecycle absence claims; Git/ref/status/mode/ACL evidence; immutable Task 057–061 history and Task 061 acceptance hashes; exact exclusion records; and bounded collision-aware rollback instructions. Exclude Owner blueprint content, credentials, candidate bytes, caches and unrelated files. Before lifecycle closure create a smaller private closure snapshot of the prompt/history and final integration state. Verify readability, hashes, modes, exclusions and restore instructions.


## Risk Assessment

- Excessive autonomy weakens safety — critical: standard authority covers only reversible in-scope local work; genuine security, credential, destructive external, architecture, production and release boundaries remain Owner-reserved.
- Simplification becomes another layer of ceremony — high: reduce concepts, reuse the existing Builder/lifecycle, define one standard execution baseline and measure interaction/validation reduction.
- Backward compatibility break — high: archived tasks remain immutable; Framework 1.1.0 task records stay valid; v2 Builder accepts/migrates existing structured inputs or fails with precise guidance; one-active-task semantics remain stable.
- Misclassification hides product defects — critical: require product-contract-first classification with explicit PRODUCT DEFECT, ACCEPTANCE/TEST DEFECT, ENVIRONMENTAL ISSUE, EXPECTED PRODUCT BEHAVIOR or INCONCLUSIVE outcomes; missing evidence remains missing.
- Diagnostic runner leaks secrets or mutates unexpectedly — critical: declared mode/targets, fixed bounded operations, no credential values/identities, privacy-safe schema, integrity check, timeout/output bounds and cleanup tests.
- Reduced validation misses regression — high: proportional validation maps checks to changed risk; security/Builder/lifecycle/release changes retain strong validation while redundant unchanged checks are reused.
- Git/lifecycle corruption — critical: snapshots, explicit staging, clean fast-forward only, direct ref verification, one active task and canonical idle remain mandatory.


## Planned Work

1. Produce a concise retrospective matrix for Tasks 057–061: useful protections, framework friction, task-specific overreach, false STOPs, redundant checks, manual Owner command count and how v2 changes each case.
2. Define Framework v2 architecture around five concepts: Standard Execution Authority, genuine Boundary STOPs, iterative Engineering Loop, proportional Evidence/Validation Tiers, and Development/Release Phase Separation.
3. Define the revised Authority Envelope: inherited reversible local authority plus explicit task targets/external permissions/reserved decisions, avoiding repeated enumeration of ordinary work.
4. Define failure classification before source change: PRODUCT DEFECT, ACCEPTANCE/TEST DEFECT, ENVIRONMENTAL ISSUE, EXPECTED PRODUCT BEHAVIOR or INCONCLUSIVE; add controlled second-environment comparison when materially useful and authorized.
5. Define the fast defect lifecycle: snapshot once, reproduce, diagnose, fix, focused test, broader retest proportional to impact, integrate and continue within one task; new numbered task only for material objective/scope/architecture change or a separately auditable phase.
6. Separate development evidence from candidate/release evidence. Preserve deterministic twin/frozen-candidate requirements only for candidate construction, and use one practical acceptance run before a separate release decision.
7. Design and implement the bounded diagnostic-runner contract/template/validator so multi-check Owner work becomes one checksum-verifiable command with privacy-safe key/value or JSON output and declared mutation/cleanup behavior.
8. Update Framework version and the smallest coherent set of core rules, Builder/rendering behavior, lifecycle/prompt/Git guidance, project-local skill and validators. Avoid rewriting the Constitution unless a demonstrated conflict requires Owner direction.
9. Add deterministic focused tests for inherited authority, ordinary iteration without gates, genuine STOP boundaries, failure classification, evidence tiers, diagnostic-runner safety, backward compatibility, single active task, snapshots, explicit staging and release-reserved actions.
10. Run migration fixtures against representative Tasks 057–061 and quantify the expected interaction reduction. Target no more than three inherent Owner interactions for a comparable practical acceptance run: private transfer/access, protected credential/receipt, and physical reboot when required, instead of repeated diagnostic fragments.
11. Run full required framework/governance/security/privacy/Git validation; document exact compatibility impact and migration steps; integrate and close Task 062 canonically.


## Rollback Plan

- Before implementation, restore only literal Task 062 targets from the verified private snapshot after checking current identities and collisions; remove only new files with proven prior absence/current Task 062 identity.
- During iteration, revert unsuccessful Task 062 changes through narrow patches or snapshot before-images, then rerun affected focused tests. Do not use broad reset, checkout, clean, history rewrite or touch unrelated/Owner content.
- If Framework 2.0.0 cannot preserve Builder/lifecycle/security compatibility, restore Framework 1.1.0 targets as one bounded rollback, retain the Task 062 failure analysis, validate Framework 1.1.0 and report the unresolved design boundary.
- After any pushed issue, use a forward correction commit only. Never force push or mutate published history.
- Post-rollback verification includes Framework version/validator, Builder/lifecycle/diversion/job suites, task/history identities, Git status/refs, permissions/ACLs, exclusions and canonical lifecycle state.


## Deliverables

- Exact Tasks 057–061 framework-friction retrospective with quantified Owner interaction and redundant-validation findings where evidence permits.
- Framework 2.0.0 architecture and normative execution model implementing default execution, local iterative correction, proportional safety/evidence, product-contract-first classification and reduced task proliferation.
- Revised concise Authority Envelope and genuine-boundary STOP semantics.
- Fast defect-correction lifecycle and clear development/candidate/acceptance/release separation.
- Bounded diagnostic-runner design plus reusable implementation/template, validator and machine-readable privacy-safe output contract.
- Migration and backward-compatibility plan for Framework 1.1.0 repositories, archived tasks, current Builder inputs and project-local skills.
- Exact updated framework/core/config/script/test/skill documentation and passing focused/full validation.
- Task 062 history, rollback evidence, validated task-scoped commits/push/ref evidence and canonical idle closure.


## Verification

- Framework reports version 2.0.0 and all internal references, schemas and documentation agree.
- Tests prove one approved task authorizes normal reversible in-scope inspect/edit/refactor/test/retry/rollback/document/integrate/close work without intermediate Owner gates.
- Tests prove ordinary failed tests/diagnostics/implementation attempts route to classification/correction, while genuine credential/security/privacy/destructive external/architecture/production/release boundaries STOP.
- Tests prove product-contract comparison precedes product-defect classification and all five failure classes are representable without manufacturing PASS.
- Tests prove proportional evidence tiers reuse unchanged facts and retain strict Builder/lifecycle/security/release validation where material.
- Diagnostic-runner tests prove fixed declared targets/actions, read-only versus mutating classification, timeout/output bounds, privacy-safe deterministic output, integrity identity, collision safety and cleanup; secret/private fixture values never appear.
- Compatibility fixtures prove archived Framework 1.1.0 tasks remain valid, current structured Builder fields migrate or validate deterministically, and exactly one active task/canonical idle remain enforced.
- Retrospective simulation shows a Tasks 059/061-style run requires only inherent Owner interactions rather than repeated command fragments, with assumptions disclosed.
- Full Framework, Builder, lifecycle, diversion, next-task, job/test-task, shell syntax, formatting, security/privacy, Git whitespace/diff/mode/ACL, rollback and repository checks pass.
- Explicit staging contains only reviewed Task 062 paths; clean fast-forward pushes and direct Forgejo refs match; no product source, candidate, tag, Release, publication, deployment, VPS action or Task 063 exists.


## Documentation Updates

- Update `ai/framework/VERSION` and the prospective framework architecture/standards needed to define Framework 2.0.0.
- Update the smallest coherent subset of `ai/core/03_AGENTS.md`, `08_JOB_TEMPLATE.md`, `11_ENGINEERING_LIFECYCLE.md`, `14_PROMPT_WORKFLOW.md`, `16_GIT_POLICY.md`, framework/config validation declarations, Builder/next-task/diversion/framework scripts, `bin/job`, `.agents/skills/qwsg-job/SKILL.md` and their focused tests where implementation requires.
- Add the bounded diagnostic-runner contract/template and its tests under clearly named reusable framework paths.
- Add a Framework 1.1.0 -> 2.0.0 migration/compatibility guide and Tasks 057–061 retrospective; update concise indexes only where required.
- Maintain and archive the Task 062 prompt/history with implementation decisions, failures/classifications, validation, rollback, commits and final state. Do not rewrite Task 057–061 histories or QWSG release evidence.


## Completion Criteria

Task 062 is complete when Framework 2.0.0 is coherently implemented and validated; approved tasks inherit broad reversible in-scope execution authority; ordinary engineering failures support autonomous diagnose/fix/retest iteration; STOP is limited to genuine boundaries; failure classification is product-contract-first; evidence and validation are proportional; diagnostic Owner work can be consolidated safely; development, candidate construction, practical acceptance and release decision are distinct; Framework 1.1.0 histories remain compatible; all focused/full tests pass; only Task 062 paths are integrated by clean fast-forward and directly verified; and the lifecycle returns to canonical idle. Completion must disclose any compatibility limitation and must not modify QWSG product behavior, construct/publish a candidate, perform external infrastructure work, release QWSG or create Task 063.


## Owner Approval Requirements

Approved by Project Owner through the Engineering Task Builder on 2026-08-26 UTC.

The structured task definition and Authority Envelope have been explicitly approved. The task is authorized to start and execute every routine operation inside that envelope without another Owner gate. Further scope changes and every Owner-reserved operation require explicit Project Owner approval.
