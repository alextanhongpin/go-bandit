package bandit

import (
	"math/rand"
	"sync"
)

// EpsilonGreedy represents the bandit data
type EpsilonGreedy struct {
	sync.RWMutex
	Epsilon float64         `json:"epsilon"`
	Counts  map[int]int64   `json:"counts"`
	Rewards map[int]float64 `json:"values"`
	N       int             `json:"n"`
}

// SelectArm chooses an arm that exploits if the value is more than the epsilon
// threshold, and explore if the value is less than epsilon
func (b *EpsilonGreedy) SelectArm() int {
	// Exploit
	b.RLock()
	rewards := b.Rewards
	b.RUnlock()
	if rand.Float64() > b.Epsilon {
		return max(rewards...)
	}
	// Explore
	return rand.Intn(len(rewards))
}

// Update will update an arm with some reward value,
// e.g. click = 1, no click = 0
func (b *EpsilonGreedy) Update(chosenArm int, reward float64) {
	b.RLock()
	count := b.Counts[chosenArm]
	b.RUnlock()

	newCount := count + 1

	// The data race exists because slices are reference types in Go. They are generally passed by value, but being reference types, any changes made to the one value is reflected in another.
	b.Lock()
	b.Counts[chosenArm] = newCount
	b.Unlock()

	n := float64(newCount)

	b.RLock()
	v := b.Rewards[chosenArm]
	b.RUnlock()

	newValue := (float64(v)*(n-1) + reward) / n

	b.Lock()
	b.Rewards[chosenArm] = newValue
	b.Unlock()
}

// SetRewards sets the values to the input specified
// func (b *EpsilonGreedy) SetRewards(rewards []float64) {
func (b *EpsilonGreedy) SetRewards(rewards map[int]float64) {
	b.Lock()
	b.Rewards = rewards
	b.Unlock()
}

// SetCounts sets the counts to the input specified
// func (b *EpsilonGreedy) SetCounts(counts []int64) {
func (b *EpsilonGreedy) SetCounts(counts map[int]int64) {
	b.Lock()
	b.Counts = counts
	b.Unlock()
}

// NewEpsilonGreedy returns a pointer to the EpsilonGreedy struct
func NewEpsilonGreedy(nArms int, epsilonDecay float64) *EpsilonGreedy {
	return &EpsilonGreedy{
		N:       nArms,
		Epsilon: epsilonDecay,
		// Rewards: make([]float64, nArms),
		// Counts:  make([]int64, nArms),
		Rewards: make(map[int]float64, nArms),
		Counts:  make(map[int]int64, nArms),
	}
}
