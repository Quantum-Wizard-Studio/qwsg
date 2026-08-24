# QWSG 1.1.0-rc.4 Acceptance and Release-Readiness Record

## Current state

- Candidate: `BUILT PRIVATELY; TRANSFERRED THROUGH OWNER-WORKSTATION FALLBACK`
- Transfer: `PASS — exact archive and sidecar verified at private destination`
- External acceptance: `STOPPED — QWSG-055-F001 BLOCKING`
- Final verdict: `PENDING GATE F`; current RC.4 evidence is `NOT READY`
- Historical `QWSG-053-F001`: `OPEN, BLOCKING`; Task 054 integrated the
  correction, but RC.4 construction must freshly prove the affected boundary.
- Historical `QWSG-051-F001`: `OPEN, BLOCKING`; RC.4 must directly retest
  guided activation and independent enabled/active/fresh Guardian evidence.
- Task 049 `QWSG-049-F002` and `QWSG-049-F003`: immutable history; fresh RC.4
  Smart Install regression passed before the later activation blocker.

No entry authorizes candidate rebuild, transfer, VPS access, credentials,
external execution, tag, Forgejo Release, upload or publication.

## Owner gate ledger

| Gate | Scope | State | Evidence |
| --- | --- | --- | --- |
| A | readiness audit and acceptance-source scaffolding integration | PASS; INTEGRATED | `f105771dbddccf363a63095ac4ad2a7a2285aa84` |
| B | private exact-commit twin construction | PASS AFTER ONE OWNER-AUTHORIZED CORRECTED RETRY | Original attempt preserved as acceptance-orchestration failure; corrected module-root twin construction and bounded validation passed |
| C | deterministic/package/provenance/security proof | PASS; PRIVACY-SAFE EVIDENCE INTEGRATED | Owner-accepted Gate B twin identity, package, provenance and exclusion proof |
| D | private transfer and destination integrity | PASS | Owner-workstation fallback; exact two-file receipt, type, size, archive/sidecar digests and sidecar verification passed |
| E | Checkpoints 01–25 external clean-host acceptance | STOPPED | Smart Install and installation passed; guided activation reached enable/start, then QWSG-055-F001 blocked continuation |
| F | privacy-safe evidence integration and verdict | AUTHORIZED — NOT READY | QWSG-055-F001 remains OPEN/BLOCKING; later checkpoints remain unexecuted |

## Source and candidate provenance

| Field | Value |
| --- | --- |
| Gate A baseline | `12069b16cc574c759a40f905d2b4981bd729716d` |
| Task 054 integration commit | `ef513dde187e4119b6aa04a3439a879056f6cc69` |
| Candidate-source commit | `f105771dbddccf363a63095ac4ad2a7a2285aa84` |
| Source version | `1.1.0-rc.4` |
| SOURCE_DATE_EPOCH / UTC build time | `1787571815` / `2026-08-24T11:43:35Z` |
| Archive | `qwsg-1.1.0-rc.4-linux-amd64.tar.gz` — PRIVATE; transferred only through the Owner-workstation fallback |
| Size / SHA-256 | `2949478` bytes / `7ad8a0f1be9bdbaa4403c3a816a6b474fd8e052934abd031047e4f82fe73a333` |
| Sidecar SHA-256 | `ca4daca22dd9f6accf5cdb062c3bb7aa88066f70b3988ae8fb1f189875d954bc`; archive verification PASS |
| Manifest SHA-256 | `ac1e31000e85ef406249675dc9dbdcda53ea848c571c7e4aeba4a3a4f72ad588`; all entries PASS |
| Binary SHA-256 | `3ec977090f90ec21a535e0a5c4fe158ab6eee883af4f230c6423514dd1cba92b` |
| Binary/manifest/archive/sidecar byte identity | PASS across two new independent exact-commit exports |
| Static Linux amd64 / embedded source | PASS; `1.1.0-rc.4`, full source commit, and controlled UTC build time exact |
| Package/layout/docs/LICENSE/exclusions | PASS; Gate C privacy-safe evidence accepted and integrated |
| Exported-source build contract | PASS in both genuine no-`.git` exports with `GOFLAGS` unset and truthful ordinary `unknown` provenance; historical finding remains OPEN/BLOCKING pending later acceptance boundaries |

The original Gate B attempt remains preserved as an acceptance-procedure and
orchestration failure: its ordinary exported-source build passed, but the
release script was launched directly while the process working directory
remained the canonical checkout. Go therefore rejected the exported absolute
package path as outside the selected module. No product defect was inferred.

## Transfer provenance

Gate D used the separately authorized Owner-workstation fallback. Exactly the
archive and sidecar were selected. Privacy-safe destination evidence proved two
regular non-symlink files, exact archive size and archive/sidecar digests, and
successful sidecar verification. No credential, private host/account identity,
or private destination path is retained here.

## External checkpoint ledger

| CP | Evidence requirement | State |
| --- | --- | --- |
| 01 | private receipt and exact two-file provenance | PASS |
| 02 | archive checksum/sidecar | PASS |
| 03 | canonical safe archive layout | PASS |
| 04 | internal manifest | PASS |
| 05 | LICENSE, root docs, RC.4/binary/source identity | PASS |
| 06 | README/INSTALL operator journey | PASS |
| 07 | clean-host `install --check` | PASS |
| 08 | F002/F003 product-guidance regression | PASS |
| 09 | immutable installation | PASS |
| 10 | setup interruption/resume and one-recipient config | PARTIAL — configuration written safely; later activation boundary failed |
| 11 | protected mode-0600 credential-file workflow | NOT EXECUTED |
| 12 | notification preflight | NOT EXECUTED |
| 13 | real external SMTP test and Owner receipt | NOT EXECUTED |
| 14 | explicit guided Guardian activation | FAIL — reached enable/start; Guardian then failed on unsafe state symlink |
| 15 | QWSG-051-F001 external correction proof | FAIL/BLOCKED — enabled, but not active and no fresh evidence |
| 16 | independent enabled/active/fresh readiness | FAIL/BLOCKED — QWSG-055-F001 |
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

- `QWSG-053-F001` remains immutable historical `OPEN, BLOCKING`. One
  Owner-authorized corrected Gate B retry passed ordinary `make build` in two
  genuine no-`.git` exports with `GOFLAGS` unset, truthful default provenance,
  exact release identity and controlled twin byte identity. This is bounded
  local Gate B evidence, not external clean-host correction or release readiness.
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
| QWSG-053-F001 | exported-source ordinary build | RELEASE BLOCKER | OPEN, BLOCKING | LOCAL GATE B RETEST PASS; external acceptance not claimed |
| QWSG-051-F001 | guided Guardian activation | RELEASE BLOCKER | OPEN, BLOCKING | RC.4 RETEST FAILED/BLOCKED by QWSG-055-F001; enabled but not active/fresh |
| QWSG-049-F002 | Smart Install user-manager guidance | RELEASE BLOCKER | OPEN, BLOCKING | RC.4 SMART INSTALL RETEST PASS |
| QWSG-049-F003 | filesystem local-semantics guidance | UX/DOCUMENTATION DEFECT; RELEASE-GATE BLOCKING | OPEN, RELEASE-GATE BLOCKING | RC.4 SMART INSTALL RETEST PASS; post-activation regression is QWSG-055-F001 evidence |
| QWSG-055-F001 | packaged user-service state-directory contract | RELEASE BLOCKER | NEW IN RC.4 | OPEN, BLOCKING — guided activation reached enable/start, then systemd created a compatibility symlink rejected by QWSG; restart limit reached |

New product or security findings stop acceptance without repair under Task 055.

## RC.4 stopped external diagnostic

Guided activation reached the validated user manager, `daemon-reload`, and
`enable --now`; the unit remained enabled. `ExecCondition` passed. The Guardian
then exited with `guardian_state_unsafe`; `ExecStopPost` could not find an exit
checkpoint; `Restart=on-failure` retried three times before the start limit and
left the unit failed.

`QWSG-055-F001` is an `OPEN, BLOCKING` `RELEASE BLOCKER`. In the user manager,
`StateDirectory=qwsg` encountered an absent state directory and the existing
same-name QWSG configuration directory. systemd's pre-v254 migration behavior
created a state-to-configuration compatibility symlink. QWSG correctly rejects
symlinked state storage, so the packaged unit and QWSG state-security contract
are incompatible on this clean-host sequence. The same symlink caused the
post-install `filesystem.local_semantics` change to
`unknown_requires_verification` with unsafe-path evidence. Acceptance stopped;
no manual state-path repair or service workaround qualifies.

## Final decision gate

`READY FOR QWSG 1.1.0 RELEASE` is permitted only when all six Owner gates and
all 25 checkpoints pass; historical affected boundaries pass fresh RC.4
retests; both real receipts are independently confirmed; reboot and lifecycle
operations pass; and no blocker/security defect remains.

Otherwise record `NOT READY FOR QWSG 1.1.0 RELEASE` with exact findings or
missing evidence. Neither verdict authorizes a tag, Forgejo Release, upload,
publication or announcement.

Task 055 terminal verdict: `complete with disclosed limitations — NOT READY
FOR QWSG 1.1.0 RELEASE`. QWSG-055-F001 remains OPEN/BLOCKING, and Checkpoints
11–13 and 17–25 were not executed after the mandatory stop.
