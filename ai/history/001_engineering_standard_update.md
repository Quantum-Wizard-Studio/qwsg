# Delivery Report 001: Engineering Task Standard and Language Policy

## Identity

- Date: `2026-07-18` UTC
- Responsible agent: Aikó/Codex
- Human authority: project owner request titled “QWSG — Engineering Standard Update”
- Status: complete and verified

## Objective and exclusions

Refine the engineering workflow, make `ai/core/08_JOB_TEMPLATE.md` the definitive backward-compatible task standard, establish English engineering artifacts and Hungarian owner communication, and require localization-ready user-facing design. Application code, software installation, server changes, and the next milestone were excluded.

## Governing documents read

The project philosophy, constitution, agent rules, engineering standards, existing job template, documentation policy, engineering history, and bootstrap delivery record governed this work.

## Verified starting state

Git was clean on `main` at `ac164d41bd8ef951cf3a4eecc307dd834b5d750d`. Relevant paths were owned by `<repository-owner>:<repository-group>`, directories retained setgid, and default ACLs granted owner and group write inheritance.

## Snapshot location

`ai/backups/20260718T211755Z_engineering_standard_update/` records the starting state, permissions, Git status, planned targets, and guarded rollback procedure.

## Planned smallest safe change

Only the constitution, agent rules, engineering standards, engineering history, job standard, documentation policy, this report, and snapshot records were changed. Existing task-template fields were retained and expanded rather than invalidated.

## Rollback procedure

From the exact project root, run the snapshot's `restore.sh`, review its bounded targets, and type the required confirmation phrase. It restores the six changed core documents from the verified baseline and removes only this history record; the snapshot remains available.

## Implementation record

- Added permanent language and localization rules to the constitution.
- Added owner-language and task-standard duties to agent rules.
- Linked engineering standards to the definitive task workflow and localization requirements.
- Expanded the original job template into a mandatory lifecycle, record schema, completion gate, and owner-delivery standard while retaining all original fields.
- Clarified engineering, owner-communication, and end-user documentation language boundaries.
- Added the milestone to the chronological engineering history.

## Language and localization review

All repository engineering documentation remains English. Owner-facing progress and delivery communication is Hungarian. The standard explicitly requires localization-ready Web Console, Installer, and future end-user documentation without changing the English engineering baseline.

## Verification evidence

Verification covers required policy phrases, cross-document consistency, preservation of original job fields, Markdown structure, rollback script syntax, Git diff checks, ownership, permissions, ACLs, excluded artifacts, and final Git status.

## Documentation updates

Updated `01_CONSTITUTION.md`, `03_AGENTS.md`, `06_ENGINEERING_STANDARDS.md`, `07_ENGINEERING_HISTORY.md`, `08_JOB_TEMPLATE.md`, and `10_DOCUMENTATION_POLICY.md`; created this report and its task snapshot.

## Unresolved issues and delivery result

No known policy conflict is unresolved. Delivery result: **complete**. Verification passed and the documentation-standard commit was created.

## Git record

Branch: `main`. Engineering-standard commit: `cd334dd3e6f177ecf968fac24b7b0bf46ab5cfc0` (`docs: establish definitive engineering task standard`). A follow-up audit commit records this immutable hash. Final working-tree status: clean after the audit commit.

## Owner delivery

The final owner-facing outcome, verification, rollback information, and completion state are communicated in Hungarian after the Git record is complete.

This report belongs to chronological engineering history. Secrets, application design, and future milestone work do not. The policy documentation will evolve during development through explicit human approval.
