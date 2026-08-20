# Current Engineering Task 049: QWSG 1.1.0 External Clean-Host Acceptance and Release Readiness

## Task Metadata

- Task ID: `049`
- Task slug: `qwsg-1-1-external-clean-host-acceptance-release-readiness`
- Status: `complete with disclosed limitations`
- Date opened: `2026-08-20` UTC
- Human authority: Project Owner
- Owner or lead-developer communication language: English

## Title

QWSG 1.1.0 External Clean-Host Acceptance and Release Readiness


## Objective

Produce a private, non-published, reproducible QWSG `1.1.0-rc.1` acceptance candidate from an exact integrated Git source commit, guide the Project Owner through a checkpointed first-use test on a disposable clean Ubuntu 24.04 amd64 VPS using only product and operator documentation, and issue an evidence-backed `READY FOR QWSG 1.1.0 RELEASE` or `NOT READY FOR QWSG 1.1.0 RELEASE` decision. Acceptance covers artifact provenance, documentation, Smart Install, guided setup and resume, real external SMTP delivery when Owner-provided service is available, explicit Guardian activation, fresh evidence, lingering/logout, physical reboot recovery, restart, uninstall, and mandatory reinstall/resume. Missing guidance is a product finding; engineering knowledge must never silently compensate. Task 049 does not publish, tag, upload, announce, or create the final 1.1.0 release.


## Scope

- Audit Tasks 043–048, current VERSION/Make/release scripts, RC/final acceptance records, reproducibility, embedded identity, manifest/sidecar, installer/uninstaller, README/INSTALL packaging, known limitations, systemd/Guardian evidence, SMTP credential boundary, release/Git policy, and Community/Pro boundary.
- Use repository precedent and select private candidate identity `1.1.0-rc.1`. It is a pre-release acceptance identity, never the published final release. Do not reuse a 1.0 RC identity or call the current `VERSION=1.0.0` binary a 1.1 candidate.
- Resolve the current build-tool limitation narrowly: extend release validation/notes/acceptance plumbing only as needed to support the 1.1 RC line. A candidate may be built only after those changes are integrated by separately explicit Git authority, so its source tree is clean and its exact commit is embedded. No uncommitted VERSION/build-script overlay may be represented as commit-pure provenance.
- Prepare but do not infer authority for the candidate source integration. If no exact integrated candidate commit exists, stop at the integration gate and request Owner authority before building/transferring the candidate.
- Build candidate twice in independent private roots with identical `SOURCE_DATE_EPOCH` derived from the exact source commit and exact full/embedded commit identity; prove byte identity. Record version, commit, epoch, archive filename, SHA-256 sidecar, internal manifest, binary `version`, layout, packaged LICENSE, README and INSTALL identities, modes, and difference from published v1.0.0.
- Keep candidate artifacts private under a unique mode-0700 `/tmp` evidence root or another Owner-approved private transfer location. Do not write them over `dist/qwsg-1.0.0-linux-amd64*`, upload to Forgejo Releases, create a tag, or announce them.
- Create a canonical Task 049 acceptance protocol and evidence record with numbered, restartable checkpoints. Each checkpoint states purpose, one bounded command/action block, expected evidence, PASS, FAIL/finding, continuation safety, and evidence to preserve. Commands target a normal operator and expose no internal architecture.
- The Owner executes all external VPS commands interactively. Aikó evaluates returned output and never claims external evidence from local fixtures. Do not access or mutate the VPS during Builder preparation.
- Enforce the documentation-only journey: obtain private candidate -> verify sidecar/manifest -> read root README/INSTALL -> install -> follow installer output -> Smart Install/readiness -> guided setup -> notification credential/test -> Guardian activation/evidence -> lingering/logout/reboot -> restart -> uninstall -> reinstall/resume. Any undiscoverable required step becomes a UX/documentation defect before manual assistance.
- Smart Install gates verify supported OS/version/architecture/systemd/user, satisfied/missing required/missing optional/unknown/incompatible states, only proven commands, manual-verification handling, no mutation, correct exits, and exact next action.
- Guided setup gates verify fresh setup, existing-state recognition, interruption/resume, no unnecessary repetition, invalid input, notification choice/configuration, protected credential workflow, controlled test, activation confirmation/decline, bounded evidence wait, final summaries and next actions. Existing automation is also smoke-tested.
- External SMTP is conditional on the Owner supplying a suitable provider and recipient. Never request credentials in chat or persist them in prompt/history/docs/logs/argv/Git. The Owner creates a private mode-0600 file locally on the VPS using hidden terminal input or provider-safe means, invokes the existing `credential set --from-file` boundary, then securely removes only the input file after store verification. Capture only redacted classifications, exit results, message metadata allowed by QWSG, and independent receipt confirmation. TCP/TLS/SMTP acceptance without recipient receipt is not delivery PASS.
- Real external SMTP delivery is mandatory for a full READY decision unless the Owner explicitly declares the provider gate unavailable. Provider unavailability is `CONDITIONAL GATE NOT EXECUTED`, not product failure, but the final decision cannot claim real-provider readiness and defaults to NOT READY for unconditional 1.1.0 publication until separately completed.
- Guardian/systemd checkpoints verify unit installed, daemon reload, enabled/active state, MainPID, invocation identity, restart count, bounded resources, QWSG process, configured cadence, canonical fresh evidence, explicit restart and recovery. ActiveState alone never passes Guardian health.
- Lingering checkpoints start with QWSG detection/guidance. QWSG never executes sudo. The Owner may apply only documented exact host-admin guidance, then revalidates. Test session logout where feasible and require a physical VPS reboot, automatic user-unit return, new process/invocation identity, bounded post-reboot cycle, fresh post-reboot canonical evidence, and notification continuity. Simulation is not PASS.
- Uninstall verifies explicit service stop/disable, binary/unit/docs removal, unchanged release-owned artifact enforcement, preservation of per-user configuration/credentials/state, no broad cleanup, and clear result. Reinstall/resume is mandatory because preserved state is a product promise: reinstall the same verified candidate, confirm safe recognition, reactivation/revalidation, and no false stale READY.
- Capture commands and output only after privacy review. Redact public IP/hostname, usernames where unnecessary, email addresses, SMTP host/account, credential references, tokens, invocation IDs when sensitive, inventory, paths and provider headers. Store no secret. Evidence differentiates local automated, staged, external host, physical reboot, and real-provider observations.
- Finding severities: `RELEASE BLOCKER` (mandatory gate/false readiness/destructive or irrecoverable failure), `SECURITY DEFECT` (secret/privilege/path/injection/privacy boundary), `FUNCTIONAL DEFECT` (documented behavior fails), `UX/DOCUMENTATION DEFECT` (operator cannot discover/complete safely), and `COSMETIC / POST-RELEASE CANDIDATE` (no material correctness/usability effect). Security defects and mandatory-gate failures block release.
- Task 049 may correct only acceptance tooling/protocol/evidence defects and narrowly scoped release-candidate plumbing identified before external execution. A product code, setup, readiness, installer, documentation, security, SMTP, systemd, reboot, or uninstall defect found during physical acceptance is recorded, acceptance stops where unsafe, final status becomes NOT READY, and correction requires a separately Owner-authorized follow-up task. No silent fix/retest loop.
- Final READY requires every mandatory gate: clean-host artifact verification/install, operator documentation journey, Smart Install, guided setup/resume, Guardian activation and fresh evidence, lingering guidance and unattended physical reboot recovery, post-reboot fresh cycle, uninstall, reinstall/resume, all local gates, no blocker/security defect, and real SMTP receipt. Conditional SMTP absence prevents unconditional READY.
- Produce a release-readiness report with exact candidate/source/artifact identities, checkpoint results, findings/severity, limitations, evidence labels, and one explicit final verdict. Passing does not authorize final VERSION/tag/Forgejo Release/upload/announcement.
- Community/Pro remains unchanged. Test one Community recipient and local Guardian; add no QWS account, API, entitlement, provisioning, managed notification, fleet or GUI behavior.
- Preserve v1.0.0, LICENSE, published artifacts/checksums/tag/source, Tasks 043–048, and `docs/architecture/QWCS_MIGRATION_BLUEPRINT.md` unchanged.


## Out of Scope

- No external VPS mutation during planning; no automated remote login/provisioning; no credential transfer through chat, Git, command arguments, logs, prompt/history, or acceptance documents.
- No public/final `v1.1.0` tag, Forgejo Release, upload, publication, signing claim, announcement, branch rewrite, or v1.0.0 modification.
- No automatic product correction during physical acceptance and no scope expansion from a finding. No Task 050 creation or implementation.
- No package/remediation/sudo automation, root provisioning, QWS API/key/account, billing, entitlement, Pro, managed email, fleet, GUI, role monitoring, analytics, AI repair, or unrelated features.
- No simulated evidence represented as external/physical/provider evidence; no READY from local tests alone or from ActiveState/TCP acceptance alone.
- No blanket Git staging, inferred commit/push authority, secrets, Owner-draft access/modification, destructive cleanup, or unbounded wait.


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

- Verify root/date/user, `main`, canonical origin, exact HEAD/origin `57d37206c221a85b83cccb974525bc3ba4e408c1`, `0/0`, empty index, clean tracked tree, canonical idle after Task 048, no Task 049 collision, and only the unchanged Owner draft visible untracked.
- Verify Framework/job/Builder/lifecycle/diversion/test-task suites; approved Builder source hash; v1.0.0 annotated object/target, LICENSE, published archive/sidecar and release evidence.
- Audit Tasks 043/044 clean-host/final-release precedent and Tasks 045–048 local/staged evidence; VERSION, Makefile, build-release, manifests, packaging, README/INSTALL, release notes/acceptance, known limitations, SMTP, systemd, readiness and security code/docs/tests.
- Record current 1.0-only release-script constraint, exact source/epoch embedding rules, clean-tree requirement, candidate collision guards, and mandatory/conditional external gates.
- Run pre-change build/full/race/vet/format/security/governance, deterministic twin builds, archive/layout/manifest/install/uninstall/docs tests, and static systemd checks. Stop on unexplained variance.


## Snapshot Requirements

Before Task 049 target changes or candidate preparation, create a unique external mode-0700 `/tmp` snapshot containing verified tracked HEAD, active lifecycle/Builder source, intended targets/absence records, Git/release/tool identities, v1.0.0 preservation, and bounded restore instructions. Record Owner draft metadata/hash only. Capture no host/private SMTP data. Verify checksums, archive readability, collisions, retention, and isolated rollback. Maintain a separate private candidate/evidence root with a manifest and exact cleanup/preservation policy.


## Risk Assessment

- **Secret/privacy exposure — critical:** Owner-only protected-file injection, redacted evidence, no chat/argv/Git/provider headers.
- **False release readiness — critical:** mandatory physical/provider gates, explicit evidence labels, no local/simulated substitution.
- **Operator coaching masks defect — high:** documentation-only journey; every undocumented intervention is a finding.
- **Candidate provenance/version ambiguity — critical:** `1.1.0-rc.1`, clean exact integrated commit, controlled epoch, twin-build identity, no overlay claim.
- **Published release damage — critical:** private roots, collision/hash/tag guards, no public action or v1.0.0 mutation.
- **External host damage — high:** disposable host, checkpoint continuation safety, no automation, bounded exact commands, rollback/uninstall.
- **False Guardian/reboot evidence — critical:** canonical freshness plus physical reboot/new invocation, not ActiveState.
- **Acceptance repair scope creep — high:** only protocol/release-plumbing corrections; product findings stop and defer.
- **Environment/provider variability — high:** classify unavailable external service separately without claiming PASS.


## Planned Work

1. Verify baseline, audit precedents/capabilities, run pre-gates, snapshot, and define mandatory/conditional evidence matrix.
2. Add narrowly version-general release-candidate plumbing and `1.1.0-rc.1` notes/acceptance scaffolding without final/public semantics.
3. Run full validation and obtain separately explicit source-integration authority; record the exact integrated candidate commit before artifact build.
4. Build two independent private candidates from that commit/epoch; prove identity, provenance, layout, docs/LICENSE/manifest/sidecar and v1.0.0 separation.
5. Produce the numbered Owner-operated clean-host protocol, redaction/evidence templates, severity table, safe-continuation rules and readiness matrix.
6. Transfer the private artifact only through an Owner-approved private channel and begin checkpoint execution one block at a time.
7. Evaluate documentation, install/Smart Install, guided setup/resume, real SMTP receipt, Guardian/systemd, lingering/logout/physical reboot, restart, uninstall and mandatory reinstall/resume.
8. Stop and classify every finding; perform no physical-acceptance product fix. Continue only where protocol marks safe.
9. Run final local/reproducibility/security/preservation/rollback/governance gates and issue exact READY or NOT READY report.
10. Complete/archive Task 049 to canonical idle without Task 050 or publication.


## Rollback Plan

- Stop only exact acceptance processes/services on the disposable VPS using the documented checkpoint rollback; preserve redacted evidence before destructive test-host teardown.
- Locally verify snapshot/candidate manifests and collisions. Restore only literal pre-task targets and remove only Task 049-created paths with proven absence/hash/ownership; never reset/clean or touch v1.0.0/Owner data.
- Candidate artifacts remain private until Owner-directed secure disposal or later release task; their retention does not authorize publication. Re-run all local, package, security, lifecycle and preservation gates after rollback simulation.


## Deliverables

- Repository/release audit and mandatory/conditional gate matrix.
- Clean-commit `1.1.0-rc.1` private candidate with twin-build reproducibility and complete provenance evidence.
- Numbered Owner-operated protocol covering the full Community journey, physical reboot, uninstall and reinstall.
- Privacy-safe SMTP injection/receipt procedure and redacted evidence templates.
- Finding register with severity, continuation decision and exact blocker policy.
- Final release-readiness report with explicit READY or NOT READY verdict and no publication action.
- Snapshot, rollback simulation, local/staged/external evidence separation, and completed lifecycle history.


## Verification

- Baseline/final Git/lifecycle/Owner/release hashes and no tag/publication checks.
- Build/full/race/vet/format/security; Framework 21, diversion 36, lifecycle 28, Builder 38; test-task/job; shell/systemd; deterministic twin candidate build.
- Candidate checks: exact clean commit/epoch/version/archive name, two-build SHA equality, sidecar, internal manifest, single-root safe layout, binary embedded version/commit, packaged LICENSE/README/INSTALL identity/cross-links/modes, install/uninstall and collision/backup behavior, distinction from v1.0.0.
- Protocol validation proves every checkpoint has purpose/action/expected/PASS/FAIL/safe-continuation/evidence, no secrets, bounded commands/waits, interruption/resume, and operator-readable instructions.
- External mandatory evidence covers clean Ubuntu 24.04 amd64/no prior QWSG, documentation-only journey, Smart Install, guided setup/resume/invalid input, protected credential/preflight/real receipt, explicit activation, systemd/resource facts, fresh evidence, lingering guidance, logout, physical reboot/new invocation/post-reboot cycle, notification continuity, restart, uninstall preservation, reinstall/resume.
- Security audit proves no credentials/private host data, automatic sudo/package/remediation, generic execution, listener, public upload/tag/release, Owner/release mutation, or unsupported claims.
- Final decision matrix rejects READY for any blocker/security defect, mandatory failure, missing physical reboot/fresh evidence/uninstall/reinstall, or missing real SMTP receipt; local tests alone never satisfy it.


## Documentation Updates

- Add `1.1.0-rc.1` private engineering acceptance notes/protocol and final readiness report; update release tooling documentation only as required for version-general RC construction.
- Reference existing README/INSTALL/Quick Start/Setup/Operations/Troubleshooting/Security/Uninstall rather than rewriting operator truth during acceptance. Record every gap as a finding.
- Record candidate provenance, checkpoint outputs in redacted form, evidence categories, findings, limitations, rollback, preservation and the separate final-publication boundary.


## Completion Criteria

Task 049 is complete only when the Task 048 idle baseline is verified; a private `1.1.0-rc.1` candidate is reproducibly built twice from one exact clean integrated commit with complete manifest/sidecar/binary/docs/LICENSE provenance and no v1.0.0 collision; the Owner executes the checkpointed documentation-only journey on a disposable clean Ubuntu 24.04 amd64 VPS; Smart Install, guided setup/resume, protected credential handling, real external email receipt, explicit Guardian activation, fresh evidence, lingering/logout, physical reboot/new invocation/post-reboot fresh cycle, restart, uninstall and reinstall/resume are truthfully evidenced; all local/security/governance/rollback gates pass; every finding is classified without silent repair; and the final report states READY or NOT READY under the mandatory matrix. Completion never authorizes final tag, public release/upload/announcement, Task 050, or v1.0.0/Owner changes. If an integrated candidate commit, safe Owner VPS participation, real SMTP receipt, physical reboot, mandatory evidence, or critical gate is unavailable, the truthful verdict is NOT READY or the task remains incomplete—not an inferred PASS.


## Owner Approval Requirements

Approved by Project Owner through the Engineering Task Builder on 2026-08-20 UTC.

The structured task definition has been explicitly approved for implementation. Further scope changes require explicit Project Owner approval.
