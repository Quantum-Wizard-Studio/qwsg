# QWSG User CLI and Snapshot Explorer — Engineering Delivery Report

## Delivery identity

- Task: `017`
- Date: `2026-07-24` UTC
- Human authority: Attila — Project Owner
- Starting commit: `2428d013c3316faf2feb9220f2edcffecf65c760`
- Framework: Reusable Engineering Framework `1.0.0`
- Delivery state: implementation and mandatory local verification passed;
  Project Owner acceptance pending

## Starting state and snapshot

Task 016 was complete and archived, Task 017 was the only approved active task,
and local `main` matched `origin/main`. Ubuntu 24.04.4 LTS, Go 1.26.5, GNU Make
4.3, and GNU install 9.4 were verified. Existing baseline Inventory output was
JSON and truthfully partial on the development host. Persistence save/load
existed; list/info, user installation, and complete contextual help did not.

The verified rollback snapshot is
`/tmp/qwsg-task017-implementation.20260724TYD2MYv`. It records 17 existing
targets, six absent repository targets, the absent `/usr/local/bin/qwsg`
system target, file modes, Git/environment state, SHA-256 checksums, retention,
and exact bounded recovery instructions. Every captured checksum passed.

## Design decisions

- Existing `qwsg inventory`, save, and load keep JSON as the compatibility
  default. `--format human` is explicit.
- `QWSG_STORE` and `QWSG_FORMAT` are visible session configuration, not
  automatic store discovery. Command options take precedence.
- Human output contains status, timestamps, schemas, and aggregate counts. It
  does not expose raw facts or imply health.
- Terminal control characters in dynamic text are escaped.
- List obtains deterministic names and validates every item with `Store.Load`.
  Info and load use that same canonical validation boundary.
- List/info/load never repair, migrate, compare, delete, prune, or rewrite.
- Version/build values are link-time variables with deterministic fallbacks.
- Installation creates one binary and supports both `PREFIX` and `DESTDIR`.

## Implementation

`cmd/qwsg` now provides root and contextual help, version/build output,
consistent usage errors, JSON/human modes, safe Inventory summaries, Snapshot
Explorer list/info/load, latest and exact-name selection, and explicit session
configuration.

The Makefile now builds with controlled metadata and installs one `0755`
binary. No external Go module or runtime dependency was added. Collectors,
Inventory schemas, canonical aggregation, and Inventory Store format and
semantics are unchanged.

Documentation now includes Ubuntu 24.04 installation, English and Hungarian
user guides, a reproducible acceptance walkthrough, updated persistence and
developer contracts, README status, and canonical indexes.

## Verification evidence

Passed:

- all Go package tests;
- `go test -race ./...`;
- vet and format checks;
- 21 framework, 36 diversion, 28 lifecycle, and 38 Task Builder assertions;
- configured Engineering Framework validations;
- production and test-task lifecycle checks;
- targeted Git diff and no-dependency/collector-change reviews;
- isolated `DESTDIR` installation with exactly one mode-`0755` binary;
- installed binary execution outside the build path;
- JSON compatibility fixture round trip;
- contextual help, argument, terminal-control, human output, list/info/load,
  empty-store, exact/latest, and status tests.

On the Ubuntu 24.04.4 development server, the installed staging binary executed
the exact documented sequence:

```text
qwsg version          -> 0
qwsg help             -> 0
qwsg inventory        -> 2
qwsg inventory save   -> 2
qwsg inventory list   -> 0
qwsg inventory info   -> 2
qwsg inventory load   -> 2
```

The `2` results were truthful partial-but-usable Inventory outcomes. Every
command produced meaningful human-readable output; stderr was empty. The store
contained only private `0700` directories and `0600` files. Info reported
`Integrity: verified`.

No daemon, scheduler, background loop, monitor, comparison, drift, health
engine, alert, notification, API, Web UI, database, listener, telemetry,
external dependency, privilege escalation, or collector modification was
introduced.

## Rollback

Restore only exact modified files from the snapshot. Remove only new Task 017
paths whose absence was recorded and whose identity remains verified. Restore
or remove an installed binary only from the recorded exact pre-install state.
Never delete an operator Inventory Store. Rerun all framework, lifecycle, Go,
race, vet, format, engineering, compatibility, install-layout, and Git checks.
Broad reset, checkout, clean, wildcard deletion, recursive prefix removal,
force-push, and history rewriting are prohibited.

## Limitations and acceptance

Task 017 is pre-alpha and verifies Ubuntu 24.04. It does not provide package
manager artifacts or broader platform guarantees. Build date remains `unknown`
unless controlled release automation supplies it. Stored snapshot checksums
detect corruption but do not authenticate data.

Implementation is ready for Project Owner review. Task 017 must not be marked
complete or archived until explicit acceptance is recorded.
