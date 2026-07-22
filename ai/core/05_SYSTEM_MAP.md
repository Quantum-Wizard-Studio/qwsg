# System Map

## Purpose

This document will map QWSG components, external boundaries, trust relationships, and operational dependencies.

## Status

An internal Slice 1 CLI runtime now exists. It is not a supported release. The product map remains `docs/PRODUCT_SYSTEM_BLUEPRINT.md`, observable behavior remains `docs/FUNCTIONAL_SPECIFICATION.md`, and technical allocation is `docs/architecture/CORE_ALPHA_ARCHITECTURE.md`.

For Slice 1, a local operator invokes the non-root Agent boundary. The discovery coordinator runs bounded read-only collectors, followed by normalization, redaction, validation, inventory assembly, optional latest-envelope persistence, and CLI/JSON presentation. A future Console consumes an Agent-owned redacted contract and cannot access collectors, shell execution, or privilege directly. Installer, remediation, network Console, e-mail, and update boundaries remain outside Slice 1.

The Canonical System Inventory is now the single internal host-information boundary. Host, OS, kernel, CPU, memory, storage, filesystem, network, and virtualization collectors register through the Collector Registry. The coordinator assembles both the authoritative canonical representation and its Inventory 1.0 compatibility projection from the same structured Results; future Health, Rule, Alert, Policy, Automation, Console, API, and reporting components consume validated canonical inventory and do not query Linux directly.

Verified components and connections belong here. Assumed services, secret endpoints, live credentials, and invented dependencies do not. The map will evolve during development as facts are verified.
