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
			genome[idx] = j // Jobs são numerados a partir de 0
			idx++
		}
	}

	shuffle(genome)

	// Embaralhar os genes para criar um indivíduo aleatório (aqui pode usar alguma função de shuffle se necessário)
	// shuffle(genome) // Se precisar de embaralhamento

	return Cromossome{genome: genome}
}

func countJobsOccurrences(genome []int) map[int]int {
	count := make(map[int]int)
	for _, job := range genome {
		count[job]++
	}
	return count
}

// Função auxiliar para ajustar os genes após o crossover
func fixGenes(offspring []int, numJobs int, numMachines int) []int {
	// Contar a ocorrência de cada gene no offspring
	jobCount := countJobsOccurrences(offspring)

	// Remover genes em excesso (se houver repetições)
	i := 0
	for i < len(offspring) {
		if jobCount[offspring[i]] > numMachines {
			jobCount[offspring[i]] = jobCount[offspring[i]] - 1
			offspring = append(offspring[:i], offspring[i+1:]...)
		} else {
			i++
		}
	}

	// Adicionar genes que estejam faltando
	for jobID := 0; jobID < numJobs; jobID++ {
		missing := numMachines - jobCount[jobID] // Quantos genes estão faltando para o job atual

		for missing > 0 {
			// Podemos adicionar o job faltante em qualquer posição; aqui, vamos adicionar no final
			offspring = append(offspring, jobID)
			missing--
		}
	}

	return offspring
}

// Print exibe o indivíduo gerado
func (ind *Cromossome) String() string {
	return fmt.Sprintf("%v", ind.genome)
}
