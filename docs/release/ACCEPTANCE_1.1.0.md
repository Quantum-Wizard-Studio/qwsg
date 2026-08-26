# QWSG 1.1.0 Final Release Acceptance

## Release verdict

- Final identity: `QWSG 1.1.0`
- Release-source commit:
  `305f4088e94b14d6cbb3114eb8cce4e32d847c16`
- Accepted behavioral source:
  `25a30718bc92882e9773a5c405ad648c0eee1a81`
- Task 061 practical acceptance: **PASS — REUSED**
- QWSG-059-F001: **EXTERNALLY CORRECTED**
- Product defects in accepted RC.6/final behavior: **NONE**
- Final release readiness: **READY FOR RELEASE**
- Publication: **PASS — QWSG 1.1.0 RELEASED**

The final release-source diff changes release identity, documentation,
deterministic release plumbing and one Framework test fixture only. It changes
no QWSG runtime source, service unit, installer, uninstaller, packaged runtime
configuration or product behavior. Task 061 external evidence therefore remains
valid without repeating practical acceptance.

## Frozen artifact identity

Two isolated `git archive` exports of the exact release-source commit were
built with independent source, output and Go cache locations using source epoch
`1787752902` (`2026-08-26T14:01:42Z`). Binary, internal manifest, archive and
sidecar bytes are identical.

- Archive: `qwsg-1.1.0-linux-amd64.tar.gz`
- Archive size: `2951638` bytes
- Archive SHA-256:
  `10a39d96b93b72a3f4799a76d769bc264afd6845a32a1ecc5531b062d6f42349`
- Sidecar: `qwsg-1.1.0-linux-amd64.tar.gz.sha256`
- Sidecar SHA-256:
  `b9414bba5a6d9bc100f7c391c11867ceda6c2139002272a0f46fbf55dc9d3cc1`
- `MANIFEST.sha256` SHA-256:
  `310d41f9a8c71599290fd1d25efb7a2da8fd210e34cbf2666e40189c988ebc3d`
- Binary SHA-256:
  `e7b5a2234221baa32a9c3fa79e0758ea49e7ed1996c99e9e1ddbc19628a5e924`
- Binary provenance: `QWSG 1.1.0`, full release-source commit, controlled UTC
  build time above
- Frozen local mode: `0400` for archive and sidecar inside a private mode-0700
  Task 063 build root

Candidate bytes are frozen and must not be rebuilt, substituted or modified.

## Package verification

- twin archive byte identity: **PASS**
- twin checksum-sidecar byte identity: **PASS**
- sidecar verification: **PASS**
- safe unique 25-member single-root archive: **PASS**
- regular-file/directory-only type boundary and no symlinks: **PASS**
- all 18 internal manifest entries: **PASS**
- required LICENSE, README, INSTALL, final notes and operator docs: **PASS**
- executable/document modes: **PASS**
- static linux-amd64 binary: **PASS**
- exact version/full commit/build-time provenance: **PASS**
- ambient Go VCS metadata absent: **PASS**
- unchanged installer/uninstaller/package contracts: **PASS**

The extracted local Smart Install assessment satisfied OS, architecture,
glibc, non-root and local-filesystem checks but correctly reported the current
development session's unreachable systemd user manager as required/not-ready.
This is an **ENVIRONMENTAL ISSUE**, not a product defect: the same environment
also emits the known static systemd bus warning. Task 061 externally proved
Smart Install ready on supported Ubuntu 24.04 amd64, and no relevant behavior
changed. That external PASS is reused.

## Reused Task 061 evidence

| Behavior | Final 1.1.0 basis |
|---|---|
| Ubuntu 24.04 amd64 Smart Install | Task 061 PASS; unchanged |
| documented installation/replacement | Task 061 PASS; unchanged artifacts/scripts |
| configuration/state/credential safety | Task 061 PASS; unchanged |
| physical reboot and pre-login autostart | Task 061 PASS; unchanged |
| systemd automatic recovery | Task 061 PASS; unchanged |
| current configuration/generation ownership | Task 061 PASS; unchanged |
| fresh canonical Guardian evidence | Task 061 PASS; unchanged |
| controlled SMTP and actual Owner receipt | Task 061 PASS; unchanged |
| notification/overall partial semantics | expected product behavior; unchanged |

Task 061's Owner-terminated explicit-restart and uninstall/reinstall items
remain disclosed evidence limitations; they are not newly invented release
requirements and no observed product defect follows from them.

## Mandatory pre-publication gates

| Gate | Result |
|---|---|
| canonical Framework/Git/source lineage | PASS |
| rollback-capable execution snapshot | PASS |
| runtime behavior unchanged from accepted RC.6 | PASS |
| focused and proportional full local validation | PASS |
| deterministic twin construction | PASS |
| frozen artifact integrity and package safety | PASS |
| Task 061 acceptance reuse | PASS |
| security/privacy review | PASS |
| exact release-source push and direct verification | PASS |
| release-boundary snapshot | PASS |
| annotated tag identity/push | PASS |
| final Forgejo Release and exact two assets | PASS |
| repository anonymous readability | PASS |
| anonymous Release-page accessibility | PASS |
| anonymous wget/curl artifact verification | PASS |

Release-boundary snapshot `/tmp/qwsg-task063-release-boundary.tkqAUW` passed.
Annotated tag object `b14347636f6c9873a5acf759c950d900a39bf1a7` is pushed and
peels exactly to the release-source commit. Forgejo Release `QWSG 1.1.0`
exists as final/non-prerelease and contains exactly the frozen archive and
sidecar with expected sizes.

Following explicit Owner authorization, the canonical repository became public.
Anonymous repository and final Release pages return successfully. Independent
clean-environment `wget` and `curl -fLO` downloads retrieved the exact archive
and sidecar; both archive copies are `2951638` bytes with SHA-256
`10a39d96b93b72a3f4799a76d769bc264afd6845a32a1ecc5531b062d6f42349`,
both sidecars have SHA-256
`b9414bba5a6d9bc100f7c391c11867ceda6c2139002272a0f46fbf55dc9d3cc1`,
and `sha256sum -c` passes for both clients. Tag object and peel remain unchanged,
the Release remains final/non-draft/non-prerelease, and its asset count remains
exactly two. No release byte changed.

Authoritative release page:
`https://git.quantumwizard.hu/Quantum_Wizard_Studio/qwsg/releases/tag/v1.1.0`.
This evidence establishes **QWSG 1.1.0 RELEASED**. It does not authorize mutable
or replacement assets, production deployment, or an unrelated announcement.
