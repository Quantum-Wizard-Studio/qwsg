# Task History 064: Approved Task

## Task metadata

- Task ID: `064`
- Task slug: `native-update-rollback-system`
- Status: `active`
- Date generated: `2026-08-26` UTC
- Human authority: Project Owner
- Preferred owner communication language: English
- Related prompt: `ai/prompts/064_CURRENT_TASK.md`

## Lifecycle state

The Engineering Task Builder generated and transactionally installed this matching prompt/history pair from validated structured owner input. Explicit approval was recorded and implementation began under Framework 2.0 Standard Execution Authority.

## Starting state

- Canonical root, ordinary user, `main`, Framework 2.0.0 and Task 064 lifecycle validated.
- Starting `HEAD == origin/main == bdcbe1827873f5422f51fe8d488679e1b06866d2`; `VERSION=1.1.0`; only the excluded Owner blueprint is unrelated/untracked.
- Pre-change snapshot: `/tmp/qwsg-task064-execution.HwBIV2`, mode `0700`; manifest SHA-256 `005e27416b4301bc1cd2c1acdc81ecd728c3158543d9bc4a2970f895bc3975de`; tracked-HEAD archive SHA-256 `3e47c3a46013f29c901f4f96a9d4a9310c8d9f9f26a286bddfc844b6055739dc`.

## Architecture and implementation

- Selected a two-phase design: anonymous canonical discovery and private
  unprivileged acquisition/verification precede a narrow fixed-destination
  privileged package transaction. Remote archive paths never select installed
  destinations.
- Added strict SemVer parsing and ordering with explicit newer/equal/older,
  unsupported-major and invalid classifications. Final releases order after
  their prereleases; malformed and ambiguous identities fail closed.
- Added canonical Forgejo Release discovery with exact repository/HTTPS asset
  origin, immutable version/tag filenames, draft/prerelease consistency,
  bounded metadata, download timeout/redirect/size controls and duplicate/
  missing asset refusal.
- Added private acquisition and local-candidate staging with regular-file,
  mode, exact filename/size and sidecar SHA-256 verification.
- Added safe archive extraction: unique single root; bounded count/sizes;
  directories and regular files only; no absolute, parent, backslash, symlink
  or duplicate members; internal manifest verification; required files; and
  `RELEASE.json` version/full-commit/build/platform provenance.
- Added fixed package-destination transactions, private integrity-bound prior
  artifact backups, atomic replacement, automatic restoration on eligible
  mutation failure, explicit rollback and tamper/incomplete metadata refusal.
  Configuration, credential and persistent Guardian/operator state paths are
  not package destinations or backup payloads.
- Added user orchestration that captures Guardian enabled/active intent, stops
  only at mutation, invokes the narrow helper through `sudo`, reloads and
  restores intent, validates installed version/configuration, records private
  local rollback metadata, and attempts package/service restoration after a
  post-mutation failure.
- The privileged helper re-copies and completely re-verifies the archive inside
  private root-owned staging before fixed-path replacement, closing the
  unprivileged-stage time-of-check/time-of-use boundary.
- Added explicit deterministic migration planning. The supported
  `1.1.0 -> 1.2.0-rc.1` path is a validated no-mutation compatibility decision
  for Configuration Source 1.0, Guardian Checkpoint 1.0, Scheduler State 1.0
  and Current Operator State 1.0–1.2; unknown paths fail closed.
- Operator workflow is `qwsg update check`, `qwsg update`, `qwsg update status`
  and `qwsg update rollback`. Because released 1.1.0 predates the command, its
  single bootstrap transition runs the fully verified newer archive binary;
  later supported updates use the installed binary.

## Focused verification

- Updater and CLI focused Go tests PASS: strict release ordering and malformed
  identities; canonical selection and hostile origin refusal; bounded private
  acquisition and size/checksum failure; migration compatibility/refusal;
  transactional preservation; explicit rollback; tampered backup refusal;
  injected post-mutation automatic restoration; CLI/help parsing; and
  privileged rollback-path bounds.
- A real development `1.2.0-rc.1` archive built with controlled placeholder
  provenance passed sidecar and the updater's archive, manifest, required-file
  and provenance verifier. Initial sidecar command used the wrong working
  directory; classified `TEST OR ACCEPTANCE DEFECT`, corrected, and retest
  PASS. No candidate was frozen or published.
- A listener-based discovery fixture failed because the managed sandbox denies
  local sockets. Classified `TEST OR ACCEPTANCE DEFECT`; replaced by direct
  deterministic metadata selection tests without weakening HTTP client bounds.
- Default Go cache was read-only. Classified `ENVIRONMENTAL ISSUE`; all Go work
  uses the configured private `/tmp` caches.

## Snapshot

- Pre-change snapshot: `/tmp/qwsg-task064-execution.HwBIV2`, mode `0700`;
  manifest SHA-256
  `005e27416b4301bc1cd2c1acdc81ecd728c3158543d9bc4a2970f895bc3975de`.
- External-acceptance boundary snapshot:
  `/tmp/qwsg-task064-acceptance-boundary.hTS8oF`, mode `0700`; manifest
  SHA-256
  `561e3fcc26834a583cef0a1cf43cb3c486bac1b70f25601c75e7d61cfef17e10`.

## Work performed

- Advanced the development identity to `1.2.0-rc.1` and implemented the
  complete discovery, acquisition, package trust, migration, transactional
  replacement, automatic recovery, explicit rollback and operator CLI design
  described above.
- Corrected an implementation defect found during integration: update planning
  initially compared the candidate with the executing archive binary rather
  than installed `/usr/local/bin/qwsg`. Added regression coverage and changed
  the comparison to the installed identity.
- Corrected rollback retention so a completed explicit rollback removes its
  consumed private backup and a later successful update discards only the
  validated superseded backup. Added tamper/refusal coverage.
- Integrated implementation commits
  `8a6ddefbc5fbf6b8ba0300ae8e3f4545b6ea7d64`,
  `ccb9aef4eea2d2cc6ba0d422353f94e74ba3a124`, and
  `b632acd8ab154e7f174b4dfb0f93f281ea261ffc` by clean fast-forward push.

## Deterministic candidate

- Two isolated canonical builds from source
  `b632acd8ab154e7f174b4dfb0f93f281ea261ffc` produced byte-identical archive
  and sidecar outputs.
- Frozen private archive `qwsg-1.2.0-rc.1-linux-amd64.tar.gz`: size `3481005`,
  SHA-256
  `975ada86f82d9b296aa00890d14f26a8f4ba44067f12561802f37106eb2664ac`.
- Sidecar SHA-256:
  `fd1ced570713b5127af2a426fff7da2aa41ce6db9cdfc05d6a8414ad39f27387`;
  manifest SHA-256
  `218dff4481619c3de7993f0601e62663b638a9f3a40434b64921ba8e86838b8e`;
  binary SHA-256
  `2a2ead6a3ec4e566574c03e079fa0e17f0c6784b5c128c14a6f6eef535252a0c`;
  provenance SHA-256
  `53514706aa1a601491a3e7f9c99c7c8a3e656123f552f550c394f5f42404cac9`.
- The selected assets were frozen mode `0400`, fully verified, never changed,
  and never tagged or published.

## External acceptance

- The initial fail-fast runner exited before its final evidence footer. A
  bounded read-only classifier proved exact installed RC.6, absent updater
  metadata and no Task 064 mutation. Exact classification:
  `ENVIRONMENTAL ISSUE` — the VPS retained Task 061 RC.6 rather than the
  assumed final 1.1.0 baseline.
- A corrected bounded runner verified RC.6, anonymously acquired and verified
  the frozen public final 1.1.0 package, then used documented replacement to
  establish official 1.1.0 without VPS reinstall/reformat. The private RC.6
  package backup remains retained for recovery.
- Authentic native `1.1.0 -> 1.2.0-rc.1` update PASS. Exact candidate identity,
  configuration/credentials, compatible state, Guardian enabled/active intent,
  configuration validation, readiness and rollback metadata checks PASS.
- Explicit `1.2.0-rc.1 -> 1.1.0` rollback PASS. Exact official 1.1.0 binary,
  configuration/credentials, compatible state, service intent, readiness and
  consumed-metadata removal checks PASS.
- Final VPS product state is operational official QWSG 1.1.0. No credential or
  private endpoint was recorded. See
  `docs/release/ACCEPTANCE_1.2.0-rc.1.md`.

## Verification

- Builder input, metadata, prompt/history identity, approval state, lifecycle,
  Framework 2.0 and one-active-task validation PASS.
- Focused and full Go tests, repository-wide race tests, vet, formatting,
  shell syntax, systemd static validation, release/build plumbing, package
  verifier, privacy checks, Git whitespace and repository safety PASS.
- Deterministic twin construction, frozen candidate verification and bounded
  external update/rollback acceptance PASS.

## Rollback

Use only the collision-aware instructions in the verified pre-change snapshot
after checking exact repository identity. Do not use broad reset/checkout/clean
and do not touch the excluded Owner blueprint. Product transactions use only
integrity-verified package backups and fixed destinations; user configuration,
credentials and compatible state remain outside package rollback payloads.

## Completion state

Implementation and external acceptance complete; final integration and
canonical lifecycle closure in progress.
