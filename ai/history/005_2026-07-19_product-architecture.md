# Task History 005: QWSG Product & System Blueprint

## Task Metadata

- Task ID: `005`
- Title: `QWSG Product & System Blueprint`
- Date: `2026-07-19` UTC
- Status: `complete with disclosed permission limitation`
- Responsible agent: `Aikó/Codex`
- Human authority: `Project Owner`
- Owner communication language: `Hungarian`
- Active prompt: `ai/prompts/005_CURRENT_TASK.md`
- Dependencies: completed Task 003 Product Definition, completed Task 004 job workflow, and owner-supplied comprehensive QWSG master plan

## Objective

Create one authoritative, implementation-neutral Product & System Blueprint that consolidates QWSG product intent, boundaries, component relationships, MVP, future direction, open questions, deferred decisions, and follow-up document sequence.

## Scope and Exclusions

Authorized work was limited to documentation, documentation inventory, a rollback snapshot, the primary blueprint, concise index and cross-reference updates, and this chronological record. No application, executable, dependency, service, system configuration, infrastructure, deployment, detailed interface, schema, framework, or implementation work was authorized or performed.

## Required Reading

Read and applied:

- `.agents/skills/qwsg-job/SKILL.md`
- `AGENTS.md`
- `ai/core/00_PROJECT_PHILOSOPHY.md`
- `ai/core/01_CONSTITUTION.md`
- `ai/core/03_AGENTS.md`
- `ai/core/05_SYSTEM_MAP.md`
- `ai/core/07_ENGINEERING_HISTORY.md`
- `ai/core/08_JOB_TEMPLATE.md`
- `ai/core/09_DELIVERY_POLICY.md`
- `ai/core/10_DOCUMENTATION_POLICY.md`
- active Prompt 005
- Task 003 prompt, history, and `docs/PRODUCT_DEFINITION.md`
- Task 004 prompt and history
- `ai/projects/QWSG/QWSG_MASTER_PLAN.md`
- relevant repository indexes and project records

## Starting State

The verified root was `<repository-root>`, branch `main`, at HEAD `8fa40acd945b5f0d5d1ee0c5e182a19bba092d2b`. `bin/job --check` passed with exactly one active prompt and Task ID `005`.

The working tree already contained uncommitted authorized results and rotations from Tasks 003 and 004, the active Prompt 005 and pending History 005, and the newly supplied master plan. These pre-existing changes were preserved. Candidate documentation directories were writable, setgid, and had default owner/group write ACLs. Existing document targets were mode `0660`; pending History 005 was pre-existing mode `0600`.

The exact comprehensive source required by the prompt was found at `ai/projects/QWSG/QWSG_MASTER_PLAN.md`. It contains the original host-problem narrative, state-transition alert model, full functional list, operating profiles, modular convenience examples, suggested implementation examples, three-component direction, and initial release proposal. No runtime or application target was needed.

## Snapshot Location and Verification

Snapshot: `ai/backups/20260719T032318Z_task005_product_system_blueprint/`.

It captures all five existing files that could be modified, the active prompt, the comprehensive master plan, initial environment and repository facts, a target manifest, SHA-256 checksums, and a root-guarded bounded restore procedure. All ten checksum entries passed. The sampled master-plan checksum matched `5edf596bd75a7a40d5bc2df65a317187a212eea4333d9d47a4d69920b9d5bc8a`. The restore script passed `bash -n` and is readable and executable.

## Risk Assessment

- Security and stability: low; only repository documentation changed and no secrets were copied.
- Data loss and rollback: low; exact existing targets were captured and the new document is explicitly bounded in rollback.
- Product direction: medium; mitigated by authority labels, traceability, open questions, and explicit implementation deferral.
- Scope creep: high; mitigated by treating early paths, commands, frameworks, thresholds, schemas, and estimates as illustrative or deferred.
- Permission and ACL: medium; all target directories were safely writable, but the pre-existing History 005 mode is `0600` and cannot be changed without separate owner approval.
- Localization: low; engineering artifacts are English and product UX requirements explicitly preserve localization and Hungarian-friendly usability.

## Planned and Performed Work

1. Validated the active task and recorded Git, repository, ownership, mode, and ACL state.
2. Located and read the comprehensive plan and all required governance, product, and history sources.
3. Created and verified the rollback snapshot before task-target edits.
4. Classified source content into settled principles, product requirements, preferred direction, illustrative examples, deferred decisions, and open questions.
5. Created `docs/PRODUCT_SYSTEM_BLUEPRINT.md` with all 42 required subject areas and a source-coverage appendix.
6. Registered the blueprint in the root README, System Map, Engineering History, and QWSG project record.
7. Recorded a bounded Task 006 Functional Specification recommendation without creating or activating a new prompt.
8. Performed content, link, placeholder, permission, checksum, Git-scope, and Markdown consistency verification.

## Files Changed

Created:

- `docs/PRODUCT_SYSTEM_BLUEPRINT.md`
- `ai/backups/20260719T032318Z_task005_product_system_blueprint/` snapshot records

Updated:

- `README.md`
- `ai/core/05_SYSTEM_MAP.md`
- `ai/core/07_ENGINEERING_HISTORY.md`
- `ai/projects/QWSG.md`
- `ai/history/005_2026-07-19_product-architecture.md`

The active prompt and comprehensive master plan were captured but not modified.

## Decisions

- Agent, Installer, and Console are product-level components with distinct responsibilities; process and deployment topology remain deferred.
- Agent-only operation is complete and useful; the Console is optional and cannot be a hidden root shell.
- The Installer owns transparent privileged lifecycle assistance and dependency consent.
- State transitions, escalation, recovery, and controlled emergency reminders are core behavior.
- The MVP is a coherent end-to-end protection slice; broad checks, channels, fleet features, remediation, and convenience extensions are post-MVP.
- SQLite is retained only as a preferred early candidate under an implementation-neutral local state-store requirement.
- Original convenience ideas such as `sl`, `cowsay`, and MOTD extensions are preserved as optional modularity examples, not core scope.
- Task 006 should specify observable functional behavior before internal architecture allocates responsibilities.

## Verification Evidence

- `bin/job --check`: passed for Task `005`.
- Snapshot `sha256sum -c SHA256SUMS`: all ten entries passed.
- Snapshot sampled checksum and readability: passed.
- Snapshot restore script `bash -n`: passed.
- Mandatory blueprint structure: all 42 numbered sections present.
- Product principles, component distinction, MVP, post-MVP, long-term direction, open questions, deferred decisions, glossary, and traceability: present.
- Placeholder scan: no unresolved template marker remains in Task 005 deliverables.
- Markdown whitespace validation with `git diff --check`: passed.
- Relative links added by Task 005 resolve.
- Secret-pattern and source review found no credential or private-key material copied into deliverables.
- All created/updated documentation except the pre-existing History 005 file is mode `0660`, owner `attila:nogroup`; snapshot restore is mode `0750`.
- No package, executable, application, dependency, service, infrastructure, production, active-prompt, or master-plan file was modified by Task 005.
- No commit or push was performed.

## Problems and Disclosed Limitation

The first two attempts stopped correctly because the required comprehensive planning source was absent. On the authorized resumption, `ai/projects/QWSG/QWSG_MASTER_PLAN.md` was present and complete enough to proceed.

History 005 existed before this task with mode `0600`, unlike the collaborative `0660` convention inherited by new documentation. The active prompt requires separate explicit owner approval for permission changes, so its mode was preserved. Content completion is unaffected, but group collaborators cannot read or edit this history until the owner authorizes normalization.

## Rollback Procedure

From exactly `<repository-root>`, run `ai/backups/20260719T032318Z_task005_product_system_blueprint/restore.sh`, review its bounded effects, and type `ROLLBACK-QWSG-TASK-005`. It removes only the Task 005-created blueprint and restores only README, System Map, Engineering History, the QWSG project record, and History 005 from verified copies. It preserves prompts, source planning material, application paths, prior task records, and unrelated working-tree changes.

## Open Questions

The blueprint records eight owner or architecture questions: initial platform versions and architectures; Console MVP sequencing; required e-mail methods; data retention; Console authentication; configuration and secrets handling; earliest safe remediation; and licensing/distribution/telemetry/commercial/hosted-service positions.

## Recommended Next Task

- Task ID: `006`
- Slug: `functional-specification`
- Title: `QWSG Functional Specification`
- Objective: create one authoritative, implementation-neutral specification of QWSG actors, operating profiles, MVP behavior, state transitions, alerts, lifecycle behavior, failure cases, and testable acceptance criteria.
- Rationale: observable required behavior must be precise before detailed architecture assigns components, interfaces, storage, privilege, and deployment mechanisms.
- Dependency: the completed Task 005 Product & System Blueprint.

No Task 006 prompt was created or activated.

## Git Record and Delivery Result

Starting and final HEAD remain `8fa40acd945b5f0d5d1ee0c5e182a19bba092d2b`; no commit was created. Task 005 is complete with the disclosed History 005 permission limitation. The deliverable is documentation-only, rollback-capable, and ready to serve as the authoritative input to a separately authorized Functional Specification milestone.
