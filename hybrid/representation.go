package hybrid

type Individual [5][]int

func newRandomIndividual(size int) Individual {
	inst := make([][]int, size)
	for i := 0; i < size; i++ {
		inst[i] = make([]int, 5)
	}
	return Individual(inst)
}

func (i Individual) expandToMatrix(instance JobShopInstance) {

}

func (i Individual) decode() []int {
	return []int{}
}
