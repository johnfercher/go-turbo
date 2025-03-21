package models

import (
	"github.com/johnfercher/go-turbo/internal/math"
)

type Transmission struct {
	Name    string    `json:"name"`
	Final   float64   `json:"final"`
	Reverse float64   `json:"reverse"`
	Gears   []float64 `json:"gears"`
}

func (t *Transmission) MaxGear() int {
	return len(t.Gears)
}

func (t *Transmission) GetGearRatioFinal(gear int) float64 {
	return t.Gears[gear] * t.Final
}

func (t *Transmission) GetSpeed(rpm float64, wheelInch float64, tireHeight float64, gear int) float64 {
	wheelTireInch := wheelInch + math.MilimetersToInch(tireHeight)
	rpmWheelTireInch := wheelTireInch * rpm
	gearFinal := t.GetGearRatioFinal(gear)
	return (rpmWheelTireInch / gearFinal) * 0.00636934576
}
