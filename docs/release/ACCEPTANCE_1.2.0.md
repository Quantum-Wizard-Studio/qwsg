# QWSG 1.2.0 Final Release Decision

## Task 071 guided-installer remediation and private RC.6

Task 069 remains `BLOCKED` by QWSG-069-F001; Task 070 remediated that blocker and produced private RC.5. During Owner-operated manual RC.5 pre-release acceptance on the clean OVH host, archive/manifest verification, first guided installation, administrator-enabled lingering, reboot-autostart, Guardian resource limits/readiness and supported uninstall preservation passed. The exact preserved-state reinstall installed the package and activated a healthy Guardian, but its immediate readiness assessment ran before the new Guardian's first canonical evidence cycle. The wizard returned exit 4 at 87%, omitted the completion phase/summary, and subsequent readiness immediately proved all mandatory domains ready with only optional notification partial. No package, configuration, state or Guardian corruption occurred. RC.5 was not promoted.

Task 071 remediates this synchronization defect with a shared bounded fresh-evidence waiter and explicit pre/post activation evidence identity. Deterministic regressions cover clean/delayed/preserved-state success, stale rejection, timeout/cancellation, optional Partial completion and localized diagnostics. RC.6 supersedes RC.5 as the next private acceptance candidate; final real-host acceptance remains pending.

Private artifact `qwsg-1.2.0-rc.6-linux-amd64.tar.gz` has SHA-256 `7f29f4cfe361412680c439eba7d8da5ae7d949b8702ce25c460baf991de99e33` and frozen source commit `6355022fd08ae4461d79edfc86943418fc1958d0` at controlled build time `2026-08-28T21:54:08Z`. An independent rebuild from a deliberately group-writable Git export under `umask 0077` matched the primary archive and sidecar byte-for-byte. Manifest, provenance, gzip/tar integrity, numeric `0/0` ownership, canonical internal modes, executable allowlist and secure/canonical extraction behavior passed. Task 071 classifies RC.6 `READY FOR FINAL ACCEPTANCE`; it is private, with no final tag or public release.

## Task 070 blocker remediation and private RC.5

Task 069 remains truthfully `BLOCKED` by `QWSG-069-F001`; RC.4 was not promoted. Task 070 implements an explicit deterministic `1.2.0-rc.2` -> `1.2.0-rc.5` compatibility path, installed-configuration preflight, independent privileged-boundary migration revalidation and a credential-free RC.2 installed-state regression fixture. The fixture proves RC.5 package replacement, byte-preservation of valid configuration/protected credential/relevant state, Guardian unit/resource/path semantics, protected rollback metadata and exact restoration of the RC.2 binary/unit state.

RC.5 supersedes RC.4 as the next private candidate. Private artifact `qwsg-1.2.0-rc.5-linux-amd64.tar.gz` has SHA-256 `ca3123f313c2fb95138fceb0fc2af17c5a8295b28e7b13e60f5d7fd57ed19faf` and frozen source commit `3f2f7822473d3d97c38c171f006e8b263250321c` at controlled build time `2026-08-28T19:01:18Z`. A rebuild from a deliberately group-writable normalized Git export under `umask 0077` matched the primary archive and sidecar byte-for-byte. Manifest, provenance, numeric `0/0` ownership, canonical modes, intended executable allowlist, archive integrity and extraction passed.

Task 070 classification: `READY FOR FINAL ACCEPTANCE`. Real OVH/Contabo update, restorative rollback, re-update, reboot, coexistence and mailbox acceptance remain pending for the subsequent final-acceptance task. Task 070 does not authorize a final tag or public release.

## Task 069 final acceptance result

`BLOCKED — QWSG 1.2.0-rc.4 is not promotable.`

Task 069 verified the immutable RC.4 candidate locally and completed real clean
OVH installation, restart, uninstall/reinstall, configuration/state
preservation and reboot-readiness acceptance. The designated Contabo host was
then tested from its actual installed QWSG `1.2.0-rc.2` state through the
candidate's canonical private-archive update command.

The mandatory update gate failed before package replacement. RC.4 correctly
classified RC.2 -> RC.4 as newer, verified the candidate, and then refused the
operation with `no deterministic compatible migration path`. Source inspection
confirmed that `internal/update.PlanMigration` contains compatibility paths
only for `1.1.0` -> RC.1/RC.2 and RC.1 -> RC.2. It contains no RC.2 -> RC.4
path and no regression test for that required transition. This is
`QWSG-069-F001`, a `PRODUCT/FRAMEWORK DEFECT — RELEASE BLOCKER`.

The failed operation remained visible independently from notification delivery:
the update returned failure while the configured credential-free Community SMTP
transport reported `Admin notification: ACCEPTED`. SMTP acceptance is not
mailbox-receipt proof; actual mailbox receipt was not claimed or requested after
the underlying mandatory update gate had already failed.

No Contabo package replacement or native rollback transaction occurred. The
pre-update RC.2 binary, unit and configuration were restored byte-identically;
configuration validation, readiness, Guardian stability, resource controls and
the named Hestia/web/database/mail/DNS/security services returned to their
recorded baseline. The Task 069 acceptance directory was removed. The OVH host
was restored to its sterile QWSG-free state after its successful acceptance
journey. No final `v1.2.0` tag, final artifact, Forgejo Release, asset upload or
external final-download claim was created.

Smallest remediation: create a new private candidate that explicitly supports
and tests the real prior-candidate transition (at minimum RC.2 -> the new
candidate) without schema mutation, then repeat notification-enabled update,
restorative rollback, re-update, reboot, full coexistence, final packaging,
Forgejo publication and external wget/curl acceptance. RC.4 remains immutable
historical evidence and must not be silently repaired or relabelled.

## Decision

`BLOCKED — QWSG 1.2.0 was not tagged, published or distributed.`

Task 067 corrected QWSG-066-F001 and produced a deterministic private RC.3
candidate, but Task 068 supersedes it because the Project Owner required
QWSG-managed administrator change notifications before final acceptance. RC.3
must not be promoted. Private RC.4 is the next acceptance subject. The
final `1.2.0` decision remains blocked until that matrix passes; no tag,
publication, Forgejo Release, or VPS mutation was performed by Task 067.

## Task 068 private RC.4 preparation

- Source commit: `4f7dcc11b5ccc9f078755946995baebd31ad6870`.
- Controlled epoch: `1787907901` (`2026-08-28T09:05:01Z`).
- Artifact: `qwsg-1.2.0-rc.4-linux-amd64.tar.gz`.
- SHA-256: `adeb591605c0d37a5fc98d541125ca388cd4561703d0f0823bba931bc7d08684`.
- Rebuild from a normalized Git export under `umask 0077` matched the primary
  build byte-for-byte. Manifest, embedded provenance, numeric `0/0` ownership,
  canonical `0755`/`0644` modes and extraction under two umasks passed.
- RC.4 is ready for the later full final-acceptance matrix; no real SMTP/VPS or
  publication action was performed by Task 068.

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
