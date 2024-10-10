package job_shop_pso

import (
	"math"
	"math/rand"
)

type Particle struct {
	position     []int
	cost         int
	bestPosition []int
	bestCost     int
}

func (p *Particle) GetBestCost() int {
	return p.bestCost
}

func (p *Particle) UpdatePosition(
	globalBestPosition []int,
	weights []float64,
	source *rand.Rand,
	jobs int,
	machines int,
) {
	particleSize := jobs * machines
	newPosition := make([]int, particleSize)
	jobCount := make(map[int]int)
	globalIndex := 0
	currentIndex := 0
	bestIndex := 0

	samples := make([]float64, particleSize)
	for i := 0; i < particleSize; i++ {
		sample := source.Float64()
		samples[i] = sample
		var selectedJob int
		if sample <= weights[0] {
			for {
				if currentIndex < len(p.position) {
					selectedJob = p.position[currentIndex]
					currentIndex++
					if jobCount[selectedJob] < machines {
						break
					}
				}
			}
		} else if sample <= weights[0]+weights[1] {
			for {
				if bestIndex < len(p.bestPosition) {
					selectedJob = p.bestPosition[bestIndex]
					bestIndex++
					if jobCount[selectedJob] < machines {
						break
					}
				}
			}
		} else {
			for {
				if globalIndex < len(globalBestPosition) {
					selectedJob = globalBestPosition[globalIndex]
					globalIndex++
					if jobCount[selectedJob] < machines {
						break
					}
				}
			}
		}
		newPosition[i] = selectedJob
		jobCount[selectedJob]++
	}

	p.position = newPosition
}

func (p *Particle) Evaluate(instance JobShopPSO) {
	// Inicializar tempos de término para cada job e máquina
	jobCompletion := make([]int, instance.numJobs)           // Tempo de término dos jobs
	machineAvailability := make([]int, instance.numMachines) // Tempo de disponibilidade das máquinas
	jobNextOperation := make([]int, instance.numJobs)        // A próxima operação a ser agendada para cada job

	// Iterar sobre o genoma do indivíduo (sequência de jobs)
	for _, jobID := range p.position {
		// Pegar a próxima operação desse job
		opIndex := jobNextOperation[jobID]
		machineID := opIndex                           // A máquina correspondente à operação
		processTime := instance.jobs[jobID][machineID] // Tempo de processamento do job na máquina

		// Encontrar o tempo de início da operação, respeitando a precedência e a disponibilidade da máquina
		startTime := int(math.Max(float64(jobCompletion[jobID]), float64(machineAvailability[machineID])))

		// Atualizar o tempo de término da operação para o job e a máquina
		jobCompletion[jobID] = startTime + processTime
		machineAvailability[machineID] = startTime + processTime

		// Incrementar a próxima operação a ser agendada para o job
		jobNextOperation[jobID]++
	}

	// O makespan é o maior tempo de conclusão entre todos os jobs
	makespan := 0
	for _, completionTime := range jobCompletion {
		if completionTime > makespan {
			makespan = completionTime
		}
	}

	p.cost = makespan

	if p.cost < p.bestCost || p.bestCost == -1 {
		p.bestPosition = p.position
		p.bestCost = p.cost
	}
}
