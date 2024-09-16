package ga

func (s *GaSolver) generateChildren(population *[][]bool, popSize uint) {
	Nc := uint(0)
	children := make([][]bool, popSize)
	for Nc < popSize {
		numberOfSelectedParents := (popSize - Nc) / 2
		selectedParents, diffDeg := selectNParents(*population, numberOfSelectedParents)

	}
}
