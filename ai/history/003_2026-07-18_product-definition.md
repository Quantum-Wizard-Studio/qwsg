# Task History 003: QWSG Product Definition

## Task ID

`003`

## Task title

QWSG Product Definition.

## Date

`2026-07-18` UTC

## Status

`complete and verified — owner review required for strategic proposals`.

## Starting state

The repository was on `main` at `4b4a0ac97fe3ac21ef1fae4beec70870cd53cad5` with an uncommitted but semantically valid prompt rotation: Prompt 002 was removed from the active directory, its archive was untracked, and Prompt 003 plus History 003 were untracked. Prompt 003 and History 003 were mode `0600`, contrary to expected group-write collaboration. No Product Definition existed. Architecture governance stated that no application architecture was approved or implemented, and the earlier Product Architecture prompt remained archived without execution.

## Snapshot location

`ai/backups/20260718T233811Z_task003_product_definition/`. It contains the exact starting Git and permission state, affected targets, copies of all uncommitted rotation records, and a guarded rollback script.

## Work performed

- Read all mandatory governance and relevant philosophy, project, roadmap, architecture-status, security, license, changelog, prompt, and history records.
- Verified that Product Architecture had not started.
- Restored group write only on the new Prompt 003 and History 003 files.
- Corrected the generated history filename from `003_2026-07-18_product-architecture.md` to `003_2026-07-18_product-definition.md` to match the approved task slug.
- Created `docs/PRODUCT_DEFINITION.md` as the consolidated product-level parent document.
- Separated established governance constraints from strategic proposals and open owner decisions.
- Added concise references in the root README, QWSG project record, roadmap, and Engineering History.
- Marked Prompt 003 complete while retaining owner review requirements.

## Files changed

- `docs/PRODUCT_DEFINITION.md` (created)
- `README.md`
- `ai/projects/QWSG.md`
- `ai/core/07_ENGINEERING_HISTORY.md`
- `ai/core/13_ROADMAP.md`
- `ai/prompts/003_CURRENT_TASK.md`
- `ai/history/003_2026-07-18_product-definition.md` (renamed and completed)
- `ai/backups/20260718T233811Z_task003_product_definition/` (created)

The pre-task prompt rotation also remains visible in Git: Prompt 002 moved to its permanent archive, and Prompt 003 plus its history entered the active workflow.

## Decisions

- Existing approved governance is labeled `Established` and may guide later work.
- Target users, personas, offline promise, optional cloud role, privacy commitments, Agent/Console product roles, editions, and commercial positioning are labeled `Proposed — owner approval required` or `Open decision` because no prior owner decision was found.
- “Free” and “Professional” are working positioning names only and grant no license or distribution rights.
- The Product Definition describes Agent and Console responsibilities only at product level and explicitly avoids topology, process, packaging, protocol, or deployment decisions.
- Product Architecture remains a separate, unstarted task.

## Verification evidence

Verification results before commit:

- All 26 required Product Definition sections are present.
- `Established`, `Proposed — owner approval required`, and `Open decision` labels are present and used to separate evidence from strategy.
- Implementation and architecture terms occur only in explicit exclusions or product/architecture boundary statements; no module, API, database, language, framework, topology, protocol, package, or deployment decision was made.
- `04_ARCHITECTURE.md` remains foundation-only, and the earlier Product Architecture prompt remains archived without execution.
- README, project record, roadmap, Engineering History, active prompt, and independent history references resolve consistently.
- The snapshot `SHA256SUMS` check passed for all eight recorded files; rollback syntax and baseline commit validation passed.
- Product Definition, Prompt 003, and History 003 are `<repository-owner>:<repository-group>`, mode `0660`; rollback is mode `0770`; directory ACLs retain owner/group write and setgid inheritance; no checked file is world-writable.
- No application manifest, dependency directory, installer, service, database, code, or server configuration was created.
- Exactly one active Markdown prompt exists: `003_CURRENT_TASK.md`.
- Markdown whitespace and final Git checks are rerun at commit and delivery.

## Problems encountered

- The task began after a valid prompt rotation that had not yet been committed; the snapshot preserves this state explicitly.
- The generated Prompt 003 and History 003 modes were `0600`, revealing that `next-task.sh` temporary-file permissions do not preserve group write. Task 003 corrected only its two files; the script itself is outside this product-definition scope.
- The generated history filename used the prior `product-architecture` slug even though Prompt 003 had been edited to `product-definition`; it was corrected and recorded.
- Commercial and edition decisions lacked approved source evidence, so they remain proposals rather than asserted facts.

## Rollback procedure

From the exact project root, run `ai/backups/20260718T233811Z_task003_product_definition/restore.sh`, review its named targets, and type `ROLLBACK-QWSG-TASK-003`. It removes only the Product Definition and corrected history path, restores four tracked references from the verified baseline commit, and restores the captured pre-task Prompt 003 and history files. It does not touch application, architecture, configuration, infrastructure, or the pre-existing Prompt 002 archive.

## Git commit hash

Task 003 engineering commit: `b165f288720802695e0c5d5622515855d5c31be9` (`docs: define QWSG product scope task 003`). A separate audit commit records this immutable hash. Final Git status: clean on `main` after the audit commit.

## Open questions

- Will the owner approve the proposed target personas?
- Must fundamental QWSG protection remain useful without mandatory vendor-cloud connectivity?
- Should “Free” and “Professional” become the official edition names?
- Which capabilities, licensing terms, prices, support commitments, and distribution rights belong to each edition?
- Does the owner approve the proposed privacy, cloud, commercial, and Agent/Console product-role principles?

## Recommended next task

Owner review and resolution of the Product Definition decision register. Do not begin Product Architecture until the owner explicitly approves the relevant product decisions and authorizes a separate task.
