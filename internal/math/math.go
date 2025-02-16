package math

import "math"

const (
	ATM                 = 14.7
	FourCylinderConst   = 0.5
	CubicFeetConversion = 1728
	LitersToCID         = 0.01639344
)

func CubicInchToCubicFeet(inch float64) float64 {
	return inch / CubicFeetConversion
}

func LitersToCubicInch(liters float64) float64 {
	return liters / LitersToCID
}

func ATMToKg(psi float64) float64 {
	return psi / ATM
}

func KgToATM(kg float64) float64 {
	return kg * ATM
}

func PressureRatio(psi float64) float64 {
	return (psi + ATM) / ATM
}

func Distance(x1, y1, x2, y2 float64) float64 {
	return math.Sqrt(((x1 - x2) * (x1 - x2)) + ((y1 - y2) * (y1 - y2)))
}

func AngleBetween(x1, y1, x2, y2 float64) float64 {
	return RadianBetween(x1, y1, x2, y2) * (180.0 / math.Pi)
}

func RadianBetween(x1, y1, x2, y2 float64) float64 {
	return math.Atan2(x2-x1, y2-y1)
}
