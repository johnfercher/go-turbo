package ports

import (
	"context"
	"github.com/johnfercher/go-turbo/internal/core/models"
)

type Reporter interface {
	Generate(ctx context.Context, file string, reports []*models.Report) error
}
