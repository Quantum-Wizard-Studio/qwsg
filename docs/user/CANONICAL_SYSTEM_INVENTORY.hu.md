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

A Task 014 nem vezet be felhasználói konfigurációt vagy collector-választó felületet. A gyűjtés egyszeri, minden collector véges idő- és kimeneti korláttal rendelkezik, és nincs történeti tárolás. Az ütemezés és a perzisztencia külön jövőbeli képesség.

## Hibaelhárítás és frissítés

A `permission_denied`, `unavailable`, `unsupported`, `timeout`, `cancelled`, `resource_limit` és `error` az adatgyűjtést írja le, nem a szerver egészségi állapotát. Csak indokolt esetben módosítsd a gazdagép hozzáférési határait; ne futtasd rootként pusztán a részleges eredmény elfedéséhez. A sémafrissítések explicit, verziózott migrációt követnek. Ez a verzió nem perzisztálja a leltárt, ezért megőrzött adatok automatikus eltávolítása sem történik.
