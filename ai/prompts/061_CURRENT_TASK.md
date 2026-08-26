# Current Engineering Task 061: QWSG 1.1.0-rc.6 Candidate Construction and Targeted External Re-Acceptance

## Task Metadata

- Task ID: `061`
- Task slug: `qwsg-1-1-0-rc-6-candidate-construction-targeted-external-re-acceptance`
- Status: `complete`
- Date opened: `2026-08-25` UTC
- Human authority: Project Owner
- Owner or lead-developer communication language: English

## Title

QWSG 1.1.0-rc.6 Candidate Construction and Targeted External Re-Acceptance


## Objective

Construct the exact QWSG 1.1.0-rc.6 linux-amd64 release candidate from the canonically closed Task 060 source state, prove deterministic candidate identity and integrity, and perform one bounded targeted external re-acceptance on the existing disposable Ubuntu 24.04 amd64 VPS. Establish whether QWSG-059-F001 is externally corrected while preserving the critical installation, state, notification, pre-login autostart, systemd recovery, explicit restart, uninstall-preservation and same-candidate reinstall behavior already proven or left incomplete in Task 059. Produce an explicit RC.6 Practical Acceptance PASS/FAIL and release-readiness recommendation without authorizing or claiming the final QWSG 1.1.0 release.


## Scope

- Verify canonical idle after Task 060, exact source commit `25a30718bc92882e9773a5c405ad648c0eee1a81`, `VERSION=1.1.0-rc.6`, Framework 1.1.0, clean tracked/index state, canonical remote/ref equality, Task 059 FAIL evidence, Task 060 correction evidence, and absence of any RC.6 candidate/tag/Release/publication.
- Create and verify the required private rollback-capable Task 061 snapshot before repository, candidate or external-host mutation. Preserve Task 059/060 and RC.5 evidence immutably.
- Construct two independent deterministic QWSG 1.1.0-rc.6 linux-amd64 release outputs from the exact authorized source commit and one explicit reproducible `SOURCE_DATE_EPOCH`, using the canonical release script in isolated private mode-0700 work/output directories. Compare twin archive and sidecar bytes before selecting one immutable acceptance candidate.
- Verify and record the selected candidate's source commit, build time/epoch, archive name, size and SHA-256, sidecar SHA-256 and check result, safe archive layout, manifest SHA-256 and all entries, binary SHA-256/version/provenance, required documentation/LICENSE set, modes, and absence of ambient VCS metadata or unauthorized content. Candidate identities are privacy-safe acceptance evidence.
- Freeze the selected candidate bytes after construction. Transfer exactly the verified archive and sidecar to the existing authorized Task 059 disposable VPS through the available private verified transfer path; destination verification must precede extraction or installation.
- Establish the VPS starting state before mutation: supported Ubuntu 24.04 amd64 ordinary-user context, installed RC.5 identity/state, service enablement/activity/lingering, preserved configuration/credential/state safety, canonical evidence/readiness status, and absence of RC.6. Retain only privacy-safe classifications and hashes.
- Replace installed RC.5 artifacts with exact RC.6 only through the documented verified release replacement workflow: stop the exact user unit, preserve the documented private backup, run RC.6 Smart Install/readiness, invoke `install.sh --replace --backup-dir` with a new bounded backup directory, reload the user manager, and restore the prior enabled/active intent. Do not manually edit runtime state, installed files, unit files, checkpoint, Scheduler state, Current Operator State, configuration or credential data.
- Revalidate installed artifact integrity, version/provenance, Smart Install, supported platform, safe configuration/state separation, state-root/path-component non-symlink properties, ownership/modes, no compatibility migration, service enablement/activity, and documented configuration preservation after replacement.
- Through documented QWSG configuration commands, perform one reversible bounded non-secret configuration-identity change after recording its original value. Do not alter notification addresses, provider identity or credentials merely to create the identity change.
- Perform one physical VPS reboot and prove pre-login user-manager/default-target/Guardian autostart, packaged unit identity, current recovered invocation/checkpoint ownership, current effective configuration identity, systemd automatic recovery behavior where naturally observed, and convergence after the required bounded cycles to fresh Guardian-running canonical evidence and Guardian readiness. Reconcile controlled SMTP acceptance plus Owner-confirmed receipt independently; valid notification-unverified and overall-partial readiness do not fail this path.
- If natural reboot does not exercise `Restart=on-failure`, perform at most one explicitly bounded user-service-process failure injection through the systemd user manager after stable ready evidence and a private state snapshot. Do not modify files or state. Require automatic bounded recovery, a new invocation, fresh evidence convergence and ready state; stop if recovery does not occur exactly as packaged.
- Perform one bounded controlled notification verification and obtain protected Owner-side receipt confirmation when needed. Under the current product contract this independently proves the configured transport; it need not create Guardian monitoring-queue evidence, persist as `notification.external=satisfied`, or make composite readiness literally `overall: ready`. Preserve the product's truthful `unknown_requires_verification` / `overall: partial` classifications separately.
- Restore the non-secret test configuration value through the documented QWSG CLI, perform the documented explicit Guardian restart, and prove fresh canonical convergence without manual state repair.
- Complete the remaining applicable practical lifecycle checks: disable/stop the exact unit, uninstall RC.6 with its exact verified archive, prove release-artifact removal plus documented preservation of configuration/credential/state, reinstall the exact unchanged RC.6 candidate, resume guided setup/configuration as documented, reactivate Guardian, and prove final installed artifact integrity and operational readiness.
- Reconcile Task 061 evidence additively against immutable Task 059/060 records. Produce RC.6 candidate identity, practical acceptance, defect, limitation, final state and release-readiness documentation; run required validation; perform task-scoped Git integration, clean fast-forward pushes, direct Forgejo ref verification, and canonical Task 061 closure.


## Out of Scope

- Do not modify product source, release scripts, candidate contents or candidate metadata after candidate construction. A need for source or byte correction is a release-blocking STOP, not an in-task fix.
- Do not rebuild or relabel RC.5, rewrite Task 059/060 evidence, convert missing evidence to PASS, or claim the targeted RC.6 run retroactively changes an earlier verdict.
- Do not reset/reinstall the VPS operating system, manually repair or delete Guardian/Scheduler/checkpoint/Current Operator State, manually edit QWSG configuration or credentials, substitute a unit, install dependencies, install Postfix, or use an ad-hoc workaround to manufacture readiness.
- Do not intentionally corrupt configuration, credentials, installed artifacts, state files, ownership, modes, symlink boundaries or network settings. The only planned failure injection is one bounded systemd-managed process termination if natural reboot supplies no automatic-recovery evidence.
- Do not expose credentials, addresses, provider/account identities, private host/network identities, tokens, passwords, private keys, private SMTP responses or message content in chat, Git, commands where avoidable, evidence, snapshots or artifacts.
- Do not broaden architecture or implement unrelated fixes/features.
- Do not create or move a tag, create a Forgejo Release, upload or publish assets, activate public downloads, deploy, announce, or claim QWSG 1.1.0 is released.
- Do not create, prepare or install Task 062.


## Authority Envelope

1. **Authorized paths/components/systems:** Task 061 may use the exact canonical Task 060 repository state, canonical deterministic release tooling, private local build/verification/rollback directories, Task 061 lifecycle/evidence/release-readiness documentation, and the existing disposable Task 059 Ubuntu 24.04 amd64 VPS. It may operate only on QWSG release artifacts, the exact QWSG user service, QWSG-owned installed artifacts, and documented per-user QWSG configuration/state necessary for the approved acceptance sequence.
2. **Routine operations:** inspect, fetch/read refs, snapshot, build two isolated deterministic release outputs, hash/compare/verify/freeze one candidate, transfer it privately, inspect the VPS privacy-safely, run documented Smart Install/replacement/setup/readiness/config/restart/uninstall/reinstall workflows, reboot once, wait bounded cycles, collect privacy-safe evidence, diagnose, retest, document, explicitly stage reviewed Task 061 paths, commit, push dry-run, clean-fast-forward push, verify refs, and close lifecycle without routine intermediate Owner gates.
3. **Permitted external actions:** one documented RC.5-to-RC.6 replacement, one reversible non-secret QWSG configuration-identity change and restoration, one physical reboot, one optional systemd-user-managed bounded process-failure injection when natural evidence is absent, one explicit Guardian restart, one RC.6 uninstall, and one exact-same-candidate reinstall are authorized. Every operation requires verified exact targets, documented preservation/rollback, and post-action validation. No operating-system reset or unrelated infrastructure change is authorized.
4. **Owner interaction is not a new engineering gate:** pause only for an inherently Owner-side private candidate transfer, protected credential entry, actual notification receipt confirmation, or physical reboot action Aikó cannot perform safely. Provide one concise exact action block, accept only privacy-safe returned classifications, then continue under the same Task 061 authority.
5. **Correction/retest authority:** recoverable procedural, transfer, evidence-ordering and test-orchestration issues follow diagnose -> smallest authorized correction -> retest -> continue. Late trustworthy evidence may be reconciled. Missing evidence remains missing. No correction may change source, candidate bytes, QWSG runtime state manually, security boundaries or acceptance meaning.
6. **Candidate immutability:** after twin equality and selection, record identities and make the candidate read-only in private retained storage. All local, transfer and destination hashes must match. Any mismatch or need to rebuild after acceptance begins triggers STOP; a distinct build is not the accepted candidate.
7. **Repository integration:** privacy-safe Task 061 evidence and required release-readiness documentation may be integrated through explicit path staging, reviewed task-scoped commits, push dry-run, clean fast-forward push to `origin/main`, and direct read-only Forgejo ref verification. Broad staging, history rewrite, force push and unrelated paths are forbidden.
8. **Lifecycle completion:** after truthful completion or accepted STOP evidence, Task 061 may finalize its prompt/history, archive the prompt, push a lifecycle-only closure commit, validate canonical idle, and report without another routine Owner gate. It may not prepare or create Task 062.
9. **Evidence and rollback:** before mutation preserve verified mode-0700 local rollback evidence and, before bounded VPS changes, privacy-safe inventories plus documented product backup/preservation evidence. Exclude Owner content, secrets, private infrastructure identity, candidate secrets, caches and unrelated data. Preserve safe real state directories, ownership/modes, symlink rejection, atomic state, generation isolation and truthful degraded/not-ready behavior.
10. **Owner-reserved operations:** source or release-tool modification; candidate-byte modification/rebuild after selection; architecture expansion; dependency installation; unplanned destructive or irreversible action; privilege/infrastructure mutation beyond exact documented QWSG install/uninstall/replacement and approved reboot; credentials except Owner entry; operating-system reinstall; tag; Forgejo Release; upload/publication; deployment; announcement; final QWSG 1.1.0 release authorization; and Task 062.
11. **Mandatory STOP conditions:** stop for source/commit/version mismatch, nondeterministic twins, candidate identity/integrity mismatch, unsafe archive or destination, inability to establish the truthful VPS baseline, required source/candidate change, security/privacy regression or uncertainty, corrected Guardian convergence failure, failure of previously proven critical RC.5 behavior, automatic recovery failure, unavailable reliable rollback, unplanned destructive/external mutation, or any need beyond this envelope. Preserve FAIL evidence and do not repair around a product defect.


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

- Verify UTC date, ordinary local user, exact repository root, branch `main`, canonical HTTPS origin, fetched `HEAD == origin/main ==` direct Forgejo `main` at `25a30718bc92882e9773a5c405ad648c0eee1a81`, ahead/behind `0/0`, empty index, clean tracked tree, zero active prompts, and canonical idle with Task 060 complete/archived.
- Verify the only unrelated state is the excluded Owner blueprint by pathname metadata only; do not read, hash, copy, stage, package or modify it.
- Verify Framework `1.1.0`, `VERSION=1.1.0-rc.6`, Task 060 implementation commit `7e3f94667fa57cae91969ab3de80e99a87eb1730`, closure commit `25a30718bc92882e9773a5c405ad648c0eee1a81`, QWSG-059-F001 root cause/correction/regressions, and exact unchanged Task 059 FAIL/RC.5 identities.
- Verify current release notes, installation/replacement/uninstall documentation, deterministic build contract, release script/plumbing, packaged unit and candidate contents before use. Run Framework, Builder, lifecycle, diversion/test-task, release-plumbing, full/race/vet/format/build-contract, shell/systemd/static, security/privacy, historical-preservation and Git baseline checks.
- Verify no RC.6 archive, sidecar, candidate directory, tag, Forgejo Release, upload, publication, VPS action, credential action or Task 062 exists before Task 061 execution.
- Before external mutation, establish the VPS facts read-only and privacy-safely. Expected baseline from Task 059 is RC.5 installed, service enabled/active, lingering enabled, preserved protected notification configuration/credential and safe state, canonical evidence degraded/partial and readiness not-ready; differences must be diagnosed and materially unsafe differences trigger STOP.


## Snapshot Requirements

Before repository, candidate or VPS mutation, create and verify a unique private mode-0700 Task 061 snapshot under `/tmp`. Include a readable exact tracked-HEAD archive; literal Task 061 lifecycle/evidence target before-images and absence records; Git/ref/status/version/mode/ACL evidence; Builder input/prompt identities; immutable Task 059/060/RC.5 hashes; release-script/build-contract/unit/install/uninstall hashes; candidate/output absence claims; and exact bounded rollback instructions. Exclude Owner blueprint content, credentials, provider/address/host identity, candidate bytes, caches, build outputs and unrelated files. Before external replacement, retain privacy-safe VPS baseline classifications and use the documented new private replacement backup directory; never copy secrets into task evidence. Before lifecycle closure create an additional bounded closure snapshot. Verify readability, hashes, modes, exclusions, collision safety and restore instructions.


## Risk Assessment

- Candidate/source identity error — critical: exact commit/version/epoch, isolated twin builds, byte comparison, archive/sidecar/manifest/binary verification and destination re-verification are mandatory.
- False external PASS — critical: every required convergence/readiness/lifecycle assertion needs trustworthy evidence; missing or contradictory proof is FAIL/limitation, never inferred success.
- Manual-state workaround — critical: only documented product/service workflows may mutate QWSG; checkpoint/Scheduler/current-state edits or deletion are forbidden.
- Configuration-change fidelity — critical: record and change one non-secret value through QWSG CLI, prove effective identity changed, reboot without repair, and restore only through the CLI after defect proof.
- Recovery-test disruption — high: use natural reboot evidence first; optional process failure injection is single, systemd-managed, bounded, performed only from stable ready state with rollback evidence, and must stop on unexpected behavior.
- Credential/privacy exposure — critical: Owner-only entry/receipt, no secret/address/provider/host identity retained, privacy-safe command construction and evidence redaction.
- State/security regression — critical: preserve real non-symlink mode-0700 state, ownership, configuration separation, atomic integrity, generation isolation and corrupt-state fail-closed behavior.
- Replacement/uninstall data loss — critical: exact archive, verified installed artifacts, documented stop/backup/replace/uninstall sequence, preservation proof and same-candidate reinstall; no purge.
- Historical corruption — critical: Task 059 remains FAIL, Task 060 remains local correction evidence, RC.5 identities remain immutable, and Task 061 is additive only.
- Publication overreach — critical: candidate construction/transfer is private acceptance work only; no tag, Release, upload, public URL, deployment, announcement or final release claim.


## Planned Work

1. Validate canonical idle, exact Task 060 source/ref/version, immutable Task 059/060 evidence, release workflow, exclusions and baseline tests; create/verify the private rollback snapshot.
2. Choose the explicit reproducible epoch from the exact source commit, construct two isolated deterministic rc.6 outputs with the canonical script, prove archive/sidecar equality, select one candidate, make it read-only and record all required identities/integrity/package evidence.
3. Establish the existing VPS read-only baseline before transfer or mutation. Transfer only the selected archive and sidecar through the private verified path and reverify exact identities at destination before extraction.
4. Safely inspect/extract the package, verify manifest/binary/docs/license/provenance, and run Smart Install. Stop on any identity, package or supported-host discrepancy.
5. Use the documented stop/private-backup/`install.sh --replace --backup-dir`/daemon-reload/start sequence to replace RC.5 with RC.6 while preserving configuration, credential and state. Verify installed artifacts and all state/security invariants.
6. Record one non-secret configuration value and effective identity; change it through `qwsg config set`; verify the identity changed without manually touching state. Obtain notification verification only if required for readiness.
7. Perform the physical reboot. Correlate boot, user manager/default target, pre-login Guardian invocation, packaged unit, checkpoint generation/configuration and cycles. Wait bounded required cycles and require fresh running evidence, complete operator state, notification satisfaction, Guardian ready and overall ready.
8. If reboot supplies no natural `Restart=on-failure` event, from stable readiness perform the single authorized systemd-managed process failure injection and require automatic recovery/new invocation/fresh complete ready evidence. Do not edit/delete state or repeat a failed injection.
9. Restore the original non-secret value via QWSG CLI; perform the documented explicit Guardian restart and require fresh complete ready convergence. Local deterministic tests remain the evidence for truthful genuine-active-failure degradation; do not damage the host to prove it.
10. Disable/stop the exact service, uninstall with the exact RC.6 archive, verify owned-artifact removal and configuration/credential/state preservation, reinstall the same immutable candidate, resume documented setup/activation, and prove final installed integrity plus operational readiness.
11. Reconcile Task 059/060/061 evidence additively; classify every acceptance area PASS/FAIL/NOT EXECUTED, defects and limitations separately, final rc.6 state, practical verdict and technical release-readiness recommendation without release authorization.
12. Run all required product/release/governance/security/privacy/preservation validations; explicitly stage only reviewed Task 061 paths, commit, dry-run/push clean fast-forward, verify direct Forgejo refs, snapshot/close lifecycle to canonical idle, and report.


## Rollback Plan

- Before candidate selection, discard only isolated Task 061 build outputs whose exact private directories are proven. After selection, retain immutable candidate evidence; do not silently rebuild or replace it.
- Before external replacement, stop on missing backup/preservation evidence. For replacement rollback, stop the exact Guardian, restore only recorded RC.5 release-owned artifacts from the documented new private backup, daemon-reload and restore prior enabled/active intent; preserve configuration/credential/state. If state compatibility is rejected, leave service stopped and retain evidence for Owner direction.
- For the non-secret test configuration, restore only the recorded original value through documented QWSG CLI and revalidate effective identity. Never edit the configuration file directly.
- For uninstall/reinstall failure, retain the exact RC.6 candidate and preserved user data; restore only verified release-owned artifacts through documented installer/backup procedure. Never recursively delete installation or user directories.
- Repository rollback restores only literal Task 061 targets from the verified snapshot after current identity/collision checks and removes only new files with proven prior absence/current Task 061 identity. Never use broad reset/checkout/restore/clean, history rewrite, force push, tag mutation or Owner-content access. Pushed corrections use a forward commit.
- After any rollback rerun candidate/source identities as applicable, installed artifact/state/security/readiness checks, focused/full product tests, release plumbing, Framework/lifecycle, privacy/historical preservation and exact Git/ref validation.


## Deliverables

- One exact immutable privately retained QWSG 1.1.0-rc.6 linux-amd64 candidate archive and sidecar constructed from source commit `25a30718bc92882e9773a5c405ad648c0eee1a81`, with deterministic twin equality and complete privacy-safe identity/integrity/package evidence.
- A precise external starting-state and documented RC.5-to-RC.6 replacement record with safe configuration/credential/state preservation and installed-artifact verification.
- External proof or truthful FAIL for the QWSG-059-F001 scenario: configuration identity change, physical reboot, pre-login autostart, recovered generation ownership, fresh complete canonical convergence and readiness without manual state repair.
- Proof of packaged automatic recovery, documented explicit restart, notification satisfaction where required, uninstall preservation and exact-same-candidate reinstall, plus local preservation evidence for genuine active failure truth.
- An additive RC.6 Practical Acceptance report that preserves Task 059/060 history, classifies all required areas, defects/limitations/final state, and states whether rc.6 is technically suitable to proceed toward separately authorized QWSG 1.1.0 release preparation.
- Complete Task 061 history, snapshots/rollback evidence, validated task-scoped commits/push/ref evidence, and canonical lifecycle closure.


## Verification

- `HEAD`, `origin/main`, direct Forgejo main, source commit, version, epoch and clean source tree match; Framework/lifecycle/release baseline passes and protected histories hash unchanged.
- Two isolated canonical release builds produce byte-identical archive and sidecar. Selected archive/sidecar are regular non-symlinks, read-only after selection, and match recorded size/SHA-256 at every transfer boundary.
- Archive has one safe canonical root, no duplicate/absolute/traversal/special entries, expected modes, complete manifest, exact binary provenance/version, required LICENSE/README/INSTALL/docs/rc.6 notes, and no secrets/private identity/ambient VCS metadata.
- VPS baseline, transfer, extraction, Smart Install and installed artifact verification are privacy-safe and exact. Replacement uses only documented commands with a new private backup and preserves configuration, credential and state safely.
- Configuration identity demonstrably changes through one documented non-secret CLI change. No checkpoint, Scheduler state, Current Operator State, unit or installed artifact is manually edited, removed or repaired.
- Physical boot identity changes. User manager/default target and Guardian start pre-login; packaged unit remains enabled/active with lingering; recovered checkpoint generation and current configuration match. After bounded cycles, canonical evidence is fresh, Guardian is running and Guardian readiness is ready. Controlled SMTP acceptance and Owner-confirmed receipt satisfy practical notification acceptance independently; current-contract `notification.external=unknown_requires_verification` and `overall: partial` are valid and must not be manufactured into PASS states.
- Natural or single controlled systemd recovery evidence proves `Restart=on-failure`, new invocation ownership, stale-generation isolation, fresh complete convergence and readiness. Genuine current-generation failure remains degraded/not-ready in deterministic local regression evidence.
- Explicit restart after documented configuration restoration converges. RC.6 uninstall removes only unchanged release-owned artifacts and preserves configuration/credentials/state; same immutable candidate reinstall restores exact artifacts and final operational readiness.
- Full Go, race, vet, formatting, ordinary build/build-contract, shell syntax, systemd static, release plumbing, installer/uninstaller, setup/state/user-service/readiness/notification, Framework, Builder, lifecycle, diversion/test-task, security/privacy/secret, historical-preservation, rollback, Git diff/mode/staging/commit/push/ref and canonical-idle checks pass.
- No source/candidate mutation after selection, no RC.6 replacement candidate, tag, Forgejo Release, upload/publication/deployment/announcement, final release claim, VPS OS reset or Task 062 occurs. Missing required evidence prevents PASS.


## Documentation Updates

- Maintain and canonically archive Task 061 prompt/history with chronological privacy-safe evidence, exact candidate identities, external sequence, diagnostics, limitations, rollback, commits and refs.
- Add the rc.6 practical acceptance/candidate identity record and update only the narrowly required release-readiness/index/plumbing documentation. Preserve Task 059 practical FAIL, Task 060 correction record, RC.5 protocols/acceptance/notes and all earlier evidence byte-for-byte.
- Record that candidate construction and private targeted acceptance are distinct from final release authorization. State explicitly that QWSG 1.1.0 is not released and external acceptance does not create tag/Release/publication authority.


## Completion Criteria

Task 061 is complete only when one exact rc.6 candidate is deterministically constructed and fully verified; the existing VPS starting state is truthfully established; documented RC.5-to-RC.6 replacement preserves security/state boundaries; the configuration-change plus reboot scenario proves or truthfully fails QWSG-059-F001 convergence; pre-login autostart, automatic recovery, controlled SMTP acceptance plus Owner-confirmed receipt, explicit restart, uninstall preservation and same-candidate reinstall are verified as applicable; all required evidence and validation pass; candidate/source/history identities remain immutable; the final rc.6 practical verdict, defects, limitations, operational state and technical release-readiness recommendation are explicit; task-scoped commits are clean-fast-forward pushed and directly verified; and Task 061 closes to canonical idle. The current product contract permits notification external unknown and overall partial after independently successful controlled delivery; those truthful states do not fail this acceptance path. PASS requires every mandatory critical boundary. Any candidate mismatch, source/byte correction need, convergence/security/critical-regression failure or missing mandatory proof produces FAIL/NOT READY and STOP evidence. Completion never releases QWSG 1.1.0 or authorizes a tag, Forgejo Release, upload, publication, deployment, announcement or Task 062.


## Owner Approval Requirements

Approved by Project Owner through the Engineering Task Builder on 2026-08-25 UTC.

The structured task definition and Authority Envelope have been explicitly approved. The task is authorized to start and execute every routine operation inside that envelope without another Owner gate. Further scope changes and every Owner-reserved operation require explicit Project Owner approval.
