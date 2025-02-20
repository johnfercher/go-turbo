package ports

import (
	"context"
	"github.com/johnfercher/go-turbo/internal/core/models"
)

type TurboRepository interface {
	Get(ctx context.Context, turbo string) (*models.Turbo, error)
	Save(ctx context.Context, name string, chart *models.Chart) error
}
