# System Map

## Purpose

This document will map QWSG components, external boundaries, trust relationships, and operational dependencies.

## Status

A supported pre-alpha one-shot user CLI and Snapshot Explorer now fronts the
implemented Inventory boundaries. The product map remains
`docs/PRODUCT_SYSTEM_BLUEPRINT.md`, observable behavior remains
`docs/FUNCTIONAL_SPECIFICATION.md`, and technical allocation is documented
under `docs/architecture/`. The canonical long-term ecosystem, edition, user
experience, deployment, licensing, privacy, automation, and AI boundaries are
defined in `docs/PRODUCT_ARCHITECTURE.md`.

All editions share one deterministic engineering flow:

Every presentation enters that flow through the Canonical Command Architecture:

```text
CLI / future Interactive Terminal / future Dashboard / future REST API
                              |
                    Command Definition 1.0
                              |
                    Canonical Command Plan
                              |
                    Pipeline Orchestration
                              |
 Inventory -> Snapshot -> Compare -> Drift -> Health -> Rule -> Report
                              |
                    Command Execution 1.0
                              |
                 Replaceable presentation
```

Only `internal/pipeline` orchestrates engines. `internal/command` owns the
presentation-neutral public request, plan, profile, parameter, projection, and
execution contracts. `internal/presentation` renders completed executions and
cannot call the pipeline or engines. `cmd/qwsg` is the first adapter, not an
alternative source of command behavior.

```text
Collectors -> Canonical Inventory -> Snapshot Store -> Comparison Engine -> Drift Engine -> Health Engine -> Rule Engine -> Report Engine
                                                                                                    |                    |
                                                                                                    v                    v
                                                                                          future Policy        future views / exports
                                              |
                 +----------------------------+---------------------------+
                 |                            |                           |
        Community local tools       Professional automation      Provider operations
```

Comparison owns facts about change, Drift owns semantic change classification,
and Health owns deterministic engineering-condition evaluation. Each Health
Record preserves its Drift and Change references. Rule owns deterministic
matching of predefined conditions and preserves Health evidence references.
Report owns deterministic presentation contracts and preserves Rule Evaluation
and Health evidence references. Health, Rule, and Report are pure offline
libraries and introduce no scheduler, monitoring, alert, policy, remediation,
process, delivery, network, or AI boundary.

Community, Professional, and Provider are orchestration and operating-context
layers above the same versioned core. No dashboard, control plane, notification
service, license service, ecosystem service, or AI adapter may become an
alternative source of engineering truth.

For Slice 1, a local operator invokes the non-root Agent boundary. The discovery coordinator runs bounded read-only collectors, followed by normalization, redaction, validation, inventory assembly, optional latest-envelope persistence, and CLI/JSON presentation. A future Console consumes an Agent-owned redacted contract and cannot access collectors, shell execution, or privilege directly. Installer, remediation, network Console, e-mail, and update boundaries remain outside Slice 1.

The Canonical System Inventory is now the single internal host-information boundary. Host, OS, kernel, CPU, memory, storage, filesystem, network, and virtualization collectors register through the Collector Registry. The coordinator assembles both the authoritative canonical representation and its Inventory 1.0 compatibility projection from the same structured Results. The explicitly invoked Inventory Store can persist and revalidate that synchronized envelope after collection; it has no collector, scheduler, monitoring, or network responsibility. Future Health, Rule, Alert, Policy, Automation, Console, API, and reporting components consume validated canonical inventory and do not query Linux directly.

System evolution follows the permanent boundary:
`Inventory -> Snapshot Store -> Comparison Engine -> Drift Engine -> Health Engine -> Rule Engine -> Report Engine`. Comparison
emits deterministic factual Change Records. Drift emits deterministic semantic
Drift Records without judging them. Health evaluates engineering condition,
Rule matches fixed conditions, and Report presents their canonical evidence.
No downstream module may independently diff Inventory snapshots, reclassify
Change Records, re-evaluate Health or Rules, or rebuild engineering summaries.

The local administrator invokes `qwsg`. Command Execution 1.0 is the canonical
machine boundary; JSON serializes it and the separate terminal renderer exposes
safe summaries. Snapshot
Explorer list/info/load operations consume only validated Inventory Store data.
The Makefile installs one binary and creates no runtime service or state.

Verified components and connections belong here. Assumed services, secret endpoints, live credentials, and invented dependencies do not. The map will evolve during development as facts are verified.
