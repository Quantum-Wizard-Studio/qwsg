# Engineering Update E001: Engineering Workflow Refinement

## Task Metadata

- Identifier: `E001`
- Date: `2026-07-18` UTC
- Responsible agent: Aikó/Codex
- Human authority: current project owner
- Preferred owner communication language for this task: Hungarian, established by current project context
- Status: complete, subject to final verification and Git recording

## Summary

E001 makes the existing Engineering Task Standard more explicitly structured, reusable, and AI-friendly; establishes a one-prompt-per-task workflow; and replaces a fixed Hungarian communication rule with a configurable preferred-language policy.

## Changes

- Reorganized `08_JOB_TEMPLATE.md` around the required fourteen-section official task structure while preserving earlier fields and lifecycle protections.
- Required the four core documents to be read before any implementation work.
- Updated the Constitution, Agent Rules, Engineering Standards, and Documentation Policy to use the current owner or lead developer's verified preferred language rather than hardcoding Hungarian.
- Strengthened localization policy so user-visible strings are not hardcoded whenever localization is technically feasible.
- Added `ai/prompts/README.md` to define naming, structure, status, approval, execution, and retention of prompts.
- Added the draft `ai/prompts/001_PRODUCT_ARCHITECTURE.md` without executing or approving Product Architecture.
- Updated Engineering History and created this E001 record.

## Reasoning

A stable section schema makes tasks easier for humans and AI tools to review, reuse, verify, and audit. Separating prompt definition from execution prevents a prepared future task from being mistaken for authorization. A configurable communication preference supports future owners and lead developers without weakening the permanent English engineering baseline.

## Recommendations

Use the official structure for every new prompt and delivery report. Verify the current communication-language preference at task start until a separate authorized task defines how project policy stores it. Do not execute Prompt 001 without explicit human approval and a fresh environment verification and snapshot.

## Verification

Verification must confirm all fourteen required headings, preservation of legacy fields, the four-document reading rule, prompt lifecycle rules, absence of hardcoded Hungarian as permanent policy, localization wording, snapshot integrity, rollback syntax, ownership, permissions, ACLs, scoped Git diff, excluded application artifacts, and final Git status.

## Rollback Information

Snapshot: `ai/backups/20260718T212711Z_E001_engineering_workflow_refinement/`. From the exact project root, run its `restore.sh`, review the bounded targets, and type `ROLLBACK-QWSG-E001`. It restores only six core documents and removes only the E001 history record and two new prompt documents. The snapshot remains available.

## Delivery Report

No application code, software, operating-system setting, or product architecture was created or modified. Product Architecture Prompt 001 remains an unexecuted draft. Unresolved issues: the policy-storage mechanism for communication-language preference is intentionally undefined and requires a future authorized governance task if needed.

## Completion Criteria

E001 is complete when documentation consistency and safety verification pass, the full engineering commit hash is recorded here, and the final working tree is clean. Git record: pending.

This English record belongs to chronological engineering history. Owner-facing completion communication is delivered separately in the preferred language.
