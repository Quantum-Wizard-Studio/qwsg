# Engineering Task Lifecycle

## Purpose and authority

This document is the authoritative lifecycle specification for QWSG engineering tasks. It is subordinate to the Project Constitution and human Project Owner, and complements the task content standard in `08_JOB_TEMPLATE.md` and storage rules in `14_PROMPT_WORKFLOW.md`.

## Governing invariants

1. **No Task Without History:** every active task prompt has exactly one matching history record with the same Task ID and slug.
2. A prepared task is not approved, active, or executed. Preparation stops at `READY FOR OWNER REVIEW`.
3. Only the Project Owner can approve a prepared task or expand its authority.
4. Production task numbers are sequential and never reused after a completed,
   superseded, or normally archived production task. An incomplete active task
   may release its number only through the explicit Project Owner-authorized
   aborted-test diversion protocol below; the diverted record retains the
   original identity and proves that production completion was not claimed.
5. Lifecycle transitions are transactional: either prompt archive, next prompt, matching history, and validation all succeed, or the original state is restored.
6. Zero active prompts is valid only as the idle state: the highest-numbered archived prompt and its unique matching history must both be complete and consistent, and no next-task prompt or history may exist.

## Lifecycle states

1. **Creation:** the Engineering Task Builder collects structured owner input and generates one prompt plus one matching history record. The compatibility workflow `next-task.sh` may still prepare an explicitly unapproved draft pair for later owner review.
2. **Owner review:** the owner supplies the substantive objective, scope, exclusions, evidence requirements, authority, communication language, and approval. Tooling validates but does not invent these fields.
3. **Approval:** the owner explicitly approves the structured definition. The builder records that approval while generating the documents; draft preparation alone grants no authority.
4. **Start:** the agent validates the executable task with `bin/job --check`, reads governing material, records the exact starting state, and stops on material variance.
5. **Snapshot:** before target changes, the agent creates and verifies a bounded rollback-capable snapshot.
6. **Implementation:** only approved scope is changed; history is updated throughout.
7. **Verification gates:** all task-mandated checks, rollback validation, documentation consistency, permissions, and Git-state checks must pass truthfully.
8. **History finalization:** the history records starting state, snapshot, decisions, changes, exact verification evidence, limitations, rollback, and completion state. The prompt and history must both be complete before rotation.
9. **Completion:** completion means the authorized objective and every mandatory gate passed. It does not authorize the next task.
10. **Idle closure or next-task generation:** after completion, the prompt may be archived without creating a successor, producing the canonical idle state. When a new task is separately authorized, invoke `./ai/scripts/task-builder.sh` and complete its structured owner workflow; it uses the latest completed archived task as the numbering baseline when idle. If an active completed prompt still exists, it is archived in the same transaction. For a separate review cycle, `./ai/scripts/next-task.sh --prepare --slug <slug>` remains available and stops with an unapproved draft.
11. **Handoff:** a builder-generated task reports `APPROVED AND READY FOR IMPLEMENTATION` but does not execute it. A compatibility draft reports `READY FOR OWNER REVIEW` and remains unapproved and unexecuted.

## Controlled failure containment and production-sequence recovery

Normal production completion gates remain strict. Failed attempts and rejected
methods are recorded inside the active history and do not by themselves change
task identity:

- `attempt-failed`: one execution attempt failed; record its inputs, outputs,
  and evidence before another attempt.
- `method-rejected`: the method is abandoned after evidence shows that retrying
  it materially unchanged is not useful.
- `experimental`: work is isolated from production authority and evidence.
- `superseded`: a separately authorized definition replaces prior intent without
  rewriting its record.
- `owner-deferred`: the Project Owner pauses work without claiming completion.
- `aborted-test`: an incomplete active production task is diverted into the
  independent test-task namespace by explicit Project Owner override.

The same materially unchanged failing method must not be repeated more than
three times by default. A different command label does not make a new method
when commands, inputs, assumptions, and expected outcome remain materially
unchanged. At the limit, record the evidence, mark the method rejected, and use
another approved method. If none remains, stop and request owner deferment or
aborted-test diversion. Critical safety failures always stop immediately.

`ai/scripts/divert-task-to-test.sh` is the only canonical diversion command. It
requires Project Owner authority, a nonempty reason, disposition
`aborted-test`, explicit production-ID release, and the exact
`DIVERT-TO-TEST` override token. It rejects complete tasks. The transaction
moves the unchanged prompt and history into `ai/test_tasks/NNN_TEST_TASK/`,
adds disposition and hash manifests, validates both namespaces, and restores
the original active files on every failure.

Test-task numbering is independent. A diversion truthfully records
`incomplete`, consumes no production number, and permits a new clean task to
reuse the released ID. It never makes failed work complete and never supplies
completion evidence to the replacement task.

## Validation modes

- `next-task.sh --check` validates either current prompt/history identity or the canonical idle state without mutation.
- `bin/job --check` validates an executable active task, or reports a valid idle state after checking the latest completed archived prompt/history pair; active-task display modes still require an active prompt.
- `bin/job --prepared-check` validates a generated draft prompt/history pair while requiring their explicit unapproved states.
- `bin/job --check-test-tasks` audits the separate test-task namespace without
  treating any test task as active production work.

These modes prevent a correct draft from being mistaken for an executable task while preserving strict execution validation.

`task-builder.sh --check-input <directory>` validates structured input without mutation. Installation validates the generated pair before mutation and runs both executable and lifecycle consistency validation after installation. Any failure restores the exact original active prompt and removes only builder-installed targets and temporary files.

## Failure and rollback

Generation checks current completion evidence, structured input, and every destination before mutation. Temporary files are created inside destination directories. The transaction archives the completed prompt, installs the next prompt, installs its matching history, and performs the appropriate validation. An error or tested injected failure removes installed next-task artifacts, restores the archived prompt, and removes temporary files. It never commits, pushes, or starts implementation. Only the builder may record approval, and only after receiving the exact explicit owner approval token defined by its input contract.

Rollback and failure reports must name exact bounded targets. Broad repository resets or cleanups are forbidden. If automatic rollback cannot prove restoration, stop and request owner direction.
