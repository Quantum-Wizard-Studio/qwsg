# Current Engineering Task 011: Core Inventory Architecture

## Task Metadata

- Task ID: `011`
- Task slug: `core-inventory-architecture`
- Status: `complete`
- Date opened: `2026-07-21` UTC
- Human authority: `Project Owner`
- Owner or lead-developer communication language: `Hungarian`
- Engineering and repository documentation language: `English`

## Priority

★★★★★

## Difficulty

Architecture Critical

## Title

Design and establish the official Inventory Architecture for the Quantum Wizard Server Guardian platform.

## Mission

The Inventory Architecture shall become the common language of the entire QWSG platform.

Every subsystem shall describe the observed system using this architecture before attempting to:

analyse
report
compare
evaluate
secure
repair
automate
or interact with it.

The Inventory Architecture is therefore not another module.

It is the foundation upon which every future module depends.

## Engineering Philosophy

Before a system can be protected,
it must first be understood.

Before it can be understood,
it must first be described.

Before it can be described,
there must exist a common language.

Task 011 creates that language.

## Objectives

Create the official Inventory Architecture specification.

Define the canonical object model.

Define relationships.

Define collector contracts.

Define serialization philosophy.

Define versioning strategy.

Define compatibility rules.

Define resource efficiency requirements.

Define future extensibility.

## Required Reading

ai/core/00_PROJECT_PHILOSOPHY.md

ai/core/01_CONSTITUTION.md

ai/core/03_AGENTS.md

ai/core/08_JOB_TEMPLATE.md

ai/core/11_ENGINEERING_LIFECYCLE.md


## New Required Core Document

Create

ai/core/12_INVENTORY_ARCHITECTURE.md

This becomes the official specification.

Future modules shall reference this document instead of redefining their own structures.

## Core Philosophy

Inventory is NOT

monitoring
logging
metrics
diagnostics

Inventory IS

A versioned digital description of the observed system.

## Digital Twin Principle

The Inventory represents a Digital Twin.

Not a simulation.

Not a virtual machine.

A structured, versioned description of reality.

Every collector contributes one piece.

The complete Inventory represents the system.

## Canonical Inventory Layers

The architecture shall define at minimum:

Host

Hardware

Operating System

Runtime

Storage

Network

Services

Applications

Users

Security

Policies

Metadata

Additional layers must remain fully extensible.

## Collector Contract

Every collector shall expose a common contract.

Minimum required fields:

Collector Name

Version

Capability

Supported Platforms

Execution Time

Timestamp

Health Status

Warnings

Errors

Collected Data

Metadata

Collectors shall never invent incompatible formats.

## Versioning

Inventory schema shall be versioned.

Example:

Inventory Schema

1.0

1.1

2.0

Collectors shall explicitly declare compatibility.

## Serialization

Architecture shall define

canonical JSON output.

Future formats may include:

YAML

MessagePack

Protocol Buffers

without changing the Inventory model.

## Resource Efficiency Contract

The Inventory Architecture must explicitly require resource-efficient execution.

The platform shall be designed for:

minimal CPU usage
minimal RAM usage
minimal disk activity
minimal process lifetime

## Core Rules

Stateless by default.

Short-lived execution preferred.

Daemon optional.

Every collector measurable.

Every collector interruptible.

Every collector timeout-aware.

No collector may allocate excessive memory.

No collector may continuously poll unless explicitly required.

## Execution Philosophy

Preferred execution model:

systemd timer

↓

oneshot service

↓

Inventory collection

↓

Policy evaluation

↓

Report generation

↓

Exit

Long-running daemons are optional future extensions.

The Core shall never depend on a daemon.

## Architecture Rules

Future modules MUST consume the Inventory.

They MUST NOT create independent host models.

They MUST NOT redefine

CPU

Memory

Storage

Network

Service

objects.

The Inventory is the single source of truth.

## Policy Compatibility

Future Policy Engine shall operate exclusively on Inventory objects.

Policies must never directly query Linux.

Collectors collect.

Policies evaluate.

## Reporting Compatibility

Reports shall be generated exclusively from Inventory.

No report generator may perform direct system discovery.

## REST API Compatibility

REST endpoints shall expose Inventory objects.

The API must not invent alternative structures.

## AI Compatibility

Future AI assistants shall consume Inventory.

The AI must never depend directly on shell commands.

Inventory becomes the AI context.

## Internationalization Contract

Internal engineering language:

English.

All user-visible messages shall be translatable.

Requirements:

English canonical
Hungarian official
unlimited future languages

No user-facing strings may be hardcoded.

## Documentation Contract

A feature is NOT complete without documentation.

Documentation must exist in:

English

Hungarian

Minimum documentation:

Overview

Installation

Configuration

Usage

Architecture

Troubleshooting

Upgrade

Removal

Developer Notes

## Future Compatibility

Task 011 prepares the foundation for:

Task 012

Collector Framework

Task 013

Platform Hardening

Task 014

Configuration Engine

Task 015

Policy Engine

Task 016

Reporting Engine

Task 017

REST API

## Out of Scope

Task 011 shall NOT implement:

new collectors

daemon

REST API

Policy Engine

Dashboard

Repair engine

Web interface

Automation

Task 011 defines architecture only.

## Deliverables

Mandatory:

12_INVENTORY_ARCHITECTURE.md

Architecture diagrams.

Inventory hierarchy.

Collector contract.

Versioning specification.

Serialization rules.

Resource contract.

Internationalization rules.

Documentation requirements.

Future compatibility rules.

## Verification

Aiko shall verify:

✔ architecture internally consistent

✔ no duplicated object definitions

✔ future modules compatible

✔ documentation complete

✔ no conflicts with Constitution

✔ no conflicts with Engineering Lifecycle

## Documentation Updates

Update:

Engineering history

Project roadmap

Architecture references

if required.

## Completion Criteria

Task 011 is complete only when:

Inventory Architecture exists.

Collector contract defined.

Digital Twin philosophy documented.

Resource contract documented.

Internationalization documented.

Documentation contract documented.

Architecture approved.

No unresolved contradictions remain.

## Closing Principle

Every component shall describe the system in a common language before attempting to analyze, report, modify, or protect it.

## Quantum Wizard Principle

A Guardian cannot protect what it does not understand.

The Inventory is how the Guardian understands the world.
