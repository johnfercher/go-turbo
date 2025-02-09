package ports

import (
	"context"
	"github.com/johnfercher/go-turbo/internal/core/models"
)

type EngineRepository interface {
	Get(ctx context.Context, engine string) (*models.Engine, error)
}
