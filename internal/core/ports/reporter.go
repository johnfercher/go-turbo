package ports

import (
	"context"
)

type Reporter interface {
	Generate(ctx context.Context, turbo [][]float64) error
}
