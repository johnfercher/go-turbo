package math

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetRate(t *testing.T) {
	t.Run("when scale1 is the same as scale2 with same offset", func(t *testing.T) {
		// Arrange
		minXScale1 := 0.0
		maxXScale1 := 1000.0

		minXScale2 := 0.0
		maxXScale2 := 1000.0

		// Act
		rate := GetRate(minXScale1, maxXScale1, minXScale2, maxXScale2)

		// Assert
		assert.Equal(t, 1.0, rate)
	})
	t.Run("when scale1 is the same as scale2 with different offset", func(t *testing.T) {
		// Arrange
		minXScale1 := 100.0
		maxXScale1 := 1100.0

		minXScale2 := 100.0
		maxXScale2 := 1100.0

		// Act
		rate := GetRate(minXScale1, maxXScale1, minXScale2, maxXScale2)

		// Assert
		assert.Equal(t, 1.0, rate)
	})
	t.Run("when scale1 is half then scale2 with same offset", func(t *testing.T) {
		// Arrange
		minXScale1 := 0.0
		maxXScale1 := 500.0

		minXScale2 := 0.0
		maxXScale2 := 1000.0

		// Act
		rate := GetRate(minXScale1, maxXScale1, minXScale2, maxXScale2)

		// Assert
		assert.Equal(t, 2.0, rate)
	})
	t.Run("when scale1 is half then scale2 with different offset", func(t *testing.T) {
		// Arrange
		minXScale1 := 100.0
		maxXScale1 := 600.0

		minXScale2 := 100.0
		maxXScale2 := 1100.0

		// Act
		rate := GetRate(minXScale1, maxXScale1, minXScale2, maxXScale2)

		// Assert
		assert.Equal(t, 2.0, rate)
	})
	t.Run("when scale1 is double then scale2 with same offset", func(t *testing.T) {
		// Arrange
		minXScale1 := 0.0
		maxXScale1 := 1000.0

		minXScale2 := 0.0
		maxXScale2 := 500.0

		// Act
		rate := GetRate(minXScale1, maxXScale1, minXScale2, maxXScale2)

		// Assert
		assert.Equal(t, 0.5, rate)
	})
	t.Run("when scale1 is double then scale2 with different offset", func(t *testing.T) {
		// Arrange
		minXScale1 := 100.0
		maxXScale1 := 1100.0

		minXScale2 := 100.0
		maxXScale2 := 600.0

		// Act
		rate := GetRate(minXScale1, maxXScale1, minXScale2, maxXScale2)

		// Assert
		assert.Equal(t, 0.5, rate)
	})
}
