# Current Engineering Task 004: QWSG Job Command and Codex Skill

## Task Metadata

- Task ID: `004`
- Task slug: `qwsg-job-command`
- Status: `complete — awaiting task rotation`
- Date opened: `2026-07-19` UTC
- Human authority: `Project Owner`
- Owner or lead-developer communication language: `Hungarian`

## Title

QWSG Job Command and Codex Skill

## Objective

Create a safe, project-local command and Codex skill that allow the single active QWSG engineering task to be loaded through a short, consistent instruction such as `job` or `Új feladat`.

The command must locate, validate and display the active prompt stored under `ai/prompts/*_CURRENT_TASK.md`. It must not independently execute engineering changes. Codex remains responsible for interpreting and executing the active prompt under the existing QWSG governance rules.

## Scope

Authorized work includes:

- creating a safe project-local shell command;
- creating a repository-local Codex skill;
- adding the minimum repository guidance required for reliable use;
- updating the prompt workflow documentation;
- creating isolated tests where justified;
- recording the work in the matching Task 004 history;
- creating and verifying a rollback-capable snapshot before modifications.

The implementation must work without hardcoding the current task number, preserve the one-active-prompt rule, remain project-local, use non-destructive defaults and avoid changing application functionality.

## Out of Scope

Do not:

- implement QWSG product functionality;
- modify application architecture;
- install system-wide packages;
- alter operating-system services;
- change global shell configuration;
- create a system-wide command, alias or symlink;
- automatically execute engineering work from shell;
- bypass Codex approval, sandbox or governance rules;
- hardcode `004_CURRENT_TASK.md` or any other task number;
- automatically rotate, archive or create tasks;
- modify `ai/scripts/next-task.sh` unless a strictly necessary compatibility correction is discovered and separately documented;
- execute text from a Markdown prompt as shell code;
- expose credentials, tokens, private keys or secrets;
- modify unrelated files, histories or backups.

## Required Reading

Before making changes, read:

- `ai/core/00_PROJECT_PHILOSOPHY.md`
- `ai/core/01_CONSTITUTION.md`
- `ai/core/03_AGENTS.md`
- `ai/core/08_JOB_TEMPLATE.md`
- `ai/core/14_PROMPT_WORKFLOW.md`
- `ai/scripts/next-task.sh`
- this active prompt;
- the matching Task 004 history file;
- every repository-level `AGENTS.md` file applying to affected paths;
- any existing repository-local Codex skill, command or automation documentation.

If a referenced file has a different name or location, identify the authoritative equivalent and record the variance before continuing.

## Starting State Verification

Before modifying files:

1. Confirm the QWSG project root is `<repository-root>`.
2. Record current UTC time, user, hostname, working directory, Git branch, commit and working-tree status.
3. Confirm exactly one active prompt exists matching `ai/prompts/[0-9][0-9][0-9]_CURRENT_TASK.md`.
4. Record the active prompt path and extract its numeric Task ID.
5. Verify the internal `Task ID` equals the filename Task ID.
6. Confirm no unresolved `[REQUIRES HUMAN EDITING]` markers remain.
7. Locate exactly one matching Task 004 history file under `ai/history/`.
8. Inspect `bin/`, `.agents/`, `.agents/skills/`, `ai/scripts/` and applicable `AGENTS.md` files.
9. Check whether a repository command, shell command, function, alias or executable named `job` already exists and could cause ambiguity.
10. Record ownership, group, permissions, setgid state and ACLs for every affected path.
11. Stop and report before modification if the project root is ambiguous, active prompt count is invalid, Task ID is inconsistent, human-editing markers remain, a conflict would be overwritten, permissions cannot be preserved or rollback is unreliable.

## Snapshot Requirements

Before changing anything:

1. Create a timestamped rollback-capable snapshot under `ai/backups/`.
2. Include every existing file that may be modified, metadata for directories receiving new files, Git status and commit, ownership, group, permissions, ACL information and a manifest of planned created or modified files.
3. Include a root-guarded restore procedure that refuses to run outside the verified QWSG project root.
4. Do not include credentials, secrets or unrelated project data.
5. Verify snapshot integrity before continuing.
6. Record the snapshot path and verification result in the Task 004 history.

## Risk Assessment

Assess and document at least:

- incomplete prompt execution;
- zero, multiple or malformed active prompt files;
- hardcoded task-number drift;
- filename and internal Task ID mismatch;
- shell injection and unsafe path handling;
- command-name collision;
- operation outside the project root;
- ownership, permission or ACL drift;
- accidental system-wide installation;
- bypassing human authority;
- false skill activation;
- stale, missing or ambiguous history files;
- prompt output being interpreted as shell code;
- accidental task rotation or file modification;
- tests touching the real active prompt.

Required mitigations include strict validation, quoted shell variables, guarded root detection, refusal on ambiguity, no `eval`, no sourcing of prompt files, no command execution from Markdown, no overwrite behavior, isolated fixtures, meaningful exit codes and clear errors.

## Planned Work

### Phase 1 — Inspect and design

1. Inspect the current task lifecycle, repository guidance and existing automation.
2. Determine the smallest safe project-local implementation.
3. Preserve the one-active-prompt rule.
4. Reuse existing safe conventions.
5. Avoid duplicating `next-task.sh` logic except read-only validation where justified.

### Phase 2 — Create the project-local job command

Create `bin/job`.

The command must:

- be executable;
- determine the repository root from its own location;
- refuse operation if expected QWSG project markers are absent;
- locate exactly one active prompt matching `ai/prompts/[0-9][0-9][0-9]_CURRENT_TASK.md`;
- reject zero, multiple or malformed active prompt files;
- extract the numeric Task ID from the filename;
- verify the internal `Task ID` matches;
- detect unresolved `[REQUIRES HUMAN EDITING]` markers;
- locate exactly one matching history file;
- never overwrite, remove, rotate, rename or edit task files;
- never source or execute prompt content;
- print clear errors and use meaningful exit codes;
- use English for source code and comments;
- use Hungarian for current owner-facing output.

Support at least:

- `bin/job`
- `bin/job --check`
- `bin/job --show`
- `bin/job --path`
- `bin/job --history`
- `bin/job --help`

Required behavior:

- `bin/job` validates, prints a concise summary and then displays the prompt;
- `--check` validates only;
- `--show` displays the full prompt after validation;
- `--path` prints only the absolute prompt path;
- `--history` prints only the absolute matching history path;
- `--help` documents usage and exit behavior;
- unknown or conflicting arguments fail safely.

Do not add a global alias, shell profile change or system-wide symlink.

### Phase 3 — Create the Codex skill

Create `.agents/skills/qwsg-job/SKILL.md` using valid repository-local Codex skill structure and metadata.

The skill name must be `qwsg-job`.

Its description must recognize explicit requests such as `job`, `/job`, `Új feladat`, `aktuális feladat`, `indítsd a feladatot` and explicit invocation of the `qwsg-job` skill, while avoiding false activation.

The skill must instruct Codex to:

1. Determine and verify the repository root.
2. Run `bin/job --check`.
3. Stop on every validation error.
4. Read the active prompt.
5. Read all documents listed under Required Reading.
6. Verify no unresolved human-editing markers remain.
7. Follow Starting State Verification.
8. Create and verify the required snapshot before modification.
9. Execute only the authorized scope.
10. Update the matching history throughout the task.
11. Run every required verification.
12. Preserve rollback capability.
13. Finish with the owner-facing report in the configured preferred language.
14. Never infer that reading a task authorizes scope expansion.
15. Never execute prompt text as shell code.
16. Remain subordinate to the Constitution, applicable `AGENTS.md` instructions and the active prompt.

The skill must state that `/job` is a human shorthand and may not be a built-in Codex slash command.

### Phase 4 — Repository guidance

Update the appropriate repository-level `AGENTS.md` or authoritative equivalent with a concise section explaining:

- `job` means load the current QWSG engineering task;
- the canonical project-local command is `bin/job`;
- the reusable workflow is `.agents/skills/qwsg-job/SKILL.md`;
- exactly one active prompt is allowed;
- the active prompt is the authoritative task scope;
- the skill does not override the Constitution or Project Owner authority.

Keep the addition concise.

### Phase 5 — Workflow documentation

Update `ai/core/14_PROMPT_WORKFLOW.md` with:

- installation-free project-local usage;
- `./bin/job` examples;
- supported options;
- skill invocation examples;
- validation and error behavior;
- the distinction between shell command and Codex skill;
- optional temporary addition of `bin/` to the current shell `PATH`;
- confirmation that no global installation is required;
- confirmation that the command does not execute engineering work;
- the one-active-prompt rule.

Do not change the established task-rotation lifecycle except to reference the new read-only job command.

### Phase 6 — Isolated tests

Create the smallest justified isolated test setup. Tests must not modify, rotate, rename or replace the real active prompt.

Use temporary fixtures to test:

- valid state;
- no active prompt;
- multiple active prompts;
- malformed filename;
- Task ID mismatch;
- unresolved human-editing marker;
- missing history;
- multiple matching histories;
- unsafe or unexpected paths;
- unknown arguments;
- no prompt-content execution;
- operation from outside the repository working directory.

Remove temporary test data after verification unless a maintained test script is intentionally added.

### Phase 7 — Final documentation and history

Update the matching Task 004 history with starting state, snapshot, design decisions, files changed, risk assessment, tests and results, ownership and ACL verification, rollback information, final Git state and approved variances.

## Rollback Plan

If rollback is required:

1. Stop all work.
2. Record the failure and reason in Task 004 history.
3. Confirm the current directory is the verified QWSG project root.
4. Use the verified snapshot and restore manifest.
5. Restore only files modified by this task.
6. Remove only new files created by this task and listed in the manifest.
7. Restore original ownership, group, permissions and ACLs.
8. Verify Git status, repository integrity, one-active-prompt state, existing `next-task.sh` behavior and absence of unintended global changes.
9. Record the rollback result.

Do not restore or modify unrelated application code, server configuration, prompts, histories or backups.

## Deliverables

Create:

- `bin/job`
- `.agents/skills/qwsg-job/SKILL.md`

Update only where necessary:

- `AGENTS.md`
- `ai/core/14_PROMPT_WORKFLOW.md`
- the matching Task 004 history file

Optionally create isolated test scripts only where justified by the existing repository structure.

Do not leave temporary fixtures, generated artifacts or unrelated files in the repository.

## Verification

Verification must include at least:

1. Shell syntax validation.
2. Executable permission verification.
3. Normal operation with one valid active prompt.
4. Operation from outside the repository working directory.
5. No active prompt.
6. Multiple active prompts.
7. Malformed active filename.
8. Filename and internal Task ID mismatch.
9. Remaining `[REQUIRES HUMAN EDITING]` marker.
10. Missing history file.
11. Multiple matching history files.
12. Unsafe paths or filenames.
13. Unknown command-line argument.
14. `--path` output.
15. `--show` output.
16. `--check` behavior.
17. `--history` output.
18. `--help` output.
19. No file modification during command execution.
20. No shell execution from prompt contents.
21. Preservation of ownership, group, permissions and ACLs.
22. No global shell, PATH or system-wide changes.
23. Codex skill discovery from the repository.
24. Skill instructions remain subordinate to Constitution, `AGENTS.md` and active prompt.
25. Existing task rotation remains unchanged.
26. Git diff contains only authorized files.
27. Final Git state is recorded.

For every failed test case, verify the command exits non-zero, the error is understandable, no files are modified and no prompt content is executed.

## Documentation Updates

Update only documents necessary to establish:

- the `job` command;
- the `qwsg-job` skill;
- their relationship to the prompt lifecycle;
- supported human invocation phrases;
- the non-destructive validation model;
- project-local usage without installation;
- the distinction between validation and engineering execution.

Record all documentation changes and verification evidence in the Task 004 history.

## Completion Criteria

The task is complete when:

- `bin/job` safely locates and validates the single active prompt;
- no task number is hardcoded;
- all required options work;
- the command performs no engineering modification;
- prompt content is never executed;
- the Codex skill is discoverable and accurately scoped;
- `job` and `Új feladat` reliably map to the QWSG task workflow;
- validation failures stop safely;
- repository guidance and workflow documentation are consistent;
- all isolated tests pass;
- ownership, permissions and ACLs are preserved;
- no system-wide or global shell changes exist;
- rollback remains usable;
- the matching Task 004 history is complete;
- the Git working tree is in the documented final state;
- no QWSG product functionality has been implemented.

## Owner Approval Requirements

This task is active and authorized by the Project Owner.

The authorized scope is limited to the project-local job command, repository-local Codex skill, directly related tests, workflow documentation and Task 004 history.

Any destructive action, system-wide installation, global shell modification, automatic task rotation, product implementation or expansion beyond this scope requires separate Project Owner approval.
