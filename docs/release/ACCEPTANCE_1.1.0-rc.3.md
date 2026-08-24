# QWSG 1.1.0-rc.3 Acceptance and Release-Readiness Record

## Current state

- Candidate: `PRIVATE BYTES CONSTRUCTED — VALIDATION STOPPED`
- External acceptance: `NOT STARTED`
- Final verdict: `NOT READY FOR QWSG 1.1.0 RELEASE`
- Historical `QWSG-051-F001`: `OPEN, BLOCKING` until the exact RC.3 external
  proof in Checkpoints 14–16 passes.
- Task 049 `QWSG-049-F002` and `QWSG-049-F003`: immutable history; fresh RC.3
  regression results are pending.

No entry in this ledger authorizes construction, transfer, VPS access,
credentials, external execution, tag, release, upload or publication.

## Owner gate ledger

| Gate | Scope | State | Evidence |
| --- | --- | --- | --- |
| A1 | repository/source readiness audit | COMPLETE | Task 053 Gate A1 report |
| A2 | exact acceptance-source integration | PENDING OWNER AUTHORIZATION | — |
| B/C/D | private twin construction and package proof | STOPPED | Twin build/package proof passed; required exported-source `make build` failed because the Git archive intentionally has no `.git` metadata and the Makefile build does not disable automatic Go VCS stamping. |
| E | private transfer and destination integrity | NOT AUTHORIZED | — |
| F | Checkpoints 01–25 | NOT AUTHORIZED | — |
| G | evidence integration and verdict | NOT AUTHORIZED | — |

## Source and candidate provenance

| Field | Value |
| --- | --- |
| Pre-A2 baseline | `3bf5a8e26e32ac1489dcdfecad5b086e4141cd91` |
| Task 052 product commit | `6bb5b62957e54e0ac3377ce1b85593408c341873` |
| Candidate-source commit | `3e7d2f9d543078b49d1afa522dc6bf3baba1c949` |
| Source version | `1.1.0-rc.3` |
| SOURCE_DATE_EPOCH / UTC build time | `1787314469` / `2026-08-21T12:14:29Z` |
| Archive | `qwsg-1.1.0-rc.3-linux-amd64.tar.gz` |
| Size / SHA-256 | `2949417` bytes / `c3ba763701b7ee0340d4928b21c23276dfdc083536b08814157366310629a0cc` |
| Binary/manifest/sidecar byte identity | PASS across both independent builds |
| Sidecar SHA-256 | `41be7497c427bae4dbccf240e846a69b7cad7c3fe5d9f8c2d8071b088f08e343`; verification PASS |
| Manifest SHA-256 | `33751c0f050b0304c9a0840a50c6f32287d2ada4793d12706f49971b93fa7dc3`; 18 entries verify PASS |
| Binary SHA-256 | `eb9926c99e90e2146ee8657797c71d1382c316e372af62186fc67d1bc6c3b044` |
| Static Linux amd64 / embedded source | PASS; static x86-64 ELF; `1.1.0-rc.3`, exact commit and UTC build identity |
| Package/layout/docs/LICENSE/exclusions | PASS: one root, 25 members, 19 regular files, only directories/regular files, byte-correct source documentation and exclusions |

## Transfer provenance

Transfer is not authorized and has not occurred. Later evidence records only
privacy-safe transport class, source/destination integrity, file count/types,
size, digest and sidecar result; never credentials or private host/account data.

## External checkpoint ledger

| CP | Evidence requirement | State |
| --- | --- | --- |
| 01 | private receipt and exact two-file provenance | NOT EXECUTED |
| 02 | archive checksum/sidecar | NOT EXECUTED |
| 03 | canonical safe archive layout | NOT EXECUTED |
| 04 | internal manifest | NOT EXECUTED |
| 05 | LICENSE, root docs, RC.3/binary/source identity | NOT EXECUTED |
| 06 | README/INSTALL operator journey | NOT EXECUTED |
| 07 | clean-host `install --check` | NOT EXECUTED |
| 08 | F002/F003 product-guidance regression | NOT EXECUTED |
| 09 | immutable installation | NOT EXECUTED |
| 10 | setup interruption/resume and one-recipient config | NOT EXECUTED |
| 11 | protected mode-0600 credential-file workflow | NOT EXECUTED |
| 12 | notification preflight | NOT EXECUTED |
| 13 | real external SMTP test and Owner receipt | NOT EXECUTED |
| 14 | explicit guided Guardian activation | NOT EXECUTED |
| 15 | F001 external correction proof | NOT EXECUTED |
| 16 | independent enabled/active/fresh readiness | NOT EXECUTED |
| 17 | process/invocation/cadence/resources/restart evidence | NOT EXECUTED |
| 18 | lingering detection/guidance | NOT EXECUTED |
| 19 | logout/session behavior | NOT EXECUTED |
| 20 | physical VPS reboot | NOT EXECUTED |
| 21 | automatic return/new identity/fresh evidence | NOT EXECUTED |
| 22 | post-reboot notification and second receipt | NOT EXECUTED |
| 23 | explicit Guardian restart/fresh recovery | NOT EXECUTED |
| 24 | safe uninstall and preserved user state | NOT EXECUTED |
| 25 | same-candidate reinstall/resume/reactivation/final readiness | NOT EXECUTED |

## F001/F002/F003 ledger

- `QWSG-051-F001` remains historical `OPEN, BLOCKING`. It can be recorded here
  as `EXTERNALLY VERIFIED CORRECTED` only after product-guided setup activation
  succeeds on RC.3 and independent readiness proves enabled, active and fresh
  canonical evidence. Manual `systemctl`, local tests or process state alone do
  not qualify.
- Task 049 F002/F003 history is never edited. RC.3 requires a fresh uncoached
  Smart Install run proving validated user-manager context and actionable safe
  guidance, plus bounded read-only filesystem evidence or precise nonmutating
  manual verification.

## SMTP and privacy ledger

| Requirement | State |
| --- | --- |
| Protected current-user mode-0600 credential boundary | NOT EXECUTED |
| Preflight | NOT EXECUTED |
| Pre-reboot controlled test | NOT EXECUTED |
| Independent pre-reboot Owner receipt | NOT CONFIRMED |
| Post-reboot controlled test | NOT EXECUTED |
| Independent post-reboot Owner receipt | NOT CONFIRMED |

Never record recipient, SMTP account, provider/host where private, credential
or reference, headers, tokens, private paths or host/account identity.

## Guardian, reboot and lifecycle-operation evidence

| Requirement | State |
| --- | --- |
| Guided activation stage/result | NOT EXECUTED |
| Enabled/active/fresh canonical evidence | NOT EXECUTED |
| Process/invocation/cadence/resources/restarts | NOT EXECUTED |
| Lingering/logout behavior | NOT EXECUTED |
| Physical reboot/new post-reboot identity | NOT EXECUTED |
| Fresh post-reboot evidence | NOT EXECUTED |
| Explicit restart and recovery | NOT EXECUTED |
| Safe uninstall with preserved config/credential/state | NOT EXECUTED |
| Reinstall/setup resume/reactivation/final readiness | NOT EXECUTED |

## Finding register

| ID | Checkpoint | Severity | Description | State |
| --- | --- | --- | --- | --- |
| QWSG-051-F001 | historical RC.2 guided activation | RELEASE BLOCKER | RC.2 lost validated user-runtime context; RC.3 external retest required. | OPEN, BLOCKING |
| QWSG-053-F001 | Gate B/C/D exported-source validation | RELEASE BLOCKER | Required `make build` in the independent Git-exported source failed before tests because Go automatic VCS stamping could not obtain repository status from an export without `.git`; the release builder uses `-buildvcs=false`, but the ordinary Makefile build does not. Per Gate B/C/D authority, validation stopped without a workaround or rebuild. | OPEN, BLOCKING |

New findings must use the established severity model and stop acceptance where
required. Do not repair product behavior inside this acceptance task.

## Final decision gate

`READY FOR QWSG 1.1.0 RELEASE` is permitted only when every mandatory ledger
entry passes, F001 is externally verified corrected, F002/F003 regressions
pass, both real receipts are independently confirmed, reboot and lifecycle
operations pass, no blocker/security defect remains, and all local/governance
gates pass.

The result is `NOT READY FOR QWSG 1.1.0 RELEASE`: required exported-source
validation failed at Gate B/C/D on open blocking `QWSG-053-F001`. Transfer and
all external checkpoints remained unexecuted. This is a completed acceptance
task with disclosed limitations, not successful release acceptance. No
release/publication action is authorized.
