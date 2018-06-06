package bandit

import (
	"math"
	"sync"
)

// AnnealingSoftmax represents the annealing-softmax algorithm
type AnnealingSoftmax struct {
	sync.RWMutex
	Counts  []int
	Rewards []float64
}

// Init will initialise the counts and rewards with the provided number of arms
func (b *AnnealingSoftmax) Init(nArms int) error {
	if nArms < 1 {
		return ErrInvalidArms
	}
	b.Counts = make([]int, nArms)
	b.Rewards = make([]float64, nArms)
	return nil
}

// SelectArm chooses an arm that exploits if the value is more than the epsilon
// threshold, and explore if the value is less than epsilon
func (b *AnnealingSoftmax) SelectArm(probability float64) int {
	b.RLock()
	defer b.RUnlock()

	nArms := len(b.Rewards)
	t := sum(b.Counts...) + 1

	temperature := 1.0 / math.Log(float64(t)+1e-7)

	var z float64
	for i := 0; i < nArms; i++ {
		reward := b.Rewards[i]
		z += math.Exp(reward / temperature)
	}

	probs := make([]float64, nArms)
	for i := 0; i < nArms; i++ {
		reward := b.Rewards[i]
		probs[i] = math.Exp(reward/temperature) / z
	}
	return categoricalProb(probability, probs...)
}

// Update will update an arm with some reward value,
// e.g. click = 1, no click = 0
func (b *AnnealingSoftmax) Update(chosenArm int, reward float64) error {
	if chosenArm < 0 || chosenArm >= len(b.Rewards) {
		return ErrArmsIndexOutOfRange
	}
	if reward < 0 {
		return ErrInvalidReward
	}

	b.Counts[chosenArm]++
	n := float64(b.Counts[chosenArm])

	oldRewards := b.Rewards[chosenArm]
	newRewards := (oldRewards*(n-1) + reward) / n
	b.Rewards[chosenArm] = newRewards

	return nil
}

// GetCounts returns the counts
func (b *AnnealingSoftmax) GetCounts() []int {
	b.RLock()
	defer b.RUnlock()

	return b.Counts
}

// GetRewards returns the rewards
func (b *AnnealingSoftmax) GetRewards() []float64 {
	b.RLock()
	defer b.RUnlock()

	return b.Rewards
}

// NewAnnealingSoftmax returns a pointer to the AnnealingSoftmax struct
func NewAnnealingSoftmax(counts []int, rewards []float64) (*AnnealingSoftmax, error) {
	if len(counts) != len(rewards) {
		return nil, ErrInvalidLength
	}

	return &AnnealingSoftmax{
		Counts:  counts,
		Rewards: rewards,
	}, nil
}
