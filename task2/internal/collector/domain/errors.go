package domain

import "errors"

var (
	ErrRepoNotFound = errors.New("repository not found")
	ErrInternal     = errors.New("internal server error")
)
