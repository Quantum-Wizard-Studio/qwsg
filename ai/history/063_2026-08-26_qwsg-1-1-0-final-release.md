# Task History 063: Approved Task

## Task metadata

- Task ID: `063`
- Task slug: `qwsg-1-1-0-final-release`
- Status: `active`
- Date generated: `2026-08-26` UTC
- Human authority: Project Owner
- Preferred owner communication language: English
- Related prompt: `ai/prompts/063_CURRENT_TASK.md`

## Lifecycle state

The Engineering Task Builder generated and transactionally installed this matching prompt/history pair from validated structured Owner input. The Project Owner authorized full Task 063 execution under Framework 2.0 Standard Execution Authority on 2026-08-26 UTC. Implementation is active.

## Starting state

- Ordinary user `attila`, canonical repository root, branch `main`.
- Framework `2.0.0` VALID; Task 063 was the sole approved active prompt/history
  pair; index and tracked worktree began clean at
  `1f26b7c16e542fdf72614269d1039bbf436d6bbf`.
- `HEAD`, `origin/main`, and direct Forgejo `refs/heads/main` matched that
  commit with `0/0` divergence. Remote/local tag `v1.1.0` was absent.
- `VERSION=1.1.0-rc.6`; final 1.1.0 notes, acceptance evidence and artifacts
  were absent. The unrelated Owner blueprint was excluded and unread.
- Accepted behavioral source
  `25a30718bc92882e9773a5c405ad648c0eee1a81` is an ancestor. The path diff
  through current HEAD contains only Task 061 evidence/lifecycle and Framework
  2 changes; behavior-bearing QWSG source and packaging differences are NONE.
- Task 061 evidence remains authoritative: practical acceptance PASS,
  QWSG-059-F001 externally corrected, no RC.6 product defect, controlled SMTP
  plus Owner receipt PASS, expected notification/overall partial semantics,
  and READY FOR RELEASE DECISION.
- Baseline execution first exposed a Framework v2 self-test fixture assumption:
  its incomplete-envelope negative case removed only a legacy label although
  the current Task 063 fixture uses the concise v2 envelope. Classified `TEST
  OR ACCEPTANCE DEFECT`; the test now removes the applicable label for either
  supported envelope form.

## Snapshot

- Execution snapshot: `/tmp/qwsg-task063-execution.1QW3IP`, mode `0700`.
- Manifest SHA-256:
  `f5d0076fcfabacd1177c0afaf4c2dc897ee4eccac05d821c1026c855fcf836ce`.
- Exact tracked-HEAD archive SHA-256:
  `35ba83361438f981a9fc60bb7252f80a50423088806f8c3df17d4aa730045d0c`.
- Contains exact HEAD payload, Task 063 lifecycle before-images, explicit
  pre-correction Framework test before-image/diff, accepted-lineage path
  evidence, release absence claims, exclusion record, hashes and bounded
  rollback instructions. Credentials, candidate bytes, caches and Owner
  content are excluded.
- The small Framework test-fixture correction occurred immediately before the
  execution snapshot while diagnosing the baseline suite. Exact pre-change
  content was recovered from immutable HEAD and included in the verified
  snapshot; rollback remained available throughout. This is recorded as a
  procedural ordering limitation, not missing rollback evidence.

## Work performed

Execution started. Corrected the concise/legacy Authority Envelope fixture
selection in `ai/tests/test-framework.sh`; no framework behavior changed.

- Changed final release identity from `1.1.0-rc.6` to `1.1.0`.
- Added final release notes, moved the completed Unreleased changelog content
  under the dated `1.1.0` heading, and changed canonical installation examples
  to the final immutable archive name.
- Extended `scripts/build-release.sh` so final `1.1.0` is accepted only with the
  same explicit epoch and full lowercase 40-character source commit required
  for the accepted RC.6 lineage. Updated positive, missing-metadata,
  short-commit and collision release-plumbing regressions.
- No `cmd/`, `internal/`, systemd unit, installer, uninstaller, release runtime
  configuration or operator behavior changed. Task 061 external product
  acceptance therefore remains valid and is reused without VPS access.

## Verification

Builder input, metadata, prompt/history identity, approval state, lifecycle
installation, canonical refs, accepted lineage, release absence and snapshot
integrity validated successfully. Full baseline validation resumes after the
test-fixture correction.

- Corrected baseline: Framework/configured validations PASS (Framework 25,
  Builder 49, v2 semantics 15, diagnostic runner 11, full Go, vet, formatting),
  repository-wide Go race PASS, lifecycle 29, diversion 36, test-task audit,
  RC.6 release plumbing, build contract, shell syntax, Git whitespace and
  systemd static verification PASS. The static systemd tool emitted the known
  sandbox bus warnings while returning success.
- Final release-source validation: final release plumbing, strict provenance
  negative checks, configured Framework/full Go/vet/format suite, build
  contract, shell syntax, Git whitespace and static unit verification PASS.
- Release-source change classification: release identity/documentation/plumbing
  only; runtime behavior mutation NONE; Task 061 acceptance reuse VALID.

## Rollback

Use `/tmp/qwsg-task063-execution.1QW3IP/ROLLBACK.txt` only after exact identity
and collision checks. Restore literal targets from the tracked-HEAD archive or
explicit before-images; never use broad reset/checkout/clean or touch excluded
Owner content.

## Completion state

`active — final release implementation in progress`
