# Current Engineering Task 053: QWSG 1.1.0-rc.3 Clean-Host Acceptance and Release Readiness

## Task Metadata

- Task ID: `053`
- Task slug: `qwsg-1-1-0-rc-3-clean-host-acceptance-release-readiness`
- Status: `complete with disclosed limitations — NOT READY FOR QWSG 1.1.0 RELEASE`
- Date opened: `2026-08-21` UTC
- Human authority: Project Owner
- Owner or lead-developer communication language: English

## Title

QWSG 1.1.0-rc.3 Clean-Host Acceptance and Release Readiness


## Objective

Prepare one private, independently identifiable, reproducible QWSG
`1.1.0-rc.3` candidate from one exact clean integrated source commit; transfer
only its archive and checksum sidecar through a separately Owner-approved
private channel; and execute a fresh Owner-operated external clean-host
acceptance from Checkpoint 01 on a freshly reinstalled supported Ubuntu 24.04
amd64 VPS. Prove on the real host whether the Task 052 guided Guardian
activation correction resolves historical blocker `QWSG-051-F001`, revalidate
the Task 050 corrections for Task 049 F002/F003, and complete the Community
operator journey through real SMTP receipt, systemd/session/physical-reboot
continuity, uninstall, reinstall, preserved-state resume and final readiness.
End with exactly `READY FOR QWSG 1.1.0 RELEASE` or `NOT READY FOR QWSG 1.1.0
RELEASE`; neither verdict authorizes final release or publication.



## Scope

- Phase A — repository and acceptance-source readiness. Audit canonical idle,
  exact HEAD/origin ancestry, `VERSION=1.1.0-rc.3`, Task 052 integration,
  release plumbing, README/INSTALL/operator docs, RC.3 notes, packaging,
  systemd unit, acceptance precedents and preservation baselines. Prepare only
  the minimum RC.3-specific protocol, acceptance ledger and version-aware
  plumbing needed for commit-pure acceptance. If any product correction or
  substantial release-source work is required, record the finding and stop for
  a narrower prerequisite task. Never build from an uncommitted overlay.
- Gate A1: report the readiness audit, exact proposed source/scaffolding path
  set and candidate-source decision. Gate A2: require separate explicit Owner
  authorization for path-based staging, review, validation, commit and clean
  fast-forward push of only necessary acceptance-source changes. No candidate
  construction before the exact clean integrated commit is directly verified.
- Phase B — private candidate construction, separately Owner-authorized.
  Export the exact approved candidate-source commit into two independent
  private mode-0700 source/build roots. Derive `SOURCE_DATE_EPOCH` from that
  commit timestamp, pass the exact lowercase full 40-character commit, and
  build `qwsg-1.1.0-rc.3-linux-amd64.tar.gz` independently in both roots.
- Phase C — deterministic twin-build proof. Require byte-identical Linux-amd64
  static binaries, internal manifests, archives and checksum sidecars, exact
  version/full commit/build time, controlled metadata and no working-tree
  input. A failed or non-identical build stops before transfer.
- Phase D — package/security proof. Verify sidecars and internal manifests;
  canonical safe single-root archive layout; relative paths; no parent,
  absolute, link or special members; correct modes; packaged byte-correct
  LICENSE; root README.md and INSTALL.md; RC.3 notes and installed docs; static
  executable; and exclusion of credentials, secrets, private keys, Builder
  inputs, prompts, acceptance evidence, snapshots, backups, caches, `.git`,
  Owner-owned content, unrelated files and RC.1/RC.2 collision. Record exact
  source, epoch, UTC build time, filename, size, SHA-256, manifest/binary
  identities and reproducibility. Stop at a separate Owner transfer gate.
- Phase E — private transfer, separately Owner-authorized. Prefer standard
  strict-host-key SSH/SCP development VPS -> Owner-approved disposable VPS
  only after destination/account and host-key trust are explicitly approved.
  Transfer exactly the RC.3 archive and `.sha256` sidecar, no wildcard or
  directory. Never weaken host-key checking, alter SSH server configuration,
  install transfer software, expose publicly or record credentials/private
  host data. If direct transfer is unavailable, stop and request authority for
  the Owner-workstation fallback. Verify destination count, regular-file/no-
  symlink type, size, SHA-256 and sidecar without extracting or executing.
- Phase F — external acceptance, separately Owner-authorized, starts at
  Checkpoint 01. The Owner operates the freshly reinstalled VPS. Provide only
  one bounded checkpoint at a time, each with purpose, exact operator action,
  expected evidence, PASS criteria, FAIL/finding criteria, continuation
  safety, and retain/redact rules. Product-visible README, INSTALL, installer,
  Smart Install, setup, readiness and installed docs are the operator guidance;
  do not coach with unstated developer/Linux knowledge to manufacture PASS.
- The RC.3 protocol must independently cover: (01) private receipt and exact
  two-file provenance; (02) archive checksum; (03) safe layout; (04) internal
  manifest; (05) LICENSE/root docs/RC.3 identity; (06) README/INSTALL journey;
  (07) `./bin/qwsg install --check`; (08) F002/F003 guidance retest; (09)
  install; (10) setup interruption/resume and Community one-recipient config;
  (11) protected credential-file workflow; (12) notification preflight; (13)
  real external notification test and independent Owner receipt; (14) explicit
  guided Guardian activation; (15) F001 correction proof; (16) readiness with
  enabled/active/fresh canonical evidence; (17) actual systemd process,
  invocation, cadence, resource and restart evidence; (18) lingering
  detection/guidance; (19) logout/session behavior; (20) physical VPS reboot;
  (21) automatic return with new process/invocation identity and fresh
  post-reboot canonical evidence; (22) post-reboot notification continuity and
  Owner-confirmed receipt; (23) explicit Guardian restart; (24) safe uninstall
  with configuration/credential/state preservation; (25) same-candidate
  reinstall, preserved-state setup/resume/reactivation and final readiness.
- F001 may move to `EXTERNALLY VERIFIED CORRECTED` only when RC.3 interactive
  `qwsg setup`, after explicit `y`, completes its fixed activation on the
  supported clean host and independent `qwsg readiness` proves enabled,
  active and fresh integrity-checked canonical Guardian evidence. Local tests,
  process state alone, manual systemctl activation or an acceptance workaround
  are insufficient. Preserve the historical RC.2 finding unchanged.
- F002/F003 remain immutable Task 049 history. RC.3 must run Smart Install from
  scratch without developer coaching and prove correct user-manager
  classification with validated runtime context, cause-specific bounded
  verification/privilege/safe-remediation/revalidation guidance, plus bounded
  read-only filesystem evidence or precise manual verification without host
  mutation. Do not infer PASS from RC.2.
- SMTP uses only the product's protected current-user mode-0600 credential-file
  boundary. Credentials never enter chat, argv, Git, history or evidence.
  Redact recipient, account, provider/host where required, credential reference,
  headers, tokens and private host/account identity. TLS/authentication is not
  receipt: unconditional readiness requires independent Owner-confirmed real
  receipt before reboot and continuity receipt after reboot.
- Phase G — evidence integration and verdict, separately Owner-authorized.
  Update only canonical RC.3 acceptance/history records with privacy-safe
  source/candidate/transfer/checkpoint/F001/F002/F003/SMTP/Guardian/reboot/
  uninstall/reinstall evidence. Stage/commit/push only an exact reviewed
  allowlist after full validation. Never overwrite RC.1/RC.2 or prior findings.
- Use the established severities: `RELEASE BLOCKER`, `SECURITY DEFECT`,
  `FUNCTIONAL DEFECT`, `UX/DOCUMENTATION DEFECT`, and `COSMETIC / POST-RELEASE
  CANDIDATE`. Record new defects truthfully. A product/security/setup/Guardian/
  SMTP/installer/uninstaller defect outside narrow acceptance-record authority
  stops at the safe checkpoint, preserves evidence, yields NOT READY when
  terminal, and requires separately authorized correction work.
- Preserve RC.1 source/artifact/evidence; RC.2 source
  `6d3f79accd4d52b94c960eefa93e2f51fbc9a48c`, artifact SHA-256
  `73d045cbc5577d3e9921a44760ba316d2094cf13fafe82f873be9f3600547315`
  and evidence; Task 049 F002/F003; Task 051 F001; v1.0.0 identities/artifacts;
  LICENSE; and excluded Owner-owned content.



## Out of Scope

- No hidden product correction, rebuild/replacement under the same candidate
  identity, acceptance workaround, arbitrary remediation or continuation after
  a stop-worthy finding. Substantial source correction requires a new bounded
  correction task and a new candidate identity as appropriate.
- No automatic sudo, package installation, lingering, reboot, SSH trust
  weakening, SSH server change, privilege escalation, arbitrary shell,
  executable/argv/environment/systemd-unit control, general remediation,
  Pro/QWS scope expansion, readiness/freshness weakening or credential access.
- No credentials, private keys, recipient/provider/private-host identifiers or
  raw sensitive evidence in chat, Git, task history, commands or reports.
- No use of RC.1/RC.2 artifact-dependent evidence as RC.3 PASS evidence. Prior
  host-independent evidence may be referenced only for chronology and protocol
  rationale; all RC.3 artifact and product-dependent checkpoints rerun.
- No final `v1.1.0` tag, Forgejo Release, public upload, signing/publication
  claim, announcement or final release. Even READY stops at a separate Owner
  release gate. No Task 054.
- Do not read, copy, hash, stage, modify, move, delete, ignore, package or
  otherwise touch Owner-owned `docs/architecture/QWCS_MIGRATION_BLUEPRINT.md`;
  record only excluded stat metadata where required.



## Required Reading

- `ai/core/00_PROJECT_PHILOSOPHY.md`
- `ai/core/01_CONSTITUTION.md`
- `ai/core/03_AGENTS.md`
- `ai/core/08_JOB_TEMPLATE.md`
- `ai/core/11_ENGINEERING_LIFECYCLE.md`
- `ai/core/14_PROMPT_WORKFLOW.md`
- `ai/core/16_GIT_POLICY.md`
- `ai/config/engineering-project.conf`

## Starting State Verification

- Verify UTC date, ordinary user, exact repository root, `main`, canonical
  HTTPS `origin`, HEAD and `origin/main` exactly
  `3bf5a8e26e32ac1489dcdfecad5b086e4141cd91`, direct remote main when authority
  permits, ahead/behind `0/0`, empty index, clean tracked tree, exactly zero
  active prompts and canonical idle with completed/archived Task 052.
- Verify the only unrelated state is the excluded Owner-owned untracked path by
  metadata only. Stop on any additional or changed tracked/untracked content.
- Verify `VERSION=1.1.0-rc.3`, matching private notes, release validate-only
  plumbing, no RC.3 archive/sidecar/tag/release, and Task 052 product commit
  `6bb5b62957e54e0ac3377ce1b85593408c341873` is an ancestor of closure HEAD.
- Verify immutable RC.1/RC.2/Task 049/Task 051 records and exact known
  identities, including F001 still historical OPEN/BLOCKING; v1.0.0 dereferenced
  target; LICENSE hash; and no Owner content access.
- Audit Tasks 049–052, RC.1/RC.2 protocols and records, current release scripts,
  package allowlist, README, INSTALL, RC.3 notes, Smart Install/readiness,
  setup/notification/userservice/Guardian/systemd behavior, installer and
  uninstaller. Determine whether only RC.3 acceptance scaffolding is needed.
- Run pre-change build, focused/full/race tests, vet, formatting, release-check,
  shell/static-systemd, package/install/uninstall, security, Framework 21,
  Builder 38, lifecycle 28, diversion 36, job/test-task, Git whitespace and
  preservation gates. Host-local probes are read-only diagnostics, never
  substitute for external acceptance.



## Snapshot Requirements

Before any target change create a unique private mode-0700 `/tmp` snapshot
containing a readable tracked-HEAD archive, Task 052 archive/history, Task 053
Builder source/input, candidate target payloads and absence records, Git/mode/
ACL/tool identities, RC.1/RC.2/F001/F002/F003/v1.0.0/LICENSE preservation and
bounded restore instructions. Record Owner content only as excluded stat
metadata without reading/copying/hashing it. Before each later phase snapshot
the exact new local/external targets and prior state needed for bounded
rollback; never store credentials or private external identity in repository
evidence.



## Risk Assessment

- **False release readiness — critical:** READY requires every mandatory RC.3
  checkpoint, F001/F002/F003 proof, two real SMTP receipts, reboot identity/
  fresh evidence, uninstall/reinstall and no open blocker/security defect.
- **Candidate/source mismatch — critical:** exact clean commit exports,
  commit-derived epoch, full embedded identity, twin byte equality and
  collision guards; never build overlays or replace bytes in place.
- **Historical evidence corruption — critical:** version-specific new records,
  hash preservation and no edits to RC.1/RC.2/Task 049/051 evidence.
- **Credential/privacy exposure — critical:** Owner-only protected file entry,
  mandatory redaction and no chat/argv/repository secrets or provider details.
- **External host mutation — high:** one checkpoint at a time, explicit Owner
  gates for transfer/install/activation/lingering/reboot/uninstall/reinstall,
  no automatic sudo/packages/lingering and bounded continuation decisions.
- **F001 false correction claim — critical:** guided activation plus independent
  enabled/active/fresh evidence on RC.3 is mandatory; no manual workaround.
- **SMTP false positive — critical:** transport success is insufficient; Owner
  must independently confirm intended receipt before and after reboot.
- **Session/reboot ambiguity — high:** capture pre/post process and invocation
  identity with fresh canonical evidence; ActiveState alone is insufficient.
- **Scope creep — high:** acceptance records may change, product defects stop
  for new authority; no hidden fix, final release or Task 054.
- **Rollback/data loss — critical:** literal targets, snapshots and preserved
  per-user state; no broad reset/clean/delete or VPS teardown authority.



## Planned Work

1. Validate canonical idle, exact Git/protected baseline and pre-change gates;
   snapshot; audit RC.3 source/release/operator/acceptance readiness.
2. Prepare RC.3-specific numbered checkpoint protocol, empty evidence ledger
   and only necessary version-aware plumbing; report Gate A1. If product work
   is needed, stop for a prerequisite correction task.
3. Under Gate A2, integrate only the exact accepted scaffolding paths using
   explicit staging, full diff/security/preservation review, commit, dry-run,
   clean fast-forward push and direct remote verification.
4. Under Gate B/C/D authorization, export the exact clean integrated commit
   twice, derive the commit epoch, build privately and independently, prove
   binary/manifest/archive/sidecar byte identity, validate package/layout/
   docs/LICENSE/static identity/exclusions, and record privacy-safe provenance.
5. Stop for Gate E; after approval transfer exactly archive plus sidecar using
   strict standard SSH/SCP and perform bounded destination integrity checks.
6. Stop for Gate F; then conduct Checkpoints 01–25 one at a time from receipt
   through final readiness, preserving evidence and stopping on each mandatory
   failure or authority boundary.
7. At F001, require product-guided activation and independent readiness proof;
   at F002/F003 require fresh Smart Install proof; at SMTP require actual Owner
   receipt before and after physical reboot.
8. Classify findings, stop safely when required, and never repair product code
   under acceptance authority.
9. Under Gate G, integrate only privacy-safe RC.3 evidence/history, run full
   regression/security/governance/preservation/rollback gates, and issue exactly
   READY or NOT READY. Stop at the separate final-release gate; do not create
   Task 054.



## Rollback Plan

- Before each phase verify snapshot identities, later-edit absence and exact
  literal targets. Restore only Task 053-owned local files from verified
  snapshots and remove only paths with recorded pre-task absence. Never use
  broad reset, checkout, clean, unresolved variables/globs or touch Owner work,
  historical evidence, tags, releases, credentials or unrelated files.
- If a build fails, preserve bounded logs and source identity; discard/retry
  only exact private failed output roots under separate safe method rules. Never
  relabel or overwrite a candidate. Non-identical twin builds stop transfer.
- If transfer fails, preserve source bytes and integrity evidence; remove or
  replace only a proven partial destination file with Owner authority. Never
  weaken SSH checks or silently switch transport.
- On the external VPS, reverse only the exact checkpoint-owned QWSG action via
  documented commands and Owner confirmation. Preserve configuration,
  credential store and state during uninstall; never undo general host admin,
  disable lingering, delete the VPS or purge user data without separate
  authority.
- After rollback rerun applicable integrity, readiness, Git/lifecycle,
  security, preservation and external-state checks; report exact remaining
  state and safe restart point.



## Deliverables

- RC.3-specific canonical Owner-operated protocol with Checkpoints 01–25 and
  explicit phase/continuation/privilege/privacy gates.
- Canonical RC.3 acceptance ledger covering source/candidate/build/transfer,
  F001/F002/F003, SMTP receipts, Guardian/systemd/session/reboot, restart,
  uninstall/reinstall/resume, findings and final verdict.
- One private `qwsg-1.1.0-rc.3-linux-amd64.tar.gz` plus sidecar, independently
  built twice from one exact clean commit with complete byte-identity and
  package/security provenance.
- Privacy-safe private transfer record or explicit stopped fallback decision.
- Fresh complete external clean-host evidence from Checkpoint 01, including
  real Owner-confirmed SMTP receipt before and after reboot.
- Truthful F001 external result without changing historical RC.2 evidence, plus
  fresh F002/F003 regression results.
- Finding register, stop/continuation decisions, snapshot/rollback and complete
  local/governance/preservation evidence.
- Exact final `READY FOR QWSG 1.1.0 RELEASE` or `NOT READY FOR QWSG 1.1.0
  RELEASE`, with no release/publication authority implied.



## Verification

- Repository/lifecycle: exact root/user/date/branch/remote/HEAD/origin/direct
  remote/ahead-behind, zero active prompts at start, sole Task 053 after later
  installation, empty index, expected status, Owner exclusion and Task 052
  ancestry/history validity.
- Source/release: exact RC.3 VERSION/notes/docs/plumbing; no conflicting
  artifact/tag/release; required acceptance scaffolding is commit-pure; no
  uncommitted candidate input.
- Reproducibility: two independent mode-0700 exact-commit exports and caches;
  commit timestamp epoch; exact full commit; byte-identical binary,
  MANIFEST.sha256, archive and sidecar; valid identities; static ELF x86-64;
  deterministic metadata and safe canonical layout.
- Package/exclusion: byte-correct LICENSE, root README/INSTALL, RC.3 notes,
  expected binary/unit/install/uninstall/config/docs only; no links/special/
  unsafe paths, secrets, credentials, evidence, prompts, Builder input,
  snapshots, backups, caches, `.git`, Owner or unrelated/RC.1/RC.2 bytes.
- Protocol quality: each checkpoint has purpose/action/expected/PASS/FAIL/
  continuation/redaction, is bounded/restartable and exposes dangerous or
  privileged actions only behind the correct Owner gate.
- Smart Install/F002/F003: correct supported-host classification, validated
  runtime context, cause-specific actionable guidance, only proven safe
  commands, privilege/revalidation, bounded read-only filesystem evidence or
  exact manual verification, no mutation or coaching.
- Setup/F001: interruption/resume and config preservation; explicit guided `y`;
  same validated manager; successful fixed activation; independent enabled,
  active and fresh canonical evidence; stage/cause diagnostics on failure; no
  manual workaround, false success or readiness weakening.
- Notification: one recipient, protected mode-0600 credential-file workflow,
  preflight, controlled test and independent real receipt; no secret evidence;
  repeat continuity receipt after reboot.
- Guardian/systemd: actual process/invocation, cadence, bounded resources and
  restart; lingering detection/guidance; logout; physical reboot; new identity,
  automatic return and fresh post-reboot evidence; explicit restart.
- Lifecycle operations: safe uninstall of unchanged release-owned artifacts,
  preserved config/credential/state, same-candidate reinstall, setup resume,
  explicit reactivation and final fresh readiness without stale PASS.
- Regression/governance: build, focused/full/race, vet, formatting,
  shell/static-systemd, package/install/uninstall, release-check, secret/diff/
  whitespace, Framework 21, Builder 38, lifecycle 28, diversion 36, job/
  test-task, snapshot/rollback and all protected hashes pass at required gates.
- READY only if every mandatory checkpoint passes, F001 is externally verified
  corrected, F002/F003 regressions pass, both SMTP receipts are independently
  confirmed, reboot/new/fresh evidence and uninstall/reinstall/resume pass, no
  blocker/security defect remains, and local/governance gates pass. Otherwise
  state exact `NOT READY FOR QWSG 1.1.0 RELEASE` with unresolved findings; do
  not infer PASS from missing evidence.



## Documentation Updates

- Add `docs/release/ACCEPTANCE_PROTOCOL_1.1.0-rc.3.md` with the canonical
  numbered Owner-operated protocol and phase gates; do not edit RC.1/RC.2.
- Add `docs/release/ACCEPTANCE_1.1.0-rc.3.md` with candidate/build/transfer,
  checkpoint, F001/F002/F003, SMTP, Guardian/reboot, uninstall/reinstall,
  finding and verdict evidence; do not overwrite historical records.
- Update `scripts/test-release-plumbing.sh` only as necessary to validate RC.3
  protocol/ledger while retaining explicit RC.1/RC.2 historical checks.
- Update RC.3 notes or operator docs only for narrow factual acceptance/source
  scaffolding corrections found in Phase A; product/behavior corrections stop
  for separate authority.
- Update Task 053 history throughout and the concise engineering history only
  at verified delivery. Archive Task 053 only under separate lifecycle
  authority. Do not create Task 054.



## Completion Criteria

Task 053 completes only when the exact clean RC.3 candidate-source commit is
integrated and directly verified; two independent commit exports produce
byte-identical binary/manifest/archive/sidecar with valid static/package/docs/
LICENSE/security provenance; transfer is separately authorized and verified;
and the Owner executes the complete RC.3 clean-host protocol from Checkpoint 01
through final readiness. F001 must be externally proven corrected by successful
product-guided activation plus independent enabled/active/fresh evidence;
F002/F003 must pass fresh Smart Install regression; real SMTP receipt must be
independently confirmed before and after physical reboot; systemd/session/
reboot/new-identity/fresh-evidence/restart and safe uninstall/preserved-state/
reinstall/resume/reactivation must pass. All findings and missing evidence are
truthfully recorded, all local/security/governance/rollback/preservation gates
pass, and the result states exactly READY or NOT READY. Any unavailable
mandatory evidence, unsafe transfer, non-reproducibility, external/product/
security defect or unresolved blocker prevents READY and stops at the proper
Owner boundary. Completion never authorizes final tag, Forgejo Release,
upload, publication, announcement, final QWSG 1.1.0 release or Task 054.


## Owner Approval Requirements

Approved by Project Owner through the Engineering Task Builder on 2026-08-21 UTC.

The structured task definition has been explicitly approved for implementation. Further scope changes require explicit Project Owner approval.
