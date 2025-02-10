package models

import (
	"fmt"
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
	Out    bool
}

func (t *TurboScore) String() string {
	return fmt.Sprintf("(%.2fB %.2fH %.2fW %.2fC) ", t.Boost, t.Health, t.Weight, t.CFM)
}
