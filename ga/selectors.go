package ga

import "sort"

func (instance *JobShopInstance) TournamentSelection(allIndividuals []*Cromossome, tournamentSize int) *Cromossome {
	tournament := make([]*Cromossome, tournamentSize)
	for i := 0; i < tournamentSize; i++ {
		tournament[i] = allIndividuals[instance.Rand.Intn(len(allIndividuals))]
	}
	sort.Slice(tournament, func(i, j int) bool {
		return tournament[i].fitness < tournament[j].fitness
	})
	return tournament[0]
}

func (instance *JobShopInstance) RouletteWheelSelection(allIndividuals []*Cromossome) *Cromossome {
	// Encontra pior valor de fitness para minimzar.
	maxFitness := 0
	for _, individual := range allIndividuals {
		if individual.fitness > maxFitness {
			maxFitness = individual.fitness
		}
	}

	// Calcula fitness invertido.
	totalInvertedFitness := 0
	for _, individual := range allIndividuals {
		totalInvertedFitness += maxFitness - individual.fitness
	}

	r := instance.Rand.Intn(totalInvertedFitness)

	// Seleciona o indivíduo baseado na roleta.
	for _, individual := range allIndividuals {
		r -= maxFitness - individual.fitness
		if r <= 0 {
			return individual
		}
	}

	return allIndividuals[len(allIndividuals)-1] // Por segurança.
}
