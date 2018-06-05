package bandit

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUCB_New(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		counts  []int
		rewards []float64
		err     error
	}{
		{nil, nil, nil},
		{make([]int, 0), make([]float64, 0), nil},
		{make([]int, 3), make([]float64, 3), nil},
		{nil, make([]float64, 3), ErrInvalidLength},
		{make([]int, 3), nil, ErrInvalidLength},
		{make([]int, 3), make([]float64, 5), ErrInvalidLength},
		{make([]int, 5), make([]float64, 3), ErrInvalidLength},
	}

	for _, tt := range tests {
		ucb, err := NewUCB(tt.counts, tt.rewards)
		if tt.err != nil {
			assert.Equal(err, tt.err, "should throw errors")
		} else {
			assert.Nil(err)
			assert.Equal(tt.counts, ucb.Counts, "counts should be equal")
			assert.Equal(tt.rewards, ucb.Rewards, "rewards should be equal")
		}
	}
}

func TestUCB_Init(t *testing.T) {
	assert := assert.New(t)
	tests := []struct {
		arms int
		err  error
	}{
		{-1, ErrInvalidArms},
		{0, ErrInvalidArms},
		{1, nil},
		{3, nil},
		{5, nil},
	}

	for _, tt := range tests {
		ucb, err := NewUCB(nil, nil)
		err = ucb.Init(tt.arms)

		if tt.err != nil {
			assert.Equal(err, tt.err, "should throw error for invalid arms length")
		} else {
			assert.Nil(err)
			assert.Equal(tt.arms, len(ucb.Counts), "counts should be of equal length with arm")
			assert.Equal(tt.arms, len(ucb.Rewards), "rewards should be of equal length with arm")
		}
	}
}

func TestUCB_SelectArm(t *testing.T) {
	assert := assert.New(t)
	b, err := NewUCB(nil, nil)
	assert.Nil(err)
	b.Init(3)

	arm := b.SelectArm(0.1)
	assert.Equal(arm, 0, "should select the unplayed arm")
	b.Update(arm, 1.0)

	arm = b.SelectArm(0.1)
	assert.Equal(arm, 1, "should select the next unplayed arm")
	b.Update(arm, 0.0)

	arm = b.SelectArm(0.1)
	assert.Equal(arm, 2, "should select the next unplayed arm")
	b.Update(arm, 1.0)

	arm = b.SelectArm(0.1)
	assert.Equal(arm, 0, "should select the correct arm")
}

func TestUCBUpdate_ValidParams(t *testing.T) {
	assert := assert.New(t)

	b, err := NewUCB(nil, nil)
	assert.Nil(err)

	tests := []struct {
		arms      int
		chosenArm int
		reward    float64
		err       error
	}{
		{1, 0, 0.0, nil},
		{1, -1, 0.0, ErrArmsIndexOutOfRange},
		{1, 1, 0.0, ErrArmsIndexOutOfRange},
		{1, 2, 0.0, ErrArmsIndexOutOfRange},
		{3, 1, 1.0, nil},
		{3, 1, -1.0, ErrInvalidReward},
		{3, 5, -1.0, ErrArmsIndexOutOfRange},
	}

	for _, tt := range tests {
		b.Init(tt.arms)
		err := b.Update(tt.chosenArm, tt.reward)
		if tt.err != nil {
			assert.Equal(tt.err, err, "should throw error for invalid params")
		} else {
			assert.Nil(err)
		}
	}
}

// func TestUCB(t *testing.T) {
// 	// rand.Seed(time.Now().UnixNano())
// 	ucb := NewUCB(3)
// 	for i := 0; i < 100000; i++ {
// 		arm := ucb.SelectArm()
// 		reward := 0.0
// 		if rand.Float64() > 0.5 {
// 			reward = 1.0
// 		}
// 		ucb.Update(arm, reward)
// 	}
// 	log.Println(ucb)
// }

// func TestSimulate3(t *testing.T) {
// 	nArms := 3
// 	means := []float64{0.1, 0.8, 0.1}
// 	bandit := NewUCB(nArms)

// 	bernoullis := make([]BernoulliArm, nArms)
// 	for i := 0; i < nArms; i++ {
// 		bernoullis[i] = BernoulliArm{p: means[i]}
// 	}
// 	index, chosenArms, rewards, cumulativeRewards := Simulate(bandit, 1000, bernoullis)

// 	file, _ := os.Create("ucb1.csv")
// 	defer file.Close()
// 	writer := csv.NewWriter(file)
// 	defer writer.Flush()

// 	writer.Write([]string{"index", "chosen_arm", "reward", "cumulative_rewards"})
// 	for _, value := range index {
// 		err := writer.Write([]string{
// 			fmt.Sprint(value),
// 			fmt.Sprint(chosenArms[value]),
// 			fmt.Sprint(rewards[value]),
// 			fmt.Sprint(cumulativeRewards[value]),
// 		})
// 		if err != nil {
// 			log.Println(err)
// 		}
// 	}
// }
