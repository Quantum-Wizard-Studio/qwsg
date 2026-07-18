# Engineering Task Standard and Job Template

## Purpose

This document defines the mandatory lifecycle, communication rules, evidence requirements, and record structure for every Quantum Wizard Server Guardian engineering task.

## Status

Definitive and mandatory from `2026-07-18`. It does not authorize work by itself; task scope still requires human authority.

## Backward compatibility

The original job fields remain required with their original meanings: Objective and exclusions, Verified starting state, Snapshot location, Planned smallest safe change, Rollback procedure, Implementation record, Verification evidence, Documentation updates, and Unresolved issues and delivery result. Existing records remain valid. New records add the requirements below without rewriting earlier history.

## Mandatory task lifecycle

1. **Authority and scope:** identify the human-authorized objective, deliverables, exclusions, and stop conditions. Do not broaden scope by assumption.
2. **Read governing documents:** read the constitution, agent rules, relevant standards, policies, system records, and prior task history before changing the project.
3. **Inspect:** record the exact relevant environment, Git, ownership, permission, ACL, dependency, and file state. Stop and report material differences.
4. **Snapshot:** create and verify a rollback-capable snapshot before modifying task targets.
5. **Plan and rollback:** document what will change, the smallest safe method, risks, verification, and an exact bounded rollback procedure before implementation.
6. **Communicate:** provide concise Hungarian progress updates, decisions, warnings, and explanations to the owner. Never assume English fluency.
7. **Implement:** make only authorized changes. Engineering artifacts are English; user-facing content is localization-ready.
8. **Verify:** verify every meaningful change and the complete deliverable. Check scope, content, tests where applicable, security, permissions, ownership, ACLs, Git diff, and rollback validity in proportion to risk.
9. **Document:** update affected engineering documents and create a chronological English delivery report.
10. **Deliver:** report the outcome to the owner in Hungarian, including verification, rollback, unresolved issues, Git state, and the explicit completion status. Do not start the next task.

## Definitive job record

Every new task record must contain the following sections. “Not applicable” requires a short reason; required sections may not be silently omitted.

1. **Identity:** task number or identifier, English title, UTC date, responsible agent, and human authority.
2. **Objective and exclusions:** requested result, deliverables, forbidden work, and completion boundary.
3. **Governing documents read:** exact policies and prior records consulted.
4. **Verified starting state:** relevant Git state, environment, files, ownership, permissions, ACLs, dependencies, and detected variances.
5. **Snapshot location:** timestamped path, contents, integrity check, and retention decision.
6. **Planned smallest safe change:** target files or systems, sequence, assumptions requiring verification, risks, and non-goals.
7. **Rollback procedure:** exact scope, safety guards, confirmation requirement for destructive actions, and verification after rollback.
8. **Implementation record:** chronological changes, decisions, deviations, and why each change was necessary.
9. **Language and localization review:** confirmation that engineering artifacts are English, owner communication is Hungarian, and applicable user-facing content is localization-ready.
10. **Verification evidence:** commands or methods, expected and actual results, permissions and ownership where relevant, tests, security checks, Git diff, and confirmation of excluded work.
11. **Documentation updates:** every updated policy, reference, changelog, and history record.
12. **Unresolved issues and delivery result:** remaining risks or blockers, exact status (`complete`, `complete with disclosed limitations`, or `blocked`), and why.
13. **Git record:** branch, commit subject and full hash when committed, plus final working-tree status.
14. **Owner delivery:** concise Hungarian summary covering outcome, verification, rollback, unresolved issues, and whether the task is safe and complete.

## Language policy

- English is mandatory for source code, comments, engineering documentation, architecture, API specifications, commit messages, changelogs, issue records, and engineering reports.
- Hungarian is mandatory for every message addressed to the project owner: progress reports, explanations, recommendations, summaries, warnings, and completion reports.
- Owner communication must explain technical decisions in Hungarian and must not assume English proficiency. Precise established engineering terminology may be retained and explained.

## Localization readiness

The Web Console, Installer, and future end-user documentation must support multiple languages by design. Requirements and reviews must distinguish engineering text from user-facing text. User-facing strings, formats, and content structures must allow translation, locale-specific formatting, and language expansion without redesign. Engineering documentation remains English.

## Completion gate

A task is complete only when authorized deliverables exist, exclusions were respected, verification passed, rollback remains usable, documentation and chronological history are updated, unresolved issues are disclosed, and the Hungarian owner delivery is provided. A task never authorizes the next milestone.

Task facts, decisions, and evidence belong in completed job records. Secrets, credentials, assumptions presented as facts, unrelated work, and undocumented deviations do not. This definitive standard will evolve during development through explicit human approval while preserving compatibility with earlier records.
