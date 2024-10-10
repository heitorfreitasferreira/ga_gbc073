package job_shop_pso

import "math/rand"

func shuffle[T any](arr []T, source *rand.Rand) {
	n := len(arr)
	if n <= 1 {
		return
	}
	for i := n - 1; i > 0; i-- {
		j := source.Intn(i + 1)
		arr[i], arr[j] = arr[j], arr[i]
	}
}
