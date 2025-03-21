package matrix

import (
	"gonum.org/v1/gonum/interp"
)

type Point struct {
	X     float64
	Y     float64
	Value float64
}

func NewPoint(x float64, y float64, value float64) Point {
	return Point{
		X:     x,
		Y:     y,
		Value: value,
	}
}

type Points []Point

func (p Points) GetValue(pressure int, flow int) float64 {
	for _, point := range p {
		if point.X == float64(flow) && point.Y == float64(pressure) {
			return point.Value
		}
	}

	return 0
}

func (p Points) Interpolate() Points {
	var newPoints Points

	valuesSegment := make(map[int]bool)
	for _, point := range p {
		valuesSegment[int(point.Value)] = true
	}

	valuesSegmentArray := make(map[int]Points)
	for value, _ := range valuesSegment {
		valuesSegmentArray[value] = p.FilterFromValue(value)
	}

	for value, valuables := range valuesSegmentArray {
		inter := interp.AkimaSpline{}

		var xs []float64
		var ys []float64

		for _, point := range valuables {
			xs = append(xs, point.X)
			ys = append(ys, point.Y)
		}

		_ = inter.Fit(xs, ys)

		minX := valuables[0].X
		maxX := valuables[len(valuables)-1].X

		for i := minX; i < maxX; i++ {
			x := inter.Predict(float64(i))
			newPoints = append(newPoints, NewPoint(float64(i), x, float64(value)))
		}
	}

	return newPoints
}

func (p Points) FilterFromValue(value int) (a Points) {
	for _, point := range p {
		if int(point.Value) == value {
			a = append(a, point)
		}
	}

	return
}
