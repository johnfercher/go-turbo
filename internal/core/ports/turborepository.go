package ports

import "context"

type TurboRepository interface {
	Get(ctx context.Context, turbo string) (string, error)
}
