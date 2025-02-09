package ports

import (
	"context"
	"github.com/johnfercher/go-turbo/internal/core/models"
)

type VERepository interface {
	Get(ctx context.Context, engine string) ([]*models.VE, error)
}
