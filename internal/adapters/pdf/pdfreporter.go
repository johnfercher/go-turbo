package pdf

import (
	"context"
	"fmt"
	"github.com/johnfercher/go-turbo/internal/core/models"
	"github.com/johnfercher/go-turbo/internal/math"
	"github.com/johnfercher/maroto/v2"
	"github.com/johnfercher/maroto/v2/pkg/components/chart"
	"github.com/johnfercher/maroto/v2/pkg/components/col"
	"github.com/johnfercher/maroto/v2/pkg/components/list"
	"github.com/johnfercher/maroto/v2/pkg/components/row"
	"github.com/johnfercher/maroto/v2/pkg/components/text"
	"github.com/johnfercher/maroto/v2/pkg/config"
	"github.com/johnfercher/maroto/v2/pkg/consts/pagesize"
	"github.com/johnfercher/maroto/v2/pkg/core"
	"github.com/johnfercher/maroto/v2/pkg/core/entity"
	"github.com/johnfercher/maroto/v2/pkg/props"
)

type pdfReporter struct {
}

func NewPdfReporter() *pdfReporter {
	return &pdfReporter{}
}

func (p *pdfReporter) Generate(ctx context.Context, file string, reports []*models.Report) error {
	cfg := config.NewBuilder().
		WithPageSize(pagesize.A4).
		WithMaxGridSize(16).
		//WithDebug(true).
		Build()

	m := maroto.New(cfg)

	for _, report := range reports {
		err := p.generate(ctx, m, report)
		if err != nil {
			return err
		}
	}

	doc, err := m.Generate()
	if err != nil {
		return err
	}

	return doc.Save(fmt.Sprintf("%s.pdf", file))
}

func (p *pdfReporter) generate(ctx context.Context, m core.Maroto, report *models.Report) error {
	var turboName string
	if report.Turbo != nil {
		turboName = report.Turbo.Name
	}

	m.AddRow(5,
		text.NewCol(4, report.Engine.Name),
		text.NewCol(4, report.Fuel.Name),
		text.NewCol(4, turboName),
		text.NewCol(4, fmt.Sprintf("%.2f", report.Boost)),
	)

	m.AddRow(5)

	if report.Turbo != nil {
		turboName = report.Turbo.Name
		matrix := p.ToTurboEfficiencyMatrix(ctx, report.Turbo.TurboScore)

		m.AddRows(
			row.New(150).Add(
				chart.NewHeatMapCol(16, "Efficiency", matrix, props.HeatMap{
					TransparentValues: []int{0},
					InvertScale:       false,
					HalfColor:         false,
				}),
			),
		)

		m.AddRow(5)
	}

	torque, torqueProps := p.getTorque(report)
	maxTorque := report.GetMaxTorque()
	hp, hpProps := p.getHP(report)
	maxHP := report.GetMaxHP()

	m.AddRows(
		row.New(70).Add(
			chart.NewTimeSeriesCol(7, torque, torqueProps),
			col.New(2),
			chart.NewTimeSeriesCol(7, hp, hpProps),
		),
	)

	m.AddRow(5)

	rng := 500
	var entries []*models.Entry
	for i, entry := range report.GetGearEntries(1) {
		if i%rng == 0 {
			entries = append(entries, entry)
		}
	}

	rows, err := list.Build(entries)
	if err != nil {
		return err
	}

	m.AddRows(rows...)

	m.AddRows(
		row.New(5).
			Add(maxTorque.Values()...).WithStyle(
			&props.Cell{
				BackgroundColor: &props.Color{
					Red:   100,
					Green: 100,
					Blue:  200,
				},
			},
		),
		row.New(5).
			Add(maxHP.Values()...).WithStyle(
			&props.Cell{
				BackgroundColor: &props.Color{
					Red:   100,
					Green: 200,
					Blue:  100,
				},
			},
		),
	)

	m.AddRows(m.AddRow(5))

	m.AddRows(p.getTransmission(report)...)

	m.AddRows(m.AddRow(10))

	return nil
}

func (p *pdfReporter) getTransmission(report *models.Report) []core.Row {
	transmission := report.Transmission

	// Header
	rows := []core.Row{
		row.New(5).Add(col.New(16)),
		row.New(5).Add(
			text.NewCol(4, transmission.Name),
			text.NewCol(4, fmt.Sprintf("Final(%f)", transmission.Final)),
			text.NewCol(4, "Gear"),
			text.NewCol(4, "Ratio"),
		)}

	// Gear ratio
	for i, gear := range transmission.Gears {
		r := row.New(5).Add(
			col.New(8),
			text.NewCol(4, fmt.Sprintf("%d", i+1)),
			text.NewCol(4, fmt.Sprintf("%f", gear)),
		)
		if i%2 == 0 {
			r.WithStyle(&props.Cell{
				BackgroundColor: &props.Color{
					Red:   200,
					Green: 200,
					Blue:  200,
				},
			})
		}
		rows = append(rows, r)
	}

	rows = append(rows, row.New(5).Add(col.New(16)))

	speedTorque, speedTorqueProps := p.getSpeedTorque(report)

	rows = append(rows, row.New(70).Add(
		chart.NewTimeSeriesCol(16, speedTorque, speedTorqueProps),
	))

	return rows
}

func (p *pdfReporter) getSpeedTorque(report *models.Report) ([]entity.TimeSeries, props.Chart) {
	var timeSeries []entity.TimeSeries
	var props props.Chart

	maxSpeed := report.GetMaxSpeed()
	stepSpeed := maxSpeed / 5

	for i := 0.0; i <= maxSpeed; i += stepSpeed {
		props.XLabels = append(props.XLabels, i)
	}

	maxTorque := report.GetMaxTorque().Crankshaft.Torque * report.Transmission.GetGearRatioFinal(0)
	pieceTorque := maxTorque / 5
	for i := 0.0; i <= maxTorque; i += pieceTorque {
		props.YLabels = append(props.YLabels, i)
	}

	for gear := 0; gear < report.GetMaxGear(); gear++ {
		var points []entity.Point
		entries := report.GetGearEntries(gear)

		for _, entry := range entries {
			points = append(points, entity.Point{
				X: entry.Speed,
				Y: entry.Crankshaft.Torque * entry.GearRatio,
			})
		}

		timeSeries = append(timeSeries, entity.NewTimeSeries(Color(gear), points))
	}

	bestGears := report.GetChangeGearBest()
	for i, gear := range bestGears {
		timeSeries[i].Labels = append(timeSeries[i].Labels, gear)
	}

	return timeSeries, props
}

func (p *pdfReporter) getHP(report *models.Report) ([]entity.TimeSeries, props.Chart) {
	var hpList []entity.TimeSeries
	var props props.Chart
	var hp = []entity.Point{}

	maxRPM := report.MaxRPM
	stepRPM := maxRPM / 5

	for i := 0.0; i <= maxRPM; i += stepRPM {
		props.XLabels = append(props.XLabels, i)
	}

	entries := report.GetGearEntries(0)
	for _, entry := range entries {
		hp = append(hp, entity.NewPoint(entry.RPM, entry.Crankshaft.HP))
	}

	maxTorque := report.GetMaxHP()
	lbsMin := math.CubicFeetToLbsMin(maxTorque.CFM)
	power := models.NewPower(lbsMin, maxTorque.RPM, report.Fuel, report.Engine)

	hpTimeSeries := entity.NewTimeSeries(Color(1), hp, entity.NewLabel("Max", entity.Point{
		X: maxTorque.RPM,
		Y: power.HP,
	}))

	hpList = append(hpList, hpTimeSeries)

	return hpList, props
}

func (p *pdfReporter) getTorque(report *models.Report) ([]entity.TimeSeries, props.Chart) {
	var torqueList []entity.TimeSeries
	var props props.Chart
	var torque = []entity.Point{}

	maxRPM := report.MaxRPM
	stepRPM := maxRPM / 5

	for i := 0.0; i <= maxRPM; i += stepRPM {
		props.XLabels = append(props.XLabels, i)
	}

	entries := report.GetGearEntries(0)
	for _, entry := range entries {
		torque = append(torque, entity.NewPoint(entry.RPM, entry.Crankshaft.Torque))
	}

	maxTorque := report.GetMaxTorque()
	lbsMin := math.CubicFeetToLbsMin(maxTorque.CFM)
	power := models.NewPower(lbsMin, maxTorque.RPM, report.Fuel, report.Engine)

	torqueTimeSeries := entity.NewTimeSeries(Color(2), torque, entity.NewLabel("Max", entity.Point{
		X: maxTorque.RPM,
		Y: power.Torque,
	}))

	torqueList = append(torqueList, torqueTimeSeries)

	return torqueList, props
}

func (p *pdfReporter) ToTurboEfficiencyMatrix(ctx context.Context, turbo [][]float64) [][]int {
	xSize := len(turbo)
	ySize := len(turbo[0])

	var matrix [][]int
	for i := 0; i < xSize; i++ {
		var line []int
		for j := 0; j < ySize; j++ {
			line = append(line, int(turbo[i][j]))
		}
		matrix = append(matrix, line)
	}

	return matrix
}
