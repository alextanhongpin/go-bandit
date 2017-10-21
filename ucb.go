package bandit

import (
	"math"
)

type UCB struct {
	Epsilon float64
	Counts  []int64
	Rewards []float64
	N       int
}

func (u *UCB) SelectArm() int {
	nArms := u.N
	for i := 0; i < nArms; i++ {
		if u.Counts[i] == 0 {
			return i
		}
	}

	totalCounts := sumInt64(u.Counts...)

	ucbValues := make([]float64, nArms)
	for i := 0; i < nArms; i++ {
		count := u.Counts[i]
		reward := u.Rewards[i]
		bonus := math.Sqrt((2.0 * math.Log(float64(totalCounts))) / float64(count))
		ucbValues[i] = bonus + reward
	}
	return max(ucbValues...)
}

func (u *UCB) Update(chosenArm int, reward float64) {
	u.Counts[chosenArm] = u.Counts[chosenArm] + 1
	n := float64(u.Counts[chosenArm])
	value := u.Rewards[chosenArm]
	newValue := ((n-1)/n)*value + (1/n)*reward
	u.Rewards[chosenArm] = newValue
}
