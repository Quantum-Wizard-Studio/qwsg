# Current Engineering Task 056: QWSG systemd User-State Directory Compatibility Correction

## Task Metadata

- Task ID: `056`
- Task slug: `qwsg-systemd-user-state-directory-compatibility-correction`
- Status: `active`
- Date opened: `2026-08-24` UTC
- Human authority: Project Owner
- Owner or lead-developer communication language: English

## Title

QWSG systemd User-State Directory Compatibility Correction


## Objective

Correct QWSG's clean-host systemd user-state-directory interaction without weakening the filesystem-security model. Before guided Guardian activation reaches systemd startup, resolve the canonical state root and securely create or validate a real current-user-owned mode-0700 non-symlink directory through the existing trusted primitive. Preserve fail-closed behavior for every unsafe existing path, integrate the correction as QWSG 1.1.0-rc.5 source, and prepare a replacement candidate path whose later clean-host acceptance restarts at Checkpoint 01 on a freshly reinstalled disposable VPS.


## Scope

- Audit guided setup, userservice activation, canonical state-root discovery, Guardian state/store initialization, assessment path selection, packaged systemd unit storage directives, installer/uninstaller handling, Task 052 runtime-context behavior, and Task 055 QWSG-055-F001 evidence.
- Implement the smallest coherent pre-activation state preparation boundary: resolve the same canonical state root used by the packaged unit, invoke the existing secure private-root primitive, and proceed to manager reachability, daemon-reload, and enable/start only after validation succeeds.
- Preserve `StateDirectory=qwsg`, `StateDirectoryMode=0700`, `WorkingDirectory=%S/qwsg`, `Environment=QWSG_STATE_DIR=%S/qwsg`, `ReadWritePaths=%S/qwsg`, and existing service hardening unless focused evidence proves a narrower unit change is necessary.
- Fail closed on existing symlinks, wrong ownership, wrong type, unsafe modes, unsafe components, unavailable roots, and state-path mismatch. Expose deterministic activation-stage/cause evidence without raw host paths.
- Add focused unit/integration tests for clean config-present/state-absent activation, safe pre-creation ordering, idempotence, unsafe-state rejection, userservice call ordering, readiness stability, Guardian evidence, restart and exit-report behavior, and installer/uninstaller compatibility.
- Advance source identity and release plumbing to `1.1.0-rc.5`; add private RC.5 notes and narrowly update operator/installation documentation where behavior or identity changes.
- Preserve Task 055 as NOT READY and QWSG-055-F001 as OPEN/BLOCKING until a separately authorized replacement candidate passes fresh external acceptance.
- Prepare exact path-based integration and later RC.5 construction/acceptance gates; do not construct candidate bytes without separate Owner authority.


## Out of Scope

- Never weaken, bypass, or special-case QWSG's symlink rejection or private state ownership/mode checks.
- Never automatically remove, replace, follow, chmod, chown, migrate, or convert an existing unsafe state path.
- No manual repair of the current test VPS, Guardian restart, RC.4 acceptance continuation, or reuse of that mutated VPS as a clean host.
- No sudo, privilege escalation, shell execution, arbitrary systemctl commands, systemd hardening removal without proof, credential handling, SMTP action, external VPS access, artifact transfer, candidate construction, tag, Forgejo Release, upload, publication, Task 057, or final release authorization.
- Never read, hash, modify, copy, stage, package, or otherwise access Owner-owned `docs/architecture/QWCS_MIGRATION_BLUEPRINT.md`; metadata-only exclusion checks are permitted.
- Task installation does not authorize implementation. Implementation does not authorize source integration, RC.5 construction, transfer, external acceptance, release, or publication; each remains a separate Owner gate.


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

- Verify UTC date, ordinary user, exact project root, branch `main`, canonical HTTPS origin, and `HEAD == origin/main ==` direct Forgejo `main` at the Task 055 lifecycle-closing commit; ahead/behind `0/0`; clean index and tracked tree; zero active prompts; canonical idle with Task 055 complete and archived.
- Verify the only unrelated state is the excluded Owner-owned untracked blueprint by pathname metadata only.
- Verify `VERSION=1.1.0-rc.4`, Task 054 correction commit `ef513dde187e4119b6aa04a3439a879056f6cc69` and Task 055 acceptance-source commit `f105771dbddccf363a63095ac4ad2a7a2285aa84` are ancestors.
- Verify Task 055 verdict `complete with disclosed limitations — NOT READY FOR QWSG 1.1.0 RELEASE`, QWSG-055-F001 OPEN/BLOCKING, the exact RC.4 candidate identity and deterministic evidence, and all unexecuted checkpoint states.
- Reproduce the architectural audit locally without touching the external host: configuration is created before activation; state is absent; the unit requests `StateDirectory=qwsg`; systemd v255 compatibility migration may create a state-to-configuration symlink; QWSG rejects it as unsafe.
- Verify RC.1, RC.2, failed RC.3, QWSG-053-F001, QWSG-051-F001, Task 049 F002/F003, v1.0.0 tag/release identities, and LICENSE remain unchanged.
- Run baseline Framework, Builder, lifecycle, diversion, active-job/test-task, build, focused, full/race/vet/format, shell, Git whitespace, and release-plumbing checks. Record any expected RC.4 terminal-ledger assertion requiring a narrow RC.5 plumbing update; stop on any unrelated failure.


## Snapshot Requirements

Before modifying any Task 056 target, create and verify a unique private mode-0700 snapshot under `/tmp` containing a readable exact tracked-HEAD archive, exact affected files and absence records, Git/mode/ACL/tool identity, Builder source/input identities, protected-history hashes, and literal bounded restore instructions. Exclude Owner content, credentials, private host identity, candidate bytes, caches and unrelated files. Verify archive readability, target hashes/modes, absence claims, and rollback commands before implementation and before every separately gated integration or construction phase.


## Risk Assessment

- Security regression — critical: accepting or repairing symlinked/wrong-owner/wrong-mode state would violate the QWSG trust model; retain strict rejection and negative tests.
- Path-contract mismatch — critical: application fallback resolution and systemd `%S/qwsg` must identify the same root; prove XDG/default behavior and fail closed on ambiguity.
- Activation ordering/race — high: state preparation must complete before any manager reload/start and remain safe under repeated activation; prove call ordering and idempotence.
- False recovery — critical: local success cannot close QWSG-055-F001 or imply release readiness; a new RC.5 candidate and fresh external Checkpoint 01 restart are mandatory.
- systemd compatibility — high: preserve useful `StateDirectory` lifecycle/hardening semantics while preventing migration symlink creation by valid pre-existence; test the packaged unit on systemd v255-compatible fixtures and verify the unit statically.
- Assessment drift — high: filesystem.local_semantics must continue to assess the canonical real state path consistently before and after setup.
- Regression breadth — high: Task 052 runtime context, Smart Install, installer/uninstaller, notification, restart and exit-report behavior must not change unintentionally.
- Authority/privacy — critical: no external host, credentials, Owner content, candidate, release or publication action is permitted during implementation without a later gate.


## Planned Work

1. Validate canonical idle, exact Git/protected baseline, Task 055 terminal evidence, release identity and rollback snapshot; stop on variance.
2. Trace the exact setup-to-activation call graph and canonical state-root selection. Establish one narrow state-preparation dependency with deterministic failure classification and no raw-path leakage.
3. Reuse `operatorstate.EnsurePrivateRoot` or a minimal shared wrapper to validate/create the canonical state root before `userservice.Controller.Activate` performs manager reachability, daemon-reload or enable/start. Do not duplicate or weaken the primitive.
4. Preserve packaged unit storage/hardening directives unless tests prove a necessary narrower adjustment. Ensure systemd encounters an existing real private directory and therefore does not create the compatibility symlink.
5. Add focused tests for ordering, clean config-present/state-absent behavior, real mode-0700 state creation, no symlink, idempotence, unsafe symlink/owner/type/mode rejection without repair, no sudo/shell/privilege expansion, deterministic activation errors, enabled/active/fresh Guardian behavior, assessment stability, restart/ExecCondition/exit-report, Task 052 runtime context, Smart Install and installer/uninstaller.
6. Advance `VERSION` and all canonical release-plumbing/documentation identities to `1.1.0-rc.5`; preserve RC.4 evidence and prohibit reuse of RC.4 identity for changed bytes.
7. Run complete local product, package, security, governance and rollback validation. Present an exact path allowlist and stop for separate source-integration authority.
8. After separately authorized integration, require separate gates for private deterministic RC.5 construction/proof, private transfer, fresh external acceptance from Checkpoint 01 on a fully reinstalled disposable VPS, final evidence integration and verdict.


## Rollback Plan

- Before integration, restore only explicit Task 056 target files from the verified snapshot after Owner authorization and remove only files whose prior absence was recorded. Never reset, clean, broadly restore, or touch Owner/unrelated content.
- If state preparation or tests reveal a mismatch, preserve evidence and stop; do not weaken checks, repair an unsafe path, alter the test VPS, or retain a partial product workaround.
- After an authorized commit, use a separately approved bounded revert only if required; never rewrite published history or reuse RC.4 identity.
- Rerun focused security/path/order tests, full validation, lifecycle, protected hashes, Git state and exclusion checks after rollback. External acceptance rollback is out of scope and the current mutated VPS must remain classified non-clean.


## Deliverables

- A smallest-safe state-preparation implementation that creates or validates the canonical real private state directory before systemd activation and fails closed on unsafe existing state.
- Focused regression coverage proving configuration-present/state-absent success, mode/owner/type/symlink safety, operation ordering, idempotence, readiness stability, Guardian enabled/active/fresh evidence, and preserved Task 052/Smart Install/installer/uninstaller/restart/exit-report behavior.
- QWSG `1.1.0-rc.5` source/release-plumbing identity and private release notes without candidate bytes.
- Updated operator/installation or architecture documentation only where the corrected storage contract requires it.
- Exact integration allowlist, validation evidence, rollback record, preservation audit, and later gated RC.5 clean-host acceptance plan requiring a full VPS reinstall and Checkpoint 01 restart.


## Verification

- Focused state-root and activation-controller unit tests prove the state root exists as a real current-user-owned mode-0700 directory before the first systemd-dependent action; no compatibility symlink appears; repeated activation is idempotent.
- Existing state symlink, symlink component, wrong owner, wrong type and unsafe mode all fail before manager/reload/start; no chmod/chown/removal/following/replacement, sudo, shell or privilege escalation occurs.
- A realistic clean-home fixture with configuration present and state absent proves guided activation, enabled and active unit state, and fresh integrity-checked canonical Guardian evidence. Filesystem local-semantics remains satisfied after setup.
- Userservice call ordering and fixed argv/environment remain deterministic; Task 052 validated runtime context remains intact; daemon-reload, enable/start, bounded restart, ExecCondition and exit-report behavior remain deterministic.
- Notification behavior and credential boundaries are unchanged. Smart Install F002/F003, installer, uninstaller, configuration preservation and reinstall/resume tests pass.
- `VERSION`, embedded identity, archive naming, documentation and release plumbing consistently identify `1.1.0-rc.5`; RC.4 evidence remains immutable and no RC.5 candidate is constructed during implementation.
- Run canonical build, focused tests, full Go tests, repository-wide race tests, vet, format, shell syntax/static checks, systemd unit verification, Git whitespace, Framework, Builder, lifecycle, diversion, active-job and test-task checks, security/exclusion audit, snapshot/rollback validation and protected-history hashes.
- Verify exact staged allowlist and complete staged diff only under a later Owner integration gate; push only after separate authority and clean fast-forward proof.
- Replacement external acceptance requires a freshly reinstalled disposable Ubuntu 24.04 amd64 VPS and restarts from Checkpoint 01; no evidence from the mutated RC.4 host substitutes for clean-host evidence.


## Documentation Updates

- Update Task 056 history throughout authorized implementation with privacy-safe decisions, validation, rollback and unresolved evidence.
- Add `docs/release/RELEASE_NOTES_1.1.0-rc.5.md` and narrowly update `VERSION`, release plumbing and installation/operator identity references required for the new source bytes.
- Update guided-activation/state-storage documentation only if the implementation changes an operator-visible or architectural contract.
- Preserve Task 055 archive/history and `docs/release/ACCEPTANCE_1.1.0-rc.4.md` as immutable NOT READY chronology; do not rewrite RC.1/RC.2/RC.3 or historical finding records.
- Do not create RC.5 acceptance evidence/protocol, tag, release, publication material, Task 057, or any Owner-content reference beyond the required exclusion rule without separate authority.


## Completion Criteria

Task 056 completes only when the smallest safe correction is implemented and integrated under separate Owner authority; clean config-present/state-absent guided activation prepares a real current-user-owned mode-0700 non-symlink canonical state root before systemd operations; unsafe existing paths fail closed without repair; Guardian activation remains enabled, active and produces fresh canonical evidence in controlled validation; filesystem.local_semantics remains satisfied; Task 052, Smart Install, installer/uninstaller, notification, restart, ExecCondition and exit-report regressions pass; full product/security/governance/rollback checks pass; and source identity is consistently `1.1.0-rc.5`. Completion does not close QWSG-055-F001 by assertion, construct or transfer RC.5, access a VPS, handle credentials, authorize release, or begin Task 057. Fresh replacement-candidate acceptance must restart at Checkpoint 01 on a fully reinstalled disposable VPS under separate Owner gates.


## Owner Approval Requirements

Approved by Project Owner through the Engineering Task Builder on 2026-08-24 UTC.

The structured task definition has been explicitly approved for implementation. Further scope changes require explicit Project Owner approval.
