# Delivery Report 000: Project Bootstrap

## Purpose and status

This report records delivery of the recoverable documentation and repository foundation for Quantum Wizard Server Guardian. Status: foundation complete and verified on `2026-07-18`; repository-local Git identity is configured and the bootstrap commit is ready to be created.

## Date

- Baseline captured: `2026-07-18T19:29:02Z`
- Completion verification: `2026-07-18` UTC

## Original state

The root contained an empty `ai/` directory, environment-managed `.agents/` and `.codex/` directories, and an initialized Git repository with no commits on `master`. The root was `attila:qwdev`, mode `2775`, with setgid and default ACL inheritance. PHP, Composer, Node.js, and npm were present before this task; none were installed or changed. Git identity was absent.

## Snapshot

The rollback-capable baseline is `ai/backups/20260718T192902Z_project_bootstrap/`. It contains the environment record, original tree, permissions, Git status, and guarded `restore.sh` procedure.

## Created foundation

- Root governance: `README.md`, `LICENSE`, `CHANGELOG.md`, `VERSION`, and `.gitignore`.
- AI records: `ai/README.md`, fourteen core documents, `ai/projects/QWSG.md`, this history report, and the bootstrap snapshot.
- Structural areas: `ai/prompts/`, `docs/` with seven documentation categories, and top-level `installer/`, `agent/`, `console/`, `modules/`, `tests/`, `scripts/`, `tools/`, and `build/`.
- Git's unborn branch was renamed from `master` to `main`.

## Decisions made

- Version is `0.0.1-prealpha`.
- Licensing remains a temporary all-rights-reserved proprietary notice pending owner approval.
- No application architecture or functionality was invented.
- Generated backup archives are ignored, while `ai/` documentation and audit records remain trackable.
- The initial inherited `default:user::r-x` ACL made newly created owner entries non-writable. During bootstrap, owner write was added only where necessary. The owner subsequently corrected project defaults to `default:user::rwx` and retained `default:group::rwx`; setgid inheritance and non-world-writable modes remain preserved.

## Verification results

- All required files and directories are present.
- `VERSION` is exactly `0.0.1-prealpha`; the changelog contains the `2026-07-18` bootstrap entry.
- Required philosophy, constitution, and agent rules were found by content checks.
- `restore.sh` passes `bash -n`, uses exact paths, requires a typed confirmation, checks the expected root, and refuses rollback over later Git history.
- Created content is owned by `attila:qwdev`; directories are mode `2771` with setgid, regular documents are mode `0660`, and the restore script is mode `0770`.
- ACL samples show `default:user::rwx` and `default:group::rwx`. A reversible creation probe produced a directory owned by `attila:qwdev` with mode `2771` and a file with mode `0660`, proving both owner and group write inheritance; the probe was removed afterward.
- No project file is world-writable.
- No Laravel application, application code, dependency, package, database, service, job, unit, or server configuration was installed or modified.

## Rollback procedure

From exactly `/home/qws/web/qwsg.quantumwizard.hu/qwsg`, run `ai/backups/20260718T192902Z_project_bootstrap/restore.sh`, review its stated effects, and type `ROLLBACK-QWSG-BOOTSTRAP`. It removes only explicitly listed bootstrap files and empty directories, restores the unborn `master` branch state (or removes the sole expected bootstrap commit), removes the bootstrap index, and restores `.git/` owner mode. The snapshot and unreachable Git objects remain for audit and recovery.

## Unresolved issues

- Repository-local Git identity is configured by the owner. Exact bootstrap commit hash: **pending creation**.
- The original ACL variance is resolved and verified. Future tasks must continue checking inheritance rather than assuming it remains unchanged.
- Empty structural directories are present in the working tree but cannot be represented by Git without adding meaningful content later.

## Recommended next task

After this bootstrap is committed and its hash recorded, the owner may separately authorize the next task. No architecture or second development task is included in this delivery.

This report belongs to the chronological delivery record; secrets and future work do not. It will remain historical while later reports evolve the project documentation.
