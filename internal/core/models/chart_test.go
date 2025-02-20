package models

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestChart_ToMatrix(t *testing.T) {
	t.Run("same scale pixel/flow, same offset", func(t *testing.T) {
		// Arrange
		xFlowRange := NewRange(0, 500)
		xPixelRange := NewRange(0, 500)

		turbo := make(map[float64][]Point)
		turbo[1.0] = []Point{
			{
				X: 0,
			},
			{
				X: 250,
			},
			{
				X: 500,
			},
		}

		c := NewChart(xPixelRange, xFlowRange).AddTurbo(turbo)

		// Act
		matrix := c.ToMatrix()

		// Assert
		assert.Equal(t, "1.00", matrix[0][0])
		assert.Equal(t, "0", matrix[0][1])
		assert.Equal(t, "250", matrix[0][2])
		assert.Equal(t, "500", matrix[0][3])
	})
	t.Run("same scale pixel/flow, different offset", func(t *testing.T) {
		// Arrange
		xFlowRange := NewRange(0, 500)
		xPixelRange := NewRange(100, 600)

		turbo := make(map[float64][]Point)
		turbo[1.0] = []Point{
			{
				X: 100,
			},
			{
				X: 350,
			},
			{
				X: 600,
			},
		}

		c := NewChart(xPixelRange, xFlowRange).AddTurbo(turbo)

		// Act
		matrix := c.ToMatrix()

		// Assert
		assert.Equal(t, "1.00", matrix[0][0])
		assert.Equal(t, "0", matrix[0][1])
		assert.Equal(t, "250", matrix[0][2])
		assert.Equal(t, "500", matrix[0][3])
	})
	t.Run("pixel double than flow, same offset", func(t *testing.T) {
		// Arrange
		xFlowRange := NewRange(0, 500)
		xPixelRange := NewRange(0, 1000)

		turbo := make(map[float64][]Point)
		turbo[1.0] = []Point{
			{
				X: 0,
			},
			{
				X: 500,
			},
			{
				X: 1000,
			},
		}

		c := NewChart(xPixelRange, xFlowRange).AddTurbo(turbo)

		// Act
		matrix := c.ToMatrix()

		// Assert
		assert.Equal(t, "1.00", matrix[0][0])
		assert.Equal(t, "0", matrix[0][1])
		assert.Equal(t, "250", matrix[0][2])
		assert.Equal(t, "500", matrix[0][3])
	})
	t.Run("pixel double than flow, different offset", func(t *testing.T) {
		// Arrange
		xFlowRange := NewRange(0, 500)
		xPixelRange := NewRange(100, 1100)

		turbo := make(map[float64][]Point)
		turbo[1.0] = []Point{
			{
				X: 100,
			},
			{
				X: 600,
			},
			{
				X: 1100,
			},
		}

		c := NewChart(xPixelRange, xFlowRange).AddTurbo(turbo)

		// Act
		matrix := c.ToMatrix()

		// Assert
		assert.Equal(t, "1.00", matrix[0][0])
		assert.Equal(t, "0", matrix[0][1])
		assert.Equal(t, "250", matrix[0][2])
		assert.Equal(t, "500", matrix[0][3])
	})
	t.Run("flow double than pixel, same offset", func(t *testing.T) {
		// Arrange
		xFlowRange := NewRange(0, 1000)
		xPixelRange := NewRange(0, 500)

		turbo := make(map[float64][]Point)
		turbo[1.0] = []Point{
			{
				X: 0,
			},
			{
				X: 250,
			},
			{
				X: 500,
			},
		}

		c := NewChart(xPixelRange, xFlowRange).AddTurbo(turbo)

		// Act
		matrix := c.ToMatrix()

		// Assert
		assert.Equal(t, "1.00", matrix[0][0])
		assert.Equal(t, "0", matrix[0][1])
		assert.Equal(t, "500", matrix[0][2])
		assert.Equal(t, "1000", matrix[0][3])
	})
	t.Run("flow double than pixel, different offset", func(t *testing.T) {
		// Arrange
		xFlowRange := NewRange(0, 1000)
		xPixelRange := NewRange(100, 600)

		turbo := make(map[float64][]Point)
		turbo[1.0] = []Point{
			{
				X: 100,
			},
			{
				X: 350,
			},
			{
				X: 600,
			},
		}

		c := NewChart(xPixelRange, xFlowRange).AddTurbo(turbo)

		// Act
		matrix := c.ToMatrix()

		// Assert
		assert.Equal(t, "1.00", matrix[0][0])
		assert.Equal(t, "0", matrix[0][1])
		assert.Equal(t, "500", matrix[0][2])
		assert.Equal(t, "1000", matrix[0][3])
	})
}
