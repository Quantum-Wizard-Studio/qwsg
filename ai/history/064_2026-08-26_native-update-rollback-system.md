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

[PENDING UNTIL TASK START]

## Work performed

None. Approved task not started.

## Verification

Builder input, metadata, prompt/history identity, approval state, and lifecycle installation validated successfully.

## Rollback

[PENDING UNTIL TASK START]

## Completion state

`approved — task not started`
