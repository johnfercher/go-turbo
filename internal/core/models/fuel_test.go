package models

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFuel_PowerDiff(t *testing.T) {
	cases := []struct {
		Name string
		Base *Fuel
		Ref  *Fuel
		Diff float64
	}{
		{
			Name: "BR Podium Gasoline to BR Gasoline",
			Base: BRPodiumGasoline(),
			Ref:  BRGasoline(),
			Diff: 1.1693514366757425,
		},
		{
			Name: "BR Alcohol to BR Gasoline",
			Base: BRAlcohol(),
			Ref:  BRGasoline(),
			Diff: 1.1693514366757425,
		},
		{
			Name: "BR Methanol to BR Gasoline",
			Base: BRMethanol(),
			Ref:  BRGasoline(),
			Diff: 1.1446710957243948,
		},
		{
			Name: "BR Alcohol to BR Podium Gasoline",
			Base: BRAlcohol(),
			Ref:  BRPodiumGasoline(),
			Diff: 1.0490583855553282,
		},
	}

	for _, c := range cases {
		t.Run(c.Name, func(t *testing.T) {
			diff := c.Base.PowerDiffTo(c.Ref)
			assert.Equal(t, c.Diff, diff)
		})
	}
}
