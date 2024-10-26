package hybrid

import "math/rand"

func crossover(ind1, ind2 Cromossome, cut int, instance JobShopInstance) (Cromossome, Cromossome) {
	if cut < 0 || cut > len(ind1.infoMatrix[0]) {
		panic("Invalid cut point")
	}
	size := len(ind1.infoMatrix[0])
	off1 := make([][]int, 5)
	off2 := make([][]int, 5)
	for i := range off1 {
		off1[i] = make([]int, size)
		off2[i] = make([]int, size)
	}
	for i := 0; i < cut; i++ {
		off1[0][i] = ind1.infoMatrix[0][i]
		off2[0][i] = ind2.infoMatrix[0][i]
	}
	for i := cut; i < size; i++ {
		off1[0][i] = ind2.infoMatrix[0][i]
		off2[0][i] = ind1.infoMatrix[0][i]
	}

	moff1 := make(map[int]int)
	moff2 := make(map[int]int)

	for i := 0; i < size; i++ {
		moff1[off1[0][i]]++
		moff2[off2[0][i]]++
	}

	for i := 0; i < size; i++ {
		v1 := off1[0][i]
		v2 := off2[0][i]
		if moff1[v1] > 1 {
			if moff2[v1] == 0 {
				off2[0][i] = v1
				moff2[v1]++
			}
		}
		if moff2[v2] > 1 {
			if moff1[v2] == 0 {
				off1[0][i] = v2
				moff1[v2]++
			}
		}
	}
	return newCromossome(instance, off1[0], 0.0), newCromossome(instance, off2[0], 0.0)
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
