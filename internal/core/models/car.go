package models

import "github.com/johnfercher/go-turbo/internal/core/models/fuel"

type Car struct {
	Engine       string
	Turbo        string
	Transmission string
	WheelInches  float64
	TireHeightMM float64
	Boost        float64
	RevMin       float64
	RevMax       float64
	Fuel         *fuel.Fuel
}
