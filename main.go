package main

import (
	"flag"
	"fmt"
	"job-shop-ga/ga"
)

func main() {
	mutationRate := flag.Float64("mut", 0.3, "Experiment mutation rate")
	crossoverRate := flag.Float64("cross", 0.4, "Experiment crossover rate")
	populationSize := flag.Int("pop", 60, "Experiment population size")
	maxGenerations := flag.Int("gen", 1000, "Experiment max generations")

	// Nome do arquivo de entrada
	filename := "./benchmark/instances/abz6"

	// Ler instância do problema
	instance, err := ga.GetInstanceFromFile(filename, *mutationRate, *crossoverRate, *populationSize, *maxGenerations)
	if err != nil {
		fmt.Println("Erro ao ler o arquivo:", err)
		return
	}
	instance.GenerateInitialPopulation()
	instance.Print()
	instance.Run()
}
