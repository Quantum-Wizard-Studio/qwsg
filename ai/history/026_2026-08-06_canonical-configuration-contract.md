# Task History 026: Canonical Configuration Contract

## Task metadata

- Task ID: `026`
- Task slug: `canonical-configuration-contract`
- Status: `complete — Canonical Configuration Contract and all required gates passed`
- Date generated: `2026-08-06` UTC
- Human authority: Project Owner
- Preferred owner communication language: English
- Related prompt: `ai/archive_prompts/026_2026-08-06_canonical-configuration-contract.md`

## Lifecycle state

The Engineering Task Builder generated and transactionally installed the approved prompt/history pair. The Project Owner invoked `job` after installation. Task 026 executed as the sole active approved task under Framework 1.0, with QWCS safe-forward-progress and minimal-intervention principles used as engineering guidance.

## Starting state

Task 026 started on `2026-08-06` UTC at the verified QWSG repository root, branch `main`, HEAD `0a8a5c7e722495b8c5eb425bca5b2d2413aaa175`. The configured `origin` URL matched policy and local `main` was equal to `origin/main`. Task 025 was the latest completed canonical implementation baseline. Framework, lifecycle, test-task, Go, vet, and formatting validation passed. Existing unstaged and untracked Task 025, QWCS architecture, backup, and task-source content was recorded and preserved.

## Snapshot

The verified rollback-capable implementation snapshot is `/tmp/qwsg-task026-implementation-Pl2YcRoi`. It contains the bounded existing-target archive, exact pre-task Rule Engine blob, absence records for new implementation, architecture, and archive targets, starting Git and lifecycle evidence, baseline validation output, deterministic SHA-256 checksums, archive inventory, and guarded restore instructions. All checksums and archive-readability checks passed.

## Work performed

- Added `internal/configuration` as a pure model, normalization, validation, resolution, identity, provenance, and serialization boundary with no filesystem, clock, execution, network, secret-resolution, or host behavior.
- Defined `qwsg.configuration-source/1.0`, Canonical Configuration Model 1.0, `qwsg.effective-configuration/1.0`, and Schedule Definition 1.0.
- Implemented exact source precedence: built-in default, primary local, activated local override, and command-specific temporary override.
- Implemented explicit whole-field replacement. Equal-precedence equal values merge source provenance; different values fail as conflicts without input-order tie-breaking.
- Added field-level provenance containing field, sorted source IDs and versions, source kind, numeric precedence, and selected, overridden, or equal-values-merged resolution.
- Added stable SHA-256 content identity and byte-stable canonical JSON for normalized sources and Effective Configuration.
- Added strict JSON decoding, unsupported-version behavior, unknown-key rejection, immutable deep-copy normalization, deterministic collection ordering, duplicate detection, reference validation, tamper detection, and bounded inputs.
- Added typed configuration for instance, locale, time zone, checks, targets, Rule Definitions, Policy Profiles, snapshot retention, schedules, execution timeout, concurrency, retries, reports, secret references, and inert optional extensions.
- Defined interval and calendar Schedule contracts with explicit IANA-style time-zone basis, daylight-saving behavior, priority, misfire policy, overlap policy, execution timeout, retry reference, Check references, and Command profile applicability.
- Added typed secret references with no secret-value field. Required unsupported extensions fail; optional extensions remain inert data.
- Added `rule.ValidateDefinitions` as a validation-only Rule-owned contract so malformed or unsupported Rule Definitions fail at configuration time without evaluating Health evidence or changing Rule evaluation behavior.
- Integrated Effective Configuration into `internal/pipeline`. The pipeline resolves all configuration before stage execution and consumes configured retention, Rule Definitions, and Policy Profiles. Compatibility fields are projected through the same resolver; combining them with an Effective Configuration fails as ambiguous.
- Preserved the existing CLI, Command plans, stage ordering, default Rule and Policy behavior, Report behavior, and public contracts.
- Added focused unit and integration tests for precedence, provenance, source-order determinism, conflicts, schedules, calendar and interval bounds, cross-references, strict decoding, identities, versions, resource limits, extensions, secret rejection, caller immutability, Rule validation, pipeline consumption, compatibility, tampering, and ambiguity.
- Created `docs/architecture/CANONICAL_CONFIGURATION_CONTRACT.md` and updated only directly affected README, architecture, system map, roadmap, Engineering History, prompt, and task-history records.
- Implemented no Scheduler, due-work calculation, timer, polling, daemon, service, lock, worker pool, Alert Engine, notification, activation, file watcher, UI, API, secret backend, network operation, process execution, remediation, AI, or host mutation.

## Verification

Builder input, metadata, prompt/history identity, approval state, and lifecycle installation passed before implementation.

Final verification passed:

- `make build`;
- `make test` across sixteen Go packages, including the new Configuration package;
- `GOCACHE=/tmp/qwsg-go-cache GOMODCACHE=/tmp/qwsg-go-modcache go test -race ./...`;
- `make vet`;
- `make fmt-check`;
- focused Configuration, Rule, and Pipeline tests;
- deterministic source-order, canonical-JSON, identity, precedence, equal-conflict, provenance, schedule, reference, privacy, immutability, compatibility, and tamper tests;
- `ai/scripts/framework-check.sh --run-validations`;
- `make engineering-test` with 21 Framework, 36 diversion, 28 lifecycle, and 38 Builder assertions;
- active and final-idle lifecycle validation;
- diverted-test audit;
- documentation reference and architecture-boundary audits;
- source import and call audits proving no clock, timer, process, network, filesystem activation, Scheduler, notification, remediation, or host-operation dependency in Configuration;
- `git diff --check` and `git diff --cached --check`;
- snapshot checksums, archive readability, target inventory, absence records, pre-task Rule blob, and restore instructions.

An initial unconfigured `go test -race ./...` attempt encountered the sandbox's read-only default Go cache. No product test executed unsuccessfully. The identical race suite passed using the repository's configured writable cache paths.

## Rollback

Verify `/tmp/qwsg-task026-implementation-Pl2YcRoi/SHA256SUMS`, `targets.tar`, `rule-engine.before.go`, target lists, absence records, and `RESTORE.md`. Refuse rollback over later Owner work. Restore only exact pre-existing targets from the bounded archive, restore the separately captured Rule Engine blob, and remove only verified Task 026-created targets and archive after confirming their recorded pre-task absence. Preserve truthful lifecycle evidence and rerun framework, lifecycle, Go, race, vet, formatting, documentation, Git-state, and snapshot-integrity checks. Broad reset, checkout, restore, clean, wildcard deletion, and repository-wide extraction are prohibited.

## Documentation updates

- `docs/architecture/CANONICAL_CONFIGURATION_CONTRACT.md`
- `README.md`
- `ai/core/04_ARCHITECTURE.md`
- `ai/core/05_SYSTEM_MAP.md`
- `ai/core/07_ENGINEERING_HISTORY.md`
- `ai/core/13_ROADMAP.md`
- `ai/archive_prompts/026_2026-08-06_canonical-configuration-contract.md`
- `ai/history/026_2026-08-06_canonical-configuration-contract.md`

## Unresolved issues and recommendations

No mandatory Task 026 issue remains. User-facing configuration file syntax, persistence, discovery, activation, secret backends, editing, and hot reload remain separately scoped. Professional Scheduler is the recommended Task 027 and must consume Effective Configuration and Schedule Definition 1.0 without redefining configuration semantics.

## Git record

Task 026 started and completed at HEAD `0a8a5c7e722495b8c5eb425bca5b2d2413aaa175` on `main`, with local and `origin/main` equal at the verified baseline. Task-scoped changes remain unstaged. No commit, push, fetch, branch, tag, release, deployment, dependency installation, privilege change, or infrastructure operation occurred. Pre-existing unrelated owner content remained untouched.

## Delivery result

`complete` — Canonical Configuration Model 1.0, Configuration Source Record 1.0, Effective Configuration 1.0, Schedule Definition 1.0, deterministic precedence and conflicts, field provenance, typed secret references, resource bounds, canonical identity and serialization, pipeline consumption, architecture documentation, and engineering tests exist; all mandatory gates passed; Scheduler and all excluded operational capabilities remain unimplemented.

## Completion state

`complete — all authorized implementation, documentation, rollback, compatibility, separation, and validation gates passed`
