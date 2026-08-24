# QWSG 1.1.0-rc.4 Acceptance and Release-Readiness Record

## Current state

- Candidate: `NOT BUILT`
- Transfer: `NOT STARTED`
- External acceptance: `NOT STARTED`
- Final verdict: `PENDING`
- Historical `QWSG-053-F001`: `OPEN, BLOCKING`; Task 054 integrated the
  correction, but RC.4 construction must freshly prove the affected boundary.
- Historical `QWSG-051-F001`: `OPEN, BLOCKING`; RC.4 must directly retest
  guided activation and independent enabled/active/fresh Guardian evidence.
- Task 049 `QWSG-049-F002` and `QWSG-049-F003`: immutable history; fresh RC.4
  Smart Install regression is pending.

No entry authorizes construction, transfer, VPS access, credentials, external
execution, tag, Forgejo Release, upload or publication.

## Owner gate ledger

| Gate | Scope | State | Evidence |
| --- | --- | --- | --- |
| A | readiness audit and acceptance-source scaffolding integration | AUDIT COMPLETE; INTEGRATION PENDING OWNER AUTHORIZATION | Task 055 Gate A report |
| B | private exact-commit twin construction | NOT AUTHORIZED | — |
| C | deterministic/package/provenance/security proof | NOT AUTHORIZED | — |
| D | private transfer and destination integrity | NOT AUTHORIZED | — |
| E | Checkpoints 01–25 external clean-host acceptance | NOT AUTHORIZED | — |
| F | privacy-safe evidence integration and verdict | NOT AUTHORIZED | — |

## Source and candidate provenance

| Field | Value |
| --- | --- |
| Gate A baseline | `12069b16cc574c759a40f905d2b4981bd729716d` |
| Task 054 integration commit | `ef513dde187e4119b6aa04a3439a879056f6cc69` |
| Candidate-source commit | PENDING GATE A INTEGRATION |
| Source version | `1.1.0-rc.4` |
| SOURCE_DATE_EPOCH / UTC build time | NOT BUILT |
| Archive | `qwsg-1.1.0-rc.4-linux-amd64.tar.gz` — NOT BUILT |
| Size / SHA-256 | NOT BUILT |
| Binary/manifest/archive/sidecar byte identity | NOT TESTED |
| Static Linux amd64 / embedded source | NOT TESTED |
| Package/layout/docs/LICENSE/exclusions | NOT TESTED |
| Exported-source build contract | LOCAL GATE A PASS; exact clean candidate-source retest required at Gate B |

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
| 05 | LICENSE, root docs, RC.4/binary/source identity | NOT EXECUTED |
| 06 | README/INSTALL operator journey | NOT EXECUTED |
| 07 | clean-host `install --check` | NOT EXECUTED |
| 08 | F002/F003 product-guidance regression | NOT EXECUTED |
| 09 | immutable installation | NOT EXECUTED |
| 10 | setup interruption/resume and one-recipient config | NOT EXECUTED |
| 11 | protected mode-0600 credential-file workflow | NOT EXECUTED |
| 12 | notification preflight | NOT EXECUTED |
| 13 | real external SMTP test and Owner receipt | NOT EXECUTED |
| 14 | explicit guided Guardian activation | NOT EXECUTED |
| 15 | QWSG-051-F001 external correction proof | NOT EXECUTED |
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

## Historical blocker retest ledger

- `QWSG-053-F001` remains immutable historical `OPEN, BLOCKING`. Gate B must
  freshly prove ordinary `make build` in genuine no-`.git` exports with
  `GOFLAGS` unset, truthful unknown defaults, exact explicit identity,
  controlled byte identity and no ambient Go `vcs.*` settings.
- `QWSG-051-F001` remains immutable historical `OPEN, BLOCKING`. It may be
  recorded here as `EXTERNALLY VERIFIED CORRECTED` only after documented guided
  activation succeeds and independent readiness proves enabled, active and
  fresh canonical evidence. Manual activation or local tests do not qualify.
- Task 049 F002/F003 history is never edited. RC.4 requires a fresh uncoached
  Smart Install run proving actionable user-manager guidance and bounded
  filesystem local-semantics verification.

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

| ID | Boundary | Severity | Historical state | RC.4 retest state |
| --- | --- | --- | --- | --- |
| QWSG-053-F001 | exported-source ordinary build | RELEASE BLOCKER | OPEN, BLOCKING | PENDING GATE B |
| QWSG-051-F001 | guided Guardian activation | RELEASE BLOCKER | OPEN, BLOCKING | NOT EXECUTED |
| QWSG-049-F002 | Smart Install user-manager guidance | RELEASE BLOCKER | OPEN, BLOCKING | NOT EXECUTED |
| QWSG-049-F003 | filesystem local-semantics guidance | UX/DOCUMENTATION DEFECT; RELEASE-GATE BLOCKING | OPEN, RELEASE-GATE BLOCKING | NOT EXECUTED |

New product or security findings stop acceptance without repair under Task 055.

## Final decision gate

`READY FOR QWSG 1.1.0 RELEASE` is permitted only when all six Owner gates and
all 25 checkpoints pass; historical affected boundaries pass fresh RC.4
retests; both real receipts are independently confirmed; reboot and lifecycle
operations pass; and no blocker/security defect remains.

Otherwise record `NOT READY FOR QWSG 1.1.0 RELEASE` with exact findings or
missing evidence. Neither verdict authorizes a tag, Forgejo Release, upload,
publication or announcement.
