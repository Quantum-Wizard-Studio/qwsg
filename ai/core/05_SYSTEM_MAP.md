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

```text
Collectors -> Canonical Inventory -> Snapshot Store -> Comparison Engine -> Drift Engine -> Health Engine -> Rule Engine
                                                                                                    |
                                                                                                    v
                                                                 future Policy / Report contracts
                                              |
                 +----------------------------+---------------------------+
                 |                            |                           |
        Community local tools       Professional automation      Provider operations
```

Comparison owns facts about change, Drift owns semantic change classification,
and Health owns deterministic engineering-condition evaluation. Each Health
Record preserves its Drift and Change references. Rule owns deterministic
matching of predefined conditions and preserves Health evidence references.
Health and Rule are pure offline libraries and introduce no scheduler,
monitoring, alert, report, policy, remediation, process, network, or AI
boundary.

Community, Professional, and Provider are orchestration and operating-context
layers above the same versioned core. No dashboard, control plane, notification
service, license service, ecosystem service, or AI adapter may become an
alternative source of engineering truth.

For Slice 1, a local operator invokes the non-root Agent boundary. The discovery coordinator runs bounded read-only collectors, followed by normalization, redaction, validation, inventory assembly, optional latest-envelope persistence, and CLI/JSON presentation. A future Console consumes an Agent-owned redacted contract and cannot access collectors, shell execution, or privilege directly. Installer, remediation, network Console, e-mail, and update boundaries remain outside Slice 1.

The Canonical System Inventory is now the single internal host-information boundary. Host, OS, kernel, CPU, memory, storage, filesystem, network, and virtualization collectors register through the Collector Registry. The coordinator assembles both the authoritative canonical representation and its Inventory 1.0 compatibility projection from the same structured Results. The explicitly invoked Inventory Store can persist and revalidate that synchronized envelope after collection; it has no collector, scheduler, monitoring, or network responsibility. Future Health, Rule, Alert, Policy, Automation, Console, API, and reporting components consume validated canonical inventory and do not query Linux directly.

System evolution follows the permanent boundary:
`Inventory -> Snapshot Store -> Comparison Engine -> Drift Engine`. Comparison
emits deterministic factual Change Records. Drift emits deterministic semantic
Drift Records without judging them. No downstream module may independently diff
Inventory snapshots or reclassify Change Records; future Health, Rules, Policy,
Reports, and Automation consume Drift and Health contracts.

The local administrator invokes `qwsg`. JSON is the machine-compatibility
boundary; the separate terminal renderer exposes safe summaries. Snapshot
Explorer list/info/load operations consume only validated Inventory Store data.
The Makefile installs one binary and creates no runtime service or state.

Verified components and connections belong here. Assumed services, secret endpoints, live credentials, and invented dependencies do not. The map will evolve during development as facts are verified.
