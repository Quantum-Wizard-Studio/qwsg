# Current Engineering Task 065: QWSG Guided Installation and Setup Experience

## Task Metadata

- Task ID: `065`
- Task slug: `guided-installation-setup-experience`
- Status: `approved`
- Date opened: `2026-08-27` UTC
- Human authority: Project Owner
- Owner or lead-developer communication language: English

## Title

QWSG Guided Installation and Setup Experience


## Objective

Design, implement, validate, integrate, and canonically close a professional QWSG-owned guided installation and setup experience for Ubuntu 24.04 LTS amd64. A technically competent operator new to QWSG must be able to securely obtain the product, select English, Magyar, or Deutsch, understand the real installation stage and each requested decision, review a preflight-first mutation plan, install and configure QWSG, safely handle optional notification credentials, activate and verify Guardian, select an honest update policy, and finish with a localized readiness summary and clear next commands. Preserve reusable engine boundaries for future concise and unattended interfaces, and advance the private development candidate without publishing a release.


## Scope

- Inspect and evolve the existing release package, installer, CLI setup/install/readiness/update flows, configuration and credential stores, user-systemd activation, lingering guidance, documentation, tests, and private candidate plumbing needed for one coherent guided journey.
- Define a QWSG-owned installation engine with explicit immutable phases/subphases, deterministic state/progress contributions, typed plan/actions/results, mutation ledger, retry/failure state, rollback outcome, and interface-neutral inputs suitable for guided, concise, and future unattended adapters without duplicating business logic.
- Implement a terminal guided installer with language selection first; a consistent localized header/dashboard when terminal capabilities safely support it; truthful phase/stage/progress display; contextual explanations, expected value shape, examples, origin guidance, required/optional status, validation, confirmation, and next-step explanation; and a clean line-oriented fallback for non-capable terminals.
- Establish stable message identifiers and maintainable English canonical/fallback, Hungarian, and German catalogs covering installer phases, prompts, guidance, validation, warnings, plan, failures, readiness, and completion. Missing translations must fall back safely to English.
- Detect Ubuntu 24.04 LTS amd64 through an extensible platform registry and gracefully explain detected versus supported systems without mutation.
- Preserve preflight-first behavior: detect fresh/existing/upgrade-related state, safely discover what can be discovered, produce an operator-readable installation plan and optional technical details, obtain confirmation before privileged mutation, and never request safely discoverable data.
- Integrate installation paths, configuration initialization, Guardian user service, state directories, lingering requirements, optional email notification setup and protected credentials, update policy, service activation, readiness verification, resumable/repeatable setup, and localized completion summary.
- Implement update-policy configuration for manual-only and update-available notification where supported. Add a clean explicit automatic-policy foundation only if semantics are honest; do not claim or enable unattended privileged updating unless fully and safely implemented and validated.
- Reuse Task 064 authoritative release discovery, archive verification, provenance, migration, transactional replacement, and rollback foundations. Design a minimal transparent bootstrap/download-and-run entry point that securely obtains and starts trusted QWSG-owned installer code; pipe-to-shell must not be mandatory.
- Add deterministic unit, integration, pseudo-terminal/transcript, localization/catalog, progress truth, plan/no-mutation, credential-safety, failure/rollback, install-root/systemd harness, bootstrap/package, and user-journey tests.
- Perform proportional local acceptance and, after strong local validation, one bounded clean Ubuntu 24.04 amd64 fresh-install acceptance if a disposable supported host is already safely available. Preserve the existing official-1.1.0 VPS baseline; use it only for bounded compatibility/update observations that do not destroy that baseline.
- Update product, installation, administration, security, troubleshooting, architecture, release-readiness, changelog, and task lifecycle records; construct only a private candidate if materially required; perform targeted clean-fast-forward Git integration and canonical closure.


## Out of Scope

- Do not begin Server Health / Assessment Engine work, new monitoring features, remediation, fleet/Pro orchestration, web setup, package repositories, broad distribution support, or Task 066.
- Do not expand the official supported-platform contract beyond Ubuntu 24.04 LTS x86_64/amd64.
- Do not publish a final or prerelease QWSG release, create or push a tag, create a Forgejo Release, upload public artifacts, deploy to production, or announce a release.
- Do not modify, rebuild, retag, replace, or rewrite official QWSG 1.1.0 artifacts, tag, Release, checksums, historical evidence, or Task 064 preserved candidate evidence/lineage.
- Do not reinstall, reformat, erase, or sacrifice the useful official QWSG 1.1.0 state on the preserved test VPS.
- Do not silently enable unattended privileged mutation, install unrelated dependencies, guess unsupported host repair, expose secrets, or store credentials in source, command history, process arguments, logs, diagnostics, transcripts, ordinary summaries, or task evidence.
- Do not create a permanently independent large shell installer or a second release-verification/trust system when the Go binary and Task 064 architecture can own the behavior.


## Authority Envelope

- **Task targets and boundaries:** Framework 2.0 Standard Execution Authority applies to the objective across QWSG installer/setup engine and adapters, localization catalogs, CLI, bootstrap/release packaging, configuration/credential/update-policy/service/readiness integration, deterministic tests and harnesses, task-scoped documentation/evidence, private candidate construction if justified, targeted Git integration, and lifecycle closure. Server Health / Assessment Engine, broader platform support, public release, and Task 066 remain excluded.
- **Permitted external actions:** anonymous read-only use of canonical public QWSG Forgejo release endpoints and official Ubuntu documentation where necessary; ordinary dependency-free network verification; normal clean-fast-forward Task 065 Git pushes; one bounded privacy-safe fresh-install acceptance on an already available disposable Ubuntu 24.04 amd64 host after local gates; and bounded read-only or reversible compatibility observations on the preserved official-1.1.0 VPS that leave it operational at official 1.1.0. If clean-host acceptance requires Owner-controlled provisioning, reinstall, authentication, or physical interaction, request only that inherent action and resume under this task afterward.
- **Owner-reserved decisions:** public release/tag/Forgejo Release/upload/deployment/announcement; alteration of official 1.1.0 or Task 064 frozen evidence; destructive VPS provisioning/reinstallation; credential disclosure or replacement; materially expanded trust/signing/channel policy; unattended privileged auto-update activation; supported-platform expansion; material product architecture expansion; Server Health / Assessment Engine; and Task 066.
- **Task-specific STOP conditions:** STOP for credential/secret exposure risk, security/privacy regression, untruthful progress or readiness that cannot be corrected in scope, unsafe privileged destination/control, unavailable required rollback, unexplained mutation of preserved official-1.1.0 state/evidence, required external clean host that is not available without Owner-controlled provisioning, material scope/architecture expansion, or another genuine Framework 2.0 boundary. Ordinary implementation, UI, localization, build, test, packaging, bootstrap, simulated-install, or acceptance failures remain in diagnose -> classify -> correct -> retest.


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

- Verify the canonical QWSG root, ordinary user, branch main, Framework 2.0.0 VALID, lifecycle canonically idle after complete Task 064, empty index, HEAD/origin/main/direct Forgejo main agreement, and only explicitly excluded unrelated Owner content present and untouched.
- Verify VERSION is 1.2.0-rc.1; QWSG 1.1.0 remains the official public release; Task 064 private candidate/evidence and official objects are unchanged; and the preserved VPS final state is operational official 1.1.0.
- Characterize the current split journey: verified archive plus install.sh, install --check, setup, notification credential/preflight/test, readiness, user-systemd/lingering guidance, and update commands. Identify reusable contracts and gaps against the Owner-directed guided experience before selecting target architecture.
- Verify current package contents, install/uninstall/replace semantics, configuration and credential ownership/modes, state roots, user-service behavior, terminal handling, test facilities, network/trust assumptions, and development-version conventions. Do not access external hosts before snapshot and local design/validation justify a bounded runner.


## Snapshot Requirements

Before modifying tracked Task 065 targets, create a unique private mode-0700 snapshot under /tmp containing a readable exact tracked-HEAD archive; exact lifecycle prompt/history and Builder input identity; literal before-images or absence records for installer/setup/bootstrap/localization/configuration/credential/update-policy/service/readiness/packaging/documentation/version targets; Git/ref/release/Task-064 preservation identities; excluded Owner pathname metadata without reading or hashing its contents; and collision-aware bounded restore instructions. Verify permissions, hashes, archive readability, exclusions, rollback preconditions, and retention. Before any external mutation or private-candidate use, add a separate private acceptance-boundary ledger recording exact source/package/runner identities, allowed mutations, clean-host expectations, credential-redaction rules, rollback/cleanup, and preservation requirements.


## Risk Assessment

- HIGH — privileged installation or service activation could damage host paths or leave Guardian unavailable. Mitigate with typed preflight plan, fixed allowlisted destinations, confirmation, narrow elevation, mutation ledger, transactional package reuse, post-action readiness, and deterministic restoration.
- HIGH — bootstrap/download trust confusion could execute hostile bytes. Reuse canonical HTTPS origin, immutable release identity, checksum/archive/manifest/provenance verification, bounded acquisition, transparent commands, and no mandatory pipe-to-shell.
- HIGH — credentials could leak through prompts, terminal echo, history, argv, logs, transcripts, diagnostics, summaries, tests, or rollback evidence. Use no-echo terminal input, protected direct storage or private file-descriptor/file handoff, redaction-by-construction, restrictive modes, and adversarial scans.
- HIGH — decorative or monotonic progress could claim unfinished work. Derive display only from explicit completed/active engine states; retries never mark work complete; rollback and failure views distinguish attempted, mutated, restored, and pending states.
- MEDIUM — translations may diverge, omit critical safety meaning, or break formatting. Use stable typed message IDs, canonical English fallback, catalog completeness/placeholder validation, golden transcripts, and native-speaker-quality review to the extent deterministically testable.
- MEDIUM — terminal control may corrupt output or become inaccessible. Capability-detect conservatively, sanitize dynamic values, support no-color/line fallback, avoid emulator-specific behavior, and test pipes/dumb terminals/PTY dimensions.
- MEDIUM — existing installations and upgrade states may be mistaken for fresh installation. Detect and classify state before mutation, refuse ambiguity, preserve Task 064 update/rollback coherence, and make repair/reconfiguration safely repeatable.
- MEDIUM — a clean external host may be unavailable. Maximize deterministic install-root/systemd/PTY acceptance locally; stop only at the inherent Owner provisioning boundary and report exactly what remains unproven.


## Planned Work

1. Validate governance/lifecycle/Git/release baseline, read relevant architecture and prior Tasks 045-064, inspect current installation/setup/update contracts, install Task 065 through the Builder, create and verify the execution snapshot, and record an architecture decision and requirements trace.
2. Define interface-neutral installation engine types: mode, platform/state detection, phases/subphases and weights, requested-input schemas, plan/actions, consent, mutation ledger, retry/failure/restoration, readiness, and completion. Specify invariants proving truthful progress and no pre-consent meaningful mutation.
3. Implement localization infrastructure with stable IDs, canonical English fallback, Hungarian and German catalogs, placeholder/catalog validators, safe dynamic interpolation, and language selection before other interactive content.
4. Implement terminal rendering/input adapters: accessible guided dashboard for capable terminals and deterministic line-mode fallback, contextual field guidance, optional detailed view, secure secret input, validation/retry, confirmations, and interrupt/resume behavior.
5. Integrate preflight/platform registry, fresh/existing/upgrade classification, understandable plan, fixed installation paths/package transaction, configuration/state initialization, systemd user service/lingering, optional notification setup/test, update policy, activation, readiness, and completion through the shared engine.
6. Design and implement the smallest transparent bootstrap/distribution path that verifies canonical artifacts and launches the QWSG-owned wizard, reusing Task 064 trust code or package verifier rather than duplicating it; document download-then-verify-then-run as the primary method.
7. Add tests at component and journey levels for all languages, fallback, progress truth, unsupported platform, fresh/existing states, no-mutation planning, consent, credential non-disclosure, optional skips, update-policy honesty, service/readiness, injected failures, restoration summaries, terminal fallback, bootstrap trust, and future adapter boundaries.
8. Update architecture/operator/security/install/troubleshooting/release/changelog documentation and development identity only as justified by candidate convention. Run proportional Framework, Go, race, vet, formatting, shell, systemd, packaging, localization, privacy, rollback, transcript, and Git checks; diagnose and correct in-scope failures.
9. If justified, construct reproducible private candidate bytes without tag/publication. Run one bounded real-user fresh-install acceptance on an available disposable supported host; otherwise preserve all local acceptance and stop only at the exact Owner-controlled clean-host boundary. Do not disturb the preserved official-1.1.0 VPS baseline.
10. Reconcile evidence and limitations, perform explicit-path staged review, clean-fast-forward commit/push and direct ref verification, finalize Task 065 history, archive it transactionally to canonical idle, and do not create Task 066.


## Rollback Plan

- Local source/documentation/lifecycle rollback uses only verified Task 065 snapshots after exact identity, mode, collision, and unrelated-change checks. Restore only literal Task 065 targets; never use broad reset/checkout/clean and never touch excluded Owner content.
- Installer tests and simulations mutate only unique private temporary install/config/state/service roots and clean exact recorded paths. Bound processes and retain sanitized failure evidence before cleanup.
- Product installation uses Task 064- compatible verified package transaction/restoration for immutable artifacts; configuration, protected credentials, and persistent state remain user-owned. The engine records whether mutation began and reports automatic restoration result; never claim restoration without verification.
- External acceptance requires an exact pre-state ledger and recovery path. Use product-owned rollback/restoration or disposable-host reset only within explicit host authority. Never reinstall/reformat the preserved 1.1.0 VPS or manually manufacture PASS. Frozen private candidate bytes, if any, are immutable after selection; a source defect produces a newly identified candidate.
- A post-commit or external publication rollback is not implied. Do not rewrite shared Git history, move/delete tags, alter Releases, or reverse external state without separate authority.


## Deliverables

- Documented shared guided-installation engine and phase/progress/failure contracts.
- Professional terminal wizard with capability-aware dashboard and clean fallback.
- English, Magyar, and Deutsch installer catalogs with canonical English fallback and validation.
- Extensible platform detection supporting only Ubuntu 24.04 LTS amd64 and localized graceful refusal elsewhere.
- Preflight-first fresh/existing/upgrade classification and understandable consented installation plan.
- Integrated installation, configuration, protected optional notification credentials, update policy, Guardian activation, lingering guidance, readiness, failure/restoration, and localized completion journeys.
- Minimal transparent bootstrap/distribution workflow reusing Task 064 trust/integrity architecture.
- Reusable interface-neutral boundaries for future concise, unattended, and fleet adapters without duplicated business logic.
- Deterministic security/localization/progress/transcript/install/bootstrap/failure tests plus proportional local and bounded external user acceptance evidence where an authorized clean host is available.
- Updated product and engineering documentation, private development candidate identity as appropriate, validated Git integration, completed Task 065 record, and canonical idle closure without Task 066.


## Verification

- Framework 2.0, Builder input/install, active-job/lifecycle/test-task, repository identity, excluded Owner path, snapshot integrity, targeted staging, and rollback-readiness checks PASS.
- Engine tests prove explicit deterministic phases/weights, percentage derived only from state, no premature completion, retry truth, failure/mutation/restoration distinctions, stable plans, and adapter-independent business behavior.
- Catalog tests prove stable IDs, English completeness/fallback, Hungarian/German coverage, placeholder compatibility, safe interpolation, no scattered wizard prose in engine logic, and localized phase/help/error/summary journeys.
- Terminal tests cover capable PTY dashboard, no-color/dumb/redirected fallback, narrow dimensions, interrupt/EOF, control-character sanitization, validation retries, optional skip, confirmations, and understandable contextual guidance for every requested value.
- Security tests prove secret input is not echoed and credential values never appear in argv, environment, logs, diagnostics, transcripts, summaries, JSON, task evidence, or world/group-readable files; protected storage ownership/modes and unsafe path/symlink refusal PASS.
- Platform and plan tests cover exact Ubuntu 24.04 amd64 support, localized detected/supported refusal elsewhere, fresh/existing/upgrade ambiguity, discoverable defaults, no meaningful pre-consent mutation, technical-details view, and exact intended privileged actions.
- Installation harness proves package verification, fixed destinations, configuration/state initialization, systemd user unit and lingering behavior, optional notification configuration/test, update-policy persistence, Guardian activation, readiness, repeat/resume, failure explanation, and verified restoration while preserving existing data.
- Bootstrap tests prove canonical HTTPS/asset identity, bounds, checksum/archive/manifest/provenance verification, transparent download-and-run flow, understandable failure, and no mandatory pipe-to-shell or parallel trust system.
- Existing qwsg update check/update/status/rollback tests remain coherent; manual policy works; notification policy is truthful; automatic policy is absent/disabled unless fully implemented and validated with explicit consent.
- Full Go tests, focused and proportional race tests, go vet, gofmt, shell syntax, systemd static checks, packaging/reproducibility, documentation/link/terminology consistency, privacy scans, git diff --check, staged allowlist, and final Git/ref/lifecycle checks PASS.
- Real-user acceptance, when an authorized disposable host exists, proves bootstrap entry, language selection, guided progression, truthful progress, contextual inputs, credential safety, plan/consent, actual installation, Guardian activation, readiness, update integration, practical failure behavior, completion summary, and resulting operator usability. Missing external-host evidence is never converted to PASS.


## Documentation Updates

- Add or update canonical guided-installer architecture, phase/progress/localization/bootstrap/security/update-policy and reusable-engine contracts without creating competing sources of truth.
- Rewrite canonical installation/operator flow and packaged README/INSTALL/administration/troubleshooting guidance around the guided entry point while retaining expert commands and Task 064 update/rollback coherence.
- Update CHANGELOG.md, VERSION/release notes only according to established private-candidate convention, docs security/privacy and supported-platform material, and add privacy-safe fresh-install acceptance evidence if performed.
- Maintain ai/history/065_2026-08-27_guided-installation-setup-experience.md throughout execution, add a concise Engineering History milestone, archive the completed prompt transactionally, and record exact commits/ref verification/canonical lifecycle state.
- Preserve official 1.1.0, Task 064 acceptance/candidate evidence, and excluded Owner content unchanged.


## Completion Criteria

Task 065 is complete only when a new supported-host operator can use a secure transparent entry point and QWSG-owned localized wizard to understand every meaningful stage and requested decision, review the pre-mutation plan, install and configure QWSG, safely skip or configure optional notification, select an honest update policy, activate Guardian, verify readiness, understand any failure/restoration, and finish with a localized actionable summary; progress is demonstrably truthful; English, Hungarian, and German plus English fallback pass; the shared engine supports future adapters without duplicated installation logic; deterministic local/harness acceptance passes; proportional real-user clean-host acceptance passes when an authorized disposable host is available, or the task stops truthfully at the inherent Owner provisioning boundary; documentation, rollback, privacy, Git integration, and lifecycle closure pass. No release/publication, official-1.1.0 mutation, Server Health / Assessment Engine work, or Task 066 is permitted. A genuine boundary yields blocked with preserved evidence; ordinary defects remain in the Task 065 iterative engineering loop.


## Owner Approval Requirements

Approved by Project Owner through the Engineering Task Builder on 2026-08-27 UTC.

The structured task definition and Authority Envelope have been explicitly approved. Framework 2.0 Standard Execution Authority permits iterative, reversible in-scope engineering without another Owner gate. Further scope changes, exceptional external actions, and Owner-reserved decisions require explicit Project Owner approval.
