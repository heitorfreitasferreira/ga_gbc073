package hybrid

import "math/rand"

func crossover(ind1, ind2 Cromossome, cut int) (Cromossome, Cromossome) {
	if cut < 0 || cut > len(ind1[0]) {
		panic("Invalid cut point")
	}
	size := len(ind1[0])
	off1 := make([][]int, 5)
	off2 := make([][]int, 5)
	for i := range off1 {
		off1[i] = make([]int, size)
		off2[i] = make([]int, size)
	}
	for i := 0; i < cut; i++ {
		off1[0][i] = ind1[0][i]
		off2[0][i] = ind2[0][i]
	}
	for i := cut; i < size; i++ {
		off1[0][i] = ind2[0][i]
		off2[0][i] = ind1[0][i]
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
	return Cromossome(off1), Cromossome(off2)
}

// Esse aqui é a versão da claudia
func singleLinearOrderCrossover(p1, p2 []int, point int) ([]int, []int) {
	n := len(p1)
	child1 := make([]int, n)
	child2 := make([]int, n)

	// Copia até o ponto de crossover
	copy(child1[:point], p1[:point])
	copy(child2[:point], p2[:point])

	// Copia o resto do outro pai
	copy(child1[point:], p2[point:])
	copy(child2[point:], p1[point:])

	// Corrige genes duplicados
	used1 := make(map[int]bool)
	used2 := make(map[int]bool)
	missing1 := make([]int, 0)
	missing2 := make([]int, 0)

	// Encontra duplicatas e faltantes
	for i := 0; i < n; i++ {
		if used1[child1[i]] {
			missing1 = append(missing1, i)
		}
		used1[child1[i]] = true

		if used2[child2[i]] {
			missing2 = append(missing2, i)
		}
		used2[child2[i]] = true
	}

	// Encontra números faltantes
	for i := 1; i <= n; i++ {
		if !used1[i] {
			child1[missing1[0]] = i
			missing1 = missing1[1:]
		}
		if !used2[i] {
			child2[missing2[0]] = i
			missing2 = missing2[1:]
		}
	}

	return child1, child2
}

// Versão da claudia
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

func mutate(ind Cromossome, source *rand.Rand) {
	start := source.Intn(len(ind[0]))
	end := source.Intn(len(ind[0])-start) + start
	j := end
	for i := start; i < (end-start)/2; i++ {
		ind[0][i], ind[0][j] = ind[0][j], ind[0][i]
		j--
	}
}
