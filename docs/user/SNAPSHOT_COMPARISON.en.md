# Snapshot Comparison

Set the private snapshot store used by the Snapshot Explorer:

```sh
export QWSG_STORE=/absolute/private/qwsg-inventory
```

Save observations at different times, then compare the latest two:

```sh
qwsg inventory save
qwsg inventory save
qwsg compare
```

JSON is the default, canonical output. Select an exact pair using names from
`qwsg inventory list`:

```sh
qwsg compare --from SNAPSHOT_1.json --to SNAPSHOT_2.json
qwsg compare --from SNAPSHOT_1.json --to SNAPSHOT_2.json --format human
```

The human report groups Added, Removed, Modified, and Unchanged facts. A
successful comparison exits `0`, even when there are no changes. Exit `1`
indicates invalid options, insufficient history, inaccessible or corrupt data,
incompatible snapshots, or an output failure.

The report describes observed differences only. It is not a health verdict,
drift decision, alert, score, recommendation, or proof of current host state.

## Protected service identity and older snapshots

Running services now have stable protected IDs. Reordering leaves their identity
unchanged; a service appearing, disappearing or replacing another remains visible
even when the count stays equal. Names stay redacted. These deterministic IDs are
pseudonymous, not anonymous: candidate names can be guessed and hashed, and equal
names can be correlated. Keep snapshots private.

A comparison involving an older, nonempty ordinal service list stops with
`service identity comparison unavailable` (exit `1`), rather than claiming exact
identity changes. Reports requiring that comparison also stop. Older snapshots
remain readable and are not migrated or rewritten. With the corrected collector,
run the two `qwsg inventory save` commands above and compare those new snapshots;
use explicit selectors if needed. No historical evidence needs deletion.
