# Task History 067: Approved Task

## Task metadata

- Task ID: `067`
- Task slug: `deterministic-release-packaging-rc3`
- Status: `complete — private RC.3 ready for final acceptance`
- Date generated: `2026-08-28` UTC
- Human authority: Project Owner: Attila
- Preferred owner communication language: English
- Related prompt: `ai/prompts/067_CURRENT_TASK.md`

## Lifecycle state

The Engineering Task Builder generated and transactionally installed this matching prompt/history pair from validated structured owner input. Explicit approval was recorded, lifecycle checks passed, and execution started under the approved Framework 2.0 envelope.

## Starting state

- Verified root, Framework 2.0.0, canonical HTTPS origin, branch `main`, and fetched `HEAD == origin/main == 90a5815058919d1d913e28a19662672ad5100b01` with `0 0` divergence.
- Before Builder installation the tracked/index worktree was clean and lifecycle was canonically idle after Task 066. After installation only the expected Task 067 prompt/history were untracked.
- Task 066 blocker evidence, RC.2 source/hash, normalized comparison hash, and ambient modes `0775`, `0660`, and `0664` were verified from the canonical history and release decision record.

## Snapshot

- Protected payload root: `/tmp/qwsg-task067-execution.Ha0m3I`, mode `0700`; retain through Owner acceptance and release rollback-window closure.
- The complete tracked HEAD archive and separate Task 067 prompt/history before-images are mode `0600`, listed in a deterministic manifest, and covered by `SHA256SUMS`.
- Checksum validation, archive readability, isolated extraction and exact `scripts/build-release.sh` Git-blob comparison passed. Restore is collision-aware and path-bounded; extraction over the live worktree and broad Git recovery commands are forbidden.

## Work performed

- Read the required governance, Task 066 history, release policy and build documentation. Inspected the builder and existing build/release tests before design.
- Confirmed QWSG-066-F001 as a `PRODUCT/FRAMEWORK DEFECT`: directories and non-executable files are not normalized before manifest/tar creation.
- Defined the canonical package policy: directories `0755`; intended executables `0755`; every other regular file `0644`.
- The first focused regression correctly found that normalizing before manifest creation still allowed the newly generated `MANIFEST.sha256` mode to inherit caller umask. Classified `PRODUCT/FRAMEWORK DEFECT`; moved the complete normalization pass after manifest creation so every final archive entry is covered.
- Added `scripts/test-release-reproducibility.sh` to construct equivalent normalized/`umask 0022` and group-writable/`umask 0002` inputs, compare artifacts byte-for-byte and by SHA-256, and enforce the exact tar mode allowlist. Integrated it into `make release-check`.
- Advanced only the private candidate metadata and canonical installation/release documentation to `1.2.0-rc.3`. RC.2 remains rejected immutable historical evidence.
- Full validation passed, then remediation commit `b6eb357ad03a02b41ac93536fc3be91ecf929803` was created as frozen RC.3 provenance.
- Built private `qwsg-1.2.0-rc.3-linux-amd64.tar.gz` with epoch `1787905053`; SHA-256 is `8543c3e09b48085b01c037d7db5106ea793374dc099b0b9be5f6cacb55af13ee` and sidecar-file SHA-256 is `ae6a9de6d8430ba51d6895dd85219ba2f86422425424f78d31b6ef2952467fdd`.
- Independently rebuilt from a normalized Git export under `umask 0077`; it matched the group-writable worktree build byte-for-byte at the same archive SHA-256.
- Independent extraction under `umask 0002` and `0077` matched recursively. Every directory was `0755`, every ordinary file `0644`, exactly `bin/qwsg`, `install.sh`, and `uninstall.sh` were `0755`, and every archive owner/group was numeric `0/0`.

## Verification

- Builder input/check/install, `bin/job --check`, lifecycle and Framework validation: PASS.
- Mandatory snapshot checksum, readability and isolated restore rehearsal: PASS.
- Focused build contract and release plumbing: PASS.
- Cross-umask (`0022` versus `0002`) artifact byte/SHA-256 identity: PASS.
- Group-writable versus normalized source regular-file and directory modes: PASS.
- Executable preservation and exact three-file executable allowlist: PASS.
- `make engineering-test test vet fmt-check release-check`: PASS; Framework 25, diversion 36, lifecycle 29 and Builder 49 assertions passed.
- `go test -race ./...`: PASS across all packages.
- RC.3 internal manifest, embedded version/full commit/build time, repeat-build identity, numeric ownership, archive modes, two-umask extraction and forbidden writable-mode checks: PASS.
- No `v1.2.0` tag, publication, Forgejo Release, or VPS mutation occurred.
- A post-freeze `release-check` first met its intentional existing-output guard because the private candidate occupied `dist/`; classified `EXPECTED BEHAVIOR`. The archive and sidecar were moved to protected temporary holding, the complete check passed, both were restored, and archive/sidecar verification passed unchanged.
- Targeted staging, staged whitespace/mode/privacy review, remediation commit `b6eb357ad03a02b41ac93536fc3be91ecf929803`, evidence commit `06bfa7addfcb3ba4fa9c1c4e57c6cb3f340fcf65`, push dry-run and clean fast-forward push passed. Fresh fetch proved `HEAD == origin/main == 06bfa7addfcb3ba4fa9c1c4e57c6cb3f340fcf65` with `0 0` divergence before lifecycle archival.

## Rollback

Verify `/tmp/qwsg-task067-execution.Ha0m3I/SHA256SUMS`, follow its collision-safe `ROLLBACK.md`, and restore only exact reviewed task paths from isolated extraction. After integration, use a forward corrective commit; never reset, clean broadly, rewrite history, or alter release refs.

## Completion state

`complete — QWSG-066-F001 corrected; private RC.3 is READY FOR FINAL ACCEPTANCE, while final QWSG 1.2.0 acceptance/publication remains separately gated`
