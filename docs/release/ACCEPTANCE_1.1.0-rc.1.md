# QWSG 1.1.0-rc.1 Acceptance and Release-Readiness Record

## Current state

- Candidate identity: `1.1.0-rc.1`.
- Candidate source commit: `ff2eb2b12499f5daf3b5ba11b1f8d7fc562f8a31`.
- Candidate archive: `qwsg-1.1.0-rc.1-linux-amd64.tar.gz`, privately transferred to the Owner-controlled disposable host; not published.
- External clean-host execution: `STARTED — STOPPED AT SMART INSTALL`.
- Real SMTP provider execution: `NOT STARTED`.
- Physical reboot execution: `NOT STARTED`.
- Release-readiness verdict: `NOT READY FOR QWSG 1.1.0 RELEASE`.

The current verdict reflects failed Smart Install operator guidance plus
unexecuted later mandatory gates. The Project Owner supplied privacy-safe
external output from the clean host. QWSG correctly refused readiness, but the
documented/product-guided journey cannot continue without prohibited developer
coaching.

## Candidate provenance ledger

| Evidence | Value | Status |
| --- | --- | --- |
| Exact source commit | `ff2eb2b12499f5daf3b5ba11b1f8d7fc562f8a31` | PASS |
| Commit-derived `SOURCE_DATE_EPOCH` | `1787237840` (`2026-08-20T14:57:20Z`) | PASS |
| Binary version and embedded commit | `1.1.0-rc.1`, exact full commit | PASS |
| Build-one archive SHA-256 | `aa139faaccc1cc85b50cfe0eedee9436539ae1c3071e01d8e9ed9283fc7f8239` | PASS |
| Build-two archive SHA-256 | `aa139faaccc1cc85b50cfe0eedee9436539ae1c3071e01d8e9ed9283fc7f8239` | PASS |
| Twin-build byte identity | archive and sidecar byte-identical | PASS |
| Sidecar and internal manifest | all entries verified | PASS |
| Safe single-root layout | regular files/directories under exact root | PASS |
| Packaged LICENSE/README/INSTALL byte identity | exact candidate-source copies | PASS |
| Packaged README/INSTALL cross-reference | exact canonical copies; each explicitly names its archive-root counterpart; repository-context Markdown targets do not resolve inside the archive | PASS with cosmetic QWSG-049-F001 |
| v1.0.0 artifact distinction/preservation | different name/hash; protected identities unchanged | PASS |

## Gate matrix

| Gate | Evidence class | Requirement | Result |
| --- | --- | --- | --- |
| Local build/test/race/vet/format/security/governance | local automated | mandatory | PASS |
| Commit-pure twin candidate build | private staged | mandatory | PASS |
| Documentation-only clean install journey | external host | mandatory | not executed |
| Smart Install classifications and next action | external host | mandatory | FAIL: QWSG-049-F002 and QWSG-049-F003 |
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
| QWSG-049-F001 | Pre-transfer packaged-document verification | COSMETIC / POST-RELEASE CANDIDATE | Exact canonical copies retain repository-context Markdown targets that do not resolve inside the archive. Both documents explicitly name their archive-root counterpart and installed paths in prose, so no operator action or guidance is missing. | Safe to continue after the separate transfer/external-execution gate. | OPEN, NON-BLOCKING |
| QWSG-049-F002 | External Smart Install | RELEASE BLOCKER | `systemd.user_manager` was correctly classified `missing_required`, but QWSG emitted only `resolve_required_findings`. It provided no cause, verification action, privilege boundary, or proven supported remediation. The clean-host journey cannot continue without developer coaching. | Stop physical acceptance. Do not supply a manual workaround. Product correction and replacement candidate require separate Owner authority. | OPEN, BLOCKING |
| QWSG-049-F003 | External Smart Install | UX/DOCUMENTATION DEFECT | `filesystem.local_semantics` was reported `unknown_requires_verification`, but QWSG supplied no bounded verification instruction or evidence needed to resolve the uncertainty. The requirement is recommended/non-blocking at runtime, but mandatory Smart Install manual-verification acceptance did not pass. | Stop with QWSG-049-F002. Include this guidance gap in the same bounded product correction. | OPEN, RELEASE-GATE BLOCKING |

## External Smart Install evidence

Evidence class: `external host`, privacy-reviewed Owner-supplied output. No host
identity, path, credential, SMTP detail, or raw inventory was retained.

```text
QWSG install readiness
Platform: ubuntu-24.04-amd64
filesystem.local_semantics         unknown_requires_verification
platform.architecture              satisfied (amd64)
platform.operating_system          satisfied (Ubuntu 24.04)
runtime.glibc                      satisfied (2.39)
runtime.non_root                   satisfied
systemd.user_manager               missing_required
systemd.version                    satisfied (255)
environment_dependencies: not_ready
installation: not_ready
Next actions:
- resolve_required_findings
- rerun_qwsg_install_check
```

Classification and fail-closed readiness are correct. Guidance is not. The
Owner correctly declined to inject a Linux/systemd workaround from engineering
knowledge. No later physical checkpoint is authorized or executed.

## Architecture assessment for blocking findings

The canonical registry already owns structured `Remediation` values and the
human/JSON presenter already renders attached commands, elevation metadata and
revalidation. However, both affected requirements have empty remediation
lists. This is absent canonical data, not a rendering-only omission.

For `systemd.user_manager`, the current fixed probe maps materially different
failure causes to the single `systemd_user_unavailable` token. That evidence is
insufficient to prove one safe remediation command, even on recognized Ubuntu
24.04 amd64. Adding a guessed command would violate the remediation-safety
contract.

For `filesystem.local_semantics`, assessment unconditionally emits
`filesystem_manual_verification_required`; no read-only probe or structured
manual-verification instruction exists. The current model requires a command
for every `Remediation`, so it cannot express a non-command verification plan.

The smallest coherent correction is product work: add versioned canonical
manual-verification guidance to the assessment model/registry and human/JSON
presentation; distinguish bounded user-manager failure evidence sufficiently
to select only proven guidance; provide exact privilege/revalidation metadata;
and add a deterministic filesystem verification plan or a justified read-only
probe. Tests must prove no guessed command, no automatic execution, and useful
next actions for both findings. This exceeds Task 049's acceptance-tooling and
release-plumbing correction authority.

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
