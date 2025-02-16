package testutils

import (
	"github.com/johnfercher/go-turbo/internal/matrix"
	"math/rand"
)

func RandomTurboMatrix(maxBoost, maxFlow int) [][]float64 {
	m := matrix.InitMatrix(maxBoost, maxFlow)
	maxRandom := rand.Intn(10)

	for i := 0; i < maxBoost; i++ {
		for j := 0; j < maxFlow; j++ {
			m[i][j] = float64(rand.Intn(maxRandom))
		}
	}

	return m
}

func IncrementalTurboMatrix(maxBoost, maxFlow int) [][]float64 {
	m := matrix.InitMatrix(maxBoost, maxFlow)

	for i := 0; i < maxFlow; i++ {
		for j := 0; j < maxBoost; j++ {
			m[i][j] = float64(i + j)
		}
	}

	return m
}
