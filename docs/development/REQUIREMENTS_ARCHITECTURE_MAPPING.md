# Requirements to Architecture Mapping

## Authority boundary

The Functional Specification owns all 125 `FR-*` and 19 `AC-*` identifiers. This document does not redefine or renumber them. Task 008 ratifies the Product Definition, Blueprint, Functional Specification, and Task 007 outputs as design baseline only within the prompt's stated limits: proposals, assumptions, and release gates remain unresolved.

## Slice 1 direct mapping

| Architecture area | Direct requirements | Acceptance inheritance | Disposition |
| --- | --- | --- | --- |
| Human/read-only authority | `FR-AUTH-001`–`003`, `FR-NFR-001` | `AC-LIFE-001` | Architecture enforces no mutation or elevation; lifecycle acceptance is not claimed complete. |
| Capability discovery | `FR-CAP-001`–`005` | `AC-AGENT-002`, `AC-DOC-002` | Slice 1 implements truthful inventory capability states; supported platform remains gated. |
| Collector contract | `FR-CHECK-001`–`007` | `AC-AGENT-002`, `AC-AGENT-004` | Direct contract input; health transitions are deferred. |
| CLI and presentation | `FR-CLI-001`–`005`, `FR-NFR-003`, `FR-NFR-005`, `FR-NFR-006` | `AC-UX-002`, `AC-UX-004` | Read-only subset only; partial exit policy is `AG-011`. |
| Data/schema/failure | `FR-DATA-002`, `FR-DATA-004`, `FR-REL-002`, `FR-REL-005` | `AC-DATA-001` | Versioning, atomic optional persistence, no silent reset. |
| Privacy/security | `FR-NFR-001`–`006`, `FR-SEC-002` | `AC-DATA-002` | Slice 1 uses stricter no-network/no-secret boundary; support bundle itself is deferred. |
| Diagnostics/capabilities | `FR-DIAG-001` | `AC-DOC-001` | Collector self-report is architecture support toward later diagnostics. |

## Full requirement-family disposition

| Family | IDs | Architecture owner or status |
| --- | --- | --- |
| Authority | `FR-AUTH-001`–`003` | Core architecture/security; Slice 1 direct. |
| Profiles | `FR-PROFILE-001`–`004` | Core architecture; Slice 1 proves Agent-only subset, presets deferred. |
| Capability | `FR-CAP-001`–`005` | Slice 1 direct; platform release gate remains. |
| Configuration | `FR-CFG-001`–`010` | Strict minimal Slice 1 boundary; full activation/threshold behavior deferred to later slices and `AG-007`. |
| Secret handling | `FR-SEC-001`–`003` | Security architecture; Slice 1 prohibits secrets, backend deferred. |
| Check contract | `FR-CHECK-001`–`007` | Collector/coordinator/data model; Slice 1 direct. |
| Disk | `FR-DISK-001`–`003` | Storage inventory supports later checks; health evaluation deferred. |
| Memory | `FR-MEM-001`–`003` | Memory inventory supports later checks; health evaluation deferred. |
| Load | `FR-LOAD-001`–`002` | CPU capacity inventory supports later checks; load health deferred. |
| Services | `FR-SVC-001`–`003` | Read-only service inventory subset; monitoring policy deferred. |
| HTTP | `FR-HTTP-001`–`003` | Excluded from Slice 1; later network-check collector architecture. |
| TLS | `FR-TLS-001`–`002` | Excluded from Slice 1; later network-check collector architecture. |
| Backups | `FR-BACKUP-001`–`003` | Excluded from Slice 1; later filesystem-check collector architecture. |
| State | `FR-STATE-001`–`007` | Data model provides unknown/freshness foundations; health state deferred. |
| Incidents | `FR-INC-001`–`004` | Deferred to state/incident slice. |
| Maintenance | `FR-MAINT-001`–`004` | Deferred to policy/notification slice. |
| Alerts | `FR-ALERT-001`–`007` | Deferred; e-mail transport `AG-002`. |
| Reports | `FR-REPORT-001`–`004` | Deferred; inventory presentation is not a Core Alpha report. |
| CLI | `FR-CLI-001`–`006` | Read-only subset direct; state-changing behavior deferred. |
| Lifecycle | `FR-LIFE-001`–`015` | Installer boundary excluded from Slice 1; architecture preserves separation. |
| Console | `FR-CONSOLE-001`–`005` | Adapter boundary only; network Console blocked by `AG-004`. |
| Data | `FR-DATA-001`–`006` | Versioning/fault semantics direct; incident/history retention deferred via `AG-003`, `AG-008`. |
| Diagnostics | `FR-DIAG-001`–`004` | Capability foundation direct; support bundle deferred. |
| Reliability | `FR-REL-001`–`005` | Failure isolation direct; scheduler and durable health deferred. |
| Non-functional | `FR-NFR-001`–`007` | Security/localization direct except Console accessibility, which is deferred. |

## Acceptance-criterion disposition

| IDs | Slice 1 contribution |
| --- | --- |
| `AC-AGENT-001` | Deferred: requires installation, scheduling, state, e-mail, report, diagnostics. |
| `AC-AGENT-002` | Partial: deterministic inventory outcome fixtures; health severities/recovery deferred. |
| `AC-AGENT-003` | Partial foundation only; monitoring behavior deferred. |
| `AC-AGENT-004` | Partial: collector failure isolation; Console/alerts deferred. |
| `AC-STATE-001`–`004` | Deferred to state and alert slices. |
| `AC-UX-001` | Partial: strict minimal configuration; activation deferred. |
| `AC-UX-002` | Partial: read-only discovery CLI/JSON/permission contracts. |
| `AC-UX-003` | Deferred and gated by `AG-004`. |
| `AC-UX-004` | Direct for Slice 1 presentation catalogs. |
| `AC-LIFE-001`–`002` | Deferred to Installer/lifecycle architecture and implementation. |
| `AC-DATA-001` | Partial: schema/storage fault behavior; restart incident continuity deferred. |
| `AC-DATA-002` | Partial: secret/redaction/no-upload rules; bundle generation deferred. |
| `AC-DOC-001` | Direct traceability obligation; no full Core Alpha completion claim. |
| `AC-DOC-002` | Deferred until supported implementation exists. |
| `AC-REL-001` | Remains binding; all release gates must close before supported release. |

## Proposal and gate controls

The following are not promoted to mandatory business policy: target personas, editions, pricing, licensing, cloud role, telemetry policy beyond the current no-default-transmission functional requirement, long-term fleet direction, and commercial support. Platform versions, Console shipment/security, e-mail transport, retention, storage technology, configuration/secrets, topology, update authenticity, and public distribution remain in the gate register.
