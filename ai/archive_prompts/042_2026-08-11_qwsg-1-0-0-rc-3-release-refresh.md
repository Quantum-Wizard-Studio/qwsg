# Current Engineering Task 042: QWSG 1.0.0-rc.3 Release Refresh

## Task Metadata

- Task ID: `042`
- Task slug: `qwsg-1-0-0-rc-3-release-refresh`
- Status: `complete`
- Date opened: `2026-08-11` UTC
- Human authority: Project Owner
- Owner or lead-developer communication language: Hungarian

## Title

QWSG 1.0.0-rc.3 Release Refresh


## Objective

Produce the canonical reproducible QWSG `1.0.0-rc.3` linux-amd64 Release Candidate from the accepted Task 041 baseline using the established Task 038/040 process. Include and artifact-verify the clean-account first-run correction without changing product semantics, architecture, licensing or publication state. End `READY FOR CLEAN-HOST ACCEPTANCE`; do not run the external VPS acceptance.


## Scope

- Advance coherent release identity from `1.0.0-rc.2` to `1.0.0-rc.3` in `VERSION`, `cmd/qwsg/main.go` fallback, README, roadmap, CHANGELOG and exact operational release references.
- Update `docs/release/QUICK_START.md` to RC.3 archive/root/checksum names; add `RELEASE_NOTES_1.0.0-rc.3.md` and a sanitized completed `ACCEPTANCE_1.0.0-rc.3.md`.
- Preserve RC.1 and RC.2 artifacts, sidecars, notes and acceptance records byte-for-byte. Never relabel the historical RC.2 artifact as containing Task 041.
- Keep `scripts/build-release.sh` unchanged unless final inspection proves a narrow RC.3 blocker; it already validates `1.0.0-rc.N`, selects version-derived notes and emits location-independent checksums.
- Produce `dist/qwsg-1.0.0-rc.3-linux-amd64.tar.gz` and matching `.sha256` with fixed recorded epoch/commit, internal manifest and the existing canonical payload.
- Perform two independent controlled builds and require byte-identical binary, manifest, archive and sidecar.
- Reuse complete Task 038/040 archive inspection, staged install/replace/rollback/uninstall, unit, compatibility, privacy/security, lifecycle, resource and endurance gates without weakening them.
- Prove from the extracted final artifact all accepted Task 039 behavior: bounded large-Report Alert reference, completed Runtime/Notification planning, read-only Console refresh, useful bounded diagnostics/Attention, lock exclusion and truthful termination freshness.
- Prove from the extracted final artifact all accepted Task 041 behavior: empty HOME with absent `.local` hierarchy, secure recursive bootstrap, `0700` directories, `0600` files, partial first baseline, later full observation, partial `check` publication, bounded bootstrap/publication diagnostics and unchanged unsafe-path protection.
- Run artifact acceptance outside the repository with no prior QWSG state and no runtime dependency on Go, Git, checkout or development tooling. Record the development-host isolation limitation honestly; final physical no-Go/reboot proof remains Owner clean-VPS work.
- Update Task 042 history, acceptance evidence and concise milestone index; archive Task 042 and return to canonical idle without a successor.


## Out of Scope

- No product feature or modification to accepted Task 039/041 Scheduler, collector, Pipeline, Alert, Runtime, Guardian, Console, Current State or storage semantics.
- No new collector, provider, transport, Dashboard, API, fleet, remote management, remediation, AI, persistence platform, installer architecture or Go runtime dependency.
- No licensing change, signing, public tag, stage, commit, push, upload, publication or announcement.
- No real `/usr/local` installation, real QWSG service/state/config/user-manager/linger mutation, or operation against the Owner clean VPS.
- No final `1.0.0` identity, additional OS/architecture/package format, Task 043 or successor creation.
- Stop if release production requires a product correction, policy/license decision, publication authority or modification of accepted Task 041 behavior.


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

- Require `bin/job --check` canonical idle with Task 041 unique latest complete and no Task 042 collision.
- Verify QWSG root/Framework, `main`, exact HEAD, canonical HTTPS origin, `0/0` upstream, empty index, full dirty-tree ownership, modes/ACLs, capacity, active processes/services and preservation of unrelated Owner content.
- Read Task 038/040 release process and acceptance, Task 039/041 prompts/histories, Release Policy, builder, installer/uninstaller, unit and complete release documentation.
- Confirm policy permits `1.0.0-rc.3`; current identity is exactly RC.2; RC.3 targets are absent; RC.2 archive SHA-256 is `0694fb7f382ea1b373aaf0b6f0171a3fc580c491c1af3b8657a3bb8697ed897b` and predates Task 041.
- Determine exact metadata references and prove the generalized release builder needs no semantic change. Record hashes/listings for all RC.1/RC.2 historical evidence.
- Verify Ubuntu 24.04/linux-amd64, Go, Make, GNU tar/gzip, SHA-256 and systemd tools without installing dependencies.
- Run pre-change focused Task 039/041, full/race/vet/format, Framework, Builder, lifecycle, diversion, `bin/job`, unit and Git checks; stop on material contradiction.


## Snapshot Requirements

- Before changes create a unique mode-`0700` external Task 042 snapshot covering every metadata/code/test/doc target, Task 042 prompt/history, builder/installer/unit, VERSION, RC.1/RC.2 artifacts and RC.3 absence.
- Record exact hashes, bytes, modes, owners, ACLs, repository identity/status/index, tools, capacity, active processes/services and unique temporary build/install/lifecycle roots.
- Include deterministic manifests, bounded archive, retention notes and guarded restore instructions refusing changed targets, later Owner edits, active processes, ambiguous artifact ownership or broad recovery.
- Before artifact/lifecycle tests record exact binary/unit hashes and initially absent roots. Retain implementation snapshot; remove only verified temporary roots after process absence.


## Risk Assessment

- **Identity/artifact divergence — critical:** derive every operational name from exact VERSION and assert binary, archive root, notes, Quick Start, changelog, manifest and sidecar agree.
- **Stale RC.2 payload — critical:** preserve historical bytes but reject RC.2 notes/identity inside RC.3 and artifact-test Task 041 behavior.
- **False reproducibility — critical:** fixed commit/epoch, independent output roots and byte comparisons for all release outputs.
- **Product regression — critical:** metadata-only source change except narrowly necessary release tests; run focused Task 039/041 and complete gates.
- **Unsafe install or Owner-data loss — critical:** staged roots only, exact snapshots/cleanup, no broad Git recovery, no overwrite of older RCs.
- **False no-Go/clean-host claim — high:** distinguish development artifact evidence from unexecuted external clean-VPS/reboot acceptance.
- **Security/privacy/license overreach — high:** retain storage/installer hardening and proprietary private-RC boundary; record only sanitized evidence.


## Planned Work

1. Validate lifecycle, policy, baseline, repository/tools, historical artifacts and all pre-change gates; create and verify rollback snapshot.
2. Apply only RC.3 identity metadata, Quick Start, changelog, new release notes/acceptance baseline and narrowly necessary identity tests. Audit that Task 039/041 implementation is untouched.
3. Build RC.3 twice with identical fixed inputs into separate clean roots; verify exact version, byte reproducibility, archive safety, manifest, checksum, modes, timestamps and payload.
4. Extract final archive and run staged clean install, collision/replacement/backup/rollback/modified-uninstall refusal/clean uninstall, version and systemd-unit checks without real-host installation.
5. Re-run applicable Task 038/040 lifecycle/endurance/security gates and artifact-level Task 039 large-report Runtime/Console/termination acceptance.
6. From the final staged binary run Task 041 empty-HOME first observe, second observe, Console and partial-check acceptance; verify ownership/modes and symlink/permissive-root refusal. Prove no Go/Git/check-out runtime dependency and disclose host isolation limits.
7. Run final focused/full/race/vet/format, Framework/Builder/lifecycle/diversion, privacy, snapshot, artifact and Git gates.
8. Complete RC.3 acceptance with exact hashes and transfer files, clean temporary roots, preserve prior RCs/snapshot, archive Task 042 and return idle.


## Rollback Plan

Stop if any Task 042 build, Guardian, service or installer process is active or if snapshot/target/Owner/artifact identities differ. Restore only snapshot-listed metadata and documentation after confirming current bytes are Task 042-owned and no later Owner edit exists. Remove new RC.3 paths only when recorded absent and current hashes match Task 042 evidence. Never alter RC.1/RC.2, private state or unrelated content; never use reset, clean, checkout, broad wildcard deletion or extraction over the repository. Re-run all identity, historical-hash, test, governance and idle checks after rollback.


## Deliverables

- coherent `1.0.0-rc.3` source/release metadata and exact release notes/acceptance record;
- reproducible linux-amd64 archive plus checksum sidecar and verified internal manifest;
- preserved RC.1/RC.2 evidence;
- staged install/uninstall, unit, security, compatibility and full release-gate evidence;
- artifact-level Task 039 and Task 041 inclusion evidence, including empty-HOME bootstrap;
- exact Owner transfer instruction for the RC.3 archive and sidecar;
- retained snapshot, completed history/archive and canonical idle closure.


## Verification

- Verify initial/final repository identity/status/index, exact changed paths, ownership/modes/ACLs, snapshot hashes and unrelated Owner-content preservation.
- Run focused Task 039/041 packages, `make build`, `go test ./...`, bounded-cache `go test -race ./...`, `go vet ./...`, `make fmt-check`, Framework, Builder, lifecycle, diversion/audit, `bin/job` and `git diff --check`.
- Assert all RC.3 identity surfaces agree and operational RC.2 references are absent while historical RC.1/RC.2 records remain unchanged.
- Require two byte-identical builds for binary, manifest, archive and sidecar; independently verify SHA-256, safe single archive root, deterministic metadata and exact payload.
- Validate extracted version, manifest and static/no-development-tool runtime contract; staged installer/uninstaller and shipped unit must pass without activation or linger changes.
- Re-run bounded Guardian recurrence/endurance and Task 039 366+ source, Runtime completion, Console `r`, lock, diagnostic, Attention, graceful/unexpected termination gates.
- With final artifact and empty HOME, require first partial baseline, second normal full pipeline, Console load and partial `check`; verify QWSG directories `0700`, files `0600`, existing ancestors unchanged, symlink/unsafe/mode rejection and privacy-safe diagnostic tokens.
- Confirm installer/unit/licensing/product semantics are unchanged, no real-host/VPS install or Git/publication action occurred, all temporary roots/processes are absent and acceptance decision is justified.


## Documentation Updates

- Update `VERSION`, `cmd/qwsg/main.go` fallback, `README.md`, `ai/core/13_ROADMAP.md`, `CHANGELOG.md` and `docs/release/QUICK_START.md` only for RC.3 identity.
- Add `docs/release/RELEASE_NOTES_1.0.0-rc.3.md` describing accepted Task 041 and unchanged boundaries.
- Add and complete `docs/release/ACCEPTANCE_1.0.0-rc.3.md` with actual reproducibility, artifact, staged-install, Task 039/041 and transfer evidence, explicitly excluding clean-VPS/reboot acceptance.
- Update only necessary release/version tests, Task 042 history/archive and concise engineering milestone. Do not rewrite prior release evidence.


## Completion Criteria

Task 042 is complete only when lifecycle begins idle at Task 041 and ends idle at Task 042; metadata-only scope and all exclusions hold; exact RC.3 identity is coherent; RC.1/RC.2 bytes remain unchanged; two builds and all manifest/checksum/archive/install/unit/full/race/governance gates pass; extracted RC.3 demonstrates accepted Task 039 and Task 041 behavior including secure empty-HOME bootstrap and truthful partial evidence; temporary roots/processes are removed; acceptance records exact artifact path/hash and says `READY FOR CLEAN-HOST ACCEPTANCE`; transfer files are exactly the RC.3 archive and sidecar; no external VPS test, licensing change, Git publication action or successor task occurred. No product blocker may remain undisclosed.


## Owner Approval Requirements

Approved by Project Owner through the Engineering Task Builder on 2026-08-11 UTC.

The structured task definition has been explicitly approved for implementation. Further scope changes require explicit Project Owner approval.
