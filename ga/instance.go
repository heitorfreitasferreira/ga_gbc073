package ga

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

// JobShopInstance representa uma instância do problema de job shop scheduling
type JobShopInstance struct {
	numJobs     int
	numMachines int
	jobs        [][]int // Matriz contendo [tempo de processamento] para cada operação
	Population  []Cromossome
}

// GetInstanceFromFile lê uma instância do problema de um arquivo de texto
func GetInstanceFromFile(filename string) (*JobShopInstance, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var numJobs, numMachines int
	var jobs [][]int

	fmt.Println("Lendo instância do arquivo", filename)

	// Ignorar cabeçalhos e ler a primeira linha relevante com o número de jobs e máquinas
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, "#") || strings.TrimSpace(line) == "" {
			continue // Ignora linhas de comentário e vazias
		}
		parts := strings.Fields(line)
		if len(parts) == 2 {
			numJobs, _ = strconv.Atoi(parts[0])
			numMachines, _ = strconv.Atoi(parts[1])
			break
		}
	}

	fmt.Printf("%d %d\n", numJobs, numMachines)

	// Inicializa a matriz de jobs
	jobs = make([][]int, numJobs)

	// Lê cada linha do arquivo contendo as operações dos jobs
	for i := 0; i < numJobs; i++ {
		if scanner.Scan() {
			line := scanner.Text()
			parts := strings.Fields(line)
			jobTimes := make([]int, numMachines)

			fmt.Println(parts)
			for j := 0; j < len(parts); j += 2 {
				machineID, _ := strconv.Atoi(parts[j])
				jobTimeInMachine, _ := strconv.Atoi(parts[j+1])
				jobTimes[machineID] = jobTimeInMachine
			}
			jobs[i] = jobTimes
		}
	}

	return &JobShopInstance{
		numJobs:     numJobs,
		numMachines: numMachines,
		jobs:        jobs,
	}, nil
}

func (instance *JobShopInstance) GenerateInitialPopulation(size int) {
	instance.Population = make([]Cromossome, size)
	for i := 0; i < size; i++ {
		instance.Population[i] = GenerateCromossome(instance)
	}
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

// Função para realizar o crossover entre dois indivíduos
func (instance *JobShopInstance) Crossover(p1, p2 Cromossome) (Cromossome, Cromossome) {
	// Escolher aleatoriamente o índice de início e término para o trecho a ser trocado
	start1 := Source.Intn(len(p1.genome))
	end1 := Source.Intn(len(p1.genome)-start1) + start1

	start2 := Source.Intn(len(p2.genome))
	end2 := Source.Intn(len(p2.genome)-start2) + start2

	// Extrair pedaços necessários para o crossover
	body1 := p1.genome[start1 : end1+1]
	body2 := p2.genome[start2 : end2+1]
	head1 := p1.genome[:start1]
	tail1 := p1.genome[end1+1:]
	head2 := p2.genome[:start2]
	tail2 := p2.genome[end2+1:]

	// Ajustar os filhos para remover excessos e adicionar genes ausentes
	o1 := fixGenes(appendMultipleSlices(head1, body2, tail1), instance.numJobs, instance.numMachines)
	o2 := fixGenes(appendMultipleSlices(head2, body1, tail2), instance.numJobs, instance.numMachines)

	return Cromossome{genome: o1}, Cromossome{genome: o2}
}

func (instance *JobShopInstance) Mutate(cromossome Cromossome) Cromossome {
	// Escolher aleatoriamente dois genes para trocar
	idx1 := Source.Intn(len(cromossome.genome))
	idx2 := Source.Intn(len(cromossome.genome))
	// Trocar os genes
	cromossome.genome[idx1], cromossome.genome[idx2] = cromossome.genome[idx2], cromossome.genome[idx1]
	return cromossome
}

func (instance *JobShopInstance) Print() {
	fmt.Println("Número de jobs:", instance.numJobs)
	fmt.Println("Número de máquinas:", instance.numMachines)
	fmt.Println("Jobs:")
	for i, job := range instance.jobs {
		fmt.Printf("Job %2d %v\n", i+1, job)
	}
	fmt.Println("População:")
	for i, ind := range instance.Population {
		fmt.Printf("Indivíduo %d Makespan %d\n", i+1, instance.CalculateMakespan(&ind))
	}
}
