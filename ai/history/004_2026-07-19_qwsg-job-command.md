# Task History 004: QWSG Job Command and Codex Skill

## Task ID

`004`

## Task title

QWSG Job Command and Codex Skill.

## Date

`2026-07-19` UTC

## Status

`complete`

## Starting state

The verified project root was `<repository-root>`; the user was `attila`, the host was `<build-host>`, the branch was `main`, and starting HEAD was `8fa40acd945b5f0d5d1ee0c5e182a19bba092d2b`. Exactly one active prompt existed as `004_CURRENT_TASK.md`, its internal Task ID was `004`, its human-editing marker occurrences were quoted validation requirements rather than unresolved fields, and exactly one matching history existed. No `job` command collision, `bin/job`, repository-level `AGENTS.md`, or prior `qwsg-job` skill existed.

The first execution stopped because the managed workspace exposed `.agents` as read-only. Before resumption, the Project Owner confirmed the permission correction. Resumption verification found `.agents/skills/qwsg-job` writable as `<repository-owner>:<repository-group>`, mode `2775`, with setgid and default ACL entries preserving owner and group write access. The working tree still contained only the authorized Task 003-to-004 rotation and Task 004 snapshot/history state.

## Snapshot location

`ai/backups/20260719T010217Z_task004_qwsg_job_command/`. All nine checksum entries passed both before the initial implementation attempt and after resumption. The guarded restore script passed `bash -n`.

## Work performed

- Read the mandatory core governance, Prompt Workflow, Task 004 prompt and history, `next-task.sh`, applicable repository guidance, and the complete `skill-creator` instructions and UI metadata reference.
- Used the official `skill-creator` initializer in an isolated temporary directory, then created and validated the repository-local skill at the authorized path.
- Created the read-only `bin/job` command with root, prompt, Task ID, human-editing marker, safe history-name, and ambiguity validation.
- Added concise repository agent guidance and documented installation-free command and skill usage.
- Ran normal, failure, immutability, permission, ACL, syntax, skill-structure, and scope verification without modifying the real prompt during tests.

## Files changed

Created:

- `bin/job`
- `.agents/skills/qwsg-job/SKILL.md`
- `.agents/skills/qwsg-job/agents/openai.yaml`
- `AGENTS.md`

Updated:

- `ai/core/14_PROMPT_WORKFLOW.md`
- `ai/prompts/004_CURRENT_TASK.md`
- `ai/history/004_2026-07-19_qwsg-job-command.md`

The existing Task 004 snapshot remains as a rollback record. `ai/scripts/next-task.sh`, application files, product architecture, dependencies, services, the operating system, and global shell configuration were not changed.

## Decisions

- Keep `bin/job` strictly read-only: it validates and displays task data but never executes, rotates, renames, or edits it.
- Resolve the repository from the command's physical location so invocation works from other working directories without accepting an unsafe root override.
- Treat only unquoted `[REQUIRES HUMAN EDITING]` occurrences as unresolved fields; quoted occurrences in validation requirements remain valid documentation.
- Require one safely named, dated history matching the active numeric Task ID.
- Keep `/job` as documented human shorthand rather than claiming it is a built-in Codex slash command.
- Preserve the existing prompt rotation lifecycle and make the skill subordinate to the Constitution, applicable `AGENTS.md`, active prompt, and Project Owner.

## Risk assessment

Strict counts and naming checks mitigate missing, multiple, malformed, stale, and ambiguous task records. Quoted variables, physical root resolution, fixed directory relationships, no `eval`, no `source`, and read-only Markdown handling mitigate path and shell injection risks. Independent fixtures mitigate risk to the real active prompt. Project-local files and explicit governance prevent system-wide installation, authority bypass, and global shell changes. Setgid and ACL verification mitigate collaborative permission drift.

## Verification evidence

- Snapshot checksums (nine files): PASS.
- Snapshot restore script syntax: PASS.
- `bash -n bin/job`: PASS.
- `skill-creator` `quick_validate.py`: PASS (`Skill is valid!`).
- Real `--check`, `--path`, `--history`, `--show`, `--help`, and default behavior: PASS.
- Invocation by absolute path while working in `/tmp`: PASS.
- Twenty isolated checks: PASS. These covered valid/default/show/path/history/help operation, immutable command behavior, no prompt, multiple prompts, malformed prompt filename, Task ID mismatch, unresolved marker, missing history, multiple histories, unsafe history name, unknown and conflicting arguments, non-execution of prompt shell payload, and unexpected project root refusal.
- Every expected failure returned non-zero without executing prompt content or changing task files.
- Skill trigger metadata covers the authorized explicit phrases and excludes generic job discussions; the workflow remains subordinate to repository governance.
- `bin/job` is executable and no world-writable task file was created.
- Ownership and ACL verification: all created files are `<repository-owner>:<repository-group>`; directories retain setgid; files retain owner/group write access; ACL masks retain group write; no created file is world-writable.
- No global alias, PATH persistence, system-wide command, package, dependency, service, product function, or architecture was created.
- `ai/scripts/next-task.sh` content remained unchanged.
- Authorized-file diff and whitespace validation: PASS.

## Problems encountered

The initial run was blocked by a managed read-only `.agents` path. The Project Owner corrected the repository-local path and requested continuation from the existing snapshot and history. Resumption verification confirmed the correction, so no new task or snapshot was created. No implementation or test problem remains.

## Rollback procedure

From the exact project root, run `ai/backups/20260719T010217Z_task004_qwsg_job_command/restore.sh`, review its bounded targets, and type `ROLLBACK-QWSG-TASK-004`. It removes only the Task 004-created command, skill, metadata, and repository guidance; restores the captured Prompt Workflow, Prompt 004, and History 004; verifies its root guard; and preserves the `.agents` permission correction separately authorized by the Project Owner.

## Git commit hash

Not created; Task 004 did not authorize a commit. Starting HEAD remains `8fa40acd945b5f0d5d1ee0c5e182a19bba092d2b`.

## Final Git state

Branch `main` remains at `8fa40acd945b5f0d5d1ee0c5e182a19bba092d2b`. The working tree is intentionally not clean because Task 004 did not authorize a commit. Final `git status --short` records the authorized Task 003-to-004 rotation plus this delivery:

```text
 M ai/core/14_PROMPT_WORKFLOW.md
 D ai/prompts/003_CURRENT_TASK.md
?? .agents/
?? AGENTS.md
?? ai/archive_prompts/003_2026-07-19_product-definition.md
?? ai/backups/20260719T010217Z_task004_qwsg_job_command/
?? ai/history/004_2026-07-19_qwsg-job-command.md
?? ai/prompts/004_CURRENT_TASK.md
?? bin/
```

## Open questions

None for Task 004.

## Recommended next task

Have the Project Owner review the command and skill delivery, then use the established rotation workflow only after separately authorizing the next engineering task. Do not begin Product Architecture or any other milestone as part of Task 004.
