package bandit

type BernoulliArm struct {
	p float64
}

func (b *BernoulliArm) Pull() bool {
	return b.p > 0.5
}
