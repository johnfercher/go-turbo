package models

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTransmission_GetSpeed(t *testing.T) {
	t.Run("", func(t *testing.T) {
		// Arrange
		transmission := &Transmission{
			Name:    "sti",
			Final:   3900,
			Reverse: 3545,
			Gears: []float64{
				3.636,
				2.375,
				1.761,
				1.346,
				1.062,
				0.842,
			},
		}

		// Act
		speed := transmission.GetSpeed(7000, 17, 45, 6)

		// Assert
		assert.Equal(t, 254.8699999952249, speed)
	})
}
