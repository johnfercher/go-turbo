package csv

import (
	"fmt"
	"github.com/johnfercher/go-turbo/internal/core/models"
	"testing"
)

func TestTurbo(t *testing.T) {
	t.Run("", func(t *testing.T) {
		// Arrange
		xFlowRange := models.NewRange(0, 675)
		xPixelRange := models.NewRange(125, 850)

		turbo := make(map[float64][]models.Point)
		turbo[1.0] = []models.Point{
			{
				X: 211,
			},
			{
				X: 229,
			},
			{
				X: 259,
			},
			{
				X: 288,
			},
			{
				X: 318,
			},
			{
				X: 335,
			},
			{
				X: 401,
			},
			{
				X: 422,
			},
			{
				X: 443,
			},
			{
				X: 479,
			},
			{
				X: 528,
			},
		}

		c := models.NewChart(xPixelRange, xFlowRange).AddTurbo(turbo)

		merged := c.ToMatrix()

		// Act
		s := Turbo(merged)

		// Assert
		fmt.Println(s)
	})
}
