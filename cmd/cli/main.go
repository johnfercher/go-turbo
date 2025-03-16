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
			Engine: "k20z3-si-2008",
			Turbo:  "kinugawa-td05-18g",
			Boost:  62,
			RevMin: 3000,
			RevMax: 9000,
			Octane:
		},
		/*{
			Engine: "ej20g-wrx-1997",
			Turbo:  "kinugawa-td05-18g",
			Boost:  100,
			RevMin: 3000,
			RevMax: 9000,
		},*/
	}

	err := accelerator.Simulate(ctx, simulations)
	if err != nil {
		log.Fatal(err)
	}
}
