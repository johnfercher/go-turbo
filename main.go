package main

import (
	"context"
	"github.com/johnfercher/go-turbo/internal/adapters/csv"
	"github.com/johnfercher/go-turbo/internal/core/models"
	"log"
)

const (
	EngineLiters = 2
)

func main() {
	ctx := context.Background()

	turboRepository := csv.NewTurboRepository()
	_, err := turboRepository.Get(ctx, "kinugawa-td05-18g")
	if err != nil {
		log.Fatal(err)
	}

	/*

		boosts := []float64{
			consts.Boost10,
			consts.Boost12,
		}

		cfmNA := getEJ20CFMna(ctx)

		for _, boost := range boosts {
			var cfmBoosted []*models.CFM
			for _, cfm := range cfmNA {
				cfmBoosted = append(cfmBoosted, cfm.AddBoostKg(boost))
			}

			fmt.Println(boost)
			fmt.Println(cfmBoosted)
		}*/
}

func getEJ20CFMna(ctx context.Context) []*models.CFM {
	veRepo := csv.NewVERepository()
	ve, err := veRepo.Get(ctx, "ej20")
	if err != nil {
		log.Fatal(err)
	}

	var cfm []*models.CFM
	for _, v := range ve {
		cfm = append(cfm, v.ToFourCylinderCFM(EngineLiters))
	}

	return cfm
}
