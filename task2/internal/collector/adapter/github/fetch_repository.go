package github

import (
	"encoding/json"
	"fmt"
	"gh-anal/internal/collector/domain"
	"net/http"
)

func (g HttpAdapter) FetchRepository(owner string, name string) (domain.GithubRepository, error) {
	link := fmt.Sprintf("https://api.github.com/repos/%s/%s", owner, name)

	res, err := g.client.Get(link)
	if err != nil {
		return domain.GithubRepository{}, domain.ErrInternal
	}
	defer res.Body.Close()

	switch res.StatusCode {
	case http.StatusOK:
		var fetched_repo Repository
		json.NewDecoder(res.Body).Decode(&fetched_repo)

		return domain.GithubRepository{
			Name:        fetched_repo.Name,
			Description: fetched_repo.Description,
			Stars:       fetched_repo.Stars,
			Forks:       fetched_repo.Forks,
			CreatedAt:   fetched_repo.CreatedAt,
			License:     fetched_repo.License.Name,
		}, nil
	case http.StatusNotFound:
		return domain.GithubRepository{}, domain.ErrRepoNotFound
	default:
		return domain.GithubRepository{}, domain.ErrInternal
	}
}
