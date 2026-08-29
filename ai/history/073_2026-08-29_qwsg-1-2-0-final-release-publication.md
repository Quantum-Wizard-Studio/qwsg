# Task History 073: Approved Task

## Task metadata

- Task ID: `073`
- Task slug: `qwsg-1-2-0-final-release-publication`
- Status: `complete — QWSG 1.2.0 RELEASED`
- Date generated: `2026-08-29` UTC
- Human authority: Project Owner
- Preferred owner communication language: Hungarian
- Related prompt: `ai/prompts/073_CURRENT_TASK.md`

## Lifecycle state

The Engineering Task Builder generated and transactionally installed this matching prompt/history pair from validated structured Owner input. Explicit approval was recorded and release-only execution is active.

## Starting state

- Canonical repository root, ordinary user, branch `main`, Framework `2.0.0`, and HTTPS origin validated.
- Before Builder installation the repository and index were clean at `HEAD == origin/main == 79585da0b2fa95f8b7aaa3e14fc15182d690a3b7`, with `0/0` divergence and canonical idle after completed Task 072. A post-install fetch left `origin/main` unchanged.
- Accepted RC.7 artifact `dist/qwsg-1.2.0-rc.7-linux-amd64.tar.gz` matched SHA-256 `fe5fdeb93efe0363376d5d4b85ac915b61817a3f7775c8158c5d62cb7db7c631` and remained readable.
- Accepted implementation source is `07d30315cebba4bd213e76fe9fb9c32e82aa9b38`; Task 072 final commit is `79585da0b2fa95f8b7aaa3e14fc15182d690a3b7`.
- The sole intervening commit is Task 072 lifecycle closure. Its path diff contains only the archived Task 072 prompt and Task 072 history. No product, packaging, runtime, installer, unit, release-tool or release-document file changed. Product-code equivalence is therefore proven.

## Snapshot

- Pre-change snapshot: `/tmp/qwsg-task073-snapshot.kFyNI3`, mode `0700`.
- Exact tracked-HEAD archive SHA-256: `9439455ee04aadd9c03cda72c5f99f19f584529f6b65bf4401ae6a81e8c7a99e`.
- The snapshot contains validated structured Builder input, baseline and absence records, a readable tracked archive, checksum manifest and collision-aware bounded rollback instructions. `sha256sum -c` passed for every recorded member; credentials, secrets, private access material, caches and unrelated content are excluded.

## Work performed

- Installed Task 073 transactionally with the canonical Builder and validated the active prompt/history pair.
- Read all required governance, Task 072 completion evidence, and the prior final-release workflow/evidence from Task 063.
- Classified the canonical promotion method: preserve the accepted RC.7 behavior-bearing tree; change final version and release documentation/plumbing; add only the required declarative `1.2.0-rc.2 -> 1.2.0` compatibility registry entry and exercise the existing deterministic installed-state fixture, package transaction and rollback machinery.
- Began the minimal final identity, release-documentation, plumbing and compatibility-metadata update. No feature, architecture or unrelated behavior change is included.
- Changed `VERSION` to final `1.2.0`; added final release notes and canonical Owner-supplied Contabo/OVH acceptance evidence; updated final archive installation examples and release/build plumbing assertions.
- Added exactly one declarative migration record, `compat-1.2.0-rc.2-to-1.2.0`, and retargeted the existing credential-free RC.2 installed-state fixture to the actual final version. The accepted migration transaction, preservation and rollback mechanism is unchanged.

## Verification

Builder input, metadata, prompt/history identity, approval state, lifecycle installation, clean baseline, fetched remote convergence, accepted artifact checksum/readability, snapshot integrity/readability and accepted-source product equivalence validated successfully.

- Focused final migration/update/rollback and CLI tests PASS.
- Full Go test and repository-wide race suites PASS; `go vet` and complete Go formatting PASS.
- Framework/configured validations, Builder/lifecycle/diversion/test-task suites PASS.
- Release plumbing, deterministic build provenance, cross-umask/source-mode reproducibility, shell syntax, static systemd verification, Git whitespace, security/redaction, notification and migration regressions PASS. Static systemd verification emitted expected sandbox bus-access warnings and returned success.
- The first focused Go attempt selected the sandbox-read-only default cache and failed before test setup. Classified `ENVIRONMENTAL ISSUE`; the unchanged command passed with task-private writable `GOCACHE`/`GOMODCACHE`.

## Release-source integration and deterministic construction

- Release commit: `348d927ffcf4c8cd4c9a50fc3eacad71d8bfe5c2`; tree `e7aa50f4ed405598b0d7b9a9224d8dd183e9a703`; epoch `1788022414`; UTC build time `2026-08-29T16:53:34Z`.
- Commit subject: `release: establish QWSG 1.2.0 source`; dry-run and actual push clean-fast-forwarded canonical `main`.
- Two independent exact-commit exports with isolated source/output/cache roots under `/tmp/qwsg-task073-final-build.T9zOvS` produced byte-identical archives and sidecars.
- Frozen artifact `qwsg-1.2.0-linux-amd64.tar.gz`: size `3524214`, SHA-256 `44768af20c8456cde09f940590b8c4446f605b2af02866e1553705a01d1a4c11`; sidecar SHA-256 `22242a8d0702e7dbe33c9beee84b8ef497e6372bf08e26747b754a01382e81c6`; MANIFEST SHA-256 `76fdd7809eaeffe1d923338a29b248e9413e5d1517375a5c651359ae62b118b3`; binary SHA-256 `c922f51873ee9e9428da4b900aedec31f08f9e16f86b7aba680a5bf4e96755a6`.
- Archive readability, all 18 MANIFEST file entries, safe 27-member single-root structure, numeric `0/0` ownership, canonical modes, exact executable allowlist, static linux-amd64 format, absence of ambient VCS metadata, and exact embedded final version/full commit/build time PASS.

## Publication and external distribution acceptance

- Release-boundary snapshot: `/tmp/qwsg-task073-release-boundary.rs8K4n`, mode `0700`; source archive, frozen checksum identity, proposed tag, publication ledger, containment instructions and all snapshot checksums verified.
- Annotated tag object `ac395b568b8e1f83c0ef85c9aa02f98c15402af0` was created and pushed as `v1.2.0`; local and direct remote peel is exactly release commit `348d927ffcf4c8cd4c9a50fc3eacad71d8bfe5c2`.
- Forgejo Release ID `3`, title `QWSG 1.2.0`, is final/non-draft/non-prerelease and targets the exact release commit. It contains exactly the frozen archive (`3524214` bytes) and sidecar (`96` bytes).
- Anonymous direct `wget` and `curl -fLO` retrieval from the actual Forgejo Release endpoint succeeded into independent directories. Both downloaded archives were non-empty, matched canonical SHA-256 `44768af20c8456cde09f940590b8c4446f605b2af02866e1553705a01d1a4c11`, passed downloaded-sidecar verification, opened as tar, passed every MANIFEST entry and contained the expected executable/package structure and provenance.
- No Contabo or OVH host was contacted or mutated. The real-host facts recorded in `docs/release/ACCEPTANCE_1.2.0.md` are explicitly Project Owner-supplied acceptance evidence.
- Final release decision: `RELEASED`.

## Rollback

Use `/tmp/qwsg-task073-snapshot.kFyNI3/ROLLBACK.md` only after exact repository and target identity checks. Restore only individually reviewed Task 073 targets from isolated archive extraction; never use broad reset, checkout, clean, wildcard removal or published-identity rewrite.

## Completion state

`complete — QWSG 1.2.0 released; annotated tag, final Forgejo Release and independent external artifact verification PASS`
