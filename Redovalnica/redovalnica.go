package Redovalnica

import "fmt"

type Student struct {
	ime     string
	priimek string
	ocene   []int
}

func DodajOceno(studenti map[string]Student, vpisnaStevilka string, ocena int) {
	student, ok := studenti[vpisnaStevilka]

	if ok {
		if ocena > 0 && ocena < 11 {
			student.ocene = append(student.ocene, ocena)
			studenti[vpisnaStevilka] = student
		} else {
			fmt.Printf("Invalid grade %d\n", ocena)
		}
	} else {
		fmt.Printf("Student %s does not exist\n", vpisnaStevilka)
	}
}

func povprecje(studenti map[string]Student, vpisnaStevilka string) float64 {
	student, ok := studenti[vpisnaStevilka]
	if ok {
		if len(student.ocene) < 6 {
			return 0.0
		}
		sum := 0.0
		for _, oc := range student.ocene {
			sum += float64(oc)
		}
		sum /= float64(len(student.ocene))

		return sum
	}
	return -1.0
}

func IzpisVsehOcen(studenti map[string]Student) {
	fmt.Printf("REDOVALNICA:\n")
	for vpisna, student := range studenti {
		fmt.Printf("%s - %s %s: %v\n", vpisna, student.ime, student.priimek, student.ocene)
	}
}

func IzpisiKoncniUspeh(studenti map[string]Student) {
	for vpisna, student := range studenti {
		povprecna := povprecje(studenti, vpisna)
		if povprecna >= 9.0 {
			fmt.Printf("%s %s: povprečna ocena %.1f -> Odličen študent!\n", student.ime, student.priimek, povprecna)
		} else if povprecna >= 6.0 && povprecna < 9.0 {
			fmt.Printf("%s %s: povprečna ocena %.1f -> Povprečen študent!\n", student.ime, student.priimek, povprecna)
		} else {
			fmt.Printf("%s %s: povprečna ocena %.1f -> Neuspešen študent!\n", student.ime, student.priimek, povprecna)
		}
	}
}
