# Task History 017: Approved Task

## Task metadata

- Task ID: `017`
- Task slug: `qwsg-user-cli-and-snapshot-explorer`
- Status: `complete — accepted by Project Owner`
- Date generated: `2026-07-24` UTC
- Human authority: Attila — Project Owner
- Preferred owner communication language: Hungarian
- Related prompt: `ai/prompts/017_CURRENT_TASK.md`

## Lifecycle state

The Engineering Task Builder generated and transactionally installed this matching prompt/history pair from validated structured owner input. Explicit approval was recorded. Execution began on `2026-07-24` UTC under the approved Task 017 scope.

## Starting state

Execution began from repository root
`/home/qws/web/qwsg.quantumwizard.hu/qwsg`, branch `main`, at
`2428d013c3316faf2feb9220f2edcffecf65c760`. Local `main` and `origin/main`
matched with ahead/behind `0 0`; the canonical HTTPS remote and tag `v0.1.0`
were present. Ubuntu 24.04.4 LTS, Go 1.26.5, GNU Make 4.3, and GNU install 9.4
were verified.

Task 016 was complete and archived, Task 017 was the only active approved
production task, no Task 018 existed, and the independent test-task namespace
was valid. All configured framework validations, Go tests, race tests, vet,
format, 21 framework assertions, 36 diversion assertions, 28 lifecycle
assertions, and 38 Task Builder assertions passed before implementation.

The baseline CLI emitted Inventory 1.0 JSON with additive canonical inventory
and truthfully exited `2` on this host because optional evidence was partial.
Save and load existed; list and info did not. `/usr/local/bin/qwsg` was absent.

## Snapshot

The verified rollback snapshot is
`/tmp/qwsg-task017-implementation.20260724TYD2MYv`. It contains every planned
existing source, test, documentation, prompt, history, and index target;
verified absence records for new documentation, audit, `go.sum`, and installed
binary paths; baseline environment and Git evidence; file state; manifest;
SHA-256 checksums; retention policy; and bounded restore instructions. All 17
captured file checksums passed.

## Work performed

Implementation is in progress. The CLI now has contextual help, controlled
version/build metadata, explicit JSON/human output modes, terminal-safe summary
rendering, deterministic Snapshot Explorer list/info/load behavior, and
session-scoped explicit store/format configuration. JSON remains the
compatibility default. Makefile build and prefix/destination-aware install
targets and initial English/Hungarian user, installation, and demonstration
documentation have been added.

The Snapshot Explorer validates every listed or inspected snapshot through the
existing Inventory Store load boundary. List/info/load do not repair, migrate,
compare, delete, prune, or rewrite. No collector, Inventory schema, canonical
aggregation, persisted format, store safety rule, or external dependency was
changed.

Canonical architecture, system map, project structure, milestone index,
README, persistence architecture, developer guide, existing Inventory guides,
and the Task 017 delivery audit were updated.

## Verification

Builder input, metadata, prompt/history identity, approval state, and lifecycle installation validated successfully.

Implementation verification passed:

- all Go package tests and `go test -race ./...`;
- vet and format checks;
- 21 framework, 36 diversion, 28 lifecycle, and 38 Task Builder assertions;
- configured framework validations and production/test-task lifecycle checks;
- isolated `DESTDIR` installation containing exactly one mode-`0755` binary;
- installed-binary version/help execution outside the build path;
- JSON compatibility, human rendering, contextual help, invalid arguments,
  terminal control escaping, empty store, exact/latest list/info/load, and
  partial-status tests;
- no `go.sum`, external dependency, collector diff, daemon, scheduler,
  monitoring, comparison, health, alert, API, Web UI, database, listener, or
  Task 018.

The exact documented acceptance sequence passed on Ubuntu 24.04.4 using the
installed staging binary and explicit `QWSG_STORE`/`QWSG_FORMAT=human`
configuration. Version/help/list exited `0`; live Inventory, save, info, and
load truthfully exited `2` for partial but usable evidence. Every command
produced meaningful stdout, stderr remained empty, store directories were
`0700`, files were `0600`, and info reported verified integrity.

Detailed delivery evidence:
`ai/audits/2026-07-24_QWSG_USER_CLI_SNAPSHOT_EXPLORER.md`.

Targeted staging included exactly 19 Task 017 paths and excluded every
pre-existing backup directory, task-maker input, and build artifact.
Implementation commit `5d34849` (`Implement user CLI and snapshot explorer`)
passed the HTTPS dry-run push and normal push to `origin/main`.

Post-push HEAD and `origin/main` matched at delivery-evidence commit `15a9115`
with ahead/behind `0 0`. A final system-prefix acceptance attempt stopped
safely because `/usr/local/bin` is not writable by `attila` and `sudo -n`
requires an interactive password. No password was requested or handled,
`/usr/local/bin/qwsg` remains absent, and no system target changed. The
Project Owner must execute the documented `sudo make install` step locally
before final acceptance; all isolated install and installed staging-binary
acceptance evidence remains successful.

The Project Owner then confirmed a second acceptance defect: interactive
`sudo make install` found the Makefile's `install: build` dependency, but Go at
`/usr/local/go/bin/go` was not present in sudo's restricted `PATH`. The command
failed before installation with `go: not found`; no system target changed.
The corrected contract builds as the ordinary user and makes `install` validate
and copy only the existing executable artifact. The privileged step no longer
invokes Go or depends on the normal user's toolchain environment.

The corrected install contract passed with `PATH=/usr/bin:/bin`, where `go`
was confirmed unavailable: it installed exactly one mode-`0755` binary from
the normal-user build artifact. The Project Owner's subsequent system install
created `/usr/local/bin/qwsg`; it is a regular mode-`0755` file and its SHA-256
exactly matches `build/qwsg`.

The complete exact acceptance sequence was then rerun through the system
`qwsg`. Version/help/list exited `0`; Inventory/save/info/load truthfully
exited `2`; every command produced meaningful stdout and empty stderr; and the
private store retained `0700` directories and `0600` files.

Install separation fix commit `f02de2d` and final system acceptance evidence
commit `1d4808e` passed dry-run and normal HTTPS push. At formal acceptance,
local and remote `main` were synchronized and all mandatory validation gates
remained successful.

## Rollback

Restore only exact modified paths from the snapshot and remove only new paths
whose absence was recorded. Restore or remove `/usr/local/bin/qwsg` only from
its separately recorded pre-install state and verified Task 017 hash. Never
remove an operator Inventory Store. Broad reset, checkout, restore, clean,
wildcard deletion, or recursive prefix removal is prohibited.

## Completion state

`complete — implementation, system installation, and mandatory verification passed; formally accepted by Attila — Project Owner on 2026-07-24 UTC`

## Project Owner acceptance

Attila — Project Owner provided the explicit acceptance statement
`Project Owner Acceptance = APPROVED` on `2026-07-24` UTC after reviewing the
Task 017 Final Owner Review Package and the corrected privileged-install
workflow. This satisfies the final completion gate. No successor task was
authorized or created by this acceptance.
