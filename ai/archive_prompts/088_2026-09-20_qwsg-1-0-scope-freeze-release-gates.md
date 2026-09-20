# Current Engineering Task 088: QWSG 1.0 Scope Freeze & Release Gates

## Task Metadata

- Task ID: `088`
- Task slug: `qwsg-1-0-scope-freeze-release-gates`
- Status: `complete`
- Date opened: `2026-09-20` UTC
- Human authority: Project Owner; explicit Task 088 authorization in the current session, including lifecycle creation, documentation, commit, push and idle closure
- Owner or lead-developer communication language: Hungarian

## Title

QWSG 1.0 Scope Freeze & Release Gates


## Objective

Canonically record the Owner-approved Community 1.0 and Pro 1.0 product maturity boundary. Measure completion by release gates, never speculative future task numbers. New capabilities cannot become 1.0 blockers without an explicit future Owner scope decision.


## Scope

One authoritative docs/PRODUCT_1_0_SCOPE_FREEZE.md containing Community C1-C9 and additional Pro P1-P5 gates, initial statuses, POST-1.0/FUTURE disposition, versioning and scope-change authority. Minimum pointers and direct normative reconciliation in Product Definition, Product Architecture, Product & System Blueprint, Functional Specification and roadmap; milestone index and Task 088 records. Preserve historical records.


## Out of Scope

No runtime/Go changes, new features, broad audit or documentation cleanup, VERSION changes, release construction/signing/index/tag/publication, production operations, or work implementing C2/C3/C7/C9/P1/P3/P4/P5. Do not create or start a successor task. Historical releases and archived records remain immutable.


## Authority Envelope

1. **Task targets and boundaries:** documentation-only Owner-approved scope freeze and direct conflict reconciliation; no runtime implementation or scope expansion.
2. **Permitted external actions:** canonical Git fetch, reviewed task commit, dry-run and fast-forward push, final synchronization verification. No production or publication actions.
3. **Owner-reserved decisions:** future scope expansion, release version/publication, entitlement design and runtime implementation.
4. **Task-specific STOP conditions:** genuine authority/safety boundary or unexpected divergence; do not expand documentation at expense of closure.


## Required Reading

- `ai/core/00_PROJECT_PHILOSOPHY.md`
- `ai/core/01_CONSTITUTION.md`
- `ai/core/03_AGENTS.md`
- `ai/core/08_JOB_TEMPLATE.md`
- `ai/core/11_ENGINEERING_LIFECYCLE.md`
- `ai/core/14_PROMPT_WORKFLOW.md`
- `ai/core/16_GIT_POLICY.md`
- `ai/core/17_EXECUTION_MODEL.md`
- `ai/core/18_BOUNDED_DIAGNOSTIC_RUNNER.md`
- `ai/config/engineering-project.conf`

## Starting State Verification

Verify main, clean HEAD/origin/main d440f20f25f240b824db34216b0e7f34db56f1e6, 0/0 divergence after fetch, valid framework and idle completed Task 087. Verify project language, file permissions/ACLs and relevant documents; reuse the completed planning audit without broad re-investigation.


## Snapshot Requirements

Before lifecycle or documentation mutation, create and verify protected external-to-Git tracked baseline archive and per-member SHA256 manifest. Retain through Owner acceptance. Record sanitized hash and bounded rollback guidance, never payload/private ACLs in Git.


## Risk Assessment

Low runtime risk: documentation only. Main risk is ambiguous precedence, silently adding blockers or claiming unperformed acceptance. Mitigate with one authority, exact Owner-approved statuses and dispositions, direct document pointers, status-count checks and explicit evidence boundaries.


## Planned Work

1. Validate/snapshot and create Task 088 through Builder.
2. Write one canonical scope/gate document. Community initial C1/C4/C5/C6 PASS, C2/C3/C7/C8 PARTIAL, C9 MISSING. Pro inherits all Community gates; P2 PASS, P3/P4 PARTIAL, P1/P5 MISSING.
3. Disposition older broader MUSTs as post-1.0 for release completion; preserve requirements and history. Record Health coverage, optional SMTP, 1.3.1 bootstrap and product-maturity/version distinction.
4. Validate stable documentation, exact gate counts, local links, scope, framework and lifecycle.
5. Finalize history, archive without successor, targeted commit/push and verify clean synchronized idle state.


## Rollback Plan

Verify external archive checksum and all member hashes; extract only to a new empty private directory. Review exact changed docs/core milestone/Task 088 paths against baseline. Restore only bounded paths after drift review and obtain destructive approval for removing new files. Never broad reset/clean or rewrite published history. Preserve task evidence and rerun framework/lifecycle/diff checks.


## Deliverables

Canonical Owner-approved scope document with C1-C9/P1-P5 and statuses, all explicitly deferred families, future evolution, no-silent-expansion rule, immutable release history distinction; minimum direct reconciliation; complete Task 088 history/archive and final commit/synchronization report.


## Verification

Inspect exact final diff, modes and privacy. Confirm only scoped Markdown changes, no old archives/history rewritten, no runtime/version/release data change. Check canonical gate uniqueness/counts/status totals and all Owner dispositions. Validate touched-document precedence and relative links. Run framework-check.sh, bin/job --check, next-task.sh --check and diverted-record validation. Framework 2.0 proportionality: no Go/race/build/release suites for unchanged runtime/tooling. Reverify snapshot, staged diff and final clean synchronized idle state.


## Documentation Updates

Create docs/PRODUCT_1_0_SCOPE_FREEZE.md; minimally reconcile docs/PRODUCT_DEFINITION.md, docs/PRODUCT_ARCHITECTURE.md, docs/PRODUCT_SYSTEM_BLUEPRINT.md, docs/FUNCTIONAL_SPECIFICATION.md and ai/core/13_ROADMAP.md; append concise milestone to ai/core/07_ENGINEERING_HISTORY.md; Task 088 prompt/history only. C8 documentation portion may close, but retain the approved aggregate baseline and do not imply runtime acceptance.


## Completion Criteria

All authorized documentation and focused checks complete; no scope expansion or excluded operations. Archive Task 088 without successor, commit and push reviewed changes, HEAD equals origin/main with 0/0 divergence and clean idle lifecycle. Final Owner report lists only remaining frozen gates. Fit current session; stop optional prose expansion early and prioritize deterministic closure.


## Owner Approval Requirements

Approved by Project Owner; explicit Task 088 authorization in the current session, including lifecycle creation, documentation, commit, push and idle closure through the Engineering Task Builder on 2026-09-20 UTC.

The structured task definition and Authority Envelope have been explicitly approved. Framework 2.0 Standard Execution Authority permits iterative, reversible in-scope engineering without another Owner gate. Further scope changes, exceptional external actions, and Owner-reserved decisions require explicit Project Owner approval.
