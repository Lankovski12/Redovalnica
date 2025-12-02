package main

// imports
import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/Lankovski12/Redovalnica/redovalnica"
	"github.com/urfave/cli/v3"
)

// globalna mapa studentov
var studenti = map[string]redovalnica.Student{
	"63220333": {Ime: "Lana", Priimek: "Tkalcic", Ocene: []int{7, 8, 9, 9, 7, 8, 8}},
	"63220123": {Ime: "Marjanca", Priimek: "Token", Ocene: []int{10, 9, 9, 9, 8, 8, 10}},
	"63220666": {Ime: "Peter", Priimek: "Miki", Ocene: []int{6, 6, 6, 2, 1}},
	"63220010": {Ime: "Pipi", Priimek: "Strel", Ocene: []int{3, 4, 2, 5, 2, 3, 6}},
	"63220999": {Ime: "Jozica", Priimek: "Marks", Ocene: []int{5, 6, 8, 9, 2, 3}},
}

func main() {
	cmd := &cli.Command{
		Name:  "Redovalnica",
		Usage: "Aplikacija za upravljanje študentskih ocen",

		// nastavimo stikala
		Flags: []cli.Flag{
			&cli.IntFlag{
				Name:  "stOcen",
				Value: 6,
				Usage: "Minimalno število ocen za pozitivno oceno",
			},
			&cli.IntFlag{
				Name:  "minOcena",
				Value: 1,
				Usage: "Najmanjša možna ocena",
			},
			&cli.IntFlag{
				Name:  "maxOcena",
				Value: 10,
				Usage: "Največja možna ocena",
			},
		},

		Action: func(ctx context.Context, cmd *cli.Command) error {
			// shranimo nove vrednosti stikal, če jih vpisemo
			minOc := cmd.Int("minOcena")
			maxOc := cmd.Int("maxOcena")
			stOcen := cmd.Int("stOcen")

			//izpis ob zagonu
			fmt.Println("=== REDOVALNICA ===")
			fmt.Println("Uporaba: ukaz [argumenti]")
			fmt.Println("Ukazi:")
			fmt.Println("  dodaj <vpisna> <ocena>  - Dodaj oceno študentu")
			fmt.Println("  izpis                   - Prikaži vse ocene")
			fmt.Println("  uspeh                   - Prikaži končni uspeh")
			fmt.Println("  izhod                   - Zapri program")
			fmt.Println("")

			scanner := bufio.NewScanner(os.Stdin)

			// neskonca zanka za potrebe delovanja v eni seji
			for {
				fmt.Print("> ")
				if !scanner.Scan() {
					break
				}

				line := strings.TrimSpace(scanner.Text())
				if line == "" {
					continue
				}

				parts := strings.Fields(line)
				command := parts[0]

				// gledamo kateri ukaz je bil vpisan in ga izvedemo
				switch command {
				case "dodaj":
					if len(parts) != 3 {
						fmt.Println("Premalo argumentov.")
						continue
					}

					vpisna := parts[1]
					ocenaString := parts[2]

					ocena, err := strconv.Atoi(ocenaString)
					if err != nil {
						fmt.Printf("Napaka, ocena ni število.")
						continue
					}

					st := redovalnica.DodajOceno(studenti, vpisna, ocena, minOc, maxOc)

					if st == 0 {
						fmt.Printf("Ocena je dodana. Preverite z 'izpis'.\n")
					}

				case "izpis":
					redovalnica.IzpisVsehOcen(studenti)

				case "uspeh":
					redovalnica.IzpisiKoncniUspeh(studenti, stOcen)

				case "izhod":
					fmt.Println("Program zaključen")
					return nil

				default:
					fmt.Printf("Neznan ukaz: %s\n", command)
				}
			}

			return nil
		},
	}

	if err := cmd.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}
}
