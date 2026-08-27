# Current Engineering Task 066: QWSG 1.2.0 Final Release & Distribution

## Task Metadata

- Task ID: `066`
- Task slug: `qwsg-1-2-0-final-release-distribution`
- Status: `complete — BLOCKED by QWSG-066-F001 release archive mode nondeterminism`
- Date opened: `2026-08-27` UTC
- Human authority: Project Owner Attila
- Owner or lead-developer communication language: English

## Title

QWSG 1.2.0 Final Release & Distribution


## Objective

Prepare, validate, and, only when every canonical release and acceptance gate permits it, finalize QWSG 1.2.0 as the first release based on the new Go-owned, interface-neutral installer/update architecture. This is primarily a release-engineering, acceptance-validation, distribution, and finalization task and does not authorize redesign of QWSG.

The task completes truthfully in exactly one of two ways: RELEASED, after every required gate passes and immutable Forgejo publication is independently verified; or BLOCKED, after a genuine release blocker is precisely documented and evidence is preserved without falsely publishing or declaring QWSG 1.2.0 complete.


## Scope

- Establish and record the exact repository, lifecycle, version, RC metadata, installer, updater, rollback, artifact, checksum, and documentation starting state.
- Confirm Task 065 canonical closure and the supplied baseline: private QWSG 1.2.0-rc.2; source commit c260dc18c2004473ec55496d16e66718fd128865; RC archive SHA-256 a34be8b18f80d877c0ccfd69dc9d9e9f197fc35fa765cdf1d5c0d72e2cb0a554; Task 065 closure commit f726d84632ebce3f2be72101b583af5beadc857e; and no public 1.2.0 tag, release, or artifact.
- Create and verify the required pre-task rollback-capable snapshot before modifying any Task 066 target and preserve deterministic rollback evidence.
- Audit QWSG 1.2.0-rc.2 for final-release readiness, classifying findings before changes and not silently correcting architectural deviations.
- Verify installer, updater, rollback, version consistency, artifact integrity, checksums, EN/HU/DE localization, CLI and guided behavior, supported unattended behavior, privilege and shell-helper boundaries, failures and interruption, idempotency, existing-install detection, permissions, service integration, notifications, Guardian activation/readiness, and documentation/release metadata.
- Perform clean-host acceptance on the designated OVH VPS where workflow and access permit, using the real Quantum Wizard Forgejo wget/curl distribution path and verifying acquisition, checksum, installation, first execution, Guardian readiness, filesystem and permission state, services/processes, configuration, uninstall/rollback, and reinstall/update behavior while preserving the host's sterile role.
- Perform full-host acceptance on the designated Contabo VPS where workflow and access permit, proving safe coexistence with Ubuntu 24.04 LTS, HestiaCP, Nginx, Apache, PHP-FPM/MultiPHP, MariaDB, PostgreSQL, Exim, Dovecot, Roundcube, SpamAssassin, ClamAV, BIND, Fail2Ban, firewall, quotas, SSL/Let's Encrypt, mail, web domains, and DNS without unrelated reconfiguration.
- Perform realistic supported update and rollback acceptance, including version transition, valid configuration and user-setting preservation, update policy, service/Guardian/notification state, rollback data, failed-update recovery, and proof that successful rollback restores the documented previous state.
- If and only if all gates pass, prepare and publish internally consistent QWSG 1.2.0 source/version output, installer/package/archive identity, update metadata, documentation, changelog/release notes, checksums, immutable Git tag, Forgejo release metadata, and downloadable artifacts; independently retrieve and checksum-verify the final artifact through normal wget/curl.
- Update all canonical lifecycle, release, user, and engineering documentation required by current policy, including supported install/update/rollback procedures, platforms, privileges, notifications, Guardian readiness, limitations, artifact verification, and Forgejo download/install instructions, preserving applicable EN/HU/DE localization expectations.
- Record concrete release evidence for starting state, snapshot, automated builds/tests, installer validation, clean/full host acceptance, update/rollback, artifact/checksum, direct retrieval, installed version, final Git/remote/tag state, Forgejo release, documentation, limitations, and lifecycle closure.


## Out of Scope

- Redesigning QWSG, expanding its architecture, or changing unrelated product behavior.
- Treating a successful build or unit-test run alone as release authorization.
- Silently fixing architectural deviations before classifying and documenting them.
- Changing unrelated OVH or Contabo host configuration, weakening firewalls/security controls, disabling HestiaCP or production-like services for convenience, or converting the clean OVH target into a permanent full test environment.
- Unnecessary root execution, hidden host changes, destructive cleanup of unrelated state, credential/secret storage in Git or evidence, or undocumented dependencies and architectural expansion.
- Publishing mutable, ambiguous, unverified, or incompletely gated canonical 1.2.0 artifacts.
- Modifying, committing, packaging, or otherwise incorporating the externally preserved Owner-owned QWCS_MIGRATION_BLUEPRINT.md draft.


## Authority Envelope

**Task targets and boundaries:** Task 066 may inspect and, after its verified snapshot, make the smallest architecturally correct changes required within QWSG release source, installer/updater/rollback implementation and tests, release packaging and metadata, EN/HU/DE user-facing localization where applicable, canonical documentation, task history, and lifecycle records. It may use the designated OVH clean-acceptance and Contabo full-acceptance systems only for the explicitly described bounded acceptance work. The protected Owner draft outside the repository and all unrelated repository/host state are excluded.

**Permitted external actions:** After Task 066 starts and its own gates authorize each phase, routine bounded access to the designated acceptance VPS systems, standard wget/curl retrieval from Quantum Wizard Forgejo, clean-host and production-like-host installation/update/rollback testing, and clean fast-forward repository integration are permitted. Final immutable tag creation, Forgejo release creation, artifact upload, and public distribution are permitted only after every mandatory technical, acceptance, evidence, security, and documentation gate passes. Credential entry remains an Owner interaction and secrets must not enter repository artifacts or reports.

**Owner-reserved decisions:** Any scope or architecture expansion; weakening a release, security, acceptance, evidence, snapshot, rollback, or documentation gate; accepting an unverified host or reboot claim; changing unrelated infrastructure; destructive recovery outside exact Task 066 targets; publishing despite a release blocker; replacing the designated distribution authority; or disclosing credentials requires explicit Project Owner authorization. A BLOCKED outcome must not be converted into RELEASED without new passing evidence.

**Task-specific STOP conditions:** Stop modification or publication on a material starting-state discrepancy, lifecycle/authority/security/privacy/rollback/project-identity boundary, unverified or failed snapshot, credential leakage risk, unexplained repository or host mutation, inability to preserve unrelated host services, ambiguous artifact identity, checksum mismatch, unsupported or non-restorative rollback, failed required acceptance gate, or any genuine release-blocking defect. Diagnose and correct recoverable in-scope failures, record evidence, and retest; if a genuine blocker remains, do not publish 1.2.0 and complete the lifecycle truthfully as BLOCKED with the smallest correct remediation proposal.


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

Before any Task 066 modification:

1. Inspect the repository, canonical engineering and release documentation, applicable AGENTS.md files, release scripts, packaging, installer, updater, rollback implementation, version declarations, RC metadata, and relevant prior history.
2. Run canonical lifecycle/framework validation and determine the exact lifecycle state.
3. Record branch, HEAD, local origin/main, independently verified remote main when the authorized workflow reaches that gate, remotes, staged/unstaged/untracked/ignored state relevant to release safety, current version declarations, RC metadata, archive/checksum state, installer state, updater state, rollback state, and public tag/release/artifact absence or presence.
4. Confirm Task 065 is canonically closed and compare the live state with the Owner-supplied baseline: HEAD/local origin/main expected at f726d84632ebce3f2be72101b583af5beadc857e; private RC version 1.2.0-rc.2; source commit c260dc18c2004473ec55496d16e66718fd128865; RC archive SHA-256 a34be8b18f80d877c0ccfd69dc9d9e9f197fc35fa765cdf1d5c0d72e2cb0a554; no public 1.2.0 tag, release, or artifact.
5. Confirm the Go-owned, interface-neutral installer architecture and narrow privileged shell-helper boundary are the actual baseline, including guided progress, EN/HU/DE localization, notification guidance, update policy, Guardian activation/readiness, update, and rollback support.
6. If the actual repository/lifecycle/release baseline differs materially, stop Task 066 modification work, preserve evidence, and report the discrepancy.


## Snapshot Requirements

Before modifying any Task 066 target, create a bounded, protected, rollback-capable pre-task snapshot according to ai/core/15_ENGINEERING_BACKUP_POLICY.md. Record UTC time, repository and Git baseline, exact captured scope and exclusions, retention, restore prerequisites, deterministic manifest, full SHA-256 checksums, archive readability where applicable, and exact bounded restore and post-restore verification commands. Keep full payload outside Git, publication-review any tracked metadata, verify checksum/readability after creation, and preserve enough evidence for deterministic rollback. Do not include the externally preserved Owner-owned QWCS migration draft or secrets. No implementation modification may begin until snapshot verification passes.


## Risk Assessment

- Critical release-integrity risk: publishing an incompletely validated, mutable, mislabeled, or checksum-inconsistent 1.2.0 artifact.
- High host-safety risk: installer/update/rollback tests on privileged and production-like systems could affect unrelated services, firewall, mail, DNS, databases, certificates, quotas, or web domains.
- High recovery risk: a nominal rollback may fail to restore the documented prior state or preserve configuration and supported user settings.
- High privilege/security risk: installer/helper boundary drift, excessive root execution, secret leakage, unsafe permissions, or hidden host modifications.
- Medium distribution risk: local artifacts may pass while real Forgejo wget/curl paths, release assets, metadata, or anonymous verification fail.
- Medium compatibility risk: clean-host behavior may not expose conflicts present under HestiaCP and a fully populated Ubuntu host.
- Medium localization/documentation risk: EN/HU/DE strings, user guidance, versions, commands, limitations, and release metadata may diverge.

Mitigate through inspect-before-change, verified snapshot and rollback, exact target boundaries, classification before correction, isolated/bounded acceptance, before/after host evidence, checksum verification, immutable versioned artifacts, least privilege, secret review, and publication only after all gates pass.


## Planned Work

1. Establish and record the exact starting state and confirm Task 065 closure and supplied RC baseline.
2. Create and verify the required Task 066 pre-modification snapshot and deterministic rollback procedure.
3. Conduct the complete 1.2.0-rc.2 release-readiness audit; classify every finding before changing anything and stop on architectural deviations requiring Owner authority.
4. Run proportional automated verification, including canonical framework/lifecycle checks, build, full tests, race, vet, formatting, CLI, installer, updater, rollback, localization, packaging, deterministic archive, integrity, permission, privilege-boundary, failure/interruption, idempotency, existing-install, service, notification, and Guardian readiness checks required by current policy.
5. Exercise the real Forgejo distribution path on the designated clean OVH host and record installation, first-run, readiness, filesystem/permissions/services/configuration, uninstall/rollback, and reinstall/update evidence while returning/preserving the host's clean acceptance role.
6. Exercise installation and operation on the designated Contabo full host and compare relevant before/after state to prove coexistence with all enumerated HestiaCP and infrastructure services without unrelated changes.
7. Validate the supported update matrix and successful and failed-update recovery; prove rollback restores the documented prior version, configuration, settings, service, Guardian, and notification state.
8. Diagnose and make only the smallest in-scope architecturally correct corrections for recoverable findings, recording attempts and retesting invalidated evidence. Do not silently fix architecture deviations.
9. Make the release decision. If any required gate fails, preserve and document a BLOCKED outcome and do not publish. If every gate passes, finalize internally consistent 1.2.0 versions, metadata, documentation, deterministic artifacts, and checksums.
10. Only after all pre-publication gates pass, create and verify the immutable 1.2.0 Git tag and Forgejo release/assets through the canonical distribution path, then independently wget/curl the public/version-specific artifact and checksum sidecar and verify SHA-256 and installed version.
11. Finalize release notes/changelog and canonical EN/HU/DE-related documentation expectations, record all evidence and limitations, verify final repository/remote/tag/release state, and close the lifecycle truthfully as RELEASED or BLOCKED.


## Rollback Plan

- Before implementation, define exact rollback boundaries for repository files, generated artifacts, Git references, release metadata, and each acceptance host action.
- Restore repository targets only from the verified Task 066 snapshot using exact bounded paths; never use broad reset, clean, checkout, or destructive repository-wide operations.
- Preserve pre-existing host configuration and state. For installer/update testing, use the product's documented uninstall/rollback mechanisms and independently verify the restored version, files, permissions, configuration, settings, services, Guardian state, notifications, and absence of unrelated changes.
- If publication preparation fails before external publication, remove only newly generated local Task 066 artifacts and restore exact modified targets as documented.
- If any external tag, Forgejo release, or asset action becomes partially applied, stop and follow canonical immutable-release recovery rules; do not overwrite or silently mutate a published canonical artifact.
- Preserve failure evidence and secrets boundaries. Do not delete snapshot payloads before Owner acceptance and the documented retention boundary.


## Deliverables

- A completed Task 066 prompt/history record with exact starting state, snapshot, decisions, work, verification evidence, rollback, limitations, and truthful RELEASED or BLOCKED result.
- A classified QWSG 1.2.0-rc.2 readiness audit covering installer, updater, rollback, versions, artifacts/checksums, localization, CLI/guided/unattended behavior, privileges/helper boundary, failures/interruption/idempotency/detection, permissions/services, notifications, Guardian readiness, documentation, and release metadata.
- Concrete automated, clean OVH, full Contabo/Hestia, update, failed-update recovery, and restorative rollback acceptance evidence.
- If RELEASED: internally consistent QWSG 1.2.0 source/version output, installer/package/archive/update metadata, changelog/release notes, documentation, immutable Git tag, Forgejo release/assets, artifact filename and SHA-256, direct wget/curl verification, and final installed-version evidence.
- If BLOCKED: precise blocker evidence, no public 1.2.0 promotion, preserved artifacts/evidence, and the smallest architecturally correct remediation proposal.
- Owner-facing completion report stating final decision, version, source commit, release tag, artifact filename, SHA-256, Forgejo status, clean-host result, full-host result, update result, rollback result, automated-test result, documentation result, final repository state, and remaining limitations.


## Verification

- Verify the exact starting Git/lifecycle/release baseline and Task 065 closure before modification; stop on material discrepancy.
- Verify the Task 066 external snapshot checksum, deterministic manifest, readability, scope/exclusions, restore instructions, and retention before implementation.
- Run every current canonical framework, lifecycle, Git-policy, build, full test, race, vet, formatting, packaging, deterministic rebuild, archive inspection, checksum, privacy/security, and release-readiness check applicable to QWSG 1.2.0.
- Verify source VERSION, CLI output, installer, archive filename/content, update metadata, documentation, changelog/release notes, tag, Forgejo metadata, and downloadable assets all use the same final identity.
- Verify installer and guided experience, supported unattended behavior, EN/HU/DE localization integrity, privilege/shell-helper boundary, failure/interruption behavior, idempotency, existing-install detection, filesystem permissions, services, notifications, Guardian activation/readiness, update, rollback, and uninstall behavior.
- On OVH, acquire through intended Forgejo wget/curl URLs, independently verify checksum, install, execute, confirm readiness/state/permissions/services/configuration, and test documented uninstall/rollback and reinstall/update paths without permanently changing its clean-host role.
- On Contabo, capture bounded before/after evidence proving no unexpected interference with HestiaCP, Nginx, Apache, PHP, MariaDB, PostgreSQL, Exim, Dovecot, Roundcube, SpamAssassin, ClamAV, BIND, Fail2Ban, firewall, quotas, SSL certificates, mail, web domains, or DNS.
- Validate supported upgrade transitions, configuration/settings preservation, update policy, services, Guardian, notifications, rollback data, failed-update recovery, and successful rollback that restores the documented prior state.
- For RELEASED, independently verify final tag target and immutability, Forgejo release metadata/assets, anonymous or ordinary wget/curl retrieval, checksum sidecar, SHA-256, archive contents, and installed final version; verify HEAD/local/remote synchronization and clean final Git state.
- Missing, stale, invalidated, or unperformed evidence is not PASS. A required failure yields BLOCKED and forbids publication.


## Documentation Updates

- Update the independent Task 066 history throughout execution and finalize all canonical lifecycle records.
- Update VERSION-bearing and release metadata, CHANGELOG/release notes, supported release/installation/update/rollback documentation, Forgejo distribution documentation, artifact/checksum verification instructions, and any architecture/system-map/inventory documents actually affected by the smallest approved change.
- Clearly document what QWSG 1.2.0 is, supported platforms and installation path, update path and supported matrix, rollback semantics, privileges and narrow helper boundary, notification configuration, Guardian activation/readiness, known limitations, checksum verification, and exact Forgejo wget/curl download/install procedure.
- Preserve English engineering artifacts and applicable EN/HU/DE user-facing localization expectations. Ensure commands, URLs, filenames, versions, checksums, tags, and claims match verified release reality.
- Record BLOCKED findings and remediation without implying publication if any gate fails.


## Completion Criteria

Task 066 is complete only when one of these evidence-backed outcomes is canonically recorded:

A. RELEASED: QWSG 1.2.0 passed every required automated, installer/updater/rollback, security, documentation, clean-host, full-host, update, restorative rollback, artifact, distribution, and release gate; an immutable canonical 1.2.0 tag and Forgejo release/assets were published through the authorized path; the artifact and checksum were independently downloaded with wget/curl and verified; installed version and final repository/remote state were proven; and the lifecycle was canonically closed.

B. BLOCKED: a genuine release blocker was precisely identified and documented with preserved evidence; QWSG 1.2.0 was not falsely tagged, published, or declared complete; the smallest architecturally correct remediation was proposed; rollback and repository/host safety were preserved; and the lifecycle was finalized according to canonical policy with truthful limitations.

Completion requires a concise Owner report containing: final decision RELEASED or BLOCKED; final version; source commit; release tag; artifact filename; SHA-256; Forgejo release/download status; clean-host acceptance; full-host acceptance; update result; rollback result; automated-test result; documentation result; final repository state; and remaining known limitations. No success claim may rely only on build or unit-test results, and missing evidence is not PASS.


## Owner Approval Requirements

Approved by Project Owner Attila through the Engineering Task Builder on 2026-08-27 UTC.

The structured task definition and Authority Envelope have been explicitly approved. Framework 2.0 Standard Execution Authority permits iterative, reversible in-scope engineering without another Owner gate. Further scope changes, exceptional external actions, and Owner-reserved decisions require explicit Project Owner approval.
