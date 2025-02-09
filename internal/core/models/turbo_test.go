package models

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRange_String(t *testing.T) {
	t.Run("should format correctly", func(t *testing.T) {
		// Arrange
		r := NewRange(100, 200, 1)

		// Act
		s := r.String()

		// Assert
		assert.Equal(t, "[100-200]1", s)
	})
}
