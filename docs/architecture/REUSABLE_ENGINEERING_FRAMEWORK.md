# Reusable Engineering Framework

## Status and version

Version `1.0.0` is the first repository-validated reusable engineering
framework. QWSG is its reference implementation and first consumer.

The approved future architecture for its next generation is defined in
`docs/architecture/QWCS_ENGINEERING_OPERATING_SYSTEM.md`. That document does
not change Framework 1.0.0 behavior, rules, validation, lifecycle, or authority;
implementation and migration require separately approved engineering work.

## Boundary

The framework is the project-local system that defines, creates, validates,
executes, records, and closes controlled engineering tasks. It is not the QWSG
runtime and is not a global service or separately published product.

The reusable core consists of:

- mandatory Constitution, lifecycle, task, Git, snapshot, rollback, approval,
  history, and delivery policies;
- deterministic task builder and lifecycle validators;
- project configuration validation;
- transaction and rollback behavior;
- regression and portability fixtures.

Project configuration supplies identity and paths:

- project name and slug;
- repository marker and relative root;
- canonical remote and primary branch;
- communication and engineering-documentation languages;
- required reading;
- validation command argv;
- lifecycle directories, numbering width, and snapshot location.

Project-specific product architecture, terminology, active task content,
history evidence, and delivery artifacts remain outside the reusable core.

## Safety model

Configuration can select project identity and approved commands but cannot
disable explicit owner approval, one-active-task enforcement, snapshots,
rollback, validation, targeted staging, history, or completion evidence. These
invariants remain code and policy requirements.

Configuration is parsed as UTF-8 data and is never sourced. Validation commands
are tab-separated argv fields and execute directly without a shell or `eval`.
Unknown, duplicate, empty, unsafe, or missing configuration fields fail closed.

## Components

| Component | Responsibility |
|---|---|
| `ai/framework/VERSION` | reusable core semantic version |
| `ai/config/engineering-project.conf` | project identity and safe configuration |
| `ai/config/engineering-validations.tsv` | project-specific validation argv |
| `ai/scripts/framework-check.sh` | configuration, identity, path, remote, branch, and command validation |
| `ai/scripts/task-builder.sh` | approved deterministic task creation |
| `ai/scripts/next-task.sh` | compatibility draft preparation and lifecycle consistency |
| `ai/scripts/divert-task-to-test.sh` | owner-authorized incomplete-task containment |
| `bin/job` | installation-free active-task access and validation |

## Compatibility

Task IDs remain three-digit and exact. Existing prompt/history paths and
interactive and `--input-dir` builder contracts remain unchanged. Builder
approval remains exactly `APPROVE`; generated tasks remain UTF-8 LF Markdown.
Historical tasks are not rewritten.

The portability boundary verified by this version is a synthetic temporary Git
fixture using a different project identity, remote URL, branch, required
reading, validation list, and lifecycle directories. It does not prove
cross-platform shell portability or migrate another production repository.
