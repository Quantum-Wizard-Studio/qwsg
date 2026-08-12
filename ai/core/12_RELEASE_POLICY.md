# Release Policy

## Purpose

This policy defines how versions of Quantum Wizard Server Guardian will be prepared, verified, documented, and recoverably released.

## Version 1.0 policy

QWSG uses semantic release identities in `VERSION`; the same value is embedded
by linker flags and used in archive names, manifests, release notes and the
changelog. A candidate is `1.0.0-rc.N`; the final identity is `1.0.0`. Task 044
records the Owner decision authorizing that identity, while tagging and
publication remain separately authorized actions.

The supported 1.0 artifact is one deterministic
`qwsg-VERSION-linux-amd64.tar.gz` plus a SHA-256 file. It contains the prebuilt
binary, fixed `/usr/local` systemd user unit, bounded install/uninstall scripts,
configuration example, license, changelog and ordinary-user documentation.
Controlled inputs are VERSION, source tree, hexadecimal BUILD_COMMIT (or the
documented `unknown` sentinel), and SOURCE_DATE_EPOCH. Assembly uses
linux-amd64, CGO disabled, `-trimpath`, disabled VCS stamping, stable ordering,
timestamps, owners and gzip metadata. Two identical builds must match.

SHA-256 detects changes after a trusted hash was obtained; it does not
authenticate the publisher. Signing, key custody, tagging and publication are
separate Owner-authorized operations.

## Release gates

A technical 1.0 release decision requires full/race/vet/format tests, Framework
and lifecycle validation, archive inspection, clean staged installation, real
isolated Guardian lifecycle, valid-state compatibility, fail-closed unsafe
state tests, upgrade/rollback/uninstall, privacy/security review, and a bounded
50-cycle endurance run. Support claims cover only environments with real
evidence. Unperformed physical reboot acceptance remains an Owner-run
pre-publication gate.

Valid Current Operator State 1.0/1.1/1.2, Scheduler State 1.0, Guardian
Checkpoint 1.0 and Configuration Source 1.0 are preserved. Unknown formats are
never silently migrated. Rollback restores exact artifacts while preserving
private state; incompatibility leaves the Guardian stopped.

QWSG 1.0 is governed by the Owner-approved QWS Community / Free License Version
1.0, a proprietary source-available Community license rather than an OSI
open-source license. The accepted local Community Guardian requires no paid
license or API entitlement. Tagging and external publication remain separate
Owner-authorized operations.
