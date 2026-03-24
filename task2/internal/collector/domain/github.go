package domain

import "time"

type GithubRepository struct {
	Name        string
	Description string
	Stars       uint32
	Forks       uint32
	CreatedAt   time.Time
	License     string
}
