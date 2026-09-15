# Task History 085: Pro Automatic Update Trigger and Guardian Integration

## Task metadata

- Task ID: `085`
- Task slug: `pro-automatic-update-trigger`
- Status: `complete`
- Date: `2026-09-15` UTC
- Agent: Codex
- Human authority: Project Owner explicit Task 085 APPROVE; Task 084 accepted.
- Preferred owner communication language: Hungarian, verified configuration.
- Related prompt: `ai/archive_prompts/085_2026-09-15_pro-automatic-update-trigger.md`
- Dependencies: accepted Tasks 082, 083 and 084.

## Objective, scope and exclusions

Integrate the automatic trigger into the existing Guardian release-check
lifecycle, with canonical capability/policy and signed candidate authority,
safe Task 084 handoff, repeated-execution protection and evidence. One session;
close/archive, commit/push and verify clean synchronized main. No advanced
scheduling, reboot, fleet/remote management, licensing backend, release, v1.3.2,
Task 086 or unrelated cleanup. Historical v1.3.1 is immutable.

## Authority and required reading

Builder installed the single prompt/history pair using the Owner's APPROVE.
Scope includes local implementation, isolated fixtures, validation, bounded
snapshot, documentation and Git integration; production installation/service
mutation and publication remain excluded. Read the project-local job skill,
AGENTS.md, all prompt Required Reading, engineering/security/delivery and backup
policies, Task 084 history and relevant Guardian/scheduler, capability/policy,
authority, awareness, transaction and installed/helper fixture code.

## Verified starting state

Main and fetched origin/main were ab66c7323999e9748ca53a952f0e80b3f2daef7e;
clean, divergence 0/0, lifecycle idle after archived Task 084. No Task 085 existed.
Framework identity and canonical HTTPS origin validated. Historical v1.3.1 object:
efb7afcac7bd65a3419f6bf3264892e09de631f6. New files retain inherited regular,
non-executable checkout permissions and Git modes. No dependencies, credentials,
live services or infrastructure changed. No unrelated worktree content existed.

## Snapshot, rollback and risks

Before mutation a protected external-to-Git local baseline archive captured 583
regular tracked members with archive SHA-256 and a per-member manifest. Digest:

`00a04baf1964f00927e1daa724930b61e90281dcc1d2081d19dd6ea4a1ace874`

Archive readability and all member hashes were verified and revalidated after
implementation. Retain until Owner acceptance. No payload/private inventory is
committed. Restore only verified exact task targets after extracting into a new
empty private directory and reviewing bounded changes; preserve audit evidence,
obtain authority before destructive restoration and rerun affected tests/lifecycle.
Never extract over a live checkout or use broad reset/clean.

High-impact service/mutation risk is bounded by generation verification, complete
stop, Guardian's existing instance lock and unchanged Task 084 inactive check.
No forked child may act as coordinator inside Guardian's cgroup. Signed authority,
canonical entitlement, real package verification, helper checks and existing
transaction locking/rollback remain authoritative. Severe/incomplete evidence
inhibits unattended retry rather than guessing recovery success.

## Plan and work performed

Inspection identified the existing 24-hour/35-second Guardian release-check
lifecycle as the smallest cadence. It remains sequential and startup-gated;
Community discovery and configured notification remain intact. Task 085 adds:

- `internal/guardian/automatic.go`: separate policy/capability/awareness decision
  and reauthenticated handoff, current no-op, Task 084 invocation and composed
  terminal result. No transaction logic is duplicated.
- `cmd/qwsg/automatic_trigger.go`: canonical Guardian composition, persisted
  admission evidence, generation-checked finite user-manager handoff, full stop,
  existing Guardian lock, resume and structured result recording. Custom config
  Guardian instances refuse. The worker rechecks canonical capability, policy,
  source, signed candidate and watermark before lifecycle changes.
- `prepareAutomaticUpdate` factors only Task 084 production preparation so the
  handoff can supply lifecycle control before calling the same core. Community
  production authority remains unchanged; no configuration Pro flag exists.
- Guardian's cgroup is killed on stop and carries NoNewPrivileges=true. Inspection
  found no existing detached handoff. A finite transient user-manager service
  outside that cgroup is the explicit architectural-constraint exception allowed
  by the task. No persistent unit, new daemon, scheduler or timer is installed.
  Fixed unit naming excludes concurrent handoffs; Task 084 locks still protect
  actual transactions. ExecStopPost requests Guardian start as manager fallback.
- Awareness's existing private atomic/fsynced writer is shared with separate
  latest-decision and last-terminal files. Task 084's receipt is composed directly.
  A no-op cannot erase the terminal receipt. Corrupt, unfinished and failed-rollback
  receipts fail closed and remain available for operator review.
- Current installed identity prevents reapplying the same release. Signed current
  authority causes no stop/staging/apply. Unsupported or invalid authority does
  not reach the transaction. Ordinary failures resume Guardian; failed rollback
  remains severe and cannot become success or unattended ordinary retry.

## Verification evidence

- Focused trigger/handoff tests PASS: Community/default, Pro/manual, authorized
  current, supported candidate, unsigned/unsupported/stale-failed-check refusal,
  handoff/stop/restart failure, changed generation, Guardian instance lock,
  existing Task 084 lock contention, successful update/repeated current no-op,
  actual rollback success and explicit rollback failure.
- Separate-process Community/Pro black-box acceptance PASS. Both run the existing
  Guardian release scheduler with the same authenticated supported future fixture.
  Community produces no handoff or mutation. Pro requests async handoff, waits
  for scheduler quiescence, invokes the real Task 084/common package transaction
  once, persists success and verifies current no-op with terminal receipt retained.
  Service control and privilege are isolated fixtures; no live production root or
  manager is changed. This is not a claim of live systemd deployment acceptance.
- Scheduler continues after failed handoff; severe/incomplete receipt inhibition
  and private atomic evidence/symlink refusal tests PASS.
- `make fmt-check vet test`: PASS after final source changes. Full Go suite includes
  Tasks 082/083/084, Guardian/scheduler and frozen-old-client manual/helper security
  regressions. No tests weakened or skipped.
- Affected race suite PASS: guardian, scheduler, updateawareness, automaticupdate,
  updateauthority, productcapability, updatepolicy and cmd/qwsg. After the final
  evidence/inhibition correction, updateawareness and cmd/qwsg race tests reran.
- Framework v2: 15 assertions PASS; bounded runner: 11 assertions PASS; diverted
  task audit PASS. Snapshot integrity, diff whitespace and file scope reviewed.
- `make engineering-test`: PASS. Checkout/export provenance contract, 25 framework,
  36 diversion, 29 lifecycle and 49 task-builder assertions passed.
- Final idle lifecycle and framework identity checks are part of closure below.

### Diagnosed attempts

ENVIRONMENTAL ISSUE: direct Go tests could not write the default build cache in
the sandbox. The same focused command with approved elevation passed.

EXPECTED BEHAVIOR: the first engineering export build omitted new unindexed Go
files because its input is the tracked-file set. Targeted task staging made the
complete proposed source reviewable/exportable; the gate was rerun unchanged.
No build or framework gate was weakened.

TEST OR ACCEPTANCE DEFECT: the first idle-closure validation required canonical
Work performed / Completion state heading labels. Archive rotation rolled back;
the existing evidence headings were corrected and idle validation rerun.

## Documentation updates

Added `docs/architecture/AUTOMATIC_UPDATE_TRIGGER.md`; updated the orchestration
contract and architecture/milestone indexes narrowly. This history and the
completed archived prompt retain lifecycle and validation evidence.

## Completion state and disclosed boundaries

All required implementation and verification gates passed. Task 085 is complete;
the prompt is archived without a successor. Targeted integration includes staged
review, commit, explicit dry-run/fast-forward push, fetch and final clean/synchronized
verification. No source changed after final tests. No release or tag change;
Task 086 has NOT been started or prepared. Deferred: advanced maintenance/weekday/
blackout/timezone/fleet rollout, reboot, remote management, production entitlement,
persistent crash recovery and long health observation. Production remains
Community until a separately authorized trusted entitlement adapter exists.
Unexpected coordinator/host loss retains incomplete evidence, never a claim of
verified rollback. The acceptance fixtures do not claim a live systemd deployment.
Final commit/remote identities belong in the Owner handoff to avoid self-reference.
