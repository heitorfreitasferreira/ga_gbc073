package ga

import (
	"fmt"
	"math/rand"
	"time"
)

// Cromossome representa um indivíduo na abordagem de representação indireta
type Cromossome struct {
	genome []int
}

// GenerateCromossome cria um indivíduo com a abordagem de representação indireta
func GenerateCromossome(instance *JobShopInstance) Cromossome {
	numGenes := instance.numJobs * instance.numMachines
	genome := make([]int, numGenes)

	// Gera um genome com a lista de jobs para cada operação de máquina
	idx := 0
	for j := 0; j < instance.numJobs; j++ {
		for m := 0; m < instance.numMachines; m++ {
			genome[idx] = j // Jobs são numerados a partir de 0
			idx++
		}
	}

	shuffle(genome)

	// Embaralhar os genes para criar um indivíduo aleatório (aqui pode usar alguma função de shuffle se necessário)
	// shuffle(genome) // Se precisar de embaralhamento

	return Cromossome{genome: genome}
}

// Função para realizar o crossover entre dois indivíduos
func Crossover(p1, p2 Cromossome) (Cromossome, Cromossome) {
	source := rand.New(rand.NewSource(time.Now().UnixNano()))

	// Escolher aleatoriamente o índice de início e término para o trecho a ser trocado
	start1 := source.Intn(len(p1.genome))
	end1 := source.Intn(len(p1.genome)-start1) + start1

	start2 := source.Intn(len(p2.genome))
	end2 := source.Intn(len(p2.genome)-start2) + start2

	// Extrair o pedaço de schedule (subsequência) dos pais
	partial1 := p1.genome[start1 : end1+1]
	partial2 := p2.genome[start2 : end2+1]

	// Gerar os filhos trocando as partes dos pais
	o1 := make([]int, len(p1.genome))
	o2 := make([]int, len(p2.genome))
	copy(o1, p1.genome)
	copy(o2, p2.genome)

	// Inserir as partes trocadas nos filhos
	for i := start1; i <= end1; i++ {
		o1[i] = partial2[i-start1]
	}
	for i := start2; i <= end2; i++ {
		o2[i] = partial1[i-start2]
	}

	// Ajustar os filhos para remover excessos e adicionar genes ausentes
	o1 = fixGenes(o1, p2.genome)
	o2 = fixGenes(o2, p1.genome)

	return Cromossome{genome: o1}, Cromossome{genome: o2}
}

// Função auxiliar para ajustar os genes após o crossover
func fixGenes(offspring []int, reference []int) []int {
	// Contar a ocorrência de cada gene no offspring
	geneCount := make(map[int]int)
	for _, gene := range offspring {
		geneCount[gene]++
	}

	// Remover genes em excesso (se houver repetições)
	for i := range offspring {
		if geneCount[offspring[i]] > countValues(offspring[i], reference) {
			geneCount[offspring[i]]--
			offspring = append(offspring[:i], offspring[i+1:]...)
			i-- // Reajustar o índice para reavaliar a posição após a remoção
		}
	}

	// Adicionar genes que estejam faltando
	for _, gene := range reference {
		if geneCount[gene] < countValues(gene, reference) {
			offspring = append(offspring, gene)
			geneCount[gene]++
		}
	}

	return offspring
}

// Função auxiliar para contar a ocorrência de um gene em um genoma
func countValues(gene int, genome []int) int {
	count := 0
	for _, g := range genome {
		if g == gene {
			count++
		}
	}
	return count
}

// Print exibe o indivíduo gerado
func (ind *Cromossome) Print() {
	fmt.Println("Indivíduo:", ind.genome)
}
