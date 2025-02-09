package math

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
