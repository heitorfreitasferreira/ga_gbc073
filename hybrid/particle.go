package hybrid

import (
	"math"
	"math/rand"
	"sort"
)

type particle struct {
	Cromossome
	pos          []float64
	vel          []float64
	pBestPos     []float64
	pBestFitness float64
}

func randomParticle(params Parameters, instance JobShopInstance, source *rand.Rand) particle {
	dimension := instance.numJobs * instance.numMachines
	position := make([]float64, dimension)
	velocity := make([]float64, dimension)

	for i := range position {
		position[i] = source.Float64() // Aqui ficou confuso q ele não falou os limites
		velocity[i] = source.Float64() // aqui tbm
	}

	return particle{
		pos:          position,
		vel:          velocity,
		pBestPos:     position,
		pBestFitness: math.MinInt,
		Cromossome:   newCromossome(instance, getSequence(position), params.Alpha),
	}
}
func (part *particle) updateCromossome(instance JobShopInstance, alpha float64) {
	part.Cromossome = newCromossome(instance, getSequence(part.pos), alpha)
}

func getSequence(pos []float64) []int {
	type pair struct {
		value float64
		index int
	}
	pairs := make([]pair, len(pos))
	for i, v := range pos {
		pairs[i] = pair{v, i}
	}

	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i].value < pairs[j].value
	})

	order := make(map[float64]int)

	for i, p := range pairs {
		order[p.value] = i
	}
	seq := make([]int, len(pos))
	for i, pos := range pos {
		seq[i] = order[pos]
	}
	return seq
}
