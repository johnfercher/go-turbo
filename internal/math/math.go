package math

import "math"

const (
	ATM                 = 14.7
	FourCylinderConst   = 0.5
	CubicFeetConversion = 1728
	LitersToCID         = 0.01639344
	milimetersToInch    = 0.0393701
)

func CubicInchToCubicFeet(inch float64) float64 {
	return inch / CubicFeetConversion
}

func CubicFeetToLbsMin(cfm float64) float64 {
	return cfm * 0.069
}

func MilimetersToInch(milimeters float64) float64 {
	return milimeters * milimetersToInch
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

func HorsepowerToTorque(horsePower float64, rpm float64) float64 {
	return horsePower / (rpm / 5252.0)
}

func FootToKgfm(ft float64) float64 {
	return ft * 0.1382549544
}

func EnvironmentAltitudePower(altitude float64) float64 {
	oceanLevel := 1.0
	offset := altitude / 300.0
	return oceanLevel - (offset * 3 / 100.0)
}

func EnvironmentTemperaturePowerLoss(temperature float64) float64 {
	return temperature / 7.0
}
