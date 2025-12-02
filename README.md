# Redovalnica

Paket Redovalnica predstavlja aplikacijo za upravljanje študenskih ocen. Hrani podatke o študentih (ime, priimek, vpisna številka, ocene) in omogoča dodajanje ocen in izpis redovalnice ter končnega uspeha. Paket vsebuje naslednje funkcije:

- DodajOceno(vpisnaStevilka, ocena): Študentu lahko vpišemo novo oceno.
- IzpisVsehOcen: Namenjen prikazu vseh študentov in njihovih ocen.
- IzpisiKoncniUspeh: Izračunamo povprečje ocen ter izpišemo povprečje vseh študentov.

Paket interno vsebuje tudi funkcijo povprečje, ki pa ni javna. <br>

Paket prav tako ponuja tri stikala, ki so namenjena spreminjanju pravil o ocenjevanju:
- stOcen: Spreminjano koliko ocen je potrebnih, da študent uspešno opravi predmet. Privzeta vrednost je 6.
- minOcena: Spremenimo lahko katera ocena je najmanjša, ki jo lahko študent pridobi. Privzeta vrednost je 1.
- maxOcena: Spremenimo lahko katera ocena je najvišja, ki jo lahko študent pridobi. Privzeta vrednost je 10.

Paket je potrebno naj prej namestiti z ukazom: <br>
```bash
git clone https://github.com/Lankovski12/Redovalnica.git
```
In namestiti odvisnosti:
```bash
go mod download
```
Ko smo paket namestili, aplikacijo naj prej zgradimo z ukazom:
```bash
 go build ./cmd/main.go
```
Ko aplikacijo zaženemo, se nam izpiše nekaj pomembnih informaciji. Na tej točki je tudi potrebno dodati stikala, katera si želimo, da bi veljala v seji. Ko program teče ne moremo več dodajati ali spreminjati stikal.
```bash
 go run ./cmd/main.go [stikala]
```
```bash
=== REDOVALNICA ===
Uporaba: ukaz [argumenti]
Ukazi:
  dodaj <vpisna> <ocena>  - Dodaj oceno študentu
  izpis                   - Prikaži vse ocene
  uspeh                   - Prikaži končni uspeh
  izhod                   - Zapri program
```
Spodaj so prikazani primeri, kako uporabljamo aplikacijo in kako spreminjamo stikala.<br>

Za uporabo funkcije IzpisVsehOcen uporabimo besedo <strong>izpis</strong>.
```bash
izpis
```
Izpis funkcije zgleda tako:
```bash
REDOVALNICA:
63220333 - Lana Tkalcic: [7 8 9 9 7 8 8]
63220123 - Marjanca Token: [10 9 9 9 8 8 10]
63220666 - Peter Miki: [6 6 6 2 1]
63220010 - Pipi Strel: [3 4 2 5 2 3 6]
63220999 - Jozica Marks: [5 6 8 9 2 3]
```
Za uporabo funkcije IzpisiKonciUspeh uporabimo besedo <strong>uspeh</strong>.
```bash
uspeh
```
Izpis funkcije zgleda tako:
```bash
Marjanca Token: povprečna ocena 9.0 -> Odličen študent
Peter Miki: povprečna ocena 0.0 -> Neuspešen študent
Pipi Strel: povprečna ocena 3.6 -> Neuspešen študent
Jozica Marks: povprečna ocena 5.5 -> Neuspešen študent
Lana Tkalcic: povprečna ocena 8.0 -> Povprečen študent
```
Za uporabo funkcije DodajOceno uporabimo besedo <strong>dodaj</strong>.
```bash
dodaj <vpisnaStevilka> <ocena>
```
Ob neuspešnem dodajanju, dobimo informacijo o napaki. Ob uspehu pa dobimo potrditev, da je bilo dodajanje uspešno.
```bash
//napaka, premalo podatkov
Premalo argumentov.
//napaka, parameter ocena ni bila število
Napaka, ocena ni število.
//ocena je bilo uspesno vpisana
Ocena je dodana. Preverite z 'izpis'.
```
Stikala moramo dodati k ukazu za pogon aplikacije.
```bash
go run ./cmd/main.go --stOcen=3 --minOcena=3
```
Za dodatne informacije o stikalih lahko napišemo naslednji ukaz.
```bash
go run ./cmd/main.go --help
//ali
go run ./cmd/main.go -h
```
```bash
NAME:
   Redovalnica - Aplikacija za upravljanje študentskih ocen

USAGE:
   Redovalnica [global options]

GLOBAL OPTIONS:
   --stOcen int    Minimalno število ocen za pozitivno oceno (default: 6)
   --minOcena int  Najmanjša možna ocena (default: 1)
   --maxOcena int  Največja možna ocena (default: 10)
   --help, -h      show help
```
Ko želimo zaključi aplikacijo preprosto napisemo <strong>izhod</strong>.
```bash
izhod
```
