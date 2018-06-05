package bandit

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSum(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		params   []int
		expected int
	}{
		{[]int{1, 2, 3, 4, 5}, 15},
		{[]int{}, 0},
		{[]int{-1, 1}, 0},
		{[]int{10, 30}, 40},
	}

	for _, tt := range tests {
		assert.Equal(tt.expected, sum(tt.params...), "should return the correct sum")
	}
}

func TestSumInt64(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		params   []int64
		expected int64
	}{
		{[]int64{1, 2, 3, 4, 5}, 15},
		{[]int64{}, 0},
		{[]int64{-1, 1}, 0},
		{[]int64{10, 30}, 40},
	}

	for _, tt := range tests {
		assert.Equal(tt.expected, sumInt64(tt.params...), "should return the correct sum")
	}
}

func TestSumFloat64(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		params   []float64
		expected float64
	}{
		{[]float64{1.1, 2.1, 3.1, 4.1, 5.1}, 15.5},
		{[]float64{}, 0},
		{[]float64{-1, 1}, 0},
		{[]float64{10.5, 30.5}, 41.0},
	}

	for _, tt := range tests {
		assert.Equal(tt.expected, sumFloat64(tt.params...), "should return the correct sum")
	}
}

func TestMax(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		params   []float64
		expected int
	}{
		{[]float64{1.1, 2.1, 3.1, 4.1, 5.1}, 4},
		{[]float64{}, 0},
		{[]float64{-1, 1}, 1},
		{[]float64{10.5, 30.5}, 1},
	}

	for _, tt := range tests {
		assert.Equal(tt.expected, max(tt.params...), "should return the max index")
	}
}

func TestCategoricalProb(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		params      []float64
		probability float64
		expected    int
	}{
		{[]float64{1.1, 2.1, 3.1, 4.1, 5.1}, 0.0, 0},
		{[]float64{1.1, 2.1, 3.1, 4.1, 5.1}, 1.0, 0},
		{[]float64{1.1, 2.1, 3.1, 4.1, 5.1}, 2.0, 1},
		{[]float64{}, 1.0, -1},
		{[]float64{-1, 1}, 1.0, 1},
		{[]float64{1.0, 2.0, 3.0, 4.0, 5.0}, 6.0, 3},
	}

	for _, tt := range tests {
		assert.Equal(tt.expected, categoricalProb(tt.probability, tt.params...), "should return the max index")
	}
}
