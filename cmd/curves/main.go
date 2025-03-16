package main

import (
	"context"
	"fmt"
	"github.com/johnfercher/go-turbo/internal/adapters/csv"
	"github.com/johnfercher/go-turbo/internal/core/models"
	"github.com/johnfercher/go-turbo/internal/core/models/fuel"
	"github.com/johnfercher/go-turbo/internal/math"
	"log"
	"os"
)

func main() {
	ctx := context.Background()
	engineRepo := csv.NewEngineRepository()

	engine, err := engineRepo.Get(ctx, "k20z3-si-2008")
	if err != nil {
		log.Fatal(err)
	}

	minRPM := 2000.0
	maxRPM := 9000.0
	rpmRange := 200.0

	VE(engine, minRPM, maxRPM, rpmRange)
	Power(engine, minRPM, maxRPM, rpmRange)
	Torque(engine, minRPM, maxRPM, rpmRange)
}

func VE(engine *models.Engine, minRPM, maxRPM, rpmRange float64) {
	var chart string
	for rpm := minRPM; rpm < maxRPM; rpm += rpmRange {
		ve := engine.GetVE(rpm)
		chart += fmt.Sprintf("%f,%f\n", rpm, ve)
	}

	os.WriteFile("ve.csv", []byte(chart), os.ModePerm)
}

func Power(engine *models.Engine, minRPM, maxRPM, rpmRange float64) {
	var chart string
	for rpm := minRPM; rpm < maxRPM; rpm += rpmRange {
		cfm := engine.GetCFM(rpm, 0)

		lbsMin := math.CubicFeetToLbsMin(cfm.Flow)

		power := models.NewPower(lbsMin, int(rpm), fuel.Gasoline100(), engine)
		chart += fmt.Sprintf("%f,%f\n", rpm, power.HP)
	}

	os.WriteFile("power.csv", []byte(chart), os.ModePerm)
}

func Torque(engine *models.Engine, minRPM, maxRPM, rpmRange float64) {
	var chart string
	for rpm := minRPM; rpm < maxRPM; rpm += rpmRange {
		cfm := engine.GetCFM(rpm, 0)

		lbsMin := math.CubicFeetToLbsMin(cfm.Flow)

		power := models.NewPower(lbsMin, int(rpm), fuel.Gasoline100(), engine)
		chart += fmt.Sprintf("%f,%f\n", rpm, power.Torque)
	}

	os.WriteFile("torque.csv", []byte(chart), os.ModePerm)
}
