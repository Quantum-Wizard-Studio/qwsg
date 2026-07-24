# Inventory Persistence and Digital Twin Foundation Delivery Report

## Delivery status

Task 016 implemented and verified the first QWSG file-backed Inventory Store
and persisted Digital Twin foundation. Technical implementation and mandatory
validation completed, and Attila — Project Owner formally accepted the delivery
on `2026-07-24` UTC. Task 017 was not created.

## Starting state and snapshot

Execution started on `main` at
`16fcafad6bb7c1a092c91d77a0f49bb6cf2b0af1`, matching `origin/main` with
ahead/behind `0 0`. Task 015 was complete and archived, Task 016 was the only
active production task, and the independent test-task audit was valid.

The verified rollback snapshot is
`/tmp/qwsg-task016-implementation.20260724TF9yI7R`. It contains every planned
existing target, verified absence for new paths, Git and lifecycle evidence,
file modes, baseline results, retention and restore instructions, and SHA-256
checksums. Every captured checksum passed.

Baseline framework-configured validation, Go tests, vet, and format passed.
`make engineering-test` exposed a deterministic framework-test blocker:
`test-framework.sh` hardcoded the archived Task 015 active-prompt path. The
approved Task 016 exception for a verified engineering-framework blocker was
used only to make that fixture select the current active task or latest archive.
The complete engineering suite then passed.

## Architecture

`internal/inventorystore` is a consumer of `internal/inventory`; collectors do
not know about or invoke persistence. The authorized boundary is:

```text
collect -> assemble -> redact -> validate -> persist -> load -> validate
```

The store requires a clean absolute root and rejects symlink path components.
Its layout uses private `0700` directories, a `0600` versioned metadata file,
and `0600` snapshot files. Store format `qwsg.digital-twin` version `1.0`
embeds the complete Inventory 1.0 envelope and its synchronized canonical
representation.

Snapshot names combine UTC completion time and a truncated SHA-256 derivative
of the opaque snapshot ID. The envelope records format, schema, snapshot,
subject, timing, and SHA-256 integrity metadata. Load rejects malformed JSON,
duplicate keys, trailing values, unknown fields, unsupported versions,
identity differences, integrity mismatches, invalid Inventory, unsafe modes,
unsafe file types, and traversal.

The checksum detects corruption but is not authentication or signing. No key
management, encryption-at-rest, or hostile-writer protection is claimed.

## Atomicity and retention

Save validates first, writes a same-directory temporary file, applies mode
`0600`, synchronizes and closes it, installs it through a no-clobber hard link,
and synchronizes the directory. A private exclusive lock rejects concurrent
writers. Failure injection before install proves that a previous last-known
valid snapshot remains readable and no new canonical snapshot remains.

Retention is immutable store metadata, defaults to 10, accepts 1 through 1000,
and always keeps at least one snapshot. At the bound, the deterministic oldest
snapshot is temporarily renamed; failed installation restores it. Final
retirement failure attempts bounded rollback of the new snapshot and restoration
of the old one and reports every rollback failure.

## CLI and compatibility

The existing `qwsg inventory` command remains one-shot and keeps its JSON shape
and status exit codes. Explicit additions are:

```text
qwsg inventory save --store <absolute-path> [--retention N]
qwsg inventory load --store <absolute-path> [--retention N]
qwsg inventory load --store <absolute-path> [--retention N] --snapshot <name>
```

Save emits the same collected Inventory after successful persistence. Load
performs no collection and emits the revalidated stored Inventory. Complete,
partial, and failed status semantics remain `0`, `2`, and `1`.

No collector was modified. `go.mod` and dependency state were unchanged.

## Tests and verification

New focused coverage includes:

- deterministic names and byte-identical persistence;
- exact round trip and CLI loading;
- runtime directory and file permissions;
- valid partial snapshot persistence;
- bounded retention and last-snapshot protection;
- invalid schema and canonical-model rejection;
- prohibited-secret rejection in both compatibility and canonical views;
- failure injection before atomic installation;
- corruption, truncation, integrity mismatch, duplicate keys, and unsupported
  persistence version;
- path traversal, root and parent symlinks, permissive files, duplicate
  snapshot conflict, writer locking, and retention mismatch.

The following passed:

- 8 Inventory Store tests;
- 3 CLI tests;
- all Go package tests;
- `go test -race ./...`;
- `make vet`;
- `make fmt-check`;
- 21 framework assertions;
- 36 diversion assertions;
- 28 lifecycle assertions;
- 38 Task Builder assertions;
- framework-configured engineering validations;
- active job, lifecycle, and test-task checks;
- build, help, and version commands;
- Go module/dependency diff check;
- collector/store import-boundary check;
- prohibited daemon, listener, scheduler, database, and external-dependency
  review;
- UTF-8, LF, permission, secret/private-host, and Git diff checks;
- no Task 017.

No mandatory validation failed after the baseline fixture blocker was corrected.

## Git delivery

Targeted staging contained exactly 19 reviewed Task 016 source, test,
architecture, English/Hungarian documentation, active prompt/history, and
delivery-report paths. It excluded all pre-existing backup directories,
`current-task-job.txt`, `current-task-maker.md`, build output, and unrelated
untracked content.

The implementation commit is:

```text
c792ba6c10097031104e10a00c9cf30730985dcb
Implement inventory persistence foundation
```

`git push --dry-run origin main` reported only the expected
`16fcafa..c792ba6` update. The normal HTTPS push succeeded. Post-push
verification confirmed `HEAD == origin/main`, ahead/behind `0 0`, no tracked or
staged diff, all mandatory suites passing, and no Task 017.

The separate delivery-evidence finalization commit is reported in the
owner-facing handoff because a commit cannot contain its own hash.

## Documentation

The canonical Inventory Architecture now defines persistence as a validated
post-redaction state adapter. A dedicated persistence architecture specifies
format, layout, integrity, atomicity, locking, retention, compatibility,
recovery, and deferred features. Canonical System Inventory architecture,
developer guidance, English user guidance, Hungarian user guidance, repository
indexes, system map, and engineering history were updated consistently.

## Exclusions confirmed

Task 016 introduced no daemon, scheduler, timer, polling loop, monitoring,
comparison, drift analysis, health scoring, alert, notification, web interface,
API, listener, upload, cloud synchronization, database, external dependency,
root requirement, privilege escalation, remediation, or collector persistence.
No Task 017 was prepared or created.

## Rollback

Verify the external snapshot checksums. Restore only exact Task 016-modified
paths and modes from its `files/` tree. Remove only Task 016-created repository
paths listed as absent before implementation. Rerun every baseline and final
gate. Never delete an operator Inventory Store during task rollback and never
use broad reset, checkout, restore, clean, wildcard deletion, or history
rewriting.

## Known limitations

- The persistence profile requires a local filesystem with hard links, atomic
  same-filesystem rename, file synchronization, and directory synchronization.
- Integrity is SHA-256 corruption detection, not authenticated integrity,
  signing, or encryption.
- Retention cannot be changed after store creation in version `1.0`.
- Stale lock and retirement artifacts fail closed and require operator review;
  automatic crash-recovery cleanup is intentionally absent.
- There is no schema migration tool. Unsupported versions fail closed.
- Store list is an internal API; the CLI exposes latest or explicit-name load,
  not a separate list command.
- Task tests use synthetic temporary fixtures; no raw live-host Inventory was
  persisted as task evidence.
- Project Owner acceptance was recorded before lifecycle completion.
