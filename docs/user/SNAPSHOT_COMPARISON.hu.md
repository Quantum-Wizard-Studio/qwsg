# Snapshot összehasonlítás

Állítsd be a Snapshot Explorer által használt privát tárolót:

```sh
export QWSG_STORE=/abszolut/privat/qwsg-inventory
```

Ments különböző időpontokban megfigyeléseket, majd hasonlítsd össze a két
legutóbbit:

```sh
qwsg inventory save
qwsg inventory save
qwsg compare
```

Az alapértelmezett, kanonikus kimenet JSON. Az `qwsg inventory list` által
kiírt nevekkel tetszőleges pár választható:

```sh
qwsg compare --from SNAPSHOT_1.json --to SNAPSHOT_2.json
qwsg compare --from SNAPSHOT_1.json --to SNAPSHOT_2.json --format human
```

Az emberi jelentés Added, Removed, Modified és Unchanged csoportokat mutat.
A sikeres összehasonlítás kilépési kódja `0` akkor is, ha nincs változás. Az
`1` hibás opciót, elégtelen előzményt, elérhetetlen vagy sérült adatot,
inkompatibilis snapshotokat vagy kimeneti hibát jelent.

A jelentés kizárólag megfigyelt eltéréseket ír le. Nem egészségítélet,
drift-döntés, riasztás, pontszám, ajánlás vagy a gép jelenlegi állapotának
bizonyítéka.
