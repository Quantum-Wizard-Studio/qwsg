# Current Engineering Task 005: QWSG Product & System Blueprint

## Task Metadata

- Task ID: `005`
- Task slug: `product-system-blueprint`
- Status: `active`
- Date opened: `2026-07-19` UTC
- Human authority: Project Owner
- Owner or lead-developer communication language: Hungarian
- Engineering documentation language: English

## Title

Quantum Wizard Server Guardian – Product & System Blueprint

## Objective

Create the first official Product & System Blueprint for Quantum Wizard Server Guardian (QWSG).

The single required outcome is a coherent, durable product-level blueprint that transforms the existing planning material into an authoritative reference for future architecture, specification, implementation, validation, packaging, deployment, and roadmap work.

This task is documentation-only. It must define what QWSG is, why it exists, who it serves, what belongs to the product, how its major parts relate at a high level, what the first usable product includes, what is deferred, and which engineering documents must follow.

The task must not implement the product or prematurely decide low-level technical details that belong to later milestones.

## Scope

Authorized work is limited to documentation and documentation-supporting repository inspection.

The task may:

- read the repository governance and engineering documents;
- read all existing QWSG product-planning material available in the repository;
- inventory relevant existing documentation before editing;
- consolidate repeated or overlapping ideas;
- identify and resolve clear contradictions when the repository evidence supports one interpretation;
- preserve unresolved questions explicitly instead of inventing answers;
- create one primary Product & System Blueprint document;
- update only the documentation indexes, engineering history, task history, or cross-references required to register the new blueprint;
- recommend the exact next engineering milestone after completion.

The blueprint must cover at least:

1. Product identity and mission.
2. Origin and problem statement.
3. Product philosophy.
4. Target users and deployment contexts.
5. Product goals and non-goals.
6. Immutable design and engineering principles.
7. Product boundaries.
8. Functional domains.
9. High-level system model.
10. Major product components.
11. Major information and control flows.
12. Agent-only, Installer-assisted, and Console-enabled operating modes.
13. Security and privilege-separation principles.
14. State-transition-based alerting principles.
15. Modularity and extension principles.
16. Configuration, secrets, state, history, logs, and audit responsibilities at a conceptual level.
17. Supported-platform strategy.
18. MVP definition.
19. Post-MVP capability groups.
20. Long-term product direction.
21. Explicitly deferred architectural decisions.
22. Open questions.
23. Recommended follow-up document sequence.

The high-level component model must at minimum consider:

- QWSG Agent;
- QWSG Installer;
- QWSG Console;
- detection and environment inventory;
- scheduler or execution coordination;
- check and module execution;
- state management;
- alert and notification management;
- reporting;
- dependency management;
- configuration management;
- secrets handling;
- storage and history;
- audit logging;
- diagnostics;
- update and uninstall responsibilities.

These items must remain conceptual. Detailed internal architecture belongs to later tasks.

## Out of Scope

The following are forbidden or deferred:

- modifying application or executable source code;
- implementing the Agent;
- implementing the Installer;
- implementing the Console;
- implementing checks or monitoring modules;
- designing detailed module interfaces;
- defining a final database schema;
- defining final REST, gRPC, WebSocket, CLI, or IPC contracts;
- selecting a final web framework;
- selecting a final programming language where the repository has not already decided;
- creating UI mockups or production interface designs;
- installing packages;
- modifying systemd, cron, web server, firewall, user, group, ACL, ownership, or permission configuration;
- changing production infrastructure;
- deploying QWSG;
- changing dependencies;
- committing or pushing changes unless repository rules and explicit owner authorization permit it;
- deleting or replacing historical planning material;
- hiding unresolved questions behind unsupported assumptions;
- expanding the milestone into the detailed Core System Architecture, Functional Specification, Module Architecture, Installer Architecture, Console Architecture, or Implementation Roadmap.

Potential later milestones include:

- Functional Specification;
- Core System Architecture;
- Module Architecture;
- Installer Architecture;
- Console Architecture;
- Security Architecture;
- Data and State Architecture;
- Implementation Roadmap.

## Required Reading

Before work begins, read and follow:

- `AGENTS.md`
- `ai/core/00_PROJECT_PHILOSOPHY.md`
- `ai/core/01_CONSTITUTION.md`
- `ai/core/03_AGENTS.md`
- `ai/core/05_SYSTEM_MAP.md`, if present
- `ai/core/07_ENGINEERING_HISTORY.md`, if present
- `ai/core/08_JOB_TEMPLATE.md`
- `ai/core/09_DELIVERY_POLICY.md`, if present
- `ai/core/10_DOCS_UPDATE.md`, if present
- the active Task 005 prompt;
- the completed Task 003 product-definition history and deliverables;
- the completed Task 004 history and deliverables;
- all repository-held QWSG product-planning notes relevant to the original comprehensive QWSG plan.

If the comprehensive planning material supplied by the Project Owner is not present inside the repository, stop before substantive drafting and report the missing source rather than reconstructing it from guesswork.

## Starting State Verification

Before modifying any file:

1. Confirm the repository root.
2. Run the repository job validation command.
3. Record the active task path and Task ID.
4. Record `git status --short`.
5. Record the current branch and current HEAD.
6. Inventory candidate documentation files relevant to the task.
7. Verify ownership, mode, and ACL state for every directory expected to be modified.
8. Confirm that no runtime or application files need modification.
9. Identify the exact source document containing the comprehensive QWSG planning material.
10. Record any mismatch between the prompt, repository state, governance documents, and available planning material.

If the active prompt is invalid, required planning input is missing, repository rules conflict, or safe write access is unavailable, stop and report the blocker.

## Snapshot Requirements

Before editing:

- create a timestamped snapshot using the repository-approved snapshot mechanism;
- include every existing file that may be modified;
- include the active prompt and relevant governance/index/history documents;
- create a manifest containing original paths, metadata, and checksums;
- verify that the snapshot can be read;
- verify at least one sampled checksum;
- create or reference a bounded restore procedure;
- do not overwrite an earlier snapshot;
- retain the snapshot according to repository policy.

The snapshot must not include secrets or unrelated production data.

## Risk Assessment

Expected risk profile:

- Security risk: low, provided no secrets or production configuration are copied into documentation.
- Stability risk: low, because this is documentation-only.
- Data-loss risk: low after verified snapshot creation.
- Compatibility risk: low.
- Permission and ACL risk: medium until the repository documentation paths are verified, because prior tasks encountered ACL and ownership differences.
- Product-direction risk: medium, because this blueprint will guide later milestones.
- Scope-creep risk: high unless detailed technical design is consciously deferred.
- Rollback risk: low with a verified documentation snapshot.

Aikó must explicitly manage the product-direction and scope-creep risks by distinguishing:

- settled principles;
- current preferred direction;
- deferred decisions;
- open questions;
- illustrative examples.

## Planned Work

Perform the smallest safe sequence below.

### Phase 1 – Validate and inventory

1. Load the `qwsg-job` skill.
2. Validate the active task with `bin/job --check`.
3. Record repository and Git state.
4. Create and verify the required snapshot.
5. Read all required governance and history documents.
6. Locate and inventory all relevant QWSG planning material.
7. Produce a private working outline mapping source material to blueprint sections.

### Phase 2 – Extract authoritative product intent

Classify the planning material into:

- product vision;
- problem statements;
- users and deployment contexts;
- design principles;
- functional domains;
- component concepts;
- MVP candidates;
- future features;
- open questions;
- examples and illustrative ideas.

Do not silently discard unusual or playful ideas. Preserve them in the appropriate future-capability or optional-extension category when they demonstrate product modularity, but do not let them distort the MVP.

### Phase 3 – Resolve structure and boundaries

Define:

- what QWSG is;
- what QWSG is not;
- the relationship between Agent, Installer, and Console;
- which capabilities are mandatory core responsibilities;
- which capabilities are optional modules;
- which capabilities are future extensions;
- which decisions require later architecture work.

Where evidence is insufficient, create an open question rather than inventing a final choice.

### Phase 4 – Draft the blueprint

Create the primary blueprint with a clear hierarchy and a concise executive summary.

The document must distinguish:

- product-level requirements;
- high-level architectural direction;
- illustrative examples;
- deferred implementation details.

The blueprint must be readable by:

- the Project Owner;
- a future technical lead;
- an implementation agent;
- a security reviewer;
- an external contributor.

### Phase 5 – Self-review

Review the draft for:

- contradictions;
- duplicated requirements;
- accidental implementation commitments;
- missing boundaries;
- unclear terminology;
- unsupported claims;
- loss of important original ideas;
- MVP bloat;
- inconsistency with repository philosophy or constitution.

Create a traceability appendix or source-coverage section showing that the major original planning domains were considered.

### Phase 6 – Register and verify

1. Update the appropriate documentation index and cross-references.
2. Create the Task 005 engineering history entry.
3. Record the recommended Task 006 title, slug, objective, and rationale.
4. Run documentation and repository verification.
5. Confirm the Git diff contains only authorized documentation changes.
6. Report all changed files and unresolved questions.

## Blueprint Content Requirements

The primary blueprint should contain a structure equivalent to the following. Headings may be refined, but no major topic may be omitted without documented justification.

1. Document Purpose and Authority
2. Executive Summary
3. Product Origin
4. Problem Statement
5. Mission
6. Vision
7. Product Philosophy
8. Target Users
9. Supported Deployment Contexts
10. Product Goals
11. Explicit Non-Goals
12. Immutable Design Principles
13. Trust, Safety, and Privilege Principles
14. Product Boundaries
15. Functional Domain Map
16. High-Level Product Model
17. Agent
18. Installer
19. Console
20. Shared Core Responsibilities
21. Modules and Extension Model
22. State and Severity Model
23. Alerting and Recovery Philosophy
24. Reporting
25. Configuration and Secrets
26. Data, History, Logs, and Audit
27. Diagnostics and Supportability
28. Installation, Update, and Removal Lifecycle
29. Supported Platform Strategy
30. Operating Profiles
31. MVP Definition
32. Post-MVP Capability Groups
33. Long-Term Direction
34. Security and Privacy Expectations
35. Reliability and Failure-Handling Expectations
36. Compatibility and Migration Expectations
37. Documentation and User-Experience Expectations
38. Open Questions
39. Deferred Decisions
40. Recommended Follow-Up Architecture Documents
41. Glossary
42. Source-Coverage or Traceability Appendix

## Product Principles That Must Be Preserved

The blueprint must preserve and formalize these established principles:

- QWSG is an independent product.
- QWSG is not tied to QUWIP, the Quantum Wizard website, Hestia, or a single VPS.
- The same product should support multiple compatible Linux servers through configuration and detected capabilities.
- Ubuntu and Debian are the initial supported distributions; wider Linux support is a later direction.
- The Agent must remain useful without the web Console.
- The Installer handles privileged bootstrap operations that a browser cannot safely perform.
- The Console provides convenient administration and visualization but must not become a hidden root shell.
- Optional dependencies must not be installed silently.
- Planned changes must be shown before installation or system modification.
- Automatic remediation is opt-in, bounded, observable, and not the default.
- Alerting is based primarily on state transitions, escalation, recovery, and controlled reminders rather than repetitive polling spam.
- Linux memory cache must not be misclassified as dangerous memory usage.
- Monitoring must include actual service outcomes where possible, not only process existence.
- Existing backup systems are monitored before QWSG attempts to become a backup product.
- Sensitive credentials must not be stored as plain text in ordinary configuration files.
- Hardware checks remain capability-dependent and optional, especially on VPS environments.
- Security checks initially report; automatic security repair belongs to a later, separately governed capability.
- Every system modification must be auditable.
- Installation, update, diagnostics, and removal are first-class product lifecycle concerns.
- Modularity must support serious operational modules as well as harmless optional convenience extensions without hard-coding every feature into the core.
- The product should remain approachable for small-server owners while preserving a path toward multi-server and professional use.
- Clear Hungarian-friendly usability may be part of the product identity, while technical architecture and repository documentation remain suitable for international development.

## MVP Framing Requirements

The blueprint must define a realistic first usable release without treating every idea in the source material as mandatory.

The MVP should evaluate and organize at least:

- Ubuntu and Debian support;
- system detection;
- dependency inspection and consent;
- interactive TUI installer;
- disk and inode checks;
- memory and swap checks;
- load checks;
- systemd service checks;
- HTTP/HTTPS checks;
- SSL expiry checks;
- backup age and size checks;
- e-mail alerting;
- state-transition handling;
- recovery notification;
- controlled emergency reminders;
- daily report;
- SQLite or an implementation-neutral local state-store requirement;
- secure standalone Console direction;
- installation, update, removal, diagnostics, and logging.

The blueprint must not finalize contested low-level choices merely because they appeared in early examples.

## Rollback Plan

If the task must be rolled back:

1. Stop all further editing.
2. Record the reason for rollback.
3. Restore only files modified by Task 005 from the verified snapshot.
4. Preserve unrelated concurrent changes.
5. Verify restored checksums where available.
6. run `git status --short`;
7. confirm that the repository returns to the recorded pre-task documentation state;
8. record the rollback in engineering history if repository policy requires it.

No production service restart or system rollback should be necessary because this task is documentation-only.

## Deliverables

Required deliverables:

1. One authoritative QWSG Product & System Blueprint document in the repository’s approved documentation location.
2. Updated documentation index or system map reference.
3. Task 005 engineering history entry.
4. Source-coverage or traceability record, embedded in the blueprint or stored alongside it.
5. A concise list of open questions and deferred decisions.
6. A recommended next task containing:
   - proposed Task ID `006`;
   - proposed slug;
   - title;
   - one-outcome objective;
   - rationale;
   - dependencies on Task 005.

Aikó may recommend splitting the next work into multiple milestones, but must not create or activate a new current task unless the repository workflow and Project Owner explicitly authorize it.

## Verification

At minimum verify:

- `bin/job --check` passes before substantive work;
- the snapshot exists and is readable;
- required source planning material was found and used;
- all mandatory blueprint subjects are covered;
- product principles are represented consistently;
- Agent, Installer, and Console responsibilities are distinct;
- MVP, post-MVP, and long-term features are distinguishable;
- open questions are not presented as settled decisions;
- no executable or runtime files changed;
- no dependencies changed;
- no infrastructure changed;
- no secrets were copied into documentation;
- documentation links and relative references resolve where practical;
- the final Git diff is limited to authorized documentation and task-history files;
- Markdown structure is valid and readable;
- no `[REQUIRES HUMAN EDITING]` placeholders remain in Task 005 deliverables;
- no accidental draft or approval-state workflow is introduced beyond the project’s one-active-task model.

Expected result:

- a stable blueprint suitable as the authoritative input for later architecture tasks;
- no runtime impact;
- a clear, bounded next milestone.

## Documentation Updates

Update only the documentation records needed to make the blueprint discoverable and historically traceable.

Expected updates may include:

- engineering history;
- documentation index;
- system map;
- project README documentation map;
- task history generated by the repository workflow.

Do not rewrite unrelated governance documents.

When updating indexes, preserve existing ordering and conventions unless a documented correction is necessary.

## Completion Criteria

Task 005 is complete only when all of the following are true:

- the starting state and snapshot are recorded;
- the comprehensive QWSG planning material was located and reviewed;
- the authoritative Product & System Blueprint exists;
- the blueprint covers every mandatory content domain or documents a justified omission;
- the product’s identity, boundaries, main components, MVP, future direction, and open decisions are clear;
- the document distinguishes principles from examples and deferred technical choices;
- Agent, Installer, and Console are defined at product level without detailed internal implementation;
- documentation indexes and engineering history are updated;
- verification passes;
- the final diff is documentation-only;
- no destructive, production, dependency, permission, service, or infrastructure change occurred;
- the next recommended engineering task is recorded;
- Aikó provides a concise completion report listing deliverables, verification results, open questions, risks, and the next-task recommendation.

## Owner Approval Requirements

This task is authorized as the active documentation milestone by the Project Owner.

Aikó may begin after the active prompt passes repository validation.

Separate explicit Project Owner approval is required before:

- expanding scope beyond documentation;
- changing executable code;
- changing dependencies;
- altering infrastructure or production configuration;
- making permission, ownership, or ACL changes;
- committing or pushing if not already authorized by repository policy;
- activating or creating the next current task;
- deleting or replacing original planning material;
- treating a deferred architectural choice as final without adequate evidence.
