package services

import (
	"context"
	"fmt"
	"github.com/johnfercher/go-turbo/internal/core/ports"
)

type Accelerator struct {
	engineRepo  ports.EngineRepository
	turboRepo   ports.TurboRepository
	pdfReporter ports.Reporter
}

func NewAccelerator(engineRepo ports.EngineRepository, turboRepo ports.TurboRepository,
	pdfReporter ports.Reporter) *Accelerator {
	return &Accelerator{
		engineRepo:  engineRepo,
		turboRepo:   turboRepo,
		pdfReporter: pdfReporter,
	}
}

func (a *Accelerator) Simulate(ctx context.Context, engineModel string, turboModel string, boost float64) error {
	engine, err := a.engineRepo.Get(ctx, engineModel)
	if err != nil {
		return err
	}

	fmt.Print(engine)

	turbo, err := a.turboRepo.Get(ctx, turboModel)
	if err != nil {
		return err
	}

	err = a.pdfReporter.Generate(ctx, turbo.TurboScore)
	if err != nil {
		return err
	}

	revLimiter := 7500
	printRev := 500
	for i := 0; i <= revLimiter; i++ {

		if i%printRev == 0 {
			cfm := engine.Get(float64(i), boost)
			health := turbo.Get(cfm.Flow, boost)
			fmt.Printf("RPM: %d, CFM: %.2f, Health: %.2f\n", i, cfm.Flow, health)
		}
	}

	return err
}
