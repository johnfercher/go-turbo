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
		{
			Engine:       "ej20g-wrx-1997",
			Transmission: "gt-5-TY752VN4BA",
			Turbo:        "kinugawa-td05-18g",
			Boost:        100,
			WheelInches:  17,
			TireHeightMM: 45,
			Fuel:         fuel.Gasoline100(),
		},
		{
			Engine:       "ej20g-wrx-1997",
			Transmission: "sti-6-TY856WB1AA",
			Turbo:        "kinugawa-td05-18g",
			Boost:        100,
			WheelInches:  17,
			TireHeightMM: 45,
			Fuel:         fuel.Gasoline100(),
		},
		{
			Engine:       "k20z3-si-2008",
			Transmission: "si-2008",
			WheelInches:  17,
			TireHeightMM: 45,
			Fuel:         fuel.Gasoline100(),
		},
	}

	err := accelerator.Simulate(ctx, 1, "si", simulations)
	if err != nil {
		log.Fatal(err)
	}
}
