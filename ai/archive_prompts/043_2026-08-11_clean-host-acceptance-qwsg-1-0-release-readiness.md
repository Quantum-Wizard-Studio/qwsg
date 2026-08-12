# Current Engineering Task 043: Clean-Host Acceptance & QWSG 1.0 Release Readiness

## Task Metadata

- Task ID: `043`
- Task slug: `clean-host-acceptance-qwsg-1-0-release-readiness`
- Status: `complete`
- Date opened: `2026-08-11` UTC
- Human authority: Project Owner
- Owner or lead-developer communication language: Hungarian

## Title

Clean-Host Acceptance & QWSG 1.0 Release Readiness


## Objective

Formally validate and record the Project Owner's real disposable-host acceptance of QWSG `1.0.0-rc.3`, reconcile it with Tasks 038, 039, 041 and 042, close the previously outstanding physical reboot gate when the supplied evidence remains internally consistent, and issue exactly one evidence-based technical decision: `READY FOR QWSG 1.0 FINAL RELEASE` or `NOT READY FOR QWSG 1.0 FINAL RELEASE` with exact blockers. This task authorizes acceptance documentation and repository validation only; it does not authorize product development or public release actions.


## Scope

- Validate the supplied clean-host evidence as owner-observed acceptance data for a freshly reinstalled Ubuntu 24.04 LTS amd64 VPS with systemd 255, ordinary user `ubuntu` UID 1000, running user manager, no prior QWSG install/state and no Go toolchain.
- Verify RC identity against the canonical local artifact `dist/qwsg-1.0.0-rc.3-linux-amd64.tar.gz`, expected SHA-256 `bc4eca323cb07d23f0d6c884886655eb5549c03c136d8c773782d7813551585c`, sidecar, internal manifest, installed version `1.0.0-rc.3`, commit `0a8a5c7e7224` and build time `2026-08-11T00:00:00Z`.
- Record successful archive-only install behavior: manifest verification, `/usr/local` installation, exact version, and the installer deliberately leaving service enablement/start to the operator.
- Record genuine empty-HOME first use: absent `~/.local/state/qwsg`, first `qwsg observe` success, eight truthful Inventory records, partial Inventory/Snapshot, one baseline and unknown-until-later condition without manual directory, permission, repository or Go preparation.
- Record second observation: complete overall execution, partial Inventory/Snapshot, complete 323-record Compare through Report stages, published Current Operator State and separately loaded degraded/current/partial Console view with 645 correlated and zero omitted concerns.
- Record user-service acceptance: enabled/running Guardian, distinct process evidence, zero restarts, bounded tasks/memory and a separate Console showing running/current/complete evidence before reboot.
- Record explicit `loginctl enable-linger`, verified `Linger=yes`, physical disposable-host reboot, automatic boot-before-login user-service start, new PID 835 and new InvocationID `90720850ce3e4d498827350861708813`, with no manual QWSG start after reboot.
- Record the completed post-reboot recurring cycle: stable PID/invocation, zero restarts, active/running state, `MemoryCurrent=13828096`, `MemoryPeak=21262336`, eight tasks, fresh Current State/checkpoint/Scheduler timestamps, `0600` state files and Console running/current/partial evidence with bounded Attention.
- Record controlled service restart with new PID 1176 and InvocationID `e2b15a8b697b4c94b7b78105cfe73101`, zero restarts and a separate running/current/partial Console.
- Record disabled/stopped service and linger before archive uninstallation, exact public-artifact removal, unit absence and intentional private user-state preservation at mode `0600`. Classify stale shell command hashing as shell behavior resolved by `hash -r`, not an uninstall defect.
- Reconcile every real-host result against the Task 038 Release Policy gates, Task 039 Runtime/Guardian/Console corrections, Task 041 clean-host bootstrap semantics and Task 042 RC.3 artifact gates.
- Assess the reported occasional double PTY Overview using existing Task 038/039/042 evidence and the clean-host observation. Treat one initial Overview plus one explicit `r` redraw as correct; classify only independently reproducible unrequested duplicate rendering or impaired operation. Do not modify Console code in this acceptance task.
- Update the canonical RC.3 engineering acceptance record and concise engineering milestone/history records with sanitized evidence and the final decision. Preserve the released RC.3 payload and its packaged documentation byte-for-byte.
- Run all configured Framework, Builder, lifecycle, diverted-task, repository, Go, race, vet, format, artifact-integrity, documentation-consistency and Git gates applicable to an acceptance-only task.
- Complete and archive Task 043 without a successor, returning to canonical idle.


## Out of Scope

- No modification of Scheduler, collector, Pipeline, Alert, Runtime, Guardian, Console, Current State, storage, installer, uninstaller or systemd-unit product semantics.
- No feature, email notification transport, setup/configuration wizard, Dashboard, API, fleet, remote execution, remediation, provider, AI, licensing enforcement or new architecture.
- No collector change to make truthful partial evidence appear complete and no weakening of path, ownership, mode, symlink, privacy or evidence validation.
- No release-identity change from `1.0.0-rc.3`, no RC rebuild, archive mutation, sidecar rewrite, manifest rewrite or relabeling of historical evidence.
- No final `1.0.0` metadata refresh, final public-license selection, signing, Git staging, commit, tag, push, upload, publication or announcement.
- No new clean-host operation, reboot, service mutation or remote VPS access; use only the Owner-supplied evidence and local read-only repository/artifact validation.
- No Task 044 or successor preparation. If a genuine product blocker is proven, stop with `NOT READY` and require a separately prepared, Owner-authorized engineering task rather than implementing it here.


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

- Require canonical idle with Task 042 as the unique latest complete archived baseline, zero active prompts and no Task 043 prompt/history/archive collision.
- Validate Framework project identity, repository root, `main`, exact HEAD, canonical HTTPS origin, `0/0` upstream relationship, empty index and complete Owner-owned dirty-tree status without changing unrelated content.
- Read the Project Constitution, Agent rules, Job Template, lifecycle/prompt/Git policies, Release Policy, Tasks 038/039/041/042 prompts and histories, all RC.1/RC.2/RC.3 acceptance records, Quick Start, Known Limitations, Operations, Troubleshooting, upgrade/uninstall guidance, installer/uninstaller, unit and relevant Console tests.
- Verify repository `VERSION` remains `1.0.0-rc.3`; archive and sidecar exist; archive SHA-256 equals `bc4eca323cb07d23f0d6c884886655eb5549c03c136d8c773782d7813551585c`; internal manifest and extracted version verify without rebuilding or installing.
- Hash all RC.1/RC.2/RC.3 artifacts, sidecars, release notes and acceptance documents before changes. Record exact documentation targets and prove no packaged RC.3 payload path will be modified.
- Confirm the Owner evidence contains no hostname, IP, secret, private configuration or unbounded journal/state payload before writing it to repository documentation. Retain only sanitized system/version, process-generation, resource, mode, timestamp-presence and outcome evidence.
- Run pre-change focused/full/race/vet/format, Framework, Builder, lifecycle, diversion/test-task, `bin/job`, artifact checksum/manifest and Git whitespace/index gates; stop on a material contradiction.


## Snapshot Requirements

- Before changing acceptance or lifecycle targets, create a unique external mode-`0700` Task 043 snapshot containing exact payloads or absence records for `docs/release/ACCEPTANCE_1.0.0-rc.3.md`, `docs/release/KNOWN_LIMITATIONS.md`, `ai/core/07_ENGINEERING_HISTORY.md`, Task 043 prompt/history, RC.3 archive/sidecar and all referenced Task 038/039/041/042 evidence.
- Record repository identity/status/index, target hashes/bytes/modes/owners/ACLs, artifact and sidecar hashes, tool versions, capacity, active QWSG/release processes and the initially absent Task 043 archive destination.
- Include deterministic manifests, a bounded readable snapshot archive, retention notes and guarded restore instructions that refuse active processes, changed targets, later Owner edits, ambiguous ownership or broad recovery.
- Preserve the Task 043 snapshot after completion. Remove only exact Task 043 temporary extraction/cache roots after verifying process absence; never remove release artifacts, Owner state or unrelated paths.


## Risk Assessment

- **Unverifiable external evidence — critical:** distinguish Owner-observed facts from local reproduction; require internal consistency across PID, InvocationID, boot, journal, state freshness and service results without fabricating raw logs.
- **False release readiness — critical:** map every Task 038/039/041/042 gate explicitly and issue `NOT READY` for any unresolved mandatory technical failure.
- **Artifact/document divergence — critical:** preserve RC.3 archive, sidecar, internal manifest and packaged documentation byte-for-byte; place post-transfer evidence only in non-payload acceptance/lifecycle records.
- **Privacy leakage — critical:** record no hostname, IP, journal dump, state payload, secret, user-specific private path beyond canonical examples or unnecessary host identity.
- **Cosmetic UX overclassification — high:** one initial render plus explicit refresh redraw is accepted behavior; do not create a blocker without repeatable unrequested duplication or operational impairment.
- **Scope expansion — high:** acceptance findings do not authorize fixes, final identity, license, commit, tag, push or publication.
- **Owner-content loss — critical:** snapshot exact targets, preserve the dirty tree and use no broad Git recovery or cleanup.


## Planned Work

1. Validate canonical idle Task 042 baseline, Framework/repository identity, RC.3 bytes, relevant historical gates, documentation targets and sanitized external-evidence completeness; run pre-change gates and create the verified snapshot.
2. Build a gate matrix mapping the supplied install, first/second observe, Guardian, linger/reboot, recurrence, restart, resources and uninstall evidence to Tasks 038, 039, 041 and 042. Mark each gate satisfied, not applicable, contradictory or blocking with a reason.
3. Assess the PTY observation against the tested Console contract. Record it as accepted redraw behavior or a bounded non-blocking post-1.0 UX observation unless objective evidence proves unsolicited duplication or practical impairment; do not change code.
4. Update `docs/release/ACCEPTANCE_1.0.0-rc.3.md` with a sanitized clean-host section, close the physical reboot/no-Go/real-uninstall gates where supported, and state the exact final technical release-readiness decision. Do not edit any file contained in the already hashed RC.3 archive.
5. Update Task 043 history and the concise engineering milestone index. Preserve `KNOWN_LIMITATIONS.md` because the verified RC.3 manifest proves it is part of the immutable payload; explain in the non-payload acceptance record that its pre-acceptance wording remains historically correct.
6. Re-run artifact hashes/manifest/version, focused/full/race/vet/format, Framework, Builder, lifecycle/diversion/test-task, privacy, documentation, snapshot and Git gates. Confirm exact changed paths, empty index and no unrelated drift.
7. If all mandatory gates pass, decide `READY FOR QWSG 1.0 FINAL RELEASE`; otherwise decide `NOT READY FOR QWSG 1.0 FINAL RELEASE` and list exact blockers. Archive Task 043 without a successor and return canonical idle.


## Rollback Plan

Stop if any QWSG/release process is active, artifact identity differs, supplied evidence contradicts a mandatory gate, a packaged RC.3 path would require modification, or current targets differ from snapshot/Task 043-owned bytes. Restore only exact snapshot-listed acceptance, milestone and lifecycle files after confirming no later Owner edit exists. Remove newly created Task 043 paths only when recorded absent and their hashes match Task 043 evidence. Never alter RC artifacts, sidecars, private state, licensing or unrelated content; never use reset, clean, checkout, broad restore, wildcard deletion or remote operations. Re-run artifact, privacy, governance, Git and idle lifecycle checks after rollback.


## Deliverables

- Sanitized canonical RC.3 clean-host acceptance evidence covering real no-Go first use, second evaluation, user-service operation, lingering, physical reboot, recurring post-reboot operation, controlled restart, resources and uninstall.
- Explicit reconciliation matrix for Tasks 038, 039, 041 and 042 with the physical disposable-host reboot gate truthfully closed or exact contradiction documented.
- Evidence-based PTY duplicate-rendering classification with no Console modification.
- Exact final technical decision: `READY FOR QWSG 1.0 FINAL RELEASE` or `NOT READY FOR QWSG 1.0 FINAL RELEASE` with blockers.
- Preserved RC.3 artifact/sidecar/manifest and packaged release documentation hashes.
- Completed Task 043 history/archive, concise engineering milestone, retained snapshot and canonical idle closure.
- Clear statement that license, final `1.0.0` identity, commit strategy, tag, push and publication remain separate Owner decisions.


## Verification

- Verify initial/final repository identity, branch/remote/upstream, complete status/index, target ownership/modes/ACLs, exact changed paths, snapshot integrity and unrelated Owner-content preservation.
- Independently verify RC.3 archive SHA-256 and sidecar, safe archive layout, internal `MANIFEST.sha256`, extracted static binary identity and exact version/commit/build time without installation or rebuild.
- Compare RC.1/RC.2/RC.3 artifact, sidecar and packaged-document hashes with pre-task evidence; require every released payload byte to remain unchanged.
- Validate the evidence matrix: clean prerequisites, installer boundary, genuine empty-HOME first partial baseline, second complete pipeline, Current State load, Guardian start, `Linger=yes`, new post-reboot PID/InvocationID, successful recurring cycle, fresh private files, restart generation, resource limits, clean uninstall and preserved private state.
- Confirm partial Inventory remains truthful and acceptable because an optional components capability is unavailable without Go; do not claim complete evidence or a QWSG Go runtime dependency.
- Confirm observed resource values remain within the shipped `MemoryMax=128M` and `TasksMax=32`, with `NRestarts=0` throughout recorded service gates.
- Confirm PTY behavior does not block release unless evidence proves an unsolicited duplicate independent of the initial render/explicit refresh or a practical usability failure. Record any cosmetic follow-up as post-1.0, not as an unapproved successor.
- Run `make build`, focused Task 039/041/Console tests, `go test ./...`, bounded-cache `go test -race ./...`, `go vet ./...`, `make fmt-check`, Framework, Builder, lifecycle, diversion, test-task audit, `bin/job`, `git diff --check` and empty-index checks.
- Scan changed documentation for secrets, hostname/IP data, raw state/journal payloads, unresolved editing markers, false publication authority and inconsistent final decisions.
- Require canonical idle at Task 043, no Task 044, no QWSG/release process, no temporary roots, no artifact mutation and no stage/commit/tag/push/publication.


## Documentation Updates

- Extend `docs/release/ACCEPTANCE_1.0.0-rc.3.md` with sanitized Owner clean-host, physical reboot, post-reboot recurrence, restart, resource and uninstall evidence plus the final technical decision.
- Preserve all files shipped inside the RC.3 archive, including `docs/release/KNOWN_LIMITATIONS.md`. Historical pre-acceptance gate wording inside the archive remains truthful for the artifact at build time.
- Update `ai/core/07_ENGINEERING_HISTORY.md` with one concise Task 043 acceptance/readiness milestone.
- Complete `ai/history/043_2026-08-11_clean-host-acceptance-qwsg-1-0-release-readiness.md` and archive the matching prompt without creating a successor.
- Do not update VERSION, command version fallback, CHANGELOG release identity, release notes, Quick Start, installer, unit, license, artifact, sidecar or manifest.


## Completion Criteria

Task 043 is complete only when lifecycle begins canonical idle at Task 042 and ends canonical idle at Task 043; RC.3 artifact identity and every packaged byte remain unchanged; the supplied clean-host evidence is sanitized, internally consistent and explicitly reconciled with Tasks 038, 039, 041 and 042; genuine empty-HOME/no-Go bootstrap, truthful partial Inventory, later full evaluation, separate Console, Guardian user service, lingering, physical reboot recovery, post-reboot recurrence, controlled restart, resource limits and uninstall are all accepted or exact blockers are named; PTY rendering is classified without unauthorized code change; all repository/governance/privacy/Git gates pass; and the record contains exactly one justified decision, `READY FOR QWSG 1.0 FINAL RELEASE` or `NOT READY FOR QWSG 1.0 FINAL RELEASE`. Final licensing, `1.0.0` metadata, commit, tag, push and publication must remain explicitly separate Owner decisions. No product fix, successor task or public-release action may occur.


## Owner Approval Requirements

Approved by Project Owner through the Engineering Task Builder on 2026-08-11 UTC.

The structured task definition has been explicitly approved for implementation. Further scope changes require explicit Project Owner approval.
