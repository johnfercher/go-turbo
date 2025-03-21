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

	torque := p.getTorque(report)
	maxTorque := report.GetMaxTorque()
	hp := p.getHP(report)
	maxHP := report.GetMaxHP()

	m.AddRows(
		row.New(70).Add(
			chart.NewTimeSeriesCol(7, torque, props.Chart{
				XLabels: []float64{1000, 3000, 6000, 9000},
			}),
			col.New(2),
			chart.NewTimeSeriesCol(7, hp, props.Chart{
				XLabels: []float64{1000, 3000, 6000, 9000},
			}),
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

	return nil
}

func (p *pdfReporter) getTransmission(report *models.Report) []core.Row {
	transmission := report.Transmission

	// Header
	rows := []core.Row{
		row.New(5),
		row.New(5).Add(
			text.NewCol(4, transmission.Name),
			text.NewCol(4, fmt.Sprintf("Final(%f)", transmission.Final)),
			text.NewCol(4, "Gear"),
			text.NewCol(4, "Ratio"),
		),
	}

	// Gear ratio
	for i, gear := range transmission.Gears {
		rows = append(rows, row.New(5).Add(
			text.NewCol(8, fmt.Sprintf("%d", i+1)),
			text.NewCol(8, fmt.Sprintf("%f", gear)),
		))
	}

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
	minSpeed := report.GetMinSpeed()

	deltaSpeed := maxSpeed - minSpeed
	pieceDelta := deltaSpeed / 5

	for i := minSpeed; i <= maxSpeed; i += pieceDelta {
		props.XLabels = append(props.XLabels, i)
	}

	maxTorque := report.GetMaxTorque().Crankshaft.Torque * report.Transmission.GetGearRatioFinal(0)
	pieceTorque := maxTorque / 5
	for i := 0.0; i < maxTorque; i += pieceTorque {
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

	return timeSeries, props
}

func (p *pdfReporter) getHP(report *models.Report) []entity.TimeSeries {
	var hpList []entity.TimeSeries
	var hp = []entity.Point{}

	for i := report.MinRPM; i < report.MaxRPM; i += 100 {
		cfm := report.Engine.GetCFM(float64(i), report.Boost)

		lbsMin := math.CubicFeetToLbsMin(cfm.Flow)

		power := models.NewPower(lbsMin, i, report.Fuel, report.Engine)
		hp = append(hp, entity.NewPoint(i, power.HP))
	}

	maxTorque := report.GetMaxHP()
	lbsMin := math.CubicFeetToLbsMin(maxTorque.CFM)
	power := models.NewPower(lbsMin, maxTorque.RPM, report.Fuel, report.Engine)

	hpTimeSeries := entity.NewTimeSeries(Color(1), hp, entity.NewLabel("Max", entity.Point{
		X: float64(maxTorque.RPM),
		Y: power.HP,
	}))

	hpList = append(hpList, hpTimeSeries)

	return hpList
}

func (p *pdfReporter) getTorque(report *models.Report) []entity.TimeSeries {
	var torqueList []entity.TimeSeries
	var torque = []entity.Point{}
	for i := 0.0; i < 9000; i += 100 {
		cfm := report.Engine.GetCFM(i, report.Boost)

		lbsMin := math.CubicFeetToLbsMin(cfm.Flow)

		power := models.NewPower(lbsMin, i, report.Fuel, report.Engine)
		torque = append(torque, entity.NewPoint(i, power.Torque))
	}

	maxTorque := report.GetMaxTorque()
	lbsMin := math.CubicFeetToLbsMin(maxTorque.CFM)
	power := models.NewPower(lbsMin, maxTorque.RPM, report.Fuel, report.Engine)

	torqueTimeSeries := entity.NewTimeSeries(Color(2), torque, entity.NewLabel("Max", entity.Point{
		X: float64(maxTorque.RPM),
		Y: power.Torque,
	}))

	torqueList = append(torqueList, torqueTimeSeries)

	return torqueList
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
