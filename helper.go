package bandit

import (
	"math"
)

func sum(values ...int) int {
	total := 0
	for _, v := range values {
		total += v
	}
	return total
}

func max(values ...float64) int {
	value := math.Inf(-1)
	index := 0
	for i, v := range values {
		if float64(v) > float64(value) {
			value = float64(v)
			index = i
		}
	}
	return index
}
