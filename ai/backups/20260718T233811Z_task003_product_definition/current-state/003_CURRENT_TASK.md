# Current Engineering Task 003: Product Definition

## Task Metadata

- Task ID: `003`
- Task slug: `product-definition`
- Status: `active`
- Date opened: `2026-07-18` UTC
- Human authority: `Approved by Project Owner`
- Owner or lead-developer communication language: `Hungarian`

## Title

QWSG Product Definition

## Objective

Create the official Product Definition document for Quantum Wizard Server Guardian.

This document shall become the single authoritative description of the product.

Its purpose is to define WHAT QWSG is, WHY it exists, WHO it serves, WHAT problems it solves, WHAT principles guide its evolution, and WHAT is intentionally outside the product scope.

This is a product-definition task only.

Do not design implementation details.
Do not define internal architecture.
Do not design APIs.
Do not implement code.

## Scope

Create a documentation-only Product Definition covering at least:

- Product purpose
- Product philosophy
- Core values
- Target users
- User personas
- Problems solved
- Problems intentionally not solved
- Product goals
- Non-goals
- Product editions
- Offline philosophy
- Cloud philosophy
- Privacy principles
- Security principles
- Commercial philosophy
- Free vs Professional positioning
- Relationship between Agent and Console
- Product boundaries
- Long-term evolution
- Guiding engineering principles

The result shall become the foundation for every future architectural decision.

## Out of Scope

Do not:

- design modules
- design APIs
- design databases
- define programming languages
- define frameworks
- implement functionality
- create installers
- define deployment architecture

Those belong to later engineering tasks.

## Required Reading

- `ai/core/00_PROJECT_PHILOSOPHY.md`
- `ai/core/01_CONSTITUTION.md`
- `ai/core/03_AGENTS.md`
- `ai/core/08_JOB_TEMPLATE.md`
- `ai/core/14_PROMPT_WORKFLOW.md`
- All existing Product Definition, roadmap, history and engineering governance documents relevant to this task.

## Starting State Verification

Before any modification:

- Verify project root.
- Verify current Git branch and working tree.
- Record Git status and current commit.
- Verify required engineering documents exist.
- Verify document versions where applicable.
- Verify ownership, permissions and ACLs of files to be modified.
- Verify that Product Architecture has not yet been started.
- Report any inconsistencies before proceeding.

## Snapshot Requirements

Before changing documentation:

- Create a timestamped snapshot under ai/backups/.
- Record Git status.
- Record permissions and ACLs.
- Record affected files.
- Verify snapshot integrity.
- Ensure rollback can be performed without data loss.

## Risk Assessment

Primary risks:

- Mixing Product Definition with Product Architecture.
- Introducing implementation decisions prematurely.
- Defining business strategy as engineering fact.
- Contradicting existing project philosophy.
- Creating documentation that becomes difficult to maintain.

Mitigation:

- Separate verified facts from proposals.
- Clearly distinguish strategic decisions requiring owner approval.
- Do not include implementation or architectural details.
- Update only documents within the approved scope.

## Planned Work

Review all existing project philosophy documents.

Extract already accepted product principles.

Separate verified facts from proposals.

Identify contradictions.

Produce one coherent Product Definition.

When multiple reasonable alternatives exist, document them and provide engineering recommendations instead of making business decisions.

Owner approval is required for strategic decisions.

## Rollback Plan

If the task must be reverted:

- Restore all modified documentation from the task snapshot.
- Verify Git status after restoration.
- Verify document consistency.
- Record rollback in task history.

No application code, configuration or infrastructure may be modified or restored during this task.


## Deliverables

Create:

docs/PRODUCT_DEFINITION.md

Update only documentation required to reference the new Product Definition.

Create the matching engineering history.

No implementation work.

## Verification

Verify that:

The Product Definition:

- is internally consistent
- contains no implementation details
- contains no architecture decisions
- clearly distinguishes facts from proposals
- can serve as the parent document for future architecture
- is understandable by both technical and non-technical stakeholders

## Documentation Updates

Update references where appropriate.

Do not modify unrelated documents.

## Completion Criteria

The task is complete when an owner-reviewable Product Definition exists and no implementation work has begun.

## Owner Approval Requirements

This task must not begin until the owner explicitly approves the edited prompt. Scope expansion and destructive actions require separate approval.
