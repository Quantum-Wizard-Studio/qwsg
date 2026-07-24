# Reusable Engineering Framework Adoption Guide

## Supported adoption boundary

Version `1.0.0` supports controlled project-local adoption. It is not a global
installer or published package. Start in an isolated temporary or new
repository; never experiment in another production repository.

## Procedure

1. Create a rollback snapshot and verify repository ownership and authority.
2. Copy the reusable policy, framework, configuration, script, and test
   components while preserving executable modes.
3. Populate `engineering-project.conf` from verified project facts.
4. Define validation commands as tab-separated argv, without shell fragments.
5. Keep mandatory approval, single-active-task, snapshot, rollback, targeted
   staging, history, and completion gates unchanged.
6. Run `framework-check.sh` before enabling builder or job workflows.
7. Use a synthetic completed baseline to test task numbering, task creation,
   rollback, and lifecycle checks.
8. Confirm the fixture cannot access or modify QWSG lifecycle data.
9. Review all project-name, path, remote, branch, and language references.
10. Record limitations before declaring adoption operational.

## Prohibited shortcuts

- Do not source configuration.
- Do not use `eval` or shell-form validation commands.
- Do not weaken mandatory policies through project fields.
- Do not copy active QWSG prompts or histories as another project's evidence.
- Do not claim portability beyond tested operating systems and shells.
- Do not install globally or migrate a production project without separate
  Project Owner authority.

## Upgrade

Update the framework version and configuration version together. Run old and new
regression suites, validate the project configuration, document compatibility
decisions, and retain a bounded restore path. Historical records remain
immutable.
