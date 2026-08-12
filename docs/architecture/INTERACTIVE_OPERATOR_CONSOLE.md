# Interactive Operator Console

Bare Console startup loads Canonical Current Operator State through its
application provider before rendering. It never reads the store directly and
never collects automatically. Explicit refresh repeats only that validated
Current State load, freshness requalification, and render.

## Contract

`Canonical Engineering and Operational Data -> Canonical Operator Presentation Model -> Interactive Operator Console`

The Console imports only `internal/presentationmodel` and the Go standard library. It accepts a validated `Overview`; it does not call or reinterpret engineering engines.

## Console Model 1.0

State contains screen, selection, viewport, terminal capabilities, locale, accepted Overview, bounded diagnostic, refresh count, and quit state. Screens are Home, Attention, Changes, Guardian, Evidence Details, and Help. Home displays the model-owned condition, attention, change and Alert counts, Guardian state, freshness/completeness, and first recommendation. Home and Attention explicitly disclose correlated and omitted attention counts when the bounded list is not exhaustive. Source identifiers appear only in Evidence Details. Every state has a text label; color and ANSI are optional.

## Interaction and refresh

The standard-library session uses injected input, output, context, and provider boundaries. `j`/`k`, Enter, `b`, `r`, `h`, and `q` provide navigation. Output, dimensions, input, diagnostics, and collections are bounded and control characters are replaced.

Refresh is explicit, read-only, and one-shot. The Console calls only
`OverviewProvider`; the application adapter loads and validates Current
Operator State, requalifies freshness at the current time, and returns the
accepted Overview. It does not acquire the Guardian operation lock, resolve or
execute `observe`, collect Inventory, run Pipeline, or publish state. Manual
active observation remains the separate `qwsg observe` command and retains its
existing Guardian lock exclusion. There is no polling, retry, cache,
background worker, or new engine ordering.

Known Runtime failure tokens are localized into practical operator reasons.
Raw errors and unbounded diagnostic detail never enter the Console.

Bare `qwsg` starts the Console only when input and output are terminals. Otherwise it prints the same deterministic Home view and exits. Existing explicit commands, help, JSON, and advanced composition remain independent.

Persistence/recovery, monitoring, notification transports, installation/supervision, REST API, Dashboard, remediation, remote execution, packaging, deployment, and release remain separately governed.
