# Task History 053: Approved Task

## Task metadata

- Task ID: `053`
- Task slug: `qwsg-1-1-0-rc-3-clean-host-acceptance-release-readiness`
- Status: `complete with disclosed limitations — NOT READY FOR QWSG 1.1.0 RELEASE`
- Date generated: `2026-08-21` UTC
- Human authority: Project Owner
- Preferred owner communication language: English
- Related prompt: `ai/prompts/053_CURRENT_TASK.md`

## Lifecycle state

The Engineering Task Builder generated and transactionally installed this matching prompt/history pair from validated structured owner input. Explicit approval was recorded. The Project Owner started Task 053 through the canonical `job` workflow on 2026-08-21 UTC. Phase A performed the read-only source/release audit and prepared the RC.3 acceptance protocol, evidence ledger and version-aware plumbing needed for a commit-pure candidate source. Gate A2 integrated that source. Gates B/C/D then stopped on `QWSG-053-F001`. On 2026-08-21 UTC the Owner authorized truthful closure as complete with disclosed limitations and `NOT READY FOR QWSG 1.1.0 RELEASE`; this is not successful release acceptance.

## Starting state

Baseline commit and direct remote `main` are
`3bf5a8e26e32ac1489dcdfecad5b086e4141cd91`, ahead/behind is `0/0`, branch is
`main`, index and tracked tree were clean, and the only unrelated untracked
path was the excluded Owner-owned migration blueprint recorded by stat metadata
only. `VERSION` is `1.1.0-rc.3`; no RC.3 archive, sidecar or tag exists; Task
052 product commit `6bb5b62957e54e0ac3377ce1b85593408c341873` is an ancestor.

## Snapshot

Private mode-0700 snapshot:
`/tmp/qwsg-task053-phase-a.YQYcaQ`. It contains a readable tracked-HEAD
archive, Task 052 archive/history, Task 053 Builder source/input and pre-edit
prompt/history, target-absence records, tool/baseline/ACL metadata, protected
hashes and bounded restore instructions. Owner content is excluded except for
stat metadata; no credentials or external identity are present.

## Work performed

Audited Tasks 049–052, RC.1/RC.2 protocols and acceptance records, current
RC.3 source/release metadata, packaging, operator documentation, Smart Install,
setup, userruntime/userservice/runner/assessment and systemd/install/uninstall
surfaces. The integrated product source is suitable for acceptance-source
preparation without further product correction. Prepared:

- `docs/release/ACCEPTANCE_PROTOCOL_1.1.0-rc.3.md` with 25 independently
  bounded checkpoints and Owner/privacy/continuation gates;
- `docs/release/ACCEPTANCE_1.1.0-rc.3.md` as an empty provenance/evidence ledger
  retaining F001 OPEN/BLOCKING and historical F002/F003;
- RC.3-aware protocol/ledger assertions in `scripts/test-release-plumbing.sh`.

At Gate A1 no candidate bytes were built and no external or credential action
occurred.

After Gate A2 integrated candidate-source commit
`3e7d2f9d543078b49d1afa522dc6bf3baba1c949`, Owner-authorized Gates B/C/D
created two independent private mode-0700 Git exports and builds. Both release
builds used commit-derived epoch `1787314469` (`2026-08-21T12:14:29Z`) and the
exact full commit. Their binary, manifest, archive and sidecar bytes matched.
The archive is `2949417` bytes with SHA-256
`c3ba763701b7ee0340d4928b21c23276dfdc083536b08814157366310629a0cc`.
Sidecar SHA-256 is
`41be7497c427bae4dbccf240e846a69b7cad7c3fe5d9f8c2d8071b088f08e343`,
manifest SHA-256 is
`33751c0f050b0304c9a0840a50c6f32287d2ada4793d12706f49971b93fa7dc3`,
and binary SHA-256 is
`eb9926c99e90e2146ee8657797c71d1382c316e372af62186fc67d1bc6c3b044`.
Sidecars and all 18 manifest entries verified. The package has one canonical
root, 25 expected members, 19 regular files, only directory/regular-file
types, deterministic time/owner metadata, byte-correct LICENSE/README/INSTALL/
RC.3 notes/operator docs, a static Linux amd64 binary with the exact embedded
version/commit/build identity, and passed the completed exclusion audit.

Required exported-source validation then stopped at its first `make build`:
Go automatic VCS stamping returned `error obtaining VCS status` because the
independent `git archive` correctly contains no `.git` metadata. The release
builder had succeeded because it explicitly uses `-buildvcs=false`; the
ordinary Makefile build does not. This is recorded as `QWSG-053-F001`, an
OPEN/BLOCKING RELEASE BLOCKER. Per Owner authority, no GOFLAGS workaround,
repair, rebuild, transfer or remaining validation was attempted.

## Verification

Pre-change build, full tests, repository race tests, vet, formatting, focused
userruntime/userservice/runner/assessment/setup tests, release plumbing,
Framework 21, Builder 38, lifecycle 28, diversion 36, job/test-task,
whitespace and shell syntax passed. The sandboxed static systemd probe could
not bind its private socket or connect to the system bus, returned success, and
is retained only as a host-local limitation rather than acceptance evidence.

## Rollback

Restore only literal Task 053-owned targets from the verified snapshot after
Owner authorization; the RC.3 protocol/ledger were recorded absent before the
task. Never reset/clean broadly or touch Owner/historical evidence.

## Completion state

`complete with disclosed limitations — NOT READY FOR QWSG 1.1.0 RELEASE`; `QWSG-053-F001` remains `OPEN, BLOCKING`. Transfer and external acceptance never began.
