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
- Documentation correction commit `6f7379d158852c1672070da2558dab82588e9697` was clean-fast-forward pushed. Two new builds were byte-identical. The selected frozen mode-0400 archive SHA-256 is `4ffab382b4f3e4bbbe3f67e5020fa099562aa43453ccd0fe5f328694e0ddf264`; sidecar SHA-256 is `5f27e3a1af066810b65eb5c16d89ad19973e5d51170a0703f6f80e507c1ff0de`. The acceptance boundary is recorded in the private execution snapshot.
- External disposable-host fresh-install acceptance remains pending because this engineering environment is not disposable and its user systemd manager is unavailable. No public release action occurred and the preserved official-1.1.0 VPS was not accessed.
- Owner-provided disposable Ubuntu 24.04.4 amd64 host preflight proved ordinary UID, working user systemd manager, absent QWSG package/configuration/state/unit, and interactive sudo. Candidate transfer, outer checksum, internal manifest and exact embedded provenance passed. English, Hungarian and German pre-mutation plan/cancel journeys showed 0 percent during active preflight, 12 percent only after completion, localized phase/plan content, and no mutation.
- The Owner supplied the inherently interactive sudo credential directly inside the VPS tmux session. Resulting exact RC.2 package, private configuration/state, enabled/active Guardian, fresh canonical evidence, notify policy, absent credential directory, and readiness/update-status checks passed. Modes were binary 0755, configuration directory/state 0700 and configuration 0600.
- A complete Hungarian rerun showed deterministic progress `0,12,22,42,57,67,75,87,97,100`, active Guardian and operational readiness. It also exposed two ordinary UX defects: the notification prompt claimed immediate configuration while the action offered provider-value guidance, and completion said only activation requested while omitting truthful overall PARTIAL readiness caused solely by optional notification/lingering. Classified `PRODUCT UX/LOCALIZATION DEFECT`; corrected catalogs and readiness-derived completion semantics locally, focused tests and vet PASS. Only the invalidated package/completion evidence will be repeated with a new candidate provenance.

## Rollback

Use only `/tmp/qwsg-task065-execution.03GLKY/ROLLBACK.md` after verifying exact identities and collision safety. Restore literal Task 065 paths only; no broad reset/checkout/clean and no access to excluded Owner content. Product package restoration remains governed by Task 064 fixed-destination integrity-verified rollback; user configuration, credentials, and persistent state remain outside package backup payloads.

## Completion state

`active — local guided-installer implementation passes focused/strong validation; external clean-host acceptance and final integration/closure remain pending`
