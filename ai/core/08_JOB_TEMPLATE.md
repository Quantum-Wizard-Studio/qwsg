# Official Engineering Task Standard

## Purpose

This document is the official, reusable, AI-friendly standard for defining, executing, verifying, documenting, and delivering every Quantum Wizard Server Guardian engineering task.

## Status

Definitive and mandatory from `2026-07-18`, refined by Engineering Update E001. It does not authorize work by itself; task scope still requires human authority.

## Backward compatibility

The original job fields remain required with their original meanings: Objective and exclusions, Verified starting state, Snapshot location, Planned smallest safe change, Rollback procedure, Implementation record, Verification evidence, Documentation updates, and Unresolved issues and delivery result. Existing records remain valid. New records add the requirements below without rewriting earlier history.

## Mandatory task lifecycle

1. **Authority and scope:** identify the human-authorized objective, deliverables, exclusions, and stop conditions. Do not broaden scope by assumption.
2. **Read governing documents:** read the constitution, agent rules, relevant standards, policies, system records, and prior task history before changing the project.
3. **Inspect:** record the exact relevant environment, Git, ownership, permission, ACL, dependency, and file state. Stop and report material differences.
4. **Snapshot:** create and verify a rollback-capable snapshot before modifying task targets.
5. **Plan and rollback:** document what will change, the smallest safe method, risks, verification, and an exact bounded rollback procedure before implementation.
6. **Communicate:** use the verified preferred language of the current project owner or lead developer for concise progress updates, decisions, warnings, and explanations. Do not hardcode or assume a human language.
7. **Implement:** make only authorized changes. Engineering artifacts are English; user-facing content is localization-ready.
8. **Verify:** verify every meaningful change and the complete deliverable. Check scope, content, tests where applicable, security, permissions, ownership, ACLs, Git diff, and rollback validity in proportion to risk.
9. **Document:** update affected engineering documents and create a chronological English delivery report.
10. **Deliver:** report the outcome in the current owner or lead developer's preferred language, including verification, rollback, unresolved issues, Git state, and the explicit completion status. Do not start the next task.

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

### 5. Required Reading

Always include `00_PROJECT_PHILOSOPHY.md`, `01_CONSTITUTION.md`, `03_AGENTS.md`, and `08_JOB_TEMPLATE.md` before implementation work. Add relevant policies, system records, prompts, prior reports, and technical references. Record what was actually read.

### 6. Environment Verification

Record relevant date, user, working directory, Git branch and status, file state, versions, ownership, permissions, ACLs, dependencies, services, and material differences from expectations. Verify facts rather than assuming them.

### 7. Snapshot Requirements

Define the timestamped snapshot path, exact captured state, integrity checks, retention, and a guarded restore procedure. Create and verify the snapshot before modifying task targets.

### 8. Risk Assessment

List security, stability, data-loss, compatibility, localization, permission, operational, and rollback risks. Rate each relevant risk and document its mitigation or acceptance authority.

### 9. Planned Work

Describe the smallest safe sequence, target paths or systems, decision points, assumptions requiring verification, expected outputs, and non-goals.

### 10. Rollback Plan

Define exact bounded targets, preconditions, confirmation for destructive steps, commands or procedures, retained evidence, and post-rollback verification. Never rely on unsafe wildcard deletion.

### 11. Verification Checklist

Provide task-specific checks with expected results. Include scope, content, tests where applicable, security, localization, ownership, permissions, ACLs, excluded work, documentation consistency, rollback validity, Git diff, and final status.

### 12. Documentation Updates

List every core document, prompt, reference, changelog, and chronological history record that must change. Record all actual modifications and justified deviations.

### 13. Delivery Report

Create an English engineering report containing the implementation record, reasoning, verification evidence, rollback information, unresolved issues, recommendations, Git record, and delivery result. Communicate the owner-facing summary separately in the verified preferred language.

### 14. Completion Criteria

Define objective pass/fail conditions. A valid result is `complete`, `complete with disclosed limitations`, or `blocked`; the reason and unresolved items must be explicit.

## Language policy

- English is mandatory for source code, comments, engineering documentation, architecture, API specifications, commit messages, changelogs, issue records, and engineering reports.
- Every message addressed to the current project owner or lead developer uses that person's configured preferred language: progress reports, explanations, recommendations, summaries, warnings, and completion reports.
- The preference must be verified from approved project context or requested when unknown. This document defines the policy concept only; it does not prescribe a configuration-file format or a permanent human language.
- Owner communication must explain technical decisions in the preferred language without assuming English proficiency. Precise established engineering terminology may be retained and explained.

## Localization readiness

The Web Console, Installer, and future end-user documentation must support multiple languages by design. Requirements and reviews must distinguish engineering text from user-facing text. User-visible strings must not be hardcoded whenever localization is technically feasible. Strings, formats, and content structures must allow translation, locale-specific formatting, and language expansion without redesign. Engineering documentation remains English.

## Prompt workflow

Every engineering task has exactly one active English prompt named `NNN_CURRENT_TASK.md` and one independent history file. The prompt follows the required task structure above, uses the semantic states `draft`, `approved`, `active`, `complete`, `superseded`, or `archived without execution`, and separates definition from execution. Prior prompts move to `ai/archive_prompts/`; archived prompts and histories remain committed. Prompt creation or archiving never grants execution authority. Detailed numbering, naming, rotation, and compatibility rules are maintained in `14_PROMPT_WORKFLOW.md`.

The general Engineering History remains a concise milestone index. It must not become an infinitely growing task log; detailed evidence belongs in the independent task history record.

## Completion gate

A task is complete only when authorized deliverables exist, exclusions were respected, verification passed, rollback remains usable, documentation and chronological history are updated, unresolved issues are disclosed, and the owner delivery is provided in the verified preferred language. A task never authorizes the next milestone.

Task facts, decisions, and evidence belong in completed job records. Secrets, credentials, assumptions presented as facts, unrelated work, and undocumented deviations do not. This definitive standard will evolve during development through explicit human approval while preserving compatibility with earlier records.
