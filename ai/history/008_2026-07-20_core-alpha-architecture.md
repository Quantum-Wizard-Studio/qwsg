# Task History 008: Core Alpha Architecture and Slice 1 Definition

## Task Metadata

- Task ID: `008`
- Date: `2026-07-20` UTC
- Status: `complete with disclosed architecture gates and permission limitation`
- Responsible agent: `Aikó/Codex`
- Human authority: `Attila (Project Owner)`
- Owner communication language: `Hungarian`
- Active prompt: `ai/prompts/008_CURRENT_TASK.md`

## Objective and boundary

Create the implementation-ready Core Alpha architecture and define `Core Alpha Slice 1: Read-only Server Discovery and System Inventory` without production code, dependency installation, host mutation, infrastructure change, commit, or push.

## Starting state

The verified root was `<repository-root>`, branch `main`, HEAD `8fa40acd945b5f0d5d1ee0c5e182a19bba092d2b`. Exactly one active prompt resolved to Task 008; Tasks 006 and 007 were complete. The worktree already contained authorized uncommitted lifecycle/documentation work. Product, test, and build placeholder directories contained no files. Toolchain and repository evidence are recorded in the snapshot.

The pre-existing Task 008 history was mode `0600` with ACL mask `---`, limiting group access. Permission changes were not authorized.

## Snapshot and rollback

Snapshot: `ai/backups/20260720T190700Z_task008_core_alpha_architecture/`. It contains start state, complete Git status, binary-capable tracked diff, permission/ACL evidence, affected-file list, preserved copies, manifest, checksums, and bounded restore. Checksums and `bash -n` passed before architecture edits.

## Work performed

1. Reconciled mandatory governance, product, specification, audit, readiness, and history sources.
2. Preserved the authority boundary: ratified design baseline did not convert labeled proposals or release gates into policy.
3. Defined system/trust/components, collector interface, normalization/validation, storage, errors, versioning, logging, extension, and test architecture.
4. Defined the exact non-root read-only Slice 1 boundary.
5. Defined the versioned inventory envelope, typed facts, provenance, freshness, unknown/redacted values, errors, and evolution.
6. Defined command, filesystem, symlink, privacy, redaction, permission, and no-network security constraints.
7. Recorded eleven gates and three evidence-supported ADRs.
8. Mapped all 125 FR and 19 AC identifiers by family and detailed Slice 1 contribution.
9. Created a staged Task 009/010 handoff and updated lifecycle indexes.

## Files created and updated

Created the three Core Alpha architecture/data/Slice documents, three ADRs, the security model, implementation plan, gate register, architecture mapping, this history, and the Task 008 snapshot. Updated README, CHANGELOG, Architecture Governance, System Map, Engineering History, Roadmap, QWSG project record, the master plan authority header, and Prompt 008. The placeholder history `008_2026-07-20_ezleszanyolcas.md` was replaced by this conventionally named record.

## Decisions

- Collectors are non-root, read-only, fact-only components behind a bounded coordinator.
- One versioned `InventorySnapshot` is canonical across CLI, optional local persistence, tests, and future Console presentation.
- The Agent owns local inventory truth; Slice 1 has no listener, outbound connection, Console, privileged helper, Installer, or remediation.
- Partial, unavailable, unsupported, permission-denied, timeout, cancelled, error, unknown, redacted, and stale meanings are explicit.
- Runtime, packaging, storage technology, configuration syntax, and unsupported technology choices were not selected.

## Proposals and open gates

Personas, editions, licensing, pricing, cloud, telemetry/business positioning, and labeled proposals remain unratified. Open gates cover platform matrix, e-mail, retention, Console, update authenticity, runtime/package, configuration/secrets, durable store, business/distribution, legacy governance/permissions, and partial-result CLI behavior.

## Risk assessment

Host stability is low risk. Security, privilege, data-model, privacy, lock-in, compatibility, authority, scope, and implementation-readiness risks are medium future-impact risks controlled by least privilege, versioned contracts, technology neutrality, explicit gates, and narrow scope. Rollback risk is low because the snapshot preserves the dirty baseline.

## Verification evidence

Final verification covered active-task validation, deliverable readability, snapshot checksums, restore syntax, 125 unique FR and 19 unique AC identifiers, valid mapping references, placeholder review, internal links, Markdown whitespace, absence of production source/manifests, pre-existing worktree preservation, permissions, and `git diff --check`. No product build or tests existed.

## Rollback

From the repository root run:

```bash
ai/backups/20260720T190700Z_task008_core_alpha_architecture/restore.sh
```

It restores only captured Task 008 targets and removes only Task 008-created architecture/history files. Unrelated pre-task changes and the snapshot remain.

## Recommended next task

Task 009 is recommended for internal Slice 1 implementation after separate approval selects a proportional runtime/package approach and stable partial-result CLI policy. Task 010 should independently verify and harden it. Neither task may silently close release gates or broaden Slice 1.

## Delivery result

**Complete with disclosed architecture gates and permission limitation.** The package is implementation-ready at contract level; Slice 1 is narrow, testable, read-only, least-privilege, traceable, and rollback-capable. No implementation, installation, infrastructure change, commit, or push occurred.
