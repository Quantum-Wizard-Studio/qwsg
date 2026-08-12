# Task History 037: Approved Task

## Task metadata

- Task ID: `037`
- Task slug: `operational-guardian-service`
- Status: `complete`
- Date generated: `2026-08-10` UTC
- Human authority: Project Owner
- Preferred owner communication language: Hungarian
- Related prompt: `ai/prompts/037_CURRENT_TASK.md`

## Lifecycle state

The Builder-installed task was loaded through the canonical `job` workflow,
implemented within its approved boundaries, validated on the supported host,
and returned to canonical idle.

## Starting state

- Task 037 started through the canonical `job` workflow on 2026-08-10 UTC.
- Repository root, `main`, HEAD `0a8a5c7e722495b8c5eb425bca5b2d2413aaa175`, canonical HTTPS origin, `0/0` relationship, empty index, sole active Task 037 prompt/history pair, and completed Task 036 baseline were verified.
- Ubuntu 24.04, Go 1.26.5, GNU Make 4.3, systemd 255, and a running user manager were verified. The engineering shell is sandboxed; system-manager access requires the separately bounded acceptance authority.
- Framework 1.0.0, lifecycle, diverted-task audit, Task 030–036 contracts, required governance, source/import boundaries, ownership, modes, ACLs, and all pre-existing Owner-owned working-tree content were inspected.
- The implementation boundary is frozen: Runtime Service remains recurrence owner, Runtime remains cycle owner, Scheduler retains scheduling persistence/locking, `internal/guardian` is operational wiring/checkpoint/lock only, and shared typed projection belongs to the existing application boundary.

## Snapshot

Verified implementation snapshot: `/tmp/qwsg-task037-implementation-20260810TEtoqqo`. It contains 37 exact pre-change target payloads, six absence records, repository/Git identity, permissions, ACLs, checksums, and bounded restore instructions. Every SHA-256 manifest entry passed before implementation.

## Work performed

Implementation added the narrow `internal/guardian` operational adapter,
private nonblocking operation lock, strict integrity-protected atomic recovery
checkpoint, canonical configuration composition, Scheduler/Pipeline/Runtime
wiring, empty Notification policy, exact Runtime Service State evidence seam,
shared typed operator projection, lifecycle/current-state publisher, and
generation-correlated demotion-only exit reporter.

`qwsg guardian run` is one foreground process with SIGINT/SIGTERM shutdown and
no internal supervisor. Explicit `observe` shares its writer lock and reports
`guardian_active` instead of racing. Presentation Model and Current Operator
State 1.2 add degraded/unavailable lifecycle truth with 1.0/1.1 reader
compatibility; stale running/stopped evidence becomes unavailable. The Console
remains a read-only consumer.

The supported systemd user unit, bounded Make installation target, operational
architecture, installation/rollback guidance, product architecture, system
map, functional specification, roadmap and user-facing README were updated.
The unit runs as the ordinary user with 0700 state, `NoNewPrivileges`, private
temporary storage, read-only system/home protection plus one state exception,
bounded restart/stop policy, `MemoryMax=128M`, `TasksMax=32`, and
`CPUQuota=25%`.

## Verification

- `go test ./...`, `go vet ./...`, formatting, `git diff --check`, Framework
  1.0.0, all framework/lifecycle/diversion/next-task/Builder suites, and
  `bin/job --check` passed before archival.
- `systemd-analyze verify --user` passed for the final production unit. A
  staged `make install-service` installed only the binary and 0644 user unit.
- An isolated foreground run completed more than five recurrence boundaries,
  published full ten-stage Guardian coverage, was consumed by a separate bare
  `qwsg`, and rejected a competing `observe` as `guardian_active`.
- A uniquely named real systemd user acceptance unit started, performed
  recurring cycles, published truthful `Guardian: degraded` for the live
  host's partial evaluation, restarted with checkpoint/Scheduler recovery, and
  stopped with systemd `inactive`, `Result=success`, and Console
  `Guardian: stopped`. Invalid configuration failed without a running claim.
- The systemd run stayed within the approved limits (observed peak 48.7 MiB,
  9 tasks, 9.428 seconds CPU over the measured multi-cycle interval, no
  restart loop). The isolated unit link and state root were removed afterward.
- Acceptance exposed and resolved two pre-release defects: systemd-created
  state needed explicit `StateDirectoryMode=0700`, and `ExecStopPost` options
  needed separate arguments for environment expansion.

## Rollback

Restore the 37 payloads and six absence records from
`/tmp/qwsg-task037-implementation-20260810TEtoqqo` using its bounded manifest
and restore instructions. Stop/disable the Guardian before binary/unit
rollback, reload the user manager, preserve private state for compatibility
assessment, and never recursively remove a prefix or unrelated user state.

## Completion state

`complete`
