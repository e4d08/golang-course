package http

import (
	"collector/internal/domain"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"
)

type HttpAdapter struct {
	client *http.Client
	apiUrl string
}

func NewHttpAdapter(apiUrl string) *HttpAdapter {
	return &HttpAdapter{
		client: &http.Client{Timeout: 5 * time.Second},
		apiUrl: apiUrl,
	}
}

func (a *HttpAdapter) FetchRepository(ctx context.Context, owner string, name string) (domain.Repository, error) {
	resp, err := a.client.Get(fmt.Sprintf("%s/%s/%s", a.apiUrl, owner, name))
	if err != nil {
		return domain.Repository{}, domain.ErrInternal
	}
	defer resp.Body.Close()

	switch resp.StatusCode {
	case 200:
	case 404:
		return domain.Repository{}, domain.ErrRepositoryNotFound
	default:
		return domain.Repository{}, errors.New(resp.Status)
	}

	var fetchedRepository Repository
	if err := json.NewDecoder(resp.Body).Decode(&fetchedRepository); err != nil {
		return domain.Repository{}, fmt.Errorf("failed to decode json: %w", err)
	}

	return domain.Repository{
		Name:        fetchedRepository.Name,
		Description: fetchedRepository.Description,
		Stars:       fetchedRepository.Stars,
		Forks:       fetchedRepository.Forks,
		CreatedAt:   fetchedRepository.CreatedAt,
		License:     fetchedRepository.License.Name,
	}, nil
}
