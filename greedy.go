package bandit

import (
	"fmt"
	"math/rand"
	"sync"

	"github.com/orcaman/concurrent-map"
)

// EpsilonGreedy represents the bandit data
type EpsilonGreedy struct {
	sync.RWMutex
	Epsilon float64            `json:"epsilon"`
	Counts  cmap.ConcurrentMap `json:"counts"`
	Rewards cmap.ConcurrentMap `json:"values"`
	N       int                `json:"n"`
}

// SelectArm chooses an arm that exploits if the value is more than the epsilon
// threshold, and explore if the value is less than epsilon
func (b *EpsilonGreedy) SelectArm() int {
	// Exploit
	rewards := make([]float64, b.N)
	for i := 0; i < b.N; i++ {
		if tmp, ok := b.Rewards.Get(fmt.Sprint(i)); ok {
			rewards[i] = tmp.(float64)
		}
	}
	if rand.Float64() > b.Epsilon {
		return max(rewards...)
	}
	// Explore
	return rand.Intn(len(rewards))
}

// Update will update an arm with some reward value,
// e.g. click = 1, no click = 0
func (b *EpsilonGreedy) Update(chosenArm int, reward float64) {
	var count int64
	if tmp, ok := b.Counts.Get(fmt.Sprint(chosenArm)); ok {
		count = tmp.(int64)
	}

	newCount := count + 1

	// The data race exists because slices are reference types in Go. They are generally passed by value, but being reference types, any changes made to the one value is reflected in another.
	b.Lock()
	b.Counts.Set(fmt.Sprint(chosenArm), newCount)
	b.Unlock()

	n := float64(newCount)

	var v float64
	if tmp, ok := b.Rewards.Get(fmt.Sprint(chosenArm)); ok {
		v = tmp.(float64)
	}

	newValue := (float64(v)*(n-1) + reward) / n

	// b.Lock()
	// b.Rewards[chosenArm] = newValue
	// b.Unlock()
	b.Rewards.Set(fmt.Sprint(chosenArm), newValue)
}

// SetRewards sets the values to the input specified
// func (b *EpsilonGreedy) SetRewards(rewards []float64) {
// func (b *EpsilonGreedy) SetRewards(rewards map[int]float64) {
// 	b.Lock()
// 	b.Rewards = rewards
// 	b.Unlock()
// }

// SetCounts sets the counts to the input specified
// func (b *EpsilonGreedy) SetCounts(counts []int64) {
// func (b *EpsilonGreedy) SetCounts(counts map[int]int64) {
// 	b.Lock()
// 	b.Counts = counts
// 	b.Unlock()
// }

// NewEpsilonGreedy returns a pointer to the EpsilonGreedy struct
func NewEpsilonGreedy(nArms int, epsilonDecay float64) *EpsilonGreedy {

	rewards := cmap.New()
	counts := cmap.New()

	for i := 0; i < nArms; n++ {
		rewards.Set(fmt.Sprint(i), 0)
		counts.Set(fmt.Sprint(i), 0)
	}
	return &EpsilonGreedy{
		N:       nArms,
		Epsilon: epsilonDecay,
		// Rewards: make([]float64, nArms),
		// Counts:  make([]int64, nArms),
		Rewards: rewards,
		Counts:  counts,
	}
}
