package bandit

import (
	"math"
)

func sum(values ...int) int {
	var total int
	for _, v := range values {
		total += v
	}
	return total
}

func sumInt64(values ...int64) int64 {
	var total int64
	for _, v := range values {
		total += v
	}
	return total
}

func sumFloat64(values ...float64) float64 {
	var total float64
	for _, v := range values {
		total += v
	}
	return total
}

func max(values ...float64) (index int) {
	value := math.Inf(-1)
	for i, v := range values {
		if float64(v) > float64(value) {
			value = float64(v)
			index = i
		}
	}
	return
}

func categoricalProb(probability float64, probs ...float64) int {
	var cumulativeProb float64
	for i := 0; i < len(probs); i++ {
		cumulativeProb += probs[i]
		if cumulativeProb > probability {
			return i
		}
	}
	return len(probs) - 1
}
