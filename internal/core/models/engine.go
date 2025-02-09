package models

import (
	"errors"
	"fmt"
	"github.com/johnfercher/go-turbo/internal/core/consts"
)

type Engine struct {
	Name      string
	Liters    float64
	Cylinders int
	VE        []*VE
	CFM       map[string][]*CFM
}

func NewEngine(name string, cylinders int, liters float64, ve []*VE) (*Engine, error) {
	e := &Engine{
		Name:      name,
		Liters:    liters,
		Cylinders: cylinders,
		VE:        ve,
		CFM:       make(map[string][]*CFM),
	}

	if cylinders != 4 {
		return nil, errors.New("cylinders != 4")
	}

	for _, boost := range consts.Boosts {
		var cfm []*CFM
		for _, v := range ve {
			cfm = append(cfm, v.ToFourCylinderCFM(e.Liters).AddBoostKg(boost))
		}
		e.CFM[KgKey(boost)] = cfm
	}

	return e, nil
}

func (e *Engine) String() string {
	s := fmt.Sprintf("Engine: %s, %.2fL, %dC\n", e.Name, e.Liters, e.Cylinders)

	for key, value := range e.CFM {
		s += fmt.Sprintf("%s Kg %v\n", key, value)
	}

	return s
}

func (e *Engine) GetBoostCFM(boost float64) ([]*CFM, error) {
	cfm, ok := e.CFM[KgKey(boost)]
	if !ok {
		return nil, errors.New("no CFM found for boost")
	}

	return cfm, nil
}
