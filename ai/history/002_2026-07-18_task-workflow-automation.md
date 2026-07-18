# Task History 002: Task Workflow Automation

## Task ID

`002` / Engineering Update `E002`.

## Task title

Task Workflow Automation.

## Date

Opened `2026-07-18` UTC.

## Status

`complete and verified`; Aikó/Codex updated this record during implementation and delivery.

## Starting state

Clean `main` at `fa156697198a2230c938fb6780e9b8e42c860f57`. `ai/prompts/README.md` was governance documentation, and `001_PRODUCT_ARCHITECTURE.md` was an unapproved, unexecuted draft. Archive and workflow-script directories did not exist. Ownership, setgid, and owner/group ACL inheritance matched policy.

## Snapshot location

`ai/backups/20260718T215938Z_E002_task_workflow_automation/`.

## Work performed

- Migrated prompt governance from `ai/prompts/README.md` to `ai/core/14_PROMPT_WORKFLOW.md` and refined it for one-active rotation.
- Archived the Product Architecture draft without execution as `ai/archive_prompts/001_2026-07-18_product-architecture-draft.md`.
- Created required archive and script directories, active Prompt 002, this independent history, and `next-task.sh`.
- Updated relevant structure, engineering, documentation, task-standard, and milestone-index records.
- Added preflight validation, safe slug normalization, no-clobber paths, temporary-file preparation, and failure restoration to the script.
- Ran all required tests in isolated temporary fixtures and removed only E002 test data.

## Files changed

- `ai/core/02_PROJECT_STRUCTURE.md`
- `ai/core/06_ENGINEERING_STANDARDS.md`
- `ai/core/07_ENGINEERING_HISTORY.md`
- `ai/core/08_JOB_TEMPLATE.md`
- `ai/core/10_DOCUMENTATION_POLICY.md`
- `ai/core/14_PROMPT_WORKFLOW.md` (migrated and refined from `ai/prompts/README.md`)
- `ai/archive_prompts/001_2026-07-18_product-architecture-draft.md` (migrated from `ai/prompts/001_PRODUCT_ARCHITECTURE.md`)
- `ai/prompts/002_CURRENT_TASK.md`
- `ai/scripts/next-task.sh`
- `ai/history/002_2026-07-18_task-workflow-automation.md`
- `ai/backups/20260718T215938Z_E002_task_workflow_automation/` snapshot records

## Decisions

- E002 receives sequential task number `002`; E001 remains a governance identifier and is not renumbered.
- The archived Product Architecture prompt retains `archived without execution` status.
- An active prompt carries its own safe task slug so rotation can archive it accurately; the user supplies the next task slug.
- A completed prompt remains the sole file in `ai/prompts/` until the next explicitly authorized rotation; the rotation archives it before any next task begins.
- Initialization derives the next unused number from committed archive and history filenames; normal rotation increments the validated active ID.
- The script prints Hungarian operator output for the current owner's verified preference. The permanent governance rule remains configurable and does not declare Hungarian universal.

## Verification evidence

Final `bash -n` syntax validation passed. Isolated fixtures passed:

- initialization with no active prompt;
- normal rotation and archive creation;
- sequential `001` to `002` increment;
- malformed active filename refusal;
- multiple active Markdown prompt refusal;
- duplicate archive destination refusal without mutation;
- unsafe slug normalization without shell execution;
- refusal when invoked outside the derived project root;
- duplicate next-history no-overwrite refusal;
- real active-prompt hash unchanged across all tests.

Final verification also checks exactly one active Markdown prompt, semantic archive status, required template fields, independent pending-history fields, ownership, group, setgid and ACL inheritance, non-world-writable modes, rollback syntax and baseline, scoped Git changes, excluded application artifacts, and clean final Git status.

## Problems encountered

No repository or product problem occurred. During review, failure cleanup was strengthened so a partial prompt move is reversed if creation of the next active record fails. Test fixtures were removed after successful execution.

## Rollback procedure

From the exact QWSG root, run the snapshot `restore.sh`, review its bounded changes, and type `ROLLBACK-QWSG-E002`. It restores baseline tracked documents and removes only named E002-created paths; it never removes task history generally.

## Git commit hash

E002 engineering commit: `ccae1e51257c1fa3f8588af44366091f454da32c` (`feat: automate engineering task rotation E002`). A separate audit commit records this immutable hash. Final Git status: clean on `main` after the audit commit.

## Open questions

No blocking question. The next task title, slug, authority, and preferred communication language require human input when rotation is authorized. Product Architecture remains unexecuted source material.

## Recommended next task

When the owner authorizes a next task, run `ai/scripts/next-task.sh` from the exact project root, enter the next task slug, edit every required field in the generated prompt, review the pending history record, and obtain explicit approval before implementation. Do not begin Product Architecture merely because its draft is archived.
