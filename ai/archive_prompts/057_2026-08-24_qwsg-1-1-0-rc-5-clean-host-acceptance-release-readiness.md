# Current Engineering Task 057: QWSG 1.1.0-rc.5 Clean-Host Acceptance and Release Readiness

## Task Metadata

- Task ID: `057`
- Task slug: `qwsg-1-1-0-rc-5-clean-host-acceptance-release-readiness`
- Status: `complete with disclosed acceptance limitations — formal certification incomplete`
- Date opened: `2026-08-24` UTC
- Human authority: Project Owner
- Owner or lead-developer communication language: English

## Title

QWSG 1.1.0-rc.5 Clean-Host Acceptance and Release Readiness


## Objective

Construct one private reproducible QWSG `1.1.0-rc.5` candidate from the exact clean source commit produced by separately authorized Task 057 acceptance-scaffolding integration, prove its deterministic provenance and package safety, transfer it only through a separately authorized private boundary, and execute a fresh Owner-operated clean-host acceptance from Checkpoint 01 on a genuinely clean disposable Ubuntu 24.04 amd64 VPS. Directly prove the Task 056 state-directory correction under real systemd user-service behavior, retest QWSG-055-F001, QWSG-051-F001 and Task 049 F002/F003, rerun every mandatory notification/session/reboot/uninstall/reinstall checkpoint, and conclude exactly `READY FOR QWSG 1.1.0 RELEASE` or `NOT READY FOR QWSG 1.1.0 RELEASE` without authorizing final release or publication.


## Scope

- Gate A — audit exact repository/source readiness and prepare only RC.5-specific acceptance protocol, an empty privacy-safe evidence ledger, Task 057 chronology, and narrowly required release-plumbing assertions. Integrate that exact path allowlist only after separate Owner authorization and a clean fast-forward push; the resulting full commit becomes the only candidate-source identity.
- Gate B — after separate authority, create two independent private mode-0700 roots, independently export the exact authorized candidate-source commit without `.git`, unset GOFLAGS, prove ordinary exported-source `make build`, derive SOURCE_DATE_EPOCH from the commit, embed the exact full lowercase commit, and independently construct RC.5 through the canonical release mechanism from each exported module root.
- Gate C — prove byte identity for binary, `MANIFEST.sha256`, archive and checksum sidecar; verify exact embedded version/commit/UTC time, static Linux amd64, safe archive root/paths/types/modes/ownership/timestamps, complete manifest, documentation, LICENSE and all security/privacy exclusions. Integrate only privacy-safe construction evidence after a separate reviewed authorization; never rebuild merely to record evidence.
- Gate D — after separate authority, privately transfer exactly the verified RC.5 archive and sidecar through an Owner-approved authenticated boundary with strict host-key verification. If direct authentication is unavailable, stop for an explicit Owner-workstation fallback. Verify destination count, regular non-symlink types, size, SHA-256 identities and sidecar PASS before extraction or execution.
- Gate E — after separate authority and Owner confirmation that the disposable VPS was freshly reinstalled/reset, execute a restartable Owner-operated 26-checkpoint protocol from candidate receipt through final reinstall/readiness. Every checkpoint defines purpose, bounded action, expected evidence, PASS, FAIL/finding, continuation safety and privacy/redaction requirements. State-changing, credential, notification, logout, reboot, restart, uninstall and reinstall boundaries require explicit Owner confirmation.
- RC.5 Checkpoints 01–10 cover private receipt, checksum, safe layout, internal manifest, LICENSE/docs/source identity, documented operator journey, clean-host Smart Install, Task 049 F002/F003 proof, immutable installation, and guided setup interruption/resume/configuration.
- Checkpoints 11–13 cover the existing protected credential boundary, notification preflight, and one actual externally confirmed SMTP receipt without credentials or private provider/recipient data entering chat, argv, Git, history or evidence.
- Checkpoint 14 performs the exact documented guided Guardian activation with no manual `systemctl` workaround. Checkpoint 15 independently proves the Task 056 boundary: canonical state root is a real non-symlink current-user-owned mode-0700 directory, no state-to-configuration compatibility symlink was created, configuration and state roots are distinct and safe, and privacy-safe systemd evidence shows no compatibility migration failure.
- Checkpoints 16–17 independently retest QWSG-055-F001 and QWSG-051-F001 through enabled, active and fresh integrity-checked Guardian evidence and complete readiness; `filesystem.local_semantics` must remain satisfied after setup/activation.
- Checkpoints 18–26 cover actual process/invocation/cadence/resource behavior, lingering guidance, logout/session behavior, physical reboot, automatic post-reboot return with new identity and fresh evidence, second Owner-confirmed notification receipt, explicit documented restart, uninstall preservation, and same-candidate reinstall/resume/reactivation/final readiness.
- Gate F — after separate authority, integrate only privacy-safe final evidence and issue exactly READY or NOT READY. READY requires every mandatory Gate and all 26 checkpoints PASS, externally verified correction evidence for QWSG-055-F001, fresh QWSG-051-F001 and F002/F003 retests, two actual receipt confirmations, reboot/restart/uninstall/reinstall success, and no open blocker or security defect.
- Record a mandatory post-acceptance release/distribution requirement: official artifacts must later be downloadable directly from `git.quantumwizard.hu`/Forgejo using stable archive and sidecar URLs, `wget`, `curl`, command-line checksum verification, and a path compatible with the future Smart Installer. Do not implement or publish that mechanism in Task 057 unless strictly required for acceptance; a newly discovered requirement for source change stops acceptance for separate authority.


## Out of Scope

- Planning, Builder installation and Gate A do not authorize candidate construction, transfer, external execution, credentials, acceptance, tagging, release creation, upload, publication, announcement or Task 058. Every Gate remains separately Owner-authorized.
- No hidden product repair, undocumented workaround, manual systemctl substitution, arbitrary host remediation, sudo/package installation, SSH trust weakening, privilege escalation, evidence rewriting, or continuation after a mandatory product/packaging/security/provenance defect.
- Never reuse the mutated RC.4 acceptance VPS as clean. Gate E requires Owner-confirmed full reinstall/reset before Checkpoint 01; otherwise stop.
- Never infer external correction of QWSG-055-F001 or QWSG-051-F001 from local tests, code inspection, service enabled state alone, process presence alone, or historical RC evidence.
- Never read, hash, modify, copy, stage, package, export, transfer or otherwise access Owner-owned `docs/architecture/QWCS_MIGRATION_BLUEPRINT.md`; pathname metadata-only exclusion checks are permitted.
- No credentials, SMTP secret, private key, token, recipient/provider identity, private host/account identity or raw private paths in chat, commands when avoidable, Git, lifecycle history, acceptance records, snapshots or artifacts.
- No final `v1.1.0` tag, Forgejo Release, public artifact upload, stable download URL mutation, publication, announcement or final release authority. READY is evidence only.
- Do not create Task 058 or implement the future Forgejo distribution mechanism under Task 057 without separate Owner authority.


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

- Verify UTC date, ordinary user, exact project root and branch `main`; canonical HTTPS origin; `HEAD == origin/main ==` direct Forgejo `main` at `05a49a2e56d254ab3eb4646dc9df04fa4b63e335`; ahead/behind `0/0`; empty index; clean tracked tree; zero active prompts; canonical idle with Task 056 complete and archived.
- Verify the only unrelated state is the excluded Owner-owned untracked blueprint using pathname metadata only.
- Verify `VERSION=1.1.0-rc.5`; Task 056 integration commit `fc52295f6cc5b61078b32535230bdc704168d13c` and lifecycle-closing commit `05a49a2e56d254ab3eb4646dc9df04fa4b63e335` are the exact baseline ancestry; checkout/export build-contract and release validate-only plumbing pass; no RC.5 archive, sidecar, acceptance ledger, protocol, tag or Forgejo Release exists.
- Verify Task 056 records truthfully state the packaged systemd unit remained unchanged, local/regression validation passed, no RC.5 candidate or external acceptance occurred, and QWSG-055-F001 remains historical OPEN/BLOCKING.
- Verify immutable RC.1, RC.2, failed RC.3 and failed RC.4 evidence; QWSG-053-F001, QWSG-051-F001, QWSG-055-F001 and Task 049 F002/F003 histories; QWSG v1.0.0 tag/release identities; LICENSE; and Owner exclusion.
- Audit the release builder/package allowlist, RC.5 notes, installation/operator documentation, Smart Install, setup/state preparation, notification, Guardian/systemd unit, installer/uninstaller, Task 055 protocol/ledger and release-plumbing validator to determine the exact Gate A allowlist.
- Run baseline Framework, Builder, lifecycle, diversion, active-job/test-task, build, focused, full/race/vet/format, shell, Git whitespace, release-plumbing, security and preservation checks; stop on unrelated failure.


## Snapshot Requirements

Before every modifying or artifact-producing phase, create and verify a unique private mode-0700 snapshot under `/tmp` containing a readable exact tracked-HEAD archive, exact phase-owned files and absence records, Git/mode/ACL/tool identity, approved Builder source/input identities where applicable, protected-history hashes, and literal bounded restore instructions. Before construction record exact source commit/epoch, empty output destinations and two independent roots. Before transfer/external work retain only privacy-safe candidate/checkpoint state. Exclude Owner content, credentials, private host/account identity, candidate bytes when not required, caches and unrelated files. Verify archive readability, hashes, modes, absence claims and collision-aware rollback before each gated mutation.


## Risk Assessment

- False release readiness — critical: READY requires every fresh mandatory RC.5 gate and checkpoint; prior RC success and local Task 056 tests are context only.
- State-directory false correction — critical: QWSG-055-F001 closes externally only when a clean real systemd environment proves a private real state directory, no compatibility symlink, guided activation, enabled/active service, fresh canonical evidence and stable filesystem assessment.
- Provenance/reproducibility failure — critical: dirty or wrong source, `.git` dependence, ambient GOFLAGS/VCS metadata, wrong commit/epoch, non-identical twins or unsafe packaging stops before transfer.
- Host cleanliness — critical: the RC.4 VPS is mutated evidence; absent Owner-confirmed full reinstall/reset, Gate E cannot start.
- Credential/privacy exposure — critical: secrets and private host/provider/recipient identity never enter chat, argv when avoidable, Git, lifecycle records, evidence or artifacts.
- External mutation — critical: every state-changing checkpoint is Owner-operated, separately confirmed, restartable and bounded; defects stop without repair.
- Historical-evidence corruption — critical: RC.1–RC.4 and F001/F002/F003 chronology remain immutable; new RC.5 evidence is additive.
- Distribution-scope creep — high: stable Forgejo wget/curl delivery is mandatory post-acceptance work but does not belong in RC.5 acceptance unless acceptance proves it indispensable.
- Owner-content exposure — critical: the excluded blueprint is never accessed beyond pathname metadata.
- Authority expansion — critical: READY authorizes no tag, Forgejo Release, upload, publication, announcement or Task 058.


## Planned Work

1. Validate canonical idle, exact Git/source/protected baseline, Task 056 ancestry and rollback snapshot; audit RC.5 readiness and prepare minimal 26-checkpoint protocol, empty ledger and release-plumbing scaffolding.
2. Stop at Gate A for exact path-based scaffolding integration authorization. The resulting clean full commit, not the pre-scaffolding baseline or working tree, becomes the only candidate-source identity.
3. Under Gate B authority, export that exact commit twice into new independent private no-`.git` module roots; prove ordinary exported builds with GOFLAGS unset; derive one commit epoch and independently construct deterministic RC.5 twins from each module root.
4. Under Gate C authority, prove binary/manifest/archive/sidecar identity, provenance, static platform, layout/types/modes/timestamps, manifest, documentation, LICENSE and exclusions; record privacy-safe evidence without rebuilding.
5. Under Gate D authority, transfer exactly the verified archive and sidecar through a strict authenticated private boundary, or stop for explicit Owner-workstation fallback; verify destination identities without extraction/execution.
6. Under Gate E authority and only after Owner-confirmed full VPS reinstall/reset, execute Checkpoints 01–26 one bounded checkpoint at a time. Use documented product workflows; collect independent Task 056/systemd evidence; stop on every mandatory defect without workaround.
7. Preserve findings and incomplete checkpoint states truthfully. Source correction requires a new commit and RC identity; do not repair or relabel RC.5.
8. Under Gate F authority, integrate privacy-safe evidence and issue exactly READY or NOT READY. Stop before final tag/release/publication authority.
9. Carry forward the mandatory post-acceptance Forgejo distribution requirement for stable archive/sidecar wget/curl URLs and checksum examples into the separately authorized release/distribution phase.


## Rollback Plan

- Restore only literal phase-owned targets from the verified snapshot after Owner authorization; remove only paths whose prior absence and current identity are proven. Never reset, clean, broadly restore, rewrite history, alter tags, touch Owner content or overwrite unrelated state.
- A failed Gate B/C preserves bounded privacy-safe logs and exact inputs, quarantines non-transferable output, and stops. Retry only under controlled-failure policy and new Owner authority; never change source, provenance or GOFLAGS to manufacture identity.
- Transfer rollback removes only the two verified destination receipt files after explicit Owner authority when required; never weaken SSH trust or expose authentication material.
- External rollback is checkpoint-specific. Preserve evidence, protected credentials, configuration and QWSG state where the protocol requires; do not manually repair product state, reuse a non-clean host as clean, or continue after a blocker.
- Re-run applicable Git/lifecycle/build/package/security/preservation and checkpoint-integrity checks after bounded rollback and report the exact safe restart point.


## Deliverables

- RC.5-specific Owner-operated 26-checkpoint protocol and initially NOT STARTED privacy-safe evidence ledger with explicit Gates A–F.
- One private reproducible RC.5 candidate and sidecar only after Gate B/C authority, with independent twin-build, provenance, package and security proof.
- Private transfer integrity record or explicit safe stop/fallback decision only after Gate D authority.
- Fresh complete clean-host evidence proving the Task 056 state-directory correction in real systemd, retesting QWSG-055-F001, QWSG-051-F001 and Task 049 F002/F003, plus required notification receipts, Guardian/session/reboot/restart and uninstall/reinstall/resume behavior.
- Exact final READY or NOT READY verdict, additive finding chronology, preservation evidence and rollback record; no final release/publication action.
- A recorded mandatory post-acceptance distribution requirement for stable Forgejo archive and sidecar URLs usable with wget/curl, documented checksum verification and future Smart Installer compatibility.


## Verification

- Repository/lifecycle/Git identity, Task 056 ancestry, exact RC.5 version, no conflicting RC.5 artifact/protocol/ledger/tag/release, and exact Gate A path allowlist.
- Ordinary checkout build plus two genuine no-`.git` exported builds with GOFLAGS unset; truthful unknown defaults, exact explicit identity, controlled byte identity and no ambient Go VCS stamping.
- Two exact-commit release exports/builds from their module roots with full commit and commit-derived epoch; byte-identical binary, manifest, archive and sidecar; exact static Linux amd64 identity and safe package layout/types/modes/timestamps/docs/LICENSE/exclusions.
- Fresh Checkpoints 01–26 on an Owner-confirmed fully reinstalled/reset disposable Ubuntu 24.04 amd64 host. The RC.4-mutated environment supplies no clean-host substitution.
- Guided setup/activation proves the canonical state root is created/validated as a real current-user-owned mode-0700 non-symlink directory before systemd-dependent activation; systemd creates no compatibility symlink; service remains enabled and active; fresh integrity-checked Guardian evidence appears; readiness and filesystem.local_semantics remain satisfied.
- QWSG-055-F001 and QWSG-051-F001 receive independent fresh boundary-specific evidence; Task 049 F002/F003 receive a fresh uncoached Smart Install retest. Historical records remain unchanged.
- Protected credential mode/type boundary, notification preflight, two independent Owner-confirmed actual receipts, logout/session behavior, physical reboot, automatic post-reboot new/fresh identity, explicit restart, uninstall preservation and same-candidate reinstall/resume/reactivation all pass.
- Full Go/race/vet/format, focused build/release/package/install/uninstall, shell/static systemd, Git whitespace, Framework, Builder, lifecycle, diversion, active-job/test-task, secret/privacy/path-type, snapshot/rollback and protected-hash validations pass at applicable gates.
- READY only if all mandatory gates and 26 checkpoints pass with no open release blocker/security defect. Otherwise NOT READY with exact unresolved findings. Neither verdict authorizes tag, Forgejo Release, upload, publication or announcement.


## Documentation Updates

- Add `docs/release/ACCEPTANCE_PROTOCOL_1.1.0-rc.5.md` and `docs/release/ACCEPTANCE_1.1.0-rc.5.md` only under Gate A; initialize every checkpoint and gate truthfully as NOT STARTED.
- Update Task 057 history throughout authorized phases with privacy-safe source/build/transfer/checkpoint/finding/verdict chronology. Never include credentials, private host/provider/recipient identity or unnecessary raw private paths.
- Update `scripts/test-release-plumbing.sh` only as narrowly required to validate the new 26-checkpoint RC.5 protocol and initial ledger while preserving immutable RC.1–RC.4 assertions. Do not modify it if final Gate A inspection proves no change is necessary.
- Preserve Task 056 archive/history and RC.1/RC.2/failed RC.3/failed RC.4 evidence without rewriting historical findings.
- Record, but do not implement in Task 057, the mandatory later distribution requirement: official Forgejo-hosted stable archive and sidecar URLs, wget and curl examples, command-line SHA-256 verification, and future Smart Installer compatibility.
- Do not create final stable release notes, tag, Forgejo Release, public artifact/upload/publication material or Task 058.


## Completion Criteria

Task 057 completes only after the exact clean Gate A candidate-source commit is fixed; deterministic twin RC.5 construction and full package/provenance/security proof pass; private transfer is separately authorized and destination-verified; and the Owner completes every mandatory fresh Checkpoint 01–26 on a genuinely clean disposable VPS. QWSG-055-F001 must be externally proven corrected by real private non-symlink state storage, absence of the systemd compatibility symlink, successful guided activation, enabled/active service, fresh canonical Guardian evidence and satisfied post-activation filesystem assessment. QWSG-051-F001 and Task 049 F002/F003 require fresh boundary-specific PASS evidence; two actual notification receipts, logout/reboot continuity, explicit restart, uninstall preservation and same-candidate reinstall/resume/reactivation must pass. The terminal verdict is exactly `READY FOR QWSG 1.1.0 RELEASE` or `NOT READY FOR QWSG 1.1.0 RELEASE`. Completion preserves all historical evidence and records the mandatory later Forgejo wget/curl distribution requirement. It never authorizes a final tag, Forgejo Release, public upload, publication, announcement or Task 058.


## Owner Approval Requirements

Approved by Project Owner through the Engineering Task Builder on 2026-08-24 UTC.

The structured task definition has been explicitly approved for implementation. Further scope changes require explicit Project Owner approval.
