package models

type Simulation struct {
	Engine string
	Turbo  string
	Boost  float64
	RevMin int
	RevMax int
	Octane float64
}
