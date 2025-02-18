package pdf

import (
	"context"
	"fmt"
	"github.com/johnfercher/go-turbo/internal/core/models"
	"github.com/johnfercher/maroto/v2"
	"github.com/johnfercher/maroto/v2/pkg/components/chart"
	"github.com/johnfercher/maroto/v2/pkg/components/row"
	"github.com/johnfercher/maroto/v2/pkg/config"
	"github.com/johnfercher/maroto/v2/pkg/consts/pagesize"
	"github.com/johnfercher/maroto/v2/pkg/props"
)

type pdfReporter struct {
}

func NewPdfReporter() *pdfReporter {
	return &pdfReporter{}
}

func (p *pdfReporter) Generate(ctx context.Context, report *models.Report) error {
	matrix := p.ToTurboEfficiencyMatrix(ctx, report.Turbo.TurboScore)

	cfg := config.NewBuilder().
		WithDebug(true).
		WithPageSize(pagesize.A4).
		Build()

	m := maroto.New(cfg)

	m.AddRows(
		row.New(200).Add(
			chart.NewHeatMapCol(12, "Efficiency", matrix, props.HeatMap{
				TransparentValues: []int{0},
				InvertScale:       false,
				HalfColor:         false,
			}),
		),
	)

	doc, err := m.Generate()
	if err != nil {
		return err
	}

	return doc.Save(fmt.Sprintf("%s-%s-%.0f.pdf", report.Engine.Name, report.Turbo.Name, report.Boost))
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
