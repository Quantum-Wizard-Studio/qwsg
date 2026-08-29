# Task History 073: Approved Task

## Task metadata

- Task ID: `073`
- Task slug: `qwsg-1-2-0-final-release-publication`
- Status: `active — release preparation and validation in progress`
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

## Rollback

Use `/tmp/qwsg-task073-snapshot.kFyNI3/ROLLBACK.md` only after exact repository and target identity checks. Restore only individually reviewed Task 073 targets from isolated archive extraction; never use broad reset, checkout, clean, wildcard removal or published-identity rewrite.

## Completion state

`active — mandatory local release gates and deterministic final construction pending`
