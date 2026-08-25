# Current Engineering Task 059: QWSG 1.1.0-rc.5 Practical Clean-Host Acceptance

## Task Metadata

- Task ID: `059`
- Task slug: `qwsg-1-1-0-rc-5-practical-clean-host-acceptance`
- Status: `complete with release-blocking product defect — practical acceptance FAIL`
- Date opened: `2026-08-25` UTC
- Human authority: Project Owner
- Owner or lead-developer communication language: English

## Title

QWSG 1.1.0-rc.5 Practical Clean-Host Acceptance


## Objective

Perform the first complete practical release acceptance of the exact verified QWSG 1.1.0-rc.5 candidate on the freshly reinstalled disposable Ubuntu 24.04 amd64 VPS under the Framework 1.1.0 Practical Release Acceptance model. Execute the complete twelve-step product journey under one bounded Owner authorization, collect privacy-safe trustworthy evidence, diagnose and correct recoverable in-scope product/test issues without altering candidate bytes, and classify practical acceptance PASS or FAIL, product defects, procedural/evidence limitations, final RC.5 operational state, and technical suitability to proceed toward QWSG 1.1.0 release preparation. This task does not release QWSG 1.1.0.


## Scope

- Validate the untouched freshly reinstalled disposable Ubuntu 24.04 amd64 VPS as a supported clean-host baseline using privacy-safe evidence.
- Receive or download the exact Task 057 RC.5 archive and sidecar through the existing private verified transfer path when no published Forgejo Release asset exists; do not require publication for acceptance.
- Preserve and verify candidate identity: version `1.1.0-rc.5`; source commit `1025d36d05b2f6f919f0ea4ec4a7029f67536000`; archive `qwsg-1.1.0-rc.5-linux-amd64.tar.gz`; archive size `2951350` bytes; archive SHA-256 `cfe300c0f1f312d80120f74a9f24bed4a64387471bf2097ddc63d94f0fb2f7b0`; sidecar SHA-256 `69f3eb4bf89dc126a7eafd08354eec37a941014171b3d1d70c6e6a4cf52e5eb0`; `MANIFEST.sha256` SHA-256 `ae51aca0bc4ddc61b0daea3a87f0acabcde5ec9fd8fadddc050f0786d6915e9e`; binary SHA-256 `5484aab96d5c3748e81b065fdb11ec8c34385589bb07ee7ea1b2b35fdffa6b93`.
- Execute the twelve steps in `docs/release/PRACTICAL_RELEASE_ACCEPTANCE.md`: fresh baseline; candidate receipt; integrity/package safety; Smart Install/readiness; documented installation; guided setup; Guardian/state contract; protected external notification; physical reboot; documented explicit Guardian restart; uninstall preservation; same-candidate reinstall and final readiness.
- Independently verify Guardian uses a safe real state directory with no compatibility symlink, correct ownership and mode, configuration/state separation, enabled and active service, fresh canonical evidence, and satisfied `filesystem.local_semantics` before reboot and where applicable after recovery/reinstall.
- Obtain exactly one Owner-confirmed actual external notification receipt through Owner-only protected credential entry without retaining protected identities or secrets.
- Reconcile trustworthy late or differently ordered evidence when candidate identity, host continuity, technical validity, and evidence independence remain provable. Missing mandatory evidence remains missing and cannot be PASS. Reporting delay alone does not invalidate clean-host status.
- Diagnose, correct where explicitly authorized, retest, and continue after recoverable in-scope product/test issues. If correction would modify product source, candidate bytes, package identity, or the accepted RC.5 artifacts, stop and report the defect before any alteration.
- Create and maintain new privacy-safe Task 059 additive acceptance evidence and the matching task history; update only the appropriate current acceptance/release-readiness documentation while keeping all historical evidence immutable.
- Perform routine validated Git integration, clean fast-forward push, post-push verification, and canonical Task 059 lifecycle closure when the truthful terminal result is fully recorded.


## Out of Scope

- Do not access the VPS during Task 059 planning or Engineering Task Builder installation. External access begins only after the Builder-approved task is installed and started under its Authority Envelope.
- Do not recreate or enforce Task 057's 26-checkpoint micro-gating model, require Owner approval between routine steps, or treat evidence reporting order as a product gate.
- Do not rewrite, relabel, delete, supersede, or otherwise mutate Task 057 or any earlier prompt, history, acceptance protocol, ledger, candidate chronology, finding, or release evidence. RC.5 results are additive only.
- Do not modify QWSG product source, rebuild/repackage RC.5, alter candidate bytes, substitute another candidate, or apply undocumented developer workarounds. A required source or candidate change is a defect and mandatory STOP.
- Do not create or move a Git tag; create a Forgejo Release; upload assets; publish, announce, or deploy QWSG; expose a stable public release URL; or claim QWSG 1.1.0 is released.
- Do not require Forgejo publication when the existing private verified RC.5 transfer path is available.
- Do not expose credentials, recipient addresses, provider identities, tokens, passwords, private keys, private host/account identities, or other secrets in chat, Git, task history, acceptance evidence, command arguments where avoidable, snapshots, logs, or release artifacts.
- Do not access Owner-owned `docs/architecture/QWCS_MIGRATION_BLUEPRINT.md` beyond pathname metadata needed to exclude it from task operations.
- Do not expand product architecture, alter production infrastructure, or begin QWSG 1.1.0 release preparation/publication work.


## Authority Envelope

1. **Authorized paths/components/systems:** Task 059 owns its active prompt/history and new privacy-safe practical RC.5 acceptance evidence; it may update the minimum current release-readiness/acceptance documentation needed for an additive verdict. It may operate the disposable Ubuntu 24.04 amd64 acceptance VPS and exact verified RC.5 archive/sidecar through the documented product workflows. The Project Owner participates only for protected credential entry, actual receipt confirmation, and decisions at a mandatory STOP.
2. **Routine operations:** After Builder installation and task start, one bounded Owner authorization covers read-only baseline inspection, exact candidate receipt/download through the existing verified private path, integrity/package checks, Smart Install/readiness, documented install, guided setup, Guardian activation and verification, protected notification preflight/send/receipt confirmation, real reboot and recovery checks, documented explicit Guardian restart, uninstall contract checks, exact same-candidate reinstall/reactivation, privacy-safe evidence capture/reconciliation, local snapshots, documentation, validation, and reporting. It also covers documented task-required ordinary-user and bounded administrative operations, including documented install/uninstall commands, service operations, and reboot, without intermediate Owner gates.
3. **Correction/retest authority:** Recoverable in-scope diagnostic, host prerequisite, product-operation, evidence-capture, or test-orchestration issues follow diagnose -> apply the smallest documented authorized correction -> retest -> continue. Corrections may adjust task-owned evidence/commands or documented host prerequisites only when they preserve clean-host validity and candidate identity. No correction may change product source, candidate bytes, package contents, provenance, security boundaries, or manufacture missing evidence; such a need is a mandatory STOP and product defect report.
4. **Repository integration:** After privacy/security review and required validation, Task 059 may explicitly stage only reviewed task-owned paths, review staged diffs and modes, create truthful task-scoped commits, run push dry-runs, clean-fast-forward push to `origin/main`, and verify local/remote refs. No broad staging, history rewrite, force push, tag operation, Release operation, or unrelated-content mutation is authorized.
5. **Lifecycle completion:** Task 059 may truthfully finalize its prompt/history and additive acceptance evidence, classify PASS or FAIL and all required limitations/states, archive the prompt, integrate lifecycle-only closure, clean-fast-forward push it, validate canonical idle, and report completion without another routine Owner gate. Lifecycle closure may record failure or disclosed limitations; it may never convert missing evidence to PASS.
6. **Permitted external actions:** Access and bounded mutation of only the disposable acceptance VPS are authorized for the twelve documented acceptance steps after task start. Exact candidate transfer/download through the existing private verified path is authorized. One protected external notification attempt and Owner receipt confirmation are authorized, with Owner-only local credential entry. Read-only verification of relevant repository/Forgejo state is permitted. No tag, Forgejo Release, asset upload, public publication, announcement, production deployment, or unrelated external-system mutation is permitted.
7. **Evidence and rollback:** Before modifying repository targets or the VPS, create and verify proportional rollback/recovery evidence that does not contaminate the clean-host baseline or retain secrets. Preserve exact candidate hashes, privacy-safe host continuity, command outcomes, package/state/service/evidence facts, notification classification, reboot/restart recovery, uninstall preservation, reinstall readiness, exact repository diffs/modes/refs, and chronological history. Evidence may be reconciled later only when trustworthy; missing mandatory evidence is never PASS. Never capture credentials or protected identities. Preserve historical evidence immutably.
8. **Owner-reserved operations:** Actual protected credential entry and actual notification receipt confirmation remain Owner-only pauses within the same Task 059 authority, not new engineering approval gates. Also reserved are material scope or architecture changes; product source/candidate modification; candidate rebuild/repackage/substitution; undocumented repair; unplanned destructive action; unresolved privilege/security/privacy decisions; unrelated infrastructure mutation; tags; Forgejo Releases; uploads; publication; deployment; announcement; and authorization to prepare or release QWSG 1.1.0.
9. **Mandatory STOP conditions:** Stop only for unavailable safe rollback/recovery; a material mismatch in clean-host baseline, candidate identity, provenance, integrity, package safety, or supported environment; a demonstrated product/package/security defect; a correction requiring product source or candidate-byte change; missing mandatory evidence at final classification; exposed or uncertain protected data; an unplanned destructive or external action; authority/scope ambiguity; meaningful risk outside the bounded VPS/repository task; or any Owner-reserved operation. Pause, rather than terminate, for Owner-only credential entry or receipt confirmation, then resume under the same authority. A recoverable in-scope issue is not itself a STOP.


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

- Before any task mutation, verify UTC date, ordinary local user, exact repository root, branch `main`, canonical HTTPS origin, `HEAD == origin/main ==` direct Forgejo `main` at the Task 058 lifecycle-closing commit, ahead/behind `0/0`, empty index, clean tracked tree, zero active prompts, and canonical idle with Task 058 complete and archived.
- Verify the only unrelated worktree state is the excluded Owner blueprint by pathname metadata only; do not read, hash, copy, modify, stage, package, or otherwise access its contents.
- Verify Framework `1.1.0`, `VERSION=1.1.0-rc.5`, the complete Authority Envelope/Builder contract, `docs/release/PRACTICAL_RELEASE_ACCEPTANCE.md`, Task 058 completion, Task 057 immutable terminal classification, and the exact candidate identities stated in this task.
- Verify no Task 059 prompt/history/evidence already exists, no RC.5 candidate bytes or historical evidence changed after Task 058, and no tag, Forgejo Release, upload, publication, announcement, or release-preparation mutation occurred.
- Run Framework, Builder-input, lifecycle, diversion, active-job, test-task, release-plumbing, security/privacy, and relevant repository baseline checks before modification.
- Do not access or test the disposable VPS during planning or Builder installation. After the approved task starts, the first external action is a privacy-safe baseline check proving the Owner-stated fresh Ubuntu 24.04 amd64 condition before candidate execution or host mutation.


## Snapshot Requirements

Before modifying repository task targets, create and verify a unique private mode-0700 snapshot under `/tmp` containing a readable exact tracked-HEAD archive; literal copies/absence records for every proposed Task 059 target; Git/ref/status evidence; modes/ACLs where relevant; Builder input/prompt identities; protected historical-evidence hashes; and exact bounded rollback instructions. Exclude Owner content, credentials, private infrastructure identity, candidate bytes, caches, and unrelated files. Before external mutation, record a privacy-safe clean-host baseline and a bounded recovery plan that does not itself prepare or contaminate the VPS. Before each materially distinct repository integration/lifecycle phase, ensure the retained snapshot still provides reliable bounded rollback. Verify archive readability, hashes, modes, absence claims, collision safety, exclusions, and restore instructions before proceeding.


## Risk Assessment

- False practical PASS — critical: all twelve mandatory steps and applicable security boundaries require trustworthy evidence; missing proof remains missing and produces FAIL/NOT READY.
- Clean-host contamination — critical: no VPS access occurs before task start; document the first baseline before candidate execution, and reinstall only if technical cleanliness is actually destroyed.
- Candidate identity drift — critical: verify every supplied identity and use the exact same archive for initial install and reinstall; any byte/source/package change stops.
- Product defect concealment — critical: documented bounded correction is allowed only without source/candidate change; no hidden developer workaround or evidence relabeling.
- Credential/privacy exposure — critical: Owner-only protected entry; retain no secrets, provider/recipient identity, or private host/account identity in chat, argv where avoidable, evidence, snapshots, Git, or artifacts.
- Remote/disruptive operation — high: installation, service actions, reboot, restart, uninstall, and reinstall intentionally mutate only the disposable VPS and require continuity/recovery evidence.
- State safety regression — critical: independently prove real non-symlink mode-0700 current-user state, configuration separation, service state, canonical evidence freshness, and local filesystem semantics.
- Uninstall data loss — critical: verify the documented preservation/removal contract before reinstall; avoid any undocumented cleanup.
- Historical corruption — critical: Task 057 and all earlier records remain byte-immutable; Task 059 evidence is new and additive.
- Publication overreach — critical: acceptance and technical suitability grant no tag, Forgejo Release, upload, publication, announcement, deployment, or final release authority.
- Rollback/evidence coupling — high: snapshots and logs must support repository rollback and truthful host diagnosis without contaminating the clean-host claim or retaining protected data.


## Planned Work

1. Validate canonical idle, Framework 1.1.0, Git/ref/version state, immutable Task 057/058 evidence, exact candidate identities, exclusions, and the Task 059 Authority Envelope; create and verify the repository snapshot. Do not access the VPS during this planning/Builder boundary.
2. After the Builder-approved task starts, establish the privacy-safe fresh Ubuntu 24.04 amd64 baseline as the first VPS interaction, including architecture, ordinary-user/service-manager context, and absence of QWSG-specific preparation.
3. Receive or download exactly the verified RC.5 archive and sidecar through the existing private path if no published Release asset exists; verify file type, size, archive/sidecar hashes, `sha256sum -c`, archive safety/layout, internal manifest, documentation, LICENSE, binary identity, and source identity.
4. Run and preserve meaningful Smart Install/readiness output. Apply only documented permitted host prerequisites, then retest readiness without bypassing product checks.
5. Install through the packaged documented workflow and verify installed version/commit, immutable files, unit, and documentation. Run guided setup as the ordinary Guardian user.
6. Activate Guardian through the documented guided path and independently verify state-root type/ownership/mode/path safety, no compatibility symlink, configuration/state separation, enabled/active service, fresh canonical evidence, and `filesystem.local_semantics`.
7. Pause for Owner-only protected notification credential entry, run documented preflight/send without secret retention, obtain one Owner-confirmed actual receipt, record only the privacy-safe classification, and resume under the same authority.
8. Perform a real VPS reboot and verify automatic Guardian recovery, enabled/active service, expected persistence, host continuity, and fresh post-reboot canonical evidence.
9. Perform the documented explicit Guardian restart and verify bounded recovery and fresh evidence.
10. Run the documented uninstaller and verify exact release-owned removal plus promised configuration, credential, and state preservation. Do not perform undocumented cleanup.
11. Reinstall the identical verified RC.5 archive, safely resume preserved configuration, reactivate as documented, and verify final operational readiness with fresh evidence.
12. Reconcile reliable out-of-order evidence without weakening proof, classify practical acceptance PASS or FAIL, product defects, procedural/evidence limitations, final RC.5 operational state, and suitability to proceed toward release preparation. Update only new/current privacy-safe acceptance records.
13. Run all proportional repository, release-plumbing, lifecycle, privacy/security, documentation, diff/mode, and rollback validations. Explicitly stage reviewed task paths, commit, push dry-run, clean-fast-forward push, verify refs, truthfully close Task 059, return to canonical idle, and report. Do not tag, publish, upload, announce, deploy, or claim release.


## Rollback Plan

- Repository rollback restores only literal Task 059-owned targets from the verified private snapshot after proving identities and collision conditions. Remove only new paths whose prior absence and current Task 059 identity are proven. Never use broad reset, checkout, restore, clean, wildcard deletion, history rewrite, force push, tag mutation, or Owner-content access.
- Prefer forward correction and retest for recoverable in-scope repository/evidence issues. If a task commit has been pushed, use a new bounded corrective commit rather than rewriting published history.
- VPS recovery follows only the documented product workflow and the task's bounded continuity plan. Uninstall is an intentional acceptance step, not a generic rollback. Do not manually erase preserved user configuration, credentials, or state, and do not reinstall the host merely because evidence was reported late.
- If host mutation invalidates the clean baseline before its required proof, preserve privacy-safe failure evidence and stop; only the Owner may authorize another host reset. If product correction requires changed source or candidate bytes, preserve evidence and stop without altering RC.5.
- Credential rollback/removal uses only documented protected local mechanisms with Owner participation and never records secrets or protected identities.
- After any rollback or bounded recovery, rerun applicable candidate, service/state, security/privacy, repository, Framework, lifecycle, Git, and evidence-integrity checks and record the exact result.


## Deliverables

- One complete privacy-safe additive Task 059 practical acceptance record covering all twelve mandatory workflow steps for the exact verified RC.5 candidate.
- Exact integrity/package verification, meaningful Smart Install output, documented install/setup results, Guardian state/service/evidence proof, one Owner-confirmed actual notification receipt classification, reboot/restart recovery evidence, uninstall preservation evidence, and same-candidate reinstall/final-readiness evidence.
- A clear terminal classification: practical acceptance PASS or FAIL; product defects, if any; procedural/evidence limitations, if any; final RC.5 operational state; and whether RC.5 is technically suitable to proceed toward QWSG 1.1.0 release preparation.
- Updated minimum current acceptance/release-readiness documentation and complete Task 059 chronological history, with all Task 057 and earlier evidence immutable.
- Proportional validation, verified rollback, exact Git integration/ref evidence, clean fast-forward push, and canonical lifecycle closure.
- No source/candidate mutation, tag, Forgejo Release, upload, publication, announcement, deployment, or claim that QWSG 1.1.0 is released.


## Verification

- Canonical Framework/Git/lifecycle baseline and zero pre-start VPS access.
- Fresh supported Ubuntu 24.04 amd64 baseline, ordinary-user context, required service manager, and absence of QWSG-specific preparation.
- Exact RC.5 version/source/archive name/size/archive hash/sidecar hash/manifest hash/binary hash and `sha256sum -c` PASS; regular non-symlink receipt files; safe archive types/paths/layout; complete manifest; required docs and LICENSE.
- Smart Install/readiness executed with meaningful preserved product output and supported-host readiness; any permitted prerequisite is documented and retested.
- Documented install and guided setup pass without product bypass; installed identity, immutable payload, unit, and documentation are correct.
- Guardian uses a real non-symlink safe current-user-owned mode-0700 canonical state root; no compatibility symlink; configuration/state separation; enabled/active service; fresh integrity-checked canonical evidence; `filesystem.local_semantics` satisfied.
- Protected notification preflight/send passes and exactly one actual receipt is Owner-confirmed without secret/provider/recipient/private-host retention.
- Physical reboot proves automatic Guardian recovery, enabled/active service, expected persistence, and fresh post-reboot evidence.
- Documented explicit Guardian restart proves bounded recovery and fresh evidence.
- Documented uninstall proves the promised release-owned removal and configuration/credential/state preservation contract.
- Reinstall uses the exact same verified RC.5 candidate and proves safe preserved-state resume, documented reactivation, final operational readiness, and fresh evidence.
- Reliable late/different-order evidence is reconciled only with identity, continuity, and independence; missing mandatory evidence is never PASS; reporting order alone does not invalidate cleanliness.
- Any defect requiring source/candidate change is reported without alteration. Product defects and procedural/evidence limitations are separately classified.
- Historical Task 057 and earlier evidence hashes/content remain unchanged; Owner blueprint content remains untouched; secret/private-identity scans pass.
- Full applicable Framework, Builder, lifecycle, diversion, active-job/test-task, release-plumbing, shell/format, documentation-link, security/privacy, Git whitespace, diff/mode, snapshot/rollback, staged-path, commit, push, ref, and canonical-idle checks pass.
- Final report says practical acceptance PASS or FAIL, final RC.5 operational state, technical suitability toward release preparation, and explicitly does not claim release or authorize publication.


## Documentation Updates

- Create a new privacy-safe Task 059 practical acceptance evidence/report artifact under the appropriate `docs/release/` location rather than rewriting the Task 057 protocol or ledger.
- Update the minimum current release-readiness/acceptance index or documentation needed to expose the additive Task 059 result, only if repository inspection proves such an update is appropriate.
- Maintain `ai/prompts/059_CURRENT_TASK.md` and its matching `ai/history/059_2026-08-25_qwsg-1-1-0-rc-5-practical-clean-host-acceptance.md` throughout execution, then archive the prompt canonically on truthful completion.
- Do not modify historical Task 057 or earlier prompts, histories, protocols, ledgers, release notes, finding chronology, or candidate identities.
- Do not update release/publication artifacts, tags, Forgejo Release metadata, announcement material, or claim QWSG 1.1.0 release.


## Completion Criteria

Task 059 is complete when the exact verified QWSG 1.1.0-rc.5 candidate has undergone all twelve Practical Release Acceptance steps on the freshly reinstalled supported Ubuntu 24.04 amd64 VPS; every mandatory evidence category is truthfully recorded or the task is explicitly classified FAIL/NOT READY for exact missing proof; product defects and procedural/evidence limitations are separately disclosed; the final installed/uninstalled/reinstalled operational state is exact; and the report states whether RC.5 is technically suitable to proceed toward QWSG 1.1.0 release preparation. Completion also requires privacy/security compliance, immutable historical evidence, proportional validation, verified rollback, task-scoped clean-fast-forward Git integration, and canonical lifecycle closure. PASS requires all twelve mandatory steps and applicable security boundaries to pass with no release blocker. A source/candidate-byte correction need, integrity/provenance/package/security defect, or unresolved mandatory evidence prevents PASS. No completion result releases QWSG 1.1.0 or authorizes a tag, Forgejo Release, upload, publication, announcement, or deployment.


## Owner Approval Requirements

Approved by Project Owner through the Engineering Task Builder on 2026-08-25 UTC.

The structured task definition and Authority Envelope have been explicitly approved. The task is authorized to start and execute every routine operation inside that envelope without another Owner gate. Further scope changes and every Owner-reserved operation require explicit Project Owner approval.
