package dp

type MemoSolver struct {
	memo [][]int
}

func InitMemoSolver(solver *MemoSolver, n, sum int) {
	solver.memo = make([][]int, n+1)
	for i := 0; i < n; i++ {
		solver.memo[i] = make([]int, sum+1)
		for j := range solver.memo[i] {
			solver.memo[i][j] = -1
		}
	}
}
func (s *MemoSolver) Solve(values []int, sum int) bool {
	n := len(values)
	return s.solve(n, sum, values) > 0
}

func (solver *MemoSolver) solve(n, sum int, values []int) int {
	if sum == 0 {
		return 1
	}

	if n == 0 {
		return 0
	}

	if solver.memo[n-1][sum] != -1 {
		return solver.memo[n-1][sum]
	}

	if values[n-1] > sum {
		solver.memo[n-1][sum] = solver.solve(n-1, sum, values)
		return solver.memo[n-1][sum]
	}
	solver.memo[n-1][sum] = solver.solve(n-1, sum, values)
	return solver.memo[n-1][sum] + solver.solve(n-1, sum-values[n-1], values)
}
