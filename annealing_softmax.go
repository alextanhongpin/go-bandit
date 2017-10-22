package bandit

import "math"

type AnnealingSoftmax struct {
	// Epsilon float64
	Counts  []int64
	Rewards []float64
	N       int
}

func (a *AnnealingSoftmax) Init() {
	a.Counts = make([]int64, a.N)
	a.Rewards = make([]float64, a.N)
}

func (a *AnnealingSoftmax) SelectArm() int {
	t := sumInt64(a.Counts...) + 1
	// var epsilon float64
	epsilon := 1.0 / math.Log(float64(t)+1e-7)

	exp := make([]float64, a.N)
	for i := 0; i < a.N; i++ {
		reward := a.Rewards[i]
		exp[i] = math.Exp(reward / epsilon)
	}
	z := sumFloat64(exp...)

	probs := make([]float64, a.N)
	for i := 0; i < a.N; i++ {
		reward := a.Rewards[i]
		probs[i] = math.Exp(reward/epsilon) / z
	}
	return categoricalProb(probs...)
}

func (a *AnnealingSoftmax) Update(chosenArm int, reward float64) {
	a.Counts[chosenArm] = a.Counts[chosenArm] + 1
	n := float64(a.Counts[chosenArm])
	value := a.Rewards[chosenArm]
	newValue := ((n-1)/n)*value + (1/n)*reward
	a.Rewards[chosenArm] = newValue
}

func NewAnnealingSoftmax(n int) *AnnealingSoftmax {
	softmax := AnnealingSoftmax{
		N: n,
	}
	softmax.Init()
	return &softmax
}
