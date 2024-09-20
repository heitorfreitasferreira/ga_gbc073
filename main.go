package main

import (
	"fmt"
	"job-shop-ga/ga"
)

func main() {
	// Nome do arquivo de entrada
	filename := "./benchmark/instances/abz9"

	// Ler instância do problema
	instance, err := ga.GetInstanceFromFile(filename)
	if err != nil {
		fmt.Println("Erro ao ler o arquivo:", err)
		return
	}
	instance.GenerateInitialPopulation(5)
	instance.Print()

	parent1 := instance.Population[0]
	parent2 := instance.Population[1]

	instance.Crossover(parent1, parent2)

	/*
		fmt.Printf("Parent 1: %v\n", parent1)
		fmt.Printf("Parent 2: %v\n", parent2)
		fmt.Printf("Offspring 1: %v\n", off1)
		fmt.Printf("Offspring 2: %v\n", off2)
	*/
}
