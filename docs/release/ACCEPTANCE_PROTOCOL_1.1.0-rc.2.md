# QWSG 1.1.0-rc.2 External Clean-Host Acceptance Protocol

## Authority and execution boundary

This is a restartable, Owner-operated protocol for a disposable clean Ubuntu
24.04 LTS amd64 VPS. Run only one numbered checkpoint at a time and return its
privacy-reviewed output for evaluation before continuing. The candidate may be
built or transferred only after the separately approved candidate-source
commit exists. QWSG does not receive remote access and no checkpoint is
external evidence until the Project Owner physically executes it.

The major phase gates are separate and non-transitive:

1. **A2 — source integration:** approve exact target staging, commit and push.
2. **B — candidate construction:** approve two private commit-export builds.
3. **C — private transfer:** approve the exact destination/account and path.
4. **D1 — external start:** approve Checkpoints 01–04, which are read-only.
5. **D2–D5:** separately approve installation/setup, SMTP, session/reboot, and
   uninstall/reinstall mutations when their preceding evidence passes.

For Gate C, prefer standard pre-existing `scp` directly from the development
VPS to the Owner-approved disposable VPS. Transfer only the archive and
sidecar, retain normal strict host-key verification, and never record
credentials, destination identifiers, account names, or host keys. Do not
change SSH server configuration, install transfer software, weaken host-key
checking, or expose the candidate publicly. If direct transfer cannot be
established safely, use the Owner-workstation two-hop fallback.

Never paste an SMTP credential, email address, hostname, public IP, provider
header, token, private configuration, raw QWSG state, or complete journal into
chat or evidence. Replace unnecessary identifiers with `[REDACTED]`. Preserve
the original private output only under Owner control.

Every checkpoint records: UTC time, candidate/source identity, evidence class
(`external-host`, `physical-reboot`, or `real-provider`), `PASS`/`FAIL`/`NOT
EXECUTED`, any finding ID, and the privacy-reviewed evidence excerpt.

## Finding and continuation policy

| Severity | Meaning | Continuation |
| --- | --- | --- |
| RELEASE BLOCKER | Mandatory gate fails, false READY, destructive behavior, or provenance ambiguity | Stop |
| SECURITY DEFECT | Secret, privilege, path, injection, or privacy boundary fails | Stop immediately |
| FUNCTIONAL DEFECT | Documented product behavior fails | Stop unless the checkpoint explicitly permits safe evidence collection |
| UX/DOCUMENTATION DEFECT | A normal operator cannot discover or safely complete the next action | Record before any coaching; continue only with Owner approval |
| COSMETIC / POST-RELEASE CANDIDATE | No material correctness or usability effect | Continue |

Product defects discovered during physical acceptance are recorded, not fixed
inside Task 051. Engineering knowledge must not silently fill a product or
documentation gap.

RC.1 and the Task 049 findings remain immutable historical evidence. RC.2 may
record `QWSG-049-F002` or `QWSG-049-F003` as externally `VERIFIED/CORRECTED`
only after all special retest criteria below pass on the real external host.

## Checkpoint 01 — Receive the private candidate

- **Purpose:** Establish that the payload arrived through the separately
  Owner-approved private channel and has the approved filename.
- **Action:** Place only `qwsg-1.1.0-rc.2-linux-amd64.tar.gz` and its `.sha256`
  sidecar in a new private directory, then run:

  ```sh
  umask 077
  mkdir qwsg-1.1.0-rc.2-acceptance
  mv qwsg-1.1.0-rc.2-linux-amd64.tar.gz qwsg-1.1.0-rc.2-linux-amd64.tar.gz.sha256 qwsg-1.1.0-rc.2-acceptance/
  cd qwsg-1.1.0-rc.2-acceptance
  ls -l
  ```

- **Expected evidence:** Exactly the archive and sidecar; no v1.0.0 payload.
- **PASS:** Both regular files exist in the private directory.
- **FAIL/finding:** Missing, extra, renamed, public, or ambiguous payload is a
  RELEASE BLOCKER.
- **Safe continuation:** Only after PASS.
- **Retain/redact:** Retain filenames, sizes, modes, and transfer provenance;
  redact account and host paths.

## Checkpoint 02 — Verify the sidecar and archive layout

- **Purpose:** Prove transport integrity and a safe single archive root.
- **Action:** From the private directory run:

  ```sh
  sha256sum -c qwsg-1.1.0-rc.2-linux-amd64.tar.gz.sha256
  tar -tzf qwsg-1.1.0-rc.2-linux-amd64.tar.gz
  ```

- **Expected evidence:** Sidecar `OK`; every member is below the single
  `qwsg-1.1.0-rc.2-linux-amd64/` root with no absolute or parent path.
- **PASS:** Both conditions hold.
- **FAIL/finding:** Checksum or unsafe/ambiguous layout is a RELEASE BLOCKER.
- **Safe continuation:** Only after PASS.
- **Retain/redact:** Retain sidecar result and member list; it contains no host
  evidence.

## Checkpoint 03 — Read product instructions before installation

- **Purpose:** Test the operator documentation entrypoint without developer
  coaching.
- **Action:** Run:

  ```sh
  tar -xzf qwsg-1.1.0-rc.2-linux-amd64.tar.gz
  cd qwsg-1.1.0-rc.2-linux-amd64
  sed -n '1,220p' README.md
  sed -n '1,260p' INSTALL.md
  sha256sum -c MANIFEST.sha256
  test -f LICENSE
  ```

- **Expected evidence:** README and INSTALL cross-reference each other, expose
  the supported platform and complete next journey, LICENSE is packaged, and
  the manifest passes.
- **PASS:** A new operator can identify verification, preflight, install,
  setup, readiness, activation, reboot, and uninstall actions from these files.
- **FAIL/finding:** Missing/mismatched files or undiscoverable action is a
  FUNCTIONAL or UX/DOCUMENTATION DEFECT.
- **Safe continuation:** Stop for manifest failure; a documentation finding
  requires Owner continuation approval before coaching.
- **Retain/redact:** Retain hashes and the exact deficient passage/finding, not
  unrelated terminal history.

## Checkpoint 04 — Run Smart Install before host mutation

- **Purpose:** Verify supported-host detection, classifications, remediation
  safety, exit behavior, and next action.
- **Action:** Run exactly:

  ```sh
  ./bin/qwsg version
  ./bin/qwsg install --check
  ./bin/qwsg install --check --format json
  ```

- **Expected evidence:** Version is `1.1.0-rc.2` with the approved full source
  commit; Ubuntu 24.04, amd64, systemd 255+, ordinary user, and user manager are
  evidence-backed; required/optional/unknown states and next action agree.
- **PASS:** No mutation occurs and every recommendation is proven and bounded.
  For `systemd.user_manager`, classification must not be a stripped-session
  false negative; every unsatisfied cause has a cause-specific explanation,
  bounded verification, privilege boundary, mandatory revalidation, and an
  exact command only when the cause and supported mapping prove it safe.
  Ambiguous states have no guessed command. For
  `filesystem.local_semantics`, QWSG uses bounded read-only evidence where
  possible; otherwise it gives a precise manual procedure and mandatory
  revalidation rather than an unconditional unexplained unknown. Neither probe
  mutates the host. Human and JSON guidance must agree.
- **FAIL/finding:** False support/readiness, guessed command, mutation, or wrong
  identity is a RELEASE BLOCKER; missing guidance is a UX defect.
- **Safe continuation:** Stop on required, incompatible, unknown, or security
  state. The Owner may perform only QWSG-displayed remediation, then rerun this
  checkpoint.
- **Retain/redact:** Retain human and redacted JSON classifications, exits, and
  recommendations; redact host/user identifiers.

## Checkpoint 05 — Install immutable artifacts

- **Purpose:** Verify the documented low-level installer and its handoff.
- **Action:** Only after Checkpoint 04 permits installation, run:

  ```sh
  sudo ./install.sh
  /usr/local/bin/qwsg version
  sha256sum README.md /usr/local/share/doc/qwsg/README.md
  sha256sum INSTALL.md /usr/local/share/doc/qwsg/INSTALL.md
  ```

- **Expected evidence:** Installer verifies the manifest, installs binary/unit/
  docs, starts no service, and prints documentation paths, `qwsg setup`, and
  `qwsg readiness`; packaged and installed documents match.
- **PASS:** All expected evidence holds.
- **FAIL/finding:** Mutation beyond owned artifacts, missing guidance, or
  identity mismatch is a RELEASE BLOCKER.
- **Safe continuation:** Only after PASS.
- **Retain/redact:** Retain installer output and hashes; redact sudo/user/host
  metadata.

## Checkpoint 06 — Guided setup, interruption, and resume

- **Purpose:** Test first-use guidance, validation, explicit decisions, and
  evidence-derived resumability.
- **Action:** As the intended ordinary user run `qwsg setup`; intentionally exit
  once at the first safe prompt, rerun `qwsg setup`, enter one visibly invalid
  non-secret value when prompted, correct it, and decline Guardian activation
  on this pass. Then run:

  ```sh
  qwsg setup --plan --format json
  qwsg config validate
  qwsg readiness
  ```

- **Expected evidence:** Valid prior values persist, no opaque progress marker
  is needed, invalid input is rejected safely, activation is not implicit, and
  the next action is correct.
- **PASS:** Resume is deterministic and readiness does not claim false READY.
- **FAIL/finding:** Lost valid state, blocking input, implicit activation, or
  false success is a FUNCTIONAL/RELEASE BLOCKER finding.
- **Safe continuation:** Continue only if configuration is valid and no
  security finding exists.
- **Retain/redact:** Retain classifications and redacted plan only; omit email,
  SMTP, paths, and configuration values.

## Checkpoint 07 — Configure the private SMTP credential

- **Purpose:** Exercise the existing protected-file credential boundary without
  echo, argv, log, chat, or repository disclosure.
- **Action:** Complete non-secret SMTP fields through `qwsg setup`. Then, in an
  interactive Bash terminal, run this block and type the credential only at the
  hidden prompt:

  ```bash
  umask 077
  secret_file=$(mktemp ./qwsg-smtp-secret.XXXXXX)
  trap 'rm -f -- "$secret_file"' EXIT HUP INT TERM
  IFS= read -r -s -p 'SMTP credential: ' smtp_secret
  printf '\n' >&2
  printf '%s' "$smtp_secret" > "$secret_file"
  unset smtp_secret
  qwsg notification credential set --from-file "$secret_file"
  rm -- "$secret_file"
  trap - EXIT HUP INT TERM
  qwsg notification preflight
  ```

- **Expected evidence:** No terminal echo; input file is private and removed;
  credential store passes canonical permission checks; preflight is truthful.
- **PASS:** Credential readiness is reported without revealing secret or
  account data.
- **FAIL/finding:** Any disclosure, unsafe mode/path, or retained input is a
  SECURITY DEFECT. Missing provider availability is `CONDITIONAL GATE NOT
  EXECUTED`, not product failure.
- **Safe continuation:** Stop immediately on security failure. Continue to real
  delivery only when preflight permits it.
- **Retain/redact:** Retain only classification, exit status, safe file-mode
  result, and input-file removal confirmation. Never retain the command's typed
  input, raw configuration, provider host, recipient, or credential.

## Checkpoint 08 — Prove real external email receipt

- **Purpose:** Verify end-to-end delivery, not merely TCP/TLS/SMTP acceptance.
- **Action:** Run `qwsg notification test`, then independently check the one
  configured administrator mailbox without displaying provider headers.
- **Expected evidence:** QWSG reports controlled test success and the exact test
  message is received by the intended recipient.
- **PASS:** Both send result and independent receipt are confirmed.
- **FAIL/finding:** Product send failure is FUNCTIONAL; credential leakage is
  SECURITY; missing receipt cannot PASS. Provider unavailability is a
  conditional gate not executed and prevents unconditional release readiness.
- **Safe continuation:** Stop for security failure; otherwise record the result
  before proceeding.
- **Retain/redact:** Retain UTC send/receipt times, QWSG-safe message identity if
  emitted, and `receipt confirmed`; redact addresses, headers, provider, and
  message body.

## Checkpoint 09 — Activate Guardian through guided setup

- **Purpose:** Verify explicit ordinary-user activation and bounded fresh
  evidence.
- **Action:** Run `qwsg setup`, accept activation only when QWSG explicitly asks,
  wait only for its bounded result, then run:

  ```sh
  qwsg readiness
  qwsg readiness --format json
  ```

- **Expected evidence:** Activation is explicit; unit state and fresh canonical
  Guardian evidence are distinct; notification and overall summaries are
  truthful.
- **PASS:** Guardian produces fresh evidence and overall readiness reflects the
  real notification result.
- **FAIL/finding:** ActiveState-only READY, unbounded wait, automatic sudo, or
  false summary is a RELEASE BLOCKER.
- **Safe continuation:** Only after a truthful result; stop on unsafe failure.
- **Retain/redact:** Retain redacted domains, evidence tokens, next actions, and
  timing; never retain raw Guardian state.

## Checkpoint 10 — Verify user-systemd facts and resources

- **Purpose:** Corroborate service installation, enablement, process identity,
  restart count, invocation, and bounded resources.
- **Action:** Run:

  ```sh
  systemctl --user show qwsg-guardian.service --no-pager --property=LoadState,UnitFileState,ActiveState,SubState,MainPID,NRestarts,InvocationID,MemoryCurrent,MemoryPeak,TasksCurrent,ExecMainStatus
  systemctl --user status qwsg-guardian.service --no-pager --lines=20
  ```

- **Expected evidence:** Loaded, enabled, active/running, valid MainPID and
  invocation, zero unexplained restarts, limits respected, installed QWSG
  process, plus Checkpoint 09 fresh evidence.
- **PASS:** Systemd facts and canonical evidence agree.
- **FAIL/finding:** Mismatch or resource/restart failure is FUNCTIONAL; service
  state alone never passes.
- **Safe continuation:** Stop on unexplained failure.
- **Retain/redact:** Retain states/counts/resources and a stable pseudonym for
  invocation; redact host, user, paths, PID if unnecessary, and journal data.

## Checkpoint 11 — Lingering guidance and logout boundary

- **Purpose:** Verify QWSG detects boot-before-login readiness and never invokes
  sudo itself.
- **Action:** Run `qwsg readiness`. If it reports lingering missing, perform
  only the exact documented/QWSG-recommended administrator command after
  reviewing the literal current username, rerun readiness, then log out of the
  SSH session. Reconnect and run:

  ```sh
  systemctl --user is-active qwsg-guardian.service
  qwsg readiness
  ```

- **Expected evidence:** Missing lingering is explained; remediation is
  operator-controlled; after reconnect the service and fresh evidence remain
  truthful.
- **PASS:** Guidance is discoverable, no automatic sudo occurs, and the logout
  boundary behaves as documented.
- **FAIL/finding:** Unsafe interpolation, hidden prerequisite, automatic
  privilege, or lost Guardian is SECURITY/FUNCTIONAL/UX as applicable.
- **Safe continuation:** Stop on unsafe or unknown guidance; otherwise continue
  after PASS.
- **Retain/redact:** Retain readiness classifications and command shape with the
  username replaced by `[USER]`.

## Checkpoint 12 — Physical reboot and post-reboot evidence

- **Purpose:** Prove unattended boot recovery; simulation is not accepted.
- **Action:** Record the redacted pre-reboot invocation identity and evidence
  time. The Owner then physically reboots the disposable VPS using its normal
  administration boundary, reconnects without manually starting QWSG, and
  runs:

  ```sh
  systemctl --user show qwsg-guardian.service --no-pager --property=ActiveState,SubState,MainPID,NRestarts,InvocationID,ExecMainStatus
  qwsg readiness
  ```

- **Expected evidence:** Automatic return, a new invocation/process identity,
  bounded post-boot Guardian cycle, and fresh post-reboot canonical evidence.
- **PASS:** Physical reboot is independently confirmed and every expected fact
  holds without manual Guardian start.
- **FAIL/finding:** Missing physical reboot, stale evidence, same invocation, or
  failed automatic return is a RELEASE BLOCKER.
- **Safe continuation:** Only after PASS.
- **Retain/redact:** Retain reboot confirmation, pseudonymous before/after
  invocation distinction, states, and freshness times.

## Checkpoint 13 — Post-reboot notification continuity

- **Purpose:** Prove protected configuration remains usable after reboot.
- **Action:** Run `qwsg notification preflight` followed by `qwsg notification
  test`, then independently confirm receipt as in Checkpoint 08.
- **Expected evidence:** Preflight passes and a new controlled message arrives.
- **PASS:** Both product result and independent receipt pass.
- **FAIL/finding:** Functional or security defect; provider outage is recorded
  separately but does not become a product PASS.
- **Safe continuation:** Stop on security failure; otherwise record result.
- **Retain/redact:** Same redaction boundary as Checkpoint 08.

## Checkpoint 14 — Explicit Guardian restart

- **Purpose:** Verify controlled ordinary-user restart and recovery.
- **Action:** Run:

  ```sh
  systemctl --user restart qwsg-guardian.service
  systemctl --user show qwsg-guardian.service --no-pager --property=ActiveState,SubState,MainPID,NRestarts,InvocationID,ExecMainStatus
  qwsg readiness
  ```

- **Expected evidence:** New invocation, active service, bounded fresh cycle,
  and no false READY before fresh evidence.
- **PASS:** Restart and evidence recovery both pass.
- **FAIL/finding:** FUNCTIONAL or RELEASE BLOCKER for false readiness.
- **Safe continuation:** Only after service/evidence reaches a truthful state.
- **Retain/redact:** Retain pseudonymous invocation change, states, and timing.

## Checkpoint 15 — Uninstall preservation

- **Purpose:** Verify narrow artifact removal and intentional user-data
  preservation.
- **Action:** From the retained verified archive directory run:

  ```sh
  systemctl --user disable --now qwsg-guardian.service
  sudo ./uninstall.sh
  test ! -e /usr/local/bin/qwsg
  test ! -e /usr/local/lib/systemd/user/qwsg-guardian.service
  test ! -e /usr/local/share/doc/qwsg/README.md
  ```

- **Expected evidence:** Release-owned binary/unit/docs are removed, modified
  artifacts are never silently deleted, and per-user configuration, credential
  store, and state remain preserved without their contents being displayed.
- **PASS:** Removal and preservation policy both hold with clear output.
- **FAIL/finding:** Unexpected deletion or residue is RELEASE BLOCKER or
  FUNCTIONAL according to impact.
- **Safe continuation:** Continue to mandatory reinstall only if preservation
  is safe.
- **Retain/redact:** Retain existence/mode classifications only; never capture
  user-data contents.

## Checkpoint 16 — Mandatory reinstall and resume

- **Purpose:** Prove preserved valid state is safely recognized and stale
  evidence does not create false READY.
- **Action:** Run:

  ```sh
  sudo ./install.sh
  qwsg setup --plan --format json
  qwsg readiness
  qwsg setup
  qwsg readiness
  ```

- **Expected evidence:** Installer succeeds from the same verified candidate;
  setup recognizes valid persisted configuration without unnecessary repeated
  questions; activation remains explicit; readiness requires new qualified
  evidence rather than trusting stale pre-uninstall state.
- **PASS:** Resume, explicit reactivation, fresh evidence, and notification
  readiness are truthful.
- **FAIL/finding:** Lost state, unsafe reuse, repeated valid questions, implicit
  activation, or stale READY is FUNCTIONAL/RELEASE BLOCKER.
- **Safe continuation:** This is the final physical product gate; stop and
  record the result.
- **Retain/redact:** Retain redacted plan/readiness classifications and timing,
  not configuration or state payloads.

## Final decision gate

`READY FOR QWSG 1.1.0 RELEASE` requires every checkpoint above to PASS, the
local reproducibility/security/governance gates to pass, a physical reboot and
fresh post-reboot cycle, uninstall/reinstall, no RELEASE BLOCKER or SECURITY
DEFECT, and actual external SMTP receipt before and after reboot. Local or
simulated results never substitute. Otherwise the verdict is `NOT READY FOR
QWSG 1.1.0 RELEASE` with exact findings. Neither verdict authorizes a tag,
public artifact, Forgejo Release, upload, announcement, or final publication.
