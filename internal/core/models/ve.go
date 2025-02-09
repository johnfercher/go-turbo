package models

import (
	"fmt"
	"github.com/johnfercher/go-turbo/internal/core/math"
)

type VE struct {
	RPM     float64 `csv:"rpm"`
	Percent float64 `csv:"percent"`
}

func NewVE(rpm float64, percent float64) *VE {
	return &VE{
		RPM:     rpm,
		Percent: percent,
	}
}

func (v *VE) String() string {
	return fmt.Sprintf("RPM(%.0f)/VE(%.2f)", v.RPM, v.Percent)
}

func (v *VE) ToFourCylinderCFM(engineLiters float64) *CFM {
	return &CFM{
		RPM:     v.RPM,
		Percent: v.Percent,
		Flow:    (math.LitersToCubicInch(engineLiters) * math.FourCylinderConst * v.RPM * v.Percent) / math.CubicFeetConversion,
	}
}
