# Engineering Prompt Workflow

## Purpose

This document defines how reusable Quantum Wizard Server Guardian engineering task prompts are created, approved, executed, and retained.

## Status

Active governance established by Engineering Update E001. It documents policy only and does not implement prompt automation or configuration files.

## One prompt, one task

Every future engineering task should have one English Markdown prompt under `ai/prompts/`. Use the sortable form `NNN_UPPER_SNAKE_CASE_TITLE.md` unless an approved task identifier requires another stable prefix. A prompt defines one bounded task; unrelated work and multiple milestones do not belong in one file.

## Required structure

Each prompt follows `ai/core/08_JOB_TEMPLATE.md` and contains Task Metadata, Objective, Scope, Out of Scope, Required Reading, Environment Verification, Snapshot Requirements, Risk Assessment, Planned Work, Rollback Plan, Verification Checklist, Documentation Updates, Delivery Report, and Completion Criteria.

Task Metadata declares a lifecycle status: `draft`, `approved`, `active`, `complete`, or `superseded`. It also records human authority and the preferred communication language when verified. The language preference is a configurable policy concept; no configuration file is defined yet.

## Lifecycle

1. Create a `draft` prompt without executing it.
2. Review scope, exclusions, risks, rollback, verification, language, and localization requirements.
3. Obtain explicit human authority before marking it `approved` or beginning work.
4. At execution, read `00_PROJECT_PHILOSOPHY.md`, `01_CONSTITUTION.md`, `03_AGENTS.md`, and `08_JOB_TEMPLATE.md`, then verify the current environment and create a new snapshot.
5. Record implementation and verification in a chronological English engineering report.
6. Mark the prompt `complete` or `superseded` only through documented project history; never erase earlier task intent.

Creating or reviewing a prompt does not authorize execution. Prompts contain instructions and acceptance criteria, not secrets, credentials, unverified environment claims, application output, or completed architecture decisions. This workflow will evolve through approved engineering-governance updates.
