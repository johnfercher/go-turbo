package ports

import (
	"context"
	"github.com/johnfercher/go-turbo/internal/core/models"
)

type RangeCalculator interface {
	Calculate(ctx context.Context, cfm []*models.CFM)
}
