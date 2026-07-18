# Project Structure

## Purpose

This document explains ownership and intended use of the repository's top-level areas.

## Status

Initial directory model; directories contain no application implementation yet.

## What belongs here

- `ai/`: governance, coordination, snapshots, and delivery history.
- `ai/prompts/`: exactly one active engineering task prompt.
- `ai/archive_prompts/`: immutable prior and unexecuted-draft prompt records.
- `ai/history/`: one independent history file per task plus concise milestone records retained for backward compatibility.
- `ai/scripts/`: bounded engineering workflow automation.
- `docs/`: audience-focused architecture, installation, administration, development, security, release, and history guidance.
- `installer/`, `agent/`, `console/`, `modules/`: future approved product components.
- `tests/`, `scripts/`, `tools/`, `build/`: future verification and engineering support.

Credentials, runtime state, vendored dependencies, generated binary backups, and undocumented binaries do not belong in version control. Active prompts, archived prompts, task histories, and Markdown audit records must remain committed. This map will evolve as architecture is approved during development.
