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
	Surge  bool
	Choke  bool
}

func (t *TurboScore) String() string {
	return fmt.Sprintf("(%.2fB %.2fH %.2fW %.2fC) ", t.Boost, t.Health, t.Weight, t.CFM)
}

func (t *TurboScore) StringBoost() string {
	return fmt.Sprintf("%.2f", t.Boost)
}

func (t *TurboScore) StringHealth() string {
	return fmt.Sprintf("%.2f", t.Health)
}

func (t *TurboScore) StringWeight() string {
	return fmt.Sprintf("%.2f", t.Weight)
}

func (t *TurboScore) StringCFM() string {
	return fmt.Sprintf("%.2f", t.CFM)
}

func (t *TurboScore) StringSurge() string {
	if t.Surge {
		return "S"
	}

	return " "
}

func (t *TurboScore) StringChoke() string {
	if t.Choke {
		return "C"
	}

	return " "
}
