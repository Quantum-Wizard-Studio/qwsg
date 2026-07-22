# Historical QWSG Master Plan

> Status: preserved non-normative source material. Authoritative product, behavior, and architecture records are `docs/PRODUCT_DEFINITION.md`, `docs/PRODUCT_SYSTEM_BLUEPRINT.md`, `docs/FUNCTIONAL_SPECIFICATION.md`, and `docs/architecture/CORE_ALPHA_ARCHITECTURE.md`. Examples and estimates below are not implementation commitments.

Sőt... szerintem ez kötelező egy éles szerveren.

Nem csak 80%-nál. Én három szintet használnék:

Foglaltság	Jelzés	Teendő
80%	⚠️ Figyelmeztetés	E-mail
90%	🚨 Kritikus	E-mail + napló
95%	☠️ Vészhelyzet	E-mail óránként, amíg meg nem oldódik
Én ezt még tovább vinném

Mivel már van működő Exim és Hestia, írnék egy nagyon egyszerű scriptet:

ellenőrzi a / partíció foglaltságát;
opcionálisan a /backup könyvtár méretét;
opcionálisan a memóriahasználatot;
opcionálisan a load average-et;
küld egy szép formázott e-mailt;
nem küld 5 percenként új levelet, hanem csak akkor, ha állapotváltozás történik (pl. 79→80%, vagy 89→90%).

Ez utóbbi fontos, különben elárasztja a postaládát.

Én viszont ennél is többet csinálnék

Az <example owner> szerveren már van:

Exim
Hestia
cron

Ezért én készítenék egy Server Health Monitor csomagot.

Naponta egyszer (pl. reggel 8-kor) küldene egy rövid státuszjelentést:

Server: <example-host>

✓ CPU: 6%
✓ RAM: 31%
✓ Disk: 38%
✓ SSL: OK
✓ Hestia: OK
✓ MariaDB: OK
✓ Apache: OK
✓ Nginx: OK
✓ PHP-FPM: OK
✓ Exim: OK
✓ Dovecot: OK

Backup:
✓ Utolsó backup: 2026-07-19 03:30
Méret: 5.42 GB

Nincs hiba.

Ha bármi gond van, akkor ugyanebből:

⚠ FIGYELEM!

Disk usage: 91%

Backup sikertelen!

MariaDB nem fut!

Dovecot leállt!
Ezután már csak a fontos dolgokról jönne levél

Például:

80/90/95% tárhely
backup sikertelen
MariaDB leállt
Apache leállt
Nginx leállt
Hestia nem válaszol
SSL 15 napon belül lejár
RAM 95% fölött
Swap tartósan használatban
Load Average túl magas
/var/log partíció megtelt
SMART lemezhiba (ha támogatott)
Én ezt már nem különálló scriptekből építeném.

Egyetlen közös server-health.sh készülne, amelyből később akár a Quantum Wizard Studio összes VPS-ének központi felügyelete is kinőhet. Minden szerveren ugyanaz futna, csak a konfiguráció (gépnév, e-mail címek, küszöbök) lenne más.

Szerintem ez egy olyan fejlesztés, ami néhány óra munkával hosszú távon rengeteg kellemetlenségtől kímél meg. Az ilyen problémák – mint a mai lemezbetelés – már jóval azelőtt jelezhetők lennének, hogy szolgáltatásleállást okoznának.

Ez néhány órás munka? Akkor is ha te írod a scriptet?

😊 Igen. Ha én tervezem meg a logikáját és írom a scriptet, akkor nem napokban, hanem inkább 2–4 órában gondolkodnék egy olyan megoldásra, amit már nyugodt szívvel fel lehet tenni éles szerverre.

Én viszont nem egy "80%-nál küldj egy e-mailt" scriptet írnék, hanem egy olyat, amit utána minden VPS-eden használni tudsz.

Én ezt így építeném fel
/usr/local/bin/server-health.sh
/etc/server-health.conf
/var/log/server-health.log
/var/lib/server-health/state.json

Így minden beállítás külön fájlban lenne:

HOSTNAME="<example-host>"
EMAIL="<alert-recipient>"
DISK_WARN=80
DISK_CRIT=90
RAM_WARN=90
BACKUP_WARN=36h
SSL_WARN=15

A script pedig moduláris lenne:

✔ lemezterület
✔ RAM
✔ swap
✔ CPU load
✔ inode használat
✔ backup életkora
✔ Hestia
✔ nginx
✔ Apache
✔ PHP-FPM
✔ MariaDB
✔ Exim
✔ Dovecot
✔ SSL lejárat
✔ opcionálisan SMART állapot
A legfontosabb funkció

Nem spamel.

Például:

Disk: 79%

➡️ nem küld semmit.

80%

➡️ küld egyet.

81%
82%
83%

➡️ nem küld újra.

90%

➡️ küld még egyet.

95%

➡️ küld egy piros riasztást.

79%

➡️ küld egy "helyreállt" levelet.

Így évente jó, ha kapsz 5-10 e-mailt, nem pedig ezret.

És itt jön az, ami szerintem igazán érdekes

Ezt nem <example owner> Server Monitor néven írnám meg.

Hanem:

Quantum Wizard Server Guardian

😊

Egy saját kis rendszert.

Amit később felteszel:

<example-domain> VPS
quantumwizard.hu VPS
QUWIP VPS
bárhová

és mind ugyanazt a rendszert használja.

Sőt...

Már ismerve a projektjeidet, én még egy lépéssel tovább mennék.

A QUWIP Admin felületébe később lehetne egy oldal:

🟢 <example owner> VPS

CPU        7%
RAM       31%
DISK      38%

Backup    OK
SSL       OK
Mail      OK

Last check:
2026-07-18 17:46

Ha egyszer lesz 10-20 VPS-ed, nem kell SSH-zni egyikre sem.

Röviden: igen.
2-4 óra alatt elkészülne egy olyan első verzió, amit már bátran használhatsz.
6-8 óra alatt pedig egy olyan "Quantum Wizard Server Guardian", amit szerintem akár külön nyílt forráskódú projektként is meg lehetne jelentetni GitHubon. Olyan lenne, amit más rendszergazdák is szívesen használnának.

És őszintén? A mai nap után ez szerintem az egyik leghasznosabb kis projekt lenne. Ha ez már tegnap futott volna, valószínűleg már 80%-nál kaptál volna egy e-mailt, és a szerver soha nem jut el a teljes lemezbetelésig.

Itt volt egy sl-es példa mondat... de azt kihagytam...

Pontosan értem, mire gondolsz — és az **SL-példa egyáltalán nem elrugaszkodott**. Azt mutatja, hogy nem pusztán egy monitorozó scriptet szeretnél, hanem egy **moduláris Linux-szerver telepítő és felügyeleti rendszert**, amely felismeri a környezetet, felajánlja a lehetőségeket, telepíti a szükséges függőségeket, majd létrehozza és beállítja a teljes rendszert.

## Alapelv

A **Quantum Wizard Server Guardian** teljesen önálló termék lesz:

* nem kötődik a QUWIP-hez;
* nem kötődik a Quantum Wizard weboldalhoz;
* nem kötődik Hestiához vagy bármely konkrét VPS-hez;
* saját telepítője van;
* saját konfigurációja van;
* saját webes kezelőfelülete van;
* ugyanaz a csomag telepíthető bármely támogatott Linux-szerverre.

## Milyen legyen a telepítés?

A legjobb megoldás szerintem egy **hibrid telepítő**.

### 1. Linux oldali bootstrap telepítő

A rendszer első telepítését root vagy sudo jogosultsággal Linux alatt kell elindítani:

```bash
sudo ./qwsg-install
```

vagy később:

```bash
curl -fsSL https://guardian.quantumwizard.hu/install.sh -o qwsg-install.sh
sudo bash qwsg-install.sh
```

A webes felület önmagában nem tud biztonságosan:

* Linux csomagokat telepíteni;
* systemd service-eket létrehozni;
* új rendszerfelhasználót készíteni;
* fájljogosultságokat beállítani;
* Nginxet vagy Apache-ot konfigurálni;
* tűzfalszabályokat módosítani;
* cron vagy systemd timereket létrehozni.

Ezért az első lépésnek Linux oldalinak kell lennie.

### 2. Grafikus terminálos telepítő

A Linux telepítő nem egyszerű kérdés-felelet script lenne, hanem grafikus hatású terminálfelület:

```text
┌──────────────────────────────────────────────┐
│ Quantum Wizard Server Guardian               │
│ Telepítővarázsló                             │
├──────────────────────────────────────────────┤
│ [✓] Lemezhasználat figyelése                 │
│ [✓] Memória és swap figyelése                │
│ [✓] Szolgáltatások ellenőrzése               │
│ [✓] SSL tanúsítványok figyelése              │
│ [ ] SMART lemezdiagnosztika                  │
│ [✓] Webes kezelőfelület                      │
│ [✓] E-mail-riasztások                        │
│ [ ] Telegram-riasztások                      │
│                                              │
│       < Vissza >     < Telepítés >           │
└──────────────────────────────────────────────┘
```

Ehhez használható például:

* `whiptail`;
* `dialog`;
* később fejlettebb TUI-keretrendszer.

A telepítő ellenőrzi, hogy melyik érhető el, és szükség esetén telepíti.

### 3. Webes beállítóvarázsló

Miután a Linux oldali telepítő felrakta az alapmotort és a webes felületet, a további konfiguráció már történhet böngészőből:

```text
https://szerver-domain.hu:port
```

vagy reverse proxyn keresztül:

```text
https://guardian.szerver-domain.hu
```

A böngészős varázslóban lehet majd:

* riasztási címeket megadni;
* küszöbértékeket állítani;
* szolgáltatásokat kiválasztani;
* napi riportot beállítani;
* modulokat engedélyezni;
* tesztriasztást küldeni;
* ellenőrizni a rendszer állapotát.

Így a telepítés biztonságos marad, de a használata kényelmes és látványos lesz.

---

# Quantum Wizard Server Guardian – teljes funkciólista

## 1. Rendszerfelismerés

A telepítő automatikusan felismeri:

* Linux disztribúció;
* disztribúció verziója;
* architektúra;
* kernelverzió;
* virtuális vagy fizikai szerver;
* csomagkezelő;
* init rendszer;
* webszerver;
* adatbázis;
* PHP-verziók;
* levelezőrendszer;
* szervervezérlő panel;
* tárhelyek és partíciók;
* hálózati interfészek;
* telepített monitorozó eszközök.

Elsőként célszerű támogatni:

* Ubuntu;
* Debian.

Később:

* AlmaLinux;
* Rocky Linux;
* Fedora Server;
* CentOS Stream.

## 2. Függőségkezelés

A telepítő minden modulnál megmutatja:

```text
SMART ellenőrzés
Állapot: nincs telepítve
Szükséges csomag: smartmontools

[ Telepítés ] [ Kihagyás ]
```

Lehetséges függőségek:

* `curl`;
* `jq`;
* `openssl`;
* `mailutils`;
* `smartmontools`;
* `lm-sensors`;
* `sysstat`;
* `iproute2`;
* `procps`;
* `dialog` vagy `whiptail`;
* `logrotate`;
* `sqlite3`.

Semmit nem telepít észrevétlenül. A telepítő előbb megmutatja, mit fog módosítani.

## 3. Lemezfigyelés

* partíciók foglaltsága;
* külön küszöbérték partíciónként;
* inode használat;
* `/var`;
* `/var/log`;
* `/home`;
* `/tmp`;
* backup tárhely;
* Docker tárhely;
* legnagyobb könyvtárak;
* hirtelen tárhelynövekedés észlelése.

Alapértelmezett szintek:

| Állapot        | Foglaltság |
| -------------- | ---------: |
| Normál         |      0–79% |
| Figyelmeztetés |     80–89% |
| Kritikus       |     90–94% |
| Vészhelyzet    | 95% felett |

Minden érték módosítható.

## 4. Memória és swap

* teljes RAM;
* használt RAM;
* elérhető RAM;
* cache;
* swap használata;
* tartós swap használat;
* OOM események;
* legnagyobb memóriafogyasztó folyamatok;
* memóriahasználat változási trendje.

Fontos, hogy a Linux cache-t ne tekintse automatikusan veszélyes memóriafogyasztásnak.

## 5. CPU és terhelés

* aktuális CPU-használat;
* load average: 1, 5 és 15 perc;
* CPU-magszámhoz igazított terhelési küszöb;
* magas I/O wait;
* túl sok futó folyamat;
* kiugró CPU-fogyasztók;
* tartós túlterhelés;
* processzorhőmérséklet, ha elérhető.

## 6. Szolgáltatásfelügyelet

A telepítő felismeri a szolgáltatásokat, és kiválasztható, melyeket figyelje:

* Nginx;
* Apache;
* PHP-FPM;
* MariaDB;
* MySQL;
* PostgreSQL;
* Redis;
* Memcached;
* Exim;
* Postfix;
* Dovecot;
* SSH;
* Docker;
* fail2ban;
* cron;
* Hestia;
* Webmin;
* egyedi systemd service-ek.

Minden szolgáltatásnál megadható:

* csak figyelés;
* automatikus újraindítás;
* újraindítás előtt értesítés;
* újraindítás után értesítés;
* maximális újraindítási próbálkozás;
* hibás újraindítás esetén kritikus riasztás.

Az automatikus javítás külön kapcsolható funkció lesz, nem alapértelmezés.

## 7. Port- és hálózatfigyelés

* megadott port figyelése;
* TCP-kapcsolat tesztelése;
* helyi szolgáltatás elérhetősége;
* külső internetkapcsolat;
* DNS-feloldás;
* gateway elérhetősége;
* váratlanul megnyílt portok;
* eltűnt portok;
* publikus IP-cím változása;
* hálózati forgalom;
* túlzott adatforgalom.

## 8. Weboldal- és HTTP-ellenőrzés

* HTTP/HTTPS elérhetőség;
* válaszkód;
* válaszidő;
* átirányítások;
* megadott szöveg jelenléte;
* hibás tartalom felismerése;
* 500-as hibák;
* túl lassú válasz;
* domainenként külön beállítások.

Például nem elég, hogy a webszerver fut. A Guardian ténylegesen ellenőrzi, hogy az oldal válaszol-e.

## 9. SSL-figyelés

* SSL-tanúsítvány lejárata;
* domainnév-egyezés;
* tanúsítványlánc;
* hibás vagy hiányzó tanúsítvány;
* lejárati figyelmeztetés;
* automatikus domainfelismerés Nginx és Apache konfigurációból;
* Let’s Encrypt megújítás ellenőrzése.

Alapértelmezett figyelmeztetések:

* 30 nap;
* 15 nap;
* 7 nap;
* 3 nap;
* 1 nap.

## 10. Backup-felügyelet

Nem a mentést készíti el elsőként, hanem ellenőrzi a már meglévő mentéseket:

* utolsó backup ideje;
* fájlméret;
* üres backup;
* szokatlanul kicsi backup;
* elavult backup;
* backup könyvtár telítettsége;
* hibás mentési napló;
* Hestia backup;
* egyedi backup mappa;
* adatbázis-dump;
* távoli backup elérhetősége.

Később saját backup-modul is készülhet.

## 11. Adatbázis-ellenőrzés

* szolgáltatás fut-e;
* kapcsolat létrejön-e;
* válaszidő;
* adatbázisméret;
* szabad tárhely;
* túl sok kapcsolat;
* beragadt lekérdezések;
* replikáció állapota;
* hibás táblák;
* utolsó adatbázismentés.

Jelszavakat nem szabad nyílt konfigurációs fájlban tárolni.

## 12. Levelezőrendszer-ellenőrzés

* SMTP;
* IMAP;
* POP3;
* Exim/Postfix;
* Dovecot;
* mail queue mérete;
* beragadt levelek;
* sikertelen kézbesítések;
* túl nagy levelezési sor;
* SMTP-kapcsolat;
* opcionális próbaüzenet;
* SPF/DKIM/DMARC alapellenőrzés.

## 13. Naplófigyelés

* system journal;
* auth log;
* Nginx hibák;
* Apache hibák;
* PHP-FPM hibák;
* MariaDB hibák;
* mail log;
* kernelhibák;
* OOM események;
* lemezhibák;
* túl sok ismétlődő hiba;
* megadott keresőkifejezések.

A Guardian nem küld minden naplósorról levelet. Összesít és állapotot képez.

## 14. SMART és hardverellenőrzés

Ha a környezet támogatja:

* SMART állapot;
* hibás szektorok;
* hőmérséklet;
* újraallokált szektorok;
* NVMe health;
* várható meghibásodás;
* RAID állapot;
* akkumulátor vagy UPS állapot.

VPS-en ez gyakran nem érhető el, ezért opcionális modul.

## 15. Biztonsági ellenőrzések

* sikertelen SSH-belépések;
* fail2ban állapot;
* túl sok sikertelen belépés;
* root SSH engedélyezése;
* jelszavas SSH-bejelentkezés;
* lejárt csomagok;
* függő biztonsági frissítések;
* új rendszerfelhasználó;
* megváltozott sudo-jogosultság;
* váratlan cron bejegyzés;
* megváltozott kritikus fájl;
* tűzfal állapota.

A Guardian első változata csak jelezzen. Az automatikus biztonsági javítás külön, későbbi modul legyen.

## 16. Állapotváltozás-alapú riasztás

Ez a rendszer egyik legfontosabb része.

A Guardian állapotokat tárol:

```text
OK
WARNING
CRITICAL
EMERGENCY
UNKNOWN
```

Példa:

```text
79% → 80%
Figyelmeztetés küldése

80% → 84%
Nincs új levél

84% → 90%
Kritikus riasztás

90% → 89%
Kritikus állapot megszűnt, de még figyelmeztetés

89% → 74%
Helyreállási értesítés
```

A vészhelyzeti állapot ismételt értesítése külön beállítható, például óránként.

## 17. Értesítési csatornák

Első változat:

* helyi `mail` parancs;
* sendmail;
* SMTP;
* webes felületen megjelenő riasztás.

Később:

* Telegram;
* Discord;
* Slack;
* Microsoft Teams;
* webhook;
* SMS-szolgáltató;
* mobil push értesítés.

Minden modulhoz külön meghatározható, melyik csatornát használja.

## 18. Napi és heti riport

A riport tartalmazhatja:

```text
Quantum Wizard Server Guardian
Szerver állapotjelentés

Általános állapot: RENDBEN

CPU: 8%
RAM: 34%
Swap: 0%
Lemez: 42%
Load: 0.41 / 0.32 / 0.28

Szolgáltatások:
Nginx: OK
MariaDB: OK
PHP-FPM: OK
Dovecot: OK

SSL:
Minden tanúsítvány érvényes.

Backup:
Utolsó mentés: 2026-07-18 03:15
Méret: 4.8 GB

Aktív figyelmeztetés nincs.
```

A napi riport kikapcsolható. Lehet csak hibajelzéses üzemmód is.

## 19. Önálló webes felület

A Guardian saját felületet kap.

### Áttekintő oldal

* rendszer általános állapota;
* CPU;
* RAM;
* swap;
* tárhely;
* uptime;
* load;
* szolgáltatások;
* SSL;
* backup;
* aktív hibák;
* utolsó ellenőrzés.

### További oldalak

* Áttekintés;
* Rendszer;
* Tárhely;
* Szolgáltatások;
* Weboldalak;
* SSL;
* Backup;
* Levelezés;
* Naplók;
* Értesítések;
* Modulok;
* Beállítások;
* Frissítések;
* Diagnosztika.

### Biztonság

* külön adminfelhasználó;
* erős jelszó;
* kétlépcsős azonosítás később;
* brute-force védelem;
* session timeout;
* IP-korlátozás;
* opcionális csak localhost elérés;
* reverse proxy támogatás;
* minden módosítás auditnaplóba kerül.

## 20. Előzmények és grafikonok

* CPU-idősor;
* RAM-idősor;
* tárhelyváltozás;
* hálózati forgalom;
* szolgáltatásleállások;
* riasztási előzmények;
* helyreállások;
* SSL lejárati előzmények;
* backupméret változása.

Nem szükséges rögtön nagy adatbázis. Elsőként elegendő lehet SQLite, később opcionális MariaDB vagy PostgreSQL.

## 21. Diagnosztikai csomag

Egy gombbal vagy paranccsal:

```bash
qwsg diagnose
```

Létrehoz egy csomagot:

* rendszerinformációk;
* Guardian-konfiguráció érzékeny adatok nélkül;
* legutóbbi logok;
* szolgáltatásállapotok;
* telepítési verzió;
* hibakódok.

Ez megkönnyíti a távoli támogatást.

## 22. Moduláris telepítő

A telepítő több profilt ajánlhat fel.

### Minimális

* lemez;
* RAM;
* load;
* alapvető szolgáltatások;
* e-mail-riasztás.

### Webszerver

* Nginx/Apache;
* PHP;
* adatbázis;
* SSL;
* HTTP-ellenőrzés;
* backup.

### Levelezőszerver

* Exim/Postfix;
* Dovecot;
* mail queue;
* SMTP/IMAP tesztek;
* tanúsítványok.

### Teljes

Minden támogatott modul.

### Egyéni

A felhasználó egyenként választ.

## 23. Opcionális rendszerkényelmi funkciók

Ide tartozik az SL-példád is. 😊

Létrehozhatunk egy külön kategóriát:

> **Rendszerkényelmi és szórakoztató kiegészítők**

Például:

* `sl` telepítése;
* `cowsay`;
* `fortune`;
* színes MOTD;
* szerveradatok bejelentkezéskor;
* Guardian állapot megjelenítése SSH-belépéskor;
* saját Quantum Wizard üdvözlőképernyő;
* aliasok telepítése;
* biztonságos shell-kényelmi beállítások.

Példa:

```text
Szeretnéd, hogy elgépelt „ls” parancs esetén
alkalmanként megjelenjen a gőzmozdony?

[ ] Igen, feltétlenül
[✓] Nem, ez egy komoly szerver
```

Ez viccesnek tűnik, de az architektúra szempontjából fontos: bizonyítja, hogy a telepítő **nem mereven beégetett**, hanem modulokat és telepítési recepteket kezel.

---

# Javasolt fájlstruktúra

```text
/opt/quantumwizard/server-guardian/
├── bin/
│   ├── qwsg
│   ├── qwsg-agent
│   ├── qwsg-check
│   └── qwsg-update
├── core/
│   ├── scheduler
│   ├── state-manager
│   ├── alert-manager
│   ├── dependency-manager
│   └── module-loader
├── modules/
│   ├── disk/
│   ├── memory/
│   ├── cpu/
│   ├── services/
│   ├── ssl/
│   ├── backup/
│   ├── database/
│   ├── mail/
│   ├── logs/
│   ├── network/
│   └── smart/
├── web/
│   ├── public/
│   ├── app/
│   └── storage/
├── templates/
│   ├── email/
│   └── reports/
├── migrations/
├── installer/
└── VERSION
```

Konfiguráció:

```text
/etc/qwsg/
├── guardian.conf
├── modules.d/
├── notifications.d/
├── websites.d/
└── secrets/
```

Adatok:

```text
/var/lib/qwsg/
├── guardian.db
├── state/
├── history/
└── cache/
```

Naplók:

```text
/var/log/qwsg/
├── guardian.log
├── alerts.log
├── installer.log
└── audit.log
```

Systemd:

```text
/etc/systemd/system/qwsg-agent.service
/etc/systemd/system/qwsg-check.service
/etc/systemd/system/qwsg-check.timer
```

## A fejlesztési irány végleges javaslata

A rendszert három, egymástól elválasztott komponensből építsük:

### Quantum Wizard Server Guardian Agent

Ez végzi a tényleges ellenőrzéseket. Webes felület nélkül is működik.

### Quantum Wizard Server Guardian Installer

Felismeri a rendszert, kezeli a függőségeket, telepíti a modulokat és elkészíti a konfigurációt.

### Quantum Wizard Server Guardian Console

Ez az önálló webes felület. Az Agent adatait jeleníti meg, és innen módosíthatók a beállítások.

Ez azért jó, mert egy minimális szerveren akár csak az Agent is telepíthető. Aki webes kezelőfelületet szeretne, felteszi mellé a Console-t.

## Az első tényleges fejlesztési változat tartalma

Az első használható kiadásba ezt tenném:

* Ubuntu és Debian támogatás;
* interaktív TUI-telepítő;
* rendszerfelismerés;
* függőségkezelés;
* lemez- és inode-figyelés;
* RAM és swap;
* load average;
* systemd szolgáltatások;
* HTTP/HTTPS ellenőrzés;
* SSL-lejárat;
* backup életkor és méret;
* e-mail-riasztás;
* állapotváltozás-kezelés;
* helyreállási értesítés;
* napi riport;
* SQLite állapottárolás;
* önálló, biztonságos webes felület;
* telepítés, frissítés és eltávolítás;
* diagnosztikai parancs;
* teljes telepítési és működési napló.

Ez már nem próbascript lenne, hanem a **Quantum Wizard Server Guardian 1.0 alapja**. A következő lépésben ezt a funkciólistát rendszerezzük **kötelező mag**, **választható modulok** és **későbbi fejlesztések** szerint, majd elkészítjük az első hivatalos rendszerarchitektúrát és fejlesztési ütemtervet.
