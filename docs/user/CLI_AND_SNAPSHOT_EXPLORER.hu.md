# QWSG felhasználói CLI és Snapshot Explorer

## Első lépések

A QWSG egyszeri, csak olvasható Linux Inventory leltárt készít, és explicit
módon validált snapshotokat menthet, illetve böngészhet. Nem monitoroz, nem
hasonlít össze, nem értékel health állapotot, nem fut daemonként, és nem
módosítja a gazdagépet.

```bash
qwsg help
qwsg version
qwsg help inventory
```

A tartós, felhasználónkénti Guardian konfigurációhoz futtasd a `qwsg setup`
parancsot, majd használd a `qwsg config show|validate|get|set` műveleteket. A
kanonikus leírás: `docs/release/SETUP_AND_CONFIGURATION.md`; a setup nem indítja
el a service-t.

A `0` kilépési kód sikeres vagy teljes Inventory eredményt, a `2` részleges, de
használható Inventory eredményt, az `1` használati, validációs, store-,
jogosultsági, korrupciós vagy futási hibát jelent.

## Élő Inventory

A kompatibilitási alapértelmezés továbbra is JSON:

```bash
qwsg inventory
```

Ez a teljes Inventory 1.0 envelope-ot és az additív `canonical_inventory`
mezőt tartalmazza. Ember számára olvasható adminisztrátori összefoglaló:

```bash
qwsg inventory --format human
```

A human kimenet mindig megjeleníti a státuszt és a megfigyelés idejét. Nem
health verdict. A strukturált collector problémák és redaction adatok a JSON
kimenetben vizsgálhatók.

## Snapshot store kijelölése

A store csak tiszta abszolút privát útvonal lehet. Megadható minden parancsnál:

```bash
qwsg inventory save --store /abszolut/privat/qwsg-inventory
qwsg inventory list --store /abszolut/privat/qwsg-inventory
```

Vagy explicit módon kijelölhető az aktuális shell session számára:

```bash
export QWSG_STORE=/abszolut/privat/qwsg-inventory
export QWSG_FORMAT=human
```

A változók nem keresnek és nem hoznak létre implicit globális store-t. A
`--store` és `--format` parancssori értékek elsőbbséget élveznek.

## Mentés és böngészés

```bash
qwsg inventory save
qwsg inventory list
qwsg inventory info
qwsg inventory load
```

A `save` egyszer gyűjt, és csak validált, privacy-safe Inventory objektumot
ment. A `list` minden megjelenített snapshotot validál, és a neveket a
legrégebbitől a legújabbig rendezi. Az `info` és `load` alapértelmezésben a
legújabb snapshotot használja. A listából kapott pontos név kijelölése:

```bash
qwsg inventory info --snapshot NÉV
qwsg inventory load --snapshot NÉV
```

A `--format json` machine-readable Inventory vagy explorer metadata kimenetet
ad. A retention alapértéke 10, létrehozáskor rögzül, és eltérő értéknél minden
megnyitáskor konzisztensen meg kell adni a `--retention N` opcióval.

A tárolt időpontok a bizonyíték megfigyelésének idejét jelentik. A list, info és
load nem gyűjt, nem javít, nem migrál, nem hasonlít össze, nem töröl, és nem
állítja, hogy az adat aktuális.

## Fail-closed működés

A QWSG elutasítja a relatív vagy nem biztonságos útvonalat, symlinket,
megengedő store jogosultságot, váratlan fájlt, nem támogatott verziót, hibás
JSON-t, duplikált kulcsot, integritási eltérést és érvénytelen Inventory
objektumot. A hibás store-t meg kell őrizni vizsgálatra. Lockot, tranzakciós
artifactot vagy operator snapshotot nem szabad vaktában törölni.
