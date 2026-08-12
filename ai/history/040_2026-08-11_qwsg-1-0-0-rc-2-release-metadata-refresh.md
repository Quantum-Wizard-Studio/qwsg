# Task History 040: Approved Task

## Task metadata

- Task ID: `040`
- Task slug: `qwsg-1-0-0-rc-2-release-metadata-refresh`
- Status: `complete`
- Date generated: `2026-08-11` UTC
- Human authority: Project Owner
- Preferred owner communication language: Hungarian
- Related prompt: `ai/archive_prompts/040_2026-08-11_qwsg-1-0-0-rc-2-release-metadata-refresh.md`

## Lifecycle state

The Engineering Task Builder generated and transactionally installed this matching prompt/history pair from validated structured owner input. The Project Owner started Task 040 through the canonical `job` workflow on 2026-08-11 UTC.

## Starting state

- Canonical lifecycle, repository root, `main`, HEAD `0a8a5c7e722495b8c5eb425bca5b2d2413aaa175`, canonical HTTPS origin, `0/0` upstream relationship, empty index, Task 039 latest-complete baseline and sole active Task 040 were verified.
- Ubuntu 24.04, linux-amd64, Go 1.26.5, GNU Make 4.3, GNU tar 1.35, gzip 1.12, SHA-256 tooling and systemd 255 were verified with ample ext4 workspace capacity.
- RC.1 archive SHA-256 remained `5b9751048e54eae584b89a4339877813ab0604297a1338ae4516b5657a8716c8`; no RC.2 target collision or active Guardian/release process existed.
- Pre-change build, full/race tests, vet, format, Framework, Builder, lifecycle, diversion, Git and job checks passed. Shipped-unit verification reached only the already documented sandbox user-bus restriction.

## Snapshot

Verified implementation snapshot: `/tmp/qwsg-task040-implementation-20260811T.xANKcT`. It contains exact pre-change payloads and RC.2 absence records, guarded restore instructions, deterministic SHA-256 manifests and readable snapshot archive. Payload manifest SHA-256 is `5bce11386aa13899687c4085b94163dcc17e78ddf59bb7eb7b1ab1b71f0ba85b`; snapshot tar SHA-256 is `92ccaef4fe30b084f27782bdd56dd91cf0a7fd52f88ebe23b62ccd4281b0d65b`.

## Work performed

Advanced the coherent repository and command identity to `1.0.0-rc.2`, added
the matching changelog, release notes and engineering-acceptance record, and
updated Quick Start references. The release builder now selects the exact
validated version-derived release-notes file and emits a location-independent
checksum sidecar. No product decision semantics or licensing content changed.

Produced the canonical linux-amd64 archive and sidecar using fixed commit and
epoch inputs. Preserved all RC.1 artifacts and records byte-for-byte. Exercised
the extracted archive only in unique staging, foreground-lifecycle and isolated
systemd roots; the development-host installation destination and real service
were not modified.

## Verification

Builder installation, starting environment, pre-change release gates and the
implementation snapshot validated successfully. Two independent assemblies
produced byte-identical binary, manifest, archive and sidecar. The final archive
SHA-256 is `0694fb7f382ea1b373aaf0b6f0171a3fc580c491c1af3b8657a3bb8697ed897b`;
its internal manifest SHA-256 is
`f2d70d8e444d0b4021007f9e4dc7052629a0877fce586310e4ce8d04ad2aa206`.

Manifest, archive safety, privacy, identity, archive-only staged installation,
collision, replacement, backup and uninstall checks passed. The extracted
static binary requires no Go, Git or checkout. Complete Go, race, vet, format,
Framework, Builder, lifecycle, diverted-task, repository and Git gates passed.

Artifact-level Task 039 acceptance used the final extracted binary. Recurring
cycles succeeded above the live 366-source boundary, reaching 456 sources;
Runtime completed without Alert source-count failure; Console `r` refreshed
read-only; explicit `observe` retained `guardian_active`; Attention correlation
was bounded; graceful stop and freshness-bounded forced-termination demotion
were truthful. An isolated shipped-unit start, recurrence, restart and graceful
stop also passed without changing linger or the real QWSG service. Sanitized
details are in `docs/release/ACCEPTANCE_1.0.0-rc.2.md`.

## Rollback

Use only the exact verified snapshot payloads after confirming no Task 040 process/unit is active, repository identity and target hashes match recorded facts, and no later Owner edit exists. Never alter RC.1, private state, LICENSE or unrelated Owner content, and never use broad Git recovery.

## Completion state

`complete`

Task 040 ended with the exact RC.2 archive and checksum sidecar ready for
Owner transfer to clean-host acceptance. Temporary build, extraction, staging,
lifecycle, systemd and Go-cache roots were removed after exact process absence;
the Builder installation and implementation rollback snapshots were retained.
No clean-host acceptance, real-host installation, licensing change, Git stage,
commit, tag, push or publication occurred.
