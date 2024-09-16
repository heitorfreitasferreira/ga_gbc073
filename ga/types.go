package ga

type SubSetSumSolver interface {
	Solver(values []int, sum int) bool
}
