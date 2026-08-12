# Task History 031: Canonical Runtime Service

## Task metadata

- Task ID: `031`
- Task slug: `canonical-runtime-service`
- Status: `complete`
- Date generated: `2026-08-07` UTC
- Human authority: Project Owner
- Preferred owner communication language: Hungarian
- Related prompt: `ai/prompts/031_CURRENT_TASK.md`

## Lifecycle state

The Builder installed the approved task, the canonical `job` workflow executed
it, every completion gate passed, and the prompt was archived without a
successor to return production to canonical idle.

## Starting state

- Repository root: `/home/qws/web/qwsg.quantumwizard.hu/qwsg`.
- Branch and HEAD: `main`, `0a8a5c7e722495b8c5eb425bca5b2d2413aaa175`.
- Canonical `origin` HTTPS URL matched Framework configuration; local/remote
  relationship was `0 0`.
- Task 031 was the sole active approved task and Task 030 the unique latest
  completed archive/history baseline.
- `internal/runtimeservice`, its tests, the Runtime Service architecture
  document, and Task 031 archive were absent.
- Framework, lifecycle, diverted-test, full tests, race, vet, format, and Git
  diff checks passed before target changes using writable `/tmp` Go caches.
- The index was empty. Existing unstaged and untracked Owner-owned work from
  Tasks 025–030, QWCS, Builder sources, and historical backups was preserved.

## Snapshot

- Implementation snapshot:
  `/tmp/qwsg-task031-implementation-20260807T115031Z/targets.tar`.
- SHA-256:
  `4cdd556cff95cb0b51aa6b4c13852292349cf09bdbf96770ef299ad39f89b0f3`.
- The readable archive contains exact pre-change working-tree versions of the
  Runtime package, directly affected permanent documentation, prompt, and
  history. Separate evidence records Git status, untracked inventory, HEAD,
  remotes, and the working-tree patch.
- Absence was verified for `internal/runtimeservice`,
  `docs/architecture/CANONICAL_RUNTIME_SERVICE.md`, and the Task 031 archive.
  Checksums, archive inventory, permissions, ACLs, and collision preconditions
  were verified before implementation.

## Work performed

- Added `internal/runtimeservice/engine.go` with Service Definition, State,
  Input, Event, Evidence, and Result 1.0 contracts; strict validation and JSON
  decoding; canonical content identities; lifecycle/outcome tokens; UTC and
  duration/counter/resource bounds.
- Implemented one explicitly invoked continuous loop over a narrow
  `RuntimeRunner`. It performs an immediate first cycle, fixed-rate nominal
  arithmetic, sequential calls, no overlap, no catch-up burst, compressed
  missed-interval evidence, and bounded per-cycle Runtime contexts.
- Mechanically forwards only Runtime-owned `NextState`, `FinalAlertState`, and
  `FinalNotificationQueue`; all immutable seed fields remain unchanged.
- Added injected Clock, Waiter, Runtime Runner, and synchronous Evidence Sink
  contracts plus local standard-library clock/timer adapters. No event history
  or persistence is retained.
- Added a separate signal adapter mapping only SIGINT and SIGTERM to context
  cancellation with scoped registration and cleanup.
- Added comprehensive no-sleep tests covering deterministic fixed-rate calls,
  exact deadlines, state handoff, missed boundaries, pre-start and active-cycle
  cancellation, signal registration, sink refusal, Runner failure privacy,
  identity tampering, limits, canonical JSON, and Event/Evidence correlation.
- Created `docs/architecture/CANONICAL_RUNTIME_SERVICE.md` and updated the
  Runtime, product, functional, roadmap, system-map, project architecture,
  README, and milestone boundaries.
- No Scheduler/Command/Pipeline/Health/Rule/Policy/Report/Alert/Notification
  package is imported by production Service code. No persistence, systemd,
  installation, supervision, monitoring, provider, API, remediation, remote,
  AI, package, deployment, or infrastructure mutation was added.

## Verification

- `make build`: PASS.
- Full `go test ./...`: PASS, including Runtime Service and Runtime regression.
- Repository-wide `go test -race ./...`: PASS.
- `make vet` and `make fmt-check`: PASS.
- `make engineering-test`: PASS — Framework 21 assertions, diversion 36,
  lifecycle 28, Builder 38.
- Framework 1.0.0, active-task/lifecycle checks, and diverted-test audit: PASS.
- Production import audit: only standard-library packages and
  `internal/runtime`; no downstream canonical engine import.
- Architecture/privacy audit: no network, database, process execution,
  persistence, monitoring, provider, interface, remediation, systemd, remote,
  AI, or installation boundary; raw dependency errors are excluded by fixed
  failure tokens and a regression test.
- Snapshot checksum/readability/inventory, permissions, ownership, ACLs,
  absence records, and rollback collision guards: PASS.
- `git diff --check` and `git diff --cached --check`: PASS; staged paths remained
  empty. No commit, push, package, installation, resident process, deployment,
  or release occurred.

## Rollback

Verify the snapshot SHA-256 and readable inventory, then compare every current
target with the snapshot and final Task 031 record. Restore only exact
pre-existing Runtime/documentation/lifecycle targets. Remove new Service files
and architecture documentation only if their recorded pre-task absence and lack
of later Owner edits remain proven. Never use broad reset, checkout, restore,
clean, wildcard deletion, broad process killing, or overwrite unrelated Owner
work. After bounded restoration rerun focused/full/race/vet/format, Framework,
Builder, lifecycle, diverted-test, architecture/privacy/import, permissions,
ACLs, Git diff, and snapshot integrity checks.

## Completion state

`complete`

Disclosed limitations are the approved boundaries: state handoff is in-memory
only; there is no restart/crash continuity, persistent evidence, configuration
activation/reload, system installation/supervision, automatic restart,
watchdog, production provider, monitoring/product health, diagnostics, or
support claim. Graceful active-cycle shutdown depends on the Runtime Runner
honoring its existing context contract. These are future Version 1.0 gates, not
unresolved defects in Task 031.
