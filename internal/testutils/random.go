package testutils

import (
	"github.com/johnfercher/go-turbo/internal/core/models"
	"math/rand"
)

func GenerateRandomRanges(n int) []models.Range {
	var turboSlices []models.Range

	bottom := 100.0
	for i := 0; i < n; i++ {
		offset := rand.Intn(100)
		turboSlices = append(turboSlices, models.Range{
			Min:   bottom,
			Max:   bottom + float64(offset),
			Score: float64(i),
		})
		bottom += float64(offset)
	}

	for i := 0; i < n; i++ {
		ra := rand.Intn(n)
		turboSlices[i], turboSlices[ra] = turboSlices[ra], turboSlices[i]
	}

	return turboSlices
}
