package services

import (
	"context"
	"github.com/johnfercher/go-turbo/internal/core/models"
	"github.com/johnfercher/go-turbo/internal/core/ports"
	"github.com/johnfercher/go-turbo/internal/math"
)

type Accelerator struct {
	engineRepo ports.EngineRepository
	turboRepo  ports.TurboRepository
	reporter   ports.Reporter
}

func NewAccelerator(engineRepo ports.EngineRepository, turboRepo ports.TurboRepository, reporter ports.Reporter) *Accelerator {
	return &Accelerator{
		engineRepo: engineRepo,
		turboRepo:  turboRepo,
		reporter:   reporter,
	}
}

func (a *Accelerator) Simulate(ctx context.Context, simulations []*models.Simulation) error {
	for _, simulation := range simulations {
		report, err := a.simulate(ctx, simulation)
		if err != nil {
			return err
		}

		err = a.reporter.Generate(ctx, report)
		if err != nil {
			return err
		}
	}

	return nil
}

func (a *Accelerator) simulate(ctx context.Context, simulation *models.Simulation) (*models.Report, error) {
	engine, err := a.engineRepo.Get(ctx, simulation.Engine)
	if err != nil {
		return nil, err
	}

	turbo, err := a.turboRepo.Get(ctx, simulation.Turbo)
	if err != nil {
		return nil, err
	}

	revLimiter := 7500
	revMin := 2000

	report := models.NewReport(engine, turbo, simulation.Boost)

	for i := revMin; i <= revLimiter; i++ {
		cfm := engine.Get(float64(i), simulation.Boost)
		health := turbo.Get(cfm.Flow, simulation.Boost)
		lbs := math.CubicFeetToLbsMin(cfm.Flow)

		report.Add(i, lbs, health)
	}

	return report, nil
}
