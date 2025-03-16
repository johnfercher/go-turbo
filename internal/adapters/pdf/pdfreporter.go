package pdf

import (
	"context"
	"fmt"
	"github.com/johnfercher/go-turbo/internal/core/models"
	"github.com/johnfercher/maroto/v2"
	"github.com/johnfercher/maroto/v2/pkg/components/chart"
	"github.com/johnfercher/maroto/v2/pkg/components/list"
	"github.com/johnfercher/maroto/v2/pkg/components/row"
	"github.com/johnfercher/maroto/v2/pkg/components/text"
	"github.com/johnfercher/maroto/v2/pkg/config"
	"github.com/johnfercher/maroto/v2/pkg/consts/pagesize"
	"github.com/johnfercher/maroto/v2/pkg/core"
	"github.com/johnfercher/maroto/v2/pkg/props"
)

type pdfReporter struct {
}

func NewPdfReporter() *pdfReporter {
	return &pdfReporter{}
}

func (p *pdfReporter) Generate(ctx context.Context, reports []*models.Report) error {
	cfg := config.NewBuilder().
		WithPageSize(pagesize.A4).
		WithMaxGridSize(16).
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

	return doc.Save(fmt.Sprintf("%s.pdf", reports[0].Engine.Name))
}

func (p *pdfReporter) generate(ctx context.Context, m core.Maroto, report *models.Report) error {
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

	m.AddRow(5,
		text.NewCol(6, report.Engine.Name),
		text.NewCol(6, report.Turbo.Name),
		text.NewCol(6, fmt.Sprintf("%.2f", report.Boost)),
	)

	rng := 500
	var entries []*models.Entry
	for i, entry := range report.Entries {
		if i%rng == 0 {
			entries = append(entries, entry)
		}
	}

	rows, err := list.Build(entries)
	if err != nil {
		return err
	}

	m.AddRows(rows...)

	maxHP := report.Entries.GetMaxHP()
	maxTorque := report.Entries.GetMaxTorque()

	m.AddRows(
		row.New(5).
			Add(
				text.NewCol(2, fmt.Sprintf("%d", maxTorque.RPM)),
				text.NewCol(2, fmt.Sprintf("%.2f", maxTorque.CFM)),
				text.NewCol(2, fmt.Sprintf("%.2f", maxTorque.LbsMin)),
				text.NewCol(2, fmt.Sprintf("%.2f", maxTorque.Health)),
				text.NewCol(2, fmt.Sprintf("%.2f Kg", maxTorque.Crankshaft.Torque)),
				text.NewCol(2, fmt.Sprintf("%.2f HP", maxTorque.Crankshaft.HP)),
				text.NewCol(2, fmt.Sprintf("%.2f Kg", maxTorque.Crankshaft.ToE85().Torque)),
				text.NewCol(2, fmt.Sprintf("%.2f HP", maxTorque.Crankshaft.ToE85().HP)),
			).WithStyle(
			&props.Cell{
				BackgroundColor: &props.Color{
					Red:   100,
					Green: 100,
					Blue:  200,
				},
			},
		),
		row.New(5).
			Add(
				text.NewCol(2, fmt.Sprintf("%d", maxHP.RPM)),
				text.NewCol(2, fmt.Sprintf("%.2f", maxHP.CFM)),
				text.NewCol(2, fmt.Sprintf("%.2f", maxHP.LbsMin)),
				text.NewCol(2, fmt.Sprintf("%.2f", maxHP.Health)),
				text.NewCol(2, fmt.Sprintf("%.2f Kg", maxHP.Crankshaft.Torque)),
				text.NewCol(2, fmt.Sprintf("%.2f HP", maxHP.Crankshaft.HP)),
				text.NewCol(2, fmt.Sprintf("%.2f Kg", maxHP.Crankshaft.ToE85().Torque)),
				text.NewCol(2, fmt.Sprintf("%.2f HP", maxHP.Crankshaft.ToE85().HP)),
			).WithStyle(
			&props.Cell{
				BackgroundColor: &props.Color{
					Red:   100,
					Green: 200,
					Blue:  100,
				},
			},
		),
	)

	return nil
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
