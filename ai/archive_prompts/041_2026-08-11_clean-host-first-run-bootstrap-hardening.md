# Current Engineering Task 041: Clean-Host First-Run Bootstrap Hardening

## Task Metadata

- Task ID: `041`
- Task slug: `clean-host-first-run-bootstrap-hardening`
- Status: `complete with disclosed validation-environment limitation`
- Date opened: `2026-08-11` UTC
- Human authority: Project Owner
- Owner or lead-developer communication language: Hungarian

## Title

Clean-Host First-Run Bootstrap Hardening


## Objective

Make the installed QWSG command safely bootstrap its private per-user state on a completely fresh supported Ubuntu 24.04 account, so the first `qwsg observe` creates the required hierarchy, persists a truthful baseline and publishes Current Operator State without manual directory preparation.

Correct only the confirmed first-run directory-creation and closely coupled partial-check diagnostic defects. Preserve the accepted Task 039 product behavior, Task 040 release process, strict storage security, installer privilege boundary and truthful partial Inventory semantics. Task 041 ends after source and temporary artifact-level clean-account acceptance; it does not refresh the public or canonical RC identity, overwrite RC.2, install on the Owner clean VPS, or authorize publication.



## Scope

- Correct the default direct-CLI state bootstrap used by `qwsg observe` when `$HOME/.local`, `$HOME/.local/state` and `$HOME/.local/state/qwsg` are all absent.
- Replace the single-level final-root creation with one narrowly reusable secure recursive bootstrap that creates only missing hierarchy and validates the final QWSG root as a clean absolute, current-user-owned, non-symlink directory with mode `0700`.
- Preserve every existing parent directory's owner and mode. Never chmod, chown, replace or otherwise mutate an existing `$HOME`, `$HOME/.local`, `$XDG_STATE_HOME` or other ancestor.
- Retain rejection of symlink path components, non-directory targets, wrong ownership, permissive QWSG-root modes, non-absolute/unclean explicit roots and unsafe races to the extent enforced by the existing bounded local-store contracts. Do not weaken fail-closed behavior.
- Keep private QWSG subdirectories at `0700` and metadata, snapshot, lock, checkpoint and Current Operator State files at `0600` through the existing Inventory Store, Guardian and Current State writers.
- Keep `packaging/release/install.sh` behavior unchanged: it installs immutable system artifacts and does not create per-user runtime state.
- Preserve systemd `StateDirectory=qwsg` behavior. The correction is for direct ordinary-user CLI first use and may reuse the same validated state-root contract.
- Preserve valid partial Inventory as truthful evidence. The absence of `/usr/local/go/bin/go` must leave the existing `components` capability unavailable; do not install Go or mark it available.
- Make first-run `observe` persist a partial baseline and publish a partial/unknown or degraded Current Operator State when the snapshot is valid, matching the existing `publishObserveBootstrap` intent.
- Apply the smallest compatible `qwsg check` correction: permit publication of valid correlated partial Inventory/Snapshot evidence, consistent with observe bootstrap, while retaining command-completion, contract, type, identity, payload, integrity and freshness gates. Do not reinterpret partial as complete.
- Ensure semantic publication ineligibility, bootstrap failure, permission failure and unsafe-path failure cannot be mislabeled as unreadable/corrupt state. Add only bounded privacy-safe tokens required by these paths; never expose raw Go errors, host paths, identity, configuration values or secrets.
- Add focused regression coverage in `cmd/qwsg/main_test.go` and the directly affected private-store tests. Modify `internal/operatorstate/store.go` only if required to expose/reuse its existing secure root-initialization contract; do not introduce a storage abstraction or new persistence layer.
- Run temporary artifact-level acceptance from a newly built package in isolated roots without overwriting `dist/qwsg-1.0.0-rc.2-linux-amd64.tar.gz` or its sidecar. The test package may retain source version `1.0.0-rc.2` only as non-release validation evidence; it is not a replacement RC.
- Update Task 041 history and only directly affected operator/release documentation, archive Task 041, and return to canonical idle without creating a successor.



## Out of Scope

- No installer-created per-user state, installer architecture change, system-wide runtime state, root-owned Guardian operation or change to systemd state-directory semantics.
- No new collector, capability, probe, dependency, Go installation, Git requirement or false complete Inventory result.
- No redesign of Inventory, Snapshot, Current State, Command, Runtime, Guardian, Alert, Console, Scheduler, Policy, Report or release architecture.
- No change to accepted Task 039 large-Report, Runtime diagnostic, Console refresh, Attention or lifecycle-freshness behavior.
- No Dashboard, REST API, provider, transport, fleet, remote management, remediation, AI, licensing system, package manager, updater or persistence platform.
- No mutation of `VERSION`, release identity, release notes, Quick Start artifact names, RC.1/RC.2 historical evidence, canonical RC.2 archive or checksum.
- No installation to real `/usr/local`, activation or mutation of the real QWSG service/state/configuration, user manager or linger state.
- No operation against the Owner clean VPS and no clean-VPS reboot acceptance.
- No Task 042 or successor-task creation, release refresh, tag, signing, stage, commit, push, upload, publication or announcement.
- Stop and report if valid partial evidence cannot be published without changing frozen canonical semantics, or if secure bootstrap requires weakening symlink, ownership, mode or path validation.



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

- Begin only when `bin/job --check` reports canonical idle with Task 040 as the unique latest completed archived prompt/history pair and no Task 041 lifecycle collision.
- Verify repository root and Framework markers, `main`, exact HEAD, canonical HTTPS origin, `0/0` upstream relationship, empty index, complete Git status, target ownership/modes/ACLs, filesystem capacity and exact preservation requirements for all unrelated Owner-owned modified/untracked content.
- Read the complete Task 039 and Task 040 prompts/histories, RC.2 release notes and acceptance record, `docs/release/SUPPORT.md`, `docs/release/QUICK_START.md`, installer, systemd unit, state architecture documents and relevant storage/CLI tests before changing a target.
- Reconfirm from source that `localStateRoot` resolves `QWSG_STATE_DIR`, then `XDG_STATE_HOME/qwsg`, then `$HOME/.local/state/qwsg`; `ensureLocalStateRoot` uses single-level `os.Mkdir`; and clean missing parents cause pre-Inventory `ENOENT` collapsed to `evaluation_failed`.
- Reconfirm that `operatorstate.Store.Publish`, Inventory Store, Guardian Store and Scheduler Store already create private layouts with atomic `0600` files, while `publishCheck` rejects valid partial stages before opening Current State Store and the fallback diagnostic incorrectly reports `state_unreadable`.
- Reconfirm that `components` probes only the allowlisted Go binary and clean release hosts without Go therefore produce truthful unavailable capability and partial aggregate Inventory.
- Record exact hashes for `cmd/qwsg/main.go`, `cmd/qwsg/main_test.go`, `internal/operatorstate/store.go` and its tests, installer, unit, relevant docs, Task 041 prompt/history targets, VERSION and canonical RC.2 artifacts. Verify canonical RC.2 archive SHA-256 remains `0694fb7f382ea1b373aaf0b6f0171a3fc580c491c1af3b8657a3bb8697ed897b`.
- Verify Ubuntu 24.04/linux-amd64 development prerequisites without installing anything. Record Go, Make, filesystem and systemd versions and ensure no Task 041 process, temporary unit or acceptance root exists.
- Run pre-change focused CLI/operator-state/Inventory Store/Guardian tests, complete tests, race tests with writable bounded caches, vet, formatting, Framework, Builder, lifecycle, diverted-task, `bin/job` and Git whitespace checks. Stop on a material baseline contradiction.



## Snapshot Requirements

- Before modifying any implementation target, create a unique mode-`0700` Task 041 implementation snapshot outside the repository.
- Capture exact pre-change bytes, modes, owners and ACLs for `cmd/qwsg/main.go`, `cmd/qwsg/main_test.go`, `internal/operatorstate/store.go`, all directly affected tests/docs, `packaging/release/install.sh`, the systemd unit, VERSION, Task 041 prompt/history and canonical RC.2 artifact metadata. Record absence for any newly authorized test/documentation path.
- Record repository root, branch, HEAD, remotes, upstream relationship, full status, empty index, relevant tools/platform, storage-root contracts, clean-HOME failure precondition, existing QWSG processes/services and exact intended temporary roots.
- Include deterministic SHA-256 manifests, a readable bounded archive, retention notes and guarded restore instructions that refuse active processes, changed target hashes, later Owner edits, ambiguous ownership, unsafe paths or broad cleanup.
- Before artifact acceptance, record the unique temporary build, extraction, staged-install, HOME, XDG state, explicit-state and store paths, their initial absence, executable hash and exact cleanup boundary. Retain the implementation snapshot after completion; remove only verified Task 041 temporary roots after process absence.



## Risk Assessment

- **Symlink traversal or path substitution — critical:** validate clean absolute paths and every existing component before recursive creation; reject symlinks and revalidate the final root after creation. Never follow a user-supplied symlink as a convenience.
- **Ownership or permission weakening — critical:** require the QWSG root to be owned by the effective user and exactly `0700`; preserve `0600` file writers; do not chmod/chown ancestors or accept permissive existing roots.
- **First-use race or partial hierarchy — high:** use the existing bounded creation/validation contract, tolerate only safe already-created directories, revalidate before use and fail with a privacy-safe token on conflict.
- **Semantic drift in partial evidence — high:** remove only the `check` stage-completeness eligibility contradiction; retain all structural, execution, typed-payload, correlation, canonical validation and freshness checks and keep overview completeness partial.
- **Misleading diagnostics — high:** classify bootstrap/publication eligibility separately from unreadable/corrupt persisted state without emitting raw filesystem or Go errors.
- **Installer privilege expansion — high:** assert installer bytes unchanged and ensure tests create state only by running the ordinary-user binary.
- **False artifact acceptance — high:** build and install only a temporary artifact into unique roots, remove Go/Git from runtime PATH, require a genuinely empty HOME, and never claim that the historical RC.2 binary contains Task 041.
- **Owner data or release artifact loss — critical:** preserve unrelated dirty-tree content and canonical RC.2 bytes; prohibit broad Git recovery, broad deletion and writes outside exact task/snapshot/temporary paths.
- **Regression of accepted release behavior — high:** rerun focused Task 039 and applicable Task 040 release/package gates; stop on product behavior changes outside first-use scope.



## Planned Work

1. Validate canonical idle, Task 040 baseline, clean-host evidence, exact repository/environment state, relevant source contracts and all pre-change gates.
2. Create and verify the bounded Task 041 implementation snapshot and guarded rollback procedure before editing code.
3. Introduce the smallest reusable secure root-ensure operation around the existing Current State private-directory checks. It shall recursively create missing parents, set only the final QWSG root to `0700`, require current-user ownership, reject unsafe/symlink/unclean paths and leave existing parents unchanged.
4. Route the direct default observation-store bootstrap through that operation before lock/store use. Preserve explicit `QWSG_STORE`, `QWSG_STATE_DIR` and `XDG_STATE_HOME` precedence and existing fail-closed behavior.
5. Align `publishCheck` with `publishObserveBootstrap` for valid partial Inventory/Snapshot evidence by removing only the stage-complete eligibility requirement. Retain exact profile, execution, stage, schema, type, snapshot identity, payload equality, Inventory validation and freshness requirements.
6. Add only necessary privacy-safe diagnostic classification for bootstrap or semantic publication failures. Prove absent first-use state is not described as corrupt/unreadable and existing corrupt/unsafe/permission cases retain their canonical tokens.
7. Add deterministic tests for fully empty HOME, missing `.local` hierarchy, partial Inventory baseline/current-state publication, exact modes/ownership, repeat use, valid existing state, symlink components, unsafe root kind/mode and feasible wrong-owner coverage. Assert existing parents are not chmod/chowned.
8. Run focused and complete source validation, then build a temporary release package with the existing canonical release script into a unique non-`dist` output directory. Verify canonical RC.2 repository artifacts remain byte-identical.
9. Stage-install the temporary package only under an isolated destination. Execute its binary with a truly empty synthetic HOME, no QWSG/XDG state variables, a PATH without Go/Git, no repository working directory and no prior state: first `observe`, second `observe`, then noninteractive Console load.
10. Verify first observe creates a partial baseline and Current State with exact private modes; second observe enters normal evaluation as evidence permits; Console reads the published state; no manual preparation, hidden host dependency or false completeness occurs. Re-run symlink/permission negatives with the packaged binary where practical.
11. Run all final build/test/race/vet/format, Framework, Builder, lifecycle, diverted-task, package, installer-boundary, security/privacy, snapshot and Git validations. Record sanitized evidence, clean exact temporary roots, preserve snapshots and RC.2, complete/archive Task 041 and return to canonical idle.



## Rollback Plan

Stop rollback if any Task 041 binary, test, Guardian, package, staged installer or temporary service is active, or if repository identity, target hashes, Owner-owned content, canonical RC.2 bytes or snapshot manifests differ from recorded facts.

Restore only snapshot-listed Task 041 implementation and documentation targets after proving their current bytes are Task 041-owned and no later Owner edit exists. Remove a Task 041-created path only when the snapshot recorded absence and its current hash matches recorded Task 041 output. Preserve all Inventory, Current State, Guardian and Owner data outside exact synthetic acceptance roots.

Never use Git reset, clean, checkout, restore, wildcard deletion or recursive removal of HOME, repository, workspace, `/tmp` generally or QWSG real state. Remove only exact unique temporary roots after process absence and identity checks. After rollback, rerun focused/full/race/vet/format and governance gates, verify canonical idle Task 040, unchanged installer/unit and exact RC.2 archive/sidecar hashes, and confirm no Task 041 residue.



## Deliverables

- secure recursive ordinary-user bootstrap for an absent default QWSG state hierarchy;
- final QWSG root `0700`, QWSG private subdirectories `0700` and private files `0600`, with existing parent ownership/modes preserved;
- unchanged installer and systemd privilege/responsibility boundaries;
- first `qwsg observe` baseline and Current Operator State publication with truthful partial Inventory on a no-Go host;
- second `qwsg observe` normal canonical evaluation path as available evidence permits;
- valid partial `qwsg check` publication consistent with existing bootstrap semantics, without false completeness;
- bounded truthful diagnostics that do not call semantic/bootstrap failures unreadable or corrupt;
- empty-HOME, repeat-use, partial-evidence, ownership/mode and unsafe-path regression coverage;
- temporary packaged-binary clean-account acceptance without Go, Git, checkout, prior state or manual directory setup;
- verified unchanged canonical RC.2 artifacts and no release-identity mutation;
- completed Task 041 history/archive, retained rollback snapshot and canonical idle closure without a successor.



## Verification

- Verify initial/final root, branch, HEAD, origin, upstream relationship, complete Git status, empty index, exact Task 041 diff, target ownership/modes/ACLs, snapshot hashes and preservation of unrelated Owner content.
- Run focused tests for `cmd/qwsg`, `internal/operatorstate`, `internal/inventorystore`, `internal/guardian`, `internal/operatorconsole`, `internal/presentationmodel`, `internal/pipeline` and any directly affected application package.
- Run `make build`, `go test ./...`, bounded-cache `go test -race ./...`, `go vet ./...`, `make fmt-check` and `git diff --check` before and after implementation as required.
- Run Framework validation/tests, Builder tests, lifecycle/next-task tests, diverted-task tests/audit and `bin/job --check` at applicable active/idle states.
- In a synthetic HOME containing none of `.local`, `.local/state` or `.local/state/qwsg`, run first observe and assert creation succeeds without preparation; QWSG root/subdirectories are current-user-owned `0700`; metadata, snapshots, locks and Current State are regular current-user-owned `0600` files.
- Assert existing ancestor owners and modes are byte-for-byte/stat-equivalent before and after bootstrap, including a conventional non-private `.local`; bootstrap must modify only missing descendants and the QWSG root.
- Verify default, `XDG_STATE_HOME` and explicit `QWSG_STATE_DIR` roots are clean absolute and deterministic. Reject relative/unclean roots, symlink components, non-directory roots, permissive QWSG roots and feasible wrong-owner fixtures without creating or modifying protected targets.
- Verify first partial observe produces a valid Inventory Store baseline and Current State whose completeness/condition remains honestly partial/unknown or degraded; it must not require the `components` collector or Go.
- Verify repeated startup does not damage valid state, does not reset metadata/retention, and takes the second observe through normal baseline comparison/evaluation as available evidence permits.
- Verify `qwsg check` publishes valid correlated partial evidence and retains partial presentation. Corrupt, incompatible, unsafe and permission errors must remain distinct; semantic ineligibility and bootstrap failures must not emit `state_unreadable` or raw errors.
- Assert `packaging/release/install.sh`, systemd unit, VERSION, release notes, Quick Start and canonical RC.2 archive/sidecar are unchanged. Independently recheck RC.2 SHA-256 `0694fb7f382ea1b373aaf0b6f0171a3fc580c491c1af3b8657a3bb8697ed897b`.
- Build a temporary archive in a new output root using the existing reproducible process; verify manifest/checksum, safe contents and staged install. Do not overwrite or relabel canonical RC.2 and do not claim the temporary package is an accepted RC.
- Run the staged binary from outside the repository with empty HOME, no prior QWSG state, no Go/Git in PATH and no repository dependency. Require successful first observe, persisted baseline/Current State, second observe progression and later `qwsg` Current State load.
- Verify privacy: no raw errors, synthetic or real HOME paths, usernames, hostnames, IPs, secret/config values or unbounded evidence enter operator diagnostics or committed documentation.
- Verify no real `/usr/local` write, real service/state/config/user-manager/linger mutation, clean-VPS contact, dependency install, license change, Git stage/commit/tag/push or publication occurred.
- Remove only exact temporary build/install/HOME/state roots after executable/process absence; retain the verified implementation snapshot and ensure Task 041 archives to canonical idle as latest completed baseline.



## Documentation Updates

- Update `cmd/qwsg/main.go`, `cmd/qwsg/main_test.go` and only directly affected private-store implementation/tests required for secure bootstrap and partial publication.
- Update `docs/release/TROUBLESHOOTING.md` only if a new bounded first-run diagnostic token requires operator meaning; do not add manual `mkdir` as the product solution.
- Update the directly affected Current Operator State or Operational Guardian architecture document only if necessary to state that ordinary-user first use recursively creates and validates the private root. Do not rewrite architecture.
- Update `CHANGELOG.md` under the existing unreleased/private RC.2 development baseline only if canonical changelog policy requires recording the accepted correction before the later RC refresh; do not change VERSION or label the historical RC.2 artifact as fixed.
- Update Task 041 history throughout execution, the concise engineering milestone index after completion, and archive the completed Task 041 prompt. Do not create successor release notes, acceptance documents or Task 042.



## Completion Criteria

Task 041 is complete only when all of the following are true:

- lifecycle begins idle with Task 040 latest complete and ends idle with Task 041 latest complete, with no successor task or collision;
- a fully empty supported-user HOME requires no manual directory preparation and first `qwsg observe` securely creates the QWSG hierarchy, Inventory Store, baseline Snapshot and Current Operator State;
- QWSG root/private directories are current-user-owned `0700`, private files are `0600`, existing ancestors retain their owners/modes, and symlink, unsafe path, wrong-owner and permissive-root checks remain fail-closed;
- installer and systemd responsibility boundaries are unchanged and no installer-created arbitrary user state is introduced;
- valid partial Inventory, including unavailable `components` on a host without Go, remains partial but can establish a truthful first baseline and operator state;
- `qwsg check` no longer calls valid partial evidence unreadable/corrupt and either publishes it under the selected compatible semantics or emits a specific privacy-safe eligibility result without semantic drift;
- first-use bootstrap failures, permission failures and unsafe paths have bounded truthful diagnostics with no raw private data;
- repeated observe uses existing state deterministically, second observe reaches normal evaluation as evidence permits, and a later Console loads the published Current State;
- source tests and a temporary installed-artifact empty-account acceptance both pass without Go, Git, repository checkout, prior state, manual `mkdir`, real-host installation or clean-VPS operation;
- complete focused/full/race/vet/format, Framework, Builder, lifecycle, diversion, security/privacy, package and Git validations pass, with exact cleanup and retained rollback evidence;
- accepted Task 039 behavior, release architecture, VERSION, licensing, installer/unit and canonical RC.1/RC.2 artifacts remain unchanged;
- no stage, commit, tag, push, signing, upload, publication, public release claim or clean-VPS acceptance occurred;
- history records actual evidence and any limitation. No additional product task is required before the separately authorized next-RC metadata/reproducibility refresh unless Task 041 validation discovers a new bounded product blocker.


## Owner Approval Requirements

Approved by Project Owner through the Engineering Task Builder on 2026-08-11 UTC.

The structured task definition has been explicitly approved for implementation. Further scope changes require explicit Project Owner approval.
