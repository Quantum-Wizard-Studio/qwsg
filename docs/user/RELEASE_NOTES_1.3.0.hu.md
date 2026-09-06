# QWSG Community 1.3.0

Az 1.3.0 hitelesített frissítésfigyelést ad a helyi Guardianhoz. A hivatalos
kiadási metaadat Ed25519-aláírását a beépített Community bizalmi kulccsal
ellenőrzi, a telepített bináris és csomag azonosságát külön vizsgálja.

- A `qwsg update check` kézzel frissíti az elérhető kiadás adatait.
  A `qwsg update status` csak helyi adatot olvas, hálózatot nem használ.
- A Guardian 24 óránként, esedékesség esetén elkülönített, korlátozott
  feladatként ellenőrzi a hivatalos kiadásokat. Az ellenőrzés hibája nem
  állítja le a megfigyelést.
- Bekapcsolt értesítési házirend és beállított Community SMTP mellett a
  hitelesített, támogatott újabb kiadásról e-mail küldhető. A sikeres küldés
  tartós nyilvántartása megakadályozza az ismétlést újabb ellenőrzés és
  Guardian-újraindítás esetén is.
- Az ellenőrzés kizárólag metaadatot tölt le. A csomag letöltését és
  telepítését továbbra is az operátor indítja a `qwsg update` paranccsal.
- A Scheduler legfeljebb 64 eredményt őriz. A 8 MiB-nál nagyobb állapotot
  feldolgozás előtt elutasítja. A Guardian 128 MiB memóriakorlátja és
  32 feladatos korlátja változatlan.

A támogatott frissítési út: 1.2.0 → 1.3.0. A konfiguráció, a hitelesítő
adatok, a támogatott helyi állapot és a szolgáltatás beállításai megmaradnak;
csomagszintű visszaállítás rendelkezésre áll. A túlméretes vagy ismeretlen
Scheduler-állapot nem törlődik automatikusan: biztonsági mentés után rendezni
kell, mielőtt az ütemezett működés egészségesnek tekinthető.

Az első produkciós elfogadás telepített 1.2.0 példánya külön engedélyezett
migrációs kompatibilitási javítást kapott. A történeti hivatalos 1.2.0 csomag
változatlan; a javított példányt a kanonikus 1.3.0 váltja fel.

Nincs telemetria, regisztráció, API-kulcs- vagy telepítésazonosító-követelmény,
új bejövő hálózati figyelő, automatikus csomagletöltés a frissítésfigyelésből,
vagy automatikus frissítéstelepítés.
