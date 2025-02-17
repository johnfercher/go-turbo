package models

import (
	"fmt"
)

type Turbo struct {
	Name       string
	TurboScore [][]float64
}

func NewTurbo(name string, turboMatrix [][]float64) *Turbo {
	return &Turbo{
		Name:       name,
		TurboScore: turboMatrix,
	}
}

func (t *Turbo) String() string {
	return fmt.Sprintf("Turbo: %s\n", t.Name)
}

func (t *Turbo) Get(cfm float64, configuredBoost float64) float64 {
	return t.TurboScore[int(cfm)][int(configuredBoost)]
}
