package matrix

import (
	"fmt"
	"github.com/johnfercher/go-turbo/internal/core/models"
)

func InitMatrix(maxBoost int, maxFlow int) [][]*models.TurboScore {
	maxFlow += 50
	maxBoost += 50

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

	fmt.Println(len(matrix))
	fmt.Println(len(matrix[0]))
	return matrix
}

func Val(matrix [][]*models.TurboScore, turbo [][]string) [][]*models.TurboScore {
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

func IsSurgeOrChoke(s string) bool {
	return IsSurge(s) || IsChoke(s)
}

func IsSurge(s string) bool {
	return s == "S"
}

func IsChoke(s string) bool {
	return s == "C"
}

func Print(matrix [][]string) {
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			fmt.Print(matrix[i][j], " ")
		}
		fmt.Println()
	}
}

func PrintBoost(turboScore [][]*models.TurboScore) {
	fmt.Println("Boost:")
	for i := 0; i < len(turboScore); i++ {
		for j := 0; j < len(turboScore[i]); j++ {
			fmt.Print(turboScore[i][j].StringBoost(), " ")
		}
		fmt.Println()
	}
}

func PrintWeight(turboScore [][]*models.TurboScore, avoidZeros ...bool) {
	av := false
	if len(avoidZeros) > 0 {
		av = avoidZeros[0]
	}

	fmt.Println("Weight:")
	for i := 0; i < len(turboScore); i++ {
		for j := 0; j < len(turboScore[i]); j++ {
			if av {
				if turboScore[i][j].Weight > 0 {
					fmt.Print(turboScore[i][j].StringWeight(), " ")
				} else {
					fmt.Print("  ")
				}
			} else {
				fmt.Print(turboScore[i][j].StringWeight(), " ")
			}
		}
		fmt.Println()
	}
}

func PrintCFM(turboScore [][]*models.TurboScore) {
	fmt.Println("CFM:")
	for i := 0; i < len(turboScore); i++ {
		for j := 0; j < len(turboScore[i]); j++ {
			fmt.Print(turboScore[i][j].StringCFM(), " ")
		}
		fmt.Println()
	}
}

func PrintHealth(turboScore [][]*models.TurboScore) {
	fmt.Println("Health:")
	for i := 0; i < len(turboScore); i++ {
		for j := 0; j < len(turboScore[i]); j++ {
			fmt.Print(turboScore[i][j].StringHealth(), " ")
		}
		fmt.Println()
	}
}

func PrintSurge(turboScore [][]*models.TurboScore) {
	fmt.Println("Surge:")
	for i := 0; i < len(turboScore); i++ {
		for j := 0; j < len(turboScore[i]); j++ {
			fmt.Print(turboScore[i][j].StringSurge(), " ")
		}
		fmt.Println()
	}
}

func PrintChoke(turboScore [][]*models.TurboScore) {
	fmt.Println("Choke:")
	for i := 0; i < len(turboScore); i++ {
		for j := 0; j < len(turboScore[i]); j++ {
			fmt.Print(turboScore[i][j].StringChoke(), " ")
		}
		fmt.Println()
	}
}
