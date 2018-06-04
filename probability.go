package bandit

import "math/rand"

// Prob represents the probability interface
type Prob interface {
	Random() float64
}

type prob struct{}

func (p *prob) Random() float64 {
	return rand.Float64()
}
