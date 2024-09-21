/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"job-shop-ga/ga"

	"github.com/spf13/cobra"
)

// runOneCmd represents the runOne command
var runOneCmd = &cobra.Command{
	Use:     "runOne",
	Short:   "Rodar uma instância do problema",
	Example: "job-shop-ga runOne --mut 0.3 --cross 0.4 --pop 60 --gen 1000 --seed 42 --instance abz6",
	Run: func(cmd *cobra.Command, args []string) {

		fileName, _ := cmd.Flags().GetString("instance")
		mutationRate, _ := cmd.Flags().GetFloat64("mut")
		crossoverRate, _ := cmd.Flags().GetFloat64("cross")
		populationSize, _ := cmd.Flags().GetInt("pop")
		maxGenerations, _ := cmd.Flags().GetInt("gen")
		seed, _ := cmd.Flags().GetInt("seed")

		ga.InitSource(int64(seed))
		// Ler instância do problema
		instance, err := ga.GetInstanceFromFile(fileName, mutationRate, crossoverRate, populationSize, maxGenerations)
		if err != nil {
			fmt.Println("Erro ao ler o arquivo:", err)
			return
		}
		instance.GenerateInitialPopulation()
		cromossome, makespan := instance.Run()

		fmt.Println("Melhor cromossomo:", cromossome)
		fmt.Println("Makespan:", makespan)

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
}
