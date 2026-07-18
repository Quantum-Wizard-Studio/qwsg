# Task 003 Product Definition Start State

## Purpose

This snapshot records the verified state before Task 003 creates the Quantum Wizard Server Guardian Product Definition.

## Status

Captured at `2026-07-18T23:38:11Z` before Product Definition target documents were modified.

## Baseline

- Root: `/home/qws/web/qwsg.quantumwizard.hu/qwsg`
- Branch: `main`
- HEAD: `4b4a0ac97fe3ac21ef1fae4beec70870cd53cad5`
- Git state: prompt rotation present but not committed (`002_CURRENT_TASK.md` deleted; archived Prompt 002, active Prompt 003, and pending History 003 untracked)
- Owner and group: `attila:qwdev`
- Directory ACL: owner and group write inherited with setgid
- Prompt 003 and History 003 mode variance: `0600`, missing expected group write
- Product Architecture: not started; architecture governance remains foundation-only and the earlier architecture prompt remains archived without execution
- Existing Product Definition: none

## Strategic evidence boundary

Existing governance approves QWSG's protective purpose, administrator-like behavior, explicit authorization for automatic correction, security and stability priority, logging, explainability, reversibility where possible, modularity, independent-server installation, product independence, an independent QWSG console, and long-term maintainability. No existing record approves edition names, Free/Professional feature boundaries, cloud services, pricing, or commercial terms. Those subjects must remain clearly marked for owner approval.

## Planned changes

Create `docs/PRODUCT_DEFINITION.md`; update only concise references in the root README, project record, roadmap, and Engineering History; correct the Task 003 history filename; update Prompt 003 and its independent history; and retain this rollback record. No implementation, architecture, API, database, framework, installer, deployment, service, dependency, or operating-system work is authorized.

This snapshot is immutable. The Product Definition will evolve only through approved product-governance work.
