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
