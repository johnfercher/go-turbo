package models

type Combustion struct {
	Humidity       float64
	AirPressure    float64
	AirTemperature float64
	Fuel           *Fuel
}
