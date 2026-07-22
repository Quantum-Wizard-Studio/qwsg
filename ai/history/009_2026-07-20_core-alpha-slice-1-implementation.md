# Engineering History 009: Core Alpha Slice 1 Implementation

## Task metadata

- Task ID: `009`
- Status: `complete`
- Date: `2026-07-20` UTC
- Human authority: `Project Owner`
- Responsible agent: `Aikó/Codex`

## Starting state and snapshot

The verified root was `<repository-root>`, branch `main`, HEAD `8fa40acd945b5f0d5d1ee0c5e182a19bba092d2b`. Task 008 was complete and Task 009 was the single valid active prompt. Go `1.26.5` linux/amd64, Make, Git, SHA-256, and ACL tools were available. The worktree contained substantial pre-existing documentation and workflow changes, all preserved.

Snapshot: `ai/backups/20260720T222559Z_task009_core_alpha_slice_1_implementation/`. It records the full dirty baseline, permissions, affected and new files, preserved copies, checksums, and bounded restore behavior.

## Decisions and risk controls

Go `1.26` is pinned in `go.mod`; the implementation uses only the standard library and builds one `qwsg` binary. The ratified exit contract is `0` complete, `1` fatal, and `2` partial. Persistence is excluded. Host-mutation, privilege, network, secret, rollback, and scope risks remain low through one-shot read-only collection, no elevation/listener/writes, privacy defaults, fixed command identities, time/output limits, and bounded rollback. Data correctness, partial ambiguity, Linux compatibility, schema compatibility, and test coverage remain medium hardening risks for Task 010.

## Work performed

Implemented the `qwsg inventory`, `qwsg version`, and `qwsg help` CLI; versioned inventory types and validation; deterministic aggregation; bounded shell-free command execution; OS, host, CPU, memory, storage, network, services, components, and capability collectors; JSON rendering; privacy redaction; unit and bounded integration tests; Make targets; development documentation; and a sanitized fixture.

## Verification

`gofmt`, `go vet ./...`, `go test ./...`, race tests, binary build, CLI contract checks, JSON parsing, local non-root read-only inventory execution, snapshot checksums, restore syntax, repository task validation, and `git diff --check` passed. The local run returned a structurally valid nine-category partial result with exit `2`; its raw host output was inspected for obvious secret markers, then discarded and not preserved. HEAD remained unchanged and no commit or push occurred.

## Limitations and open gates

No platform is declared supported. Service inventory depends on accessible systemd metadata; component detection is deliberately restricted to allowlisted Go discovery. Container/VPS variants, restrictive permissions, parser fuzzing, resource measurement, and broader platform fixtures remain Task 010 work. All architecture and release gates other than the Task 009 runtime/exit decisions remain open.

## Rollback

From the repository root run `ai/backups/20260720T222559Z_task009_core_alpha_slice_1_implementation/restore.sh`. It restores only Task 009-modified pre-existing files and removes only enumerated Task 009-created files/build output.

## Completion state

`complete — Task 010 independent verification and hardening is recommended`
