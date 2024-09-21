/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "job-shop-ga",
	Short: "Programa para resolver el problema de Job Shop Scheduling con algoritmos genéticos",
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.PersistentFlags().Int64("seed", 42, "Seed para el generador de números aleatorios")

	rootCmd.PersistentFlags().Float64("mut", 0.3, "Experiment mutation rate")
	rootCmd.PersistentFlags().Float64("cross", 0.4, "Experiment crossover rate")
	rootCmd.PersistentFlags().Int("pop", 60, "Experiment population size")
	rootCmd.PersistentFlags().Int("gen", 1000, "Experiment max generations")
	rootCmd.PersistentFlags().Bool("csv", false, "Exportar estatísticas a un archivo CSV")
}
