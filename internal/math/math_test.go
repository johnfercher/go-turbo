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

func TestDistance(t *testing.T) {
	t.Run("find correct distance", func(t *testing.T) {
		assert.Equal(t, 5.0, Distance(0, 0, 3, 4))
	})
}

func TestAngleBetween(t *testing.T) {
	t.Run("find correct angle", func(t *testing.T) {
		assert.Equal(t, 45.0, AngleBetween(0, 0, 3, 3))
		assert.Equal(t, -135.0, AngleBetween(3, 3, 0, 0))
	})
}
