package models

import (
	"errors"
	"fmt"
)

type Range struct {
	Min    float64
	Max    float64
	Boost  float64
	Health float64
}

func (r Range) GreaterThan(a Range) bool {
	return r.Min > a.Min
}

func NewRange(min, max, boost, health float64) *Range {
	return &Range{
		Min:    min,
		Max:    max,
		Boost:  boost,
		Health: health,
	}
}

func (r *Range) String() string {
	return fmt.Sprintf("[%.0f-%.0f]%.0f~|%.0f♥", r.Min, r.Max, r.Boost, r.Health)
}

type Turbo struct {
	Name   string
	Slices map[string][]*Range
}

func (t *Turbo) String() string {
	s := fmt.Sprintf("Turbo: %s\n", t.Name)
	for key, value := range t.Slices {
		s += fmt.Sprintf("%s Kg %s\n", key, value)
	}
	return s
}

func (t *Turbo) GetBootsCFMRange(boost float64) ([]*Range, error) {
	r, ok := t.Slices[KgKey(boost)]
	if !ok {
		return nil, errors.New("no range for boost")
	}

	return r, nil
}
