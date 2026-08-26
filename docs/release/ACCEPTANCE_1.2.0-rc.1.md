# QWSG 1.2.0-rc.1 Native Update Acceptance

## Status

- Development candidate only; not tagged, published, deployed, or announced.
- Practical native update and rollback acceptance: **PASS**.
- Source commit: `b632acd8ab154e7f174b4dfb0f93f281ea261ffc`.
- External baseline after environment reconciliation: official QWSG `1.1.0`.
- Final external installed state: official QWSG `1.1.0` after explicit rollback.

## Frozen private candidate

- Archive: `qwsg-1.2.0-rc.1-linux-amd64.tar.gz`
- Archive size: `3481005` bytes
- Archive SHA-256:
  `975ada86f82d9b296aa00890d14f26a8f4ba44067f12561802f37106eb2664ac`
- Sidecar SHA-256:
  `fd1ced570713b5127af2a426fff7da2aa41ce6db9cdfc05d6a8414ad39f27387`
- `MANIFEST.sha256` SHA-256:
  `218dff4481619c3de7993f0601e62663b638a9f3a40434b64921ba8e86838b8e`
- Binary SHA-256:
  `2a2ead6a3ec4e566574c03e079fa0e17f0c6784b5c128c14a6f6eef535252a0c`
- `RELEASE.json` SHA-256:
  `53514706aa1a601491a3e7f9c99c7c8a3e656123f552f550c394f5f42404cac9`
- Two isolated canonical builds produced byte-identical archives and sidecars.
  The selected files were frozen mode `0400` and were not changed afterward.

## Package verification

The frozen candidate passed archive and sidecar identity, safe unique archive
layout, regular-file/directory type restrictions, complete internal manifest,
binary version and source provenance, required documentation, modes, platform,
and deterministic source reconciliation. The candidate remains private because
Task 064 does not authorize a tag, Forgejo Release, or publication.

## External lifecycle evidence

The first fail-fast acceptance runner returned without a diagnostic footer.
A bounded read-only classifier proved that no Task 064 mutation occurred and
that the VPS still contained the exact accepted `1.1.0-rc.6` installation from
Task 061, rather than the expected final `1.1.0` installation. Classification:
**ENVIRONMENTAL ISSUE** caused by an incorrect baseline assumption, not a QWSG
product or updater defect.

The corrected bounded runner verified the exact RC.6 identity, anonymously
downloaded and verified the immutable public QWSG 1.1.0 archive and sidecar,
and used the documented replacement workflow to establish the official 1.1.0
baseline without reinstalling or reformatting the VPS. Its private RC.6 package
backup was retained as a recovery precaution.

From that exact baseline, the newer verified archive binary performed the
one-time bootstrap command:

```text
qwsg update --archive <private-frozen-archive> --version 1.2.0-rc.1
```

External results:

- official 1.1.0 identity and integrity: PASS;
- native update `1.1.0 -> 1.2.0-rc.1`: PASS;
- configuration and protected credentials preserved: PASS;
- compatible Guardian/operator state preserved: PASS;
- Guardian service enabled/active intent preserved: PASS;
- post-update configuration validation and readiness: PASS;
- `qwsg update status` rollback availability: PASS;
- explicit `qwsg update rollback`: PASS;
- exact official 1.1.0 binary restored: PASS;
- post-rollback configuration, state, service intent, and readiness: PASS;
- local current-update metadata removed after successful rollback: PASS;
- runner temporary-data cleanup: PASS.

No credential value, address, private provider identity, or private host
identity is contained in this evidence.

## Acceptance interpretation

The Task 064 native update foundation satisfies its intended development
contract. Release discovery can use only published canonical releases; because
this RC is intentionally unpublished, external acceptance correctly used the
private frozen archive path. QWSG 1.2.0-rc.1 is a validated development
candidate, not a published release.
