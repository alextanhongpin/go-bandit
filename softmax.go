package bandit

import (
	"math"
	"sync"
)

// Softmax represents the softmax algorithm
type Softmax struct {
	sync.RWMutex
	Temperature float64
	Counts      []int
	Rewards     []float64
}

// Init will initialise the counts and rewards with the provided number of arms
func (b *Softmax) Init(nArms int) error {
	if nArms < 1 {
		return ErrInvalidArms
	}
	b.Counts = make([]int, nArms)
	b.Rewards = make([]float64, nArms)
	return nil
}

// SelectArm chooses an arm that exploits if the value is more than the epsilon
// threshold, and explore if the value is less than epsilon
func (b *Softmax) SelectArm(probability float64) int {
	b.RLock()
	defer b.RUnlock()

	nArms := len(b.Rewards)
	var z float64
	for i := 0; i < nArms; i++ {
		reward := b.Rewards[i]
		z += math.Exp(reward / b.Temperature)
	}

	probs := make([]float64, nArms)
	for i := 0; i < nArms; i++ {
		reward := b.Rewards[i]
		probs[i] = math.Exp(reward/b.Temperature) / z
	}
	return categoricalProb(probability, probs...)
}

// GetCounts returns the counts
func (b *Softmax) GetCounts() []int {
	return b.Counts
}

// GetRewards returns the rewards
func (b *Softmax) GetRewards() []float64 {
	return b.Rewards
}

// Update will update an arm with some reward value,
// e.g. click = 1, no click = 0
func (b *Softmax) Update(chosenArm int, reward float64) error {
	// It is safe to call len(ch) concurrently, but the obtained value is unavoidably racy per se.
	// In our case, the len is fixed, so we can call the lock/unlock later.
	// This is a micro-optimization, mainly because defer can be expensive.
	if chosenArm < 0 || chosenArm >= len(b.Rewards) {
		return ErrArmsIndexOutOfRange
	}
	if reward < 0 {
		return ErrInvalidReward
	}

	b.Lock()
	defer b.Unlock()

	b.Counts[chosenArm]++
	n := float64(b.Counts[chosenArm])

	oldRewards := b.Rewards[chosenArm]
	newRewards := (oldRewards*(n-1) + reward) / n
	b.Rewards[chosenArm] = newRewards

	return nil
}

// NewSoftmax returns a pointer to the Softmax struct
func NewSoftmax(temperature float64, counts []int, rewards []float64) (*Softmax, error) {
	if temperature < 0 {
		return nil, ErrInvalidTemperature
	}
	if len(counts) != len(rewards) {
		return nil, ErrInvalidLength
	}

	return &Softmax{
		Temperature: temperature,
		Counts:      counts,
		Rewards:     rewards,
	}, nil
}
