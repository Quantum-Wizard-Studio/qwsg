# Canonical Command Architecture

Task 035 adds the `observe` live Report profile. It uses the existing dependency closure exactly once: Inventory→Snapshot→Compare→Drift→Health→Rule→Policy→Report. The application establishes an Inventory/Snapshot baseline on the first invocation and runs the full profile only when prior evidence exists. It never compares a snapshot with itself.

`check` remains Inventory→Snapshot. Its validated typed result may still publish limited Current Operator State; publication does not add a stage or turn Inventory completion into Health.

The `observe` application boundary distinguishes a canonical Pipeline failure,
an Operator projection/validation failure, and a Current Operator State
publication failure with bounded privacy-safe tokens. This classification does
not change Command definitions, stage ordering, or engine results.

The Console refresh adapter resolves the existing `status` definition and performs one explicit Pipeline execution. It adds no profile, order, retry, polling, or interpretation; advanced CLI and JSON composition remain unchanged.

## Status and purpose

Task 024 establishes the permanent, presentation-independent Canonical Command
Architecture of Quantum Wizard Server Guardian. Command Definition 1.0 is the
single public request contract for every presentation interface. The current
CLI is its first reference adapter; it is not the architecture.

The mandatory separation is:

```text
Canonical Engineering Engines
              ↑
Canonical Pipeline Orchestration
              ↑
Canonical Command Architecture
              ↑
Replaceable Presentation Adapters
```

No CLI, Interactive Terminal, Dashboard, REST API, or other presentation may
collect evidence, select or order engines, evaluate health or rules, generate
reports, or implement orchestration. Replacing a presentation cannot change
command identity, planning, stage selection, engine behavior, or execution
results.

## Components and responsibility boundaries

- `internal/command` owns Command Definition 1.0, profiles, grammar,
  normalization, validation, stable identity, deterministic planning, parameter
  projection, and presentation-neutral Command Execution contracts. It imports
  no engineering engine or presentation package.
- `internal/pipeline` is the only orchestration layer. It converts a validated
  command plan into calls to the existing Inventory, Snapshot, Compare, Drift,
  Health, Rule, and Report boundaries. It neither parses CLI syntax nor renders
  output.
- `internal/presentation` consumes a completed Command Execution. Its terminal
  and JSON renderers cannot plan or execute stages and cannot reinterpret
  canonical evidence.
- `cmd/qwsg` maps process arguments and environment defaults to canonical
  definitions, invokes the orchestrator, selects a renderer, and maps failures
  to process output. Existing `inventory` and `compare` compatibility commands
  remain available.
- The canonical engine packages retain exclusive ownership of their existing
  semantics. Orchestration calls them; it does not copy them.

The dependency direction is one-way. An engine cannot depend on command or
presentation code. The command model cannot depend on a presentation.

## Public Contracts and Parameter Model

The public contract is identified by `qwsg.command/1.0`. A definition contains:

- a stable content-derived command ID;
- command and optional profile identity;
- source selection (`live` or `store`);
- optional store and exact from/to snapshot selectors;
- requested pipeline targets;
- explicit engine inclusion and exclusion;
- filters, grouping, sorting, output format, and presentation selection;
- fixed contract version.

Normalization creates non-null collections, deduplicates and canonically orders
engine selections, sorts projection parameters, validates fixed registries and
bounds, rejects contradictions, and derives the stable SHA-256 identity from
canonical JSON. Unknown versions, stages, options, formats, presentations,
filters, grouping fields, sorting fields, duplicate singleton parameters,
control characters, and excessive values fail explicitly.

No default is hidden in a presentation. Environment variables are resolved by
the CLI adapter into explicit selection values before normalization.

## Command Taxonomy and Command Profiles

Two usage models are equally canonical.

### Simple Command Model

Simple commands resolve to immutable predefined profiles:

| Profile | Source | Required stages |
|---|---|---|
| `status` | live | Inventory |
| `check` | live | Inventory → Snapshot |
| `observe` | live | Inventory → Snapshot → Compare → Drift → Health → Rule → Policy → Report |
| `changes` | store | Compare |
| `health` | store | Compare → Drift → Health |
| `report` | store | Compare → Drift → Health → Rule → Policy → Report |

The current `qwsg` invocation without arguments retains contextual help for
backward compatibility. Named profiles provide complete bounded workflows
without exposing internal package knowledge.

### Advanced Command Model and Command Grammar

`qwsg analyze` is the CLI reference mapping for the advanced grammar:

```text
--source live|store
--store DIRECTORY
--from SNAPSHOT --to SNAPSHOT
--pipeline STAGE[,STAGE]
--include STAGE[,STAGE]
--exclude STAGE[,STAGE]
--filter FIELD=VALUE
--group FIELD
--sort FIELD
--output json|human
--presentation structured|terminal
```

Options are a structured serialization of Command Definition 1.0, not
independent switches with private behavior. Simple and advanced requests with
equivalent selections produce the same stage plan and invoke the same
orchestrator.

## Canonical Analysis Pipeline

The permanent order is:

```text
Inventory → Snapshot → Compare → Drift → Health → Rule → Policy → Report
```

Planning computes the transitive dependency closure of every requested stage,
then emits each required stage exactly once in canonical order. An excluded
stage required by the closure is a contradiction and fails before execution.

Stored-snapshot analysis begins at Compare because validated Inventory and
Snapshot evidence already exists. Live analysis begins at Inventory. A live
comparison resolves its baseline before saving the new observation, preserving
previous/current semantics.

The orchestrator stops at the first failed stage and emits an explicit bounded
diagnostic. It never silently skips a required dependency or presents a partial
result as complete.

## Pipeline Orchestration and Engine Adapters

Pipeline stages call only established public boundaries:

- Inventory: validated collector result supplied through the collector adapter;
- Snapshot: Inventory Store validation and optional explicit save;
- Compare: `comparison.Compare`;
- Drift: `drift.Classify`;
- Health: `health.Evaluate`;
- Rule: `rule.Evaluate`;
- Policy: `policy.Evaluate`;
- Report: `report.GeneratePolicy`.

Every stage result records stage name, contract name and version, record count,
completeness, and the unchanged canonical value. Command orchestration does not
reclassify, score, summarize, or otherwise reinterpret it.

The reference `report` profile supplies a versioned deterministic observation
Rule Definition and Policy Profile. Those definitions are configuration
consumed by their canonical engines; command and presentation layers perform
neither Rule matching nor Policy interpretation.

Task 025 adds `policy` as the single compatible Command 1.0 stage extension.
The public `report` profile retains its name and presentation contract while
its dependency closure gains the mandatory Policy stage. Existing Inventory,
Snapshot, Compare, Drift, Health, and Rule stage meanings are unchanged.

Task 027 Scheduler creates no alternative command model. An empty Schedule
Check scope resolves the complete referenced profile through
`command.ResolveProfile`; its one-cycle adapter submits that definition only to
the existing Pipeline Orchestrator. Non-empty Check scope is explicitly
inapplicable until a versioned Command contract can represent it.

Task 028 Alert creates no Command or Pipeline stage. Callers supply already
validated canonical engine outputs directly to the pure Alert decision
boundary. Alert cannot plan commands, execute the Pipeline, change stage order,
or become an alternative orchestration path.

## Command Execution model

Completed execution uses `qwsg.command-execution/1.0`. It includes a
content-derived execution identity, command and plan identities, strictly
ordered stage results, deterministic view rows/groups, explicit diagnostics,
and completeness. Validation recomputes stage order, the projected view, and
identity. Presentation adapters receive this value only after validation.

Filters, grouping, and sorting operate solely on the fixed stage-metadata
registry:

```text
stage contract version record_count complete
```

This model never mutates or filters canonical engine evidence. Command
Definition 1.0 supports repeated conjunctive filters, one grouping field, and
ordered stable sort keys. The registry is versioned and can be extended only
compatibly.

## Human-readable Terminal Presentation

`structured` selects the stable Command Execution contract. `terminal` selects
a localization-ready human view. `json` and `human` select serialization
format. Unsupported values fail rather than falling back silently.

The terminal renderer displays only already projected stage rows and escapes
control characters. JSON serializes the complete canonical execution. Neither
renderer imports pipeline or engine packages.

Future presentation adapters must:

1. create or accept Command Definition 1.0;
2. submit it to the same planner and orchestrator;
3. consume the same validated Command Execution;
4. perform presentation only.

They must not directly invoke collectors, stores, comparison, drift, health,
rule, policy, or report engines.

## Compatibility Strategy

- Command Definition 1.0 and Command Execution 1.0 use exact versions.
- Additive fields require deterministic defaults and old-reader safety.
- New stages, profiles, parameter fields, or semantics require an explicit
  contract version or documented compatible extension.
- Unknown versions and tokens fail closed.
- Existing `inventory`, Snapshot Explorer, and `compare` JSON contracts remain
  unchanged and tested.
- Profile names are public compatibility commitments. Changed behavior requires
  compatibility evidence and an explicit migration.
- Canonical identities are content-derived; presentation choice cannot alter
  engineering behavior.

## Future integrations

### Future Interactive Terminal integration

The future Interactive Terminal edits definitions and renders executions.
Keyboard interaction, panels, navigation, and localization are adapter
concerns.

### Future Dashboard integration

The future Dashboard serializes definitions and visualizes executions. Caching,
layouts, and visualization cannot alter command behavior.

### Future REST API integration

The future REST API transports versioned definitions and executions. Routes,
authentication, authorization, HTTP mapping, and pagination are adapter
concerns.

None of these interfaces may call engineering engines directly.

Task 024 implements none of these future interfaces.

## Security, privacy, and operational boundaries

Command values reject NUL and terminal line controls and are bounded by count
and size. Terminal output escapes control characters. Existing engine privacy
and redaction contracts remain intact because canonical values are not
reinterpreted. Execution is local and explicitly invoked.

The Command architecture adds no daemon, scheduler, monitoring, Alert decision,
Automation, remediation, host mutation, remote execution, network listener,
Dashboard, REST service, Interactive Terminal, AI, or machine learning.

Runtime may inspect a successful Command Execution only through a validated
Scheduler Execution Trace and must validate it against its canonical Command
Plan before projecting exact typed values. Runtime cannot execute stages or
reinterpret Command semantics.

## Operator presentation consumer

The Canonical Operator Presentation Model is a downstream read-only consumer
of validated Command Execution and typed canonical outputs. It does not extend
Command Definition, planning, profiles, pipeline stages, or `Execution.View`.
Future interfaces retain Command Definition/Execution for advanced composition
and use the shared operator model for the beginner overview.
