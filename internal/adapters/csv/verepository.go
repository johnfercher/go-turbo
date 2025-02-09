package csv

import (
	"bytes"
	"context"
	"github.com/gocarina/gocsv"
	"github.com/johnfercher/go-turbo/internal/core/models"
	"os"
)

type VERepository struct {
}

func NewVERepository() *VERepository {
	return &VERepository{}
}

func (v *VERepository) Get(ctx context.Context, engine string) ([]*models.VE, error) {
	b, err := os.ReadFile("data/ve/" + engine + ".csv")
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
