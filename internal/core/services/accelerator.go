package services

import (
	"context"
	"github.com/johnfercher/go-turbo/internal/core/models"
	"github.com/johnfercher/go-turbo/internal/core/ports"
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

func (a *Accelerator) Simulate(ctx context.Context, file string, simulations []*models.Simulation) error {
	var reports []*models.Report
	for _, simulation := range simulations {
		report, err := a.simulate(ctx, simulation)
		if err != nil {
			return err
		}

		reports = append(reports, report)
	}

	return a.reporter.Generate(ctx, file, reports)
}

func (a *Accelerator) simulate(ctx context.Context, simulation *models.Simulation) (*models.Report, error) {
	engine, err := a.engineRepo.Get(ctx, simulation.Engine)
	if err != nil {
		return nil, err
	}

	var turbo *models.Turbo
	if simulation.Turbo != "" {
		turbo, err = a.turboRepo.Get(ctx, simulation.Turbo)
		if err != nil {
			return nil, err
		}
	}

	report := models.NewReport(engine, turbo, simulation.Fuel, simulation.Boost)

	for i := simulation.RevMin; i <= simulation.RevMax; i++ {
		cfm := engine.Get(float64(i), simulation.Boost)

		var health = 0.0
		if simulation.Boost > 0 {
			health = turbo.Get(cfm.Flow, simulation.Boost)
		}

		report.Add(i, cfm.Flow, health)
	}

	return report, nil
}
