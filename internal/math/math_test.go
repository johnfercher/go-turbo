package math

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCubicInchToCubicFeet(t *testing.T) {
	t.Run("convert correctly", func(t *testing.T) {
		// Assert
		assert.Equal(t, 1.0, CubicInchToCubicFeet(1728))
	})
}

func TestLitersToCubicInch(t *testing.T) {
	t.Run("convert correctly", func(t *testing.T) {
		// Assert
		assert.Equal(t, 122.00001952000314, LitersToCubicInch(2))
	})
}

func TestATMToKg(t *testing.T) {
	t.Run("convert correctly", func(t *testing.T) {
		// Assert
		assert.Equal(t, 1.0, ATMToKg(14.7))
	})
}

func TestKgToATM(t *testing.T) {
	t.Run("convert correctly", func(t *testing.T) {
		// Assert
		assert.Equal(t, 14.7, KgToATM(1))
	})
}

func TestPressureRatio(t *testing.T) {
	t.Run("convert correctly", func(t *testing.T) {
		assert.Equal(t, 2.0, PressureRatio(14.7))
	})
}
