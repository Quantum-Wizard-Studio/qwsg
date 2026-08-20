# Current Engineering Task 047: QWSG Smart Install and Readiness Assessment Foundation

## Task Metadata

- Task ID: `047`
- Task slug: `qwsg-smart-install-readiness-assessment-foundation`
- Status: `complete`
- Date opened: `2026-08-20` UTC
- Human authority: Project Owner
- Owner or lead-developer communication language: English

## Title

QWSG Smart Install and Readiness Assessment Foundation


## Objective

Implement the common, versioned, evidence-backed foundation for QWSG Smart Install. From an unpacked release archive or installed binary, an ordinary operator must be able to determine whether QWSG can run, what mandatory or optional conditions are satisfied or missing, what remains uncertain, what exact proven operator-run remediation is available on the supported platform, and what to do next. Generalize Task 046's five-state readiness vocabulary rather than create an incompatible taxonomy; adapt notification preflight to the common contract. Add deterministic human and machine-readable assessment and composite readiness for installation/environment, configuration, notification, Guardian/service, and overall operation. Assessment is read-only, non-root, bounded, privacy-safe, and non-mutating. Keep `packaging/release/install.sh` a narrow immutable-artifact installer. Solve bootstrap by running the archive's bundled QWSG binary before system installation. Structure detection and remediation-plan data for a future separately authorized Detect -> Plan -> Ask -> Execute -> Verify -> Continue engine, but implement only Detect, Plan, and operator guidance.


## Scope

- Audit the complete installer, archive, uninstaller, Task 045 setup/configuration, Task 046 notification preflight/SMTP, Guardian/runtime, systemd user service, supported-platform, clean-host, security/privacy, diagnostics, Inventory collector/capability, presentation, packaging, and troubleshooting architecture before design or modification.
- Establish one common canonical assessment contract with schema/model version, stable requirement IDs, deterministic order, privacy-safe evidence, and exactly `satisfied`, `missing_required`, `missing_optional`, `unknown_requires_verification`, and `incompatible`. Preserve meanings across human and JSON output. Do not reuse unrelated Inventory or Drift classifications.
- Move or adapt SMTP-local `Classification` and `Finding` into the common model without breaking `qwsg notification preflight`, configuration validation, Guardian preconditions, or existing JSON. Notification detection remains owned by `internal/smtpnotification`; the common package owns shared types, validation, aggregation, ordering, and presentation-neutral data.
- Define one deterministic versioned registry, outside CLI code. Entries cover stable ID, purpose/message key, requirement class, required/optional/recommended disposition, workflow/capability, platform predicate, justified version constraints, bounded probe ID, remediation mapping, elevation, required revalidation, and privacy/security class.
- Treat remediation mappings and commands as canonical registry data. A recommendation identifies the exact recognized platform/version/architecture, exact display command and structured argv/steps, privilege requirement, managed-stack/control-panel compatibility guard, and rerun requirement. No external data becomes executable text.
- Distinguish mandatory runtime dependencies, install-time dependencies, optional feature dependencies, environment capabilities, configuration requirements, and recommended non-blocking capabilities. Do not label a dependency mandatory merely because a current script uses it.
- Freeze initial supported remediation to the production contract: Ubuntu 24.04 LTS, Linux amd64/x86-64, systemd 255+, glibc-compatible userspace, ordinary non-root runtime user, working systemd user manager, and required local filesystem semantics. Other distributions receive no guessed package commands. Unsupported architecture/init/root runtime and incompatible versions classify truthfully; insufficient evidence is unknown.
- Record the actual minimal dependency set. The release binary has no external Go module or Go runtime dependency. Shell, `sha256sum`, archive extraction, `install`, `cp`, `mkdir`, `rm`, `rmdir`, and `awk` are release/install/uninstall tooling concerns, not automatically Guardian runtime dependencies. systemd/user manager, glibc, atomic rename/advisory lock/ownership/private modes, writable private config/state, and outbound DNS/TCP/TLS trust when email is enabled are capabilities or runtime conditions, not all package dependencies.
- Implement injected probe abstractions with fixtures/clocks/runner. Prefer direct Go/runtime/syscall/filesystem evidence. Any external production probe must map from a registry-owned ID to an absolute allowlisted executable and fixed bounded argv, with deadline, output limit, stable error parsing, no shell, no caller-controlled executable/arguments, no secret input, and no mutation.
- Assess OS identity/version from bounded `/etc/os-release`, architecture, ordinary/root identity, systemd version, user-manager availability, filesystem/path safety and private-storage feasibility, configuration existence/validity, installed artifact visibility, Guardian unit installed/enabled/active state, fresh canonical Guardian evidence, and notification readiness through Task 046. Unit/process state alone must not prove Guardian monitoring.
- Treat lingering as distinct boot-before-login readiness. It may be reported only from bounded evidence and never changed. Missing lingering does not make logged-in local Guardian operation impossible, but prevents a ready claim for unattended boot-before-login.
- Do not perform broad package inventory or infer remediation from package presence. Where package ownership, a control panel, managed service stack, or coexistence is ambiguous, return unknown with manual verification and no command.
- Reuse configured SMTP, credential availability, TLS mode/system trust, and explicit delivery-test evidence conservatively. Default readiness makes no SMTP connection. Do not install or recommend Postfix, Exim, Sendmail, a relay, or sendmail-compatible facility.
- Define immutable structured results with registry/model version, timestamp, phase, recognized platform, ordered findings, stable evidence/reason tokens, optional registry-owned recommendations, domain summaries, blockers, manual verification, and next actions. Exclude secrets, credential references, raw host identifiers/paths/output/errors, destinations, and sensitive values.
- Aggregate deterministically. Missing/incompatible mandatory requirements block the affected workflow; mandatory unknown prevents a ready claim; optional absence affects only its feature. A disabled optional feature is not a core runtime failure.
- Define readiness domains for installation, environment/dependencies, configuration, notification, Guardian/service, and overall operation. Use stable summary values `ready`, `partial`, `not_ready`, and `unknown`, distinct from the five finding classifications, with documented precedence and evidence rules.
- Overall readiness must permit Guardian core `ready`, notification `not_ready`, overall `partial`; notification verified plus ready core may yield overall `ready`. Disabled, configured-but-unverified, verified, invalid, and delivery-failed notification states remain distinct. Mandatory absence never yields ready.
- Add the smallest coherent command surface: `qwsg install --check [--format human|json]` for pre-install/environment assessment and `qwsg readiness [--format human|json]` for composite post-install operation. Preserve `qwsg setup`, `qwsg config`, and `qwsg notification preflight|test`; focused notification preflight becomes an adapter/view over common findings.
- Make `qwsg install --check` work as both `./bin/qwsg install --check` from an unpacked archive and installed `qwsg install --check`. It must not require installation, write configuration/state, invoke setup, contact SMTP, enable services, or mutate the host. Document stable ready/blocked/usage/internal exits consistent with canonical CLI policy.
- Keep `install.sh` deterministic, noninteractive, and narrowly privileged. Its checks remain limited to artifact integrity/platform/copy safety. Update guidance to run bundled preflight as the intended non-root user before `sudo ./install.sh`, then setup/readiness before explicit activation. Do not duplicate registry/detection in shell or run rich assessment as root.
- Document the journey: download/checksum/unpack -> bundled preflight -> operator remediation/revalidation -> low-level artifact install -> setup/operator values and credentials -> readiness/focused notification test -> explicit user-unit install/enable -> revalidation. Mandatory blockers stop; optional findings do not block unrelated core operation.
- Keep administrator email, SMTP provider, sender, username, credential, future API key, and VPS purpose/profile as operator configuration/product choices, not detected dependencies. Guide them through setup without inventing values.
- Do not persist a VPS-purpose field. Document a future versioned setup-profile extension point, avoiding premature role/profile schema obligations. Do not implement role-specific monitoring.
- Make remediation-plan data consumable by a future Pro coordinator while exposing no executor. Future automation must retain separate Detect, Plan, Ask, Execute, Verify, Continue stages, explicit authority, audit, rollback, and revalidation.
- Community receives the complete assessment, supported remediation guidance, readiness reporting, one administrator recipient, and local Guardian operation. Add no entitlement checks and do not reserve Smart Install by crippling Community.
- Update only the minimum canonical architecture, CLI, platform, installation, setup, operations, notification, security/privacy, troubleshooting, package, roadmap/history, and lifecycle documentation.
- Treat Tasks 045–047 as possible evidence for a later separately authorized 1.1.0 line, but do not change `VERSION`, release, rebuild/republish v1.0.0, tag, push, or publish.
- Preserve byte-for-byte `v1.0.0`, its release source/artifacts/checksums, LICENSE, completed Task 045/046 records, and `docs/architecture/QWCS_MIGRATION_BLUEPRINT.md`.


## Out of Scope

- No package install/remove/update, repository change, package-manager execution, remediation execution, root provisioning, sudo/polkit/setuid, privilege escalation, service mutation, lingering change, firewall/SSH/kernel/control-panel/mail-server mutation, or arbitrary host configuration.
- No generic command runner, shell evaluation, caller-supplied probe argv, untrusted PATH execution, remote script, plugin probe, broad network scan, inbound listener, or automatic SMTP test.
- No guessed remediation for unsupported distributions; no inferred package equivalents; no mail infrastructure or certificate/system package recommendation without exact proven mapping and compatibility.
- No interactive setup redesign, secret collection in assessment, raw editor workflow, invented personal/business values, SMTP provider selection, VPS-profile persistence, role monitoring, log/firewall/traffic analysis, or AI remediation.
- No Pro executor/download/orchestration, remote provisioning, fleet management, QWS API/key/account, managed notification, billing/entitlement, multiple recipients, GUI, or public release.
- No replacement or duplication of Inventory, Health, Presentation, Configuration, Notification, Runtime, Guardian, or systemd truth engines.
- No claim of external clean-host, control-panel, or real-provider evidence unless actually performed; fixture, staged, host-local, external-host, and provider evidence remain distinct.
- No modification of QWSG 1.0.0, LICENSE, published artifacts, tag, or release source. No Task 048 creation or implementation.
- No Git staging, commit, push, tag, publication, cleanup/reset, or unrelated worktree authority. Do not substantively read/change the Owner draft.


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

- Verify root, UTC date/user, `main`, exact HEAD and `origin/main` `3aa332be188bbb62789c5262a23a7ef0bd3ee8b9`, `0/0` ahead/behind, canonical origin, empty index, clean tracked tree, and pushed Task 046.
- Verify canonical idle after unique completed Task 046, no Task 047 collision, and passing `bin/job --check`, `next-task.sh --check`, Framework, lifecycle, diversion, test-task, and Builder validation.
- Record visible/relevant ignored paths. Verify the only unrelated visible untracked item is the Owner draft; record hash/mode/owner/ACL/size without substantive access or mutation.
- Verify local/remote annotated v1.0.0 object/target, release-source commit, LICENSE, retained archive/sidecar hashes, and release evidence unchanged.
- Read mandatory governance; relevant Tasks 037–046; Product Definition, Functional Specification, Product Architecture, System Map, roadmap, Inventory/capability, runner/security, configuration/setup, notification/SMTP, Guardian/runtime/service, presentation/operator state, packaging/install/uninstall, platform/acceptance, operations, troubleshooting, and privacy material; inspect all relevant code/tests.
- Record archive/bootstrap contents, installer/uninstaller tools, setup paths/mutations, service/lingering flow, Guardian prerequisites, Task 046 classification limitations, collector registry/descriptors, bounded runner, presentation states, and JSON conventions.
- Produce an evidence table classifying every dependency as runtime, install-time, optional feature, environment, configuration, or recommendation, including why script tools do/do not imply product dependencies.
- Audit exact Ubuntu package/remediation evidence. If mapping is unproven, keep unknown and omit a command.
- Run pre-change build, focused config/SMTP/Guardian/collector/runner/CLI tests, full/race/vet/format/whitespace/security, packaging/install/uninstall staging, systemd verification where supported, and all governance suites. Stop on unexplained differences.


## Snapshot Requirements

Before any Task 047 implementation modification, create a unique external mode-`0700` `/tmp` snapshot with payloads or absence records for every intended target. Record root/remote/branch/HEAD/upstream/index/status, hashes, owners/modes/ACLs, tool versions, relevant service facts, privacy-safe config/state metadata, and immutable release identities. Include and verify checksums, archive readability, collision guards, retention, and an exact bounded restore procedure.

- Record but do not copy the Owner draft; preserve only metadata/hash/exclusion.
- Capture no credential, private configuration, raw inventory/state, SMTP destination, or sensitive output. Tests use generated fixture values in isolated private roots.
- Use separate isolated HOME/XDG/install/state/fixture roots for staged/systemd/filesystem/probe acceptance. Do not mutate real user services, lingering, packages, repositories, firewall, panel, or mail stack.
- Retain the snapshot through completion and record it in history. Snapshot creation grants no excluded authority.


## Risk Assessment

- **Host mutation/privilege escalation — critical:** no executor; read-only allowlisted bounded probes and mutation audits.
- **Unsafe remediation — critical:** commands only from tested supported registry mappings; uncertain/managed-stack cases get no command.
- **Vocabulary divergence — high:** one shared five-state model; summary states are explicitly separate.
- **False readiness — critical:** evidence precedence blocks mandatory unknown/missing/incompatible findings and isolates optional notification.
- **Bootstrap/root-context error — high:** bundled Go binary assesses as intended user; shell only verifies/copies artifacts.
- **Generic execution/injection — critical:** compiled probe IDs and exact argv, no shell/PATH/metadata execution, bounded output/time.
- **Privacy leakage — high:** stable tokens and minimal platform facts; no identifiers, secrets, raw paths/config/output/errors.
- **Platform/control-panel conflict — critical:** strict supported predicate and compatibility guards; package presence never proves safe remediation.
- **Filesystem probe damage — high:** no writes by default; isolated temporary roots only in explicit tests.
- **Notification/systemd regression — high:** common-model adapter with Task 045/046 and Guardian/service regression suites; unit state never substitutes for Guardian evidence.
- **Schema lock-in — high:** version contracts and defer VPS profile persistence.
- **Release/Owner damage — critical:** hash/exclude immutable release and Owner draft; no Git/release authority.


## Planned Work

1. Verify baseline, audit dependencies, run pre-change gates, and create the verified implementation snapshot.
2. Document shared findings, registry, probes, remediation plan, domain aggregation, privacy, bootstrap, and provisioning separation.
3. Add a common assessment package with versioned types, registry validation, deterministic order, evidence tokens, remediation records, summaries, aggregation, and serialization validation.
4. Adapt notification preflight to common findings while preserving SMTP ownership and all behavior.
5. Add the initial immutable Ubuntu 24.04 amd64/systemd 255+ registry/platform adapter, encoding only proven requirements/remediation and validating all references/safety constraints.
6. Implement bounded read-only direct/allowlisted probes with injected fixtures for platform, identity, user manager, filesystem/path, config, service/Guardian, failures, and timeouts.
7. Build installation/environment assessment and composite readiness from existing canonical evidence without a new truth store.
8. Add `qwsg install --check` and `qwsg readiness` human/JSON UX, exits, blockers, recommendations, manual verification, rerun, and next actions.
9. Update installer/archive guidance while preserving low-level install/uninstall behavior and eliminating shell-side registry duplication.
10. Add model/registry/probe/presentation/CLI/security tests and all setup/notification/Guardian/package regressions.
11. Run release-style staged acceptance from archive preflight through install, setup, notification and Guardian readiness fixtures, and uninstall preservation; label evidence truthfully.
12. Update documentation/history, run final gates, verify rollback/preservation, archive Task 047, and restore canonical idle without Task 048.


## Rollback Plan

- Stop only exact recorded Task 047 test processes; mutate no real service.
- Verify snapshot manifests, hashes, readability, absence records, target inventory, and restore instructions; stop if later changes may exist.
- Restore only exact pre-existing targets and remove only exact Task 047-created paths with proven prior absence/ownership. Never use broad reset/checkout/clean, unresolved globs, or repository-wide extraction.
- Preserve pre-task config/state, Tasks 045/046, ignored Builder data, release evidence, LICENSE, v1.0.0, and Owner draft.
- Remove only validated isolated temporary roots and retain diagnostic evidence.
- Rerun focused/full/race/vet/format, assessment/notification/Guardian/package/security/privacy/no-mutation, governance, Git, ACL/mode, release, and Owner checks; record results.


## Deliverables

- Dependency audit with correct six-way requirement classification.
- Common versioned five-state assessment contract and deterministic registry with safe platform-gated remediation data.
- Bounded read-only probe abstraction and strict supported-platform adapter without generic execution/mutation.
- Notification preflight adapter with no Task 046 regression.
- Archive/installed `qwsg install --check` and composite `qwsg readiness` in human/JSON formats.
- Readiness aggregation distinguishing ready core plus unavailable notification from fully ready operation.
- Documented two-phase bootstrap and unchanged low-level installer privilege boundary.
- Focused/full/adversarial tests, staged lifecycle acceptance, snapshot, rollback proof, and operator documentation.


## Verification

- Verify baseline/final repository/Git/lifecycle/ACL state, Owner draft preservation, and immutable release identities.
- Run build; focused assessment/collector/runner/config/SMTP/Guardian/CLI tests; full tests; bounded-cache race; vet; format; diff; security/static; package/systemd checks.
- Run Framework, Builder, lifecycle, diversion, and test-task suites; validate active state and simulate completion/archive to canonical idle.
- Registry tests cover versions, unique IDs, order, complete metadata, constraints/references, platform gates, remediation/elevation/revalidation, unsafe/duplicate/dangling rejection, and deterministic JSON.
- Classification tests cover all five states; mandatory/optional missing/unknown; incompatibility; six requirement classes; and ready/partial/not-ready/unknown precedence.
- Platform tests cover supported Ubuntu 24.04 amd64/systemd 255, wrong architecture, old systemd, unsupported/malformed/missing OS data, uncertain version, root user, unavailable user manager, container ambiguity, and glibc/filesystem uncertainty.
- Remediation tests prove exact supported command generation/order and prove no command for unsupported/unknown/ambiguous/control-panel/managed-stack/SMTP infrastructure cases; test escaping/injection resistance.
- Probe tests prove exact allowlist/argv, no shell/PATH/metadata execution, bounded timeout/cancellation/output, stable failures, malformed evidence handling, privacy, determinism, and no host mutation.
- Filesystem/security tests cover unsafe/symlink/owner/mode conditions and trace absence of package/repository/service/lingering/firewall/network/config writes or secret/raw-host output.
- CLI tests cover archive/installed invocation, human/JSON parity, ready/blockers/optional/incompatible/unknown/internal exits, deterministic ordering, exact rerun/next action, and localization-ready rendering.
- Task 045 regressions cover absent/valid/invalid/unsafe config and idempotent setup; assessment never writes configuration.
- Task 046 regressions cover disabled, configured/unverified, verified evidence, invalid, missing credential, no network/TLS failure, preflight parity, privacy/cardinality/provider behavior, and continued monitoring.
- Guardian/systemd tests distinguish unit absent/installed/enabled/active, manager and lingering states, configuration precondition, and fresh/stale/absent canonical evidence without service mutation.
- Packaging acceptance covers release archive -> bundled preflight -> staged install -> installed preflight/setup/readiness -> notification states -> Guardian fixtures -> uninstall, preserving hashes/modes/layout/data and proving no shell registry duplication.
- Label fixture, staged-local, host-local read-only, external clean-host, control-panel, and real-provider evidence separately. Do not require or claim unperformed external/package/provider work.
- Verify snapshot/rollback simulation, history completeness, canonical idle simulation, Community/Pro boundary, and all preserved assets.


## Documentation Updates

- Add a canonical Smart Install and Readiness Assessment architecture document owning registry, classifications, probes, remediation, aggregation, bootstrap, security, and future provisioning.
- Update Community email/configuration activation, Guardian/system map, Product Architecture/roadmap/history references for the common adapter without duplicating SMTP truth.
- Add/update operator Smart Install guidance plus Installation, Quick Start, Setup, Operations, Support, Troubleshooting, Security/Privacy, Upgrade/Uninstall, and CLI documentation with one coherent journey.
- Document supported platform, six requirement classes, five finding states, four domain summary states, command eligibility/no-command rules, privilege/revalidation, user-systemd/lingering distinctions, notification partial readiness, and unknown/unsupported behavior.
- Document archive bootstrap, low-level installer, explicit activation, Community/Pro separation, future provisioning stages, and deferral of VPS-purpose persistence.
- Record decisions, dependency evidence, validation/acceptance labels, snapshot/rollback, limitations, later 1.1.0 recommendation, release/Owner preservation, and idle completion in history.


## Completion Criteria

Task 047 is complete only when the Task 046 idle baseline is verified; one common versioned five-state contract and deterministic registry exist; every dependency is correctly classified; supported detection/remediation is strict to Ubuntu 24.04 amd64/systemd 255+ with no guesses; probes are read-only, bounded, allowlisted, private, fixture-testable, and non-generic; notification preflight uses the common model without regression; bundled/installed `qwsg install --check` and composite `qwsg readiness` work in deterministic human/JSON forms; mandatory blockers cannot yield ready; optional notification failure permits ready core and partial overall; verified notification plus ready core permits overall ready; recommendations are proven data and never executed; installer remains low-level/noninteractive/narrowly privileged with no duplicate truth; setup owns operator values/secrets; no VPS profile field or Pro executor exists; all tests/security/privacy/no-mutation/package/systemd/governance/rollback gates pass; staged acceptance is truthful; docs are operator-ready; history is complete and archived; canonical idle returns with no Task 048; and QWSG 1.0.0, LICENSE, release artifacts/tag/source, Tasks 045/046, and Owner draft remain unchanged. Any mutation, automatic remediation, privilege escalation, generic execution, guessed command, false readiness, privacy leak, regression, unbounded probe, duplicated truth, release mutation, or missing mandatory evidence blocks completion.


## Owner Approval Requirements

Approved by Project Owner through the Engineering Task Builder on 2026-08-20 UTC.

The structured task definition has been explicitly approved for implementation. Further scope changes require explicit Project Owner approval.
