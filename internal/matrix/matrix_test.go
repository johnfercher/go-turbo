package matrix_test

import (
	"github.com/johnfercher/go-turbo/internal/matrix"
	"github.com/johnfercher/go-turbo/internal/testutils"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNormalizeWeights(t *testing.T) {
	t.Run("normalize weights", func(t *testing.T) {
		// Arrange
		m := testutils.IncrementalTurboMatrix(3, 3)

		matrix.PrintWeight(m)

		// Act
		m = matrix.NormalizeWeights(m)

		matrix.PrintWeight(m)

		// Assert
		assert.Equal(t, 0.0, m[0][0].Weight)
		assert.Equal(t, 100.0, m[0][1].Weight)
		assert.Equal(t, 75.0, m[0][2].Weight)
		assert.Equal(t, 100.0, m[1][0].Weight)
		assert.Equal(t, 75.0, m[1][1].Weight)
		assert.Equal(t, 50.0, m[1][2].Weight)
		assert.Equal(t, 75.0, m[2][0].Weight)
		assert.Equal(t, 50.0, m[2][1].Weight)
		assert.Equal(t, 25.0, m[2][2].Weight)
	})
}
