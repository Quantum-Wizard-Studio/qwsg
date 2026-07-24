# Reusable Engineering Framework 1.0.0 Delivery Report

## Delivery result

Task 015 implemented and verified Reusable Engineering Framework `1.0.0` with
QWSG as the reference implementation. The framework remains project-local and
does not change QWSG runtime behavior or production infrastructure.

Final Project Owner acceptance remains the lifecycle completion gate. No next
task was created or prepared.

## Starting state and snapshot

Implementation started on `main` at
`1f6d4c34a4f4dda761469caed585a9aac24226f7`. The existing `origin/main`
tracking reference matched local `main`; the canonical remote was
`https://git.quantumwizard.hu/Quantum_Wizard_Studio/qwsg.git`; `v0.1.0`
existed. Baseline framework, lifecycle, builder, diversion, Go, vet, and format
tests passed.

The verified external rollback snapshot is
`/tmp/qwsg-task015-framework-20260724T002404Z`. It includes the affected
framework scope, active task records, Git and lifecycle evidence, permissions,
embedded-constant inventory, manifest, exact restore procedure, and SHA-256
checksums. Every checksum passed.

## Architecture and configuration

The reusable core has a semantic version in `ai/framework/VERSION`.
`ai/config/engineering-project.conf` represents project identity, repository
marker, canonical HTTPS remote, primary branch, owner and documentation
languages, required reading, lifecycle directories, three-digit numbering,
snapshot location, and mandatory rollback.

`ai/config/engineering-validations.tsv` represents project-specific validation
commands as tab-separated argv. Configuration and commands are parsed as UTF-8
data and are never sourced, evaluated, or passed through a shell.

`ai/scripts/framework-check.sh` rejects:

- missing, empty, duplicate, unknown, unsafe, non-UTF-8, CRLF, or NUL-bearing
  configuration;
- version mismatch;
- unsafe paths, non-HTTPS remotes, wrong branch or remote, missing project
  marker, required reading, lifecycle directories, or validation executables;
- optional rollback or incompatible task-number width;
- multiple active tasks, missing canonical task sections, unresolved
  placeholders, missing generated approval evidence, and executable broad
  staging instructions outside Out of Scope.

Project configuration cannot disable approval, one-active-task enforcement,
snapshots, rollback, targeted staging, history, validation, or completion
evidence.

## Compatibility

The existing interfaces remain compatible:

- `bin/job`;
- `qwtask` and interactive Task Builder input;
- deterministic `--input-dir`;
- read-only `--check-input`;
- exact `APPROVE`;
- `next-task.sh`;
- aborted-test diversion;
- three-digit Task IDs;
- prompt/history paths;
- transactional installation and rollback.

When `ai/framework/VERSION` exists, live workflow entry points require framework
validation. Legacy isolated lifecycle fixtures without the version marker retain
their bounded compatibility behavior.

Generated tasks now reference canonical lifecycle, prompt workflow, Git policy,
and project configuration documents.

## Canonical Git policy

`ai/core/16_GIT_POLICY.md` is the canonical source for repository verification,
HTTPS synchronization, targeted staging, review, dry-run push, commit/push
evidence, branch/tag protection, untracked-file protection, and bounded
recovery.

For QWSG:

- remote: `origin`;
- URL: `https://git.quantumwizard.hu/Quantum_Wizard_Studio/qwsg.git`;
- branch: `main`;
- HTTPS is canonical;
- Forgejo SSH port 2222 is not required.

The remote fetch dry-run passed after sandbox-approved network access. No
force-push, history rewrite, branch/tag mutation, or unrelated untracked-file
operation occurred.

## Git delivery

Targeted staging contained exactly 29 Task 015 framework, configuration,
policy, documentation, test, active prompt/history, and delivery-record files.
No backup directory, build output, `current-task-job.txt`, or
`current-task-maker.md` was staged.

The implementation commit is:

```text
5ad9bad88209449176a2bd65378b2801c711b7bc
Implement reusable engineering framework v1.0.0
```

`git push --dry-run origin main` reported the expected
`1f6d4c3..5ad9bad` update. The subsequent normal HTTPS push succeeded. Post-push
verification confirmed `HEAD == origin/main`, ahead/behind `0 0`, clean tracked
and staged diffs, and all framework/lifecycle regression gates passing.

The separate delivery-evidence finalization commit is reported in the
owner-facing handoff because a commit cannot contain its own hash.

## Portability evidence

`ai/tests/test-framework.sh` creates a unique temporary Git fixture with a
different project name, slug, branch, HTTPS remote, owner language, required
reading, validation command, and lifecycle paths. It validates the synthetic
identity and direct argv execution, exercises negative configuration cases, and
removes the fixture through its bounded trap.

This demonstrates project-local configuration portability only. It does not
claim cross-platform shell portability, global installation, packaging, or
migration of another production repository.

## Verification evidence

The following passed:

- framework validator and configured validation runner;
- 21 framework and portability assertions;
- 36 aborted-test diversion assertions;
- 28 lifecycle assertions;
- 38 Task Builder assertions;
- `make test`;
- `make vet`;
- `make fmt-check`;
- `bin/job --check`;
- `bin/job --check-test-tasks`;
- `ai/scripts/next-task.sh --check`;
- live Task Builder `--help` and `--check-input`;
- shell syntax checks;
- UTF-8 and LF checks;
- permission and executable-bit checks;
- `git diff --check`;
- no Task 016 or Task 017.

## Rollback

If a mandatory gate later fails, restore only Task 015-modified paths from the
verified snapshot, remove only Task 015-created framework/configuration/test
documentation files, restore captured executable modes, and rerun every
baseline test. Broad Git reset, checkout, restore, clean, wildcard deletion, and
live-worktree overlay are prohibited.

## Known limitations

- Framework version `1.0.0` is validated on the current Bash/Linux environment.
- Configuration remains project-local and uses a fixed three-digit QWSG
  compatibility profile.
- The framework is not packaged, globally installed, or published separately.
- Project-specific validation argv intentionally cannot express shell pipelines,
  redirections, substitutions, or compound commands.
- Final completion and archival require explicit Project Owner acceptance.
