package domain

import "context"

type CollectorAdapter interface {
	GetRepository(ctx context.Context, owner string, name string) (Repository, error)
}
