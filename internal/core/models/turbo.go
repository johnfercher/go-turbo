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

type TurboSlice struct {
	PSI    float64
	Ranges []*Range
}

func NewTurboSlice(psi float64, ranges ...*Range) *TurboSlice {
	return &TurboSlice{
		PSI:    psi,
		Ranges: ranges,
	}
}

func (t *TurboSlice) String() string {
	s := fmt.Sprintf("PSI: %.1f", t.PSI)
	for _, r := range t.Ranges {
		s += fmt.Sprintf(" %s", r)
	}

	return s
}

type Turbo struct {
	Name   string
	Slices []*TurboSlice
}

func NewTurbo(slices ...*TurboSlice) *Turbo {
	return &Turbo{
		Slices: slices,
	}
}

func (t *Turbo) String() string {
	s := fmt.Sprintf("Turbo: %s\n", t.Name)
	for _, r := range t.Slices {
		s += fmt.Sprintf("%s\n", r)
	}
	return s
}
