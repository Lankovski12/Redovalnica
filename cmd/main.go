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
	// Podatki študentov
	studenti := map[string]redovalnica.Student{
		"63220333": {Ime: "Lana", Priimek: "Tkalcic", Ocene: []int{7, 8, 9, 9, 7, 8, 8}},
		"63220123": {Ime: "Marjanca", Priimek: "Token", Ocene: []int{10, 9, 9, 9, 8, 8, 10}},
		"63220666": {Ime: "Peter", Priimek: "Miki", Ocene: []int{1, 2, 3, 2, 1}},
		"63220010": {Ime: "Pipi", Priimek: "Strel", Ocene: []int{3, 4, 2, 5, 2, 3, 6}},
		"63220999": {Ime: "Jozica", Priimek: "Marks", Ocene: []int{5, 6, 8, 9, 2, 3}},
	}

	cmd := &cli.Command{
		Name:  "redovalnica",
		Usage: "Aplikacija za upravljanje študentskih ocen",

		// TRI STIKALA iz navodil
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

		// UKAZI, ki kličejo funkcije iz paketa redovalnica
		Commands: []*cli.Command{
			{
				Name:  "dodaj",
				Usage: "Dodaj oceno študentu",
				Action: func(ctx context.Context, cmd *cli.Command) error {
					// Preberi argumente
					args := cmd.Args()
					if args.Len() != 2 {
						fmt.Println("Uporaba: dodaj <vpisna_stevilka> <ocena>")
						return nil
					}

					vpisna := args.Get(0)
					ocena, err := strconv.Atoi(args.Get(1))
					if err != nil {
						fmt.Println("Napaka: ocena mora biti število")
						return nil
					}

					// DODAJ OCENO - uporabi funkcijo iz paketa
					// Tukaj morate še povezati stikala s funkcijo!
					// Za zdaj delamo brez stikal
					redovalnica.DodajOceno(studenti, vpisna, ocena)
					return nil
				},
			},
			{
				Name:  "izpis",
				Usage: "Izpiši vse ocene",
				Action: func(ctx context.Context, cmd *cli.Command) error {
					// Kliče funkcijo iz paketa
					redovalnica.IzpisVsehOcen(studenti)
					return nil
				},
			},
			{
				Name:  "uspeh",
				Usage: "Izpiši končni uspeh",
				Action: func(ctx context.Context, cmd *cli.Command) error {
					// Kliče funkcijo iz paketa
					redovalnica.IzpisiKoncniUspeh(studenti)
					return nil
				},
			},
		},

		// Privzeta akcija (če ni ukaza)
		Action: func(ctx context.Context, cmd *cli.Command) error {
			fmt.Println("=== REDOVALNICA ===")
			fmt.Println("Uporaba: redovalnica [ukaz]")
			fmt.Println("")
			fmt.Println("Ukazi:")
			fmt.Println("  dodaj   - Dodaj oceno študentu")
			fmt.Println("  izpis   - Prikaži vse ocene")
			fmt.Println("  uspeh   - Prikaži končni uspeh")
			fmt.Println("")
			fmt.Println("Primer: redovalnica --stOcen=5 izpis")
			return nil
		},
	}

	// Zaženi program
	if err := cmd.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}
}
