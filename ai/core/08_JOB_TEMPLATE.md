# Official Engineering Task Standard

## Purpose

This document is the official, reusable, AI-friendly standard for defining, executing, verifying, documenting, and delivering every Quantum Wizard Server Guardian engineering task.

## Status

Definitive and mandatory from `2026-07-18`, revised for Framework 2.0.0. It does not authorize work by itself; task scope still requires human authority.

## Backward compatibility

The original job fields remain required with their original meanings: Objective and exclusions, Verified starting state, Snapshot location, Planned smallest safe change, Rollback procedure, Implementation record, Verification evidence, Documentation updates, and Unresolved issues and delivery result. Existing records remain valid. New records add the requirements below without rewriting earlier history.

## Mandatory task lifecycle

1. **Authority and scope:** identify the human-authorized objective, Authority
   Envelope, deliverables, exclusions, and boundaries. Builder approval grants
   Standard Execution Authority from `17_EXECUTION_MODEL.md`; it does not
   authorize reserved work or broaden scope by assumption.
2. **Read governing documents:** read the constitution, agent rules, relevant standards, policies, system records, and prior task history before changing the project. Validate the versioned project configuration and canonical Git policy; never source configuration or treat project-overridable data as authority to weaken mandatory safeguards.
3. **Inspect:** record the exact relevant environment, Git, ownership,
   permission, ACL, dependency, and file state. Diagnose and correct recoverable
   in-scope differences; stop on a material authority, safety, privacy, security,
   destructive, external-system, or rollback difference.
4. **Snapshot:** create and verify a rollback-capable snapshot before modifying task targets.
5. **Plan and rollback:** document what will change, the smallest safe method, risks, verification, and an exact bounded rollback procedure before implementation.
6. **Communicate:** use the verified preferred language of the current project owner or lead developer for concise progress updates, decisions, warnings, and explanations. Do not hardcode or assume a human language.
7. **Implement:** make only authorized changes. Engineering artifacts are English; user-facing content is localization-ready.
8. **Verify:** verify every meaningful change and the complete deliverable.
   Check scope, content, tests where applicable, security, permissions,
   ownership, ACLs, Git diff, and rollback validity in proportion to risk.
   Recoverable failures follow diagnose -> correct -> retest -> continue while
   they remain inside the Authority Envelope.
9. **Document:** update affected engineering documents and create a chronological English delivery report.
10. **Deliver and hand off:** report the outcome in the current owner or lead developer's preferred language, including verification, rollback, unresolved issues, Git state, and explicit completion status. When required by the active task, prepare the next unapproved task under `11_ENGINEERING_LIFECYCLE.md`; never start it.

## Required task structure

Every new task record must contain the following sections. “Not applicable” requires a short reason; required sections may not be silently omitted.

### 1. Task Metadata

Task identifier, English title, UTC date, status, responsible agent, human authority, preferred owner or lead-developer communication language, related prompt path, and dependencies.

### 2. Objective

The single intended outcome, deliverables, acceptance intent, and completion boundary.

### 3. Scope

Exact files, components, systems, people, and actions authorized for the task.

### 4. Out of Scope

Explicit exclusions, forbidden actions, deferred work, and the next milestone that must not begin.

### 5. Authority Envelope

Framework 2.0 tasks use a concise envelope with four categories:

1. **Task targets and boundaries:** the objective, authorized components, and
   exclusions that bound Standard Execution Authority.
2. **Permitted external actions:** external systems, people, privilege, or
   physical actions explicitly authorized; `none` is valid.
3. **Owner-reserved decisions:** task-specific decisions retained by the Owner.
4. **Task-specific STOP conditions:** additional boundaries beyond the mandatory
   STOP semantics in `17_EXECUTION_MODEL.md`; `none beyond the standard` is
   valid.

Snapshot, rollback, ordinary local edits, diagnose/fix/retest iterations,
proportional validation, task-scoped documentation, safe targeted Git
integration, and lifecycle closure are inherited. Ambiguity never expands the
task's targets or external authority. Legacy nine-category Framework 1.1
envelopes remain valid and are not rewritten.

### 6. Required Reading

Always include `00_PROJECT_PHILOSOPHY.md`, `01_CONSTITUTION.md`, `03_AGENTS.md`, and `08_JOB_TEMPLATE.md` before implementation work. Add relevant policies, system records, prompts, prior reports, and technical references. Record what was actually read.

### 7. Environment Verification

Record relevant date, user, working directory, Git branch and status, file state, versions, ownership, permissions, ACLs, dependencies, services, and material differences from expectations. Verify facts rather than assuming them.

### 8. Snapshot Requirements

Define the timestamped snapshot path, exact captured state, integrity checks, retention, and a guarded restore procedure. Create and verify the snapshot before modifying task targets.

### 9. Risk Assessment

List security, stability, data-loss, compatibility, localization, permission, operational, and rollback risks. Rate each relevant risk and document its mitigation or acceptance authority.

### 10. Planned Work

Describe the smallest safe sequence, target paths or systems, decision points, assumptions requiring verification, expected outputs, and non-goals.

### 11. Rollback Plan

Define exact bounded targets, preconditions, confirmation for destructive steps, commands or procedures, retained evidence, and post-rollback verification. Never rely on unsafe wildcard deletion.

### 12. Verification Checklist

Provide task-specific checks with expected results. Include scope, content, tests where applicable, security, localization, ownership, permissions, ACLs, excluded work, documentation consistency, rollback validity, Git diff, and final status.

### 13. Documentation Updates

List every core document, prompt, reference, changelog, and chronological history record that must change. Record all actual modifications and justified deviations.

### 14. Delivery Report

Create an English engineering report containing the implementation record, reasoning, verification evidence, rollback information, unresolved issues, recommendations, Git record, and delivery result. Communicate the owner-facing summary separately in the verified preferred language.

### 15. Completion Criteria

Define objective pass/fail conditions. A valid result is `complete`, `complete with disclosed limitations`, or `blocked`; the reason and unresolved items must be explicit.

## Language policy

- English is mandatory for source code, comments, engineering documentation, architecture, API specifications, commit messages, changelogs, issue records, and engineering reports.
- Every message addressed to the current project owner or lead developer uses that person's configured preferred language: progress reports, explanations, recommendations, summaries, warnings, and completion reports.
- The preference must be verified from approved project context or requested when unknown. This document defines the policy concept only; it does not prescribe a configuration-file format or a permanent human language.
- Owner communication must explain technical decisions in the preferred language without assuming English proficiency. Precise established engineering terminology may be retained and explained.

## Localization readiness

The Web Console, Installer, and future end-user documentation must support multiple languages by design. Requirements and reviews must distinguish engineering text from user-facing text. User-visible strings must not be hardcoded whenever localization is technically feasible. Strings, formats, and content structures must allow translation, locale-specific formatting, and language expansion without redesign. Engineering documentation remains English.

## Prompt workflow

The authoritative state-transition, completion-gate, transactional preparation, and No Task Without History rules are defined in `11_ENGINEERING_LIFECYCLE.md`.

Every current engineering task has one active English prompt named `NNN_CURRENT_TASK.md` and one independent history file. When the latest task is complete and archived and no next task is authorized, `ai/prompts/` is empty; this is the canonical idle state, not a missing record. The prompt follows the required task structure above, uses the semantic states `draft`, `approved`, `active`, `complete`, `superseded`, or `archived without execution`, and separates unapproved preparation from approved execution. Prior prompts move to `ai/archive_prompts/`; archived prompts and histories remain committed. Draft preparation or prompt archiving never grants execution authority; Builder approval grants Standard Execution Authority bounded by the recorded Authority Envelope. Detailed numbering, naming, rotation, and compatibility rules are maintained in `14_PROMPT_WORKFLOW.md`.

The general Engineering History remains a concise milestone index. It must not become an infinitely growing task log; detailed evidence belongs in the independent task history record.

The official creation path is `ai/scripts/task-builder.sh`. It collects or reads structured owner-authored fields, generates metadata and approval text, validates the complete documents, and installs the prompt/history pair transactionally. Its deterministic input-directory mode keeps every field in a separate text file so content is read as data and multi-line values remain lossless. The older `next-task.sh` draft-preparation path remains supported when approval must occur in a later, separate review step; its generated placeholders must never be manually assembled into an executable task without the required owner review and validation.

Project identity, canonical remote and branch, communication and documentation
languages, required reading, lifecycle directories, and project-specific
validation argv are declared in `ai/config/engineering-project.conf` and its
referenced validation file. `ai/scripts/framework-check.sh` validates them.
Explicit approval, one-active-task enforcement, snapshots, rollback, targeted
staging, history, and completion evidence are mandatory core rules and cannot be
disabled by project configuration.

Failure classification, proportional validation, development/release
separation, and evidence reuse follow `17_EXECUTION_MODEL.md`. External
multi-check diagnostics should use `18_BOUNDED_DIAGNOSTIC_RUNNER.md` when
practical.

## Completion gate

A task is complete only when authorized deliverables exist, exclusions were respected, verification passed, rollback remains usable, documentation and chronological history are updated, unresolved issues are disclosed, and the owner delivery is provided in the verified preferred language. A task never authorizes the next milestone.

Task facts, decisions, and evidence belong in completed job records. Secrets, credentials, assumptions presented as facts, unrelated work, and undocumented deviations do not. This definitive standard will evolve during development through explicit human approval while preserving compatibility with earlier records.
