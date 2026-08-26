# Current Engineering Task 064: QWSG Native Update and Rollback System

## Task Metadata

- Task ID: `064`
- Task slug: `native-update-rollback-system`
- Status: `active`
- Date opened: `2026-08-26` UTC
- Human authority: Project Owner
- Owner or lead-developer communication language: English

## Title

QWSG Native Update and Rollback System


## Objective

Design, implement, validate and externally prove the smallest durable native QWSG update and rollback system, advancing development identity to QWSG 1.2.0-rc.1. An installed QWSG must be able to identify itself, discover and validate an authoritative newer public release, preserve user-owned configuration, credentials, compatible state and Guardian service intent, apply a verified update transactionally, recover automatically after a post-mutation failure where deterministic, and support an explicit safe rollback. Complete one bounded practical upgrade and rollback acceptance using the existing disposable VPS's officially released QWSG 1.1.0 installation as the preserved baseline, without reinstalling or reformatting that host.


## Scope

- Inspect and evolve the existing CLI, installer replacement/backup contract, release packaging, configuration and credential stores, Guardian/operator state, user-systemd integration, readiness, version provenance and public Forgejo distribution contract.
- Select and document the native update command architecture; prefer a coherent `qwsg update check`, `qwsg update`, status/plan support and an explicit rollback interface when source analysis confirms those are the safest operator semantics.
- Add strict semantic release identity comparison for newer, equal, older, unsupported and invalid identities, including prerelease/final ordering appropriate to QWSG policy.
- Discover releases only from the canonical public QWSG Forgejo source through bounded HTTPS behavior; select immutable version/tag assets and reject mutable, ambiguous or unsupported responses.
- Download into private unprivileged staging, enforce size/time/redirect bounds, verify checksum sidecar, safe archive layout/types, internal manifest, binary version/source provenance, required documentation and platform compatibility before privileged mutation.
- Implement a two-phase update transaction with a validated plan, narrow privileged installation boundary, private rollback backup and metadata, preservation of configuration, credential store, compatible state and service enabled/active intent, necessary Guardian stop/daemon-reload/start sequencing, post-install validation, and deterministic automatic restoration on eligible failure.
- Define versioned deterministic configuration/state migration interfaces and journals even when 1.1.0 to 1.2.0-rc.1 requires only a no-op/compatibility migration; migrations must be testable, failure-safe and rollback-aware.
- Add explicit rollback discovery/status/application with exact installed/backup identity checks, bounded retention and refusal of unsafe, incomplete, tampered or incompatible rollback data.
- Add deterministic local unit, integration and simulated-distribution tests covering success and failure injection without requiring a live public release during ordinary development.
- Construct only the private deterministic 1.2.0-rc.1 acceptance candidate needed for the authorized external test after source validation; freeze and transfer it through a privacy-safe path, then exercise authentic 1.1.0 -> 1.2.0-rc.1 update and rollback behavior on the existing disposable VPS.
- Record privacy-safe evidence, update product/operator/release-readiness documentation, perform validated Git integration, clean-fast-forward push and canonical Task 064 lifecycle closure.


## Out of Scope

- Do not modify, rebuild, retag, replace or rewrite QWSG 1.1.0 artifacts, `v1.1.0`, its Forgejo Release, checksums, evidence or historical acceptance records.
- Do not reinstall, reformat or otherwise manufacture a clean state on the existing test VPS; preserve its installed QWSG 1.1.0 baseline and persistent user data.
- Do not publish QWSG 1.2.0-rc.1, create or push a tag, create a Forgejo Release, upload a public asset, deploy or announce a new release.
- Do not implement unattended background auto-update, update channels, fleet orchestration, package-manager repositories, signing/PKI, arbitrary mirrors, delta packages, Windows/macOS support or production deployment unless separately authorized.
- Do not expose or persist credentials in source, chat, logs, evidence, process arguments or artifacts. Public update discovery/download must not use Forgejo credentials.
- Do not begin the separate QWSG Server Health / Assessment Engine direction and do not create Task 065.
- Do not broaden product architecture beyond the minimum durable native update, migration and rollback foundations required by this task.


## Authority Envelope

- **Task targets and boundaries:** Framework 2.0 Standard Execution Authority applies to the Task 064 objective across the QWSG CLI, update/download/trust/version/migration/transaction components, installer/release packaging needed by the native updater, deterministic tests and simulated distribution fixtures, 1.2.0-rc.1 development identity/notes, task-scoped documentation/evidence and lifecycle. Multiple inspect/design/implement/fail/classify/correct/retest/refactor cycles remain inside this task. Released 1.1.0 objects and the Server Health / Assessment Engine remain excluded.
- **Permitted external actions:** read-only use of the canonical public QWSG Forgejo repository/Release API and assets; normal clean-fast-forward Task 064 Git pushes; private deterministic construction and privacy-safe transfer of one frozen 1.2.0-rc.1 acceptance candidate; bounded access to the existing disposable Ubuntu 24.04 amd64 VPS only after local validation; documented update from its preserved official 1.1.0 installation, validation, explicit rollback, and restoration/re-update only as required to prove the acceptance contract. Owner-only VPS authentication or physical interaction may be requested as one bounded interaction and resumes under the same authority. No protected Forgejo credential is permitted for product downloads.
- **Owner-reserved decisions:** material product or architecture expansion; destructive VPS replacement/reformat; credential disclosure/replacement; unsupported external infrastructure mutation; security/trust-policy change beyond SHA-256 plus authoritative HTTPS release identity; signing/PKI/channel policy; tag/Forgejo Release/publication/deployment/announcement; alteration of released 1.1.0 objects; and Task 065.
- **Task-specific STOP conditions:** STOP for credential/secret exposure risk; security/privacy regression; unsafe privileged destination control; unavailable required rollback; corruption or unexplained mutation of preserved 1.1.0 baseline; required modification of released 1.1.0 or frozen 1.2.0-rc.1 candidate bytes; material architecture/scope expansion; an external update/rollback action that cannot be bounded and recovered safely; or another genuine Framework 2.0 boundary. Ordinary implementation, build, test, simulated update, migration, rollback or acceptance failures enter diagnose -> classify -> correct -> retest within this task.


## Required Reading

- `ai/core/00_PROJECT_PHILOSOPHY.md`
- `ai/core/01_CONSTITUTION.md`
- `ai/core/03_AGENTS.md`
- `ai/core/08_JOB_TEMPLATE.md`
- `ai/core/11_ENGINEERING_LIFECYCLE.md`
- `ai/core/14_PROMPT_WORKFLOW.md`
- `ai/core/16_GIT_POLICY.md`
- `ai/core/17_EXECUTION_MODEL.md`
- `ai/core/18_BOUNDED_DIAGNOSTIC_RUNNER.md`
- `ai/config/engineering-project.conf`

## Starting State Verification

- Verify canonical repository root, ordinary-user execution, branch `main`, Framework 2.0.0 VALID, canonical lifecycle idle after Task 063, zero active prompts, empty index, `HEAD == origin/main ==` direct Forgejo main at `bdcbe1827873f5422f51fe8d488679e1b06866d2`, and only the unrelated excluded Owner blueprint present and untouched.
- Verify `VERSION=1.1.0`; `v1.1.0` tag object `b14347636f6c9873a5acf759c950d900a39bf1a7` peels to release source `305f4088e94b14d6cbb3114eb8cce4e32d847c16`; public final Release and its two frozen assets retain the Task 063 identities; Task 063 is complete and immutable.
- Verify the current CLI has no native update command, characterize installer `--replace`/`--backup-dir`, uninstall preservation, release archive/manifest/provenance, configuration/credential/state roots, Guardian user-service management, readiness and version output through source and focused baseline tests.
- Verify existing public Forgejo distribution endpoints anonymously and record protocol/API assumptions without downloading or mutating the VPS during starting-state inspection.
- Do not access the test VPS before the active task has a verified snapshot and local architecture/validation identifies the minimum external acceptance runner.


## Snapshot Requirements

Before modifying tracked task targets, create a unique private mode-0700 snapshot under `/tmp` containing an exact tracked-HEAD archive, hashes/modes and literal before-images of update-adjacent CLI/installer/release/configuration/state/service/version/documentation/lifecycle targets, exact Git/ref/released-1.1.0 identities, excluded-path record without reading its content, and collision-aware bounded rollback instructions. Exclude credentials, caches, downloaded artifacts and Owner content. Before external mutation, create a second private acceptance-boundary snapshot/ledger containing the exact validated 1.2.0-rc.1 source/candidate identities, simulated-update validation, immutable 1.1.0 baseline expectations, VPS preservation/rollback plan, allowed mutations and cleanup. Verify both snapshots' permissions, hashes, readability, exclusions and guarded restoration paths.


## Risk Assessment

- HIGH — privileged replacement or rollback could damage installation paths. Mitigate with a prevalidated immutable plan, fixed allowlisted destinations, symlink/type/ownership checks, private backups, atomic writes where applicable, narrow elevation and failure injection.
- HIGH — remote bytes could control privileged mutation. Mitigate by authoritative HTTPS source pinning, bounded download, sidecar/archive/manifest/provenance/platform verification and complete extraction safety before invoking any privileged installation step.
- HIGH — configuration, credentials or persistent state could be lost or exposed. Treat them as user-owned non-package data, never place secrets in backups/evidence, preserve modes/ownership, test exact non-mutation and migrate only through versioned deterministic code.
- HIGH — failed service/update sequencing could leave Guardian unavailable. Capture enabled/active intent, stop only at the mutation boundary, reload/start conditionally, validate readiness, and automatically restore the prior installation when deterministic.
- MEDIUM — semantic version or prerelease ordering could choose an invalid target/downgrade. Use a strict tested parser/policy and explicit operator override only where safely designed; default update rejects equal/older/unsupported/ambiguous identities.
- MEDIUM — rollback metadata or backup tampering could install unsafe bytes. Store private integrity-bound metadata, validate every backup artifact and source/destination rule, and refuse incomplete or foreign rollback state.
- MEDIUM — environment-specific external behavior could be mistaken for a product defect. Use deterministic local HTTP/system/install harnesses first, classify against product contract, then run one bounded real 1.1.0 upgrade/rollback acceptance without reformatting.
- LOW — new user-visible commands may resist localization. Keep stable machine tokens/JSON, route human text through the existing CLI conventions and document localization boundaries.


## Planned Work

1. Validate the canonical idle/release baseline, read governing and Task 063 distribution evidence, inspect update-adjacent source/contracts, create and verify the pre-change snapshot, and record a concise architecture decision before implementation.
2. Define strict QWSG release identity/order and canonical distribution discovery contracts, including immutable tag/assets, bounded HTTP behavior and explicit failure classification.
3. Implement unprivileged update check/plan/download/stage verification so no remote content is executed or controls privileged destinations before all integrity, provenance, platform and compatibility gates pass.
4. Implement private installed-version/update/rollback metadata and a narrow transaction boundary that captures package artifacts and service intent, preserves user-owned configuration/credentials/state, performs fixed-path replacement, validates the result, and restores deterministically on eligible failure.
5. Implement versioned migration planning/application/rollback interfaces with an explicit 1.1.0-to-1.2.0-rc.1 compatible/no-op path and deterministic failure semantics.
6. Implement explicit operator rollback/status behavior and bounded backup retention/refusal rules; integrate Guardian stop/reload/start only when needed.
7. Add focused unit tests, malicious package/archive fixtures, simulated authoritative HTTP distribution, unprivileged staging, destdir/systemd harnesses and injected failures proving no pre-verification mutation, preservation, automatic restore, explicit rollback, service-intent recovery and migration safety.
8. Advance development identity to `1.2.0-rc.1`, update changelog/release notes/operator/security/upgrade documentation and package contents, then run strong local Framework/Go/race/vet/format/shell/systemd/release-plumbing/security/rollback validation. Diagnose and correct ordinary failures autonomously.
9. Construct two isolated deterministic private 1.2.0-rc.1 outputs, require byte identity, select/freeze one acceptance candidate and fully verify it. Do not tag, publish or modify it after freeze.
10. Use one bounded privacy-safe external acceptance runner against the existing installed QWSG 1.1.0 VPS: record baseline identity/data/service intent; transfer and verify the frozen candidate; exercise authentic native check/update; verify configuration, protected credentials, compatible state, service intent and readiness; exercise explicit rollback to exact 1.1.0 and validate preservation/operation; restore the intended final development-test state only through documented updater behavior. Do not reinstall/reformat or manually repair runtime state to manufacture PASS.
11. Reconcile evidence, document limitations, perform explicit-path validated Git integration and clean-fast-forward push, directly verify canonical refs, archive Task 064 and return lifecycle to idle.


## Rollback Plan

- Local development rollback uses only the verified Task 064 pre-change snapshot after exact HEAD/target/collision checks. Restore literal task targets and modes from the snapshot; never use broad reset/checkout/clean and never touch the excluded Owner blueprint.
- Test harness mutations remain inside unique private temporary roots and clean up exact created paths. Simulated servers/processes use bounded timeouts and recorded PIDs.
- Update transaction rollback restores only the integrity-verified prior QWSG package artifacts and recorded service intent; user configuration, credentials and persistent state are preserved rather than replaced. Failed migrations invoke their deterministic reverse/restore contract before package recovery.
- External acceptance begins with a privacy-safe 1.1.0 baseline and verified rollback availability. If an update fails after mutation, use the product's automatic rollback; if explicit rollback is under test, use only the documented Task 064 command. Do not manually overwrite runtime files, reinstall the OS or erase preserved data. STOP if trustworthy recovery becomes unavailable.
- Frozen 1.2.0-rc.1 candidate bytes are never edited. A candidate defect returns to Task 064 source development only after preserving failure evidence; rebuild identity is new and the superseded candidate remains identified, not overwritten.


## Deliverables

- A documented native update architecture and operator workflow integrated with the existing CLI and installation model.
- Strict installed/available version comparison and canonical public Forgejo discovery/download implementation.
- Fully verified private staging and package trust pipeline before privileged mutation.
- Transactional replacement with preservation of configuration, protected credentials, compatible state and service intent; deterministic automatic recovery on eligible failure.
- Versioned deterministic migration framework with 1.1.0 compatibility coverage.
- Integrity-bound local update/rollback metadata and explicit safe rollback capability.
- Deterministic unit/integration/simulated-distribution/failure-injection tests and strong validation evidence.
- Private deterministic QWSG 1.2.0-rc.1 acceptance candidate identity without tag, Forgejo Release or publication.
- Bounded external evidence from the preserved installed QWSG 1.1.0 VPS proving native update, preservation, service/readiness behavior and explicit rollback.
- Updated English operator, security, release, upgrade/rollback, changelog and Task 064 engineering history; validated integration and canonical idle closure.


## Verification

- Framework 2.0, Builder, lifecycle, one-active-task, diverted-task audit, repository identity/cleanliness and excluded Owner path checks PASS.
- Strict version parser/order tests cover newer/equal/older, final/prerelease, unsupported and invalid identities without accidental downgrade.
- Discovery tests use controlled HTTP fixtures and prove canonical origin/tag/asset selection, time/size/redirect bounds, anonymous access and rejection of malformed/ambiguous metadata.
- Package verification rejects wrong sidecar/archive/manifest/binary version/source/platform, unsafe paths/types/symlinks/duplicates, missing docs and provenance mismatch before mutation.
- Transaction tests prove fixed privileged destinations, private rollback backups/metadata, no remote destination control, preservation of configuration/credential/state ownership and modes, service-intent capture/recovery, atomic/failure-safe behavior and automatic restoration after injected post-mutation failures.
- Migration tests prove deterministic version selection, no-op compatibility, idempotence where required, failure rollback and refusal of unknown schema paths.
- Explicit rollback tests prove exact prior identity restoration, readiness/service behavior, tamper/incomplete-backup refusal and bounded retention semantics.
- Full Go tests, targeted race and repository-wide race as proportional, vet, formatting, shell syntax, systemd static verification, release/build plumbing, privacy/secret scan, documentation consistency, Git diff/staging and rollback-readiness checks PASS.
- Two isolated 1.2.0-rc.1 builds are byte-identical; selected private candidate passes archive, sidecar, manifest, binary provenance, documentation, modes and package safety verification and remains frozen/unpublished.
- External acceptance on the preserved VPS proves starting QWSG 1.1.0 identity, authentic native update to exact 1.2.0-rc.1 candidate, preservation of configuration/credentials/compatible state/service intent, operational validation, explicit rollback to exact 1.1.0, and final documented state. Missing required evidence is not PASS.
- Final `HEAD == origin/main ==` direct Forgejo main, Task 064 archived, zero active prompts and canonical idle lifecycle PASS; released 1.1.0 tag/assets/evidence remain unchanged.


## Documentation Updates

- Update `CHANGELOG.md`, `VERSION`, relevant `README.md`/`INSTALL.md` and packaged operator documentation for QWSG 1.2.0-rc.1 native update/rollback behavior.
- Add or update the canonical update architecture, distribution trust, migration, security/privacy, troubleshooting, operations and upgrade/rollback contracts at paths selected after repository inspection; avoid duplicate competing documentation.
- Add private 1.2.0-rc.1 source/candidate readiness and external acceptance evidence without modifying historical 1.1.0 records.
- Maintain `ai/history/064_2026-08-26_native-update-rollback-system.md` throughout execution and add only a concise milestone to `ai/core/07_ENGINEERING_HISTORY.md` if completion merits it.
- Archive the completed Task 064 prompt transactionally and record exact implementation/evidence/closure commits and final repository state.


## Completion Criteria

Task 064 is complete only when the native updater can safely identify an installed version, discover and verify an authoritative update, plan and apply a bounded transaction, preserve user configuration/credentials/compatible state/service intent, support deterministic versioned migrations, automatically recover from eligible post-mutation failure, and explicitly roll back from integrity-verified local metadata; deterministic local and simulated tests pass; a frozen unpublished 1.2.0-rc.1 candidate is constructed and verified; and bounded external evidence on the preserved QWSG 1.1.0 VPS proves authentic update plus preservation and rollback without OS reinstall or manual repair. Released QWSG 1.1.0 objects/history must remain unchanged, security/privacy and rollback boundaries must pass, limitations must be explicit, validated commits must be clean-fast-forward pushed/directly verified, Task 064 must be archived and lifecycle canonical idle. A genuine boundary yields `blocked` with preserved evidence; ordinary defects remain in the Task 064 diagnose/correct/retest loop. No QWSG 1.2.0-rc.1 publication and no Server Health / Assessment Engine work are permitted.


## Owner Approval Requirements

Approved by Project Owner through the Engineering Task Builder on 2026-08-26 UTC.

The structured task definition and Authority Envelope have been explicitly approved. Framework 2.0 Standard Execution Authority permits iterative, reversible in-scope engineering without another Owner gate. Further scope changes, exceptional external actions, and Owner-reserved decisions require explicit Project Owner approval.
