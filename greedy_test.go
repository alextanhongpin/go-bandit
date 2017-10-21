package bandit

import (
	"testing"
)

func TestNewBandit(t *testing.T) {
	t.Skip("test bandit")
	// nArms := 10
	// epsilon := float64(0.1)
	// // rand.Seed(time.Now().UnixNano())

	// bandit := New(nArms, epsilon)
	// wantN := nArms
	// gotN := bandit.n
	// if wantN != gotN {
	// 	t.Errorf("want %v, got %v", wantN, gotN)
	// }

	// wantEpsilon := epsilon
	// gotEpsilon := bandit.epsilon
	// 	if wantEpsilon != gotEpsilon {
	// 		t.Errorf("want %v, got %v", wantEpsilon, gotEpsilon)
	// 	}

	// 	arm, _ := bandit.SelectArm()
	// 	wantArm := 0
	// 	gotArm := arm
	// 	if wantArm != gotArm {
	// 		t.Errorf("want %v, got %v", wantArm, gotArm)
	// 	}

	// 	bandit.Update(arm, 1)
	// 	wantValues := float64(1)
	// 	gotValues := float64(bandit.values[arm])
	// 	if wantValues != gotValues {
	// 		t.Errorf("want %v, got %v", wantValues, gotValues)
	// 	}
	// }

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
}
