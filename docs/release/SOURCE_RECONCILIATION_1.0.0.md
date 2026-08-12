# QWSG 1.0 Source Reconciliation

## Status and boundary

This is the pre-staging classification for the accepted Task 025–043 working
tree relative to `main` and `origin/main` commit
`0a8a5c7e722495b8c5eb425bca5b2d2413aaa175`. It is not a staging instruction,
commit authorization, tag or publication action. The Owner-approved legal
license and final `1.0.0` metadata pass are now present and validated. The
resulting exact lists below remain proposals until the next Owner gate.

The audited state contains 208 tracked paths, 28 tracked modifications, 111
untracked paths after Task 044 installation and 90 ignored files. The Git index
is empty. No audited work-tree path outside `.git` is a symlink or special file.

## Classification

### Canonical release source and tests

All existing tracked source plus the accepted untracked Go packages under
`internal/alert`, `internal/configuration`, `internal/guardian`,
`internal/notification`, `internal/operatorconsole`, `internal/operatorstate`,
`internal/policy`, `internal/presentationmodel`, `internal/runtime`,
`internal/runtimeservice` and `internal/scheduler`, together with
`internal/app/operator.go` and `internal/report/policy.go`, are canonical.
Their tests are canonical source evidence. `cmd/`, `go.mod`, `Makefile` and
other tracked build inputs remain canonical.

### Canonical architecture and operator documentation

Tracked documentation changes and accepted untracked documents under
`docs/architecture/`, `docs/release/` and `docs/user/` are canonical except for
`docs/architecture/QWCS_MIGRATION_BLUEPRINT.md`. The QWCS Engineering Operating
System document is canonical because tracked core/framework architecture
explicitly references it. The migration blueprint is a separate unreferenced
Owner draft and remains excluded pending explicit scope.

### Canonical engineering evidence

Task prompt archives `025` through `043`, histories `025` through `043`, the
active Task 044 prompt/history and existing tracked governance records are
canonical lifecycle evidence. The untracked
`ai/backups/20260721T_task013_engineering_task_builder/SHA256SUMS` is canonical
audit metadata because Task 013 history explicitly requires it for snapshot
verification; the ignored backup archive it identifies remains local.

### Canonical packaging and release tooling

`packaging/release/`, `packaging/systemd/` and `scripts/build-release.sh` are
canonical release source. Generated archives and binaries are not source.

### Generated or local content excluded from a source commit

- `build/` and `dist/`, including RC.1–RC.3 archives and checksum sidecars;
- root `qwsg`, a locally compiled static executable;
- `current-task-job.txt` and `current-task-maker.md`, repeatedly identified by
  prior task histories as preserved local Owner input;
- ignored caches, coverage, logs, credentials patterns and dependency output;
- ignored backup archives, payloads, preserved working copies, raw diffs,
  status dumps, host evidence and other snapshot payloads under `ai/backups/`;
- `docs/architecture/QWCS_MIGRATION_BLUEPRINT.md`, preserved unchanged as an
  unrelated Owner-owned draft.

Exclusion means “do not stage”; it does not authorize deletion, movement,
permission changes or content changes.

## Exact proposed staging rule

The release-source candidate is the union of:

1. every currently tracked path, subject to final diff/privacy/mode review;
2. the accepted untracked source/test, architecture, release, user,
   packaging and release-script paths classified above;
3. Task `025`–`044` canonical prompt/history records and the Task 013 snapshot
   checksum metadata;
4. Task 044 final legal/license, metadata, acceptance and reconciliation paths
   after their separate gates pass.

It explicitly excludes every path listed in the local/generated section. No
glob or directory-wide Git staging command may implement this rule. The sorted
exact allowlist contains 140 paths and has SHA-256
`7d0bdf1da05ba49d9027ba2808acb70e2e490b40071b84f5ca9b1b0edac010bb`.
The sorted exact exclusion list contains 94 paths and has SHA-256
`c09a72cbe45dd2b2961c3bc72c7974ad8f62836136f28697b7870a1cc903fe3b`.
Together they classify all 234 modified, untracked and ignored status paths,
with zero overlap. Their exact bytes are retained in the verified pre-staging
snapshot. The first `git add` still requires explicit Project Owner
authorization and must name only the reviewed allowlist paths.

## License gate

The Owner approved and supplied the exact QWS Community / Free License Version
1.0 text. It grants bounded Community use, internal modification and unmodified
official-release redistribution while reserving modified redistribution,
rebranding and Pro Services. It expressly identifies QWSG as proprietary
source-available software, not OSI open source. Community capability policy,
future Pro/API architecture and binding legal distribution rights remain
distinct and are now consistent. Staging and commit still require the next
explicit Owner authorization.
