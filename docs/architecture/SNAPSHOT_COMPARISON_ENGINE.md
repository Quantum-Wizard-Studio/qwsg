# Snapshot Comparison Engine

## Architectural boundary

The Comparison Engine is the only supported source of system-change facts:

```text
Inventory -> Snapshot Store -> Comparison Engine -> Drift Engine
                                                     |
                                   Health / future Rules / Policy
```

The Canonical Drift Engine consumes validated Change Records and classifies
their semantic type. The Canonical Health Engine consumes validated Drift
Records and evaluates their engineering condition. Future Rule, Policy, Alert,
Reporting, CLI,
e-mail, and Web UI components consume Drift and Health contracts. They must not
compare Inventory snapshots directly or maintain a competing change model.

`internal/comparison` imports the Inventory domain model. It accepts two
validated Inventory 1.0 snapshots containing canonical
`canonical-system-inventory-v1` data, requires the same subject and canonical
contract, and returns a validated comparison result. It performs no collection,
persistence, host access, networking, scheduling, privilege escalation, or
mutation.

## Semantic comparison profile

Profile `canonical-layer-resource-fact-v1` independently walks every canonical
layer and compares:

- layer availability status;
- resource presence and canonical kind;
- every canonical fact and its typed semantic attributes.

Layer, resource, and fact unions are sorted before comparison. Observation and
provenance timestamps, collection duration, request identifiers, labels,
relationships, and generic metadata are intentionally excluded because they
describe acquisition or annotation rather than a canonical state value in this
profile. Expanding this profile requires an explicit, versioned contract change.

An absent value becoming present is `added`; present becoming absent is
`removed`; unequal typed values are `modified`; equal typed values are
`unchanged`. The engine detects facts only and makes no judgement.

## CLI selection

```text
qwsg compare --store /absolute/private/store
qwsg compare --store /absolute/private/store --from SNAPSHOT --to SNAPSHOT
qwsg compare --store /absolute/private/store --format json
qwsg compare --store /absolute/private/store --format human
```

Without selectors, the two newest names returned by the deterministic store
listing are compared from previous to latest. Explicit selection requires both
`--from` and `--to`; each is an exact snapshot filename. `QWSG_STORE` and
`QWSG_FORMAT` provide the same configuration. JSON is the default and canonical
interface.

Successful comparison exits `0`, including a comparison with no detected
changes. Usage, insufficient history, store integrity, source incompatibility,
validation, or output failure exits `1`. Source Inventory completeness is
reported in comparison metadata and never converted into a health judgement.

Human output groups Added, Removed, Modified, and Unchanged records. It is
generated only from the canonical result, escapes terminal control characters,
and contains no diagnosis, recommendation, severity, or score.

The permanent Compare-to-Drift boundary and its compatibility rules are defined
in `CANONICAL_DRIFT_ENGINE.md`. Comparison remains factual and does not assign a
Drift category.

## Security and resource properties

The engine is read-only and side-effect free. Snapshot loading retains the
store's absolute private-path, permissions, symlink, size, duplicate-key,
checksum, strict-decoding, and Inventory validation controls. Complexity is
bounded by the already bounded Inventory documents plus sorted unions; no
external dependency, daemon, scheduler, database, or network communication is
introduced.
