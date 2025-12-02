package redovalnica

import "fmt"

// definiramo stukturo student
type Student struct {
	Ime     string
	Priimek string
	Ocene   []int
}

// vidna funkcija, ki studentu doda oceno
func DodajOceno(studenti map[string]Student, vpisnaStevilka string, ocena int, minOcena int, maxOcena int) int {
	student, ok := studenti[vpisnaStevilka]

	if !ok {
		fmt.Printf("Študent %s ne obstaja\n", vpisnaStevilka)
		return 1
	}

	if ocena < minOcena || ocena > maxOcena {
		fmt.Printf("Neveljavna ocena %d (dovoljeno: %d-%d)\n", ocena, minOcena, maxOcena)
		return 2
	}

	student.Ocene = append(student.Ocene, ocena)
	studenti[vpisnaStevilka] = student
	return 0
}

// skrita funkcija, ki izracuna povprecje vseh ocen
func povprecje(studenti map[string]Student, vpisnaStevilka string, stOcen int) float64 {
	student, ok := studenti[vpisnaStevilka]

	if !ok {
		return -1.0
	}

	if len(student.Ocene) < stOcen {
		return 0.0
	}

	sum := 0.0
	for _, oc := range student.Ocene {
		sum += float64(oc)
	}
	return sum / float64(len(student.Ocene))
}

// javna funkcija, ki izpise vse ocene za vse studente
func IzpisVsehOcen(studenti map[string]Student) {
	fmt.Printf("REDOVALNICA:\n")

	for vpisna, student := range studenti {
		fmt.Printf("%s - %s %s: %v\n", vpisna, student.Ime, student.Priimek, student.Ocene)
	}
}

// javna funkcija, ki izracuna uspeh vseh studentov in jih izpise
func IzpisiKoncniUspeh(studenti map[string]Student, stOcen int) {
	for vpisna, student := range studenti {
		povprecna := povprecje(studenti, vpisna, stOcen)

		if povprecna >= 9.0 {
			fmt.Printf("%s %s: povprečna ocena %.1f -> Odličen študent\n", student.Ime, student.Priimek, povprecna)
		} else if povprecna >= 6.0 && povprecna < 9.0 {
			fmt.Printf("%s %s: povprečna ocena %.1f -> Povprečen študent\n", student.Ime, student.Priimek, povprecna)
		} else {
			fmt.Printf("%s %s: povprečna ocena %.1f -> Neuspešen študent\n", student.Ime, student.Priimek, povprecna)
		}
	}
}
