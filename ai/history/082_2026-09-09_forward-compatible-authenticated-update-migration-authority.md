# Task History 082: Forward-Compatible Authenticated Update Migration Authority

## Task metadata

- Task ID: `082`
- Task slug: `forward-compatible-authenticated-update-migration-authority`
- Status: `complete`
- Date: `2026-09-09` UTC
- Human authority: Project Owner; conversation task with explicit `APPROVE`
- Preferred owner communication language: Hungarian
- Related prompt: `ai/archive_prompts/082_2026-09-09_forward-compatible-authenticated-update-migration-authority.md`

## Starting state

Verified before mutation: canonical `main`, HEAD and fetched `origin/main`
`fa570723c4d6b0a1f6758a9c2ea9ddf610716765`, ahead/behind 0/0, clean worktree.
`v1.3.1` resolves to `45009fe6169bff00842a4c4e9561bf339a5db81e`.
Framework 2.0.0 and idle lifecycle valid; Task 081 accepted and complete.
Task 082 was installed transactionally by the canonical Task Builder using
the supplied Owner approval. No Task 083 was created.

Read the Constitution, philosophy, agent rules, task template, lifecycle,
prompt/Git/execution/diagnostic/backup policies and project configuration;
then engineering, delivery, documentation, security/release policies, prior
Task 081 evidence and affected update/release architecture/runbooks.
Go toolchain: `go1.26.5 linux/amd64`; unchanged module Go version `1.26`.
No new dependency or toolchain installation was required.

## Snapshot

Protected pre-task snapshot identifier:
`qwsg-task082-baseline-20260909T162657Z` (configured local `/tmp` storage).
It contains tracked HEAD as a readable tar archive, a per-member SHA-256
manifest, Git/lifecycle baseline and a guarded restore procedure. Initial
worktree was clean; ignored files, Git internals and external host state are
outside this rollback payload.

Archive SHA-256:
`419be32ef938955cd0edd6bde6a909f62785e1b6fda47ce6e7dbe19b7c7ab5a3`.
Archive checksum/readability passed at creation. Every captured member was
rehash-verified again before closure. Retain through Owner acceptance and the
rollback window. No payload, credentials or private host capture is committed.

## Work performed

1. Added `qwsg.release-index/2` exact-source capability declarations and
   `qwsg.migration/1` constraints. The sole compiled capability is
   `preserve-package-v1`, using the existing bounded package transaction.
   Future targets are signed data, never new local version-pair entries.
2. Preserved historical `/1` canonical signing bytes, routes and installation
   classification. New metadata cannot mix historical routes with capability
   authority; duplicate/conflicting source-platform selectors refuse.
3. Kept discovery read-only and introduced `internal/updateauthority` as the
   small binding boundary. Private candidate authority binds signed source,
   target, platform, size, digest and source commit to verified package content.
   Mutable awareness/evaluation records cannot grant execution authority.
4. Replaced production update's unsigned Forgejo selection with authenticated
   index selection. Downloads use the exact signed size/digest. Equal/older
   releases remain no-update results. Known awareness generation restricts
   metadata replay; future-index clock skew uses the existing 15-minute bound.
5. Local archives now require both checksum and full production-signed
   `.release-index.json` companions. The root helper re-stages, independently
   reauthenticates, reclassifies actual installed source and checks capability,
   artifact identity/digest and package provenance before existing apply.
6. Preserved configuration preflight, fixed package destinations, service
   intent, post-validation and deterministic package rollback. No migration
   scripts, commands, remote parameters or arbitrary executable mechanism
   are accepted. Package install/uninstall scripts are never invoked by the
   capability transaction.
7. Added mandatory frozen-old-client acceptance as a `release-check`
   prerequisite, explicit `/2` publication fixtures, canonical signing
   reproducibility coverage and unit/security regression coverage.
8. Updated architecture/system map, release/operator guidance, EN/HU user
   documentation and changelog. Source VERSION, production key/trust,
   published release assets, historical index/signatures and tags are unchanged.

The common authorization/package-verification result is available for later
Pro policy orchestration. No Pro scheduler, entitlement, automatic install,
remote-management feature, Task 083 work or new release publication occurred.

## Decisions and trust boundaries

One existing preservation capability is sufficient; no general migration
interpreter or speculative transformation registry is introduced. Exact source
selectors keep selection auditable and deterministic. Unknown well-formed
capabilities/constraints are visible as unsupported; malformed/ambiguous/schema
violations refuse. Remote authority may attest compatibility of its signed
package but cannot create local executable migration behavior.

Already published 1.3.0/1.3.1 binaries cannot acquire a new parser retroactively.
A separately authorized bootstrap installation of a Task 082-capable client is
required. From that point, later compatible target versions need no backport.
The existing production release/signing state is not falsified or rewritten.

## Verification

All checks below passed using task-private caches under `/tmp` and the existing
Go module cache. No test was deleted, disabled or weakened for PASS.

| Check | Evidence/result |
|---|---|
| `make fmt-check vet test` | Full formatting, vet and all Go packages PASS |
| `go test -race ./...` | Full package race suite PASS |
| `make engineering-test` | Checkout/export build provenance/reproducibility PASS; 25 framework, 36 diversion, 29 lifecycle and 49 task-builder assertions PASS |
| `ai/tests/test-framework-v2.sh` | 15 semantic assertions PASS |
| `ai/tests/test-bounded-diagnostic-runner.sh` | 11 assertions PASS |
| `make release-check release-authority-check` | Frozen-old-client gate, release plumbing, archive reproducibility across umask/modes, Linux publication and Windows offline-signer reproducibility, `/1` and `/2` canonical signing-input identity PASS |
| Final focused authority tests | Capability constraints, conflicting selectors, unsigned/tampered metadata, schema/source/target/platform/downgrade refusals, authenticity cloning, watermark and non-upgrade/zero-candidate refusal PASS |
| Final old-client test, normal and race invocation | All 11 scenarios PASS; exact expected refusal gates asserted |
| Historical real-package tests with explicit archive environment | `TestVerifyExternalReleasePackage`, `TestRealRelease130CleanMigrationRollback`, `TestRealRelease131CleanMigrationRollback` with `-race` PASS |
| Framework/lifecycle/test-task validators | PASS; closure returns to idle with Task 082 complete |
| Git/snapshot checks | Diff whitespace, object connectivity, fetched 0/0 starting relationship, historical protected paths unchanged, snapshot per-member hashes/readability PASS |

The full Go and race runs were followed by targeted normal/race checks for
added watermark and stricter black-box refusal assertions. No product code
changed after full validation. Release artifacts were built only in isolated
verification locations; no frozen or published artifact was overwritten.

Historical archive SHA-256 identities were rechecked before real-package tests:

- 1.2.0: `44768af20c8456cde09f940590b8c4446f605b2af02866e1553705a01d1a4c11`
- 1.3.0: `8e19f624ccd1a49f32f6127889390e3389aec6eda958c7fbaaaa657d7b7cf9a0`
- 1.3.1: `0ab726bcde36182232ff89e3ea33e2d5ab77d8cad2b94ce56c9534a8147d9cd7`

These match the protected historical evidence; 1.3.1 also matches its existing
signed production index. Live-host inventory acceptance remains its existing
separately invoked test, not a Task 082 production-deployment claim.

### Mandatory black-box evidence

`TestOldBinaryForwardAuthenticatedUpdate` builds and freezes the old executable
before choosing the future target from its digest. It asserts no historical
migration-table route exists, builds the future target afterward and signs its
capability declaration using an explicit fixture key. Separate processes invoke
the old CLI, which executes the actual privileged helper dispatch. The target
binary performs only normal post-install identity/configuration checks.

The final normal run selected `1.113.109` after freezing the old binary. Earlier
runs selected different targets as the test-linked source changed; none were
added to migrationPaths. Cases: online, archive, rollback-after-failure,
unknown-capability, unsigned, source-mismatch, provenance-mismatch,
digest-mismatch, archive-missing-authority, helper-reauthentication and
archive-forged-sidecar. Discovery alone leaves the installed binary and helper
untouched. Successful transactions and injected failure restore the old binary
and preserve private state. Test-only host seams are absent from shipped
executables; no production environment/flag overrides are introduced.

### Diagnosed attempts

- `ENVIRONMENTAL ISSUE`: sandbox denied local socket creation for existing TLS
  tests. The authorized rerun outside sandbox passed; test semantics unchanged.
- `TEST OR ACCEPTANCE DEFECT`: the preexisting unknown-schema test used `/2`,
  now intentionally supported. Its unknown-schema input was changed to `/999`;
  strict rejection assertions retained and `/2` contract coverage added.
- `ENVIRONMENTAL ISSUE`: checkout/export verification copies only Git-followed
  files, so the first run omitted the new untracked authority package. Reviewed
  new paths were explicitly staged; the unchanged canonical export test and
  subsequent engineering suite passed. No dependency or test bypass was used.

## Rollback

Verify the snapshot checksum and per-member manifest. Extract into a new empty
private directory only. Compare the exact Task 082 changed paths with the
baseline before any individually reviewed restoration. Never extract over a
live worktree, reset broadly, rewrite published history or delete unrelated
files. Restore/removal must be bounded to the reviewed task path set; preserve
this audit history. Re-run Git diff, framework/lifecycle and relevant tests.

Runtime rollback remains the existing protected package before-image journal.
Both the forward-compatible black-box transaction and historical real packages
passed exact rollback checks. Configuration/credentials/private state remain
outside package mutation. No production rollback was performed.

## Repository integration and completion state

Task-scoped paths were reviewed for scope, permissions/executable bits,
whitespace, private material and backup exclusion. The final commit contains
implementation, tests, documentation, this report and completed prompt archival.
The commit's own identity and dry-run/push/post-push relationship are reported
in the Owner handoff to avoid a self-referential hash in this record.

Task 082 is complete. Canonical idle closure creates no Task 083. Remaining
external work is only the explicitly out-of-scope bootstrap/release publication
and later Pro orchestration; neither is claimed delivered here.
