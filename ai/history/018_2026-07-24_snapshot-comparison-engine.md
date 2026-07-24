# Task History 018: Approved Task

## Task metadata

- Task ID: `018`
- Task slug: `snapshot-comparison-engine`
- Status: `active — implementation and verification in progress`
- Date generated: `2026-07-24` UTC
- Human authority: Attila — Project Owner
- Preferred owner communication language: Hungarian
- Related prompt: `ai/prompts/018_CURRENT_TASK.md`

## Lifecycle state

The Engineering Task Builder generated and transactionally installed this matching prompt/history pair from validated structured owner input. Explicit approval was recorded and the canonical `job` workflow started implementation.

## Starting state

Task 017 was complete at commit `83ae0119df62e2a4d091b362df2b2772c7499b67`; `main` matched `origin/main`. Baseline tests, race checks, format/vet, engineering checks, framework tests, and lifecycle validation passed. `qwsg compare` was not yet a command. Inventory 1.0, Canonical Inventory v1, the validated file-backed store, and the user CLI/Snapshot Explorer were present. Pre-existing untracked backups and local job helper files were preserved outside Task 018 scope.

## Snapshot

The verified implementation rollback snapshot is `/tmp/qwsg-task018-implementation.20260724Tm7qzQr`. It records 20 pre-existing tracked targets and the absence of the new comparison package, comparison architecture/schema documents, and delivery audit. The builder review and rollback evidence remain under `/tmp/qwsg-task018-*`.

## Work performed

- Added the dedicated deterministic Comparison Engine and validated Change Record model.
- Added previous/latest and explicit `--from/--to` CLI comparison with canonical JSON and record-derived human output.
- Added engine and CLI tests for comparison semantics, determinism, validation, selection, and presentation.
- Documented the architecture, schema, security boundary, developer contract, bilingual operation, and acceptance sequence.
- Preserved Inventory, collector, store, installation, and privilege contracts unchanged.

## Verification

Builder input, metadata, prompt/history identity, approval state, and lifecycle installation validated successfully. Targeted and repository-wide Go tests, race checks, vet, format, diff checks, all 21 framework assertions, 36 diversion assertions, 28 lifecycle assertions, 38 Task Builder assertions, active-task validation, diverted-task audit, and configured framework validations pass.

On Ubuntu 24.04, a normal-user build produced the one-shot binary. Two fresh private-store observations exited truthfully with partial Inventory code `2`; default and explicit comparisons both exited `0`. Repeating the same default comparison produced identical SHA-256 `ccfd1b23f66a0441c473f6ce6efa6f80eb05fe39ea6e9ad644029c98d3982e7c`. Canonical JSON reported `qwsg.comparison` `1.0`; human output grouped all four record types. The temporary store retained modes `0700`/`0600`.

After delivery commit `eec367f`, the normal-user build embedded commit `eec367f3d775`. With `PATH=/usr/bin:/bin`, artifact-only `make install DESTDIR=<isolated-stage>` succeeded without Go, the staged binary exposed compare help, and `cmp` proved it identical to `build/qwsg`. The real `sudo make install` correctly reached the privileged boundary but this non-interactive engineering session could not provide the Project Owner's sudo password. `/usr/local/bin/qwsg` therefore truthfully remains the Task 017 build `b0f4dd4e9fe5` until the owner runs the documented single local install command; installed compare acceptance remains pending.

## Rollback

Restore Task 018-owned pre-existing targets from `/tmp/qwsg-task018-implementation.20260724Tm7qzQr` using its manifest, then remove only the six targets recorded absent there. Re-run the manifest hash check, repository tests, framework validation, and lifecycle validation. Never delete or alter an operator-selected Inventory store or unrelated untracked repository data.

## Completion state

`active — implementation and verification in progress`
