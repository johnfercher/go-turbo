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

func TestTurboSlice_String(t *testing.T) {
	t.Run("should format correctly", func(t *testing.T) {
		// Arrange
		slice := NewTurboSlice(1.0,
			NewRange(100, 200, 3),
			NewRange(200, 300, 2),
			NewRange(300, 400, 1),
			NewRange(400, 500, 2),
			NewRange(500, 600, 3),
		)

		// Act
		s := slice.String()

		// Assert
		assert.Equal(t, "PSI: 1.0 [100-200]3 [200-300]2 [300-400]1 [400-500]2 [500-600]3", s)
	})
}

func TestTurbo_String(t *testing.T) {
	t.Run("should format correctly", func(t *testing.T) {
		// Arrange
		turbo := NewTurbo(
			NewTurboSlice(1.0,
				NewRange(100, 200, 3),
				NewRange(200, 300, 2),
				NewRange(300, 400, 1),
				NewRange(400, 500, 2),
				NewRange(500, 600, 3),
			),
			NewTurboSlice(1.2,
				NewRange(120, 220, 3),
				NewRange(220, 320, 2),
				NewRange(320, 420, 1),
				NewRange(420, 520, 2),
				NewRange(520, 620, 3),
			),
		)

		// Act
		s := turbo.String()

		// Assert
		assert.Equal(t, "PSI: 1.0 [100-200]3 [200-300]2 [300-400]1 [400-500]2 [500-600]3\nPSI: 1.2 [120-220]3 [220-320]2 [320-420]1 [420-520]2 [520-620]3\n", s)
	})
}
