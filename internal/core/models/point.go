package models

import "fmt"

type Point struct {
	X, Y float64
}

func (p *Point) String() string {
	return fmt.Sprintf("(%f, %f)", p.X, p.Y)
}
