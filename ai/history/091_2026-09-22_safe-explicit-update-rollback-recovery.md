# Task History 091: Safe Explicit Update, Rollback & Recovery

## Task metadata

- Task ID: `091`
- Task slug: `safe-explicit-update-rollback-recovery`
- Status: `complete`
- Date generated: `2026-09-22` UTC
- Human authority: Project Owner explicitly authorizes Task 091 C7 implementation, tests, documentation, lifecycle completion, commit and push in this session.
- Preferred owner communication language: Hungarian
- Related prompt: `ai/archive_prompts/091_2026-09-22_safe-explicit-update-rollback-recovery.md`

## Lifecycle state

The Engineering Task Builder generated and transactionally installed this matching prompt/history pair from validated structured owner input. Explicit session Owner approval was recorded. Implementation, all twelve acceptance cases and mandatory validation passed. Closure archives only Task 091, with no successor.

## Starting state

See the verified baseline, snapshot and bounded restore record below.

## Snapshot

See the verified baseline, snapshot and bounded restore record below.

## Work performed

Implemented the bounded C7 changes detailed below.

## Verification

Builder input, metadata, prompt/history identity, approval state, and lifecycle installation validated successfully.

## Rollback

See the verified baseline, snapshot and bounded restore record below.

## Completion state

Complete: C7 moves PARTIAL → PASS on the evidence below. Community totals are
7 PASS / 1 PARTIAL / 1 MISSING; inherited Pro totals 8 PASS / 3 PARTIAL / 3 MISSING.
No other gate changes. Reviewed task paths are integrated by commit, canonical
push dry-run and fast-forward push, followed by fetch and final clean equality
verification reported in the Owner handoff. No next task is authorized.

## Verified baseline, snapshot and transaction map

Fetched origin; main/HEAD/origin/main equal a27935e44fcbe60012cf6869490eefe3d0f71a89,
divergence 0/0, clean worktree, framework valid, idle after completed 090.
Required reading and engineering/backup/delivery/documentation policies read.
Before lifecycle mutation, protected baseline archive verified: 610 member hashes,
SHA256 8550d5284fb9a6494294159539f4501c457ca6e772a898b5f18abc2e19d2b657.
Payload retained outside Git until Owner acceptance. Restore only to a NEW private
directory after archive/member verification; review exact task paths and drift,
never broad reset/clean. Revalidate tests, framework, lifecycle and Git afterward.

Current manual ordering, inspected before source changes:
1. CLI checks manual capability/non-root and takes the common mutation lease.
2. Reads installed identity and prior current.json; authenticates signed metadata,
   migration declaration and awareness watermark; stages/verifies package.
3. Validates installed configuration; queries enabled/active using yes/no commands.
4. Stops active Guardian without explicit inactive verification.
5. Privileged helper independently reauthenticates source/metadata/package.
6. Apply interleaves backup copy and destination replacement, then writes its
   transaction journal only after all replacements; returned failure restores
   the in-memory write set and validates hashes, but interruption has no journal.
7. Manual coordinator reloads/enables/starts Guardian before version/config check.
8. On failure it discards rollback, reload, enable and start results; generic exit 1.
9. On success it saves current.json, discards prior backup and reports success.
10. Explicit rollback already has durable intent and Guardian recovery evidence,
    common lease, package validation before runtime recovery and nonzero failure.
    However helper deletes backup before coordinator validation/recovery, and
    low-level restore verifies sources incrementally while mutating destinations.

Confirmed PRODUCT/FRAMEWORK DEFECTS: manual outcome loss and ordering; missing
pre-mutation durable complete rollback set; early rollback-source destruction;
invalid later backup source can permit partial restore. Plan: prepare/sync all
backup sources and journal before mutation, prevalidate entire rollback source,
preserve backup through validation/retry, reuse Guardian recovery semantics and
record explicit manual phases/results before destructive boundaries. No C3
inventory primitive or Pro policy expansion. Builder's first input rejected only
Markdown label syntax; corrected before successful transactional installation.


## Implementation and acceptance evidence

- Manual coordination consumes the existing authenticated helper receipt and
  Guardian Recovery type. Every installed package destination is compared with
  its authenticated staged counterpart (bytes/mode), followed by identity/config
  validation and daemon reload, before Guardian start. Runtime verification
  requires active/running, success and positive MainPID; inactive intent remains
  inactive. Failed updates remain exit 1 after successful recovery; lock conflict
  remains exit 3. No-op authenticated online current remains exit 0 without handoff.
- `update-result.json` durably records prepare/mutation/validation/recovery phases,
  apply receipt, rollback execution/validation, runtime result and intervention.
  `manual-attempt.json` preserves pre-mutation refusal/no-op phase and exit;
  explicit `rollback-result.json` is separately displayed by update status.
  Unknown/incomplete results inhibit another manual update. Missing helper
  receipt is conservative, never proof of no mutation. Final evidence failure
  cannot print successful update or emit a success notification.
- The complete root-owned backup write set is copied/fsynced and prevalidated;
  atomic prepared journal plus directory/ancestor fsync precede replacement.
  Replacements sync files/directories. Whole-source rollback prevalidation
  precedes any restore. Restored bytes/modes/absence validate independently.
  Existing complete journals remain compatible, prepared journals allow explicit
  interrupted recovery. Sources survive failed rollback, validation and recovery.
- Explicit rollback consumes pending transaction evidence if normal current.json
  was not committed. Original Guardian intent survives retry after runtime
  failure. Pre-mutation interrupted stop can recover without nonexistent package
  backup. Normal completed rollback removes its local pointer with directory
  fsync; repeated invocation cannot silently apply a different transaction.
- Common mutation lease remains held through helper, validation, recovery and
  evidence. Signed release-index, migration capabilities, local archive authority,
  watermark and helper reauthentication remain mandatory. No new update lock,
  runtime fault flags, dependency, scheduled execution or Community automation.
  C3 code is unchanged; its bounded private evidence reader is reused.

| Case | Result | Evidence |
| --- | --- | --- |
| 1 Normal explicit success | PASS | TestManualC7TransactionMatrix/success executes signed staging and actual privileged dispatcher/Apply in private roots, validates all installed package bytes/modes and identity/config, then verified Guardian recovery and durable success. OldBinaryForwardAuthenticatedUpdate online/archive also pass. |
| 2 Failed mutation recovered | PASS | Matrix/mutation_failure preserves nonzero command and update_failed_recovered with successful rollback/validation/runtime. TestApplyAutomaticallyRestoresAfterMutationFailure exercises actual partial mutation and helper restoration. |
| 3 Apparent success fails validation | PASS | Matrix/validation_failure corrupts installed unit after successful helper return, detects exact package mismatch before start, restores old package, reports recovered failure. |
| 4 Rollback execution failure | PASS | Matrix/rollback_failure retains failed rollback, intervention and blocked Guardian; no start may erase it. |
| 5 Rollback validation failure | PASS | Matrix/rollback_validation_failure returns successful restore but rejects restored state; failure remains separate, Guardian blocked. |
| 6 Runtime recovery failure | PASS | Matrix/recovery_failure preserves validated rollback plus failed Guardian recovery; retry succeeds using original running intent. ManualRollbackRetryPreservesOriginalRuntimeIntent covers failure after explicit rollback of a successful update. |
| 7 Initially inactive | PASS | Matrix/inactive and existing ManualRollbackRecoveryEvidence/stopped preserve verified inactivity without start/stop. |
| 8 Common mutation collision | PASS | Existing ManualOwnersExcludeAllMutationPathsAndReleaseOnFailure, ManualMutationOwnershipThroughHelperAndRecovery and AutomaticOwnsCommonMutationBoundary test actual lease/process contenders through helper/recovery; exit 3 and no mutation. |
| 9 Invalid authority | PASS | ManualC7RefusalAndCurrent missing/invalid/conflicting authority refuses before mutation/handoff and saves authentication attempt evidence; old-client unsigned/unknown/source/digest/provenance/forged-sidecar/helper refusal regressions preserved. |
| 10 Invalid rollback source | PASS | ManualC7InvalidRollbackAndRepeat missing/invalid sources preserve current installation; InvalidLaterRollbackSourceDoesNotMutate detects a corrupt later source before restoring even the first destination. |
| 11 Repeated/current no-op | PASS | ManualC7RefusalAndCurrent/current returns 0 with no helper or Guardian handoff. Explicit rollback repeat refuses after consumed pointer without changing installed files; low-level prepared rollback is idempotent. |
| 12 Durable recovery evidence | PASS | Matrix reads persisted final states for all distinct failures. InterruptedMutationPreservesRecoveryIntent and InterruptionBeforeMutationRecoversWithoutBackup prove honest incomplete intent and explicit recovery. PreparedJournalPrecedesMutation, InterruptedMutationUsesPreparedRollbackJournal, JournalSyncFailureNeverCrossesMutationBoundary and DestinationSyncFailurePreservesRecoverableBackup exercise actual prepare/replacement/durability boundaries; terminal-evidence fault prevents success. |

## Validation attempts and classification

Focused update/authentication/mutation/Guardian tests PASS. Focused CLI matrix,
rollback and recovery tests PASS. Old-client binary acceptance PASS using only
synthetic private test packages, no release construction/publication.

ENVIRONMENTAL ISSUE: default Go build cache writes were sandbox-blocked;
authorized reruns passed. Full local format/vet passed. First full suite hit
sandbox denial creating httptest TCP listener; reran test/engineering with
required escalation rather than changing tests or weakening checks.

TEST OR ACCEPTANCE DEFECT: mutation fixture initially intercepted only the old
unstructured helper, so it reached local sudo (which refused without a terminal).
Updated test seam for the structured helper, retained collision assertions.
Conflicting-authority fixture initially attempted to sign structurally invalid
metadata with a validating signer; changed it to deliver the malformed document
and prove fail-closed refusal. Old-client assertions now require the specific
recovered-failure outcome instead of the obsolete generic rollback-attempt text.

Race PASS for update, authority, mutation, automaticupdate and Guardian. CLI race
initial cache sandbox failure was rerun alone with required escalation: PASS.
Framework v2: 15 assertions PASS. Bounded diagnostic runner: 11 assertions PASS.
Full Go suite PASS on authorized rerun. Engineering validation PASS:
checkout/export build provenance contract, 25 framework, 36 diversion,
29 lifecycle and 49 task-builder assertions. Source was frozen before broad validation; successful checks
are reused rather than rerun for ceremony.

## Scope and remaining operational boundaries

No VERSION change, release artifacts for publication, signatures, tags, release
index update, production/VPS operation or new task. Only C7 is eligible to change.
C8 documentation closure, C9 real released 1.3.1 upgrade acceptance, and
P1/P3/P4/P5 remain outside scope. Root backups retained for safe retry are not an
automatic retention policy. Supported fsync filesystem guarantees are required.
After coordinator interruption, the operator must first establish that any
privileged helper has exited before initiating recovery; no autonomous replay
or process-loss supervision architecture was added.

Documentation changes: Native Update and Rollback architecture, directly affected
operator recovery section, frozen C7 gate (only after validation), task history,
prompt closure and chronological milestone index. Snapshot payload stays outside
Git. Exact final integration commit is reported in the Owner handoff to avoid a
self-referential commit identifier in tracked history.


## Final review and integration

Snapshot archive checksum and all 610 member hashes reverified PASS. Final source
format/vet/full Go tests and relevant race suites PASS. Private payloads/logs
remain outside Git. Reviewed explicit staging set: the new manual coordinator
and its matrix tests; update_commands, update_mutation and forward_update tests;
low-level transaction and durability tests; native update/rollback architecture,
operator operations, C7 scope register, milestone index, archived prompt and this
history. No unrelated untracked paths, permission changes, dependencies or
production state are included. Final framework/lifecycle and staged diff checks
are required before the authorized commit and push.
