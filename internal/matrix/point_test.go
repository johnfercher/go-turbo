package matrix_test

import (
	"github.com/johnfercher/go-turbo/internal/matrix"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPoints_Interpolate(t *testing.T) {
	t.Run("avoid same x interpolation, 1", func(t *testing.T) {
		// Arrange
		points := matrix.Points{
			matrix.NewPoint(0, 0, 10),
			matrix.NewPoint(0, 5, 5),
		}

		// Act
		newPoints := points.Interpolate(50)

		// Assert
		assert.Equal(t, len(newPoints), len(points))
	})
	t.Run("avoid same x interpolation, 2", func(t *testing.T) {
		// Arrange
		points := matrix.Points{
			matrix.NewPoint(0, 0, 10),
			matrix.NewPoint(0, 50, 5),
			matrix.NewPoint(50, 0, 3),
			matrix.NewPoint(50, 0, 2),
		}

		// Act
		newPoints := points.Interpolate(50)

		// Assert
		assert.Equal(t, len(newPoints), len(points))
	})
}
