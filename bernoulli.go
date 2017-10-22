package bandit

import "math/rand"

type BernoulliArm struct {
	p float64
}

func (b *BernoulliArm) Pull() bool {
	return rand.Float64() > b.p
}
