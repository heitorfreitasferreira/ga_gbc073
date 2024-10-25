/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"job-shop-ga/hybrid"
	"math/rand"

	"github.com/spf13/cobra"
)

// hybridCmd represents the hybrid command
var hybridCmd = &cobra.Command{
	Use:   "hybrid",
	Short: "PSO/GA",
	Long:  `Inicializa a população do GA com o PSO`,
	Run: func(cmd *cobra.Command, args []string) {
		fileName, _ := cmd.Flags().GetString("instance")
		// mutationRate, _ := cmd.Flags().GetFloat64("mut")
		// crossoverRate, _ := cmd.Flags().GetFloat64("cross")
		// populationSize, _ := cmd.Flags().GetInt("pop")
		// maxGenerations, _ := cmd.Flags().GetInt("gen")
		// seed, _ := cmd.Flags().GetInt("seed")

		// Ler instância do problema
		instance, err := hybrid.GetInstanceFromFile(fileName)
		if err != nil {
			fmt.Println("Erro ao ler o arquivo:", err)
			return
		}
		source := rand.New(rand.NewSource(42))
		ind := hybrid.NewRandomIndividual(*instance, source)
		ind.ExpandToMatrix(*instance)
		fmt.Println(ind)
	},
}

func init() {
	rootCmd.AddCommand(hybridCmd)
	hybridCmd.Flags().String("instance", "./benchmark/instances/ft06", "Nome da instância do problema")

	hybridCmd.Flags().Int("n_particles", 10, "Number of particles")
	hybridCmd.Flags().Int("max_iter", 100, "Number of iterations")
	hybridCmd.Flags().Float64("w", 0.5, "Inertia")
	hybridCmd.Flags().Float64("c_cogni", 1.0, "Cognitive component")
	hybridCmd.Flags().Float64("c_social", 2.0, "Social component")
}
