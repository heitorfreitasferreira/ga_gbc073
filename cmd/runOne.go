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

// runOneCmd represents the runOne command
var runOneCmd = &cobra.Command{
	Use:     "runOne",
	Short:   "Rodar uma instância do problema",
	Example: "job-shop-ga runOne --mut 0.3 --cross 0.4 --pop 60 --gen 1000 --seed 42 --instance abz8",
	Run: func(cmd *cobra.Command, args []string) {

		fileName, _ := cmd.Flags().GetString("instance")
		mutationRate, _ := cmd.Flags().GetFloat64("mut")
		crossoverRate, _ := cmd.Flags().GetFloat64("cross")
		populationSize, _ := cmd.Flags().GetInt("pop")
		maxGenerations, _ := cmd.Flags().GetInt("gen")
		seed, _ := cmd.Flags().GetInt("seed")
		mod, _ := cmd.Flags().GetString("mod")

		// Ler instância do problema
		instance, err := ga.GetInstanceFromFile(fileName, mutationRate, crossoverRate, populationSize, maxGenerations, rand.New(rand.NewSource(int64(seed))))
		if err != nil {
			fmt.Println("Erro ao ler o arquivo:", err)
			return
		}

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

		if csv, _ := cmd.Flags().GetBool("csv"); csv {
			instance.ToCsv()
		}
	},
}

func init() {
	rootCmd.AddCommand(runOneCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// runOneCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	runOneCmd.Flags().String("instance", "./benchmark/instances/abz6", "Nome da instância do problema")
	runOneCmd.Flags().String("mod", "", "Modificação a ser executada")
}
