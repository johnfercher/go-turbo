package main

import (
	"context"
	"fmt"
	"github.com/johnfercher/go-turbo/internal/adapters/csv"
	"log"
)

func main() {
	ctx := context.Background()

	engineRepo := csv.NewEngineRepository()
	turboRepo := csv.NewTurboRepository()

	engine, err := engineRepo.Get(ctx, "ej20")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(engine)

	turbo, err := turboRepo.Get(ctx, "kinugawa-td05-18g")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(turbo)
}
