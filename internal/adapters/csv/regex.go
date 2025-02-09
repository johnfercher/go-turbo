package csv

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
