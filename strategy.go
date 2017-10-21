package bandit

// Strategy represents the different algorithm that can be used to implement bandit algorithm
type Strategy interface {
	SelectArm() (arm int, exploit bool)
	Update(arm int, reward float64)
}
