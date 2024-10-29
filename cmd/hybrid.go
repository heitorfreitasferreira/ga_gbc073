/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"job-shop-ga/hybrid"
	"math/rand"
	"os"

	"github.com/spf13/cobra"
)

// hybridCmd represents the hybrid command
var hybridCmd = &cobra.Command{
	Use:   "hybrid",
	Short: "PSO/GA",
	Long:  `Inicializa a população do GA com o PSO`,
	Run: func(cmd *cobra.Command, args []string) {
		instances := []string{
			"ft06", "ft10", "ft20", "la01", "la06", "la11", "la16", "la21",
		}
		modifications := []string{
			"original", "lower_mutation", "c1_and_c2_sum_4", "pop_size_iter",
		}
		parameters := []hybrid.Parameters{
			{
				GaParams: hybrid.GaParams{
					CrossoverRate: 0.65,
					MutationRate:  0.95,
					GA_MAX_ITER:   200,
				},
				POPULATION_SIZE: 100,
				PsoParams: hybrid.PsoParams{
					C1:           1.0,
					C2:           2.0,
					OMEGA_MIN:    0.4,
					OMEGA_MAX:    1.2,
					Alpha:        1.0,
					PSO_MAX_ITER: 20,
				},
			},
			{
				GaParams: hybrid.GaParams{
					CrossoverRate: 0.65,
					MutationRate:  0.50,
					GA_MAX_ITER:   200,
				},
				POPULATION_SIZE: 100,
				PsoParams: hybrid.PsoParams{
					C1:           1.0,
					C2:           2.0,
					OMEGA_MIN:    0.4,
					OMEGA_MAX:    1.2,
					Alpha:        1.0,
					PSO_MAX_ITER: 20,
				},
			},
			{
				GaParams: hybrid.GaParams{
					CrossoverRate: 0.65,
					MutationRate:  0.95,
					GA_MAX_ITER:   200,
				},
				POPULATION_SIZE: 100,
				PsoParams: hybrid.PsoParams{
					C1:           1.5,
					C2:           2.5,
					OMEGA_MIN:    0.4,
					OMEGA_MAX:    1.2,
					Alpha:        1.0,
					PSO_MAX_ITER: 20,
				},
			},
			{
				GaParams: hybrid.GaParams{
					CrossoverRate: 0.65,
					MutationRate:  0.95,
					GA_MAX_ITER:   350,
				},
				POPULATION_SIZE: 150,
				PsoParams: hybrid.PsoParams{
					C1:           1.0,
					C2:           2.0,
					OMEGA_MIN:    0.4,
					OMEGA_MAX:    1.2,
					Alpha:        1.0,
					PSO_MAX_ITER: 50,
				},
			},
		}

		// Create stats folder
		statsFolder := "./benchmark/stats/"
		err := os.MkdirAll(statsFolder, os.ModePerm)
		if err != nil {
			fmt.Println("Error creating stats folder")
			os.Exit(1)
		}

		filePrefix := "./benchmark/instances/"
		for _, instanceName := range instances {
			filePath := filePrefix + instanceName
			for j, params := range parameters {
				modificationName := modifications[j]
				fmt.Printf("Excutando instância %v (%v)...\n", instanceName, modificationName)

				instance, err := hybrid.GetInstanceFromFile(filePath)
				if err != nil {
					fmt.Printf("Erro ao ler o arquivo: %v", filePath)
					os.Exit(1)
				}

				source := rand.New(rand.NewSource(int64(4502730368040452047)))

				rGa, rPso := hybrid.Run(instance, source, params)
				modificationFolder := statsFolder + modificationName + "/"
				// Create modification folder
				err = os.MkdirAll(modificationFolder, os.ModePerm)
				if err != nil {
					fmt.Println("Erro ao criar diretório para salvar dados.")
					os.Exit(1)
				}

				fileName := modificationFolder + instanceName
				rPso.SaveCsv(fileName + "_pso")
				rGa.SaveCsv(fileName + "_ga")
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(hybridCmd)
}
