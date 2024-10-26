package hybrid

import (
	"math"
	"strconv"
)

type Cromossome struct {
	infoMatrix
	fitness  float64
	makespan int
}

type infoMatrix [5][]int

func (ind infoMatrix) String() string {
	str := ""
	for i := 0; i < 5; i++ {
		for j := 0; j < len(ind[0]); j++ {
			str += strconv.Itoa(ind[i][j]) + ", "
		}
		str += "\n"
	}
	return str
}

func newCromossome(instance JobShopInstance, sequence []int, alpha float64) Cromossome {
	size := instance.numJobs * instance.numMachines
	inst := make([][]int, 5)
	for i := range inst {
		inst[i] = make([]int, size)
	}
	inst[0] = sequence

	matrix := infoMatrix(inst)
	matrix.expandToMatrix(instance)
	fitnessValue, makespanValue := fitness(matrix, instance, alpha)

	return Cromossome{infoMatrix: matrix, fitness: fitnessValue, makespan: makespanValue}

}

func (matrix infoMatrix) expandToMatrix(instance JobShopInstance) {
	// Assumindo q ind[0] está certo e instanciado e é o individuo

	for i, operation := range matrix[0] {
		job := int(math.Floor(float64(operation) / float64(instance.numMachines)))
		jobOperation := operation % instance.numMachines
		machine := instance.jobs[job][jobOperation][0]
		time := instance.jobs[job][jobOperation][1]

		matrix[1][i] = job
		matrix[2][i] = jobOperation
		matrix[3][i] = machine
		matrix[4][i] = time
	}
}

func (ind infoMatrix) calcMakespan(instance JobShopInstance) int {
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
