# QWSG 1.1.0-rc.5 Acceptance and Release-Readiness Record

## Current state

- Candidate: `BUILT PRIVATELY; NOT TRANSFERRED`
- Transfer: `NOT STARTED`
- External acceptance: `NOT STARTED`
- Final verdict: `PENDING`
- Notification receipts: `NOT CONFIRMED`
- Candidate-source commit: `1025d36d05b2f6f919f0ea4ec4a7029f67536000`
- Historical QWSG-055-F001: `OPEN, BLOCKING`

This local Gate B evidence authorizes no rebuild, evidence integration,
transfer, external execution, credentials, tag, Forgejo Release, upload or
publication.

## Owner gate ledger

| Gate | Scope | State |
| --- | --- | --- |
| A | readiness audit and acceptance-scaffolding integration | PASS; INTEGRATED AS `1025d36d05b2f6f919f0ea4ec4a7029f67536000` |
| B | private exact-commit twin construction | PASS; PRIVACY-SAFE LOCAL EVIDENCE RECORDED |
| C | deterministic/package/provenance/security proof and evidence integration | NOT STARTED |
| D | private transfer and destination integrity | NOT STARTED |
| E | Checkpoints 01–26 fresh external clean-host acceptance | NOT STARTED |
| F | privacy-safe evidence integration and final verdict | NOT STARTED |

## Source and candidate provenance

| Field | State |
| --- | --- |
| Gate A baseline | `05a49a2e56d254ab3eb4646dc9df04fa4b63e335` |
| Task 056 integration | `fc52295f6cc5b61078b32535230bdc704168d13c` |
| Candidate-source commit | `1025d36d05b2f6f919f0ea4ec4a7029f67536000` |
| Source version | `1.1.0-rc.5` |
| SOURCE_DATE_EPOCH / UTC | `1787594463` / `2026-08-24T18:01:03Z` |
| Archive filename, size and SHA-256 | `qwsg-1.1.0-rc.5-linux-amd64.tar.gz`; `2951350` bytes; `cfe300c0f1f312d80120f74a9f24bed4a64387471bf2097ddc63d94f0fb2f7b0` |
| Sidecar SHA-256 / verification | `69f3eb4bf89dc126a7eafd08354eec37a941014171b3d1d70c6e6a4cf52e5eb0`; PASS in both private roots |
| Manifest SHA-256 / verification | `ae51aca0bc4ddc61b0daea3a87f0acabcde5ec9fd8fadddc050f0786d6915e9e`; every entry PASS in both twins; no unmanifested payload |
| Binary SHA-256 / embedded identity | `5484aab96d5c3748e81b065fdb11ec8c34385589bb07ee7ea1b2b35fdffa6b93`; QWSG `1.1.0-rc.5`, exact full commit, controlled UTC time |
| Twin byte identity | PASS for ordinary binaries and release binary, manifest, archive and sidecar across two independent no-`.git` exports |
| Static platform/package/docs/LICENSE/exclusions | PASS; stripped static ELF64 x86-64, canonical safe root/content/types/modes/timestamps, source-identical docs/LICENSE, no excluded content |

## External checkpoint ledger

| CP | Evidence requirement | State |
| --- | --- | --- |
| 01 | private receipt and exact two-file provenance | NOT EXECUTED |
| 02 | archive checksum/sidecar | NOT EXECUTED |
| 03 | canonical safe archive layout | NOT EXECUTED |
| 04 | internal manifest | NOT EXECUTED |
| 05 | LICENSE, docs, RC.5/binary/source identity | NOT EXECUTED |
| 06 | README/INSTALL operator journey | NOT EXECUTED |
| 07 | clean-host Smart Install | NOT EXECUTED |
| 08 | Task 049 F002/F003 regression | NOT EXECUTED |
| 09 | immutable installation | NOT EXECUTED |
| 10 | setup interruption/resume/configuration | NOT EXECUTED |
| 11 | protected credential workflow | NOT EXECUTED |
| 12 | notification preflight | NOT EXECUTED |
| 13 | first actual notification and Owner receipt | NOT EXECUTED |
| 14 | guided Guardian activation without manual workaround | NOT EXECUTED |
| 15 | Task 056 real state-directory/systemd compatibility proof | NOT EXECUTED |
| 16 | QWSG-055-F001 external correction proof | NOT EXECUTED |
| 17 | QWSG-051-F001 and complete readiness proof | NOT EXECUTED |
| 18 | process/invocation/cadence/resources | NOT EXECUTED |
| 19 | lingering detection/guidance | NOT EXECUTED |
| 20 | logout/session behavior | NOT EXECUTED |
| 21 | physical VPS reboot | NOT EXECUTED |
| 22 | automatic post-reboot return/new identity/fresh evidence | NOT EXECUTED |
| 23 | post-reboot notification and second receipt | NOT EXECUTED |
| 24 | explicit Guardian restart/recovery | NOT EXECUTED |
| 25 | uninstall with preserved user state | NOT EXECUTED |
| 26 | same-candidate reinstall/resume/reactivation/final readiness | NOT EXECUTED |

## Historical finding and retest ledger

| Finding | Immutable historical state | Required fresh RC.5 evidence | Current RC.5 state |
| --- | --- | --- | --- |
| QWSG-055-F001 | OPEN, BLOCKING | Checkpoints 14–16: real safe state directory, no compatibility symlink, guided activation, enabled/active/fresh evidence and satisfied filesystem assessment | NOT EXECUTED |
| QWSG-051-F001 | OPEN, BLOCKING | Checkpoint 17: documented guided workflow and independent enabled/active/fresh readiness | NOT EXECUTED |
| QWSG-053-F001 | OPEN, BLOCKING | Gates B/C: ordinary no-`.git` builds without GOFLAGS and exact deterministic provenance | LOCAL GATE B RETEST PASS; historical state remains OPEN/BLOCKING pending later acceptance boundaries |
| QWSG-049-F002 | OPEN, BLOCKING | Checkpoints 07–08: fresh actionable user-manager guidance | NOT EXECUTED |
| QWSG-049-F003 | OPEN, RELEASE-GATE BLOCKING | Checkpoints 07–08 and 15–17: fresh local-semantics assessment before and after activation | NOT EXECUTED |

Historical evidence remains unchanged. An external RC.5 PASS is additive and
does not rewrite the original finding chronology.

## Task 056 / QWSG-055-F001 proof ledger

| Requirement | State |
| --- | --- |
| Canonical state root is a real directory | NOT EXECUTED |
| No state-root or unsafe path-component symlink | NOT EXECUTED |
| Intended ordinary-user ownership | NOT EXECUTED |
| Mode 0700 | NOT EXECUTED |
| Configuration and state roots distinct | NOT EXECUTED |
| No systemd compatibility symlink/migration | NOT EXECUTED |
| Guided activation without manual systemctl repair | NOT EXECUTED |
| Service enabled and active | NOT EXECUTED |
| Fresh integrity-checked canonical Guardian evidence | NOT EXECUTED |
| Post-activation filesystem.local_semantics satisfied | NOT EXECUTED |

QWSG-055-F001 remains `OPEN, BLOCKING` until every row passes with fresh
independent evidence on the genuinely clean RC.5 host.

## SMTP and privacy ledger

| Requirement | State |
| --- | --- |
| Protected regular current-user mode-0600 credential boundary | NOT EXECUTED |
| Preflight | NOT EXECUTED |
| Pre-reboot controlled test | NOT EXECUTED |
| Independent pre-reboot Owner receipt | NOT CONFIRMED |
| Post-reboot controlled test | NOT EXECUTED |
| Independent post-reboot Owner receipt | NOT CONFIRMED |

Never record credentials, recipient/provider identity, headers, tokens,
private host/account identity or raw private paths.

## Clean-host declaration

Gate E is blocked until the Owner confirms the disposable Ubuntu 24.04 amd64
VPS was freshly installed/reset before Checkpoint 01. The mutated RC.4
environment is not clean evidence by assertion.

## Mandatory post-acceptance distribution follow-up

A later separately authorized release/distribution phase must provide stable
official archive and SHA-256 sidecar URLs from `git.quantumwizard.hu`/Forgejo,
`wget` and `curl` examples, command-line checksum verification, future Smart
Installer compatibility, and ordinary installation without workstation-mediated
artifact copying. Gate A does not implement or publish this mechanism.

## Final decision gate

`READY FOR QWSG 1.1.0 RELEASE` requires Gates A–F and all 26 checkpoints PASS,
every required historical retest, both actual notification receipts,
reboot/restart/uninstall/reinstall success, and no open blocker or security
defect. Otherwise the verdict is `NOT READY FOR QWSG 1.1.0 RELEASE`. The current
verdict remains `PENDING`. Neither outcome authorizes a tag, Forgejo Release,
upload, publication or announcement.
