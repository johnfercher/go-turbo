package main

import (
	"context"
	"github.com/johnfercher/go-turbo/internal/adapters/csv"
	"github.com/johnfercher/go-turbo/internal/adapters/pdf"
	"github.com/johnfercher/go-turbo/internal/core/models"
	"github.com/johnfercher/go-turbo/internal/core/models/fuel"
	"github.com/johnfercher/go-turbo/internal/core/services"
	"log"
)

func main() {
	ctx := context.Background()

	engineRepo := csv.NewEngineRepository()
	turboRepo := csv.NewTurboRepository()
	transmissionRepo := csv.NewTransmissionRepository()
	pdfReporter := pdf.NewPdfReporter()

	accelerator := services.NewSimulator(engineRepo, turboRepo, transmissionRepo, pdfReporter)

	simulations := []*models.Car{
		/*{
			Engine: "k20z3-si-2008",
			RevMin: 3000,
			RevMax: 9000,
			Fuel:   fuel.Gasoline100(),
		},*/
		{
			Engine:       "k20z3-si-2008",
			Transmission: "si-2008",
			Turbo:        "kinugawa-td05-18g",
			RevMin:       1500,
			RevMax:       9000,
			Boost:        150,
			WheelInches:  17,
			TireHeightMM: 45,
			Fuel:         fuel.BREthanol(),
		},
		{
			Engine:       "ej20g-wrx-1997",
			Transmission: "sti-6-TY856WB1AA",
			Turbo:        "kinugawa-td05-18g",
			Boost:        100,
			RevMin:       1500,
			RevMax:       8500,
			WheelInches:  17,
			TireHeightMM: 45,
			Fuel:         fuel.BREthanol(),
		},
		/*{
			Engine: "k20z3-si-2008",
			Turbo:  "kinugawa-td05-18g",
			Boost:  100,
			RevMin: 3000,
			RevMax: 9000,
			Fuel:   fuel.BREthanol(),
		},
		{
			Engine: "ej20g-wrx-1997",
			Turbo:  "kinugawa-td05-18g",
			Boost:  100,
			RevMin: 3000,
			RevMax: 9000,
			Fuel:   fuel.Gasoline100(),
		},
		{
			Engine: "ej20g-wrx-1997",
			Turbo:  "kinugawa-td05-18g",
			Boost:  100,
			RevMin: 3000,
			RevMax: 9000,
			Fuel:   fuel.BREthanol(),
		},*/
	}

	err := accelerator.Simulate(ctx, 1, "si", simulations)
	if err != nil {
		log.Fatal(err)
	}
}
