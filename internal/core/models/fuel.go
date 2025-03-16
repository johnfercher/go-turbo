package models

type Fuel struct {
	Name       string
	MinOctane  float64
	MaxOctane  float64
	LatentHeat float64
	HeatPower  float64
	AFRatio    float64
}

func (f *Fuel) PowerDiffTo(ref *Fuel) float64 {
	return 0.0
}

func (f *Fuel) AVGOctane() float64 {
	return (f.MinOctane + f.MaxOctane) / 2.0
}

func BRGasoline() *Fuel {
	return &Fuel{
		Name:       "BR Gasoline",
		MinOctane:  92,
		MaxOctane:  92,
		LatentHeat: 350,
		AFRatio:    13.56,
		HeatPower:  11.220,
	}
}

func Gasoline100() *Fuel {
	return &Fuel{
		Name:       "Gasoline 100%",
		MinOctane:  87,
		MaxOctane:  87,
		LatentHeat: 350,
		AFRatio:    14.57,
		HeatPower:  11.220,
	}
}

func VpGasolineC16() *Fuel {
	return &Fuel{
		Name:       "VP Gasoline C16",
		MinOctane:  120,
		MaxOctane:  120,
		LatentHeat: 350,
		AFRatio:    14.57,
		HeatPower:  11.220,
	}
}

func BRPodiumGasoline() *Fuel {
	return &Fuel{
		Name:       "BR Podium Gasoline",
		MinOctane:  102,
		MaxOctane:  102,
		LatentHeat: 350,
		AFRatio:    13.56,
		HeatPower:  11.220,
	}
}

func BRAlcohol() *Fuel {
	return &Fuel{
		Name:       "BR Alcohol",
		MinOctane:  110,
		MaxOctane:  110,
		LatentHeat: 900,
		AFRatio:    8.5, // Hidratado
		HeatPower:  6.507,
	}
}

func BRMethanol() *Fuel {
	return &Fuel{
		Name:       "BR Methanol",
		MinOctane:  115,
		MaxOctane:  115,
		LatentHeat: 1100,
		AFRatio:    6.47,
		HeatPower:  5.330,
	}
}

func E85() *Fuel {
	return &Fuel{
		Name:       "E85",
		MinOctane:  100,
		MaxOctane:  105,
		LatentHeat: 900,
		AFRatio:    8.99,
		HeatPower:  6.507,
	}
}

func E100() *Fuel {
	return &Fuel{
		Name:       "Alcohol 100%",
		MinOctane:  110,
		MaxOctane:  110,
		LatentHeat: 900,
		AFRatio:    8.99,
		HeatPower:  6.507,
	}
}
