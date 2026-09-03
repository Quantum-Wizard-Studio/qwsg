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
készít elő és nem telepít frissítést, és nem indítja újra a QWSG-t.

Ha az `update.policy` értéke `notify`, továbbá a Community e-mail engedélyezett
és érvényesen beállított, a Guardian egy hitelesített, újabb, alkalmazható
stable kiadás felismerésekor egy tömör kezelői e-mailt küld. Az üzenet megadja
a telepített és elérhető verziót, a stable csatornát, a metaadat hitelesítését,
a kanonikus forrást, valamint azt, hogy a telepítés nem automatikus. A sikeres
SMTP-átvételt a QWSG a hitelesített kiadási verzióhoz, artifact digesthez és
aláíróazonossághoz köti; ezért ugyanaz a kiadás későbbi ellenőrzés és Guardian-
újraindítás után sem küld újabb értesítést. Egy másik hitelesített újabb kiadás
ismét jogosult értesítésre.

Kikapcsolt e-mail vagy `update.policy=manual` mellett nincs kézbesítési kísérlet.
A kézbesítési hiba nem teszi hibássá a Guardiant és nem kerül sikerként
rögzítésre; újrapróbálás csak egy későbbi ütemezett kiadásellenőrzés után
lehetséges. SMTP-jelszó vagy címzett nem kerül a kiadástudatossági állapotba.

Azonnali hitelesített kézi ellenőrzéshez használja a `qwsg update check`
parancsot; ez nem küld frissítési értesítést. A tárolt eredményt hálózati
kapcsolat és értesítés nélkül a `qwsg update status` mutatja. A telepítés
továbbra is a külön, kifejezetten indított `qwsg update` művelet.
