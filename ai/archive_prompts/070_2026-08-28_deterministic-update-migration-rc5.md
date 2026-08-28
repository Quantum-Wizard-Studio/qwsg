# Current Engineering Task 070: Deterministic RC.2 to RC.5 Update Migration & Release Candidate

## Task Metadata

- Task ID: `070`
- Task slug: `deterministic-update-migration-rc5`
- Status: `complete`
- Date opened: `2026-08-28` UTC
- Human authority: Project Owner: Attila
- Owner or lead-developer communication language: English

## Title

Deterministic RC.2 to RC.5 Update Migration & Release Candidate


## Objective

Remediate release blocker QWSG-069-F001 with the smallest safe reusable deterministic compatibility/migration mechanism supporting the real QWSG 1.2.0-rc.2 to 1.2.0-rc.5 update, prove preservation, validation, rollback and lifecycle-notification behavior through automated regression fixtures, and produce a private reproducible QWSG 1.2.0-rc.5 candidate for subsequent final acceptance. This task is not final acceptance and does not authorize public release.


## Scope

- Establish the exact Task 069 completed baseline, mandatory protected snapshot and verified isolated rollback before task-target changes.
- Define and document a canonical fail-closed compatibility/migration contract covering source/target identification, supported-source decision, deterministic path selection, preflight, configuration/state and protected-credential preservation or explicit transformation, package-replacement boundary, post-update configuration/Guardian/systemd validation, rollback metadata, failure handling and Task 068 lifecycle-event integration.
- Implement a reusable but narrowly scoped deterministic QWSG 1.2.0-rc.2 -> 1.2.0-rc.5 path; treat RC.3/RC.4 only if existing architecture genuinely requires it and never use a brittle version bypass.
- Create or derive a credential-free deterministic RC.2 installed-state fixture matching the compatibility-relevant real Task 069 predecessor characteristics without live-host or private data dependencies.
- Prove actual file/state results for successful migration and RC.5 -> restored RC.2 rollback, including byte identity where deterministic, configuration/state/credential protection, unit/Guardian semantics, resource controls, paths, failure safety, idempotency/repeated invocation and no mutation for unsupported sources.
- Reuse Task 068 notification composition/dispatch abstractions for update and rollback SUCCESS/FAILED events, version direction, operation identity/action requirements, transport-result separation, redaction and deterministic local/mock transport tests; localize user-facing migration/update text in EN/HU/DE where applicable.
- Run all focused and existing relevant regression, race, security, framework and repository suites; preserve Task 067 deterministic packaging guarantees.
- Advance private candidate identity to 1.2.0-rc.5 only after behavioral gates pass; build, inspect and independently reproduce its archive across repeated builds, umasks and source-mode variations, recording artifact, SHA-256 and source provenance.
- Update relevant canonical updater/operator/security/release/limitations documentation and docs/release/ACCEPTANCE_1.2.0.md without erasing Task 069 failure evidence; perform task-scoped Git integration and canonical Task 070 closure.


## Out of Scope

- No final QWSG 1.2.0 publication, final v1.2.0 tag, public candidate, Forgejo Release or acceptance waiver.
- No mutation of OVH or Contabo; Task 070 uses local fixtures/mocks only and real VPS acceptance belongs to the subsequent final acceptance task.
- No real SMTP/mailbox delivery, QUWIP or Telegram transport.
- No unrelated updater redesign, enterprise migration framework, invented migration paths, arbitrary filesystem tamper monitoring, unrelated product features, Hestia changes or infrastructure cleanup.
- No secrets, private-host data or credentials in fixtures, repository, logs, tests, history, notifications, diagnostics or artifacts.


## Authority Envelope

**Task targets and boundaries:** Framework 2.0 Standard Execution Authority applies to the existing QWSG updater/rollback/version/trust/transaction components, configuration and QWSG-managed state contracts, installed Guardian/user-unit lifecycle, Task 068 notification composition/dispatch and EN/HU/DE localization, deterministic credential-free RC.2 fixtures/tests, private RC.5 version/release metadata and artifacts, narrowly relevant documentation, Task 070 lifecycle/snapshot evidence and task-scoped Git integration. Work is limited to remediating QWSG-069-F001 with a reusable deterministic RC.2 -> RC.5 path and private release candidate.

**Permitted external actions:** Fetch and compare canonical origin, then perform reviewed task-scoped push dry-run and clean fast-forward push. Protected local snapshots, isolated local test/install roots, deterministic mock transports and private artifact construction/inspection are permitted. OVH, Contabo, real SMTP/mailbox delivery and public release systems must not be mutated.

**Owner-reserved decisions:** Final v1.2.0 tagging/publication, public RC.5 publication, Forgejo Release/assets, final acceptance, acceptance waivers, external VPS or infrastructure mutation, production notification delivery, new transports, material architecture/scope expansion, destructive external actions, force push, history rewriting and changes outside this remediation remain reserved to Project Owner Attila.

**Task-specific STOP conditions:** Stop for a material mismatch from HEAD/origin a28f29b69712a919b34ecee6e6cd66f504d755d2 or Task 069's clean idle BLOCKED baseline; unavailable deterministic rollback; secret/privacy/security exposure; need to mutate OVH/Contabo; unsafe source/version guessing or package replacement before compatibility/preflight; architectural conflict requiring a brittle bypass or material redesign; inability to prove preservation, Guardian/systemd safety, rollback, notification truth or deterministic packaging; release-blocking defect; or an explicit Owner/lifecycle boundary. Ordinary in-scope failures enter diagnose, correct and retest.


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

- Verify exact QWSG root, ordinary non-root execution, Framework 2.0.0, branch main, clean full worktree/index, canonical HTTPS origin, fresh fetch, HEAD == origin/main == a28f29b69712a919b34ecee6e6cd66f504d755d2, 0/0 divergence and canonical idle lifecycle after Task 069 archived BLOCKED.
- Verify Task 069 finding QWSG-069-F001 and preserved failure evidence: actual installed predecessor 1.2.0-rc.2 was refused by RC.4 before replacement with `Update refused: no deterministic compatible migration path`; RC.4 was not promoted and no final v1.2.0 tag/release exists.
- Verify RC.4 reference artifact identity where locally available: qwsg-1.2.0-rc.4-linux-amd64.tar.gz SHA-256 adeb591605c0d37a5fc98d541125ca388cd4561703d0f0823bba931bc7d08684. Treat missing disposable local artifact as non-material if authoritative Task 069 evidence remains intact.
- Confirm OVH is intentionally QWSG-free and Contabo restored to RC.2 only from existing evidence; do not contact or mutate either host.


## Snapshot Requirements

- Before modifying task targets, create a protected mode-0700 snapshot outside Git per ai/core/15_ENGINEERING_BACKUP_POLICY.md with mode-0600 payload/evidence.
- Capture complete tracked baseline plus exact Builder-created Task 070 prompt/history before-images, commit/remote/lifecycle state, bounded scope/exclusions/retention, deterministic manifest and SHA-256 checksums, collision-safe exact restore instructions and post-restore validations.
- Verify checksums/archive readability and rehearse restoration only into an empty isolated protected directory. Do not extract over the live worktree, store secrets/private host evidence, or touch Task 069 external-host snapshots.


## Risk Assessment

- High risk: update transactions can replace executable/unit state and preserve sensitive configuration; incorrect compatibility selection or rollback can strand an installation, weaken Guardian boundaries or expose credentials.
- Controls: architecture-first inspection, explicit declarative compatibility path, fail-closed unknown/malformed state, preflight before mutation, bounded transaction boundary, protected rollback metadata, exact-result assertions, credential-free fixtures/redaction tests, local mock notifications, resource/path validation, targeted staging and complete regression/race/reproducibility checks.
- No supported path may be guessed. Operation result and notification transport result remain independently truthful. Package replacement occurs only after deterministic compatibility and preflight success.


## Planned Work

1. Read all required governance and relevant Tasks 064, 067, 068 and 069 evidence plus updater/rollback/release documentation; verify and record the exact baseline.
2. Create, verify and rehearse the mandatory protected snapshot and exact bounded rollback path.
3. Inspect installed-state detection, migration planning, transaction/preflight, configuration/state preservation, package replacement, post-validation, Guardian/systemd, rollback metadata and notification seams; define the minimal canonical migration contract.
4. Build a deterministic credential-free RC.2 installed-state fixture representing the Task 069 predecessor and add explicit declarative RC.2 -> RC.5 compatibility/path selection without RC.2-special-case bypass logic.
5. Implement preflight, preservation/transformation (documenting no schema transformation if none is needed), package boundary, post-update validation and protected deterministic rollback metadata/restoration using existing architecture.
6. Integrate update/rollback SUCCESS/FAILED lifecycle events through Task 068 abstractions with version direction, operation/action semantics, transport separation, redaction, EN/HU/DE where applicable and deterministic duplicate behavior.
7. Add the complete required automated matrix: source detection; path decision/selection; unknown/malformed fail-closed/no-mutation; preflight boundary; successful migration/resulting RC.5; config/credential/state/unit/Guardian preservation; post-validation; metadata; actual RC.2 restore/file verification; update/rollback failure; all four event outcomes; version reporting; transport separation; redaction/localization; idempotency/repeated update.
8. Run focused, full, race, formatting/vet, security/privacy, release and framework validations; diagnose/correct/retest all recoverable in-scope failures.
9. Update private identity/docs to RC.5 after gates pass, freeze the candidate source commit, build and inspect archive integrity/manifest/provenance/modes/numeric 0/0 ownership/extraction, and require byte/SHA identity across repeated, cross-umask and source-mode builds.
10. Record exact evidence and limitations, review diffs/modes/privacy, stage explicit task paths, commit, push dry-run and clean fast-forward push, verify remote equality, then close Task 070 canonically as READY FOR FINAL ACCEPTANCE or BLOCKED.


## Rollback Plan

- Before integration, verify the protected snapshot and restore only individually reviewed Task 070 target files from isolated extraction; remove only explicitly identified Task 070-generated private artifacts when necessary; rerun focused migration/rollback/notification/release tests, Git status, Framework and lifecycle validation.
- Runtime regression tests operate only in isolated temporary roots and must prove their own RC.5 -> exact RC.2 rollback including actual version/files/state; cleanup is bounded to explicit test roots.
- After commit/push, use forward corrective commits or exact snapshot-backed path restoration under the same bounded authority; never broad reset/checkout/clean, history rewrite, tag mutation or unrelated-state disturbance. If exact safe restoration cannot be proven, stop and report BLOCKED.


## Deliverables

- Documented canonical deterministic compatibility/migration contract and reusable fail-closed implementation with supported QWSG 1.2.0-rc.2 -> 1.2.0-rc.5 path.
- Credential-free deterministic RC.2 installed-state fixture and comprehensive regression tests proving compatibility, no unsupported mutation, preservation, Guardian/systemd semantics, failures, rollback and lifecycle notifications.
- Actual successful RC.2 -> RC.5 -> rollback -> RC.2 state proof, protected rollback metadata, credential/redaction/security evidence and EN/HU/DE user-facing behavior where applicable.
- Updated updater/operator/security/release/acceptance/limitations documentation preserving Task 069 failure evidence.
- Private reproducible qwsg-1.2.0-rc.5-linux-amd64.tar.gz with exact SHA-256, source commit/provenance, canonical metadata, completed history, clean Git integration and idle lifecycle.


## Verification

- Validate Builder installation, active prompt/history identity, Framework/lifecycle and exact Git baseline before implementation and at completion.
- Test all thirty Owner-mandated migration cases, inspecting resulting file/state/version/unit/Guardian state rather than exit codes alone and comparing deterministic preserved/restored data byte-for-byte where possible.
- Prove protected credentials are preserved but never emitted; unsupported/malformed/preflight failures cause no package mutation; rollback failures are visible/safe; operation truth remains separate from notification delivery; update/rollback SUCCESS/FAILED events carry correct old/new/restored versions.
- Prove expected user-unit definition/state, config validation, Guardian readiness, enabled/active semantics, resource controls, one service/process, current executable path and absence of unexpected system-level paths in isolated acceptance tests.
- Run relevant focused Go/shell tests, go test ./..., go test -race ./..., formatting, vet, build/release checks, engineering/framework/lifecycle suites and all configured repository validation.
- For RC.5 require repeated-build byte identity, cross-umask and source-mode reproducibility, safe archive paths/types, canonical directory/file permissions and intended executables, numeric 0/0 ownership, integrity, manifest/provenance/version consistency, SHA-256 identity and correct extraction behavior.
- Review unstaged/staged paths, modes, generated/ignored outputs and privacy; run git diff --cached --check; verify final clean main, HEAD == origin/main, 0/0 divergence, no final tag/publication or VPS mutation, preserved private artifact hash and idle Task 070 lifecycle.


## Documentation Updates

- Document deterministic compatibility decisions, supported sources/path, explicit RC.2 -> RC.5 migration, fail-closed unknown state, configuration/state and credential preservation or schema transformation, pre/post validation, Guardian/systemd expectations, rollback metadata/restoration, notification semantics and known limitations.
- Update docs/release/ACCEPTANCE_1.2.0.md to retain that Task 069 was BLOCKED by QWSG-069-F001 and RC.4 was not promoted, record Task 070 remediation and RC.5 supersession, and state that real final acceptance remains pending.
- Record RC.5 artifact filename, SHA-256, reproducibility and source commit only after freezing; maintain Task 070 history with baseline, snapshot, architecture, attempts/classifications, exact tests, security, rollback, artifact, Git/lifecycle evidence and final READY FOR FINAL ACCEPTANCE or BLOCKED classification.


## Completion Criteria

Task 070 is COMPLETE only when QWSG-069-F001 is remediated by a reusable deterministic RC.2 -> RC.5 migration; unsupported/unknown/malformed sources still fail closed without mutation; valid configuration, protected credentials, relevant state and Guardian/user-unit safety are proven; deterministic rollback restores actual RC.2 version/files/state; notification integration, transport separation, redaction, localization and applicable idempotency pass; a credential-free RC.2 fixture prevents regression; all focused/full/race/framework/release tests pass; Task 067 packaging remains deterministic; private RC.5 is produced with recorded SHA-256/source provenance; documentation preserves Task 069 evidence; Git is clean/equal and lifecycle closes canonically. Final classification must explicitly be READY FOR FINAL ACCEPTANCE or BLOCKED. No final v1.2.0 tag or public release may be created.


## Owner Approval Requirements

Approved by Project Owner: Attila through the Engineering Task Builder on 2026-08-28 UTC.

The structured task definition and Authority Envelope have been explicitly approved. Framework 2.0 Standard Execution Authority permits iterative, reversible in-scope engineering without another Owner gate. Further scope changes, exceptional external actions, and Owner-reserved decisions require explicit Project Owner approval.
