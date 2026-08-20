# Current Engineering Task 050: QWSG Smart Install Actionable Remediation and Verification Guidance

## Task Metadata

- Task ID: `050`
- Task slug: `qwsg-smart-install-actionable-remediation-verification-guidance`
- Status: `complete`
- Date opened: `2026-08-20` UTC
- Human authority: Project Owner
- Owner or lead-developer communication language: English

## Title

QWSG Smart Install Actionable Remediation and Verification Guidance


## Objective

Correct QWSG-049-F002 and QWSG-049-F003 with the smallest coherent extension of the canonical Assessment Model and Registry: preserve fail-closed classifications, make supported-host findings explainable and actionable in equivalent deterministic human/JSON output, distinguish systemd user-manager evidence without guessing, and replace unexplained filesystem uncertainty with justified read-only evidence or a precise structured manual-verification plan. Assessment remains non-mutating and never executes guidance. Prepare the integrated source for a later separately authorized private `1.1.0-rc.2` acceptance candidate; do not build, transfer, tag, publish, or resume physical acceptance in Task 050.


## Scope

- Preserve Task 047's single canonical assessment/registry ownership and extend Assessment Model/Registry from 1.0 to a backward-conscious 1.1 contract. Keep the five classifications (`satisfied`, `missing_required`, `missing_optional`, `unknown_requires_verification`, `incompatible`), requirement classes, dispositions, domain summaries, deterministic ordering, and existing inert `remediation` compatibility field.
- Add registry-owned structured actionable guidance that does not require a command. A finding may carry a validated guidance plan selected by stable platform, requirement, classification, and evidence token. The same plan drives human and JSON output.
- Model at least stable explanation token, blocking effect, verification actions, operator actions, privilege requirement (`none`, `administrator`, or `manual_verification`), optional proven structured command/argv plus display form, manual-verification flag, safety-note tokens, and mandatory revalidation action. Bound counts/lengths, reject controls/newlines/unsafe argv, and preserve localization-ready tokens.
- Keep existing `Remediation` data for compatibility and integrate it into guidance rather than creating duplicate CLI truth. Unsupported/ambiguous platforms receive no command. A command is displayed only from a validated exact registry mapping.
- Correct the systemd user-manager probe boundary. Audit the current runner behavior that strips `XDG_RUNTIME_DIR` and `DBUS_SESSION_BUS_ADDRESS`; reproduce and test that this can make a reachable user manager appear unavailable. Do not classify a manager missing solely because the probe removed required safe session context.
- Add narrow, validated user-runtime context for the fixed user-manager probe without inheriting arbitrary environment. Derive the expected `/run/user/<effective-uid>` identity internally; reject symlink, wrong owner, unsafe type/mode, missing runtime directory, timeout, output overflow, and ambiguous bus state with distinct privacy-safe evidence tokens. Never accept caller-selected path, executable, arguments, environment, unit, or username.
- Parse only documented/stable `systemctl --user is-system-running` states under `LANG=C`/`LC_ALL=C`. Treat proven `running` and the existing justified `degraded` state as available; treat transient states as retry/verification guidance; distinguish runtime-directory unavailable/unsafe, bus/manager unavailable, probe timeout/failure, and unrecognized output without exposing raw stderr. Do not depend on unbounded or locale-variable error text for a remediation decision.
- For every safely distinguishable user-manager state, supply a registry-owned explanation, bounded verification action, privilege boundary, revalidation command (`qwsg install --check`), and only a remediation command proven by Ubuntu 24.04 amd64 evidence. Ambiguous states must say manual verification is required and provide no plausible-looking command.
- Do not assume package absence. Do not recommend `apt`, package names, PAM edits, `loginctl`, lingering changes, service changes, sudo, logout/reboot, or session repair unless the exact state-to-action mapping is independently proven for the supported platform, compatibility reviewed, and fixture/host-local evidence establishes safety. Task 050 may legitimately ship no command for an ambiguous cause while still making verification actionable.
- Replace the unconditional filesystem finding with a deterministic filesystem evidence function behind the existing Host/probe abstraction. Prefer bounded direct Go read-only facts for the effective QWSG configuration/state locations: safe nearest-existing-ancestor resolution, no symlink traversal, current-user ownership where required, `statfs`/mount-type evidence, and a documented allowlist only for filesystems whose needed atomic rename, advisory flock, Unix ownership/mode, and private `0700/0600` semantics are justified.
- Do not claim semantics from a filesystem name alone without canonical evidence. Unknown, remote, pseudo, overlay/container, inaccessible, unsafe-path, or unsupported filesystem evidence remains `unknown_requires_verification` or `incompatible` as justified, but must include a precise structured verification plan and no speculative install command.
- If write behavior is needed to prove rename/flock/mode semantics, it must not occur in default assessment. Design any isolated user-owned temporary behavioral verification as a separately explicit, bounded, opt-in verification action with guaranteed cleanup, symlink/path protections, no system configuration mutation, and fixtures; implement it only if repository and platform evidence proves it is necessary and safe. Default `qwsg install --check` remains read-only.
- Generate finding-specific next actions from canonical guidance instead of only `resolve_required_findings`. Preserve `rerun_qwsg_install_check` and make revalidation explicit. Guided setup may consume the same guidance but must not duplicate or execute it.
- Human output must state what is wrong/uncertain, blocking effect, what to verify, operator action, privilege requirement, exact command only when proven, safety note, and revalidation. Avoid internal jargon and raw host details.
- JSON must expose equivalent guidance deterministically with schema/model/registry versioning suitable for future GUI and authorized Pro planning. Keep existing finding/remediation fields readable and classifications compatible; document the schema evolution.
- Verify notification preflight/adaptation, guided setup/readiness, Guardian/user-service behavior, installer boundary, Community one-recipient behavior, and all QWSG 1.0 observation/runtime behavior remain unchanged.
- Update canonical Smart Install, installation, troubleshooting, support, security/privacy and relevant setup documentation. Record that Community receives complete actionable guidance and Future Pro may consume, but not execute in Task 050, the same Detect/Explain/Verify/Plan data.
- Preserve immutable historical `1.1.0-rc.1` commit `ff2eb2b12499f5daf3b5ba11b1f8d7fc562f8a31`, artifact SHA-256 `aa139faaccc1cc85b50cfe0eedee9436539ae1c3071e01d8e9ed9283fc7f8239`, Task 049 evidence/findings, v1.0.0, LICENSE, published artifacts/tag/source, and the Owner draft.
- Because product/package bytes change, prepare source identity and release documentation for `1.1.0-rc.2` without building an artifact. Update only version-coupled release plumbing/tests/docs required to prevent any changed bytes from being called RC.1. RC.2 construction and external acceptance require a later separate task/Owner gate.


## Out of Scope

- No Task 049 physical acceptance continuation, external VPS access/mutation, manual workaround, Owner coaching around F002/F003, SMTP credential request/handling, real provider action, or external evidence claim.
- No automatic command execution, remediation executor, package installation/removal, apt/repository change, sudo/password handling, privilege escalation, PAM/system configuration edit, automatic login/session repair, lingering change, service mutation during assessment, firewall/SSH/control-panel change, or arbitrary shell/process execution.
- No caller-controlled executable, argv, environment, filesystem target, username, unit name, command template, shell interpretation, plugin probe, inbound listener, network probe, remote provisioning, QWS API/key/account, billing, entitlement, Pro execution, fleet, GUI, role monitoring, analytics, AI remediation, or automatic repair.
- No rewriting, rebuilding, replacing, transferring, or relabeling `1.1.0-rc.1`; no RC.2 artifact build; no public/final v1.1.0 tag, Forgejo Release, upload, publication, signing claim, announcement, or release readiness claim.
- No Task 051 creation/installation, no inferred Git staging/commit/push authority, no blanket staging, and no modification of Owner-owned content.
- Do not divert Task 049 into `ai/test_tasks`: it is real production acceptance evidence, not an aborted test, and diversion would release/reuse production ID 049 rather than create Task 050.


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

- Before Task 050 installation, require explicit Owner authorization to finalize Task 049 as `complete with disclosed limitations — NOT READY FOR QWSG 1.1.0 RELEASE`, despite its originally broader completion gate, because the external mandatory Smart Install gate truthfully failed. Preserve F001/F002/F003, private RC.1 provenance, checkpoint evidence, unexecuted gates and restart position; do not claim acceptance PASS.
- Use the normal canonical Builder transaction only after Task 049 prompt/history both satisfy the completed-task validator. Archive Task 049 and install Task 050 atomically. Do not use diversion, manual prompt replacement, concurrent active prompts, or a fabricated pause/resume state.
- Verify repository root, UTC date, ordinary user, branch `main`, canonical HTTPS origin, HEAD/origin at the exact Owner-approved Task 049 closure baseline (planning baseline `ff2eb2b12499f5daf3b5ba11b1f8d7fc562f8a31`; stop if closure authority changes it without an updated Owner-approved baseline), ahead/behind `0/0`, empty index, and exact expected Task 049 evidence/Owner-draft status.
- Verify sole active Task 050 after Builder installation, matching history, Task 049 archived once with truthful terminal NOT READY result, no Task 051, and Framework/job/lifecycle/diversion/test-task validity.
- Verify `VERSION=1.1.0-rc.1` at planning baseline, immutable RC.1 commit/artifact provenance, annotated v1.0.0 object/target, published v1.0.0 archive/checksum, LICENSE SHA-256 `e66792f9831509ecc888f0c06d64b89ae87814fe37842e12df6b716b404f4126`, and Owner-draft SHA-256 `c83a0b2e5269de51acd9485b44fbe4f752e1f7f172c180e4126f26e91fdf8a73`.
- Audit Task 047 assessment/model/registry/host/fixed-runner/tests/docs, Task 048 setup-flow/user-service orchestration, Task 049 F002/F003 evidence and protocol, runner environment policy, SMTP adapter, CLI human/JSON presentation, filesystem/path/state requirements, release/version plumbing and security boundaries.
- Reproduce in fixtures and, where safely available, host-local read-only evidence that stripping validated user-runtime context changes `systemctl --user` reachability. Do not treat sandbox denial as external product evidence.
- Run pre-change build, focused/full/race tests, vet, formatting, release-check, Framework 21, lifecycle 28, Builder 38, diversion 36, test-task/job, shell/systemd static, Git whitespace, security/exclusion and RC.1/v1.0.0 preservation gates. Stop on unexplained variance.


## Snapshot Requirements

Before target changes, create a unique external mode-0700 `/tmp` implementation snapshot containing a readable complete tracked-HEAD archive, exact Task 050 prompt/history/Builder source, Task 049 archived evidence identities, all intended target payloads or absence records, Git/status/diff/mode/ACL/tool identities, RC.1/v1.0.0/LICENSE preservation metadata, and bounded restore instructions. Record Owner draft metadata/hash only. Verify checksum manifest, archive readability, collision absence, isolated rollback and retention. Capture no external private host data or SMTP secret.


## Risk Assessment

- **False remediation — critical:** distinct causes remain distinct; commands require exact supported registry proof; ambiguous states receive verification/manual guidance only.
- **False missing user manager — critical:** current runner strips required runtime context; validate only canonical UID-derived context and test healthy/missing/unsafe/transient/ambiguous cases.
- **Assessment mutation — critical:** default assessment performs no writes, service changes, sudo, package action or configuration change; any optional behavioral filesystem probe requires separate explicit action and cleanup proof.
- **Path/symlink/ownership — critical:** use bounded UID-derived paths, nearest-existing-ancestor logic and no symlink/special/wrong-owner acceptance.
- **Schema compatibility — high:** version the extension, retain classifications/remediation compatibility, deterministic JSON and Task 046/048 consumers.
- **Sensitive output — high:** emit tokens and bounded facts, never raw stderr, usernames, runtime paths, mount inventory, config, credentials or host identifiers.
- **Lifecycle corruption — critical:** Task 049 cannot coexist, be diverted, or later reactivate; require truthful Owner-authorized terminal closure and normal transactional Builder rotation.
- **RC identity collision — critical:** RC.1 is immutable evidence; changed source uses RC.2 metadata only and produces no artifact in Task 050.
- **Scope expansion — high:** no general host doctor, command engine, provisioning, Pro executor, system repair, external acceptance or release work.


## Planned Work

1. Validate the Owner-authorized Task 049 terminal closure and transactional Task 050 installation; verify baseline, governance, protected identities, pre-change gates and snapshot.
2. Specify Assessment/Registry 1.1 actionable-guidance schema, validation, compatibility, deterministic ordering and token/presentation ownership. Extend existing remediation rather than duplicate it.
3. Correct the bounded runner/user-runtime context boundary and implement evidence-specific systemd user-manager findings. Prove healthy context is not stripped and ambiguous failures never gain a guessed command.
4. Implement read-only filesystem evidence for actual QWSG path ancestors and justified local filesystem allowlist. Provide structured manual verification for unresolved semantics; keep default assessment non-mutating.
5. Attach guidance by exact platform/requirement/classification/evidence and derive specific next actions/revalidation. Update human and JSON output from the same plan.
6. Integrate guidance into setup-flow consumption without adding execution or wizard state. Verify notification/readiness/Guardian compatibility.
7. Advance source/release documentation to `1.1.0-rc.2` identity without building an artifact; preserve RC.1 evidence byte-for-byte.
8. Run focused/full/race/vet/format/release/security/governance tests, staged fixture acceptance, no-write proof, rollback simulation, exact diff/security audit and preservation checks.
9. Finalize and archive Task 050 only after correction acceptance passes and Owner later authorizes source integration. Do not create Task 051. A future separately authorized acceptance-continuation task will build RC.2 and restart the external protocol.


## Rollback Plan

- Stop Task 050 processes/tests, verify snapshot and current exact target ownership/hashes, and confirm no external/QWSG service/remediation process was started.
- Restore only pre-existing Task 050 target files from the verified tracked archive after checking for later edits; remove only Task 050-created paths with explicit snapshot absence records. Never reset/clean/checkout broadly, touch Task 049/RC.1 private evidence, Owner data, v1.0.0, user runtime state, systemd services or host configuration.
- Re-run focused/full/race/vet/format, release-check, no-write/security, Framework/lifecycle/Builder/diversion/job, RC.1/v1.0.0/LICENSE/Owner preservation, Git diff/index and isolated rollback validation.


## Deliverables

- Assessment Model/Registry 1.1 structured actionable-guidance contract with compatibility and security validation.
- Evidence-specific systemd user-manager detection/guidance that corrects stripped safe runtime context and never guesses remediation.
- Read-only filesystem-semantics evidence plus precise structured manual verification for unresolved cases.
- Equivalent deterministic human and JSON guidance with finding-specific next actions, privilege and revalidation.
- Focused fixtures/security/no-write tests and full regression evidence for Tasks 045–049, notification, setup, readiness, Guardian and QWSG 1.0 behavior.
- Updated Smart Install/installation/troubleshooting/support/security/setup architecture and operator documentation.
- RC.2 source identity/release-note/plumbing preparation with no artifact, tag, transfer or publication; immutable RC.1 preservation evidence.
- Snapshot, rollback simulation, staged local acceptance, exact changed-path/security audit and completed Task 050 history.


## Verification

- Model/registry tests cover schema/version, deterministic ordering, command-free guidance, explanation/verification/operator/privilege/manual/safety/revalidation fields, exact evidence selection, bounds, invalid tokens/control characters, hostile argv/display text, unsupported platform and backward-compatible remediation/classification behavior.
- User-manager fixtures cover validated healthy runtime context; stripped/missing/unsafe/wrong-owner/symlink runtime directory; running/degraded; transient states; unavailable manager/bus; timeout/output limit/unknown output; supported versus ambiguous guidance; no raw stderr/privacy leak; and no command for uncertain causes.
- Filesystem fixtures cover supported local type, unknown/remote/pseudo/overlay/inaccessible/unsafe path, nearest existing ancestor, ownership/mode/symlink hazards, deterministic evidence, and a useful manual plan whenever automatic proof is insufficient. Default assessment creates/modifies/removes no file.
- CLI tests prove actionable human answers to all seven operator questions, equivalent deterministic JSON, stable exit `4` for blockers, specific next actions, explicit privilege/revalidation, no generic-only dead end, and no command when certainty is insufficient.
- Security tests prove no shell, caller-selected command/argv/env/path/user/unit, automatic sudo/package/service/session/lingering/configuration action, host mutation, inbound/network action, secret/raw host output or future Pro executor.
- Task 045 setup/config, Task 046 SMTP/preflight/one-recipient, Task 047 classification/readiness, Task 048 setup-flow/activation, Task 049 findings fixtures, QWSG 1.0 observation/Guardian, packaging/install/uninstall and Community/Pro regression gates pass.
- Build, focused/full/race Go tests, vet, formatting, release-check updated for RC.2 source identity, shell syntax, static systemd verification, Framework 21, lifecycle 28, Builder 38, diversion 36, test-task/job, Git whitespace/index, exact diff, secret/privacy scan, snapshot/rollback and preservation checks pass.
- Staged local fixture acceptance reproduces the RC.1 F002/F003 inputs and proves corrected actionable output without changing the development host. No external clean-host, real SMTP, physical reboot or RC.2 artifact evidence is claimed.


## Documentation Updates

- Update `docs/architecture/SMART_INSTALL_READINESS.md` for Assessment/Registry 1.1, actionable guidance, user-runtime context, filesystem evidence and future executor boundary.
- Update `docs/installation/INSTALL.md`, `docs/release/TROUBLESHOOTING.md`, `docs/release/SUPPORT.md`, `docs/release/SECURITY_AND_PRIVACY.md`, and where needed `docs/release/SETUP_AND_CONFIGURATION.md` so supported operators receive the same canonical explanation/verification/privilege/revalidation journey.
- Add `docs/release/RELEASE_NOTES_1.1.0-rc.2.md`; update version-coupled release checks/docs to RC.2 without altering RC.1 notes, candidate, evidence or acceptance record.
- Update Task 050 history and concise milestone index only at verified delivery. Do not rewrite Task 049 findings as passed and do not create a successor acceptance task.


## Completion Criteria

Task 050 is complete only when Task 049 is truthfully preserved as terminal `NOT READY` evidence and archived through the canonical Builder; Assessment/Registry 1.1 supplies safe evidence-specific actionable guidance in equivalent human/JSON form; the healthy user-manager false-negative boundary is corrected; every user-manager uncertainty has a bounded verification/privilege/revalidation plan and no guessed command; filesystem semantics is satisfied from justified read-only evidence or carries a precise manual plan rather than unexplained unknown; default assessment remains non-mutating; all focused/full/race/security/governance/package/rollback/preservation gates pass; RC.2 source identity is prepared without artifact/release action; F002/F003 are corrected in local evidence but not falsely closed as external PASS; documentation/history are complete; and no external acceptance, SMTP credential action, RC.2 build, Task 051, tag, upload, publication, Pro executor or unrelated work occurs.


## Owner Approval Requirements

Approved by Project Owner through the Engineering Task Builder on 2026-08-20 UTC.

The structured task definition has been explicitly approved for implementation. Further scope changes require explicit Project Owner approval.
