package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/Lankovski12/Redovalnica/redovalnica"
	"github.com/urfave/cli/v3"
)

func main() {
	//definiramo mapo studentov
	studenti := make(map[string]redovalnica.Student)

	//dodamo nekaj studentov
	studenti["63220333"] = redovalnica.Student{Ime: "Lana", Priimek: "Tkalcic", Ocene: []int{7, 8, 9, 9, 7, 8, 8}}
	studenti["63220123"] = redovalnica.Student{Ime: "Marjanca", Priimek: "Token", Ocene: []int{10, 9, 9, 9, 8, 8, 10}}
	studenti["63220666"] = redovalnica.Student{Ime: "Peter", Priimek: "Miki", Ocene: []int{6, 6, 6, 2, 1}}
	studenti["63220010"] = redovalnica.Student{Ime: "Pipi", Priimek: "Strel", Ocene: []int{3, 4, 2, 5, 2, 3, 6}}
	studenti["63220999"] = redovalnica.Student{Ime: "Jozica", Priimek: "Marks", Ocene: []int{5, 6, 8, 9, 2, 3}}

	cmd := &cli.Command{
		Name:  "redovalnica",
		Usage: "Aplikacija za upravljanje študentskih ocen",

		//definiramo stikala
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

		Commands: []*cli.Command{
			{
				Name:  "dodaj",
				Usage: "Dodaj oceno študentu",
				Action: func(ctx context.Context, cmd *cli.Command) error {
					args := cmd.Args()
					if args.Len() != 2 {
						fmt.Println("Uporaba: redovalnica dodaj <vpisnaStevilka> <ocena>")
						return nil
					}

					vpisna := args.Get(0)
					ocenaStr := args.Get(1)

					ocena, err := strconv.Atoi(ocenaStr)
					if err != nil {
						fmt.Printf("Napaka, vpisi stevilo.")
					}

					redovalnica.DodajOceno(studenti, vpisna, ocena)

					fmt.Printf("Ukaz izveden. Preverite z 'izpis'.\n")

					return nil
				},
			},
			{
				Name:  "izpis",
				Usage: "Izpiši vse ocene",
				Action: func(ctx context.Context, cmd *cli.Command) error {
					redovalnica.IzpisVsehOcen(studenti)
					return nil
				},
			},
			{
				Name:  "uspeh",
				Usage: "Izpiši končni uspeh",
				Action: func(ctx context.Context, cmd *cli.Command) error {
					redovalnica.IzpisiKoncniUspeh(studenti)
					return nil
				},
			},
		},

		Action: func(ctx context.Context, cmd *cli.Command) error {
			fmt.Println("=== REDOVALNICA ===")
			fmt.Println("Uporaba: go run ./main.go [stikalo] [ukaz]")
			fmt.Println("")
			fmt.Println("Ukazi:")
			fmt.Println("  dodaj [vpisnaStevilka] [ocena]   - Dodaj oceno študentu")
			fmt.Println("  izpis   - Prikaži vse ocene")
			fmt.Println("  uspeh   - Prikaži končni uspeh")
			fmt.Println("")
			fmt.Println("Primer: go run ./cmd/main.go --stOcen=5 izpis")
			return nil
		},
	}

	if err := cmd.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}
}
