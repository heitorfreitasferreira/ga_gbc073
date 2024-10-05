package pso

import (
	"fmt"
	"math/rand"
	"os"
)

func Pso(
	cost func([]float64) (float64, []int), // "fitness", retorna também a posição discreta caso seja otimização discreta, se não for pode retornar tudo 0
	bounds [][]float64, // limites para cada dimensão (exemplo: [[-10, 10], [-10, 10], [-10, 10]] para 3 dimensões)
	n_particles int,
	particle_dimention int, // "tamanho do genoma"
	w []float64, // inércia da iteração, um valor para cada iteração caso queira variar, se não quiser passar um vetor com um único valor
	c_cogni []float64, // componente cognitivo, um valor para cada iteração
	c_social []float64, // componente social, um valor para cada iteração
	max_iter int,
	initialPosFactory func() []float64, // função que retorna a posição inicial de um particle
	initialVelFactory func() []float64, // função que retorna a velocidade inicial de um particle
	csvFile *os.File, // arquivo para salvar os resultados
	randSource *rand.Rand, // fonte de números aleatórios
) ([]float64, float64, []int) { // retorna a melhor posição, o erro da melhor posição e a melhor posição discreta caso seja otimização discreta
	particles := make([]*particle, n_particles)
	pos_best_g := make([]float64, particle_dimention)
	best_g_discrete := make([]int, particle_dimention)
	type stats struct {
		bestError       []float64
		bestPos         [][]float64
		bestPosDiscrete [][]int
	}
	s := stats{
		bestError:       make([]float64, max_iter),
		bestPos:         make([][]float64, max_iter),
		bestPosDiscrete: make([][]int, max_iter),
	}

	// Inicializa os particles
	for i := 0; i < n_particles; i++ {
		particles[i] = &particle{}
		initParticle(particles[i], initialPosFactory(), initialVelFactory(), bounds, cost, randSource)
	}

	err_best_g := -1.0
	fmt.Fprintf(csvFile, "Iteration,Best Error\n")
	for i := 0; i < max_iter; i++ {
		for j := 0; j < n_particles; j++ {
			particles[j].evaluate()

			if particles[j].errI < err_best_g || err_best_g == -1 {
				copy(pos_best_g, particles[j].position)
				copy(best_g_discrete, particles[j].discr_position)
				err_best_g = particles[j].errI
			}
		}

		for j := 0; j < n_particles; j++ {
			particles[j].updateVelocity(pos_best_g, w[i], c_cogni[i], c_social[i])
			particles[j].updatePosition()
		}
		// Atualizando o histórico
		s.bestError[i] = err_best_g
		s.bestPos[i] = make([]float64, len(pos_best_g))
		copy(s.bestPos[i], pos_best_g)
		s.bestPosDiscrete[i] = make([]int, len(best_g_discrete))
		copy(s.bestPosDiscrete[i], best_g_discrete)
	}

	headers, data := encodeToCsv(s.bestPos, s.bestError, s.bestPosDiscrete, ",")
	statsToCsv(headers, data, ";", csvFile)
	return pos_best_g, err_best_g, best_g_discrete
}
