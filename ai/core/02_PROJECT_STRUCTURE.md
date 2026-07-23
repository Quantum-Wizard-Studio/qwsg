# Project Structure

## Purpose

This document explains ownership and intended use of the repository's top-level areas.

## Status

Initial directory model; directories contain no application implementation yet.

## What belongs here

- `ai/`: governance, coordination, snapshots, and delivery history.
- `ai/prompts/`: zero or one active engineering task prompt; empty is the canonical idle state after completion and archival.
- `ai/archive_prompts/`: immutable prior and unexecuted-draft prompt records.
- `ai/history/`: one independent history file per task plus concise milestone records retained for backward compatibility.
- `ai/test_tasks/`: independently numbered, owner-authorized experimental or
  aborted task evidence isolated from production task numbering and active-task
  selection.
- `ai/scripts/`: bounded engineering workflow automation.
- `docs/`: audience-focused architecture, installation, administration, development, security, release, and history guidance.
- `installer/`, `agent/`, `console/`, `modules/`: future approved product components.
- `tests/`, `scripts/`, `tools/`, `build/`: future verification and engineering support.

Credentials, runtime state, vendored dependencies, full backup payloads, host-state evidence, generated binary backups, and undocumented binaries do not belong in version control. Active prompts, archived prompts, task histories, sanitized backup manifests/checksums/restore documentation, and Markdown audit records may remain committed under the Engineering Backup Policy. This map will evolve as architecture is approved during development.
