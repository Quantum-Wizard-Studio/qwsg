# Engineering Task Builder Architecture

## Purpose

The Engineering Task Builder is the official creation boundary between Project Owner intent and executable QWSG engineering tasks. It replaces manual assembly of `CURRENT_TASK` documents with validated deterministic generation.

## Boundary and interface

`ai/scripts/task-builder.sh` is a dependency-free repository-local command. It runs only from the verified QWSG root and never executes a generated task. It supports:

- interactive structured collection, with `.` on its own line terminating each multi-line field;
- deterministic `--input-dir <directory>` generation;
- read-only `--check-input <directory>` validation.

The input-directory contract uses one regular text file per field:

```text
slug title authority language objective scope out-of-scope starting-state
snapshot risk planned-work rollback deliverables verification documentation
completion approval
```

`slug`, `title`, `authority`, `language`, and `approval` are single-line values. The remaining substantive fields preserve multi-line content. `approval` must equal `APPROVE`. Input files are read as data and are never sourced, evaluated, or interpreted as shell commands.

## Ownership of information

The Project Owner supplies every substantive and authority-bearing value. The builder generates only mechanical lifecycle data: the next sequential Task ID, UTC date, approved state, filenames, mandatory required-reading section, approval record, and initial history lifecycle text. It does not invent scope, acceptance criteria, risks, or implementation work.

## Generation pipeline

1. Verify repository identity and exactly one completed current prompt/history pair.
2. Validate all structured fields, slug syntax, absence of unresolved editing markers, and explicit approval.
3. Calculate the next three-digit ID and collision-free destinations.
4. Render the complete prompt and matching history into same-directory temporary files.
5. Validate required metadata, identity, approval state, and required prompt sections before mutation.
6. Archive the completed prompt and install the new pair with no-clobber moves.
7. Run `bin/job --check` and `next-task.sh --check` against the installed state.
8. Report an approved but unstarted task and stop.

Identical input, completed-task identity, and UTC date produce byte-identical generated documents.

## Transaction and rollback

The transaction tracks the archive, prompt, and history moves independently. A failure at any point removes only artifacts installed by the current builder invocation, restores the exact archived active prompt, and removes temporary files. Failure injection after each move is covered by tests. Destination collisions fail before mutation, and broad cleanup or Git reset is never used.

The builder preserves the existing `next-task.sh` workflow for compatibility. That command continues to generate an explicitly unapproved draft for a separate review cycle; it is not the official direct approved-generation path.

## Security and operational properties

- No dependency installation, network access, privilege change, commit, push, or task execution occurs.
- Owner content is never executed.
- Exact approval is mandatory and recorded in both lifecycle state and prompt approval text.
- Existing file ownership and inherited ACL behavior are preserved through same-directory creation.
- Generated task content remains English engineering documentation; the configured owner communication language is metadata for later owner-facing messages.
