# Redovalnica

Paket Redovalnica predstavlja preprosto aplikacijo za upravljanje študenskih ocen. Hranimo podatke o študentih (ime, priimek, vpisna številka, ocene) in omogoča dodajanje ocen in izpis redovalnice ter končnega uspega. Paket vsebuje naslednje funkcije:

- DodajOceno(vpisnaStevilka, ocena): Študentu lahko vpišemo novo oceno.
- IzpisVsehOcen: Namenjen prikazu vseh študenov in njihovih ocen.
- IzpisiKoncniUspeh: Izračunamo povprečje ocen ter izpišemo povprečje vseh študentov.

Paket interno vsebuje tudi funkcijo povprečje, ki pa ni javna. <br>

Paket prav tako ponuja tri stikala, ki so namenjena spreminjanju pravil o ocenjevanju:
- stOcen: Spreminjano koliko ocen je potrebnih, da študent uspešno opravi predmet. Prevzeta vrednost je 6.
- minOcena: Spremenimo lahko katera ocena je najmanjša, ki jo lahko študent pridobi. Prevzeta vrednost je 1.
- maxOcena: Spremenimo lahko katera ocena je najvišja, ki jo lahko študent pridobi. Prevzeta vrednost je 10.

Paket je potrebno naj prej namestiti z ukazom: <br>
```bash
git clone https://github.com/Lankovski12/Redovalnica.git
```
In namestiti odvisnosti:
```bash
go mod download
```
Ko smo paket namestili, paket naj prej zgradimo z ukazom:
```bash
 go build ./cmd/main.go
```
Ko aplikacijo zaženemo, se nam izpiše nekaj pomebnih informaciji.
```bash
 go run ./cmd/main.go
```
```bash
=== REDOVALNICA ===
Uporaba: go run ./main.go [stikalo] [ukaz]

Ukazi:
  dodaj [vpisnaStevilka] [ocena]   - Dodaj oceno študentu
  izpis   - Prikaži vse ocene
  uspeh   - Prikaži končni uspeh

Primer: go run ./cmd/main.go --stOcen=5 izpis
```
Za pomoč, kako uporabljati paket, lahko uporabimo ukaz:
```bash
 go run ./cmd/main.go -h
//ali
go run ./cmd/main.go --help
```

Spodaj so prikazani primeri kako uporabljamo aplikacijo in kako spreminjamo stikala.<br>
Za uporabo funkcije IzpisVsehOcen uporabimo besedo <strong>izpis</strong>.
```bash
go run ./cmd/main.go izpis
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
Za uporabo funkcije IzpisiKonciUspeh upoaabimo besedo <strong>uspeh</strong>.
```bash
go run ./cmd/main.go uspeh
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
go run ./cmd/main.go dodaj <vpisnaStevilka> <ocena>
```
Sama funkcija ne izpiše ničesar, vendar izpiše napako oziroma upešen vpis ocene:
```bash
//napaka, premalo podatkov
Uporaba: go run ./cmd/main.go dodaj <vpisnaStevilka> <ocena>
//napaka, parameter ocena ni bila število
Napaka, vpisi stevilo.
//ocena je bilo uspesno vpisana
Ocena je dodana. Preverite z 'izpis'.
```
Stikala pa spreminjamo tako:
```bash
go run ./cmd/main.go --stOcen=3 uspeh
```

