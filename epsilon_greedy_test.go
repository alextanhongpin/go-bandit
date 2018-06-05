package bandit

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEpsilonGreedy_New(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		epsilon float64
		counts  []int
		rewards []float64
		err     error
	}{
		{0.1, nil, nil, nil},
		{0.1, []int{0, 0, 0}, []float64{0.0, 0.0, 0.0}, nil},
		{-0.1, nil, nil, ErrInvalidEpsilon},
		{1.1, nil, nil, ErrInvalidEpsilon},
		{1.1, []int{0, 0}, nil, ErrInvalidEpsilon},
		{1.0, []int{0, 0}, nil, ErrInvalidLength},
		{1.0, nil, []float64{0.0}, ErrInvalidLength},
		{1.0, []int{0, 0, 0, 0, 0}, []float64{0.0}, ErrInvalidLength},
	}

	for i, tt := range tests {
		_, err := NewEpsilonGreedy(tt.epsilon, tt.counts, tt.rewards)
		if tt.err != nil {
			assert.Equal(tt.err, err, "should throw the correct error for test %d", i+1)
		} else {
			assert.Nil(err)
		}
	}
}

func TestEpsilonGreedy_SelectArm(t *testing.T) {
	assert := assert.New(t)
	b, err := NewEpsilonGreedy(0.1, nil, nil)
	assert.Nil(err)
	b.Init(3)
	arm := b.SelectArm(0.1)
	assert.Equal(2, arm, "arm should be equal last item")
	b.Update(arm, 1.0)
	arm = b.SelectArm(1)
	assert.Equal(2, arm, "should select the best arm")
}
func TestEpsilonInit_WithValidParams(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		epsilon         float64
		counts          []int
		rewards         []float64
		arms            int
		expectedCounts  []int
		expectedRewards []float64
	}{
		{0.1, nil, nil, 3, []int{0, 0, 0}, []float64{0.0, 0.0, 0.0}},
		{0.1, []int{1, 2, 3}, []float64{1.0, 2.0, 3.0}, 3, []int{0, 0, 0}, []float64{0.0, 0.0, 0.0}},
	}

	for _, tt := range tests {
		b, err := NewEpsilonGreedy(tt.epsilon, tt.counts, tt.rewards)
		assert.Nil(err)

		err = b.Init(tt.arms)
		assert.Nil(err)

		assert.Equal(tt.arms, len(b.Counts), "should have length of %d", tt.arms)
		assert.Equal(tt.arms, len(b.Rewards), "should have length of %d", tt.arms)
		assert.Equal(tt.expectedCounts, b.Counts, "should be initialized to zero values")
		assert.Equal(tt.expectedRewards, b.Rewards, "should be initialized to zero values")
	}
}

func TestEpsilonInit_WithInvalidParams(t *testing.T) {
	assert := assert.New(t)
	b, err := NewEpsilonGreedy(0.1, nil, nil)
	assert.Nil(err)

	tests := []struct {
		arms int
		err  bool
	}{
		{100, false},
		{10, false},
		{1, false},
		{0, true},
		{-1, true},
		{-10, true},
		{-100, true},
	}
	for _, tt := range tests {
		err = b.Init(tt.arms)
		if tt.err {
			assert.Equal(ErrInvalidArms, err, "should throw error when arms length is invalid")
		} else {
			assert.Equal(tt.arms, len(b.Counts), "should have length of %d", tt.arms)
			assert.Equal(tt.arms, len(b.Rewards), "should have length of %d", tt.arms)
		}
	}
}

func TestEpsilonUpdate_Independent(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		arms            int
		chosenArm       int
		reward          float64
		expectedRewards float64
		expectedCounts  int
		err             bool
		errMsg          string
	}{
		{1, 1, 1.0, 0.0, 0, true, "arms index is out of range"},
		{1, -1, 1.0, 0.0, 0, true, "arms index is out of range"},
		{3, 1, 0.0, 0.0, 1, false, ""},
		{3, 1, 1.0, 1.0, 1, false, ""},
		{3, 1, -1.0, 1.0, 1, true, "reward must be greater than zero"},
		{3, 5, -1.0, 1.0, 1, true, "arms index is out of range"},
	}

	for _, tt := range tests {
		b, err := NewEpsilonGreedy(0.1, nil, nil)
		assert.Nil(err)

		err = b.Init(tt.arms)
		assert.Nil(err)

		err = b.Update(tt.chosenArm, tt.reward)
		if tt.err {
			assert.Error(err, tt.errMsg)
		} else {
			assert.Nil(err)
			assert.Equal(tt.expectedCounts, b.Counts[tt.chosenArm], "counts should be equal")
			assert.Equal(tt.expectedRewards, b.Rewards[tt.chosenArm], "counts should be equal")
		}
	}
}

func TestUpdate_Continuous(t *testing.T) {
	assert := assert.New(t)

	b, err := NewEpsilonGreedy(0.1, nil, nil)
	assert.Nil(err)

	err = b.Init(5)
	assert.Nil(err)

	tests := []struct {
		chosenArm       int
		reward          float64
		expectedRewards float64
		expectedCounts  int
	}{
		{1, 1.0, 1.0, 1},
		{1, 1.0, 1.0, 2},
		{1, 1.0, 1.0, 3},
		{1, 0.0, 0.75, 4},
		{4, 1.0, 1.0, 1},
		{4, 0.0, 0.5, 2},
	}

	for _, tt := range tests {
		err = b.Update(tt.chosenArm, tt.reward)
		assert.Nil(err)

		assert.Equal(tt.expectedCounts, b.Counts[tt.chosenArm], "counts should be equal")
		assert.Equal(tt.expectedRewards, b.Rewards[tt.chosenArm], "rewards should be equal")
	}

}
