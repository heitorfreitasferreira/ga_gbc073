package ga

import (
	"fmt"
	"math"
)

// Cromossome representa um indivíduo na abordagem de representação indireta
type Cromossome struct {
	genome  []int
	fitness int
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

	shuffle(genome, instance.Rand)

	// Embaralhar os genes para criar um indivíduo aleatório (aqui pode usar alguma função de shuffle se necessário)
	// shuffle(genome) // Se precisar de embaralhamento

	return Cromossome{genome: genome, fitness: -1}
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

// CalcularFitness avalia a aptidão de um indivíduo com base no tempo total decorrido
func (instance *JobShopInstance) CalculateMakespan(cromossome *Cromossome) int {
	// Inicializar tempos de término para cada job e máquina
	jobCompletion := make([]int, instance.numJobs)           // Tempo de término dos jobs
	machineAvailability := make([]int, instance.numMachines) // Tempo de disponibilidade das máquinas
	unavailableMachinesByJob := make([][]bool, instance.numJobs)
	for i := 0; i < instance.numJobs; i++ {
		unavailableMachinesByJob[i] = make([]bool, instance.numMachines)
	}

	// Iterar sobre o genoma do indivíduo (sequência de jobs)
	for i := 0; i < len(cromossome.genome); i++ {
		jobID := cromossome.genome[i] // Identifica o job
		bestMachineID := -1
		bestMachineTime := math.MaxInt
		for machineID, unavailable := range unavailableMachinesByJob[jobID] {
			if unavailable {
				continue
			}

			machineTime := instance.jobs[jobID][machineID] + machineAvailability[machineID]
			if machineTime < bestMachineTime {
				bestMachineID = machineID
				bestMachineTime = machineTime
			}
		}
		unavailableMachinesByJob[jobID][bestMachineID] = true

		if bestMachineID == -1 {
			// Descobrir maquina com menor tempo ocioso ao começar em jobCompletion[jobID]
			panic("Não foi possível encontrar uma máquina disponível para o job")
		}

		processTime := instance.jobs[jobID][bestMachineID] // Tempo de processamento do job na máquina atual

		// O job só pode começar quando a máquina estiver disponível e o job estiver pronto (término da operação anterior)
		startTime := int(math.Max(float64(machineAvailability[bestMachineID]), float64(jobCompletion[jobID])))

		// Atualizar o tempo de término da operação para o job e a máquina
		jobCompletion[jobID] = startTime + processTime
		machineAvailability[bestMachineID] = startTime + processTime
	}

	// O fitness é o maior tempo de conclusão entre todos os jobs
	fitness := 0
	for _, completionTime := range jobCompletion {
		if completionTime > fitness {
			fitness = completionTime
		}
	}

	return fitness
}

func (instance *JobShopInstance) Mutate(cromossome *Cromossome) {
	// Escolher aleatoriamente dois genes para trocar
	idx1 := instance.Rand.Intn(len(cromossome.genome))
	idx2 := instance.Rand.Intn(len(cromossome.genome))
	// Trocar os genes
	cromossome.genome[idx1], cromossome.genome[idx2] = cromossome.genome[idx2], cromossome.genome[idx1]
}

// Print exibe o indivíduo gerado
func (ind *Cromossome) String() string {
	return fmt.Sprintf("%v", ind.genome)
}
