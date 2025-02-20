package models

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_ToMatrix(t *testing.T) {
	t.Run("should map correctly", func(t *testing.T) {
		// Arrange
		xFlowRange := NewRange(0, 500)
		xPixelRange := NewRange(100, 600)

		// Act
		rate := xFlowRange.GetRate(xPixelRange)

		// Act
		assert.Equal(t, 1.0, rate)
	})
}
