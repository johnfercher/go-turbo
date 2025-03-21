package csv

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/gocarina/gocsv"
	"github.com/johnfercher/go-turbo/internal/core/models"
	"os"
)

type EngineRepository struct {
}

func NewEngineRepository() *EngineRepository {
	return &EngineRepository{}
}

func (v *EngineRepository) Get(ctx context.Context, engine string) (*models.Engine, error) {
	details, err := v.loadEngineDetails(ctx, engine)
	if err != nil {
		return nil, err
	}

	ve, err := v.loadVE(engine)
	if err != nil {
		return nil, err
	}

	e, err := models.NewEngine(details.Name, details.Cylinders, details.Liters,
		details.EfficiencyRatio, details.CompressionRatio, details.BoostGainRatio,
		ve, details.MinRpm, details.MaxRpm)
	if err != nil {
		return nil, err
	}

	return e, nil
}

func (v *EngineRepository) loadVE(engine string) ([]*models.VE, error) {
	b, err := os.ReadFile("data/engine/" + engine + "/ve.csv")
	if err != nil {
		return nil, err
	}

	var entries []*models.VE
	err = gocsv.Unmarshal(bytes.NewBuffer(b), &entries)
	if err != nil {
		return nil, err
	}

	return entries, nil
}

func (v *EngineRepository) loadEngineDetails(ctx context.Context, engine string) (*EngineDetails, error) {
	b, err := os.ReadFile("data/engine/" + engine + "/details.json")
	if err != nil {
		return nil, err
	}

	var details EngineDetails
	err = json.Unmarshal(b, &details)
	if err != nil {
		return nil, err
	}

	return &details, nil
}
