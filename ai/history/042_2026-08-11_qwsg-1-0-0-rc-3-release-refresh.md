# Task History 042: Approved Task

## Task metadata

- Task ID: `042`
- Task slug: `qwsg-1-0-0-rc-3-release-refresh`
- Status: `complete with disclosed external clean-host gate`
- Date generated: `2026-08-11` UTC
- Human authority: Project Owner
- Preferred owner communication language: Hungarian
- Related prompt: `ai/archive_prompts/042_2026-08-11_qwsg-1-0-0-rc-3-release-refresh.md`

## Lifecycle state

The Engineering Task Builder generated and transactionally installed this matching prompt/history pair from validated structured owner input. The Project Owner started Task 042 through the canonical `job` workflow on 2026-08-11 UTC.

## Starting state

- Canonical Task 041 idle baseline and deterministic Task 042 Builder installation were verified before execution. Task 042 is the sole active approved task.
- Repository root, `main`, HEAD `0a8a5c7e722495b8c5eb425bca5b2d2413aaa175`, canonical HTTPS origin, `0/0` upstream relationship and empty index were verified.
- Ubuntu 24.04, linux-amd64, Go 1.26.5, Make 4.3, GNU tar 1.35, gzip 1.12, systemd 255 and sufficient workspace capacity were verified. No QWSG or release process was active.
- Source identity was exactly `1.0.0-rc.2`; RC.3 destinations were absent. The RC.2 archive SHA-256 matched `0694fb7f382ea1b373aaf0b6f0171a3fc580c491c1af3b8657a3bb8697ed897b`.
- Release Policy permits `1.0.0-rc.N`. The generalized builder already selects version-derived notes and emits location-independent sidecars, so no script change is required.
- Focused/full/race Go tests, vet, formatting, Framework, Builder, lifecycle, diversion, test-task audit, `bin/job` and Git whitespace gates passed before metadata changes.

## Snapshot

Verified implementation snapshot: `/tmp/qwsg-task042-implementation-20260811T.Tjyarw`. It contains exact target payloads and absence records, RC.1/RC.2 evidence, identities, ACL/mode records, guarded restore instructions and deterministic SHA-256 evidence. Payload-manifest file SHA-256 is `f667d8250da96bc9ea93e04cab4215c782dba2018008069dc3a90c759483061b`; snapshot archive SHA-256 is `279549e92df8bedc0fc0283a37e979a177f472d9a4f0d32c5b45c8647818acf8`.

## Work performed

Advanced the coherent repository, command and release-document identity to
`1.0.0-rc.3`. Added exact RC.3 release notes and engineering acceptance,
updated Quick Start, changelog, README, roadmap and milestone index, and kept
the already generalized release builder unchanged. No accepted Task 039/041
product semantics, installer, unit, license or historical RC evidence changed.

Produced `dist/qwsg-1.0.0-rc.3-linux-amd64.tar.gz` and its matching sidecar
from fixed commit/epoch inputs. Two independent clean build roots were
byte-identical for binary, manifest, archive and sidecar. The final archive
SHA-256 is
`bc4eca323cb07d23f0d6c884886655eb5549c03c136d8c773782d7813551585c`;
the internal manifest SHA-256 is
`1b764edc8686f12618c73a8870ce714ce02378152784a94817d9ff2aabc333f4`.

Archive-only staging covered clean installation, collision refusal, explicit
RC.2 replacement, backup integrity, exact rollback, modified-artifact refusal,
archive restoration and clean uninstall. The extracted static binary completed
empty-HOME bootstrap, first partial baseline, second full pipeline, Current
State Console load, partial `check`, strict modes and unsafe-path rejection.
Real host artifact evidence reached exactly 366 Policy/Report sources. An
isolated foreground Guardian recorded 51 successful cycles; separate Console,
direct PTY `r`, manual-observe lock, graceful stop and SIGKILL freshness
demotion passed.

The first post-build checksum command was an `attempt-failed` procedural invocation: the location-independent sidecar was checked from the repository root instead of its containing `dist/` directory. Both independent archives and sidecars had already compared byte-identical. No artifact changed; validation resumed from the sidecar directory without repeating either build.

The first staged rollback sequence was also an `attempt-failed` test ordering: it uninstalled RC.3 before trying a clean RC.2 reinstall, while the intentionally preserved historical RC.2 release-note file still occupied its owned path. The package and host remained isolated. The accepted rollback method instead restores the installer-created exact backup while RC.3 is installed and removes only the newly introduced RC.3 release-note path.

The modified-artifact refusal test correctly stopped at the changed systemd unit after removing earlier hash-identical payload entries. Restoring only that unit was insufficient for a following clean uninstall. As in the Task 040 acceptance method, the complete staged payload is restored from the verified archive before the owned uninstall is repeated; this is isolated test recovery, not a product change.

A proposed faster endurance continuation changed the Guardian interval from 5s to 2s against an existing checkpoint and failed closed with `guardian_checkpoint_invalid`. This is correct definition-compatibility behavior. The method was rejected and endurance completed with the original 5s/4s definition.

## Verification

Builder installation, starting environment, historical RC identity, pre-change
gates and implementation snapshot validated successfully. `make build`, focused
Task 039/041 tests, complete Go tests, repository-wide race tests, vet, format,
Framework (21), Builder (38), lifecycle (28), diversion (36), test-task audit,
`bin/job`, Git whitespace and empty-index gates passed.

Two release assemblies matched byte-for-byte. Archive safety, deterministic
metadata, internal manifest, external sidecar, static version identity,
staged install/upgrade/rollback/uninstall, systemd unit and snapshot checks
passed. RC.1/RC.2 artifacts, sidecars and release documents remained
byte-identical. All non-snapshot Task 042 temporary roots and processes were
removed after validation; the implementation and Builder snapshots remain.

The development host physically contains Go. Static/archive-only execution and
a minimal PATH were verified, but physical absence of the Go tree is not
claimed. The Owner's clean Ubuntu 24.04 no-Go installation, Guardian/reboot and
uninstall acceptance remains the explicitly deferred release gate, not a
product blocker found by Task 042. No real installation, licensing change,
stage, commit, tag, push or publication occurred.

## Rollback

Use only the exact verified snapshot after confirming no Task 042 process is active, current targets are Task 042-owned and no later Owner edit exists. Restore only recorded paths and remove newly created RC.3 paths only when their hashes match Task 042 evidence. Never use broad Git recovery or alter RC.1/RC.2, private state or unrelated Owner content.

## Completion state

`complete with disclosed external clean-host gate`

## Release decision

`READY FOR CLEAN-HOST ACCEPTANCE`

Transfer exactly:

- `dist/qwsg-1.0.0-rc.3-linux-amd64.tar.gz`
- `dist/qwsg-1.0.0-rc.3-linux-amd64.tar.gz.sha256`
