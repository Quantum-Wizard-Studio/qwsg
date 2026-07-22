# Engineering Standards

## Purpose

This document defines implementation quality, review, testing, compatibility, and operational safety expectations.

## Status

Active standard: follow the constitution and `08_JOB_TEMPLATE.md`, prefer small reversible changes, and verify proportionately to risk.

## Language and localization

- All engineering artifacts are written in English: source code, comments, documentation, architecture, API specifications, commit messages, changelogs, issue records, and engineering reports.
- All communication intended for the current project owner or lead developer uses that person's configured preferred language, including progress, explanations, recommendations, summaries, warnings, and completion messages. The preference is a project policy concept; no configuration file is defined yet.
- Technical decisions are explained in the preferred language without assuming proficiency in English, while established engineering terminology remains precise.
- The Web Console, Installer, and future end-user documentation are designed for localization from their first approved design. User-visible strings must not be hardcoded whenever localization is technically feasible.
- Engineering documentation remains English regardless of the languages offered to users.

## Workflow authority

`08_JOB_TEMPLATE.md` is the definitive engineering task standard. Its required sequence and record fields apply to documentation, code, infrastructure, release, and operational tasks unless a stricter approved policy applies. Deviations require explicit human approval and must be recorded.

At most one active task prompt is stored under `ai/prompts/` as `NNN_CURRENT_TASK.md`. The directory is empty in the valid idle state after the latest completed task is archived and before a separately authorized next task is created. Prior prompts are retained under `ai/archive_prompts/`, and every task receives a separate history file under `ai/history/`. Sequential numbering, lifecycle states, rotation safety, and naming rules are defined in `14_PROMPT_WORKFLOW.md`. Creating or archiving a prompt does not authorize or prove execution.

Approved coding, testing, review, localization, and compatibility rules belong here. Product requirements, secrets, and one-off task notes do not. Standards will evolve during development as the technology stack is approved.
