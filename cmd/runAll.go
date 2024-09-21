/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"job-shop-ga/ga"
	"log"
	"math/rand"
	"os"
	"sync"

	"github.com/spf13/cobra"
)

// runAllCmd represents the runAll command
var runAllCmd = &cobra.Command{
	Use:   "runAll",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		folder := "./benchmark/instances"
		files, err := os.ReadDir(folder)
		if err != nil {
			log.Fatalf("failed reading directory: %s", err)
		}
		mutationRate, _ := cmd.Flags().GetFloat64("mut")
		crossoverRate, _ := cmd.Flags().GetFloat64("cross")
		populationSize, _ := cmd.Flags().GetInt("pop")
		maxGenerations, _ := cmd.Flags().GetInt("gen")
		seed, _ := cmd.Flags().GetInt("seed")
		wg := sync.WaitGroup{}
		for _, file := range files {
			if file.IsDir() {
				continue
			}

			wg.Add(1)
			go func() {
				defer wg.Done()
				fileName := folder + "/" + file.Name()
				instance, err := ga.GetInstanceFromFile(fileName, mutationRate, crossoverRate, populationSize, maxGenerations, rand.New(rand.NewSource(int64(seed))))
				if err != nil {
					fmt.Println("Erro ao ler o arquivo:", err)
					return
				}
				instance.GenerateInitialPopulation()
				_, _ = instance.Run()

				instance.ToCsv()

			}()
		}
		wg.Wait()
	},
}

func init() {
	rootCmd.AddCommand(runAllCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// runAllCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// runAllCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
