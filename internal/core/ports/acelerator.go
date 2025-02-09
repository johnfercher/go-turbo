package ports

import (
	"context"
)

type Accelerator interface {
	Simulate(ctx context.Context, engineModel string, turboModel string, boost float64) error
}
