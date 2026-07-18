# Engineering Task Prompt 001: Product Architecture

## Task Metadata

- Identifier: `001_PRODUCT_ARCHITECTURE`
- Status: `draft — not authorized for execution`
- Engineering language: English
- Owner or lead-developer communication language: verify at execution time
- Human authority: required before execution
- Expected delivery report: assign a chronological path when authorized

## Objective

Define an approved, evidence-based product architecture for Quantum Wizard Server Guardian without implementing application functionality.

## Scope

When explicitly authorized, identify required product boundaries, components, responsibilities, interfaces, data flows, trust boundaries, deployment assumptions, and decision records needed for a maintainable independent QWSG system.

## Out of Scope

- Executing this prompt before explicit human approval
- Application code or prototypes
- Installing dependencies or software
- Modifying operating-system or server configuration
- Deploying services, databases, jobs, or infrastructure
- Treating architectural assumptions as verified facts
- Beginning implementation or a later milestone

## Required Reading

Before any architecture work, read:

1. `ai/core/00_PROJECT_PHILOSOPHY.md`
2. `ai/core/01_CONSTITUTION.md`
3. `ai/core/03_AGENTS.md`
4. `ai/core/08_JOB_TEMPLATE.md`

Also read all then-current architecture, structure, system-map, security, roadmap, project, and prior history documents relevant to the authorized scope.

## Environment Verification

At execution time, record UTC time, user, project root, Git branch and status, relevant file tree, versions, ownership, permissions, ACLs, existing architecture records, and material differences from this prompt's assumptions. Stop and report differences that affect scope.

## Snapshot Requirements

Before changing architecture documentation, create a timestamped snapshot under `ai/backups/` containing the starting state, target documents, permissions, Git status, and a guarded root-checked restore procedure. Verify the snapshot before implementation.

## Risk Assessment

Assess at minimum architectural lock-in, unsafe automation authority, privilege boundaries, secret exposure, server stability, product coupling, deployment portability, maintainability, localization, rollback, and unverified environmental assumptions. Record severity, mitigation, and owner decisions.

## Planned Work

After authorization and verification, propose the smallest documentation-only sequence for architecture discovery, alternatives, decisions, diagrams where useful, security review, owner review, and acceptance. Separate verified constraints from proposals and decisions.

## Rollback Plan

Define exact architecture-document targets and restore them only from the task snapshot after explicit confirmation. Never delete application, server, legacy, or unrelated project data. Verify Git status and restored content afterward.

## Verification Checklist

- Required documents were read and recorded.
- Human authority and preferred communication language were verified.
- Environment and existing architecture state were inspected.
- Snapshot and rollback procedure were validated.
- Architecture remains independent of QUWIP and other Quantum Wizard products.
- Automatic corrective action requires explicit authorization.
- Security, stability, modularity, portability, maintainability, and localization were addressed.
- Alternatives, assumptions, decisions, and unresolved issues are distinguishable.
- No code, dependency, software, service, database, job, or server configuration was created or changed.
- Documentation consistency, ownership, permissions, ACLs, Git diff, and final status pass.

## Documentation Updates

When authorized, identify the minimum architecture, system-map, structure, security, roadmap, project, changelog, and history documents that require updates. Do not pre-authorize changes merely because they are listed here.

## Delivery Report

Create an English chronological engineering report containing summary, changes, reasoning, alternatives, decisions, verification, rollback, unresolved issues, recommendations, and Git record. Provide the owner-facing report separately in the verified preferred language.

## Completion Criteria

The future task is complete only when the authorized architecture documentation and decisions exist, owner review requirements are satisfied, verification passes, rollback remains usable, history is updated, unresolved issues are explicit, and no implementation work has begun.

This file is a future engineering prompt only. Its current creation does not execute or approve Product Architecture, and it will evolve only through authorized engineering work.
