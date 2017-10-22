package bandit

import (
	"log"
	"math/rand"
	"testing"
	"time"
)

func TestSoftmax(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	softmax := NewSoftmax(3, 0.1)

	// arm := softmax.SelectArm()
	// wantArm := 1
	// gotArm := arm
	// if wantArm != gotArm {
	// 	t.Errorf("want %v, got %v", wantArm, gotArm)
	// }
	// softmax.Update(arm, 1)

	for i := 0; i < 1000; i++ {
		arm := softmax.SelectArm()
		reward := 0.0
		if rand.Float64() > 0.5 {
			reward = 1.0
		}
		softmax.Update(arm, reward)
	}
	log.Println(softmax)
}
