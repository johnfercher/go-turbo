package csv

import (
	"bytes"
	"context"
	"encoding/csv"
	"fmt"
	"github.com/johnfercher/go-turbo/internal/core/models"
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

	reader := csv.NewReader(bytes.NewBuffer(b))
	reader.FieldsPerRecord = -1
	csvData, err := reader.ReadAll()
	if err != nil {
		panic(err)
	}

	data := t.addSurgeFirstColumnSurge(csvData)
	data = t.addChokeColumns(data)
	data = t.addWeightsOutsideSurgeChoke(data)
	turbo := t.buildTurboMatrixWithSurgeAndChoke(data)
	turbo = t.normalizeWeights(turbo)
	turbo = t.computeBoost(turbo)
	turbo = t.computeHealth(turbo)

	for i := 0; i < len(turbo); i++ {
		for j := 0; j < len(turbo[i]); j++ {

		}
	}

	//models.PrintBoost(turbo)
	models.PrintWeight(turbo)
	//models.PrintCFM(turbo)
	models.PrintHealth(turbo)
	//models.PrintSurge(turbo)
	//models.PrintChoke(turbo)

	return nil, nil
}

func (t *TurboRepository) normalizeWeights(turbo [][]*models.TurboScore) [][]*models.TurboScore {
	// Add all max to invalid points
	maxWeight := 0.0
	for i := 0; i < len(turbo); i++ {
		for j := 0; j < len(turbo[i]); j++ {
			if turbo[i][j].Weight > maxWeight {
				maxWeight = turbo[i][j].Weight
			}
		}
	}

	for i := 0; i < len(turbo); i++ {
		for j := 0; j < len(turbo[i]); j++ {
			if turbo[i][j].Weight == 0 {
				turbo[i][j].Weight = maxWeight + 1
			}
		}
	}

	base := 1 / maxWeight
	for i := 0; i < len(turbo); i++ {
		for j := 0; j < len(turbo[i]); j++ {
			turbo[i][j].Weight = 1 + base - (turbo[i][j].Weight / maxWeight)
		}
	}

	return turbo
}

func (t *TurboRepository) computeHealth(turbo [][]*models.TurboScore) [][]*models.TurboScore {
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

		for j := 1; j <= firstMaxWeightIndex+1; j++ {
			if !turbo[i][j].Surge && !turbo[i][j].Choke {
				turbo[i][j].Health = 1
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

		for j := 1; j <= firstMaxWeightIndex+1; j++ {
			if !turbo[i][j].Surge && !turbo[i][j].Choke {
				turbo[i][j].Boost = turbo[i][j].Weight
			}
		}

		for j := firstMaxWeightIndex + 2; j < len(turbo[i]); j++ {
			if !turbo[i][j].Surge {
				turbo[i][j].Boost = 1.0
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
					Boost:  1,
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

func (t *TurboRepository) addWeightsOutsideSurgeChoke(data [][]string) [][]string {
	for i := 0; i < len(data); i++ {
		found := 0
		max := 0
		base := 0.0
		for j := 0; j < len(data[i]); j++ {
			if models.IsBaseRange(data[i][j]) {
				base = models.GetScoreFromBaseRange(data[i][j])
				break
			} else {
				found++
			}
		}
		max = found
		for j := 0; j < max; j++ {
			if data[i][j] != "S" && data[i][j] != "C" {
				data[i][j] = fmt.Sprintf("%s-%d", data[i][j], found+int(base))
			}
			found--
		}
		offset := 1
		for j := max + 2; j < len(data[i]); j++ {
			if data[i][j] != "S" && data[i][j] != "C" {
				data[i][j] = fmt.Sprintf("%s-%d", data[i][j], offset+int(base))
			}
			offset++
		}
	}

	/*for i := 0; i < len(data); i++ {
		for j := 0; j < len(data[i]); j++ {
			fmt.Print(data[i][j], " ")
		}
		fmt.Println()
	}*/

	return data
}

func (t *TurboRepository) addChokeColumns(data [][]string) [][]string {
	for i := 0; i < len(data); i++ {
		allEmpty := true
		for j := 0; j < len(data[i]); j++ {
			if data[i][j] != "" && data[i][j] != "S" {
				allEmpty = false
				break
			}
		}

		if allEmpty {
			for j := 0; j < len(data[i]); j++ {
				d := data[i][j]
				if d == "" {
					data[i][j] = "S"
				}
			}
		} else {
			for j := 0; j < len(data[i]); j++ {
				d := data[i][j]
				if d == "" {
					data[i][j] = "C"
				}
			}
		}

	}

	return data
}

func (t *TurboRepository) addSurgeFirstColumnSurge(data [][]string) [][]string {
	var minPadded [][]string
	for i := 1; i < len(data); i++ {
		var line []string
		line = append(line, "S")
		for j := 1; j < len(data[i]); j++ {
			line = append(line, data[i][j])
		}
		minPadded = append(minPadded, line)
	}
	return minPadded
}
