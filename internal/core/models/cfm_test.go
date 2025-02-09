package models

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCFM_AddBoost(t *testing.T) {
	t.Run("should add a new boost", func(t *testing.T) {
		// Arrange
		ve := NewVE(2000, 0.99)
		cfm := ve.ToFourCylinderCFM(2)

		// Act
		cfm1KG := cfm.AddBoostKg(14.7)

		// Assert
		assert.Equal(t, cfm.RPM, cfm1KG.RPM)
		assert.Equal(t, cfm.Percent, cfm1KG.Percent)
		assert.Equal(t, 139.79168903333692, cfm1KG.Flow)

	})
}
