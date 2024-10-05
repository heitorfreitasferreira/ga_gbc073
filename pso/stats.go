package pso

import (
	"fmt"
	"os"
	"strings"
)

func statsToCsv(
	headers []string,
	data [][]string,
	csvSeparator string,
	file *os.File,
) {
	if len(headers) == 0 || len(data) == 0 {
		panic("headers and data must not be empty")
	}
	if len(headers) != len(data[0]) {
		panic("headers and data must have the same length")
	}

	sb := strings.Builder{}
	sb.WriteString(fmt.Sprintf("%s\n", strings.Join(headers, csvSeparator)))
	for _, row := range data {
		sb.WriteString(fmt.Sprintf("%s\n", strings.Join(row, csvSeparator)))
	}

	file.WriteString(sb.String())
}

func encodeToCsv(bestPos [][]float64, bestErr []float64, bestPosDiscrete [][]int, separator string) (headers []string, data [][]string) {
	if len(bestPos) != len(bestErr) || len(bestPos) != len(bestPosDiscrete) {
		panic("all slices must have the same length")
	}
	headers = []string{"Best Position", "Best Error", "Best Position Discrete"}

	data = make([][]string, len(bestPos))

	for i := 0; i < len(bestPos); i++ {
		data[i] = []string{
			strings.Join(float64ArrayToStringArray(bestPos[i]), separator),
			fmt.Sprintf("%f", bestErr[i]),
			strings.Join(intArrayToStringArray(bestPosDiscrete[i]), separator),
		}
	}
	return
}

func float64ArrayToStringArray(arr []float64) []string {
	strArr := make([]string, len(arr))
	for i, v := range arr {
		strArr[i] = fmt.Sprintf("%f", v)
	}
	return strArr
}

func intArrayToStringArray(arr []int) []string {
	strArr := make([]string, len(arr))
	for i, v := range arr {
		strArr[i] = fmt.Sprintf("%d", v)
	}
	return strArr
}
