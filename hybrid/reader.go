package hybrid

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type JobShopInstance struct {
	Name        string
	numJobs     int
	numMachines int
	jobs        [][][2]int // Matriz[Job][Operação][Máquina, Tempo]
}

func (instance JobShopInstance) String() string {
	machineMatrix := ""
	for i := range instance.jobs {
		for j := range instance.jobs[i] {
			machineMatrix += fmt.Sprintf("%2d %2d|", instance.jobs[i][j][0], instance.jobs[i][j][1])
		}
		machineMatrix += "\n"
	}
	return fmt.Sprintf("Instance: %s\nJobs: %d\nMachines: %d\n\n%s", instance.Name, instance.numJobs, instance.numMachines, machineMatrix)
}

// GetInstanceFromFile lê uma instância do problema de um arquivo de texto
func GetInstanceFromFile(filename string) (*JobShopInstance, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var numJobs, numMachines int
	var jobs [][][2]int

	// Ignorar cabeçalhos e ler a primeira linha relevante com o número de jobs e máquinas
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, "#") || strings.TrimSpace(line) == "" {
			continue // Ignora linhas de comentário e vazias
		}
		parts := strings.Fields(line)
		if len(parts) == 2 {
			numJobs, _ = strconv.Atoi(parts[0])
			numMachines, _ = strconv.Atoi(parts[1])
			break
		}
	}

	// Inicializa a matriz de jobs
	jobs = make([][][2]int, numJobs)

	// Lê cada linha do arquivo contendo as operações dos jobs
	for i := 0; i < numJobs; i++ {
		if scanner.Scan() {
			line := scanner.Text()
			parts := strings.Fields(line)

			jobs[i] = make([][2]int, len(parts)/2)
			for j := 0; j < len(parts); j += 2 {
				machineID, _ := strconv.Atoi(parts[j])
				jobTimeInMachine, _ := strconv.Atoi(parts[j+1])
				jobs[i][j/2][0] = machineID
				jobs[i][j/2][1] = jobTimeInMachine
			}
		}
	}
	filenameParts := strings.Split(filename, "/")

	return &JobShopInstance{
		Name:        strings.Split(filename, "/")[len(filenameParts)-1],
		numJobs:     numJobs,
		numMachines: numMachines,
		jobs:        jobs, // Análogo à tabela 1s
	}, nil
}
