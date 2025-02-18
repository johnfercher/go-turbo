package ports

import (
	"context"
	"github.com/johnfercher/go-turbo/internal/core/models"
)

type Reporter interface {
	Generate(ctx context.Context, reports []*models.Report) error
}
