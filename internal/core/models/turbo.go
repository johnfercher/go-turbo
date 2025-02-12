package models

import (
	"fmt"
	"github.com/johnfercher/go-turbo/internal/core/consts"
	"gonum.org/v1/gonum/interp"
)

type Turbo struct {
	Name        string
	TurboScore  map[string][]*TurboScore
	BoostInter  map[string]interp.AkimaSpline
	HealthInter map[string]interp.AkimaSpline
	WeightInter map[string]interp.AkimaSpline
}

func NewTurbo(name string, turboMatrix [][]*TurboScore) (*Turbo, error) {
	t := &Turbo{
		Name:        name,
		TurboScore:  make(map[string][]*TurboScore),
		BoostInter:  make(map[string]interp.AkimaSpline),
		HealthInter: make(map[string]interp.AkimaSpline),
		WeightInter: make(map[string]interp.AkimaSpline),
	}

	for i, turboLine := range turboMatrix {
		if isAllSurgeChoke(turboLine) {
			continue
		}

		boost := consts.Boosts[i]

		t.TurboScore[KgKey(boost)] = turboLine

		// Boost
		boostInter, err := buildBoostInter(turboLine)
		if err != nil {
			return nil, err
		}
		t.BoostInter[KgKey(boost)] = boostInter

		// Health
		healthInter, err := buildHealthInter(turboLine)
		if err != nil {
			return nil, err
		}
		t.HealthInter[KgKey(boost)] = healthInter

		// Weight
		weightInter, err := buildWeightInter(turboLine)
		if err != nil {
			return nil, err
		}
		t.WeightInter[KgKey(boost)] = weightInter
	}

	return t, nil
}

func (t *Turbo) String() string {
	return fmt.Sprintf("Turbo: %s\n", t.Name)
}

func (t *Turbo) Get(cfm float64, configuredBoost float64) (surge bool, choke bool, boost float64, health float64) {
	turboSlice := t.TurboScore[KgKey(configuredBoost)]

	boostInter := t.BoostInter[KgKey(configuredBoost)]
	healthInter := t.HealthInter[KgKey(configuredBoost)]

	for _, turbo := range turboSlice {
		if cfm <= turbo.CFM {
			surge = turbo.Surge
			break
		}
	}

	boost = boostInter.Predict(cfm)
	if boost < 0 {
		boost = 0
	}
	if boost > configuredBoost {
		boost = configuredBoost
	}
	health = healthInter.Predict(cfm)
	if health < 0 {
		health = 0
	}
	if health > 1.0 {
		health = 1.0
	}

	return
}

func buildBoostInter(turboScore []*TurboScore) (interp.AkimaSpline, error) {
	inter := interp.AkimaSpline{}
	xs := []float64{}
	ys := []float64{}

	for _, turboS := range turboScore {
		xs = append(xs, turboS.CFM)
		ys = append(ys, turboS.Boost)
	}

	xs, ys = filterXY(xs, ys)

	return inter, inter.Fit(xs, ys)
}

func buildHealthInter(turboScore []*TurboScore) (interp.AkimaSpline, error) {
	inter := interp.AkimaSpline{}
	xs := []float64{}
	ys := []float64{}

	for _, turboS := range turboScore {
		xs = append(xs, turboS.CFM)
		ys = append(ys, turboS.Health)
	}

	xs, ys = filterXY(xs, ys)

	return inter, inter.Fit(xs, ys)
}

func buildWeightInter(turboScore []*TurboScore) (interp.AkimaSpline, error) {
	inter := interp.AkimaSpline{}
	xs := []float64{}
	ys := []float64{}

	for _, turboS := range turboScore {
		xs = append(xs, turboS.CFM)
		ys = append(ys, turboS.Weight)
	}

	xs, ys = filterXY(xs, ys)

	return inter, inter.Fit(xs, ys)
}

func filterXY(xs []float64, ys []float64) ([]float64, []float64) {
	lastChangeIndex := 0
	lastValue := -1.0
	for index, value := range xs {
		if value != lastValue {
			lastValue = value
			lastChangeIndex = index
		}
	}

	return xs[:lastChangeIndex+1], ys[:lastChangeIndex+1]
}

func isAllSurgeChoke(turboScore []*TurboScore) bool {
	for _, t := range turboScore {
		if !t.Surge && !t.Choke {
			return false
		}
	}

	return true
}
