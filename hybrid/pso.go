package hybrid

import (
	"math"
	"math/rand"
)

type pso struct {
	particles     []particle
	gBest         []float64
	gBestFitness  float64
	gBestMakespan int
	iteration     int
	instance      *JobShopInstance

	POPULATION_SIZE int
	PsoParams
}

func newPso(instance *JobShopInstance, params Parameters, source *rand.Rand) *pso {
	dimension := instance.numJobs * instance.numMachines
	psoInst := &pso{
		particles:       make([]particle, params.POPULATION_SIZE),
		gBest:           make([]float64, dimension),
		gBestFitness:    -1.0,
		gBestMakespan:   math.MaxInt,
		instance:        instance,
		PsoParams:       params.PsoParams,
		POPULATION_SIZE: params.POPULATION_SIZE,
	}

	// Inicializa população conforme seção 2.3
	for i := range psoInst.particles {
		psoInst.particles[i] = randomParticle(params, *instance, source)
	}

	return psoInst
}

func (psoInst *pso) getInitialPopulation(source *rand.Rand) ([]Individual, Result) {
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
		bestMakespans[i] = psoInst.gBestMakespan
	}

	// Mapeia os resultados do PSO para a representação do GA
	initialPopulation := make([]Individual, psoInst.POPULATION_SIZE)

	for i, particle := range psoInst.particles {
		initialPopulation[i] = particle.Individual
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

	part.updateCromossome(*psoInst.instance, psoInst.Alpha)
}

func (h *pso) evalParticle(p *particle) {
	if p.fitness > p.pBestFitness {
		p.pBestFitness = p.fitness
		copy(p.pBestPos, p.pos)
	}

	if p.fitness > h.gBestFitness {
		h.gBestFitness = p.fitness
		h.gBestMakespan = p.makespan
		copy(h.gBest, p.pos)
	}
}

func (ind *Individual) setFitness(instance JobShopInstance, alpha float64) {
	makespan := ind.calcMakespan(instance)

	M := 0.0
	for i := range ind.infoMatrix[0] {
		M += float64(ind.infoMatrix[4][i])
	}
	M /= float64(instance.numMachines)

	ind.fitness, ind.makespan = alpha*M/float64(makespan), makespan
}
