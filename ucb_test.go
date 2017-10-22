package bandit

import (
	"log"
	"math/rand"
	"testing"
	"time"
)

func TestUCB(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	ucb := NewUCB(3)
	for i := 0; i < 100000; i++ {
		arm := ucb.SelectArm()
		reward := 0.0
		if rand.Float64() > 0.5 {
			reward = 1.0
		}
		ucb.Update(arm, reward)
	}
	log.Println(ucb)
}
