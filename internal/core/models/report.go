package models

import (
	"fmt"
	"github.com/johnfercher/go-turbo/internal/math"
	"github.com/johnfercher/maroto/v2/pkg/components/row"
	"github.com/johnfercher/maroto/v2/pkg/components/text"
	"github.com/johnfercher/maroto/v2/pkg/core"
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
	Engine  *Engine
	Turbo   *Turbo
	Boost   float64
	Entries Entries
	maxLbs  float64
}

func NewReport(engine *Engine, turbo *Turbo, boost float64) *Report {
	return &Report{
		Engine: engine,
		Turbo:  turbo,
		Boost:  boost,
	}
}

func (r *Report) Add(rpm int, cfm float64, health float64) {
	lbsMin := math.CubicFeetToLbsMin(cfm)

	power := NewPower(lbsMin, rpm, r.Engine.CompressionRatio, r.Engine.BoostGainRatio)

	e := &Entry{
		RPM:        rpm,
		LbsMin:     lbsMin,
		CFM:        cfm,
		Health:     health,
		Crankshaft: power,
	}

	r.Entries = append(r.Entries, e)
}

type Entry struct {
	RPM        int
	CFM        float64
	LbsMin     float64
	Health     float64
	Crankshaft *Power
}

func (e Entry) GetHeader() core.Row {
	return row.New(5).
		Add(
			text.NewCol(2, "RPM", props.Text{Color: whiteColor}),
			text.NewCol(2, "CFM", props.Text{Color: whiteColor}),
			text.NewCol(2, "LBS/Min", props.Text{Color: whiteColor}),
			text.NewCol(2, "Health", props.Text{Color: whiteColor}),
			text.NewCol(2, "Torque", props.Text{Color: whiteColor}),
			text.NewCol(2, "Power", props.Text{Color: whiteColor}),
			text.NewCol(2, "Torque E85", props.Text{Color: whiteColor}),
			text.NewCol(2, "Power E85", props.Text{Color: whiteColor}),
		).WithStyle(
		&props.Cell{
			BackgroundColor: darkGrayColor,
		})
}

func (e Entry) GetContent(i int) core.Row {
	r := row.New(5).
		Add(
			text.NewCol(2, fmt.Sprintf("%d", e.RPM)),
			text.NewCol(2, fmt.Sprintf("%.2f", e.CFM)),
			text.NewCol(2, fmt.Sprintf("%.2f", e.LbsMin)),
			text.NewCol(2, fmt.Sprintf("%.2f", e.Health)),
			text.NewCol(2, fmt.Sprintf("%.2f Kg", e.Crankshaft.Torque)),
			text.NewCol(2, fmt.Sprintf("%.2f HP", e.Crankshaft.HP)),
			text.NewCol(2, fmt.Sprintf("%.2f Kg", e.Crankshaft.ToE85().Torque)),
			text.NewCol(2, fmt.Sprintf("%.2f HP", e.Crankshaft.ToE85().HP)),
		)

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

type Power struct {
	HP     float64
	Torque float64
}

func NewPower(lbsMin float64, rpm int, compressionRatio float64, boostGainRatio float64) *Power {
	hp := lbsMin * compressionRatio * boostGainRatio
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
