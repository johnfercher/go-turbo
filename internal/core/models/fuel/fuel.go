package fuel

type Fuel struct {
	Name       string
	Octane     float64
	LatentHeat float64
	HeatPower  float64
	AFRatio    float64
}

func (f *Fuel) Power() float64 {
	afHeatPower := f.HeatPower / f.AFRatio
	return afHeatPower * f.Octane
}

func (f *Fuel) Compare(ref *Fuel) float64 {
	fix := 0.85
	value := f.Power() / ref.Power()
	if value > 1 {
		sup := value - 1
		sup *= fix
		calc := 1 + sup
		return calc
	}

	sub := 1 - value
	sub *= fix
	calc := 1 - sub
	return calc
}

func Mix(name string, a *Fuel, percentA float64, b *Fuel, percentB float64) *Fuel {
	return &Fuel{
		Name:       name,
		Octane:     (a.Octane * percentA / 100.0) + (b.Octane * percentB / 100.0),
		LatentHeat: (a.LatentHeat * percentA / 100.0) + (b.LatentHeat * percentB / 100.0),
		HeatPower:  (a.HeatPower * percentA / 100.0) + (b.HeatPower * percentB / 100.0),
		AFRatio:    (a.AFRatio * percentA / 100.0) + (b.AFRatio * percentB / 100.0),
	}
}

func Gasoline100() *Fuel {
	return &Fuel{
		Name:       "Gasoline 100%",
		Octane:     87,
		LatentHeat: 350,
		AFRatio:    14.57,
		HeatPower:  11200,
	}
}

func E85() *Fuel {
	gasoline100 := Gasoline100()
	e100 := E100()

	return Mix("E85", e100, 85, gasoline100, 15)
}

func E100() *Fuel {
	return &Fuel{
		Name:       "Ethanol 100%",
		Octane:     110,
		LatentHeat: 900,
		AFRatio:    8.99,
		HeatPower:  6507,
	}
}

func BRGasoline() *Fuel {
	gasoline100 := Gasoline100()
	e100 := E100()

	return Mix("BR Gasoline", e100, 22, gasoline100, 78)
}

func BRPodiumGasoline() *Fuel {
	gasoline100 := Gasoline100()
	e100 := E100()

	return Mix("BR Podium Gasoline", e100, 67, gasoline100, 33)
}

func BREthanol() *Fuel {
	return &Fuel{
		Name:       "BR Ethanol",
		Octane:     110,
		LatentHeat: 900,
		AFRatio:    8.5, // Hidratado
		HeatPower:  6507,
	}
}

func BRMethanol() *Fuel {
	return &Fuel{
		Name:       "BR Methanol",
		Octane:     115,
		LatentHeat: 1100,
		AFRatio:    6.47,
		HeatPower:  5330,
	}
}
