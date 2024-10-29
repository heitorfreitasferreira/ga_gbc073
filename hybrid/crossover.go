package hybrid

import "math/rand"

func crossover(ind1, ind2 Individual, cut int, instance JobShopInstance, alpha float64) (Individual, Individual) {
	if cut < 0 || cut > len(ind1.infoMatrix[0]) {
		panic("Invalid cut point")
	}

	size := len(ind1.infoMatrix[0])

	off1 := make([]int, size)
	off2 := make([]int, size)

	// Perform crossover
	for i := 0; i < cut; i++ {
		off1[i] = ind1.infoMatrix[0][i]
		off2[i] = ind2.infoMatrix[0][i]
	}
	for i := cut; i < size; i++ {
		off1[i] = ind2.infoMatrix[0][i]
		off2[i] = ind1.infoMatrix[0][i]
	}

	// Find duplicate and missing genes
	superfluous1 := make([]int, 0)
	superfluous2 := make([]int, 0)
	lacking1 := make([]int, 0)
	lacking2 := make([]int, 0)

	// Create maps to track gene counts
	count1 := make(map[int]int)
	count2 := make(map[int]int)

	// Count occurrences
	for i := 0; i < size; i++ {
		count1[off1[i]]++
		count2[off2[i]]++
	}

	// Find all numbers from 1 to size
	for i := 0; i < size; i++ {
		// Check first offspring
		if count1[i] > 1 {
			superfluous1 = append(superfluous1, i)
		} else if count1[i] == 0 {
			lacking1 = append(lacking1, i)
		}

		// Check second offspring
		if count2[i] > 1 {
			superfluous2 = append(superfluous2, i)
		} else if count2[i] == 0 {
			lacking2 = append(lacking2, i)
		}
	}

	// Replace duplicate genes with missing genes
	for i := 0; i < len(superfluous1); i++ {
		for j := 0; j < size; j++ {
			if off1[j] == superfluous1[i] && count1[off1[j]] > 1 {
				off1[j] = lacking1[i]
				count1[superfluous1[i]]--
				break
			}
		}
	}

	for i := 0; i < len(superfluous2); i++ {
		for j := 0; j < size; j++ {
			if off2[j] == superfluous2[i] && count2[off2[j]] > 1 {
				off2[j] = lacking2[i]
				count2[superfluous2[i]]--
				break
			}
		}
	}

	return newCromossome(instance, off1, alpha), newCromossome(instance, off2, alpha)
}

func inverseMutation(sequence []int, source *rand.Rand) {
	n := len(sequence)
	pos1 := source.Intn(n)
	pos2 := source.Intn(n)
	if pos1 > pos2 {
		pos1, pos2 = pos2, pos1
	}

	// Inverte a subsequência
	for i, j := pos1, pos2; i < j; i, j = i+1, j-1 {
		sequence[i], sequence[j] = sequence[j], sequence[i]
	}
}
