package csv

import (
	"context"
	"encoding/json"
	"github.com/johnfercher/go-turbo/internal/core/models"
	"os"
)

type TransmissionRepository struct {
}

func NewTransmissionRepository() *TransmissionRepository {
	return &TransmissionRepository{}
}

func (v *TransmissionRepository) Get(ctx context.Context, name string) (*models.Transmission, error) {
	b, err := os.ReadFile("data/transmission/" + name + "/ratio.json")
	if err != nil {
		return nil, err
	}

	transmission := &models.Transmission{}
	err = json.Unmarshal(b, transmission)
	if err != nil {
		return nil, err
	}

	return transmission, nil
}
