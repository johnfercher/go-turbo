package csv

import (
	"context"
	"github.com/johnfercher/go-turbo/internal/core/models"
	"github.com/johnfercher/go-turbo/internal/matrix"
	"os"
)

type TurboRepository struct {
}

func NewTurboRepository() *TurboRepository {
	return &TurboRepository{}
}

func (t *TurboRepository) Get(ctx context.Context, turboFile string) (*models.Turbo, error) {
	b, err := os.ReadFile("data/turbo/" + turboFile + ".csv")
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
	turbo = matrix.InterpolateX(turbo)

	/*var s string
	xSize := len(turbo)
	ySize := len(turbo[0])
	for i := 0; i < ySize; i++ {
		for j := 0; j < xSize; j++ {
			s += fmt.Sprintf("%.0f ", turbo[j][i])
		}
		s += fmt.Sprintf("\n")
	}

	os.WriteFile("save.csv", []byte(s), 0644)*/

	return models.NewTurbo(turboFile, turbo), nil
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
