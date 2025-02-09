package services

import (
	"context"
	"fmt"
	"github.com/johnfercher/go-turbo/internal/core/ports"
)

type Accelerator struct {
	engineRepo ports.EngineRepository
	turboRepo  ports.TurboRepository
}

func NewAccelerator(engineRepo ports.EngineRepository, turboRepo ports.TurboRepository) *Accelerator {
	return &Accelerator{
		engineRepo: engineRepo,
		turboRepo:  turboRepo,
	}
}

func (a *Accelerator) Simulate(ctx context.Context, engineModel string, turboModel string, boost float64) error {
	engine, err := a.engineRepo.Get(ctx, engineModel)
	if err != nil {
		return err
	}

	_, err = a.turboRepo.Get(ctx, turboModel)
	if err != nil {
		return err
	}

	cfm, err := engine.GetBoostCFM(boost)
	if err != nil {
		return err
	}

	fmt.Println(cfm)

	return err
}
