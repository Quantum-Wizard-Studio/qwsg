# Current Engineering Task 045: QWSG Setup & Configuration Foundation

## Task Metadata

- Task ID: `045`
- Task slug: `qwsg-setup-configuration-foundation`
- Status: `complete`
- Date opened: `2026-08-12` UTC
- Human authority: Project Owner
- Owner or lead-developer communication language: English

## Title

QWSG Setup & Configuration Foundation


## Objective

Establish one safe, deterministic, documented local configuration and setup foundation for QWSG. A clean-account operator must be able to discover, create, inspect, validate, and deliberately modify the canonical per-user configuration, then activate and verify the existing non-root Guardian through a coherent documented path. The implementation must reuse the accepted Configuration Contract 1.0 as its semantic core, add a secure filesystem activation and persistence boundary, resolve effective runtime configuration with explicit provenance and precedence, integrate Guardian fail-safely, and preserve all accepted QWSG 1.0 local Community behavior. The result must remain usable without an account, API key, subscription, Internet access, or QWS remote service, and must provide stable extension points for later Community email notifications and opt-in Pro/API capabilities without implementing them.



## Scope

- Re-audit and record every configuration source and consumer: Configuration Source Records and Effective Configuration, environment variables, CLI flags, Guardian scheduler and recurrence settings, systemd user-unit environment, state-root and store resolution, compiled defaults, precedence, validation, path/permission boundaries, packaging samples, install/setup instructions, and duplicated or contradictory concepts.
- Preserve `internal/configuration` Configuration Contract 1.0 as the single semantic model unless an evidence-backed defect requires the smallest compatible amendment. Do not introduce an independent setup schema or hidden UI configuration universe.
- Add a bounded per-user configuration activation/store layer that discovers the canonical XDG-based configuration location, distinguishes configuration from runtime state, reads deterministically, rejects unsafe paths and files, and writes atomically.
- Define the canonical default location as `${XDG_CONFIG_HOME}/qwsg/config.json` when `XDG_CONFIG_HOME` is a clean absolute path, otherwise `${HOME}/.config/qwsg/config.json`; define an explicit `--config` path for bounded command/automation use. Retain `--config-source` only as a compatibility alias if repository evidence proves it is required, with conflicts rejected rather than guessed.
- Define one explicit precedence model: compiled defaults, then the discovered primary local Source Record, then an explicitly selected configuration file replacing discovery, then narrowly supported command-temporary overrides represented as Configuration Sources. Environment variables that remain operational inputs must be documented separately and must not silently outrank canonical model fields. Equal-precedence conflicts remain fail-closed.
- Implement deterministic effective-configuration resolution and provenance display. Reconcile Guardian cadence and cycle timeout with the effective configuration so runtime behavior and reported configuration cannot disagree.
- Define schema/version, unknown-field, malformed-file, unsupported-version, compatibility, and migration behavior. Unknown fields, invalid identities, conflicting sources, unsupported major versions, trailing data, and materially invalid values must fail closed with bounded actionable diagnostics. No silent rewrite or lossy migration is allowed.
- Enforce configuration-directory mode `0700` and file mode `0600`, current-user ownership, ordinary regular files, clean absolute paths, rejection of symlink components and special files, bounded size, no group/other writability, and race-aware revalidation at use. Do not weaken existing private state-root rules.
- Implement atomic configuration persistence using a same-directory private temporary regular file, complete write, file sync, atomic rename, and parent-directory sync. Preserve the prior valid file on interruption or failure and reject unsafe replacement targets.
- Introduce the smallest coherent CLI surface consistent with the existing command dispatcher: `qwsg setup`, `qwsg config show`, `qwsg config validate`, `qwsg config get <key>`, and `qwsg config set <key> <value>`. Limit mutable keys to a documented typed allowlist owned by Task 045; reject arbitrary JSON paths and unknown keys.
- Make `qwsg setup` safe for first use and repeat use: detect absence versus existing valid configuration, show defaults and destination, validate proposed changes, show a final summary before persistence, preserve valid existing values unless explicitly changed, refuse unsafe storage, and produce clear next steps. Interactive operation requires a terminal; `--accept-defaults` and bounded `--set key=value` or equivalent provide deterministic noninteractive automation.
- Provide human-readable output and stable JSON output for show/validate/setup results where appropriate, using existing localization/presentation conventions and bounded diagnostics. Never render raw credential values.
- Integrate Guardian with automatic canonical configuration discovery and explicit override support. Resolve and validate configuration before mutating checkpoint/lifecycle state. A materially invalid configuration must prevent a healthy claim, produce an actionable bounded diagnostic, and avoid an uncontrolled systemd restart loop.
- Preserve the systemd user-service sandbox, non-root execution, resource limits, start limiting, private runtime state, and no-network behavior. Add a validation precondition or equivalent bounded fail-safe only if tested systemd semantics prove it prevents invalid-configuration restart churn without masking failure.
- Preserve the division between system-wide package installation, per-user setup/configuration, per-user runtime state, and explicit Guardian activation. Keep `install.sh` noninteractive; update its successful next-step guidance rather than making it configure or start a user service.
- Define and document a public/non-secret configuration boundary. Configuration may contain typed secret references only. Reserve a separate private credential-provider/store boundary under the per-user QWSG configuration domain for future work, but implement no credential backend and accept no real credential.
- Keep future extension fields or documented extension points for notification enablement, operator-owned SMTP, an ordered notification-recipient collection, severity/policy preferences, API endpoint, API credential reference, entitlement metadata, and remote opt-in inactive and non-required. The model must be able to represent one Community administrator email address and multiple future Pro recipients without a schema redesign. Task 045 must neither send email nor enforce recipient entitlement. Unsupported future functionality must not affect local Community operation.
- Record the Owner-approved notification product rule precisely: basic local email notification is a Community capability requiring no QWS account, API key, subscription, Internet-dependent QWS service, or paid entitlement; Community may configure exactly one administrator notification email address; Pro has no product-entitlement limit on notification recipients. Global parser, file-size, and runtime resource-safety bounds still apply. The entitlement boundary is recipient cardinality, not access to basic local email notification.
- Define and execute real clean-host-style acceptance using a released-style staged package install in an isolated clean HOME/account environment and, where the engineering host permits, the existing user-systemd acceptance method. Truthfully distinguish isolated acceptance from a real separate physical host.
- Update only the minimum canonical architecture, Setup, Configuration Reference, Quick Start, Operations, Troubleshooting, Security/Privacy, Upgrade, packaging, CLI, and Task 045 lifecycle documentation necessary to eliminate contradictory operator instructions.
- Preserve the immutable `v1.0.0` tag, release-source commit `177535e44b2ce5ed9efd73ab0793ffe6881f0cd6`, published artifact and SHA-256 `edfba7366adf2c1ce0a8ce56369bb0dc5ad11326c4e3d1e301625a5313292fa5`, and unrelated Owner draft `docs/architecture/QWCS_MIGRATION_BLUEPRINT.md` byte-for-byte.



## Out of Scope

- No final email notification subsystem, SMTP transport, message delivery, recipient dispatch, notification queue, notification templates, email-address validation beyond any future-model structural contract, or external account registration.
- No Community one-recipient enforcement, Pro unlimited-recipient enforcement, entitlement lookup, plan detection, or licensing decision in Task 045. The accepted rule is documented and made representable only.
- No hosted QWS notification delivery, QWS cloud service, Pro backend, API implementation, API-key validation, entitlement enforcement, subscription, billing, fleet management, remote administration, browser administration panel, or public release automation.
- No credential value, SMTP password, API key, token, private key, or real secret in general JSON, tests, fixtures, logs, diagnostics, snapshots, or documentation.
- No root requirement for normal QWSG use, arbitrary shell execution, uncontrolled network activity, systemd sandbox weakening, or change to the accepted local Community guarantee.
- No interactive behavior in `install.sh`; no implicit user-service enable/start/restart by setup; no hidden configuration outside the canonical model.
- No arbitrary generic JSON mutation language. `config set` is limited to the explicitly documented Task 045 typed key registry.
- No automatic migration that silently changes semantics, no acceptance of unknown fields for convenience, and no fallback to stale configuration while claiming healthy operation after material validation failure.
- No modification, retagging, rebuilding, republishing, or replacement of QWSG 1.0.0 source, tag, LICENSE, archive, checksum, Forgejo Release, or historical evidence.
- No Task 046 creation or implementation. Community Email Notifications may be recommended as the next separately Owner-authorized task only after Task 045 completion.
- No broad Git cleanup, staging, commit, push, tag, or publication authority beyond any later explicit Task 045 lifecycle gate. Task installation itself grants no implementation or Git authority.
- No deletion, relocation, chmod, ownership change, staging, or content inspection beyond what is necessary to hash/preserve the unrelated Owner draft.



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

- Verify canonical idle: no active prompt; Task 044 is the unique latest complete archive/history pair; Framework, lifecycle, diverted-task, test-task, Builder, and `bin/job --check` validations pass.
- Verify repository root, current UTC date/user, branch `main`, exact HEAD and `origin/main` both `c427bee165db17bb9416a46b44965bb54ff530a6`, `0` ahead and `0` behind, canonical origin, empty Git index, and clean tracked working tree.
- Record all untracked and ignored paths. Verify the only preserved unrelated untracked Owner path is `docs/architecture/QWCS_MIGRATION_BLUEPRINT.md`, record its pre-task hash/mode/owner/ACL, and do not read or alter its substantive content.
- Verify annotated `v1.0.0` remains unchanged, peels to `177535e44b2ce5ed9efd73ab0793ffe6881f0cd6`, and the published/local retained artifact identity remains SHA-256 `edfba7366adf2c1ce0a8ce56369bb0dc5ad11326c4e3d1e301625a5313292fa5` where local evidence exists.
- Read the mandatory governance plus Task 044 archive/history, configuration architecture and tests, CLI dispatcher and tests, Guardian service/checkpoint/scheduler code and tests, operator state/store/presentation code, packaging/install/uninstall/user-unit files, Quick Start/Operations/Troubleshooting/Security/Upgrade/release documentation, and relevant Task 039–044 histories.
- Record the exact current configuration audit before design: Source Record and Effective Configuration semantics, every environment variable and CLI flag, defaults, precedence, consumers, state/store roots, schedule sources, systemd access, path checks, ownership/modes, error behavior, and operator steps.
- Run pre-change `make build`, focused configuration/Guardian/CLI/operator tests, `go test ./...`, bounded-cache race tests, `go vet ./...`, formatting, whitespace, security/static scans, packaging and lifecycle validations. Stop and report any unexplained baseline deviation.



## Snapshot Requirements

### Mandatory Engineering Safety Rule

Before any Task 045-owned repository modification, create a unique external mode-`0700` snapshot under `/tmp` containing a verified manifest and exact payloads or absence records for every intended source, test, packaging, documentation, and lifecycle target. The snapshot must also record Git branch/HEAD/origin/upstream/index/status, file hashes, modes, owners, ACLs, relevant tool versions, service/process state, and the immutable release identities. It must preserve every pre-existing file that may be changed and include a safe bounded restore procedure. Do not proceed if the repository root or target set cannot be identified safely.

- Record the Owner draft by path, hash, metadata, and explicit exclusion only; do not copy or expose its contents unnecessarily.
- Verify the snapshot manifest and restore mapping before implementation. Keep rollback targets explicit and refuse symlinks, special files, changed post-snapshot hashes, active writers, ambiguous ownership, or broad repository restoration.
- Before clean-host/service acceptance, create separate isolated state/config/HOME/install roots with recorded ownership/modes and cleanup boundaries. Never target the real home, root filesystem, or published release assets for destructive cleanup.
- Preserve the verified snapshot through completion and record its location and integrity in Task 045 history. Snapshot creation is not permission to modify excluded content.



## Risk Assessment

- **Configuration corruption or loss — critical:** same-directory atomic writes, fsync, guarded replacement, injected interruption/failure tests, and verified rollback preserve the last valid file.
- **Symlink/path/ownership attack — critical:** clean absolute XDG resolution, no symlink components, current-user ownership, private modes, regular-file and race-aware checks fail closed.
- **Split effective/runtime behavior — critical:** project every runtime cadence/timeout input through one Configuration Source resolution and test provenance against actual Guardian behavior.
- **Invalid-config restart churn or false health — critical:** validate before checkpoint mutation, emit a stable bounded configuration diagnostic, and test service start-limit/precondition behavior without sandbox weakening.
- **Backward-compatibility regression — high:** preserve Configuration Contract 1.0, built-in defaults, existing explicit flags where defensible, QWSG 1.0 observation/bootstrap/partial-evidence semantics, and provide explicit deprecation/conflict behavior.
- **Secret leakage — critical:** only typed opaque secret references in public configuration; no credential backend or raw-value rendering; scan fixtures, output, diagnostics, snapshots, and Git diff.
- **State/configuration conflation — high:** use XDG config for persistent operator intent and the existing private state root only for runtime evidence/checkpoints.
- **Uncontrolled scope growth — high:** close the command/key surface and defer notification delivery, remote services, licensing, fleet, and UI work.
- **Installer privilege confusion — high:** keep installation noninteractive/system-wide and setup per-user; never configure a invoking sudo user's home from `install.sh`.
- **Owner-content or release-history damage — critical:** hash and exclude the Owner draft; never change `v1.0.0`, its source commit, artifact, LICENSE, or published release.
- **Localization or automation lock-in — medium:** reuse presentation/localization boundaries and provide deterministic JSON/noninteractive interfaces.
- **Schema migration dead end — high:** reject unsupported major versions, define additive compatibility explicitly, and keep setup/CLI/UI as adapters to one contract.
- **Incorrect notification product boundary — high:** model recipients as a future collection capable of one or many entries, but defer cardinality enforcement; document that Community includes basic local email with exactly one administrator recipient and Pro has no entitlement-imposed recipient cap without gating the base capability. Retain independent global resource-safety bounds.



## Planned Work

1. Validate the exact idle/release/Git baseline, complete the configuration-source/consumer audit, run pre-change gates, and create the verified bounded snapshot.
2. Write the architecture decision for canonical configuration location, activation, precedence, provenance, strict parsing, compatibility, security, atomic persistence, state separation, and future secret references. Include the Owner-approved future notification-recipient cardinality rule and a representation that supports one or many recipients without activating or enforcing it. Resolve the schedule/runtime mismatch before editing code.
3. Add a focused configuration filesystem/store package with dependency-injected environment/filesystem boundaries where needed. Implement safe location resolution, absence detection, strict load, ownership/mode/symlink checks, bounded reads, deterministic save, and atomic replacement.
4. Add a setup application layer that produces a reviewable plan, preserves an existing valid Source Record, applies only explicit typed changes, validates the complete result, recalculates canonical identity deterministically, and separates plan from commit.
5. Add `setup` and the bounded `config show|validate|get|set` CLI surface with human and JSON presentation, terminal/noninteractive rules, actionable exit statuses, secret redaction, and no shell/service side effects.
6. Route Guardian discovery, explicit file selection, CLI temporary schedule overrides, and runtime recurrence through the canonical resolver. Resolve configuration before lifecycle/checkpoint mutation and add bounded invalid-configuration state/diagnostics.
7. Test the least invasive systemd integration. Retain all sandbox/resource directives; use a validation precondition only if direct tests prove correct no-loop and visible-failure semantics. Otherwise document and test bounded start-limit behavior.
8. Keep `install.sh` noninteractive. Update packaged sample/config documentation and install completion text to direct the operator to `qwsg setup`, review, explicit unit installation/activation, verification, and Console use.
9. Add exhaustive unit/integration tests for clean absence, defaults, first/repeated setup, explicit mutation, strict validation, unsafe files/paths, atomic failures, deterministic effective configuration, Guardian valid/invalid startup, Console diagnostics, and compatibility regressions.
10. Run isolated released-style clean-HOME/package acceptance through install, setup, config review, Guardian activation, verification, Console, restart/reboot-equivalent where applicable, and uninstall preservation. Record exactly which tests used real user systemd and do not claim a separate clean host unless actually used.
11. Update the minimum canonical architecture/operator/release/security/upgrade documents, Task 045 history, and Engineering History milestone. Remove or redirect duplicated instructions rather than copying divergent procedures.
12. Run focused/full/race/vet/format/security/whitespace/packaging/Framework/Builder/lifecycle/diversion/job validations, audit the complete diff and release/Owner exclusions, exercise bounded rollback validation, and close/archive Task 045 to canonical idle only after all criteria pass. Do not create Task 046.



## Rollback Plan

Before any commit, stop Task 045 processes and restore only the explicit Task 045 target paths whose current identities match the recorded implementation outputs, using the external verified snapshot. For newly created Task 045 paths, remove only individually enumerated regular files after verifying their hashes and parent directories; never use a repository-wide reset, clean, checkout, wildcard, or unresolved environment variable. Restore modes and ACLs from the manifest, then rerun configuration, Guardian, lifecycle, Git, Owner-draft, release-tag, and artifact-identity checks. Preserve investigation and failure evidence outside the repository. If an authorized local commit later exists, do not rewrite it automatically; report the commit and request Owner direction. Published `v1.0.0` objects and assets are immutable and never rollback targets.



## Deliverables

- Complete Task 045 configuration audit and architecture decision with a single source/precedence/location/security/compatibility model.
- Secure per-user configuration location and atomic filesystem store integrated with Configuration Contract 1.0.
- Deterministic effective-configuration resolver/provenance shared by setup/config commands and Guardian runtime.
- Safe rerunnable `qwsg setup` plus bounded `qwsg config show`, `validate`, `get`, and `set` interfaces with automation-capable JSON/noninteractive operation.
- Guardian configuration discovery and fail-safe invalid-configuration behavior without root, network dependency, false health, checkpoint corruption, uncontrolled restart, or sandbox reduction.
- Explicit public configuration versus future private credential-reference boundary; no credential backend or secret values.
- Documented future notification-recipient representation and product rule: one Community administrator address with basic local email available offline and account-free; no Pro entitlement-imposed recipient limit, subject to global safety bounds; no Task 045 delivery or entitlement enforcement.
- Noninteractive installer guidance and a coherent canonical operator path from installation through setup, review, explicit activation, verification, and Console use.
- Focused/full regression and isolated released-style clean-environment acceptance evidence.
- Minimal noncontradictory Setup, Configuration Reference, Operations, Troubleshooting, Security/Privacy, Upgrade, packaging, architecture, and lifecycle documentation.
- Completed Task 045 prompt/history/archive and canonical idle state, with Task 046 only recommended—not created—as the future Community Email Notifications milestone.



## Verification

- Verify no prior configuration: clean account/HOME with unset and valid XDG variants resolves deterministically; defaults are visible; no file is written until explicit setup confirmation.
- Verify first setup, `--accept-defaults`, noninteractive typed settings, review summary, resulting schema/identity, directory `0700`, file `0600`, current ownership, and effective output.
- Verify repeated setup is idempotent, preserves all unspecified valid values and extensions, requires explicit changes, and produces byte-identical output for an identical requested model.
- Verify `config show|validate|get|set` human and JSON contracts, documented exit statuses, typed key allowlist, locale handling, bounded errors, and no raw secret rendering.
- Verify malformed JSON, trailing data, unknown fields, unsupported versions, invalid stable identity, duplicate/equal-precedence conflict, excessive size/count/string/duration, and invalid values all fail closed.
- Verify relative/unclean paths, symlink file and every symlinked parent, nonregular files, wrong owner, group/other-writable file or directory, permissive modes, replacement races, and unsafe XDG/HOME inputs are rejected without modifying existing data.
- Inject failures before/during write, sync, rename, and directory sync; prove interrupted writes preserve the prior valid file and leave no accepted partial configuration.
- Verify deterministic canonical JSON, Source/Effective identities, provenance, precedence, explicit-file-versus-discovery behavior, CLI temporary overrides, and actual Guardian cadence/timeout agree.
- Verify Guardian with valid configuration starts and updates lifecycle normally. With invalid configuration it resolves before checkpoint mutation, never claims healthy, emits actionable bounded diagnostics, and cannot enter an uncontrolled restart loop.
- Verify Console presents invalid/unavailable Guardian configuration truthfully, remains read-only, does not trigger observe, and does not leak configuration or secret material.
- Re-run all relevant QWSG 1.0 inventory, snapshot, compare, drift, health, rule, policy, report, first-run bootstrap, partial evidence, later full evaluation, Guardian checkpoint/restart/reboot, resource/sandbox, store safety, Console, install/upgrade/rollback/uninstall preservation, and no-network Community tests.
- Run focused package tests, `make build`, `go test ./...`, bounded-cache `go test -race ./...`, `go vet ./...`, `make fmt-check`, static/security scans, `git diff --check`, package manifest/layout/license tests, systemd unit verification, Framework, Builder, lifecycle, diversion/test-task, and `bin/job --check` gates.
- Perform isolated released-style clean-HOME installation and per-user setup/activation verification. Validate archive/install inputs without modifying the published 1.0.0 artifact and report whether user-systemd or a separate clean host was actually exercised.
- Verify uninstall preserves per-user configuration and state by default and documents deliberate removal separately.
- Audit all diffs, modes, owners, ACLs, generated files, snapshots, logs, fixtures, and diagnostics for secrets, host-private data, arbitrary execution, network additions, scope expansion, and release-history changes.
- Verify the Owner draft hash/mode/owner/ACL and content remain unchanged; index contains only later explicitly authorized Task 045 paths; `v1.0.0`, release-source commit, artifact hash, LICENSE, and Forgejo release remain unchanged.
- Verify architecture and configuration compatibility tests prove a future ordered recipient collection can represent one and multiple recipients without changing the base Configuration Contract, while Task 045 exposes no working email delivery and performs no Community/Pro entitlement decision.
- Exercise the bounded rollback in an isolated copy, then verify the real lifecycle can close Task 045 complete/archive and return to empty-prompt canonical idle without creating Task 046.



## Documentation Updates

- Add one canonical Setup guide and one canonical Configuration Reference, or a single clearly partitioned canonical document if that removes duplication. Document commands, locations, precedence, defaults, schema/compatibility, validation, permissions, atomicity, automation, secret references, failure behavior, and next steps.
- Update `docs/release/QUICK_START.md`, `docs/OPERATIONS.md`, `docs/TROUBLESHOOTING.md`, `docs/SECURITY_AND_PRIVACY.md`, `docs/UPGRADE_AND_ROLLBACK.md`, `packaging/release/INSTALL.md`, the CLI reference, and configuration/Guardian architecture only where necessary; link to canonical material instead of duplicating it.
- Update the packaged example Source Record to exactly match setup-generated valid content, while stating that the documentation sample is not automatically activated.
- Update installer completion text and systemd guidance without making installation interactive or implicitly enabling/starting Guardian.
- Document future Community SMTP and managed Pro/API extension points as inactive typed references/opt-ins sharing this model. State that Community basic local email permits exactly one administrator recipient without a QWS account/API key/subscription/remote service, while Pro has no entitlement-imposed recipient limit subject to global safety bounds; do not document nonexistent delivery or entitlement enforcement as available.
- Record implementation, tests, real acceptance environment, deviations, rollback, Git state, unresolved issues, and delivery result in `ai/history/045_<UTC-date>_qwsg-setup-configuration-foundation.md`; archive the matching prompt only on genuine completion and add one concise Engineering History milestone.
- Preserve Task 044 records, all QWSG 1.0.0 release documents/artifacts/identity, LICENSE, and `docs/architecture/QWCS_MIGRATION_BLUEPRINT.md` unchanged.



## Completion Criteria

Task 045 is complete only when the initial canonical idle and immutable release baseline are verified; the configuration audit is recorded; one documented Configuration Contract 1.0-based location/precedence/provenance/security/compatibility model exists; the future recipient representation supports the Owner-approved one-address Community and unlimited-recipient Pro boundary without Task 045 email delivery or entitlement enforcement; safe deterministic loading and atomic persistence pass adversarial tests; setup and bounded config commands work interactively and noninteractively; Guardian uses the same effective configuration before lifecycle mutation and fails visibly without uncontrolled restart behavior; state, public configuration, and future secret material have explicit separate boundaries; install remains noninteractive; a clean isolated operator can follow the canonical install → setup → review → activate → verify → Console path; all focused/full/race/vet/format/security/packaging/systemd/lifecycle/regression gates pass; accepted local Community/no-network behavior and uninstall preservation remain; the Owner draft and immutable v1.0.0 source/tag/artifact/LICENSE/release are unchanged; rollback is verified; documentation and Task 045 history are complete; and Task 045 is archived with no active prompt in canonical idle. Any unsafe path behavior, data-loss case, secret exposure, split runtime/effective setting, false Guardian health, uncontrolled restart, incorrect Community/Pro notification boundary, release-identity change, unexplained regression, or unperformed mandatory acceptance keeps the task active or blocked with the exact reason. Task 046 is not created by Task 045 completion.


## Owner Approval Requirements

Approved by Project Owner through the Engineering Task Builder on 2026-08-12 UTC.

The structured task definition has been explicitly approved for implementation. Further scope changes require explicit Project Owner approval.
