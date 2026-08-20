# QWSG 1.1.0-rc.1 Acceptance and Release-Readiness Record

## Current state

- Candidate identity: `1.1.0-rc.1`.
- Candidate source commit: `[PENDING SEPARATE OWNER-AUTHORIZED INTEGRATION]`.
- Candidate archive: `[NOT BUILT]`.
- External clean-host execution: `NOT STARTED`.
- Real SMTP provider execution: `NOT STARTED`.
- Physical reboot execution: `NOT STARTED`.
- Release-readiness verdict: `NOT READY FOR QWSG 1.1.0 RELEASE`.

The current verdict reflects unexecuted mandatory gates, not a product failure.
No candidate may be built or transferred until its exact clean integrated
source commit exists and the Project Owner separately authorizes the candidate
build gate.

## Candidate provenance ledger

| Evidence | Value | Status |
| --- | --- | --- |
| Exact source commit | pending | NOT EXECUTED |
| Commit-derived `SOURCE_DATE_EPOCH` | pending | NOT EXECUTED |
| Binary version and embedded commit | pending | NOT EXECUTED |
| Build-one archive SHA-256 | pending | NOT EXECUTED |
| Build-two archive SHA-256 | pending | NOT EXECUTED |
| Twin-build byte identity | pending | NOT EXECUTED |
| Sidecar and internal manifest | pending | NOT EXECUTED |
| Safe single-root layout | pending | NOT EXECUTED |
| Packaged LICENSE/README/INSTALL identity | pending | NOT EXECUTED |
| v1.0.0 artifact distinction/preservation | pending | NOT EXECUTED |

## Gate matrix

| Gate | Evidence class | Requirement | Result |
| --- | --- | --- | --- |
| Local build/test/race/vet/format/security/governance | local automated | mandatory | pending |
| Commit-pure twin candidate build | private staged | mandatory | not executed |
| Documentation-only clean install journey | external host | mandatory | not executed |
| Smart Install classifications and next action | external host | mandatory | not executed |
| Guided setup interruption/resume | external host | mandatory | not executed |
| Protected credential boundary | external host | mandatory | not executed |
| Actual external SMTP receipt | real provider | mandatory for unconditional READY | not executed |
| Explicit Guardian activation and fresh evidence | external host | mandatory | not executed |
| Lingering guidance and logout | external host | mandatory | not executed |
| Physical reboot/new invocation/fresh cycle | physical reboot | mandatory | not executed |
| Post-reboot notification receipt | real provider | mandatory | not executed |
| Explicit restart and recovery | external host | mandatory | not executed |
| Uninstall preservation | external host | mandatory | not executed |
| Reinstall/resume without stale READY | external host | mandatory | not executed |

## Finding register

| ID | Checkpoint | Severity | Summary | Continuation | Status |
| --- | --- | --- | --- | --- | --- |
| — | — | — | No physical acceptance findings yet | — | NOT STARTED |

Use only the severities and continuation rules in
`ACCEPTANCE_PROTOCOL_1.1.0-rc.1.md`. Record an undocumented intervention before
providing it. Never include secrets or private host/provider data.

## Evidence index

For each checkpoint record only privacy-reviewed evidence with its UTC time,
candidate/source identity, evidence class, PASS/FAIL/NOT EXECUTED state,
finding IDs, and retained private-evidence reference. Local fixture evidence,
external host evidence, physical reboot evidence, and real-provider evidence
must remain distinguishable.

## Publication boundary

Task 049 produces technical acceptance evidence only. It creates no public or
final `v1.1.0` tag, Forgejo Release, upload, publication, signing claim, or
announcement. Final release work requires separate Project Owner authority.
