package main
/*
TODO: Update the implementation
import (
	"errors"
	"fmt"
	"log"
	"math"
	"math/rand"
	"sync"
)

var (
	ErrArmsIndexOutOfRange = errors.New("arms index is out of range")
	ErrInvalidReward       = errors.New("reward must be greater than zero")
	ErrNotImplemented      = errors.New("not implemented")
)

type Bandit interface {
	SelectArm(probability float64) uint
	Update(arm uint, reward float64) error
	Counts() []int
	Rewards() []float64
	Arms() uint
}

type bandit struct {
	sync.RWMutex
	counts  []int
	rewards []float64
	arms    uint
}

func newBandit(arms uint) bandit {
	return bandit{
		counts:  make([]int, arms),
		rewards: make([]float64, arms),
		arms:    arms,
	}
}

// SelectArm chooses an arm that exploits if the value is more than the epsilon
// threshold, and explore if the value is less than epsilon
func (b *bandit) SelectArm(probability float64) uint {
	log.Fatal(ErrNotImplemented)
	return 0
}

// Update will update an arm with some reward value,
// e.g. click = 1, no click = 0
func (b *bandit) Update(arm uint, reward float64) error {
	b.Lock()
	defer b.Unlock()

	if arm >= b.arms {
		return ErrArmsIndexOutOfRange
	}
	if reward < 0 {
		return ErrInvalidReward
	}
	b.counts[arm]++
	n := float64(b.counts[arm])
	oldRewards := b.rewards[arm]
	b.rewards[arm] = (oldRewards*(n-1) + reward) / n

	return nil
}

// Counts returns the counts
func (b *bandit) Counts() []int {
	b.RLock()
	cpy := make([]int, len(b.counts))
	copy(cpy, b.counts)
	b.RUnlock()

	return cpy
}

// Rewards returns the rewards
func (b *bandit) Rewards() []float64 {
	b.RLock()
	cpy := make([]float64, len(b.rewards))
	copy(cpy, b.rewards)
	b.RUnlock()

	return cpy
}

func (b *bandit) Arms() uint {
	return b.arms
}

type AnnealingSoftmax struct {
	bandit
}

func NewAnnealingSoftmax(arms uint) *AnnealingSoftmax {
	return &AnnealingSoftmax{
		bandit: newBandit(arms),
	}
}

func (b *AnnealingSoftmax) SelectArm(probability float64) uint {
	b.RLock()
	defer b.RUnlock()

	nArms := len(b.rewards)
	t := sum(b.counts...) + 1

	temperature := 1.0 / math.Log(float64(t)+1e-7)
	var z float64
	for i := 0; i < nArms; i++ {
		reward := b.rewards[i]
		z += math.Exp(reward / temperature)
	}

	probs := make([]float64, nArms)
	for i := 0; i < nArms; i++ {
		reward := b.rewards[i]
		probs[i] = math.Exp(reward/temperature) / z
	}
	return uint(categoricalProb(probability, probs...))
}

type EpsilonGreedy struct {
	epsilon float64
	bandit
}

func NewEpsilonGreedy(arms uint, epsilon float64) *EpsilonGreedy {
	return &EpsilonGreedy{
		epsilon: epsilon,
		bandit:  newBandit(arms),
	}
}

func (b *EpsilonGreedy) SelectArm(probability float64) uint {
	b.RLock()
	defer b.RUnlock()

	// Exploit
	if probability > b.epsilon {
		return uint(max(b.rewards...))
	}

	// Explore
	return uint(rand.Intn(len(b.rewards)))
}

type UCB struct {
	bandit
}

func NewUCB(arms uint) *UCB {
	return &UCB{
		bandit: newBandit(arms),
	}
}

func (b *UCB) SelectArm(probability float64) uint {
	b.RLock()
	defer b.RUnlock()

	nArms := len(b.counts)

	// Select unplayed arms
	for i := 0; i < nArms; i++ {
		if b.counts[i] == 0 {
			return uint(i)
		}
	}

	totalCounts := sum(b.counts...)
	ucbValues := make([]float64, nArms)

	for i := 0; i < nArms; i++ {
		count := b.counts[i]
		reward := b.rewards[i]
		bonus := math.Sqrt((2.0 * math.Log(float64(totalCounts))) / float64(count))
		ucbValues[i] = bonus + reward
	}
	return uint(max(ucbValues...))
}

func main() {
	bandits := []Bandit{
		NewAnnealingSoftmax(3),
		NewEpsilonGreedy(3, 0.1),
		NewUCB(3),
	}
	for _, bandit := range bandits {
		for i := 0; i < 1e3; i++ {
			arm := bandit.SelectArm(rand.Float64())
			reward := 1.0
			if rand.Float64() < 0.2 {
				reward = 0.0
			}
			bandit.Update(arm, reward)
		}
		counts, rewards := bandit.Counts(), bandit.Rewards()
		fmt.Println(counts, rewards)
	}
}

func sum(values ...int) int {
	var total int
	for _, v := range values {
		total += v
	}
	return total
}

func max(values ...float64) (index int) {
	value := float64(math.Inf(-1))
	for i, v := range values {
		if v > value {
			value = v
			index = i
		}
	}
	return
}

func categoricalProb(probability float64, probs ...float64) int {
	var cumulativeProb float64
	for i := 0; i < len(probs); i++ {
		cumulativeProb += probs[i]
		if cumulativeProb > probability {
			return i
		}
	}
	return len(probs) - 1
}
*/
