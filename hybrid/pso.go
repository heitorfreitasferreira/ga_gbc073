package hybrid

import (
	"math"
	"math/rand"
)

type pso struct {
	particles     []particle
	gBest         []float64
	gBestSequence []int
	gBestFitness  float64
	iteration     int
	instance      *JobShopInstance

	POPULATION_SIZE int
	PsoParams
}

func newPso(instance *JobShopInstance, params PsoParams, popSize int, source *rand.Rand) *pso {
	dimension := instance.numJobs * instance.numMachines
	psoInst := &pso{
		particles:       make([]particle, popSize),
		gBest:           make([]float64, dimension),
		gBestSequence:   make([]int, dimension),
		gBestFitness:    math.MaxFloat64,
		instance:        instance,
		PsoParams:       params,
		POPULATION_SIZE: popSize,
	}

	// Inicializa população conforme seção 2.3
	for i := range psoInst.particles {
		psoInst.particles[i] = newParticle(dimension, source)
		psoInst.particles[i].updateSequence()
		psoInst.evalParticle(&psoInst.particles[i])
	}

	return psoInst
}

func (psoInst *pso) getInitialPopulation(source *rand.Rand) ([]Cromossome, Result) {
	// Executa o PSO conforme seção 2.3

	bestMakespans := make([]int, psoInst.PSO_MAX_ITER)
	bestFitness := make([]float64, psoInst.PSO_MAX_ITER)
	for i := 0; i < psoInst.PSO_MAX_ITER; i++ {
		psoInst.iteration = i
		for j := range psoInst.particles {
			psoInst.updateParticle(&psoInst.particles[j], source)
			psoInst.evalParticle(&psoInst.particles[j])
		}
		bestFitness[i] = psoInst.gBestFitness
		ind := newIndividual(*psoInst.instance, psoInst.gBestSequence)
		ind.expandToMatrix(*psoInst.instance)
		bestMakespans[i] = ind.calcMakespan(*psoInst.instance)
	}

	// Mapeia os resultados do PSO para a representação do GA
	initialPopulation := make([]Cromossome, psoInst.POPULATION_SIZE)

	for i := range psoInst.particles {
		initialPopulation[i] = newIndividual(*psoInst.instance, psoInst.particles[i].seq)
		initialPopulation[i].expandToMatrix(*psoInst.instance)
	}

	return initialPopulation, Result{BestMakespans: bestMakespans, BestFitness: bestFitness}
}

// Implementa as equações (1), (2) e (3) do artigo
func (psoInst *pso) updateParticle(part *particle, source *rand.Rand) {
	// Equação (1): Peso de inércia
	omega := psoInst.OMEGA_MAX - ((psoInst.OMEGA_MAX - psoInst.OMEGA_MIN) * float64(psoInst.iteration) / float64(psoInst.PSO_MAX_ITER))

	for i := range part.pos {
		r1, r2 := source.Float64(), source.Float64()

		// Equação (2)
		part.vel[i] = omega*part.vel[i] +
			psoInst.C1*r1*(part.pBestPos[i]-part.pos[i]) +
			psoInst.C2*r2*(psoInst.gBest[i]-part.pos[i])
		// Equação (3)
		part.pos[i] += part.vel[i]
	}

	part.updateSequence()
}

func (h *pso) evalParticle(p *particle) {
	ind := newIndividual(*h.instance, p.seq)
	ind.expandToMatrix(*h.instance)

	p.fitness = fitness(ind, *h.instance, h.Alpha)

	if p.fitness > p.pBestFitness {
		p.pBestFitness = p.fitness
		copy(p.pBestPos, p.pos)
	}

	if p.fitness > h.gBestFitness {
		h.gBestFitness = p.fitness
		copy(h.gBest, p.pos)
		copy(h.gBestSequence, p.seq)
	}
}

func fitness(ind Cromossome, instance JobShopInstance, alpha float64) float64 {
	makespan := ind.calcMakespan(instance)

	M := 0.0
	for i := range ind {
		M += float64(ind[i][4])
	}
	M /= float64(instance.numMachines)

	return alpha * M / float64(makespan)
}
