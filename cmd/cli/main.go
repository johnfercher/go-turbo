package main

import (
	"context"
	"github.com/johnfercher/go-turbo/internal/adapters/csv"
	"github.com/johnfercher/go-turbo/internal/core/consts"
	"github.com/johnfercher/go-turbo/internal/core/services"
	"log"
)

func main() {
	ctx := context.Background()

	engineRepo := csv.NewEngineRepository()
	turboRepo := csv.NewTurboRepository()

	accelerator := services.NewAccelerator(engineRepo, turboRepo)

	err := accelerator.Simulate(ctx, "ej20", "kinugawa-td05-18g", consts.Boost12)
	if err != nil {
		log.Fatal(err)
	}
}
