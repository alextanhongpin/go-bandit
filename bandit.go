package bandit

// Bandit represents the bandit interface
type Bandit interface {
	Init(nArms int)
	SelectArm() int
	Update(chosenArm int, reward float64)
}
