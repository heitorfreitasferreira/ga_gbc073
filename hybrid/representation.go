package hybrid

import (
	"fmt"
	"math"
)

type Individual struct {
	infoMatrix
	fitness  float64
	makespan int
}

type infoMatrix [5][]int

func (ind infoMatrix) String() string {
	str := ""
	for i := 0; i < 5; i++ {
		for j := 0; j < len(ind[0]); j++ {
			str += fmt.Sprintf("%2d ", ind[i][j])
		}
		str += "\n"
	}
	return str
}

func newCromossome(instance JobShopInstance, sequence []int, alpha float64) Individual {
	size := instance.numJobs * instance.numMachines
	inst := make([][]int, 5)
	for i := range inst {
		inst[i] = make([]int, size)
	}
	inst[0] = sequence

	matrix := infoMatrix(inst)
	matrix.expandToMatrix(instance)

	crom := Individual{infoMatrix: matrix}
	crom.expandToMatrix(instance)
	crom.setFitness(instance, alpha)
	return crom
}

func (matrix infoMatrix) expandToMatrix(instance JobShopInstance) {
	// Assumindo q ind[0] está certo e instanciado e é o individuo

	for i, operation := range matrix[0] {
		job := int(math.Floor(float64(operation) / float64(instance.numMachines)))
		if job > instance.numJobs-1 { // -1 pois a contagem começa em 0
			panic(fmt.Errorf("job nr {%d} para um problema com {%d} jobs \n\n%v\nMatrix expandida:\n%v\n%v\n%v\n%v\n%v", job, instance.numJobs, instance, matrix[0], matrix[1], matrix[2], matrix[3], matrix[4]))
		}
		jobOperation := operation % instance.numMachines
		if jobOperation > instance.numMachines-1 { // -1 pois a contagem começa em 0
			panic(fmt.Errorf("operação nr {%d} para um problema com {%d} máquinas \n\n%v\nMatrix expandida:\n%v\n%v\n%v\n%v\n%v", jobOperation, instance.numMachines, instance, matrix[0], matrix[1], matrix[2], matrix[3], matrix[4]))
		}
		machine := instance.jobs[job][jobOperation][0]

		if machine > instance.numMachines-1 { // -1 pois a contagem começa em 0
			panic(fmt.Errorf("máquina nr {%d} para um problema com {%d} máquinas \n\n%v\nMatrix expandida:\n%v\n%v\n%v\n%v\n%v", machine, instance.numMachines, instance, matrix[0], matrix[1], matrix[2], matrix[3], matrix[4]))
		}

		time := instance.jobs[job][jobOperation][1]

		matrix[1][i] = job
		matrix[2][i] = jobOperation
		matrix[3][i] = machine
		matrix[4][i] = time
	}
}

func (ind infoMatrix) calcMakespan(instance JobShopInstance) int {

	machineTime := make([]int, instance.numMachines)
	jobTime := make([]int, instance.numJobs)

	type tuple struct {
		generalOperation, job, jobOperation, machine, time int
	}
	oldOrder := make([]tuple, len(ind[0]))
	for i := range ind[0] {
		oldOrder[i] = tuple{
			generalOperation: ind[0][i],
			job:              ind[1][i],
			jobOperation:     ind[2][i],
			machine:          ind[3][i],
			time:             ind[4][i],
		}
	}

	actual := make([]tuple, len(ind[0]))

	used := make([]bool, len(ind[0]))
	nextOrderToExecute := make([]int, instance.numJobs)
	scheduled := 0
	i := 0
	for scheduled < len(ind[0]) {
		if i >= len(ind[0]) {
			i = 0
		}
		if used[i] {
			i++
			continue
		}
		if nextOrderToExecute[oldOrder[i].job] == oldOrder[i].jobOperation {
			actual[scheduled] = oldOrder[i]
			used[i] = true
			scheduled++
			nextOrderToExecute[oldOrder[i].job]++
		}
		i++
	}

	var greatestFinishingTime int
	for _, tpl := range actual {
		job := tpl.job
		machine := tpl.machine
		time := tpl.time

		startTime := int(math.Max(float64(machineTime[machine]), float64(jobTime[job])))
		finishTime := startTime + time

		machineTime[machine] = finishTime
		jobTime[job] = finishTime
		if finishTime > greatestFinishingTime {
			greatestFinishingTime = finishTime
		}
	}
	return greatestFinishingTime
}
