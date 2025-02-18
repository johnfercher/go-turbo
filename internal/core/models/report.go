package models

import (
	"fmt"
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

func (r *Report) Add(rpm int, lbsMin float64, health float64) {
	e := &Entry{
		RPM:      rpm,
		LbsMin:   lbsMin,
		Health:   health,
		MinHP:    lbsMin * 9.5,
		MaxHP:    lbsMin * 10.5,
		MinHPE85: lbsMin * 9.5 * 1.15,
		MaxHPE85: lbsMin * 10.5 * 1.15,
	}

	r.Entries = append(r.Entries, e)
}

type Entry struct {
	RPM      int
	LbsMin   float64
	Health   float64
	MinHP    float64
	MaxHP    float64
	MinHPE85 float64
	MaxHPE85 float64
}

func (e Entry) GetHeader() core.Row {
	return row.New(5).
		Add(
			text.NewCol(2, "RPM", props.Text{Color: whiteColor}),
			text.NewCol(2, "LBS/Min", props.Text{Color: whiteColor}),
			text.NewCol(2, "Health", props.Text{Color: whiteColor}),
			text.NewCol(3, "HP", props.Text{Color: whiteColor}),
			text.NewCol(3, "HPE85", props.Text{Color: whiteColor}),
		).WithStyle(
		&props.Cell{
			BackgroundColor: darkGrayColor,
		})
}

func (e Entry) GetContent(i int) core.Row {
	r := row.New(5).
		Add(
			text.NewCol(2, fmt.Sprintf("%d", e.RPM)),
			text.NewCol(2, fmt.Sprintf("%.2f", e.LbsMin)),
			text.NewCol(2, fmt.Sprintf("%.2f", e.Health)),
			text.NewCol(3, fmt.Sprintf("(%.2f~%.2f)HP", e.MinHP, e.MaxHP)),
			text.NewCol(3, fmt.Sprintf("(%.2f~%.2f)HP", e.MinHPE85, e.MaxHPE85)),
		)

	if i%2 != 0 {
		r.WithStyle(&props.Cell{
			BackgroundColor: lightGrayColor,
		})
	}

	return r
}

type Entries []*Entry

func (e Entries) GetTop() *Entry {
	maxLbs := 0.0
	maxIndex := 0

	for i, entry := range e {
		if entry.LbsMin > maxLbs {
			maxLbs = entry.LbsMin
			maxIndex = i
		}
	}

	return e[maxIndex]
}
