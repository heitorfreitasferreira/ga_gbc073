package ga

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

// JobShopInstance representa uma instância do problema de job shop scheduling
type JobShopInstance struct {
	numJobs        int
	numMachines    int
	jobs           [][]int // Matriz contendo [tempo de processamento] para cada operação
	Population     []*Cromossome
	mutationRate   float64
	crossoverRate  float64
	populationSize int
	maxGenerations int

	evolutionStats
}

// GetInstanceFromFile lê uma instância do problema de um arquivo de texto
func GetInstanceFromFile(filename string, mutationRate, crossoverRate float64, populationSize, maxGenerations int) (*JobShopInstance, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var numJobs, numMachines int
	var jobs [][]int

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

	// Inicializa a matriz de jobs
	jobs = make([][]int, numJobs)

	// Lê cada linha do arquivo contendo as operações dos jobs
	for i := 0; i < numJobs; i++ {
		if scanner.Scan() {
			line := scanner.Text()
			parts := strings.Fields(line)
			jobTimes := make([]int, numMachines)

			for j := 0; j < len(parts); j += 2 {
				machineID, _ := strconv.Atoi(parts[j])
				jobTimeInMachine, _ := strconv.Atoi(parts[j+1])
				jobTimes[machineID] = jobTimeInMachine
			}
			jobs[i] = jobTimes
		}
	}

	return &JobShopInstance{
		numJobs:        numJobs,
		numMachines:    numMachines,
		jobs:           jobs,
		mutationRate:   mutationRate,
		crossoverRate:  crossoverRate,
		populationSize: populationSize,
		maxGenerations: maxGenerations,
		evolutionStats: evolutionStats{
			best:   make([]int, maxGenerations),
			worst:  make([]int, maxGenerations),
			median: make([]int, maxGenerations),
			avg:    make([]float64, maxGenerations),
			stdDev: make([]float64, maxGenerations),
		},
	}, nil
}

func (instance *JobShopInstance) GenerateInitialPopulation() {
	instance.Population = make([]*Cromossome, instance.populationSize)
	for i := 0; i < instance.populationSize; i++ {

		ind := GenerateCromossome(instance)
		instance.Population[i] = &ind
	}
}

// Função para realizar o crossover entre dois indivíduos
func (instance *JobShopInstance) Crossover(p1, p2 *Cromossome) (Cromossome, Cromossome) {
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

func (instance *JobShopInstance) Run() ([]int, int) {
	instance.GenerateInitialPopulation()

	for i := 0; i < instance.maxGenerations; i++ {
		// Emabaralha a população
		shuffle(instance.Population)

		children := make([]*Cromossome, 0)
		for j := 0; j < int(float64(instance.populationSize)*instance.crossoverRate); j += 2 {
			parent1 := instance.Population[j]
			parent2 := instance.Population[j+1]
			child1, child2 := instance.Crossover(parent1, parent2)
			children = append(children, &child1, &child2)
		}

		// Mutação na população
		for j := 0; j < int(float64(instance.populationSize)*instance.mutationRate); j++ {
			instance.Mutate(instance.Population[j])
		}

		allIndividuals := append(instance.Population, children...)

		// Calcular o makespan para cada indivíduo
		for _, individual := range allIndividuals {
			individual.fitness = instance.CalculateMakespan(individual)
		}

		sort.Slice(allIndividuals, func(i, j int) bool {
			return allIndividuals[i].fitness < allIndividuals[j].fitness
		})

		// Copy the best populationSize individuals to the next generation
		for j := 0; j < instance.populationSize; j++ {
			instance.Population[j] = allIndividuals[j]
		}
		instance.calculateStats(i)
	}

	return instance.Population[0].genome, instance.Population[0].fitness
}

func (instance *JobShopInstance) Print() {
	fmt.Println("Número de jobs:", instance.numJobs)
	fmt.Println("Número de máquinas:", instance.numMachines)
	fmt.Println("Taxa de mutação:", instance.mutationRate)
	fmt.Println("Taxa de crossover:", instance.crossoverRate)
	fmt.Println("Tamanho da população:", instance.populationSize)
	fmt.Println("Número máximo de gerações:", instance.maxGenerations)
	fmt.Println("Jobs:")
	for i, job := range instance.jobs {
		fmt.Printf("Job %2d %v\n", i+1, job)
	}
}
