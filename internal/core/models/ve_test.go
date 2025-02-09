package models

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestVE_ToCFM(t *testing.T) {
	t.Run("should convert to correctly", func(t *testing.T) {
		// Arrange
		ve := NewVE(2000, 0.99)

		// Act
		cfm := ve.ToFourCylinderCFM(2)

		// Assert
		assert.Equal(t, ve.RPM, cfm.RPM)
		assert.Equal(t, ve.Percent, cfm.Percent)
		assert.Equal(t, 69.89584451666846, cfm.Flow)
	})
}
