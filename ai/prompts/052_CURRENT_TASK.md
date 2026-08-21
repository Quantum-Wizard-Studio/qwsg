# Current Engineering Task 052: QWSG Guided Guardian Activation Runtime Context and Diagnostics

## Task Metadata

- Task ID: `052`
- Task slug: `qwsg-guided-guardian-activation-runtime-context-diagnostics`
- Status: `active`
- Date opened: `2026-08-21` UTC
- Human authority: Project Owner
- Owner or lead-developer communication language: English

## Title

QWSG Guided Guardian Activation Runtime Context and Diagnostics


## Objective

Correct external release blocker `QWSG-051-F001` with the smallest coherent
product change: make guided Guardian activation communicate with the current
user's systemd manager through the same deterministically UID-derived,
filesystem-validated runtime context used by Smart Install/readiness, while
preserving the bounded runner's sanitized environment and fixed command/unit
contract. Add deterministic privacy-safe activation diagnostics that identify
the failed fixed stage and safely distinguish runtime-context/user-manager,
timeout, output-bound, and other fixed-operation failures. Prepare source and
release metadata for a later separately authorized private
`QWSG 1.1.0-rc.3` candidate without building, transferring, externally testing,
tagging, releasing, or publishing it.



## Scope

- Before installation, require explicit Project Owner authority to close Task
  051 truthfully as `complete with disclosed limitations — NOT READY FOR QWSG
  1.1.0 RELEASE`. Preserve `QWSG-051-F001` as OPEN/BLOCKING RELEASE BLOCKER,
  exact stop at guided Guardian activation, RC.2 source
  `6d3f79accd4d52b94c960eefa93e2f51fbc9a48c`, RC.2 archive SHA-256
  `73d045cbc5577d3e9921a44760ba316d2094cf13fafe82f873be9f3600547315`,
  reproducibility/package evidence, privacy-safe completed external facts,
  incomplete transfer provenance, unexecuted later checkpoints, immutable RC.1
  evidence, and Task 049 F002/F003 history. Never claim RC.2 acceptance PASS or
  continue the external workflow.
- Use the canonical Builder transaction only after Task 051 prompt/history
  satisfy the truthful completed-task validator. Archive Task 051 once and
  install Task 052 plus matching history atomically. Do not divert, manually
  replace prompts, fabricate completion evidence, or create Task 053.
- Extract effective-UID runtime-directory derivation and validation now private
  to `internal/assessment` into a small shared package tentatively named
  `internal/userruntime`. It owns deterministic
  `/run/user/<decimal-effective-uid>` construction and lstat validation:
  canonical identity, directory type, no symlink, effective-UID ownership, and
  no group/other permission bits. Return stable bounded reason tokens and a
  trusted context value, never caller-provided environment.
- Keep `internal/runner.Bounded` as defense in depth. It continues replacing
  child environment with fixed PATH and C locale and accepts only the
  construction-time canonical `XDG_RUNTIME_DIR=/run/user/<uid>` form. Add no
  ambient/general/caller-controlled environment facility.
- Refactor assessment to consume the shared validator without changing Task
  047 classifications, Task 050 evidence/guidance tokens, human/JSON output,
  read-only behavior, or supported-host semantics.
- Refactor `internal/userservice` to consume the same validated context and
  attach it only to fixed absolute `/usr/bin/systemctl --user` operations.
  Preserve exact unit `qwsg-guardian.service`, bounded time/output, no shell,
  no arbitrary argv, and `daemon-reload` then `enable --now` ordering.
- Add a bounded fixed user-manager reachability/diagnostic probe where needed
  to distinguish inability to reach the manager from generic fixed-operation
  failure. Share stable reachability semantics with assessment where practical;
  never classify from locale-variable raw stderr. Running/degraded remain
  reachable; transient, unavailable, timeout, output-limit and unrecognized
  states remain distinct.
- Add typed deterministic activation failures carrying stage and cause tokens.
  Required stages: runtime-context preparation, manager/reachability,
  daemon-reload, and enable/start. Required causes: context missing/unsafe,
  manager unreachable, timeout, output limit, and other safely distinguishable
  fixed-operation failure. Preserve identity through wrapping; expose no raw
  stderr, environment, username, path, inventory, credential or unbounded data.
- Update guided setup human diagnostics to state failed stage, privacy-safe
  explanation, configuration preservation and a safe next action derived from
  the typed cause. Do not invent a command for ambiguity or claim activation.
- Preserve setup success: after explicit `y`, valid configuration and installed
  unit on a supported host with validated running manager execute the fixed
  sequence, wait within the existing bound for fresh canonical evidence, then
  independently report readiness. Do not weaken readiness/freshness/integrity.
- Preserve Task 048 resumability, Task 046 one-recipient notification,
  installer/uninstaller, non-root operation, systemd hardening/resource limits,
  deterministic JSON conventions, QWSG 1.0 and Community/Pro boundaries.
- Advance `VERSION`, private release notes, known limitations, installation
  filename references and version-coupled release validation to
  `1.1.0-rc.3` only after correction gates pass. RC.3 is future acceptance
  source identity, not an artifact/readiness claim. Preserve RC.1/RC.2 records.
- Recommend a separately authorized future Task 053 that builds RC.3 twice
  from one exact clean integrated commit and restarts clean-host acceptance at
  Checkpoint 01. Historical host-independent evidence may be referenced only
  for chronology; all artifact-identity and product-dependent external gates
  must be re-executed for RC.3. Do not create Task 053.



## Out of Scope

- No Task 051 continuation, workaround, readiness remediation execution,
  external VPS/SSH action, SMTP credential/provider action or external claim.
- No RC.3 build, reproducibility construction, transfer, execution, tag,
  Forgejo Release, upload, publication, signing, announcement or final release.
- No arbitrary XDG runtime or DBus inheritance; no caller/user/config/CLI
  environment, path, UID, executable, argv, unit or username. Add no DBus
  address unless necessary and deterministically derived from validated state.
- No shell, arbitrary command, sudo, privilege escalation, package/system-unit
  change, lingering/login repair, network/provisioning, generic controller or
  remediation executor.
- No runner/readiness/Guardian/installer/notification/security/JSON or
  Community/Pro weakening; no RC.1/RC.2 or Task 049 evidence alteration.
- Do not modify Owner-owned `docs/architecture/QWCS_MIGRATION_BLUEPRINT.md`.
- No inferred staging/commit/push, broad staging, or Task 053.



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

- Require Owner-authorized Task 051 terminal prompt/history state `complete
  with disclosed limitations — NOT READY FOR QWSG 1.1.0 RELEASE`, preserving
  QWSG-051-F001 and exact RC.2/checkpoint evidence. Rotate only through the
  canonical Builder into sole active Task 052 with matching history.
- Verify root/date/user, `main`, canonical HTTPS origin, exact Owner-approved
  closure HEAD/origin, `0/0`, empty index, expected Task 051 closure/Task 052
  lifecycle changes and only the preserved Owner-owned unrelated path.
- Verify RC.2 source/artifact/reproducibility and F001 stop evidence; immutable
  RC.1 source/artifact, Task 049 F002/F003, v1.0.0, LICENSE; and no RC.3
  artifact/tag/release.
- Audit assessment host/fixtures, runner policy/tests, userservice/controller,
  guided setup/readiness/setup-flow, systemd unit, release plumbing, and Tasks
  047–051 histories/architecture.
- Reproduce in fixtures that assessment receives validated UID-derived context
  while current userservice does not. Host-local checks are read-only, not
  external evidence.
- Run pre-change build, focused/full/race, vet, format, release, shell/static
  systemd, package/install/uninstall, security, Framework 21, Builder 38,
  lifecycle 28, diversion 36, job/test-task, Git and preservation gates.



## Snapshot Requirements

Before target changes create a unique mode-0700 `/tmp` snapshot containing a
readable tracked-HEAD archive, Task 051 closure and Task 052 prompt/history/
Builder source, target payloads/absence records, Git/mode/ACL/tool identities,
RC.1/RC.2/v1.0.0/LICENSE preservation and bounded restore instructions. Record
Owner content only as excluded metadata without reading/copying/hashing it.
Verify checksums, readability, collision absence, rollback and retention; no
external private host data or credential.



## Risk Assessment

- **False activation success — critical:** unchanged readiness and fresh
  canonical evidence independently gate success.
- **Environment injection — critical:** UID derivation, filesystem validation,
  typed context, runner backstop; no ambient/caller values.
- **Shared regression — critical:** one validator, separate fixed operation
  allowlists, unchanged assessment semantics.
- **Misdiagnosis — high:** typed stage/cause, no raw stderr/guessed commands,
  cause-specific safe action and preserved configuration.
- **Privilege/mutation — critical:** ordinary-user fixed unit only; no sudo,
  system unit, lingering, package or arbitrary behavior.
- **Bounds — high:** preserve timeout, output limit and process-group cleanup.
- **RC collision — critical:** changed bytes become RC.3 only; no Task 052 build.
- **Lifecycle — critical:** Task 051 closes NOT READY with blocker open through
  canonical transaction, never diversion/fabricated PASS.
- **Scope — high:** no generic systemd/DBus/executor/external retest/Task 053.



## Planned Work

1. Validate Owner-authorized Task 051 closure and Builder rotation; verify
   baseline, protected identities, pre-gates and snapshot.
2. Specify shared `internal/userruntime` context/result and stable validation,
   reachability and activation stage/cause tokens.
3. Move UID runtime validation from assessment; adapt assessment compatibly.
4. Give userservice the same trusted context only for fixed systemctl IDs; add
   bounded reachability diagnosis and typed stage/cause errors.
5. Render deterministic safe setup diagnostics and retain fresh-evidence wait.
6. Add context, controller, CLI, readiness and security fixtures; perform
   staged local acceptance without changing the development user service.
7. Advance source/release metadata/docs to private RC.3; build no artifact.
8. Run focused/full/race/vet/format/release/security/governance/package,
   rollback, diff and preservation gates.
9. Stop for separate source-integration authority; after integration close
   Task 052 only under separate authority. Do not create/build Task 053/RC.3.



## Rollback Plan

- Stop Task 052 tests, verify snapshot/targets and absence of external/service
  mutation. Restore only literal targets after later-edit checks; remove only
  created paths with absence records. Never reset/clean/checkout broadly or
  touch external/SSH state, Owner content, RC.1/RC.2, v1.0.0 or LICENSE.
- Re-run focused/full/race/vet/format/release/security/package/governance,
  lifecycle, snapshot, diff/index and preservation checks after rollback.



## Deliverables

- Shared validated effective-UID user-runtime context with stable outcomes.
- Fixed userservice activation using that context only for exact operations.
- Typed activation stage/cause errors and safe actionable setup diagnostics.
- Security/failure/success/CLI/readiness fixtures and full regression evidence.
- Updated runtime/setup architecture and relevant operator documentation.
- Private RC.3 source metadata/plumbing, no artifact/acceptance claim, immutable
  RC.1/RC.2 preservation.
- Snapshot, rollback simulation, exact security/diff audit and Task 052 history.



## Verification

- Context: effective UID/canonical decimal path; valid/missing/symlink/special/
  wrong-owner/unsafe-mode/lstat; no ambient XDG/DBus/caller override.
- Runner: fixed sanitized environment, canonical trusted XDG only, rejection of
  arbitrary/duplicate/noncanonical/DBus/HOME, timeout/output/process cleanup.
- Assessment: all Task 047/050 classifications/evidence/guidance unchanged for
  running/degraded/transient/unavailable/unsafe/timeout/output/failure/unknown.
- Userservice: exact order/executable/unit; manager reachability states;
  daemon-reload and enable-start success/failure/timeout/output; early stop;
  typed stage/cause; no raw stderr/caller inputs.
- Setup: explicit confirmation; successful fixed activation; preserved config
  on failure; distinct deterministic diagnostics/safe actions; bounded wait;
  no false success; decline/noninteractive/resume compatibility.
- Readiness: enabled/active/fresh evidence independently required after success;
  no weakening or stale READY.
- Regression: notification/credential, Guardian cadence/resources/unit,
  installer/uninstaller, config/setup-flow, JSON, QWSG 1.0, Community/Pro.
- Security: no shell/sudo/arbitrary executable/argv/environment/path/UID/unit/
  user, ambient DBus, privilege/system/lingering/package/network/remediation,
  raw host output, secret or private evidence.
- Build, focused/full/race, vet, format, shell/static systemd, package/install/
  uninstall, RC.3 release-check, Framework 21, Builder 38, lifecycle 28,
  diversion 36, job/test-task, Git, snapshot/rollback/diff and all protected
  identity gates pass.



## Documentation Updates

- Update `docs/architecture/SMART_SETUP_GUIDED_ACTIVATION.md` and, only as
  needed, `docs/architecture/SMART_INSTALL_READINESS.md` for shared validation.
- Update installation, setup, troubleshooting, support and security/privacy
  docs for cause-specific guided failure and preserved configuration.
- Add `docs/release/RELEASE_NOTES_1.1.0-rc.3.md`; update `VERSION`, known
  limitations and version-coupled checks only as required. Preserve RC.1/RC.2.
- Update Task 052 history and concise engineering history at delivery; preserve
  Task 051 terminal NOT READY and do not create Task 053.



## Completion Criteria

Task 052 completes only when Task 051 is archived truthfully as NOT READY with
F001 OPEN/BLOCKING and exact RC.2/external evidence; one shared effective-UID
runtime validator safely serves assessment and guided activation; assessment
remains compatible; userservice uses that context only for fixed QWSG systemctl
operations; typed safe diagnostics distinguish runtime/manager, daemon-reload,
enable/start, timeout, output and other bounded failures; supported fixtures
complete explicit activation and unchanged readiness only after enabled/active/
fresh evidence; every failure preserves config and gives a safe non-guessed
action; all full/race/security/package/governance/rollback/preservation gates
pass; source advances to `1.1.0-rc.3` with no artifact/acceptance claim; and no
external action, credential, RC.3 build/transfer, tag/release/publication,
Task 053 or unrelated work occurs. Integration and Task 052 closure each
require separate explicit Owner authority.


## Owner Approval Requirements

Approved by Project Owner through the Engineering Task Builder on 2026-08-21 UTC.

The structured task definition has been explicitly approved for implementation. Further scope changes require explicit Project Owner approval.
