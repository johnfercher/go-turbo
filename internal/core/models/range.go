package models

import (
	"fmt"
	"github.com/johnfercher/go-turbo/internal/math"
)

type Range struct {
	Begin float64
	End   float64
}

func NewRange(begin, end float64) *Range {
	_begin := begin
	_end := end

	if end < begin {
		_begin = end
		_end = begin
	}

	return &Range{
		Begin: _begin,
		End:   _end,
	}
}

func (r *Range) GetRate(o *Range) float64 {
	return math.GetRate(r.Begin, r.End, o.Begin, o.End)
}

func (r *Range) String() string {
	return fmt.Sprintf("[%f, %f]", r.Begin, r.End)
}
