# Task History 020: Approved Task

## Task metadata

- Task ID: `020`
- Task slug: `canonical-drift-engine`
- Status: `complete — Canonical Drift Engine and all required gates passed`
- Date generated: `2026-07-24` UTC
- Human authority: Project Owner through the Task 020 implementation instruction
- Preferred owner communication language: Hungarian
- Related prompt: `ai/prompts/020_CURRENT_TASK.md`

## Lifecycle state

The Engineering Task Builder generated and transactionally installed this
matching prompt/history pair from the Project Owner's approved Task 020 brief.
`bin/job --check` and lifecycle consistency validation passed before work began.

## Starting state

The task started from canonical idle state after completed and archived Task
019. The verified repository root is
`/home/qws/web/qwsg.quantumwizard.hu/qwsg`; framework version is `1.0.0`;
branch is `main`; canonical origin is
`https://git.quantumwizard.hu/Quantum_Wizard_Studio/qwsg.git`. Fetch confirmed
HEAD `81c08f9b25ab1f86fca6453f3549a2f544909d44` is one expected Task 019
commit ahead of `origin/main`
`88d75a894983b4824efa26b32919086b8ae0d9c2`, matching Task 019 delivery
evidence. No tracked modification existed.

Pre-existing owner-owned untracked backup directories,
`current-task-job.txt`, `current-task-maker.md`, and the local `qwsg` binary
were present exactly as documented by Task 019 and remain outside scope.
Repository identity, remote, lifecycle, required reading, Comparison contracts,
permissions, ownership, ACLs, and Go project structure were inspected without
other material difference.

## Snapshot

The pre-Builder lifecycle snapshot is
`/tmp/qwsg-task020-lifecycle-20260724T132945Z`. It contains a complete Git
bundle and the idle prompt/archive/history state. SHA-256:

- repository bundle:
  `2cfe9ae119e301557022f4adee0a2dda4b0e500c4d2f2a8f079b552481b404cc`;
- lifecycle archive:
  `c3bb85825738e799439299e8e1207f51204c6e463fea9cab9386e3ed4dce977b`.

The post-Builder, pre-implementation snapshot is
`/tmp/qwsg-task020-implementation-20260724T132945Z`. It contains a complete Git
bundle, bounded archive of every pre-existing planned target, file list,
status, modes, ownership, ACL evidence, and explicit absence records for
`internal/drift` and
`docs/architecture/CANONICAL_DRIFT_ENGINE.md`. SHA-256:

- repository bundle:
  `2cfe9ae119e301557022f4adee0a2dda4b0e500c4d2f2a8f079b552481b404cc`;
- target archive:
  `fe4f119bc0a42bf4d071c63e9a7fa3535b084285a729525b26b887cc518b9160`.

Both bundles verified as complete histories; archive listings and checksum
verification passed. Snapshots are private under `/tmp` and retained through
owner acceptance.

## Work performed

- Added `internal/drift` as an isolated, pure semantic classifier over
  `comparison.ChangeRecord`.
- Defined public Drift schema, engine, and taxonomy version `1.0`, stable Drift
  IDs, explicit scope, four change classifications, deterministic
  confidence-basis-point semantics, bounded metadata, and per-record dependency
  versions.
- Implemented the required twelve categories plus a conservative `extension`
  category for unregistered future canonical layers.
- Added validation, fixed precedence, stable ordering, duplicate rejection,
  source-value privacy, and one-change/one-drift enforcement.
- Added taxonomy, determinism, ordering, immutability, invalid-contract,
  privacy, metadata-bound, and real Compare-to-Drift pipeline tests.
- Added the permanent Canonical Drift Architecture and updated Comparison,
  architecture governance, system map, roadmap, README, and milestone index.
- Preserved the Comparison contract and CLI behavior; added no runtime command,
  Health, risk, policy, alert, daemon, scheduler, AI, network, or remediation
  behavior.

## Verification

Builder input, metadata, prompt/history identity, approval state, and lifecycle installation validated successfully.

Focused Drift and Comparison package tests passed after implementation. A first
sandboxed test invocation could not write the existing external Go build cache;
it changed no repository state. The same tests then passed with approved cache
access.

`go test ./...` and `go test -race ./...` passed for all eight packages,
including the new Drift and unchanged Comparison packages. `make vet`,
`make fmt-check`, and `git diff --check` passed. The configured
`framework-check.sh --run-validations` gate passed active-job, lifecycle,
test-task audit, repository-wide Go tests, vet, and formatting. `make
engineering-test` passed 21 framework, 36 diversion, 28 lifecycle, and 38
Builder assertions. Direct active lifecycle, diverted-test audit, taxonomy
topic, documentation consistency, privacy/offline boundary, modes, ownership,
ACL, and complete Git-state reviews passed.

Tests prove all required categories, conservative extension handling, all four
change classifications, exact one-to-one record production, deterministic IDs,
ordering and byte-identical JSON, input immutability, bounded output metadata,
absence of source values, duplicate and invalid-contract rejection, and a real
`comparison.Compare` to Drift pipeline. Comparison tests passed unchanged.

Both snapshot checksum manifests and the implementation Git bundle were
reverified immediately before lifecycle completion. No dependency was added.
No pre-existing owner-owned untracked path was modified or staged.

## Rollback

Verify `SHA256SUMS`, the bundle, and target archive before restoration. Compare
current paths and obtain confirmation before destructive action. Restore only
the archived pre-existing Task 020 targets; remove only the two new paths whose
absence is explicitly recorded. Lifecycle rollback uses only the bounded
lifecycle archive. Re-run framework, lifecycle, Go, race, vet, formatting,
documentation, permissions, ACL, Git, and consistency checks. Broad reset,
checkout, clean, wildcard removal, and extraction over the live worktree are
forbidden.

## Completion state

`complete — all authorized implementation, documentation, rollback, compatibility, and validation gates passed`
