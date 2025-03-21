package models

import (
	"fmt"
)

type Chart struct {
	XPixelRange *Range
	XFlowRange  *Range
	Turbo       map[float64][]Point
}

func NewChart(xPixelRange, xFlowRange *Range) *Chart {
	return &Chart{
		XPixelRange: xPixelRange,
		XFlowRange:  xFlowRange,
	}
}

func (c *Chart) AddTurbo(turbo map[float64][]Point) *Chart {
	c.Turbo = turbo
	return c
}

func (c *Chart) ToMatrix() [][]string {
	var matrix [][]string
	for pressure, flow := range c.Turbo {
		var line []string
		line = append(line, fmt.Sprintf("%.1f", pressure))
		for _, point := range flow {
			pixelX := float64(point.X) - c.XPixelRange.Begin
			rate := c.XPixelRange.GetRate(c.XFlowRange)
			line = append(line, fmt.Sprintf("%.0f", pixelX*rate))
		}
		matrix = append(matrix, line)
	}
	return matrix
}

func (c *Chart) String() string {
	return fmt.Sprintf("pixel: %s - flow: %s", c.XPixelRange.String(), c.XFlowRange.String())
}
