# Szakadt csöves '88 – natív port

A **Szakadt csöves '88** (Pató Pál Software, SZY & KAMARAS, 1989) DOS-os
játék natív újraírása Go-ban. Egyetlen futtatható fájl, Linuxon (PC-n és
ARM-on, pl. Raspberry Pi-n), Windowson és macOS-en fut, DOSBox és emulátor
nélkül.

A port az eredeti `CSOVES.EXE` visszafejtésével készült, és ugyanazt a
képet adja, mint az eredeti program.

A port AI segítségével (Claude) készült.

## Futtatás

```sh
go build -o csoves ./cmd/csoves
./csoves
```

Kapcsolók: `-scale N` (ablaknagyítás, alapból 2), `-norandomize` (a
véletlenszám-generátor mindig ugyanonnan indul).

## Irányítás

Mint az eredetiben: **-** feljebb, **+** lejjebb viszi a menücsíkot,
**Enter** választ, **Space** mondja ki az összerakott káromkodást. A címkép
után **I**/**N** a súgóhoz.

A port kiegészítése: a menücsík a **↑** (fel) és **↓** (le) nyíllal is
mozgatható, hogy numerikus billentyűzet nélküli (pl. TKL) billentyűzeten is
kényelmes legyen. A **-** és **+** ugyanúgy működik, mint az eredetiben.

**F11** vagy **Alt+Enter**: teljes képernyő.

## Fordítás más rendszerekre

Linuxról keresztfordítva:

```sh
# Windows (64 bit)
GOOS=windows GOARCH=amd64 CGO_ENABLED=0 go build -trimpath -ldflags "-s -w -H windowsgui" -o csoves.exe ./cmd/csoves

# macOS, Apple Silicon / Intel
GOOS=darwin GOARCH=arm64 CGO_ENABLED=0 go build -trimpath -ldflags "-s -w" -o csoves-mac-arm64 ./cmd/csoves
GOOS=darwin GOARCH=amd64 CGO_ENABLED=0 go build -trimpath -ldflags "-s -w" -o csoves-mac-amd64 ./cmd/csoves

# Linux ARM, pl. Raspberry Pi (64 bites, ill. 32 bites OS; GOARM=6 minden Pi-n fut)
GOOS=linux GOARCH=arm64 CGO_ENABLED=0 go build -trimpath -ldflags "-s -w" -o csoves-linux-arm64 ./cmd/csoves
GOOS=linux GOARCH=arm GOARM=6 CGO_ENABLED=0 go build -trimpath -ldflags "-s -w" -o csoves-linux-armv6 ./cmd/csoves
```

Hogy melyik ARM-változat kell: a célgépen a `uname -m` parancs `aarch64`-et
ír 64 bites rendszeren, `armv6l`/`armv7l`-t 32 bitesen.

A macOS-bináris nincs aláírva; első indítás előtt a Macen:
`xattr -d com.apple.quarantine csoves-mac-arm64 && chmod +x csoves-mac-arm64`.

Linuxon az Ebitengine alapból cgo-t használ, ehhez kellenek az X11/OpenGL
fejlécfájlok (Arch: `libx11 libxrandr libxcursor libxinerama libxi mesa`).
Futtatáshoz asztali környezet (X11 vagy Wayland) és OpenGL-driver kell.

## Felépítés

| Hely | Tartalom |
|---|---|
| `cmd/csoves` | Belépési pont: Ebitengine-ablak, billentyűzet-leképezés, headless mód |
| `internal/vid` | A két videómód, ahogy a DOSBox mutatja: 40 oszlopos szöveges mód (VGA 8×16 font, attribútumok, villogás) és SCREEN 1 (CGA 320×200, 4 szín, 8×8 font) |
| `internal/dos` | BIOS-szerű billentyűpuffer, Delay, Halt |
| `internal/game` | A játék. `main_gen.go` és `vars_gen.go` az eredeti gépi kódból fordított Go kód; `qb.go` (PRINT, COLOR, LOCATE, CLS, INPUT$, INKEY$, RND, stringfüggvények) és `draw.go` (DRAW, PAINT) a QuickBASIC 4.5 futtatókörnyezetének szükséges része |
| `internal/charset` | Az eredeti kódlap (CP437 + CWI-2 ékezetek) ↔ UTF-8, a billentyűzethez |
| `assets` | A VGA-kártya fontjai (8×8 és 8×16) |

A játéklogika szándékosan követi az eredeti BASIC program szerkezetét: a
modulszintű változók a `vars` struktúra mezői (a CLEAR ezt nullázza), a SUB-ok
metódusok, a GOSUB/RETURN egy visszatérési verem, a számítások SINGLE
pontossággal és ugyanabban a sorrendben történnek, és a véletlenszámok is a
QuickBASIC generátorából jönnek – így a játék ugyanazokra a billentyűkre
ugyanúgy reagál, mint az eredeti.

A headless mód tesztekhez használható:

```sh
./csoves -headless -norandomize -keys "Enter n Enter shot:a.png + + Enter" -shot vege.png
```
