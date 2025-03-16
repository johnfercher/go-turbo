package models

import (
	"errors"
	"fmt"
	"gonum.org/v1/gonum/interp"
)

type Engine struct {
	Name             string
	Liters           float64
	Cylinders        int
	CompressionRatio float64
	EfficiencyRatio  float64
	BoostGainRatio   float64
	VE               []*VE
	VEInter          interp.AkimaSpline
}

func NewEngine(name string, cylinders int, liters float64, efficiencyRatio float64, compressionRatio float64, boostGainRatio float64, ve []*VE) (*Engine, error) {
	e := &Engine{
		Name:             name,
		Liters:           liters,
		Cylinders:        cylinders,
		EfficiencyRatio:  efficiencyRatio,
		CompressionRatio: compressionRatio,
		BoostGainRatio:   boostGainRatio,
		VE:               ve,
	}

	interp := interp.AkimaSpline{}

	var xs []float64
	var ys []float64

	for _, v := range ve {
		xs = append(xs, v.RPM)
		ys = append(ys, v.Percent)
	}

	err := interp.Fit(xs, ys)
	if err != nil {
		return nil, err
	}

	e.VEInter = interp

	if cylinders != 4 {
		return nil, errors.New("cylinders != 4")
	}

	return e, nil
}

func (e *Engine) String() string {
	s := fmt.Sprintf("Engine: %s, %.2fL, %dC\n", e.Name, e.Liters, e.Cylinders)
	s += fmt.Sprintf("VE: %v\n", e.VE)

	return s
}

func (e *Engine) Get(RPM float64, boost float64) *CFM {
	percent := e.VEInter.Predict(RPM)
	ve := NewVE(RPM, percent)
	return ve.ToFourCylinderCFM(e.Liters).AddBoostKg(boost)
}
