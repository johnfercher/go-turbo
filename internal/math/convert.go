package math

import "math"

func GetRate(minA, maxA, minB, maxB float64) float64 {
	flowRange := maxA - minA
	borderRange := math.Abs(minB - maxB)
	return borderRange / flowRange
}
