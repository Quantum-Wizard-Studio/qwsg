# QWSG Project Bootstrap Start State

## Purpose and status

This immutable bootstrap record captures the verified state before the Quantum Wizard Server Guardian foundation was created. Status: complete baseline, recorded at `2026-07-18T19:29:02Z`.

## Environment

- Date (UTC): `2026-07-18T19:29:02Z`
- User: `attila` (`uid=1000`, primary group `qwdev`)
- Working directory: `/home/qws/web/qwsg.quantumwizard.hu/qwsg`
- Operating system: GNU/Linux (Ubuntu kernel build)
- Kernel: `Linux 6.8.0-124-generic x86_64`
- Git: `2.43.0`
- PHP: `8.3.31`
- Composer: `2.7.1`
- Node.js: `v24.18.0`
- npm: `12.0.1`

## Original project state

The project contained an empty `ai/` directory, sandbox-owned `.agents/` and `.codex/` directories, and an initialized Git repository with no commits on `master`. The root was owned by `attila:qwdev`, mode `2775`, with setgid and collaborative ACL inheritance. Git user name and email were not configured.

The default ACL has `default:user::r-x`; therefore newly created directories initially lack owner write permission even while their group ACL permits writing. This behavior was discovered while creating the snapshot. The only pre-foundation filesystem change was creation of `ai/backups/`; its owner write bit was then minimally enabled so the required snapshot could be stored.

## Planned creation

The bootstrap will create only documentation, governance files, empty structural directories, a safe rollback record, and Git metadata needed to rename the unborn branch to `main`. It will not install dependencies or implement application functionality. Exact intended paths are defined by the task and enumerated by `restore.sh`.

## Scope

This record belongs to the bootstrap audit trail. Future environment changes do not belong here and must be recorded in later chronological history documents. This record remains fixed; the broader project documentation will evolve during development.
