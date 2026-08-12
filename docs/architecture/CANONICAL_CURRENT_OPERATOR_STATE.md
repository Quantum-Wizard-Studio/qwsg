# Canonical Current Operator State

## Boundary

`Canonical execution -> typed Overview projection -> Canonical Current Operator State -> OverviewProvider -> Canonical Operator Presentation Model freshness requalification -> replaceable interface`

Current Operator State 1.1 is one durable local handoff record, not a general persistence platform. It stores one already validated canonical Operator Overview in a versioned envelope. The envelope adds stable state identity, SHA-256 payload integrity, explicit `inventory_snapshot` or `operator_evaluation` coverage, Command provenance, observation/publication/freshness times, and bounded canonical JSON. It never persists `command.StageResult.Value` or interprets condition, attention, Health, Alert, Guardian, completeness, or recommendations. Readers retain compatibility with coherent legacy State/Overview 1.0 records; new publications use 1.1 so bounded-attention disclosure is explicit.

## Publication

The existing `check` profile remains `live Inventory -> Snapshot`. After its validated execution, the application requires exactly two complete, correctly ordered and versioned stage results with concrete `inventory.Snapshot` values and matching snapshot identities. It projects Command plus Inventory observations through `presentationmodel.Project`, normalizes the envelope, then publishes it. Inventory coverage is current evidence but not Health: condition remains `unknown`, Health/Policy/Alert/Runtime/Service evidence remains absent, and Guardian remains not observed.

`status`, `health`, `report`, and advanced definitions do not implicitly publish. Failed, incomplete, untyped, mismatched, or invalid executions cannot replace state. Publication precedes terminal rendering, so presentation failure cannot create a partial state.

`observe` is the explicit full-evaluation publisher. An empty Inventory Store causes one persisted `check` bootstrap and limited unknown state. With a valid baseline, the application requires eight complete typed and correlated stage values, projects them through the existing Presentation Model, and publishes `operator_evaluation` coverage. Corrupt, incompatible, unsafe, or unreadable stores fail rather than becoming a new baseline. One-shot evaluation supplies no Runtime Service evidence, so Guardian remains not observed.

Application failures are classified at the privacy boundary as
`evaluation_pipeline_failed`, `operator_projection_failed`, or
`current_state_publication_failed`. Raw Go errors, host paths, identifiers,
and evidence values are never rendered as diagnostics.

## Storage and durability

The application resolves the state directory from explicit `QWSG_STATE_DIR`, then `XDG_STATE_HOME/qwsg`, then `$HOME/.local/state/qwsg`. The resolved location is an application storage input, not engineering policy. It must be clean and absolute. Symlink components are rejected; the state directory and file must be owned by the current user with modes `0700` and `0600`.

On direct ordinary-user first use, QWSG creates a missing state hierarchy
recursively and validates the final QWSG root before any store or operation lock
uses it. Existing ancestors are inspected for symlinks but their ownership and
modes are not changed. The installer remains outside this per-user runtime
responsibility.

Publication writes bounded canonical bytes to an exclusive same-directory temporary file, syncs and closes it, atomically renames it to `current-operator-state.json`, then syncs the directory. Pre-rename failures preserve the old record and remove the temporary file. Loading performs bounded strict decoding, version, identity, digest, provenance, timestamp, Overview, ownership, mode, and path validation.

## Consumption and time

Bare `qwsg` performs one bounded load and never collects automatically. Missing state uses the existing unavailable Overview. A valid state is passed to `presentationmodel.RequalifyFreshness` with the current time and exclusive freshness deadline. Before the deadline the exact Overview is retained; at and after it freshness becomes stale, complete evidence becomes partial, healthy/unknown condition degrades, recommendations are recomputed, and identity changes. A stale Overview can never be upgraded without new evidence.

Corrupt, incompatible, unsafe, permission-invalid, and unreadable states fail closed with distinct bounded localized diagnostics. The Console still receives only validated Overview values and imports neither storage nor engineering packages.

Interactive Console refresh performs this same load, validation, freshness
requalification, and render path. It never acquires an operation lock, invokes
`observe`, collects evidence, runs Pipeline, or publishes replacement state.

Scheduler state, Alert/Notification/Runtime history, incidents, monitoring databases, retention, remote storage, APIs, Dashboard, services, and remediation remain outside this contract.
