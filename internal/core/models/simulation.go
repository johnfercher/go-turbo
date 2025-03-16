package models

import "github.com/johnfercher/go-turbo/internal/core/models/fuel"

type Simulation struct {
	Engine string
	Turbo  string
	Boost  float64
	RevMin int
	RevMax int
	Fuel   *fuel.Fuel
}
