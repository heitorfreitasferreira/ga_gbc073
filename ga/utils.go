package ga

import (
	"math"
	"math/rand"
)

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

func max(arr []int) int {
	maxVal := math.MinInt
	for _, value := range arr {
		if value > maxVal {
			maxVal = value
		}
	}
	return maxVal
}

// AppendMultipleSlices appends multiple slices into one.
func appendMultipleSlices(slices ...[]int) []int {
	var result []int
	for _, slice := range slices {
		result = append(result, slice...)
	}
	return result
}
