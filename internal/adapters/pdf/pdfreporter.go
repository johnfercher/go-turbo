package pdf

import (
	"context"
	"github.com/johnfercher/go-turbo/internal/core/consts"
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

func (p *pdfReporter) Generate(ctx context.Context, turbo [][]*models.TurboScore, reportType consts.ReportType) error {
	matrix := p.ToTurboEfficiencyMatrix(ctx, turbo)

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

	return doc.Save("current.pdf")
}

func (p *pdfReporter) ToTurboEfficiencyMatrix(ctx context.Context, turbo [][]*models.TurboScore) [][]int {
	xSize := len(turbo)
	ySize := len(turbo[0])

	var matrix [][]int
	for i := 0; i < xSize; i++ {
		var line []int
		for j := 0; j < ySize; j++ {
			line = append(line, int(turbo[i][j].Weight))
		}
		matrix = append(matrix, line)
	}

	return matrix
}
