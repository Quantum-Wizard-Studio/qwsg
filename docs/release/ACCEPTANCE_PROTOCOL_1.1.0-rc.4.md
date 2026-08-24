# QWSG 1.1.0-rc.4 External Clean-Host Acceptance Protocol

## Authority and execution boundary

This protocol applies only to one private RC.4 archive built reproducibly from
one exact clean integrated commit. It restarts at Checkpoint 01. The Owner
operates the disposable Ubuntu 24.04 amd64 VPS and receives one bounded
checkpoint at a time. Each phase requires its explicit Owner gate: A source
scaffolding integration; B private twin construction; C deterministic/package/
security proof; D private transfer; E external execution; and F evidence
integration/verdict. A PASS or READY does
not authorize a tag, Forgejo Release, upload, publication or announcement.

Use only product-visible README, INSTALL, installer, Smart Install, setup,
readiness and installed documentation. Do not supply hidden Linux knowledge to
manufacture a PASS. Never request or record credentials, recipient/provider
identity, private host/account data, private SSH material, tokens or headers.
No automatic sudo, package installation, lingering, reboot or arbitrary shell
remediation is permitted.

For every checkpoint report: purpose; exact Owner command/action; expected
evidence; PASS criteria; FAIL/finding criteria; continuation safety; and
retain/redact requirements. Stop on a mandatory failure or new product/security
defect outside the acceptance-record authority.

## Candidate and transfer gates

- **Gate A:** integrate only the reviewed RC.4 acceptance-source allowlist.
- **Gate B:** export the authorized clean commit twice into independent
  mode-0700 roots; derive `SOURCE_DATE_EPOCH` from its commit timestamp; embed
  the full commit; first prove ordinary checkout/export builds pass without
  `GOFLAGS`, with truthful defaults, exact explicit identity and no ambient
  Go VCS settings, preserving historical `QWSG-053-F001` unchanged.
- **Gate C:** prove binary, manifest, archive and sidecar byte identity;
  verify static Linux amd64 identity, safe layout, documentation, LICENSE and
  exclusions.
- **Gate D:** transfer exactly the archive and sidecar by Owner-approved strict
  SSH/SCP. Never weaken host-key verification. Verify two regular non-symlink
  destination files, size, digest and sidecar without extraction or execution.
- **Gate E:** begin the following external protocol only after explicit Owner
  authority. Separate state-changing actions where indicated.

## Finding policy

Use `RELEASE BLOCKER`, `SECURITY DEFECT`, `FUNCTIONAL DEFECT`,
`UX/DOCUMENTATION DEFECT`, or `COSMETIC / POST-RELEASE CANDIDATE`. Preserve all
evidence safely and stop where continuation could obscure the finding or exceed
authority. Never repair or replace candidate bytes during acceptance.

## Checkpoint 01 — Private candidate receipt

- **Purpose:** establish the RC.4 transfer identity before touching bytes.
- **Action:** list only the approved destination directory entries and file
  types using the exact bounded command supplied at Gate E.
- **Expected evidence:** exactly the RC.4 archive and `.sha256` sidecar, both regular
  files and neither a symlink.
- **PASS:** count, names and types match the Gate D record.
- **FAIL/finding:** missing, extra, linked or differently named content.
- **Safe continuation:** read-only; stop on FAIL.
- **Retain/redact:** retain names/types; redact host/account/path identity.

## Checkpoint 02 — Archive checksum

- **Purpose:** prove received-byte identity.
- **Action:** run `sha256sum -c qwsg-1.1.0-rc.4-linux-amd64.tar.gz.sha256` in
  the bounded candidate directory.
- **Expected evidence:** the exact archive reports `OK` and size/digest match Gate D.
- **PASS:** the expected evidence is observed exactly.

- **FAIL/finding:** any mismatch or sidecar ambiguity; stop without extraction.
- **Safe continuation:** read-only.
- **Retain/redact:** retain filename, size, SHA-256 and result only.

## Checkpoint 03 — Safe archive layout

- **Purpose:** reject unsafe or noncanonical members before extraction.
- **Action:** inspect the archive member/type listing with the protocol-provided
  fixed `tar` commands; do not extract yet.
- **Expected evidence:** one canonical root; relative normalized paths; no absolute,
  parent, duplicate, link or special members; expected modes only.
- **PASS:** the expected evidence is observed exactly.

- **FAIL/finding:** any unsafe, extra or ambiguous member; stop.
- **Safe continuation:** read-only.
- **Retain/redact:** retain normalized member/type/mode evidence.

## Checkpoint 04 — Internal manifest

- **Purpose:** prove every packaged file against `MANIFEST.sha256`.
- **Action:** extract only after Checkpoint 03, enter the single root and run
  the documented manifest verification command.
- **Expected evidence:** every entry verifies, with no unmanifested package file
  except the manifest itself as defined by packaging.
- **PASS:** the expected evidence is observed exactly.

- **FAIL/finding:** mismatch, omission or extraction anomaly; stop.
- **Safe continuation:** extraction affects only the disposable candidate directory.
- **Retain/redact:** retain manifest result and root identity.

## Checkpoint 05 — LICENSE, documentation and RC.4 identity

- **Purpose:** verify required legal/operator material and candidate identity.
- **Action:** inspect root `LICENSE`, `README.md`, `INSTALL.md`, RC.4 release
  notes and binary version/source metadata using their documented commands.
- **Expected evidence:** byte-correct LICENSE; usable root docs; RC.4 notes; binary
  version `1.1.0-rc.4`; exact Gate C source commit/build time.
- **PASS:** the expected evidence is observed exactly.

- **FAIL/finding:** absent, stale, conflicting or incorrect identity/content.
- **Safe continuation:** read-only; documentation defects are recorded truthfully.
- **Retain/redact:** retain hashes and public identities only.

## Checkpoint 06 — README and INSTALL operator journey

- **Purpose:** prove an ordinary supported-platform operator can determine the
  next safe action without developer coaching.
- **Action:** Owner follows only the packaged README/INSTALL sequence up to the
  Smart Install command.
- **Expected evidence:** prerequisites, privilege boundaries, paths and next steps
  are clear, bounded and internally consistent.
- **PASS:** the expected evidence is observed exactly.

- **FAIL/finding:** missing/unsafe/ambiguous guidance is a finding; stop if blocking.
- **Safe continuation:** no host mutation yet.
- **Retain/redact:** retain cited document sections, not private host details.

## Checkpoint 07 — Smart Install

- **Purpose:** assess the clean host before installation.
- **Action:** run exactly `./bin/qwsg install --check` without added environment
  overrides or developer coaching.
- **Expected evidence:** supported Ubuntu 24.04 amd64 facts are correctly classified;
  all mandatory requirements are satisfied or have safe product-visible action.
- **PASS:** the expected evidence is observed exactly.

- **FAIL/finding:** false classification, unexplained unknown, unsafe action or NOT READY.
- **Safe continuation:** read-only; do not install on blocking FAIL.
- **Retain/redact:** retain complete privacy-reviewed assessment output.

## Checkpoint 08 — F002/F003 regression proof

- **Purpose:** freshly retest the historical Task 049 corrections.
- **Action:** evaluate the Checkpoint 07 product output only.
- **Expected evidence:** `systemd.user_manager` uses validated runtime context and
  supplies cause-specific bounded verification, privilege, only proven-safe
  remediation and mandatory revalidation; `filesystem.local_semantics` uses
  bounded read-only evidence or precise manual verification without mutation.
- **PASS:** the expected evidence is observed exactly.

- **FAIL/finding:** false negative, guessed command, unconditional unexplained unknown,
  missing next action or host mutation. Historical records remain unchanged.
- **Safe continuation:** read-only; blocking failure stops.
- **Retain/redact:** retain classifications/guidance, redact host identifiers.

## Checkpoint 09 — Immutable installation

- **Purpose:** install only the verified release-owned artifacts.
- **Action:** after a separate Owner confirmation, follow the packaged install
  command exactly; no automatic sudo or package installation.
- **Expected evidence:** binary, user unit and installed docs match the manifest and
  installer reports a bounded handoff.
- **PASS:** the expected evidence is observed exactly.

- **FAIL/finding:** unexpected mutation, privilege behavior or identity mismatch; stop.
- **Safe continuation:** state-changing and separately confirmed.
- **Retain/redact:** retain installed paths/modes/hashes, not private identity.

## Checkpoint 10 — Guided setup interruption and resume

- **Purpose:** prove resumability and Community one-recipient configuration.
- **Action:** start `qwsg setup`, interrupt at the documented safe point, rerun,
  reject representative invalid input, then complete configuration without
  entering SMTP secrets in chat or argv.
- **Expected evidence:** safe interruption, preserved progress, deterministic
  validation and one-recipient Community configuration.
- **PASS:** the expected evidence is observed exactly.

- **FAIL/finding:** corruption, secret exposure, lost state or unsafe defaults.
- **Safe continuation:** configuration-only; stop before credentials on FAIL.
- **Retain/redact:** retain field status only; redact addresses/provider data.

## Checkpoint 11 — Protected credential-file workflow

- **Purpose:** establish the existing protected local secret boundary.
- **Action:** Owner enters the SMTP credential only on the VPS through the
  documented product workflow and verifies its mode/ownership without showing
  contents.
- **Expected evidence:** current-user protected mode-0600 file/reference, no secret
  in argv/output/history/evidence.
- **PASS:** the expected evidence is observed exactly.

- **FAIL/finding:** permissive mode, copied secret or acceptance-only mechanism; stop.
- **Safe continuation:** sensitive and Owner-only.
- **Retain/redact:** retain mode/type/result only.

## Checkpoint 12 — Notification preflight

- **Purpose:** validate configuration safely before sending.
- **Action:** run the product-documented notification preflight.
- **Expected evidence:** bounded success with protected credential resolution.
- **PASS:** the expected evidence is observed exactly.

- **FAIL/finding:** configuration/auth/TLS/diagnostic defect; do not infer receipt.
- **Safe continuation:** may contact configured provider; Owner confirmation required.
- **Retain/redact:** redact recipient, account, provider, headers and tokens.

## Checkpoint 13 — Real external SMTP receipt

- **Purpose:** prove end-to-end Community notification delivery.
- **Action:** run the documented controlled test and have the Owner independently
  confirm receipt at the intended external mailbox.
- **Expected evidence:** command succeeds and Owner confirms actual receipt.
- **PASS:** the expected evidence is observed exactly.

- **FAIL/finding:** TLS/auth success without receipt is FAIL; stop if unresolved.
- **Safe continuation:** sends one test message.
- **Retain/redact:** record confirmation/time class only; no message/header data.

## Checkpoint 14 — Guided Guardian activation

- **Purpose:** execute the exact formerly failing supported workflow.
- **Action:** run/resume `qwsg setup`; at `Activate QWSG Guardian now? [y/N]:`
  explicitly enter `y`. Do not manually invoke `systemctl` as a workaround.
- **Expected evidence:** setup reports successful fixed activation and preserves the
  validated systemd user-runtime boundary.
- **PASS:** the expected evidence is observed exactly.

- **FAIL/finding:** any activation-stage failure, generic/unactionable diagnostic or
  workaround need creates a finding and stops.
- **Safe continuation:** enables/starts only the fixed QWSG user service.
- **Retain/redact:** retain stage/result and safe diagnostic only.

## Checkpoint 15 — QWSG-051-F001 correction proof

- **Purpose:** independently decide whether RC.4 corrects the blocker.
- **Action:** run `qwsg readiness` exactly as product guidance directs.
- **Expected evidence:** unit installed, service enabled and active, user manager
  reachable, plus fresh integrity-checked canonical Guardian evidence.
- **PASS:** the expected evidence is observed exactly.

- **FAIL/finding:** missing/stale evidence or failed enabled/active state. Local tests,
  manual activation and process state alone never suffice.
- **Safe continuation:** read-only; F001 remains historical OPEN/BLOCKING on FAIL.
- **Retain/redact:** retain classifications and evidence identities safely.

## Checkpoint 16 — Independent final readiness at activation boundary

- **Purpose:** prove installation, environment, configuration and Guardian form
  one coherent ready state.
- **Action:** preserve the complete privacy-reviewed readiness result and its
  canonical evidence verification.
- **Expected evidence:** every mandatory readiness group is ready with fresh evidence.
- **PASS:** the expected evidence is observed exactly.

- **FAIL/finding:** any mandatory missing/unknown/invalid/stale state; stop.
- **Safe continuation:** read-only.
- **Retain/redact:** omit private paths/account/provider values.

## Checkpoint 17 — systemd process, invocation and resource behavior

- **Purpose:** prove actual execution rather than ActiveState alone.
- **Action:** run only the bounded product/documented user-service status,
  process and journal evidence commands supplied for this checkpoint.
- **Expected evidence:** expected executable/argv, cadence, restart behavior and
  configured resource limits; identity correlates with fresh Guardian evidence.
- **PASS:** the expected evidence is observed exactly.

- **FAIL/finding:** wrong process/invocation, unbounded resources or inconsistent state.
- **Safe continuation:** read-only; no arbitrary unit control.
- **Retain/redact:** retain QWSG unit/process facts, redact host identifiers.

## Checkpoint 18 — Lingering detection and guidance

- **Purpose:** verify product handling of boot persistence prerequisites.
- **Action:** inspect only the product-visible lingering assessment/guidance.
- **Expected evidence:** accurate detection, bounded privilege explanation and safe
  manual Owner action only when needed; QWSG does not enable lingering itself.
- **PASS:** the expected evidence is observed exactly.

- **FAIL/finding:** automatic mutation, guessed command or missing revalidation.
- **Safe continuation:** read-only unless Owner separately performs documented admin.
- **Retain/redact:** retain state/guidance only.

## Checkpoint 19 — Logout/session behavior

- **Purpose:** prove Guardian behavior across loss of the interactive session.
- **Action:** Owner logs out normally, waits the bounded documented interval,
  reconnects, then runs the approved readiness/evidence checks.
- **Expected evidence:** behavior matches lingering state and documented contract;
  fresh canonical evidence remains coherent.
- **PASS:** the expected evidence is observed exactly.

- **FAIL/finding:** unexplained stop, false readiness or session dependency.
- **Safe continuation:** session-changing; no server configuration change.
- **Retain/redact:** retain timing/state classes, not login/host identity.

## Checkpoint 20 — Physical VPS reboot

- **Purpose:** test actual boot continuity.
- **Action:** only after a separate explicit Owner confirmation, the Owner
  performs a physical VPS reboot and reconnects normally.
- **Expected evidence:** reboot completes and the approved account returns without
  acceptance automation or hidden remediation.
- **PASS:** the expected evidence is observed exactly.

- **FAIL/finding:** unsafe/unavailable reboot or ambiguous host identity; stop.
- **Safe continuation:** disruptive and explicitly gated.
- **Retain/redact:** record confirmation/time class, redact infrastructure data.

## Checkpoint 21 — Automatic post-reboot Guardian return

- **Purpose:** prove new invocation identity and fresh post-reboot evidence.
- **Action:** run readiness plus bounded service/process/evidence inspection.
- **Expected evidence:** enabled service returns automatically with a new process or
  invocation identity and fresh integrity-checked canonical evidence.
- **PASS:** the expected evidence is observed exactly.

- **FAIL/finding:** stale/pre-reboot evidence, manual start need or identity ambiguity.
- **Safe continuation:** read-only; no workaround.
- **Retain/redact:** retain comparative identities without host-private data.

## Checkpoint 22 — Post-reboot notification continuity

- **Purpose:** prove protected credentials and delivery survive reboot.
- **Action:** run the documented controlled notification test and obtain a
  second independent Owner receipt confirmation.
- **Expected evidence:** test succeeds and the intended message is actually received.
- **PASS:** the expected evidence is observed exactly.

- **FAIL/finding:** transport-only success, missing receipt or secret boundary change.
- **Safe continuation:** sends one test message.
- **Retain/redact:** same strict redaction as Checkpoints 11–13.

## Checkpoint 23 — Explicit Guardian restart

- **Purpose:** prove supported restart and recovery behavior.
- **Action:** after explicit Owner confirmation, use only the documented fixed
  QWSG Guardian restart action, then run readiness.
- **Expected evidence:** clean restart, new invocation identity and fresh canonical
  evidence without configuration loss.
- **PASS:** the expected evidence is observed exactly.

- **FAIL/finding:** wrong unit control, stale evidence or failed recovery.
- **Safe continuation:** state-changing but bounded to the QWSG user service.
- **Retain/redact:** retain transition/evidence identity safely.

## Checkpoint 24 — Safe uninstall with preserved user state

- **Purpose:** prove release-owned removal and user-data preservation.
- **Action:** after explicit Owner confirmation, run the packaged uninstall
  workflow and inspect only documented paths/modes.
- **Expected evidence:** release-owned binary/unit/docs removed; service safely
  handled; configuration, protected credential and state preserved unchanged.
- **PASS:** the expected evidence is observed exactly.

- **FAIL/finding:** user-data loss, unrelated mutation or incomplete unsafe removal.
- **Safe continuation:** destructive to installed artifacts only; snapshot first.
- **Retain/redact:** retain path classes/hashes, never credential contents.

## Checkpoint 25 — Same-candidate reinstall, resume and final readiness

- **Purpose:** prove preserved-state recovery without stale PASS.
- **Action:** after explicit Owner confirmation, reinstall the same verified
  RC.4 candidate, run `qwsg setup` to detect/resume preserved state, explicitly
  reactivate Guardian through product guidance, then run final readiness.
- **Expected evidence:** preserved config/credential/state are reused safely; setup
  resumes; activation succeeds; final enabled/active/fresh evidence is newly
  generated and valid.
- **PASS:** the expected evidence is observed exactly.

- **FAIL/finding:** stale evidence accepted, lost state, secret re-exposure, activation
  failure or any mandatory readiness gap.
- **Safe continuation:** final state-changing checkpoint; stop on any discrepancy.
- **Retain/redact:** retain preservation hashes/state classes and fresh evidence
  identity only.

## Final decision gate

`READY FOR QWSG 1.1.0 RELEASE` requires all 25 checkpoints PASS, externally
verified F001 correction, fresh F002/F003 regression PASS, two independently
confirmed SMTP receipts, post-reboot new/fresh identity evidence, successful
restart and uninstall/reinstall/resume/reactivation, no open blocker/security
defect, and all local/governance/preservation gates PASS.

Otherwise record `NOT READY FOR QWSG 1.1.0 RELEASE` with exact unresolved
findings or missing mandatory evidence. Gate F may integrate privacy-safe
records only after separate Owner authorization. Neither verdict authorizes
final release or publication.
