# Automatikus kiadásellenőrzés

A QWSG Guardian automatikusan ellenőrzi a hivatalos Community kiadási indexet,
amikor a kiadástudatossági ellenőrzés esedékes. Az alapértelmezett időköz 24
óra. Első használatkor a hiányzó állapot miatt az ellenőrzés a Guardian első
helyi megfigyelési ciklusának befejezése után esedékes. Később az új ellenőrzés
az utolsó rögzített próbálkozás után 24 órával válik esedékessé. A határidő
előtti újraindítás nem ismétli meg a kérést, hanem kivárja a hátralévő időt.

A hitelesítő adat nélküli kérés kizárólag a beállított hivatalos QWSG HTTPS
kiadási indexhez kapcsolódik. A QWSG megköveteli az előírt médiatípust,
ellenőrzi az aláírt indexet a beépített Ed25519 bizalmi kulccsal, és alkalmazza
a visszaállítási, illetve jövőbeli-index védelmet. Nem küld telepítési
azonosítót, gépnevet, leltárt, fiókazonosítót, API-kulcsot, e-mail-címet vagy
telemetriai adatot. Mint minden HTTPS-kapcsolatnál, a célrendszer és a hálózat
láthatja a szokásos kapcsolati metaadatokat, például a forrás IP-címét.

Hálózati vagy hitelesítési hiba esetén csak egy korlátozott helyi hibakategória
kerül rögzítésre. A Guardian helyi felügyelete folytatódik, azonnali ismétlési
hurok nincs. Az automatikus ellenőrzés soha nem tölt le kiadási csomagot, nem
készít elő és nem telepít frissítést, nem indítja újra a QWSG-t, és nem küld
értesítést. A frissítési értesítés és deduplikáció a jövőbeli Task 080 része.

Azonnali hitelesített kézi ellenőrzéshez használja a `qwsg update check`
parancsot. A tárolt eredményt hálózati kapcsolat nélkül a `qwsg update status`
mutatja. A telepítés továbbra is a külön, kifejezetten indított `qwsg update`
művelet.
