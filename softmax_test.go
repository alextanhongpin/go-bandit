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

	arm := softmax.SelectArm()
	wantArm := 1
	gotArm := arm
	if wantArm != gotArm {
		t.Errorf("want %v, got %v", wantArm, gotArm)
	}
	softmax.Update(arm, 1)

	for i := 0; i < 10000; i++ {
		arm := softmax.SelectArm()
		s1 := rand.NewSource(time.Now().UnixNano())
		r1 := rand.New(s1)
		reward := 0.0
		if r1.Float64() > 0.5 {
			reward = 1.0
		}
		softmax.Update(arm, reward)
	}
	log.Println(softmax)
}
