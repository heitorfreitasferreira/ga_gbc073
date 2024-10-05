package pso

import "math/rand"

type particle struct {
	position       []float64
	discr_position []int
	velocity       []float64
	posBest        []float64
	errBest        float64
	errI           float64
	bounds         [][]float64

	cost func([]float64) (float64, []int)
	rng  *rand.Rand
	i    int
}

func initParticle(p *particle, x0 []float64, v0 []float64, b [][]float64, costFunc func([]float64) (float64, []int), rng *rand.Rand) {
	p.position = x0
	p.velocity = v0
	p.posBest = x0
	p.errBest = -1
	p.errI = -1
	p.bounds = b
	p.cost = costFunc
	p.rng = rng
	p.i = 0
}

func (p *particle) updatePosition() {
	for i := 0; i < len(p.position); i++ {
		p.position[i] += p.velocity[i]
		if p.position[i] < float64(p.bounds[i][0]) {
			p.position[i] = float64(p.bounds[i][0])
		}
		if p.position[i] > float64(p.bounds[i][1]) {
			p.position[i] = float64(p.bounds[i][1])
		}
	}
	if p.errI < p.errBest || p.errBest == -1 {
		p.posBest = p.position
		p.errBest = p.errI
	}
	p.i++
}

func (p *particle) updateVelocity(pos_best_g []float64, w float64, c1 float64, c2 float64) {
	for i := 0; i < len(p.velocity); i++ {
		r1 := p.rng.Float64()
		r2 := p.rng.Float64()

		vel_cognitive := c1 * r1 * (p.posBest[i] - p.position[i])
		vel_social := c2 * r2 * (pos_best_g[i] - p.position[i])
		p.velocity[i] = w*p.velocity[i] + vel_cognitive + vel_social
	}
}

func (p *particle) evaluate() {
	p.errI, p.discr_position = p.cost(p.position)

	if p.errI < p.errBest || p.errBest == -1 {
		p.posBest = p.position
		p.errBest = p.errI
	}
}
