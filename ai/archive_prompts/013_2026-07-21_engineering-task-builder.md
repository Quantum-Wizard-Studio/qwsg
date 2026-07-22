# Current Engineering Task 013: Engineering Task Builder

## Task Metadata

- Task ID: `013`
- Task slug: `engineering-task-builder`
- Status: `complete`
- Date opened: `2026-07-21` UTC
- Human authority: `Project Owner`
- Owner or lead-developer communication language: `Hungarian`
- Engineering and repository documentation language: `English`

## Title

Design and implement the official Engineering Task Builder for the Quantum Wizard Server Guardian platform.

## Objective

Design and implement the official Engineering Task Builder used by the QWSG Engineering Lifecycle.

The Engineering Task Builder shall replace manual CURRENT_TASK creation with a validated, transactional and deterministic task generation workflow.

The builder shall generate fully approved, lifecycle-compliant CURRENT_TASK documents from structured owner input while preserving repository consistency, metadata integrity and Engineering Lifecycle rules.

The generated task shall require no manual structural editing and shall be immediately ready for implementation.

## Scope

This task includes:

- Design of the Engineering Task Builder workflow.
- Interactive structured owner input.
- Multi-line input support.
- Automatic metadata generation.
- Automatic CURRENT_TASK generation.
- Automatic Owner Approval generation.
- Lifecycle validation before installation.
- Transaction-safe file generation.
- Repository consistency validation.
- Integration with the existing Engineering Lifecycle.
- Architecture documentation.
- Unit tests.

## Out of Scope

This task shall NOT implement:

- graphical user interface
- web interface
- remote task management
- AI-assisted task generation
- project planning
- roadmap management
- task execution
- automatic implementation

Task 013 establishes the Engineering Task Builder only.

## Required Reading

- `ai/core/00_PROJECT_PHILOSOPHY.md`
- `ai/core/01_CONSTITUTION.md`
- `ai/core/03_AGENTS.md`
- `ai/core/08_JOB_TEMPLATE.md`
- `ai/core/11_ENGINEERING_LIFECYCLE.md`

## Starting State Verification

Before implementation:

- Verify Task 012 is complete.
- Verify Engineering Lifecycle integrity.
- Execute bin/job --check.
- Verify repository consistency.
- Verify existing lifecycle scripts.
- Create verified snapshot.
- Record current test status.

## Snapshot Requirements

Create a complete verified snapshot before any modification.

The snapshot shall include:

- lifecycle scripts
- Engineering documentation
- prompt templates
- history handling
- validation logic
- tests

Generate rollback instructions and record the snapshot location in Task History.

## Risk Assessment

Primary risks:

- invalid task metadata
- inconsistent lifecycle state
- broken history generation
- incorrect approval state
- transaction failure
- repository inconsistency

The Engineering Task Builder shall preserve full compatibility with the existing Engineering Lifecycle.


## Planned Work

1. Analyse the current lifecycle implementation.

2. Design the Engineering Task Builder workflow.

3. Design structured owner input.

4. Implement multi-line input mode.

5. Implement automatic metadata generation.

6. Implement CURRENT_TASK generation.

7. Implement automatic approval generation.

8. Implement lifecycle validation.

9. Implement transactional installation.

10. Preserve compatibility with existing lifecycle validation.

11. Update documentation.

12. Execute complete verification.


## Rollback Plan

If verification fails:

- restore verified snapshot
- verify restored hashes
- rerun baseline tests
- document rollback
- leave repository in verified pre-task state

## Deliverables

Mandatory deliverables:

- Engineering Task Builder
- Interactive owner workflow
- Automatic CURRENT_TASK generation
- Automatic metadata generation
- Automatic approval generation
- Lifecycle integration
- Updated Engineering documentation
- Updated Engineering History
- Tests

## Verification

Aiko shall verify:

- bin/job --check succeeds
- generated CURRENT_TASK passes lifecycle validation
- generated metadata is internally consistent
- generated approval state is correct
- no REQUIRES HUMAN EDITING markers remain
- transaction rollback works correctly
- history generation remains compatible
- existing lifecycle remains compatible
- Go tests successful
- git diff --check clean

## Documentation Updates

Update when required:

- ai/core/08_JOB_TEMPLATE.md
- ai/core/11_ENGINEERING_LIFECYCLE.md
- ai/core/07_ENGINEERING_HISTORY.md
- ai/core/13_ROADMAP.md

Create additional documentation describing the Engineering Task Builder workflow.

## Completion Criteria

Task 013 is complete only when:

- Engineering Task Builder exists.
- Structured owner input implemented.
- Multi-line input supported.
- CURRENT_TASK generated automatically.
- Metadata generated automatically.
- Approval generated automatically.
- Lifecycle validation successful.
- Repository consistency preserved.
- Documentation updated.
- Verification successful.
- Snapshot and rollback documented.
- No unresolved lifecycle inconsistencies remain.

## Engineering Principle

> Engineering documents shall never be manually assembled from independent fragments. Every lifecycle document shall be generated from structured, validated input through a deterministic process that guarantees repository consistency.

## Owner Approval Requirements

Approved by the Project Owner.

This task has been reviewed and explicitly approved for implementation.

Implementation may begin in accordance with the Engineering Lifecycle.

Further scope changes require explicit Project Owner approval.
