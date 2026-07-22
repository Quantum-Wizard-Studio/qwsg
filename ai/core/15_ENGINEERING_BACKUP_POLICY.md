# Engineering Backup Policy

## Purpose and authority

This policy defines safe engineering snapshots and backup records for QWSG. It is subordinate to the Project Constitution and applies to task snapshots, milestone snapshots, rollback evidence, retention, and publication review.

## Current state and target state

The repository currently contains early tracked backup directories with mixed payload, host-state evidence, manifests, checksums, and restore scripts. No managed external artifact storage is implemented or claimed.

The target state separates full rollback payloads from sanitized Git-tracked metadata. Full payloads belong in access-controlled external artifact storage when such storage is approved and implemented. Until then, new payloads must remain outside Git in a protected local location with documented retention and restore procedures.

## Required snapshots

A verified snapshot is mandatory before source, governance, lifecycle, infrastructure, security, migration, destructive, or broad repository changes; before any operation whose rollback cannot be reconstructed safely; and whenever an approved task explicitly requires one. Read-only inspection alone does not require a snapshot.

Every snapshot must identify its purpose, UTC creation time, repository and Git baseline, captured scope, exclusions, retention decision, and restore preconditions. It must be complete for the authorized rollback boundary and readable without proprietary or undocumented tooling.

## Payload and metadata separation

Full payloads include copied source trees, preserved files, patches containing unpublished content, repository archives, database or host dumps, permission/ACL captures, and other material needed for restoration. Full payloads must not be committed.

Git may track only publication-reviewed metadata necessary for engineering traceability: a sanitized manifest, SHA-256 checksum record, restore documentation, scope description, and verification result. A checksum for an external payload may be tracked, but the metadata must not reveal a sensitive external path, credential, username, host identity, UID/GID, ACL, secret, or private content.

## Manifest, integrity, and restore requirements

Each snapshot requires a deterministic manifest describing every payload object by repository-relative or sanitized logical name, type, size where useful, and SHA-256 digest. SHA-256 verification must pass after creation and before restore. Compressed repository snapshots require both checksum validation and an archive readability/listing test.

Restore documentation must name exact bounded targets, prerequisites, verification commands, and collision behavior. It must forbid extracting over a live worktree unless an explicitly reviewed procedure requires that action. Rollback validation must demonstrate that the payload is readable, referenced objects exist, checksums pass, restore commands are syntactically valid, and post-restore Git/tests/lifecycle checks are defined. A destructive restore is never executed merely to test documentation.

## Storage, access, retention, and deletion

Payload storage must use least-privilege access, encryption at rest where available, protected transport, access logging where available, and separation from public repository hosting. Credentials for storage never belong in a manifest, restore script, repository, or archive.

Retention is risk-based and must be stated when the snapshot is created. At minimum, retain a task snapshot until owner acceptance and retain a milestone snapshot until the milestone has been independently backed up and its rollback window is closed. Legal, security, incident, or release requirements may require longer retention.

Deletion of a payload before its documented retention boundary requires Project Owner approval. Deletion after the boundary requires a recorded review confirming that no active rollback, audit, incident, or release dependency remains. Repository metadata and early tracked backups are historical engineering records and must not be deleted or rewritten solely to enforce this policy; any later migration or removal requires a dedicated inventory, verified replacement evidence, and explicit owner approval.

## Privacy and prohibited Git content

Do not commit repository snapshot archives, tarballs, compressed payloads, copied source payloads, host-state dumps, pre-task evidence, raw Git diffs that expose private work, filesystem inventories, ownership/UID/GID/ACL captures, credential files, environment files, tokens, keys, passwords, or secrets. Hostnames, usernames, absolute server paths, IP addresses, e-mail addresses, infrastructure domains, and permission identities must be removed unless they are technically necessary, publication-safe engineering metadata and have passed explicit review.

Sanitized manifests, hashes, and restore documentation may be committed only after privacy and secret review. `.gitignore` is a safety layer, not authorization to store ignored secrets or a substitute for access control, retention, or deletion review.
