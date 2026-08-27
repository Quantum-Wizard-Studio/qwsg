# Task History 066: Approved Task

## Task metadata

- Task ID: `066`
- Task slug: `qwsg-1-2-0-final-release-distribution`
- Status: `complete — BLOCKED by QWSG-066-F001 release archive mode nondeterminism`
- Date generated: `2026-08-27` UTC
- Human authority: Project Owner Attila
- Preferred owner communication language: English
- Related prompt: `ai/prompts/066_CURRENT_TASK.md`

## Lifecycle state

The Engineering Task Builder generated and transactionally installed this matching prompt/history pair from validated structured owner input. The Project Owner explicitly authorized execution. Required reading, canonical lifecycle validation, starting-state verification, remote synchronization and the pre-modification snapshot gate passed; release-readiness audit work is active.

## Starting state

- Canonical root, Framework 2.0.0, branch `main`, remote URL, active Task 066 prompt/history identity and project markers validated.
- Starting `HEAD == fetched origin/main == f726d84632ebce3f2be72101b583af5beadc857e`, with `0 0` ahead/behind. The index and tracked worktree were clean; only the expected Builder-created Task 066 prompt/history were untracked.
- Task 065 is canonically complete. `VERSION=1.2.0-rc.2`; final private RC.2 source commit is `c260dc18c2004473ec55496d16e66718fd128865`; frozen archive SHA-256 is `a34be8b18f80d877c0ccfd69dc9d9e9f197fc35fa765cdf1d5c0d72e2cb0a554` and sidecar SHA-256 is `23218e15a85ab5ee644031bf5ad6469d33a247fc1e271f550de1d52d83e88e11`.
- The Go-owned, interface-neutral guided installer, narrow privileged shell helper, EN/HU/DE catalogs, notification guidance, manual/notify policy, Guardian activation/readiness, native update and rollback foundations are present and consistent with the Task 065 record.
- Fresh remote verification found no local or remote `v1.2.0*` tag. Anonymous Forgejo API and release-page requests for `v1.2.0` returned HTTP 404; no final public release or asset existed.
- Ignored historical backups, local task-source documents, build/dist output and binary remained excluded. The Owner-owned QWCS migration draft had already been moved to protected storage outside the repository under separate authority and was absent from the worktree before Task 066 installation.

## Snapshot

- Protected payload root: `/tmp/qwsg-task066-execution.AJSQ69`, mode `0700`; retain through Project Owner acceptance and release rollback-window closure.
- `repository-head.tar.gz` captures complete tracked HEAD and is SHA-256 `292f063d1989dd4d5712019f482dc570dee979fa6bb09f71ae609539edb1f07d`.
- Separate mode-`0600` before-images capture the Builder-created Task 066 prompt and history. `SHA256SUMS`, deterministic `MANIFEST.tsv`, `START_STATE.md` and collision-safe `ROLLBACK.md` are included.
- `sha256sum -c` passed for every payload object; compressed-archive listing/readability passed; path traversal, absolute paths, symlinks and special-file surprises were absent. Restore instructions require protected temporary extraction and exact reviewed paths only, never extraction over the live worktree or broad Git cleanup.

## Work performed

Established the exact starting state, independently fetched canonical remote `main`, verified absence of final tag/Release, and created and verified the mandatory pre-modification snapshot. No release identity, source, artifact, host, tag, or publication mutation has occurred yet.

- First full-validation attempt stopped at `make engineering-test`: `scripts/test-build-contract.sh` still required `1.2.0-rc.1` while canonical `VERSION` is `1.2.0-rc.2`. Classified `TEST OR ACCEPTANCE DEFECT` after source/history review; the contract test was narrowly advanced to RC.2 before retesting. No product or candidate bytes changed.
- Full governance and automated product validation then passed with isolated writable caches. A literal intermediate `go test ./...` invocation selected the sandbox-read-only default Go cache and failed before product execution; classified `ENVIRONMENTAL ISSUE`, corrected to the canonical `/tmp` caches, and both full and repository-wide race suites passed.
- Rebuilt RC.2 twice with source commit `c260dc18c2004473ec55496d16e66718fd128865` and epoch `1787839370`. Both archives were byte-identical to each other and to the Task 065 frozen candidate at SHA-256 `a34be8b18f80d877c0ccfd69dc9d9e9f197fc35fa765cdf1d5c0d72e2cb0a554`; sidecars matched SHA-256 `23218e15a85ab5ee644031bf5ad6469d33a247fc1e271f550de1d52d83e88e11`. Outer checksum, safe archive paths/types, internal manifest, required files, systemd static verification exit, embedded version and provenance passed.
- **QWSG-066-F001 — PRODUCT/FRAMEWORK DEFECT, RELEASE BLOCKER.** Package inspection found RC.2 directories mode `0775`, ordinary documentation/unit/configuration files mode `0660`, and generated metadata mode `0664`. The release builder normalizes only executable files and inherits every other mode from the ambient worktree/umask.
- An exact `c260dc1` source export with identical logical bytes, commit, epoch and toolchain but ordinary `0644` non-executable modes produced archive SHA-256 `5b32df3b090658cfb9a08a7d670848c65af4d5d048dc053e3ad0973d11f0082a`, not the frozen hash. Extracted payload content compared byte-identical; representative `README.md` modes were `0660` frozen versus `0644` clean export. Cross-mode deterministic construction and least-privilege archive permissions therefore fail.
- Per the approved STOP boundary, no final identity, tag, Forgejo Release/asset, OVH mutation, Contabo mutation or announcement was attempted. Task 065's prior clean-host claims remain valid but missing Task 066 clean/full/update/restorative rollback/physical reboot/notification/published-download gates are not PASS.
- Added `docs/release/ACCEPTANCE_1.2.0.md` with the blocked decision, evidence and smallest remediation: normalize package modes before manifest/tar creation, add cross-umask/mode reproducibility tests, issue new private RC.3 provenance and repeat every mandatory acceptance gate.

## Verification

- Builder input, metadata, prompt/history identity, approval state and lifecycle installation: PASS.
- `bin/job --check`, `ai/scripts/next-task.sh --check` and Framework validation: PASS.
- Fresh remote branch comparison and remote tag query: PASS, exact expected baseline.
- Snapshot checksum, mode, archive readability/type/path safety and rollback-document review: PASS.
- Framework/lifecycle/Builder/diversion suites: PASS (`25`, `29`, `49`, `36` assertions); Framework v2 and bounded diagnostic suites: PASS (`15`, `11`).
- Full Go, repository-wide race, vet, formatting, build contract, release plumbing, shell syntax and Git whitespace: PASS after the classified test/cache corrections above.
- Frozen RC.2 same-environment two-build and Task 065 byte identity: PASS. Archive checksum, manifest, provenance and safe layout: PASS.
- Cross-mode equivalent-source deterministic archive and least-privilege archive modes: FAIL — QWSG-066-F001, release-blocking.
- Clean OVH final acceptance: NOT EXECUTED after mandatory STOP; prior Task 065 evidence is partial input only.
- Full Contabo coexistence: NOT EXECUTED after mandatory STOP.
- Final update/restorative rollback acceptance: NOT EXECUTED after mandatory STOP; local update/rollback tests passed but are insufficient for release.
- Forgejo final wget/curl distribution: NOT EXECUTED; tag/Release/assets remain absent.

## Rollback

Use only `/tmp/qwsg-task066-execution.AJSQ69/ROLLBACK.md` after verifying `SHA256SUMS`, live identities and exact target collisions. Repository restoration is path-bounded; acceptance hosts require their own pre-state and product rollback evidence. Public refs/Releases/assets are never rewritten as a local rollback shortcut.

## Completion state

`complete — BLOCKED; QWSG 1.2.0-rc.2 was not promoted because archive modes are ambient-state-dependent and unnecessarily group-writable; no tag, Release, public artifact or VPS mutation occurred`
