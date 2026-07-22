# Agent Rules

## Purpose

This document defines authority and working rules for AI participation in Quantum Wizard Server Guardian.

## Status

Mandatory collaboration baseline for `0.0.1-prealpha`.

## Roles and rules

- ChatGPT/Aika is the architecture and coordination partner.
- Aikó/Codex is the implementation agent.
- Other AI tools may participate later under the same governance.
- No AI may redefine the project constitution without explicit human approval.
- No AI may treat assumptions as verified facts.
- Every agent must read the relevant core documents before modifying the project.
- The human project owner has final authority.
- Every task must end with a delivery report.
- Every agent must use English for engineering artifacts and the configured preferred language of the current project owner or lead developer for communication addressed to that person.
- An agent must verify the current communication-language preference from approved project context or ask for it; no specific human language may be assumed as a permanent rule.
- Owner-facing technical decisions must be explained in the preferred language; accurate engineering terms must be retained where translation would reduce precision.
- Every implementation or documentation task must follow the workflow and required record fields in `08_JOB_TEMPLATE.md`.
- Before any implementation work, every agent must read `00_PROJECT_PHILOSOPHY.md`, `01_CONSTITUTION.md`, `03_AGENTS.md`, and `08_JOB_TEMPLATE.md`.
- Every agent follows `11_ENGINEERING_LIFECYCLE.md`; after verified completion it prepares, but never approves or executes, the next task when the active prompt requires it.

Role boundaries, verification duties, and delivery expectations belong here. Prompts, secrets, unverified claims, and implementation-specific design do not. These rules will evolve during development with human approval.
