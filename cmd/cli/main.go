package main

import (
	"context"
	"github.com/johnfercher/go-turbo/internal/adapters/csv"
	"github.com/johnfercher/go-turbo/internal/adapters/pdf"
	"github.com/johnfercher/go-turbo/internal/core/models"
	"github.com/johnfercher/go-turbo/internal/core/services"
	"log"
)

func main() {
	ctx := context.Background()

	engineRepo := csv.NewEngineRepository()
	turboRepo := csv.NewTurboRepository()
	pdfReporter := pdf.NewPdfReporter()

	accelerator := services.NewAccelerator(engineRepo, turboRepo, pdfReporter)

	simulations := []*models.Simulation{
		{
			Engine: "ej20",
			Turbo:  "kinugawa-td05-18g",
			Boost:  80,
			RevMin: 3000,
			RevMax: 7500,
		},
		{
			Engine: "ej20",
			Turbo:  "kinugawa-td05-18g",
			Boost:  100,
			RevMin: 3000,
			RevMax: 7500,
		},
		{
			Engine: "ej20",
			Turbo:  "kinugawa-td05-18g",
			Boost:  120,
			RevMin: 3000,
			RevMax: 7500,
		},
		{
			Engine: "ej20",
			Turbo:  "kinugawa-td05-18g",
			Boost:  139,
			RevMin: 3000,
			RevMax: 7500,
		},
	}

	err := accelerator.Simulate(ctx, simulations)
	if err != nil {
		log.Fatal(err)
	}
}
