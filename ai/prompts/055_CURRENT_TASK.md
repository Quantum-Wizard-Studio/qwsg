# Current Engineering Task 055: QWSG 1.1.0-rc.4 Clean-Host Acceptance and Release Readiness

## Task Metadata

- Task ID: `055`
- Task slug: `qwsg-1-1-0-rc-4-clean-host-acceptance-release-readiness`
- Status: `approved`
- Date opened: `2026-08-24` UTC
- Human authority: Project Owner
- Owner or lead-developer communication language: English

## Title

QWSG 1.1.0-rc.4 Clean-Host Acceptance and Release Readiness


## Objective

Construct one private, reproducible QWSG `1.1.0-rc.4` candidate from one exact clean integrated source commit, prove its deterministic identity and safe package, transfer it only through a separately authorized private channel, and execute a fresh Owner-operated clean-host acceptance from Checkpoint 01. Directly retest historical `QWSG-053-F001`, `QWSG-051-F001`, and Task 049 F002/F003, then conclude exactly `READY FOR QWSG 1.1.0 RELEASE` or `NOT READY FOR QWSG 1.1.0 RELEASE` without authorizing final release or publication.


## Scope

- Gate A — audit exact repository/source readiness and prepare only RC.4-specific acceptance protocol, empty evidence ledger, and narrowly required version-aware scaffolding. Integrate that scaffolding only after a separate reviewed path-based Owner authorization and clean fast-forward push.
- Gate B — after separate authorization, derive the exact clean candidate-source commit and its commit timestamp, create two independent private mode-0700 source exports with no `.git`, and build twice with the exact full lowercase commit and commit-derived `SOURCE_DATE_EPOCH`.
- Gate C — require byte-identical binaries, manifests, archives, and checksum sidecars; exact embedded version/commit/UTC identity; static Linux amd64; safe canonical layout, modes and deterministic metadata; byte-correct LICENSE and documentation; and complete security/exclusion proof. Stop before transfer.
- Gate D — after separate authorization, privately transfer exactly the archive and sidecar using strict host-key SSH/SCP when the authenticated boundary works, or stop for the Owner-workstation fallback. Verify destination type, count, size and hashes without execution.
- Gate E — after separate authorization, run a restartable Owner-operated clean-host protocol from Checkpoint 01. Every checkpoint records purpose, bounded action, expected evidence, PASS, FAIL/finding, continuation safety, and retain/redact requirements.
- Checkpoints 01–25 cover: receipt/provenance; archive checksum; safe layout; manifest; LICENSE/docs/binary/source identity; documented journey; install readiness; Smart Install F002/F003; install; setup interruption/resume; protected notification configuration; notification preflight; actual notification and Owner-confirmed receipt; guided Guardian activation; QWSG-051-F001 correction proof; independent enabled/active/fresh readiness; actual Guardian process/cadence/resources/restart evidence; lingering/session behavior; logout; physical reboot; automatic return with new process identity and fresh evidence; post-reboot notification and Owner-confirmed receipt; explicit restart; uninstall with preserved configuration/credentials/state; same-candidate reinstall, resume, reactivation and final readiness.
- `QWSG-053-F001` requires ordinary `make build` in genuine no-`.git` exports with no `GOFLAGS`, truthful unknown defaults, exact explicit identity, controlled byte identity, and no ambient Go `vcs.*` settings. Release construction still uses explicit commit/epoch and `-buildvcs=false`.
- `QWSG-051-F001` requires guided activation to succeed through the documented product workflow and independent evidence that the service is enabled, active, and producing fresh integrity-checked canonical Guardian state. No manual activation workaround counts.
- Task 049 F002/F003 require a fresh uncoached Smart Install run proving actionable systemd user-manager guidance and bounded filesystem local-semantics verification. Historical evidence is chronology only.
- Notification evidence uses the protected credential boundary and privacy-safe classifications; actual readiness requires independent Owner-confirmed receipt before and after reboot where the protocol requires it.
- Gate F — after separate authorization, integrate only privacy-safe evidence and issue exactly READY or NOT READY. READY requires every mandatory RC.4 gate and checkpoint to pass and authorizes no tag, Forgejo Release, upload, publication, or announcement.


## Out of Scope

- Planning and installation do not authorize candidate construction, transfer, external execution, credentials, acceptance, tagging, release creation, upload, publication, or Task 056.
- No hidden product repair, undocumented workaround, readiness weakening, manual systemctl substitution, arbitrary remediation, sudo/package installation, SSH trust weakening, privilege escalation, or continuation after a release-blocking product/security defect.
- No reuse of RC.1/RC.2/RC.3 artifact-dependent PASS evidence as RC.4 evidence; RC.4 has new source identity and bytes and starts at Checkpoint 01.
- Never read, hash, copy, modify, stage, package, transfer, or otherwise touch Owner-owned `docs/architecture/QWCS_MIGRATION_BLUEPRINT.md`; metadata-only exclusion checks are permitted.
- No final `v1.1.0` tag, Forgejo Release, public artifact, publication claim, or Task 056.


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

- Verify UTC date, ordinary user, exact root and `main`; canonical HTTPS origin; `HEAD == origin/main ==` direct Forgejo main at `12069b16cc574c759a40f905d2b4981bd729716d`; ahead/behind `0/0`; empty index; clean tracked tree; zero active prompts; canonical idle after completed Task 054.
- Verify the only unrelated state is the excluded Owner-owned untracked path by metadata only.
- Verify `VERSION=1.1.0-rc.4`, Task 054 integration commit `ef513dde187e4119b6aa04a3439a879056f6cc69` is an ancestor, ordinary checkout/export build-contract and release validate-only plumbing pass, and no RC.4 archive, sidecar, tag or release exists.
- Verify immutable RC.1/RC.2/failed RC.3 evidence; historical `QWSG-053-F001` and `QWSG-051-F001` OPEN/BLOCKING records; Task 049 F002/F003; v1.0.0 identities; LICENSE; and Owner exclusion.
- Audit current release builder, package allowlist, installation/operator documentation, Smart Install, setup, notification, Guardian/systemd, installer/uninstaller and acceptance precedents to determine the exact Gate A scaffolding allowlist.


## Snapshot Requirements

Before each modifying phase create and verify a unique private mode-0700 snapshot containing a readable exact tracked-HEAD archive, affected lifecycle/scaffolding files and absence records, Git/mode/ACL/tool identity, Builder source/input, protected preservation hashes, and literal bounded restore instructions. Before construction preserve exact source commit/epoch and output absence; before transfer/external checkpoints preserve only privacy-safe state. Exclude candidate bytes where unnecessary, all credentials/private host identity, and Owner content except metadata-only exclusion records.


## Risk Assessment

- False release readiness — critical: READY requires every mandatory fresh RC.4 checkpoint and cannot be inferred from prior candidates.
- Provenance/reproducibility failure — critical: any dirty source, wrong commit/epoch, `.git` dependence, non-identical twin output, collision, or unsafe package stops before transfer.
- Historical blocker false closure — critical: F001/F002/F003 require direct boundary-specific proof with no workaround.
- Credential/privacy exposure — critical: secrets, private host identity and sensitive evidence never enter chat, argv, Git, lifecycle records, snapshots or artifacts.
- External host mutation — critical: each Owner-operated action is bounded, restartable and gated; product/security defects stop acceptance.
- Owner-content exposure — critical: the excluded blueprint is never read, hashed, copied, staged or packaged.
- Authority expansion — critical: READY does not authorize tag, release, upload, publication, announcement or Task 056.


## Planned Work

1. Validate canonical idle, exact Git/protected baseline and Task 054 ancestry; snapshot; audit RC.4 readiness and prepare minimal protocol/ledger/scaffolding.
2. Stop at Gate A for exact path-based scaffolding integration authorization.
3. Under Gate B authority, export the exact clean integrated commit twice and construct private deterministic twins using explicit full commit and commit-derived epoch.
4. Under Gate C authority, prove binary/manifest/archive/sidecar identity, package provenance, layout, documentation, LICENSE, static target and exclusions; stop before transfer.
5. Under Gate D authority, privately transfer the exact two files and verify destination integrity; otherwise use only an explicitly approved Owner-workstation fallback.
6. Under Gate E authority, execute Checkpoints 01–25 one at a time from fresh receipt through Smart Install, setup, notification, Guardian, reboot, uninstall/reinstall/resume and final readiness.
7. Record findings truthfully and stop on product/security/reproducibility/privacy failure without hidden repair.
8. Under Gate F authority, integrate privacy-safe evidence and issue exactly READY or NOT READY; stop before final release authority.


## Rollback Plan

- Restore only literal phase-owned targets from the verified snapshot after Owner authorization; remove only outputs proven absent before the phase. Never reset, clean, overwrite candidate identity, alter tags, or touch Owner/historical content.
- Failed construction preserves bounded logs and exact inputs, then stops or retries only under the controlled failure policy; non-identical or unsafe bytes are never transferred or relabeled.
- Transfer/external rollback is checkpoint-specific and preserves configuration, protected credential files and QWSG state; no VPS teardown or general host administration is authorized.
- Rerun applicable Git, lifecycle, build, package, security, preservation and external integrity gates after rollback and report the exact safe restart point.


## Deliverables

- RC.4-specific Owner-operated 25-checkpoint protocol and privacy-safe evidence ledger with explicit Gates A–F.
- One private reproducible RC.4 candidate plus sidecar only after construction authorization, with twin-build/package/security evidence.
- Private transfer record or explicit stopped fallback decision only after transfer authorization.
- Fresh complete clean-host evidence retesting QWSG-053-F001, QWSG-051-F001 and Task 049 F002/F003, including actual notification receipts, Guardian/session/reboot and reinstall/resume behavior.
- Exact final READY or NOT READY verdict, preservation evidence and rollback record; no final release/publication action.


## Verification

- Repository/lifecycle/Git identity, Task 054 ancestry, exact RC.4 version and no conflicting artifact/tag/release.
- Ordinary checkout build plus two genuine no-`.git` export builds without GOFLAGS; truthful defaults, exact explicit identity, controlled byte identity and absence of ambient VCS settings.
- Two exact-commit release exports/builds with commit-derived epoch; byte-identical binary, manifest, archive and sidecar; exact static Linux amd64 identity and safe package/layout/modes/docs/LICENSE/exclusions.
- Fresh Checkpoints 01–25 with independent PASS evidence; guided activation proves enabled/active/fresh; Smart Install retests F002/F003; actual Owner-confirmed notification receipts; logout/reboot/new identity/fresh state/restart; uninstall preservation and same-candidate reinstall/resume.
- Full Go/race/vet/format, focused build/release/package/install/uninstall, shell/static, secret/privacy, Git whitespace, Framework, Builder, lifecycle, diversion, job/test-task, snapshot/rollback and protected-hash gates.
- READY only if all mandatory gates pass and no blocker/security defect remains; otherwise NOT READY with exact findings. No candidate construction during planning or installation.


## Documentation Updates

- Add RC.4-specific acceptance protocol and evidence ledger only under Gate A if absent and required; never overwrite RC.1/RC.2/RC.3 records.
- Update Task 055 history throughout authorized phases with privacy-safe source/build/transfer/checkpoint/finding/verdict evidence.
- Update only narrowly required release-plumbing assertions for RC.4 acceptance scaffolding while preserving prior candidate checks.
- Do not create final release notes beyond existing private RC.4 notes, tags, Forgejo Release records, publication material, Task 056, or Owner-content references.


## Completion Criteria

Task 055 completes only when exact clean RC.4 source is integrated, deterministic twin construction and complete package/security proof pass, private transfer is separately authorized and verified, and the Owner completes all mandatory fresh Checkpoints 01–25. QWSG-053-F001 must pass exported-source determinism/provenance; QWSG-051-F001 must pass documented guided activation plus independent enabled/active/fresh evidence; Task 049 F002/F003 boundaries must pass fresh Smart Install regression; actual notification receipts, reboot continuity, restart, uninstall preservation and reinstall/resume must pass. The final state is exactly READY FOR QWSG 1.1.0 RELEASE or NOT READY FOR QWSG 1.1.0 RELEASE. Completion never authorizes final tag, Forgejo Release, upload, publication, announcement or Task 056.


## Owner Approval Requirements

Approved by Project Owner through the Engineering Task Builder on 2026-08-24 UTC.

The structured task definition has been explicitly approved for implementation. Further scope changes require explicit Project Owner approval.
