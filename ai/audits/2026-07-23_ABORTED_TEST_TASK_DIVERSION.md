# Aborted Test Task Diversion Delivery Record

## Purpose

This record documents the Project Owner-authorized controlled failure
containment and production-sequence recovery applied to the interrupted Forgejo
production task on `2026-07-23`.

## Authority and starting state

Attila — Project Owner explicitly directed that the incomplete Task 015 attempt
must not be represented as complete and must no longer block the production
Engineering Lifecycle.

The verified starting state was:

- branch `main`;
- HEAD `b4316050436bf8be4062f0e1d4ba7c371c334223`;
- canonical remote
  `https://git.quantumwizard.hu/Quantum_Wizard_Studio/qwsg.git`;
- active production Task 015, slug `forgejo-installation`;
- prompt status `approved`;
- history status `active — production snapshot verified`;
- completion state
  `active — OAuth2 JWT startup gate diagnosed; second scoped repair pending`;
- no Task 016 or Task 017.

The Task 015 history truthfully records failed startup and repair attempts,
pending service validation, and incomplete delivery. No completion claim was
added.

## Snapshot and rollback

The verified pre-change snapshot is:

```text
/tmp/qwsg-aborted-task-diversion-20260723T205834Z
```

It contains repository and lifecycle evidence, the original Task 015 prompt and
history, applicable governance and lifecycle tooling, the referenced repair
script, the verified Task 015 infrastructure snapshot, permissions, a manifest,
SHA-256 checksums, and bounded restoration instructions. Every checksum passed.

The diversion command moves the original prompt and history transactionally.
Failure before or after test-task installation restores both exact active files
and removes only the transaction-owned partial test directory. Failure injection
after prompt move, history move, and final install passed in isolated tests.

## Implementation

The reusable command is:

```text
ai/scripts/divert-task-to-test.sh
```

It requires:

- Project Owner authority;
- a nonempty reason;
- disposition exactly `aborted-test`;
- explicit production-ID release exactly `yes`;
- override token exactly `DIVERT-TO-TEST`.

It rejects complete tasks, missing or ambiguous production lifecycle evidence,
destination collisions, invalid existing test-task records, and malformed
authority or confirmation. Owner disposition prose can be supplied as a
validated regular UTF-8 data file and is never evaluated as shell code.

Diverted records use the independent namespace:

```text
ai/test_tasks/NNN_TEST_TASK/
```

Production validators ignore that namespace. `bin/job --check-test-tasks`
explicitly validates unique contiguous test IDs, original task identity,
authority, reason, incomplete result, production-number release, required
records, and SHA-256 evidence.

The lifecycle documentation now distinguishes failed attempts, rejected
methods, experiments, aborted tests, supersession, and owner deferment. The same
materially unchanged failing method has a default maximum of three attempts.
This does not weaken any production completion or safety gate.

## Applied diversion

The command assigned `001_TEST_TASK` and moved:

```text
ai/prompts/015_CURRENT_TASK.md
  -> ai/test_tasks/001_TEST_TASK/prompt/original-prompt.md

ai/history/015_2026-07-22_forgejo-installation.md
  -> ai/test_tasks/001_TEST_TASK/history/original-history.md
```

The before/after SHA-256 values are identical:

- prompt:
  `616852c086d52d405609cb97c6ddf2d9bbd647148923c84d6e58b49568e4fead`;
- history:
  `37c7d289048a59ad9e4b7d955713c4e613f3e9af2c046566df269601733eeab5`.

The disposition is `aborted-test`, the result is `incomplete`, production
completion is not claimed, and production ID 015 is released.

## Verification

The following passed:

- `ai/tests/test-divert-task-to-test.sh`: 36 assertions;
- `ai/tests/test-next-task.sh`: 28 assertions;
- `ai/tests/test-task-builder.sh`: 36 assertions;
- `make test`;
- `make vet`;
- `make fmt-check`;
- `bin/job --check`;
- `ai/scripts/next-task.sh --check`;
- `bin/job --check-test-tasks`;
- `git diff --check`;
- preserved prompt/history comparison against the external snapshot;
- test-task `SHA256SUMS`.

An isolated builder fixture proved that test-task records do not affect
production numbering and that a clean successor can reuse production ID 015.
No replacement production task was installed in the real repository.

## Delivery result

The production lifecycle is valid and idle at completed Task 014. The next
production ID is 015. Task 016 and Task 017 do not exist. The interrupted Forgejo
attempt remains fully auditable as `001_TEST_TASK`, with incomplete status and
no retroactive success claim.

No task implementation, commit, push, secret handling, or unrelated untracked
file modification occurred.
