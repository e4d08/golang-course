package github

import (
	"net/http"
	"time"
)

type Repository struct {
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Stars       uint32    `json:"stargazers_count"`
	Forks       uint32    `json:"forks_count"`
	CreatedAt   time.Time `json:"created_at"`
	License     License   `json:"license"`
}

type License struct {
	Name string `json:"name"`
}

type HttpAdapter struct {
	client *http.Client
}

func NewHttpAdapter(timeout time.Duration) HttpAdapter {
	client := http.Client{Timeout: timeout}

	return HttpAdapter{
		client: &client,
	}
}
