package fuel

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFuel_GetPotential(t *testing.T) {
	cases := []struct {
		name      string
		fuel      *Fuel
		potential float64
	}{
		{
			name:      "E100",
			fuel:      E100(),
			potential: 79618.46496106786,
		},
		{
			name:      "gasoline 100%",
			fuel:      Gasoline100(),
			potential: 66877.14481811941,
		},
		{
			name:      "E85",
			fuel:      E85(),
			potential: 78185.27755164342,
		},
		{
			name:      "br gasoline",
			fuel:      BRGasoline(),
			potential: 70154.07515889195,
		},
		{
			name:      "br podium gasoline",
			fuel:      BRPodiumGasoline(),
			potential: 76165.88925715974,
		},
		{
			name:      "br ethanol",
			fuel:      BREthanol(),
			potential: 84208.23529411765,
		},
		{
			name:      "br methanol",
			fuel:      BRMethanol(),
			potential: 94737.2488408037,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			assert.Equal(t, c.potential, c.fuel.Power())
		})
	}
}

func TestFuel_Compare(t *testing.T) {
	cases := []struct {
		name      string
		base      *Fuel
		ref       *Fuel
		potential float64
	}{
		{
			name:      "br gasoline x gasoline 100%",
			base:      BRGasoline(),
			ref:       Gasoline100(),
			potential: 1.0416493676162741,
		},
		{
			name:      "br podium gasoline x br gasoline",
			base:      BRPodiumGasoline(),
			ref:       BRGasoline(),
			potential: 1.0728402729556892,
		},
		{
			name:      "br ethanol x br gasoline",
			base:      BREthanol(),
			ref:       BRGasoline(),
			potential: 1.1702828536743628,
		},
		{
			name:      "br ethanol x br podium gasoline",
			base:      BREthanol(),
			ref:       BRPodiumGasoline(),
			potential: 1.0897513860612036,
		},
		{
			name:      "br methanol x br ethanol",
			base:      BRMethanol(),
			ref:       BREthanol(),
			potential: 1.106280121931356,
		},
		{
			name:      "br methanol x br gasoline",
			base:      BRMethanol(),
			ref:       BRGasoline(),
			potential: 1.2978543667249312,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			assert.Equal(t, c.potential, c.base.Compare(c.ref))
		})
	}
}
