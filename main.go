package main

import (
	"fmt"
	"job-shop-ga/ga"
)

func main() {
	// Nome do arquivo de entrada
	filename := "./benchmark/instances/abz5"

	// Ler instância do problema
	instance, err := ga.GetInstanceFromFile(filename)
	if err != nil {
		fmt.Println("Erro ao ler o arquivo:", err)
		return
	}
	instance.GenerateInitialPopulation(5)
	instance.Print()

}
