# Task History 045: Approved Task

## Task metadata

- Task ID: `045`
- Task slug: `qwsg-setup-configuration-foundation`
- Status: `complete`
- Date generated: `2026-08-12` UTC
- Human authority: Project Owner
- Preferred owner communication language: English
- Related prompt: `ai/prompts/045_CURRENT_TASK.md`

## Lifecycle state

The Engineering Task Builder generated and transactionally installed this matching prompt/history pair from validated structured owner input. Explicit approval was recorded; implementation has not started.

## Starting state

- Implementation authority was received on `2026-08-12` UTC for the exact installed Task 045 prompt.
- Approved Builder source SHA-256 was reverified as `cadd0da6263b73e5b85defd0b53855c64e86e5dceff827d23ed12c28f52263ee`.
- Repository root, `main`, canonical origin, active prompt/history identity, empty index, clean tracked tree, and `0/0` upstream relationship were verified at `c427bee165db17bb9416a46b44965bb54ff530a6`.
- `v1.0.0` remained annotated object `7abc83da6185199606c2d76ac7d3504ddd78cf68`, peeled target `177535e44b2ce5ed9efd73ab0793ffe6881f0cd6`; the retained final archive remained SHA-256 `edfba7366adf2c1ce0a8ce56369bb0dc5ad11326c4e3d1e301625a5313292fa5`.
- The excluded Owner draft remained SHA-256 `c83a0b2e5269de51acd9485b44fbe4f752e1f7f172c180e4126f26e91fdf8a73` and was not read or modified.
- Pre-change build, focused tests, full tests, race tests, vet, formatting, Framework, Builder, lifecycle, and diversion baselines passed.

## Snapshot

- Verified implementation snapshot: `/tmp/qwsg-task045-implementation-20260812T.FB2cFR` (mode `0700`).
- It contains a complete tracked-source archive, separate active lifecycle capture, planned-target and absence manifests, Git/environment/tool/process evidence, ACLs, immutable release identities, Owner-draft hash/metadata-only evidence, SHA-256 manifest, and bounded rollback instructions.
- Both archives listed successfully and every snapshot manifest entry verified before product-source modification.

## Work performed

- Audited Configuration Contract 1.0, every CLI/environment input, Guardian scheduling/checkpoint integration, XDG state resolution, systemd environment and sandbox, packaging sample, installer/uninstaller, and operator documentation.
- Confirmed configuration inputs before Task 045: immutable Source Records; compiled model defaults; Guardian `--config-source`, `--interval`, and `--cycle-timeout`; command `--store`, `--format`, and selectors; `QWSG_STATE_DIR`, `XDG_STATE_HOME`, `HOME`, `QWSG_STORE`, `QWSG_FORMAT`, `QWSG_LOCALE`, and systemd `INVOCATION_ID`. The principal defects were absent discovery/persistence/setup, unsafe ordinary config reads, split schedule/runtime authority, generic diagnostics, and configuration resolution after checkpoint mutation.
- Added `internal/configurationstore`, selecting `$XDG_CONFIG_HOME/qwsg/config.json` or `$HOME/.config/qwsg/config.json`, with explicit `--config` replacement, bounded reads, exact current-user `0700/0600` checks, symlink/special-file rejection across every component, identity revalidation, and primary-source enforcement.
- Added atomic same-directory persistence with private temporary files, complete write, file sync, rename, and parent sync. Injected failures proved pre-rename preservation and complete-only behavior on directory-sync uncertainty.
- Hardened Configuration Source JSON by rejecting duplicate object field names in addition to existing unknown-field, trailing-data, schema/model, identity, reference, and bounds checks.
- Added `qwsg setup` and `qwsg config show|validate|get|set`. Setup supports interactive confirmation, `--accept-defaults`, repeated idempotent use, bounded `--set key=value`, explicit paths, and human/JSON output. The mutation registry is limited to `locale`, `time_zone`, `snapshot_retention`, `guardian.interval`, and `guardian.cycle_timeout`.
- Unified Guardian recurrence with Effective Configuration. Discovered primary values and temporary duration flags now resolve through the same Source precedence/provenance mechanism consumed by Runtime Service. Inconsistent interval/timeout is rejected before persistence or start.
- Moved Guardian configuration resolution before lock/checkpoint mutation. A valid configuration identity change starts a fresh Runtime/Alert/Notification state epoch while retaining Inventory and Current Operator evidence; invalid configuration changes no state and cannot claim health.
- Added the systemd `ExecCondition` validation gate. Live local user-systemd evidence showed valid input starts with `NRestarts=0`; malformed input is skipped due to `exec-condition`, remains inactive, and has `NRestarts=0`. Existing sandbox/resource directives were unchanged.
- Kept release installation noninteractive and added only per-user next-step guidance. The packaged sample is the exact setup-generated default Source Record and remains inactive documentation.
- Documented one canonical setup/configuration reference, activation architecture, precedence, compatibility, troubleshooting, security, upgrade, Guardian, CLI, and installer relationships.
- Recorded the accepted future product boundary as an inert optional recipient extension: Community basic local email has one administrator recipient without account/API/subscription/QWS remote service; Pro has no entitlement-imposed recipient cap subject to global safety bounds. No delivery, credential backend, or entitlement behavior was implemented.
- Created no Task 046, release, tag, commit, push, publication, external call, email transport, credential implementation, or remote-service feature.

## Verification

- Pre-change `make build`, focused configuration/Guardian/operator/CLI tests, `go test ./...`, `go test -race ./...`, `go vet ./...`, formatting, Framework, Builder, lifecycle, and diversion baselines passed.
- Focused Task 045 tests passed for absent/default configuration, first/repeated/interactive/noninteractive setup, typed get/set, JSON/human output, malformed/unknown/duplicate fields, explicit absence, inconsistent timing, symlinks behind missing parents, permissive files/directories, special targets, deterministic identity, inert one/many recipient representation, atomic fault injection, Console invalid-config refusal, and Guardian pre-state validation.
- Final full `go test ./...`, full race suite, `go vet ./...`, `make build`, `make fmt-check`, and `git diff --check` passed.
- `systemd-analyze verify --user packaging/systemd/qwsg-guardian.service` passed. Actual host-local user-systemd acceptance passed for valid start, bounded resource state, Console consumption, graceful stop, invalid-config skip, and zero restart churn. This was not an external clean host.
- Isolated staged-package acceptance root: `/tmp/qwsg-task045-staged-acceptance.or9axS`. Release-style build, archive sidecar, internal manifest, static binary, packaged LICENSE, install, setup, validation, two observations, Console, foreground Guardian run/stop, modes, and uninstall preservation passed.
- Two independent non-public temporary builds were byte-identical. Their temporary archive SHA-256 was `9798315088488dfdcf4f25f55b4451cc96bfd9cab4ad80e14192cb011be81612`; it embeds development baseline `c427bee165db17bb9416a46b44965bb54ff530a6`. This is acceptance evidence, not the immutable published artifact.
- Live read-only observation acceptance passed with truthful partial host inventory. Static scans found no new network/process-execution code, credentials, private keys, or raw secret output in Task 045 implementation.
- Framework tests passed (21 assertions), lifecycle tests passed (28), diversion tests passed (36), Builder tests passed (38), test-task audit passed, and `bin/job` lifecycle validation passed.
- Bounded rollback simulation passed at `/tmp/qwsg-task045-rollback-sim.scRpcF`, reproducing the exact pre-task tracked tree after restoring only changed targets and deleting only verified Task 045-created files.
- The published archive remained SHA-256 `edfba7366adf2c1ce0a8ce56369bb0dc5ad11326c4e3d1e301625a5313292fa5`; `v1.0.0`, its annotated object/target, LICENSE, and Task 044 evidence remained unchanged.
- Owner draft `docs/architecture/QWCS_MIGRATION_BLUEPRINT.md` remained SHA-256 `c83a0b2e5269de51acd9485b44fbe4f752e1f7f172c180e4126f26e91fdf8a73`, with its original mode/owner/ACL and no content access or modification.

## Rollback

Use `/tmp/qwsg-task045-implementation-20260812T.FB2cFR` and its `ROLLBACK.md`. Stop exact acceptance writers, verify recorded current hashes and immutable identities, extract the tracked archive only into a new private temporary directory, restore only enumerated changed regular files, and remove only enumerated new files whose hashes match this delivery. Broad reset, checkout, clean, wildcard deletion, tag changes, release changes, and Owner-draft operations are forbidden. The snapshot manifest and both archives verify successfully.

## Completion state

`complete`

Task 045 met its approved objective and exclusions. Automated/sandbox, staged-package, and actual host-local user-systemd acceptance passed. No external clean-host acceptance was performed or claimed. The remaining limitations are deliberate: no email delivery, SMTP credentials, entitlement enforcement, API/licensing/billing, remote service, fleet, or web UI; no hot reload; configuration changes take effect on the next Guardian start; and only the five documented public keys are mutable through Task 045 CLI. The repository is ready for a separately authorized next task after canonical idle closure; Task 046 was not created.
