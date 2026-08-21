# QWSG 1.1.0-rc.2 Acceptance and Release-Readiness Record

## Current state

- Candidate identity: `1.1.0-rc.2`.
- Candidate source commit: `PENDING OWNER GATE A2`.
- Commit-derived build epoch: `PENDING`.
- Candidate archive: `qwsg-1.1.0-rc.2-linux-amd64.tar.gz` — `NOT BUILT`.
- Archive SHA-256: `PENDING`.
- Private transfer: `NOT EXECUTED`.
- External clean-host execution: `NOT EXECUTED`.
- F002/F003 external retest: `NOT EXECUTED`.
- Real SMTP receipt: `NOT EXECUTED`.
- Physical reboot: `NOT EXECUTED`.
- Release-readiness verdict: `NOT READY FOR QWSG 1.1.0 RELEASE` while mandatory
  gates remain unexecuted.

This record is the independent RC.2 ledger. QWSG `1.1.0-rc.1`, source
`ff2eb2b12499f5daf3b5ba11b1f8d7fc562f8a31`, archive SHA-256
`aa139faaccc1cc85b50cfe0eedee9436539ae1c3071e01d8e9ed9283fc7f8239`,
and all Task 049 evidence remain immutable historical records.

## Owner gate ledger

| Gate | Authority | State | Evidence |
| --- | --- | --- | --- |
| A1 — readiness audit and scaffolding | Task 051 approved scope | IN PROGRESS | local repository only |
| A2 — source staging/commit/push | separate explicit Owner authorization | NOT AUTHORIZED | none |
| B — private twin candidate construction | separate explicit Owner authorization | NOT AUTHORIZED | none |
| C — private transfer | separate explicit Owner authorization | NOT AUTHORIZED | none |
| D1 — external read-only Checkpoints 01–04 | separate explicit Owner authorization | NOT AUTHORIZED | none |
| D2 — install and guided setup | separate explicit Owner authorization | NOT AUTHORIZED | none |
| D3 — protected SMTP and receipt | separate explicit Owner authorization | NOT AUTHORIZED | none |
| D4 — Guardian, logout and physical reboot | separate explicit Owner authorization | NOT AUTHORIZED | none |
| D5 — uninstall and reinstall | separate explicit Owner authorization | NOT AUTHORIZED | none |
| Final publication | separate release task/Owner authority | OUT OF SCOPE | none |

Each later gate is independent. Passing or authorizing one does not authorize
the next.

## Candidate provenance ledger

| Evidence | Value | Status |
| --- | --- | --- |
| Exact clean integrated source commit | pending | NOT EXECUTED |
| Commit-derived `SOURCE_DATE_EPOCH` | pending | NOT EXECUTED |
| Build-one private export/root | pending | NOT EXECUTED |
| Build-two private export/root | pending | NOT EXECUTED |
| Binary version/full commit/build time | pending | NOT EXECUTED |
| Binary static linux-amd64 identity | pending | NOT EXECUTED |
| Build-one archive SHA-256 | pending | NOT EXECUTED |
| Build-two archive SHA-256 | pending | NOT EXECUTED |
| Binary/manifest/archive/sidecar byte identity | pending | NOT EXECUTED |
| Sidecar and internal manifest | pending | NOT EXECUTED |
| Safe canonical archive layout and modes | pending | NOT EXECUTED |
| LICENSE/root README/root INSTALL/RC.2 notes | pending | NOT EXECUTED |
| Secret/private/Builder/snapshot/backup/Owner exclusion | pending | NOT EXECUTED |
| RC.1 and v1.0.0 preservation | pending | NOT EXECUTED |

Candidate bytes may be constructed only from two independent exports of the
exact Gate A2 commit with its committer timestamp as `SOURCE_DATE_EPOCH` and
its full lowercase 40-character identity as `BUILD_COMMIT`. No uncommitted
overlay is candidate source.

## Private transfer ledger

| Evidence | Value | Status |
| --- | --- | --- |
| Owner-approved destination and account | private; not retained here | NOT EXECUTED |
| Direct standard SSH/SCP safety assessment | pending | NOT EXECUTED |
| Transfer path | direct VPS-to-VPS preferred; Owner-workstation fallback | NOT EXECUTED |
| Exact transferred allowlist | RC.2 archive and sidecar only | NOT EXECUTED |
| Destination checksum | pending | NOT EXECUTED |

Normal strict host-key verification is mandatory. No credential, host key,
private host/account identifier, SSH configuration change, new transfer
software, public exposure, or additional payload belongs in this record.

## External checkpoint ledger

Execute only through
`ACCEPTANCE_PROTOCOL_1.1.0-rc.2.md`, one Owner-operated checkpoint at a time.

| Checkpoint | Mandatory coverage | Evidence class | Result |
| --- | --- | --- | --- |
| 01 | private candidate receipt/provenance | external-host | NOT EXECUTED |
| 02 | checksum and safe archive layout | external-host | NOT EXECUTED |
| 03 | manifest, LICENSE, README and INSTALL usability | external-host | NOT EXECUTED |
| 04 | Smart Install plus F002/F003 retest | external-host | NOT EXECUTED |
| 05 | immutable installation and handoff | external-host | NOT EXECUTED |
| 06 | guided setup interruption/resume/invalid input | external-host | NOT EXECUTED |
| 07 | one-recipient configuration/protected credential/preflight | external-host | NOT EXECUTED |
| 08 | real external test-email receipt | real-provider | NOT EXECUTED |
| 09 | explicit Guardian activation/fresh evidence | external-host | NOT EXECUTED |
| 10 | systemd process/invocation/cadence/resources/restarts | external-host | NOT EXECUTED |
| 11 | lingering guidance and logout/session behavior | external-host | NOT EXECUTED |
| 12 | physical reboot/new identity/fresh post-reboot evidence | physical-reboot | NOT EXECUTED |
| 13 | post-reboot notification receipt continuity | real-provider | NOT EXECUTED |
| 14 | explicit Guardian restart and fresh recovery | external-host | NOT EXECUTED |
| 15 | safe uninstall and preserved user data | external-host | NOT EXECUTED |
| 16 | reinstall/resume/reactivation/no stale READY | external-host | NOT EXECUTED |

## F002/F003 special retest

Run `./bin/qwsg install --check` on the real supported external VPS without
developer coaching.

`QWSG-049-F002` may be marked externally `VERIFIED/CORRECTED` for RC.2 only
when `systemd.user_manager` has the correct evidence-based classification, no
stripped-session false negative, cause-specific explanation, bounded
verification, explicit privilege boundary, mandatory revalidation, and a
remediation command only when proven safe. Ambiguous evidence must never gain
a guessed command.

`QWSG-049-F003` may be marked externally `VERIFIED/CORRECTED` for RC.2 only
when `filesystem.local_semantics` presents actual bounded read-only evidence
where possible and, otherwise, a precise bounded manual verification procedure
with mandatory revalidation. The assessment must not mutate the host or end in
an unconditional unexplained unknown.

Human and JSON output must agree. Historical RC.1 findings are never edited or
closed by this RC.2 result.

## SMTP and privacy ledger

Community acceptance uses one administrator recipient. SMTP credentials remain
on the external VPS within the protected credential boundary and never enter
chat, argv, Git, task history, or acceptance evidence. Retained evidence must
redact recipient, account, provider/host identity, provider headers, credential
references, tokens, and private host/account identifiers.

SMTP acceptance requires both QWSG controlled-test success and independent
Owner confirmation that the intended message arrived. TLS/authentication alone
is insufficient. A second independently confirmed receipt after physical
reboot is mandatory.

## Guardian, reboot and reinstall evidence

PASS requires explicit activation, configured cadence, fresh canonical
Guardian evidence correlated with actual systemd process/invocation state,
bounded resource and restart behavior, product-visible lingering guidance,
logout/session behavior, physical reboot, new post-reboot process/invocation
identity, fresh post-reboot evidence, notification continuity, explicit
restart, safe uninstall, preserved per-user configuration/credential/state,
and same-candidate reinstall/resume with explicit reactivation and no stale
READY. ActiveState alone is never Guardian-health evidence.

## Finding register

| ID | Checkpoint | Severity | Summary | Continuation | Status |
| --- | --- | --- | --- | --- | --- |
| none | — | — | No RC.2 external finding recorded | — | — |

New findings use `RELEASE BLOCKER`, `SECURITY DEFECT`, `FUNCTIONAL DEFECT`,
`UX/DOCUMENTATION DEFECT`, or `COSMETIC / POST-RELEASE CANDIDATE`. Product or
security defects outside Task 051 authority stop at the safe checkpoint and
require a separately authorized correction task; acceptance does not silently
repair them.

## Final decision gate

`READY FOR QWSG 1.1.0 RELEASE` requires every mandatory checkpoint and all
local reproducibility/security/governance gates to PASS, external F002/F003
verification, real SMTP receipt before and after physical reboot, fresh
Guardian/reboot/restart evidence, uninstall/reinstall, and no open RELEASE
BLOCKER or SECURITY DEFECT.

While any mandatory gate is unexecuted, failed, or unresolved, the verdict is
`NOT READY FOR QWSG 1.1.0 RELEASE` with exact findings. Neither verdict
authorizes a final tag, Forgejo Release, public upload, announcement, signing
claim, publication, or Task 052.
