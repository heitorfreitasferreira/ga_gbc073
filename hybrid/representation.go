package hybrid

import (
	"math"
	"strconv"
)

type Cromossome [5][]int

func (ind Cromossome) String() string {
	str := ""
	for i := 0; i < 5; i++ {
		for j := 0; j < len(ind[0]); j++ {
			str += strconv.Itoa(ind[i][j]) + ", "
		}
		str += "\n"
	}
	return str
}

func newIndividual(instance JobShopInstance, sequence []int) Cromossome {
	size := instance.numJobs * instance.numMachines
	inst := make([][]int, 5)
	for i := range inst {
		inst[i] = make([]int, size)
	}
	inst[0] = sequence
	return Cromossome(inst)
}

func (ind Cromossome) expandToMatrix(instance JobShopInstance) {
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

func (ind Cromossome) calcMakespan(instance JobShopInstance) int {
	// Assumindo q ind[0] está certo e instanciado e é o individuo
	machineTime := make([]int, instance.numMachines)
	jobTime := make([]int, instance.numJobs)
	for i := range ind[0] {
		job := ind[1][i]
		machine := ind[3][i]
		time := ind[4][i]

		startTime := int(math.Max(float64(machineTime[machine]), float64(jobTime[job])))
		finishTime := startTime + time

		machineTime[machine] = finishTime
		jobTime[job] = finishTime
	}

	max := math.MinInt

	for _, time := range jobTime {
		if time > max {
			max = time
		}
	}
	return max
}
