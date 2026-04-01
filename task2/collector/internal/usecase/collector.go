package usecase

import (
	"collector/internal/domain"
	"context"
)

type CollectorService struct {
	github domain.GithubAdapter
}

func NewCollectorService(github domain.GithubAdapter) *CollectorService {
	return &CollectorService{github: github}
}

func (s *CollectorService) GetRepository(ctx context.Context, owner string, name string) (domain.Repository, error) {
	repo, err := s.github.FetchRepository(ctx, owner, name)
	if err != nil {
		return domain.Repository{}, err
	}

	return repo, nil
}
