# Current Engineering Task 051: QWSG 1.1.0-rc.2 Clean-Host Acceptance Retest

## Task Metadata

- Task ID: `051`
- Task slug: `qwsg-1-1-0-rc-2-clean-host-acceptance-retest`
- Status: `active`
- Date opened: `2026-08-21` UTC
- Human authority: Project Owner
- Owner or lead-developer communication language: English

## Title

QWSG 1.1.0-rc.2 Clean-Host Acceptance Retest


## Objective

Prepare one private, independently identifiable, reproducible QWSG `1.1.0-rc.2` candidate from an exact clean canonical Git commit; transfer only its archive and checksum sidecar through a separately Owner-approved private channel; and restart the external clean-host acceptance at Checkpoint 01 on a real supported Ubuntu 24.04 amd64 VPS. First prove the Task 050 corrections for Task 049 findings `QWSG-049-F002` and `QWSG-049-F003` from product-visible guidance without developer coaching, then, only while mandatory gates pass, complete the Community operator journey through SMTP receipt, Guardian activation, logout, physical reboot, restart, uninstall, reinstall, and resume. Preserve RC.1 and all Task 049 evidence unchanged. End with exactly `READY FOR QWSG 1.1.0 RELEASE` or `NOT READY FOR QWSG 1.1.0 RELEASE`; neither result authorizes publication.


## Scope

- Phase A — repository and RC.2 readiness: audit the canonical repository, Tasks 049/050, `VERSION`, release scripts/tests, README, INSTALL, RC.2 notes, packaging, license, acceptance precedent, Git identity, tags, lifecycle, and Owner-owned paths. Determine the exact candidate-source commit. Permit only narrow RC.2 acceptance/release-plumbing corrections needed before a commit-pure candidate; substantial product work triggers a prerequisite-task recommendation and stop.
- Treat the existing untracked `docs/architecture/QWCS_MIGRATION_BLUEPRINT.md` as Owner-owned: do not read, stage, copy, hash into public/task evidence, modify, move, delete, ignore, or package it. A dirty ambient worktree cannot be described as clean; candidate purity must be proven from exact committed exported source roots and an empty task index/target diff, with Owner-owned content excluded.
- Prepare RC.2-specific acceptance protocol and record paths without modifying or overwriting `docs/release/ACCEPTANCE_PROTOCOL_1.1.0-rc.1.md`, `docs/release/ACCEPTANCE_1.1.0-rc.1.md`, Task 049 records, or private RC.1 artifacts. Integrate all authorized Phase A source/lifecycle/scaffolding changes through a separately explicit Owner staging/commit/push gate before any build. Do not build candidate bytes from uncommitted overlays.
- Phase B — after Owner source-integration authorization and verification of one exact clean candidate commit, export that commit twice into independent private mode-0700 roots. Derive `SOURCE_DATE_EPOCH` from the commit timestamp and pass the exact lowercase 40-character commit as `BUILD_COMMIT`. Build twice and prove byte-identical binary, internal manifest, archive, and checksum sidecar; static Linux amd64; safe single-root layout; stable metadata; packaged LICENSE, root README.md, root INSTALL.md, RC.2 notes; and absence of secrets, credentials, private evidence, Builder inputs, snapshots, backups, Owner-owned content, or RC.1 collision.
- Record exact source commit, commit-derived epoch, `qwsg-1.1.0-rc.2-linux-amd64.tar.gz`, archive SHA-256, sidecar, internal manifest identity, binary version/build date/full embedded commit, package layout and modes. Keep candidate artifacts private and keep RC.1 untouched.
- Phase C — after a separate explicit Owner transfer gate, determine the safest bounded standard SSH/SCP procedure supported by both hosts without assuming credentials, destination, account, or host-key material. Prefer direct development VPS to Owner-approved disposable VPS transfer of only the archive and sidecar. Require strict normal host-key verification, pre-existing standard tools and Owner-controlled authentication; do not record private identifiers. If safe direct transfer is unavailable, use development VPS to Owner workstation to test VPS. Record privacy-safe provenance and post-transfer checksum.
- Phase D — restart acceptance at RC.2 Checkpoint 01. Provide exactly one bounded Owner-operated checkpoint at a time. Every checkpoint includes purpose, exact command/action, expected evidence, PASS, FAIL/finding, continuation safety, and retain/redact rules. Evidence is external only when physically executed by the Owner on the disposable supported host.
- Cover private receipt; checksum; layout; manifest; LICENSE; README/INSTALL usability; Smart Install and F002/F003; immutable install; guided setup interruption/resume and invalid input; one Community recipient; protected SMTP credential; preflight; real test delivery and independent receipt; explicit Guardian activation; fresh canonical evidence; systemd process/invocation/cadence/resource/restart evidence; lingering and logout; physical reboot; new process/invocation and fresh post-reboot evidence; post-reboot notification receipt; explicit restart; safe uninstall with preserved user state; reinstall/resume/reactivation with no stale READY.
- Product-visible guidance is the only operator coaching source: archive README/INSTALL, installer and Smart Install output, readiness, guided setup, and installed docs. Record any undiscoverable required action as a finding before assistance.
- Record privacy-reviewed RC.2 provenance, transfer provenance, checkpoint results, F002/F003 external retest result, SMTP receipt confirmation, Guardian/logout/reboot evidence, uninstall/reinstall evidence, findings, and final verdict in new canonical RC.2 records. Historical RC.1 evidence remains immutable.
- Community/Pro boundary remains unchanged: one Community recipient and local Guardian only; no QWS account, API, entitlement, managed notification, fleet, GUI, or Pro behavior.


## Out of Scope

- No Task 051 installation, implementation, build, transfer, VPS access, credentials, staging, commit, push, tag, release, upload, publication, or announcement occurs during Builder preparation. Once installed, each later phase still requires its stated Owner gate.
- No final `v1.1.0` tag, Forgejo Release, public artifact, announcement, signing claim, or final publication. A complete acceptance PASS stops at a separate Owner release gate.
- No overwrite, relabel, rebuild, transfer, modification, or reinterpretation of RC.1 or Task 049 evidence. F002/F003 historical findings remain unchanged even if RC.2 externally verifies the corrections.
- No SSH credential request or storage; no credentials, destination identifiers, private host/account data, or host keys in Git, chat, task history, or acceptance records. No SSH server changes, weakened host-key checking, newly installed transfer software, public exposure, or files beyond the RC.2 archive and sidecar.
- No automatic sudo, package installation, lingering, reboot, arbitrary remediation shell, external-host automation, or unbounded command/wait. The Owner operates the disposable host and authorizes privileged actions separately at the relevant checkpoint.
- No SMTP credential through chat, argv, Git, task/evidence records, logs, or provider output. No raw recipient, SMTP account/host, provider headers, credential references, tokens, or private host identities in retained evidence.
- No silent product repair during physical acceptance. Only exact narrow corrections explicitly authorized by Task 051 may occur before candidate-source integration. A new product/security defect outside that authority stops acceptance and requires a separate correction task.
- No broad staging, reset, clean, checkout, deletion, or modification of unrelated/Owner-owned content, including `docs/architecture/QWCS_MIGRATION_BLUEPRINT.md`.


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

- Verify UTC date, ordinary user, repository root, `main`, canonical HTTPS `origin`, HEAD and `origin/main` at Owner-approved baseline `af8154140ba934cfa0b89aa7071633c87323ecb9`, ahead/behind `0/0`, empty index, canonical idle after complete Task 050, zero active prompts, no Task 051 collision, and no `v1.1.0` or RC.2 tag/release. Fetch only when separately authorized; stop on unexplained identity or divergence.
- Record the known planning variance: the ambient worktree is not clean because `docs/architecture/QWCS_MIGRATION_BLUEPRINT.md` is untracked Owner-owned content. Preserve it untouched and excluded. Stop on any additional unexplained tracked, staged, untracked, ignored-sensitive, or lifecycle difference.
- Verify `VERSION=1.1.0-rc.2`; RC.2 release notes and known limitations; release validation accepts RC.2 and rejects collisions/missing commit-derived metadata; Task 050 product commit `6fd9d65dd0690ac7409200dc3fd949e9be14e981` is an ancestor of lifecycle-closing commit `af8154140ba934cfa0b89aa7071633c87323ecb9`; and no RC.2 artifact exists in repository release outputs.
- Verify immutable RC.1 provenance: source `ff2eb2b12499f5daf3b5ba11b1f8d7fc562f8a31`, archive SHA-256 `aa139faaccc1cc85b50cfe0eedee9436539ae1c3071e01d8e9ed9283fc7f8239`, Task 049 F001/F002/F003 records, and stopped Checkpoint 04 state. Verify v1.0.0 tag/release identities and LICENSE preservation.
- Read Tasks 049/050 prompts and histories; RC.1 acceptance record and 16-checkpoint protocol; RC.2 notes; release/build/package/install/uninstall plumbing; README/INSTALL and referenced operator documents; Smart Install model/registry/runner/CLI tests; setup, SMTP, Guardian, systemd, state/evidence, privacy and release policies.
- Run pre-change build, focused/full/race tests, vet, formatting, release-check, shell syntax, static systemd, archive/package/install/uninstall tests, security/privacy scans, Git whitespace, Framework 21, Builder 38, lifecycle 28, diversion 36, test-task and job validation. Record exact versions and stop on unexplained failure.


## Snapshot Requirements

Before modifying any Task 051 target, create a unique mode-0700 `/tmp` snapshot containing a readable complete tracked-HEAD archive, exact Task 051 prompt/history and Builder source, intended target payloads or verified absence records, Git/status/diff/index/mode/ACL/tool identities, Task 049/050 and RC.1/v1.0.0/LICENSE preservation identities, and bounded restore instructions. Record the Owner-owned untracked path only as excluded metadata without copying or hashing its contents. Verify checksums, archive readability, collisions, retention and isolated rollback. Maintain distinct private mode-0700 candidate-build, transfer-staging, and redacted-evidence roots; never mix them with Builder inputs, snapshots, credentials, Owner content, or repository artifacts.


## Risk Assessment

- **False release readiness — critical:** require all mandatory real-host/provider/reboot/uninstall/reinstall gates and exact final decision matrix; local or simulated evidence never substitutes.
- **Candidate provenance ambiguity — critical:** integrate source first, use two independent commit exports, exact full commit and commit epoch, collision guards, and byte comparisons; no overlays.
- **Owner-content or secret disclosure — critical:** preserve/exclude the untracked Owner path; package allowlist and privacy scans; protected SMTP boundary and redacted evidence only.
- **RC.1 evidence damage or identity collision — critical:** immutable hash/path checks, RC.2-specific names and roots, no overwrite/relabel.
- **Operator coaching masks a defect — high:** product-visible guidance only; record a finding before any intervention.
- **Unsafe transfer — high:** separate Owner gate, standard pre-existing SSH/SCP only, strict host-key verification, two-file allowlist, private fallback through Owner workstation.
- **External host damage/privilege escalation — critical:** one bounded checkpoint at a time, Owner execution, no automatic sudo/packages/lingering/reboot/remediation, explicit continuation safety.
- **False F002/F003 closure — critical:** external RC.2 evidence required; validate evidence, cause, privilege, safe command certainty, manual procedure, no mutation, and revalidation.
- **SMTP false positive or leak — critical:** independent receipt required before and after reboot; TLS/auth success is insufficient; no secret/private/provider evidence retained.
- **Guardian stale evidence — critical:** correlate systemd facts with canonical fresh evidence and cadence; require physical reboot, new invocation/process identity, post-reboot freshness, restart, and notification continuity.
- **Acceptance scope expansion — high:** narrow pre-build plumbing only; substantial product correction or newly found out-of-authority defect stops for a prerequisite/follow-up task.
- **Lifecycle/Git corruption — critical:** canonical Builder, snapshots, targeted staging, separate explicit integration authority, clean exported commits, no publication inference.


## Planned Work

1. Gate A1: validate baseline, governance, immutable identities, Owner-owned exclusion, release plumbing and Task 049/050 evidence; run pre-gates and create/verify the bounded snapshot.
2. Audit whether RC.2 is commit-ready. Prepare only RC.2-specific acceptance protocol/record scaffolding and narrow release/document corrections if required. If substantial product work is required, record `NOT READY`, recommend a prerequisite task, and stop.
3. Gate A2 — Owner source integration: present the exact reviewed Task 051 path allowlist, diffs, validations and proposed commit. Do not stage/commit/push without separate explicit authority. After integration, verify one exact candidate-source commit, canonical remote relationship, empty index, no tracked diff, and Owner-content exclusion.
4. Gate B — Owner candidate construction: from two separately created private mode-0700 exports of the exact commit, derive its committer epoch, build RC.2 twice with identical controlled inputs, compare binary/manifests/archives/sidecars byte-for-byte, audit package contents/layout/modes/static linkage/embedded identity and exclusions, and record exact provenance. Never overwrite RC.1.
5. Prepare the RC.2 checkpoint protocol by adapting proven Task 049 structure while restarting at Checkpoint 01 and splitting actions where needed for safe one-step Owner operation. Preserve empty result fields until real evidence exists.
6. Gate C — Owner transfer: inspect available standard transfer capabilities without credentials; propose literal source/destination file allowlist and strict host-key procedure. After explicit approval, transfer only archive/sidecar direct VPS-to-VPS; otherwise use Owner-workstation fallback. Verify receipt hash and record redacted provenance.
7. Gate D1 — Owner external start: establish clean supported host and execute RC.2 Checkpoints 01 onward one at a time. Verify receipt, checksum, layout, manifest, LICENSE, README/INSTALL, then run `./bin/qwsg install --check` without developer coaching.
8. At the former failure point, evaluate Task 050 guidance for `systemd.user_manager` and `filesystem.local_semantics`. Stop on false classification, generic dead end, guessed/unsafe remediation, missing verification/privilege/revalidation, unexplained unknown, or mutation. Mark F002/F003 externally VERIFIED/CORRECTED only after complete PASS.
9. Gate D2 — Owner host-mutation continuation: only after pre-install mandatory gates pass and separate approval, install immutable artifacts and exercise guided setup interruption/resume, invalid input, one-recipient Community configuration and explicit activation decisions.
10. Gate D3 — Owner SMTP: only after separate readiness and Owner provider availability, guide protected on-host credential entry without chat disclosure; run preflight and controlled test; require independent Owner receipt confirmation.
11. Gate D4 — Owner Guardian/session/reboot: verify fresh canonical Guardian evidence, actual systemd process/invocation, cadence and bounded resources/restarts; evaluate lingering guidance; test logout; obtain separate Owner physical-reboot action; verify new invocation/process identity, fresh post-reboot evidence, and notification receipt continuity; then explicitly restart and reverify.
12. Gate D5 — Owner uninstall/reinstall: explicitly disable/stop, run safe matching uninstaller, verify owned artifact removal and preserved per-user configuration/credentials/state without displaying contents; reinstall the same verified RC.2, resume setup, explicitly reactivate, require new qualified evidence and no stale READY.
13. Maintain new RC.2 acceptance record and finding register after each checkpoint. Preserve raw private evidence only under Owner control; commit only bounded redacted evidence after separate Owner authority. Stop safely on every blocking/security/out-of-scope defect.
14. Run final local/reproducibility/security/preservation/rollback/governance gates and issue exactly READY or NOT READY. Complete/archive Task 051 to canonical idle only under lifecycle authority; do not tag, release, upload, announce, publish, or start a successor.


## Rollback Plan

- Before every phase, identify the exact Task 051-created local paths, candidate roots, transfer targets, and external actions. Preserve evidence before rollback and require the relevant Owner gate for external or destructive action.
- Locally verify snapshot manifests and later-edit absence. Restore only literal Task 051 targets from the verified snapshot and remove only paths with proven pre-task absence/Task 051 ownership. Never use broad reset/clean/checkout, touch Owner-owned content, RC.1/Task 049 evidence, v1.0.0, LICENSE, credentials, or unrelated files.
- Candidate and redacted evidence roots remain private until explicit Owner retention/disposal authority. Their existence does not authorize transfer or publication. Never delete an ambiguous root or use an unresolved variable/glob.
- On the external VPS, stop only the exact QWSG test service/process through documented commands; use the matching verified uninstaller only at its checkpoint; preserve per-user configuration/state/credential store. Do not undo host administration, delete the VPS, disable lingering, remove packages, or purge user data unless separately authorized.
- If transfer fails, leave source candidate unchanged, preserve checksum evidence, and remove/replace only the proven partial destination file with Owner approval. Never weaken host verification to retry.
- After rollback or stopped acceptance, rerun targeted integrity, Git/lifecycle, RC.1/v1.0.0/Owner-content preservation, and external-state checks appropriate to the reached checkpoint; report exact remaining state and safe restart point.


## Deliverables

- Exact RC.2 readiness audit and candidate-source identity decision, with any prerequisite-task recommendation if substantial correction is required.
- New RC.2-specific, numbered, restartable Owner-operated acceptance protocol and canonical acceptance/readiness record; RC.1 records unchanged.
- One private reproducible `qwsg-1.1.0-rc.2-linux-amd64.tar.gz` plus sidecar built twice from one exact integrated commit, with complete byte-identity, manifest, static binary, package, exclusion and provenance evidence.
- Owner-approved private direct VPS-to-VPS transfer procedure and privacy-safe provenance, or documented Owner-workstation fallback.
- External checkpoint evidence covering all 30 required acceptance outcomes, including explicit F002/F003 retest, real SMTP receipts, Guardian/session/physical-reboot continuity, uninstall and reinstall/resume.
- Finding register with severity, continuation/stop decision, status and correction authority boundary.
- Snapshot, rollback, security/redaction, immutable RC.1 and lifecycle/Git validation evidence.
- Final exact `READY FOR QWSG 1.1.0 RELEASE` or `NOT READY FOR QWSG 1.1.0 RELEASE` verdict, stopping at the separate Owner release gate.


## Verification

- Repository/source: canonical root/remote/branch; exact HEAD/origin and ancestry; empty index; tracked target cleanliness; truthful disclosure and exclusion of Owner-owned untracked content; Task 050 integration; no conflicting RC.2 tag/artifact; RC.1/v1.0.0/LICENSE hashes unchanged.
- Release readiness: exact `VERSION=1.1.0-rc.2`; matching notes; README/INSTALL usable in archive root; release-check/collision guards; exact full commit and explicit commit-derived epoch; no uncommitted overlay.
- Reproducibility: two independent `git archive` exports of the exact candidate commit into separate mode-0700 roots; identical controlled environment; byte-identical binary, `MANIFEST.sha256`, archive and sidecar; valid sidecar/internal manifest; deterministic metadata; safe single root/no absolute-parent/link/special members; static linux/amd64 executable; binary reports exact version/full commit/build time.
- Package allowlist: LICENSE, root README.md, root INSTALL.md, RC.2 notes, expected binary/unit/install/uninstall/config/operator docs only; no credential, secret, private evidence, host identifiers, Builder input, prompt snapshot, backup, `.git`, Owner content, RC.1 artifact, or public upload.
- Protocol structure: every numbered checkpoint contains purpose/action/expected/PASS/FAIL/safe-continuation/retain-redact; commands are bounded, operator-readable, restartable and one checkpoint at a time; privilege and external mutations have explicit gates.
- F002: on real supported host, `./bin/qwsg install --check` gives correct `systemd.user_manager` classification without stripped-session false negative; cause-specific explanation, bounded verification, privilege, only proven safe remediation, no guessed command for ambiguity, mandatory rerun, and no mutation.
- F003: `filesystem.local_semantics` uses bounded read-only evidence where possible; never returns an unconditional unexplained unknown; when automatic proof is unavailable it provides precise bounded manual verification and mandatory rerun without host mutation.
- Operator journey: fresh RC.2 receipt/checksum/layout/manifest/LICENSE/docs; Smart Install; immutable install; setup interruption/resume/invalid input; one recipient; protected credential; notification preflight; real Owner-confirmed receipt; explicit activation; fresh evidence; actual systemd process/invocation/cadence/resources/restarts; lingering/logout; physical reboot/new identity/fresh cycle; post-reboot receipt; explicit restart; safe uninstall/preservation; same-candidate reinstall/resume/reactivation/no stale READY.
- Security/privacy: no credential through chat/argv/Git/history/evidence; redact recipient/account/provider/headers/references/tokens/private host/account; no automatic sudo/packages/lingering/remediation/reboot, SSH weakening/server change/software install, public exposure, arbitrary shell, or unbounded wait.
- Regression/governance: build, focused/full/race tests, vet, formatting, shell/static systemd, package/install/uninstall, release-check, Git whitespace, secret scan, Framework 21, Builder 38, lifecycle 28, diversion 36, test-task/job, snapshot checksum and isolated rollback all pass.
- Decision matrix: READY only if every mandatory checkpoint passes, F002/F003 are externally verified/corrected, real SMTP receipt passes before and after reboot, physical reboot and new/fresh evidence pass, uninstall/reinstall pass, no open blocker/security defect exists, and all local gates pass. Otherwise exact NOT READY with unresolved findings; neither verdict authorizes publication.


## Documentation Updates

- Add `docs/release/ACCEPTANCE_PROTOCOL_1.1.0-rc.2.md` with the numbered Owner-operated protocol and phase/Owner gates.
- Add `docs/release/ACCEPTANCE_1.1.0-rc.2.md` with candidate/source/transfer provenance, checkpoint ledger, F002/F003 retest, SMTP confirmations, Guardian/reboot, uninstall/reinstall, findings and final verdict.
- Update `docs/release/RELEASE_NOTES_1.1.0-rc.2.md` only if audit shows a narrow factual candidate/acceptance clarification is needed; do not rewrite Task 050 corrections or claim acceptance before evidence.
- Update `scripts/test-release-plumbing.sh` only as needed to validate the new RC.2 protocol/record without weakening existing gates; retain RC.1 historical validation or replace hard-coded current-protocol assertions with version-aware checks that explicitly preserve RC.1.
- Update `ai/history/051_<date>_qwsg-1-1-0-rc-2-clean-host-acceptance-retest.md` throughout and the concise `ai/core/07_ENGINEERING_HISTORY.md` at verified delivery. Archive `ai/prompts/051_CURRENT_TASK.md` only through canonical lifecycle completion.
- Do not alter RC.1 acceptance/protocol/notes or Task 049 prompt/history. Reference existing operator docs rather than silently editing product guidance during external acceptance; record newly discovered gaps as findings.


## Completion Criteria

Task 051 is complete only when the Task 050 canonical-idle baseline and Owner-owned exclusion are verified; any necessary narrow Phase A source/scaffolding changes are separately integrated; one exact clean candidate-source commit is established; RC.2 is built twice from independent commit exports with controlled commit epoch and full embedded identity; all binary/manifest/archive/sidecar bytes match; package/layout/LICENSE/root-doc/static/exclusion checks pass; RC.1 remains immutable; transfer occurs only through a separately approved private two-file path; and RC.2 restarts at Checkpoint 01 on a real supported clean Ubuntu 24.04 amd64 VPS. The Owner must complete every mandatory product-visible-guidance checkpoint, including external F002/F003 correction proof, setup resume, protected one-recipient SMTP with independently confirmed receipt before and after reboot, explicit Guardian activation, fresh canonical/systemd/cadence/resource evidence, lingering/logout, physical reboot with new identity and fresh evidence, explicit restart, safe uninstall with preserved user data, and reinstall/resume/reactivation without stale READY. Every finding is truthfully classified and no unauthorized repair occurs. Final local/security/reproducibility/governance/rollback gates pass and the record states exactly `READY FOR QWSG 1.1.0 RELEASE` or `NOT READY FOR QWSG 1.1.0 RELEASE`. If substantial source correction, safe transfer, external host participation, real SMTP receipt, physical reboot, mandatory evidence, or any critical gate is unavailable, do not infer PASS: stop safely, issue NOT READY when evidence supports a terminal verdict or leave the task incomplete with the exact Owner gate required. Completion never authorizes final tag, release, upload, announcement, publication, Task 052, or alteration of RC.1/Owner content.


## Owner Approval Requirements

Approved by Project Owner through the Engineering Task Builder on 2026-08-21 UTC.

The structured task definition has been explicitly approved for implementation. Further scope changes require explicit Project Owner approval.
