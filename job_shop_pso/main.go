package job_shop_pso

import (
	"bufio"
	"math"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

type JobShopPSO struct {
	currIteration int
	maxIterations int
	swarmSize     int
	bestParticle  Particle
	weigths       []float64
	source        *rand.Rand
	swarm         []Particle

	jobs        [][]int
	numJobs     int
	numMachines int
}

func GetInstanceFromFile(
	filename string,
	source *rand.Rand,
	weights []float64,
	swarmSize int,
	maxIterations int,
) (*JobShopPSO, error) {
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

	return &JobShopPSO{
		numJobs:     numJobs,
		numMachines: numMachines,
		source:      source,
		bestParticle: Particle{
			cost:     math.MaxInt,
			bestCost: math.MaxInt,
		},
		jobs:          jobs,
		weigths:       weights,
		swarmSize:     swarmSize,
		swarm:         make([]Particle, swarmSize),
		maxIterations: maxIterations,
	}, nil
}

func (pso *JobShopPSO) Run() (Particle, []Particle) {
	bestParticles := make([]Particle, pso.maxIterations)

	// Step 2
	for i := 0; i < pso.swarmSize; i++ {
		pso.swarm[i] = pso.generateParticle()
		pso.swarm[i].Evaluate(*pso)
		// Find best particle
		if pso.swarm[i].bestCost < pso.bestParticle.bestCost {
			pso.bestParticle = pso.swarm[i]
		}
	}

	bestParticles[pso.currIteration] = pso.bestParticle
	pso.currIteration++

	for pso.currIteration < pso.maxIterations {
		pso.update()
		bestParticles[pso.currIteration] = pso.bestParticle

		pso.currIteration++
	}

	return pso.bestParticle, bestParticles
}

func (pso *JobShopPSO) generateParticle() Particle {
	posSize := pso.numJobs * pso.numMachines
	genome := make([]int, posSize)

	idx := 0
	for j := 0; j < pso.numJobs; j++ {
		for m := 0; m < pso.numMachines; m++ {
			genome[idx] = j // Jobs são numerados a partir de 0
			idx++
		}
	}

	shuffle(genome, pso.source)

	newParticle := Particle{
		position:     genome,
		cost:         math.MaxInt,
		bestPosition: genome,
		bestCost:     math.MaxInt,
	}

	return newParticle
}

func (pso *JobShopPSO) update() {
	for i := 0; i < pso.swarmSize; i++ {
		pso.swarm[i].UpdatePosition(
			pso.bestParticle.position,
			pso.weigths,
			pso.source,
			pso.numJobs,
			pso.numMachines,
		)

		pso.swarm[i].Evaluate(*pso)

		if pso.bestParticle.bestCost > pso.swarm[i].bestCost {
			pso.bestParticle = pso.swarm[i]
		}
	}
}
