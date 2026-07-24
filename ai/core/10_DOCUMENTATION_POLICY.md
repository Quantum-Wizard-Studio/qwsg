# Documentation Policy

## Purpose

This policy keeps Quantum Wizard Server Guardian understandable and maintainable.

## Status

Active permanent language and documentation baseline.

## Language boundaries

- Engineering documentation is written in English. This includes architecture, API specifications, changelogs, issue records, and engineering reports, as well as code-facing documentation and comments.
- Messages intended for the current project owner or lead developer use that person's configured preferred communication language. Progress reports, explanations, recommendations, summaries, warnings, and completion reports are owner communication, not engineering artifacts.
- The preferred language must be verified from validated project configuration
  or requested rather than permanently assumed.
- Technical decisions must be explained in the preferred language without assuming English fluency, while engineering terminology remains accurate.
- End-user documentation is designed for localization and may be published in multiple languages. Its source and translation workflow must keep engineering governance in English.

Accurate purpose, status, scope, decisions, procedures, language ownership, localization requirements, and verification evidence belong in documentation. Secrets, stale claims, unexplained generated output, and personal sensitive data do not. Documentation must change with completed work and will evolve throughout development.

Each engineering task must have its own zero-padded, date-and-slug history file. Previous prompts belong in `ai/archive_prompts/`; prompt archives and task histories are permanent Git-tracked records. General history documents may index milestones but must not accumulate detailed records for every task. Existing valid historical files remain unchanged where backward compatibility requires it. Generated binary backup archives may remain ignored, but engineering Markdown records must not be ignored.
