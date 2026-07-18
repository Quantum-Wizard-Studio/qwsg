# Project Structure

## Purpose

This document explains ownership and intended use of the repository's top-level areas.

## Status

Initial directory model; directories contain no application implementation yet.

## What belongs here

- `ai/`: governance, coordination, snapshots, and delivery history.
- `docs/`: audience-focused architecture, installation, administration, development, security, release, and history guidance.
- `installer/`, `agent/`, `console/`, `modules/`: future approved product components.
- `tests/`, `scripts/`, `tools/`, `build/`: future verification and engineering support.

Credentials, runtime state, vendored dependencies, generated backups, and undocumented binaries do not belong in version control. This map will evolve as architecture is approved during development.
