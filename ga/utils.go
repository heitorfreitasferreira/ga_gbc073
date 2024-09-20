package ga

import (
	"math"
	"math/rand"
)

var Source = rand.New(rand.NewSource(42))

func shuffle(arr []int) {
	n := len(arr)
	for i := n - 1; i > 0; i-- {
		j := Source.Intn(i + 1)
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
