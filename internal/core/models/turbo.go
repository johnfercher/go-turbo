package models

import "fmt"

type Range struct {
	Min   float64
	Max   float64
	Score float64
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
	PSI    string
	Ranges []*Range
}

type Turbo struct {
	Slices []*TurboSlice
}
