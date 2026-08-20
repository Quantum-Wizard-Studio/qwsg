# Task History 049: Approved Task

## Task metadata

- Task ID: `049`
- Task slug: `qwsg-1-1-external-clean-host-acceptance-release-readiness`
- Status: `complete with disclosed limitations`
- Date generated: `2026-08-20` UTC
- Human authority: Project Owner
- Preferred owner communication language: English
- Related prompt: `ai/prompts/049_CURRENT_TASK.md`

## Lifecycle state

The Engineering Task Builder generated and transactionally installed this matching prompt/history pair from validated structured owner input. Explicit approval was recorded. The Project Owner authorized implementation preparation on 2026-08-20 UTC, then separately authorized the exact candidate-source integration, clean fast-forward push, private twin build, transfer, and first external execution. Candidate construction is complete. External Smart Install failed its operator-guidance gate at **Checkpoint 04 — Run Smart Install before host mutation**; all later physical execution and SMTP credential handling remain stopped.

On 2026-08-20 UTC, the Project Owner explicitly authorized truthful terminal
closure as `complete with disclosed limitations`. The final verdict is **NOT
READY FOR QWSG 1.1.0 RELEASE**. This lifecycle closure does not represent a
successful external acceptance: `QWSG-049-F002` remains a release blocker,
`QWSG-049-F003` remains release-gate blocking, and no operator workaround was
used. The immutable historical candidate remains QWSG `1.1.0-rc.1`, source
commit `ff2eb2b12499f5daf3b5ba11b1f8d7fc562f8a31`, artifact SHA-256
`aa139faaccc1cc85b50cfe0eedee9436539ae1c3071e01d8e9ed9283fc7f8239`.

## Starting state

- Verified repository root `/home/qws/web/qwsg.quantumwizard.hu/qwsg`, UTC date
  `2026-08-20`, ordinary user `attila`, branch `main`, canonical HTTPS origin,
  HEAD and `origin/main` both
  `57d37206c221a85b83cccb974525bc3ba4e408c1`, ahead/behind `0/0`, empty
  index, and clean tracked tree.
- Task 049 is the sole active approved prompt/history pair. Approved Builder
  source SHA-256 is
  `de37cccc4c210e43c8d7ae4bf3eeba95b2f27f5d65e68089c1353cd4d04c32cf`.
- The only visible untracked state before implementation was the Task 049
  prompt/history plus the preserved Owner draft. LICENSE SHA-256 remained
  `e66792f9831509ecc888f0c06d64b89ae87814fe37842e12df6b716b404f4126`;
  Owner-draft SHA-256 remained
  `c83a0b2e5269de51acd9485b44fbe4f752e1f7f172c180e4126f26e91fdf8a73`.
  Annotated `v1.0.0` remained object
  `7abc83da6185199606c2d76ac7d3504ddd78cf68`, peeling to
  `177535e44b2ce5ed9efd73ab0793ffe6881f0cd6`; published archive identity
  remained
  `edfba7366adf2c1ce0a8ce56369bb0dc5ad11326c4e3d1e301625a5313292fa5`.
- Pre-change Framework 21, diversion 36, lifecycle 28, Builder 38, test-task,
  full Go, repository race, vet, formatting, shell syntax, and job gates passed.
  Static systemd verification emitted the known sandbox bus warning and will
  be repeated through the host-local verification boundary. Two independent
  pre-change v1.0.0 builds from exact HEAD/epoch were byte-identical; this was
  local automated evidence only.

## Snapshot

Verified mode-0700 implementation snapshot:
`/tmp/qwsg-task049-implementation.mtkVTh`. It contains a readable tracked-HEAD
archive, exact active prompt/history/Builder source, Git and release identities,
status/diff evidence, checksums, and bounded restore instructions. The Owner
draft is represented by metadata/hash only. All snapshot manifest entries pass.

Separate mode-0700 private evidence root:
`/tmp/qwsg-task049-private-evidence.773atT`. At this gate it contains only a
verified policy and checksum manifest—no candidate, host evidence, provider
data, or credential. Its policy requires Owner authority before later content
and bounded audited disposal.

## Work performed

- Changed the private candidate source identity to `1.1.0-rc.1` and narrowly
  generalized `scripts/build-release.sh` to accept only the historical 1.0
  identities plus the approved numeric 1.1 RC identity. Added a validate-only
  mode so pre-integration gates cannot create a candidate-shaped archive.
- Added a deterministic release-plumbing check and Make target. Actual archive
  construction remains gated on a clean integrated candidate-source commit.
- Added private RC release notes and aligned the packaged installation command
  examples and known-notification limitation with the accepted post-1.0
  product. README wording no longer incorrectly identifies the product version
  while retaining the unchanged QWS Community / Free License Version 1.0.
- Added a 16-checkpoint Owner-operated clean-host protocol. Every checkpoint
  includes purpose, bounded action, expected evidence, PASS, FAIL/finding,
  continuation safety, and retention/redaction rules. It covers provenance,
  documentation, Smart Install, guided setup/resume, protected credential
  injection, real receipt, activation/fresh evidence, systemd, lingering,
  logout, physical reboot, post-reboot receipt, restart, uninstall, and
  mandatory reinstall/resume.
- Added an acceptance/readiness ledger whose mandatory gates are explicitly
  `NOT EXECUTED` and whose present verdict is `NOT READY FOR QWSG 1.1.0
  RELEASE` solely because physical/provider gates have not begun.
- No candidate archive, external transfer/action, SMTP credential, product
  repair, tag, release, publication, Task 050, staging, commit, or push was
  performed.

The Project Owner then authorized the exact 12-path integration. Pre-stage
allowlist/content/security checks passed; the content-manifest SHA-256 was
`a7841bb2b3b49d8241ef7e0f7d16972f860c30c259139189be8402d9ce4f8cde`.
Commit `ff2eb2b12499f5daf3b5ba11b1f8d7fc562f8a31` with subject
`release: prepare QWSG 1.1.0-rc.1 acceptance source` has parent
`57d37206c221a85b83cccb974525bc3ba4e408c1` and exactly the approved 12
paths. Post-commit gates passed, the dry-run passed, and `main` was pushed as a
single clean fast-forward so `origin/main` equals the candidate-source commit.

Two independent `git archive` source roots built the private candidate with
commit epoch `1787237840` (`2026-08-20T14:57:20Z`). The first invocation method
ran the release script outside the exported Go module and was recorded as
`attempt-failed`; Go stopped before archive creation. Running inside each
exported source root produced byte-identical archive and sidecar. Archive
`qwsg-1.1.0-rc.1-linux-amd64.tar.gz` is SHA-256
`aa139faaccc1cc85b50cfe0eedee9436539ae1c3071e01d8e9ed9283fc7f8239`.

Sidecar, internal manifest, exact single-root layout, static linux-amd64
binary, version, full embedded commit, build time, LICENSE/README/INSTALL byte
identity, release notes, modes, secret scan, and isolated installer/uninstaller
passed. Installed documentation is correctly mode `0644` and the installer
prints documentation, setup, and readiness guidance.

Pre-transfer link-resolution verification found `QWSG-049-F001`, classified
`COSMETIC / POST-RELEASE CANDIDATE`: the exact canonical document copies retain
repository-context Markdown targets that do not resolve inside the archive.
Both documents explicitly name their archive-root counterpart and installed
paths in prose, satisfying the approved byte-identity and operator-guidance
contract without hidden engineering knowledge. No required action is missing,
so the finding is non-blocking. Candidate transfer and all external checkpoints
remain stopped solely because they require the next explicit Owner gate.

The Owner then supplied privacy-safe external Smart Install output from the
clean Ubuntu 24.04 amd64 host. QWSG correctly classified
`systemd.user_manager` as `missing_required`, kept environment/installation
`not_ready`, and refused continuation. Finding `QWSG-049-F002` is a `RELEASE
BLOCKER`: the product supplied only `resolve_required_findings`, with no cause,
verification action, privilege boundary, or proven remediation. The Owner
correctly did not mask the missing product guidance with developer knowledge.

The simultaneous `filesystem.local_semantics =
unknown_requires_verification` produced no bounded manual-verification
instruction. Finding `QWSG-049-F003` is a `UX/DOCUMENTATION DEFECT`; although
the requirement is recommended/non-blocking at runtime, the mandatory Smart
Install manual-verification acceptance gate fails and therefore blocks release
readiness.

Repository audit proves these are not hidden presentation data. Both Registry
1.0 requirements have empty remediation lists. The presenter already renders
canonical remediation when present. The user-manager probe collapses distinct
failures into `systemd_user_unavailable`, so current evidence cannot safely
select one exact command. Filesystem semantics is emitted unconditionally as
unknown and has neither a probe nor a structured manual plan. The current model
cannot express non-command verification because remediation validation requires
commands.

The smallest correction changes product assessment/model/registry/probe/
presentation behavior and tests, not Task 049 acceptance tooling. It requires
separate Owner authority. No workaround, source correction, next checkpoint,
SMTP credential request, external mutation, release action, or Task 050 was
performed. Acceptance remains `NOT READY FOR QWSG 1.1.0 RELEASE`.

## Verification

Builder input, metadata, prompt/history identity, approval state, and lifecycle
installation validated successfully. Candidate-source preparation passes:

- `make build` with local binary identity `1.1.0-rc.1`; full `go test ./...`;
  repository-wide `go test -race ./...`; vet; formatting; shell syntax; Git
  whitespace; Framework 21; lifecycle 28; Builder 38; diversion 36; test-task
  audit; and active-job validation;
- host-local static `systemd-analyze verify --user` with no diagnostic;
- release validate-only checks for exact RC identity, mandatory explicit epoch,
  mandatory full lowercase 40-character commit, archive/sidecar collision
  refusal, README/INSTALL references, and all seven required fields across all
  16 protocol checkpoints;
- snapshot manifest/readability and isolated rollback simulation at
  `/tmp/qwsg-task049-rollback-simulation.uiZZUy` against every modified tracked
  path and every newly created target;
- empty index, unchanged HEAD/origin and `0/0`, unchanged v1.0.0 tag/target,
  LICENSE and Owner-draft hashes, no candidate archive in the repository, and
  no private-key/access-key marker in the Task 049 source set.

The exact proposed candidate-source integration set is 12 paths:

1. `Makefile`
2. `README.md`
3. `VERSION`
4. `ai/history/049_2026-08-20_qwsg-1-1-external-clean-host-acceptance-release-readiness.md`
5. `ai/prompts/049_CURRENT_TASK.md`
6. `docs/installation/INSTALL.md`
7. `docs/release/ACCEPTANCE_1.1.0-rc.1.md`
8. `docs/release/ACCEPTANCE_PROTOCOL_1.1.0-rc.1.md`
9. `docs/release/KNOWN_LIMITATIONS.md`
10. `docs/release/RELEASE_NOTES_1.1.0-rc.1.md`
11. `scripts/build-release.sh`
12. `scripts/test-release-plumbing.sh`

The proposed single commit subject is `release: prepare QWSG 1.1.0-rc.1
acceptance source`. Explicit path-based staging, commit, any push, and the
subsequent clean-commit twin candidate build all require the next Project Owner
authorization. `current-task-job.txt`, the Owner draft, build output, snapshots,
caches, private evidence, and every unrelated/generated path are excluded.

## Rollback

Verify `/tmp/qwsg-task049-implementation.mtkVTh/SHA256SUMS` and review current
Task 049 path ownership. Restore only literal Task 049 tracked targets from the
snapshot's tracked-HEAD archive; restore lifecycle files individually from its
`lifecycle/` directory when their identities require it. Remove only Task
049-created paths after proving their pre-task absence. Never use broad Git
reset/clean/checkout and never alter the Owner draft, LICENSE, v1.0.0, or
published evidence. Re-run all validation and preservation gates after a
rollback simulation.

## Completion state

`complete with disclosed limitations — NOT READY FOR QWSG 1.1.0 RELEASE`

External acceptance stopped at **Checkpoint 04 — Run Smart Install before
host mutation**. `QWSG-049-F002` remains `OPEN, BLOCKING` and
`QWSG-049-F003` remains `OPEN, RELEASE-GATE BLOCKING`. Checkpoints 05–16 were
not begun. No manual workaround, external continuation, SMTP credential
handling, RC replacement, tag, release, upload, or publication occurred.
