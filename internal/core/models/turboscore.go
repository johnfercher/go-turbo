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

func PrintBoost(turboScore [][]*TurboScore) {
	fmt.Println("Boost:")
	for i := 0; i < len(turboScore); i++ {
		for j := 0; j < len(turboScore[i]); j++ {
			fmt.Print(turboScore[i][j].StringBoost(), " ")
		}
		fmt.Println()
	}
}

func PrintWeight(turboScore [][]*TurboScore) {
	fmt.Println("Weight:")
	for i := 0; i < len(turboScore); i++ {
		for j := 0; j < len(turboScore[i]); j++ {
			fmt.Print(turboScore[i][j].StringWeight(), " ")
		}
		fmt.Println()
	}
}

func PrintCFM(turboScore [][]*TurboScore) {
	fmt.Println("CFM:")
	for i := 0; i < len(turboScore); i++ {
		for j := 0; j < len(turboScore[i]); j++ {
			fmt.Print(turboScore[i][j].StringCFM(), " ")
		}
		fmt.Println()
	}
}

func PrintHealth(turboScore [][]*TurboScore) {
	fmt.Println("Health:")
	for i := 0; i < len(turboScore); i++ {
		for j := 0; j < len(turboScore[i]); j++ {
			fmt.Print(turboScore[i][j].StringHealth(), " ")
		}
		fmt.Println()
	}
}

func PrintSurge(turboScore [][]*TurboScore) {
	fmt.Println("Surge:")
	for i := 0; i < len(turboScore); i++ {
		for j := 0; j < len(turboScore[i]); j++ {
			fmt.Print(turboScore[i][j].StringSurge(), " ")
		}
		fmt.Println()
	}
}

func PrintChoke(turboScore [][]*TurboScore) {
	fmt.Println("Choke:")
	for i := 0; i < len(turboScore); i++ {
		for j := 0; j < len(turboScore[i]); j++ {
			fmt.Print(turboScore[i][j].StringChoke(), " ")
		}
		fmt.Println()
	}
}
