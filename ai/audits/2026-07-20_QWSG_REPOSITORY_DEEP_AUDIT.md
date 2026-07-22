# QWSG Repository Deep Audit

## 1. Executive Summary

**VERIFIED:** QWSG is a governance- and specification-complete pre-alpha repository, not an implemented server-guardian product. The `agent/`, `installer/`, `console/`, `modules/`, `tests/`, `scripts/`, `tools/`, and `build/` areas contain zero files. The only executable implementation is engineering workflow automation: `bin/job` and `ai/scripts/next-task.sh`.

**DOCUMENTED:** Product intent, a product/system blueprint, and 125 uniquely identified functional requirements exist. **VERIFIED:** detailed application architecture, security architecture, data contracts, dependency manifests, package/build configuration, product tests, and product runtime code do not.

**INFERRED:** Core Alpha implementation should not begin directly. A bounded architecture-and-contract milestone is required first. The proposed narrow first vertical slice can then begin if it is explicitly described as an implementation slice rather than a complete Core Alpha release.

## 2. Audit Method

The audit read the active Task 007 prompt, repository instructions, all mandatory governance and product documents, prior prompts and histories, the original master plan, scripts, backup metadata, and relevant root files. It inventoried the filesystem, permissions, ACLs, Git state, runtimes, manifests, entry points, and tests. It ran only read-only or static checks. Evidence labels are `VERIFIED`, `DOCUMENTED`, `INFERRED`, `PROPOSED`, and `UNKNOWN`.

Snapshot baseline: `ai/backups/20260720T183612Z_task007_repository_deep_audit/`. All manifest checksums passed; `restore.sh` passed `bash -n` and is executable. An earlier Task 007 snapshot at `20260720T170841Z` was retained as historical evidence but rejected as the current rollback baseline because it predates Task 006 completion and would restore unrelated lifecycle files.

## 3. Verified Repository State

- **VERIFIED:** root `<repository-root>`; branch `main`; HEAD `8fa40acd945b5f0d5d1ee0c5e182a19bba092d2b`.
- **VERIFIED:** the working tree was already dirty with authorized, uncommitted Task 003–007 records and documentation. These changes were preserved.
- **VERIFIED:** version `0.0.1-prealpha`; temporary proprietary license; no functional release claim.
- **VERIFIED:** one active prompt exists: `ai/prompts/007_CURRENT_TASK.md`; `bin/job --check` validates it and its matching history.
- **VERIFIED:** top-level product and engineering placeholder directories are setgid and group-writable by ACL. Task histories 005–007 are mode `0600`, inconsistent with collaborative group access documented in earlier history.
- **VERIFIED:** no dependency installation, service change, network access, production mutation, or rollback execution occurred.

## 4. Component Inventory

| Component | Path | Files | Language/runtime | Entry/config/test artifacts | Maturity | Evidence and missing elements |
| --- | --- | ---: | --- | --- | --- | --- |
| Agent | `agent/` | 0 | None | None | PLACEHOLDER | Directory exists; no source, entry point, configuration, or tests. |
| Installer | `installer/` | 0 | None | None | PLACEHOLDER | Directory exists; no lifecycle implementation. |
| Console | `console/` | 0 | None | None | PLACEHOLDER | Directory exists; no web application or assets. |
| Modules | `modules/` | 0 | None | None | PLACEHOLDER | No checks or module contract implementation. |
| Product CLI | planned `qwsg` | 0 | None selected | None | DOCUMENTED | `docs/FUNCTIONAL_SPECIFICATION.md` defines behavior; no executable exists. |
| Job command | `bin/job` | 1 | Bash 5.2-compatible | `bin/job` | TESTED | Read-only validation/options work; Task 004 records isolated negative tests. It is engineering tooling, not product CLI. |
| Task automation | `ai/scripts/next-task.sh` | 1 | Bash | interactive entry point | TESTED | Static syntax passed; prior Task 002 documents isolated lifecycle tests. Not executed here because it mutates prompts. |
| Tests | `tests/` | 0 | None | None | PLACEHOLDER | No maintained product or workflow test files. |
| Build system | `build/` | 0 | None | None | PLACEHOLDER | No Makefile, pipeline, build manifest, or artifacts. |
| Packaging | repository-wide | 0 | None | None | ABSENT | No Debian package metadata, archive recipe, container file, or signed distribution. |
| Configuration | repository-wide | 0 runtime files | None selected | Only documented model | DOCUMENTED | Syntax, schema representation, storage, and migration are undecided. |
| Logging | repository-wide | 0 runtime files | None selected | Only requirements | DOCUMENTED | No format, sink, rotation, retention, or implementation. |
| State storage | repository-wide | 0 | None selected | SQLite only a preferred candidate | DOCUMENTED | No schema, migration, durability, locking, or corruption strategy design. |
| Monitoring | `agent/`, `modules/` | 0 | None | Requirements only | DOCUMENTED | Required checks are specified but absent. |
| Notifications | repository-wide | 0 | None | E-mail required by spec | DOCUMENTED | Transport selection and implementation are release gates. |
| Service integration | repository-wide | 0 | None | systemd behavior documented | DOCUMENTED | No service units or privilege model. |
| Update | repository-wide | 0 | None | lifecycle requirements only | DOCUMENTED | Integrity/authenticity mechanism missing. |
| Uninstall | repository-wide | 0 | None | lifecycle requirements only | DOCUMENTED | Ownership manifest and retention/export design missing. |
| Rollback support | `ai/backups/` | 10 snapshot directories | Bash restore scripts | 10 scripts; 6 checksum sets | PARTIAL | Engineering-task rollback exists but formats and cross-task safety are inconsistent. No product rollback exists. |

## 5. Runtime and Dependency Inventory

**VERIFIED host tools:** Bash 5.2.21, Git 2.43.0, Python 3.12.3, PHP 8.3.31, Node.js 24.18.0, npm 12.0.1, Composer 2.7.1. Go, Rust, SQLite CLI, and ShellCheck were not detected.

**VERIFIED repository dependencies:** none. There is no `composer.json`, lock file, `package.json`, npm lock file, Python manifest, `go.mod`, `Cargo.toml`, Dockerfile, systemd unit, CI configuration, or Makefile. The detected PHP, Node, Python, npm, and Composer versions are environmental facts, not QWSG dependencies. **UNKNOWN:** future production runtime, framework, database, and dependency requirements because architecture deliberately defers them.

## 6. Existing Implementation Status

**VERIFIED:** meaningful product implementation does not exist. `README.md`, `ai/core/02_PROJECT_STRUCTURE.md`, `ai/core/04_ARCHITECTURE.md`, and the empty product directories agree on this point. `bin/job` is a functioning project-development command; it must not be mistaken for the specified `qwsg` product CLI. `next-task.sh` is functioning workflow automation documented by prior history, but no maintained regression suite is present.

## 7. Test and Build Results

| Command/check | Exit | Result |
| --- | ---: | --- |
| `bin/job --check` | 0 | Active Task 007 and matching history validated. |
| `bin/job --path`, `--history`, `--help` | 0 | Documented modes worked. |
| `bin/job --bad-option` | 2 | Correctly rejected an unknown option with an understandable error. |
| `bash -n bin/job` | 0 | Syntax valid. |
| `bash -n ai/scripts/next-task.sh` | 0 | Syntax valid; not executed because it mutates lifecycle files. |
| `bash -n` on all 10 backup restore scripts | 0 each | Static syntax valid; no restore was executed. |
| `sha256sum -c` on all 6 checksum sets | 0 each | All listed snapshot files matched. Four older backups have no checksum set. |
| Markdown relative-link scan excluding snapshots | 0 | 0 broken relative links. |
| Functional ID scan | 0 | 125 unique `FR-*`; no duplicates. 19 unique `AC-*`; no duplicates. |
| `git diff --check` | 0 | No whitespace errors in tracked diff. |

Skipped: product unit/integration tests, build, packaging, CI, ShellCheck, PHP/Python/Node tests, database validation, and service tests. **VERIFIED reason:** no product implementation, manifests, test files, build commands, or ShellCheck binary exist. Dependencies were not installed.

## 8. Documentation Authority Map

| Document | Intended/current authority | Owns | Duplication/update trigger/currentness |
| --- | --- | --- | --- |
| `docs/PRODUCT_DEFINITION.md` | Parent product intent, but only established statements are currently authoritative | identity, values, strategic proposals/open decisions | Duplicated by Blueprint; update on owner product decision; **not fully approved/currently conditional**. |
| `docs/PRODUCT_SYSTEM_BLUEPRINT.md` | Declares authoritative product-level blueprint subordinate to Constitution and Product Definition | product components, boundaries, MVP and deferred decisions | Duplicates Product Definition and master plan; update on approved product/system direction; current but assumes unresolved proposals. |
| `docs/FUNCTIONAL_SPECIFICATION.md` | Authoritative observable Core Alpha contract | `FR-*`, `AC-*`, workflows and release gates | Duplicates Blueprint behavior more precisely; update on governed behavior change; current. |
| `ai/core/04_ARCHITECTURE.md` | Governance shell for approved architecture | future architecture decisions | No substantive architecture yet; update after separately approved architecture; current as a truthful placeholder. |
| `ai/core/05_SYSTEM_MAP.md` | Concise pointer/map | verified runtime boundaries | Duplicates cross-links only; update when runtime exists; current. |
| `ai/core/13_ROADMAP.md` | Approved milestone order | sequencing and gates | Duplicates task summaries; update after milestone approval/completion; current but too terse to authorize next implementation. |
| `ai/projects/QWSG.md` | Concise contributor context/index | project status and authoritative pointers | Duplicates README/status; update on milestone; current. |
| `ai/projects/QWSG/QWSG_MASTER_PLAN.md` | Preserved source material, not normative | original ideas and examples | Extensively superseded by Blueprint/spec; update trigger should be none; current role is implied, not declared inside the file. |

**PROPOSED hierarchy:** Constitution and core policies -> owner-approved Product Definition decisions -> Product & System Blueprint -> Functional Specification -> separately approved architecture/security/data decisions -> implementation and tests. The Roadmap authorizes sequence; System Map and project records index verified current state; the Master Plan remains immutable source evidence. Owner ratification is required before treating currently proposed Product Definition choices as settled parent requirements.

## 9. Documentation Contradictions

### DOC-001 — HIGH

- Documents: `docs/PRODUCT_DEFINITION.md`, `docs/PRODUCT_SYSTEM_BLUEPRINT.md`, `docs/FUNCTIONAL_SPECIFICATION.md`.
- Conflict: the Product Definition labels target users, offline promise, privacy commitments, Agent/Console roles, and edition direction as proposals requiring owner approval, while the Blueprint and Functional Specification convert several into mandatory direction.
- Evidence: Product Definition “Status and authority” and “Owner decision register”; Blueprint Sections 1, 8, 17, 33–34; Functional Specification Sections 1–5 and 20.
- Recommendation: explicitly ratify the applicable product proposals or label downstream requirements provisional.
- Required human decision: **YES**.

### DOC-002 — HIGH

- Documents: `ai/core/04_ARCHITECTURE.md`, Functional Specification Sections 7, 15, 17, 22, Roadmap.
- Conflict: the functional contract is detailed enough for behavior design, but no application, security, data/state, module, installer, or Console architecture is approved. “Implementation may begin” in Task 006 history can be misunderstood as authorization to code without these boundaries.
- Evidence: Architecture status says foundation only; Functional Specification leaves topology, storage, secrets, update integrity, and Console security unresolved.
- Recommendation: complete a bounded Core Alpha Architecture package before code.
- Required human decision: **NO** for identifying the gap; **YES** to authorize the architecture task.

### DOC-003 — MEDIUM

- Documents: archived Prompt 005 and History 005.
- Conflict: archived Prompt 005 still says `Status: active`; its history says complete.
- Evidence: `ai/archive_prompts/005_2026-07-19_product-system-blueprint.md`; `ai/history/005_2026-07-19_product-architecture.md`.
- Recommendation: future lifecycle metadata correction under a governance task; preserve history.
- Required human decision: **NO**.

### DOC-004 — MEDIUM

- Documents: Blueprint Section 40, Functional Specification, History 006.
- Conflict: Blueprint recommends Task 006 as future work although Task 006 is complete.
- Recommendation: treat the section as dated recommendation or add a non-destructive supersession note in a later documentation task.
- Required human decision: **NO**.

### DOC-005 — HIGH

- Documents: archive/history records for identifier 001.
- Conflict: archived Prompt 001 is an unexecuted Product Architecture draft, while `ai/history/001_engineering_standard_update.md` is a different completed engineering task. The unified numeric identity is ambiguous.
- Recommendation: define a legacy-ID namespace/alias rule without renaming immutable records.
- Required human decision: **YES**, because identity policy and historical meaning are governance decisions.

### DOC-006 — MEDIUM

- Documents: Prompt/History 005 and 006 filenames and metadata.
- Conflict: Task 005 archive slug is `product-system-blueprint`, but history slug is `product-architecture`; Task 006 archive date is `2026-07-20`, history filename date is `2026-07-19`, while its metadata spans both dates.
- Recommendation: document backward-compatible filename exceptions and make future tooling validate semantic identity independently from legacy filename differences.
- Required human decision: **NO**.

### DOC-007 — MEDIUM

- Documents: Blueprint MVP, Functional Specification Sections 16 and 22.
- Conflict: Blueprint includes a secure standalone Console direction in the first usable release; the Functional Specification makes Console optional and leaves shipment timing a release gate.
- Recommendation: decide delivery sequence and distinguish product-direction inclusion from release inclusion.
- Required human decision: **YES**, before release planning; not before the first Agent slice.

### DOC-008 — MEDIUM

- Documents: `ai/README.md`, `ai/core/02_PROJECT_STRUCTURE.md`, prompt workflow.
- Conflict: `ai/README.md` describes `ai/prompts/` as reusable instructions, while current policy defines exactly one authoritative active task there.
- Recommendation: update the workspace README in a future governance-doc task.
- Required human decision: **NO**.

### DOC-009 — LOW

- Documents: master plan and Functional Specification.
- Conflict: source example uses `qwsg diagnose`; normative specification uses `qwsg diagnostics`.
- Recommendation: preserve the master plan as non-authoritative source; use the normative command.
- Required human decision: **NO**.

### DOC-010 — INFORMATIONAL

- Documents: README, Project Structure, Architecture, Blueprint, Functional Specification.
- Apparent conflict: extensive feature language may look implemented, but authoritative status text consistently says no runtime exists.
- Recommendation: keep prominent status labels and never market documentation as shipped functionality.
- Required human decision: **NO**.

## 10. Task Workflow and Backup Review

**VERIFIED strengths:** one active prompt is enforced by `bin/job`; Task ID and one matching history are validated; prompt Markdown is read as data; `next-task.sh` uses sanitized slugs, temporary files, no-clobber moves, rollback-on-failure logic, and root checks. The current active task passes validation.

**VERIFIED gaps:** `bin/job` validates only one exact Task ID metadata form and does not validate prompt semantic status, slug-to-history agreement, history placeholders, or archive lifecycle status. `next-task.sh` archives without changing semantic status, explaining the stale Task 005 `active` state. There is no maintained automated regression suite.

Backup formats are inconsistent. The first four backups have no `SHA256SUMS` or manifest; Task 003 lacks Git diff/log/tree; Task 005 lacks affected-files, Git diff/log/tree/permissions; Task 006 lacks affected-files/tree/permissions. All restore scripts pass Bash syntax, but older scripts are interactive while Task 007 explicitly requires non-interactive restoration. Point-in-time restore scripts may overwrite later legitimate lifecycle work if used after subsequent tasks. The first Task 007 snapshot is a concrete example: its checksums pass, but its preserved CHANGELOG and Engineering History predate Task 006 completion.

**PROPOSED:** standardize snapshot schema, declare whether restores are task-local or repository-timepoint operations, detect drift before overwriting preserved files, and require manifest/checksum/static-validation evidence. Do not execute historical restore scripts against the active tree without a task-specific impact review.

## 11. Security and Operational Risks

| Risk | Severity | Evidence | Recommendation |
| --- | --- | --- | --- |
| Coding before privilege/security boundaries | HIGH | No approved security architecture; Console and lifecycle require privilege-sensitive behavior. | Architecture and threat model before implementation. |
| State semantics without durable data contract | HIGH | Requirements mandate restart continuity, incidents, audit, corruption handling; no schemas exist. | Define versioned contracts and recovery behavior first. |
| Restore scripts overwriting later work | HIGH | Old Task 007 snapshot is valid cryptographically but stale semantically. | Add drift guards and per-target hashes. |
| Dirty, largely uncommitted baseline | MEDIUM | HEAD predates Tasks 004–007; many authoritative files are untracked. | Owner-reviewed commit strategy before implementation branching. |
| Collaborative records mode `0600` | MEDIUM | Histories 005–007 are not group-readable despite repository collaboration model. | Authorize and normalize future creation modes; avoid retroactive change without approval. |
| No dependency locking/build/test infrastructure | MEDIUM | No manifests, lock files, CI, or tests. | Choose stack, lock dependencies, create reproducible checks. |
| Temporary proprietary license unresolved | MEDIUM for release | LICENSE and release gate 8. | Resolve before public distribution. |

## 12. Quantum Creator Alignment Summary

**DOCUMENTED:** philosophy alignment is strong: human sovereignty, local-first operation, explainability, meaningful silence, reversibility, proportional technology, and ethical edition boundaries are explicitly represented. **NOT VERIFIABLE:** implementation alignment because no product implementation exists. See `ai/audits/2026-07-20_QUANTUM_CREATOR_CONFORMANCE.md`.

## 13. Core Alpha Readiness

**INFERRED:** requirements are substantially ready; implementation is not. The repository needs approved architecture, versioned contracts, security boundaries, stack/package decisions, and test scaffolding. See `docs/development/CORE_ALPHA_READINESS.md`.

## 14. What Must Remain Unchanged

- Human final authority and explicit consent for mutation.
- Agent-only and local-first usefulness.
- No automatic remediation in Core Alpha.
- `UNKNOWN` must never mean healthy.
- Transition-based alerts and recovery semantics.
- Separation of routine monitoring from privileged lifecycle work.
- Secrets redaction, localization readiness, bounded rollback, and evidence-based claims.
- Current authoritative product/governance documents during this read-only audit.

## 15. What Requires Fine-Tuning

- Explicit authority hierarchy and Product Definition ratification.
- Core Alpha architecture and contract package.
- Lifecycle metadata validation and legacy numbering rules.
- Standard snapshot schema and drift-safe restore semantics.
- Collaborative file-mode creation behavior.
- A maintained test harness for engineering workflow scripts.
- README/workspace status wording and dated blueprint recommendations.

## 16. What Conflicts with the Philosophy

No implemented product behavior conflicts because no product exists. **DOCUMENTED conflict risk:** treating proposed product decisions as settled weakens human sovereignty; using stale restore scripts without drift checks weakens reversibility; beginning privileged implementation without architecture would weaken explainability and Guardian-not-ruler boundaries.

## 17. Required Human Decisions

### DEC-001 — Ratify product authority

- Question: Which Product Definition proposals are approved as binding inputs to the Blueprint and Functional Specification?
- Why it matters: downstream documents currently use several proposals as requirements.
- Options: approve all referenced proposals; approve a specified subset; revise downstream requirements.
- Engineering recommendation: approve local-first operation, data minimization, Agent/Console responsibility split, human control, and ethical edition safety; defer pricing/naming.
- Postponement: architecture carries an unresolved authority defect.

### DEC-002 — Approve the next architecture milestone

- Question: May the next task produce a bounded Core Alpha Architecture package covering component/trust boundaries, data/state contracts, configuration/secrets, privilege model, runtime/packaging choice, and test strategy?
- Why it matters: implementation cannot safely allocate mandatory behavior without these decisions.
- Options: one consolidated architecture task; a short ordered series of architecture tasks; postpone implementation.
- Engineering recommendation: one bounded architecture package for the first vertical slice, with explicit deferred sections for Console and full lifecycle.
- Postponement: Core Alpha coding remains not ready.

### DEC-003 — Name and scope the first implementation slice

- Question: Should discovery, disk/inode monitoring, state transitions, local logging, config validation, `status`/`check`, dry-run preview, and isolated tests be authorized as an internal “Core Alpha Slice 1,” not as the complete Core Alpha release?
- Why it matters: the Functional Specification defines Core Alpha more broadly, including memory/load/services/HTTP/TLS/backups/e-mail/reporting/lifecycle.
- Options: approve the narrow slice; implement the entire functional scope at once; redefine Core Alpha.
- Engineering recommendation: approve the narrow vertical slice while preserving all remaining requirements as later slices and release gates.
- Postponement: milestone language and acceptance claims remain ambiguous.

## 18. Recommended Next Milestone

**PROPOSED:** Task 008 — `core-alpha-architecture`: ratify authority inputs and create the minimum implementable architecture for Slice 1. It should define runtime and packaging, process/privilege boundaries, module/check contract, observation/state/incident schemas, configuration and secret-reference boundary, logging/audit formats, filesystem ownership, CLI/JSON contracts, failure model, test harness, and migration/versioning rules. It must not implement product functionality.

## 19. Evidence Appendix

- Repository inventory: `find` reports 0 files in all product component and test/build directories, 1 file in `bin/`, 3 product documents in `docs/`, and the remainder predominantly governance/snapshot records under `ai/`.
- Runtime evidence: version commands listed in Section 5.
- Dependency evidence: manifest search returned only `.agents/skills/qwsg-job/agents/openai.yaml`; no product manifest.
- Permissions evidence: product placeholders are mode `2771`; `ai/` and `docs/` are setgid; histories 005–007 are `0600`.
- Git evidence: snapshot `git-status-before.txt`, `git-log-before.txt`, and `git-diff-before.patch`.
- Snapshot evidence: current snapshot `SHA256SUMS`, `manifest.txt`, and statically validated `restore.sh`.

## 20. Commands Executed

Read-only commands included `bin/job --check|--path|--history|--help`, `sed`, `rg`, `find`, `wc`, `stat`, `getfacl`, `git status`, `git log`, `git diff`, runtime `--version` commands, `bash -n`, `sha256sum -c`, a read-only Python Markdown-link/ID scan, and `git diff --check`. Snapshot generation used bounded `mkdir`, `cp`, inventories, checksums, and `chmod` only under the authorized Task 007 snapshot directory. No build, install, network, service, or rollback command ran.

## 21. Files Created or Modified

Created: this report; `ai/audits/2026-07-20_QUANTUM_CREATOR_CONFORMANCE.md`; `docs/development/REQUIREMENTS_TRACEABILITY_MATRIX.md`; `docs/development/CORE_ALPHA_READINESS.md`; and `ai/backups/20260720T183612Z_task007_repository_deep_audit/`.

Modified only as required by policy: `ai/history/007_2026-07-20_repository-deep-audit.md`, `ai/core/07_ENGINEERING_HISTORY.md`, and `CHANGELOG.md`.

## 22. Rollback Instructions

From the exact repository root run:

```bash
ai/backups/20260720T183612Z_task007_repository_deep_audit/restore.sh
```

The script removes only the four Task 007 report outputs and restores the preserved pre-audit Task 007 history, Engineering History, and CHANGELOG. It does not remove either Task 007 snapshot directory.
