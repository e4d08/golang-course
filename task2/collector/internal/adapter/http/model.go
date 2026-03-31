package http

import "time"

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
