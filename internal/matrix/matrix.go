package matrix

import (
	"github.com/johnfercher/go-turbo/internal/core/consts"
	"github.com/johnfercher/go-turbo/internal/core/models"
)

func BuildEmptyTurbo() [][]*models.TurboScore {
	var matrix [][]*models.TurboScore
	for i := 0; i < consts.TurboMaxLines; i++ {
		var arr []*models.TurboScore
		for j := 0; j < consts.TurboMaxColumns; j++ {
			arr = append(arr, &models.TurboScore{})
		}
		matrix = append(matrix, arr)
	}

	return matrix
}
