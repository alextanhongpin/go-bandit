package bandit

import (
	"encoding/csv"
	"fmt"
	"log"
	"math/rand"
	"os"
	"testing"
	"time"
)

func TestAnnealingBandit(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	softmax := NewAnnealingSoftmax(3)
	for i := 0; i < 100000; i++ {
		arm := softmax.SelectArm()
		reward := 0.0
		if rand.Float64() > 0.5 {
			reward = 1.0
		}
		softmax.Update(arm, reward)
	}
	log.Println(softmax)
}

func TestSimulate2(t *testing.T) {
	nArms := 3
	means := []float64{0.1, 0.8, 0.1}
	bandit := NewAnnealingSoftmax(nArms)

	bernoullis := make([]BernoulliArm, nArms)
	for i := 0; i < nArms; i++ {
		bernoullis[i] = BernoulliArm{p: means[i]}
	}
	index, chosenArms, rewards, cumulativeRewards := Simulate(bandit, 1000, bernoullis)

	file, _ := os.Create("annealing_softmax.csv")
	defer file.Close()
	writer := csv.NewWriter(file)
	defer writer.Flush()

	writer.Write([]string{"index", "chosen_arm", "reward", "cumulative_rewards"})
	for _, value := range index {
		err := writer.Write([]string{
			fmt.Sprint(value),
			fmt.Sprint(chosenArms[value]),
			fmt.Sprint(rewards[value]),
			fmt.Sprint(cumulativeRewards[value]),
		})
		if err != nil {
			log.Println(err)
		}
	}
}
