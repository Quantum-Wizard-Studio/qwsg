# Installed Package Classification

Task 075 establishes one local, read-only classification boundary for QWSG
installation identity. `internal/installation.Classify` is the canonical
consumer-facing mechanism. Guided installation and native update orchestration
must use it rather than independently interpreting a file named `qwsg`.

## States and evidence precedence

Classification is deterministic and fail closed. It examines only the
release-owned installed package layout and returns a bounded state, reason,
version, and (only when proven) declared migration identifier.

| State | Meaning |
|---|---|
| `no_installation` | No canonical package artifact exists. Unrelated configuration or state is not package evidence. |
| `verified_supported_installation` | The canonical package layout is complete, regular and non-symlinked; strict `qwsg.release/1` provenance is valid; and executable identity exactly agrees with it. |
| `supported_upgrade_source` | A verified installation additionally has an exact, locally declared migration to the supplied candidate version. |
| `legacy_installation` | A binary-only artifact reports a syntactically valid major-zero QWSG identity. It is recognized for safe operator handling, never trusted as a package. |
| `unknown_unverified_installation` | Binary-only evidence, an unsupported package version, or an installed source without a declared route to a newer candidate cannot establish a supported installation decision. |
| `inconsistent_incomplete_installation` | Canonical artifacts are partial, unsafe types, malformed, unreadable, or disagree on provenance. |

Evidence precedence is: safe canonical layout; strict installed release
provenance; exact executable-to-provenance agreement; supported major; then,
when a candidate is supplied, the existing declared migration registry.
Configuration, credentials, Guardian checkpoints, readiness, rollback state,
PATH lookup and arbitrary filesystem names do not strengthen package identity.
The installed QWSG 1.2 layout has no installed package manifest, so Task 075
does not invent duplicate metadata or falsely require archive-only
`MANIFEST.sha256` after installation.

The classifier executes only the canonical installed binary's bounded
`version` operation, with discarded stderr, a two-second timeout and a 4096-byte
output limit. Version output alone is never sufficient: schema, version,
commit, build timestamp, and platform must match strict installed
`RELEASE.json` evidence. Results expose stable reason tokens, not filesystem
paths, configuration contents, credentials, host identity, or command output.

## Installer and updater decisions

On `no_installation`, guided installation may present a fresh package install.
A verified current/newer package may enter setup without rewriting package
files. A supported older package is directed to the explicit existing update
workflow. Legacy, unknown, incomplete, inconsistent, and unsafe evidence stops
before consent or mutation; the operator must back up and resolve it explicitly.

The updater obtains its installed source version through the same classifier.
Candidate package verification, declared migration planning, privileged apply,
post-install validation and rollback remain the Task 064/1.2 architecture; no
second updater was introduced. The local migration registry remains
authoritative for executable compatibility.

## Field regression and future release awareness

The server.quantumwizard.hu field case—a lone `/usr/local/bin/qwsg` reporting
`QWSG 0.0.1-prealpha`—now classifies as legacy and cannot cause the wizard to
claim that a verified package is already installed or silently overwrite the
artifact. Deterministic classifier and guided-decision tests preserve this
boundary.

Installed-version trust remains separate from remote-release trust. Task 075
does not implement release-index retrieval, Guardian scheduling, notification,
telemetry, or installation automation. Future discovery may consume only a
verified local classification result as its installed-version input.
