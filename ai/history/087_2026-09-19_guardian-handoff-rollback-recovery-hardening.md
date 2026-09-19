# Task History 087: Guardian Handoff & Rollback Recovery Hardening

## Task metadata

- Task ID: `087`
- Task slug: `guardian-handoff-rollback-recovery-hardening`
- Status: `complete`
- Date: `2026-09-19` UTC
- Agent: Codex
- Human authority: Project Owner explicitly authorized creation, implementation, validation, commit, push and closure in this session.
- Preferred owner communication language: Hungarian, verified project configuration.
- Related prompt: `ai/archive_prompts/087_2026-09-19_guardian-handoff-rollback-recovery-hardening.md`

## Objective and boundaries

Make Guardian restart a consequence of evaluated recovery intent and validated
package state; strengthen durable automatic and manual rollback evidence.
Tasks 082–086 authorization and common mutation exclusion remain mandatory.
No release, version bump, production acceptance/mutation, sudo provisioning,
service configuration, entitlement, retention, inventory or unrelated work.

## Required reading and starting state

Read AGENTS.md, the project-local qwsg-job skill, all prompt Required Reading,
backup/engineering/security/delivery policies, Task 086 history and the existing
handoff/orchestration contracts and relevant source/tests. Framework identity valid.
Starting main HEAD and freshly fetched origin/main both
`9cd2cec207dd7f78b8fab58f7a5cf16a9118d67b`; divergence 0/0, clean worktree,
valid idle lifecycle, latest completed archived Task 086. Builder transaction
assigned Task 087 under the explicit Owner authority.

## Snapshot and rollback

Before lifecycle/source changes, protected local baseline archive outside Git;
594 regular tracked members with SHA-256 manifest. Archive readable and every
member digest reverified. Archive SHA-256:
`0afe2448ff9f3c79bb4026cbeeeb08ff0747b3ff26b16ef68a82255ef6ec57fd`.
Retain until Owner acceptance. Payload/private storage paths are not committed.
Verify archive and member hashes, extract only into a new empty private directory,
compare exact task changed paths, then review bounded restoration with destructive
authority where necessary. Never extract over a live checkout or broad reset/clean.
Post-restore gates: affected tests, framework/lifecycle and exact Git state.

## Ownership map before editing

- Task 085 release-check trigger evaluates capability/policy and signed awareness;
  only canonical generation/state can request the transient systemd worker.
- ExecStart runs the internal automatic-handoff worker; ExecStopPost independently
  requested Guardian start on worker exit. This competed with Go Resume intent.
- automaticServiceControl owned in-memory stopRequested and instance lock; Go
  AutomaticHandoff stopped/resumed outside Task 084 Run's mutation lease.
- Task 084 Run owned package authorization, staging, helper receipt classification,
  validation, rollback and commit; Task 086 exclusion covered its transaction.
- Automatic terminal evidence composed Task 084 result in automatic-result.json;
  a start command return alone was recovery evidence. Pending/failed rollback
  receipts already inhibited unattended retries.
- Manual rollback owned its common lease, helper and start calls, but lacked
  package revalidation, verified recovery and a separate durable terminal receipt.

## Work performed and risks

- Removed transient ExecStopPost fallback. Go is the sole recovery owner.
- Wrapped preflight performs the handoff only after authenticated staging, inside
  Run's common mutation lease. An optional trusted host finalizer keeps recovery
  and terminal persistence under that lease without nested acquisition.
- Service observation distinguishes active/inactive/unknown; active additionally
  requires expected InvocationID. Durable pending intent precedes stop. Originally
  inactive Guardian remains stopped. No service enablement policy is changed.
- Recovery permits start only after known non-mutating failure, validated update
  or validated rollback. Incomplete/failed rollback and abnormal apply do not start.
  A panic/nonterminal transaction cannot use a missing receipt as proof of safety;
  interrupted mutable phases explicitly retain unknown mutation evidence.
- Shared recovery uses a separate 30-second context; successful start must also
  yield ActiveState=active, SubState=running, Result=success and positive MainPID.
  This is immediate service evidence, not a long-duration health guarantee.
- Automatic evidence retains package result independently from intent, stop request,
  verified running/preserved stopped/failed/blocked state and overall outcome.
  recovery_failed never becomes success; evidence failure leaves pending evidence.
  Contenders cannot overwrite another worker's terminal/pending receipt. Severe,
  pending and unknown recovery states fail closed for unattended retry.
- Manual rollback now persists rollback-result.json before side effects and at
  completion. Helper success requires canonical installed identity/configuration
  validation, then required recovery. Failed helper/validation blocks restart;
  failed recovery retains rollback metadata. No success output precedes evidence.
  The existing local atomic writer now also fsyncs its parent directory.
- Tests use isolated package roots and explicit service/privilege fixtures; no
  production mutation, dependency, credential or global policy change occurred.

## Verification

All required checks passed on the final source:

- Focused scenarios: initially running/stopped, preflight/staging refusal,
  validated automatic rollback for both intents, failed rollback and rollback
  validation, failed start and successful-but-inactive start, missing service
  properties, generation change and instance exclusion.
- Worker loss: no ExecStopPost fallback, durable pending intent, retry inhibition;
  panic during apply preserves unknown mutation without unsafe restart.
  Intent and terminal evidence failures fail closed.
- Manual rollback: helper refusal, installed-version validation refusal, unknown
  service, running/stopped recovery and recovery failure; failure retains current
  metadata and separate package/recovery truth.
- Task 086: manual/automatic/rollback contention, independent processes, re-entry,
  failure release, exclusion through stop/recovery/terminal evidence. Existing
  Community/Pro and frozen-old-client authenticated forward-update/helper/rollback
  security regressions pass; only added service boundaries are fixture-controlled.
- `make fmt-check vet test`: PASS on final source, including `go vet ./...` and
  `go test ./...`.
- Full affected race suite PASS: `go test -race ./internal/updatemutation
  ./internal/automaticupdate ./internal/update ./internal/updateauthority
  ./internal/productcapability ./internal/updatepolicy ./internal/updateawareness
  ./internal/guardian ./internal/scheduler ./cmd/qwsg`. After final review,
  affected package race passed again; after the final interrupted-mutation
  correction, recovery/worker/interruption/rollback/automatic/manual/concurrency
  scenarios passed again under race. No races reported.
- `make engineering-test`: PASS with source frozen: checkout/export build contract,
  25 framework, 36 diversion, 29 lifecycle and 49 task-builder assertions.
- Framework v2: 15 assertions PASS; diagnostic runner: 11 assertions PASS;
  diverted-task audit PASS; framework identity and final idle lifecycle PASS.
- Snapshot archive and all 594 file hashes reverified; no special payload objects;
  restore inspection script syntax valid. Exact task diff, modes, whitespace,
  private-key/path exclusions and bounded staging reviewed. No unrelated files.

Failed attempts and corrections:
Fixture corrections: private directory mode, version-output contract and explicit
new service-boundary mocks. These were TEST OR ACCEPTANCE DEFECTS, corrected
without weakening assertions. TEST OR ACCEPTANCE DEFECT: two build-contract
attempts captured earlier source snapshots while final review still changed source;
checkout/export bytes correctly differed. Rejected those invalidated attempts and
reran the unchanged validator with source frozen. ENVIRONMENTAL ISSUE: sandbox refused the default
Go build cache during nested CLI builds; approved elevated Go tests and repository
Makefile cache configuration resolved it. No validation failure is counted as PASS.
The first closure validation correctly refused the missing literal Work performed
heading; the transaction restored the active pair. Corrected the history heading
and reran closure validation without changing source or weakening the validator.

## Documentation and deferred findings

Updated existing automatic trigger/orchestration contracts and task evidence.
Persistent restart/replay after worker or host loss remains outside scope: the
worker leaves incomplete evidence, Guardian may remain stopped, and operator
review is required. Long health observation and production VPS acceptance remain
excluded. Historical clients require consistent deployment for common exclusion.

## Completion state

All implementation and mandatory verification gates passed. Task 087 is complete
and archived without a successor. Tasks 082–086 authentication, migration,
capability/policy, trigger, common exclusion, privileged helper and rollback
authority remain mandatory. Recovery evidence grants no authority; Community
remains manual. No release, tag, version bump or production deployment.

Implementation and idle closure use one task-scoped commit. Integration protocol:
explicit reviewed staging, staged whitespace/privacy review, commit, canonical
`git push --dry-run origin main`, `git push origin main`, fetch and verification
of HEAD/origin equality, 0/0 divergence, clean worktree and idle lifecycle.
Actual commit identity and final repository evidence are supplied in the Owner
report to avoid a self-referential commit hash in this record.
