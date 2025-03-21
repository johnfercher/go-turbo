package services

import (
	"context"
	"github.com/johnfercher/go-turbo/internal/core/models"
	"github.com/johnfercher/go-turbo/internal/core/ports"
)

type Simulator struct {
	engineRepo       ports.EngineRepository
	turboRepo        ports.TurboRepository
	transmissionRepo ports.TransmissionRepository
	reporter         ports.Reporter
}

func NewSimulator(engineRepo ports.EngineRepository, turboRepo ports.TurboRepository, transmissionRepo ports.TransmissionRepository,
	reporter ports.Reporter) *Simulator {
	return &Simulator{
		engineRepo:       engineRepo,
		turboRepo:        turboRepo,
		transmissionRepo: transmissionRepo,
		reporter:         reporter,
	}
}

func (a *Simulator) Simulate(ctx context.Context, rpmIterator float64, file string, cars []*models.Car) error {
	var reports []*models.Report
	for _, simulation := range cars {
		report, err := a.simulate(ctx, rpmIterator, simulation)
		if err != nil {
			return err
		}

		reports = append(reports, report)
	}

	return a.reporter.Generate(ctx, file, reports)
}

func (a *Simulator) simulate(ctx context.Context, rpmIterator float64, car *models.Car) (*models.Report, error) {
	engine, err := a.engineRepo.Get(ctx, car.Engine)
	if err != nil {
		return nil, err
	}

	transmission, err := a.transmissionRepo.Get(ctx, car.Transmission)
	if err != nil {
		return nil, err
	}

	var turbo *models.Turbo
	if car.Turbo != "" {
		turbo, err = a.turboRepo.Get(ctx, car.Turbo)
		if err != nil {
			return nil, err
		}
	}

	report := models.NewReport(car.RevMin, car.RevMax, engine, turbo, transmission, car.Fuel, car.Boost)

	for gear := 0; gear < transmission.MaxGear(); gear++ {
		for i := car.RevMin; i <= car.RevMax; i += rpmIterator {
			cfm := engine.GetCFM(float64(i), car.Boost)

			var health = 0.0
			if car.Boost > 0 {
				health = turbo.Get(cfm.Flow, car.Boost)
			}

			report.Add(i, cfm.Flow, health, car.WheelInches, car.TireHeightMM, gear)
		}
		report.NextGear()
	}

	return report, nil
}
