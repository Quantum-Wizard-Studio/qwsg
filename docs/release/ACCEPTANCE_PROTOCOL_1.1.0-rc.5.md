# QWSG 1.1.0-rc.5 External Clean-Host Acceptance Protocol

## Authority and execution boundary

This protocol applies only to one private RC.5 archive reproducibly constructed
from the exact clean Task 057 Gate A integration commit. Gates A–F are separate
Owner authorities: A scaffolding integration; B twin construction; C
deterministic/package/security proof and privacy-safe evidence integration; D
private transfer; E external execution; F final evidence integration/verdict.
Every external run starts at Checkpoint 01 on an Owner-confirmed freshly
reinstalled/reset disposable Ubuntu 24.04 amd64 VPS. The mutated RC.4 host is
not clean evidence without that reset.

The Owner operates the host one bounded checkpoint at a time using documented
product workflows. A mandatory product, packaging, provenance, installation,
Guardian, notification, reboot, uninstall, reinstall, privacy or security
failure stops acceptance. Never repair the host or product, manually substitute
`systemctl`, weaken SSH trust, or provide a hidden workaround to manufacture a
PASS. Source correction requires a new commit and RC identity.

Credentials, recipient/provider identity, private host/account identity,
private SSH material, tokens and raw private paths never enter chat, Git,
lifecycle history or canonical evidence. State-changing, credential,
notification, logout, reboot, restart, uninstall and reinstall actions require
explicit Owner confirmation at their checkpoint.

## Candidate and transfer gates

- **Gate A:** integrate only reviewed RC.5 protocol, empty ledger, Task 057
  lifecycle chronology and narrowly required release-plumbing assertions.
- **Gate B:** export the resulting exact full commit twice into independent
  private no-`.git` module roots; unset GOFLAGS; prove ordinary exported builds;
  derive SOURCE_DATE_EPOCH from the commit; build twins with the full commit.
- **Gate C:** prove byte-identical binary, manifest, archive and sidecar; exact
  provenance; static Linux amd64; safe layout, modes and deterministic metadata;
  documentation, LICENSE and exclusion integrity. Record only privacy-safe proof.
- **Gate D:** privately transfer exactly the verified archive and sidecar using
  strict host-key verification or a separately authorized Owner-workstation
  fallback; verify destination count, type, size, hashes and sidecar before use.
- **Gate E:** execute Checkpoints 01–26 only after explicit authority and clean
  VPS confirmation.
- **Gate F:** integrate privacy-safe evidence and issue exactly READY or NOT
  READY. Neither verdict authorizes tag, Forgejo Release, upload or publication.

## Finding and preservation policy

Use `RELEASE BLOCKER`, `SECURITY DEFECT`, `FUNCTIONAL DEFECT`,
`UX/DOCUMENTATION DEFECT`, or `COSMETIC / POST-RELEASE CANDIDATE`. Preserve
RC.1, RC.2, failed RC.3, failed RC.4, QWSG-055-F001, QWSG-053-F001,
QWSG-051-F001 and Task 049 F002/F003 unchanged. RC.5 evidence is additive.
QWSG-055-F001 remains historical OPEN/BLOCKING until Checkpoints 14–17 pass
completely with fresh independent real-systemd evidence.

## Checkpoint 01 — Private candidate receipt

- **Purpose:** establish the transferred RC.5 file set before touching bytes.
- **Action:** list only the approved receipt directory names and file types with the bounded Gate E command.
- **Expected evidence:** exactly the RC.5 archive and `.sha256` sidecar, both regular non-symlink files.
- **PASS:** count, names and types match the Gate D record.
- **FAIL/finding:** missing, extra, linked or differently named content; stop.
- **Safe continuation:** read-only.
- **Retain/redact:** retain names/types only; redact host, account and path identity.

## Checkpoint 02 — Archive checksum

- **Purpose:** prove received-byte identity.
- **Action:** run `sha256sum -c` against the transferred RC.5 sidecar in the bounded receipt directory.
- **Expected evidence:** the exact archive reports `OK`; size and digest match Gate D.
- **PASS:** sidecar verification and recorded identities match exactly.
- **FAIL/finding:** any mismatch or ambiguity; stop without extraction.
- **Safe continuation:** read-only.
- **Retain/redact:** retain filename, size, SHA-256 and result only.

## Checkpoint 03 — Safe archive layout

- **Purpose:** reject unsafe or noncanonical members before extraction.
- **Action:** inspect fixed archive member/type/mode listings without extracting.
- **Expected evidence:** one canonical root; normalized relative paths; no parent, absolute, duplicate, link or special members; expected modes only.
- **PASS:** every layout and type rule passes.
- **FAIL/finding:** any unsafe, extra or ambiguous member; stop.
- **Safe continuation:** read-only.
- **Retain/redact:** retain normalized members, types and modes.

## Checkpoint 04 — Internal manifest

- **Purpose:** prove every packaged file after safe extraction.
- **Action:** extract only after Checkpoint 03 and run the packaged manifest verification from the canonical root.
- **Expected evidence:** every entry passes and no unmanifested payload exists outside the defined manifest exception.
- **PASS:** complete manifest verification succeeds.
- **FAIL/finding:** mismatch, omission or extraction anomaly; stop.
- **Safe continuation:** mutation is bounded to the disposable candidate directory.
- **Retain/redact:** retain manifest identity and aggregate result.

## Checkpoint 05 — LICENSE, documentation and RC.5 identity

- **Purpose:** verify legal/operator material and exact candidate provenance.
- **Action:** inspect packaged LICENSE, README, INSTALL, RC.5 notes and binary version metadata using documented commands.
- **Expected evidence:** preserved LICENSE; correct docs; version `1.1.0-rc.5`; exact Gate C commit and controlled UTC time.
- **PASS:** content identities and embedded provenance match Gate C.
- **FAIL/finding:** absent, stale, conflicting or incorrect material; stop if mandatory.
- **Safe continuation:** read-only.
- **Retain/redact:** retain public identities and hashes only.

## Checkpoint 06 — README and INSTALL operator journey

- **Purpose:** prove an ordinary operator can identify the safe supported path.
- **Action:** Owner follows only packaged README/INSTALL through the Smart Install command.
- **Expected evidence:** prerequisites, privilege boundaries, paths and next steps are clear and consistent.
- **PASS:** no developer coaching or hidden step is needed.
- **FAIL/finding:** unsafe, missing or ambiguous guidance; stop if blocking.
- **Safe continuation:** no host mutation.
- **Retain/redact:** retain cited public sections only.

## Checkpoint 07 — Clean-host Smart Install

- **Purpose:** assess the genuinely clean host before installation.
- **Action:** run exactly `./bin/qwsg install --check` without environment overrides or developer coaching.
- **Expected evidence:** Ubuntu 24.04 amd64 facts are correctly classified and mandatory requirements are satisfied or have safe product-visible action.
- **PASS:** assessment is coherent and installation-ready.
- **FAIL/finding:** false classification, unsafe action or blocking NOT READY; stop.
- **Safe continuation:** read-only.
- **Retain/redact:** retain complete privacy-reviewed assessment output.

## Checkpoint 08 — Task 049 F002/F003 regression proof

- **Purpose:** freshly retest user-manager guidance and filesystem local semantics.
- **Action:** evaluate only the uncoached Checkpoint 07 product output and bounded evidence.
- **Expected evidence:** actionable cause-specific user-manager guidance and accurate non-mutating local-filesystem classification/manual verification.
- **PASS:** both affected boundaries pass fresh RC.5 evidence.
- **FAIL/finding:** guessed remediation, unexplained unknown, missing revalidation or host mutation; stop if blocking.
- **Safe continuation:** read-only.
- **Retain/redact:** retain classifications and guidance; historical records stay unchanged.

## Checkpoint 09 — Immutable installation

- **Purpose:** install only verified release-owned artifacts.
- **Action:** after explicit Owner confirmation, follow the packaged install workflow exactly.
- **Expected evidence:** binary, unit and docs match the package and installer reports the documented handoff.
- **PASS:** exact owned artifacts install with expected modes and no unrelated mutation.
- **FAIL/finding:** privilege, identity, mode or scope discrepancy; stop.
- **Safe continuation:** state-changing and separately confirmed.
- **Retain/redact:** retain public path classes, modes and hashes only.

## Checkpoint 10 — Guided setup interruption, resume and configuration

- **Purpose:** prove resumability and safe configuration before activation.
- **Action:** interrupt at the documented safe point, resume, reject bounded invalid input, and complete configuration without exposing secrets; defer activation until Checkpoint 14.
- **Expected evidence:** safe resume, deterministic validation, configuration written privately, and no premature Guardian activation.
- **PASS:** configuration is valid and activation boundary remains pending.
- **FAIL/finding:** corruption, lost progress, unsafe defaults or secret exposure; stop.
- **Safe continuation:** configuration-only.
- **Retain/redact:** retain field states; redact addresses/provider data and paths.

## Checkpoint 11 — Protected credential workflow

- **Purpose:** establish the existing protected local secret boundary.
- **Action:** Owner enters credentials only through the documented on-host workflow and verifies type/mode/ownership without showing contents.
- **Expected evidence:** current-user-owned regular mode-0600 credential boundary with no secret in argv/output/history/evidence.
- **PASS:** boundary and reference behavior match documentation.
- **FAIL/finding:** permissive mode, symlink, copied secret or acceptance-only mechanism; stop.
- **Safe continuation:** sensitive and Owner-only.
- **Retain/redact:** retain type/mode/ownership/result only.

## Checkpoint 12 — Notification preflight

- **Purpose:** validate protected notification configuration safely.
- **Action:** run the product-documented preflight after explicit Owner confirmation.
- **Expected evidence:** bounded success with protected credential resolution and safe diagnostics.
- **PASS:** configuration, TLS/auth boundary and diagnostics pass.
- **FAIL/finding:** configuration, credential, TLS or diagnostic defect; stop.
- **Safe continuation:** may contact the configured provider.
- **Retain/redact:** redact recipient, account, provider, headers and tokens.

## Checkpoint 13 — First actual notification receipt

- **Purpose:** prove end-to-end Community notification delivery before reboot.
- **Action:** run one documented controlled test and obtain independent Owner receipt confirmation.
- **Expected evidence:** command succeeds and the intended external mailbox receives the message.
- **PASS:** both product result and Owner receipt confirmation exist.
- **FAIL/finding:** transport-only success or missing receipt; stop.
- **Safe continuation:** sends one controlled message.
- **Retain/redact:** record confirmation/time class only; no message/header data.

## Checkpoint 14 — Guided Guardian activation

- **Purpose:** execute the exact supported activation workflow corrected by Task 056.
- **Action:** resume documented `qwsg setup` and answer `y` at its Guardian activation prompt; do not manually invoke `systemctl`.
- **Expected evidence:** QWSG prepares state and completes its guided activation through the validated user-manager boundary.
- **PASS:** setup reports successful guided activation without workaround.
- **FAIL/finding:** any state-preparation, runtime, manager, reload, start or fresh-evidence failure; stop.
- **Safe continuation:** enables/starts only the QWSG user service.
- **Retain/redact:** retain privacy-safe activation stage/result only.

## Checkpoint 15 — Task 056 state-directory compatibility proof

- **Purpose:** independently prove the QWSG-055-F001 architectural correction under real systemd.
- **Action:** use bounded non-mutating type/ownership/mode/component checks plus privacy-safe user-unit state and journal classification after guided activation.
- **Expected evidence:** canonical state root is a real non-symlink directory; every required component is safe; intended ordinary-user ownership and mode 0700; configuration and state are distinct; no compatibility symlink or migration message; no manual repair.
- **PASS:** every state-path and systemd criterion passes exactly.
- **FAIL/finding:** symlink, unsafe component, wrong type/owner/mode, root overlap, compatibility migration, ambiguous evidence or repair need; QWSG-055-F001 remains OPEN/BLOCKING and acceptance stops.
- **Safe continuation:** read-only after Checkpoint 14.
- **Retain/redact:** retain classifications, modes and public unit facts; omit raw private paths and host identity.

## Checkpoint 16 — QWSG-055-F001 external correction proof

- **Purpose:** decide the historical blocker using the complete corrected boundary.
- **Action:** combine Checkpoints 14–15 with independent enabled, active and fresh canonical Guardian evidence; make no manual service change.
- **Expected evidence:** guided activation succeeded; service remains enabled and active; state contract is safe; fresh integrity-checked Guardian evidence exists; filesystem.local_semantics remains satisfied.
- **PASS:** all criteria pass, permitting only the additive status `EXTERNALLY VERIFIED CORRECTED IN RC.5` while historical OPEN/BLOCKING evidence remains immutable.
- **FAIL/finding:** any missing, stale, unsafe or inferred criterion; stop.
- **Safe continuation:** read-only evaluation.
- **Retain/redact:** retain finding decision and bounded evidence identities only.

## Checkpoint 17 — QWSG-051-F001 and complete readiness proof

- **Purpose:** freshly retest guided activation context and coherent operational readiness.
- **Action:** run exactly `qwsg readiness` and preserve its privacy-reviewed canonical evidence verification.
- **Expected evidence:** installation/environment/configuration ready; user manager valid; unit installed/enabled/active; fresh Guardian evidence; filesystem.local_semantics satisfied.
- **PASS:** every mandatory readiness group passes and QWSG-051-F001 receives additive external RC.5 correction evidence.
- **FAIL/finding:** missing/unknown/invalid/stale state or manual activation dependency; stop.
- **Safe continuation:** read-only.
- **Retain/redact:** retain classifications and evidence identities, not private paths.

## Checkpoint 18 — Guardian process, cadence and resource behavior

- **Purpose:** prove actual bounded execution rather than ActiveState alone.
- **Action:** run only approved product/documented status, process and bounded journal/resource checks.
- **Expected evidence:** expected executable/argv, cadence, restart policy and resource limits correlate with fresh evidence.
- **PASS:** process and runtime contract match exactly.
- **FAIL/finding:** wrong invocation, unbounded behavior or inconsistent state; stop.
- **Safe continuation:** read-only.
- **Retain/redact:** retain unit/process facts; redact host identifiers.

## Checkpoint 19 — Lingering detection and guidance

- **Purpose:** verify boot-persistence prerequisite handling.
- **Action:** inspect only product-visible lingering assessment and guidance.
- **Expected evidence:** accurate state, bounded privilege explanation and safe Owner action only when needed; QWSG does not enable lingering itself.
- **PASS:** detection, guidance and revalidation are correct.
- **FAIL/finding:** automatic mutation, guessed action or missing revalidation; stop if blocking.
- **Safe continuation:** read-only unless a separately confirmed documented admin action is needed.
- **Retain/redact:** retain state/guidance only.

## Checkpoint 20 — Logout and session behavior

- **Purpose:** prove Guardian behavior across loss of the interactive session.
- **Action:** Owner logs out, waits the bounded documented interval, reconnects, and runs approved readiness/evidence checks.
- **Expected evidence:** behavior matches lingering state and fresh canonical evidence remains coherent.
- **PASS:** no unexplained session dependency or evidence gap.
- **FAIL/finding:** unexplained stop, false readiness or stale evidence; stop.
- **Safe continuation:** session-changing without host configuration mutation.
- **Retain/redact:** retain timing/state classes only.

## Checkpoint 21 — Physical VPS reboot

- **Purpose:** test actual boot continuity.
- **Action:** after explicit Owner confirmation, the Owner physically reboots the disposable VPS and reconnects normally.
- **Expected evidence:** reboot completes and the approved account returns without acceptance automation or hidden remediation.
- **PASS:** clean reboot/reconnect boundary is confirmed.
- **FAIL/finding:** unsafe/unavailable reboot or ambiguous host continuity; stop.
- **Safe continuation:** disruptive and separately confirmed.
- **Retain/redact:** retain confirmation/time class; redact infrastructure identity.

## Checkpoint 22 — Automatic post-reboot Guardian return

- **Purpose:** prove automatic return with new invocation identity and fresh evidence.
- **Action:** run readiness plus bounded service/process/evidence inspection without manual start.
- **Expected evidence:** enabled service returns automatically with new identity and fresh integrity-checked canonical Guardian evidence; state remains a safe real directory.
- **PASS:** all automatic-return and freshness criteria pass.
- **FAIL/finding:** stale evidence, manual start need, state regression or ambiguity; stop.
- **Safe continuation:** read-only.
- **Retain/redact:** retain comparative identities and state classes safely.

## Checkpoint 23 — Post-reboot notification receipt

- **Purpose:** prove protected credentials and delivery survive reboot.
- **Action:** run one documented controlled notification test and obtain a second independent Owner receipt confirmation.
- **Expected evidence:** command succeeds and the intended message is actually received.
- **PASS:** product result and second Owner receipt confirmation both exist.
- **FAIL/finding:** transport-only success, missing receipt or credential-boundary change; stop.
- **Safe continuation:** sends one controlled message.
- **Retain/redact:** apply the same strict redaction as Checkpoints 11–13.

## Checkpoint 24 — Explicit Guardian restart

- **Purpose:** prove supported restart and recovery behavior.
- **Action:** after explicit Owner confirmation, use only the documented fixed QWSG restart action and then run readiness.
- **Expected evidence:** clean restart, new invocation identity, safe real state root and fresh canonical evidence without configuration loss.
- **PASS:** all restart/recovery criteria pass.
- **FAIL/finding:** wrong control path, stale evidence, state regression or failed recovery; stop.
- **Safe continuation:** state-changing but bounded to the QWSG user service.
- **Retain/redact:** retain transition/evidence identity safely.

## Checkpoint 25 — Safe uninstall with preserved user state

- **Purpose:** prove release-owned removal and user-data preservation.
- **Action:** after explicit Owner confirmation and bounded snapshot, run the packaged uninstall workflow and inspect only documented paths/modes.
- **Expected evidence:** release-owned binary/unit/docs are removed while configuration, protected credentials and state remain unchanged and safe.
- **PASS:** removal and preservation hashes/classes match.
- **FAIL/finding:** user-data loss, unrelated mutation or unsafe incomplete removal; stop.
- **Safe continuation:** destructive only to verified release-owned artifacts.
- **Retain/redact:** retain path classes/hashes; never credential contents.

## Checkpoint 26 — Same-candidate reinstall, resume and final readiness

- **Purpose:** prove preserved-state recovery without stale PASS.
- **Action:** after explicit Owner confirmation, reinstall the same verified RC.5 candidate, resume setup, reactivate through product guidance and run final readiness.
- **Expected evidence:** preserved config/credential/state are reused safely; real state directory remains non-symlink mode 0700; activation succeeds; new enabled/active/fresh evidence is valid.
- **PASS:** every final readiness, preservation and freshness criterion passes.
- **FAIL/finding:** stale evidence accepted, lost state, secret re-exposure, compatibility symlink, activation failure or readiness gap; stop.
- **Safe continuation:** final state-changing checkpoint.
- **Retain/redact:** retain preservation hashes/state classes and fresh evidence identity only.

## Final decision gate

`READY FOR QWSG 1.1.0 RELEASE` requires Gates A–F and all 26 checkpoints PASS,
fresh external QWSG-055-F001 and QWSG-051-F001 correction proof, fresh Task 049
F002/F003 evidence, exported-source QWSG-053-F001 proof, two independently
confirmed notification receipts, reboot/restart/uninstall/reinstall success,
and no open blocker or security defect. Otherwise record `NOT READY FOR QWSG
1.1.0 RELEASE` with exact unresolved evidence. Neither verdict authorizes tag,
Forgejo Release, upload, publication or announcement.

## Mandatory post-acceptance distribution follow-up

A separately authorized release/distribution phase must provide stable official
archive and SHA-256 sidecar URLs from `git.quantumwizard.hu`/Forgejo, command-line
checksum verification, `wget` and `curl` examples, future Smart Installer
compatibility, and normal installation without workstation-mediated copying.
Task 057 Gate A does not implement or publish that mechanism.
