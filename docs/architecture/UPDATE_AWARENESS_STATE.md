# Update Awareness State

## Boundary

Task 077 implements `qwsg.update-awareness/1` in
`internal/updateawareness`. It is an ordinary-user, read-only awareness record,
not Guardian checkpoint, Current Operator State, scheduler, rollback,
notification, or installation state.

```text
Task 075 verified installed identity
  + Task 076 authenticated CheckResult
  -> deterministic awareness transition
  -> private atomic state
  -> update check/status presentation
```

No production endpoint or trust anchor is activated. Until a later Owner
decision supplies both, `update check` records `source_authority_refused` and
fails closed; unsigned Forgejo data is never awareness truth.

## Record and integrity

The record is `<canonical-state-root>/update/awareness.json`, schema
`qwsg.update-awareness/1`, bounded to 64 KiB. State directories are
current-user-owned mode `0700`; record and lock are current-user-owned,
single-link regular files mode `0600`. Clean absolute paths are required.
Symlinks, special files, hard links, wrong owners and permissive modes fail
closed.

The SHA-256 claim covers normalized typed JSON. Loading uses bounded strict JSON
decoding and validates schema, digest, timestamps, versions, relationships,
Task 075 classification, Task 076 Ed25519 evidence, artifact identity and cache
validators. This local digest is not a secret MAC and never replaces Task 076
publisher authenticity.

Publication takes a non-blocking advisory lock, writes and syncs a
same-directory temporary file, atomically renames it, then syncs the directory.
Pre-rename failure preserves the prior record. Post-rename sync uncertainty is
reported while the renamed integrity-valid record remains recoverable.

## Semantics

States are `current`, `update_available`,
`update_available_unsupported_source`, `withdrawn`, and `unknown`; absence is
`never_checked`. Latest attempt time/outcome/failure and last successful
authenticated observation/freshness are separate. Default freshness is 48
hours.

A failure preserves authenticated success only for the same source, channel
and installed identity and never refreshes freshness. Changed installed
identity is reported as `installed_identity_changed`. A `304` refresh requires
a matching prior Ed25519-authenticated observation. An authenticated withdrawal
makes the cached release non-actionable while retaining historic identity.

`qwsg update check` is the sole explicit network refresh and never acquires or
installs. `qwsg update status` opens no network source. Awareness never changes
Guardian health/readiness. Notification transition/delivery identity and
Guardian scheduling remain separate future approval boundaries.
