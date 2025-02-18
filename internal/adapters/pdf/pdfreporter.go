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

	return doc.Save(fmt.Sprintf("%s-%s.pdf", reports[0].Engine.Name, reports[0].Turbo.Name))
}

func (p *pdfReporter) generate(ctx context.Context, m core.Maroto, report *models.Report) error {
	matrix := p.ToTurboEfficiencyMatrix(ctx, report.Turbo.TurboScore)

	m.AddRows(
		row.New(200).Add(
			chart.NewHeatMapCol(12, "Efficiency", matrix, props.HeatMap{
				TransparentValues: []int{0},
				InvertScale:       false,
				HalfColor:         false,
			}),
		),
	)

	m.AddRow(5,
		text.NewCol(4, report.Engine.Name),
		text.NewCol(4, report.Turbo.Name),
		text.NewCol(4, fmt.Sprintf("%.2f", report.Boost)),
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

	e := report.Entries.GetTop()

	r := row.New(5).
		Add(
			text.NewCol(2, fmt.Sprintf("%d", e.RPM)),
			text.NewCol(2, fmt.Sprintf("%.2f", e.LbsMin)),
			text.NewCol(2, fmt.Sprintf("%.2f", e.Health)),
			text.NewCol(3, fmt.Sprintf("(%.2f~%.2f)HP", e.MinHP, e.MaxHP)),
			text.NewCol(3, fmt.Sprintf("(%.2f~%.2f)HP", e.MinHPE85, e.MaxHPE85)),
		).WithStyle(
		&props.Cell{
			BackgroundColor: &props.Color{
				Red:   100,
				Green: 200,
				Blue:  100,
			},
		},
	)

	m.AddRows(r)

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
