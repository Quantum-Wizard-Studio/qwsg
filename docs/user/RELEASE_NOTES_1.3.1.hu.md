# QWSG Community 1.3.1

Ez a javítókiadás az 1.3.0 valós produkciós frissítésének elfogadásakor feltárt
két hibát javítja. A történeti 1.3.0 kiadás változatlan marad.

- Telepített verzióváltás után a QWSG teljes, hitelesített indexlekérésből
  számolja újra a verziók viszonyát. A HTTP 304 válasz nem tarthat meg korábbi,
  téves „újabb frissítés érhető el” besorolást. A korábban hibásan mentett
  1.3.0 állapot is újraértékelődik: azonos verzióknál current/equal az eredmény.
- A sikeres frissítési értesítés nyilvántartása verzióváltás, sikertelen
  ellenőrzés és újraindítás után is megmarad. Ugyanarról a már sikeresen
  értesített kiadásról önmagában a verzióváltás miatt nem küld újabb levelet.
- Sikertelen újraellenőrzéskor az állapot unknown lesz; a korábbi hitelesített
  bizonyíték és értesítési rekord megmarad, a visszagörgetés elleni védelem él.

A javítás támogatott útja 1.3.0 → 1.3.1. Mivel az 1.3.0 még nem ismeri ezt
az explicit útvonalat, az ellenőrzött 1.3.1 archívum binárisával indítható a
meglévő natív frissítő. A telepítés operátori művelet, visszaállítható;
a konfiguráció, hitelesítő adatok és támogatott állapot megmaradnak.

Korábban elveszett értesítési rekord csak megőrzött, ellenőrzött sikeres
küldési bizonyítékból állítható helyre. A program nem talál ki korábbi küldést.
Az Ed25519-hitelesítés és erőforráskorlátok változatlanok. Nincs telemetria,
regisztráció, új hálózati figyelő vagy automatikus privilegizált frissítés.
