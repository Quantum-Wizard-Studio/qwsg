# QWSG 1.1.0-rc.2 Acceptance and Release-Readiness Record

## Current state

- Candidate identity: `1.1.0-rc.2`.
- Candidate source commit: `6d3f79accd4d52b94c960eefa93e2f51fbc9a48c`.
- Commit-derived build epoch: `1787293383` (`2026-08-21T06:23:03Z`).
- Candidate archive: `qwsg-1.1.0-rc.2-linux-amd64.tar.gz` — privately built
  twice; not transferred or published.
- Archive size: `2945629` bytes.
- Archive SHA-256:
  `73d045cbc5577d3e9921a44760ba316d2094cf13fafe82f873be9f3600547315`.
- Private transfer: `NOT EXECUTED`.
- External clean-host execution: `STARTED — STOPPED AT GUIDED GUARDIAN ACTIVATION`.
- F002/F003 external retest: `NOT EXECUTED`.
- Real SMTP receipt: `NOT EXECUTED`.
- Physical reboot: `NOT EXECUTED`.
- Release-readiness verdict: `NOT READY FOR QWSG 1.1.0 RELEASE` because guided
  Guardian activation failed and later mandatory gates remain unexecuted.

This record is the independent RC.2 ledger. QWSG `1.1.0-rc.1`, source
`ff2eb2b12499f5daf3b5ba11b1f8d7fc562f8a31`, archive SHA-256
`aa139faaccc1cc85b50cfe0eedee9436539ae1c3071e01d8e9ed9283fc7f8239`,
and all Task 049 evidence remain immutable historical records.

## Owner gate ledger

| Gate | Authority | State | Evidence |
| --- | --- | --- | --- |
| A1 — readiness audit and scaffolding | Task 051 approved scope | COMPLETE | local repository validation |
| A2 — source staging/commit/push | separate explicit Owner authorization | COMPLETE | commit `6d3f79accd4d52b94c960eefa93e2f51fbc9a48c`; `origin/main` verified |
| B — private twin candidate construction | separate explicit Owner authorization | COMPLETE | two private commit exports; byte-identical result |
| C — private transfer | separate explicit Owner authorization | BLOCKED BEFORE TRANSFER | host key reconciled; protected interactive authentication unavailable in this session |
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
| Exact clean integrated source commit | `6d3f79accd4d52b94c960eefa93e2f51fbc9a48c`; parent `af8154140ba934cfa0b89aa7071633c87323ecb9` | PASS |
| Commit-derived `SOURCE_DATE_EPOCH` | `1787293383` (`2026-08-21T06:23:03Z`) | PASS |
| Build-one private export/root | `/tmp/qwsg-task051-rc2-build-one.Qr9pC5` (mode 0700) | PASS |
| Build-two private export/root | `/tmp/qwsg-task051-rc2-build-two.3AHmql` (mode 0700) | PASS |
| Binary version/full commit/build time | `1.1.0-rc.2`; exact commit; `2026-08-21T06:23:03Z` | PASS |
| Binary static linux-amd64 identity | ELF 64-bit x86-64, statically linked, stripped; no dynamic executable | PASS |
| Build-one archive SHA-256 | `73d045cbc5577d3e9921a44760ba316d2094cf13fafe82f873be9f3600547315` | PASS |
| Build-two archive SHA-256 | `73d045cbc5577d3e9921a44760ba316d2094cf13fafe82f873be9f3600547315` | PASS |
| Archive filename and size | `qwsg-1.1.0-rc.2-linux-amd64.tar.gz`; `2945629` bytes | PASS |
| Binary/manifest/archive/sidecar byte identity | all byte-identical | PASS |
| Sidecar and internal manifest | both sidecars verify; both complete manifests verify and match | PASS |
| Binary and manifest SHA-256 | binary `bade6dfc92418784f18ce4d5f495c1c6580f3fad56bac00beae7aeed4422643f`; manifest `2c016761d54f69355631665e1099fec0bb489fb97943b61536c7416dd412f798` | PASS |
| Safe canonical archive layout and modes | single RC.2 root; relative paths; regular files/directories only; deterministic owner/time/order | PASS |
| LICENSE/root README/root INSTALL/RC.2 notes | present and byte-identical to exact source export | PASS |
| Secret/private/Builder/snapshot/backup/cache/Owner exclusion | fixed package allowlist audited; excluded | PASS |
| RC.1 and v1.0.0 preservation | historical identities unchanged; no RC.1 collision | PASS |

Candidate bytes may be constructed only from two independent exports of the
exact Gate A2 commit with its committer timestamp as `SOURCE_DATE_EPOCH` and
its full lowercase 40-character identity as `BUILD_COMMIT`. No uncommitted
overlay is candidate source.

Gate B validation completed at `2026-08-21T06:31:28Z`. An initial invocation
from the development working directory failed safely before output because Go
module resolution requires the exported source root as the working directory.
Both output directories were confirmed empty, and the successful independent
builds were launched from their respective unmodified exports with identical
authorized provenance inputs and separate caches.

## Private transfer ledger

| Evidence | Value | Status |
| --- | --- | --- |
| Owner-approved destination and account | private; not retained here | NOT EXECUTED |
| Direct standard SSH/SCP safety assessment | verified ED25519 key accepted under strict checking; password authentication then required | BLOCKED |
| Transfer path | direct VPS-to-VPS preferred; Owner-workstation fallback | NOT EXECUTED |
| Exact transferred allowlist | RC.2 archive and sidecar only | NOT EXECUTED |
| Destination checksum | pending | NOT EXECUTED |

The first strict SSH connection attempt stopped before authentication, remote
command execution, directory creation, or file transfer because the presented
host key conflicted with the existing trusted entry. The fingerprint and
destination identity are not retained in repository evidence. Independent
Owner verification and separate authorization to update the exact known-host
entry are required before Gate C may resume. No bypass or alternate transfer
method was attempted.

The Owner independently verified the presented ED25519 fingerprint and
authorized exact destination known-host reconciliation. Only the three stale
destination entries were removed, and the new ED25519 key was accepted after
the strict interactive prompt displayed the exact verified fingerprint. SSH
then required password authentication. This agent session has no protected
Owner-only credential input boundary, so the prompt was aborted without a
password, authentication, remote command, directory creation, or transfer.

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
| 09 | explicit Guardian activation/fresh evidence | external-host | FAIL: QWSG-051-F001 |
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
| QWSG-051-F001 | Guided Guardian activation | RELEASE BLOCKER | `qwsg setup` preserved the valid configuration but its explicitly confirmed internal Guardian activation failed even though readiness independently reported the user manager running, unit installed, installation/environment/configuration ready, and only enabled/active/canonical-evidence missing. The activation controller strips the validated user-runtime environment used by the readiness probe, so its fixed `systemctl --user` call cannot reliably reach the same user bus. Its generic error also does not identify whether daemon-reload or enable/start failed. | Stop external acceptance. Do not use the readiness command as an operator workaround. Product correction and replacement candidate require separate Owner authority. | OPEN, BLOCKING |

New findings use `RELEASE BLOCKER`, `SECURITY DEFECT`, `FUNCTIONAL DEFECT`,
`UX/DOCUMENTATION DEFECT`, or `COSMETIC / POST-RELEASE CANDIDATE`. Product or
security defects outside Task 051 authority stop at the safe checkpoint and
require a separately authorized correction task; acceptance does not silently
repair them.

### QWSG-051-F001 architecture assessment

Owner-supplied privacy-safe evidence establishes that guided setup wrote the
configuration, explicit activation was accepted, activation failed, and the
product instructed the Owner to run `qwsg readiness`. The complete raw
readiness output was not included; the supplied canonical facts are sufficient
for the boundary diagnosis: configuration and environment were ready,
`systemd.user_manager` was satisfied/running, and the installed unit was
present, while enabled, active and canonical Guardian evidence remained
missing.

Readiness constructs its systemd runner through the assessment host boundary,
which validates the current UID runtime directory and supplies only the trusted
`XDG_RUNTIME_DIR=/run/user/<uid>` value to fixed `systemctl --user` probes.
Guided setup instead calls `userservice.New().Activate`. That controller creates
a separate bounded runner with no trusted runtime environment. The common
runner replaces the child environment with only `PATH`, `LANG` and `LC_ALL`,
so both activation calls omit `XDG_RUNTIME_DIR`. This explains how readiness
can reach and classify the running user manager while the setup-owned
`daemon-reload` or `enable --now` call fails. The controller discards bounded
stderr and returns only a generic stage-independent error, so external evidence
cannot distinguish the first call from the second.

The command displayed by readiness is product guidance but is not accepted as
a workaround for this finding. Checkpoint 09 requires the explicitly confirmed
guided activation path itself plus fresh canonical evidence. Task 051 is an
acceptance/retest task and does not authorize this product-source correction or
an in-acceptance repair/rebuild/retest loop.

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
