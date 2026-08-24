# QWSG 1.1.0-rc.5 Acceptance and Release-Readiness Record

## Current state

- Candidate: `BUILT PRIVATELY; TRANSFERRED AND DESTINATION-VERIFIED`
- Transfer: `PASS — OWNER-WORKSTATION FALLBACK; READ-ONLY RECOVERY VERIFIED`
- Task classification: `COMPLETE WITH DISCLOSED ACCEPTANCE LIMITATIONS`
- Core result: `RC.5 CORE CLEAN-HOST FUNCTIONAL PROOF: ACHIEVED`
- Formal result: `FORMAL 26-CHECKPOINT RELEASE-READINESS ACCEPTANCE: INCOMPLETE`
- External acceptance: `TERMINATED INCOMPLETE BY OWNER PROCESS DECISION`
- Final verdict: `NOT READY FOR QWSG 1.1.0 RELEASE — FORMAL CERTIFICATION TERMINATED INCOMPLETE BY OWNER DECISION`
- Notification receipts: `NOT CONFIRMED`
- Candidate-source commit: `1025d36d05b2f6f919f0ea4ec4a7029f67536000`
- Historical QWSG-055-F001: `OPEN, BLOCKING`

The Owner confirmed the host was fully reinstalled/reset before acceptance.
Gate D destination verification passed before extraction but was reported only
after setup and activation had begun. The reporting-order deviation is
preserved; bounded recovery reverified the unchanged archive and sidecar.

## Owner gate ledger

| Gate | Scope | State |
| --- | --- | --- |
| A | readiness audit and acceptance-scaffolding integration | PASS; INTEGRATED AS `1025d36d05b2f6f919f0ea4ec4a7029f67536000` |
| B | private exact-commit twin construction | PASS; PRIVACY-SAFE LOCAL EVIDENCE RECORDED |
| C | deterministic/package/provenance/security proof and evidence integration | PASS; INTEGRATED AS `957d2ffd88139cffbb127bb336abeb2282c1e8db` |
| D | private transfer and destination integrity | PASS; OWNER-WORKSTATION FALLBACK; DESTINATION RECOVERY PASS |
| E | Checkpoints 01–26 fresh external clean-host acceptance | TERMINATED BY OWNER; FORMAL SEQUENCE INCOMPLETE |
| F | privacy-safe evidence integration and final verdict | COMPLETE WITH DISCLOSED ACCEPTANCE LIMITATIONS; NOT READY |

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

The Owner authorized a full clean-host restart from Checkpoint 01 on a newly
reinstalled Ubuntu 24.04 amd64 VPS. The earlier out-of-sequence RC.5 attempt
remains preserved as additive chronology, including its successful Gate D
recovery and Task 056/QWSG-055-F001 proof, but it does not substitute for any
checkpoint in this restarted run.

## Preserved RC.5 core functional proof

Before the formal restart, the exact verified RC.5 candidate produced genuine
external evidence on an Owner-confirmed freshly reinstalled Ubuntu 24.04 amd64
VPS. Transfer integrity, archive checksum and internal manifest verification
passed. Smart Install reported ready; installation, guided setup and guided
Guardian activation completed without manual systemctl intervention. Fresh
canonical Guardian evidence was produced, the service remained enabled and
active, and `filesystem.local_semantics` remained satisfied.

Independent evidence proved the canonical state root was a real non-symlink,
current-user-owned mode-0700 directory reached through safe components;
configuration and state roots were distinct; systemd created no compatibility
symlink or migration; and Guardian reported success with zero restarts. This is
the achieved RC.5 core clean-host functional proof. It is not a substitute for
the incomplete formal 26-checkpoint certification.

| CP | Evidence requirement | State |
| --- | --- | --- |
| 01 | private receipt and exact two-file provenance | PASS; EXACT TWO REGULAR NON-SYMLINK FILES |
| 02 | archive checksum/sidecar | PASS; SIZE AND BOTH DIGESTS EXACT, SIDECAR OK |
| 03 | canonical safe archive layout | PASS; EXACT 25-MEMBER SAFE LISTING AND METADATA |
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
| QWSG-055-F001 | OPEN, BLOCKING | Checkpoints 14–16: real safe state directory, no compatibility symlink, guided activation, enabled/active/fresh evidence and satisfied filesystem assessment | ADDITIVE EXTERNALLY VERIFIED CORRECTED IN RC.5; HISTORICAL OPEN/BLOCKING IMMUTABLE; FORMAL RESTART INCOMPLETE |
| QWSG-051-F001 | OPEN, BLOCKING | Checkpoint 17: documented guided workflow and independent enabled/active/fresh readiness | ADDITIVE EXTERNALLY VERIFIED CORRECTED IN RC.5; HISTORICAL OPEN/BLOCKING IMMUTABLE; FORMAL RESTART INCOMPLETE |
| QWSG-053-F001 | OPEN, BLOCKING | Gates B/C: ordinary no-`.git` builds without GOFLAGS and exact deterministic provenance | LOCAL GATE B RETEST PASS; historical state remains OPEN/BLOCKING pending later acceptance boundaries |
| QWSG-049-F002 | OPEN, BLOCKING | Checkpoints 07–08: fresh actionable user-manager guidance | NOT EXECUTED |
| QWSG-049-F003 | OPEN, RELEASE-GATE BLOCKING | Checkpoints 07–08 and 15–17: fresh local-semantics assessment before and after activation | NOT EXECUTED |

Historical evidence remains unchanged. An external RC.5 PASS is additive and
does not rewrite the original finding chronology.

## Task 056 / QWSG-055-F001 proof ledger

| Requirement | State |
| --- | --- |
| Canonical state root is a real directory | NOT EXECUTED IN CLEAN-HOST RESTART |
| No state-root or unsafe path-component symlink | NOT EXECUTED IN CLEAN-HOST RESTART |
| Intended ordinary-user ownership | NOT EXECUTED IN CLEAN-HOST RESTART |
| Mode 0700 | NOT EXECUTED IN CLEAN-HOST RESTART |
| Configuration and state roots distinct | NOT EXECUTED IN CLEAN-HOST RESTART |
| No systemd compatibility symlink/migration | NOT EXECUTED IN CLEAN-HOST RESTART |
| Guided activation without manual systemctl repair | NOT EXECUTED IN CLEAN-HOST RESTART |
| Service enabled and active | NOT EXECUTED IN CLEAN-HOST RESTART |
| Fresh integrity-checked canonical Guardian evidence | NOT EXECUTED IN CLEAN-HOST RESTART |
| Post-activation filesystem.local_semantics satisfied | NOT EXECUTED IN CLEAN-HOST RESTART |

The earlier RC.5 attempt passed every row and remains additive evidence. The
Owner-authorized clean-host restart must prove every row again before the
current run may close QWSG-055-F001. Historical evidence remains immutable.

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

The Owner confirmed the disposable Ubuntu 24.04 amd64 VPS was fully
reinstalled/reset after RC.4 and had no QWSG-specific preparation before the
first RC.5 execution. The later formal restart encountered previously created
acceptance state at Checkpoint 04. The Owner terminated the disproportionately
procedural sequence rather than reinstalling solely to reproduce chronology.
This is a procedural/evidence limitation and does not demonstrate a product
defect.

## Mandatory post-acceptance distribution follow-up

A later separately authorized release/distribution phase must provide stable
official archive and SHA-256 sidecar URLs from `git.quantumwizard.hu`/Forgejo,
`wget` and `curl` examples, command-line checksum verification, future Smart
Installer compatibility, and ordinary installation without workstation-mediated
artifact copying. Gate A does not implement or publish this mechanism.

## Final decision gate

Task 057 is `COMPLETE WITH DISCLOSED ACCEPTANCE LIMITATIONS`. `RC.5 CORE
CLEAN-HOST FUNCTIONAL PROOF: ACHIEVED`, while `FORMAL 26-CHECKPOINT
RELEASE-READINESS ACCEPTANCE: INCOMPLETE`. The terminal verdict is `NOT READY
FOR QWSG 1.1.0 RELEASE` because formal certification was terminated incomplete
by Owner decision. No product defect is inferred solely from missing procedural
evidence. This verdict authorizes no tag, Forgejo Release, upload, publication
or announcement.
