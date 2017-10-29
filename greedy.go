package bandit

import (
	"math/rand"
	"sync"
)

// EpsilonGreedy represents the bandit data
type EpsilonGreedy struct {
	sync.Mutex
	Epsilon float64   `json:"epsilon"`
	Counts  []int64   `json:"counts"`
	Rewards []float64 `json:"values"`
	N       int       `json:"n"`
}

// SelectArm chooses an arm that exploits if the value is more than the epsilon
// threshold, and explore if the value is less than epsilon
func (b *EpsilonGreedy) SelectArm() int {
	// Exploit
	b.Lock()
	rewards := b.Rewards
	b.Unlock()
	if rand.Float64() > b.Epsilon {
		return max(rewards...)
	}
	// Explore
	return rand.Intn(len(rewards))
}

// Update will update an arm with some reward value,
// e.g. click = 1, no click = 0
func (b *EpsilonGreedy) Update(chosenArm int, reward float64) {
	b.Lock()
	count := b.Counts[chosenArm]
	b.Unlock()

	newCount := count + 1

	b.Lock()
	b.Counts[chosenArm] = newCount
	b.Unlock()

	n := float64(newCount)

	b.Lock()
	v := float64(b.Rewards[chosenArm])
	b.Unlock()

	newValue := (v*(n-1) + reward) / n

	b.Lock()
	b.Rewards[chosenArm] = newValue
	b.Unlock()
}

// SetRewards sets the values to the input specified
func (b *EpsilonGreedy) SetRewards(rewards []float64) {
	b.Lock()
	b.Rewards = rewards
	b.Unlock()
}

// SetCounts sets the counts to the input specified
func (b *EpsilonGreedy) SetCounts(counts []int64) {
	b.Lock()
	b.Counts = counts
	b.Unlock()
}

// NewEpsilonGreedy returns a pointer to the EpsilonGreedy struct
func NewEpsilonGreedy(nArms int, epsilonDecay float64) *EpsilonGreedy {
	return &EpsilonGreedy{
		N:       nArms,
		Epsilon: epsilonDecay,
		Rewards: make([]float64, nArms),
		Counts:  make([]int64, nArms),
	}
}
