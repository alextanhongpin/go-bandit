package bandit

import "math"

type Softmax struct {
	Epsilon float64
	Counts  []int64
	Rewards []float64
	N       int
}

func (s *Softmax) Initialize() {
	s.Counts = make([]int64, s.N)
	s.Rewards = make([]float64, s.N)
}

func (s *Softmax) SelectArm() int {
	exp := make([]float64, s.N)
	for i := 0; i < s.N; i++ {
		reward := s.Rewards[i]
		exp[i] = math.Exp(reward / s.Epsilon)
	}
	z := sumFloat64(exp...)

	probs := make([]float64, s.N)
	for i := 0; i < s.N; i++ {
		reward := s.Rewards[i]
		probs[i] = math.Exp(reward/s.Epsilon) / z
	}
	return categoricalProb(probs...)
}

func (s *Softmax) Update(chosenArm int, reward float64) {
	s.Counts[chosenArm] = s.Counts[chosenArm] + 1
	n := float64(s.Counts[chosenArm])
	value := s.Rewards[chosenArm]
	newValue := (((n - 1) / n) * value) + ((1 / n) * reward)
	s.Rewards[chosenArm] = newValue
}

func NewSoftmax(n int, epsilon float64) *Softmax {
	softmax := Softmax{
		N:       n,
		Epsilon: epsilon,
	}
	softmax.Initialize()
	return &softmax
}
