# Task History 090: Local Evidence Interruption Integrity & Recovery

## Task metadata

- Task ID: `090`
- Task slug: `local-evidence-interruption-integrity-recovery`
- Status: `complete`
- Date generated: `2026-09-21` UTC
- Human authority: Project Owner explicitly authorizes Task 090 C3 implementation, tests, documentation, lifecycle closure, commit and push in this session.
- Preferred owner communication language: Hungarian
- Related prompt: `ai/archive_prompts/090_2026-09-21_local-evidence-interruption-integrity-recovery.md`

## Lifecycle state

The Builder installed the session-authorized task pair transactionally. Implementation, all nine acceptance cases and mandatory validation passed. Closure archives only Task 090, with no successor.

## Starting state

See the verified transaction map and snapshot evidence below.

## Snapshot

See the verified transaction map and snapshot evidence below.

## Work performed

Implemented kernel-owned exclusion with legacy-writer guard, bounded transaction recovery, preservation of installed evidence on sync/cleanup failure and bounded private evidence/checkpoint input. Added deterministic and process-exit acceptance tests.

## Verification

Builder input, metadata, prompt/history identity, approval state and lifecycle
installation validated successfully. Stable-source gofmt, vet, full Go tests,
focused race and engineering/build contract PASS. Engineering validation:
checkout/export binary/provenance contract PASS; 25 framework, 36 diversion,
29 lifecycle and 49 task-builder assertions PASS. Framework v2 and diagnostic
runner PASS; rollback archive/member hashes and documentation links PASS.
Detailed matrix and diagnosed attempts follow.

## Rollback

Verify protected archive checksum and every manifest member; extract into a NEW empty private directory only. Compare exact reviewed Task 090 paths and intervening drift before bounded restoration. Preserve new paths unless deletion is explicitly authorized. No broad reset/clean or published history rewrite; revalidate tests, framework/lifecycle, permissions and Git.

## Completion state

Complete: all C3 requirements and mandatory local gates passed. C3 moves from
PARTIAL to PASS; no other gate changes. Community totals 6 PASS / 2 PARTIAL /
1 MISSING; inherited Pro totals 7 PASS / 4 PARTIAL / 3 MISSING. No unresolved
Task 090 implementation blocker. Idle closure is validated before integration.
Integration uses reviewed explicit Task 090 paths, staged diff/mode/privacy
review, commit, canonical dry-run and fast-forward push, then fetch and clean
HEAD/origin/main equality with divergence 0/0. Exact resulting commit and push
outcome are provided in the Owner handoff to avoid a self-referential hash.

## Pre-implementation transaction map

Baseline verified after fetch: main/HEAD/origin/main at
`d8b8e598aebb25cbd126227a890b5d346faf9f0a`, divergence 0/0, clean;
framework valid, idle after completed 089. Builder installs approved 090 under
explicit session Owner authority. Required prompt reading and backup/engineering/
delivery/documentation policies read. Protected baseline snapshot verified:
603 tracked member hashes, readable archive SHA-256
`1897f47274f065f58450dafbcf22455bbc4ecd8072922f8c0b2c570989362103`.
Retain private payload and restore notes outside Git until Owner acceptance.

Actual inventory states: store.json describes immutable format/retention;
non-hidden snapshot JSON files are committed objects, validated by envelope,
identity, canonical model and SHA-256. Save creates layout before O_EXCL empty
.write.lock; oldest snapshot is renamed to .retire-<original> at capacity and
directory synced. atomicInstall creates .tmp-inventory-*, writes/fsyncs/closes,
links to final name (no overwrite), syncs directory, then removes temp. Save
removes retirement artifact and syncs. Before commit there are retention-1
visible snapshots plus one retired; after commit there are retention visible
plus one retired. Current errors attempt restoration, even removal of new
committed evidence if cleanup fails. Reads do not lock; any hidden artifact
blocks List; no recovery. Metadata installation uses the same atomic helper.
ReadFile follows only pre-stat inventory limits; directory enumeration unbounded.
Guardian checkpoint writes private temporary, fsync, rename, directory fsync;
Load uses unbounded ReadFile then 4 MiB check. Guardian already has kernel flock
for its operation lifecycle. Operator current-state reader already uses limit+1.
Updater transaction/backup state belongs to C7 and will remain unchanged.

Plan: separate permanent kernel lock inode, never unlink it; fail explicitly on
legacy .write.lock because empty legacy files cannot distinguish live owners.
Under lock classify bounded directory entries, validate all committed/retired
objects before recovery, restore retired if pre-commit count or finish retention
if post-commit count, preserve conflicting/unknown states. Never remove committed
new evidence on cleanup/sync error. Readers recover under same lock; no hidden
artifact can be promoted. Read limits enforced during input, formats unchanged.

## Implementation record

- Inventory lock acquisition now precedes layout/metadata mutation. Permanent
  private `.write.flock` uses nonblocking kernel flock and is never unlinked.
  Complete versioned `.write.lock` guard is atomically installed to exclude
  legacy O_EXCL writers, including mixed-version races. Empty/unknown legacy
  locks are preserved with explicit quiesce/preserve/move/retry instructions.
  All reads hold the same lock through recovery and evidence selection/loading.
- Recovery bounds root and snapshot enumeration, classifies known private
  temporary/retired objects before mutation, validates all evidence involved in
  retirement, restores pre-commit retirement or finalizes post-commit retention.
  Unknown files, conflicting/multiple retirements, invalid counts, corrupt
  evidence, absent metadata with existing evidence and unsafe types fail closed.
- Installed evidence is never removed as rollback after a sync/cleanup error.
  Recovery re-syncs observed final links before removing retired evidence;
  fsync failures cannot be returned as success. Existing no-clobber installation,
  SHA-256 envelope, canonical validation and retention policy are preserved.
- `internal/evidenceio` supplies bounded limit+1 reads and no-follow private
  regular-file opening to inventory/metadata and Guardian checkpoint readers
  only. No updater/shared update transaction helper changed. Guardian checkpoint
  atomic replacement, generation handling and operation locks remain unchanged.
- Tests use existing pre-install and a matching directory-sync fault hook,
  abrupt subprocess exit bypassing defers, real kernel lock collisions, and
  reconstructed actual post-install states. No production crash machinery,
  journal, dependency, daemon, telemetry or network lookup is introduced.

## Acceptance matrix

| Case | Result | Evidence |
| --- | --- | --- |
| 1 Normal persistence | PASS | Existing deterministic canonical round-trip/permissions and retention regressions, plus normal writes throughout recovery fixtures. |
| 2 A + interruption before B commit | PASS | TestInterruptedWriterProcess exits inside the actual pre-install boundary after retirement/temp fsync, bypasses defers; A is restored byte-for-byte, incomplete B never promoted. Existing returned-error pre-install regression also passes. |
| 3 Interruption after durable boundary | PASS | TestDurableInstallWithInterruptedRetirementCleanup uses actual synced atomicInstall, leaves retirement and extra temp hard link; B remains committed, unchanged, for N=1 and N=2. Directory-sync fault separately proves no false success or destructive reversal during durability uncertainty. |
| 4 Stale lock | PASS | Abrupt child exit releases kernel flock automatically; subsequent load and save succeed. Legacy/malformed locks fail with deterministic operator recovery, remain unchanged, and work after test-established quiescence and preservation outside store. |
| 5 Live writer | PASS | Real kernel-lock collision rejects another Store; blocked actual Save transaction excludes a concurrent Save. Permanent O_EXCL guard rejects legacy writer acquisition. |
| 6 Interrupted retention | PASS | Pre-install retired A restores; post-install retirement cleans up without reverting B; multiple/conflicting/wrong-count/corrupt/unsafe states remain unchanged and fail. |
| 7 Oversize | PASS | Infinite counting Reader consumes exactly limit+1 for 4096, 4 MiB and 16 MiB; oversized inventory/metadata and sparse 256 MiB checkpoint are rejected. Within-limit content and reader errors tested. |
| 8 Malformed/incomplete | PASS | Truncated temporary files never promoted; corrupt retired/committed transition refuses without mutation; existing duplicate-key/hash/schema/truncation checks pass; incomplete checkpoint rejects. |
| 9 Idempotence | PASS | Repeated pre/post-commit recovery preserves valid bytes and stable authoritative state, removes recognized artifacts, and does not progressively modify ambiguous states. |

Additional checks: legacy format 1.0 evidence reads unchanged without permanent
lock files; interrupted metadata initialization recovers; missing metadata never
silently resets a populated store; directory limits fail without cleanup;
symlink/content/permissive kernel locks fail; private file permissions and
snapshot filenames remain enforced. Inventory Partial remains valid observed
partial evidence, distinct from an incomplete persistence artifact.

## Validation attempts and classification

- ENVIRONMENTAL ISSUE: initial focused go test could not write the default Go
  cache in the sandbox. Authorized rerun passed. No source change to evade it.
- Full `make fmt-check vet test`: PASS on stable source, including all CLI,
  Community/pipeline, canonical evidence, privacy, Guardian, updater and retention
  regression packages. Focused race for inventorystore, evidenceio, guardian,
  collector, comparison, pipeline and operatorstate: PASS.
- TEST OR ACCEPTANCE DEFECT / ENVIRONMENTAL ISSUE: first build-contract export
  omitted new untracked package files (git ls-files export), leading to module
  resolution and sandbox-denied DNS. Explicitly staged the reviewed new source
  paths and reran engineering validation with authorized sandbox escalation.
  No build/check policy was weakened; full Go validation was not needlessly rerun.
- Framework v2: 15 assertions PASS; bounded diagnostic runner: 11 assertions
  PASS; active job/lifecycle and diverted-test audit PASS. Final archive/hash/
  all 603 snapshot members and affected local documentation links reverified.

## Documentation and boundary review

Updated the persistence architecture's actual lock/commit/recovery contract and
paired English/Hungarian operator guidance. No historical snapshot migration or
semantic change. Private directories and 0600 files are preserved; bounded
no-follow reads strengthen checkpoint handling. Task 089 pseudonyms/redaction,
authenticated release, mutation exclusion, Guardian recovery and Community/Pro
boundaries remain intact. No raw evidence content added to logs for recovery.

Updater backup/transaction durability and rollback remain C7. C8 general
readiness/recovery closure and C9 release work remain unchanged. Checkpoint
replacement and temporary-file lifecycle were inspected; this task changes its
input bound/type checks only, not Guardian generation/restart semantics or a
general persistence subsystem. The supported filesystem contract does not claim
arbitrary hardware repair, hostile same-user mutation handling or automatic
resolution of legacy ownerless locks.
