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

Future task prompts are stored under `ai/prompts/`, with one prompt per engineering task. Prompt structure, status, execution boundary, and naming rules are defined in `ai/prompts/README.md`. Creating a prompt does not authorize or execute its task.

Approved coding, testing, review, localization, and compatibility rules belong here. Product requirements, secrets, and one-off task notes do not. Standards will evolve during development as the technology stack is approved.
