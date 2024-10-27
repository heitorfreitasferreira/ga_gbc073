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
	// Cria um array de índices
	indices := make([]int, len(pos))
	for i := range indices {
		indices[i] = i
	}

	// Ordena os índices com base nos valores de pos
	sort.Slice(indices, func(i, j int) bool {
		return pos[indices[i]] < pos[indices[j]]
	})

	// Cria um array para os resultados
	seq := make([]int, len(pos))
	for i, index := range indices {
		seq[index] = i
	}

	return seq
}
