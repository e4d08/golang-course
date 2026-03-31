package domain

import "context"

type GithubAdapter interface {
	FetchRepository(ctx context.Context, owner string, name string) (Repository, error)
}
