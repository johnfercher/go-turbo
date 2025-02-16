package csv

import (
	"context"
	"github.com/johnfercher/go-turbo/internal/core/consts"
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
	turbo = matrix.InterpolateLimits(turbo, data)
	turbo = matrix.NormalizeWeights(turbo)

	//turbo = t.computeBoost(turbo)
	//turbo = t.computeHealth(turbo)

	//matrix.PrintBoost(turbo)
	//models.PrintWeight(turbo)
	//models.PrintCFM(turbo)
	//models.PrintHealth(turbo)
	//models.PrintSurge(turbo)
	//models.PrintChoke(turbo)

	return models.NewTurbo(turboFile, turbo)
}

func (t *TurboRepository) computeHealth(turbo [][]*models.TurboScore) [][]*models.TurboScore {
	for i := 0; i < len(turbo); i++ {
		firstMaxWeight := 0.0
		firstMaxWeightIndex := 0
		beginChokeIndex := 0
		for j := 0; j < len(turbo[i]); j++ {
			if turbo[i][j].Choke && beginChokeIndex == 0 {
				beginChokeIndex = j
			}
			if turbo[i][j].Surge || turbo[i][j].Choke {
				continue
			}
			currentWeight := turbo[i][j].Weight
			if currentWeight > firstMaxWeight {
				firstMaxWeight = currentWeight
				firstMaxWeightIndex = j
			}
		}

		entireChokeSurge := true
		for j := 0; j < len(turbo[i]); j++ {
			if !turbo[i][j].Surge && !turbo[i][j].Choke {
				entireChokeSurge = false
				break
			}
		}

		if entireChokeSurge {
			continue
		}

		rateChoke := 1 / (float64(beginChokeIndex) - float64(firstMaxWeightIndex+1))
		chokeIndex := 0.0
		for j := 1; j < len(turbo[i]); j++ {
			if j < firstMaxWeightIndex+1 {
				if !turbo[i][j].Surge && !turbo[i][j].Choke {
					turbo[i][j].Health = 1
				}
			} else {
				if !turbo[i][j].Choke {
					turbo[i][j].Health = 1 - (rateChoke * chokeIndex)
					chokeIndex++
				}
			}
		}
	}

	return turbo
}

func (t *TurboRepository) computeBoost(turbo [][]*models.TurboScore) [][]*models.TurboScore {
	for i := 0; i < len(turbo); i++ {
		firstMaxWeight := 0.0
		firstMaxWeightIndex := 0
		for j := 0; j < len(turbo[i]); j++ {
			if turbo[i][j].Surge || turbo[i][j].Choke {
				continue
			}
			currentWeight := turbo[i][j].Weight
			if currentWeight > firstMaxWeight {
				firstMaxWeight = currentWeight
				firstMaxWeightIndex = j
			}
		}

		entireChokeSurge := true
		for j := 0; j < len(turbo[i]); j++ {
			if !turbo[i][j].Surge && !turbo[i][j].Choke {
				entireChokeSurge = false
				break
			}
		}

		if entireChokeSurge {
			continue
		}

		beginValidIndex := 0
		for j := 0; j < len(turbo[i]); j++ {
			if !turbo[i][j].Surge {
				beginValidIndex = j
				break
			}
		}

		maxBoost := consts.Boosts[i]
		rateBoost := float64(beginValidIndex) / float64(firstMaxWeightIndex)
		stepBoost := rateBoost * maxBoost
		for j := 1; j < len(turbo[i]); j++ {
			if j <= firstMaxWeightIndex {
				if !turbo[i][j].Surge && !turbo[i][j].Choke {
					turbo[i][j].Boost = stepBoost * float64(j)
				}
			} else {
				turbo[i][j].Boost = maxBoost
			}
		}
	}

	return turbo
}

func (t *TurboRepository) buildTurboMatrixWithSurgeAndChoke(data [][]string) [][]*models.TurboScore {
	var turbo [][]*models.TurboScore
	for i := 0; i < len(data); i++ {
		var line []*models.TurboScore

		maxCFM := 0.0
		for j := 0; j < len(data[i]); j++ {
			if data[i][j] != "S" && data[i][j] != "C" {
				cfm := models.GetFlowFromBaseRange(data[i][j])
				if cfm > maxCFM {
					maxCFM = cfm
				}
			}
		}

		// Define base surge/choke
		for j := 0; j < len(data[i]); j++ {
			if data[i][j] == "S" {
				line = append(line, &models.TurboScore{
					Boost:  0,
					Health: 1,
					CFM:    0,
					Weight: 0,
					Surge:  true,
					Choke:  false,
				})
			} else if data[i][j] == "C" {
				line = append(line, &models.TurboScore{
					Boost:  0,
					Health: 0,
					CFM:    maxCFM,
					Weight: 0,
					Surge:  false,
					Choke:  true,
				})
			} else {
				cfm := models.GetFlowFromBaseRange(data[i][j])
				score := models.GetScoreFromBaseRange(data[i][j])
				line = append(line, &models.TurboScore{
					Health: 0,
					CFM:    cfm,
					Weight: score,
					Surge:  false,
					Choke:  false,
				})
			}
		}

		turbo = append(turbo, line)
	}
	return turbo
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
