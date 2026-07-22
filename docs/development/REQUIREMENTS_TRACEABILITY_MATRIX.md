# QWSG Requirements Traceability Matrix

## Status and interpretation

This audit-time matrix maps product promises to their architectural home and repository evidence. Every implementation status is evidence-based. `docs/FUNCTIONAL_SPECIFICATION.md` supplies the functional IDs; `docs/PRODUCT_SYSTEM_BLUEPRINT.md` supplies conceptual components. There is no approved detailed architecture or product implementation.

| Feature group | Philosophy/product reason | Blueprint | Functional requirements | Architecture element | Repository path/status | Test coverage | Gap |
| --- | --- | --- | --- | --- | --- | --- | --- |
| System discovery | Verify, do not assume; portable independent servers | 15, 18, 29–31 | `FR-CAP-001`–`005` | Detection/capability boundary | `agent/`, `installer/`: PLACEHOLDER | Acceptance only: `AC-AGENT-001` | Support matrix, adapters, normalized inventory contract absent. |
| Disk monitoring | Early prevention and explainable protection | 21, 31 | `FR-DISK-001`–`003` | Check/module + state | `modules/`: PLACEHOLDER | `AC-AGENT-002`–`003` | Filesystem policy and fixtures not implemented. |
| Inode monitoring | Prevent capacity failure not visible in bytes | 21, 31 | `FR-DISK-001`–`003` | Check/module + state | `modules/`: PLACEHOLDER | `AC-AGENT-003` | Same as disk; independent incident contract needs design. |
| CPU and memory | Correct Linux semantics, proportional evidence | 12, 21, 31 | `FR-MEM-001`–`003`, `FR-LOAD-001`–`002` | Check/module | `modules/`: PLACEHOLDER | `AC-AGENT-002`–`003` | No collectors, capacity adapters, fixtures. |
| Service monitoring | Outcome-oriented protection | 21, 31 | `FR-SVC-001`–`003` | systemd adapter/check | `modules/`: PLACEHOLDER | `AC-AGENT-003` | Unit permission and transient-state design absent. |
| SSL monitoring | Prevent certificate outage | 21, 31 | `FR-TLS-001`–`002` | TLS check | `modules/`: PLACEHOLDER | `AC-AGENT-003` | Client trust source and endpoint model absent. |
| Backup monitoring | Explain backup freshness without owning backups | 11, 21, 31 | `FR-BACKUP-001`–`003` | Backup evidence check | `modules/`: PLACEHOLDER | `AC-AGENT-003` | Artifact-set contract and privacy handling absent. |
| HTTP monitoring | Verify service outcome, not process presence | 21, 31 | `FR-HTTP-001`–`003` | HTTP check | `modules/`: PLACEHOLDER | `AC-AGENT-003` | Client, credential boundary, origin policy absent. |
| State transitions | Meaningful interpretation and silence | 22–23 | `FR-STATE-*`, `FR-INC-*` | State/incident engine | `agent/`: PLACEHOLDER | `AC-STATE-001`–`003` | Versioned schemas, persistence, clock/restart algorithm absent. |
| Alerting | Warn on meaningful change | 23, 31 | `FR-ALERT-001`–`007` | Alert manager/channel | no path implementation: DOCUMENTED | `AC-STATE-002`–`004` | E-mail transport, queue/delivery contract absent. |
| Recovery notifications | Explain restored or remaining risk | 23 | `FR-STATE-003`, `FR-INC-004`, `FR-ALERT-001`–`003` | Incident/alert engine | `agent/`: PLACEHOLDER | `AC-STATE-001`–`003` | No implementation or fixtures. |
| Local state | Local-first continuity and data ownership | 26 | `FR-DATA-001`–`006` | State store | no implementation: DOCUMENTED | `AC-DATA-001` | Technology, schema, locking, corruption and migration absent. |
| Logging/audit | Transparency and traceability | 26 | `FR-AUTH-002`, `FR-DATA-005`–`006`, `FR-NFR-003` | Logging/audit boundary | no implementation: DOCUMENTED | Partial acceptance references | Formats, sinks, rotation, integrity, retention absent. |
| Configuration | Operator-controlled validated policy | 25 | `FR-CFG-001`–`010`, `FR-SEC-001`–`003` | Config/secrets boundary | no implementation: DOCUMENTED | `AC-UX-001` | Syntax, schema, permissions, activation mechanism absent. |
| CLI | Explainable local control | 17, 27 | `FR-CLI-001`–`006` | CLI/application service boundary | `bin/job` is unrelated; product CLI ABSENT | `AC-UX-002` | Command parser, JSON schema, RBAC and dry-run vocabulary absent. |
| Console | Optional localized administration | 19, 30–31 | `FR-CONSOLE-001`–`005` | Web trust boundary | `console/`: PLACEHOLDER | `AC-UX-003`–`004` | Security architecture and shipment decision unresolved. |
| Installer | Consent-based privileged lifecycle | 18, 28 | `FR-LIFE-001`–`015` | Privileged installer boundary | `installer/`: PLACEHOLDER | `AC-LIFE-001`–`002` | Package, ownership manifest, plan format and privilege design absent. |
| Update | Integrity and recoverability | 28, 36 | `FR-LIFE-008`–`010` | Update/migration boundary | ABSENT | `AC-LIFE-001` | Authenticity, compatibility and downgrade policy absent. |
| Uninstall | Reversibility and data ownership | 28 | `FR-LIFE-013`–`015` | Removal/export boundary | ABSENT | `AC-LIFE-002` | Owned-artifact manifest/export format absent. |
| Diagnostics | Explainability and safe support | 27 | `FR-DIAG-001`–`004` | Diagnostic/redaction service | ABSENT | `AC-DATA-002` | Bundle schema, redaction engine and identifying-data policy absent. |
| Database monitoring | Later service depth | 21, 32 | Explicit Core Alpha exclusion, Section 21 | Future module | `modules/`: PLACEHOLDER | None required for Alpha | Post-MVP requirements not yet specified. |
| Mail monitoring | Later service depth | 21, 32 | Explicit Core Alpha exclusion, Section 21 | Future module | `modules/`: PLACEHOLDER | None required for Alpha | Post-MVP requirements not yet specified. |
| Log analysis | Interpret rather than spam raw logs | 21, 32 | Explicit Core Alpha exclusion, Section 21 | Future module | `modules/`: PLACEHOLDER | None required for Alpha | Post-MVP requirements not yet specified. |
| Multi-server support | Scale without weakening local control | 6, 33 | Core Alpha single-server boundary | Future fleet coordination | ABSENT | None | Business, identity, protocol and topology decisions deferred. |
| Cloud connectivity | Optional additive capability | 33–34 | `FR-NFR-004`; no vendor transmission by default | Future optional remote boundary | ABSENT | None | Owner approval, privacy/data contracts and topology deferred. |
| Automatic remediation | Act only with explicit bounded authority | 12–14 | Excluded from Core Alpha, Sections 3 and 21 | Future remediation boundary | ABSENT | None | Must remain absent until separately specified and authorized. |

## Coverage conclusion

**DOCUMENTED:** every important promised feature has a philosophical/product rationale and a conceptual home. Core Alpha features generally have functional and acceptance requirements. **VERIFIED:** no feature has implementation or executable test coverage. Post-MVP database, mail, log, fleet, cloud, and remediation groups do not yet have normative feature-level requirements, which is appropriate until separately authorized.
