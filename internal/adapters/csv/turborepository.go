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

	data := t.getDataWithMinPadding(csvData)
	data = t.getDataWithMaxPadding(data)
	data = t.getDataWithWeights(data)

	for i := 0; i < len(data); i++ {
		for j := 0; j < len(data[i]); j++ {
			fmt.Print(data[i][j], " ")
		}
		fmt.Println()
	}

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

		for j := 0; j < len(data[i]); j++ {
			if data[i][j] == "S" {
				line = append(line, &models.TurboScore{
					Boost:  0,
					Health: 1,
					CFM:    0,
					Weight: 0,
					Out:    true,
				})
			} else if data[i][j] == "C" {
				line = append(line, &models.TurboScore{
					Boost:  1,
					Health: 0,
					CFM:    maxCFM,
					Weight: 0,
					Out:    true,
				})
			} else {
				cfm := models.GetFlowFromBaseRange(data[i][j])
				score := models.GetScoreFromBaseRange(data[i][j])
				line = append(line, &models.TurboScore{
					Health: 0,
					CFM:    cfm,
					Weight: score,
					Out:    false,
				})
			}
		}

		turbo = append(turbo, line)
	}

	for i := 0; i < len(turbo); i++ {
		minWeight := 10000.0
		minWeightIndex := 0
		for j := 0; j < len(turbo[i]); j++ {
			if turbo[i][j].Out {
				continue
			}
			currentWeight := turbo[i][j].Weight
			if currentWeight < minWeight {
				minWeight = currentWeight
				minWeightIndex = j
			}
		}

		maxBottom := 0.0
		maxBottomIndex := 0
		for j := 0; j < len(turbo[i]); j++ {
			if turbo[i][j].Weight != 0 {
				maxBottom = turbo[i][j].Weight
				maxBottomIndex = j
				break
			}
		}

		fmt.Println(minWeight, minWeightIndex, maxBottom, maxBottomIndex)
	}

	for i := 0; i < len(turbo); i++ {
		for j := 0; j < len(turbo[i]); j++ {
			fmt.Print(turbo[i][j])
		}
		fmt.Println()
	}

	return nil, nil
}

func (t *TurboRepository) getDataWithWeights(data [][]string) [][]string {
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

	return data
}

func (t *TurboRepository) getDataWithMaxPadding(data [][]string) [][]string {
	for i := 0; i < len(data); i++ {
		for j := 0; j < len(data[i]); j++ {
			d := data[i][j]
			if d == "" {
				data[i][j] = "C"
			}
		}
	}

	return data
}

func (t *TurboRepository) getDataWithMinPadding(data [][]string) [][]string {
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
