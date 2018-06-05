package bandit

import (
	"log"
	"math/rand"
)

type BernoulliArm struct {
	p float64
}

func (b *BernoulliArm) Pull() bool {
	return rand.Float64() > b.p
}

func Simulate(b Bandit, pulls int, arms []BernoulliArm) (index, chosenArms []int, rewards, cumulativeRewards []float64) {

	index = make([]int, pulls)
	chosenArms = make([]int, pulls)
	rewards = make([]float64, pulls)
	cumulativeRewards = make([]float64, pulls)

	for i := 0; i < pulls; i++ {
		arm := b.SelectArm(rand.Float64())
		chosenArm := arms[arm]
		reward := 0.0
		if chosenArm.Pull() == true {
			reward = 1.0
		}
		b.Update(arm, reward)

		index[i] = i
		chosenArms[i] = arm
		rewards[i] = float64(reward)
		if i == 0 {
			cumulativeRewards[i] = float64(reward)
		} else {
			cumulativeRewards[i] = cumulativeRewards[i-1] + float64(reward)
		}
	}
	log.Println(b)
	return
}
