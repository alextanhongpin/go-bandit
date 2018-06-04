package bandit

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewEpsilonGreedy_WithEpsilon(t *testing.T) {
	assert := assert.New(t)

	b, err := NewEpsilonGreedy(0.1, nil, nil)
	assert.Nil(err)
	assert.Nil(b.Counts)
	assert.Nil(b.Rewards)
	assert.Equal(0.1, b.Epsilon, "epsilon should be correct")
}

func TestNewEpsilonGreedy_WithAllParams(t *testing.T) {
	assert := assert.New(t)

	b, err := NewEpsilonGreedy(0.1, []int{0, 0, 0}, []float64{0.0, 0.0, 0.0})
	assert.Nil(err)
	assert.Equal(0.1, b.Epsilon, "epsilon should be correct")
	assert.Equal([]int{0, 0, 0}, b.Counts, "counts should be equal")
	assert.Equal([]float64{0.0, 0.0, 0.0}, b.Rewards, "rewards should be equal")
}

func TestNewEpsilonGreedy_WithIncorrectParams(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		epsilon float64
		counts  []int
		rewards []float64
		err     bool
		errMsg  string
	}{
		{-0.1, nil, nil, true, "epsilon must be in range 0 to 1"},
		{1.1, nil, nil, true, "epsilon must be in range 0 to 1"},
		{1.1, []int{0, 0}, nil, true, "counts and rewards must be of equal length"},
		{1.0, []int{0, 0}, nil, true, "counts and rewards must be of equal length"},
		{1.0, nil, []float64{0.0}, true, "counts and rewards must be of equal length"},
		{1.0, []int{0, 0, 0, 0, 0}, []float64{0.0}, true, "counts and rewards must be of equal length"},
	}

	for _, tt := range tests {
		b, err := NewEpsilonGreedy(tt.epsilon, tt.counts, tt.rewards)
		assert.Nil(b)
		if tt.err {
			assert.Error(err, tt.errMsg)
		}
	}
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
