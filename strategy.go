package bandit

import (
	"log"
)

// Strategy represents the different algorithm that can be used to implement bandit algorithm
type Strategy interface {
	SelectArm() (arm int)
	Update(arm int, reward float64)
}

func Simulate(s Strategy, pulls int, arms []BernoulliArm) (index, chosenArms []int, rewards, cumulativeRewards []float64) {

	index = make([]int, pulls)
	chosenArms = make([]int, pulls)
	rewards = make([]float64, pulls)
	cumulativeRewards = make([]float64, pulls)

	for i := 0; i < pulls; i++ {
		arm := s.SelectArm()
		chosenArm := arms[arm]
		reward := 0.0
		if chosenArm.Pull() == true {
			reward = 1.0
		}
		s.Update(arm, reward)

		index[i] = i
		chosenArms[i] = arm
		rewards[i] = float64(reward)
		if i == 0 {
			cumulativeRewards[i] = float64(reward)
		} else {
			cumulativeRewards[i] = cumulativeRewards[i-1] + float64(reward)
		}
	}
	log.Println(s)
	return
}
