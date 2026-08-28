# Current Engineering Task 068: QWSG Change Notifications & 1.2.0-rc.4

## Task Metadata

- Task ID: `068`
- Task slug: `qwsg-change-notifications-rc4`
- Status: `approved`
- Date opened: `2026-08-28` UTC
- Human authority: Project Owner: Attila
- Owner or lead-developer communication language: English

## Title

QWSG Change Notifications & 1.2.0-rc.4


## Objective

Integrate administrator notifications for meaningful QWSG-managed installation, update, rollback, version, configuration, and cleanly supportable Guardian lifecycle changes into the existing notification/configuration architecture, preserve operation truth independently from delivery results, and produce a deterministic private QWSG 1.2.0-rc.4 candidate for later final acceptance.


## Scope

- Establish and record the canonical starting state, mandatory protected snapshot, and deterministic bounded rollback before modifying task targets.
- Inspect the existing notification implementation, SMTP/email support, configuration model and managed configuration paths, installer notification guidance, updater, rollback, Guardian lifecycle, service management, localization, security model, and relevant tests before design.
- Reuse and narrowly extend existing notification transports/configuration/localization abstractions; do not introduce a parallel notification subsystem.
- Support administrator notification for successful installation, safe/meaningful post-capability installation failure, update success/failure, rollback success/failure, installed-version transition, authorized QWSG-managed configuration change, and cleanly integrable non-noisy Guardian activation/start or administrative transition.
- Include privacy-safe operational identity, event type, old/new or restored versions, result, timestamp, existing operation/event identifier where available, action-required state, and concise safe failure reason where available.
- Provide localized user-facing notification content for English, Hungarian, and German using existing localization conventions.
- Reuse supported configuration paths so administrators can enable/disable lifecycle notifications, configure/validate the recipient, and use or cleanly extend a deterministic test-notification mechanism.
- Define and expose operation result separately from notification delivery result; transport failure must neither corrupt a successful operation nor hide an operation failure.
- Add deterministic duplicate/idempotency protection where the architecture supports event/operation identity, without loops or repeated spam.
- Add automated mock/local-boundary coverage for all mandated success/failure/disabled/redaction/localization/duplicate cases without real external email.
- After all required validation passes, advance private candidate metadata to `1.2.0-rc.4`, build and independently verify a reproducible RC.4 archive with Task 067 canonical modes, metadata consistency, repeated-build byte/SHA-256 identity, and documented provenance.
- Update only relevant operator, configuration, security/privacy, installation, update/rollback, release, limitations, and Task 068 lifecycle documentation; perform task-scoped Git integration and canonical lifecycle closure.


## Out of Scope

- Do not implement continuous filesystem integrity or arbitrary external tamper monitoring.
- Do not redesign unrelated notification, configuration, installer, updater, rollback, Guardian, service, localization, or security subsystems.
- Do not create a second transport/configuration subsystem or embed large duplicated localization logic in scripts.
- Do not require or perform real external email delivery in automated tests.
- Do not perform the complete OVH or Contabo/Hestia final acceptance matrix and do not modify either VPS merely to complete this task.
- Do not promote RC.3, publish RC.4 or final 1.2.0, create a final `v1.2.0` tag or public Forgejo Release, claim final acceptance, or bypass later acceptance gates.
- Do not store or expose SMTP passwords, API keys, authorization tokens, private keys, credentials, or confidential configuration values in logs, history, Git, release artifacts, notifications, or diagnostics.


## Authority Envelope

**Task targets and boundaries:** Framework 2.0 Standard Execution Authority applies to existing QWSG notification/SMTP abstractions, configuration model and authorized configuration commands, installer/update/rollback/Guardian lifecycle integration points, service/report evidence, EN/HU/DE localization catalogs, deterministic tests/mocks, private RC.4 version and artifact metadata, narrowly relevant documentation, Task 068 records, snapshot/rollback evidence, and task-scoped Git integration. Work is limited to QWSG-managed lifecycle/configuration notifications and a private reproducible RC.4 candidate.

**Permitted external actions:** Fetch and compare the canonical `origin`, then perform reviewed task-scoped push dry-run and clean fast-forward push. Local protected snapshot, isolated build/test environments, deterministic mock transports, and private artifact construction/inspection are permitted. No VPS mutation or external email delivery is authorized or required.

**Owner-reserved decisions:** Final `v1.2.0` tagging/publication, public RC.4 publication, Forgejo Release/assets, acceptance waivers, external VPS mutation, production notification delivery, product/architecture expansion beyond the existing notification system, destructive external actions, force push, history rewriting, and identity changes beyond `1.2.0-rc.4` remain reserved to Project Owner Attila.

**Task-specific STOP conditions:** Stop for a material baseline/identity discrepancy, unavailable rollback, secret/privacy/security exposure, an integration requiring a parallel or unsafe notification subsystem, an unresolved semantic conflict between operation and delivery results, unsafe duplicate/loop behavior, inability to satisfy required EN/HU/DE/redaction/event coverage or deterministic packaging, an external/credential boundary, a release-blocking defect, or a canonical lifecycle requirement explicitly requiring Owner action. If safe architectural integration cannot be proven, report BLOCKED and do not declare RC.4 valid.


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

- Verify repository root, Framework 2.0, branch `main`, clean full working tree, canonical HTTPS `origin`, fetched `HEAD == origin/main == 1f876d130d3ab4f204b8c1758411ef1f95889991`, `0 0` divergence, and canonical idle lifecycle after completed Task 067.
- Verify private RC.3 artifact `qwsg-1.2.0-rc.3-linux-amd64.tar.gz`, SHA-256 `8543c3e09b48085b01c037d7db5106ea793374dc099b0b9be5f6cacb55af13ee`, source commit `b6eb357ad03a02b41ac93536fc3be91ecf929803`, and Task 067 deterministic packaging evidence where locally available.
- Verify no final `v1.2.0` tag or publication action occurred and classify any material difference before implementation.


## Snapshot Requirements

- Before task-target changes, create a protected mode-`0700` local snapshot outside Git following `ai/core/15_ENGINEERING_BACKUP_POLICY.md`.
- Capture the complete tracked baseline plus exact Builder-created Task 068 prompt/history before-images; record UTC time, purpose, commit/remote/lifecycle state, bounded scope, exclusions, retention through Owner acceptance and release rollback-window closure, deterministic manifest, SHA-256 checksums, archive readability, collision-safe exact restore instructions, and post-restore validations.
- Verify checksums and rehearse restoration only into an isolated protected directory; never extract over the live worktree or commit payload/private host evidence.


## Risk Assessment

- High risk: lifecycle hooks and notifications touch update/rollback truth, sensitive configuration, localization, delivery reliability, duplicate behavior, and release-candidate provenance.
- Controls: architecture-first inspection, reuse of existing boundaries, default-safe configuration, strict secret redaction, deterministic mock transport, separate operation/delivery outcomes, idempotent event identity where supported, bounded integration points, protected rollback, focused and full/race tests, reproducible archive inspection, targeted staging, and prohibition on external delivery/publication.
- Notification failure must be visible but cannot corrupt or rewrite an otherwise successful lifecycle result; underlying failures remain visible even when notification delivery fails.


## Planned Work

1. Read all governing and relevant prior/release records; verify and record exact baseline and remote state.
2. Create, verify, and safely rehearse the mandatory rollback snapshot.
3. Inspect notification/SMTP/configuration/installer/update/rollback/Guardian/service/localization/security implementations and tests; document the narrowest existing extension seam.
4. Define event/result/privacy/localization/idempotency contracts and supported integration points without changing unrelated architecture.
5. Implement reusable lifecycle notification composition/delivery and configuration behavior, then integrate installation, update, rollback, version transition, authorized configuration change, and cleanly supportable Guardian transitions.
6. Add deterministic tests for every mandated event and outcome, disabled mode, transport failure, secret redaction, EN/HU/DE, version direction, and duplicate safety; diagnose/correct/retest in scope.
7. Run focused and repository-wide/full/race/security tests. Only after they pass, update canonical private metadata to RC.4.
8. Build private RC.4 from a frozen source commit, independently audit contents/modes/ownership/metadata/provenance, rebuild under a materially different controlled condition, and require byte-identical SHA-256.
9. Update narrow documentation and Task 068 history with exact feature, security, delivery, artifact, test, limitation and Git evidence.
10. Review exact paths/diffs/modes/privacy, use targeted staging, commit, dry-run and clean-fast-forward push, verify remote identity, then canonically archive the completed task to idle.


## Rollback Plan

- Before integration, verify the protected snapshot and restore only individually reviewed Task 068 target files from isolated extraction; remove only explicitly identified Task 068-generated private build outputs when required; then rerun focused notification/release tests, Git status, Framework, and lifecycle checks.
- After commits or push, use forward corrective commits or exact snapshot-backed path restoration under the same bounded authority; never broad reset/checkout/clean, history rewrite, tag mutation, or unrelated-state disturbance.
- Notification tests use local deterministic transports and require no external rollback. If exact repository/configuration restoration cannot be proven or sensitive/external state would be affected, stop for Owner direction.


## Deliverables

- Clean integration into existing notification/configuration/localization architecture for required QWSG-managed lifecycle and configuration events.
- Explicit operation-result versus delivery-result contract, safe secret handling, non-looping duplicate behavior, EN/HU/DE content, supported configuration and deterministic test-notification behavior.
- Automated deterministic tests covering mandated success/failure/update/rollback/version/configuration/disabled/transport/redaction/localization/idempotency cases.
- Updated task-scoped operator/security/release documentation recording RC.3 supersession and RC.4 behavior.
- Private reproducible `qwsg-1.2.0-rc.4-linux-amd64.tar.gz`, exact SHA-256 and provenance, canonical-mode/ownership inspection, repeat-build identity, completed history, Git integration, and idle lifecycle.


## Verification

- Validate Builder installation, active prompt/history identity, Framework and lifecycle before implementation and at completion.
- Run focused notification/configuration/installer/update/rollback/Guardian/localization tests and deterministic mock-transport tests for all required events, success/failure semantics, disabled delivery, transport failure, secret redaction, EN/HU/DE, old/new/restored version reporting, and duplicate behavior.
- Prove operation results remain truthful and independently visible when notification delivery fails, with no loops or duplicate spam.
- Run shell syntax, formatting, vet, full Go tests, race tests, build/release contracts, engineering/framework suites, privacy/security review, and every configured repository-wide validation.
- For RC.4 verify version/release metadata consistency, embedded full source commit/build time, manifest, safe paths/types, canonical `0755` directories, exactly intended executable `0755` files, other files `0644`, numeric `0/0` ownership, no writable/privileged surprises, and extraction permissions.
- Repeat candidate build from logically identical content under a materially different mode/umask condition; require `cmp` and SHA-256 identity.
- Review unstaged/staged path lists, executable bits, secrets/privacy, generated/ignored artifacts and `git diff --cached --check`; stage only explicit task paths.
- Verify final clean branch, `HEAD == origin/main`, `0 0` divergence, idle Task 068 lifecycle, preserved private artifact hash, and absence of tags/publication/VPS changes.


## Documentation Updates

- Document administrator change-notification purpose, supported events, enable/disable and recipient configuration, validation/test usage, delivery-failure versus operation-result semantics, privacy/redaction, EN/HU/DE localization, update and rollback version direction, duplicate behavior, and known limitations.
- Update 1.2.0 release documentation to state that private RC.3 is superseded by RC.4 because the Owner required lifecycle/change notifications; record exact RC.4 artifact SHA-256 and provenance after freezing.
- Maintain Task 068 history throughout with starting/snapshot/architecture decisions, attempts/classifications, exact verification, artifact/Git/lifecycle results, and limitations. Do not record secrets or unrelated architecture changes.


## Completion Criteria

Task 068 is COMPLETE only when the existing notification architecture is cleanly extended; required QWSG-managed events and old/new/restored version semantics are covered; operation and delivery failures remain separately truthful; secrets/redaction and duplicate safety pass; EN/HU/DE behavior passes; configuration/test behavior is supported; all focused, automated, full/race and repository-wide tests pass; Task 067 deterministic packaging remains intact; a private reproducible RC.4 is created with recorded SHA-256/provenance and safe canonical metadata; documentation records RC.3 supersession and limitations; Git integration passes; and lifecycle closes canonically. If safe in-architecture integration or required proof fails, report BLOCKED and do not declare RC.4 ready.


## Owner Approval Requirements

Approved by Project Owner: Attila through the Engineering Task Builder on 2026-08-28 UTC.

The structured task definition and Authority Envelope have been explicitly approved. Framework 2.0 Standard Execution Authority permits iterative, reversible in-scope engineering without another Owner gate. Further scope changes, exceptional external actions, and Owner-reserved decisions require explicit Project Owner approval.
