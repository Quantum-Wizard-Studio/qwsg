# Task History 038: Approved Task

## Task metadata

- Task ID: `038`
- Task slug: `qwsg-version-1-0-release-hardening-and-release-candidate`
- Status: `complete`
- Date generated: `2026-08-10` UTC
- Human authority: Project Owner
- Preferred owner communication language: Hungarian
- Related prompt: `ai/archive_prompts/038_2026-08-10_qwsg-version-1-0-release-hardening-and-release-candidate.md`

## Lifecycle state

The Builder-installed task was executed through the canonical `job` workflow,
completed as the Version 1.0 technical release gate, and returned to canonical
idle without creating Task 039.

## Starting state

- Task 038 started through the canonical `job` workflow on 2026-08-10 UTC.
- Repository root, `main`, HEAD `0a8a5c7e722495b8c5eb425bca5b2d2413aaa175`, canonical HTTPS origin, `0/0` relationship, empty index, sole active Task 038 prompt/history pair, and completed Task 037 baseline were verified.
- Ubuntu 24.04, linux-amd64, Go 1.26.5, systemd 255, Framework 1.0.0, ownership/modes/ACLs, and the pre-existing Owner-owned modified/untracked worktree were inspected.
- The pre-change build, full tests, race tests, vet, format, Framework, lifecycle, diversion and Builder suites passed. The sandbox denied the user/system bus; real systemd acceptance therefore remains isolated and requires the bounded host-service authority already granted by the task.
- The release audit confirmed the approved blockers: pre-alpha version/policy, no self-contained release archive or checksums, repository/Go-dependent installation, install-prefix/unit mismatch, incomplete uninstall/upgrade/support documentation, and unexecuted Task 038 acceptance. No missing canonical business engine was found.

## Snapshot

Verified implementation snapshot: `/tmp/qwsg-task038-implementation-20260810TcNbal7`. It contains exact payloads or absence records for every planned release, documentation, Console and lifecycle target, plus repository/tool identity and SHA-256 verification. Restore is bounded to recorded targets after drift and active-service checks.

## Work performed

Version was advanced coherently to `1.0.0-rc.1`. Task 038 added deterministic
linux-amd64 archive assembly, internal/external SHA-256 evidence, fixed-prefix
safe install/uninstall tooling, a validated Source Record example, and complete
Quick Start, operations, troubleshooting, upgrade/rollback/uninstall, support,
security, limitation, release-note and acceptance documentation.

The `/usr/local` release boundary now refuses alternate Make prefixes rather
than installing a unit that points to another executable. Installation verifies
the archive, refuses collisions by default, requires an explicit new backup
directory for replacement, and never enables/starts service or changes linger.
Uninstall removes only hash-identical release artifacts and preserves all user
configuration/state.

The interactive Console now emits one initial Overview and a terminal clear
only before subsequent interactive redraws. Canonical engines, Runtime Service,
Guardian, state and presentation ownership were unchanged. Historical Alpha
gates were reconciled with the frozen local 1.0 product; optional providers,
Dashboard/API, fleet, remediation and AI are post-1.0.

## Verification

- `make build`, full tests, race tests, vet, format, Framework 1.0.0, all
  Builder/lifecycle/diversion suites, Git whitespace and systemd unit checks
  passed.
- Two final controlled builds were byte-identical. Final archive SHA-256 is
  `5b9751048e54eae584b89a4339877813ab0604297a1338ae4516b5657a8716c8`.
  Manifest, traversal-free archive, embedded version, fixed prefix, staged
  install/reinstall/rollback/uninstall and modified-artifact safety passed.
- A unique isolated systemd user unit completed 78 recurrence results. Peak
  memory was 8.5 MiB, tasks/threads 8, FDs 7, endurance restart count 0,
  snapshots stayed at retention 10, and Scheduler state remained bounded.
- Separate Console, graceful stop/stopped evidence, restart, SIGKILL plus one
  bounded restart, invalid configuration, explicit-observe/single-instance
  refusal, foreground SIGINT, legacy/current state compatibility and truthful
  lifecycle output passed.
- Real PTY evidence showed one initial Overview and one clear/redraw for the
  accepted quit action. The acceptance unit/link and all non-snapshot temporary
  roots were removed; real QWSG service/state and `Linger=no` were unchanged.
- A broader capability-dropping systemd hardening experiment failed safely in
  the supported user-manager environment and was rejected; the already proven
  Task 037 least-privilege sandbox remains shipped. Detailed sanitized evidence
  is in `docs/release/ACCEPTANCE_1.0.0-rc.1.md`.

## Rollback

Use only the exact payload and absence records in `/tmp/qwsg-task038-implementation-20260810TcNbal7`; reject restoration if a target changed after Task 038 or any acceptance service is active. Do not use Git reset/clean/checkout/restore and never remove user state.

## Completion state

`complete`

## Release decision

`READY FOR QWSG 1.0 RELEASE`

No technical release blocker remains. A disposable-host physical reboot
procedure and public-license confirmation remain explicit Owner gates before
public claims/publication; neither authorizes Task 039, tagging or publication.
