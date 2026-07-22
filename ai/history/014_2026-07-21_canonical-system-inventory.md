# Task History 014: Canonical System Inventory v1

## Task metadata

- Task ID: `014`
- Task slug: `canonical-system-inventory`
- Status: `complete`
- Date opened: `2026-07-21` UTC
- Responsible agent: Aikó/Codex
- Human authority: Project Owner
- Preferred owner communication language: Hungarian
- Related prompt: `ai/prompts/014_CURRENT_TASK.md`
- Dependencies: completed Tasks 012 and 013; canonical Inventory Architecture

## Starting state

The approved Task 014 prompt and matching history passed `bin/job --check` and `next-task.sh --check`. Task 012 Collector Framework contracts are unchanged: Registry registration, deterministic dependency planning, availability checks, timeout/cancellation, panic isolation, and structured Results all passed their existing tests. Task 013 lifecycle tooling passed 33 builder assertions and 23 compatibility lifecycle assertions.

Baseline `go test ./...`, `go vet ./...`, `make fmt-check`, and `git diff --check` passed using writable Go caches under `/tmp`. Go version is `go1.26.5 linux/amd64`. Relevant source files are owned by `attila:nogroup`, mode `0660`, with inherited owner/group-write ACLs. The worktree is intentionally dirty with inherited earlier-task work; those changes are preserved and are not attributed to Task 014.

## Snapshot

- Location: `ai/backups/20260721T_task014_canonical_system_inventory/`
- Archive: `repository-before-task014.tar.gz`
- Scope: complete worktree excluding `.git`, prior `ai/backups`, and generated `build`; includes all Go sources, Collector Framework and Registry, lifecycle scripts, engineering documentation, tests, configuration, active prompt/history, and recorded Git status.
- Integrity: recorded in `SHA256SUMS` and passed `sha256sum -c`.
- Readability and completeness: full archive listing succeeded; representative collector, lifecycle, and active-task paths were extracted from the listing successfully.
- Retention: preserve through owner acceptance and all dependent inventory work.

## Risk assessment

| Risk | Rating | Mitigation |
|---|---|---|
| Inventory 1.0 regression | High | Preserve the legacy envelope and category projection; add the canonical model as an additive validated representation. |
| Nondeterministic enumeration | High | Sort collectors, layers, resources, facts where represented as arrays, interfaces, mounts, and devices by stable canonical keys. |
| Distribution-specific parsing | High | Use bounded Linux kernel interfaces and tolerant fixture-tested parsers; represent unavailable evidence explicitly. |
| Privacy exposure | High | Hash host/device identities and redact hostnames, interface names, addresses, mount paths, and raw device names. |
| Collector failure propagation | High | Retain Registry isolation, finite timeouts, cancellation, and dependency semantics unchanged. |
| Competing host models | High | Make the canonical system inventory an additive authoritative projection governed by `12_INVENTORY_ARCHITECTURE.md`, not an independent discovery path. |

## Planned work

Extend the inventory domain with canonical layers, resources, facts, collector summaries, deterministic validation, and a loss-aware legacy Category projection. Complete the nine approved Linux collectors, keep all registration through `DefaultRegistry`, add fixture-driven collector/model/integration tests, preserve legacy behavior, and update architecture, developer, history, and delivery documentation.

## Rollback

Preconditions: stop QWSG execution, verify the exact snapshot hash and archive readability, inspect the current diff, and obtain Project Owner confirmation because restoration overwrites Task 014 targets.

1. Verify `ai/backups/20260721T_task014_canonical_system_inventory/SHA256SUMS`.
2. Compare the archive with the exact Task 014 source, test, documentation, prompt, and history targets.
3. Restore the complete pre-task worktree represented by the archive, as explicitly required by Task 014; do not use Git reset or delete unrelated/unverified paths.
4. Verify restored hashes and recorded repository status.
5. Re-run Go tests/vet, lifecycle tests/checks, format validation, and `git diff --check`.

The snapshot is the sole complete rollback source. No partial rollback is accepted by the active task.

## Work performed

Implemented the additive canonical domain model in `internal/inventory/canonical.go`: full snapshot identity and timing, producer, deterministic Collector Results, canonical layers, resources, typed facts, relationships, issues, redactions, and metadata. Validation enforces envelope agreement with Inventory 1.0, timestamp order, deterministic unique collector/layer/resource identities, required resource semantics, prohibited-secret rejection, and relationship referential integrity.

Completed the approved Host, Operating System, Kernel, CPU, Memory, Storage, Filesystem, Network, and Virtualization collectors in `internal/collector`. All use bounded ordinary-user Linux evidence, context cancellation, finite Registry timeouts, privacy-safe identifiers or redaction, deterministic item ordering, and structured Results. Filesystem depends on Storage and Virtualization depends on Host, exercising preserved dependency-aware orchestration. Existing capability, service, and component collectors remain for backward compatibility without Task 014 scope expansion.

The Registry now enforces every descriptor's normalized output byte limit after collection while retaining timeout, cancellation, availability, dependency, and panic isolation. `internal/app` assembles the canonical and legacy Inventory 1.0 representations from the exact same Result set, validates their envelope agreement, and introduces no second discovery path.

Added parser, collector-set, dependency, privacy-ID, output-limit, deterministic JSON, canonical validation, secret/relationship adversarial, envelope-consistency, and end-to-end integration tests. Updated architecture, system map, roadmap, changelog, developer guidance, and English/Hungarian user documentation.

## Verification

Implementation verification passed before owner handoff:

- All nine Task 014 collectors returned `available` during a live non-root Linux CLI run.
- The live CLI emitted parseable JSON containing the unchanged Inventory `1.0` envelope and `canonical_inventory.schema_name == "qwsg.inventory"`; overall exit `2` truthfully reflected only the pre-existing optional service collector's unavailable systemd environment.
- Required categories, deterministic layer/resource order, canonical/legacy envelope agreement, and privacy redactions passed automated assertions.
- `go test ./...`, `go test -race ./...`, and `go vet ./...`: passed.
- `make fmt-check`: passed.
- `bash ai/tests/test-task-builder.sh`: 33 assertions passed.
- `bash ai/tests/test-next-task.sh`: 23 assertions passed.
- `bin/job --check` and `ai/scripts/next-task.sh --check`: passed for Task 014 active state.
- `git diff --check`: clean.
- Snapshot SHA-256, readability, and representative-path completeness checks: passed.
- Scope review confirmed no health/rule/alert/remediation/process/container-monitoring/persistence/API/UI/configuration/scheduling/history/comparison implementation.

Closure verification after explicit Project Owner acceptance reconfirmed all lifecycle, snapshot, Go, race, vet, formatting, repository consistency, live canonical JSON, privacy, deterministic-ordering, and required-collector gates. The Task 014 implementation has no known technical verification failure.

## Documentation updates

- Added `docs/architecture/CANONICAL_SYSTEM_INVENTORY_V1.md`.
- Added `docs/development/CANONICAL_SYSTEM_INVENTORY.md`.
- Added English and Hungarian user guides under `docs/user/`.
- Updated `ai/core/04_ARCHITECTURE.md`, `05_SYSTEM_MAP.md`, `07_ENGINEERING_HISTORY.md`, `12_INVENTORY_ARCHITECTURE.md`, and `13_ROADMAP.md`.
- Updated `README.md`, `CHANGELOG.md`, and `docs/development/SLICE_1_DEVELOPMENT.md`.
- Updated this record and the active prompt through the owner-acceptance gate to their truthful final `complete` state.

## Project Owner acceptance

The Project Owner explicitly accepted the Task 014 implementation after reviewing the delivery outcome. This acceptance satisfies the final governance gate in the approved Completion Criteria and authorizes formal closure of the prompt and independent history.

Acceptance preserves the disclosed operational observation: the live CLI reports aggregate `partial` only because the pre-existing optional service collector cannot access systemd in the verification environment. This is a truthful, isolated compatibility-collector result. All nine Task 014 canonical collectors—Host, Operating System, Kernel, CPU, Memory, Storage, Filesystem, Network, and Virtualization—executed successfully and returned `available`.

## Unresolved issues

None within Task 014 scope. The optional legacy service collector's environment-specific systemd unavailability remains explicitly disclosed and does not affect any of the nine canonical collectors. The inherited dirty worktree remains present; no commit or push was performed during implementation or closure.

## Completion state

`complete`

Engineering delivery result: `complete`.

The implementation, documentation, rollback evidence, every agent-executable verification gate, and explicit Project Owner acceptance are complete. Canonical System Inventory v1 is the authoritative internal host-information subsystem, all nine required canonical collectors are operational, and Inventory 1.0 compatibility is preserved. The disclosed aggregate CLI `partial` status is caused exclusively by the pre-existing optional service collector's systemd environment. The next task was not generated, approved, or started.

## Final Delivery Report

Task 014 delivered the production-oriented Canonical System Inventory v1 subsystem, its nine approved Linux collectors, canonical aggregation and validation contracts, Registry integration and output-limit enforcement, legacy Inventory 1.0 compatibility, privacy controls, deterministic behavior, automated unit/integration/adversarial tests, architecture and developer documentation, bilingual user guidance, verified snapshot and rollback evidence, and chronological engineering records.

Final verification passed and the Project Owner explicitly accepted the implementation. No excluded engine, remediation, persistence, API, UI, scheduling, historical comparison, or non-Linux capability was introduced. Repository changes remain uncommitted at formal closure, and no next task has been created or started.
