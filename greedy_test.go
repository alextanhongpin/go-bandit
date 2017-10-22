package bandit

import (
	"log"
	"math/rand"
	"testing"
	"time"
)

func TestNewBandit(t *testing.T) {
	// t.Skip("test bandit")
	nArms := 3
	epsilon := float64(0.1)
	rand.Seed(time.Now().UnixNano())

	bandit := NewEpsilonGreedy(nArms, epsilon)
	wantN := nArms
	gotN := bandit.N
	if wantN != gotN {
		t.Errorf("want %v, got %v", wantN, gotN)
	}

	wantEpsilon := epsilon
	gotEpsilon := bandit.Epsilon
	if wantEpsilon != gotEpsilon {
		t.Errorf("want %v, got %v", wantEpsilon, gotEpsilon)
	}

	arm := bandit.SelectArm()
	wantArm := 0
	gotArm := arm
	if wantArm != gotArm {
		t.Errorf("want %v, got %v", wantArm, gotArm)
	}

	bandit.Update(arm, 1)
	wantValues := float64(1)
	gotValues := float64(bandit.Rewards[arm])
	if wantValues != gotValues {
		t.Errorf("want %v, got %v", wantValues, gotValues)
	}

	for i := 0; i < 1000; i++ {
		arm := bandit.SelectArm()
		reward := 0
		if rand.Float64() > 0.5 {
			reward = 1
		}
		bandit.Update(arm, float64(reward))
	}
	log.Println(bandit)
}

// func TestPull(t *testing.T) {
// 	t.Skip("skipping test pull")
// 	nArms := 5
// 	epsilon := float64(0.1)
// 	bandit := New(nArms, epsilon)
// 	exploitCount := 0
// 	for i := 0; i < 100000; i++ {
// 		arm, exploit := bandit.SelectArm()
// 		reward := 0
// 		if bernoulliArm() {
// 			reward = 1
// 		}
// 		if exploit == true {
// 			exploitCount++
// 		}
// 		bandit.Update(arm, float64(reward))
// 	}

// 	log.Printf("got %#v:", bandit)
// 	log.Println("exploit_count", exploitCount)
// 	if 1 != 0 {
// 		t.Errorf("want %v, got %v", 1, 0)
// 	}
