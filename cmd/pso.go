/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"encoding/csv"
	"fmt"
	"job-shop-ga/job_shop_pso"
	"math/rand"
	"os"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
)

// psoCmd represents the pso command
var psoCmd = &cobra.Command{
	Use:   "pso",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fileNames := []string{
			"./benchmark/instances/ft06",
			"./benchmark/instances/ft10",
			"./benchmark/instances/ft20",
			"./benchmark/instances/abz9",
			"./benchmark/instances/swv12",
			"./benchmark/instances/ta75",
		}

		seed := 42
		allWeights := [][]float64{{0.3, 0.5, 0.2}, {0.5, 0.2, 0.3}, {0.2, 0.5, 0.3}}
		swarmSizes := []int{50, 100, 150}
		maxIterations := 50

		for _, fileName := range fileNames {
			for _, swarmSize := range swarmSizes {
				for _, weights := range allWeights {
					instanceName := strings.Split(fileName, "instances/")[1]
					fmt.Println("Running instance", instanceName)
					fmt.Println("\tWeights:", weights)
					fmt.Println("\tSwarm size:", swarmSize)

					randSource := rand.New(rand.NewSource(int64(seed)))
					instance, error := job_shop_pso.GetInstanceFromFile(
						fileName,
						randSource,
						weights,
						swarmSize,
						maxIterations,
					)

					if error != nil {
						fmt.Println("Erro ao ler o arquivo:", error)
						continue
					}

					_, bestParticles := instance.Run()

					file, err := os.Create(
						fmt.Sprintf(
							"./benchmark/stats_pso_mods/%s_%.2f-%.2f-%.2f_%d.csv",
							instanceName,
							weights[0],
							weights[1],
							weights[2],
							swarmSize,
						),
					)
					if err != nil {
						fmt.Println("Erro ao criar arquivo de estatísticas:", err)
						return
					}
					defer file.Close()

					writer := csv.NewWriter(file)
					defer writer.Flush()

					writer.Write([]string{"iteration", "makespan"})
					for i := 0; i < len(bestParticles); i++ {
						writer.Write([]string{
							strconv.Itoa(i),
							strconv.Itoa(bestParticles[i].GetBestCost()),
						})
					}
				}
			}

		}
	},
}

func init() {
	rootCmd.AddCommand(psoCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// psoCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// psoCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	psoCmd.Flags().Int("n_particles", 10, "Number of particles")
	psoCmd.Flags().Int("max_iter", 100, "Number of iterations")
	psoCmd.Flags().Float64("w", 0.5, "Inertia")
	psoCmd.Flags().Float64("c_cogni", 1.0, "Cognitive component")
	psoCmd.Flags().Float64("c_social", 2.0, "Social component")

}
