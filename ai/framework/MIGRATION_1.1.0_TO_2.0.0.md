# Migration from Framework 1.1.0 to 2.0.0

Framework 2.0.0 is lifecycle-compatible with existing repositories and task
history. Existing nine-category Authority Envelopes remain valid and are not
rewritten. New tasks may use the concise four-category v2 envelope:

- **Task targets and boundaries**
- **Permitted external actions**
- **Owner-reserved decisions**
- **Task-specific STOP conditions**

Ordinary reversible engineering authority is inherited from
`17_EXECUTION_MODEL.md`; it no longer has to be repeated in every task.

Repositories migrate by updating the framework version/configuration, adding
the v2 required-reading documents and diagnostic runner, and adopting the v2
validator/Builder. One-active-task, history, snapshot, rollback, deterministic
Builder, targeted staging, clean integration, credential protection, candidate
freeze, and release authority remain unchanged. Prepared v1 drafts and active
v1 tasks remain governed by their recorded authority until completion.

The old three-repeat rule is removed. Repeating a materially unchanged method
without a reason is still poor engineering, but classification and evidence—not
an arbitrary count—determine whether to retry, change method, or escalate.
