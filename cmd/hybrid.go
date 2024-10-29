/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"job-shop-ga/hybrid"
	"math/rand"
	"strings"

	"github.com/spf13/cobra"
)

// hybridCmd represents the hybrid command
var hybridCmd = &cobra.Command{
	Use:   "hybrid",
	Short: "PSO/GA",
	Long:  `Inicializa a população do GA com o PSO`,
	Run: func(cmd *cobra.Command, args []string) {
		// Flags
		filePath, _ := cmd.Flags().GetString("instance")
		crossoverRate, _ := cmd.Flags().GetFloat64("cross")
		mutationRate, _ := cmd.Flags().GetFloat64("mut")
		populationSize, _ := cmd.Flags().GetInt("pop")
		maxGenerations, _ := cmd.Flags().GetInt("ga_gen")
		c1, _ := cmd.Flags().GetFloat64("c1")
		c2, _ := cmd.Flags().GetFloat64("c2")
		maxPsoIters, _ := cmd.Flags().GetInt("pso_gen")
		alpha, _ := cmd.Flags().GetFloat64("alpha")
		omegaMax, _ := cmd.Flags().GetFloat64("omega_max")
		omegaMin, _ := cmd.Flags().GetFloat64("omega_min")

		// Ler instância do problema
		instanceName := strings.Split(filePath, "/")[len(strings.Split(filePath, "/"))-1]
		instance, err := hybrid.GetInstanceFromFile(filePath)
		if err != nil {
			fmt.Println("Erro ao ler o arquivo:", err)
			return
		}
		// Randomly generate a seed and print it.
		// TODO: set seed as 42 when program is working properly.
		seed := rand.Int()
		fmt.Println("Seed:", seed)
		source := rand.New(rand.NewSource(int64(4502730368040452047)))

		// Parâmetros do GA
		params := hybrid.Parameters{
			GaParams: hybrid.GaParams{
				CrossoverRate: crossoverRate,
				MutationRate:  mutationRate,
				GA_MAX_ITER:   maxGenerations,
			},
			POPULATION_SIZE: populationSize,
			PsoParams: hybrid.PsoParams{
				C1:           c1,
				C2:           c2,
				OMEGA_MIN:    omegaMax,
				OMEGA_MAX:    omegaMin,
				Alpha:        alpha,
				PSO_MAX_ITER: maxPsoIters,
			},
		}

		rGa, rPso := hybrid.Run(instance, source, params)

		rPso.SaveCsv(instanceName + "_pso")
		rGa.SaveCsv(instanceName + "_ga")
	},
}

func init() {
	rootCmd.AddCommand(hybridCmd)
	hybridCmd.Flags().String("instance", "./benchmark/instances/ft06", "Nome da instância do problema")

	// GA flags
	hybridCmd.Flags().Float64("cross", 0.65, "Taxa de crossover")
	hybridCmd.Flags().Float64("mut", 0.95, "Taxa de mutação")
	hybridCmd.Flags().Int("pop", 100, "Tamanho da população")
	hybridCmd.Flags().Int("ga_gen", 20, "Número de gerações")

	// PSO flags
	hybridCmd.Flags().Float64("w", 0.5, "Inertia")
	hybridCmd.Flags().Float64("c1", 1.0, "Cognitive component")
	hybridCmd.Flags().Float64("c2", 2.0, "Social component")
	hybridCmd.Flags().Int("pso_gen", 200, "Number of pso iterations")
	hybridCmd.Flags().Float64("alpha", 1.0, "alpha")
	hybridCmd.Flags().Float64("omega_min", 0.4, "min inertia")
	hybridCmd.Flags().Float64("omega_max", 1.2, "max inertia")
}
