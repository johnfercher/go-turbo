package models

import (
	"fmt"
	"github.com/johnfercher/go-turbo/internal/core/math"
)

type CFM struct {
	RPM     float64 `csv:"rpm"`
	Percent float64 `csv:"percent"`
	Flow    float64 `csv:"flow"`
}

func (c *CFM) String() string {
	return fmt.Sprintf("RPM(%.0f)/VE(%.2f)/CFM(%.2f)", c.RPM, c.Percent, c.Flow)
}

func (c *CFM) AddBoostKg(kg float64) *CFM {
	return &CFM{
		RPM:     c.RPM,
		Percent: c.Percent,
		Flow:    c.Flow * math.PressureRatio(math.KgToATM(kg)),
	}
}
