package bandit

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSoftmax_New(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		epsilon float64
		counts  []int
		rewards []float64
		err     error
	}{
		{0.1, nil, nil, nil},
		{0.1, []int{0, 0, 0}, []float64{0.0, 0.0, 0.0}, nil},
		{-0.1, nil, nil, ErrInvalidTemperature},
		{1.1, nil, nil, nil},
		{1.1, []int{0, 0}, nil, ErrInvalidLength},
		{1.0, []int{0, 0}, nil, ErrInvalidLength},
		{1.0, nil, []float64{0.0}, ErrInvalidLength},
		{1.0, []int{0, 0, 0, 0, 0}, []float64{0.0}, ErrInvalidLength},
	}

	for i, tt := range tests {
		_, err := NewSoftmax(tt.epsilon, tt.counts, tt.rewards)
		if tt.err != nil {
			assert.Equal(tt.err, err, "should throw the correct error for test %d", i+1)
		} else {
			assert.Nil(err)
		}
	}
}

func TestSoftmax_Init(t *testing.T) {
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
		softmax, err := NewSoftmax(0.1, nil, nil)
		err = softmax.Init(tt.arms)

		if tt.err != nil {
			assert.Equal(err, tt.err, "should throw error for invalid arms length")
		} else {
			assert.Nil(err)
			assert.Equal(tt.arms, len(softmax.Counts), "counts should be of equal length with arm")
			assert.Equal(tt.arms, len(softmax.Rewards), "rewards should be of equal length with arm")
		}
	}
}

func TestSoftmax_SelectArm(t *testing.T) {
	assert := assert.New(t)
	b, err := NewSoftmax(0.1, nil, nil)
	assert.Nil(err)
	b.Init(3)

	arm := b.SelectArm(0.1)
	assert.Equal(arm, 0, "should select the unplayed arm")
	b.Update(arm, 1.0)
}

func TestSoftmax_UpdateArm(t *testing.T) {
	assert := assert.New(t)
	b, err := NewSoftmax(0.1, nil, nil)
	assert.Nil(err)

	err = b.Init(3)
	assert.Nil(err)

	tests := []struct {
		arm            int
		reward         float64
		expectedCounts int
		expectedReward float64
	}{
		{0, 1.0, 1, 1.0},
		{0, 0.0, 2, 0.5},
		{0, 1.0, 3, 2.0 / 3.0},
		{0, 1.0, 4, 0.75},
	}
	for i, tt := range tests {
		b.Update(tt.arm, tt.reward)
		assert.Equal(tt.expectedCounts, b.Counts[tt.arm], "counts should be equal for test %d", i+1)
		assert.Equal(tt.expectedReward, b.Rewards[tt.arm], "rewards should be equal for test %d", i+1)
	}
}

func TestSoftmax_UpdateArmWithInvalidParams(t *testing.T) {
	assert := assert.New(t)

	b, err := NewSoftmax(0.1, nil, nil)
	assert.Nil(err)

	tests := []struct {
		arms      int
		chosenArm int
		reward    float64
		err       error
	}{
		{1, 3, 0.0, ErrArmsIndexOutOfRange},
		{3, 3, 0.0, ErrArmsIndexOutOfRange},
		{3, 2, 1.0, nil},
		{3, 2, 0.0, nil},
		{3, 2, -1.0, ErrInvalidReward},
		{3, 3, -1.0, ErrArmsIndexOutOfRange},
	}
	for _, tt := range tests {
		err = b.Init(tt.arms)
		assert.Nil(err)

		err = b.Update(tt.chosenArm, tt.reward)
		if tt.err != nil {
			assert.Equal(tt.err, err, "should throw the correct error")
		} else {
			assert.Nil(err)
		}
	}
}
