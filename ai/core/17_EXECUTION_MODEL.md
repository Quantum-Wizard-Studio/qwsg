# Framework 2.0 Execution Model

## Standard Execution Authority

Project Owner approval of a Builder-generated task grants the agent Standard
Execution Authority for the task's stated objective and scope. Without further
permission, the agent may repeatedly inspect, snapshot, diagnose, classify,
edit, refactor, build locally, test, correct, retest, update task-scoped
documentation, use a verified rollback, and perform the task's authorized Git
integration and lifecycle closure.

This authority is bounded by the task's targets and exclusions. It does not
expand product scope, architecture, external-system access, or release
authority. A task envelope records deviations from this standard and any
task-specific external actions; it need not enumerate every ordinary command.

## Default engineering loop

`SNAPSHOT -> INSPECT/REPRODUCE -> DIAGNOSE -> CLASSIFY -> IMPLEMENT -> FOCUSED
TEST -> CORRECT -> PROPORTIONAL RETEST -> INTEGRATE -> CONTINUE`

The loop may repeat within one task. A failed check, failed diagnostic,
incomplete evidence item, or recoverable implementation mistake is evidence to
investigate, not a lifecycle transition and not a reason to create a new task.

Before changing implementation in response to an unexpected result, classify
it as exactly one of:

- `PRODUCT/FRAMEWORK DEFECT`
- `TEST OR ACCEPTANCE DEFECT`
- `ENVIRONMENTAL ISSUE`
- `EXPECTED BEHAVIOR`
- `INCONCLUSIVE`

Compare the observation with the documented product or framework contract.
Correct the component in which the defect actually exists. `INCONCLUSIVE`
authorizes further safe diagnosis inside the existing scope; it becomes a STOP
only when uncertainty cannot be resolved safely within existing authority.

## Boundary-based STOP semantics

Stop and obtain Owner direction only at a genuine boundary:

- credential or secret exposure risk;
- a security or privacy regression or unresolved safety uncertainty;
- irreversible work without a reliable rollback;
- unauthorized destructive external mutation;
- material architecture, objective, or product-scope expansion;
- production deployment or an external infrastructure change not authorized by
  the task;
- tag, Release, publication, or public artifact action not explicitly
  authorized;
- required rollback being unavailable; or
- uncertainty that safe in-scope diagnosis cannot resolve.

Owner interaction may also be inherently required for a private credential,
physical action, or external decision. That interaction is not a new
engineering approval gate; after it, work continues under the same task.

## Proportional validation and evidence

Use the lowest level that reliably answers the engineering question:

1. **Lightweight:** inspection, formatting, narrow static checks, and a concise
   change record for low-risk local work.
2. **Focused:** deterministic regression tests and directly affected checks for
   ordinary implementation changes.
3. **Strong:** full local validation, rollback proof, migration/state safety,
   and integration checks for high-impact or completion work.
4. **Deterministic release:** isolated reproducible construction, provenance,
   byte identities, integrity, and frozen-candidate controls.
5. **External acceptance:** bounded real-environment validation only when it
   materially proves behavior unavailable locally.

Preserve reliable evidence and reuse it until a relevant mutation invalidates
it. Do not repeat checks solely for ceremony. Missing required evidence remains
missing; proportionality never permits manufacturing PASS.

## Development and release separation

The normal path is:

`DEVELOP -> FULL LOCAL VALIDATION -> BUILD FROZEN CANDIDATE -> ONE PRACTICAL
ACCEPTANCE RUN -> RELEASE DECISION`

Development authority does not imply candidate construction, candidate
construction does not imply publication, and practical acceptance does not
authorize release. Candidate bytes are immutable after freeze. A genuine
candidate defect returns to development; a test defect returns to acceptance
tooling; an environmental issue returns to environment diagnosis. A new task
is appropriate only when the objective materially changes, architecture or
scope needs an Owner decision, a separately auditable phase begins, or the
current task is complete.
