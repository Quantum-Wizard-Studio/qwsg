# Task History 065: Approved Task

## Task metadata

- Task ID: `065`
- Task slug: `guided-installation-setup-experience`
- Status: `active — local implementation and validation in progress`
- Date generated: `2026-08-27` UTC
- Human authority: Project Owner
- Preferred owner communication language: English
- Related prompt: `ai/prompts/065_CURRENT_TASK.md`

## Lifecycle state

The Engineering Task Builder generated and transactionally installed this matching prompt/history pair from validated structured owner input. Explicit approval was recorded and Framework 2.0 Standard Execution Authority is active. Local implementation and validation are in progress; external clean-host acceptance has not yet been claimed.

## Starting state

- Canonical repository root, ordinary user, branch `main`, Framework 2.0.0 and idle-after-064 lifecycle validated before Task 065 installation.
- Starting `HEAD == origin/main == 5e261e9dbd4af7699d8840df20588e29a0569685`; index and tracked worktree were clean. `VERSION=1.2.0-rc.1`.
- The only unrelated path was excluded Owner content at `docs/architecture/QWCS_MIGRATION_BLUEPRINT.md`; its content was not read, copied, hashed, modified, staged, or packaged.
- Existing installation was a split archive `install.sh` -> `qwsg setup` -> `qwsg readiness` journey with read-only `install --check`, protected credential commands, user-systemd activation, and Task 064 update/rollback foundations.

## Snapshot

- Private snapshot: `/tmp/qwsg-task065-execution.03GLKY`, mode `0700`.
- Tracked-HEAD archive SHA-256 is recorded and verified in `tracked-head.sha256`; archive readability passed.
- Snapshot contains lifecycle/Builder identities, literal relevant before-images, Git/ref identity, checksums, excluded-path metadata without content access, and bounded collision-aware `ROLLBACK.md`.
- The first snapshot command named a nonexistent root `INSTALL.md`; classified `TEST OR ACCEPTANCE DEFECT`, corrected to canonical `docs/installation/INSTALL.md`, and the completed snapshot verified.

## Work performed

- Added `internal/installer` as the interface-neutral engine for deterministic phases/weights/state-derived progress, installation/update-policy vocabulary, strict Ubuntu 24.04 amd64 platform detection, stable message IDs, canonical English fallback, and complete English/Hungarian/German catalogs.
- Added `qwsg install --guided`, language selection, capability-aware ANSI dashboard, `TERM=dumb`/explicit line fallback, read-only preflight, understandable pre-mutation plan and consent, narrow packaged `install.sh` invocation, configuration initialization, optional-notification guidance, manual/notify policy persistence, Guardian activation, readiness, localized failure guidance, and localized completion summary.
- Preserved `install --check` and concise expert setup paths. Automatic privileged updates remain explicitly unavailable; notify is stored preference only.
- Advanced private development identity to `1.2.0-rc.2`; added explicit no-mutation migration compatibility from 1.1.0 and private 1.2.0-rc.1; updated release plumbing.
- Added guided-installer architecture, candidate notes, changelog, README, and canonical installation flow documentation.
- Focused failures corrected inside the task: Go composite-literal assertion parsing (`TEST OR ACCEPTANCE DEFECT`), incorrect assessment aggregate field assumption (`PRODUCT DEFECT` during implementation), invalid-option compatibility ordering/message (`PRODUCT DEFECT`), and English leakage from legacy setup/readiness output after language selection (`PRODUCT LOCALIZATION DEFECT`).

## Verification

- Builder input, generated prompt/history, approval state, lifecycle installation and `bin/job --check`: PASS.
- Installer model/catalog/platform and CLI/update focused tests: PASS.
- All Go package tests: PASS.
- Focused race tests for installer/update/CLI: PASS.
- `go vet ./...`, gofmt, release-plumbing regression and `git diff --check`: PASS.
- Pseudo-terminal Hungarian line-mode acceptance proved language selection, truthful 0-percent preflight, localized stage/explanation, strict supported-platform identity and graceful existing assessment blocker. The managed engineering session lacks a reachable user systemd manager, so this run truthfully stopped before plan/mutation and is not clean-host acceptance.
- Implementation commit `0788b1ad822a71bd7538129a7b7eb9ec0385e7cf` was clean-fast-forward pushed to canonical `main` after dry-run.
- Two deterministic RC.2 builds from that commit were byte-identical. The archive SHA-256 was `03cc1e2d77768e9ae16b4e9bfd901036968d652cc6d9b2f915ffacb70f812fbe`; sidecar SHA-256 was `933bf3cba245198f0e2f8a3cfa55b4fb250d53d0e125eaa5fc093a0838b1cb53`; manifest and embedded version/provenance passed.
- Package review then found the packaged living `KNOWN_LIMITATIONS.md` still named RC.1 and omitted guided-installer limitations. Classified `PRODUCT DOCUMENTATION DEFECT`; those candidate bytes are preserved as superseded and are not eligible for external acceptance. Current packaged setup/limitations documentation is being corrected before a new provenance identity is built.
- External disposable-host fresh-install acceptance remains pending; no public release action occurred.

## Rollback

Use only `/tmp/qwsg-task065-execution.03GLKY/ROLLBACK.md` after verifying exact identities and collision safety. Restore literal Task 065 paths only; no broad reset/checkout/clean and no access to excluded Owner content. Product package restoration remains governed by Task 064 fixed-destination integrity-verified rollback; user configuration, credentials, and persistent state remain outside package backup payloads.

## Completion state

`active — local guided-installer implementation passes focused/strong validation; external clean-host acceptance and final integration/closure remain pending`
