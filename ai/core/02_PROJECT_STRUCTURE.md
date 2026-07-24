# Project Structure

## Purpose

This document explains ownership and intended use of the repository's top-level areas.

## Status

The repository contains the pre-alpha one-shot Inventory application, validated
Inventory Store, user CLI, Snapshot Explorer, Comparison Engine, tests, installation workflow,
and governed engineering framework.

## What belongs here

- `ai/`: governance, coordination, snapshots, and delivery history.
- `ai/prompts/`: zero or one active engineering task prompt; empty is the canonical idle state after completion and archival.
- `ai/archive_prompts/`: immutable prior and unexecuted-draft prompt records.
- `ai/history/`: one independent history file per task plus concise milestone records retained for backward compatibility.
- `ai/test_tasks/`: independently numbered, owner-authorized experimental or
  aborted task evidence isolated from production task numbering and active-task
  selection.
- `ai/framework/`: reusable engineering core version identity.
- `ai/config/`: validated project identity and non-executable validation argv.
- `ai/scripts/`: bounded engineering workflow automation.
- `cmd/qwsg/`: the one-shot user CLI, help, safe rendering, JSON compatibility,
  Snapshot Explorer, and exit-policy boundary.
- `internal/collector/`: bounded read-only evidence acquisition.
- `internal/inventory/`: Inventory 1.0 and canonical model assembly/validation.
- `internal/inventorystore/`: explicitly invoked file-backed Digital Twin
  persistence, integrity, atomicity, and retention.
- `internal/comparison/`: deterministic canonical Snapshot Comparison Engine
  and Change Record validation; the sole system-evolution boundary.
- `docs/`: audience-focused architecture, installation, administration, development, security, release, and history guidance.
- `installer/`, `agent/`, `console/`, `modules/`: future approved product components.
- `tests/`, `scripts/`, `tools/`, `build/`: future verification and engineering support.

Credentials, runtime state, vendored dependencies, full backup payloads, host-state evidence, generated binary backups, and undocumented binaries do not belong in version control. Active prompts, archived prompts, task histories, sanitized backup manifests/checksums/restore documentation, and Markdown audit records may remain committed under the Engineering Backup Policy. This map will evolve as architecture is approved during development.
