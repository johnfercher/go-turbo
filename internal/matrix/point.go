package matrix

type Point struct {
	X     int
	Y     int
	Value int
}

type Points []Point

func (p Points) GetValue(pressure int, flow int) int {
	for _, point := range p {
		if point.X == flow && point.Y == pressure {
			return point.Value
		}
	}

	return 0
}
