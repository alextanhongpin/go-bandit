package bandit

import (
	"errors"
	"math/rand"
	"sync"
)

var (
	ErrInvalidEpsilon      = errors.New("epsilon must be in range 0 to 1")
	ErrInvalidLength       = errors.New("counts and rewards must be of equal length")
	ErrInvalidArms         = errors.New("arms must be greater than zero")
	ErrArmsIndexOutOfRange = errors.New("arms index is out of range")
	ErrInvalidReward       = errors.New("reward must be greater than zero")
)

// EpsilonGreedy represents the bandit data
type EpsilonGreedy struct {
	sync.RWMutex
	Epsilon float64   `json:"epsilon"`
	Counts  []int     `json:"counts"`
	Rewards []float64 `json:"values"`
}

// Init will initialise the counts and rewards with the provided number of arms
func (b *EpsilonGreedy) Init(nArms int) error {
	if nArms < 1 {
		return ErrInvalidArms
	}
	b.Counts = make([]int, nArms)
	b.Rewards = make([]float64, nArms)
	return nil
}

// SelectArm chooses an arm that exploits if the value is more than the epsilon
// threshold, and explore if the value is less than epsilon
func (b *EpsilonGreedy) SelectArm(prob Prob) int {
	// Exploit
	if prob.Random() > b.Epsilon {
		return max(b.Rewards...)
	}
	// Explore
	return rand.Intn(len(b.Rewards))
}

// Update will update an arm with some reward value,
// e.g. click = 1, no click = 0
func (b *EpsilonGreedy) Update(chosenArm int, reward float64) error {
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

// NewEpsilonGreedy returns a pointer to the EpsilonGreedy struct
func NewEpsilonGreedy(epsilon float64, counts []int, rewards []float64) (*EpsilonGreedy, error) {
	if epsilon < 0 || epsilon > 1 {
		return nil, ErrInvalidEpsilon
	}
	if len(counts) != len(rewards) {
		return nil, ErrInvalidLength
	}

	return &EpsilonGreedy{
		Epsilon: epsilon,
		Rewards: rewards,
		Counts:  counts,
	}, nil
}
