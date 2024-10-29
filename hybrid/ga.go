package hybrid

import (
	"fmt"
	"math/rand"
	"os"
	"sort"
	"strings"
)

// const (
//
//	OMEGA_MIN       = 0.4
//	OMEGA_MAX       = 1.2
//	C1              = 2.0 // Coeficiente de aceleração para pbest
//	C2              = 2.0 // Coeficiente de aceleração para gbest
//	PSO_MAX_ITER    = 20  // M1 no artigo
//	GA_MAX_ITER     = 200 // M2 no artigo
//	POPULATION_SIZE = 100
//	CROSSOVER_RATE  = 0.65
//	MUTATION_RATE   = 0.95
//
// )

type GaParams struct {
	CrossoverRate, MutationRate float64
	GA_MAX_ITER                 int
}

type PsoParams struct {
	C1, C2, OMEGA_MIN, OMEGA_MAX float64
	Alpha                        float64
	PSO_MAX_ITER                 int
}

type Parameters struct {
	GaParams
	PsoParams
	POPULATION_SIZE int
}

type Result struct {
	BestMakespans []int
	BestFitness   []float64
}

func (res Result) SaveCsv(filePath string) {
	file, err := os.Create(filePath + ".csv")
	if err != nil {
		fmt.Println("Erro ao criar arquivo de estatísticas:", err)
		panic(err)
	}
	defer file.Close()

	strBuilder := strings.Builder{}

	strBuilder.WriteString("iteration,makespan,fitness\n")

	for i := range res.BestMakespans {
		strBuilder.WriteString(fmt.Sprintf("%d,%d,%f\n", i, res.BestMakespans[i], res.BestFitness[i]))
	}

	file.WriteString(strBuilder.String())
}

func Run(inst *JobShopInstance, source *rand.Rand, params Parameters) (Result, Result) {
	var res, resPso Result
	psoInst := newPso(inst, params, source)
	cromossomes := make([]Individual, params.POPULATION_SIZE)
	cromossomes, resPso = psoInst.getInitialPopulation(source)

	for i := 0; i < params.GA_MAX_ITER; i++ {
		// Crossover
		for range cromossomes {
			if source.Float64() < params.CrossoverRate {
				p1Idx, p2Idx := source.Intn(params.POPULATION_SIZE), source.Intn(params.POPULATION_SIZE)
				cut := source.Intn(inst.numJobs * inst.numMachines)
				c1, c2 := crossover(cromossomes[p1Idx], cromossomes[p2Idx], cut, *inst, params.Alpha)

				cromossomes = append(cromossomes, c1)
				cromossomes = append(cromossomes, c2)
			}
		}
		// Mutação aleatória
		for i, ind := range cromossomes {
			if source.Float64() < params.MutationRate {
				inverseMutation(ind.infoMatrix[0], source)
				cromossomes[i] = newCromossome(*inst, ind.infoMatrix[0], params.Alpha)
			}
		}

		// Seleção elitista
		sort.Slice(cromossomes, func(i, j int) bool {
			return cromossomes[i].fitness > cromossomes[j].fitness
		})
		cromossomes = cromossomes[:params.POPULATION_SIZE]
		res.BestMakespans = append(res.BestMakespans, cromossomes[0].makespan)
		res.BestFitness = append(res.BestFitness, cromossomes[0].fitness)
	}

	return res, resPso
}
