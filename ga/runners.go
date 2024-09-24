package ga

import "sort"

func (instance *JobShopInstance) Run() ([]int, int) {
	instance.GenerateInitialPopulation()

	for i := 0; i < instance.maxGenerations; i++ {
		// Emabaralha a população
		shuffle(instance.Population, instance.Rand)

		children := make([]*Cromossome, 0)
		for j := 0; j < int(float64(instance.populationSize)*instance.crossoverRate); j += 2 {
			parent1 := instance.Population[j]
			parent2 := instance.Population[j+1]
			child1, child2 := instance.Crossover(parent1, parent2)
			children = append(children, &child1, &child2)
		}

		// Mutação na população como no artigo
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
		instance.Population = allIndividuals[:instance.populationSize]

		instance.calculateStats(i)
	}

	return instance.Population[0].genome, instance.Population[0].fitness
}

// Nessa modificação fazemos a mutação somente nos filhos.
func (instance *JobShopInstance) RunModMutation() ([]int, int) {
	instance.Name += "_modMutation"
	instance.GenerateInitialPopulation()

	for i := 0; i < instance.maxGenerations; i++ {
		// Emabaralha a população
		shuffle(instance.Population, instance.Rand)

		children := make([]*Cromossome, 0)
		for j := 0; j < int(float64(instance.populationSize)*instance.crossoverRate); j += 2 {
			parent1 := instance.Population[j]
			parent2 := instance.Population[j+1]
			child1, child2 := instance.Crossover(parent1, parent2)
			children = append(children, &child1, &child2)
		}

		for j := 0; j < len(children); j++ {
			if instance.Rand.Float64() < instance.mutationRate {
				instance.Mutate(children[j])
			}
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
		instance.Population = allIndividuals[:instance.populationSize]

		instance.calculateStats(i)
	}

	return instance.Population[0].genome, instance.Population[0].fitness
}

// Nessa modificação usamos o método de seleção por torneio.
func (instance *JobShopInstance) RunModTournament() ([]int, int) {
	instance.Name += "_modTournament"
	instance.GenerateInitialPopulation()

	for i := 0; i < instance.maxGenerations; i++ {
		// Emabaralha a população
		shuffle(instance.Population, instance.Rand)

		children := make([]*Cromossome, 0)
		for j := 0; j < int(float64(instance.populationSize)*instance.crossoverRate); j += 2 {
			parent1 := instance.Population[j]
			parent2 := instance.Population[j+1]
			child1, child2 := instance.Crossover(parent1, parent2)
			children = append(children, &child1, &child2)
		}

		// Mutação na população como no artigo
		for j := 0; j < int(float64(instance.populationSize)*instance.mutationRate); j++ {
			instance.Mutate(instance.Population[j])
		}

		allIndividuals := append(instance.Population, children...)

		// Calcular o makespan para cada indivíduo
		for _, individual := range allIndividuals {
			individual.fitness = instance.CalculateMakespan(individual)
		}

		winners := make([]*Cromossome, instance.populationSize)
		for j := 0; j < instance.populationSize; j++ {
			winners[j] = instance.TournamentSelection(allIndividuals, 2)
		}

		instance.Population = winners

		instance.calculateStats(i)
	}

	return instance.Population[0].genome, instance.Population[0].fitness
}

// Nessa modificação usamos o método de seleção por torneio e realizamos a mutação somente nos filhos.
func (instance *JobShopInstance) RunModTournamentMutation() ([]int, int) {
	instance.Name += "_modTournamentMutation"
	instance.GenerateInitialPopulation()

	for i := 0; i < instance.maxGenerations; i++ {
		// Emabaralha a população
		shuffle(instance.Population, instance.Rand)

		children := make([]*Cromossome, 0)
		for j := 0; j < int(float64(instance.populationSize)*instance.crossoverRate); j += 2 {
			parent1 := instance.Population[j]
			parent2 := instance.Population[j+1]
			child1, child2 := instance.Crossover(parent1, parent2)
			children = append(children, &child1, &child2)
		}

		for j := 0; j < len(children); j++ {
			if instance.Rand.Float64() < instance.mutationRate {
				instance.Mutate(children[j])
			}
		}

		allIndividuals := append(instance.Population, children...)

		// Calcular o makespan para cada indivíduo
		for _, individual := range allIndividuals {
			individual.fitness = instance.CalculateMakespan(individual)
		}

		winners := make([]*Cromossome, instance.populationSize)
		for j := 0; j < instance.populationSize; j++ {
			winners[j] = instance.TournamentSelection(allIndividuals, 2)
		}

		instance.Population = winners

		instance.calculateStats(i)
	}

	return instance.Population[0].genome, instance.Population[0].fitness
}

// Nessa modificação usamos o método de seleção por roleta.
func (instance *JobShopInstance) RunModRoulette() ([]int, int) {
	instance.Name += "_modRoulette"
	instance.GenerateInitialPopulation()

	for i := 0; i < instance.maxGenerations; i++ {
		// Emabaralha a população
		shuffle(instance.Population, instance.Rand)

		children := make([]*Cromossome, 0)
		for j := 0; j < int(float64(instance.populationSize)*instance.crossoverRate); j += 2 {
			parent1 := instance.Population[j]
			parent2 := instance.Population[j+1]
			child1, child2 := instance.Crossover(parent1, parent2)
			children = append(children, &child1, &child2)
		}

		// Mutação na população como no artigo
		for j := 0; j < int(float64(instance.populationSize)*instance.mutationRate); j++ {
			instance.Mutate(instance.Population[j])
		}

		allIndividuals := append(instance.Population, children...)

		// Calcular o makespan para cada indivíduo
		for _, individual := range allIndividuals {
			individual.fitness = instance.CalculateMakespan(individual)
		}

		winners := make([]*Cromossome, instance.populationSize)
		for j := 0; j < instance.populationSize; j++ {
			winners[j] = instance.RouletteWheelSelection(allIndividuals)
		}

		instance.Population = winners

		instance.calculateStats(i)
	}

	return instance.Population[0].genome, instance.Population[0].fitness
}

// Nessa modificação usamos o método de seleção por roleta e realizamos a mutação somente nos filhos.
func (instance *JobShopInstance) RunModRouletteMutation() ([]int, int) {
	instance.Name += "_modRouletteMutation"
	instance.GenerateInitialPopulation()

	for i := 0; i < instance.maxGenerations; i++ {
		// Emabaralha a população
		shuffle(instance.Population, instance.Rand)

		children := make([]*Cromossome, 0)
		for j := 0; j < int(float64(instance.populationSize)*instance.crossoverRate); j += 2 {
			parent1 := instance.Population[j]
			parent2 := instance.Population[j+1]
			child1, child2 := instance.Crossover(parent1, parent2)
			children = append(children, &child1, &child2)
		}

		for j := 0; j < len(children); j++ {
			if instance.Rand.Float64() < instance.mutationRate {
				instance.Mutate(children[j])
			}
		}

		allIndividuals := append(instance.Population, children...)

		// Calcular o makespan para cada indivíduo
		for _, individual := range allIndividuals {
			individual.fitness = instance.CalculateMakespan(individual)
		}

		winners := make([]*Cromossome, instance.populationSize)
		for j := 0; j < instance.populationSize; j++ {
			winners[j] = instance.RouletteWheelSelection(allIndividuals)
		}

		instance.Population = winners

		instance.calculateStats(i)
	}

	return instance.Population[0].genome, instance.Population[0].fitness
}
