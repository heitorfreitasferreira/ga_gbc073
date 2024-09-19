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
	jobs        [][]int // Matriz contendo [máquina, tempo de processamento] para cada operação
	population  []Cromossome
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
	instance.population = make([]Cromossome, size)
	for i := 0; i < size; i++ {
		instance.population[i] = GenerateCromossome(instance)
	}
}

// CalcularFitness avalia a aptidão de um indivíduo com base no tempo total decorrido
func (instance *JobShopInstance) CalculateFitness(cromossome *Cromossome) int {
	// Inicializar tempos de término para cada job e máquina
	jobCompletion := make([]int, instance.numJobs)           // Tempo de término dos jobs
	machineAvailability := make([]int, instance.numMachines) // Tempo de disponibilidade das máquinas

	// Iterar sobre o genoma do indivíduo (sequência de jobs)
	for i := 0; i < len(cromossome.genome); i++ {
		jobID := cromossome.genome[i] - 1     // Identifica o job (ajuste para 0-indexado)
		machineID := i % instance.numMachines // Identifica a máquina (cíclico sobre as máquinas)

		processTime := instance.jobs[jobID][machineID] // Tempo de processamento do job na máquina atual

		// O job só pode começar quando a máquina estiver disponível e o job estiver pronto (término da operação anterior)
		startTime := int(math.Max(float64(machineAvailability[machineID]), float64(jobCompletion[jobID])))

		// Atualizar o tempo de término da operação para o job e a máquina
		jobCompletion[jobID] = startTime + processTime
		machineAvailability[machineID] = startTime + processTime
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

func (instance *JobShopInstance) Print() {
	fmt.Println("Número de jobs:", instance.numJobs)
	fmt.Println("Número de máquinas:", instance.numMachines)
	fmt.Println("Jobs:")
	for i, job := range instance.jobs {
		fmt.Printf("Job %2d %v\n", i+1, job)
	}
	fmt.Println("População:")
	for i, ind := range instance.population {
		fmt.Printf("Indivíduo %d Fitness %d\n", i+1, instance.CalculateFitness(&ind))
	}
}
