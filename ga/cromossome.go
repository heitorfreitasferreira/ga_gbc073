package ga

import (
	"fmt"
)

// Cromossome representa um indivíduo na abordagem de representação indireta
type Cromossome struct {
	genome []int
}

// GenerateCromossome cria um indivíduo com a abordagem de representação indireta
func GenerateCromossome(instance *JobShopInstance) Cromossome {
	numGenes := instance.numJobs * instance.numMachines
	genome := make([]int, numGenes)

	// Gera um genome com a lista de jobs para cada operação de máquina
	idx := 0
	for j := 0; j < instance.numJobs; j++ {
		for m := 0; m < instance.numMachines; m++ {
			genome[idx] = j + 1 // Jobs são numerados a partir de 1
			idx++
		}
	}

	shuffle(genome)

	// Embaralhar os genes para criar um indivíduo aleatório (aqui pode usar alguma função de shuffle se necessário)
	// shuffle(genome) // Se precisar de embaralhamento

	return Cromossome{genome: genome}
}

// Print exibe o indivíduo gerado
func (ind *Cromossome) Print() {
	fmt.Println("Indivíduo:", ind.genome)
}
