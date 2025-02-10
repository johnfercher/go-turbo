package models

import (
	"regexp"
	"strconv"
	"strings"
)

var baseRange = regexp.MustCompile(`\d+\-\d`)

func IsBaseRange(s string) bool {
	return baseRange.MatchString(s)
}

func GetFlowFromBaseRange(s string) float64 {
	flow := strings.Split(s, "-")[0]
	v, _ := strconv.ParseFloat(strings.TrimSpace(flow), 64)
	return v
}

func GetScoreFromBaseRange(s string) float64 {
	flow := strings.Split(s, "-")[1]
	v, _ := strconv.ParseFloat(strings.TrimSpace(flow), 64)
	return v
}

type TurboScore struct {
	Boost  float64
	Health float64
	Weight float64
	CFM    float64
}

func NewTurboScoreFromString(s string) *TurboScore {
	if s == "" {
		return &TurboScore{}
	}

	if IsBaseRange(s) {
		score := GetScoreFromBaseRange(s)
		cfm := GetFlowFromBaseRange(s)

		return &TurboScore{
			Boost:  1.0,
			Health: 1.0,
			Weight: score,
			CFM:    cfm,
		}
	}

	v, _ := strconv.ParseFloat(strings.TrimSpace(s), 64)

	return &TurboScore{
		CFM: v,
	}

}
