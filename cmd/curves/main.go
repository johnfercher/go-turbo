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

	engine, err := engineRepo.Get(ctx, "ej20g-wrx-1997")
	if err != nil {
		log.Fatal(err)
	}

	boost := 100.0
	minRPM := 2000.0
	maxRPM := 9000.0
	rpmRange := 200.0

	VE(engine, minRPM, maxRPM, rpmRange)
	Power(engine, minRPM, maxRPM, rpmRange, boost)
	Torque(engine, minRPM, maxRPM, rpmRange, boost)
}

func VE(engine *models.Engine, minRPM, maxRPM, rpmRange float64) {
	var chart string
	for rpm := minRPM; rpm < maxRPM; rpm += rpmRange {
		ve := engine.GetVE(rpm)
		chart += fmt.Sprintf("%f,%f\n", rpm, ve)
	}

	os.WriteFile("ve.csv", []byte(chart), os.ModePerm)
}

func Power(engine *models.Engine, minRPM, maxRPM, rpmRange, boost float64) {
	var chart string
	for rpm := minRPM; rpm < maxRPM; rpm += rpmRange {
		cfm := engine.GetCFM(rpm, boost)

		lbsMin := math.CubicFeetToLbsMin(cfm.Flow)

		power := models.NewPower(lbsMin, int(rpm), fuel.Gasoline100(), engine)
		chart += fmt.Sprintf("%f,%f\n", rpm, power.HP)
	}

	os.WriteFile("power.csv", []byte(chart), os.ModePerm)
}

func Torque(engine *models.Engine, minRPM, maxRPM, rpmRange, boost float64) {
	var chart string
	for rpm := minRPM; rpm < maxRPM; rpm += rpmRange {
		cfm := engine.GetCFM(rpm, boost)

		lbsMin := math.CubicFeetToLbsMin(cfm.Flow)

		power := models.NewPower(lbsMin, int(rpm), fuel.Gasoline100(), engine)
		chart += fmt.Sprintf("%f,%f\n", rpm, power.Torque)
	}

	os.WriteFile("torque.csv", []byte(chart), os.ModePerm)
}
