# Current Engineering Task 010: Automated Task Lifecycle and Safe Next-Task Preparation

## Task Metadata

- Task ID: `010`
- Task slug: `automated-task-lifecycle`
- Status: `complete`
- Date opened: `2026-07-21` UTC
- Human authority: Project Owner
- Owner or lead-developer communication language: Hungarian
- Engineering and repository documentation language: English

## Title

Automated Task Lifecycle, History Safety, and Transactional Next-Task Preparation

## Objective

Improve the Quantum Engineering Framework so that every successfully completed
engineering task safely prepares the next task without relying on manual creation
of the matching history record.

The existing `ai/scripts/next-task.sh` implementation must be audited and
hardened rather than replaced without justification.

After a task has been fully completed, its history finalized, and all required
verification gates have passed, the executing agent must invoke the
repository-approved next-task preparation workflow.

The workflow must prepare, but must never execute or approve, the next task.

## Governing Principle

A task and its matching history record form one indivisible lifecycle unit.

A task must never exist without exactly one matching history record.

The next-task preparation operation must be transactional:

- either every required change succeeds;
- or the repository is restored to the exact pre-operation state.

No ambiguous or partially rotated task state is permitted.

## Mandatory Starting-State Procedure

Before changing any file:

1. Confirm the repository root and current working directory.
2. Record the current Git branch, HEAD commit, and working-tree status.
3. Inspect the current active prompt.
4. Locate the exactly matching Task 010 history file.
5. Inspect all existing Task 009 and Task 010 prompt, history, and archive files.
6. Inspect the current implementations and contracts of at least:

   - `ai/scripts/next-task.sh`
   - `bin/job`
   - `ai/core/00_PROJECT_PHILOSOPHY.md`
   - `ai/core/01_CONSTITUTION.md`
   - `ai/core/03_AGENTS.md`
   - `ai/core/08_JOB_TEMPLATE.md`
   - every other core document that currently defines task completion,
     history, verification, archival, or handoff behavior

7. Detect repository-specific conventions instead of assuming them.
8. Record all findings in the Task 010 history file before implementation.

Do not overwrite an existing Task 010 history file.

If Task 010 does not have exactly one matching history file, stop and repair
that lifecycle inconsistency before implementation.

## Snapshot Requirements

Before modifying files, create a timestamped Task 010 snapshot beneath the
repository-approved backup location.

The snapshot must include every file that may be modified, together with:

- a manifest;
- integrity hashes;
- original ownership and permissions where relevant;
- a guarded restoration script;
- a `START_STATE.md` record.

The restoration script must restore only the bounded Task 010 changes and must
refuse unsafe execution outside the expected repository.

Validate the restoration script syntax before implementation.

Record the snapshot location and verification evidence in the Task 010 history.

## Existing Implementation to Preserve

The current `ai/scripts/next-task.sh` already contains valuable safety behavior,
including:

- strict Bash execution;
- project-root enforcement;
- repository marker checks;
- prompt and history directory checks;
- active prompt filename validation;
- Task slug validation;
- sequential Task ID generation;
- destination collision detection;
- temporary-file creation;
- EXIT-trap cleanup;
- rollback of prompt rotation on failure;
- no-clobber moves;
- simultaneous prompt and history scaffold creation.

Preserve these properties unless a documented, safer implementation replaces
them.

Do not weaken existing protections.

## Scope

### 1. Lifecycle Constitution

Create:


ai/core/11_ENGINEERING_LIFECYCLE.md

This document must become the authoritative lifecycle specification and define:

task creation;
task approval;
task start;
starting-state recording;
snapshot requirements;
implementation;
verification gates;
history finalization;
completion;
next-task preparation;
owner review;
next-task approval;
failure and rollback behavior.

It must explicitly establish:

No Task Without History.

It must also establish:

Preparing a task does not approve or execute that task.

2. Core Documentation Integration

Update the appropriate existing authoritative core documents so that they refer
to the new lifecycle specification.

At minimum, evaluate whether updates are required in:

ai/core/01_CONSTITUTION.md
ai/core/03_AGENTS.md
ai/core/08_JOB_TEMPLATE.md
repository README or engineering entry-point documentation

Avoid copying the entire lifecycle rule into multiple documents.

Use one authoritative source and concise cross-references where possible.

3. Harden next-task.sh

Extend ai/scripts/next-task.sh into a safe next-task preparation tool.

The implementation must retain backward compatibility with its current
interactive workflow unless a documented repository requirement prevents it.

The tool must support non-interactive agent execution so that task completion
does not depend on an interactive shell alias.

The shell alias:

qwtask

currently resolves to:

<repository-root>/ai/scripts/next-task.sh

However, automated agents must be able to invoke the repository script directly
from the project root.

Design and implement clear command modes. The exact option names may be adapted
to existing repository conventions, but the behavior must include equivalents
of:

./ai/scripts/next-task.sh --check

for read-only lifecycle validation, and:

./ai/scripts/next-task.sh --prepare --slug <next-task-slug>

for non-interactive transactional next-task preparation.

No-argument interactive use should remain available for the human owner where
safe and practical.

4. Completion Preconditions

Before rotating to the next task, the tool must verify the current task state.

At minimum, it must verify:

exactly one valid active prompt exists;
the active prompt contains exactly one valid Task ID;
the active prompt contains exactly one valid Task slug;
exactly one matching history record exists for the active Task ID;
the matching history belongs to the same Task;
the current history is finalized and does not remain pending,
draft, or task not started;
required completion evidence is present according to repository rules;
destination prompt, history, and archive paths do not conflict;
Task numbering is sequential and unambiguous;
no temporary or partially rotated lifecycle state already exists.

Do not trust filenames alone when the documents contain authoritative Task ID
metadata.

5. Verification-Gate Integration

Inspect the current bin/job behavior before deciding how it participates in
task completion and post-rotation validation.

The final design must avoid a circular failure such as:

the completed Task passes validation;
the next draft Task is created;
validation rejects the repository merely because the new Task is correctly
awaiting human editing and approval.

Define and document separate semantics where necessary for:

validating a completed current task;
validating a newly prepared draft task;
validating an executable approved task.

Modify bin/job only if required by the accepted lifecycle design.

Any modification must remain backward compatible where reasonably possible and
must be covered by tests.

6. Transaction Safety

The complete rotation must be all-or-nothing.

Transaction scope includes:

archiving the completed active prompt;
creating the next active prompt;
creating exactly one matching next-task history scaffold;
performing required post-rotation validation.

If any operation fails, restore the original active-task state.

The rollback must also remove a newly created history file if later steps fail.

The current implementation appears to track rollback of the active prompt and
new prompt, but the complete behavior must be audited carefully, including
whether a newly installed history file is removed after a later failure.

Do not leave:

an archived old prompt without a new active prompt;
a new prompt without history;
a new history without prompt;
multiple active prompts;
multiple histories for one Task ID;
stale temporary files.
7. History Scaffold

The newly prepared history file must clearly state that the task:

has been prepared;
has not started;
is not approved;
contains no implementation evidence yet.

It must contain the same Task ID and slug as the new active prompt.

Include machine-verifiable metadata if this is consistent with existing
repository conventions.

8. Human Authority Boundary

The tool must never:

invent the substantive objective of the next task;
approve the next task;
begin implementation;
expand scope;
make destructive project changes unrelated to rotation;
commit or push Git changes.

After preparation it must stop in a clearly reported:

READY FOR OWNER REVIEW

state.

9. Completion Output

Successful preparation should report, in a concise and readable format:

completed Task ID;
archived prompt path;
new Task ID;
new prompt path;
new history path;
lifecycle validation result;
owner-review status.

The output must make failure states equally clear.

Test Requirements

Create repository-appropriate automated tests for the lifecycle script.

At minimum cover:

successful interactive preparation;
successful non-interactive preparation;
read-only --check;
missing active prompt;
malformed active prompt filename;
invalid or missing Task slug;
missing matching history;
duplicate matching histories;
history still pending;
conflicting archive destination;
conflicting next prompt;
conflicting next history;
malformed Task ID metadata;
prompt/history Task ID mismatch;
interruption or injected failure after archiving;
failure after installing the new prompt;
failure after installing the new history;
complete rollback to the original state;
no orphan temporary files;
correct generation of matching prompt and history metadata.

Tests must use isolated temporary fixtures and must not rotate the real
repository during test execution.

Required Verification

Run all repository-relevant checks, including at minimum:

bash -n ai/scripts/next-task.sh

and, where applicable:

shellcheck ai/scripts/next-task.sh

If ShellCheck is unavailable, record that fact without installing packages
unless explicitly authorized.

Also run:

git diff --check

Run all lifecycle-script tests.

Run all existing framework validation.

Run the complete current Go validation suite:

gofmt
go vet ./...
go test ./...
go test -race ./...

Build the QWSG binary.

Run any existing JSON fixture validation, snapshot integrity validation,
restore-script syntax validation, documentation-link validation, and
bin/job validation required by the repository.

Do not claim a green verification that was not actually executed.

Documentation and History

Continuously update the existing Task 010 history record.

The final history must include:

exact starting state;
snapshot path;
files changed;
implementation decisions;
lifecycle semantics;
command interfaces;
tests added;
verification commands and results;
failures encountered;
rollback procedure;
open questions;
recommended next task.

The history status must be finalized before next-task preparation is invoked.

Final Lifecycle Demonstration

Only after:

implementation is complete;
documentation is complete;
Task 010 history is finalized;
all mandatory checks are green;
no unresolved verification gate remains;

perform one real next-task preparation operation.

Prepare Task 011 with the slug:

core-alpha-platform-hardening

Use the new non-interactive repository command directly rather than relying on
the interactive shell alias.

Expected semantic operation:

./ai/scripts/next-task.sh --prepare --slug core-alpha-platform-hardening

Adapt the exact syntax only if the implemented, documented interface differs.

This final operation must:

archive the completed Task 010 active prompt;
preserve the finalized Task 010 history;
create exactly one Task 011 active prompt scaffold;
create exactly one matching Task 011 history scaffold;
validate the prepared lifecycle state;
stop without executing Task 011.

After rotation, verify and report that the repository is:

READY FOR OWNER REVIEW
Out of Scope
Implementing Task 011 platform hardening.
Starting monitoring, daemon, remediation, or web-console work.
Changing QWSG inventory behavior.
Installing new operating-system packages without explicit approval.
Git commit or push.
Automatically approving or executing Task 011.
Moving the script to a new reusable cross-project framework location in this
task unless strictly required for correctness.

A future task may extract the lifecycle tooling into a reusable Quantum
Engineering Framework component after this repository-local implementation is
proven stable.

Deliverables
ai/core/11_ENGINEERING_LIFECYCLE.md
hardened ai/scripts/next-task.sh
lifecycle integration updates in existing authoritative documents
bin/job changes only if justified
isolated automated tests
updated and finalized Task 010 history
verified Task 010 snapshot and rollback
archived Task 010 prompt after successful completion
prepared ai/prompts/011_CURRENT_TASK.md
exactly one matching Task 011 history scaffold
final lifecycle validation report
Completion Criteria

Task 010 is complete only when:

the lifecycle constitution exists and is integrated;
next-task.sh has safe read-only and non-interactive preparation behavior;
no Task can be prepared without a matching history scaffold;
an unfinished current Task cannot be rotated;
prompt and history metadata are cross-validated;
the entire operation is transaction-safe;
failure-injection tests prove rollback;
all mandatory checks pass;
Task 010 history is finalized;
Task 011 is prepared with slug core-alpha-platform-hardening;
Task 011 has exactly one prompt and exactly one matching pending history;
Task 011 remains unapproved and unexecuted;
the final state is reported as READY FOR OWNER REVIEW.
Owner Approval Requirements

This Task 010 prompt is approved for execution after the Project Owner confirms
the generated Task 010 prompt and history files exist and the prompt has been
placed into ai/prompts/010_CURRENT_TASK.md.

Task 011 must not begin without separate explicit Project Owner approval.

Do not commit or push.
