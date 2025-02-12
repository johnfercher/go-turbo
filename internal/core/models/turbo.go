package models

import (
	"github.com/johnfercher/go-turbo/internal/core/consts"
	"gonum.org/v1/gonum/interp"
)

type Turbo struct {
	Name        string
	TurboScore  map[string][]*TurboScore
	BoostInter  map[string]interp.AkimaSpline
	HealthInter map[string]interp.AkimaSpline
	WeightInter map[string]interp.AkimaSpline
	CFMInter    map[string]interp.AkimaSpline
}

func NewTurbo(name string, turboScore [][]*TurboScore) *Turbo {
	t := &Turbo{
		Name:       name,
		TurboScore: make(map[string][]*TurboScore),
	}

	for i, v := range turboScore {
		boost := consts.Boosts[i]
		t.TurboScore[KgKey(boost)] = v
	}

	return t
}
