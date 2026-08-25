# QWSG 1.1.0-rc.5 Practical Release Acceptance

## Verdict

- Practical Acceptance: **FAIL**
- Candidate: `QWSG 1.1.0-rc.5`
- Release readiness: **NOT READY FOR RELEASE**
- Release/publication authority: not granted or exercised

Task 059 performed the first twelve-step Practical Release Acceptance run under
Framework 1.1.0 on a freshly reinstalled disposable Ubuntu 24.04 amd64 VPS.
This is additive RC.5 evidence. Task 057 and every earlier acceptance record
remain immutable.

## Candidate identity

- Source commit: `1025d36d05b2f6f919f0ea4ec4a7029f67536000`
- Archive: `qwsg-1.1.0-rc.5-linux-amd64.tar.gz`
- Archive size: `2951350` bytes
- Archive SHA-256: `cfe300c0f1f312d80120f74a9f24bed4a64387471bf2097ddc63d94f0fb2f7b0`
- Sidecar SHA-256: `69f3eb4bf89dc126a7eafd08354eec37a941014171b3d1d70c6e6a4cf52e5eb0`
- `MANIFEST.sha256` SHA-256: `ae51aca0bc4ddc61b0daea3a87f0acabcde5ec9fd8fadddc050f0786d6915e9e`
- Binary SHA-256: `5484aab96d5c3748e81b065fdb11ec8c34385589bb07ee7ea1b2b35fdffa6b93`

## Twelve-step result

| Step | Practical acceptance area | Result |
|---:|---|---|
| 1 | Fresh supported host baseline | PASS |
| 2 | Exact candidate receipt | PASS |
| 3 | Integrity and package safety | PASS |
| 4 | Smart Install/readiness | PASS |
| 5 | Documented installation | PASS |
| 6 | Guided setup | PASS |
| 7 | Guardian and state contract | PASS |
| 8 | Protected external notification and actual receipt | PASS |
| 9 | Physical reboot and automatic recovery | **FAIL — QWSG-059-F001** |
| 10 | Documented explicit Guardian restart | NOT EXECUTED |
| 11 | Uninstall preservation | NOT EXECUTED |
| 12 | Same-candidate reinstall | NOT EXECUTED |

## Passed product evidence

- The fresh host was supported Ubuntu 24.04 amd64 in an ordinary-user context
  with QWSG-specific installation, configuration, state and unit state absent.
- Exactly two regular non-symlink candidate receipt files matched the preserved
  archive size, archive hash, sidecar hash and sidecar verification.
- The safe 25-member archive layout, all 18 manifest entries, LICENSE,
  documentation, binary hash, version, source commit and controlled build time
  passed.
- Smart Install reported the environment and installation domains ready,
  including satisfied `filesystem.local_semantics`, supported platform,
  running user manager, systemd 255 and glibc 2.39.
- The packaged installer, guided setup and guided Guardian activation passed
  without a developer workaround or manual service substitution.
- The canonical state root and required components were non-symlink, private,
  current-user-owned and correctly moded. Configuration and state remained
  separate; no compatibility migration occurred; service enablement/activity
  and fresh pre-notification canonical evidence passed.
- Protected credential storage, TCP, STARTTLS, certificate trust, AUTH PLAIN,
  QWSG SMTP acceptance and one actual Owner-confirmed receipt passed. An
  initially incorrect credential was diagnosed and corrected within scope; it
  was not a product defect.

No credential, address, provider/account identity, host/network identity,
private response, message/header content or private path is retained here.

## Release-blocking finding

### QWSG-059-F001 — post-reboot recovered Guardian fails canonical-evidence convergence

The physical boot identity and Guardian process identity changed. Lingering,
the packaged unit, enablement, activation and the user default-target dependency
were present. A distinct Guardian invocation automatically started before
login, encountered one privacy-safely unclassified transient failure, and
systemd recovered it through the packaged `Restart=on-failure` policy.

The recovered process remained active and enabled. Its active checkpoint
matched the current effective configuration and systemd invocation and recorded
a completed cycle. Its canonical state also matched the current configuration
and was fresh, but remained `degraded` with partial completeness and
`inspect_evidence` / `verify_guardian_operation` recommendations after more
than one full configured five-minute interval. Readiness therefore remained
not-ready instead of converging to fresh running evidence.

This is not a missing linger setting, disabled/inactive service, wrong unit,
missing dependency, configuration-identity mismatch, stale evidence or early
readiness check. Correcting restart/exit/cycle evidence reconciliation requires
product source changes and new candidate bytes. Task 059 therefore stopped
without altering RC.5.

## Limitations and final state

- The exact terminal cause of the first pre-login invocation was not safely
  classified. The automatic restart and persistent convergence failure were
  independently established.
- Steps 10–12 are missing by design after the mandatory product-defect STOP and
  are never represented as PASS.
- At STOP, RC.5 remained installed with Guardian enabled and active, lingering
  enabled, canonical state degraded/partial, and overall readiness not-ready.
- No post-finding restart, configuration rewrite, notification resend,
  uninstall, reinstall, source/candidate modification, tag, Forgejo Release,
  upload, publication, deployment or announcement occurred.

RC.5 is not technically suitable to proceed directly toward QWSG 1.1.0 release
preparation. A separately authorized correction and replacement-candidate
acceptance are required. QWSG 1.1.0 is not released.
