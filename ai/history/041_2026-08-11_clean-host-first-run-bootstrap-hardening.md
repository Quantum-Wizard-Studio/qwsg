# Task History 041: Approved Task

## Task metadata

- Task ID: `041`
- Task slug: `clean-host-first-run-bootstrap-hardening`
- Status: `complete with disclosed validation-environment limitation`
- Date generated: `2026-08-11` UTC
- Human authority: Project Owner
- Preferred owner communication language: Hungarian
- Related prompt: `ai/archive_prompts/041_2026-08-11_clean-host-first-run-bootstrap-hardening.md`

## Lifecycle state

The Engineering Task Builder generated and transactionally installed this matching prompt/history pair from validated structured owner input. The Project Owner started Task 041 through the canonical `job` workflow on 2026-08-11 UTC.

## Starting state

- Canonical Task 040 idle baseline, Task 041 prompt/history identity, repository root, `main`, HEAD `0a8a5c7e722495b8c5eb425bca5b2d2413aaa175`, canonical HTTPS origin, `0/0` upstream relationship and empty index were verified.
- Ubuntu 24.04, linux-amd64, Go 1.26.5, Make 4.3, systemd 255 and ext4 capacity were verified. No QWSG process was active.
- The clean-host failure path, single-level state-root creation, partial `components` evidence, partial-check rejection and bounded diagnostic mappings were reconfirmed from source.
- Pre-change focused/full/race tests, vet, format, Framework, Builder, lifecycle, diversion, `bin/job` and Git whitespace checks passed with bounded writable Go caches. The initial default-cache attempt failed only because the sandbox home cache is read-only and was not a product failure.

## Snapshot

Verified implementation snapshot: `/tmp/qwsg-task041-implementation-20260811T.vcXFCi`. It contains exact pre-change target payloads, identities, ACL evidence, repository state, guarded restore instructions and canonical RC.2 artifact bytes. Payload manifest SHA-256 is `6ec3b52c147ac203a7257541a5470d2bf8e17acaa72d1b1cfae8ecfc223a774f`; snapshot tar SHA-256 is `cde2c8735e6b6c3ee71aca5dc4d249489d353569b0f3399bc011ad7a610430ce`.

## Work performed

Added `operatorstate.EnsurePrivateRoot` as a narrow reuse of the existing clean-
absolute-path, symlink, ownership and exact-mode validation. Direct observe
bootstrap now uses it to recursively create missing parents while chmodding
only the final QWSG root. Existing ancestors remain unchanged.

Removed only the partial-stage completeness rejection from `publishCheck`;
execution, profile, schema, typed payload, snapshot identity, canonical
validation, equality and freshness gates remain. Valid partial evidence now
publishes a partial Current State. Added bounded `state_bootstrap_failed` and
`state_publication_failed` paths while preserving existing corrupt,
incompatible, permission and unsafe-path tokens.

Added empty-HOME, partial-check, repeat-bootstrap, ancestor-preservation,
symlink, permissive-root and privacy-safe diagnostic regressions. Updated only
the directly affected Current State architecture and troubleshooting text.
Installer, systemd unit, collectors, VERSION, licensing and release artifacts
were not changed.

## Verification

Builder installation, starting state, pre-change gates and implementation
snapshot validated successfully. Focused and complete Go tests, repository-wide
race tests, vet, formatting, Framework (21), Builder (38), lifecycle (28),
diversion (36), test-task audit, `bin/job`, Git whitespace and empty-index gates
passed.

A temporary reproducible package was built outside `dist`, verified through its
manifest/checksum and installed only under an isolated `--destdir`. Its archive
SHA-256 was `f85fc6df0ef7dfab59e86974ae27fc237d24c68ff32d6a440e0474e025b15dea`;
this is validation evidence, not a replacement RC. From a genuinely empty
synthetic HOME and repository-external working directory, the staged static
binary completed first observe with truthful partial Inventory, created one
baseline, then completed the normal eight-stage second observe and loaded the
published current/partial state in a separate Console invocation. A separate
partial `check` published successfully. QWSG roots/subdirectories were owned by
the ordinary user at `0700`; Current State, metadata, snapshots and lock files
were `0600`. Artifact-level symlink and permissive-root attempts were rejected
without creating protected descendants.

The validation host's Go installation was not altered. Attempts to hide it with
an unprivileged mount namespace and a minimal ordinary-UID chroot were denied by
host sandbox policy before QWSG ran; a ptrace fault-injection attempt was also
denied. These rejected isolation methods created no QWSG state and are not
counted as acceptance. No-Go behavior remains directly established by the
accepted clean-host evidence and unchanged collector contract; partial
bootstrap itself is covered with an unavailable-components fixture. The next
RC must repeat the final proof on the actual no-Go clean host.

Canonical RC.2 archive SHA-256 remained
`0694fb7f382ea1b373aaf0b6f0171a3fc580c491c1af3b8657a3bb8697ed897b`.
`packaging/release/install.sh`, the systemd unit, VERSION, canonical sidecar and
historical artifacts remained byte-identical.

## Rollback

Use only the exact verified snapshot after confirming no Task 041 process is active, current targets are Task 041-owned and no later Owner edit exists. Restore only snapshot-listed paths; never use broad Git recovery or alter unrelated Owner content, real QWSG state, installer/unit, VERSION or canonical RC artifacts.

## Completion state

`complete with disclosed validation-environment limitation`

The confirmed release-blocking first-run defect and coupled partial-check
diagnostic defect are corrected, source and staged-artifact empty-HOME behavior
passed, strict path security remained intact, and no additional product blocker
was found. The new Release Candidate must be produced separately and its final
no-Go proof repeated on the Owner clean host because this development sandbox
refused every safe method of hiding its installed Go tree. No successor task,
release refresh, real installation, clean-host operation, stage, commit, tag,
push or publication occurred.
