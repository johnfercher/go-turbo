package models

import "fmt"

type Range struct {
	Min   float64
	Max   float64
	Score float64
}

func (r Range) GreaterThan(a Range) bool {
	return r.Min > a.Min
}

func NewRange(min, max, score float64) *Range {
	return &Range{
		Min:   min,
		Max:   max,
		Score: score,
	}
}

func (r *Range) String() string {
	return fmt.Sprintf("[%.0f-%.0f]%.0f", r.Min, r.Max, r.Score)
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
