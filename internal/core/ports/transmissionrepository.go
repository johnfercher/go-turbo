package ports

import (
	"context"
	"github.com/johnfercher/go-turbo/internal/core/models"
)

type TransmissionRepository interface {
	Get(ctx context.Context, name string) (*models.Transmission, error)
}
