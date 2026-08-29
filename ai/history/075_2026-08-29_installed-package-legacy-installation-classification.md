# Task History 075: Approved Task

## Task metadata

- Task ID: `075`
- Task slug: `installed-package-legacy-installation-classification`
- Status: `complete — implementation validated for Owner review`
- Date generated: `2026-08-29` UTC
- Human authority: Project Owner
- Preferred owner communication language: English
- Related prompt: `ai/prompts/075_CURRENT_TASK.md`

## Lifecycle state

The Engineering Task Builder generated and transactionally installed this matching prompt/history pair from validated structured owner input. Explicit approval was recorded, and the Project Owner separately directed canonical execution.

## Starting state

- Repository: `/home/qws/web/qwsg.quantumwizard.hu/qwsg`
- Branch: `main`
- HEAD: `19571803edb8d2f4bf47688ce277cc574a216d9b`
- `origin/main`: `19571803edb8d2f4bf47688ce277cc574a216d9b`
- Divergence: `0/0`
- Initial index: clean
- Initial work tree: only the Builder-installed Task 075 prompt/history pair
- Protected tag object: `ac395b568b8e1f83c0ef85c9aa02f98c15402af0`
- Protected tag peel/release commit: `348d927ffcf4c8cd4c9a50fc3eacad71d8bfe5c2`
- Protected release artifact SHA-256: `44768af20c8456cde09f940590b8c4446f605b2af02866e1553705a01d1a4c11`

## Snapshot

- Retained Builder snapshot: `/tmp/qwsg-task075-builder-install.Jqyr6l`
- Execution snapshot: `/tmp/qwsg-task075-execution.1xhnPr`
- Snapshot directory mode `0700`; evidence files mode `0600`
- The execution snapshot records Git/lifecycle/release identities, target
  payloads or prior-absence facts, metadata/ACL evidence, checksums, a readable
  bundle, and bounded literal rollback instructions.
- Snapshot checksum and readability validation: PASS before production change.

## Work performed

- Inventoried guided installation, native update, package provenance,
  installed destinations, version output, migration registry, configuration,
  state and rollback paths. The weak decisions were the guided installer's
  `/usr/local/bin/qwsg` `Lstat` check and updater's direct version-output trust.
- Added `internal/installation`, the single read-only classification boundary.
  It checks safe canonical layout, strict installed `qwsg.release/1`, exact
  executable/provenance agreement, supported major identity and optional exact
  declared migration route.
- Added bounded safe structured states/reasons. Configuration, credentials,
  Guardian/checkpoint/readiness/rollback state, PATH and arbitrary file names
  are deliberately excluded as package evidence.
- Integrated guided installation. No package artifacts select fresh install; a
  verified package may continue setup without package copying; a declared
  upgrade source is directed to `qwsg update`; legacy, unknown and inconsistent
  states stop before consent or mutation with localized guidance.
- Integrated the existing updater's installed-version and post-install checks
  with the classifier. Candidate verification, migration planning, privileged
  transaction and rollback remain unchanged.
- Added deterministic classification, spoofing, unsafe-path, updater and guided
  decision tests, including the `0.0.1-prealpha` field regression.
- Added the canonical architecture contract and synchronized installer,
  discovery, setup, upgrade, system-map, roadmap and engineering-history docs.

## Verification

- Focused `go test ./internal/installation ./cmd/qwsg ./internal/update`: PASS.
  This includes clean host, verified current/older package, exact supported
  upgrade, unsupported source, pre-alpha binary-only, arbitrary/version-only
  binary, partial/malformed/mismatched package, unsafe file and ancestor
  symlink, stale state, guided refusal and native-updater identity boundaries.
- `make fmt-check vet test`: PASS for formatting, `go vet ./...`, and every Go
  package.
- `go test -race ./...`: PASS for every Go package.
- `make engineering-test`: PASS after exact Task 075 staging made the new
  package visible to its Git-index-based isolated build fixture: Framework 25,
  diversion 36, lifecycle 29 and Builder 49 assertions, plus deterministic
  build-contract checks.
- The first `make engineering-test` attempt failed because the build-contract
  fixture intentionally exports only indexed files and the new classifier was
  not yet staged. Classified as an acceptance-fixture precondition, corrected
  by targeted staging, and the unchanged suite then passed.
- `scripts/test-release-reproducibility.sh`: PASS across source modes and
  umasks. `make release-check` stopped at its intentional clean-output
  precondition because the immutable published final artifact already exists
  in `dist/`; classified expected behavior. The artifact was not removed or
  overwritten merely to exercise release construction.
- `bin/job --check`, Framework validation, `git diff --check`, staged diff
  check, targeted scope/mode/ACL inspection and bounded secret-pattern scan:
  PASS.
- Protected evidence remains exact: annotated tag object
  `ac395b568b8e1f83c0ef85c9aa02f98c15402af0`, peel/release commit
  `348d927ffcf4c8cd4c9a50fc3eacad71d8bfe5c2`, artifact size `3524214`, and
  SHA-256
  `44768af20c8456cde09f940590b8c4446f605b2af02866e1553705a01d1a4c11`.
- No external host was mutated or contacted during Task 075 execution. No
  release discovery, Guardian update scheduling, update-availability
  notification, unattended update, Pro or Task 076 implementation was added.

## Rollback

Restore only recorded pre-existing targets from the verified execution
snapshot and remove only Task 075-created paths whose absence is recorded.
Reapply recorded metadata, then re-run focused/full/framework/lifecycle/Git and
release-identity checks. Do not use broad reset/checkout/clean, mutate refs, or
touch published release evidence.

## Completion state

`complete — implementation and mandatory validation passed; pending canonical idle closure and Owner review`
