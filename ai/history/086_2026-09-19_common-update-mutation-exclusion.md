# Task History 086: Common Update Mutation Exclusion

## Task metadata

- Task ID: `086`
- Task slug: `common-update-mutation-exclusion`
- Status: `complete`
- Date: `2026-09-19` UTC
- Agent: Codex
- Human authority: Project Owner explicitly authorized task creation, execution, commit, push and closure in this session.
- Preferred owner communication language: Hungarian, verified project configuration.
- Related prompt: `ai/archive_prompts/086_2026-09-19_common-update-mutation-exclusion.md`

## Objective, scope and exclusions

Share one nonblocking process-safe exclusion boundary across manual authenticated
update, automatic update and manual rollback. Preserve Tasks 082-085 and all
security gates. No release/version bump, deployment, production acceptance,
Guardian redesign, privilege provisioning, retention, entitlement or unrelated work.

## Authority and required reading

Read AGENTS.md, the project-local job skill, all active prompt Required Reading,
backup, engineering, security and delivery policies, Task 085 history, mutation,
helper, transaction, awareness and orchestration code and focused fixtures.
Builder installed Task 086 from the explicitly authorized structured definition.

## Starting state

Verified main HEAD and freshly fetched origin/main:
`8dea8fcc649e89c995c3120052af6b564adb6ff2`; clean worktree, divergence 0/0.
Framework identity valid; lifecycle idle after complete archived Task 085.
No dependencies, credentials, server settings or production services changed.

## Snapshot

Before lifecycle/source mutation: protected local baseline tar outside Git,
589 regular members with per-member SHA-256 manifest; archive readable and all
member digests verified. Archive SHA-256:
`acf29560500dbec72aa2f068fcebe8fd2851515efce77d4786f09417e7bf718f`.
Retain until Owner acceptance. Payload and private paths are not committed.

## Plan, risks and rollback

Extract the existing secure nonblocking flock without changing its inode name;
place ownership at top-level coordinators, before state-dependent work and service
mutation, through helper completion, validation, recovery and record finalization.
High-impact coordination risk: lock ownership never establishes trust or authority.
Verify snapshot digest and members, extract into a new empty private directory,
compare exact changed paths, and review bounded restoration with destructive
approval where necessary. Never extract over a live checkout or broadly reset/clean.
Rerun affected tests and Git/lifecycle validation after any restoration.

## Work performed

- Existing manual update: executeUpdate -> synchronous privileged-apply ->
  update.Apply; first package mutation is destination creation/replacement after
  backup. Failure may call privileged-rollback; success records current.json and
  discards superseded backup. All are inside the coordinator scope.
- Automatic: automaticupdate.Run -> Host.Apply -> privileged-apply-report -> same
  independently authenticating helper/update.Apply. Existing atomic process guard
  and automatic.lock previously covered only this coordinator. The lease extends
  through Host.Validate, Commit and rollback/recovery validation.
- Manual rollback: runUpdateRollback -> privileged-rollback -> update.Rollback ->
  restore (destination removal/replacement after transaction/integrity checks).
  Ownership starts before loading current.json and lasts through its removal.
- Extracted existing private owner/mode-checked, O_NOFOLLOW/O_CLOEXEC nonblocking
  flock to internal/updatemutation. Retained automatic.lock for compatibility;
  never unlink it. A typed lease documents single coordinator ownership and a
  sentinel distinguishes contention. Existing automatic atomic guard remains.
- Manual contention returns exit 3 plus transaction_conflict and retry guidance;
  automatic returns transaction_conflict with MutationStarted=false. Unsafe lock
  errors fail closed as lock-unavailable, separately from contention. No retry loop.
- Helpers are synchronous internal steps under coordinator ownership and never
  reacquire recursively. Root/helper authentication and rollback validation remain
  independent requirements. Awareness/check/status retain their separate state
  protection and do not acquire mutation authority.

## Verification

- Focused suite PASS: all manual update/automatic update/manual rollback pairings
  in both directions, same-process re-entry and independent child processes using
  the actual command/Run entry points. Contenders do not enter helpers, mutate
  installed files or change rollback records. Deterministic callbacks/channels,
  bounded child processes and isolated fixtures; no destructive live testing.
- Ownership through manual helper/recovery calls and automatic post-validation/
  rollback PASS. Injected metadata, service-stop and helper failure release the
  lease; a later legitimate automatic transaction succeeds. Existing automatic
  concurrency/re-entry tests remain passing.
- Independent process flock, historical lock inode compatibility, retained inode,
  release, unsafe directory/lock mode, symlink and nonregular target tests PASS.
- Check and status remain usable while the mutation lease is held; awareness
  persists its authenticated result under its separate existing state lock.
- `make fmt-check vet test`: PASS (full `go test ./...`, `go vet ./...`). Includes
  unchanged Tasks 082-085 authentication, migration, Community/Pro capability and
  policy, trigger/handoff, privileged-helper and rollback security regressions,
  plus frozen-old-client forward authenticated update acceptance.
- `go test -race ./internal/updatemutation ./internal/automaticupdate
  ./internal/update ./internal/updateauthority ./internal/productcapability
  ./internal/updatepolicy ./internal/updateawareness ./internal/guardian
  ./internal/scheduler ./cmd/qwsg`: PASS.
- After adding the final separate-process command probes, their focused ordinary
  and race tests PASS; formatting and full vet rerun PASS. Production source did
  not change after full tests/race. No test weakened or skipped.
- `make engineering-test`: PASS; checkout/export build provenance, 25 framework,
  36 diversion, 29 lifecycle and 49 task-builder assertions.
- Framework v2: 15 assertions PASS; bounded diagnostic runner: 11 assertions PASS;
  diverted task audit PASS. Snapshot archive and all 589 member hashes reverified.
- Reviewed exact source/docs/lifecycle staging, whitespace, privacy and modes;
  new source/test records are ordinary nonexecutable files with inherited checkout
  permissions. No payloads, private inventory, credentials or unrelated paths added.

ENVIRONMENTAL ISSUE: default Go cache write was refused by the sandbox on a
subsequent focused run; rerun requested through the approved elevated Go test rule.
No validation gate is weakened.

## Documentation and deferred findings

Updated the orchestration ownership contract, this task history, archived prompt
and chronological milestone. No new release/version or next task.
All callers for an instance must share its canonical user/state directory. The
existing root-only helpers remain synchronous internal steps rather than public
coordinators; arbitrary root administration is outside cooperative exclusion.
Persistent recovery after coordinator/host loss and broader Guardian handoff/
restart coordination remain separate work. Historical manual binaries do not
participate; consistent client deployment is required. No deployment was performed.

## Completion state

All scoped implementation and mandatory validation gates passed. Task 086 is
complete and archived without a successor. Authentication, migration, capability,
policy, helper validation and rollback authority remain mandatory; Community stays
manual by default and lock ownership grants no Pro capability or trust.
Implementation and idle closure are integrated in one task-scoped commit, with
explicit dry-run push, fast-forward push, fetch and final synchronized clean idle
verification. Final commit identities are reported in the Owner handoff to avoid
self-reference. No release, tag, version bump or production installation.
