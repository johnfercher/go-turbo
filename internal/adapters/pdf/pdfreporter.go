package pdf

import (
	"context"
	"github.com/johnfercher/go-turbo/internal/core/consts"
	"github.com/johnfercher/go-turbo/internal/core/models"
	"github.com/johnfercher/maroto/v2"
	"github.com/johnfercher/maroto/v2/pkg/components/chart"
	"github.com/johnfercher/maroto/v2/pkg/config"
	"github.com/johnfercher/maroto/v2/pkg/consts/orientation"
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
		WithPageSize(pagesize.A1).
		WithOrientation(orientation.Horizontal).
		Build()

	m := maroto.New(cfg)

	m.AddRows(chart.NewHeatMapRow(300, "Efficiency", matrix, props.HeatMap{
		TransparentValues: []int{0},
		InvertScale:       true,
	}))

	doc, err := m.Generate()
	if err != nil {
		return err
	}

	return doc.Save("current.pdf")
}

func (p *pdfReporter) ToTurboEfficiencyMatrix(ctx context.Context, turbo [][]*models.TurboScore) [][]int {
	var matrix [][]int
	for i := 0; i < len(turbo); i++ {
		var line []int
		for j := 0; j < len(turbo[i]); j++ {
			line = append(line, int(turbo[i][j].Weight))
		}
		matrix = append(matrix, line)
	}

	return matrix
}
