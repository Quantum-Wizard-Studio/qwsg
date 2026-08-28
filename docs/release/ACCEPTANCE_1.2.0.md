# QWSG 1.2.0 Final Release Decision

## Decision

`BLOCKED — QWSG 1.2.0 was not tagged, published or distributed.`

Task 067 corrected QWSG-066-F001 and produced a deterministic private RC.3
candidate, but Task 068 supersedes it because the Project Owner required
QWSG-managed administrator change notifications before final acceptance. RC.3
must not be promoted. Private RC.4 is the next acceptance subject. The
final `1.2.0` decision remains blocked until that matrix passes; no tag,
publication, Forgejo Release, or VPS mutation was performed by Task 067.

## Task 067 remediation and private RC.3

- Canonical modes: directories `0755`; `bin/qwsg`, `install.sh`, and
  `uninstall.sh` `0755`; all other regular files `0644`.
- Source commit: `b6eb357ad03a02b41ac93536fc3be91ecf929803`.
- Controlled build epoch: `1787905053` (`2026-08-28T08:17:33Z`).
- Private artifact: `qwsg-1.2.0-rc.3-linux-amd64.tar.gz`.
- SHA-256: `8543c3e09b48085b01c037d7db5106ea793374dc099b0b9be5f6cacb55af13ee`.
- Rebuild from a normalized Git export under `umask 0077` matched the
  group-writable worktree build byte-for-byte with the same SHA-256.
- Extraction under `umask 0002` and `0077` produced identical canonical modes;
  archive ownership was numeric `0/0`, no regular file was group/world
  writable, and exactly the three intended files were executable.
- Automated `release-check` now includes the cross-umask/source-mode SHA-256
  regression and exact mode allowlist.

Task 066 audited the exact private `1.2.0-rc.2` candidate before promotion. A
release-builder permission-normalization defect makes the archive depend on
ambient worktree modes and leaves the frozen candidate's extracted
documentation, unit and configuration example group-writable. This violates
the deterministic-release and least-privilege gates. Source correction would
create new candidate bytes, so RC.2 cannot be promoted truthfully.

## Audited candidate

- Version: `1.2.0-rc.2` (private and unpublished).
- Source commit: `c260dc18c2004473ec55496d16e66718fd128865`.
- Controlled build epoch: `1787839370` (`2026-08-27T14:02:50Z`).
- Frozen archive: `qwsg-1.2.0-rc.2-linux-amd64.tar.gz`.
- Frozen archive SHA-256:
  `a34be8b18f80d877c0ccfd69dc9d9e9f197fc35fa765cdf1d5c0d72e2cb0a554`.
- Frozen sidecar SHA-256:
  `23218e15a85ab5ee644031bf5ad6469d33a247fc1e271f550de1d52d83e88e11`.
- Starting/final publication state: no `v1.2.0*` tag and no Forgejo `v1.2.0`
  Release or assets.

Two Task 066 builds in the repository environment reproduced the frozen
archive byte-for-byte and matched the Task 065 candidate. The outer checksum,
internal manifest, safe member paths/types, embedded version/source/build
identity and packaged required documents passed.

## Release blocker QWSG-066-F001

**Classification:** `PRODUCT/FRAMEWORK DEFECT — RELEASE BLOCKER`.

`scripts/build-release.sh` explicitly sets executable modes but does not
normalize directory modes or non-executable regular-file modes before creating
`MANIFEST.sha256` and the tar stream. GNU tar therefore preserves ambient modes
from the build worktree and its mode-`0002` environment.

The frozen RC.2 archive contains:

- package directories with mode `0775`;
- packaged documentation, `LICENSE`, `CHANGELOG.md`, the systemd user unit and
  configuration example with mode `0660`; and
- generated `MANIFEST.sha256` and `RELEASE.json` with mode `0664`.

An independent export of the exact source commit was given ordinary
non-executable file mode `0644` and built with the identical commit, epoch,
toolchain and logical file content. That archive had SHA-256
`5b32df3b090658cfb9a08a7d670848c65af4d5d048dc053e3ad0973d11f0082a`,
not the frozen candidate hash. A recursive content comparison of both extracted
payloads passed with no byte-content difference; a representative `README.md`
mode was `0660` in the frozen candidate and `0644` in the clean export.

The packaged installer later writes installed targets as `0755` or `0644`, but
that does not repair the public archive contract: extracted release files
remain unnecessarily group-writable, and equivalent source checkouts do not
produce one canonical byte identity.

## Verification completed before STOP

- Canonical Task 066 lifecycle, Framework and remote baseline: PASS.
- Required protected snapshot, checksums, archive readability and bounded
  rollback instructions: PASS.
- Fresh `origin/main` comparison: PASS at
  `f726d84632ebce3f2be72101b583af5beadc857e`, ahead/behind `0/0`.
- Full Framework, diversion, lifecycle and Builder suites: PASS (`25`, `36`,
  `29` and `49` assertions).
- Framework v2 and bounded diagnostic suites: PASS (`15` and `11` assertions).
- Full Go tests with writable isolated caches: PASS.
- Repository-wide Go race tests: PASS.
- `go vet ./...`, formatting, shell syntax, release plumbing and Git whitespace:
  PASS.
- RC.2 two-build reproducibility inside the same ambient mode environment and
  exact match to the Task 065 frozen candidate: PASS.
- Cross-mode equivalent-source archive reproducibility: FAIL, blocking.

The first aggregate validation attempt also exposed a stale
`scripts/test-build-contract.sh` RC identity (`1.2.0-rc.1`). It was classified
as a `TEST OR ACCEPTANCE DEFECT`, narrowly corrected to RC.2 and passed. A
literal Go-test retry initially selected the sandbox-read-only default Go cache;
that was classified `ENVIRONMENTAL ISSUE` and passed with the canonical
writable caches. Neither observation changes RC.2 product bytes.

## External acceptance disposition

Task 066 stopped before release-affecting host mutations after proving
QWSG-066-F001. The prior Task 065 clean-host guided-install evidence remains
valid for its exact claims, but it does not satisfy Task 066's complete final
release acceptance, real published Forgejo retrieval, twelve-step physical
acceptance, full Contabo coexistence, or final update/restorative-rollback
gates. No OVH or Contabo system was modified during Task 066. Missing gates are
not recorded as PASS.

## Smallest correct remediation

Prepare a new private candidate identity, not a rewritten RC.2:

1. normalize every package directory to `0755`, every non-executable regular
   payload file to `0644`, and only the binary/install/uninstall entry points to
   `0755` before manifest generation and archive assembly;
2. add regression coverage that builds identical exported source under
   deliberately different ambient umasks/modes and requires byte-identical
   archives plus an exact tar-mode allowlist;
3. construct a new deterministic candidate (normally `1.2.0-rc.3`) with new
   source provenance and checksums; and
4. repeat all local, clean OVH, full Contabo, update, restorative rollback,
   physical reboot, protected-notification and real Forgejo distribution gates
   before any final `1.2.0` publication.

RC.2 and its recorded hashes remain immutable historical evidence. They must
not be republished under another identity or silently replaced.
