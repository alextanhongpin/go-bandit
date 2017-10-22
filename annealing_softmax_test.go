package bandit

import (
	"log"
	"math/rand"
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
