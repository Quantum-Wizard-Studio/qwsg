# Foundation Milestone Housekeeping Report

## Scope and baseline

This report records the Project Owner-authorized repository housekeeping that closes the Task 004–014 foundation milestone without creating Task 015, changing Guardian runtime behavior, pushing, or tagging. The starting HEAD was `8fa40acd945b5f0d5d1ee0c5e182a19bba092d2b`. The shared uncommitted worktree could not be reconstructed honestly as task-by-task Git history.

## Snapshot and rollback

Before modification, a complete repository snapshot including `.git`, tracked, untracked, and ignored state was written outside the repository at `/tmp/qwsg-foundation-pre-housekeeping-20260722T000000Z/`. `SHA256SUMS` validates the archive and restore documentation. The repository archive SHA-256 is `17f8f3442e3ffd6b9ff8cfa4de6e9ad5802b53710a514cf3dc649a06d7b67b9c`; the archive listing test passed. `RESTORE.md` requires extraction into an empty parent and explicitly forbids overlaying a live worktree.

## Lifecycle anomaly resolution

The former `ai/archive_prompts/008_2026-07-20_core-alpha-slice-1-implementation.md` was a misnamed Task 009 duplicate. Its SHA-256, `901faf8d7fb1c703af46a33abdc231eb29c72c87fb9174d06e089ba172161496`, exactly matched the Task 009 pre-implementation snapshot copy of `009_CURRENT_TASK.md`; its internal ID and slug also identified Task 009. The actual Task 008 prompt was recovered from the verified Task 008 snapshot, where its SHA-256 was `d7769ccda1431d7e3d0175b77172803116ee07ecf3025084b80054bdbaa0e725`, and installed as `ai/archive_prompts/008_2026-07-20_core-alpha-architecture.md`. Its status was synchronized with the official completed Task 008 history. No missing narrative was invented.

`ai/history/009-1_2026-07-20_core-alpha-slice-1-implementation.md` was an unresolved scaffold, not evidence of execution. It is retained with an explicit `never executed — superseded` classification and points to `ai/history/009_2026-07-20_core-alpha-slice-1-implementation.md` as the sole official completed Task 009 history.

Task 014's complete prompt was moved unchanged in substance to `ai/archive_prompts/014_2026-07-22_canonical-system-inventory.md`. Its completed history and embedded Final Delivery Report remain in `ai/history/014_2026-07-21_canonical-system-inventory.md`. `ai/prompts/` is empty and no Task 015 prompt or history exists.

The lifecycle now defines and validates a canonical idle state: zero active prompts is valid only when the highest-numbered archived prompt and its unique history are complete and metadata-consistent. `bin/job --check` and `next-task.sh --check` validate that state. Display modes still require an active task. The task generators can use the latest completed archive as a future numbering baseline, but they do not create a task without an explicit invocation and owner input.

## Backup policy and ignored payloads

`ai/core/15_ENGINEERING_BACKUP_POLICY.md` separates full rollback payloads from publication-reviewed metadata, states honestly that managed external artifact storage is not implemented, and defines manifests, SHA-256, restore documentation, retention, access protection, privacy, deletion approval, early tracked-backup handling, and rollback verification.

`.gitignore` excludes repository-root build output, Go coverage and test binaries, compressed snapshots, copied backup payload areas, raw diffs, host-state, permissions, filesystem trees, and pre-task evidence without using broad rules for release fixtures. Existing backup directories and the milestone snapshot remain outside this foundation commit.

## Publication and privacy review

Foundation audit, history, archive, project, architecture, lifecycle, and testdata documents were checked for absolute host paths, usernames, UID/GID data, infrastructure hostnames/domains, IP addresses, e-mail addresses, credentials, tokens, API keys, and private keys. Historical absolute repository paths were replaced with `<repository-root>`; personal cache paths with `<user-cache>`; build-host and repository ownership identities with role placeholders. Project Owner authority was retained by role rather than personal name.

Publication-safe engineering metadata intentionally retained includes the QWSG project name, repository-relative paths, task identifiers and dates, Git hashes, Go version evidence, the public-facing QWSG domain where technically relevant, architecture decisions, test results, and sanitized inventory fixtures. No credential or secret is required for historical understanding.

## Staging boundary

The controlled foundation staging includes QWSG Go source and tests; CLI and engineering tooling; repository skill; Makefile and module definition; sanitized testdata; canonical engineering, architecture, security, development, user, project, and audit documentation; Task 004–014 official histories; corrected archived prompts through Task 014; lifecycle automation and tests; backup policy; `.gitignore`; README; and CHANGELOG.

Excluded from staging are build output, every repository backup directory from Task 004 onward, the external milestone snapshot, archive payloads, host-state and pre-task evidence, generated/temporary files, the erroneous Task 009 duplicate under a Task 008 filename, secrets/credentials, and any Task 015 artifact.

## Verification record

Pre-staging lifecycle checks and the updated isolated suites passed: Task Builder `36` assertions and next-task lifecycle `28` assertions. `bin/job --check` and `ai/scripts/next-task.sh --check` passed in the canonical idle state. `go test ./...`, `go test -race ./...`, and `go vet ./...` passed with isolated caches under `/tmp`; their first literal invocation could not initialize the sandbox-read-only default user cache and performed no build. `git diff --cached --check` passed after line-ending and whitespace normalization.

No dedicated secret scanner was installed, and none was installed during housekeeping. Added-line scans covered AWS, Google, GitHub, OpenAI-style, Slack, and private-key signatures; credential-like assignments; absolute host paths, usernames, hostnames, IP addresses, and e-mail addresses; and credential-like filenames. No finding remained. Staged-path checks found no backup/build payload, archive, Go test/coverage output, credential file, or Task 015 artifact. The commit hash is reported after commit creation because a commit cannot contain its own hash.
