package ports

import (
	"context"
	"github.com/johnfercher/go-turbo/internal/core/consts"
	"github.com/johnfercher/go-turbo/internal/core/models"
)

type Reporter interface {
	Generate(ctx context.Context, turbo [][]*models.TurboScore, reportType consts.ReportType) error
}
