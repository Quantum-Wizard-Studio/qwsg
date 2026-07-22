# Roadmap

## Purpose

This document will order approved outcomes for the safe development of QWSG.

## Status

Milestone 0 and engineering-governance updates E001–E002 are complete. Tasks 003, 005, 006, and 007 established the product, blueprint, functional, and audit baseline. Task 008 established the Core Alpha architecture and the first implementation milestone: `Core Alpha Slice 1: Read-only Server Discovery and System Inventory`. No supported product release exists.

Task 009 completed the bounded Slice 1 implementation using Go and the stable `0/1/2` complete/fatal/partial exit policy. Task 010 established the automated task lifecycle and did not promote Slice 1 to a supported release. The Functional Specification and architecture gate register remain controlling for release blockers.

Task 011 established the platform-wide Inventory Architecture as the common digital-twin language for collectors, policy, reporting, APIs, Console, automation, and AI consumers. Task 012 implemented the internal Collector Framework against that contract while preserving the bounded legacy `1.0` Inventory projection: descriptors, capabilities, dependencies, Registry-only execution, deterministic planning, availability checks, timeouts, cancellation, and failure isolation are covered by tests. Task 013 established the official Engineering Task Builder. Task 014 implemented Canonical System Inventory v1 with nine production-oriented Linux collectors and a validated authoritative canonical representation while preserving Inventory 1.0 compatibility. No supported release was introduced. Tasks 015–017 remain future, separately authorized milestones.

Approved goals, dependencies, risks, and acceptance criteria belong here. The Product Definition is the parent product record for future roadmap decisions after owner review. Fixed promises, speculative dates, credentials, and implementation performed without a task do not. The roadmap will evolve during development under owner direction.
