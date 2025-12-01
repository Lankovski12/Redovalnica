package redovalnica

import "fmt"

type Student struct {
	Ime     string
	Priimek string
	Ocene   []int
}

var (
	stOcen   = 6
	minOcena = 1
	maxOcena = 10
)

func DodajOceno(studenti map[string]Student, vpisnaStevilka string, ocena int) {
	student, ok := studenti[vpisnaStevilka]

	if ok {
		if ocena >= minOcena && ocena <= maxOcena {
			student.Ocene = append(student.Ocene, ocena)
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
		if len(student.Ocene) < stOcen {
			return 0.0
		}
		sum := 0.0
		for _, oc := range student.Ocene {
			sum += float64(oc)
		}
		sum /= float64(len(student.Ocene))

		return sum
	}
	return -1.0
}

func IzpisVsehOcen(studenti map[string]Student) {
	fmt.Printf("REDOVALNICA:\n")
	for vpisna, student := range studenti {
		fmt.Printf("%s - %s %s: %v\n", vpisna, student.Ime, student.Priimek, student.Ocene)
	}
}

func IzpisiKoncniUspeh(studenti map[string]Student) {
	for vpisna, student := range studenti {
		povprecna := povprecje(studenti, vpisna)
		if povprecna >= 9.0 {
			fmt.Printf("%s %s: povprečna ocena %.1f -> Odličen študent!\n", student.Ime, student.Priimek, povprecna)
		} else if povprecna >= 6.0 && povprecna < 9.0 {
			fmt.Printf("%s %s: povprečna ocena %.1f -> Povprečen študent!\n", student.Ime, student.Priimek, povprecna)
		} else {
			fmt.Printf("%s %s: povprečna ocena %.1f -> Neuspešen študent!\n", student.Ime, student.Priimek, povprecna)
		}
	}
}
