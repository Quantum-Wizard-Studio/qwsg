# Current Engineering Task 069: QWSG 1.2.0 Final Acceptance & Release

## Task Metadata

- Task ID: `069`
- Task slug: `qwsg-1-2-0-final-acceptance-release`
- Status: `complete — BLOCKED by QWSG-069-F001; RC.4 not released`
- Date opened: `2026-08-28` UTC
- Human authority: Project Owner: Attila
- Owner or lead-developer communication language: English

## Title

QWSG 1.2.0 Final Acceptance & Release


## Objective

Execute the complete final acceptance and release gate for the immutable private QWSG 1.2.0-rc.4 candidate. Prove the candidate through all canonical regression, clean OVH installation, real Contabo update/notification/rollback/reboot/coexistence, and external distribution checks. If and only if every mandatory gate passes, promote consistent final 1.2.0 metadata, reproducibly build, tag, publish through Quantum Wizard Forgejo, independently retrieve and verify the released artifact, and close truthfully as RELEASED. If a release-blocking candidate defect or mandatory-gate failure is found, preserve evidence, restore affected environments, do not publish final 1.2.0, and close truthfully as BLOCKED with the smallest recommended remediation and a new candidate requirement.


## Scope

- Treat `qwsg-1.2.0-rc.4-linux-amd64.tar.gz` with SHA-256 `adeb591605c0d37a5fc98d541125ca388cd4561703d0f0823bba931bc7d08684` and source commit `4f7dcc11b5ccc9f078755946995baebd31ad6870` as the immutable object under test.
- Verify canonical idle-after-068 repository, Git, version, artifact, provenance, deterministic archive metadata, ownership and permission baseline before mutation; create and rehearse the mandatory protected Task 069 snapshot and bounded rollback.
- Run all applicable Framework, lifecycle, Builder, diversion, Go, race, vet, format, build, release, reproducibility, shell, archive, permissions, Git, notification, EN/HU/DE, update and rollback regression gates.
- Perform real distributed-path clean installation, uninstall/rollback, reinstall, service/resource, reboot/readiness and unrelated-mutation acceptance on the designated sterile OVH VPS.
- Record a secret-free Contabo pre-update baseline covering installed QWSG/configuration/notification/Guardian/systemd/resource state and relevant HestiaCP, web, PHP, database, mail, DNS, security, quota, firewall, SSL and domain infrastructure.
- Use only the canonical QWSG update mechanism to update the actual installed Contabo version to RC.4; verify integrity, configuration and credential protection, Guardian/systemd/resources, host-service coexistence and truthful operation evidence.
- Exercise real lifecycle/change notifications through the existing provider-neutral event/dispatch architecture and canonical Community SMTP configuration. Separately prove operation result, notification transport result and—where an already-authorized mechanism permits—mailbox receipt. Request only minimal Owner confirmation if mailbox receipt is inherently human-observable.
- Validate safe controlled non-destructive failure notification behavior, including operation/delivery separation and secret redaction.
- Perform and state-verify real restorative Contabo rollback, rollback notification, canonical re-update where practical, controlled reboot, and full production-like coexistence validation without simplifying or repairing unrelated infrastructure.
- Make the ACCEPTED or BLOCKED RC.4 decision. No RC.4 code/source/package repair is permitted during acceptance.
- Only after ACCEPTED, prepare consistent final 1.2.0 metadata, rerun required gates, reproducibly build and inspect the canonical immutable final artifact, update release documentation, commit, create the canonical annotated tag, synchronize authorized remote state, create the Quantum Wizard Forgejo Release, and publish the artifact/checksum.
- Validate actual Forgejo user-facing distribution with direct wget and/or curl retrieval, HTTP/file/size/SHA-256/archive/provenance checks, and installability of the downloaded bytes; complete final Git/release/documentation/lifecycle verification.
- Maintain Task 069 history and privacy-safe acceptance evidence throughout; perform task-scoped Git integration and canonical closure as RELEASED or BLOCKED.


## Out of Scope

- Do not add product features or silently repair, rebuild, relabel or mutate RC.4 to make acceptance pass.
- Do not publish final 1.2.0, create its tag or represent RC.4 as final unless every mandatory final-release gate passes.
- Do not implement QUWIP, Telegram or another notification transport, and do not collapse lifecycle-event composition/dispatch into direct SMTP coupling.
- Do not simplify the Contabo production-like environment, convert OVH into the permanent full-service test server, clean up unrelated warnings, or alter unrelated Hestia/mail/web/database/DNS/SSL/firewall/quota infrastructure for a green result.
- Do not hard-code, request, disclose or record SMTP passwords, protected credentials, tokens, private keys or other secrets in Git, task evidence, logs, shell history, notifications or artifacts.
- Do not weaken security controls, firewall rules, privilege boundaries or resource limits; do not perform destructive failure injection, broad host mutation, force push, history rewrite or ambiguous/mutable publication.
- Do not claim mailbox receipt, UI state, service health, rollback restoration, download identity or any PASS solely from an exit code or unavailable evidence.


## Authority Envelope

**Task targets and boundaries:** Framework 2.0 Standard Execution Authority applies to Task 069 lifecycle/evidence; immutable RC.4 inspection and complete local regression; designated OVH clean-host acceptance; designated Contabo baseline, canonical QWSG update, notification, controlled failure, rollback, re-update, reboot and coexistence acceptance; final 1.2.0 metadata/source/release documentation only after RC.4 acceptance; deterministic final packaging; task-scoped Git integration; annotated final tag; authorized Forgejo Release/artifact/checksum publication; direct external retrieval and downloaded-artifact installation verification; and canonical RELEASED or BLOCKED closure. Candidate repair, new features, unrelated infrastructure mutation and new notification transports are excluded.

**Permitted external actions:** Fetch/compare/push the canonical Git remote; securely access and mutate only QWSG-managed state on the designated OVH and Contabo acceptance hosts; acquire private candidate bytes through the canonical safe distribution architecture; install/update/rollback/uninstall/reinstall QWSG; manage the authorized QWSG user service; execute controlled host reboots after preflight; send acceptance notifications through already-configured canonical credentials without exposing them; inspect non-secret health evidence for named coexistence services; and, only after ACCEPTED, create/push the canonical annotated `v1.2.0` tag, create the Quantum Wizard Forgejo Release, upload the immutable artifact/checksum and perform anonymous/intended wget/curl retrieval. Ordinary least-privilege escalation required by canonical install/service workflows is authorized. No unrelated external mutation is authorized.

**Owner-reserved decisions:** Acceptance waivers; product or architecture expansion; repair of an RC.4 defect and construction of a replacement release candidate; disclosure or entry of private credentials not already available through authorized protected mechanisms; confirmation of actual mailbox receipt when it cannot be proven automatically; destructive or unrelated infrastructure changes; release-policy exceptions; force push, history rewriting, identity changes outside final `1.2.0`, and any action beyond the named repository/OVH/Contabo/Forgejo targets remain reserved to Project Owner Attila.

**Task-specific STOP conditions:** Stop before mutation on a material approved-baseline, identity, authority, host-target, candidate-byte/provenance or rollback discrepancy. Stop final promotion on any mandatory regression/acceptance failure, any need to modify RC.4 source/code/package, inability to verify actual required state, secret/privacy/security exposure, destructive or unrelated-host risk, loss of deterministic rollback, unsupported update/rollback path, or an unresolved QWSG-caused coexistence regression. At the mailbox evidence boundary, pause only for the smallest non-secret Owner receipt confirmation and then continue. A blocked candidate must not be tagged or published as final; preserve evidence, restore authorized host state where required, and close canonically as BLOCKED.


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

- Before any Task 069 modification or external-host mutation verify the exact QWSG root, Framework 2.0.0, canonical idle lifecycle with Task 068 complete and archived, branch `main`, clean tracked/full worktree, canonical HTTPS origin, fetched `HEAD == origin/main == 48e4d47b6987afcaa5986c86c4d321fd2d6c7f94`, and `0 0` divergence.
- Verify version and release metadata identify private candidate `1.2.0-rc.4`; verify artifact `qwsg-1.2.0-rc.4-linux-amd64.tar.gz`, SHA-256 `adeb591605c0d37a5fc98d541125ca388cd4561703d0f0823bba931bc7d08684`, source commit `4f7dcc11b5ccc9f078755946995baebd31ad6870`, sidecar/provenance and Task 068 freeze evidence.
- Independently verify archive readability, safe paths/types, numeric ownership `0/0`, directories `0755`, `bin/qwsg`, `install.sh` and `uninstall.sh` at `0755`, every other regular file `0644`, embedded version/source/build metadata, internal manifest and deterministic reproducibility evidence.
- Verify no final `v1.2.0` tag, final artifact or Forgejo Release is already presented as released unless this is explained by canonical state.
- Resolve and record exact authorized OVH and Contabo host identities/access paths without committing secrets. Before Contabo mutation determine actual installed QWSG version and supported update path rather than assuming it.
- Any material difference from this approved baseline is a pre-mutation STOP and must be reported.


## Snapshot Requirements

- Before modifying task targets, create a protected mode-`0700` Task 069 snapshot root outside Git following `ai/core/15_ENGINEERING_BACKUP_POLICY.md`.
- Capture the complete tracked repository baseline plus exact Builder-created Task 069 prompt/history before-images, Git refs/status/remote/version/candidate metadata, RC.4 artifact and checksum/provenance evidence, and privacy-safe pre-mutation state required to restore every QWSG-managed OVH and Contabo target that will change.
- For each host mutation, create the canonical local/remote QWSG snapshot or update backup required by the supported workflow, record exact paths, ownership/modes, service state and guarded restore prerequisites without secrets or unrelated host payload.
- Record UTC time, purpose, scope, exclusions, retention through Owner acceptance and release rollback-window closure, deterministic manifests and SHA-256 checksums. Verify checksum/readability and rehearse repository restoration only into an isolated protected directory; rehearse host rollback through safe non-mutating validation where the canonical framework permits.
- Never extract over the live worktree, overwrite existing targets ambiguously, commit snapshot payloads, or collect protected credentials.


## Risk Assessment

- Critical release-integrity risk: RC.4 could be mutated or promoted despite a missing gate. Mitigate with frozen hash/provenance checks, phase gates, evidence-backed ACCEPTED/BLOCKED decision and strict no-repair/no-publication-on-failure rules.
- High external-host risk: install/update/rollback/reboot can affect real systems. Mitigate with exact target verification, secret-free preflight/baselines, least privilege, canonical mechanisms, snapshots, deterministic rollback, state verification and no unrelated cleanup.
- High coexistence risk on Contabo: production-like Hestia, web, database, mail, DNS, firewall, quota and SSL services must remain intact. Compare before/after evidence and distinguish pre-existing conditions from QWSG-caused regressions.
- High credential/privacy risk: real notifications use protected SMTP configuration. Never print, copy, log or commit secrets; verify protection/redaction and keep operation, transport and mailbox evidence separate.
- High publication risk: tag, push and Forgejo assets are externally visible. Permit them only after every acceptance gate, exact final metadata/reproducibility review and immutable checksum establishment.
- Medium reboot/connectivity risk: record expected state and ensure no unrelated operation is active; verify return, services and resource controls after each controlled reboot.
- Rollback risk: command success is insufficient. Require version/configuration/service/file/resource and coexistence state verification, with bounded recovery and STOP if restoration cannot be proven.


## Planned Work

1. Read all required governance, prior Task 063-068 release/acceptance records, current release/update/rollback/notification/install documentation and relevant implementation. Verify canonical repository/candidate/remote/host baseline and record evidence before mutation.
2. Create, verify and rehearse the mandatory protected repository and host rollback snapshots.
3. Run the complete canonical pre-acceptance regression matrix: Framework/lifecycle/Builder/diversion, full/race/vet/format/build/release/reproducibility/shell/archive/permissions/Git, notification EN/HU/DE, update and rollback gates. Stop final release on a mandatory failure.
4. On designated OVH, prove clean starting state and intended distributed candidate acquisition, independent checksum/archive/mode verification, least-privilege installation, files/version/configuration, Guardian/systemd/resources, restart, uninstall/rollback, reinstall, reboot readiness and absence of unrelated mutation.
5. On Contabo, capture installed-version/configuration/notification/Guardian/systemd/resource/filesystem and named production-like infrastructure baseline without secrets; derive the actual supported canonical update path.
6. Update Contabo only through QWSG's canonical mechanism. State-verify RC.4, preserved valid configuration and protected credentials, Guardian/systemd/resources, integrity evidence and unrelated-host health.
7. Validate real administrator update notification identity/event/version direction/result/time/action status. Record operation, transport and mailbox receipt levels separately; ask the Owner only for minimal non-secret mailbox confirmation if automation cannot prove receipt.
8. Exercise safe controlled non-destructive failure notification behavior, proving underlying operation visibility, independent delivery failure and redaction.
9. Perform and state-verify real restorative Contabo rollback and rollback notification; canonically re-update to RC.4 where practical and verify it independently.
10. Perform controlled reboot acceptance on required hosts, then verify QWSG readiness/version/configuration/resources and the Contabo coexistence matrix across Hestia, web/PHP, databases, mail/antispam/antivirus, DNS, Fail2Ban, firewall, quotas, SSL and existing domains.
11. Decide ACCEPTED or BLOCKED. If blocked, preserve evidence, restore required state, document the exact gate/remediation and close without tag/publication.
12. Only if accepted, change metadata to final 1.2.0, update release notes/docs, rerun mutation-invalidated gates, build twice under controlled differing conditions, require byte/SHA identity, inspect final permissions/ownership/provenance and verify CLI output.
13. Perform targeted Git review/staging/commit, create the canonical annotated tag, push authorized branch/tag, create the Forgejo Release and attach immutable artifact/checksum.
14. Retrieve the real hosted artifact with wget/curl, verify HTTP/filename/size/SHA/archive/provenance and install the downloaded bytes on the appropriate acceptance target. Verify final local/remote/tag/release/docs/worktree/lifecycle state and close truthfully as RELEASED.


## Rollback Plan

- Before every mutation identify the exact repository or QWSG-managed host targets, verify the protected snapshot/checksums, and record expected post-restore state. Restore only reviewed explicit paths or use the documented canonical QWSG rollback/uninstall/update mechanism; never use broad reset, checkout, clean, wildcard deletion or extraction over live state.
- Repository pre-publication rollback restores only Task 069-reviewed paths from isolated verified snapshot material and reruns focused/full/release/lifecycle checks. After Git integration use forward corrective commits unless an Owner-reserved ref action is required; never rewrite history or silently move a published tag.
- OVH rollback uses the canonical uninstall/rollback behavior and exact pre-test inventory to return the sterile target to its required clean acceptance state, followed by file/service/unrelated-mutation verification.
- Contabo rollback uses the supported real QWSG rollback and captured pre-update backup; verify actual version, valid configuration, protected notification configuration, Guardian/systemd/files/resources and all relevant unrelated services. If final intended state is RC.4, re-update only through the canonical path and verify again.
- Reboot recovery verifies host reachability and expected state; do not repair unrelated pre-existing warnings. If rollback is unavailable, ambiguous, destructive, leaks secrets or cannot prove restoration, stop and obtain Owner direction.
- If publication has not occurred and RC.4 is blocked, do not create final refs/assets. If an externally visible publication step partially fails after acceptance, preserve immutable checksums and exact remote evidence, avoid mutable replacement, diagnose within the authorized release model and stop before any unsafe ref/asset mutation.


## Deliverables

- Complete privacy-safe Task 069 evidence for canonical baseline, protected snapshots, rollback rehearsal and every mandatory local regression gate.
- OVH clean distributed-install, service/resource, uninstall/reinstall, reboot/readiness and unrelated-mutation acceptance result.
- Contabo actual-version baseline; canonical RC.4 update; real notification, safe failure, restorative rollback, rollback notification, canonical re-update/final state, reboot and full coexistence results.
- Explicit three-level notification evidence: QWSG operation, transport acceptance/failure and actual mailbox receipt, without fabricated or secret-bearing claims; EN/HU/DE verification.
- Evidence-backed final RC.4 decision. If BLOCKED: exact failed gate/evidence/environment/mutation/restoration/repository/lifecycle state and smallest remediation, with no final tag or publication.
- If ACCEPTED: consistent final 1.2.0 source/metadata/docs, deterministic immutable final archive and checksum, annotated Git tag, synchronized branch/tag, Quantum Wizard Forgejo Release and assets.
- Real wget/curl distribution-path retrieval with byte/checksum/archive/provenance/installability verification of the hosted artifact.
- Updated release/task documentation, completed Task 069 history, clean reviewed repository and canonical lifecycle closure as RELEASED or BLOCKED.


## Verification

- Validate Builder input/install, prompt/history identity, required reading, Framework, lifecycle, Git policy, exact root/branch/remote and pre/post repository consistency.
- Independently verify RC.4 filename/hash/source/version, sidecar and archive integrity, safe paths/types, embedded metadata/manifests, numeric `0/0` ownership and canonical modes; prove correspondence to recorded source and prior reproducibility evidence.
- Run every applicable canonical engineering/framework/lifecycle/Builder/diversion test, shell syntax check, full Go test, race test, vet, formatting, build contract, release plumbing/check, deterministic repeated build, archive/permission/Git check, notification/localization/redaction/duplicate test, update test and rollback test.
- OVH: verify clean preflight, real distributed acquisition, independent checksum, extraction/modes, installer privilege boundaries, installed files/version/config, Guardian validation/readiness, systemd start/restart, resources, uninstall/rollback, reinstall, reboot persistence and no unrelated mutation.
- Contabo before/after each mutation: verify actual version, configuration/protected credentials, Guardian PID/service/files/systemd/resources and relevant HestiaCP, Nginx, Apache, PHP, MariaDB, PostgreSQL, Exim, Dovecot, Roundcube, SpamAssassin, ClamAV, BIND, Fail2Ban, firewall, quotas, SSL and domain health. Record pre-existing conditions separately.
- Update/rollback/re-update must use canonical mechanisms and state verification, not exit status alone. Verify version direction, configuration contract, service/resources, integrity evidence and absence of Hestia/firewall/SSL/mail/web/database regressions.
- Notifications: verify identity, privacy-safe host, event, old/new/restored versions, SUCCESS/FAILED, timing and action state; independently record operation result, provider/SMTP transport result and actual mailbox receipt level; verify safe failure injection, delivery failure separation and no secret leakage; prove EN/HU/DE contracts.
- Reboots: verify expected pre-state/no competing operation, host return, version/config validation, Guardian readiness/systemd/resources and Contabo critical service recovery.
- Final release only after ACCEPTED: verify final 1.2.0 metadata/source/docs consistency, repeated-build byte and SHA identity, final CLI output, archive metadata, targeted Git diff/staging, annotated tag target/local-remote agreement, branch synchronization, Forgejo Release/assets/checksum and immutability.
- External distribution: use real wget/curl endpoint; verify accessibility/HTTP success, filename, nonzero expected size, SHA-256 byte identity, archive/provenance and installation from the downloaded artifact.
- Final verify no RC.4 is falsely presented as final, no secrets/unrelated changes, clean worktree, `HEAD == origin/main`, intended tag/release identity, and canonical RELEASED or BLOCKED lifecycle closure.


## Documentation Updates

- Maintain the independent Task 069 history chronologically with exact baseline, snapshots, commands/results, classifications, host mutations, rollback/restoration, notification evidence levels, acceptance decision, artifact/Git/Forgejo/download results and unresolved limitations; never include secrets.
- Update final 1.2.0 release notes, version/release metadata, installation/update/rollback/notification documentation and distribution links only as required by proven final release state.
- Record RC.4 provenance and acceptance disposition without presenting private candidate bytes as final. If blocked, document the exact candidate blocker and smallest new-candidate remediation without implementing it.
- Preserve the provider-neutral lifecycle event/dispatch architecture and document operation-versus-delivery-versus-mailbox semantics, EN/HU/DE results, actual supported update/rollback behavior and known limitations.
- Update the general engineering milestone index only with the concise final Task 069 outcome required by repository policy, and archive/close lifecycle records canonically.


## Completion Criteria

Task 069 is complete only with one truthful terminal result. RELEASED requires every mandatory local, OVH, Contabo, notification, rollback, reboot, coexistence, deterministic packaging, Git, Forgejo, external wget/curl, downloaded-byte/installability, documentation and final lifecycle gate to pass; final version must be 1.2.0 with an intended source commit, canonical annotated tag, immutable artifact filename and independently verified SHA-256, synchronized repository/release state, clean worktree and canonical closure. BLOCKED requires stopping promotion at the failed mandatory gate, preserving exact evidence, performing and verifying required bounded restoration, creating no final v1.2.0 tag or Forgejo final release, and recording candidate version/hash, affected environment, mutations, rollback state, repository/lifecycle state and smallest recommended remediation/new-candidate requirement. Mailbox receipt must never be fabricated; when it is the sole inherently human evidence boundary, minimal non-secret Owner confirmation is required before a RELEASED decision.


## Owner Approval Requirements

Approved by Project Owner: Attila through the Engineering Task Builder on 2026-08-28 UTC.

The structured task definition and Authority Envelope have been explicitly approved. Framework 2.0 Standard Execution Authority permits iterative, reversible in-scope engineering without another Owner gate. Further scope changes, exceptional external actions, and Owner-reserved decisions require explicit Project Owner approval.
