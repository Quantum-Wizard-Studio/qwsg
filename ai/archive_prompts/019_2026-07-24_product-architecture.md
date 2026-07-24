# Current Engineering Task 019: Canonical QWSG Product Architecture

## Task Metadata

- Task ID: `019`
- Task slug: `product-architecture`
- Status: `complete`
- Date opened: `2026-07-24` UTC
- Human authority: Project Owner through the Task 019 Aiko Engineering Prompt
- Owner or lead-developer communication language: Hungarian

## Title

Canonical QWSG Product Architecture


## Objective

Create the canonical, engineering-grade QWSG Product Architecture that defines the long-term product ecosystem and becomes the primary architectural reference for future implementation tasks.

Preserve this immutable project philosophy verbatim:

> The Community Edition exists to earn trust. The Professional Edition exists to save time. The Provider Edition exists to operate at scale. Every edition shares the same deterministic engineering core.

Refine, reorganize, and expand the supplied architecture draft as a Chief System Architect, prioritizing long-term consistency, extensibility, engineering quality, and conservative architectural decisions.


## Scope

- Review the owner-supplied architecture draft in `current-task-job.txt` as source material.
- Create or establish one canonical Product Architecture document.
- Define Product Vision, Engineering Philosophy, Community Edition, Professional Edition, Provider Edition, edition comparison, Product Identity, Workspace Architecture, Terminal Experience, future Terminal UI philosophy, Web Dashboard vision, Licensing philosophy, Privacy model, Deployment models, Automation philosophy, AI separation, Future ecosystem, long-term roadmap structure, and Engineering Principles.
- Make Community explicitly a complete professional Linux engineering toolkit; Professional an automation, convenience, and scale extension; and Provider an operational-capability extension.
- Define deterministic core, offline capability, privacy-first design, AI independence, reproducible behavior, and engineering-before-automation as permanent principles.
- Update affected canonical architecture, system-map, roadmap, README, and milestone-history documentation where appropriate.
- Maintain documentation consistency and lifecycle evidence.


## Out of Scope

- No production source code, tests, binaries, runtime configuration, dependencies, infrastructure, or behavior changes.
- Do not implement the daemon, scheduler, Drift Engine, Health Engine, alert engine, email notifications, Web Dashboard, licensing enforcement, remote agents, fleet management, or Provider services.
- Do not make Community a crippled or artificially restricted edition.
- Do not assign Professional any improvement to engineering correctness, inventory quality, deterministic analysis, or core evidence.
- Do not begin any future implementation phase or authorize Task 020.
- Do not modify unrelated owner-owned untracked files.


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

- Verify the exact QWSG repository root, framework identity, primary branch, canonical remote, HEAD, complete Git status, and local-to-remote relationship.
- Verify the Foundation Phase is complete through canonical architecture, roadmap, and Task 018 lifecycle records.
- Verify the repository is in the canonical idle state with Task 018 archived and complete and no active prompt before Task 019 creation.
- Verify `current-task-job.txt` is a readable owner-supplied architecture draft and preserve its opening philosophy verbatim.
- Inventory existing product, architecture, system-map, roadmap, README, licensing, privacy, deployment, UI, automation, and AI statements before editing.
- Stop on material environmental, lifecycle, authority, or scope differences.


## Snapshot Requirements

Before modifying task targets, create a UTC-stamped rollback-capable snapshot outside Git under `/tmp`.

Capture the repository Git baseline and exact contents of every tracked documentation or lifecycle target to be modified. Record verified absence for any new target. Include the owner-supplied draft as read-only source evidence without modifying it. Produce SHA-256 integrity evidence and verify archive readability.

Retain the snapshot through owner acceptance. Restore only the explicit Task 019 target paths after collision review; never extract over a live worktree or use broad reset, checkout, clean, or wildcard deletion. Re-run framework, lifecycle, documentation, and Git-state checks after any restore.


## Risk Assessment

- Architectural ambiguity or premature implementation commitment: medium; separate durable product boundaries from non-binding examples and explicitly label future capabilities.
- Edition degradation or paywalling engineering correctness: high; define one shared deterministic engineering core and invariant Community completeness.
- Contradiction with canonical contracts and completed Foundation architecture: high; inspect and cross-reference all relevant canonical documents before writing.
- Privacy, licensing, and deployment overclaim: medium; specify principles and trust boundaries without claiming unimplemented mechanisms.
- Roadmap instability: medium; organize by capability streams and architectural gates rather than fragile dates.
- Documentation drift or duplicate authority: medium; establish one primary Product Architecture and make other documents concise references.
- Accidental production or behavior change: low but prohibited; restrict the changed path set to documentation and lifecycle records and verify the final diff.
- Rollback/data-loss risk: low; use a bounded verified external snapshot and exact restore procedure.


## Planned Work

1. Validate the framework, lifecycle, Git baseline, owner draft, and required reading.
2. Inventory all existing product and architecture claims and resolve authority boundaries.
3. Create and verify a bounded external snapshot of every intended target before editing.
4. Design the canonical Product Architecture around a shared deterministic core and explicit edition capability boundaries.
5. Define product identity, user workspace, terminal and future TUI experience, Web Dashboard vision, licensing, privacy, deployment, automation, AI separation, ecosystem boundaries, and architectural governance.
6. Reorganize the long-term roadmap into durable capability streams and gates without scheduling or implementing future features.
7. Update concise cross-references in affected canonical documents and the milestone index only where they materially improve consistency.
8. Update Task 019 history with decisions, exact changes, verification evidence, rollback, limitations, and delivery result.
9. Run every configured framework validation plus documentation, scope, permission, secret, lifecycle, and Git-diff review.
10. Complete and archive Task 019 into the canonical idle state after all gates pass; do not create a successor.


## Rollback Plan

Use the verified external Task 019 snapshot and exact target manifest. Before restoring, confirm the repository root, expected Task 019 target list, archive checksum, archive readability, and absence of unrelated overlapping changes.

Restore only previously existing Task 019 documentation and lifecycle targets from the snapshot. Remove a newly created Product Architecture file only if its verified pre-task absence is recorded and Project Owner authority covers Task 019 rollback. Never overwrite unrelated work or use broad Git reset, checkout, restore, clean, or wildcard deletion.

After restoration, verify file hashes where recorded, `ai/scripts/framework-check.sh`, `bin/job --check`, configured validation commands, documentation consistency, permissions, and `git status --short --branch`. A destructive restore is documented but not executed merely as a test.


## Deliverables

- A polished canonical QWSG Product Architecture suitable as the primary architectural reference for future implementation tasks.
- Explicit edition definitions and comparison preserving Community completeness, Professional automation, and Provider operations.
- Canonical product identity, workspace, terminal/TUI, Web Dashboard, licensing, privacy, deployment, automation, AI separation, ecosystem, roadmap, and engineering-principle decisions.
- Updated roadmap and affected canonical cross-references where appropriate.
- Updated concise engineering milestone index if the new canonical architecture is milestone-worthy.
- Complete Task 019 lifecycle and delivery history with snapshot, decisions, verification, rollback, Git state, and limitations.
- Canonical idle lifecycle state after successful completion, with no successor task created.


## Verification

- Confirm the immutable opening philosophy appears verbatim in the canonical Product Architecture.
- Confirm every required topic is defined and internally consistent.
- Confirm Community is explicitly complete, Professional improves automation but never engineering correctness, and Provider adds operational scale.
- Confirm deterministic core, offline capability, privacy-first design, AI independence, reproducibility, and engineering-before-automation are permanent cross-edition principles.
- Confirm all existing canonical contracts and Foundation architecture remain authoritative and uncontradicted.
- Confirm roadmap and cross-references point to the new Product Architecture without duplicating or weakening it.
- Run `ai/scripts/framework-check.sh`.
- Run every command configured in `ai/config/engineering-validations.tsv`.
- Run lifecycle consistency checks required by the active prompt and completion workflow.
- Review `git diff --check`, changed path lists, file modes, ownership, permissions, ACLs where relevant, documentation links, private-host/secret evidence, and complete Git status.
- Confirm no production code, tests, binaries, dependencies, runtime configuration, or behavior changed and unrelated untracked files remain untouched.
- Verify the snapshot checksum, readability, bounded restore instructions, and retention statement.


## Documentation Updates

- Create the canonical Product Architecture at the repository's established architecture-document location.
- Update `ai/core/13_ROADMAP.md` to reference and reflect the product architecture's durable capability streams and gates.
- Update `ai/core/04_ARCHITECTURE.md`, `ai/core/05_SYSTEM_MAP.md`, and `README.md` only where concise canonical references or product identity alignment are appropriate.
- Update `ai/core/07_ENGINEERING_HISTORY.md` with a concise Task 019 milestone entry if warranted.
- Update `ai/history/019_2026-07-24_product-architecture.md` throughout execution and delivery.
- Update `ai/prompts/019_CURRENT_TASK.md` status only through the governed completion workflow, then archive it under `ai/archive_prompts/`.
- Do not rewrite completed historical records or duplicate the Product Architecture in secondary documents.


## Completion Criteria

- One canonical Product Architecture exists and can govern future implementation without redefining product philosophy.
- The immutable philosophy and all required topics are complete, explicit, technically coherent, and consistent with canonical contracts.
- Edition boundaries preserve a fully professional Community toolkit, automation-focused Professional capabilities, and operations-focused Provider capabilities.
- Roadmap and affected canonical references are consistent with the Product Architecture.
- No production source code or behavior changed.
- All framework, lifecycle, documentation, scope, security/privacy, permissions, rollback, and Git-diff validations pass.
- Task history truthfully records decisions, exact evidence, rollback, unresolved items, and final Git state.
- Task 019 prompt and history reach `complete`, the prompt is archived, and the repository returns to canonical idle state without creating Task 020.


## Owner Approval Requirements

Approved by Project Owner through the Task 019 Aiko Engineering Prompt through the Engineering Task Builder on 2026-07-24 UTC.

The structured task definition has been explicitly approved for implementation. Further scope changes require explicit Project Owner approval.
