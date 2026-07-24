# Kanonikus rendszerleltár

## Áttekintés

A Kanonikus rendszerleltár a QWSG csak olvasható leírása az aktuális Linux gazdagépről. Megfigyelt gazdagép-, operációsrendszer-, kernel-, CPU-, memória-, tároló-, fájlrendszer-, hálózati és virtualizációs tényeket rögzít. Nem értékeli az egészségi állapotot, nem futtat szabályzatot, nem figyel folyamatosan, és nem módosítja a szervert.

## Előfeltételek és jogosultság

A QWSG-t normál felhasználóként, Linux rendszeren kell futtatni. Root jogosultság és hálózati hozzáférés nem szükséges. Egyes opcionális adatok a gazdagép jogosultságai vagy izolációja miatt nem érhetők el; a QWSG ezt jelzi a jogosultság automatikus emelése helyett.

## Használat

Fordítás: `make build`, majd futtatás:

```bash
build/qwsg inventory
```

A parancs JSON-t ír a szabványos kimenetre. A `0` kilépési kód teljes, a `2` részleges, de használható, az `1` sikertelen leltárt jelent. A felső szintű Inventory 1.0 nézet kompatibilitási céllal megmarad. Az új integrációk a `canonical_inventory` mezőt használják.

## Adatvédelem és értelmezés

A gazdagépnevek, hálózati és hardvercímek, interfésznevek, csatolási útvonalak, nyers eszköznevek, gépazonosítók és szolgáltatásazonosítók kimaradnak, maszkoltak, vagy adatvédelmi szempontból biztonságos azonosítót kapnak. A hiányzó vagy maszkolt érték nem jelent üres, hamis vagy egészséges állapotot. Minden collector és réteg státuszát, valamint a strukturált problémákat is ellenőrizni kell.

## Konfiguráció és korlátok

A gyűjtés továbbra is egyszeri, és minden collector véges idő- és
kimeneti korláttal rendelkezik. Egy validált eredmény explicit mentéséhez adj
meg egy abszolút, privát könyvtárat:

```bash
build/qwsg inventory save --store /abszolut/privat/qwsg-inventory
build/qwsg inventory load --store /abszolut/privat/qwsg-inventory
```

A store alapértelmezésben a legutóbbi 10 snapshotot őrzi meg. Létrehozáskor az
`--retention N` opcióval 1 és 1000 közötti fix érték adható meg; megnyitáskor
ugyanezt kell használni. A könyvtárak jogosultsága `0700`, a fájloké `0600`.
A mentés és visszatöltés ugyanazt az Inventory JSON-t írja ki, és megőrzi a
státusz kilépési kódját, beleértve a részleges, de használható eredmény `2`
kódját.

A perzisztencia kézi művelet. Nem vezet be monitoringot, összehasonlítást,
health scoringot, daemont, ütemezést, riasztást, értesítést, hálózati
szolgáltatást, adatbázist vagy feltöltést.

## Hibaelhárítás és frissítés

A `permission_denied`, `unavailable`, `unsupported`, `timeout`, `cancelled`,
`resource_limit` és `error` az adatgyűjtést írja le, nem a szerver egészségi
állapotát. Csak indokolt esetben módosítsd a gazdagép hozzáférési határait; ne
futtasd rootként pusztán a részleges eredmény elfedéséhez.

A store elutasítja a nem biztonságos jogosultságot vagy útvonalat, a nem
támogatott verziót, a hibás JSON-t, a duplikált kulcsot, az integritási eltérést
és az érvénytelen Inventory objektumot. A checksum korrupciót észlel, de nem
kriptográfiai aláírás. Elakadt lock vagy tranzakciós fájlt ne törölj kézzel a
store ellenőrzése nélkül. A sémamigráció és a hitelesített tárolás későbbi
képesség.
