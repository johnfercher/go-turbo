package matrix

import (
	"fmt"
	"math/rand"
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

func (p Points) Interpolate(dx, dy int) Points {
	var newPoints Points

	values := make(map[int]bool)
	for _, point := range p {
		values[int(point.Value)] = true
	}

	valuablesMap := make(map[int]Points)
	for value, _ := range values {
		valuablesMap[value] = p.FilterFromValue(value)
	}

	for _, valuables := range valuablesMap {
		newValuables := Interpolate(valuables, dx, dy)
		newPoints = append(newPoints, newValuables...)
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

func Interpolate(p Points, dx int, dy int) Points {
	newPointsMap := make(map[string]Point)

	for _, point := range p {
		newPointsMap[getKey(point)] = point
	}

	sizeP := len(p)

	for i := 0; i < sizeP; i++ {
		for j := 0; j < sizeP; j++ {
			if i == j {
				continue
			}

			if p[i].Y == p[j].Y {
				continue
			}

			x1, y1, w1 := p[i].X, p[i].Y, p[i].Value
			x2, y2, w2 := p[j].X, p[j].Y, p[j].Value

			halfX := (x1 + x2) / 2.0
			halfY := (y1 + y2) / 2.0
			halfW := (w1 + w2) / 2.0

			haveRandX := rand.Intn(100)%2 == 0
			if haveRandX {
				halfX += float64(rand.Intn(dx))
			}

			haveRandY := rand.Intn(100)%2 == 0
			if haveRandY {
				halfY += float64(rand.Intn(dy))
			}

			newPoint := NewPoint(halfX, halfY, halfW)
			_, ok := newPointsMap[getKey(newPoint)]
			if !ok {
				newPointsMap[getKey(newPoint)] = newPoint
			}
		}
	}

	var newPoints Points
	for _, point := range newPointsMap {
		newPoints = append(newPoints, point)
	}

	return newPoints
}

func getKey(p Point) string {
	return fmt.Sprintf("%d-%d", int(p.X), int(p.Y))
}
