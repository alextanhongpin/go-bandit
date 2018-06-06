package bandit

import "errors"

var (
	ErrInvalidEpsilon      = errors.New("epsilon must be in range 0 to 1")
	ErrInvalidTemperature  = errors.New("temperature must be greater than zero")
	ErrInvalidLength       = errors.New("counts and rewards must be of equal length")
	ErrInvalidArms         = errors.New("arms must be greater than zero")
	ErrArmsIndexOutOfRange = errors.New("arms index is out of range")
	ErrInvalidReward       = errors.New("reward must be greater than zero")
)

// Bandit represents the bandit interface
type Bandit interface {
	Init(nArms int) error
	SelectArm(probability float64) int
	Update(chosenArm int, reward float64) error
	GetCounts() []int
	GetRewards() []float64
}
