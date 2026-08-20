# Current Engineering Task 048: QWSG Smart Setup and Guided Guardian Activation

## Task Metadata

- Task ID: `048`
- Task slug: `qwsg-smart-setup-guided-guardian-activation`
- Status: `complete`
- Date opened: `2026-08-20` UTC
- Human authority: Project Owner
- Owner or lead-developer communication language: English

## Title

QWSG Smart Setup and Guided Guardian Activation


## Objective

Orchestrate the canonical setup/configuration, Community email, Smart Install/readiness, user-systemd Guardian, and Current Operator State capabilities from Tasks 045–047 into one deterministic operator journey. Interactive `qwsg setup` must reconstruct progress from real evidence, preserve valid values, ask only for operator-owned choices, guide secure credentials and an explicit notification test, offer separately confirmed ordinary-user Guardian activation, wait boundedly for fresh Guardian evidence, and finish with evidence-backed READY/PARTIAL/NOT READY/UNKNOWN plus an exact next action. Do not persist wizard progress, duplicate subsystem truth, equate accepted configuration or service state with operation, execute remediation, elevate privilege, or break existing automation.


## Scope

- Audit and reuse Task 045 configuration resolution/storage, typed setters/defaults and setup/config commands; Task 046 one-recipient SMTP configuration, credential store, preflight/test/provider and delivery evidence; Task 047 Assessment Model/Registry, classifications, summaries, probes, recommendations, install check and readiness; and Guardian user unit, lifecycle, checkpoint, presentation and Current Operator State.
- Add one versioned presentation-independent setup-flow model with deterministic phase, action, reason and evidence tokens; blocking/optional state; confirmation/input requirements; optional registry-owned remediation; retry/revalidation; and privacy-safe JSON. Human prompts and future GUI/automation are adapters over this model.
- Derive progress on every invocation from host assessment, Configuration Source, credential availability, notification evidence, unit installed/enabled/active facts, lingering, and fresh integrity-checked Guardian evidence. Add no wizard progress file, success flag, or competing readiness store.
- Use ordered logical phases: environment, base configuration, notification decision/configuration, credential readiness, notification verification, activation readiness, Guardian activation, fresh-evidence verification, and final readiness. Stable identifiers must be versioned and documented.
- Stop before activation for mandatory `missing_required` or `incompatible`; mandatory `unknown_requires_verification` prevents READY and yields bounded manual guidance. Display only Task 047 platform-gated registry remediation and never execute it.
- Evolve interactive bare `qwsg setup` into the guided journey when stdin/stdout are terminals. Preserve the existing configuration review/write as its base phase; reuse valid values and change them only by explicit operator choice.
- Preserve `qwsg setup --accept-defaults`, `--set`, explicit `--config`, existing human/JSON behavior, `qwsg config ...`, `qwsg notification preflight|test|credential set --from-file`, `qwsg install --check`, and `qwsg readiness`. These deterministic forms never prompt or activate services.
- Add an explicit read-only setup plan/status option under `qwsg setup` with human/JSON output for automation/future GUI. It never prompts, writes, contacts SMTP, or mutates services. Bare nonterminal setup continues to fail deterministically unless explicit noninteractive options are supplied.
- Guide only canonical keys: base values and notification enablement, one administrator recipient, host/port, implicit TLS or required STARTTLS, auth, sender, username, credential reference and timeout. Reuse immediate canonical validation; never invent addresses, providers, identities, credentials, or business choices; fail closed on invalid/unsafe existing configuration.
- Add hidden interactive credential entry only if a bounded no-echo controlling-terminal adapter can be proven safe without adding an external module/network dependency. Pass bytes directly to the existing credential store and avoid argv/environment/output/log/config/state/JSON exposure. Otherwise guide the protected mode-0600 `--from-file` path. Preserve that automation path unmodified.
- Offer the existing controlled notification test only after separate confirmation; never contact SMTP merely because setup starts. Distinguish disabled, incomplete, missing credential, configured-unverified, verified, temporary failure, and incompatible/auth/TLS/config failure. Never create a fake Guardian incident.
- If explicit test success cannot currently be reconstructed truthfully on rerun, add only the smallest notification-owned private integrity-checked evidence, versioned and bound to relevant configuration identity/freshness, with atomic path protections and invalidation on relevant change. It is notification evidence, never wizard success.
- Add an injected narrow controller for only absolute `/usr/bin/systemctl --user daemon-reload` and fixed `qwsg-guardian.service` enable/start operations. Require explicit activation confirmation; bound execution, argv, environment and output; allow no shell, PATH selection, caller-controlled unit/argv, sudo, loginctl, package command, or general service runner. Decline/failure preserves valid configuration and cannot yield READY.
- Keep privileged actions guidance-only. Never invoke sudo or change lingering. An exact lingering command may be printed only if added as Ubuntu-24.04-amd64 registry data with strictly validated current-user substitution, structured argv/unambiguous display, and revalidation; otherwise provide manual verification without a command.
- Model lingering as boot-before-login evidence: disabled may permit current logged-in operation but prevents unattended logout/reboot readiness; unavailable remains unknown. Unit enablement alone never proves reboot persistence.
- After activation, revalidate unit installed/enabled/active through canonical probes and then require fresh integrity-checked Guardian coverage. Use injected clock/waiter/cancellation and non-tight bounded polling capped from validated cycle timeout plus documented allowance. Timeout/stale/invalid/absent evidence yields retry guidance, not READY.
- Refactor operational assessment composition out of CLI-local code as needed so setup and `qwsg readiness` share evidence, aggregation and rendering inputs. Do not fork readiness truth or vocabulary.
- Final output covers environment/dependencies, configuration, notification, Guardian service/evidence, and overall readiness. Ready Guardian core plus notification not ready is overall PARTIAL; mandatory missing/incompatible is NOT READY; mandatory unknown is UNKNOWN; only complete qualified evidence yields READY.
- Every non-ready result selects a deterministic next internal step, operator question, registry recommendation, manual verification, retry/revalidation, or documentation reference. Never speculate.
- Keep evidence/state, planner, action execution and terminal presentation separate for future GUI consumption. Keep `install.sh` the low-level deterministic noninteractive copier.
- Establish `README.md` as the primary operator/end-user entrypoint: concise product purpose, Community capabilities, basic use, principal commands, setup/readiness journey, and a prominent link to the dedicated installation guide. Move or link developer/build/architecture detail to its canonical specialist documentation so it does not dominate the user introduction.
- Keep `docs/installation/INSTALL.md` as the single canonical source installation guide and package it as archive-root `INSTALL.md`. It must cover the supported platform and prerequisites, archive/checksum/manifest verification, low-level installation, bundled Smart Install preflight, guided setup, Community notification/credential/test flow, explicit Guardian activation, readiness, logout/reboot/lingering considerations, uninstall/rollback, installation troubleshooting, documentation locations, and a clear link back to `README.md`.
- Make `README.md` and the canonical installation guide cross-reference each other clearly in the repository, release archive, and installed documentation. Avoid a second independently maintained installation narrative; archive-root `INSTALL.md` is a release copy of the canonical source.
- Update release assembly to include archive-root `README.md` and `INSTALL.md`, verify their identity against canonical source, and update the low-level installer to install both as read-only documentation under `/usr/local/share/doc/qwsg/README.md` and `/usr/local/share/doc/qwsg/INSTALL.md` (respecting staged `DESTDIR`). Documentation remains available after the archive is removed and is removed only according to the existing uninstall code/documentation preservation contract.
- Keep installer documentation UX deterministic and noninteractive. After copying artifacts, print concise archive/installed documentation locations, the exact next guided setup/readiness commands, and how to read the guides. Do not automatically invoke a pager/editor, require an optional viewer, prompt, or dump long documentation. Any optional display mechanism must use QWSG/portable existing capabilities, be explicitly selected, bounded, and add no installation dependency.
- Continue deferring VPS profile because it changes no current behavior. Document future brand/intro only as a presentation extension; add no questionnaire, profile, animation, graphics, GUI, or assets.
- Community receives the full guided local journey, one recipient, controlled test, activation and readiness. Future Pro may automate authorized orchestration over the same model, but Task 048 adds no Pro function.
- Update minimum canonical architecture/operator/security/installation/systemd/notification/readiness/troubleshooting/history documentation. Preserve v1.0.0, LICENSE, release artifacts/tag/source, Tasks 045–047, and the Owner draft byte-for-byte.


## Out of Scope

- No package/repository/remediation execution, sudo/password automation, root provisioning, privilege escalation, lingering/firewall/SSH/kernel/control-panel/mail-server mutation, or automatic repair.
- No generic runner, shell evaluation, arbitrary service/process execution, inbound listener, automatic SMTP assessment connection, or unbounded wait.
- No duplicate config/credential/notification/readiness/Guardian store or validator; no wizard phase/success persistence or fake readiness.
- No QWS API/key/account, billing, entitlement, managed email, unlimited Community recipients, remote provisioning/fleet, web/desktop GUI, role monitoring, firewall/log/traffic analytics, AI remediation, VPS profile, or branding implementation.
- No secret echo/argv/environment/config/state/log/JSON/report leakage, silent activation, fake incident, guessed remediation, or external-evidence claim not actually performed.
- No interactive installer documentation wizard, automatic pager/editor launch, mandatory external documentation viewer, duplicated install guide, or unconditional long terminal documentation dump.
- No v1.0.0/LICENSE/tag/artifact/release/Owner-draft modification; no Task 049; no staging, commit, push, tag, release, publication, reset, cleanup, or unrelated work.


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

- Verify root/date/user, `main`, canonical origin, HEAD and `origin/main` `bedf5761922f55732f0fc80e5769425a0a21da01`, ahead/behind `0/0`, empty index, clean tracked tree, and pushed Task 047 integration.
- Verify canonical idle after unique completed Task 047, no Task 048 collision, and passing job, next-task, Framework, Builder, lifecycle, diversion and test-task validators.
- Verify the only unrelated visible untracked path is `docs/architecture/QWCS_MIGRATION_BLUEPRINT.md`; record hash/mode/owner/ACL/size without substantive access or mutation. Inventory/exclude ignored Builder/build/dist/local paths.
- Verify v1.0.0 annotated object `7abc83da6185199606c2d76ac7d3504ddd78cf68`, target `177535e44b2ce5ed9efd73ab0793ffe6881f0cd6`, LICENSE SHA-256 `e66792f9831509ecc888f0c06d64b89ae87814fe37842e12df6b716b404f4126`, archive SHA-256 `edfba7366adf2c1ce0a8ce56369bb0dc5ad11326c4e3d1e301625a5313292fa5`, sidecar and release evidence.
- Read governance/AGENTS/skill, Tasks 037/045/046/047, configuration, email, Smart Install, Guardian/runtime/operator-state/presentation, runner/security, package/systemd, operator/release docs and relevant tests.
- Produce an ownership/interface map and record interactive/option/exit contracts, mutations, unit search/sandbox, config/state modes, notification-evidence limits, lingering and fresh-evidence timing. Identify duplication/refactor boundaries first.
- Audit the current operator suitability and cross-links of root `README.md` and `docs/installation/INSTALL.md`; release archive contents/names; `scripts/build-release.sh`; installer document copy loop/output; installed `/usr/local/share/doc/qwsg` layout; uninstall handling; and every documentation reference to installation. Record which document is canonical versus a packaged copy.
- Run pre-change build; focused setup/config/store/credential/SMTP/assessment/Guardian/operator-state/presentation/runner/CLI tests; full/race/vet/format/whitespace/security; staged package/install/uninstall; systemd static verification; and governance suites. Stop on unexplained variance.


## Snapshot Requirements

Before implementation, create a unique external mode-0700 `/tmp` snapshot containing a verified complete tracked-HEAD archive and payload/absence records for every intended target. Record Git/upstream/index/status, hashes, modes/owners/ACLs, tools/fixed command paths, privacy-safe systemd facts, immutable release identities, lifecycle files, and approved Builder source hash. Include checksums, archive readability, collision guards, retention and exact bounded restore instructions; verify isolated rollback before completion.

- Record Owner draft metadata/hash/exclusion only; do not copy or substantively read it.
- Capture no credential, address, private config/state, inventory, host identifier, or sensitive output. Use generated fixtures in isolated private HOME/XDG/config/state/install/archive/SMTP/systemd roots.
- Do not mutate real packages, repositories, firewall, lingering, panels, mail stack, or real user service in automated acceptance. Snapshot grants no excluded authority.


## Risk Assessment

- **False READY — critical:** canonical evidence only; config/unit/SMTP syntax never substitutes for delivery/fresh Guardian evidence.
- **Privilege/host mutation — critical:** confirmed fixed ordinary-user QWSG systemctl only; no sudo/loginctl/package/general runner.
- **Secret leakage — critical:** proven no-echo adapter or protected-file fallback; adversarial disclosure tests.
- **CLI regression — high:** terminal-gated orchestration; preserve all deterministic paths and exits.
- **Opaque/stale progress — high:** no wizard file; reconstruct every rerun.
- **Notification-evidence ambiguity — high:** minimal subsystem-owned identity-bound evidence only if required.
- **Injection/execution — critical:** absolute executable, fixed unit/argv/environment, bounded timeout/output, no metadata execution.
- **Wait/hang — high:** injected bounded cancellation-aware non-tight waiter.
- **Lingering overclaim — critical:** distinguish session operation from boot-before-login.
- **Partial corruption — high:** validate before atomic store/action; failure never creates success.
- **Vocabulary divergence — high:** reuse/refactor Task 047 model and aggregation.
- **Scope/release/Owner damage — critical:** hashes, snapshots, exact exclusions and path boundaries.
- **Documentation drift/operator dead end — high:** one canonical install-guide source, verified packaged copies, reciprocal links, installed read-only locations, and concise deterministic next-step output.


## Planned Work

1. Verify baseline, audit interfaces/evidence, run gates, create snapshot and rollback simulation.
2. Document orchestration phases, resumability, ownership, next-action precedence and GUI/action boundary.
3. Refactor reusable operational readiness collection without changing existing CLI contracts.
4. Add versioned setup-flow model/planner over injected canonical evidence.
5. Add read-only human/JSON setup plan/status rendering.
6. Evolve terminal bare setup into resumable guidance while retaining deterministic setup/config paths.
7. Add guided canonical base/notification input and atomic validation/persistence.
8. Add proven hidden credential adapter or protected-file guidance fallback.
9. Integrate confirmed notification test and minimal durable notification evidence only if necessary.
10. Add narrow confirmed user-service activation controller and revalidation.
11. Add lingering guidance and bounded fresh-evidence wait.
12. Present shared final readiness and deterministic next actions.
13. Add comprehensive unit/integration/adversarial/regression/staged/rollback tests.
14. Rework `README.md` as the operator entrypoint; complete the canonical install guide; package/install both guides with identity checks; and add concise deterministic installer documentation/next-step output.
15. Update remaining docs/history, complete gates, archive Task 048 and return idle without Task 049.


## Rollback Plan

- Stop only exact Task 048 fixture processes/isolated units. Verify snapshot manifests, target hashes, absence/collision records and ownership before restoration; stop on later-edit ambiguity.
- Restore literal pre-existing targets and remove only exact created targets with proven prior absence/current ownership/hash. Never reset, checkout, clean, glob, broadly extract, or mutate host services.
- Preserve real config/credentials/state, Tasks 045–047, Builder data, release evidence, LICENSE, tag/artifacts and Owner draft. Re-run all focused/full/security/package/governance/Git/preservation gates after rollback simulation.


## Deliverables

- Canonical Tasks 045–047 ownership map and orchestration architecture.
- Versioned resumable setup-flow/next-action model with privacy-safe deterministic human/JSON output and no progress file.
- Guided terminal setup plus read-only plan/status; unchanged noninteractive contracts.
- Guided canonical base and one-recipient notification configuration.
- Safe credential strategy/implementation only when proven; protected-file compatibility.
- Controlled notification verification with truthful evidence semantics.
- Explicit fixed ordinary-user Guardian activation; no privileged execution.
- Truthful lingering and bounded fresh Guardian evidence verification.
- Shared final readiness and exact next actions.
- Full tests, staged Community acceptance, snapshot/rollback evidence, docs and lifecycle history.
- Operator-first `README.md`; complete reciprocal-linked canonical installation guide; verified archive-root `README.md`/`INSTALL.md`; installed read-only copies with discoverable locations; and noninteractive installer next-step guidance.


## Verification

- Verify baseline/final Git/lifecycle/ACL, Builder hash, Owner draft and immutable release identities.
- Run build, focused subsystem/flow/CLI tests, full tests, race, vet, format, whitespace, security/static, package/archive/install/uninstall and systemd checks; all Framework/Builder/lifecycle/diversion/test-task suites; simulate completion to idle without Task 049.
- Planner tests cover every phase, schemas/tokens/order/JSON, classifications/summaries, mandatory/optional precedence, actions, resume states and no duplicated state.
- Compatibility tests prove all existing setup/config/notification/install/readiness forms and exits, no unexpected prompts/activation, and valid values preserved.
- Interaction tests cover terminal/nonterminal, EOF/interrupt/cancel, invalid/retry, confirmations/declines, no automation blocking and privacy-safe output.
- Config/credential tests cover safe defaults, immediate validation, unsafe paths, atomic preservation, one recipient, no echo/disclosure, bounds/modes/owner/symlink/hard-link/special-file, redirected rejection and protected-file compatibility.
- Notification tests cover disabled/incomplete/missing credential/preflight/unverified/verified/auth/TLS/timeout/failure, evidence binding/invalidation/freshness, no automatic network/fake alert and failure preservation.
- Activation tests prove exact absolute executable/fixed argv/unit/environment, time/output/cancellation bounds, confirmation/decline/failure, malicious metadata rejection, no shell/PATH/sudo/loginctl/package/general service action and canonical revalidation.
- Lingering/Guardian tests cover enabled/disabled/unavailable, safe recommendation handling, logged-in versus reboot semantics, unit states, running without evidence, fresh/stale/invalid/absent evidence, bounded success/timeout/interruption and no tight polling/false READY.
- Readiness tests prove parity and READY/PARTIAL/NOT READY/UNKNOWN precedence, including core-ready plus notification-unready PARTIAL.
- Task 045–047 and QWSG 1.0 regressions pass; audits prove no excluded mutation/execution/listener/secret/Pro/profile/branding/release behavior.
- Staged journey covers archive checks -> install check -> isolated install -> guided config/credential -> controlled SMTP fixture -> explicit activation fixture -> fresh evidence -> final readiness -> rerun/restart -> uninstall preservation. Label fixture, staged-local, host-local systemd, external host/control-panel and real-provider evidence separately; claim only performed evidence.
- Documentation/package tests verify README/INSTALL reciprocal links, required operator topics, archive-root filenames and byte identity with canonical sources, installed `/usr/local/share/doc/qwsg` modes/contents, concise installer location/next-command output, no prompt/pager/editor/long dump, archive-removal survivability, and documented uninstall/rollback behavior.


## Documentation Updates

- Add canonical Smart Setup/Guided Activation architecture and update Smart Install, Configuration Activation, Community Email, Guardian, Product Architecture/System Map/roadmap/history references without duplicating truth.
- Update Installation, Quick Start, Setup, Operations, Support, Troubleshooting, Security/Privacy, Upgrade/Uninstall and CLI docs for journey, resume, noninteractive use, credentials/test, activation, lingering, bounded evidence, readiness and recovery.
- Make root `README.md` primarily operator-facing and make `docs/installation/INSTALL.md` the complete dedicated installation source. Cross-link them explicitly; package the latter as `INSTALL.md`; document repository, archive and installed locations; and ensure every path ends with what to run next and how to check readiness.
- Document Community completeness, future Pro consumption, deferred VPS profile, presentation-only future brand extension, evidence limitations, possible later 1.1.0, snapshot/rollback and preservation.


## Completion Criteria

Task 048 is complete only when the Task 047 idle baseline is verified; canonical subsystem ownership remains intact; one versioned planner derives resumable progress/actions solely from evidence; terminal setup guides the full Community journey without a progress file; existing automation is compatible; blockers stop safely with only proven guidance; valid values persist; one recipient is enforced; secrets never leak; notification READY requires controlled delivery evidence; activation is separately confirmed and fixed to bounded ordinary-user QWSG systemctl; no sudo/loginctl/package/remediation executes; lingering semantics are truthful; fresh Current Operator State, not service state, proves Guardian operation; final/setup and standalone readiness share aggregation; core-ready notification-unready is PARTIAL and mandatory missing/unknown cannot be READY; every failure has a safe next action; `README.md` is an operator-first entrypoint and reciprocally links the complete canonical installation guide; release archives contain verified root `README.md`/`INSTALL.md`; installation preserves read-only copies under `/usr/local/share/doc/qwsg` and prints concise locations/next commands/readiness guidance without interaction or viewer dependency; all tests, staged/rollback evidence, docs/history and governance gates pass; idle returns without Task 049; and v1.0.0, LICENSE, artifacts/tag/source, Tasks 045–047 and Owner draft remain unchanged. Any duplicated truth, documentation drift/dead end, opaque progress, false READY, silent/privileged/generic execution, unbounded wait, secret leak, CLI regression, speculative remediation, scope expansion or missing evidence blocks completion.


## Owner Approval Requirements

Approved by Project Owner through the Engineering Task Builder on 2026-08-20 UTC.

The structured task definition has been explicitly approved for implementation. Further scope changes require explicit Project Owner approval.
