package matrix

import (
	"fmt"
	"github.com/johnfercher/go-turbo/internal/core/models"
)

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
