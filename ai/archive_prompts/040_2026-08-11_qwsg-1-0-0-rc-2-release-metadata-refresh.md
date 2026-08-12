# Current Engineering Task 040: QWSG 1.0.0-rc.2 Release Metadata Refresh

## Task Metadata

- Task ID: `040`
- Task slug: `qwsg-1-0-0-rc-2-release-metadata-refresh`
- Status: `complete`
- Date opened: `2026-08-11` UTC
- Human authority: Project Owner
- Owner or lead-developer communication language: Hungarian

## Title

QWSG 1.0.0-rc.2 Release Metadata Refresh


## Objective

Create the canonical, reproducible QWSG `1.0.0-rc.2` linux-amd64 Release Candidate from the accepted Task 039 baseline by applying only the smallest coherent release/version metadata updates and reusing the complete Task 038 release process.

Task 040 is procedural release work, not product development. It shall preserve the accepted Task 039 implementation unchanged, prove that the packaged executable contains its release-blocking corrections, and end with a self-contained archive ready for transfer to an Owner-controlled clean Ubuntu 24.04 x86-64 VPS. It shall not install on the development host, execute clean-host acceptance, change licensing, tag, push, publish, or create a successor task.


## Scope

- Advance the single canonical release identity from `1.0.0-rc.1` to `1.0.0-rc.2` consistently across `VERSION`, embedded command fallback metadata, changelog, release documentation and artifact names.
- Preserve `1.0.0-rc.1` release notes, acceptance evidence and artifacts as historical records. Do not relabel or overwrite them as RC.2.
- Change `scripts/build-release.sh` only as required to package the release-notes document matching the exact validated `VERSION`. Prefer one narrowly validated version-derived filename over another RC.2-specific hard-coded branch. Do not generalize the release format, platform matrix or packaging architecture.
- Update `docs/release/QUICK_START.md` to the exact RC.2 archive, checksum and extracted-directory names.
- Add `docs/release/RELEASE_NOTES_1.0.0-rc.2.md` describing the accepted Task 039 corrections and unchanged private-RC/product boundaries.
- Add `docs/release/ACCEPTANCE_1.0.0-rc.2.md` as the sanitized Task 040 build and local artifact-acceptance record. It must distinguish completed development-host artifact validation from the not-yet-executed Owner clean-VPS/reboot acceptance.
- Update `CHANGELOG.md` with the coherent RC.2 identity and only the already accepted Task 039 corrections plus release refresh facts.
- Update the default version string in `cmd/qwsg/main.go` only to keep non-release builds and source metadata coherent with `VERSION`; release builds continue to embed the canonical identity through existing linker flags.
- Update directly affected version/release tests or documentation references only where required to reject identity drift and prove the RC.2 package contract. Do not change product semantics.
- Reuse the Task 038 deterministic linux-amd64 archive builder and controlled `SOURCE_DATE_EPOCH`/`BUILD_COMMIT` inputs. Produce `dist/qwsg-1.0.0-rc.2-linux-amd64.tar.gz` and its matching `.sha256` sidecar without overwriting RC.1.
- Require the internal `MANIFEST.sha256` to cover the prebuilt QWSG binary, systemd user unit, installer, uninstaller, configuration example, license, changelog, Quick Start, operations, troubleshooting, upgrade/rollback/uninstall, support, security/privacy, known limitations and exact RC.2 release notes.
- Prove that extraction and supported installation need no Go, Git, repository checkout or development tooling. Build tools are permitted only on the development host during artifact production.
- Run the complete Task 038 release gates against the final RC.2 bytes: controlled duplicate builds, archive and manifest inspection, staged archive-only installation, collision/replacement/rollback/uninstall checks, version verification, unit validation, compatibility, security/privacy, bounded lifecycle/endurance and repository governance checks.
- Add artifact-level evidence for the accepted Task 039 behavior using the extracted final RC.2 binary and isolated private roots: bounded large-Report Alert/Runtime success at real 366-or-more scale, read-only Console refresh while Guardian owns the lock, privacy-safe diagnostic behavior, explicit-observe `guardian_active` exclusion, graceful stop and freshness-bounded unexpected termination truthfulness.
- Keep the clean Ubuntu 24.04 VPS untouched. End by reporting the exact archive and checksum sidecar that the Owner must transfer for the separately authorized clean-host acceptance.
- Update Task 040 history throughout execution, archive the completed prompt, return to canonical idle with Task 040 as the latest completed baseline, and retain rollback evidence.


## Out of Scope

- No modification of accepted Scheduler, Pipeline, Health, Rule, Policy, Report, Alert, Notification, Runtime, Runtime Service, Current State, Presentation, Console or Guardian semantics.
- No reimplementation or redesign of Task 039 behavior. A product-behavior failure discovered during artifact validation is a stop condition, not authority for an unplanned fix.
- No new feature, collector, provider, transport, API, Dashboard, network listener, fleet, remote management, remediation, AI, licensing enforcement, updater, package repository, installer architecture or persistence platform.
- No DEB, RPM, container, AppImage, signing system, trust infrastructure, additional OS/architecture support or public distribution workflow.
- No change to `LICENSE`, public licensing position, commercial terms, final `1.0.0` identity or support claims.
- No installation into `/usr/local` or other real development-host destinations; no activation, enablement or modification of the real QWSG user service, state, configuration, user manager or linger setting.
- No clean-VPS installation, reboot, acceptance or remote-host operation.
- No mutation or deletion of the historical RC.1 archive, sidecar, release notes or acceptance record.
- No Git stage, commit, tag, branch, fetch, push, external upload, publication or release announcement.
- No Task 041 or other successor. Stop and report if a genuine product blocker, version-policy conflict, licensing decision or publication authority becomes necessary.


## Required Reading

- `ai/core/00_PROJECT_PHILOSOPHY.md`
- `ai/core/01_CONSTITUTION.md`
- `ai/core/03_AGENTS.md`
- `ai/core/08_JOB_TEMPLATE.md`
- `ai/core/11_ENGINEERING_LIFECYCLE.md`
- `ai/core/14_PROMPT_WORKFLOW.md`
- `ai/core/16_GIT_POLICY.md`
- `ai/config/engineering-project.conf`

## Starting State Verification

- Begin only when `bin/job --check` reports canonical idle with Task 039 as the unique latest completed archived prompt/history pair and no Task 040 lifecycle collision.
- Verify repository root, Framework markers, `main`, exact HEAD, canonical HTTPS origin, `0/0` upstream relationship, empty index, complete Git status, ownership/modes/ACLs, filesystem capacity and preservation of every unrelated Owner-owned modified/untracked path.
- Read the complete Task 038 prompt/history/acceptance evidence, Release Policy, release builder, installer/uninstaller, unit, release documents and Task 039 prompt/history/implementation evidence before changing a target.
- Confirm Task 039 is accepted, no product correction remains, and the existing RC.1 archive predates Task 039 and cannot be reused or renamed.
- Confirm Release Policy permits `1.0.0-rc.N`; `VERSION` is exactly `1.0.0-rc.1`; `scripts/build-release.sh`, Quick Start, release notes, acceptance documentation and the command fallback contain RC.1-specific references; and `1.0.0-rc.2` has no policy or filename collision.
- Record hashes and archive listings for existing RC.1 artifacts. Verify their bytes are preserved throughout Task 040.
- Verify Ubuntu 24.04, linux-amd64/x86-64, Go, GNU tar, gzip, SHA-256 tooling, Make and systemd analysis prerequisites. Do not install a missing dependency without new Owner authority.
- Run pre-change build, full tests, race tests, vet, formatting, Framework, Builder, lifecycle, diverted-task, `bin/job`, Git whitespace and shipped-unit validations. Stop on a material baseline contradiction.
- Confirm the existing temporary proprietary license remains sufficient only for the private technical RC and requires no decision for local RC.2 creation.
- Confirm no QWSG acceptance Guardian, temporary service or prior release-build process is active before snapshot or build work.


## Snapshot Requirements

- Before modifying any target, create a unique mode-`0700` Task 040 implementation snapshot outside the repository with an exact manifest for `VERSION`, `cmd/qwsg/main.go`, `CHANGELOG.md`, `scripts/build-release.sh`, `docs/release/QUICK_START.md`, RC.1 release/acceptance records, intended RC.2 release/acceptance paths, directly affected tests, Task 040 prompt/history and existing `dist` RC.1/RC.2 artifact presence.
- Record absence for new RC.2 paths, exact hashes for all existing targets, repository root/branch/HEAD/remotes/upstream/status/index, tool and platform versions, modes/owners/ACLs, filesystem capacity and active-process/service preconditions.
- Include deterministic SHA-256 manifests, a readable bounded snapshot archive, retention notes and restore instructions that refuse changed targets, later Owner edits, active release/Guardian processes, ambiguous artifact identity or unsafe path state.
- Before any temporary lifecycle, staged-install or endurance operation, record exact unique unit/process names, executable and unit hashes, isolated state/config/store/install roots, pre-existing absence and exact cleanup boundaries.
- Retain the verified implementation snapshot after completion. Acceptance roots and second-build directories are temporary and must be removed only after exact process absence and evidence capture.


## Risk Assessment

- **Artifact/version divergence — critical:** derive archive name and embedded binary identity from exact `VERSION`; package only the matching release-notes filename; assert VERSION, binary, archive root, Quick Start, changelog, manifest and sidecar agree.
- **Packaging stale RC.1 content — critical:** preserve historical RC.1 files but explicitly reject RC.1 release notes or filenames inside RC.2 and prove the final archive includes Task 039 code behavior.
- **False reproducibility claim — critical:** use fixed commit and epoch inputs for two clean output directories and compare binary, archive, internal manifest and external sidecar byte-for-byte.
- **Owner content or historical artifact loss — critical:** snapshot exact targets, never clean/reset broadly, never overwrite RC.1, and remove only exact Task 040 temporary roots after process checks.
- **Unintended product change — high:** constrain code changes to release identity fallback and release validation; audit diffs/imports and run Task 039 focused regressions plus full suites.
- **Installer damage — high:** install only into unique staged roots from the unpacked archive; test collision/replacement/rollback/uninstall there; never invoke real `/usr/local` installation.
- **False clean-host readiness — high:** distinguish local artifact readiness from unexecuted VPS installation/reboot evidence and transfer only the archive plus trusted checksum sidecar.
- **License/publication overreach — high:** retain the private proprietary notice, do not sign/tag/push/publish, and make no public-release claim.
- **Host/private evidence exposure — high:** record only bounded counts, hashes, versions and privacy-safe tokens; exclude raw Inventory, hostnames, IPs, usernames, private paths, journals, configuration bodies and secrets from repository artifacts.
- **Long-running validation residue — medium:** use bounded intervals/timeouts and exact PID/unit/root identities; stop immediately on cleanup ambiguity.


## Planned Work

1. Validate canonical idle, Task 039 baseline acceptance, release policy, exact repository/tool state, RC.1 historical artifacts and all pre-change test gates.
2. Create and verify the bounded Task 040 implementation snapshot and exact rollback procedure.
3. Apply only coherent RC.2 metadata changes: `VERSION`, command fallback, changelog, Quick Start, new RC.2 release notes/acceptance baseline, and narrowly version-derived release-note selection in the existing release builder. Add only necessary release-identity tests.
4. Audit the diff to prove no Task 039 product behavior or unrelated Owner content changed. Run focused Task 039 regression packages and release-script checks.
5. Build the final repository binary and the first controlled RC.2 release archive using recorded fixed epoch and commit inputs. Verify exact identity, contents, safe paths, modes, timestamps, owners, manifest and checksum.
6. Build RC.2 again into a separate new output directory with identical controlled inputs; compare binary, internal manifest, archive and checksum bytes. Reject any difference.
7. Extract only the verified final archive into an isolated root. Verify the archive-only journey without Go/Git/repository dependency: version, manifest, staged clean install, collision refusal, explicit replacement/backup, rollback, modified-artifact refusal, bounded uninstall and preserved state/configuration behavior.
8. Validate the shipped systemd unit and run the Task 038 bounded isolated lifecycle/endurance gates required for release confidence without touching the real installed service or linger state.
9. With the extracted final RC.2 binary, rerun artifact-level Task 039 acceptance: at least two recurring full cycles including a real 366-or-more Policy Report evaluation, completed Runtime/Notification planning, separate Console, interactive read-only `r`, explicit-observe exclusion, useful bounded diagnostics/Attention, graceful stop and unexpected-termination freshness demotion.
10. Run complete build/full/race/vet/format, Framework, Builder, lifecycle, diverted-task, Git, privacy, source/import, snapshot and artifact validations against final bytes.
11. Complete the sanitized RC.2 acceptance record with actual hashes/results and the explicit unexecuted clean-VPS gate. Verify exact transfer files, remove only Task 040 temporary acceptance roots, preserve RC.1 and the Task 040 snapshot, archive Task 040 and return to canonical idle.


## Rollback Plan

Stop rollback if any Task 040 build, Guardian, temporary service or staged installer process is active, or if target bytes, snapshot hashes, repository identity, Owner-owned content or artifact ownership differ from the recorded manifest.

Restore only exact snapshot-listed Task 040 targets after verifying current bytes are Task 040-owned and no later Owner edit exists. Restore pre-existing metadata files atomically with their recorded modes. Remove a newly created RC.2 documentation or artifact path only when the snapshot recorded absence and its current hash matches Task 040 evidence. Never remove, rename or overwrite RC.1 artifacts or historical records.

Do not use Git reset, clean, checkout, restore, broad recursive deletion, wildcard artifact removal or extraction over the live repository. Preserve failure artifacts and sanitized evidence until the Owner chooses rollback or diagnosis. After rollback, rerun focused/full tests, build, release-script validation, Framework/lifecycle/Builder/diversion, `bin/job`, Git diff/index/status, snapshot hashes and process/residue checks; require the exact Task 039 canonical idle baseline and RC.1 bytes to be restored.


## Deliverables

- coherent repository release identity `1.0.0-rc.2` with no RC.1/RC.2 metadata mismatch;
- minimally generalized version-matched release-note selection in the existing Task 038 builder;
- preserved immutable RC.1 artifacts and documentation;
- `docs/release/RELEASE_NOTES_1.0.0-rc.2.md` and completed `docs/release/ACCEPTANCE_1.0.0-rc.2.md`;
- updated RC.2 Quick Start and changelog identity;
- deterministic `dist/qwsg-1.0.0-rc.2-linux-amd64.tar.gz` and `dist/qwsg-1.0.0-rc.2-linux-amd64.tar.gz.sha256`;
- internal `MANIFEST.sha256` containing every required release payload with matching hashes;
- artifact-only staged install/replace/rollback/uninstall and unit validation evidence;
- artifact-level evidence that the final RC.2 binary contains all accepted Task 039 corrections;
- full Task 038 gate, reproducibility, privacy/security, compatibility, lifecycle/endurance and repository validation evidence;
- exact Owner transfer instruction naming only the RC.2 archive and trusted checksum sidecar;
- verified Task 040 snapshot, completed history/archive and canonical idle closure without a successor task.


## Verification

- Verify initial/final root, branch, HEAD, origin, upstream relationship, complete status, empty index, exact changed paths, ownership/modes/ACLs and byte preservation of every unrelated Owner path and RC.1 artifact.
- Run `make build`, `go test ./...`, `go test -race ./...` with bounded writable caches, `go vet ./...`, `make fmt-check`, focused Task 039 tests for Alert, Runtime, Presentation, Current State, Console, Guardian, application and CLI, plus any narrowly added release tests.
- Run Framework validation/tests, Builder tests, lifecycle/next-task tests, diverted-task tests/audit, `bin/job --check`, `git diff --check`, staged-path checks, snapshot SHA-256 verification and canonical idle checks at applicable phases.
- Assert `VERSION`, `qwsg version`, archive basename/root, RC.2 release-note heading/name, Quick Start commands, changelog and manifest all identify exactly `1.0.0-rc.2`; reject unexpected `1.0.0-rc.1` operational instructions or packaged release notes while retaining historical files outside the RC.2 payload.
- Use one recorded `SOURCE_DATE_EPOCH` and one recorded hexadecimal `BUILD_COMMIT` for both controlled builds. Require byte-identical extracted binary, `MANIFEST.sha256`, archive and `.sha256` sidecar, and independently verify the final external checksum.
- Inspect the archive for one safe relative root, deterministic ordering/timestamps/numeric ownership, expected modes, no unsafe link/device/path/traversal entry, no secret/private-host data and exactly the canonical Task 038 payload set with RC.2 release notes.
- From the unpacked archive, run `sha256sum -c MANIFEST.sha256`, verify `bin/qwsg version`, and perform all installer/uninstaller tests under a unique `--destdir`. Verify no repository, Go or Git is referenced or required at install/runtime; installer never enables/starts service or changes linger.
- Verify collision refusal, explicit replacement with a new private backup, binary/unit/document hashes, fixed `/usr/local` unit paths under staging, rollback to exact prior staged bytes, modified-artifact uninstall refusal, successful matching uninstall and preserved non-owned state/configuration.
- Run `systemd-analyze verify --user` on the shipped unit and any exact temporary acceptance unit. Preserve Task 038 sandbox/security/resource limits and prove no listener, root Guardian, remote execution, arbitrary shell, remediation or new privilege.
- Re-exercise the complete applicable Task 038 release gates, including isolated service lifecycle, compatibility/fail-closed state behavior and at least 50 bounded recurrence boundaries with retention, memory, task/thread, FD, Scheduler/checkpoint/current-state and restart-loop checks. Record sanitized actual measurements.
- Prove the extracted RC.2 binary contains Task 039 behavior. Synthetic regressions must cover 64, 65, 366 and larger reports; real isolated host evidence must include a 366-or-more Policy evaluation cycle, successful Alert, completed Runtime and Notification planning reachability, and no source-count `alert_evaluation_failed`/`runtime_not_completed`.
- During extracted-binary Guardian operation, require separate Console `running` with current evidence after a completed cycle; real interactive `r` must reload without `Refresh failed`, lock conflict, observation, new snapshot/cycle or state publication; explicit `qwsg observe` must still fail with `guardian_active`.
- Require bounded understandable Attention and privacy-safe Runtime causes without raw errors, paths, hostname/IP, secret/config/source payload or unbounded duplicate rows.
- Require exact graceful termination to clear checkpoint active state and prevent a later running claim. Require exact unexpected termination to leave no running claim at or after the exclusive freshness boundary even if checkpoint `active=true` remains.
- Verify RC.2 acceptance does not install to real development-host locations, alter the real service/state/config/user manager/linger, contact the clean VPS, modify licensing, or perform Git stage/commit/tag/push/publication.
- Verify all temporary RC.2 second-build, extraction, staged-install, service and Guardian roots and processes are absent after evidence capture; retain only canonical `dist` outputs and the implementation snapshot.
- Complete the RC.2 acceptance document with exact version, artifact filename/path, archive SHA-256, manifest evidence, reproducibility, validations, Task 039 behavior, clean-VPS transfer files and explicit `READY FOR CLEAN-HOST ACCEPTANCE` or a bounded failure decision.


## Documentation Updates

- Update `VERSION`, the default version metadata in `cmd/qwsg/main.go`, and `CHANGELOG.md` for the private `1.0.0-rc.2` identity.
- Update `scripts/build-release.sh` only to select the exact release-notes file derived from validated `VERSION`; retain the Task 038 artifact structure and reproducibility controls.
- Update `docs/release/QUICK_START.md` to exact RC.2 checksum, archive and directory names.
- Add `docs/release/RELEASE_NOTES_1.0.0-rc.2.md` describing Task 039 corrections, unchanged frozen scope, private licensing and unexecuted clean-host/publication gates.
- Add and complete `docs/release/ACCEPTANCE_1.0.0-rc.2.md` with sanitized actual build, reproducibility, artifact, staged-install, lifecycle/resource, Task 039 behavior and transfer evidence. Do not rewrite the RC.1 acceptance record.
- Update only directly affected README/install/support references if exact source inspection proves they contain an operational RC.1 filename or contradict RC.2. Do not perform general documentation cleanup.
- Update any narrowly required release/version tests, Task 040 history, archived prompt and concise engineering milestone index. Do not alter accepted Task 039 evidence except to reference it as the baseline.


## Completion Criteria

Task 040 is complete only when all of the following are true:

- canonical lifecycle began idle with Task 039 latest complete and ends idle with Task 040 latest complete, no successor and no lifecycle collision;
- no new product behavior or architecture was introduced, accepted Task 039 source semantics remain unchanged, and exact code diff is limited to release identity fallback plus necessary release validation;
- `VERSION`, embedded binary, archive/root names, RC.2 release notes, Quick Start, changelog, manifest and checksum consistently identify `1.0.0-rc.2`;
- historical RC.1 archive, sidecar, release notes and acceptance evidence remain byte-identical and are not packaged as the RC.2 release notes;
- the existing Task 038 builder produces `dist/qwsg-1.0.0-rc.2-linux-amd64.tar.gz` and its matching `.sha256` without requiring a new packaging architecture;
- the RC.2 archive contains the prebuilt linux-amd64 binary, systemd user unit, installer, uninstaller, validated configuration example, license, changelog, internal manifest and the complete canonical ordinary-user documentation set including exact RC.2 release notes and Quick Start;
- two controlled clean builds are byte-identical for binary, internal manifest, archive and sidecar; all internal and external SHA-256 checks pass;
- archive inspection and staged archive-only clean install/replacement/rollback/uninstall pass without Go, Git, repository checkout, unsafe path, silent service activation, linger change or real-host installation;
- complete build/full/race/vet/format, focused Task 039, Framework, Builder, lifecycle, diversion, Git, unit, privacy/security, compatibility and Task 038 release-gate validation passes;
- the extracted final RC.2 binary completes at least two real recurring Guardian cycles and at least one real 366-or-more Policy evaluation without source-count Alert/Runtime partial failure, reaches Notification planning, publishes truthful Current State and remains bounded;
- a separate Console shows current truthful Guardian state; interactive `r` is read-only and successful; explicit observe remains safely excluded; diagnostics and Attention are useful, bounded and privacy-safe;
- graceful stop and unexpected termination are truthful through checkpoint/lifecycle/freshness semantics and create no false running claim;
- the RC.2 acceptance record contains exact artifact path/name/hash, manifest, reproducibility and validation results, Task 039 inclusion evidence, transfer files and a justified readiness decision while clearly marking clean-VPS/reboot acceptance unexecuted;
- all Task 040 temporary processes, units, roots, locks and staged installs are absent; real development-host service/state/config/linger, clean VPS, unrelated Owner content, license and public release state remain unchanged;
- no dependency installation, Git stage/commit/tag/push, signing, upload, publication or public claim occurred;
- final delivery identifies exactly the RC.2 archive and `.sha256` sidecar for Owner transfer and reports `READY FOR CLEAN-HOST ACCEPTANCE` only if every local artifact gate passed.


## Owner Approval Requirements

Approved by Project Owner through the Engineering Task Builder on 2026-08-11 UTC.

The structured task definition has been explicitly approved for implementation. Further scope changes require explicit Project Owner approval.
