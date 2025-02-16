package matrix

import (
	"fmt"
	"github.com/johnfercher/go-turbo/internal/core/models"
)

func InitMatrix(maxBoost int, maxFlow int) [][]*models.TurboScore {
	var matrix [][]*models.TurboScore
	for i := 0; i < maxFlow; i++ {
		var flow []*models.TurboScore
		for j := 0; j < maxBoost; j++ {
			flow = append(flow, &models.TurboScore{
				Boost: float64(i),
				CFM:   float64(j),
			})
		}
		matrix = append(matrix, flow)
	}

	return matrix
}

func InterpolateLimits(matrix [][]*models.TurboScore, turbo [][]string) [][]*models.TurboScore {
	step := 20.0

	// Find marks
	var halfBottomPoints Points
	var halfTopPoints Points

	for i := 0; i < len(turbo); i++ {
		halfBetterIndex := 0
		better := 1000.0
		for j := 0; j < len(turbo[i]); j++ {
			turboString := turbo[i][j]
			if !IsSurgeOrChoke(turboString) {
				score := models.GetScoreFromBaseRange(turboString)
				if score < better {
					better = score
					halfBetterIndex = j
				}
			}
		}

		for j := 0; j <= halfBetterIndex; j++ {
			turboString := turbo[i][j]
			if !IsSurgeOrChoke(turboString) {
				score := models.GetScoreFromBaseRange(turboString)
				flow := models.GetFlowFromBaseRange(turboString)

				stepX := float64(i) * step

				halfBottomPoints = append(halfBottomPoints, Point{
					X:     stepX,
					Y:     flow,
					Value: score,
				})
			}
		}
		for j := halfBetterIndex + 1; j < len(turbo[i]); j++ {
			turboString := turbo[i][j]
			if !IsSurgeOrChoke(turboString) {
				score := models.GetScoreFromBaseRange(turboString)
				flow := models.GetFlowFromBaseRange(turboString)

				stepX := float64(i) * step

				halfTopPoints = append(halfTopPoints, Point{
					X:     stepX,
					Y:     flow,
					Value: score,
				})
			}
		}
	}

	halfBottomPoints = halfBottomPoints.Interpolate()
	halfTopPoints = halfTopPoints.Interpolate()

	var points Points
	points = append(points, halfBottomPoints...)
	points = append(points, halfTopPoints...)

	// Fill marks
	for _, p := range points {
		x, y, w := p.X, p.Y, p.Value
		matrix[int(y)][int(x)].Weight = w
	}

	return matrix
}

func NormalizeWeights(turbo [][]*models.TurboScore) [][]*models.TurboScore {
	// Add all max to invalid points
	maxWeight := 0.0
	for i := 0; i < len(turbo); i++ {
		for j := 0; j < len(turbo[i]); j++ {
			if turbo[i][j].Weight > maxWeight {
				maxWeight = turbo[i][j].Weight
			}
		}
	}

	fmt.Println(maxWeight)

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
			turbo[i][j].Weight = (1 + base - (turbo[i][j].Weight / maxWeight)) * 100.0
		}
	}

	return turbo
}
