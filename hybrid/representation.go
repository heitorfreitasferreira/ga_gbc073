package hybrid

import (
	"job-shop-ga/ga"
	"math"
	"math/rand"
	"strconv"
)

type Individual [5][]int

func (ind Individual) String() string {
	str := ""
	for i := 0; i < 5; i++ {
		for j := 0; j < len(ind[0]); j++ {
			str += strconv.Itoa(ind[i][j]) + ", "
		}
		str += "\n"
	}
	return str
}
func NewRandomIndividual(instance JobShopInstance, source *rand.Rand) Individual {
	size := instance.numJobs * instance.numMachines
	inst := make([][]int, 5)
	for i := range inst {
		inst[i] = make([]int, size)
	}
	for i := 0; i < size; i++ {
		inst[0][i] = i
	}
	ga.Shuffle(inst[0], source)
	return Individual(inst)
}

func (ind Individual) ExpandToMatrix(instance JobShopInstance) {
	// Assumindo q ind[0] está certo e instanciado e é o individuo

	for i, operation := range ind[0] {
		job := int(math.Floor(float64(operation) / float64(instance.numMachines)))
		jobOperation := operation % instance.numMachines
		machine := instance.jobs[job][jobOperation][0]
		time := instance.jobs[job][jobOperation][1]

		ind[1][i] = job
		ind[2][i] = jobOperation
		ind[3][i] = machine
		ind[4][i] = time
	}

}
