# QWSG Operátori Konzol

Futtasd a `qwsg observe` parancsot. Az első futás privát alapállapotot készít és ismeretlen marad; egy későbbi ismétlés végrehajtja a teljes operátori értékelést. Ezután egy külön processzben indított `qwsg` megjeleníti a minősített eredményt. A `qwsg check` továbbra is a korlátozott Inventory/Snapshot művelet. A `QWSG_STATE_DIR` privát abszolút állapotgyökeret választ, a `QWSG_STORE` felülírhatja az Inventory Store helyét.

Ha egy értékelés több figyelmet igénylő tényt ad, mint amennyit a korlátozott
Console felsorolhat, a QWSG a legsúlyosabb és legfontosabb tényeket tartja meg,
és kiírja a korrelált vagy elhagyott további tények számát. Az aktuális, de
részleges eredmény a bizonyíték vizsgálatát javasolja, nem állítja, hogy az
ismétlés megjavítja. A megfigyelési hiba jelzi, hogy az értékelés, a projekció
vagy az állapotközlés hibázott, privát gépadatok kiírása nélkül.

Futtasd a `qwsg` parancsot. Terminálban megnyitja a csak olvasható Konzolt; pipe-ban vagy scriptben egy rövid nézetet ír ki, majd kilép.

- `j` / `k`: mozgás
- Enter: megnyitás
- `b`: vissza
- `r`: a Current State újratöltése és frissességi újraminősítése megfigyelés indítása nélkül
- `h`: súgó
- `q`: kilépés

A kezdőnézet megmutatja a szerver állapotát, a szükséges figyelmet, a változásokat, a riasztásokat, a Guardian állapotát, a bizonyíték minőségét és a javasolt lépést. Az azonosítók csak a Részletekben jelennek meg. A haladó parancsok és a JSON-kimenet változatlanul elérhetők.

Az állapot csak a megvalósított kanonikus összehasonlítási és mérnöki bizonyítékot minősíti; nem bizonyít minden lehetséges üzemeltetési problémát. Egy sikeres one-shot megfigyelés nem folyamatos service-bizonyíték, ezért a Guardian „nincs megfigyelve” marad, amíg valódi Runtime Service nem publikál aktuális lifecycle evidence-t.

Az `r` csak olvas: nem foglalja le a Guardian műveleti zárát, nem gyűjt
Inventory-t, nem futtat Pipeline-t és nem publikál állapotot. Kézi aktív
megfigyeléshez külön a `qwsg observe` parancsot használd. A Runtime hibái
korlátozott, gyakorlati okként jelennek meg (például Alert-értékelési,
Notification-tervezési vagy kézbesítési hiba, időtúllépés, megszakítás), nyers
privát hibaszöveg nélkül.
