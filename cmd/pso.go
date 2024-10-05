/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"job-shop-ga/pso"
	"math"
	"math/rand"
	"os"

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
		nParticles, _ := cmd.Flags().GetInt("n_particles")
		maxIter, _ := cmd.Flags().GetInt("max_iter")
		w, _ := cmd.Flags().GetFloat64("w")
		cCogni, _ := cmd.Flags().GetFloat64("c_cogni")
		cSocial, _ := cmd.Flags().GetFloat64("c_social")

		fmt.Printf("nParticles: %d, maxIter: %d, w: %f, cCogni: %f, cSocial: %f\n", nParticles, maxIter, w, cCogni, cSocial)
		file, err := os.Create("pso.csv")
		if err != nil {
			panic(err)
		}
		defer file.Close()
		fileInfo, err := file.Stat()
		if err != nil {
			panic(err)
		}
		fileSize := fileInfo.Size()
		if fileSize > 0 {
			// apaaga o conteúdo do arquivo
			file.Truncate(0)
		}
		randomSource := rand.New(rand.NewSource(42))
		maxDiscrete := 10.
		costFunc := func(x []float64) (float64, []int) {
			total := 0.0
			discrete := make([]int, len(x))
			for i, v := range x {
				total += math.Pow(v, 2)
				discrete[i] = int(math.Floor(1 + (v * maxDiscrete)))
			}
			return total, discrete
		}

		bounds := [][]float64{{-10, 10}, {-10, 10}}
		getInitialPos := func() []float64 {
			return []float64{5, 5}
		}
		getInitialVel := func() []float64 {
			x := randomSource.Float64()*2 - 1
			y := randomSource.Float64()*2 - 1
			return []float64{x, y}
		}
		getWs := func() []float64 {
			r := make([]float64, maxIter)
			for i := 0; i < maxIter; i++ {
				r[i] = w
			}
			return r
		}
		getCogni := func() []float64 {
			r := make([]float64, maxIter)
			for i := 0; i < maxIter; i++ {
				r[i] = cCogni
			}
			return r
		}
		getSocial := func() []float64 {
			r := make([]float64, maxIter)
			for i := 0; i < maxIter; i++ {
				r[i] = cSocial
			}
			return r
		}
		pso.Pso(costFunc, bounds, nParticles, 2, getWs(), getCogni(), getSocial(), maxIter, getInitialPos, getInitialVel, file, randomSource)
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
