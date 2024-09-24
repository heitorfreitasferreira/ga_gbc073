/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"job-shop-ga/ga"
	"math/rand"

	"github.com/spf13/cobra"
)

// runExpCmd represents the runOne command
var runExpCmd = &cobra.Command{
	Use:     "runExp",
	Short:   "Rodar experimentos propostos",
	Example: "job-shop-ga runExp",
	Run: func(cmd *cobra.Command, args []string) {
		fileNames := []string{"./benchmark/instances/abz9", "./benchmark/instances/swv12", "./benchmark/instances/ta75"}

		mutationRate := 0.3
		crossoverRate := 0.4
		populationSize := 60
		maxGenerations := 50000
		seed := 42

		mods := []string{"", "mut", "tour", "mutTour"}
		for _, fileName := range fileNames {
			for _, mod := range mods {
				// Ler instância do problema
				instance, err := ga.GetInstanceFromFile(
					fileName,
					mutationRate,
					crossoverRate,
					populationSize,
					maxGenerations,
					rand.New(rand.NewSource(int64(seed))),
				)

				if err != nil {
					fmt.Println("Erro ao ler o arquivo:", err)
					return
				}

				fmt.Println("Running instance", fileName, "with mod", mod)

				switch mod {
				case "":
					instance.Run()
				case "mut":
					instance.RunModMutation()
				case "tour":
					instance.RunModTournament()
				case "mutTour":
					instance.RunModTournamentMutation()
				}

				instance.ToCsv()
			}

			fmt.Println("Running instance", fileName, "with crossover 0.7")
			crossoverRate = 0.7
			instance, err := ga.GetInstanceFromFile(
				fileName,
				mutationRate,
				crossoverRate,
				populationSize,
				maxGenerations,
				rand.New(rand.NewSource(int64(seed))),
			)

			if err != nil {
				fmt.Println("Erro ao ler o arquivo:", err)
				return
			}

			// Muda taxa de crossover
			instance.Name += "_crossover_0.7"
		}

	},
}

func init() {
	rootCmd.AddCommand(runExpCmd)
}
