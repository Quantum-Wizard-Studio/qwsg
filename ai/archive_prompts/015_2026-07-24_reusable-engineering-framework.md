# Current Engineering Task 015: Reusable Engineering Framework Consolidation v1

## Task Metadata

- Task ID: `015`
- Task slug: `reusable-engineering-framework`
- Status: `complete`
- Date opened: `2026-07-23` UTC
- Human authority: Attila — Project Owner
- Owner or lead-developer communication language: Hungarian

## Title

Reusable Engineering Framework Consolidation v1


## Objective

Consolidate, harden, and generalize the existing QWSG Engineering Lifecycle, Engineering Task Builder, task execution workflow, snapshot and rollback system, Git safety rules, validation framework, and documentation structure into a reliable reusable engineering framework.

The framework shall remain fully operational for the current Quantum Wizard Server Guardian project while becoming sufficiently project-neutral and configurable to serve as the engineering foundation for future Quantum Wizard Studio software projects.

The purpose is to complete the reusable “software-building software” layer once, validate it thoroughly within QWSG, and prepare it for controlled reuse without duplicating or manually rebuilding the engineering workflow for every future project.

The implementation must preserve the existing QWSG workflow, including:

* deterministic task creation;
* one active engineering task at a time;
* explicit Project Owner approval;
* pre-change snapshots;
* documented rollback procedures;
* targeted Git staging;
* controlled commits and pushes;
* task-specific engineering history;
* delivery and completion evidence;
* automatic transition to the next task through the existing lifecycle tooling.

The framework must support project-specific objectives, architecture, terminology, repositories, languages, validation commands, and documentation while keeping the core safety and lifecycle rules consistent.

The result shall not be a separate speculative platform or an unfinished rewrite. It shall be a working, tested evolution of the existing QWSG engineering system.


## Scope

The scope includes:

* Audit the current engineering framework as implemented in the QWSG repository.
* Identify which lifecycle rules are universal and which rules are QWSG-specific.
* Preserve all current working behavior before introducing generalization.
* Create a verified pre-change repository snapshot and documented rollback procedure.
* Consolidate canonical engineering rules into clearly defined core policy documents.
* Permanently document the verified HTTPS Git workflow.
* Establish canonical rules for:

  * repository state verification;
  * remote synchronization;
  * targeted staging;
  * commit review;
  * dry-run push verification;
  * normal push behavior;
  * prohibited destructive Git operations;
  * branch and tag safety;
  * untracked-file protection;
  * completion evidence.
* Preserve HTTPS as the canonical QWSG Git transport.
* Record that Forgejo SSH access on port 2222 is not required for normal QWSG development.
* Review the Engineering Constitution, Engineering Lifecycle, Prompt Workflow, Agent Rules, Job Template, Task Builder, validation scripts, task transition scripts, and related documentation for duplication, contradiction, and project-specific coupling.
* Define the reusable engineering framework as a versioned core system.
* Separate reusable framework configuration from QWSG-specific project configuration where this can be done safely.
* Introduce a clear project identity and project configuration model.
* Ensure the framework can represent, at minimum:

  * project name;
  * project slug;
  * project repository path;
  * canonical Git remote;
  * primary branch;
  * communication language;
  * documentation language;
  * required reading;
  * project-specific validation commands;
  * project-specific directories;
  * task numbering;
  * task and history naming rules;
  * snapshot location;
  * rollback requirements.
* Keep all safety-critical defaults enabled.
* Ensure project-specific configuration cannot silently weaken mandatory safety rules.
* Preserve deterministic task numbering and the rule that only one active task may exist.
* Preserve explicit Project Owner approval before task activation.
* Preserve transactional task installation and rollback behavior.
* Preserve task-specific history files instead of replacing them with one monolithic history file.
* Preserve automatic archival of completed task prompts and records.
* Preserve the existing `bin/job`, `qwtask`, `ai/scripts/task-builder.sh`, and `ai/scripts/next-task.sh` workflows unless a controlled compatible improvement is required.
* Improve the Task Builder so that reusable framework rules are referenced from canonical core documents rather than unnecessarily duplicated in every generated task.
* Keep each generated task self-contained enough to remain understandable and safely executable.
* Define which mandatory sections must always remain in generated engineering tasks.
* Define which content may be inherited from canonical policy documents.
* Ensure generated task metadata remains deterministic and machine-verifiable.
* Ensure generated task files remain human-readable and editable using UTF-8 with LF line endings.
* Maintain compatibility with the current exact Task ID header validation rule.
* Add framework-level validation that detects:

  * missing required files;
  * invalid project configuration;
  * unresolved placeholders;
  * inconsistent task IDs;
  * multiple active tasks;
  * missing owner approval;
  * missing snapshot requirements;
  * missing rollback requirements;
  * missing verification sections;
  * missing completion evidence requirements;
  * unsafe Git instructions;
  * accidental broad staging instructions;
  * project-specific paths embedded in reusable core logic where configuration should be used.
* Add or update automated assertions for all modified framework components.
* Verify that all existing Task Builder and lifecycle assertions continue to pass.
* Add regression tests for current QWSG behavior.
* Add portability tests using a safe temporary fixture project or sandbox directory.
* The fixture shall demonstrate that the framework can initialize or validate a second project identity without modifying or contaminating the live QWSG task state.
* Ensure temporary fixture data is removed after testing.
* Create documentation explaining how the reusable framework can later be adopted by another project.
* Document the boundary between:

  * reusable engineering core;
  * project configuration;
  * project-specific architecture;
  * current engineering task;
  * task history;
  * generated delivery evidence.
* Version the resulting reusable framework.
* Record every change, validation result, compatibility decision, and known limitation.
* Update the official QWSG Engineering History.
* Produce a final Engineering Delivery Report.

The implementation must maintain the QWSG repository as the reference implementation and first validated consumer of the reusable engineering framework.


## Out of Scope

The following items are explicitly out of scope:

* Starting QWSG runtime feature development unrelated to the engineering framework.
* Modifying the canonical system inventory collectors.
* Modifying the Collector Framework.
* Changing Inventory 1.0 runtime behavior.
* Implementing continuous monitoring, alerts, email delivery, daemon mode, web interfaces, APIs, dashboards, or paid services.
* Migrating QUWIP, QuantumWizard.hu, AlexTamas.hu, or any other existing project to the framework.
* Modifying another production repository.
* Creating a shared global installation outside the QWSG repository.
* Publishing the framework as a public standalone product or package.
* Creating an external package registry release.
* Creating a new Forgejo repository for the framework.
* Splitting the framework into a separate repository.
* Creating a Composer, npm, Go, Debian, Snap, Docker, or other distributable package.
* Rewriting the complete engineering system from scratch.
* Replacing working Bash tooling merely for stylistic reasons.
* Introducing Python, Node.js, PHP, Go, or another runtime dependency unless technically necessary and explicitly justified.
* Adding unnecessary third-party dependencies.
* Weakening snapshot, rollback, validation, owner approval, or Git safety requirements.
* Making safety-critical policies optional.
* Permitting multiple active engineering tasks.
* Permitting tasks to bypass explicit Project Owner approval.
* Automatically committing or pushing without the required validation and review sequence.
* Running broad Git staging commands such as:

  * `git add .`
  * `git add -A`
  * `git add --all`
* Automatically deleting, moving, cleaning, committing, or ignoring pre-existing untracked files.
* Running destructive Git commands without explicit Project Owner approval.
* Force-pushing, rebasing, squashing, rewriting published history, or deleting existing tags.
* Changing the canonical QWSG `main` branch.
* Changing the current Forgejo repository URL.
* Replacing HTTPS Git operations with SSH.
* Repairing or reconfiguring the Forgejo SSH service on port 2222.
* Changing server firewall rules.
* Changing HestiaCP, Nginx, Apache, MariaDB, Forgejo, mail, or operating-system infrastructure.
* Modifying production services outside the repository.
* Creating user credentials, access tokens, deploy keys, or SSH keys.
* Recording secrets in project files, logs, snapshots, or reports.
* Performing an uncontrolled server reboot.
* Creating artificial historical commits.
* Renumbering or rewriting Tasks 001–015.
* Reformatting historical task files solely for consistency.
* Deleting existing engineering backups, histories, prompts, or delivery evidence.
* Automatically adding the currently untracked historical backup, history, or prompt directories to Git without explicit scope review.
* Treating the temporary portability fixture as a production project.
* Claiming full cross-platform portability beyond what is actually verified.
* Preparing the next QWSG product-development task.
* Creating Task 017.


## Required Reading

- `ai/core/00_PROJECT_PHILOSOPHY.md`
- `ai/core/01_CONSTITUTION.md`
- `ai/core/03_AGENTS.md`
- `ai/core/08_JOB_TEMPLATE.md`
- `ai/core/11_ENGINEERING_LIFECYCLE.md`

## Starting State Verification

The implementation starts from the following verified baseline:

* QWSG Engineering Lifecycle v1 is operational.
* The Engineering Task Builder is operational.
* The Task Builder supports:

  * interactive task creation;
  * deterministic `--input-dir` operation;
  * multi-line field input;
  * automatic metadata;
  * automatic Owner Approval generation;
  * explicit `APPROVE` confirmation;
  * transactional task installation;
  * rollback on installation failure.
* Existing Task Builder assertions pass.
* Existing lifecycle assertions pass.
* The canonical task lifecycle enforces one active task at a time.
* Completed tasks are archived.
* Task-specific history files are used.
* Pre-change snapshots and restore instructions are required.
* QWSG foundation milestone commit exists:
  `b4316050436bf8be4062f0e1d4ba7c371c334223`
* Annotated release tag exists:
  `v0.1.0`
* Forgejo is installed and operational.
* Canonical QWSG repository remote:
  `https://git.quantumwizard.hu/Quantum_Wizard_Studio/qwsg.git`
* Canonical remote name:
  `origin`
* Canonical primary branch:
  `main`
* Git HTTPS read and write access has been verified.
* The following Git operations have been verified successfully:

  * fetch;
  * dry-run push;
  * temporary remote branch creation;
  * remote branch verification;
  * temporary remote branch deletion;
  * deletion verification.
* The temporary Git test branch was completely removed.
* The `main` branch remained unchanged during the Git transport test.
* The local `main` branch is synchronized with `origin/main`.
* Forgejo SSH on port 2222 is not required for normal development.
* The repository contains existing untracked engineering backup, history, and prompt paths.
* These untracked paths must not be automatically staged, deleted, moved, cleaned, or committed.
* The current working repository remains operational.
* No Task 017 or later has been created.
* The current engineering framework is primarily embedded within the QWSG repository and contains some QWSG-specific assumptions.
* The objective is controlled consolidation and generalization, not replacement of the working lifecycle.


## Snapshot Requirements

Before modifying any tracked project file, create a verified engineering snapshot.

The snapshot shall include, at minimum:

* Current branch name.
* Current commit hash.
* Current tag references.
* Current configured Git remotes.
* Current local-to-remote branch relationship.
* Complete `git status --short --branch` output.
* Complete list of tracked files that may be modified.
* Complete list of relevant untracked paths without copying unrelated large backup contents into the task snapshot.
* Copies of every framework file that may be modified.
* Current hashes of:

  * Task Builder;
  * lifecycle scripts;
  * validators;
  * job command;
  * canonical engineering core documents;
  * templates;
  * relevant Makefile targets.
* Current Task Builder test results.
* Current lifecycle test results.
* Current task-state validation results.
* Current file permissions and executable bits for modified scripts.
* Current UTF-8 and line-ending state for modified text files.
* Current framework directory structure.
* Current project-specific constants and embedded paths identified during audit.

The snapshot shall:

* use a unique UTC timestamp;
* include a manifest;
* include SHA-256 hashes;
* include a starting-state record;
* include exact restore instructions;
* remain separate from unrelated historical backups;
* contain no secrets, credentials, access tokens, private keys, or sensitive runtime data;
* be verified before any implementation change begins.

No framework modification may begin until:

* the snapshot exists;
* the manifest is complete;
* integrity verification succeeds;
* rollback instructions are readable and executable;
* current lifecycle validation passes or any pre-existing failure is explicitly documented.


## Risk Assessment

This task modifies the software that controls how future engineering tasks are created, approved, executed, validated, archived, committed, and delivered.

A defect could affect every later QWSG task and any future project that adopts the framework.

The primary risks are:

* Breaking the currently operational Task Builder.
* Breaking `qwtask`.
* Breaking `bin/job`.
* Breaking task numbering.
* Allowing multiple active tasks.
* Producing malformed task metadata.
* Losing explicit Project Owner approval.
* Generating tasks with unresolved placeholders.
* Removing required snapshot or rollback requirements.
* Creating contradictory rules across core documents.
* Over-generalizing the framework until it becomes difficult to use.
* Leaving QWSG-specific assumptions hidden inside supposedly reusable code.
* Creating configuration that can weaken mandatory safety rules.
* Introducing incompatible paths or directory layouts.
* Changing script permissions or line endings.
* Accidentally staging existing untracked files.
* Accidentally committing backup data.
* Automatically pushing unintended changes.
* Producing a framework that works only in tests but not in the live QWSG workflow.
* Producing documentation that claims portability not demonstrated by evidence.
* Making future project adoption unnecessarily complex.
* Creating duplicate sources of truth.
* Coupling the reusable framework to the current task content.
* Losing deterministic behavior.
* Invalidating historical tasks.
* Making rollback incomplete.

Risk mitigation requirements:

* Preserve the working baseline before changes.
* Audit before refactoring.
* Make incremental changes.
* Validate after each phase.
* Keep the live QWSG workflow operational at all times.
* Use compatibility-preserving changes where possible.
* Do not rewrite functioning components without demonstrated need.
* Maintain one canonical source of truth for each rule.
* Keep mandatory safety rules outside project-overridable configuration.
* Use targeted Git staging only.
* Review the staged diff before commit.
* Perform dry-run push verification before real push.
* Do not continue past a failed mandatory validation.
* Restore the previous framework if compatibility cannot be demonstrated.
* Record every design decision and limitation.


## Planned Work

The implementation shall be performed in the following controlled phases:

### Phase 1 — Engineering Preparation

* Verify the starting repository state.
* Verify `main` and `origin/main`.
* Record the current remote URL.
* Record existing untracked paths.
* Run all current framework tests.
* Create and verify the engineering snapshot.
* Confirm rollback capability.
* Confirm no unrelated task is active.

### Phase 2 — Framework Inventory

* Inventory all engineering framework files.
* Map relationships between:

  * Constitution;
  * Agent Rules;
  * Architecture;
  * Engineering Lifecycle;
  * Prompt Workflow;
  * Job Template;
  * Task Builder;
  * task transition script;
  * job command;
  * validators;
  * tests;
  * task history;
  * archives;
  * backups.
* Identify duplicate rules.
* Identify contradictory rules.
* Identify QWSG-specific constants.
* Identify reusable core behavior.
* Identify project-specific configuration candidates.
* Record the inventory before implementation.

### Phase 3 — Canonical Policy Consolidation

* Define the canonical reusable engineering core.
* Consolidate mandatory safety rules into appropriate core policy documents.
* Establish a canonical Git policy.
* Document HTTPS as the QWSG transport.
* Document targeted staging requirements.
* Document prohibited destructive Git operations.
* Document branch and tag rules.
* Document untracked-file protection.
* Document commit and push evidence requirements.
* Ensure other documents reference canonical policies rather than redefining them inconsistently.
* Preserve essential task-level safety context.

### Phase 4 — Project Configuration Model

* Design a minimal, deterministic project configuration model.
* Separate project identity from universal engineering safety rules.
* Define mandatory project configuration fields.
* Define validation rules.
* Define safe defaults.
* Ensure project configuration cannot disable:

  * snapshot requirements;
  * rollback requirements;
  * owner approval;
  * single active task enforcement;
  * validation;
  * targeted staging;
  * completion evidence.
* Populate the QWSG project configuration from verified current values.
* Avoid introducing secrets into configuration.

### Phase 5 — Task Builder Compatibility Update

* Update the Task Builder only where required.
* Preserve current interactive usage.
* Preserve deterministic `--input-dir` usage.
* Preserve explicit `APPROVE`.
* Preserve transactional installation.
* Preserve rollback on failure.
* Preserve exact Task ID formatting.
* Preserve UTF-8 LF output.
* Make generated tasks reference canonical policies.
* Prevent excessive duplication without making generated tasks unsafe or unclear.
* Validate unresolved-placeholder detection.
* Validate task metadata generation.
* Validate owner approval generation.

### Phase 6 — Lifecycle and Validation Hardening

* Update lifecycle validators where required.
* Validate project configuration.
* Validate required core policy files.
* Validate single active task state.
* Validate task numbering.
* Validate owner approval.
* Validate snapshot requirements.
* Validate rollback requirements.
* Validate verification requirements.
* Validate completion evidence requirements.
* Detect unsafe Git instructions.
* Detect broad staging commands in generated task content where applicable.
* Detect accidental project-specific hardcoding in reusable logic.
* Preserve current QWSG lifecycle behavior.

### Phase 7 — Portability Fixture

* Create a temporary isolated fixture outside the live task state.
* Assign it a separate test project identity.
* Use only synthetic, non-production data.
* Validate project configuration loading.
* Validate task generation or framework validation in the fixture.
* Verify that the fixture does not modify:

  * QWSG current task;
  * QWSG task numbering;
  * QWSG history;
  * QWSG archive;
  * QWSG backups;
  * QWSG Git remote.
* Remove temporary fixture artifacts after evidence is recorded.
* Document exactly what portability was and was not verified.

### Phase 8 — Regression Testing

* Run all original Task Builder assertions.
* Run all original lifecycle assertions.
* Add tests for new behavior.
* Add compatibility tests for QWSG.
* Add negative tests for invalid configuration.
* Add negative tests for missing approval.
* Add negative tests for unsafe Git instructions.
* Add negative tests for unresolved placeholders.
* Add negative tests for multiple active tasks.
* Add rollback tests where safe.
* Verify script executable permissions.
* Verify UTF-8 LF output.
* Verify shell syntax and static checks.
* Verify deterministic results.

### Phase 9 — Live QWSG Workflow Verification

* Verify the current QWSG task remains valid.
* Verify `bin/job --check`.
* Verify the lifecycle state.
* Verify the next-task tooling without creating Task 017.
* Verify Task Builder help and validation modes.
* Verify no live task or history file is accidentally replaced.
* Verify no unrelated untracked file is modified.
* Verify the QWSG remote remains unchanged.
* Verify no automatic push occurred.

### Phase 10 — Documentation

* Document the reusable framework architecture.
* Document the project configuration model.
* Document canonical policies.
* Document QWSG as the reference implementation.
* Document future adoption procedure.
* Document compatibility guarantees.
* Document known limitations.
* Document test coverage.
* Update Engineering History.
* Produce the Engineering Delivery Report.
* Produce implementation and validation evidence.

### Phase 11 — Git Review and Delivery

* Review `git status --short`.
* Review unstaged changes.
* Review staged changes.
* Stage only explicitly approved Task 015 files.
* Do not stage unrelated existing untracked files.
* Run final validation.
* Create a clear Task 015 commit.
* Perform a dry-run push.
* Push only after all mandatory checks pass.
* Record the commit hash and push result.
* Confirm final repository state.


## Rollback Plan

Rollback shall restore the complete pre-Task-015 engineering framework.

If any mandatory validation fails, implementation shall stop immediately.

Rollback may include:

* Restore modified core policy documents from the verified snapshot.
* Restore the previous Task Builder.
* Restore the previous lifecycle scripts.
* Restore the previous validators.
* Restore the previous `bin/job`.
* Restore the previous Makefile targets.
* Remove newly created framework files.
* Remove newly created project configuration files.
* Remove temporary portability fixture files.
* Restore executable permissions.
* Restore original UTF-8 LF file state.
* Restore original task templates.
* Restore the previous documentation references.
* Re-run all original Task Builder assertions.
* Re-run all original lifecycle assertions.
* Re-run current task validation.
* Confirm that Task 015 remains correctly represented or restore its pre-implementation state as documented.
* Confirm that no Task 017 was created.
* Confirm that the Git remote remains unchanged.
* Confirm that the `main` branch and published history were not rewritten.
* Confirm that existing untracked files remain untouched.

Rollback requirements:

* No rollback step may use broad destructive Git commands.
* No rollback step may delete unrelated untracked files.
* No rollback step may modify historical Tasks 001–015.
* No rollback step may change the Forgejo server.
* No rollback step may modify unrelated production services.
* No rollback step may expose secrets.
* Every rollback action shall be recorded.
* After rollback, the original framework tests must pass.
* After rollback, the current QWSG lifecycle must remain operational.


## Deliverables

The completed task shall deliver:

### Reusable Engineering Core

* A clearly defined reusable engineering framework.
* Versioned framework identity.
* Canonical boundaries between reusable and project-specific components.
* Consolidated mandatory safety policies.
* Reduced policy duplication.
* No unresolved contradictions among canonical engineering documents.

### Git Policy

* Canonical HTTPS Git workflow.
* Repository verification rules.
* Fetch and synchronization rules.
* Targeted staging rules.
* Commit review rules.
* Dry-run push rules.
* Push rules.
* Prohibited destructive-operation rules.
* Branch and tag safety rules.
* Untracked-file protection.
* Completion evidence requirements.

### Project Configuration

* A validated project configuration model.
* QWSG project configuration.
* Safe mandatory defaults.
* Protection against project-level weakening of core safety rules.
* Documentation for future project adoption.

### Task Builder

* Fully operational Task Builder.
* Existing interactive workflow preserved.
* Deterministic input workflow preserved.
* Explicit approval preserved.
* Transactional installation preserved.
* Rollback preserved.
* Canonical policy references integrated.
* Required task structure preserved.
* No unresolved placeholders.
* UTF-8 LF output preserved.

### Lifecycle

* One active task enforcement preserved.
* Deterministic task numbering preserved.
* Task-specific history preserved.
* Archive behavior preserved.
* Snapshot requirements preserved.
* Rollback requirements preserved.
* Completion evidence preserved.
* Validation strengthened.

### Testing

* Existing assertions passing.
* New framework assertions passing.
* QWSG compatibility tests passing.
* Invalid-configuration tests passing.
* Safety-rule tests passing.
* Portability fixture test passing.
* Temporary fixture removed after testing.

### Documentation

* Reusable framework architecture.
* Project configuration reference.
* Git policy.
* Adoption guide.
* Compatibility and limitation record.
* Updated Engineering History.
* Final Engineering Delivery Report.
* Implementation Record.
* Validation Record.
* Rollback Record if rollback was used.


## Verification

The task shall not be considered complete until all verification steps have passed.

### Engineering Baseline

* Starting state recorded.
* Snapshot created.
* Snapshot integrity verified.
* Restore procedure verified.
* Existing tests recorded before modification.

### Core Policies

* Every mandatory policy has one canonical source of truth.
* Cross-references are valid.
* No unresolved contradiction remains.
* QWSG-specific details are not incorrectly presented as universal rules.
* Universal safety rules cannot be overridden by project configuration.

### Git Safety

* Canonical remote remains:
  `https://git.quantumwizard.hu/Quantum_Wizard_Studio/qwsg.git`
* Remote name remains:
  `origin`
* Primary branch remains:
  `main`
* HTTPS workflow documented.
* Targeted staging enforced.
* Broad staging commands are not introduced.
* Destructive Git operations are not introduced.
* Existing untracked paths remain untouched.
* Dry-run push procedure documented and followed.
* No force-push occurred.
* No published history was rewritten.

### Project Configuration

* QWSG configuration validates.
* Required fields are present.
* Invalid configuration is rejected.
* Missing required configuration is rejected.
* Unsafe overrides are rejected.
* No secret is stored.
* Paths resolve correctly.

### Task Builder

* Interactive mode works.
* Deterministic input mode works.
* Approval is mandatory.
* Rejection or missing approval prevents activation.
* Transactional installation works.
* Failed installation rolls back.
* Generated Task ID format is exact.
* Generated files use UTF-8 LF.
* Required sections are present.
* No unresolved placeholder remains.
* Canonical policy references are correct.

### Lifecycle

* Exactly one active task is allowed.
* Multiple active tasks are rejected.
* Task numbering remains deterministic.
* Task-specific history remains operational.
* Archive behavior remains operational.
* Completed-task validation remains operational.
* `bin/job --check` passes.
* Next-task validation works without creating Task 017.

### Regression

* All pre-existing Task Builder assertions pass.
* All pre-existing lifecycle assertions pass.
* All newly added assertions pass.
* Shell syntax validation passes.
* Static checks pass.
* File permission checks pass.
* No unrelated repository file changed.

### Portability

* Temporary fixture uses a separate project identity.
* Fixture configuration validates.
* Framework behavior works in the fixture to the documented extent.
* QWSG live state remains unchanged.
* Fixture data is removed after verification.
* Portability claims match actual evidence.

### Documentation

* Framework architecture documented.
* Project configuration documented.
* Git policy documented.
* Adoption procedure documented.
* Known limitations documented.
* Engineering History updated.
* Delivery Report completed.
* Final documentation reflects the implemented state.

### Final Repository State

* Only Task 015 files are staged.
* Staged diff reviewed.
* Final tests pass.
* Commit created.
* Dry-run push succeeds.
* Real push succeeds if approved by the active workflow.
* Final commit hash recorded.
* Final `git status --short --branch` recorded.
* Local `main` and `origin/main` are synchronized.
* No Task 017 exists.


## Documentation Updates

The following documentation shall be created or updated as appropriate:

### Canonical Engineering Documentation

* Project Philosophy references where required.
* Engineering Constitution.
* Agent Rules.
* Engineering Architecture.
* Engineering Lifecycle.
* Prompt Workflow.
* Job Template.
* Canonical Git Policy.
* Reusable Framework Architecture.
* Project Configuration Reference.
* Framework Adoption Guide.

### Operational Documentation

* Task Builder usage.
* Lifecycle commands.
* Validation commands.
* Snapshot procedure.
* Restore procedure.
* Git review and push procedure.
* Troubleshooting guidance.
* Portability fixture procedure.
* Upgrade and maintenance guidance for the framework.

### Engineering Records

* Engineering History.
* Task 015 Implementation Record.
* Task 015 Validation Record.
* Task 015 Delivery Report.
* Rollback Record if required.

### Documentation Requirements

* Documentation shall describe the final implemented state.
* Documentation shall distinguish universal rules from QWSG-specific configuration.
* Documentation shall not claim unverified portability.
* Documentation shall contain no secrets.
* Documentation shall avoid unnecessary duplication.
* Canonical policies shall be clearly identified.
* Cross-references shall resolve correctly.
* Commands shall be safe and reproducible.
* Historical task records shall not be rewritten.


## Completion Criteria

This task is complete only when all of the following are true:

### Framework

* The reusable engineering core is clearly defined.
* QWSG remains its working reference implementation.
* Universal and project-specific concerns are separated.
* Mandatory safety rules remain enforced.
* No critical behavior has regressed.
* No duplicate source of truth remains for core policies.

### Task Builder

* Task creation remains operational.
* Deterministic mode remains operational.
* Approval remains mandatory.
* Transactional installation remains operational.
* Rollback remains operational.
* Generated tasks remain valid and readable.
* Required sections remain present.
* Canonical policy references are used correctly.

### Lifecycle

* One active task rule remains enforced.
* Task numbering remains correct.
* History and archive behavior remain correct.
* Snapshot and rollback rules remain mandatory.
* Completion evidence remains mandatory.
* Existing QWSG workflow remains usable.

### Portability

* A second synthetic project identity has been safely validated.
* The fixture did not alter live QWSG state.
* The demonstrated reuse boundary is documented.
* Known limitations are documented.
* No migration of another real project occurred.

### Git

* HTTPS remains canonical.
* Git policy is permanently documented.
* Targeted staging was used.
* Existing untracked files were preserved.
* No destructive operation occurred.
* No force-push occurred.
* Commit and push evidence is complete.

### Validation

* All original assertions pass.
* All new assertions pass.
* All compatibility checks pass.
* All safety checks pass.
* No unresolved critical or high-severity issue remains.
* No unresolved validation failure remains.
* No undocumented workaround remains.

### Documentation

* Canonical documentation is complete.
* Engineering History is updated.
* Delivery Report is complete.
* Adoption guide is complete.
* Configuration reference is complete.
* Validation evidence is complete.
* Final state matches documentation.

### Final Acceptance

The task is complete only after:

* all mandatory verification has passed;
* the QWSG engineering workflow remains fully operational;
* the reusable framework has been demonstrated safely;
* all implementation evidence has been recorded;
* the final repository state is clean and understood;
* the Project Owner has reviewed the Delivery Report and formally accepted the implementation.

Only after successful acceptance may Task 015 be archived and the next engineering task be created.


## Owner Approval Requirements

Approved by Attila — Project Owner through the Engineering Task Builder on 2026-07-23 UTC.

The structured task definition has been explicitly approved for implementation. Further scope changes require explicit Project Owner approval.
