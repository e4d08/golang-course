package usecase

import (
	"gh-anal/internal/collector/domain"
)

type GithubAdapter interface {
	FetchRepository(owner string, name string) (domain.GithubRepository, error)
}

type CollectorService struct {
	github GithubAdapter
}

func NewCollectorService(github GithubAdapter) CollectorService {
	return CollectorService{
		github: github,
	}
}

func (c CollectorService) GetRepository(name string, owner string) (domain.GithubRepository, error) {
	result, err := c.github.FetchRepository(name, owner)
	if err != nil {
		return result, err
	}

	return result, nil
}
