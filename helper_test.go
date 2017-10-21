package bandit

import "testing"

func TestSum(t *testing.T) {
	testTable := []struct {
		in  []int
		out int
	}{
		{[]int{1, 2, 3, 4, 5}, 15},
		{[]int{-2, -1, 0, 1, 2}, 0},
		{[]int{10, 1, 20, 2, 30}, 63},
	}

	for _, v := range testTable {
		want := v.out
		got := sum(v.in...)
		if got != want {
			t.Errorf("want %v, got %v", want, got)
		}
	}
}

func TestMax(t *testing.T) {
	testTable := []struct {
		in  []float64
		out int
	}{
		{[]float64{-100, -50, 0, 50, 100}, 4},
		{[]float64{1, 0}, 0},
		{[]float64{1000, 1000}, 0},
		{[]float64{-100, -1000}, 0},
	}

	for _, v := range testTable {
		want := v.out
		got := max(v.in...)

		if got != want {
			t.Errorf("want %v, got %v", want, got)
		}
	}
}
