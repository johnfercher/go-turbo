package models

import (
	"fmt"
	"github.com/johnfercher/go-turbo/internal/core/models/fuel"
	"github.com/johnfercher/go-turbo/internal/math"
	"github.com/johnfercher/maroto/v2/pkg/components/row"
	"github.com/johnfercher/maroto/v2/pkg/components/text"
	"github.com/johnfercher/maroto/v2/pkg/core"
	"github.com/johnfercher/maroto/v2/pkg/core/entity"
	"github.com/johnfercher/maroto/v2/pkg/props"
)

var darkGrayColor = &props.Color{
	Red:   50,
	Green: 50,
	Blue:  50,
}

var lightGrayColor = &props.Color{
	Red:   200,
	Green: 200,
	Blue:  200,
}

var whiteColor = &props.Color{
	Red:   255,
	Green: 255,
	Blue:  255,
}

type Report struct {
	MinRPM       float64
	MaxRPM       float64
	Engine       *Engine
	Turbo        *Turbo
	Transmission *Transmission
	Fuel         *fuel.Fuel
	Boost        float64
	gearEntries  []Entries
	maxLbs       float64
	gear         int
}

func NewReport(minRPM float64, maxRPM float64, engine *Engine, turbo *Turbo, transmission *Transmission, fuel *fuel.Fuel, boost float64) *Report {
	return &Report{
		MinRPM:       minRPM,
		MaxRPM:       maxRPM,
		Engine:       engine,
		Turbo:        turbo,
		Transmission: transmission,
		Fuel:         fuel,
		Boost:        boost,
		gearEntries:  make([]Entries, transmission.MaxGear()),
		gear:         0,
	}
}

func (r *Report) GetMaxHP() *Entry {
	if len(r.gearEntries) == 0 {
		return nil
	}

	return r.gearEntries[0].GetMaxHP()
}

func (r *Report) GetMaxTorque() *Entry {
	if len(r.gearEntries) == 0 {
		return nil
	}

	return r.gearEntries[0].GetMaxTorque()
}

func (r *Report) GetGearEntries(gear int) Entries {
	return r.gearEntries[gear]
}

func (r *Report) GetChangeGearBest() []entity.Label {
	minSpeed := r.GetMinSpeed()
	maxSpeed := r.GetMaxSpeed()

	var labels []entity.Label
	for j := 0; j < len(r.gearEntries)-1; j++ {
		a := r.GetGearEntries(j)
		b := r.GetGearEntries(j + 1)

		found := false
		for i := minSpeed; i < maxSpeed; i++ {
			aRPM, aT, aF := a.GetRPMTorqueInSpeed(i)
			_, bT, bF := b.GetRPMTorqueInSpeed(i)

			if aF && bF && bT >= aT {
				found = true
				label := entity.NewLabel(fmt.Sprintf("%.0f / %.0f", aRPM, i), entity.NewPoint(i, bT))
				labels = append(labels, label)
				break
			}
		}
		if !found {
			rpm, torque, speed := a.GetMaxRPMTorqueInSpeed()
			label := entity.NewLabel(fmt.Sprintf("%.0f / %.0f", rpm, speed), entity.NewPoint(speed, torque))
			labels = append(labels, label)
		}
	}

	return labels
}

func (r *Report) GetMaxSpeed() float64 {
	lastGear := r.gearEntries[len(r.gearEntries)-1]
	return lastGear[len(lastGear)-1].Speed
}

func (r *Report) GetMinSpeed() float64 {
	firstGear := r.gearEntries[0]
	return firstGear[0].Speed
}

func (r *Report) GetMaxGear() int {
	return len(r.gearEntries)
}

func (r *Report) Add(rpm float64, cfm float64, health float64, wheelInch float64, tireHeight float64, speed int) {
	lbsMin := math.CubicFeetToLbsMin(cfm)

	power := NewPower(lbsMin, rpm, r.Fuel, r.Engine)

	e := &Entry{
		RPM:        rpm,
		LbsMin:     lbsMin,
		CFM:        cfm,
		Health:     health,
		Crankshaft: power,
		Speed:      r.Transmission.GetSpeed(rpm, wheelInch, tireHeight, speed),
		Gear:       speed,
		GearRatio:  r.Transmission.Gears[speed] * r.Transmission.Final,
	}

	r.gearEntries[r.gear] = append(r.gearEntries[r.gear], e)
}

func (r *Report) NextGear() bool {
	if r.gear < r.Transmission.MaxGear() {
		r.gear++
		return true
	}

	return false
}

type Entry struct {
	RPM        float64
	CFM        float64
	LbsMin     float64
	Health     float64
	Crankshaft *Power
	Speed      float64
	Gear       int
	GearRatio  float64
}

func (e Entry) GetHeader() core.Row {
	return row.New(5).
		Add(
			text.NewCol(2, "RPM", props.Text{Color: whiteColor}),
			text.NewCol(2, "CFM", props.Text{Color: whiteColor}),
			text.NewCol(3, "LBS/Min", props.Text{Color: whiteColor}),
			text.NewCol(3, "Health", props.Text{Color: whiteColor}),
			text.NewCol(3, "Torque", props.Text{Color: whiteColor}),
			text.NewCol(3, "Power", props.Text{Color: whiteColor}),
		).WithStyle(
		&props.Cell{
			BackgroundColor: darkGrayColor,
		})
}

func (e Entry) Values() []core.Col {
	return []core.Col{
		text.NewCol(2, fmt.Sprintf("%.0f", e.RPM)),
		text.NewCol(2, fmt.Sprintf("%.2f", e.CFM)),
		text.NewCol(3, fmt.Sprintf("%.2f", e.LbsMin)),
		text.NewCol(3, fmt.Sprintf("%.2f", e.Health)),
		text.NewCol(3, fmt.Sprintf("%.2f Kg", e.Crankshaft.Torque)),
		text.NewCol(3, fmt.Sprintf("%.2f HP", e.Crankshaft.HP)),
	}
}

func (e Entry) GetContent(i int) core.Row {
	r := row.New(5).
		Add(e.Values()...)
	if i%2 != 0 {
		r.WithStyle(&props.Cell{
			BackgroundColor: lightGrayColor,
		})
	}

	return r
}

type Entries []*Entry

func (e Entries) GetMaxHP() *Entry {
	maxHP := 0.0
	maxIndex := 0

	for i, entry := range e {
		if entry.Crankshaft.HP > maxHP {
			maxHP = entry.Crankshaft.HP
			maxIndex = i
		}
	}

	return e[maxIndex]
}

func (e Entries) GetMaxTorque() *Entry {
	maxTorque := 0.0
	maxIndex := 0

	for i, entry := range e {
		if entry.Crankshaft.Torque > maxTorque {
			maxTorque = entry.Crankshaft.Torque
			maxIndex = i
		}
	}

	return e[maxIndex]
}

func (e Entries) GetRPMTorqueInSpeed(speed float64) (float64, float64, bool) {
	for _, entry := range e {
		if entry.Speed >= speed {
			return entry.RPM, entry.Crankshaft.Torque * entry.GearRatio, true
		}
	}
	return 0, 0, false
}

func (e Entries) GetMaxRPMTorqueInSpeed() (float64, float64, float64) {
	max := e[len(e)-1]

	return max.RPM, max.Crankshaft.Torque * max.GearRatio, max.Speed
}

type Power struct {
	HP     float64
	Torque float64
}

func NewPower(lbsMin float64, rpm float64, f *fuel.Fuel, engine *Engine) *Power {
	powerGain := f.Compare(fuel.Gasoline100())
	hp := lbsMin * engine.CompressionRatio * powerGain * engine.EfficiencyRatio

	return &Power{
		HP:     hp,
		Torque: math.FootToKgfm(math.HorsepowerToTorque(hp, rpm)),
	}
}

func (p *Power) ToE85() *Power {
	return &Power{
		HP:     p.HP * 1.15,
		Torque: p.Torque * 1.15,
	}
}
