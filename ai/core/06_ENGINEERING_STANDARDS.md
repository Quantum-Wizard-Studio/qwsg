# Engineering Standards

## Purpose

This document defines implementation quality, review, testing, compatibility, and operational safety expectations.

## Status

Active standard: follow the constitution and `08_JOB_TEMPLATE.md`, prefer small reversible changes, and verify proportionately to risk.

## Language and localization

- All engineering artifacts are written in English: source code, comments, documentation, architecture, API specifications, commit messages, changelogs, issue records, and engineering reports.
- All communication intended for the project owner is written in Hungarian, including progress, explanations, recommendations, summaries, warnings, and completion messages.
- Technical decisions are explained in Hungarian without assuming English fluency, while established engineering terminology remains precise.
- The Web Console, Installer, and future end-user documentation are designed for localization from their first approved design. User-facing strings must not be inseparably embedded in implementation logic.
- Engineering documentation remains English regardless of the languages offered to users.

## Workflow authority

`08_JOB_TEMPLATE.md` is the definitive engineering task standard. Its required sequence and record fields apply to documentation, code, infrastructure, release, and operational tasks unless a stricter approved policy applies. Deviations require explicit human approval and must be recorded.

Approved coding, testing, review, localization, and compatibility rules belong here. Product requirements, secrets, and one-off task notes do not. Standards will evolve during development as the technology stack is approved.
