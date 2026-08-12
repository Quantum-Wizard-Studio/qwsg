# Current Engineering Task 038: QWSG Version 1.0 Release Hardening and Release Candidate

## Task Metadata

- Task ID: `038`
- Task slug: `qwsg-version-1-0-release-hardening-and-release-candidate`
- Status: `complete`
- Date opened: `2026-08-10` UTC
- Human authority: Project Owner
- Owner or lead-developer communication language: Hungarian

## Title

QWSG Version 1.0 Release Hardening and Release Candidate


## Objective


Treat Task 038 as the final engineering release gate for the existing local QWSG product. Freeze the implemented Version 1.0 scope, close only verified release blockers, and produce a reproducible, installable, rollback-capable QWSG 1.0 Release Candidate for the narrowly supported platform.

The candidate shall let an ordinary operator install prebuilt QWSG artifacts without repository or Go knowledge, run the existing one-shot and continuously supervised workflows, understand truthful Console and service diagnostics, upgrade or roll back without corrupting valid state, and uninstall owned artifacts under an explicit data-retention policy. It shall preserve every canonical engine boundary and add no optional product family.

Task 038 shall end with exactly one evidence-backed engineering decision: `READY FOR QWSG 1.0 RELEASE` or `NOT READY FOR QWSG 1.0 RELEASE`. `READY` requires a complete Release Candidate artifact set and all applicable acceptance gates below. `NOT READY` shall identify only genuine blockers and the smallest correction. External publication, tagging, licensing selection, and final release authorization remain separate Owner actions and are not implied by technical RC readiness.



## Scope


- Establish a versioned QWSG 1.0 capability inventory from the repository and freeze it to the implemented local product: Inventory, Snapshot, Comparison, Drift, Health, Rule, Policy, Report, Configuration, Scheduler, Alert, provider-neutral Notification contracts, Runtime, Runtime Service, Operator Presentation Model, Current Operator State, Interactive Operator Console, canonical `observe`, and the Operational Guardian Service.
- Audit every historical release gate against current implementation and current Owner authority. Classify findings as `BLOCKER`, `SHOULD`, or `POST-1.0`; only BLOCKER work may change product/release artifacts. Reconcile stale pre-alpha documents that still describe already-implemented topology or optional provider/network families as mandatory, without silently claiming unimplemented behavior.
- Define the narrow supported platform as Ubuntu 24.04 LTS, systemd 255 or later, Linux amd64/x86-64, glibc-compatible userspace, an ordinary non-root user with a working systemd user manager, a local filesystem supporting atomic rename, advisory `flock`, owner/mode checks and private `0700`/`0600` state. Other Linux distributions and CPU architectures are experimental or unsupported unless Task 038 produces equivalent real evidence; do not advertise them as supported.
- Replace the pre-alpha release-policy shell with a concrete Version 1.0 policy covering semantic version identity, Release Candidate naming, artifact contents, deterministic build inputs, checksums, compatibility, support claims, approval, rollback, known limitations and publication separation. Use the existing VERSION/ldflags behavior rather than adding a competing version system.
- Prepare one coherent RC identity, expected to be `1.0.0-rc.1` unless repository compatibility analysis proves a different existing-policy-compatible spelling is required. Update VERSION, version output, changelog/release notes and documentation coherently. Do not tag, sign, publish or push.
- Build the smallest professional artifact set for `linux-amd64`: one statically inspectable/versioned QWSG binary built from the standard-library-only module, the systemd user unit, deterministic install and uninstall tooling or an equivalently simple no-Go operator workflow, default/example configuration guidance, Quick Start/operator/troubleshooting/upgrade/rollback/uninstall documentation, release notes, a manifest, and SHA-256 checksums. A tar archive is preferred; do not add DEB, RPM, AppImage, container or package-repository machinery.
- Make release installation prefix-safe and artifact-safe. Resolve the current hard-coded `/usr/local/bin/qwsg` unit/install-prefix mismatch. Installation shall verify platform, artifact checksum, exact destinations and collisions; separate privileged artifact copying from ordinary-user runtime/service activation; install but never silently enable/start the service; preserve and report pre-existing artifacts; establish no root-running Guardian and no broad directory mutation.
- Provide a bounded uninstall path that stops/disables only the exact QWSG user unit when explicitly requested, removes only manifest-owned release artifacts after identity checks, never deletes shared dependencies, and preserves private operator state/configuration by default. Document an explicit, separately confirmed purge procedure if offered; do not automatically remove user data.
- Define canonical configuration discovery for the released service without hidden precedence. The built-in operational configuration remains sufficient for default operation. If a user configuration file is supported, document its exact Source Record 1.0 path, unit override mechanism, strict permissions and failure diagnostics. Ship a validated example rather than inventing a new configuration engine or secret backend.
- Validate the actual systemd lifecycle using a uniquely named temporary user unit, isolated installation/configuration/state roots, bounded journal window and exact cleanup: install/link, daemon reload, start, recurring cycles, separate Console, graceful stop, stopped evidence, restart, checkpoint/Scheduler recovery, abnormal termination, demotion, bounded restart/start limit, invalid configuration, interrupted cycle, stale evidence, single-instance refusal and no false running claim.
- Define reboot behavior precisely. Enabling the user unit is separate from starting it; boot-before-login requires an administrator-authorized lingering setting for the exact runtime user. Do not alter the engineering host's existing linger state. Validate enablement, default target linkage and isolated user-manager recovery. If no disposable host/VM reboot is authorized and available, provide an exact Owner-run Ubuntu 24.04 reboot acceptance procedure and label physical reboot as unexecuted evidence, not proven support.
- Define and exercise the supported upgrade/rollback boundary for the current 1.0/1.1/1.2 stored contracts: binary/unit replacement under stop/restart, valid Current Operator State compatibility, Scheduler state compatibility, Guardian checkpoint/configuration identity compatibility, interrupted upgrade recovery, incompatible schema fail-closed behavior, restoration of the previous binary/unit, and successful operation without rewriting unknown state. No broad migration framework is authorized.
- Perform bounded security/privacy review of executable/unit/install/uninstall/config/state/journal/Console boundaries: non-root runtime, no listeners, no remote execution, no arbitrary shell execution, strict path/symlink/ownership/mode behavior, state integrity, configuration privacy, secret-reference-only handling, systemd sandbox, controlled signals, log tokens, no host-sensitive raw errors, and privileged installation separation. Fix only exploitable or release-blocking defects.
- Perform bounded endurance/resource acceptance across at least 50 short isolated Guardian recurrence boundaries after warm-up. Record RSS/MemoryCurrent, CPU, file descriptors, tasks/threads, goroutine evidence through in-process tests, snapshot counts/bytes, Scheduler state size/result cap, checkpoint/current-state sizes, journal behavior and restart count. Prove bounded retention, no monotonic goroutine/FD/task growth, no restart loop and compliance with documented systemd limits. Use Task 037's approximately 48.7 MiB peak and 9 tasks as comparison evidence, not a permanent assertion.
- Investigate the reported duplicate initial interactive Console Overview rendering against a real PTY and deterministic session tests. If confirmed, correct only the render/redraw boundary and test startup, refresh, navigation, terminal fallback, EN/HU localization, overflow disclosure, stale/partial wording, Guardian lifecycle wording and recommendations. Do not redesign the Console.
- Produce ordinary-user English documentation for product purpose, support matrix, release verification, installation, first baseline and second full `qwsg observe`, bare `qwsg`, Guardian enable/start/stop/restart/status/logs, configuration, condition/attention/change/Alert/evidence interpretation, diagnostics, upgrade, rollback, uninstall, data preservation, privacy/security and known limitations. Keep Hungarian operator strings and existing Hungarian direction accurate; add only bounded Hungarian Quick Start parity if the existing documentation structure requires it for a coherent local product.
- Execute the full release journey from the built RC artifacts, not the repository binary alone: clean isolated supported-host layout, checksum verification, install, version verification, baseline, full observation, Guardian recurring cycles, separate Console, stop/stopped, restart/recovery, bounded failure, upgrade/reinstall, state compatibility, rollback, operation after rollback, uninstall, owned-artifact absence and documented preserved-data result.
- Update only directly affected release, installation, operator, architecture, roadmap, specification, changelog, prompt/history/archive and acceptance evidence. Preserve all unrelated Owner-owned working-tree content exactly.



## Out of Scope


- New Inventory collectors or new Health, Rule, Policy, Report, Scheduler, Alert, Notification, Runtime or Runtime Service business semantics.
- Concrete SMTP, Telegram, Discord, webhook or other provider transports. Provider-neutral Notification architecture and locally visible Alerts are the Version 1.0 boundary unless the existing implementation itself cannot operate without a provider.
- Web Dashboard, REST API, public listener, fleet/remote management, cloud service, remediation, infrastructure management, licensing enforcement, billing, telemetry or AI.
- General persistence, database, migration, updater, package manager, signing service, trust-root/key-custody system, DEB/RPM/AppImage/container packaging or multi-distribution build farm.
- Root-running Guardian, privileged collectors, new capabilities, new users/groups, hidden configuration precedence, automatic lingering changes or modification of an unrelated user/system service.
- Physical/VM reboot of a non-disposable host, alteration of the Owner's real QWSG installation/state, or destructive cleanup outside exact Task 038 acceptance roots.
- Selection of the final public license, prices, editions or commercial terms. The existing temporary proprietary notice may govern a private RC, but the Owner must confirm or replace it before public publication.
- Git stage, commit, tag, branch, fetch, push, external upload, package publication, release publication or Version 1.0 public announcement.
- Task 039 preparation or installation. A successor is permitted only if Task 038 ends `NOT READY` with a genuine unresolved blocker that cannot safely be corrected within this release-hardening scope.



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


- Begin only from canonical idle with no active prompt, Task 037 as the unique latest complete archived prompt/history pair, no Task 038 prompt/history/archive collision, and `bin/job --check` success.
- Verify repository root `/home/qws/web/qwsg.quantumwizard.hu/qwsg`, Framework 1.x markers/configuration, canonical HTTPS origin, `main`, exact HEAD, `origin/main...main`, complete Git status, empty index, ownership, modes and ACLs. Preserve the large pre-existing Owner-owned modified/untracked worktree; do not interpret it as Task 038 output.
- Verify the Task 037 implementation snapshot and history, final systemd acceptance evidence, supported user manager and its cleanup. Re-read Tasks 025–037 and every relevant canonical architecture, installation, product, functional, roadmap and release-policy document.
- Record current verified facts: VERSION is `0.0.1-prealpha`; Go module requires Go 1.26 and uses only the standard library; build embeds VERSION/commit/date; the unit currently names `/usr/local/bin/qwsg`; Make installs a binary/unit but offers no release archive, checksum manifest or complete uninstall target; LICENSE is a temporary proprietary notice; Release Policy is still a pre-alpha shell; Ubuntu 24.04/systemd 255/linux-amd64 is the only real operational acceptance baseline.
- Re-run current build, full tests, race tests where feasible, vet, format, Framework, Builder, lifecycle, diverted-task, Git and systemd unit verification before implementation. Record any baseline failure as a release BLOCKER and stop if it contradicts Task 037 completion.
- Inspect Runtime/Guardian import and call boundaries, configuration resolution, all durable formats, Current Operator State legacy reads, checkpoint failure behavior, systemd unit sandbox, install prefix behavior, Console PTY/session behavior, retention/resource caps, privacy-sensitive logs and absence of listeners.
- Produce a Version 1.0 capability/gate table before modifying targets. Explicitly classify old Product Definition/Functional Specification gates in current product terms: local Console ships and is not network exposed; default built-in configuration is implemented; local file state is bounded; continuous systemd topology exists; provider transports remain optional; public licensing/signing remains outside private technical RC authority; only narrow tested platform claims are permitted.
- Confirm authority for isolated temporary user-service acceptance and release artifact generation inside the repository or `/tmp`. Stop for Owner input only if safe execution requires changing real linger state, rebooting a non-disposable host, choosing public licensing terms, signing/publishing, or broadening the frozen product.



## Snapshot Requirements


- Before implementation create `/tmp/qwsg-task038-implementation-<UTC>-<random>` with mode `0700` and an exact bounded manifest of every intended target: VERSION, LICENSE, CHANGELOG, Makefile, Release Policy, README, installation/user/troubleshooting/release documents, systemd unit, release tooling/tests, Console/session targets, configuration examples, directly affected core/product/roadmap/specification files, Task 038 prompt/history, and verified absence records for new release artifacts/scripts/docs/archive.
- Capture repository root, HEAD, branch, remotes, ahead/behind, complete status, staged paths, owner/group/mode/ACL, tool/platform versions, systemd user-manager state, linger observation, unit hashes, current test results and pre-change artifact hashes. Do not copy unrelated host state or secrets.
- Add a separate service-acceptance snapshot before linking any temporary unit. It shall identify the exact unique unit name, unit hash, executable hash, isolated roots, pre-existing absence, manager state, journal cursor/time boundary, cleanup commands and stop conditions.
- Verify every snapshot payload and absence record with SHA-256 before changes. Store guarded restore instructions that reject target drift, later Owner edits, ambiguous identity, active service state or broad deletion.
- Retain snapshots through delivery. Snapshot existence is rollback evidence, not authority to overwrite Owner changes.



## Risk Assessment


- **Unsupported release claim — critical:** narrow the support matrix to actually tested Ubuntu 24.04/systemd 255/linux-amd64; label all other contexts experimental/unsupported and disclose unperformed physical reboot evidence.
- **Artifact/version mismatch — critical:** derive filenames, embedded version, manifest, checksums, changelog and release notes from one validated VERSION value; rebuild twice with controlled inputs and compare hashes where reproducibility is claimed.
- **Installation privilege/path error — critical:** validate exact prefix/unit executable correlation, collision backups and staged layouts; runtime stays non-root; installer never silently enables, starts or mutates linger.
- **State loss during upgrade/uninstall — critical:** stop exact service, snapshot exact private state, validate schemas, preserve data by default, refuse unknown formats, restore only verified binary/unit artifacts and never recursively delete user state.
- **False Guardian lifecycle after crash/restart — critical:** exercise real abnormal exit, correlated ExecStopPost demotion, restart bounds, freshness and checkpoint generations under isolated systemd acceptance.
- **Historical requirement contradiction — high:** reconcile obsolete pre-alpha gates through explicit current capability/classification evidence; do not pretend missing provider/fleet/network features exist or weaken canonical safety requirements.
- **Release authenticity limitation — high:** SHA-256 detects corruption but does not authenticate an untrusted distribution channel. Document this limitation; no signing/key system is invented. Public distribution remains Owner-authorized and license/signing policy controlled.
- **Console redraw regression — medium:** reproduce under PTY before editing; limit changes to deterministic screen redraw and retain noninteractive output/tests.
- **Endurance/storage growth — high:** measure warm-up and at least 50 cycles, verify hard caps/retention and state sizes, stop on monotonic FD/goroutine/task growth, resource-limit breach or restart loop.
- **Acceptance host damage — critical:** use a unique Task 038 unit and isolated roots, never real QWSG state, record identities before control, and perform exact cleanup with post-cleanup absence checks.
- **Dirty worktree overlap — high:** compare every target with the implementation snapshot before edit, preserve unrelated Owner content, keep the index empty and do not use Git reset/clean/checkout/restore.
- **Public-license authority — high:** technical RC creation is authorized; license selection, tag and public publication are not. READY is an engineering decision, not legal/publication approval.



## Planned Work


### Phase 1 — Release audit and freeze

1. Load through `job`, reread authority/governance, reverify canonical idle baseline, platform, Git, owner content and Builder-installed scope; create and verify the implementation snapshot.
2. Build the capability/release-gate matrix. Audit release policy, LICENSE, VERSION, changelog, supported-platform claims, install/unit prefix, artifact completeness, configuration/state compatibility, security, Console behavior and product requirements. Freeze BLOCKER/SHOULD/POST-1.0 classifications before changes.
3. Freeze Version 1.0 compatibility and support contracts: Ubuntu 24.04 LTS, systemd 255+, linux-amd64, non-root user service, current configuration and stored schema versions, built-in default operation, optional providers and no network Console.

### Phase 2 — Minimal hardening and artifacts

4. Correct only confirmed release defects, including install/unit prefix correlation, ordinary-user no-Go installation, privacy-safe actionable diagnostics and duplicated interactive startup rendering if reproduced. Add regression tests before or with each correction.
5. Implement deterministic release assembly from controlled existing build inputs. Produce the binary, unit, example configuration/guidance, manifest, install/uninstall tooling, release notes and SHA-256 checksums in one versioned linux-amd64 archive. Validate archive paths/modes, no secrets, no repository-only dependencies and repeat-build identity under the documented reproducibility model.
6. Establish concrete release/upgrade/rollback/uninstall policy and documentation. Preserve existing valid 1.0/1.1/1.2 state; reject incompatible/tampered formats; never silently migrate or purge unknown/user data.
7. Update VERSION to the coherent RC identity and align version output, changelog, Release Policy, README, support matrix, Quick Start, installation, configuration, Guardian operations, diagnostics, security/privacy, limitations and English operator documentation. Keep localization catalogs and Hungarian guidance accurate.

### Phase 3 — Executable release acceptance

8. Run staged clean-install acceptance exclusively from the RC archive/checksums. Prove version, baseline, second observation, service installation, recurring cycles, Current State and separate Console behavior.
9. Use the predeclared isolated systemd user unit/root to verify graceful and abnormal lifecycle, restart/start-limit, invalid configuration, interrupted cycle, stale evidence, single-instance exclusion, checkpoint recovery, logs, sandbox and exact cleanup.
10. Exercise upgrade/reinstall and rollback using two exact artifact identities and preserved valid state. Tamper/incompatibility cases must fail closed while the previous valid data remains recoverable.
11. Run the 50-cycle endurance window and security/privacy audits; record sanitized measurements and caps. Execute the bounded enablement/user-manager recovery test and prepare the exact Owner-run physical reboot procedure if a disposable reboot environment is absent.
12. Execute the complete release journey and all automated/build/race/vet/format/systemd/Framework/Builder/lifecycle/Git/snapshot/install/uninstall validations. Rebuild release artifacts from clean controlled inputs and verify manifests/checksums.

### Phase 4 — Decision and lifecycle close

13. Update Task 038 history with actual artifact hashes, acceptance evidence, support claims, limitations, cleanup and rollback. Do not claim checks not performed.
14. Issue `READY FOR QWSG 1.0 RELEASE` only if every BLOCKER is closed. Otherwise issue `NOT READY FOR QWSG 1.0 RELEASE` with only unresolved blockers and smallest corrections. Classify remaining suggestions and deferred families without creating a successor automatically.
15. Mark and archive Task 038, verify canonical idle with Task 038 as latest complete baseline, empty index, unchanged branch/HEAD/ahead-behind, preserved Owner worktree and no temporary service residue. Do not stage, commit, tag, push or publish.



## Rollback Plan


Rollback is exact, service-safe, state-preserving and artifact-manifest-driven. Stop if unit identity, executable hash, state root, snapshot checksum or target ownership differs from recorded acceptance facts.

For an acceptance rollback, stop/disable/reset/unlink only the unique Task 038 temporary user unit after proving its unit hash and MainPID executable. Reload only the user manager, close only the recorded journal window, and remove only exact temporary unit/config/install/state roots bearing the Task 038 marker and manifest. Never alter the real `qwsg-guardian.service`, linger setting, real QWSG state/configuration, another service or a broad home/system path.

For installed-artifact rollback, stop the exact Guardian, restore the recorded prior binary and unit atomically with their original owner/mode/hash, daemon-reload, validate the restored version and restart only when it was previously active. Preserve Current Operator State, Inventory Store, Scheduler state and Guardian checkpoint. If the restored binary rejects a newer schema, leave the service stopped and retain/export state for Owner review; do not rewrite or delete it.

For repository rollback, verify the Task 038 implementation snapshot and every current target identity. Restore only captured pre-change payloads and remove only paths with both a recorded absence and exact Task 038 identity. Do not use wildcard deletion, Git reset/clean/checkout/restore, replacement from HEAD, or overwrite later Owner edits. Release artifacts under a task-created `dist` path may be removed only by their exact manifest after the service acceptance is inactive.

After rollback run version/build/tests, schema reads, systemd unit verification, staged install/uninstall, Framework/Builder/lifecycle/diverted-task checks, `bin/job --check`, Git status/index/ahead-behind, permissions/ACLs, snapshot checksums and acceptance residue checks. Retain snapshots, failure evidence and rollback report.



## Deliverables


- a frozen Version 1.0 capability inventory and BLOCKER/SHOULD/POST-1.0 release-gate matrix;
- a concrete Version 1.0 Release Policy, coherent `1.0.0-rc.1` identity or documented compatible equivalent, changelog and release notes;
- one reproducible versioned `linux-amd64` RC archive containing the QWSG binary, prefix-correct systemd user unit, manifest, checksum, installer/uninstaller, configuration example/guidance and ordinary-user documentation;
- checksum verification and deterministic artifact assembly tooling/tests without external dependencies or publication;
- safe clean install, enable/start guidance, exact uninstall and preserved-user-data policy;
- documented supported/experimental/unsupported platform contract and exact reboot/linger model;
- tested upgrade/reinstall/rollback and stored-state compatibility matrix;
- confirmed/fixed interactive Console startup redraw defect if reproducible, with EN/HU and fallback regressions;
- security/privacy review and bounded 50-cycle endurance/resource evidence;
- real isolated systemd lifecycle, failure, recovery, separate-process Console and cleanup evidence;
- complete Quick Start, operations, configuration, troubleshooting, upgrade, rollback, uninstall, security/privacy and known-limitations documentation;
- an explicit `READY FOR QWSG 1.0 RELEASE` or `NOT READY FOR QWSG 1.0 RELEASE` engineering decision;
- completed Task 038 history/archive and canonical idle lifecycle.

No provider transport, network interface, package ecosystem, updater/signing infrastructure, public license selection, Git tag/commit/push or external publication is a deliverable.



## Verification


- Validate repository identity, Task 037 idle baseline, Task 038 prompt/history identity, Framework 1.x, full Git status, empty index, main/HEAD/origin relationship, ownership/modes/ACLs and preservation of every unrelated Owner path before and after work.
- Run `make build`, `go test ./...`, `go test -race ./...` with bounded writable caches, `go vet ./...`, `make fmt-check`, Framework checks/tests, Builder tests, lifecycle/next-task/diverted-task tests, `bin/job --check`, `git diff --check`, staged-path and snapshot-manifest checks.
- Prove by imports/source/call tests that release work changes no canonical engine decision and that Guardian/Runtime/Runtime Service/Scheduler/Current State/Presentation/Console ownership remains unchanged.
- Verify VERSION, embedded `qwsg version`, archive filename, manifest, checksums, changelog and release notes agree exactly. Build twice with controlled commit/date and compare binary/archive/checksum identities under the documented reproducibility contract.
- Inspect the archive listing for safe relative paths, stable order/timestamps/modes, absence of symlinks unless explicitly validated, no secrets/private host data, and no repository/Go requirement at install/runtime.
- Install solely from an unpacked RC archive into a clean staged root. Verify collision refusal/backup behavior, executable/unit modes and ownership, prefix-correct ExecStart/ExecStopPost, checksum verification, no silent enable/start/linger change, and bounded uninstall with state preserved by default.
- Run `systemd-analyze verify --user` on the shipped and acceptance units. Audit `systemd-analyze security` where meaningful, unit limits/sandbox, no listener/socket, ordinary UID, state/config permissions, symlink rejection and privacy-safe journal output.
- Execute real isolated lifecycle: start; at least five recurring cycles; separate bare Console; graceful SIGTERM stop and stopped evidence; restart and running/degraded current evidence; SIGINT foreground test; SIGKILL/abnormal exit; correlated demotion; bounded restart/start-limit; invalid configuration; Current State publication failure; stale demotion; second-instance and explicit observe refusal; interrupted-cycle and checkpoint recovery.
- Verify no false healthy/running claim in missing, partial, failed, incompatible, stale, stopped or unavailable cases. Preserve last qualified engineering facts through lifecycle demotion.
- Test valid legacy Current Operator State 1.0/1.1 and current 1.2, Scheduler 1.0 and Guardian checkpoint 1.0 across reinstall/restart/rollback. Unknown schema, corrupt digest, unsafe path, wrong owner/mode and configuration mismatch must fail closed without overwrite.
- Reproduce interactive Console startup in a real PTY. Assert exactly one initial Overview, one redraw per accepted action/refresh, correct navigation/quit, bounded output, noninteractive single render, English/Hungarian tokens, attention overflow disclosure, truthful stale/partial/Guardian wording and no raw private diagnostics.
- Across at least 50 short isolated recurrence boundaries, sample after warm-up and assert RSS remains below the shipped 128 MiB service limit, tasks remain below 32, no monotonic FD/goroutine/task growth, sequential cycles, bounded snapshot retention, Scheduler results capped at 4096, checkpoint/current state within encoded limits, no restart loop and documented CPU behavior. Record exact observations and environment.
- Verify user-service enablement and default target linkage without changing linger. Document `loginctl show-user` evidence and run isolated manager restart/recovery where safe. If no disposable reboot occurs, verify the Owner-run reboot script/procedure is complete and ensure all support text says reboot acceptance remains operator-run before public release.
- Complete the archive-only user journey: checksum, install, version, first baseline, second full observation, Guardian start/cycles, Console, stop, restart/recovery, bounded failure, upgrade/reinstall, compatibility, rollback, post-rollback operation, uninstall, owned-artifact absence and preserved-data report.
- Audit privacy/security: no credentials, secret values, raw config, usernames, private paths or host identifiers in artifacts/logs/Console; no network listener, remote command, arbitrary shell, remediation, privilege escalation, root runtime or unbounded path deletion.
- Verify Task 038 acceptance unit/config/state/install roots are absent afterward; real QWSG service/state and linger are unchanged; implementation and service snapshots still pass SHA-256; rollback remains executable.
- Confirm documentation contains no pre-alpha contradiction, unsupported distro/architecture, implemented provider claim, public-license claim, physical-reboot claim, checksum-authenticity claim or external publication claim.



## Documentation Updates


Update only directly affected sections and create missing bounded release documents:

- `ai/core/12_RELEASE_POLICY.md` for Version 1.0/RC identity, gates, artifacts, checksums, reproducibility, compatibility, approval, publication and rollback;
- `VERSION`, `CHANGELOG.md`, `LICENSE` only to the extent authorized (retain the temporary proprietary terms unless the Owner separately decides otherwise), and version/release tests;
- `README.md` and a concise release Quick Start for ordinary-user installation and first useful result;
- `docs/installation/INSTALL.md` plus dedicated upgrade/rollback/uninstall and troubleshooting/diagnostics guidance where separation improves usability;
- a versioned support matrix and known Version 1.0 limitations document;
- `docs/architecture/OPERATIONAL_GUARDIAN_SERVICE.md`, Current State/Presentation/Console/Configuration documents only for confirmed release behavior, compatibility and diagnostics;
- `docs/PRODUCT_DEFINITION.md`, `docs/PRODUCT_ARCHITECTURE.md`, `docs/FUNCTIONAL_SPECIFICATION.md`, architecture gate/traceability records and `ai/core/13_ROADMAP.md` to reconcile current implemented Version 1.0 scope and explicitly deferred provider/network/commercial families;
- `ai/core/04_ARCHITECTURE.md`, `ai/core/05_SYSTEM_MAP.md`, and `ai/core/07_ENGINEERING_HISTORY.md` for the final supported local topology and Task 038 milestone;
- English operator documentation and existing Hungarian user guidance/catalogs for actual commands/status meanings, without translating engineering records;
- release artifact manifest, release notes, sanitized acceptance report, Task 038 prompt/history/archive and exact rollback evidence.

Record every actual file, command class, artifact name/hash, measurement, platform fact, limitation and justified omission in Task 038 history. Do not store raw host evidence, journal content, usernames, temporary absolute paths, secrets or publication credentials.



## Completion Criteria


Task 038 is complete only when:

- Version 1.0 scope is frozen to implemented local QWSG and every remaining item is classified BLOCKER/SHOULD/POST-1.0 without optional-scope inflation;
- every BLOCKER is either corrected and verified or produces an explicit `NOT READY` decision; `READY` is forbidden with an unresolved technical, safety, installation, compatibility, platform or documentation blocker;
- one coherent RC version is embedded in the executable and all artifact/release metadata;
- an ordinary Ubuntu 24.04 linux-amd64 operator can verify and install the prebuilt archive without repository or Go knowledge, and installation never silently enables, starts, elevates runtime or changes lingering;
- the shipped unit is prefix-correct, validates under systemd, runs non-root within documented security/resource bounds and exposes truthful lifecycle evidence;
- the complete real user journey passes from RC archive through observation, continuous Guardian, Console, failures, upgrade, rollback and uninstall with preserved-data reporting;
- valid supported stored state survives restart, reinstall, upgrade and declared rollback; corrupt/incompatible/unsafe state fails closed and remains recoverable without silent migration or deletion;
- actual graceful/abnormal systemd lifecycle, restart limits, stale demotion, single-writer behavior, interrupted-cycle recovery and no-false-running behavior pass in isolated acceptance;
- the 50-cycle endurance test shows bounded memory, CPU, tasks, goroutines, FDs, snapshots and state files with no monotonic leak or restart loop;
- security/privacy review finds no release-blocking privilege, path, listener, command, secret, log, installer or uninstall defect;
- the duplicated interactive initial rendering is either reproduced and fixed with regression evidence or disproved with PTY evidence and documented; other Console behavior remains compatible and localized;
- the support matrix claims only Ubuntu 24.04 LTS, systemd 255+, linux-amd64 and verified filesystem/user-manager assumptions; physical reboot is claimed only if executed on a disposable supported host, otherwise exact Owner-run acceptance remains a disclosed pre-publication procedure;
- Release Policy, README/Quick Start, installation, configuration, operations, diagnostics, upgrade, rollback, uninstall, privacy/security, support and limitations documentation are coherent for an ordinary user;
- release artifacts are deterministic under documented inputs, manifest/checksums pass, staged install/uninstall leaves only explicitly preserved data, and no acceptance residue remains;
- full/race/vet/format/build, systemd, artifact, install, compatibility, security, endurance, Framework, Builder, lifecycle, diversion, Git, snapshot and rollback validations pass;
- unrelated Owner content, real QWSG state/service, linger and external infrastructure are unchanged; index remains empty; no stage, commit, tag, push or publication occurred;
- Task 038 history records the explicit release decision, exact evidence and known limitations, the prompt is archived, and `bin/job --check` reports canonical idle with Task 038 latest.

A technical `READY FOR QWSG 1.0 RELEASE` decision does not authorize a Git tag, signing, public distribution or public release. The temporary proprietary LICENSE must be explicitly confirmed or replaced by the Owner before public publication. Concrete notification transports, Dashboard/API, fleet, remote, cloud, remediation, AI, licensing enforcement and packaging ecosystems are known POST-1.0 capabilities and do not require Task 039.


## Owner Approval Requirements

Approved by Project Owner through the Engineering Task Builder on 2026-08-10 UTC.

The structured task definition has been explicitly approved for implementation. Further scope changes require explicit Project Owner approval.
