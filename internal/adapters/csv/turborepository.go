package csv

import (
	"context"
	"fmt"
	"github.com/johnfercher/go-turbo/internal/core/models"
	"github.com/johnfercher/go-turbo/internal/matrix"
	"os"
	"strings"
)

type TurboRepository struct {
}

func NewTurboRepository() *TurboRepository {
	return &TurboRepository{}
}

func (t *TurboRepository) Get(ctx context.Context, turboFile string) (*models.Turbo, error) {
	fileName := TurboFilePath(turboFile)

	b, err := os.ReadFile(fileName)
	if err != nil {
		return nil, err
	}

	data := Load(b)

	maxFlow := t.getMaxValue(data)
	maxBoost := 200
	padding := 50
	maxFlow += padding
	maxBoost += padding

	turbo := matrix.InitMatrix(200, maxFlow)
	turbo = matrix.InterpolateLimitsY(turbo, data)
	turbo = matrix.NormalizeWeights(turbo)
	//turbo = matrix.InterpolateCurves(turbo)
	turbo = matrix.InterpolateX(turbo)

	return models.NewTurbo(turboFile, turbo), nil
}

func (t *TurboRepository) Save(ctx context.Context, name string, chart *models.Chart) error {
	s := Turbo(chart.ToMatrix())
	return os.WriteFile(TurboFilePath(name), []byte(s), os.ModePerm)
}

func (t *TurboRepository) getMaxValue(data [][]string) int {
	max := 0.0
	for i := 0; i < len(data); i++ {
		for j := 1; j < len(data[i]); j++ {
			if models.IsBaseRange(data[i][j]) {
				flow := models.GetFlowFromBaseRange(data[i][j])
				if flow > max {
					max = flow
				}
			}
		}
	}

	return int(max)
}

func (t *TurboRepository) getMaxBoost(data [][]string) int {
	max := 0
	step := 0.2
	for i := 0; i < len(data); i++ {
		for j := 1; j < len(data[i]); j++ {
			if models.IsBaseRange(data[i][j]) {
				max = i
				break
			}
		}
	}

	return int(100 * float64(max) * step)
}

func TurboFilePath(name string) string {
	return fmt.Sprintf("data/turbo/%s.csv", name)
}

func Turbo(turbo [][]string) string {
	s := "kg,col1,col2,col3,col4,col5,col6,col7,col8,col9,col10,col11,col12,col13,col14,col15,col16\n"

	for i := 0; i < len(turbo); i++ {
		var line []string
		for j := 0; j < len(turbo[i]); j++ {
			line = append(line, turbo[i][j])
		}
		s += strings.Join(line, ",") + "\n"
	}

	return s
}
