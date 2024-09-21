package ga

import (
	"encoding/csv"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
)

type evolutionStats struct {
	best, worst, median []int
	stdDev, avg         []float64
}

func (instance *JobShopInstance) calculateStats(generation int) {
	best := math.MinInt
	worst := math.MaxInt
	sum := 0
	var varianceSum float64
	mkspans := make([]int, instance.populationSize)

	for i, individual := range instance.Population {
		if individual.fitness < best {
			best = individual.fitness
		}
		if individual.fitness > worst {
			worst = individual.fitness
		}
		sum += individual.fitness
		mkspans[i] = individual.fitness
	}

	// Ordenar os makespans para calcular a mediana
	sort.Ints(mkspans)

	avg := float64(sum) / float64(instance.populationSize)

	// Calcular o desvio padrão
	for _, mkspan := range mkspans {
		varianceSum += math.Pow(float64(mkspan)-avg, 2)
	}

	stdDev := math.Sqrt(varianceSum / float64(instance.populationSize))

	// Atualizar as estatísticas de evolução
	instance.evolutionStats.best[generation] = best
	instance.evolutionStats.worst[generation] = worst
	instance.evolutionStats.avg[generation] = avg
	instance.evolutionStats.median[generation] = mkspans[instance.populationSize/2]
	instance.evolutionStats.stdDev[generation] = stdDev
}

func (s evolutionStats) save(instanceName string) {
	file, err := os.Create(fmt.Sprintf("./stats/%s.csv", instanceName))
	if err != nil {
		fmt.Println("Erro ao criar arquivo de estatísticas:", err)
		return
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	writer.Write([]string{"generation", "best", "worst", "median", "average", "std_dev"})

	for i := 0; i < len(s.best); i++ {
		writer.Write([]string{
			strconv.Itoa(i),
			strconv.Itoa(s.best[i]),
			strconv.Itoa(s.worst[i]),
			strconv.Itoa(s.median[i]),
			strconv.FormatFloat(s.avg[i], 'f', 3, 64),
			strconv.FormatFloat(s.stdDev[i], 'f', 3, 64),
		})
	}
}
