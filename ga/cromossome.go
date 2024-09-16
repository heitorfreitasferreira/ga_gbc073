package ga

import "math/rand/v2"

type GaSolver struct {
	Generations    uint
	PopulationSize uint

	population [][]bool

	SettingDifferenceDegree float64 // Diferencial do artigo é isso ir diminuindo
	coolingRatio            float64
}

func (s *GaSolver) Solve(values []uint, sum uint) bool {
	s.initPopulation(uint(len(values)))
	return false
}
func (s *GaSolver) initPopulation(cromSize uint) {
	s.population = make([][]bool, s.PopulationSize)
	for i := uint(0); i < s.PopulationSize; i++ {
		s.population[i] = make([]bool, cromSize)
		for j := range s.population[i] {
			s.population[i][j] = rand.Float64() < 0.5
		}
	}
}

func selectNParents(population [][]bool, n uint) ([][][]bool, []float64) {
	selected := make([][][]bool, n)
	diffDeg := make([]float64, n)
	for i := range selected {
		// Cada posição do vetor selected é um par de pais
		selected[i] = make([][]bool, 2)
	}
	takenIndexes := make(map[uint]bool)
	for i := uint(0); i < n; i++ {
		var p1, p2 []bool
		for {
			p1Index := uint(rand.IntN(len(population)))
			if takenIndexes[p1Index] {
				continue
			}
			p2Index := uint(rand.IntN(len(population)))
			if takenIndexes[p2Index] {
				continue
			}
			p1 = population[p1Index]
			p2 = population[p2Index]
			takenIndexes[p1Index] = true
			takenIndexes[p2Index] = true
			break
		}

		selected[i][0] = p1
		selected[i][1] = p2
		diffDeg[i] = differenceDegree(p1, p2)
	}
	return selected, diffDeg
}

func initializeRandomPopulation(pop *[][]bool, cromSize uint, popSize uint) {

}

func get_fitness_calculator(items []uint, C uint) func(cromossome []bool) uint {
	return func(cromossome []bool) uint {
		total := uint(0)
		for i, taken := range cromossome {
			if taken {
				total += items[i]
			}
		}

		feasible := C-total >= 0

		if !feasible {
			return total
		}
		return C - total
	}
}

func differenceDegree(p1, p2 []bool) float64 {
	var diff uint
	for i := range p1 {
		if p1[i] != p2[i] {
			diff++
		}
	}
	return float64(diff) / float64(len(p1))
}
