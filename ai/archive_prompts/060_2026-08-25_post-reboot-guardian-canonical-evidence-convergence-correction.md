# Current Engineering Task 060: Post-Reboot Guardian Canonical Evidence Convergence Correction

## Task Metadata

- Task ID: `060`
- Task slug: `post-reboot-guardian-canonical-evidence-convergence-correction`
- Status: `complete`
- Date opened: `2026-08-25` UTC
- Human authority: Project Owner
- Owner or lead-developer communication language: English

## Title

Post-Reboot Guardian Canonical Evidence Convergence Correction


## Objective

Identify the exact root cause of `QWSG-059-F001` and implement the smallest safe correction so a Guardian invocation that fails transiently during boot and is automatically recovered by systemd converges to fresh, complete canonical evidence and overall readiness. Preserve all working RC.5 behavior already externally proven, advance source/release identity to the expected replacement candidate `QWSG 1.1.0-rc.6`, and establish deterministic local regression evidence without constructing candidate artifacts or accessing the VPS.


## Scope

- Diagnose `QWSG-059-F001` from immutable privacy-safe Task 059 evidence and the actual Guardian/runtime-service/systemd/canonical-state implementation. Reproduce the failure in a bounded local controlled environment before correcting it where technically feasible.
- Inspect and modify only the smallest necessary product components governing Guardian invocation generation, systemd restart/exit reporting, checkpoint ownership, runtime-service lifecycle publication, current-operator-state convergence, readiness classification, and directly coupled tests.
- Ensure a failed boot-time invocation cannot leave or reapply degraded exit evidence over a later recovered invocation's successful fresh cycle, and that successful recovery converges deterministically to `guardian=running`, current/complete canonical evidence, and overall readiness when all other requirements are satisfied.
- Preserve exact generation/invocation isolation, stale-exit rejection, checkpoint integrity, atomic state publication, private state-directory safety, and truthful failure evidence. Do not hide a genuinely failed active invocation or manufacture ready evidence.
- Add focused regression tests for the actual Task 059 sequence: pre-login start, transient invocation failure, systemd-style replacement generation, successful recovered cycle, delayed/late old-exit reporting where relevant, fresh canonical convergence, and resistance to stale-generation corruption.
- Preserve clean-host installation, guided setup/activation, safe real state-directory handling, no compatibility symlink/migration, pre-login Guardian autostart, systemd automatic recovery, SMTP notification, and protected credential handling already externally proven in RC.5.
- Advance version-coupled source identity, release notes, documentation and release-plumbing assertions from `1.1.0-rc.5` to expected `1.1.0-rc.6` only where required by changed product bytes. Record RC.5 as failed practical acceptance and preserve all historical identities/evidence immutably.
- Run focused and complete local validation, perform task-scoped Git integration and clean fast-forward push, and canonically close Task 060 under its Authority Envelope.


## Out of Scope

- Do not access, inspect, restart, reconfigure, uninstall, reinstall, reset, or otherwise modify the disposable acceptance VPS.
- Do not modify, rebuild, repackage, relabel, transfer, or reuse the accepted RC.5 candidate bytes. RC.5 remains immutable failed acceptance evidence.
- Do not construct an RC.6 release archive, checksum sidecar, release candidate, deterministic twin, or externally transferable artifact. Ordinary bounded local test builds are permitted only as validation outputs and are not release candidates.
- Do not perform external clean-host/reboot/notification/uninstall/reinstall acceptance; replacement-candidate construction and acceptance require a later separately authorized task.
- Do not create or move a tag, create a Forgejo Release, upload/publish assets, activate a public download URL, deploy, announce, or claim QWSG 1.1.0 is released.
- Do not weaken canonical evidence integrity/freshness, generation isolation, systemd recovery, state permissions, credential privacy, notification security, snapshot/rollback, historical immutability, or release gates to make tests pass.
- Do not introduce unrelated architecture, features, providers, service topology, dependency installation, or broad refactoring.
- Do not read, hash, copy, modify, stage, package, or otherwise access Owner-owned `docs/architecture/QWCS_MIGRATION_BLUEPRINT.md` beyond pathname metadata exclusion checks.
- Do not create or install Task 061.


## Authority Envelope

1. **Authorized paths/components/systems:** Task 060 may modify the minimum necessary Guardian, runtime-service, systemd exit/restart handoff, checkpoint/current-state/readiness components and directly coupled tests; version/release-note/release-plumbing files required for `1.1.0-rc.6`; Task 060 prompt/history; and narrowly necessary architecture/operations documentation. Work is confined to this repository and private local rollback/test storage. No VPS is authorized.
2. **Routine operations:** inspect, analyze, snapshot, reproduce locally, edit, format, build ordinary test binaries, run focused/integration/full/race/vet/security/release-plumbing/governance tests, diagnose, correct, retest, document, privacy-review, explicitly stage reviewed Task 060 paths, review staged diffs/modes, commit, push dry-run, clean-fast-forward push, verify refs, and report without intermediate Owner gates.
3. **Correction/retest authority:** Recoverable in-scope implementation, fixture, test, documentation, and validation failures follow diagnose -> smallest safe correction -> retest -> continue. Up to the controlled-failure policy limit, vary the diagnostic method when evidence rejects an approach. Do not weaken evidence semantics, bypass a failing product boundary, or change external/historical evidence to manufacture PASS.
4. **Repository integration:** After required validation, explicit path staging, task-scoped commits, push dry-run, clean fast-forward push to `origin/main`, and direct read-only Forgejo ref verification are authorized. Broad staging, history rewrite, force push, tag operations, Release operations, publication, and unrelated paths are forbidden.
5. **Lifecycle completion:** Task 060 may truthfully finalize its prompt/history, integrate the correction and evidence, archive its prompt, push lifecycle-only closure, validate canonical idle, and report completion without another routine Owner gate. It may not create or install Task 061.
6. **Permitted external actions:** Read-only access to the canonical Forgejo repository/ref and official technical documentation is permitted when needed. No VPS, SMTP provider, credential, candidate transfer, external acceptance, tag, Release, upload, publication, deployment, or announcement action is permitted.
7. **Evidence and rollback:** Before changes, preserve a private mode-0700 snapshot with tracked HEAD, exact target before-images/absence claims, Task 059/RC.5 preservation hashes, Git/ref/mode/ACL evidence, and bounded rollback. Retain the failing regression, corrected regression, generation/isolation behavior, exact diffs/modes, validation results, commit/ref identities, privacy/security scans, and chronological history. Missing evidence is never PASS.
8. **Owner-reserved operations:** Material scope or architecture expansion; dependency installation; unplanned destructive/irreversible work; unresolved security/privacy decisions; privilege escalation; VPS or other infrastructure mutation; credentials; candidate construction/transfer; external acceptance; tags; Forgejo Releases; asset upload/publication; deployment; announcement; QWSG 1.1.0 release authority; and Task 061.
9. **Mandatory STOP conditions:** Stop for unavailable reliable rollback; inability to reproduce or technically ground the defect before a semantic correction; a required change outside the bounded Guardian/evidence/recovery scope; meaningful security/privacy or canonical-integrity uncertainty; unplanned destructive/external mutation; a need for credentials/elevated authority/Owner-reserved work; or meaningful risk outside Task 060. Ambiguity never expands authority. Routine in-scope failures remain diagnose/correct/retest events.


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

- Verify UTC date, ordinary user, exact repository root, branch `main`, canonical HTTPS origin, `HEAD == origin/main ==` direct Forgejo `main` at `07b744e604e8a3b8e7ef1dc80147399c1af1e06b`, ahead/behind `0/0`, empty index, clean tracked tree, zero active prompts, and canonical idle with Task 059 complete and archived.
- Verify the only unrelated state is the excluded Owner blueprint by pathname metadata only. Do not access its contents.
- Verify Framework `1.1.0`, `VERSION=1.1.0-rc.5`, Task 059 Practical Acceptance FAIL, Steps 1–8 PASS, Step 9 FAIL, Steps 10–12 NOT EXECUTED, and exact `QWSG-059-F001` privacy-safe evidence.
- Verify Task 057 and earlier candidate/protocol/history evidence remains immutable; RC.5 candidate identities remain exactly preserved; no RC.6 bytes, tag, Release, upload, publication, VPS mutation, or Task 060 artifacts exist.
- Read the Guardian, runtime-service, systemd unit, checkpoint/current-state, readiness, setup/state, notification, installer/uninstaller, release-plumbing code/tests and documentation necessary to trace the Task 059 sequence.
- Run Framework, Builder, lifecycle, diversion/test-task, release-plumbing, build, focused/full/race/vet/format, systemd/static, security/privacy, historical-preservation, and Git baseline validation before modification.


## Snapshot Requirements

Before modifying task targets, create and verify a unique private mode-0700 snapshot under `/tmp` containing a readable exact tracked-HEAD archive; literal copies and absence records for every proposed target; Git/ref/status evidence; modes/ACLs where relevant; Builder input/prompt identities; immutable Task 059/RC.5 and earlier protected-evidence hashes; and exact bounded rollback instructions. Exclude Owner content, credentials, private infrastructure identity, candidate bytes, caches, build outputs, and unrelated files. Verify readability, hashes, modes, absence claims, collision safety, exclusions, and restore instructions before implementation. Create an additional bounded snapshot before lifecycle closure if required for reliable rollback.


## Risk Assessment

- False root cause — critical: reproduce and correlate generation, checkpoint, exit report, recovered cycle and canonical state before changing semantics.
- False readiness — critical: a correction must restore running/complete/current evidence only after a valid recovered cycle; genuine current failure remains degraded/not-ready.
- Stale-generation overwrite — critical: late exit evidence from an old invocation must never mutate the recovered invocation's state; preserve exact generation isolation and test reordered events.
- Recovery regression — critical: retain systemd pre-login autostart and `Restart=on-failure`; do not remove or mask recovery behavior.
- State/security regression — critical: preserve non-symlink private state preparation, ownership/modes, atomic integrity, configuration/state separation, and no compatibility migration.
- Notification/credential regression — high: preserve optional SMTP behavior, protected credential files, TLS/auth boundaries and privacy; no external provider access occurs.
- Candidate identity error — critical: changed product bytes require consistent `1.1.0-rc.6` source identity/plumbing, but no candidate artifact may be constructed.
- Historical corruption — critical: Task 059 and RC.5 evidence remain immutable and additive; never relabel RC.5 PASS.
- Test realism — high: local controlled recovery must model the actual boot/restart ordering without claiming external proof; later clean-host acceptance remains required.
- Owner-content/privacy exposure — critical: exclude blueprint content, private host/provider/account identity, credentials, acceptance secrets, caches and artifacts from snapshots/Git.


## Planned Work

1. Validate canonical idle, Git/ref/version/protected evidence and Framework; read Task 059 plus relevant Guardian/runtime/systemd/state/readiness code and tests; create/verify the Task 060 snapshot.
2. Build a precise event/generation timeline for `QWSG-059-F001`: pre-login invocation, transient failure, systemd replacement, checkpoint generation, exit reporting, recovered cycle publication, and persistent degraded canonical state.
3. Add the smallest controlled failing regression that reproduces the demonstrated convergence failure, including late/reordered old-generation exit evidence and successful new-generation cycle where evidence supports that boundary. Reject speculative changes that do not reproduce the defect.
4. Identify the exact ownership/race/transition defect and implement the smallest safe correction in the canonical Guardian evidence boundary. Preserve truthful degraded evidence for the active generation and ignore only demonstrably stale superseded-generation mutations.
5. Add focused positive/negative tests proving recovered-cycle convergence to running/current/complete readiness, stale-generation rejection, genuine-current failure preservation, atomic state/checkpoint integrity, restart behavior, and state/security invariants.
6. Run setup/state/user-service/readiness/notification/installer/uninstaller and systemd static regressions needed to preserve the externally proven RC.5 behavior.
7. Advance source identity and required release notes/plumbing consistently to `1.1.0-rc.6` without constructing candidate bytes or external artifacts. Document that external replacement-candidate acceptance remains pending.
8. Run focused and full/race/vet/format/build/release/governance/security/privacy/preservation validation. Diagnose, correct and retest in-scope failures.
9. Review exact paths/diffs/modes, explicitly stage only Task 060 paths, commit, push dry-run, clean-fast-forward push, verify Forgejo refs, truthfully close Task 060 to canonical idle, and report. Do not access the VPS or create a candidate/tag/Release/publication/Task 061.


## Rollback Plan

- Restore only literal Task 060-owned targets from the verified snapshot after proving current identities and collision conditions. Remove only new files whose prior absence and current Task 060 identity are proven.
- Never use broad reset, checkout, restore, clean, wildcard deletion, history rewrite, force push, tag mutation, Owner-content access, or external/VPS rollback.
- Prefer a bounded forward correction and retest for recoverable in-scope failures. If a task commit is already pushed, use a new corrective commit rather than rewriting published history.
- Version identity rollback must restore the complete consistent pre-Task-060 RC.5 source/plumbing set; never leave mixed RC.5/RC.6 identities or relabel any candidate bytes.
- After rollback, rerun focused/full product tests, race/vet/format, release plumbing, Framework/lifecycle, systemd/static, security/privacy, historical preservation, Git state and exact target validation.


## Deliverables

- A technically grounded root-cause report for `QWSG-059-F001` with an exact generation/event/state timeline.
- The smallest safe product correction that makes a systemd-recovered Guardian converge to fresh running/complete canonical evidence after a successful recovered cycle while preserving genuine failure truth.
- Focused regression coverage for transient boot failure, replacement invocation, stale/late exit evidence, recovered cycle convergence, generation isolation, and negative failure cases.
- Preserved clean-host installation, guided setup, state-directory safety, pre-login autostart, automatic recovery, SMTP notification and protected credential behavior through local regression evidence.
- Consistent expected replacement source identity `QWSG 1.1.0-rc.6`, release notes and release-plumbing validation, with no candidate construction or publication.
- Complete Task 060 documentation/history, verified rollback, validated task-scoped commits/push/ref evidence, and canonical lifecycle closure.


## Verification

- Exact Task 059 failure reproduction or a technically equivalent deterministic model grounded in its event/generation evidence before semantic correction.
- Tests prove old/superseded invocation exit evidence cannot degrade a newer active generation and cannot overwrite a newer successful canonical publication.
- Tests prove a recovered invocation's successful cycle publishes matching current, complete, `guardian=running` evidence and readiness becomes ready when all other requirements are satisfied.
- Tests prove a genuinely failing current invocation remains degraded/not-ready and missing/invalid/stale evidence never becomes PASS.
- Checkpoint generation, invocation identity, active state, completed-cycle identity, report-exit idempotence, atomic current-state integrity and reordered-event behavior pass.
- Packaged systemd unit retains pre-login enablement and bounded `Restart=on-failure`; state root remains real non-symlink, current-user-owned mode-0700 with safe components and configuration/state separation.
- Guided setup/activation, user-runtime context, readiness, notification/preflight/credential storage, installer/uninstaller, restart and exit-report focused regressions pass.
- Version/source/release-note/plumbing identities consistently report `1.1.0-rc.6`; no RC.6 archive, sidecar, tag, Release or external artifact exists.
- Full `go test ./...`, `go test -race ./...`, `go vet ./...`, formatting, ordinary build, shell syntax, systemd static verification, release plumbing, Framework, Builder, lifecycle, diversion/test-task, security/privacy/secret, historical-preservation, rollback, Git whitespace/diff/mode/staging/commit/push/ref and canonical-idle checks pass.
- Task 059/RC.5 evidence and excluded Owner content remain unchanged; no VPS access or external notification occurs.


## Documentation Updates

- Maintain Task 060 prompt/history throughout execution and archive canonically on truthful completion.
- Add or update the smallest code-adjacent architecture/operations documentation needed to explain corrected generation ownership and canonical evidence convergence.
- Add `1.1.0-rc.6` release notes and update version/release-plumbing documentation/assertions required by changed bytes.
- Preserve Task 059 practical acceptance report, prompt/history, Task 057 records, RC.5 notes/protocol/ledger and all earlier evidence immutably.
- State explicitly that RC.6 candidate construction and external acceptance remain pending and that QWSG 1.1.0 is not released.


## Completion Criteria

Task 060 is complete when `QWSG-059-F001` has a technically demonstrated root cause; a minimal correction prevents stale/superseded invocation evidence from defeating a successfully recovered current Guardian while preserving truthful current-generation failures; focused recovery/convergence/generation/security regressions and all complete validation pass; source/release identity is consistently prepared as `1.1.0-rc.6`; Task 059/RC.5 and Owner evidence remain immutable; no candidate archive/sidecar, VPS access, credential action, tag, Forgejo Release, upload, publication, deployment, announcement or Task 061 occurs; task-scoped commits are clean-fast-forward pushed and verified; and Task 060 is canonically closed to idle. Completion prepares source for later replacement-candidate construction and acceptance only; it does not prove external correction or release QWSG 1.1.0.


## Owner Approval Requirements

Approved by Project Owner through the Engineering Task Builder on 2026-08-25 UTC.

The structured task definition and Authority Envelope have been explicitly approved. The task is authorized to start and execute every routine operation inside that envelope without another Owner gate. Further scope changes and every Owner-reserved operation require explicit Project Owner approval.
