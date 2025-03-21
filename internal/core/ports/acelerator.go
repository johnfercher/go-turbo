package ports

import (
	"context"
	"github.com/johnfercher/go-turbo/internal/core/models"
)

type Accelerator interface {
	Simulate(ctx context.Context, rpmIterator int, file string, cars []*models.Car) error
}
