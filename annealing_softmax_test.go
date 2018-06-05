package bandit

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAnnealingSoftmax_New(t *testing.T) {
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
		softmax, err := NewAnnealingSoftmax(tt.counts, tt.rewards)
		if tt.err != nil {
			assert.Equal(err, tt.err, "should throw errors")
		} else {
			assert.Nil(err)
			assert.Equal(tt.counts, softmax.Counts, "counts should be equal")
			assert.Equal(tt.rewards, softmax.Rewards, "rewards should be equal")
		}
	}
}

func TestAnnealingSoftmax_Init(t *testing.T) {
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
		softmax, err := NewAnnealingSoftmax(nil, nil)
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

func TestAnnealingSoftmax_SelectArm(t *testing.T) {
	assert := assert.New(t)
	b, err := NewAnnealingSoftmax(nil, nil)
	assert.Nil(err)
	b.Init(3)

	arm := b.SelectArm(0.1)
	assert.Equal(arm, 0, "should select the unplayed arm")
	err = b.Update(arm, 1.0)
	assert.Nil(err)

}

func TestAnnealingSoftmax_UpdateWithValidParams(t *testing.T) {
	assert := assert.New(t)

	b, err := NewAnnealingSoftmax(nil, nil)
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
