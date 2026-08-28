# Current Engineering Task 071: Guided Installer Readiness Synchronization & RC.6

## Task Metadata

- Task ID: `071`
- Task slug: `guided-installer-readiness-synchronization-rc6`
- Status: `approved`
- Date opened: `2026-08-28` UTC
- Human authority: Project Owner: Attila
- Owner or lead-developer communication language: English

## Title

Guided Installer Readiness Synchronization & RC.6


## Objective

Remediate the major guided-installer synchronization/control-flow defect manually reproduced during Owner-operated RC.5 OVH pre-release acceptance: after successful Guardian activation, wait deterministically for the newly activated Guardian's first fresh canonical evidence before classifying mandatory readiness, preserve genuine failure semantics, add deterministic clean/reinstall/delayed/timeout/optional regression coverage, and produce private reproducible QWSG 1.2.0-rc.6. This task is not final acceptance and authorizes no public release.


## Scope

- Verify the exact completed Task 070 baseline and manual RC.5 evidence, then create and rehearse a protected rollback snapshot before task-target changes.
- Inspect and narrowly correct guided installer activation -> readiness -> completion synchronization by extracting/reusing the existing setup-path bounded fresh-evidence mechanism; maintain one coherent definition of freshness, wait bound, polling/cancellation, success, timeout and diagnostics.
- Ensure successful mandatory readiness with optional notification missing reaches 97%/100%, renders the accurate summary and exits 0; genuine mandatory timeout/NotReady remains explicit exit 4 without false success.
- Reject stale preserved evidence from the prior Guardian lifecycle; safely support clean install and preserved configuration/state reinstall.
- Add deterministic tests for clean immediate success, delayed fresh evidence, stale rejection, preserved-state reinstall, bounded timeout/cancellation, failure diagnostics, completion rendering/exit status, optional notification Partial, EN/HU/DE, secret safety and existing setup behavior.
- Preserve Task 068 notification result separation and Task 070 deterministic migration/rollback architecture; update the required path to 1.2.0-rc.2 -> 1.2.0-rc.6 through the declared contract and retain fail-closed unsupported behavior.
- Run focused/full/race/vet/format/framework/Builder/lifecycle/release/migration/notification suites and Task 067 deterministic packaging regressions.
- Advance private identity to 1.2.0-rc.6 only after behavior passes; build, independently reproduce and audit the private RC.6 candidate.
- Update canonical installer/operations/release/acceptance documentation without erasing Task 069/070/RC.5 evidence; integrate task-scoped commits/push and close Task 071 canonically.


## Out of Scope

- No final v1.2.0 tag, final/public release, Forgejo Release/assets, acceptance waiver or public candidate.
- No OVH or Contabo access or mutation and no real SMTP/mailbox testing.
- No QUWIP/Telegram/new SMTP transport, new product feature, Hestia/infrastructure change, arbitrary filesystem monitoring, unrelated Guardian/installer redesign or cleanup.
- No arbitrary sleep, fixed long delay, stale-evidence acceptance, ignored NotReady, exit-code masking, disabled readiness checks, unbounded retry or shell bypass.
- Do not fabricate a prior Task 071 record; this remediation is the first canonical Task 071.


## Authority Envelope

**Task targets and boundaries:** Framework 2.0 Standard Execution Authority applies to the guided installer and existing setup fresh-evidence wait seam, operational readiness classification, installer localization/tests, Task 070 migration/rollback fixtures, private RC.6 identity/release plumbing/artifacts, narrowly relevant documentation, Task 071 lifecycle/snapshot evidence and task-scoped Git integration. Work is limited to the RC.5 87% synchronization defect and required RC.6 preservation regressions.

**Permitted external actions:** Fetch/compare canonical origin and perform reviewed task-scoped push dry-run plus clean fast-forward push. Protected local snapshots, isolated deterministic test roots/mocks/clocks and private artifact construction/inspection are permitted. No acceptance VPS, SMTP, public release or other external system mutation is permitted.

**Owner-reserved decisions:** Final acceptance, final v1.2.0 tagging/publication, public RC.6/Forgejo Release, acceptance waivers, OVH/Contabo/infrastructure mutation, production notification delivery, new transports, material architecture/scope expansion, destructive external action, force push and history rewriting remain reserved to Project Owner Attila.

**Task-specific STOP conditions:** Stop for mismatch from clean/equal commit 641be36ae3098042331858ec2cc7b8c5108967b6 and canonical idle-after-070 baseline; unavailable rollback; secret/privacy/security exposure; need to contact OVH/Contabo; architecture requiring a second polling system, arbitrary delay, stale-evidence weakening or material redesign; inability to preserve real mandatory-failure exit 4, Task 070 migration/rollback or deterministic packaging; release blocker; or explicit Owner/lifecycle boundary. Ordinary in-scope failures enter diagnose/correct/retest.


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

- Verify exact QWSG root, ordinary non-root user, Framework 2.0.0, main, clean full worktree/index, canonical HTTPS origin, fresh fetch, HEAD == origin/main == 641be36ae3098042331858ec2cc7b8c5108967b6, 0/0 divergence and canonical idle lifecycle after completed/archived Task 070; verify no pre-existing Task 071 record.
- Verify Task 070 COMPLETE/READY FOR FINAL ACCEPTANCE, RC.5 artifact `dist/qwsg-1.2.0-rc.5-linux-amd64.tar.gz` SHA-256 `ca3123f313c2fb95138fceb0fc2af17c5a8295b28e7b13e60f5d7fd57ed19faf`, source `3f2f7822473d3d97c38c171f006e8b263250321c`, and no final v1.2.0 tag/publication.
- Record Owner manual evidence: exact RC.5 clean install/reboot/uninstall preservation PASS; preserved-state reinstall installed and activated successfully but immediate readiness returned exit 4 at 87%, omitting completion/summary; subsequent mandatory readiness/Guardian PASS with only optional notification Partial and no corruption.
- Treat manual OVH activity as factual evidence only; do not access either VPS.


## Snapshot Requirements

- Before modifying targets, create a protected mode-0700 snapshot outside Git per backup policy with mode-0600 payload/evidence.
- Capture complete tracked baseline and exact Builder-created Task 071 prompt/history before-images; record UTC, purpose, commit/remote/lifecycle, scope/exclusions/retention, deterministic manifest/checksums, collision-safe exact restore and post-restore validation.
- Verify checksums/readability and rehearse extraction only into an empty protected directory; never extract over the live worktree or include secrets/private host evidence.


## Risk Assessment

- High risk: installer completion truth, fresh/stale operational evidence, asynchronous Guardian startup, exit semantics, localization and release provenance.
- Controls: reuse one canonical bounded waiter; dependency-injected deterministic clock/poll/evidence tests; lifecycle-generation freshness contract; explicit timeout/cancellation; no arbitrary sleep correctness; optional-versus-required classification; no weakening of state integrity/resource/security; full race/release/migration/notification regression; targeted staging and protected rollback.
- Installation and notification outcomes remain independently truthful; package/runtime success does not manufacture readiness, and transient startup must not manufacture failure before the allowed bound.


## Planned Work

1. Read governance and relevant Tasks 065/068/069/070 evidence plus installer/setup/readiness/Guardian/release contracts; verify baseline and no fabricated predecessor.
2. Create, verify and rehearse the mandatory snapshot and bounded rollback.
3. Reproduce control flow locally through deterministic seams; compare guided install's immediate report with setup's bounded evidence wait and define the narrow shared contract.
4. Extract/reuse a cancellable bounded fresh-evidence waiter that accepts only canonical integrity/freshness satisfied evidence and provides explicit success/timeout/cancellation outcomes.
5. Integrate it after guided Guardian activation; only then build operational readiness. Preserve optional Partial completion/exit 0 and genuine mandatory timeout/NotReady diagnostic/exit 4/no summary.
6. Add deterministic guided-flow harness/tests covering all thirty-five required cases, especially clean install, preserved stale-state reinstall, immediate/delayed evidence, timeout/cancellation, 97%/100%/summary, exit codes, EN/HU/DE and no false pass.
7. Update declared migration target and fixture from RC.5 to RC.6; prove RC.2 detection/path, preservation, rollback metadata/restoration, unsupported fail-closed and notification integration.
8. Run focused/full/race/vet/format/framework/Builder/lifecycle/release/security/privacy suites; correct/retest in scope.
9. Advance docs/private identity to RC.6, freeze source commit, construct and independently reproduce candidate under different umask/source modes, audit integrity/provenance/permissions/ownership/extraction and record SHA-256.
10. Preserve RC.5 failure evidence, finalize Task history/classification, review/stage explicit paths, commit, dry-run/push clean fast-forward, verify remote equality and canonically archive Task 071 idle.


## Rollback Plan

- Before integration, verify snapshot checksums and restore only individually reviewed Task 071 targets from isolated extraction; remove only explicitly identified Task 071-generated private outputs when needed; rerun focused installer/migration/notification/release tests, Git, Framework and lifecycle checks.
- Tests use isolated temporary roots and deterministic fakes; they must clean only their exact roots and require no host rollback.
- After commit/push use forward corrective commits or exact snapshot-backed path restoration; never broad reset/checkout/restore/clean, history rewrite, tag mutation or unrelated disturbance. Stop if exact safe restoration cannot be proven.


## Deliverables

- One shared canonical bounded fresh-Guardian-evidence wait contract reused by setup and guided installer, with deterministic timeout/cancellation and stale rejection.
- Guided clean and preserved-state reinstall regressions proving delayed success reaches 97%/100%, summary and exit 0; timeout/mandatory failure remains diagnostic exit 4 without false summary; optional notification Partial remains successful.
- EN/HU/DE diagnostics, state/security/resource boundaries and Task 068 notification truth preserved.
- RC.2 -> RC.6 declared migration plus preservation/rollback/fail-closed regression.
- Updated historical acceptance/installer/operations/security/release documentation and private reproducible `qwsg-1.2.0-rc.6-linux-amd64.tar.gz` with exact hash/source provenance.
- Completed Task 071 history, clean pushed repository and canonical idle lifecycle ready for possible future Task 072 final acceptance.


## Verification

- Validate Builder installation, prompt/history/Framework/lifecycle identity and baseline before implementation and completion.
- Test all Owner-required clean/reinstall/immediate/delayed/stale/timeout/cancellation/optional/exit/render/localization/security/setup/migration/rollback/notification cases using deterministic injected boundaries, not wall-clock sleeps.
- Prove fresh canonical evidence belongs to the newly activated lifecycle; stale preserved state alone never passes. Prove wait bound/no infinite loop and useful failure/resume instruction.
- Run focused Go tests, full `go test ./...`, `go test -race ./...`, vet, formatting, build/release, Framework v2, Builder, lifecycle, diversion, migration, notification and all configured validations.
- For RC.6 require repeated byte identity, cross-umask/source-mode reproduction, safe paths/types, canonical 0755 directories/intended executables, other files 0644, numeric 0/0 ownership, archive/manifest/provenance/version integrity, SHA identity and extraction behavior.
- Review staged/unstaged paths, modes, privacy/secrets/generated outputs and `git diff --cached --check`; verify final clean/equal main 0/0, idle Task 071, artifact hash, no final tag/publication or VPS access.


## Documentation Updates

- Update canonical installer, operations, troubleshooting, security/limitations and RC.6 release notes for bounded post-activation synchronization, stale evidence, timeout/exit and optional Partial semantics.
- Update `docs/release/ACCEPTANCE_1.2.0.md` without erasure: Task 069 blocker; Task 070 remediation/RC.5; Owner manual RC.5 clean install/reboot/uninstall preservation PASS; preserved-state reinstall 87% exit-4 synchronization defect despite functioning package/Guardian; RC.5 not promoted; Task 071 remediation; RC.6 supersession and pending final acceptance.
- Record exact artifact/hash/reproducibility/source only after freeze. Maintain Task 071 history with baseline, snapshot, root cause, implementation, test matrix, failures/classifications, security, rollback, artifact, Git/lifecycle evidence and final READY FOR FINAL ACCEPTANCE or BLOCKED.


## Completion Criteria

Task 071 is COMPLETE only when the RC.5 87% premature termination is remediated by deterministic shared fresh-evidence synchronization; stale preserved evidence cannot pass; clean and preserved-state guided paths reach 100%/summary/exit 0 after mandatory readiness; optional notification alone remains Partial and successful; real timeout/NotReady remains bounded diagnostic exit 4 without false success; EN/HU/DE and security pass; Task 070 migration remains intact as explicit RC.2 -> RC.6 with deterministic rollback/fail-closed behavior; all focused/full/race/framework/release tests pass; private reproducible RC.6 is produced with hash/provenance; historical evidence is preserved; Git is clean/equal and lifecycle closes idle. Classify RC.6 READY FOR FINAL ACCEPTANCE or BLOCKED. No final tag/public release or VPS access is permitted.


## Owner Approval Requirements

Approved by Project Owner: Attila through the Engineering Task Builder on 2026-08-28 UTC.

The structured task definition and Authority Envelope have been explicitly approved. Framework 2.0 Standard Execution Authority permits iterative, reversible in-scope engineering without another Owner gate. Further scope changes, exceptional external actions, and Owner-reserved decisions require explicit Project Owner approval.
