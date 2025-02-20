package models

import "fmt"

type Point struct {
	X, Y int
}

func (p *Point) String() string {
	return fmt.Sprintf("(%d, %d)", p.X, p.Y)
}
