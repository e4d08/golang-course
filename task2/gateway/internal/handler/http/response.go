package http

import (
	"time"
)

type GetRepositoryResponse struct {
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Stars       uint32    `json:"stargazers_count"`
	Forks       uint32    `json:"forks_count"`
	CreatedAt   time.Time `json:"created_at"`
	License     string    `json:"license"`
}

type ErrorResponse struct {
	Error string `json:"error"`
}
