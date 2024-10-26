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

func newParticle(params Parameters, instance JobShopInstance, source *rand.Rand) particle {
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
		Cromossome:   newCromossome(instance, position, alpha),
	}
}

func (p particle) updateSequence() {
	type pair struct {
		value float64
		index int
	}
	pairs := make([]pair, len(p.pos))
	for i, v := range p.pos {
		pairs[i] = pair{v, i}
	}

	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i].value < pairs[j].value
	})

	order := make(map[float64]int)

	for i, p := range pairs {
		order[p.value] = i
	}
	for i, pos := range p.pos {
		p.seq[i] = order[pos]
	}
}
