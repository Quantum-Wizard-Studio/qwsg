# Task History 012: Core Collector Framework

## 1. Task Metadata

- Task ID: `012`
- Title: `Core Collector Framework`
- Task slug: `core-collector-framework`
- Date: `2026-07-21` UTC
- Status: `complete`
- Responsible agent: `Aikó/Codex`
- Human authority: `Project Owner`
- Preferred owner communication language: `Hungarian`
- Active prompt: `ai/prompts/012_CURRENT_TASK.md`
- Dependency: completed Task 011 and the canonical Inventory Architecture

## 2. Objective

Implement the only supported internal framework for identifying, registering, checking, executing, isolating, and integrating Inventory collectors while preserving the existing Task 011-compatible Inventory output.

## 3. Scope and Exclusions

Authorized targets are the Go collector and Inventory integration, their unit tests, architecture references, roadmap, Engineering History, and this record. The work includes the Collector contract, Registry, lifecycle, capability and dependency checks, standard result, deterministic execution, timeouts, cancellation, panic isolation, and migration of existing collectors.

No new Inventory domain, policy/configuration engine, API, Console, daemon, remediation, dynamic plugin, remote collector, dependency installation, service change, commit, or push is authorized.

## 4. Required Reading

Read before implementation:

- `AGENTS.md`
- `.agents/skills/qwsg-job/SKILL.md`
- `ai/prompts/012_CURRENT_TASK.md`
- `ai/core/00_PROJECT_PHILOSOPHY.md`
- `ai/core/01_CONSTITUTION.md`
- `ai/core/03_AGENTS.md`
- `ai/core/08_JOB_TEMPLATE.md`
- `ai/core/11_ENGINEERING_LIFECYCLE.md`
- `ai/core/12_INVENTORY_ARCHITECTURE.md`
- `ai/history/011_2026-07-21_core-inventory-architecture.md`
- relevant current Go implementation and tests

## 5. Starting State and Environment Verification

- UTC start verification: `2026-07-21T15:59:26Z`
- User: `attila` (`uid=1000`, `gid=1000`, supplementary group `nogroup`)
- Working directory: `<repository-root>`
- Git branch: `main`
- `bin/job --check`: passed for approved Task 012
- Task 011 prompt and history: both `complete`
- Inventory Architecture: `ai/core/12_INVENTORY_ARCHITECTURE.md` exists and its collector contract was inspected
- The worktree contains extensive inherited modified and untracked work from prior tasks. Task 011 records the same condition; it is preserved without cleanup, reset, or unrelated rewriting.
- Relevant directories and files are owned by `attila:nogroup`; setgid project directories and inspected ACL masks preserve owner/group write behavior. No permission or ACL change is authorized.
- Direct `go test ./...` could not initialize the default cache because `<user-cache>` is read-only in the managed environment. The canonical repository commands use isolated caches under `/tmp` and passed: `make fmt-check`, `make vet`, `make test`, and `make build`.

No material variance blocks the authorized work.

## 6. Snapshot

Verified snapshot: `ai/backups/20260721T160300Z_task012_core_collector_framework/repository-baseline.tar.gz`

- Scope: complete repository baseline except `.git`, prior `ai/backups`, and generated `build`; includes source, tests, configuration, AI documentation, prompts, histories, archives, and product documentation.
- SHA-256: `cd28241d5926eb68414b9edff9873050fd265c2b2d3ae07200ee5c394c126165`
- Size: `177378` bytes
- Ownership/mode: `attila:nogroup`, `0660`
- Integrity: gzip/tar listing succeeded and required source/document/config path classes were present.

## 7. Risk Assessment and Plan

| Risk | Rating | Mitigation |
| --- | --- | --- |
| Inventory output regression | High | Keep the existing `inventory.Category` projection and run full CLI/model tests. |
| Duplicate or unstable registration | High | Validate descriptors, reject duplicate IDs, and sort immutable registry snapshots. |
| Collector coupling | High | Express dependencies as IDs and resolve them only in the Registry. |
| Timeout/cancellation leaks | High | Derive bounded contexts, propagate cancellation, and require collectors to honor context. |
| One collector corrupts the run | High | Convert errors, availability failures, dependency failures, and panics into isolated results. |
| Resource expansion | Medium | Retain bounded command/file acquisition and finite descriptor limits; execute deterministically without unbounded concurrency. |
| Dirty-worktree rollback damage | High | Use the complete archive and exact-path restore only; no Git reset or broad cleanup. |

Smallest safe sequence: implement the contract and Registry; migrate current collectors; route application collection exclusively through the Registry; add focused tests; update governing documentation; run every required verification gate.

## 8. Rollback Plan

Preconditions: stop any QWSG process, verify the exact snapshot path and SHA-256 above, inspect the current diff, and obtain owner confirmation because rollback overwrites Task 012 work.

1. Extract the verified archive into a new temporary directory under `/tmp`, never directly over the repository.
2. Compare its manifest and expected Task 012 targets with the live worktree.
3. Copy back only the exact files created or modified by Task 012; remove only Task 012-created paths after explicit confirmation.
4. Re-run `bin/job --check`, `make fmt-check`, `make vet`, `make test`, `make build`, permission/ACL inspection, and bounded Git diff inspection.

No wildcard deletion, repository reset, unrelated cleanup, or direct full-tree archive extraction over the live worktree is permitted.

## 9. Work performed

Implemented the canonical internal framework in `internal/collector/framework.go`:

- `Descriptor` declares stable identity, implementation and contract versions, Inventory compatibility, capability, supported platforms, privilege, finite timeouts and output limits, dependencies, and sensitivity classes.
- `Request`, `Availability`, `Warning`, and `Result` provide the lifecycle inputs and standard structured execution outcome. Collector health remains separate from host health.
- `Registry.Register` validates descriptors and rejects duplicate IDs.
- Registry snapshots are immutable during a run and planned in deterministic dependency-aware order: dependencies precede dependents and stable collector ID orders otherwise equivalent work.
- Missing or unusable dependencies and unavailable/unsupported capability checks produce isolated structured results rather than invoking the collector.
- Each invocation receives a derived bounded context and absolute deadline. Parent cancellation and collector timeout are distinguished.
- Collector and availability-check panics are recovered as isolated internal collector errors. A failure cannot invalidate unrelated completed collectors.

Migrated existing collectors in `internal/collector/collector.go` to expose descriptors, availability, and standard results. `DefaultRegistry` registers every existing inventory collector. `internal/app` now accepts a Registry and obtains every category exclusively from Registry results; direct collector iteration was removed. The CLI initializes the default Registry and retains the existing schema `1.0`, category projection, redaction policy, and complete/partial/failed exit behavior.

Added focused framework tests for registration, duplicate rejection, stable execution, dependency planning, availability/dependency failures, timeout, parent cancellation, and panic isolation. Existing collector, application, Inventory, runner, and CLI tests remain passing.

## 10. Verification

- `bin/job --check`: passed for Task 012 after completion-state update.
- Registry operation and duplicate rejection: focused tests passed.
- Deterministic execution and dependency ordering: focused tests passed for 20 uncached repetitions.
- Availability and dependency checks: focused tests passed.
- Timeout, cancellation, collector panic, and availability-check isolation paths: focused tests passed.
- Existing Inventory compatibility: application, Inventory, and CLI packages passed for five uncached repetitions.
- Canonical live output: `./build/qwsg inventory` emitted validated schema `1.0` JSON and returned the permitted partial exit `2` in this service-restricted environment.
- `make fmt-check`: passed.
- `make vet`: passed.
- `make test`: all packages passed.
- `make build`: passed and produced `build/qwsg`.
- Snapshot SHA-256 remained `cd28241d5926eb68414b9edff9873050fd265c2b2d3ae07200ee5c394c126165`; archive integrity passed.
- `git diff --check`: passed.
- Target ownership, mode, and ACLs were inspected: files remain `attila:nogroup`, `0660`, with the inherited effective owner/group write ACL pattern.
- Scope inspection found no dependency installation, service mutation, new Inventory domain, policy/configuration/API/UI/daemon/remediation/plugin/remote-collector work, commit, push, or unrelated cleanup.

The repository remains intentionally dirty with inherited work described in Section 5 plus Task 012 changes and its verified backup. Git's ordinary diff summary does not enumerate untracked implementation files; verification therefore used explicit target inspection and the full Go test/build gates as well as bounded Git checks.

## 11. Documentation Updates

- Updated `ai/core/04_ARCHITECTURE.md` with the implemented Registry boundary and legacy Inventory projection relationship.
- Updated `ai/core/13_ROADMAP.md` with the completed Task 012 outcome and retained future-task boundary.
- Updated `ai/core/07_ENGINEERING_HISTORY.md` with a concise Task 012 milestone.
- Updated `ai/prompts/012_CURRENT_TASK.md` to `complete` after implementation and verification.
- Finalized this independent chronological record.

## 12. Delivery and Completion State

Delivery result: `complete`.

All authorized deliverables exist, existing collectors use the Registry framework, Inventory compatibility is preserved, required documentation and rollback evidence are current, and mandatory verification passed. No unresolved architectural inconsistency or implementation blocker remains inside Task 012 scope.

No commit or push was performed. No next task was prepared or started because Task 012 does not authorize lifecycle rotation.
