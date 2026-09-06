# Community 1.3.1 — produkciós elfogadás

A Task 081 javítókiadása a valós 1.3.0 → 1.3.1 operátori frissítéssel került
telepítésre. A történeti 1.3.0 kiadás változatlan: annak elfogadásakor feltárt
két hibát az új 1.3.1 javítja.

- Az éles, aláírt 1.3.1 indexből egy értesítés keletkezett; ismételt
  ellenőrzéskor nem történt újabb küldés. Az SMTP-szolgáltató elfogadta a levelet.
- Frissítés és szolgáltatás-újraindulás után a sikeres értesítési rekord megmaradt.
- Két valós HTTP 304 válasz után is current/equal maradt az 1.3.1 állapot.
- A konfiguráció és hitelesítő adatok változatlanok. A helyi állapotlekérdezés
  nem használ hálózatot. A Guardian és Scheduler erőforráskorlátai megmaradtak.
- A telepítés operátori művelet volt; nem történt automatikus frissítés.
  A natív rollback rendelkezésre áll, determinisztikus izolált próbája sikeres.

A pre-install próbát az ellenőrzött 1.3.1 Guardian-kód korlátozott futtatása
végezte a valós production állapoton; ez nem a régi 1.3.0 napi időzítőjének
elfogadási bizonyítéka. A részletes eredmények és pontos kiadásazonosítók az
[elfogadási jegyzőkönyvben](../release/ACCEPTANCE_1.3.1.md) találhatók.
