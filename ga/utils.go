package ga

import (
	"math"
	"math/rand"
	"time"
)

func shuffle(arr []int) {
	source := rand.New(rand.NewSource(time.Now().UnixNano()))
	n := len(arr)
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
